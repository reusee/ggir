package glib

/*
#include <glib.h>
#include <glib-object.h>
#include <glib-unix.h>
#include <glib/gstdio.h>
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
