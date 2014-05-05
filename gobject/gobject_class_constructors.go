package gobject

/*
#include <glib-object.h>
#include <gobject/gvaluecollector.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "reflect"
import "errors"

func init() {
	_ = unsafe.Pointer(nil)
	_ = reflect.ValueOf(nil)
	_ = errors.New("")
}

var UnusedFix_ bool

// g_object_new_valist is not generated due to varargs

/*
Creates a new instance of a #GObject subtype and sets its properties.

Construction parameters (see #G_PARAM_CONSTRUCT, #G_PARAM_CONSTRUCT_ONLY)
which are not explicitly specified are set to their default values.
*/
func ObjectNewv(object_type C.GType, n_parameters uint, parameters *C.GParameter) (return__ *Object) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_object_newv(object_type, C.guint(n_parameters), parameters)
	if __cgo__return__ != nil {
		return__ = NewObjectFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}
