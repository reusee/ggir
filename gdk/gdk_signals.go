package gdk

/*
#include <gdk/gdk.h>
#include <gdk/gdkprivate.h>
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

func (self *TraitDevice) OnChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "changed", func(_self reflect.Value) {
	})
}

func (self *TraitDeviceManager) OnDeviceAdded(f func(*Device)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "device-added", func(_self reflect.Value, device reflect.Value) {
	})
}

func (self *TraitDeviceManager) OnDeviceChanged(f func(*Device)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "device-changed", func(_self reflect.Value, device reflect.Value) {
	})
}

func (self *TraitDeviceManager) OnDeviceRemoved(f func(*Device)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "device-removed", func(_self reflect.Value, device reflect.Value) {
	})
}

func (self *TraitDisplay) OnClosed(f func(bool)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "closed", func(_self reflect.Value, is_error reflect.Value) {
	})
}

func (self *TraitDisplay) OnOpened(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "opened", func(_self reflect.Value) {
	})
}

func (self *TraitDisplayManager) OnDisplayOpened(f func(*Display)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "display-opened", func(_self reflect.Value, display reflect.Value) {
	})
}

func (self *TraitFrameClock) OnAfterPaint(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "after-paint", func(_self reflect.Value) {
	})
}

func (self *TraitFrameClock) OnBeforePaint(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "before-paint", func(_self reflect.Value) {
	})
}

func (self *TraitFrameClock) OnFlushEvents(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "flush-events", func(_self reflect.Value) {
	})
}

func (self *TraitFrameClock) OnLayout(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "layout", func(_self reflect.Value) {
	})
}

func (self *TraitFrameClock) OnPaint(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "paint", func(_self reflect.Value) {
	})
}

func (self *TraitFrameClock) OnResumeEvents(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "resume-events", func(_self reflect.Value) {
	})
}

func (self *TraitFrameClock) OnUpdate(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "update", func(_self reflect.Value) {
	})
}

func (self *TraitKeymap) OnDirectionChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "direction-changed", func(_self reflect.Value) {
	})
}

func (self *TraitKeymap) OnKeysChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "keys-changed", func(_self reflect.Value) {
	})
}

func (self *TraitKeymap) OnStateChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "state-changed", func(_self reflect.Value) {
	})
}

func (self *TraitScreen) OnCompositedChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "composited-changed", func(_self reflect.Value) {
	})
}

func (self *TraitScreen) OnMonitorsChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "monitors-changed", func(_self reflect.Value) {
	})
}

func (self *TraitScreen) OnSizeChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "size-changed", func(_self reflect.Value) {
	})
}

func (self *TraitWindow) OnCreateSurface(f func(int, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "create-surface", func(_self reflect.Value, width reflect.Value, height reflect.Value) {
	})
}

func (self *TraitWindow) OnFromEmbedder(f func(float64, float64, unsafe.Pointer, unsafe.Pointer)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "from-embedder", func(_self reflect.Value, embedder_x reflect.Value, embedder_y reflect.Value, offscreen_x reflect.Value, offscreen_y reflect.Value) {
	})
}

func (self *TraitWindow) OnPickEmbeddedChild(f func(float64, float64)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "pick-embedded-child", func(_self reflect.Value, x reflect.Value, y reflect.Value) {
	})
}

func (self *TraitWindow) OnToEmbedder(f func(float64, float64, unsafe.Pointer, unsafe.Pointer)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "to-embedder", func(_self reflect.Value, offscreen_x reflect.Value, offscreen_y reflect.Value, embedder_x reflect.Value, embedder_y reflect.Value) {
	})
}
