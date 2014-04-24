package main

import (
	"encoding/xml"
	"reflect"
)

type BaseInfo struct {
	XMLName       xml.Name
	Doc           *Doc          `xml:"doc"`
	DocDeprecated *Doc          `xml:"doc-deprecated"`
	DocVersion    *Doc          `xml:"doc-version"`
	NotHandled    []*NotHandled `xml:",any"`
}

type Doc struct {
	Text string `xml:",chardata"`
}

type NotHandled struct {
	XMLName xml.Name
	Xml     string `xml:",innerxml"`
}

type Repository struct {
	BaseInfo
	Package   *Package    `xml:"package"`
	CInclude  *CInclude   `xml:"include"`
	Namespace *Namespace  `xml:"namespace"`
	Constants []*Constant `xml:"constant"`
}

type Package struct {
	BaseInfo
	Name  string `xml:"name,attr"`
	CType string `xml:"type"`
}

type CInclude struct {
	BaseInfo
	Name string `xml:"name,attr"`
}

type Namespace struct {
	BaseInfo
	Name        string       `xml:"name,attr"`
	ErrorDomain *ErrorDomain `xml:"errordomain"`
	Aliases     []*Alias     `xml:"alias"`
	Constants   []*Constant  `xml:"constant"`
	Bitfields   []*Bitfield  `xml:"bitfield"`
	Functions   []*Function  `xml:"function"`
	Callbacks   []*Function  `xml:"callback"`
	Enums       []*Enum      `xml:"enumeration"`
	Unions      []*Union     `xml:"union"`
	Classes     []*Class     `xml:"class"`
	Interfaces  []*Class     `xml:"interface"`
	Records     []*Class     `xml:"record"`
	Boxeds      []*Boxed     `xml:"boxed"`
}

type ErrorDomain struct {
	BaseInfo
	Name       string      `xml:"name,attr"`
	GetQuark   string      `xml:"get-quark,attr"`
	Codes      string      `xml:"codes,attr"`
	Annotation *Annotation `xml:"annotation"`
}

type Annotation struct {
	Key   string `xml:"key,attr"`
	Value string `xml:"value,attr"`
}

type Alias struct {
	BaseInfo
	Name  string `xml:"name,attr"`
	Left  string `xml:"type,attr"`
	Right *Type  `xml:"type"`
}

type Constant struct {
	BaseInfo
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
	CName string `xml:"type,attr"`
	Type  *Type  `xml:"type"`
}

type Type struct {
	BaseInfo
	Name  string `xml:"name,attr"`
	CType string `xml:"type,attr"`
	Type  string `xml:"type"`
	Array *Array `xml:"array"`
}

type Class struct {
	BaseInfo
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
}

type Implement struct {
	BaseInfo
	Interface string `xml:"name,attr"`
}

type Prerequisite struct {
	BaseInfo
	Name string `xml:"name,attr"`
}

type Property struct {
	BaseInfo
	Name              string `xml:"name,attr"`
	Version           string `xml:"version,attr"`
	Writable          string `xml:"writable,attr"`
	TransferOwnership string `xml:"transfer-ownership,attr"`
	Type              *Type  `xml:"type"`
	Array             *Array `xml:"array"`
}

type Field struct {
	BaseInfo
	Name     string    `xml:"name,attr"`
	Writable string    `xml:"writable,attr"`
	Type     *Type     `xml:"type"`
	Array    *Array    `xml:"array"`
	Callback *Function `xml:"callback"`
}

type Union struct {
	BaseInfo
	Name         string      `xml:"name,attr"`
	CType        string      `xml:"type,attr"`
	Fields       []*Field    `xml:"field"`
	Methods      []*Function `xml:"method"`
	Record       *Class      `xml:"record"`
	Functions    []*Function `xml:"function"`
	Constructors []*Function `xml:"constructor"`
}

type Array struct {
	BaseInfo
	Name           string `xml:"name,attr"`
	ZeroTerminated string `xml:"zero-terminated,attr"`
	CType          string `xml:"type,attr"`
	FizedSize      string `xml:"fixed-sized,attr"`
	Type           *Type  `xml:"type"`
	Array          *Array `xml:"array"`
}

type Function struct {
	BaseInfo
	Name              string        `xml:"name,attr"`
	CType             string        `xml:"type,attr"` // for callback
	Version           string        `xml:"version,attr"`
	Deprecated        string        `xml:"deprecated,attr"`
	DeprecatedVersion string        `xml:"deprecated-version,attr"`
	When              string        `xml:"when,attr"` // for signal
	CIdentifier       string        `xml:"identifier,attr"`
	MovedTo           string        `xml:"moved-to,attr"` // in Gdk-3.0
	Return            *Return       `xml:"return-value"`
	InstanceParam     *Param        `xml:"parameters>instance-parameter"`
	Params            []*Param      `xml:"parameters>parameter"`
	DocStability      *DocStability `xml:"doc-stability"`
}

type DocStability struct {
	Space string `xml:"space,attr"`
}

type Return struct {
	BaseInfo
	Array *Array `xml:"array"`
	Type  *Type  `xml:"type"`
}

type Param struct {
	BaseInfo
	Name     string   `xml:"name,attr"`
	Transfer string   `xml:"transfer-ownership,attr"`
	Array    *Array   `xml:"array"`
	Type     *Type    `xml:"type"`
	Varargs  *Varargs `xml:"varargs"`
}

type Varargs struct {
	BaseInfo
}

type Bitfield struct {
	BaseInfo
	Name      string      `xml:"name,attr"`
	CType     string      `xml:"type,attr"`
	Members   []*Member   `xml:"member"`
	Functions []*Function `xml:"function"`
}

type Member struct {
	BaseInfo
	Name        string `xml:"name,attr"`
	Value       string `xml:"value,attr"`
	CIdentifier string `xml:"identifier,attr"`
}

type Enum struct {
	BaseInfo
	Name            string      `xml:"name,attr"`
	CType           string      `xml:"type,attr"`
	GlibErrorDomain string      `xml:"error-domain,attr"`
	Members         []*Member   `xml:"member"`
	Functions       []*Function `xml:"function"`
}

type Boxed struct {
	GlibName      string `xml:"name,attr"`
	CSymbolPrefix string `xml:"symbol-prefix,attr"`
	GlibTypeName  string `xml:"type-name,attr"`
	GlibGetType   string `xml:"get-type,attr"`
}

func (self *Generator) Parse(contents []byte) *Repository {
	// unmarshal
	var repo Repository
	err := xml.Unmarshal(contents, &repo)
	checkError(err)

	// dump
	//repo.dump()
	dumpNotHandled(reflect.ValueOf(repo))

	return &repo
}
