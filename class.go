package main

func (self *Generator) GenClassTypes(ns *Namespace) {
	for _, c := range ns.Classes {
		p("%s\n", c.Name)
		p("%s\n", c.CSymbolPrefix)
		p("%s\n", c.CType)
		p("%s\n", c.Parent)
		p("<<<\n")
	}
}
