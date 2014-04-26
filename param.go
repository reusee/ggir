package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"
)

func (self *Param) CollectInfo() {
	// name
	self.GoName = self.Name
	if isGoReservedWord(self.GoName) {
		self.GoName += "_"
	}

	// type info
	if self.Array != nil { // array type
		self.IsArray = true
		self.CType = self.Array.CType
		self.CTypeName = self.Array.Name // for spec only
		self.ElementCType = self.Array.Type.CType
		self.ElementCTypeName = self.Array.Type.Name // for spec only
		self.ElementGoType = cTypeToGoType(self.ElementCType)
	} else if self.Type != nil { // non-array type
		self.CType = self.Type.CType
		self.CTypeName = self.Type.Name // for spec only
	}
	self.GoType = cTypeToGoType(self.CType) // go type in function signature
	if self.GoType == "" {
		log.Fatal("ctype is not converted to gotype " + self.CType)
	}

	// skip void type
	if self.CTypeName == "none" || self.CType == "void" || self.GoType == "C.void" {
		self.IsVoid = true
		return
	}

	// check direction
	if self.Direction == "out" {
		if !strings.HasSuffix(self.CType, "*") || !strings.HasPrefix(self.GoType, "*") {
			log.Fatal("non-pointer out param " + self.CType)
		}
		// convert to non-pointer type
		self.CType = self.CType[:len(self.CType)-1]
		self.GoType = self.GoType[1:]
	} else if self.Direction == "" { // in param
		self.Direction = "in"
	}

	// type spec, for type mapping
	spec := fmt.Sprintf("%v", self.IsArray)
	if self.GoType != "" {
		spec += " GoType " + self.GoType
	}
	if self.CTypeName != "" {
		spec += " TypeName " + self.CTypeName
	}
	if self.ElementGoType != "" {
		spec += " ElemType " + self.ElementGoType
	}
	if self.ElementCTypeName != "" {
		spec += " ElemName " + self.ElementCTypeName
	}
	self.TypeSpec = spec

	// type mapping
	self.MappedType = self.MapType()

	// cgo param expression
	self.CgoParamExpr = self.GetCgoParamExpr()
}

func (self *Param) MapType() (ret string) {
	// skip complex types
	//FIXME map to registered boxed types
	if strings.HasPrefix(self.GoType, "C.") && unicode.IsUpper(rune(self.GoType[2])) {
		ret = self.GoType
	}
	if strings.HasPrefix(self.GoType, "*C.") && unicode.IsUpper(rune(self.GoType[3])) {
		ret = self.GoType
	}
	if strings.HasPrefix(self.GoType, "**C.") && unicode.IsUpper(rune(self.GoType[4])) {
		ret = self.GoType
	}

	switch self.TypeSpec {

	// bool
	case "false GoType C.gboolean TypeName gboolean":
		ret = "bool"

	// numeric
	case "false GoType C.int TypeName gint",
		"false GoType C.gint TypeName gint":
		ret = "int"
	case "false GoType C.guint TypeName guint",
		"false GoType C.uint TypeName guint":
		ret = "uint"
	case "false GoType C.gint8 TypeName gint8":
		ret = "int8"
	case "false GoType C.uint8_t TypeName guint8",
		"false GoType C.guint8 TypeName guint8":
		ret = "uint8"
	case "false GoType C.uint16_t TypeName guint16":
		ret = "uint16"
	case "false GoType C.gint32 TypeName gint32":
		ret = "int32"
	case "false GoType C.guint32 TypeName guint32",
		"false GoType C.ulong TypeName gulong":
		ret = "uint32"
	case "false GoType C.goffset TypeName gint64",
		"false GoType C.gsize TypeName gsize",
		"false GoType C.glong TypeName glong",
		"false GoType C.gulong TypeName gulong",
		"false GoType C.gint64 TypeName gint64",
		"false GoType C.gssize TypeName gssize":
		ret = "int64"
	case "false GoType C.guint64 TypeName guint64":
		ret = "uint64"
	case "false GoType C.gfloat TypeName gfloat":
		ret = "float32"
	case "false GoType C.double TypeName gdouble",
		"false GoType C.gdouble TypeName gdouble",
		"false GoType C.double TypeName long double":
		ret = "float64"

	// numeric slice
	case "true GoType *C.gint ElemType C.gint ElemName gint",
		"true GoType *C.int ElemType C.int ElemName gint":
		ret = "[]int"
	case "true GoType *C.gfloat ElemType C.gfloat ElemName gfloat",
		"true GoType *C.float ElemType C.float ElemName gfloat":
		ret = "[]float32"

	// string
	case "true GoType *C.gchar ElemType C.gchar ElemName utf8",
		"false GoType *C.gchar TypeName utf8",
		"false GoType *C.gchar TypeName filename",
		"false GoType *C.char TypeName utf8":
		ret = "string"

	// string slice
	case "true GoType **C.gchar ElemType *C.gchar ElemName utf8",
		"false GoType **C.gchar TypeName utf8",
		"true GoType **C.gchar ElemName utf8",
		"true GoType **C.gchar ElemName filename":
		ret = "[]string"

	// char
	case "false GoType C.char TypeName gchar",
		"false GoType C.gchar TypeName gchar":
		ret = "byte"
	case "false GoType C.gunichar TypeName gunichar":
		ret = "rune"

	// bytes
	case "true GoType *C.guchar ElemType C.guchar ElemName guint8",
		"false GoType *C.guchar TypeName guint8",
		"false GoType *C.uint8_t TypeName guint8",
		"true GoType *C.guint8 ElemType C.guint8 ElemName guint8",
		"true GoType *C.guchar ElemName guint8",
		"true GoType *C.gchar ElemName guint8":
		ret = "[]byte"

	// untyped pointer
	case "false GoType *C.void TypeName gpointer",
		"false GoType C.gpointer TypeName gpointer":
		ret = "unsafe.Pointer"

	// do not map pointer to basic types
	case "false GoType *C.gunichar2 TypeName guint16",
		"false GoType **C.char TypeName utf8",
		"false GoType ***C.gchar TypeName utf8",
		"false GoType *C.gpointer TypeName gpointer",
		"false GoType *C.gsize TypeName gsize",
		"false GoType *C.glong TypeName glong",
		"false GoType *C.gunichar TypeName gunichar",
		"false GoType *C.guint TypeName guint",
		"false GoType *C.gboolean TypeName gboolean",
		"false GoType **C.gchar TypeName filename",
		"false GoType *C.gint64 TypeName gint64",
		"false GoType *C.guint8 TypeName guint8",
		"false GoType *C.gint TypeName gint":
		ret = self.GoType

	}
	return
}

func (self *Param) GetCgoParamExpr() string {
	if self.GoType == self.MappedType {
		return self.GoName
	}
	//p("%s -> %s\n", self.MappedType, self.CType)
	return ""
}
