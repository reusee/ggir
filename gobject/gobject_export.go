package gobject

/*
#include <glib-object.h>
*/
import "C"

import (
	"reflect"
	"unsafe"
)

//export closureMarshal
func closureMarshal(closure *C.GClosure, ret *C.GValue, nParams C.guint, params *C.GValue, hint, data C.gpointer) {
	// callback value
	f := *((*interface{})(unsafe.Pointer(data)))
	fValue := reflect.ValueOf(f)
	fType := fValue.Type()

	// convert GValue to reflect.Value
	var paramSlice []C.GValue
	h := (*reflect.SliceHeader)(unsafe.Pointer(&paramSlice))
	h.Len = int(nParams)
	h.Cap = h.Len
	h.Data = uintptr(unsafe.Pointer(params))
	var arguments []reflect.Value
	for i, gv := range paramSlice {
		if i == fType.NumIn() {
			break
		}
		goValue := fromGValue(&gv)
		var arg reflect.Value
		switch fType.In(i).Kind() {
		case reflect.Ptr:
			p := goValue.(unsafe.Pointer)
			arg = reflect.NewAt(fType.In(i), unsafe.Pointer(&p)).Elem()
		default:
			panic("FIXME") //TODO
		}
		arguments = append(arguments, arg)
	}

	// call
	fValue.Call(arguments[:fType.NumIn()])

	//TODO set return value

}
