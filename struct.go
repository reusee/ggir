package main

import (
	"bytes"
	"regexp"
	"strings"
)

func (self *Generator) GenStructs(ns *Namespace) {
	/*
		for _, s := range ns.Records {
			p("%s\n", s.Name)
			p("%s\n", s.CType)
			p("%v\n", s.Constructors)
			p("%v\n", s.Methods)
			p("%v\n", s.Functions)
			p("<<<\n")
		}
	*/
}

func (self *Generator) GenStructTypes(ns *Namespace) {
	output := new(bytes.Buffer)
	w(output, "package %s\n\n", self.PackageName)
	w(output, "/*\n")
	for _, include := range self.Includes {
		w(output, "#include <%s>\n", include)
	}
	w(output, "*/\n")
	w(output, "import \"C\"\n")

	for _, s := range ns.Records {
		goType := cTypeToGoType(s.CType)
		typeName := s.Name

		// skip
		for _, t := range self.TypesIgnorePatterns {
			matched, err := regexp.MatchString(t, typeName)
			checkError(err)
			if matched {
				goto next
			}
		}
		if s.IsGTypeStruct != "" {
			goto next
		}
		if strings.HasSuffix(s.Name, "Private") {
			goto next
		}

		// type mapping
		TypeMapping["*"+goType] = "*" + typeName
		InParamMapping[fs("*%s -> *%s", typeName, goType)] = func(param *Param) {
			param.CgoParam = fs("(*%s)(unsafe.Pointer(%s))", goType, param.GoName)
		}
		OutParamMapping[fs("*%s -> *%s", goType, typeName)] = func(param *Param) {
			param.CgoAfterStmt += fs("%s = (*%s)(unsafe.Pointer(__cgo__%s))", param.GoName, typeName, param.GoName)
		}

		// type declaration
		w(output, "type %s %s\n", typeName, goType)

	next:
	}

	self.formatAndOutput("struct_types", output.Bytes())
}
