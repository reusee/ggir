package gobject

/*
#include <glib-object.h>
#include <gobject/gvaluecollector.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func init() {
	_ = unsafe.Pointer(nil)
}

type Binding struct {
	*_TraitBinding
	*_TraitObject
	CPointer unsafe.Pointer
}

func NewBindingFromCPointer(p unsafe.Pointer) *Binding {
	return &Binding{
		&_TraitBinding{(*C.GBinding)(p)},
		&_TraitObject{(*C.GObject)(p)},
		p,
	}
}

type InitiallyUnowned struct {
	*_TraitInitiallyUnowned
	*_TraitObject
	CPointer unsafe.Pointer
}

func NewInitiallyUnownedFromCPointer(p unsafe.Pointer) *InitiallyUnowned {
	return &InitiallyUnowned{
		&_TraitInitiallyUnowned{(*C.GInitiallyUnowned)(p)},
		&_TraitObject{(*C.GObject)(p)},
		p,
	}
}

type Object struct {
	*_TraitObject
	CPointer unsafe.Pointer
}

func NewObjectFromCPointer(p unsafe.Pointer) *Object {
	return &Object{
		&_TraitObject{(*C.GObject)(p)},
		p,
	}
}

type ParamSpec struct {
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecFromCPointer(p unsafe.Pointer) *ParamSpec {
	return &ParamSpec{
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecBoolean struct {
	*_TraitParamSpecBoolean
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecBooleanFromCPointer(p unsafe.Pointer) *ParamSpecBoolean {
	return &ParamSpecBoolean{
		&_TraitParamSpecBoolean{(*C.GParamSpecBoolean)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecBoxed struct {
	*_TraitParamSpecBoxed
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecBoxedFromCPointer(p unsafe.Pointer) *ParamSpecBoxed {
	return &ParamSpecBoxed{
		&_TraitParamSpecBoxed{(*C.GParamSpecBoxed)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecChar struct {
	*_TraitParamSpecChar
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecCharFromCPointer(p unsafe.Pointer) *ParamSpecChar {
	return &ParamSpecChar{
		&_TraitParamSpecChar{(*C.GParamSpecChar)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecDouble struct {
	*_TraitParamSpecDouble
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecDoubleFromCPointer(p unsafe.Pointer) *ParamSpecDouble {
	return &ParamSpecDouble{
		&_TraitParamSpecDouble{(*C.GParamSpecDouble)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecEnum struct {
	*_TraitParamSpecEnum
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecEnumFromCPointer(p unsafe.Pointer) *ParamSpecEnum {
	return &ParamSpecEnum{
		&_TraitParamSpecEnum{(*C.GParamSpecEnum)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecFlags struct {
	*_TraitParamSpecFlags
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecFlagsFromCPointer(p unsafe.Pointer) *ParamSpecFlags {
	return &ParamSpecFlags{
		&_TraitParamSpecFlags{(*C.GParamSpecFlags)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecFloat struct {
	*_TraitParamSpecFloat
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecFloatFromCPointer(p unsafe.Pointer) *ParamSpecFloat {
	return &ParamSpecFloat{
		&_TraitParamSpecFloat{(*C.GParamSpecFloat)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecGType struct {
	*_TraitParamSpecGType
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecGTypeFromCPointer(p unsafe.Pointer) *ParamSpecGType {
	return &ParamSpecGType{
		&_TraitParamSpecGType{(*C.GParamSpecGType)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecInt struct {
	*_TraitParamSpecInt
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecIntFromCPointer(p unsafe.Pointer) *ParamSpecInt {
	return &ParamSpecInt{
		&_TraitParamSpecInt{(*C.GParamSpecInt)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecInt64 struct {
	*_TraitParamSpecInt64
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecInt64FromCPointer(p unsafe.Pointer) *ParamSpecInt64 {
	return &ParamSpecInt64{
		&_TraitParamSpecInt64{(*C.GParamSpecInt64)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecLong struct {
	*_TraitParamSpecLong
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecLongFromCPointer(p unsafe.Pointer) *ParamSpecLong {
	return &ParamSpecLong{
		&_TraitParamSpecLong{(*C.GParamSpecLong)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecObject struct {
	*_TraitParamSpecObject
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecObjectFromCPointer(p unsafe.Pointer) *ParamSpecObject {
	return &ParamSpecObject{
		&_TraitParamSpecObject{(*C.GParamSpecObject)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecOverride struct {
	*_TraitParamSpecOverride
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecOverrideFromCPointer(p unsafe.Pointer) *ParamSpecOverride {
	return &ParamSpecOverride{
		&_TraitParamSpecOverride{(*C.GParamSpecOverride)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecParam struct {
	*_TraitParamSpecParam
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecParamFromCPointer(p unsafe.Pointer) *ParamSpecParam {
	return &ParamSpecParam{
		&_TraitParamSpecParam{(*C.GParamSpecParam)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecPointer struct {
	*_TraitParamSpecPointer
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecPointerFromCPointer(p unsafe.Pointer) *ParamSpecPointer {
	return &ParamSpecPointer{
		&_TraitParamSpecPointer{(*C.GParamSpecPointer)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecString struct {
	*_TraitParamSpecString
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecStringFromCPointer(p unsafe.Pointer) *ParamSpecString {
	return &ParamSpecString{
		&_TraitParamSpecString{(*C.GParamSpecString)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecUChar struct {
	*_TraitParamSpecUChar
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecUCharFromCPointer(p unsafe.Pointer) *ParamSpecUChar {
	return &ParamSpecUChar{
		&_TraitParamSpecUChar{(*C.GParamSpecUChar)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecUInt struct {
	*_TraitParamSpecUInt
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecUIntFromCPointer(p unsafe.Pointer) *ParamSpecUInt {
	return &ParamSpecUInt{
		&_TraitParamSpecUInt{(*C.GParamSpecUInt)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecUInt64 struct {
	*_TraitParamSpecUInt64
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecUInt64FromCPointer(p unsafe.Pointer) *ParamSpecUInt64 {
	return &ParamSpecUInt64{
		&_TraitParamSpecUInt64{(*C.GParamSpecUInt64)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecULong struct {
	*_TraitParamSpecULong
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecULongFromCPointer(p unsafe.Pointer) *ParamSpecULong {
	return &ParamSpecULong{
		&_TraitParamSpecULong{(*C.GParamSpecULong)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecUnichar struct {
	*_TraitParamSpecUnichar
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecUnicharFromCPointer(p unsafe.Pointer) *ParamSpecUnichar {
	return &ParamSpecUnichar{
		&_TraitParamSpecUnichar{(*C.GParamSpecUnichar)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecValueArray struct {
	*_TraitParamSpecValueArray
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecValueArrayFromCPointer(p unsafe.Pointer) *ParamSpecValueArray {
	return &ParamSpecValueArray{
		&_TraitParamSpecValueArray{(*C.GParamSpecValueArray)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type ParamSpecVariant struct {
	*_TraitParamSpecVariant
	*_TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecVariantFromCPointer(p unsafe.Pointer) *ParamSpecVariant {
	return &ParamSpecVariant{
		&_TraitParamSpecVariant{(*C.GParamSpecVariant)(p)},
		&_TraitParamSpec{(*C.GParamSpec)(p)},
		p,
	}
}

type TypeModule struct {
	*_TraitTypeModule
	*_TraitObject
	CPointer unsafe.Pointer
}

func NewTypeModuleFromCPointer(p unsafe.Pointer) *TypeModule {
	return &TypeModule{
		&_TraitTypeModule{(*C.GTypeModule)(p)},
		&_TraitObject{(*C.GObject)(p)},
		p,
	}
}
