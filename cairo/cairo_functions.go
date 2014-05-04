package cairo

/*
#include <cairo.h>
#include <stdlib.h>
#cgo pkg-config: cairo
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

// cairo_image_surface_create is not generated due to explicit ignore
