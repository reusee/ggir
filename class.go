package main

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
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

		w(output, "type %s struct {\n", c.Name)
		currentClass := c
		for currentClass != nil {
			w(output, "*_Trait%s\n", currentClass.Name)
			parent := currentClass.Parent
			currentClass = classes[parent]
			if parent != "" && currentClass == nil {
				p("==fixme== external class %s\n", parent)
			}
		}
		w(output, "CPointer unsafe.Pointer\n")
		w(output, "}\n\n")

		// type mapping
		TypeMapping["*"+goType] = "*" + typeName
		InParamMapping[fs("*%s -> *%s", typeName, goType)] = func(param *Param) {
			param.CgoParam = fs("(*%s)(%s.CPointer)", goType, param.GoName)
		}
		OutParamMapping[fs("*%s -> *%s", goType, typeName)] = func(param *Param) {
			param.CgoAfterStmt += fs("_ = __cgo__%s", param.GoName) //FIXME
		}
	}

	f, err := os.Create(filepath.Join(self.outputDir, self.PackageName+"_class_types.go"))
	checkError(err)
	formatted, err := format.Source(output.Bytes())
	if err != nil {
		f.Write(output.Bytes())
		f.Close()
		checkError(err)
	}
	f.Write(formatted)
	f.Close()
}

func (self *Generator) GenClasses(ns *Namespace) {
	/*
		for _, c := range ns.Classes {
			p("%s\n", c.Name)
			p("%s\n", c.CSymbolPrefix)
			p("%s\n", c.CType)
			p("%s\n", c.Parent)
			p("%s\n", c.Abstract)
			p("%v\n", c.Prerequisite)
			p("%v\n", c.Implements)
			p("%v\n", c.Constructors)
			p("%v\n", c.VirtualMethods)
			p("%v\n", c.Methods)
			p("%v\n", c.Functions)
			p("%v\n", c.Properties)
			p("%v\n", c.Fields)
			p("%v\n", c.Signals)
			p("<<<\n")
		}
	*/
}
