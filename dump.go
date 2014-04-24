package main

import "strings"

func (self *Repository) dump() {
	p("package %s\n", self.Package.Name)
	p("cinlucde %s\n", self.CInclude.Name)
	self.Namespace.dump()
}

func (self *Namespace) dump() {
	p("namespace %s\n", self.Name)
	for _, alias := range self.Aliases {
		alias.dump()
	}
	for _, c := range self.Constants {
		c.dump()
	}
	for _, record := range self.Records {
		record.dump()
	}
	for _, b := range self.Bitfields {
		b.dump()
	}
	for _, f := range self.Functions {
		f.dump(false)
	}
	for _, e := range self.Enums {
		e.dump()
	}
}

func (self *Namespace) dumpNotHandled() {
	for _, n := range self.NotHandled {
		n.dump()
	}
	for _, e := range self.Aliases {
		for _, n := range e.NotHandled {
			n.dump()
		}
	}
	for _, e := range self.Constants {
		for _, n := range e.NotHandled {
			n.dump()
		}
	}
	for _, e := range self.Records {
		for _, n := range e.NotHandled {
			n.dump()
		}
	}
	for _, e := range self.Bitfields {
		for _, n := range e.NotHandled {
			n.dump()
		}
	}
	for _, e := range self.Functions {
		for _, n := range e.NotHandled {
			n.dump()
		}
	}
	for _, e := range self.Enums {
		for _, n := range e.NotHandled {
			n.dump()
		}
	}
	for _, e := range self.Callbacks {
		for _, n := range e.NotHandled {
			n.dump()
		}
	}
	for _, e := range self.Unions {
		for _, n := range e.NotHandled {
			n.dump()
		}
	}
	for _, e := range self.Classes {
		for _, n := range e.NotHandled {
			n.dump()
		}
	}
	for _, e := range self.Interfaces {
		for _, n := range e.NotHandled {
			n.dump()
		}
	}
}

func (f *Function) dump(isMethod bool) {
	if isMethod {
		p("method %s", f.CIdentifier)
	} else {
		p("function %s", f.CIdentifier)
	}
	p(" (")
	for _, param := range f.Params {
		if param.Type != nil {
			p("%s %s, ", param.Name, param.Type.CType)
		}
		if param.Array != nil {
			p("%s ARRAY of %s, ", param.Name, param.Array.Type.CType)
		}
	}
	p(")")
	if f.Return.Type != nil {
		p(" -> %s", f.Return.Type.CType)
	}
	if f.Return.Array != nil {
		p(" -> ARRAY of %s", f.Return.Array.Type.CType)
	}
	p("\n")
}

func (alias *Alias) dump() {
	p("alias %s %s\n", alias.Left, alias.Right.CType)
}

func (c *Constant) dump() {
	p("const %s %s %s %s\n", c.Name, c.Value, c.CName, c.Type.CType)
}

func (record *Class) dump() {
	p("record %s\n", record.CType)
	for _, field := range record.Fields {
		p("field %s", field.Name)
		if field.Type != nil {
			p(" %s", field.Type.CType)
		}
		if field.Array != nil {
			p(" ARRAY of %s", field.Array.Type.CType)
		}
		p("\n")
	}
	for _, f := range record.Functions {
		f.dump(false)
	}
	for _, m := range record.Methods {
		m.dump(true)
	}
	p("\n")
}

func (b *Bitfield) dump() {
	p("bitfield %s %s\n", b.Name, b.CType)
	for _, m := range b.Members {
		p("%s %s %s\n", m.Name, m.CIdentifier, m.Value)
	}
	p("\n")
}

func (e *Enum) dump() {
	p("enum %s %s\n", e.Name, e.CType)
	for _, m := range e.Members {
		p("%s %s %s\n", m.Name, m.Value, m.CIdentifier)
	}
	p("\n")
}

func (n *NotHandled) dump() {
	p("%s\n", strings.Repeat("=", 64))
	p("%v\n", n)
}
