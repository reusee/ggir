package gdkpixbuf

/*
#include <gdk-pixbuf/gdk-pixbuf.h>
#include <gdk-pixbuf/gdk-pixdata.h>
#include <stdlib.h>
*/
import "C"
import "github.com/reusee/ggir/gobject"
import "unsafe"
import "reflect"

func init() {
	_ = unsafe.Pointer(nil)
	_ = reflect.ValueOf(nil)
}

var _ = gobject.UnusedFix_

func (self *TraitPixbufLoader) OnAreaPrepared(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "area-prepared", func(_self reflect.Value) {
	})
}

func (self *TraitPixbufLoader) OnAreaUpdated(f func(int, int, int, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "area-updated", func(_self reflect.Value, x reflect.Value, y reflect.Value, width reflect.Value, height reflect.Value) {
	})
}

func (self *TraitPixbufLoader) OnClosed(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "closed", func(_self reflect.Value) {
	})
}

func (self *TraitPixbufLoader) OnSizePrepared(f func(int, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "size-prepared", func(_self reflect.Value, width reflect.Value, height reflect.Value) {
	})
}
