package main

import "bytes"

func (self *Generator) GenTraits(ns *Namespace) {
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

	// classes
	for _, c := range ns.Classes {
		goType := cTypeToGoType(c.CType)
		w(output, "type _Trait%s struct { CPointer *%s }\n", c.Name, goType)
		w(output, "type Is%s interface { Get%sPointer() *%s }\n", c.Name, c.Name, goType)
		w(output, "func (self *_Trait%s) Get%sPointer() *%s {\n return self.CPointer\n }\n",
			c.Name, c.Name, goType)
		for _, m := range c.Methods {
			self.GenFunction(m, output, c)
		}
		//p("%v\n", c.VirtualMethods) FIXME
		//p("%v\n", c.Methods) FIXME
		//p("%v\n", c.Functions) FIXME

		w(output, "\n")
	}
	// interfaces FIXME
	//p("%v\n", c.VirtualMethods) FIXME
	//p("%v\n", c.Methods) FIXME
	//p("%v\n", c.Functions) FIXME

	self.formatAndOutput("traits", output.Bytes())
}
