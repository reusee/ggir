package main

import (
	"bytes"
	"strings"
	"unicode"
)

func (self *Generator) GenEnums(ns *Namespace) {
	self.genEnums(ns.Enums, "enums")
}

func (self *Generator) GenFlags(ns *Namespace) {
	self.genEnums(ns.Bitfields, "flags")
}

func (self *Generator) genEnums(enums []*Enum, file string) {
	output := new(bytes.Buffer)
	w(output, "package %s\n\n", self.PackageName)
	w(output, "/*\n")
	for _, include := range self.Includes {
		w(output, "#include <%s>\n", include)
	}
	w(output, "*/\n")
	w(output, "import \"C\"\n")

	w(output, "var (\n")
	for _, f := range enums {
		w(output, "\t// %s\n", f.Name)
		for _, m := range f.Members {
			parts := strings.Split(m.CIdentifier, "_")
			goName := strings.Join(parts[1:], "_")
			if unicode.IsDigit(rune(goName[0])) {
				goName = strings.ToUpper(self.PackageName) + "_" + goName
			}
			w(output, "%s = C.%s(C.%s)\n", goName, f.CType, m.CIdentifier)
		}
		w(output, "\n")
	}
	w(output, ")\n")

	self.formatAndOutput(file, output.Bytes())
}
