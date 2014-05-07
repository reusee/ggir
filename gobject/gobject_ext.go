package gobject

/*
#include <glib-object.h>
#include <stdlib.h>

static inline GType gvalue_get_type(GValue *v) {
	return G_VALUE_TYPE(v);
}

static inline GType gtype_get_fundamental(GType t) {
	return G_TYPE_FUNDAMENTAL(t);
}

extern void closureMarshal(GClosure*, GValue*, guint, GValue*, gpointer, gpointer);

GClosure* new_closure(void *data) {
	GClosure *closure = g_closure_new_simple(sizeof(GClosure), NULL);
	g_closure_set_meta_marshal(closure, data, (GClosureMarshal)(closureMarshal));
	return closure;
}

*/
import "C"
import (
	"fmt"
	"sync"
	"unsafe"
)

var refHolder []interface{}
var refHolderLock sync.Mutex

func (self *TraitObject) Connect(signal string, cb interface{}) uint64 {
	return Connect(unsafe.Pointer(self.CPointer), signal, cb)
}

func Connect(obj unsafe.Pointer, signal string, cb interface{}) uint64 {
	cbp := &cb
	refHolderLock.Lock()
	refHolder = append(refHolder, cbp) //FIXME deref
	refHolderLock.Unlock()
	closure := C.new_closure(unsafe.Pointer(cbp))
	cSignal := (*C.gchar)(unsafe.Pointer(C.CString(signal)))
	defer C.free(unsafe.Pointer(cSignal))
	id := C.g_signal_connect_closure(C.gpointer(obj), cSignal, closure, C.gboolean(0))
	return uint64(id)
}

func fromGValue(v *C.GValue) (ret interface{}) {
	valueType := C.gvalue_get_type(v)
	fundamentalType := C.gtype_get_fundamental(valueType)
	switch fundamentalType {
	case C.G_TYPE_OBJECT:
		ret = unsafe.Pointer(C.g_value_get_object(v))
	case C.G_TYPE_STRING:
		ret = fromGStr(C.g_value_get_string(v))
	case C.G_TYPE_UINT:
		ret = int(C.g_value_get_uint(v))
	case C.G_TYPE_BOXED:
		ret = unsafe.Pointer(C.g_value_get_boxed(v))
	case C.G_TYPE_BOOLEAN:
		ret = C.g_value_get_boolean(v) == C.gboolean(1)
	default:
		fmt.Printf("from type %s %T\n", fromGStr(C.g_type_name(fundamentalType)), v)
		panic("FIXME") //TODO
	}
	return
}

func fromGStr(s *C.gchar) string {
	return C.GoString((*C.char)(unsafe.Pointer(s)))
}
