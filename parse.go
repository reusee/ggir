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
	Name       string      `xml:"name,attr"`
	Aliases    []*Alias    `xml:"alias"`
	Constants  []*Constant `xml:"constant"`
	Bitfields  []*Bitfield `xml:"bitfield"`
	Functions  []*Function `xml:"function"`
	Callbacks  []*Function `xml:"callback"`
	Enums      []*Enum     `xml:"enumeration"`
	Unions     []*Union    `xml:"union"`
	Classes    []*Class    `xml:"class"`
	Interfaces []*Class    `xml:"interface"`
	Records    []*Class    `xml:"record"`

	NotHandled []*NotHandled `xml:",any"`
}

type NotHandled struct {
	XMLName xml.Name
	Xml     string `xml:",innerxml"`
}

type Alias struct {
	Name          string `xml:"name,attr"`
	Left          string `xml:"type,attr"`
	Right         *Type  `xml:"type"`
	Doc           *Doc   `xml:"doc"`
	DocDeprecated *Doc   `xml:"doc-deprecated"`

	NotHandled []*NotHandled `xml:",any"`
}

type Constant struct {
	Name          string `xml:"name,attr"`
	Value         string `xml:"value,attr"`
	CName         string `xml:"type,attr"`
	Type          *Type  `xml:"type"`
	Doc           *Doc   `xml:"doc"`
	DocDeprecated *Doc   `xml:"doc-deprecated"`

	NotHandled []*NotHandled `xml:",any"`
}

type Type struct {
	Name  string `xml:"name,attr"`
	CType string `xml:"type,attr"`
}

type Class struct {
	Name           string        `xml:"name,attr"`
	CSymbolPrefix  string        `xml:"symbol-prefix,attr"`
	CType          string        `xml:"type,attr"`
	Parent         string        `xml:"parent,attr"`
	Prerequisite   *Prerequisite `xml:"prerequisite"` // for interface
	GlibTypeName   string        `xml:"type-name,attr"`
	GlibGetType    string        `xml:"get-type,attr"`
	GlibTypeStruct string        `xml:"type-struct,attr"`
	Implements     []*Implement  `xml:"implements"`
	Constructors   []*Function   `xml:"constructor"`
	VirtualMethods []*Function   `xml:"virtual-method"`
	Methods        []*Function   `xml:"method"`
	Functions      []*Function   `xml:"function"`
	Properties     []*Property   `xml:"property"`
	Union          *Union        `xml:"union"` // for record
	Fields         []*Field      `xml:"field"`
	Signals        []*Function   `xml:"signal"`
	Doc            *Doc          `xml:"doc"`
	DocDeprecated  *Doc          `xml:"doc-deprecated"`

	NotHandled []*NotHandled `xml:",any"`
}

type Implement struct {
	Interface string `xml:"name,attr"`
}

type Prerequisite struct {
	Name string `xml:"name,attr"`
}

type Property struct {
	Name              string `xml:"name,attr"`
	Version           string `xml:"version,attr"`
	Writable          string `xml:"writable,attr"`
	TransferOwnership string `xml:"transfer-ownership,attr"`
	Doc               *Doc   `xml:"doc"`
	DocDeprecated     *Doc   `xml:"doc-deprecated"`
	Type              *Type  `xml:"type"`
	Array             *Array `xml:"array"`
}

type Field struct {
	Name     string `xml:"name,attr"`
	Writable string `xml:"writable,attr"`
	Type     *Type  `xml:"type"`
	Array    *Array `xml:"array"`
}

type Union struct {
	Name          string      `xml:"name,attr"`
	CType         string      `xml:"type,attr"`
	Fields        []*Field    `xml:"field"`
	Methods       []*Function `xml:"method"`
	Record        *Class      `xml:"record"`
	Doc           *Doc        `xml:"doc"`
	DocDeprecated *Doc        `xml:"doc-deprecated"`

	NotHandled []*NotHandled `xml:",any"`
}

type Array struct {
	Name           string `xml:"name,attr"`
	ZeroTerminated string `xml:"zero-terminated,attr"`
	CType          string `xml:"type,attr"`
	FizedSize      string `xml:"fixed-sized,attr"`
	Type           *Type  `xml:"type"`
}

type Function struct {
	Name              string   `xml:"name,attr"`
	CType             string   `xml:"type,attr"` // for callback
	Version           string   `xml:"version,attr"`
	Deprecated        string   `xml:"deprecated,attr"`
	DeprecatedVersion string   `xml:"deprecated-version,attr"`
	When              string   `xml:"when,attr"` // for signal
	CIdentifier       string   `xml:"identifier,attr"`
	Return            *Return  `xml:"return-value"`
	InstanceParam     *Param   `xml:"parameters>instance-parameter"`
	Params            []*Param `xml:"parameters>parameter"`
	Doc               *Doc     `xml:"doc"`
	DocDeprecated     *Doc     `xml:"doc-deprecated"`

	NotHandled []*NotHandled `xml:",any"`
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
	Name          string    `xml:"name,attr"`
	CType         string    `xml:"type,attr"`
	Members       []*Member `xml:"member"`
	Doc           *Doc      `xml:"doc"`
	DocDeprecated *Doc      `xml:"doc-deprecated"`

	NotHandled []*NotHandled `xml:",any"`
}

type Member struct {
	Name        string `xml:"name,attr"`
	Value       string `xml:"value,attr"`
	CIdentifier string `xml:"identifier,attr"`
}

type Enum struct {
	Name            string      `xml:"name,attr"`
	CType           string      `xml:"type,attr"`
	GlibErrorDomain string      `xml:"error-domain,attr"`
	Members         []*Member   `xml:"member"`
	Functions       []*Function `xml:"function"`
	Doc             *Doc        `xml:"doc"`
	DocDeprecated   *Doc        `xml:"doc-deprecated"`

	NotHandled []*NotHandled `xml:",any"`
}

type Doc struct {
	Text string `xml:",chardata"`
}

func (self *Generator) Parse(contents []byte) *Repository {
	// unmarshal
	var repo Repository
	err := xml.Unmarshal(contents, &repo)
	checkError(err)

	// dump
	repo.dump()
	repo.Namespace.dumpNotHandled()

	return &repo
}
