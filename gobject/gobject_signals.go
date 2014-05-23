package gobject

/*
#include <glib-object.h>
#include <gobject/gvaluecollector.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "reflect"

func init() {
	_ = unsafe.Pointer(nil)
	_ = reflect.ValueOf(nil)
}
func (self *TraitObject) OnNotify(f func(*ParamSpec)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "notify", func(_self reflect.Value, pspec reflect.Value) {
	})
}
