package gdkpixbuf

/*
#include <gdk-pixbuf/gdk-pixbuf.h>
#include <gdk-pixbuf/gdk-pixdata.h>
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

type Pixbuf struct {
	*TraitPixbuf
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPixbufFromCPointer(p unsafe.Pointer) *Pixbuf {
	ret := &Pixbuf{
		NewTraitPixbuf(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Pixbuf) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type PixbufAnimation struct {
	*TraitPixbufAnimation
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPixbufAnimationFromCPointer(p unsafe.Pointer) *PixbufAnimation {
	ret := &PixbufAnimation{
		NewTraitPixbufAnimation(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *PixbufAnimation) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type PixbufAnimationIter struct {
	*TraitPixbufAnimationIter
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPixbufAnimationIterFromCPointer(p unsafe.Pointer) *PixbufAnimationIter {
	ret := &PixbufAnimationIter{
		NewTraitPixbufAnimationIter(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *PixbufAnimationIter) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type PixbufLoader struct {
	*TraitPixbufLoader
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPixbufLoaderFromCPointer(p unsafe.Pointer) *PixbufLoader {
	ret := &PixbufLoader{
		NewTraitPixbufLoader(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *PixbufLoader) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type PixbufSimpleAnim struct {
	*TraitPixbufSimpleAnim
	*TraitPixbufAnimation
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPixbufSimpleAnimFromCPointer(p unsafe.Pointer) *PixbufSimpleAnim {
	ret := &PixbufSimpleAnim{
		NewTraitPixbufSimpleAnim(p),
		NewTraitPixbufAnimation(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *PixbufSimpleAnim) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}
