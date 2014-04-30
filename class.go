package main

import (
	"bytes"
	"regexp"
)

func (self *Generator) GenClassTypes(ns *Namespace) {
	output := new(bytes.Buffer)
	w(output, "package %s\n\n", self.PackageName)
	w(output, "/*\n")
	for _, include := range self.Includes {
		w(output, "#include <%s>\n", include)
	}
	w(output, "*/\n")
	w(output, "import \"C\"\n")
	w(output, "import \"unsafe\"\n")
	w(output, `func init() {
		_ = unsafe.Pointer(nil)
	}
	`)

	// generate class map
	classes := make(map[string]*Class)
	for _, c := range ns.Classes {
		classes[c.Name] = c
	}

	for _, c := range ns.Classes {
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
		for currentClass != nil {
			w(output, "*_Trait%s\n", currentClass.Name)
			constructExpr += fs("&_Trait%s{(*%s)(p)},\n",
				currentClass.Name, cTypeToGoType(currentClass.CType))
			parent := currentClass.Parent
			currentClass = classes[parent]
			if parent != "" && currentClass == nil {
				//p("==fixme== external class %s\n", parent) FIXME
			}
		}
		w(output, "CPointer unsafe.Pointer\n")
		w(output, "}\n")

		// wrapper
		w(output, `func New%sFromCPointer(p unsafe.Pointer) *%s {
			return &%s{
				%sp,
			}
		}
		`, typeName, typeName, typeName, constructExpr)

		// type mapping
		TypeMapping["*"+goType] = "*" + typeName
		InParamMapping[fs("*%s -> *%s", typeName, goType)] = func(param *Param) {
			param.CgoParam = fs("(*%s)(%s.CPointer)", goType, param.GoName)
		}
		OutParamMapping[fs("*%s -> *%s", goType, typeName)] = func(param *Param) {
			param.CgoAfterStmt += fs("%s = New%sFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__%s).Pointer()))",
				param.GoName, typeName, param.GoName)
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
	w(output, `func init() {
		_ = unsafe.Pointer(nil)
		_ = reflect.ValueOf(nil)
		_ = errors.New("")
	}
	`)

	for _, c := range ns.Classes {
		for _, ctor := range c.Constructors {
			ctor.Name = c.Name + "_" + ctor.Name
			ctor.Return.Type.CType = c.CType + "*"
			ctor.IsConstructor = true
			self.GenFunction(ctor, output, nil)
		}
	}

	self.formatAndOutput("class_constructors", output.Bytes())
}
