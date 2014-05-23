package main

import (
	"bytes"
	"strings"
)

func (self *Generator) GenSignals(ns *Namespace) {
	// common mappings
	self.SignalParamTypeMappings["gint"] = "int"
	self.SignalParamTypeMappings["guint"] = "uint"
	self.SignalParamTypeMappings["gint64"] = "int64"
	self.SignalParamTypeMappings["gboolean"] = "bool"
	self.SignalParamTypeMappings["gchar*"] = "string"
	self.SignalParamTypeMappings["gpointer"] = "unsafe.Pointer"
	self.SignalParamTypeMappings["gdouble"] = "float64"
	self.SignalParamTypeMappings["Array utf8"] = "[]string"
	self.SignalParamTypeMappings["Array GLib.Pid"] = "[]int"
	self.SignalParamTypeMappings["Array GLib.Quark"] = "[]uint32"

	// collect signal param type mapping info. local/external class/struct/enum
	for _, c := range ns.Classes {
		self.SignalParamTypeMappings[c.Name] = "*" + c.Name
		self.SignalParamTypeMappings["Array "+c.Name] = "*" + c.Name
	}
	/* FIXME implement interfaces first
	for _, c := range ns.Interfaces {
		self.SignalParamTypeMappings[c.Name] = "*" + c.Name
		self.SignalParamTypeMappings["Array "+c.Name] = "*" + c.Name
	}
	*/
	for _, s := range ns.Records {
		self.SignalParamTypeMappings[s.Name] = "*" + s.Name
		self.SignalParamTypeMappings["Array "+s.Name] = "*" + s.Name
	}
	for _, u := range ns.Unions {
		self.SignalParamTypeMappings[u.Name] = "*C." + u.CType
		self.SignalParamTypeMappings["Array "+u.Name] = "*C." + u.CType
	}
	for _, e := range ns.Enums {
		self.SignalParamTypeMappings[e.Name] = "C." + e.CType
		self.SignalParamTypeMappings["Array "+e.Name] = "*C." + e.CType
	}
	for _, e := range ns.Bitfields {
		self.SignalParamTypeMappings[e.Name] = "C." + e.CType
		self.SignalParamTypeMappings["Array "+e.Name] = "*C." + e.CType
	}
	for _, ext := range self.ExternalPackages {
		ns := ext.Repo.Namespace
		for _, c := range ns.Classes {
			self.SignalParamTypeMappings[ns.Name+"."+c.Name] = "*" + strings.ToLower(ns.Name) + "." + c.Name
			self.SignalParamTypeMappings["Array "+ns.Name+"."+c.Name] = "*" + strings.ToLower(ns.Name) + "." + c.Name
		}
		for _, c := range ns.Interfaces {
			self.SignalParamTypeMappings[ns.Name+"."+c.Name] = "*" + strings.ToLower(ns.Name) + "." + c.Name
			self.SignalParamTypeMappings["Array "+ns.Name+"."+c.Name] = "*" + strings.ToLower(ns.Name) + "." + c.Name
		}
		for _, s := range ns.Records {
			self.SignalParamTypeMappings[ns.Name+"."+s.Name] = "*" + strings.ToLower(ns.Name) + "." + s.Name
			self.SignalParamTypeMappings["Array "+ns.Name+"."+s.Name] = "*" + strings.ToLower(ns.Name) + "." + s.Name
		}
		for _, u := range ns.Unions {
			self.SignalParamTypeMappings[ns.Name+"."+u.Name] = "*C." + u.CType
			self.SignalParamTypeMappings["Array "+ns.Name+"."+u.Name] = "*C." + u.CType
		}
		for _, e := range ns.Enums {
			self.SignalParamTypeMappings[ns.Name+"."+e.Name] = "C." + e.CType
			self.SignalParamTypeMappings["Array "+ns.Name+"."+e.Name] = "*C." + e.CType
		}
		for _, e := range ns.Bitfields {
			self.SignalParamTypeMappings[ns.Name+"."+e.Name] = "C." + e.CType
			self.SignalParamTypeMappings["Array "+ns.Name+"."+e.Name] = "*C." + e.CType
		}
	}

	// output
	output := new(bytes.Buffer)
	w(output, "package %s\n\n", self.PackageName)
	w(output, "/*\n")
	for _, include := range self.Includes {
		w(output, "#include <%s>\n", include)
	}
	w(output, "*/\n")
	w(output, "import \"C\"\n")
	for _, ext := range self.ExternalPackages { // external import
		w(output, "import \"%s\"\n", ext.Import)
	}
	w(output, "import \"unsafe\"\n")
	w(output, "import \"reflect\"\n")
	w(output, `func init() {
		_ = unsafe.Pointer(nil)
		_ = reflect.ValueOf(nil)
	}
	`)
	for _, ext := range self.ExternalPackages {
		w(output, "var _ = %s.UnusedFix_\n", ext.Name)
	}

	// gen
	for _, c := range ns.Classes {
		self.genClassSignals(c, output)
	}
}

func (self *Generator) genClassSignals(c *Class, output *bytes.Buffer) {
	for _, signal := range c.Signals {
		self.genSignal(signal, c, output)
	}
}

func (self *Generator) genSignal(signal *Function, class *Class, output *bytes.Buffer) {
	// check deprecation
	if signal.Deprecated != "" {
		return
	}

	// map param types
	for _, param := range signal.Params {
		// get param spec
		var spec string
		if param.Type != nil {
			spec = param.Type.CType
			if spec == "" {
				spec = param.Type.Name
			}
		} else if param.Array != nil {
			spec = "Array " + param.Array.Type.Name
		}
		if spec == "" {
			panic(fs("param has no type %s %s", signal.Name, param.Name))
		}
		// map
		mappedType, ok := self.SignalParamTypeMappings[spec]
		if !ok {
			p("FIXME signal param type mapping -> %s %s -> %s\n", signal.Name, param.Name, spec)
			return
		}
		param.MappedType = mappedType
		// param name
		param.GoName = param.GetGoName()
	}

	parts := strings.Split(signal.Name, "-")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}
	signalName := strings.Join(parts, "")
	// signature
	w(output, "func (self *Trait%s) On%s(f func(", class.Name, signalName)
	for _, param := range signal.Params {
		w(output, "%s, ", param.MappedType)
	}
	w(output, ")) {\n")
	// body
	w(output, `gobject.Connect(unsafe.Pointer(self.CPointer), "%s", func(_self reflect.Value,`, signal.Name)
	for _, param := range signal.Params {
		w(output, "%s reflect.Value,", param.GoName)
	}
	w(output, ") {\n")
	w(output, "})\n")
	w(output, "}\n\n")

	self.formatAndOutput("signals", output.Bytes())
}
