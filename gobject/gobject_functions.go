package gobject

/*
#include <glib-object.h>
#include <gobject/gvaluecollector.h>
#include <stdlib.h>
#cgo pkg-config: gobject-2.0
*/
import "C"

import (
	"errors"
	"reflect"
	"unsafe"
)

func init() {
	var _ unsafe.Pointer
	var _ reflect.SliceHeader
	_ = errors.New("")
}

/*
Provide a copy of a boxed structure @src_boxed which is of type @boxed_type.
*/
func BoxedCopy(boxed_type C.GType, src_boxed unsafe.Pointer) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_boxed_copy(boxed_type, (C.gconstpointer)(src_boxed))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Free the boxed structure @boxed which is of type @boxed_type.
*/
func BoxedFree(boxed_type C.GType, boxed unsafe.Pointer) {
	C.g_boxed_free(boxed_type, (C.gpointer)(boxed))
	return
}

/*
This function creates a new %G_TYPE_BOXED derived type id for a new
boxed type with name @name. Boxed type handling functions have to be
provided to copy and free opaque boxed structures of this type.
*/
func BoxedTypeRegisterStatic(name string, boxed_copy C.GBoxedCopyFunc, boxed_free C.GBoxedFreeFunc) (return__ C.GType) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	return__ = C.g_boxed_type_register_static(__cgo__name, boxed_copy, boxed_free)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

func CclosureMarshalBOOLEANBOXEDBOXED(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_BOOLEAN__BOXED_BOXED(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`gboolean (*callback) (gpointer instance, gint arg1, gpointer user_data)` where the #gint parameter
denotes a flags type.
*/
func CclosureMarshalBOOLEANFLAGS(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_BOOLEAN__FLAGS(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`gchar* (*callback) (gpointer instance, GObject *arg1, gpointer arg2, gpointer user_data)`.
*/
func CclosureMarshalSTRINGOBJECTPOINTER(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_STRING__OBJECT_POINTER(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, gboolean arg1, gpointer user_data)`.
*/
func CclosureMarshalVOIDBOOLEAN(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__BOOLEAN(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, GBoxed *arg1, gpointer user_data)`.
*/
func CclosureMarshalVOIDBOXED(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__BOXED(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, gchar arg1, gpointer user_data)`.
*/
func CclosureMarshalVOIDCHAR(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__CHAR(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, gdouble arg1, gpointer user_data)`.
*/
func CclosureMarshalVOIDDOUBLE(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__DOUBLE(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, gint arg1, gpointer user_data)` where the #gint parameter denotes an enumeration type..
*/
func CclosureMarshalVOIDENUM(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__ENUM(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, gint arg1, gpointer user_data)` where the #gint parameter denotes a flags type.
*/
func CclosureMarshalVOIDFLAGS(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__FLAGS(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, gfloat arg1, gpointer user_data)`.
*/
func CclosureMarshalVOIDFLOAT(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__FLOAT(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, gint arg1, gpointer user_data)`.
*/
func CclosureMarshalVOIDINT(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__INT(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, glong arg1, gpointer user_data)`.
*/
func CclosureMarshalVOIDLONG(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__LONG(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, GObject *arg1, gpointer user_data)`.
*/
func CclosureMarshalVOIDOBJECT(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__OBJECT(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, GParamSpec *arg1, gpointer user_data)`.
*/
func CclosureMarshalVOIDPARAM(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__PARAM(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, gpointer arg1, gpointer user_data)`.
*/
func CclosureMarshalVOIDPOINTER(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__POINTER(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, const gchar *arg1, gpointer user_data)`.
*/
func CclosureMarshalVOIDSTRING(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__STRING(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, guchar arg1, gpointer user_data)`.
*/
func CclosureMarshalVOIDUCHAR(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__UCHAR(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, guint arg1, gpointer user_data)`.
*/
func CclosureMarshalVOIDUINT(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__UINT(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, guint arg1, gpointer arg2, gpointer user_data)`.
*/
func CclosureMarshalVOIDUINTPOINTER(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__UINT_POINTER(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, gulong arg1, gpointer user_data)`.
*/
func CclosureMarshalVOIDULONG(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__ULONG(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, GVariant *arg1, gpointer user_data)`.
*/
func CclosureMarshalVOIDVARIANT(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__VARIANT(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A marshaller for a #GCClosure with a callback of type
`void (*callback) (gpointer instance, gpointer user_data)`.
*/
func CclosureMarshalVOIDVOID(closure *C.GClosure, return_value *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_VOID__VOID(closure, return_value, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
A generic marshaller function implemented via
[libffi](http://sourceware.org/libffi/).
*/
func CclosureMarshalGeneric(closure *C.GClosure, return_gvalue *C.GValue, n_param_values uint, param_values *C.GValue, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	C.g_cclosure_marshal_generic(closure, return_gvalue, C.guint(n_param_values), param_values, (C.gpointer)(invocation_hint), (C.gpointer)(marshal_data))
	return
}

/*
Creates a new closure which invokes @callback_func with @user_data as
the last parameter.
*/
func CclosureNew(callback_func C.GCallback, user_data unsafe.Pointer, destroy_data C.GClosureNotify) (return__ *C.GClosure) {
	return__ = C.g_cclosure_new(callback_func, (C.gpointer)(user_data), destroy_data)
	return
}

/*
A variant of g_cclosure_new() which uses @object as @user_data and
calls g_object_watch_closure() on @object and the created
closure. This function is useful when you have a callback closely
associated with a #GObject, and want the callback to no longer run
after the object is is freed.
*/
func CclosureNewObject(callback_func C.GCallback, object IsObject) (return__ *C.GClosure) {
	return__ = C.g_cclosure_new_object(callback_func, object.GetObjectPointer())
	return
}

/*
A variant of g_cclosure_new_swap() which uses @object as @user_data
and calls g_object_watch_closure() on @object and the created
closure. This function is useful when you have a callback closely
associated with a #GObject, and want the callback to no longer run
after the object is is freed.
*/
func CclosureNewObjectSwap(callback_func C.GCallback, object IsObject) (return__ *C.GClosure) {
	return__ = C.g_cclosure_new_object_swap(callback_func, object.GetObjectPointer())
	return
}

/*
Creates a new closure which invokes @callback_func with @user_data as
the first parameter.
*/
func CclosureNewSwap(callback_func C.GCallback, user_data unsafe.Pointer, destroy_data C.GClosureNotify) (return__ *C.GClosure) {
	return__ = C.g_cclosure_new_swap(callback_func, (C.gpointer)(user_data), destroy_data)
	return
}

// g_clear_object is not generated due to explicit ignore

/*
This function is meant to be called from the `complete_type_info`
function of a #GTypePlugin implementation, as in the following
example:

|[<!-- language="C" -->
static void
my_enum_complete_type_info (GTypePlugin     *plugin,
                            GType            g_type,
                            GTypeInfo       *info,
                            GTypeValueTable *value_table)
{
  static const GEnumValue values[] = {
    { MY_ENUM_FOO, "MY_ENUM_FOO", "foo" },
    { MY_ENUM_BAR, "MY_ENUM_BAR", "bar" },
    { 0, NULL, NULL }
  };

  g_enum_complete_type_info (type, info, values);
}
]|
*/
func EnumCompleteTypeInfo(g_enum_type C.GType, const_values *C.GEnumValue) (info C.GTypeInfo) {
	C.g_enum_complete_type_info(g_enum_type, &info, const_values)
	return
}

/*
Returns the #GEnumValue for a value.
*/
func EnumGetValue(enum_class *C.GEnumClass, value int) (return__ *C.GEnumValue) {
	return__ = C.g_enum_get_value(enum_class, C.gint(value))
	return
}

/*
Looks up a #GEnumValue by name.
*/
func EnumGetValueByName(enum_class *C.GEnumClass, name string) (return__ *C.GEnumValue) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	return__ = C.g_enum_get_value_by_name(enum_class, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Looks up a #GEnumValue by nickname.
*/
func EnumGetValueByNick(enum_class *C.GEnumClass, nick string) (return__ *C.GEnumValue) {
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	return__ = C.g_enum_get_value_by_nick(enum_class, __cgo__nick)
	C.free(unsafe.Pointer(__cgo__nick))
	return
}

/*
Registers a new static enumeration type with the name @name.

It is normally more convenient to let [glib-mkenums][glib-mkenums],
generate a my_enum_get_type() function from a usual C enumeration
definition  than to write one yourself using g_enum_register_static().
*/
func EnumRegisterStatic(name string, const_static_values *C.GEnumValue) (return__ C.GType) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	return__ = C.g_enum_register_static(__cgo__name, const_static_values)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
This function is meant to be called from the complete_type_info()
function of a #GTypePlugin implementation, see the example for
g_enum_complete_type_info() above.
*/
func FlagsCompleteTypeInfo(g_flags_type C.GType, const_values *C.GFlagsValue) (info C.GTypeInfo) {
	C.g_flags_complete_type_info(g_flags_type, &info, const_values)
	return
}

/*
Returns the first #GFlagsValue which is set in @value.
*/
func FlagsGetFirstValue(flags_class *C.GFlagsClass, value uint) (return__ *C.GFlagsValue) {
	return__ = C.g_flags_get_first_value(flags_class, C.guint(value))
	return
}

/*
Looks up a #GFlagsValue by name.
*/
func FlagsGetValueByName(flags_class *C.GFlagsClass, name string) (return__ *C.GFlagsValue) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	return__ = C.g_flags_get_value_by_name(flags_class, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Looks up a #GFlagsValue by nickname.
*/
func FlagsGetValueByNick(flags_class *C.GFlagsClass, nick string) (return__ *C.GFlagsValue) {
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	return__ = C.g_flags_get_value_by_nick(flags_class, __cgo__nick)
	C.free(unsafe.Pointer(__cgo__nick))
	return
}

/*
Registers a new static flags type with the name @name.

It is normally more convenient to let [glib-mkenums][glib-mkenums]
generate a my_flags_get_type() function from a usual C enumeration
definition than to write one yourself using g_flags_register_static().
*/
func FlagsRegisterStatic(name string, const_static_values *C.GFlagsValue) (return__ C.GType) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	return__ = C.g_flags_register_static(__cgo__name, const_static_values)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

func GtypeGetType() (return__ C.GType) {
	return__ = C.g_gtype_get_type()
	return
}

/*
Creates a new #GParamSpecBoolean instance specifying a %G_TYPE_BOOLEAN
property.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecBoolean(name string, nick string, blurb string, default_value bool, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	__cgo__default_value := C.gboolean(0)
	if default_value {
		__cgo__default_value = C.gboolean(1)
	}
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_boolean(__cgo__name, __cgo__nick, __cgo__blurb, __cgo__default_value, flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecBoxed instance specifying a %G_TYPE_BOXED
derived property.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecBoxed(name string, nick string, blurb string, boxed_type C.GType, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_boxed(__cgo__name, __cgo__nick, __cgo__blurb, boxed_type, flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecChar instance specifying a %G_TYPE_CHAR property.
*/
func GParamSpecChar(name string, nick string, blurb string, minimum int8, maximum int8, default_value int8, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_char(__cgo__name, __cgo__nick, __cgo__blurb, C.gint8(minimum), C.gint8(maximum), C.gint8(default_value), flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecDouble instance specifying a %G_TYPE_DOUBLE
property.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecDouble(name string, nick string, blurb string, minimum float64, maximum float64, default_value float64, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_double(__cgo__name, __cgo__nick, __cgo__blurb, C.gdouble(minimum), C.gdouble(maximum), C.gdouble(default_value), flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecEnum instance specifying a %G_TYPE_ENUM
property.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecEnum(name string, nick string, blurb string, enum_type C.GType, default_value int, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_enum(__cgo__name, __cgo__nick, __cgo__blurb, enum_type, C.gint(default_value), flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecFlags instance specifying a %G_TYPE_FLAGS
property.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecFlags(name string, nick string, blurb string, flags_type C.GType, default_value uint, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_flags(__cgo__name, __cgo__nick, __cgo__blurb, flags_type, C.guint(default_value), flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecFloat instance specifying a %G_TYPE_FLOAT property.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecFloat(name string, nick string, blurb string, minimum float32, maximum float32, default_value float32, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_float(__cgo__name, __cgo__nick, __cgo__blurb, C.gfloat(minimum), C.gfloat(maximum), C.gfloat(default_value), flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecGType instance specifying a
%G_TYPE_GTYPE property.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecGtype(name string, nick string, blurb string, is_a_type C.GType, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_gtype(__cgo__name, __cgo__nick, __cgo__blurb, is_a_type, flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecInt instance specifying a %G_TYPE_INT property.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecInt(name string, nick string, blurb string, minimum int, maximum int, default_value int, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_int(__cgo__name, __cgo__nick, __cgo__blurb, C.gint(minimum), C.gint(maximum), C.gint(default_value), flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecInt64 instance specifying a %G_TYPE_INT64 property.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecInt64(name string, nick string, blurb string, minimum int64, maximum int64, default_value int64, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_int64(__cgo__name, __cgo__nick, __cgo__blurb, C.gint64(minimum), C.gint64(maximum), C.gint64(default_value), flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecLong instance specifying a %G_TYPE_LONG property.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecLong(name string, nick string, blurb string, minimum int64, maximum int64, default_value int64, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_long(__cgo__name, __cgo__nick, __cgo__blurb, C.glong(minimum), C.glong(maximum), C.glong(default_value), flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecBoxed instance specifying a %G_TYPE_OBJECT
derived property.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecObject(name string, nick string, blurb string, object_type C.GType, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_object(__cgo__name, __cgo__nick, __cgo__blurb, object_type, flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new property of type #GParamSpecOverride. This is used
to direct operations to another paramspec, and will not be directly
useful unless you are implementing a new base type similar to GObject.
*/
func GParamSpecOverride(name string, overridden IsParamSpec) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_override(__cgo__name, overridden.GetParamSpecPointer())
	C.free(unsafe.Pointer(__cgo__name))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecParam instance specifying a %G_TYPE_PARAM
property.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecParam(name string, nick string, blurb string, param_type C.GType, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_param(__cgo__name, __cgo__nick, __cgo__blurb, param_type, flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecPointer instance specifying a pointer property.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecPointer(name string, nick string, blurb string, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_pointer(__cgo__name, __cgo__nick, __cgo__blurb, flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecPool.

If @type_prefixing is %TRUE, lookups in the newly created pool will
allow to specify the owner as a colon-separated prefix of the
property name, like "GtkContainer:border-width". This feature is
deprecated, so you should always set @type_prefixing to %FALSE.
*/
func GParamSpecPoolNew(type_prefixing bool) (return__ *C.GParamSpecPool) {
	__cgo__type_prefixing := C.gboolean(0)
	if type_prefixing {
		__cgo__type_prefixing = C.gboolean(1)
	}
	return__ = C.g_param_spec_pool_new(__cgo__type_prefixing)
	return
}

/*
Creates a new #GParamSpecString instance.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecString(name string, nick string, blurb string, default_value string, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	__cgo__default_value := (*C.gchar)(unsafe.Pointer(C.CString(default_value)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_string(__cgo__name, __cgo__nick, __cgo__blurb, __cgo__default_value, flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	C.free(unsafe.Pointer(__cgo__default_value))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecUChar instance specifying a %G_TYPE_UCHAR property.
*/
func GParamSpecUchar(name string, nick string, blurb string, minimum uint8, maximum uint8, default_value uint8, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_uchar(__cgo__name, __cgo__nick, __cgo__blurb, C.guint8(minimum), C.guint8(maximum), C.guint8(default_value), flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecUInt instance specifying a %G_TYPE_UINT property.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecUint(name string, nick string, blurb string, minimum uint, maximum uint, default_value uint, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_uint(__cgo__name, __cgo__nick, __cgo__blurb, C.guint(minimum), C.guint(maximum), C.guint(default_value), flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecUInt64 instance specifying a %G_TYPE_UINT64
property.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecUint64(name string, nick string, blurb string, minimum uint64, maximum uint64, default_value uint64, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_uint64(__cgo__name, __cgo__nick, __cgo__blurb, C.guint64(minimum), C.guint64(maximum), C.guint64(default_value), flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecULong instance specifying a %G_TYPE_ULONG
property.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecUlong(name string, nick string, blurb string, minimum uint64, maximum uint64, default_value uint64, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_ulong(__cgo__name, __cgo__nick, __cgo__blurb, C.gulong(minimum), C.gulong(maximum), C.gulong(default_value), flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecUnichar instance specifying a %G_TYPE_UINT
property. #GValue structures for this property can be accessed with
g_value_set_uint() and g_value_get_uint().

See g_param_spec_internal() for details on property names.
*/
func GParamSpecUnichar(name string, nick string, blurb string, default_value rune, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_unichar(__cgo__name, __cgo__nick, __cgo__blurb, C.gunichar(default_value), flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecValueArray instance specifying a
%G_TYPE_VALUE_ARRAY property. %G_TYPE_VALUE_ARRAY is a
%G_TYPE_BOXED type, as such, #GValue structures for this property
can be accessed with g_value_set_boxed() and g_value_get_boxed().

See g_param_spec_internal() for details on property names.
*/
func GParamSpecValueArray(name string, nick string, blurb string, element_spec IsParamSpec, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_value_array(__cgo__name, __cgo__nick, __cgo__blurb, element_spec.GetParamSpecPointer(), flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GParamSpecVariant instance specifying a #GVariant
property.

If @default_value is floating, it is consumed.

See g_param_spec_internal() for details on property names.
*/
func GParamSpecVariant(name string, nick string, blurb string, type_ *C.GVariantType, default_value *C.GVariant, flags C.GParamFlags) (return__ *ParamSpec) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__nick := (*C.gchar)(unsafe.Pointer(C.CString(nick)))
	__cgo__blurb := (*C.gchar)(unsafe.Pointer(C.CString(blurb)))
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_variant(__cgo__name, __cgo__nick, __cgo__blurb, type_, default_value, flags)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__nick))
	C.free(unsafe.Pointer(__cgo__blurb))
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Registers @name as the name of a new static type derived from
#G_TYPE_PARAM. The type system uses the information contained in
the #GParamSpecTypeInfo structure pointed to by @info to manage the
#GParamSpec type and its instances.
*/
func ParamTypeRegisterStatic(name string, pspec_info *C.GParamSpecTypeInfo) (return__ C.GType) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	return__ = C.g_param_type_register_static(__cgo__name, pspec_info)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Transforms @src_value into @dest_value if possible, and then
validates @dest_value, in order for it to conform to @pspec.  If
@strict_validation is %TRUE this function will only succeed if the
transformed @dest_value complied to @pspec without modifications.

See also g_value_type_transformable(), g_value_transform() and
g_param_value_validate().
*/
func ParamValueConvert(pspec IsParamSpec, src_value *C.GValue, dest_value *C.GValue, strict_validation bool) (return__ bool) {
	__cgo__strict_validation := C.gboolean(0)
	if strict_validation {
		__cgo__strict_validation = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_param_value_convert(pspec.GetParamSpecPointer(), src_value, dest_value, __cgo__strict_validation)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks whether @value contains the default value as specified in @pspec.
*/
func ParamValueDefaults(pspec IsParamSpec, value *C.GValue) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_param_value_defaults(pspec.GetParamSpecPointer(), value)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets @value to its default value as specified in @pspec.
*/
func ParamValueSetDefault(pspec IsParamSpec, value *C.GValue) {
	C.g_param_value_set_default(pspec.GetParamSpecPointer(), value)
	return
}

/*
Ensures that the contents of @value comply with the specifications
set out by @pspec. For example, a #GParamSpecInt might require
that integers stored in @value may not be smaller than -42 and not be
greater than +42. If @value contains an integer outside of this range,
it is modified accordingly, so the resulting value will fit into the
range -42 .. +42.
*/
func ParamValueValidate(pspec IsParamSpec, value *C.GValue) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_param_value_validate(pspec.GetParamSpecPointer(), value)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Compares @value1 with @value2 according to @pspec, and return -1, 0 or +1,
if @value1 is found to be less than, equal to or greater than @value2,
respectively.
*/
func ParamValuesCmp(pspec IsParamSpec, value1 *C.GValue, value2 *C.GValue) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_param_values_cmp(pspec.GetParamSpecPointer(), value1, value2)
	return__ = int(__cgo__return__)
	return
}

/*
Creates a new %G_TYPE_POINTER derived type id for a new
pointer type with name @name.
*/
func PointerTypeRegisterStatic(name string) (return__ C.GType) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	return__ = C.g_pointer_type_register_static(__cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
A predefined #GSignalAccumulator for signals intended to be used as a
hook for application code to provide a particular value.  Usually
only one such value is desired and multiple handlers for the same
signal don't make much sense (except for the case of the default
handler defined in the class structure, in which case you will
usually want the signal connection to override the class handler).

This accumulator will use the return value from the first signal
handler that is run as the return value for the signal and not run
any further handlers (ie: the first handler "wins").
*/
func SignalAccumulatorFirstWins(ihint *C.GSignalInvocationHint, return_accu *C.GValue, handler_return *C.GValue, dummy unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_signal_accumulator_first_wins(ihint, return_accu, handler_return, (C.gpointer)(dummy))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
A predefined #GSignalAccumulator for signals that return a
boolean values. The behavior that this accumulator gives is
that a return of %TRUE stops the signal emission: no further
callbacks will be invoked, while a return of %FALSE allows
the emission to continue. The idea here is that a %TRUE return
indicates that the callback handled the signal, and no further
handling is needed.
*/
func SignalAccumulatorTrueHandled(ihint *C.GSignalInvocationHint, return_accu *C.GValue, handler_return *C.GValue, dummy unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_signal_accumulator_true_handled(ihint, return_accu, handler_return, (C.gpointer)(dummy))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Adds an emission hook for a signal, which will get called for any emission
of that signal, independent of the instance. This is possible only
for signals which don't have #G_SIGNAL_NO_HOOKS flag set.
*/
func SignalAddEmissionHook(signal_id uint, detail C.GQuark, hook_func C.GSignalEmissionHook, hook_data unsafe.Pointer, data_destroy C.GDestroyNotify) (return__ uint64) {
	var __cgo__return__ C.gulong
	__cgo__return__ = C.g_signal_add_emission_hook(C.guint(signal_id), detail, hook_func, (C.gpointer)(hook_data), data_destroy)
	return__ = uint64(__cgo__return__)
	return
}

/*
Calls the original class closure of a signal. This function should only
be called from an overridden class closure; see
g_signal_override_class_closure() and
g_signal_override_class_handler().
*/
func SignalChainFromOverridden(instance_and_params *C.GValue, return_value *C.GValue) {
	C.g_signal_chain_from_overridden(instance_and_params, return_value)
	return
}

// g_signal_chain_from_overridden_handler is not generated due to varargs

/*
Connects a closure to a signal for a particular object.
*/
func SignalConnectClosure(instance IsObject, detailed_signal string, closure *C.GClosure, after bool) (return__ uint64) {
	__cgo__detailed_signal := (*C.gchar)(unsafe.Pointer(C.CString(detailed_signal)))
	__cgo__after := C.gboolean(0)
	if after {
		__cgo__after = C.gboolean(1)
	}
	var __cgo__return__ C.gulong
	__cgo__return__ = C.g_signal_connect_closure(C.gpointer(unsafe.Pointer(instance.GetObjectPointer())), __cgo__detailed_signal, closure, __cgo__after)
	C.free(unsafe.Pointer(__cgo__detailed_signal))
	return__ = uint64(__cgo__return__)
	return
}

/*
Connects a closure to a signal for a particular object.
*/
func SignalConnectClosureById(instance IsObject, signal_id uint, detail C.GQuark, closure *C.GClosure, after bool) (return__ uint64) {
	__cgo__after := C.gboolean(0)
	if after {
		__cgo__after = C.gboolean(1)
	}
	var __cgo__return__ C.gulong
	__cgo__return__ = C.g_signal_connect_closure_by_id(C.gpointer(unsafe.Pointer(instance.GetObjectPointer())), C.guint(signal_id), detail, closure, __cgo__after)
	return__ = uint64(__cgo__return__)
	return
}

/*
Connects a #GCallback function to a signal for a particular object. Similar
to g_signal_connect(), but allows to provide a #GClosureNotify for the data
which will be called when the signal handler is disconnected and no longer
used. Specify @connect_flags if you need `..._after()` or
`..._swapped()` variants of this function.
*/
func SignalConnectData(instance IsObject, detailed_signal string, c_handler C.GCallback, data unsafe.Pointer, destroy_data C.GClosureNotify, connect_flags C.GConnectFlags) (return__ uint64) {
	__cgo__detailed_signal := (*C.gchar)(unsafe.Pointer(C.CString(detailed_signal)))
	var __cgo__return__ C.gulong
	__cgo__return__ = C.g_signal_connect_data(C.gpointer(unsafe.Pointer(instance.GetObjectPointer())), __cgo__detailed_signal, c_handler, (C.gpointer)(data), destroy_data, connect_flags)
	C.free(unsafe.Pointer(__cgo__detailed_signal))
	return__ = uint64(__cgo__return__)
	return
}

/*
This is similar to g_signal_connect_data(), but uses a closure which
ensures that the @gobject stays alive during the call to @c_handler
by temporarily adding a reference count to @gobject.

When the @gobject is destroyed the signal handler will be automatically
disconnected.  Note that this is not currently threadsafe (ie:
emitting a signal while @gobject is being destroyed in another thread
is not safe).
*/
func SignalConnectObject(instance unsafe.Pointer, detailed_signal string, c_handler C.GCallback, gobject unsafe.Pointer, connect_flags C.GConnectFlags) (return__ uint64) {
	__cgo__detailed_signal := (*C.gchar)(unsafe.Pointer(C.CString(detailed_signal)))
	var __cgo__return__ C.gulong
	__cgo__return__ = C.g_signal_connect_object((C.gpointer)(instance), __cgo__detailed_signal, c_handler, (C.gpointer)(gobject), connect_flags)
	C.free(unsafe.Pointer(__cgo__detailed_signal))
	return__ = uint64(__cgo__return__)
	return
}

// g_signal_emit is not generated due to varargs

// g_signal_emit_by_name is not generated due to varargs

// g_signal_emit_valist is not generated due to varargs

/*
Emits a signal.

Note that g_signal_emitv() doesn't change @return_value if no handlers are
connected, in contrast to g_signal_emit() and g_signal_emit_valist().
*/
func SignalEmitv(instance_and_params *C.GValue, signal_id uint, detail C.GQuark, return_value *C.GValue) {
	C.g_signal_emitv(instance_and_params, C.guint(signal_id), detail, return_value)
	return
}

/*
Returns the invocation hint of the innermost signal emission of instance.
*/
func SignalGetInvocationHint(instance IsObject) (return__ *C.GSignalInvocationHint) {
	return__ = C.g_signal_get_invocation_hint(C.gpointer(unsafe.Pointer(instance.GetObjectPointer())))
	return
}

/*
Blocks a handler of an instance so it will not be called during any
signal emissions unless it is unblocked again. Thus "blocking" a
signal handler means to temporarily deactive it, a signal handler
has to be unblocked exactly the same amount of times it has been
blocked before to become active again.

The @handler_id has to be a valid signal handler id, connected to a
signal of @instance.
*/
func SignalHandlerBlock(instance IsObject, handler_id uint64) {
	C.g_signal_handler_block(C.gpointer(unsafe.Pointer(instance.GetObjectPointer())), C.gulong(handler_id))
	return
}

/*
Disconnects a handler from an instance so it will not be called during
any future or currently ongoing emissions of the signal it has been
connected to. The @handler_id becomes invalid and may be reused.

The @handler_id has to be a valid signal handler id, connected to a
signal of @instance.
*/
func SignalHandlerDisconnect(instance IsObject, handler_id uint64) {
	C.g_signal_handler_disconnect(C.gpointer(unsafe.Pointer(instance.GetObjectPointer())), C.gulong(handler_id))
	return
}

/*
Finds the first signal handler that matches certain selection criteria.
The criteria mask is passed as an OR-ed combination of #GSignalMatchType
flags, and the criteria values are passed as arguments.
The match @mask has to be non-0 for successful matches.
If no handler was found, 0 is returned.
*/
func SignalHandlerFind(instance IsObject, mask C.GSignalMatchType, signal_id uint, detail C.GQuark, closure *C.GClosure, func_ unsafe.Pointer, data unsafe.Pointer) (return__ uint64) {
	var __cgo__return__ C.gulong
	__cgo__return__ = C.g_signal_handler_find(C.gpointer(unsafe.Pointer(instance.GetObjectPointer())), mask, C.guint(signal_id), detail, closure, (C.gpointer)(func_), (C.gpointer)(data))
	return__ = uint64(__cgo__return__)
	return
}

/*
Returns whether @handler_id is the id of a handler connected to @instance.
*/
func SignalHandlerIsConnected(instance IsObject, handler_id uint64) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_signal_handler_is_connected(C.gpointer(unsafe.Pointer(instance.GetObjectPointer())), C.gulong(handler_id))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Undoes the effect of a previous g_signal_handler_block() call.  A
blocked handler is skipped during signal emissions and will not be
invoked, unblocking it (for exactly the amount of times it has been
blocked before) reverts its "blocked" state, so the handler will be
recognized by the signal system and is called upon future or
currently ongoing signal emissions (since the order in which
handlers are called during signal emissions is deterministic,
whether the unblocked handler in question is called as part of a
currently ongoing emission depends on how far that emission has
proceeded yet).

The @handler_id has to be a valid id of a signal handler that is
connected to a signal of @instance and is currently blocked.
*/
func SignalHandlerUnblock(instance IsObject, handler_id uint64) {
	C.g_signal_handler_unblock(C.gpointer(unsafe.Pointer(instance.GetObjectPointer())), C.gulong(handler_id))
	return
}

/*
Blocks all handlers on an instance that match a certain selection criteria.
The criteria mask is passed as an OR-ed combination of #GSignalMatchType
flags, and the criteria values are passed as arguments.
Passing at least one of the %G_SIGNAL_MATCH_CLOSURE, %G_SIGNAL_MATCH_FUNC
or %G_SIGNAL_MATCH_DATA match flags is required for successful matches.
If no handlers were found, 0 is returned, the number of blocked handlers
otherwise.
*/
func SignalHandlersBlockMatched(instance IsObject, mask C.GSignalMatchType, signal_id uint, detail C.GQuark, closure *C.GClosure, func_ unsafe.Pointer, data unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_signal_handlers_block_matched(C.gpointer(unsafe.Pointer(instance.GetObjectPointer())), mask, C.guint(signal_id), detail, closure, (C.gpointer)(func_), (C.gpointer)(data))
	return__ = uint(__cgo__return__)
	return
}

func SignalHandlersDestroy(instance IsObject) {
	C.g_signal_handlers_destroy(C.gpointer(unsafe.Pointer(instance.GetObjectPointer())))
	return
}

/*
Disconnects all handlers on an instance that match a certain
selection criteria. The criteria mask is passed as an OR-ed
combination of #GSignalMatchType flags, and the criteria values are
passed as arguments.  Passing at least one of the
%G_SIGNAL_MATCH_CLOSURE, %G_SIGNAL_MATCH_FUNC or
%G_SIGNAL_MATCH_DATA match flags is required for successful
matches.  If no handlers were found, 0 is returned, the number of
disconnected handlers otherwise.
*/
func SignalHandlersDisconnectMatched(instance IsObject, mask C.GSignalMatchType, signal_id uint, detail C.GQuark, closure *C.GClosure, func_ unsafe.Pointer, data unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_signal_handlers_disconnect_matched(C.gpointer(unsafe.Pointer(instance.GetObjectPointer())), mask, C.guint(signal_id), detail, closure, (C.gpointer)(func_), (C.gpointer)(data))
	return__ = uint(__cgo__return__)
	return
}

/*
Unblocks all handlers on an instance that match a certain selection
criteria. The criteria mask is passed as an OR-ed combination of
#GSignalMatchType flags, and the criteria values are passed as arguments.
Passing at least one of the %G_SIGNAL_MATCH_CLOSURE, %G_SIGNAL_MATCH_FUNC
or %G_SIGNAL_MATCH_DATA match flags is required for successful matches.
If no handlers were found, 0 is returned, the number of unblocked handlers
otherwise. The match criteria should not apply to any handlers that are
not currently blocked.
*/
func SignalHandlersUnblockMatched(instance IsObject, mask C.GSignalMatchType, signal_id uint, detail C.GQuark, closure *C.GClosure, func_ unsafe.Pointer, data unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_signal_handlers_unblock_matched(C.gpointer(unsafe.Pointer(instance.GetObjectPointer())), mask, C.guint(signal_id), detail, closure, (C.gpointer)(func_), (C.gpointer)(data))
	return__ = uint(__cgo__return__)
	return
}

/*
Returns whether there are any handlers connected to @instance for the
given signal id and detail.

One example of when you might use this is when the arguments to the
signal are difficult to compute. A class implementor may opt to not
emit the signal if no one is attached anyway, thus saving the cost
of building the arguments.
*/
func SignalHasHandlerPending(instance IsObject, signal_id uint, detail C.GQuark, may_be_blocked bool) (return__ bool) {
	__cgo__may_be_blocked := C.gboolean(0)
	if may_be_blocked {
		__cgo__may_be_blocked = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_signal_has_handler_pending(C.gpointer(unsafe.Pointer(instance.GetObjectPointer())), C.guint(signal_id), detail, __cgo__may_be_blocked)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Lists the signals by id that a certain instance or interface type
created. Further information about the signals can be acquired through
g_signal_query().
*/
func SignalListIds(itype C.GType) (n_ids uint, return__ []uint) {
	var __cgo__n_ids C.guint
	var __cgo__return__ *C.guint
	__cgo__return__ = C.g_signal_list_ids(itype, &__cgo__n_ids)
	n_ids = uint(__cgo__n_ids)
	defer func() {
		__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&return__))
		__header__return__.Len = int(n_ids)
		__header__return__.Cap = int(n_ids)
		__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	}()
	return
}

/*
Given the name of the signal and the type of object it connects to, gets
the signal's identifying integer. Emitting the signal by number is
somewhat faster than using the name each time.

Also tries the ancestors of the given type.

See g_signal_new() for details on allowed signal names.
*/
func SignalLookup(name string, itype C.GType) (return__ uint) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_signal_lookup(__cgo__name, itype)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = uint(__cgo__return__)
	return
}

/*
Given the signal's identifier, finds its name.

Two different signals may have the same name, if they have differing types.
*/
func SignalName(signal_id uint) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_signal_name(C.guint(signal_id))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

// g_signal_new is not generated due to varargs

// g_signal_new_class_handler is not generated due to varargs

// g_signal_new_valist is not generated due to varargs

/*
Creates a new signal. (This is usually done in the class initializer.)

See g_signal_new() for details on allowed signal names.

If c_marshaller is %NULL @g_cclosure_marshal_generic will be used as
the marshaller for this signal.
*/
func SignalNewv(signal_name string, itype C.GType, signal_flags C.GSignalFlags, class_closure *C.GClosure, accumulator C.GSignalAccumulator, accu_data unsafe.Pointer, c_marshaller C.GSignalCMarshaller, return_type C.GType, n_params uint, param_types *C.GType) (return__ uint) {
	__cgo__signal_name := (*C.gchar)(unsafe.Pointer(C.CString(signal_name)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_signal_newv(__cgo__signal_name, itype, signal_flags, class_closure, accumulator, (C.gpointer)(accu_data), c_marshaller, return_type, C.guint(n_params), param_types)
	C.free(unsafe.Pointer(__cgo__signal_name))
	return__ = uint(__cgo__return__)
	return
}

/*
Overrides the class closure (i.e. the default handler) for the given signal
for emissions on instances of @instance_type. @instance_type must be derived
from the type to which the signal belongs.

See g_signal_chain_from_overridden() and
g_signal_chain_from_overridden_handler() for how to chain up to the
parent class closure from inside the overridden one.
*/
func SignalOverrideClassClosure(signal_id uint, instance_type C.GType, class_closure *C.GClosure) {
	C.g_signal_override_class_closure(C.guint(signal_id), instance_type, class_closure)
	return
}

/*
Overrides the class closure (i.e. the default handler) for the
given signal for emissions on instances of @instance_type with
callback @class_handler. @instance_type must be derived from the
type to which the signal belongs.

See g_signal_chain_from_overridden() and
g_signal_chain_from_overridden_handler() for how to chain up to the
parent class closure from inside the overridden one.
*/
func SignalOverrideClassHandler(signal_name string, instance_type C.GType, class_handler C.GCallback) {
	__cgo__signal_name := (*C.gchar)(unsafe.Pointer(C.CString(signal_name)))
	C.g_signal_override_class_handler(__cgo__signal_name, instance_type, class_handler)
	C.free(unsafe.Pointer(__cgo__signal_name))
	return
}

/*
Internal function to parse a signal name into its @signal_id
and @detail quark.
*/
func SignalParseName(detailed_signal string, itype C.GType, force_detail_quark bool) (signal_id_p uint, detail_p C.GQuark, return__ bool) {
	__cgo__detailed_signal := (*C.gchar)(unsafe.Pointer(C.CString(detailed_signal)))
	var __cgo__signal_id_p C.guint
	__cgo__force_detail_quark := C.gboolean(0)
	if force_detail_quark {
		__cgo__force_detail_quark = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_signal_parse_name(__cgo__detailed_signal, itype, &__cgo__signal_id_p, &detail_p, __cgo__force_detail_quark)
	C.free(unsafe.Pointer(__cgo__detailed_signal))
	signal_id_p = uint(__cgo__signal_id_p)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Queries the signal system for in-depth information about a
specific signal. This function will fill in a user-provided
structure to hold signal-specific information. If an invalid
signal id is passed in, the @signal_id member of the #GSignalQuery
is 0. All members filled into the #GSignalQuery structure should
be considered constant and have to be left untouched.
*/
func GSignalQuery(signal_id uint) (query C.GSignalQuery) {
	C.g_signal_query(C.guint(signal_id), &query)
	return
}

/*
Deletes an emission hook.
*/
func SignalRemoveEmissionHook(signal_id uint, hook_id uint64) {
	C.g_signal_remove_emission_hook(C.guint(signal_id), C.gulong(hook_id))
	return
}

func SignalSetVaMarshaller(signal_id uint, instance_type C.GType, va_marshaller C.GSignalCVaMarshaller) {
	C.g_signal_set_va_marshaller(C.guint(signal_id), instance_type, va_marshaller)
	return
}

/*
Stops a signal's current emission.

This will prevent the default method from running, if the signal was
%G_SIGNAL_RUN_LAST and you connected normally (i.e. without the "after"
flag).

Prints a warning if used on a signal which isn't being emitted.
*/
func SignalStopEmission(instance IsObject, signal_id uint, detail C.GQuark) {
	C.g_signal_stop_emission(C.gpointer(unsafe.Pointer(instance.GetObjectPointer())), C.guint(signal_id), detail)
	return
}

/*
Stops a signal's current emission.

This is just like g_signal_stop_emission() except it will look up the
signal id for you.
*/
func SignalStopEmissionByName(instance IsObject, detailed_signal string) {
	__cgo__detailed_signal := (*C.gchar)(unsafe.Pointer(C.CString(detailed_signal)))
	C.g_signal_stop_emission_by_name(C.gpointer(unsafe.Pointer(instance.GetObjectPointer())), __cgo__detailed_signal)
	C.free(unsafe.Pointer(__cgo__detailed_signal))
	return
}

/*
Creates a new closure which invokes the function found at the offset
@struct_offset in the class structure of the interface or classed type
identified by @itype.
*/
func SignalTypeCclosureNew(itype C.GType, struct_offset uint) (return__ *C.GClosure) {
	return__ = C.g_signal_type_cclosure_new(itype, C.guint(struct_offset))
	return
}

// g_source_set_closure is not generated due to explicit ignore

/*
Sets a dummy callback for @source. The callback will do nothing, and
if the source expects a #gboolean return value, it will return %TRUE.
(If the source expects any other type of return value, it will return
a 0/%NULL value; whatever g_value_init() initializes a #GValue to for
that type.)

If the source is not one of the standard GLib types, the
@closure_callback and @closure_marshal fields of the #GSourceFuncs
structure must have been filled in with pointers to appropriate
functions.
*/
func SourceSetDummyCallback(source *C.GSource) {
	C.g_source_set_dummy_callback(source)
	return
}

/*
Return a newly allocated string, which describes the contents of a
#GValue.  The main purpose of this function is to describe #GValue
contents for debugging output, the way in which the contents are
described may change between different GLib versions.
*/
func StrdupValueContents(value *C.GValue) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strdup_value_contents(value)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Adds a #GTypeClassCacheFunc to be called before the reference count of a
class goes from one to zero. This can be used to prevent premature class
destruction. All installed #GTypeClassCacheFunc functions will be chained
until one of them returns %TRUE. The functions have to check the class id
passed in to figure whether they actually want to cache the class of this
type, since all classes are routed through the same #GTypeClassCacheFunc
chain.
*/
func TypeAddClassCacheFunc(cache_data unsafe.Pointer, cache_func C.GTypeClassCacheFunc) {
	C.g_type_add_class_cache_func((C.gpointer)(cache_data), cache_func)
	return
}

/*
Registers a private class structure for a classed type;
when the class is allocated, the private structures for
the class and all of its parent types are allocated
sequentially in the same memory block as the public
structures. This function should be called in the
type's get_type() function after the type is registered.
The private structure can be retrieved using the
G_TYPE_CLASS_GET_PRIVATE() macro.
*/
func TypeAddClassPrivate(class_type C.GType, private_size int64) {
	C.g_type_add_class_private(class_type, C.gsize(private_size))
	return
}

func TypeAddInstancePrivate(class_type C.GType, private_size int64) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_type_add_instance_private(class_type, C.gsize(private_size))
	return__ = int(__cgo__return__)
	return
}

/*
Adds a function to be called after an interface vtable is
initialized for any class (i.e. after the @interface_init
member of #GInterfaceInfo has been called).

This function is useful when you want to check an invariant
that depends on the interfaces of a class. For instance, the
implementation of #GObject uses this facility to check that an
object implements all of the properties that are defined on its
interfaces.
*/
func TypeAddInterfaceCheck(check_data unsafe.Pointer, check_func C.GTypeInterfaceCheckFunc) {
	C.g_type_add_interface_check((C.gpointer)(check_data), check_func)
	return
}

/*
Adds the dynamic @interface_type to @instantiable_type. The information
contained in the #GTypePlugin structure pointed to by @plugin
is used to manage the relationship.
*/
func TypeAddInterfaceDynamic(instance_type C.GType, interface_type C.GType, plugin *C.GTypePlugin) {
	C.g_type_add_interface_dynamic(instance_type, interface_type, plugin)
	return
}

/*
Adds the static @interface_type to @instantiable_type.
The information contained in the #GInterfaceInfo structure
pointed to by @info is used to manage the relationship.
*/
func TypeAddInterfaceStatic(instance_type C.GType, interface_type C.GType, info *C.GInterfaceInfo) {
	C.g_type_add_interface_static(instance_type, interface_type, info)
	return
}

func TypeCheckClassCast(g_class *C.GTypeClass, is_a_type C.GType) (return__ *C.GTypeClass) {
	return__ = C.g_type_check_class_cast(g_class, is_a_type)
	return
}

func TypeCheckClassIsA(g_class *C.GTypeClass, is_a_type C.GType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_type_check_class_is_a(g_class, is_a_type)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Private helper function to aid implementation of the
G_TYPE_CHECK_INSTANCE() macro.
*/
func TypeCheckInstance(instance *C.GTypeInstance) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_type_check_instance(instance)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

func TypeCheckInstanceCast(instance *C.GTypeInstance, iface_type C.GType) (return__ *C.GTypeInstance) {
	return__ = C.g_type_check_instance_cast(instance, iface_type)
	return
}

func TypeCheckInstanceIsA(instance *C.GTypeInstance, iface_type C.GType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_type_check_instance_is_a(instance, iface_type)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

func TypeCheckIsValueType(type_ C.GType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_type_check_is_value_type(type_)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

func TypeCheckValue(value *C.GValue) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_type_check_value(value)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

func TypeCheckValueHolds(value *C.GValue, type_ C.GType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_type_check_value_holds(value, type_)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Return a newly allocated and 0-terminated array of type IDs, listing
the child types of @type.
*/
func TypeChildren(type_ C.GType) (n_children uint, return__ *C.GType) {
	var __cgo__n_children C.guint
	return__ = C.g_type_children(type_, &__cgo__n_children)
	n_children = uint(__cgo__n_children)
	return
}

/*
Registers a private structure for an instantiatable type.

When an object is allocated, the private structures for
the type and all of its parent types are allocated
sequentially in the same memory block as the public
structures.

Note that the accumulated size of the private structures of
a type and all its parent types cannot exceed 64 KiB.

This function should be called in the type's class_init() function.
The private structure can be retrieved using the
G_TYPE_INSTANCE_GET_PRIVATE() macro.

The following example shows attaching a private structure
MyObjectPrivate to an object MyObject defined in the standard
GObject fashion in the type's class_init() function.

Note the use of a structure member "priv" to avoid the overhead
of repeatedly calling MY_OBJECT_GET_PRIVATE().

|[<!-- language="C" -->
typedef struct _MyObject        MyObject;
typedef struct _MyObjectPrivate MyObjectPrivate;

struct _MyObject {
 GObject parent;

 MyObjectPrivate *priv;
};

struct _MyObjectPrivate {
  int some_field;
};

static void
my_object_class_init (MyObjectClass *klass)
{
  g_type_class_add_private (klass, sizeof (MyObjectPrivate));
}

static void
my_object_init (MyObject *my_object)
{
  my_object->priv = G_TYPE_INSTANCE_GET_PRIVATE (my_object,
                                                 MY_TYPE_OBJECT,
                                                 MyObjectPrivate);
}

static int
my_object_get_some_field (MyObject *my_object)
{
  MyObjectPrivate *priv;

  g_return_val_if_fail (MY_IS_OBJECT (my_object), 0);

  priv = my_object->priv;

  return priv->some_field;
}
]|
*/
func TypeClassAddPrivate(g_class unsafe.Pointer, private_size int64) {
	C.g_type_class_add_private((C.gpointer)(g_class), C.gsize(private_size))
	return
}

func TypeClassAdjustPrivateOffset(g_class unsafe.Pointer, private_size_or_offset *C.gint) {
	C.g_type_class_adjust_private_offset((C.gpointer)(g_class), private_size_or_offset)
	return
}

/*
Gets the offset of the private data for instances of @g_class.

This is how many bytes you should add to the instance pointer of a
class in order to get the private data for the type represented by
@g_class.

You can only call this function after you have registered a private
data area for @g_class using g_type_class_add_private().
*/
func TypeClassGetInstancePrivateOffset(g_class unsafe.Pointer) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_type_class_get_instance_private_offset((C.gpointer)(g_class))
	return__ = int(__cgo__return__)
	return
}

/*
This function is essentially the same as g_type_class_ref(),
except that the classes reference count isn't incremented.
As a consequence, this function may return %NULL if the class
of the type passed in does not currently exist (hasn't been
referenced before).
*/
func TypeClassPeek(type_ C.GType) (return__ *C.GTypeClass) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_type_class_peek(type_)
	return__ = (*C.GTypeClass)(unsafe.Pointer(__cgo__return__))
	return
}

/*
A more efficient version of g_type_class_peek() which works only for
static types.
*/
func TypeClassPeekStatic(type_ C.GType) (return__ *C.GTypeClass) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_type_class_peek_static(type_)
	return__ = (*C.GTypeClass)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Increments the reference count of the class structure belonging to
@type. This function will demand-create the class if it doesn't
exist already.
*/
func TypeClassRef(type_ C.GType) (return__ *C.GTypeClass) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_type_class_ref(type_)
	return__ = (*C.GTypeClass)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Creates and initializes an instance of @type if @type is valid and
can be instantiated. The type system only performs basic allocation
and structure setups for instances: actual instance creation should
happen through functions supplied by the type's fundamental type
implementation.  So use of g_type_create_instance() is reserved for
implementators of fundamental types only. E.g. instances of the
#GObject hierarchy should be created via g_object_new() and never
directly through g_type_create_instance() which doesn't handle things
like singleton objects or object construction.

Note: Do not use this function, unless you're implementing a
fundamental type. Also language bindings should not use this
function, but g_object_new() instead.
*/
func TypeCreateInstance(type_ C.GType) (return__ *C.GTypeInstance) {
	return__ = C.g_type_create_instance(type_)
	return
}

/*
If the interface type @g_type is currently in use, returns its
default interface vtable.
*/
func TypeDefaultInterfacePeek(g_type C.GType) (return__ *C.GTypeInterface) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_type_default_interface_peek(g_type)
	return__ = (*C.GTypeInterface)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Increments the reference count for the interface type @g_type,
and returns the default interface vtable for the type.

If the type is not currently in use, then the default vtable
for the type will be created and initalized by calling
the base interface init and default vtable init functions for
the type (the @base_init and @class_init members of #GTypeInfo).
Calling g_type_default_interface_ref() is useful when you
want to make sure that signals and properties for an interface
have been installed.
*/
func TypeDefaultInterfaceRef(g_type C.GType) (return__ *C.GTypeInterface) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_type_default_interface_ref(g_type)
	return__ = (*C.GTypeInterface)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Decrements the reference count for the type corresponding to the
interface default vtable @g_iface. If the type is dynamic, then
when no one is using the interface and all references have
been released, the finalize function for the interface's default
vtable (the @class_finalize member of #GTypeInfo) will be called.
*/
func TypeDefaultInterfaceUnref(g_iface *C.GTypeInterface) {
	C.g_type_default_interface_unref(C.gpointer(unsafe.Pointer(g_iface)))
	return
}

/*
Returns the length of the ancestry of the passed in type. This
includes the type itself, so that e.g. a fundamental type has depth 1.
*/
func TypeDepth(type_ C.GType) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_type_depth(type_)
	return__ = uint(__cgo__return__)
	return
}

/*
Ensures that the indicated @type has been registered with the
type system, and its _class_init() method has been run.

In theory, simply calling the type's _get_type() method (or using
the corresponding macro) is supposed take care of this. However,
_get_type() methods are often marked %G_GNUC_CONST for performance
reasons, even though this is technically incorrect (since
%G_GNUC_CONST requires that the function not have side effects,
which _get_type() methods do on the first call). As a result, if
you write a bare call to a _get_type() macro, it may get optimized
out by the compiler. Using g_type_ensure() guarantees that the
type's _get_type() method is called.
*/
func TypeEnsure(type_ C.GType) {
	C.g_type_ensure(type_)
	return
}

/*
Frees an instance of a type, returning it to the instance pool for
the type, if there is one.

Like g_type_create_instance(), this function is reserved for
implementors of fundamental types.
*/
func TypeFreeInstance(instance *C.GTypeInstance) {
	C.g_type_free_instance(instance)
	return
}

/*
Lookup the type ID from a given type name, returning 0 if no type
has been registered under this name (this is the preferred method
to find out by name whether a specific type has been registered
yet).
*/
func TypeFromName(name string) (return__ C.GType) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	return__ = C.g_type_from_name(__cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Internal function, used to extract the fundamental type ID portion.
Use G_TYPE_FUNDAMENTAL() instead.
*/
func TypeFundamental(type_id C.GType) (return__ C.GType) {
	return__ = C.g_type_fundamental(type_id)
	return
}

/*
Returns the next free fundamental type id which can be used to
register a new fundamental type with g_type_register_fundamental().
The returned type ID represents the highest currently registered
fundamental type identifier.
*/
func TypeFundamentalNext() (return__ C.GType) {
	return__ = C.g_type_fundamental_next()
	return
}

/*
Returns the #GTypePlugin structure for @type.
*/
func TypeGetPlugin(type_ C.GType) (return__ *C.GTypePlugin) {
	return__ = C.g_type_get_plugin(type_)
	return
}

/*
Obtains data which has previously been attached to @type
with g_type_set_qdata().

Note that this does not take subtyping into account; data
attached to one type with g_type_set_qdata() cannot
be retrieved from a subtype using g_type_get_qdata().
*/
func TypeGetQdata(type_ C.GType, quark C.GQuark) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_type_get_qdata(type_, quark)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Returns an opaque serial number that represents the state of the set
of registered types. Any time a type is registered this serial changes,
which means you can cache information based on type lookups (such as
g_type_from_name()) and know if the cache is still valid at a later
time by comparing the current serial with the one at the type lookup.
*/
func TypeGetTypeRegistrationSerial() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_type_get_type_registration_serial()
	return__ = uint(__cgo__return__)
	return
}

// g_type_init is not generated due to deprecation attr

// g_type_init_with_debug_flags is not generated due to deprecation attr

/*
Adds @prerequisite_type to the list of prerequisites of @interface_type.
This means that any type implementing @interface_type must also implement
@prerequisite_type. Prerequisites can be thought of as an alternative to
interface derivation (which GType doesn't support). An interface can have
at most one instantiatable prerequisite type.
*/
func TypeInterfaceAddPrerequisite(interface_type C.GType, prerequisite_type C.GType) {
	C.g_type_interface_add_prerequisite(interface_type, prerequisite_type)
	return
}

/*
Returns the #GTypePlugin structure for the dynamic interface
@interface_type which has been added to @instance_type, or %NULL
if @interface_type has not been added to @instance_type or does
not have a #GTypePlugin structure. See g_type_add_interface_dynamic().
*/
func TypeInterfaceGetPlugin(instance_type C.GType, interface_type C.GType) (return__ *C.GTypePlugin) {
	return__ = C.g_type_interface_get_plugin(instance_type, interface_type)
	return
}

/*
Returns the #GTypeInterface structure of an interface to which the
passed in class conforms.
*/
func TypeInterfacePeek(instance_class *C.GTypeClass, iface_type C.GType) (return__ *C.GTypeInterface) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_type_interface_peek(C.gpointer(unsafe.Pointer(instance_class)), iface_type)
	return__ = (*C.GTypeInterface)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Returns the prerequisites of an interfaces type.
*/
func TypeInterfacePrerequisites(interface_type C.GType) (n_prerequisites uint, return__ *C.GType) {
	var __cgo__n_prerequisites C.guint
	return__ = C.g_type_interface_prerequisites(interface_type, &__cgo__n_prerequisites)
	n_prerequisites = uint(__cgo__n_prerequisites)
	return
}

/*
Return a newly allocated and 0-terminated array of type IDs, listing
the interface types that @type conforms to.
*/
func TypeInterfaces(type_ C.GType) (n_interfaces uint, return__ *C.GType) {
	var __cgo__n_interfaces C.guint
	return__ = C.g_type_interfaces(type_, &__cgo__n_interfaces)
	n_interfaces = uint(__cgo__n_interfaces)
	return
}

/*
If @is_a_type is a derivable type, check whether @type is a
descendant of @is_a_type. If @is_a_type is an interface, check
whether @type conforms to it.
*/
func TypeIsA(type_ C.GType, is_a_type C.GType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_type_is_a(type_, is_a_type)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the unique name that is assigned to a type ID.  Note that this
function (like all other GType API) cannot cope with invalid type
IDs. %G_TYPE_INVALID may be passed to this function, as may be any
other validly registered type ID, but randomized type IDs should
not be passed in and will most likely lead to a crash.
*/
func TypeName(type_ C.GType) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_type_name(type_)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

func TypeNameFromClass(g_class *C.GTypeClass) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_type_name_from_class(g_class)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

func TypeNameFromInstance(instance *C.GTypeInstance) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_type_name_from_instance(instance)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Given a @leaf_type and a @root_type which is contained in its
anchestry, return the type that @root_type is the immediate parent
of. In other words, this function determines the type that is
derived directly from @root_type which is also a base class of
@leaf_type.  Given a root type and a leaf type, this function can
be used to determine the types and order in which the leaf type is
descended from the root type.
*/
func TypeNextBase(leaf_type C.GType, root_type C.GType) (return__ C.GType) {
	return__ = C.g_type_next_base(leaf_type, root_type)
	return
}

/*
Return the direct parent type of the passed in type. If the passed
in type has no parent, i.e. is a fundamental type, 0 is returned.
*/
func TypeParent(type_ C.GType) (return__ C.GType) {
	return__ = C.g_type_parent(type_)
	return
}

/*
Get the corresponding quark of the type IDs name.
*/
func TypeQname(type_ C.GType) (return__ C.GQuark) {
	return__ = C.g_type_qname(type_)
	return
}

/*
Queries the type system for information about a specific type.
This function will fill in a user-provided structure to hold
type-specific information. If an invalid #GType is passed in, the
@type member of the #GTypeQuery is 0. All members filled into the
#GTypeQuery structure should be considered constant and have to be
left untouched.
*/
func GTypeQuery(type_ C.GType) (query C.GTypeQuery) {
	C.g_type_query(type_, &query)
	return
}

/*
Registers @type_name as the name of a new dynamic type derived from
@parent_type.  The type system uses the information contained in the
#GTypePlugin structure pointed to by @plugin to manage the type and its
instances (if not abstract).  The value of @flags determines the nature
(e.g. abstract or not) of the type.
*/
func TypeRegisterDynamic(parent_type C.GType, type_name string, plugin *C.GTypePlugin, flags C.GTypeFlags) (return__ C.GType) {
	__cgo__type_name := (*C.gchar)(unsafe.Pointer(C.CString(type_name)))
	return__ = C.g_type_register_dynamic(parent_type, __cgo__type_name, plugin, flags)
	C.free(unsafe.Pointer(__cgo__type_name))
	return
}

/*
Registers @type_id as the predefined identifier and @type_name as the
name of a fundamental type. If @type_id is already registered, or a
type named @type_name is already registered, the behaviour is undefined.
The type system uses the information contained in the #GTypeInfo structure
pointed to by @info and the #GTypeFundamentalInfo structure pointed to by
@finfo to manage the type and its instances. The value of @flags determines
additional characteristics of the fundamental type.
*/
func TypeRegisterFundamental(type_id C.GType, type_name string, info *C.GTypeInfo, finfo *C.GTypeFundamentalInfo, flags C.GTypeFlags) (return__ C.GType) {
	__cgo__type_name := (*C.gchar)(unsafe.Pointer(C.CString(type_name)))
	return__ = C.g_type_register_fundamental(type_id, __cgo__type_name, info, finfo, flags)
	C.free(unsafe.Pointer(__cgo__type_name))
	return
}

/*
Registers @type_name as the name of a new static type derived from
@parent_type. The type system uses the information contained in the
#GTypeInfo structure pointed to by @info to manage the type and its
instances (if not abstract). The value of @flags determines the nature
(e.g. abstract or not) of the type.
*/
func TypeRegisterStatic(parent_type C.GType, type_name string, info *C.GTypeInfo, flags C.GTypeFlags) (return__ C.GType) {
	__cgo__type_name := (*C.gchar)(unsafe.Pointer(C.CString(type_name)))
	return__ = C.g_type_register_static(parent_type, __cgo__type_name, info, flags)
	C.free(unsafe.Pointer(__cgo__type_name))
	return
}

/*
Registers @type_name as the name of a new static type derived from
@parent_type.  The value of @flags determines the nature (e.g.
abstract or not) of the type. It works by filling a #GTypeInfo
struct and calling g_type_register_static().
*/
func TypeRegisterStaticSimple(parent_type C.GType, type_name string, class_size uint, class_init C.GClassInitFunc, instance_size uint, instance_init C.GInstanceInitFunc, flags C.GTypeFlags) (return__ C.GType) {
	__cgo__type_name := (*C.gchar)(unsafe.Pointer(C.CString(type_name)))
	return__ = C.g_type_register_static_simple(parent_type, __cgo__type_name, C.guint(class_size), class_init, C.guint(instance_size), instance_init, flags)
	C.free(unsafe.Pointer(__cgo__type_name))
	return
}

/*
Removes a previously installed #GTypeClassCacheFunc. The cache
maintained by @cache_func has to be empty when calling
g_type_remove_class_cache_func() to avoid leaks.
*/
func TypeRemoveClassCacheFunc(cache_data unsafe.Pointer, cache_func C.GTypeClassCacheFunc) {
	C.g_type_remove_class_cache_func((C.gpointer)(cache_data), cache_func)
	return
}

/*
Removes an interface check function added with
g_type_add_interface_check().
*/
func TypeRemoveInterfaceCheck(check_data unsafe.Pointer, check_func C.GTypeInterfaceCheckFunc) {
	C.g_type_remove_interface_check((C.gpointer)(check_data), check_func)
	return
}

/*
Attaches arbitrary data to a type.
*/
func TypeSetQdata(type_ C.GType, quark C.GQuark, data unsafe.Pointer) {
	C.g_type_set_qdata(type_, quark, (C.gpointer)(data))
	return
}

func TypeTestFlags(type_ C.GType, flags uint) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_type_test_flags(type_, C.guint(flags))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the location of the #GTypeValueTable associated with @type.

Note that this function should only be used from source code
that implements or has internal knowledge of the implementation of
@type.
*/
func TypeValueTablePeek(type_ C.GType) (return__ *C.GTypeValueTable) {
	return__ = C.g_type_value_table_peek(type_)
	return
}

/*
Registers a value transformation function for use in g_value_transform().
A previously registered transformation function for @src_type and @dest_type
will be replaced.
*/
func ValueRegisterTransformFunc(src_type C.GType, dest_type C.GType, transform_func C.GValueTransform) {
	C.g_value_register_transform_func(src_type, dest_type, transform_func)
	return
}

/*
Returns whether a #GValue of type @src_type can be copied into
a #GValue of type @dest_type.
*/
func ValueTypeCompatible(src_type C.GType, dest_type C.GType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_value_type_compatible(src_type, dest_type)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Check whether g_value_transform() is able to transform values
of type @src_type into values of type @dest_type. Note that for
the types to be transformable, they must be compatible and a
transform function must be registered.
*/
func ValueTypeTransformable(src_type C.GType, dest_type C.GType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_value_type_transformable(src_type, dest_type)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}
