package main

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
)

func (self *Generator) GenConstants(ns *Namespace) {
	for _, c := range ns.Constants {
		for _, pattern := range self.ConstantIgnorePatterns {
			matched, err := regexp.MatchString(pattern, c.CName)
			checkError(err)
			if matched {
				c.Ignored = true
			}
		}
	}

	output := new(bytes.Buffer)
	w(output, "package %s\n\n", self.PackageName)
	w(output, "/*\n")
	for _, include := range self.Includes {
		w(output, "#include <%s>\n", include)
	}
	w(output, "\n")
	for _, c := range ns.Constants {
		if c.Ignored {
			continue
		}
		w(output, "%s __%s = %s;\n", c.Type.CType, c.CName, c.CName)
	}
	w(output, "*/\n")
	w(output, "import \"C\"\n")

	w(output, "var (\n")
	for _, c := range ns.Constants {
		if c.Ignored {
			w(output, "\t// %s is not generated due to explicit ignore\n", c.CName)
			continue
		}
		w(output, "%s = C.__%s\n", c.Name, c.CName)
	}
	w(output, ")\n")

	f, err := os.Create(filepath.Join(self.outputDir, self.PackageName+"_constants.go"))
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
