package gdk

/*
#include <gdk/gdk.h>
#include <gdk/gdkprivate.h>
#include <stdlib.h>
#include <glib-object.h>
#cgo pkg-config: gobject-2.0
*/
import "C"
import "unsafe"
import "github.com/reusee/ggir/gobject"
import "runtime"

func init() {
	_ = unsafe.Pointer(nil)
	_ = runtime.Compiler
}

type AppLaunchContext struct {
	*TraitAppLaunchContext
	CPointer unsafe.Pointer
}

func NewAppLaunchContextFromCPointer(p unsafe.Pointer) *AppLaunchContext {
	ret := &AppLaunchContext{
		NewTraitAppLaunchContext(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *AppLaunchContext) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Cursor struct {
	*TraitCursor
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCursorFromCPointer(p unsafe.Pointer) *Cursor {
	ret := &Cursor{
		NewTraitCursor(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Cursor) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Device struct {
	*TraitDevice
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDeviceFromCPointer(p unsafe.Pointer) *Device {
	ret := &Device{
		NewTraitDevice(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Device) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DeviceManager struct {
	*TraitDeviceManager
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDeviceManagerFromCPointer(p unsafe.Pointer) *DeviceManager {
	ret := &DeviceManager{
		NewTraitDeviceManager(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DeviceManager) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Display struct {
	*TraitDisplay
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDisplayFromCPointer(p unsafe.Pointer) *Display {
	ret := &Display{
		NewTraitDisplay(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Display) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DisplayManager struct {
	*TraitDisplayManager
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDisplayManagerFromCPointer(p unsafe.Pointer) *DisplayManager {
	ret := &DisplayManager{
		NewTraitDisplayManager(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DisplayManager) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DragContext struct {
	*TraitDragContext
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDragContextFromCPointer(p unsafe.Pointer) *DragContext {
	ret := &DragContext{
		NewTraitDragContext(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DragContext) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type FrameClock struct {
	*TraitFrameClock
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFrameClockFromCPointer(p unsafe.Pointer) *FrameClock {
	ret := &FrameClock{
		NewTraitFrameClock(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FrameClock) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Keymap struct {
	*TraitKeymap
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewKeymapFromCPointer(p unsafe.Pointer) *Keymap {
	ret := &Keymap{
		NewTraitKeymap(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Keymap) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Screen struct {
	*TraitScreen
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewScreenFromCPointer(p unsafe.Pointer) *Screen {
	ret := &Screen{
		NewTraitScreen(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Screen) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Visual struct {
	*TraitVisual
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewVisualFromCPointer(p unsafe.Pointer) *Visual {
	ret := &Visual{
		NewTraitVisual(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Visual) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Window struct {
	*TraitWindow
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewWindowFromCPointer(p unsafe.Pointer) *Window {
	ret := &Window{
		NewTraitWindow(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Window) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}
