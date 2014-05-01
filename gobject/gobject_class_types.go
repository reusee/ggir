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
	*TraitBinding
	*TraitObject
	CPointer unsafe.Pointer
}

func NewBindingFromCPointer(p unsafe.Pointer) *Binding {
	return &Binding{
		NewTraitBinding(p),
		NewTraitObject(p),
		p,
	}
}

type InitiallyUnowned struct {
	*TraitInitiallyUnowned
	*TraitObject
	CPointer unsafe.Pointer
}

func NewInitiallyUnownedFromCPointer(p unsafe.Pointer) *InitiallyUnowned {
	return &InitiallyUnowned{
		NewTraitInitiallyUnowned(p),
		NewTraitObject(p),
		p,
	}
}

type Object struct {
	*TraitObject
	CPointer unsafe.Pointer
}

func NewObjectFromCPointer(p unsafe.Pointer) *Object {
	return &Object{
		NewTraitObject(p),
		p,
	}
}

type ParamSpec struct {
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecFromCPointer(p unsafe.Pointer) *ParamSpec {
	return &ParamSpec{
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecBoolean struct {
	*TraitParamSpecBoolean
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecBooleanFromCPointer(p unsafe.Pointer) *ParamSpecBoolean {
	return &ParamSpecBoolean{
		NewTraitParamSpecBoolean(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecBoxed struct {
	*TraitParamSpecBoxed
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecBoxedFromCPointer(p unsafe.Pointer) *ParamSpecBoxed {
	return &ParamSpecBoxed{
		NewTraitParamSpecBoxed(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecChar struct {
	*TraitParamSpecChar
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecCharFromCPointer(p unsafe.Pointer) *ParamSpecChar {
	return &ParamSpecChar{
		NewTraitParamSpecChar(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecDouble struct {
	*TraitParamSpecDouble
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecDoubleFromCPointer(p unsafe.Pointer) *ParamSpecDouble {
	return &ParamSpecDouble{
		NewTraitParamSpecDouble(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecEnum struct {
	*TraitParamSpecEnum
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecEnumFromCPointer(p unsafe.Pointer) *ParamSpecEnum {
	return &ParamSpecEnum{
		NewTraitParamSpecEnum(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecFlags struct {
	*TraitParamSpecFlags
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecFlagsFromCPointer(p unsafe.Pointer) *ParamSpecFlags {
	return &ParamSpecFlags{
		NewTraitParamSpecFlags(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecFloat struct {
	*TraitParamSpecFloat
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecFloatFromCPointer(p unsafe.Pointer) *ParamSpecFloat {
	return &ParamSpecFloat{
		NewTraitParamSpecFloat(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecGType struct {
	*TraitParamSpecGType
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecGTypeFromCPointer(p unsafe.Pointer) *ParamSpecGType {
	return &ParamSpecGType{
		NewTraitParamSpecGType(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecInt struct {
	*TraitParamSpecInt
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecIntFromCPointer(p unsafe.Pointer) *ParamSpecInt {
	return &ParamSpecInt{
		NewTraitParamSpecInt(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecInt64 struct {
	*TraitParamSpecInt64
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecInt64FromCPointer(p unsafe.Pointer) *ParamSpecInt64 {
	return &ParamSpecInt64{
		NewTraitParamSpecInt64(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecLong struct {
	*TraitParamSpecLong
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecLongFromCPointer(p unsafe.Pointer) *ParamSpecLong {
	return &ParamSpecLong{
		NewTraitParamSpecLong(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecObject struct {
	*TraitParamSpecObject
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecObjectFromCPointer(p unsafe.Pointer) *ParamSpecObject {
	return &ParamSpecObject{
		NewTraitParamSpecObject(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecOverride struct {
	*TraitParamSpecOverride
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecOverrideFromCPointer(p unsafe.Pointer) *ParamSpecOverride {
	return &ParamSpecOverride{
		NewTraitParamSpecOverride(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecParam struct {
	*TraitParamSpecParam
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecParamFromCPointer(p unsafe.Pointer) *ParamSpecParam {
	return &ParamSpecParam{
		NewTraitParamSpecParam(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecPointer struct {
	*TraitParamSpecPointer
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecPointerFromCPointer(p unsafe.Pointer) *ParamSpecPointer {
	return &ParamSpecPointer{
		NewTraitParamSpecPointer(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecString struct {
	*TraitParamSpecString
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecStringFromCPointer(p unsafe.Pointer) *ParamSpecString {
	return &ParamSpecString{
		NewTraitParamSpecString(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecUChar struct {
	*TraitParamSpecUChar
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecUCharFromCPointer(p unsafe.Pointer) *ParamSpecUChar {
	return &ParamSpecUChar{
		NewTraitParamSpecUChar(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecUInt struct {
	*TraitParamSpecUInt
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecUIntFromCPointer(p unsafe.Pointer) *ParamSpecUInt {
	return &ParamSpecUInt{
		NewTraitParamSpecUInt(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecUInt64 struct {
	*TraitParamSpecUInt64
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecUInt64FromCPointer(p unsafe.Pointer) *ParamSpecUInt64 {
	return &ParamSpecUInt64{
		NewTraitParamSpecUInt64(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecULong struct {
	*TraitParamSpecULong
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecULongFromCPointer(p unsafe.Pointer) *ParamSpecULong {
	return &ParamSpecULong{
		NewTraitParamSpecULong(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecUnichar struct {
	*TraitParamSpecUnichar
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecUnicharFromCPointer(p unsafe.Pointer) *ParamSpecUnichar {
	return &ParamSpecUnichar{
		NewTraitParamSpecUnichar(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecValueArray struct {
	*TraitParamSpecValueArray
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecValueArrayFromCPointer(p unsafe.Pointer) *ParamSpecValueArray {
	return &ParamSpecValueArray{
		NewTraitParamSpecValueArray(p),
		NewTraitParamSpec(p),
		p,
	}
}

type ParamSpecVariant struct {
	*TraitParamSpecVariant
	*TraitParamSpec
	CPointer unsafe.Pointer
}

func NewParamSpecVariantFromCPointer(p unsafe.Pointer) *ParamSpecVariant {
	return &ParamSpecVariant{
		NewTraitParamSpecVariant(p),
		NewTraitParamSpec(p),
		p,
	}
}

type TypeModule struct {
	*TraitTypeModule
	*TraitObject
	CPointer unsafe.Pointer
}

func NewTypeModuleFromCPointer(p unsafe.Pointer) *TypeModule {
	return &TypeModule{
		NewTraitTypeModule(p),
		NewTraitObject(p),
		p,
	}
}
