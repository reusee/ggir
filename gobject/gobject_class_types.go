package gobject

/*
#include <glib-object.h>
#include <gobject/gvaluecollector.h>
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

type Binding struct {
	*TraitBinding
	*TraitObject
	CPointer unsafe.Pointer
}

func NewBindingFromCPointer(p unsafe.Pointer) *Binding {
	ret := &Binding{
		NewTraitBinding(p),
		NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Binding) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type InitiallyUnowned struct {
	*TraitInitiallyUnowned
	*TraitObject
	CPointer unsafe.Pointer
}

func NewInitiallyUnownedFromCPointer(p unsafe.Pointer) *InitiallyUnowned {
	ret := &InitiallyUnowned{
		NewTraitInitiallyUnowned(p),
		NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *InitiallyUnowned) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Object struct {
	*TraitObject
	CPointer unsafe.Pointer
}

func NewObjectFromCPointer(p unsafe.Pointer) *Object {
	ret := &Object{
		NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Object) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpec struct {
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecFromCPointer(p unsafe.Pointer) *ParamSpec {
	ret := &ParamSpec{
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpec) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecBoolean struct {
	*TraitParamSpecBoolean
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecBooleanFromCPointer(p unsafe.Pointer) *ParamSpecBoolean {
	ret := &ParamSpecBoolean{
		NewTraitParamSpecBoolean(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecBoolean) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecBoxed struct {
	*TraitParamSpecBoxed
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecBoxedFromCPointer(p unsafe.Pointer) *ParamSpecBoxed {
	ret := &ParamSpecBoxed{
		NewTraitParamSpecBoxed(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecBoxed) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecChar struct {
	*TraitParamSpecChar
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecCharFromCPointer(p unsafe.Pointer) *ParamSpecChar {
	ret := &ParamSpecChar{
		NewTraitParamSpecChar(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecChar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecDouble struct {
	*TraitParamSpecDouble
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecDoubleFromCPointer(p unsafe.Pointer) *ParamSpecDouble {
	ret := &ParamSpecDouble{
		NewTraitParamSpecDouble(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecDouble) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecEnum struct {
	*TraitParamSpecEnum
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecEnumFromCPointer(p unsafe.Pointer) *ParamSpecEnum {
	ret := &ParamSpecEnum{
		NewTraitParamSpecEnum(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecEnum) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecFlags struct {
	*TraitParamSpecFlags
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecFlagsFromCPointer(p unsafe.Pointer) *ParamSpecFlags {
	ret := &ParamSpecFlags{
		NewTraitParamSpecFlags(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecFlags) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecFloat struct {
	*TraitParamSpecFloat
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecFloatFromCPointer(p unsafe.Pointer) *ParamSpecFloat {
	ret := &ParamSpecFloat{
		NewTraitParamSpecFloat(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecFloat) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecGType struct {
	*TraitParamSpecGType
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecGTypeFromCPointer(p unsafe.Pointer) *ParamSpecGType {
	ret := &ParamSpecGType{
		NewTraitParamSpecGType(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecGType) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecInt struct {
	*TraitParamSpecInt
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecIntFromCPointer(p unsafe.Pointer) *ParamSpecInt {
	ret := &ParamSpecInt{
		NewTraitParamSpecInt(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecInt) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecInt64 struct {
	*TraitParamSpecInt64
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecInt64FromCPointer(p unsafe.Pointer) *ParamSpecInt64 {
	ret := &ParamSpecInt64{
		NewTraitParamSpecInt64(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecInt64) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecLong struct {
	*TraitParamSpecLong
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecLongFromCPointer(p unsafe.Pointer) *ParamSpecLong {
	ret := &ParamSpecLong{
		NewTraitParamSpecLong(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecLong) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecObject struct {
	*TraitParamSpecObject
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecObjectFromCPointer(p unsafe.Pointer) *ParamSpecObject {
	ret := &ParamSpecObject{
		NewTraitParamSpecObject(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecObject) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecOverride struct {
	*TraitParamSpecOverride
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecOverrideFromCPointer(p unsafe.Pointer) *ParamSpecOverride {
	ret := &ParamSpecOverride{
		NewTraitParamSpecOverride(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecOverride) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecParam struct {
	*TraitParamSpecParam
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecParamFromCPointer(p unsafe.Pointer) *ParamSpecParam {
	ret := &ParamSpecParam{
		NewTraitParamSpecParam(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecParam) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecPointer struct {
	*TraitParamSpecPointer
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecPointerFromCPointer(p unsafe.Pointer) *ParamSpecPointer {
	ret := &ParamSpecPointer{
		NewTraitParamSpecPointer(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecPointer) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecString struct {
	*TraitParamSpecString
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecStringFromCPointer(p unsafe.Pointer) *ParamSpecString {
	ret := &ParamSpecString{
		NewTraitParamSpecString(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecString) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecUChar struct {
	*TraitParamSpecUChar
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecUCharFromCPointer(p unsafe.Pointer) *ParamSpecUChar {
	ret := &ParamSpecUChar{
		NewTraitParamSpecUChar(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecUChar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecUInt struct {
	*TraitParamSpecUInt
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecUIntFromCPointer(p unsafe.Pointer) *ParamSpecUInt {
	ret := &ParamSpecUInt{
		NewTraitParamSpecUInt(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecUInt) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecUInt64 struct {
	*TraitParamSpecUInt64
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecUInt64FromCPointer(p unsafe.Pointer) *ParamSpecUInt64 {
	ret := &ParamSpecUInt64{
		NewTraitParamSpecUInt64(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecUInt64) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecULong struct {
	*TraitParamSpecULong
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecULongFromCPointer(p unsafe.Pointer) *ParamSpecULong {
	ret := &ParamSpecULong{
		NewTraitParamSpecULong(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecULong) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecUnichar struct {
	*TraitParamSpecUnichar
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecUnicharFromCPointer(p unsafe.Pointer) *ParamSpecUnichar {
	ret := &ParamSpecUnichar{
		NewTraitParamSpecUnichar(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecUnichar) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecValueArray struct {
	*TraitParamSpecValueArray
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecValueArrayFromCPointer(p unsafe.Pointer) *ParamSpecValueArray {
	ret := &ParamSpecValueArray{
		NewTraitParamSpecValueArray(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecValueArray) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ParamSpecVariant struct {
	*TraitParamSpecVariant
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecVariantFromCPointer(p unsafe.Pointer) *ParamSpecVariant {
	ret := &ParamSpecVariant{
		NewTraitParamSpecVariant(p),
		NewTraitParamSpec(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ParamSpecVariant) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TypeModule struct {
	*TraitTypeModule
	*TraitObject
	CPointer unsafe.Pointer
}

func NewTypeModuleFromCPointer(p unsafe.Pointer) *TypeModule {
	ret := &TypeModule{
		NewTraitTypeModule(p),
		NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TypeModule) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}
