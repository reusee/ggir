package main

import "strings"

func (self *Generator) GenFunction(fn *Function) {
	// skip varargs functions
	for _, param := range fn.Params {
		if param.Varargs != nil {
			w(self.funcsOutput, "// %s is not generated due to varargs\n\n", fn.CIdentifier)
			return
		}
	}

	// doc
	if fn.Doc != nil {
		w(self.funcsOutput, "/*\n%s\n*/\n", fn.Doc.Text)
	}

	// receiver TODO

	// name
	goFuncName := convertFuncName(fn.Name)
	w(self.funcsOutput, "func %s", goFuncName)

	// collect params
	var inParams, outParams []*Param
	for _, param := range fn.Params {
		param.GoName = convertParamName(param.Name)
		convertParamType(param)
		if param.Direction == "out" {
			outParams = append(outParams, param)
		} else if param.Direction == "inout" {
			inParams = append(inParams, param)
			outParams = append(outParams, param)
		} else {
			inParams = append(inParams, param)
			param.Direction = "in"
		}
	}

	// input param
	w(self.funcsOutput, "(") // start in param
	for _, param := range inParams {
		w(self.funcsOutput, "%s %s, ", param.GoName, param.GoType)
	}
	w(self.funcsOutput, ")") // end in param

	// output param
	w(self.funcsOutput, "(") // start out param
	for _, param := range outParams {
		w(self.funcsOutput, "__out__%s %s, ", param.GoName, param.GoType)
	}

	// return value
	convertParamType(fn.Return)
	if fn.Return.GoType != "C.void" {
		w(self.funcsOutput, "__return__ %s, ", fn.Return.GoType)
		outParams = append(outParams, fn.Return) // add return value to out params
	}

	// throws error
	throws := fn.Throws == "1"
	if throws {
		w(self.funcsOutput, "__err__ error, ")
	}
	w(self.funcsOutput, ")") // end out param

	// body
	w(self.funcsOutput, "{\n") // body start
	//p("%s\n", fn.CIdentifier) //TODO
	w(self.funcsOutput, "return\n}\n") // body end

	// blank line
	w(self.funcsOutput, "\n")
}

func convertFuncName(n string) string {
	parts := strings.Split(n, "_")
	for i, p := range parts {
		parts[i] = strings.Title(p)
	}
	return strings.Join(parts, "")
}

func convertParamName(n string) string {
	if isGoReservedWord(n) {
		return n + "_"
	}
	return n
}

func convertParamType(param *Param) {
	var cType string
	if param.Array != nil {
		cType = param.Array.CType
	} else if param.Type != nil {
		cType = param.Type.CType
	} else {
		panic("param has no type")
	}
	param.GoType = cTypeToGoType(cType)
}

func cTypeToGoType(t string) string {
	t = strings.Replace(t, "const", "", -1) // const and non-const pointer is compatible in Go 1.2
	t = strings.TrimSpace(t)
	t = strings.Replace(t, "volatile", "", -1) // no volatile
	t = strings.TrimSpace(t)
	t = strings.Replace(t, "long double", "double", -1) // no long double, FIXME
	t = strings.TrimSpace(t)
	pointerDepth := 0 // pointer depth
	for strings.HasSuffix(t, "*") {
		pointerDepth++
		t = t[:len(t)-1]
	}
	t = "C." + t                              // cgo type
	t = strings.Repeat("*", pointerDepth) + t // pointer
	return strings.TrimSpace(t)
}
