package cairo

/*
#include <cairo.h>
#include <stdlib.h>
#include <glib-object.h>
#cgo pkg-config: gobject-2.0
*/
import "C"
import "unsafe"
import "runtime"

func init() {
	_ = unsafe.Pointer(nil)
	_ = runtime.Compiler
}
