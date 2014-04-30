package glib

/*
#include <stdio.h>
#include <glib.h>
#include <glib-object.h>
#include <glib-unix.h>
#include <glib/gstdio.h>
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
