package main

import (
	"fmt"
	"strings"
)

var typeStat = make(map[string][]string)

func (self *Generator) GenFunction(fn *Function) {
	// convert
	fn.Convert()

	// skip varargs functions
	if fn.IsVarargs {
		w(self.funcsOutput, "// %s is not generated due to varargs\n\n", fn.CIdentifier)
		return
	}

	// doc
	if fn.Doc != nil {
		w(self.funcsOutput, "/*\n%s\n*/\n", fn.Doc.Text)
	}

	// receiver FIXME

	// function name
	w(self.funcsOutput, "func %s", fn.GoName)

	// collect params
	var inParams, outParams []*Param
	for _, param := range fn.Params {
		param.Convert()
		if param.Direction == "out" {
			outParams = append(outParams, param)
			if param.GoType[0] == '*' {
				param.GoType = param.GoType[1:]
			}
		} else { // inout param is treated as in param
			inParams = append(inParams, param)
			param.Direction = "in"
		}
	}
	fn.Return.Convert() // return value
	if fn.Return.GoType != "C.void" {
		fn.Return.GoName = "__return__"
		outParams = append(outParams, fn.Return) // add return value to out params
	}
	if fn.Throws == "1" { // error value
		errParam := &Param{
			GoName: "__err__",
			GoType: "error",
		}
		outParams = append(outParams, errParam)
	}

	// function signature
	w(self.funcsOutput, "(")
	for _, param := range inParams { // in
		if param.MappedType != "" {
			w(self.funcsOutput, "%s %s, ", param.GoName, param.MappedType)
		} else {
			typeSpec := fmt.Sprintf("%s %s %s %s %s %s", param.CType, param.CTypeName, param.GoType,
				param.ElementCType, param.ElementCTypeName, param.ElementGoType)
			typeStat[typeSpec] = append(typeStat[typeSpec], param.Name+" @ "+fn.Name)
			w(self.funcsOutput, "%s %s, ", param.GoName, param.GoType)
		}
	}
	w(self.funcsOutput, ") (")
	for _, param := range outParams { // out
		if param.MappedType != "" {
			w(self.funcsOutput, "%s %s, ", param.GoName, param.MappedType)
		} else {
			w(self.funcsOutput, "%s %s, ", param.GoName, param.GoType)
		}
	}
	w(self.funcsOutput, ")")

	// body
	w(self.funcsOutput, "{\n") // body start
	//p("%s\n", fn.CIdentifier) //FIXME
	w(self.funcsOutput, "return\n}\n") // body end

	// blank line
	w(self.funcsOutput, "\n")
}

func (self *Function) Convert() {
	// varargs
	for _, param := range self.Params {
		if param.Varargs != nil || param.Type != nil && param.Type.Name == "va_list" {
			self.IsVarargs = true
			break
		}
	}
	// name
	parts := strings.Split(self.Name, "_")
	for i, p := range parts {
		parts[i] = strings.Title(p)
	}
	self.GoName = strings.Join(parts, "")
}

func (self *Param) Convert() {
	// name
	self.GoName = self.Name
	if isGoReservedWord(self.GoName) {
		self.GoName += "_"
	}
	// type info
	if self.Array != nil { // array type
		self.IsArray = true
		self.CType = self.Array.CType
		self.CTypeName = self.Array.Name
		self.ElementCType = self.Array.Type.CType
		self.ElementCTypeName = self.Array.Type.Name
		self.ElementGoType = cTypeToGoType(self.ElementCType)
	} else if self.Type != nil { // non-array type
		self.CType = self.Type.CType
		self.CTypeName = self.Type.Name
	}
	self.GoType = cTypeToGoType(self.CType)
	// type mapping
	self.MapType()
}
