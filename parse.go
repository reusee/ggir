package main

import (
	"reflect"
	"strings"

	"./xml"
)

type BaseInfo struct {
	XMLName           xml.Name
	Version           string        `xml:"version,attr"`
	Deprecated        string        `xml:"deprecated,attr"`
	DeprecatedVersion string        `xml:"deprecated-version,attr"`
	Introspectable    string        `xml:"introspectable,attr"`
	Stability         string        `xml:"stability,attr"`
	Doc               *Doc          `xml:"doc"`
	DocDeprecated     *Doc          `xml:"doc-deprecated"`
	DocVersion        *Doc          `xml:"doc-version"`
	DocStability      *DocStability `xml:"doc-stability"`
	NotHandled        []*NotHandled `xml:",any"`
}

type Doc struct {
	Space      string `xml:"space,attr"`
	Whitespace string `xml:"whitespace,attr"`
	Text       string `xml:",chardata"`
}

type DocStability struct {
	Space string `xml:"space,attr"`
}

type NotHandled struct {
	XMLName xml.Name
	Xml     string `xml:",innerxml"`
}

type Repository struct {
	BaseInfo
	Glib                string      `xml:"glib,attr"`
	CSymbolPrefixes     string      `xml:"symbol-prefixes,attr"`
	CIdentifierPrefixes string      `xml:"identifier-prefixes,attr"`
	C                   string      `xml:"c,attr"`
	Xmlns               string      `xml:"xmlns,attr"`
	Package             *Package    `xml:"package"`
	CInclude            *CInclude   `xml:"include"`
	Namespace           *Namespace  `xml:"namespace"`
	Constants           []*Constant `xml:"constant"`
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
	Name                string         `xml:"name,attr"`
	CSymbolPrefixes     string         `xml:"symbol-prefixes,attr"`
	CIdentifierPrefixes string         `xml:"identifier-prefixes,attr"`
	SharedLibrary       string         `xml:"shared-library,attr"`
	CPrefix             string         `xml:"prefix,attr"`
	CSymbolPrefix       string         `xml:"symbol-prefix,attr"`
	Functions           []*Function    `xml:"function"`
	Callbacks           []*Function    `xml:"callback"`
	Classes             []*Class       `xml:"class"`
	Interfaces          []*Class       `xml:"interface"`
	Records             []*Class       `xml:"record"`
	Bitfields           []*Bitfield    `xml:"bitfield"`
	Enums               []*Enum        `xml:"enumeration"`
	Unions              []*Union       `xml:"union"`
	Constants           []*Constant    `xml:"constant"`
	Boxeds              []*Boxed       `xml:"boxed"`
	ErrorDomains        []*ErrorDomain `xml:"errordomain"`
	Aliases             []*Alias       `xml:"alias"`
}

type ErrorDomain struct {
	BaseInfo
	Name       string      `xml:"name,attr"`
	GetQuark   string      `xml:"get-quark,attr"`
	Codes      string      `xml:"codes,attr"`
	Annotation *Annotation `xml:"annotation"`
}

type Annotation struct {
	BaseInfo
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
	Name        string `xml:"name,attr"`
	Value       string `xml:"value,attr"`
	CName       string `xml:"type,attr"`
	CIdentifier string `xml:"identifier,attr"`
	Type        *Type  `xml:"type"`
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
	Name             string        `xml:"name,attr"`
	CSymbolPrefix    string        `xml:"symbol-prefix,attr"`
	CType            string        `xml:"type,attr"`
	Parent           string        `xml:"parent,attr"`
	GlibTypeName     string        `xml:"type-name,attr"`
	GlibGetType      string        `xml:"get-type,attr"`
	GlibTypeStruct   string        `xml:"type-struct,attr"`
	GlibFundamental  string        `xml:"fundamental,attr"`
	GlibGetValueFunc string        `xml:"get-value-func,attr"` // for Gst.MiniObject
	GlibSetValueFunc string        `xml:"set-value-func,attr"` // for Gst.MiniObject
	GlibRefFunc      string        `xml:"ref-func,attr"`       // for Gst.MiniObject
	GlibUnrefFunc    string        `xml:"unref-func,attr"`     // for Gst.MiniObject
	Disguised        string        `xml:"disguised,attr"`
	Foreign          string        `xml:"foreign,attr"`
	IsGTypeStruct    string        `xml:"is-gtype-struct-for,attr"`
	Abstract         string        `xml:"abstract,attr"`
	Prerequisite     *Prerequisite `xml:"prerequisite"` // for interface
	Implements       []*Implement  `xml:"implements"`
	Constructors     []*Function   `xml:"constructor"`
	VirtualMethods   []*Function   `xml:"virtual-method"`
	Methods          []*Function   `xml:"method"`
	Functions        []*Function   `xml:"function"`
	Properties       []*Property   `xml:"property"`
	Union            *Union        `xml:"union"` // for record
	Fields           []*Field      `xml:"field"`
	Signals          []*Function   `xml:"signal"`
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
	Writable          string `xml:"writable,attr"`
	TransferOwnership string `xml:"transfer-ownership,attr"`
	Construct         string `xml:"construct,attr"`
	ConstructOnly     string `xml:"construct-only,attr"`
	Readable          string `xml:"readable,attr"`
	Type              *Type  `xml:"type"`
	Array             *Array `xml:"array"`
}

type Field struct {
	BaseInfo
	Name     string    `xml:"name,attr"`
	Writable string    `xml:"writable,attr"`
	Private  string    `xml:"private,attr"`
	Readable string    `xml:"readable,attr"`
	Bits     string    `xml:"bits,attr"`
	Type     *Type     `xml:"type"`
	Array    *Array    `xml:"array"`
	Callback *Function `xml:"callback"`
}

type Union struct {
	BaseInfo
	Name          string      `xml:"name,attr"`
	CType         string      `xml:"type,attr"`
	GlibTypeName  string      `xml:"type-name,attr"`
	GlibGetType   string      `xml:"get-type,attr"`
	CSymbolPrefix string      `xml:"symbol-prefix,attr"`
	Fields        []*Field    `xml:"field"`
	Methods       []*Function `xml:"method"`
	Record        *Class      `xml:"record"`
	Functions     []*Function `xml:"function"`
	Constructors  []*Function `xml:"constructor"`
}

type Array struct {
	BaseInfo
	Name           string `xml:"name,attr"`
	ZeroTerminated string `xml:"zero-terminated,attr"`
	CType          string `xml:"type,attr"`
	Length         string `xml:"length,attr"`
	FixedSize      string `xml:"fixed-size,attr"`
	Type           *Type  `xml:"type"`
	Array          *Array `xml:"array"`
}

type Function struct {
	BaseInfo
	Name          string   `xml:"name,attr"`
	CType         string   `xml:"type,attr"` // for callback
	CIdentifier   string   `xml:"identifier,attr"`
	MovedTo       string   `xml:"moved-to,attr"` // in Gdk-3.0
	Throws        string   `xml:"throws,attr"`
	Shadows       string   `xml:"shadows,attr"`
	ShadowedBy    string   `xml:"shadowed-by,attr"`
	Invoker       string   `xml:"invoker,attr"`    // for virtual-method
	When          string   `xml:"when,attr"`       // for signal
	Action        string   `xml:"action,attr"`     // for signal
	NoHooks       string   `xml:"no-hooks,attr"`   // for signal
	NoRecurse     string   `xml:"no-recurse,attr"` // for signal
	Detailed      string   `xml:"detailed,attr"`   // for signal
	Return        *Param   `xml:"return-value"`
	InstanceParam *Param   `xml:"parameters>instance-parameter"`
	Params        []*Param `xml:"parameters>parameter"`
	GoName        string
	IsVarargs     bool
}

type Param struct {
	BaseInfo
	Name              string   `xml:"name,attr"`
	TransferOwnership string   `xml:"transfer-ownership,attr"`
	Direction         string   `xml:"direction,attr"`
	CallerAllocates   string   `xml:"caller-allocates,attr"`
	AllowNone         string   `xml:"allow-none,attr"`
	Scope             string   `xml:"scope,attr"`
	Destroy           string   `xml:"destroy,attr"`
	Closure           string   `xml:"closure,attr"`
	Skip              string   `xml:"skip,attr"`
	Array             *Array   `xml:"array"`
	Type              *Type    `xml:"type"`
	Varargs           *Varargs `xml:"varargs"`
	GoName            string
	MappedType        string
	IsArray           bool
	IsVoid            bool
	CType             string
	CTypeName         string
	GoType            string
	ElementCType      string // for array
	ElementCTypeName  string // for array
	ElementGoType     string // for array
	LenParamName      string // for array
	IsZeroTerminated  bool   // for array
	TypeSpec          string
	CgoParam          string
	CgoBeforeStmt     string
	CgoAfterStmt      string
}

type Varargs struct {
	BaseInfo
}

type Bitfield struct {
	BaseInfo
	Name         string      `xml:"name,attr"`
	CType        string      `xml:"type,attr"`
	GlibGetType  string      `xml:"get-type,attr"`
	GlibTypeName string      `xml:"type-name,attr"`
	Members      []*Member   `xml:"member"`
	Functions    []*Function `xml:"function"`
}

type Member struct {
	BaseInfo
	Name        string `xml:"name,attr"`
	Value       string `xml:"value,attr"`
	CIdentifier string `xml:"identifier,attr"`
	Nick        string `xml:"nick,attr"`
}

type Enum struct {
	BaseInfo
	Name            string      `xml:"name,attr"`
	CType           string      `xml:"type,attr"`
	GlibErrorDomain string      `xml:"error-domain,attr"`
	GlibTypeName    string      `xml:"type-name,attr"`
	GlibGetType     string      `xml:"get-type,attr"`
	Members         []*Member   `xml:"member"`
	Functions       []*Function `xml:"function"`
}

type Boxed struct {
	BaseInfo
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
	dumpNotHandled(reflect.ValueOf(repo))

	return &repo
}

func dumpNotHandled(v reflect.Value) {
	if !v.IsValid() {
		return
	}
	t := v.Type()
	kind := t.Kind()
	if kind == reflect.Ptr {
		dumpNotHandled(v.Elem())
	} else if kind == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			dumpNotHandled(v.Index(i))
		}
	} else if kind == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			if field.Name == "NotHandled" {
				for _, n := range v.Field(i).Interface().([]*NotHandled) {
					p("%s %s\n", n.XMLName.Local, strings.Repeat("=", 64))
					p("%s\n", n.Xml)
				}
			} else if field.Name == "NotHandledAttr" {
				for _, attr := range v.Field(i).Interface().([]xml.Attr) {
					p("%s = %s %s\n", attr.Name.Local, attr.Value, strings.Repeat("=", 64))
				}
			} else {
				dumpNotHandled(v.Field(i))
			}
		}
	}
}
