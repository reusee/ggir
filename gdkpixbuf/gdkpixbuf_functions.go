package gdkpixbuf

/*
#include <gdk-pixbuf/gdk-pixbuf.h>
#include <gdk-pixbuf/gdk-pixdata.h>
#include <stdlib.h>
#cgo pkg-config: gdk-pixbuf-2.0
*/
import "C"

import "github.com/reusee/ggir/gobject"
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

var _ = gobject.UnusedFix_

func PixbufErrorQuark() (return__ C.GQuark) {
	return__ = C.gdk_pixbuf_error_quark()
	return
}
