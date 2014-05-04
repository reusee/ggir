package main

func (self *Generator) GenSignals(ns *Namespace) {
	// common mappings
	self.SignalParamTypeMappings["gint"] = "int"
	self.SignalParamTypeMappings["guint"] = "uint"
	self.SignalParamTypeMappings["gboolean"] = "bool"
	self.SignalParamTypeMappings["gchar*"] = "string"
	self.SignalParamTypeMappings["gpointer"] = "unsafe.Pointer"
	self.SignalParamTypeMappings["gdouble"] = "float64"
	// collect signal param type mapping info. local/external class/struct/enum
	for _, c := range ns.Classes {
		self.SignalParamTypeMappings[c.Name] = "*" + c.Name
	}
	for _, c := range ns.Interfaces {
		self.SignalParamTypeMappings[c.Name] = "*" + c.Name
	}
	for _, s := range ns.Records {
		self.SignalParamTypeMappings[s.Name] = "*" + s.Name
	}
	for _, e := range ns.Enums {
		self.SignalParamTypeMappings[e.Name] = "C." + e.CType
	}
	for _, e := range ns.Bitfields {
		self.SignalParamTypeMappings[e.Name] = "C." + e.CType
	}
	for _, u := range ns.Unions {
		self.SignalParamTypeMappings[u.Name] = "*C." + u.CType
	}
	for _, ext := range self.ExternalPackages {
		ns := ext.Repo.Namespace
		for _, c := range ns.Classes {
			self.SignalParamTypeMappings[ns.Name+"."+c.Name] = "*" + ns.Name + "." + c.Name
		}
		for _, c := range ns.Interfaces {
			self.SignalParamTypeMappings[ns.Name+"."+c.Name] = "*" + ns.Name + "." + c.Name
		}
		for _, s := range ns.Records {
			self.SignalParamTypeMappings[ns.Name+"."+s.Name] = "*" + ns.Name + "." + s.Name
		}
		for _, e := range ns.Enums {
			self.SignalParamTypeMappings[ns.Name+"."+e.Name] = "C." + e.CType
		}
		for _, e := range ns.Bitfields {
			self.SignalParamTypeMappings[ns.Name+"."+e.Name] = "C." + e.CType
		}
		for _, u := range ns.Unions {
			self.SignalParamTypeMappings[ns.Name+"."+u.Name] = "*C." + u.CType
		}
	}
	// gen
	for _, c := range ns.Classes {
		self.genClassSignals(c)
	}
}

func (self *Generator) genClassSignals(c *Class) {
	for _, signal := range c.Signals {
		self.genSignal(signal)
	}
}

func (self *Generator) genSignal(signal *Function) {
	// check deprecation
	if signal.Deprecated != "" {
		return
	}
	// map param types
	for _, param := range signal.Params {
		// get param type
		paramType := param.Type.CType
		if paramType == "" {
			paramType = param.Type.Name
		}
		if paramType == "" {
			panic(fs("param has no type %s %s", signal.Name, param.Name))
		}
		mappedType, ok := self.SignalParamTypeMappings[paramType]
		if !ok {
			p("FIXME signal param type mapping -> %s\n", paramType)
			return
		}
		_ = mappedType
	}
}
