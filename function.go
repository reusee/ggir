package main

import (
	"bytes"
	"io"
	"regexp"
	"strings"
)

func (self *Generator) GenFunctions(ns *Namespace) {
	output := new(bytes.Buffer)
	w(output, "package %s\n\n", self.PackageName)
	w(output, "/*\n")
	for _, include := range self.Includes {
		w(output, "#include <%s>\n", include)
	}
	if len(self.PkgConfigs) > 0 {
		w(output, "#cgo pkg-config: %s\n", strings.Join(self.PkgConfigs, " "))
	}
	w(output, "*/\n")
	w(output, "import \"C\"\n\n")
	w(output, `import ("unsafe";	"reflect"; "errors")
func init() {
	var _ unsafe.Pointer
	var _ reflect.SliceHeader
	_ = errors.New("")
}
`)
	for _, fn := range ns.Functions {
		self.GenFunction(fn, output, nil)
	}

	self.formatAndOutput("functions", output.Bytes())
}

func (self *Generator) GenFunction(fn *Function, output io.Writer, receiver *Class) {
	// skip deprecated
	if fn.Deprecated != "" {
		w(output, "// %s is not generated due to deprecation attr\n\n", fn.CIdentifier)
		return
	}

	// skip explicit ignored
	for _, pattern := range self.FunctionIgnorePatterns {
		matched, err := regexp.MatchString(pattern, fn.CIdentifier)
		checkError(err)
		if matched {
			w(output, "// %s is not generated due to explicit ignore\n\n", fn.CIdentifier)
			return
		}
	}

	// collect info
	fn.CollectInfo()

	// rename
	for _, rename := range self.FunctionRenames {
		matched, err := regexp.MatchString(rename.From, fn.GoName)
		checkError(err)
		if matched {
			fn.GoName = regexp.MustCompile(rename.From).ReplaceAllString(fn.GoName, rename.To)
		}
	}

	// patch param direction
	for _, param := range fn.Params {
		spec := fs("%s %s", fn.CIdentifier, param.Name)
		if dir, ok := self.ParamDirections[spec]; ok {
			param.Direction = dir
		}
	}

	// skip varargs functions
	if fn.IsVarargs {
		w(output, "// %s is not generated due to varargs\n\n", fn.CIdentifier)
		return
	}

	//FIXME skip function with inout param due to broken rule
	for _, param := range fn.Params {
		if param.Direction == "inout" {
			p("===fixme=== function with inout param %s\n", fn.CIdentifier)
			w(output, "// %s is not generated due to inout param\n\n", fn.CIdentifier)
			return
		}
	}

	// collect params and return info
	for _, param := range fn.Params {
		param.Function = fn
		param.Generator = self
		param.CollectInfo(false, fn)

		//FIXME skip functions with long double param
		if param.CType == "long double" {
			w(output, "// %s is not generated due to long double param\n\n", fn.CIdentifier)
			return
		}
	}
	fn.Return.Function = fn
	fn.Return.Generator = self
	fn.Return.CollectInfo(true, fn)

	// generate doc
	if fn.Doc != nil {
		w(output, "/*\n%s\n*/\n", fn.Doc.Text)
	}

	// generate signature
	w(output, "func ")

	// generate receiver
	if receiver != nil {
		w(output, "(self *Trait%s)", receiver.Name)
	}

	// generate function name
	w(output, fn.GoName)

	// generate parameters
	w(output, "(")
	for _, param := range fn.Params {
		if !param.IsVoid && param.Direction == "in" {
			w(output, "%s %s, ", param.GoName, param.MappedType)
		}
	}
	w(output, ")")

	// generate returns
	w(output, "(")
	for _, param := range fn.Params {
		if !param.IsVoid && param.Direction == "out" {
			w(output, "%s %s, ", param.GoName, param.MappedType)
		}
	}
	if !fn.Return.IsVoid { // return
		w(output, "%s %s, ", fn.Return.GoName, fn.Return.MappedType)
	}
	if fn.Throws != "" { // error
		w(output, "__err__ error")
	}
	w(output, ")")

	// body start
	w(output, "{\n")

	// statements before cgo call
	for _, param := range fn.Params {
		if param.CgoBeforeStmt != "" {
			w(output, "%s\n", param.CgoBeforeStmt)
		}
	}

	// error param
	if fn.Throws != "" {
		w(output, "var __cgo_error__ *C.GError\n")
	}

	// cgo return value
	if !fn.Return.IsVoid {
		if fn.Return.GoType != fn.Return.MappedType { // mapped
			if fn.IsConstructor { // cgo return and constructor return may be different type
				w(output, "var __cgo__return__ interface{}\n")
				w(output, "__cgo__return__ = ")
			} else {
				w(output, "var __cgo__return__ %s\n", fn.Return.GoType)
				w(output, "__cgo__return__ = ")
			}
		} else { // not mapped
			w(output, "%s = ", fn.Return.GoName)
		}
	}

	// cgo call
	w(output, "C.%s(", fn.CIdentifier)
	if receiver != nil {
		cReceiverType := cTypeToGoType(fn.InstanceParam.Type.CType)
		if cReceiverType != "*"+cTypeToGoType(receiver.CType) { // GObject's receiver type is C.gpointer
			w(output, fs("(%s)(unsafe.Pointer(self.CPointer)),", cReceiverType))
		} else {
			w(output, "self.CPointer,")
		}
	}
	for _, param := range fn.Params {
		if !param.IsVoid {
			w(output, "%s,", param.CgoParam)
		}
	}
	if fn.Throws != "" { // error param
		w(output, "&__cgo_error__")
	}
	w(output, ")\n")

	// statements after cgo call
	for _, param := range fn.Params {
		if param.CgoAfterStmt != "" {
			w(output, "%s\n", param.CgoAfterStmt)
		}
	}
	if fn.Return.CgoAfterStmt != "" {
		w(output, "%s\n", fn.Return.CgoAfterStmt)
	}

	// convert GError to error
	if fn.Throws != "" {
		w(output, `if __cgo_error__ != nil {
	__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
}
`)
	}

	// body end
	w(output, "return\n}\n")

	// blank line
	w(output, "\n")
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
