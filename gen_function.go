package main

import "strings"

var typeStat = make(map[string][]string)

func (self *Generator) GenFunction(fn *Function) {
	// collect info
	fn.CollectInfo()

	// skip varargs functions
	if fn.IsVarargs {
		w(self.funcsOutput, "// %s is not generated due to varargs\n\n", fn.CIdentifier)
		return
	}

	// skip function with inout param due to broken rule FIXME
	//base64_decode_inplace的两个inout参数，text不是pointer，out_len是pointer，规则不一
	for _, param := range fn.Params {
		if param.Direction == "inout" {
			w(self.funcsOutput, "// %s is not generated due to inout params\n\n", fn.CIdentifier)
			return
		}
	}

	// collect params and return info
	for _, param := range fn.Params {
		param.CollectInfo()
		if !param.IsVoid && param.MappedType == "" { // add rules for these specs
			typeStat[param.TypeSpec] = append(typeStat[param.TypeSpec], param.Name+" @ "+fn.Name)
		}
	}
	fn.Return.CollectInfo()
	if !fn.Return.IsVoid {
		if fn.Return.MappedType == "" { // add rules for return type
			typeStat[fn.Return.TypeSpec] = append(typeStat[fn.Return.TypeSpec], fn.Return.Name+" @ "+fn.Name)
		}
		fn.Return.GoName = "__return__"
		fn.Return.Direction = "out"
	}

	// generate doc
	if fn.Doc != nil {
		w(self.funcsOutput, "/*\n%s\n*/\n", fn.Doc.Text)
	}

	// generate signature
	w(self.funcsOutput, "func ")

	// generate receiver FIXME

	// generate function name
	w(self.funcsOutput, fn.GoName)

	// generate parameters
	w(self.funcsOutput, "(")
	for _, param := range fn.Params {
		if !param.IsVoid && param.Direction == "in" {
			w(self.funcsOutput, "%s %s, ", param.GoName, param.MappedType)
		}
	}
	w(self.funcsOutput, ")")

	// generate returns
	w(self.funcsOutput, "(")
	for _, param := range fn.Params {
		if !param.IsVoid && param.Direction == "out" {
			w(self.funcsOutput, "%s %s, ", param.GoName, param.MappedType)
		}
	}
	if !fn.Return.IsVoid { // return
		w(self.funcsOutput, "%s %s, ", fn.Return.GoName, fn.Return.MappedType)
	}
	if fn.Throws != "" { // error
		w(self.funcsOutput, "__err__ error")
	}
	w(self.funcsOutput, ")")

	// generate body
	w(self.funcsOutput, "{\n") // body start
	//FIXME
	w(self.funcsOutput, "return\n}\n") // body end

	// blank line
	w(self.funcsOutput, "\n")
}

func (self *Function) CollectInfo() {
	// is varargs
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
