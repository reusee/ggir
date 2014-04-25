package main

import "strings"

func cTypeToGoType(t string) string {
	if t == "" {
		return ""
	}
	t = strings.Replace(t, "const", "", -1) // const and non-const pointer is compatible in Go 1.2
	t = strings.TrimSpace(t)
	t = strings.Replace(t, "volatile", "", -1) // no volatile
	t = strings.TrimSpace(t)
	t = strings.Replace(t, "long double", "double", -1) // no long double, FIXME
	t = strings.TrimSpace(t)
	pointerDepth := 0 // pointer depth
	for strings.HasSuffix(t, "*") {
		pointerDepth++
		t = t[:len(t)-1]
	}
	t = "C." + t                              // cgo type
	t = strings.Repeat("*", pointerDepth) + t // pointer
	return strings.TrimSpace(t)
}

func (self *Param) MapType() (ret string) {
	// skip G* types
	if strings.HasPrefix(self.GoType, "C.G") {
		ret = self.GoType
	}
	if strings.HasPrefix(self.GoType, "*C.G") {
		ret = self.GoType
	}
	if strings.HasPrefix(self.GoType, "**C.G") {
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
	case "false GoType C.guint TypeName guint":
		ret = "uint"
	case "false GoType C.gint32 TypeName gint32":
		ret = "int32"
	case "false GoType C.guint32 TypeName guint32":
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
	case "false GoType C.double TypeName gdouble",
		"false GoType C.gdouble TypeName gdouble",
		"false GoType C.double TypeName long double":
		ret = "float64"

	// string
	case "true GoType *C.gchar ElemType C.gchar ElemName utf8",
		"false GoType *C.gchar TypeName utf8",
		"false GoType *C.gchar TypeName filename",
		"false GoType *C.char TypeName utf8":
		ret = "string"

	// string slice
	case "true GoType **C.gchar ElemType *C.gchar ElemName utf8",
		"false GoType **C.gchar TypeName utf8":
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
		"true GoType *C.guint8 ElemType C.guint8 ElemName guint8",
		"true GoType *C.guchar ElemName guint8",
		"true GoType *C.gchar ElemName guint8":
		ret = "[]byte"

	// untyped pointer
	case "false GoType *C.void TypeName gpointer",
		"false GoType C.gpointer TypeName gpointer":
		ret = "unsafe.Pointer"

	// do not map
	case "false GoType *C.gunichar2 TypeName guint16",
		"false GoType **C.char TypeName utf8",
		"false GoType ***C.gchar TypeName utf8",
		"false GoType *C.gpointer TypeName gpointer",
		"false GoType *C.gsize TypeName gsize",
		"false GoType *C.glong TypeName glong",
		"false GoType *C.gunichar TypeName gunichar",
		"false GoType *C.guint TypeName guint",
		"false GoType *C.gint TypeName gint":
		ret = self.GoType

	}
	return
}
