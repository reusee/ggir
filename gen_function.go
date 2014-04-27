package main

import (
	"regexp"
	"strings"
)

var typeStat = make(map[string][]string)

func (self *Generator) GenFunction(fn *Function) {
	// skip deprecated
	if fn.Deprecated != "" {
		w(self.funcsOutput, "// %s is not generated due to deprecation attr\n\n", fn.CIdentifier)
		return
	}
	for _, f := range self.FunctionDeprecated {
		if fn.CIdentifier == f {
			w(self.funcsOutput, "// %s is not generated due to explicit deprecation\n\n", fn.CIdentifier)
			return
		}
	}

	// skip explicit ignored
	for _, pattern := range self.FunctionIgnorePatterns {
		matched, err := regexp.MatchString(pattern, fn.CIdentifier)
		checkError(err)
		if matched {
			w(self.funcsOutput, "// %s is not generated due to explicit ignore\n\n", fn.CIdentifier)
			return
		}
	}

	// collect info
	fn.CollectInfo()

	// skip varargs functions
	if fn.IsVarargs {
		w(self.funcsOutput, "// %s is not generated due to varargs\n\n", fn.CIdentifier)
		return
	}

	//FIXME skip function with inout param due to broken rule
	for _, param := range fn.Params {
		if param.Direction == "inout" {
			w(self.funcsOutput, "// %s is not generated due to inout param\n\n", fn.CIdentifier)
			return
		}
	}

	// collect params and return info
	for _, param := range fn.Params {
		param.CollectInfo(false)

		//FIXME skip functions with long double param
		if param.CType == "long double" {
			w(self.funcsOutput, "// %s is not generated due to long double param\n\n", fn.CIdentifier)
			return
		}

		if !param.IsVoid && param.MappedType == "" { // add rules for these specs
			p("%s %s %s\n", fn.Name, param.Name, param.CType)
			typeStat[param.TypeSpec] = append(typeStat[param.TypeSpec], param.Name+" @ "+fn.Name)
		}
	}
	fn.Return.CollectInfo(true)
	if !fn.Return.IsVoid {
		if fn.Return.MappedType == "" { // add rules for return type
			typeStat[fn.Return.TypeSpec] = append(typeStat[fn.Return.TypeSpec], fn.Return.Name+" @ "+fn.Name)
		}
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

	// body start
	w(self.funcsOutput, "{\n")

	// statements before cgo call
	for _, param := range fn.Params {
		if param.CgoBeforeStmt != "" {
			w(self.funcsOutput, "%s\n", param.CgoBeforeStmt)
		}
	}

	// error param
	if fn.Throws != "" {
		w(self.funcsOutput, "var __cgo_error__ *C.GError\n")
	}

	// cgo return value
	if !fn.Return.IsVoid {
		if fn.Return.GoType != fn.Return.MappedType { // mapped
			w(self.funcsOutput, "var __cgo__return__ %s\n", fn.Return.GoType)
			w(self.funcsOutput, "__cgo__return__ = ")
		} else { // not mapped
			w(self.funcsOutput, "%s = ", fn.Return.GoName)
		}
	}

	// cgo call
	w(self.funcsOutput, "C.%s(", fn.CIdentifier)
	for _, param := range fn.Params {
		if !param.IsVoid {
			w(self.funcsOutput, "%s,", param.CgoParam)
		}
	}
	if fn.Throws != "" { // error param
		w(self.funcsOutput, "&__cgo_error__")
	}
	w(self.funcsOutput, ")\n")

	// statements after cgo call
	for _, param := range fn.Params {
		if param.CgoAfterStmt != "" {
			w(self.funcsOutput, "%s\n", param.CgoAfterStmt)
		}
	}
	if fn.Return.CgoAfterStmt != "" {
		w(self.funcsOutput, "%s\n", fn.Return.CgoAfterStmt)
	}

	// convert GError to error
	if fn.Throws != "" {
		w(self.funcsOutput, `if __cgo_error__ != nil {
	__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
}
`)
	}

	// body end
	w(self.funcsOutput, "return\n}\n")

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
