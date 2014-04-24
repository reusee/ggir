package main

import "encoding/xml"

type Repository struct {
	Package   *Package   `xml:"package"`
	CInclude  *CInclude  `xml:"include"`
	Namespace *Namespace `xml:"namespace"`
}

type Package struct {
	Name string `xml:"name,attr"`
}

type CInclude struct {
	Name string `xml:"name,attr"`
}

type Namespace struct {
	Name      string      `xml:"name,attr"`
	Aliases   []*Alias    `xml:"alias"`
	Constants []*Constant `xml:"constant"`
	Records   []*Record   `xml:"record"`
	Bitfields []*Bitfield `xml:"bitfield"`
	Functions []*Function `xml:"function"`
	Enums     []*Enum     `xml:"enumeration"`
	Callbacks []*Callback `xml:"callback"`
	Unions    []*Union    `xml:"union"`

	NotHandled []*NotHandled `xml:",any"`
}

type NotHandled struct {
	Xml string `xml:",innerxml"`
}

type Alias struct {
	Name  string `xml:"name,attr"`
	Left  string `xml:"type,attr"`
	Right *Type  `xml:"type"`
}

type Constant struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
	CName string `xml:"type,attr"`
	Type  *Type  `xml:"type"`
}

type Type struct {
	Name  string `xml:"name,attr"`
	CType string `xml:"type,attr"`
}

type Record struct {
	Name          string      `xml:"name,attr"`
	CType         string      `xml:"type,attr"`
	GlibTypeName  string      `xml:"type-name,attr"`
	GlibGetType   string      `xml:"get-type,attr"`
	CSymbolPrefix string      `xml:"symbol-prefix,attr"`
	Fields        []*Field    `xml:"field"`
	Functions     []*Function `xml:"function"`
	Methods       []*Function `xml:"method"`
}

type Field struct {
	Name     string `xml:"name,attr"`
	Writable string `xml:"writable,attr"`
	Type     *Type  `xml:"type"`
	Array    *Array `xml:"array"`
}

type Union struct {
	Name   string   `xml:"name,attr"`
	CType  string   `xml:"type,attr"`
	Fields []*Field `xml:"field"`
}

type Array struct {
	Name           string `xml:"name,attr"`
	ZeroTerminated string `xml:"zero-terminated,attr"`
	CType          string `xml:"type,attr"`
	FizedSize      string `xml:"fixed-sized,attr"`
	Type           *Type  `xml:"type"`
}

type Function struct {
	Name        string   `xml:"name,attr"`
	CIdentifier string   `xml:"identifier,attr"`
	Return      *Return  `xml:"return-value"`
	Params      []*Param `xml:"parameters>parameter"`
}

type Return struct {
	Array *Array `xml:"array"`
	Type  *Type  `xml:"type"`
}

type Param struct {
	Name     string `xml:"name,attr"`
	Transfer string `xml:"transfer-ownership,attr"`
	Array    *Array `xml:"array"`
	Type     *Type  `xml:"type"`
}

type Bitfield struct {
	Name    string    `xml:"name,attr"`
	CType   string    `xml:"type,attr"`
	Members []*Member `xml:"member"`
}

type Member struct {
	Name        string `xml:"name,attr"`
	Value       string `xml:"value,attr"`
	CIdentifier string `xml:"identifier,attr"`
}

type Enum struct {
	Name            string    `xml:"name,attr"`
	CType           string    `xml:"type,attr"`
	GlibErrorDomain string    `xml:"error-domain,attr"`
	Members         []*Member `xml:"member"`
}

type Callback struct {
	Name   string   `xml:"name,attr"`
	CType  string   `xml:"type,attr"`
	Return *Return  `xml:"return-value"`
	Params []*Param `xml:"parameters>parameter"`
}

func (self *Generator) Parse(contents []byte) *Repository {
	// unmarshal
	var repo Repository
	err := xml.Unmarshal(contents, &repo)
	checkError(err)

	// dump
	repo.dump()
	for _, alias := range repo.Namespace.Aliases {
		alias.dump()
	}
	for _, c := range repo.Namespace.Constants {
		c.dump()
	}
	for _, record := range repo.Namespace.Records {
		record.dump()
	}
	for _, b := range repo.Namespace.Bitfields {
		b.dump()
	}
	for _, f := range repo.Namespace.Functions {
		f.dump(false)
	}
	for _, e := range repo.Namespace.Enums {
		e.dump()
	}
	for _, n := range repo.Namespace.NotHandled {
		n.dump()
	}

	return &repo
}
