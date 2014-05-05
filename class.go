package main

import (
	"bytes"
	"regexp"
	"strings"
)

func (self *Generator) GenClassTypes(ns *Namespace) {
	output := new(bytes.Buffer)
	w(output, "package %s\n\n", self.PackageName)
	w(output, "/*\n")
	for _, include := range self.Includes {
		w(output, "#include <%s>\n", include)
	}
	w(output, "#include <glib-object.h>\n")
	w(output, "#cgo pkg-config: gobject-2.0\n")
	w(output, "*/\n")
	w(output, "import \"C\"\n")
	w(output, "import \"unsafe\"\n")
	for _, ext := range self.ExternalPackages { // external import
		w(output, "import \"%s\"\n", ext.Import)
	}
	w(output, "import \"runtime\"\n")
	w(output, `func init() {
		_ = unsafe.Pointer(nil)
		_ = runtime.Compiler
	}
	`)
	for _, ext := range self.ExternalPackages {
		w(output, "var _ = %s.UnusedFix_\n", ext.Name)
	}

	// generate class map
	classes := make(map[string]*Class)
	for _, c := range ns.Classes {
		classes[c.Name] = c
	}
	for _, ext := range self.ExternalPackages {
		for _, c := range ext.Repo.Namespace.Classes {
			classes[ext.Repo.Namespace.Name+"."+c.Name] = c
			c.Namespace = ext.Repo.Namespace.Name
		}
	}

	for _, c := range ns.Classes {
		if c.CType == "" {
			continue
		}
		goType := cTypeToGoType(c.CType)
		typeName := c.Name

		// skip
		skip := false
		for _, t := range self.TypesIgnorePatterns {
			matched, err := regexp.MatchString(t, typeName)
			checkError(err)
			if matched {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		//FIXME interfaces
		//p("%v\n", c.Prerequisite)
		//p("%v\n", c.Implements)

		// type declaration
		w(output, "type %s struct {\n", typeName)
		currentClass := c
		constructExpr := ""
		var pkg string
		for currentClass != nil {
			if currentClass.Namespace != "" { // external class
				pkg = strings.ToLower(currentClass.Namespace) + "."
			} else {
				pkg = ""
			}
			w(output, "*%sTrait%s\n", pkg, currentClass.Name)
			constructExpr += fs("%sNewTrait%s(p),\n", pkg, currentClass.Name)
			parent := currentClass.Parent
			if currentClass.Namespace != "" && currentClass.Parent != "" { // external class
				parent = currentClass.Namespace + "." + currentClass.Parent
			}
			currentClass = classes[parent]
			if parent != "" && currentClass == nil {
				p("FIXME external class -> %s\n", parent)
			}
		}
		w(output, "CPointer unsafe.Pointer\n")
		w(output, "}\n")

		// wrapper
		w(output, `func New%sFromCPointer(p unsafe.Pointer) *%s {
			ret := &%s{
				%sp,
			}
			C.g_object_ref_sink(C.gpointer(p))
			runtime.SetFinalizer(ret, func(p *%s) {
				C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
			})
			return ret
		}
		`, typeName, typeName, typeName, constructExpr, typeName)

		// type mapping
		// typed pointer
		TypeMapping[fs("in GoType *%s TypeName %s", goType, typeName)] = "Is" + typeName
		TypeMapping[fs("out GoType *%s TypeName %s", goType, typeName)] = "*" + typeName
		InParamMapping[fs("Is%s -> *%s", typeName, goType)] = func(param *Param) {
			param.CgoParam = fs("%s.Get%sPointer()", param.GoName, typeName)
		}
		OutParamMapping[fs("*%s -> *%s", goType, typeName)] = func(param *Param) {
			param.CgoAfterStmt += fs(`if __cgo__%s != nil {
				%s = New%sFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__%s).Pointer()))
			}`, param.GoName, param.GoName, typeName, param.GoName)
		}
		// untyped pointer
		TypeMapping[fs("in GoType C.gpointer TypeName %s", typeName)] = "Is" + typeName
		TypeMapping[fs("out GoType C.gpointer TypeName %s", typeName)] = "*" + typeName
		InParamMapping[fs("Is%s -> C.gpointer", typeName)] = func(param *Param) {
			param.CgoParam = fs("C.gpointer(unsafe.Pointer(%s.Get%sPointer()))", param.GoName, typeName)
		}
		OutParamMapping[fs("C.gpointer -> *%s", typeName)] = func(param *Param) {
			param.CgoAfterStmt += fs(`%s = New%sFromCPointer(unsafe.Pointer(__cgo__%s))`,
				param.GoName, typeName, param.GoName)
		}
	}

	// register external types mapping
	for _, ext := range self.ExternalPackages {
		ns := ext.Repo.Namespace
		for _, c := range ns.Classes {
			typeName := c.Name
			namespace := ns.Name
			packageName := strings.ToLower(namespace)
			// untyped pointer
			TypeMapping[fs("in GoType C.gpointer TypeName %s.%s", namespace, typeName)] = fs("%s.Is%s",
				packageName, typeName)
			TypeMapping[fs("out GoType C.gpointer TypeName %s.%s", namespace, typeName)] = fs("*%s.%s",
				packageName, typeName)
			InParamMapping[fs("%s.Is%s -> C.gpointer", packageName, typeName)] = func(param *Param) {
				param.CgoParam = fs("C.gpointer(unsafe.Pointer(reflect.ValueOf(%s.Get%sPointer()).Pointer()))",
					param.GoName, typeName)
			}
			OutParamMapping[fs("C.gpointer -> *%s.%s", packageName, typeName)] = func(param *Param) {
				param.CgoAfterStmt += fs("%s = %s.New%sFromCPointer(unsafe.Pointer(__cgo__%s))",
					param.GoName, packageName, typeName, param.GoName)
			}
		}
	}

	self.formatAndOutput("class_types", output.Bytes())
}

func (self *Generator) GenClassConstructors(ns *Namespace) {
	output := new(bytes.Buffer)
	w(output, "package %s\n\n", self.PackageName)
	w(output, "/*\n")
	for _, include := range self.Includes {
		w(output, "#include <%s>\n", include)
	}
	w(output, "*/\n")
	w(output, "import \"C\"\n")
	w(output, "import \"unsafe\"\n")
	w(output, "import \"reflect\"\n")
	w(output, "import \"errors\"\n")
	for _, ext := range self.ExternalPackages { // external import
		w(output, "import \"%s\"\n", ext.Import)
	}
	w(output, `func init() {
		_ = unsafe.Pointer(nil)
		_ = reflect.ValueOf(nil)
		_ = errors.New("")
	}
	var UnusedFix_ bool
	`)
	for _, ext := range self.ExternalPackages {
		w(output, "var _ = %s.UnusedFix_\n", ext.Name)
	}

	for _, c := range ns.Classes {
		for _, ctor := range c.Constructors {
			ctor.Name = c.Name + "_" + ctor.Name
			ctor.Return.Type.CType = c.CType + "*"
			ctor.Return.Type.Name = c.Name
			ctor.IsConstructor = true
			self.GenFunction(ctor, output, nil)
		}
	}

	self.formatAndOutput("class_constructors", output.Bytes())
}
