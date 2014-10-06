package gobject

/*
#include <glib-object.h>
#include <gobject/gvaluecollector.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "reflect"
import "errors"

func init() {
	_ = unsafe.Pointer(nil)
	_ = reflect.ValueOf(nil)
	_ = errors.New("")
}

type TraitBinding struct{ CPointer *C.GBinding }
type IsBinding interface {
	GetBindingPointer() *C.GBinding
}

func (self *TraitBinding) GetBindingPointer() *C.GBinding {
	return self.CPointer
}
func NewTraitBinding(p unsafe.Pointer) *TraitBinding {
	return &TraitBinding{(*C.GBinding)(p)}
}

/*
Retrieves the flags passed when constructing the #GBinding.
*/
func (self *TraitBinding) GetFlags() (return__ C.GBindingFlags) {
	return__ = C.g_binding_get_flags(self.CPointer)
	return
}

/*
Retrieves the #GObject instance used as the source of the binding.
*/
func (self *TraitBinding) GetSource() (return__ *Object) {
	var __cgo__return__ *C.GObject
	__cgo__return__ = C.g_binding_get_source(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewObjectFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Retrieves the name of the property of #GBinding:source used as the source
of the binding.
*/
func (self *TraitBinding) GetSourceProperty() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_binding_get_source_property(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Retrieves the #GObject instance used as the target of the binding.
*/
func (self *TraitBinding) GetTarget() (return__ *Object) {
	var __cgo__return__ *C.GObject
	__cgo__return__ = C.g_binding_get_target(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewObjectFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Retrieves the name of the property of #GBinding:target used as the target
of the binding.
*/
func (self *TraitBinding) GetTargetProperty() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_binding_get_target_property(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Explicitly releases the binding between the source and the target
property expressed by @binding.

This function will release the reference that is being held on
the @binding instance; if you want to hold on to the #GBinding instance
after calling g_binding_unbind(), you will need to hold a reference
to it.
*/
func (self *TraitBinding) Unbind() {
	C.g_binding_unbind(self.CPointer)
	return
}

type TraitInitiallyUnowned struct{ CPointer *C.GInitiallyUnowned }
type IsInitiallyUnowned interface {
	GetInitiallyUnownedPointer() *C.GInitiallyUnowned
}

func (self *TraitInitiallyUnowned) GetInitiallyUnownedPointer() *C.GInitiallyUnowned {
	return self.CPointer
}
func NewTraitInitiallyUnowned(p unsafe.Pointer) *TraitInitiallyUnowned {
	return &TraitInitiallyUnowned{(*C.GInitiallyUnowned)(p)}
}

type TraitObject struct{ CPointer *C.GObject }
type IsObject interface {
	GetObjectPointer() *C.GObject
}

func (self *TraitObject) GetObjectPointer() *C.GObject {
	return self.CPointer
}
func NewTraitObject(p unsafe.Pointer) *TraitObject {
	return &TraitObject{(*C.GObject)(p)}
}

/*
Increases the reference count of the object by one and sets a
callback to be called when all other references to the object are
dropped, or when this is already the last reference to the object
and another reference is established.

This functionality is intended for binding @object to a proxy
object managed by another memory manager. This is done with two
paired references: the strong reference added by
g_object_add_toggle_ref() and a reverse reference to the proxy
object which is either a strong reference or weak reference.

The setup is that when there are no other references to @object,
only a weak reference is held in the reverse direction from @object
to the proxy object, but when there are other references held to
@object, a strong reference is held. The @notify callback is called
when the reference from @object to the proxy object should be
"toggled" from strong to weak (@is_last_ref true) or weak to strong
(@is_last_ref false).

Since a (normal) reference must be held to the object before
calling g_object_add_toggle_ref(), the initial state of the reverse
link is always strong.

Multiple toggle references may be added to the same gobject,
however if there are multiple toggle references to an object, none
of them will ever be notified until all but one are removed.  For
this reason, you should only ever use a toggle reference if there
is important state in the proxy object.
*/
func (self *TraitObject) AddToggleRef(notify C.GToggleNotify, data unsafe.Pointer) {
	C.g_object_add_toggle_ref(self.CPointer, notify, (C.gpointer)(data))
	return
}

// g_object_add_weak_pointer is not generated due to inout param

/*
Creates a binding between @source_property on @source and @target_property
on @target. Whenever the @source_property is changed the @target_property is
updated using the same value. For instance:

|[
  g_object_bind_property (action, "active", widget, "sensitive", 0);
]|

Will result in the "sensitive" property of the widget #GObject instance to be
updated with the same value of the "active" property of the action #GObject
instance.

If @flags contains %G_BINDING_BIDIRECTIONAL then the binding will be mutual:
if @target_property on @target changes then the @source_property on @source
will be updated as well.

The binding will automatically be removed when either the @source or the
@target instances are finalized. To remove the binding without affecting the
@source and the @target you can just call g_object_unref() on the returned
#GBinding instance.

A #GObject can have multiple bindings.
*/
func (self *TraitObject) BindProperty(source_property string, target IsObject, target_property string, flags C.GBindingFlags) (return__ *Binding) {
	__cgo__source_property := (*C.gchar)(unsafe.Pointer(C.CString(source_property)))
	__cgo__target_property := (*C.gchar)(unsafe.Pointer(C.CString(target_property)))
	var __cgo__return__ *C.GBinding
	__cgo__return__ = C.g_object_bind_property((C.gpointer)(unsafe.Pointer(self.CPointer)), __cgo__source_property, C.gpointer(unsafe.Pointer(target.GetObjectPointer())), __cgo__target_property, flags)
	C.free(unsafe.Pointer(__cgo__source_property))
	C.free(unsafe.Pointer(__cgo__target_property))
	if __cgo__return__ != nil {
		return__ = NewBindingFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Complete version of g_object_bind_property().

Creates a binding between @source_property on @source and @target_property
on @target, allowing you to set the transformation functions to be used by
the binding.

If @flags contains %G_BINDING_BIDIRECTIONAL then the binding will be mutual:
if @target_property on @target changes then the @source_property on @source
will be updated as well. The @transform_from function is only used in case
of bidirectional bindings, otherwise it will be ignored

The binding will automatically be removed when either the @source or the
@target instances are finalized. To remove the binding without affecting the
@source and the @target you can just call g_object_unref() on the returned
#GBinding instance.

A #GObject can have multiple bindings.

The same @user_data parameter will be used for both @transform_to
and @transform_from transformation functions; the @notify function will
be called once, when the binding is removed. If you need different data
for each transformation function, please use
g_object_bind_property_with_closures() instead.
*/
func (self *TraitObject) BindPropertyFull(source_property string, target IsObject, target_property string, flags C.GBindingFlags, transform_to C.GBindingTransformFunc, transform_from C.GBindingTransformFunc, user_data unsafe.Pointer, notify C.GDestroyNotify) (return__ *Binding) {
	__cgo__source_property := (*C.gchar)(unsafe.Pointer(C.CString(source_property)))
	__cgo__target_property := (*C.gchar)(unsafe.Pointer(C.CString(target_property)))
	var __cgo__return__ *C.GBinding
	__cgo__return__ = C.g_object_bind_property_full((C.gpointer)(unsafe.Pointer(self.CPointer)), __cgo__source_property, C.gpointer(unsafe.Pointer(target.GetObjectPointer())), __cgo__target_property, flags, transform_to, transform_from, (C.gpointer)(user_data), notify)
	C.free(unsafe.Pointer(__cgo__source_property))
	C.free(unsafe.Pointer(__cgo__target_property))
	if __cgo__return__ != nil {
		return__ = NewBindingFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a binding between @source_property on @source and @target_property
on @target, allowing you to set the transformation functions to be used by
the binding.

This function is the language bindings friendly version of
g_object_bind_property_full(), using #GClosures instead of
function pointers.
*/
func (self *TraitObject) BindPropertyWithClosures(source_property string, target IsObject, target_property string, flags C.GBindingFlags, transform_to *C.GClosure, transform_from *C.GClosure) (return__ *Binding) {
	__cgo__source_property := (*C.gchar)(unsafe.Pointer(C.CString(source_property)))
	__cgo__target_property := (*C.gchar)(unsafe.Pointer(C.CString(target_property)))
	var __cgo__return__ *C.GBinding
	__cgo__return__ = C.g_object_bind_property_with_closures((C.gpointer)(unsafe.Pointer(self.CPointer)), __cgo__source_property, C.gpointer(unsafe.Pointer(target.GetObjectPointer())), __cgo__target_property, flags, transform_to, transform_from)
	C.free(unsafe.Pointer(__cgo__source_property))
	C.free(unsafe.Pointer(__cgo__target_property))
	if __cgo__return__ != nil {
		return__ = NewBindingFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
This is a variant of g_object_get_data() which returns
a 'duplicate' of the value. @dup_func defines the
meaning of 'duplicate' in this context, it could e.g.
take a reference on a ref-counted object.

If the @key is not set on the object then @dup_func
will be called with a %NULL argument.

Note that @dup_func is called while user data of @object
is locked.

This function can be useful to avoid races when multiple
threads are using object data on the same key on the same
object.
*/
func (self *TraitObject) DupData(key string, dup_func C.GDuplicateFunc, user_data unsafe.Pointer) (return__ unsafe.Pointer) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_object_dup_data(self.CPointer, __cgo__key, dup_func, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__key))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
This is a variant of g_object_get_qdata() which returns
a 'duplicate' of the value. @dup_func defines the
meaning of 'duplicate' in this context, it could e.g.
take a reference on a ref-counted object.

If the @quark is not set on the object then @dup_func
will be called with a %NULL argument.

Note that @dup_func is called while user data of @object
is locked.

This function can be useful to avoid races when multiple
threads are using object data on the same key on the same
object.
*/
func (self *TraitObject) DupQdata(quark C.GQuark, dup_func C.GDuplicateFunc, user_data unsafe.Pointer) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_object_dup_qdata(self.CPointer, quark, dup_func, (C.gpointer)(user_data))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
This function is intended for #GObject implementations to re-enforce
a [floating][floating-ref] object reference. Doing this is seldom
required: all #GInitiallyUnowneds are created with a floating reference
which usually just needs to be sunken by calling g_object_ref_sink().
*/
func (self *TraitObject) ForceFloating() {
	C.g_object_force_floating(self.CPointer)
	return
}

/*
Increases the freeze count on @object. If the freeze count is
non-zero, the emission of "notify" signals on @object is
stopped. The signals are queued until the freeze count is decreased
to zero. Duplicate notifications are squashed so that at most one
#GObject::notify signal is emitted for each property modified while the
object is frozen.

This is necessary for accessors that modify multiple properties to prevent
premature notification while the object is still being modified.
*/
func (self *TraitObject) FreezeNotify() {
	C.g_object_freeze_notify(self.CPointer)
	return
}

/*
Gets a named field from the objects table of associations (see g_object_set_data()).
*/
func (self *TraitObject) GetData(key string) (return__ unsafe.Pointer) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_object_get_data(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Gets a property of an object. @value must have been initialized to the
expected type of the property (or a type to which the expected type can be
transformed) using g_value_init().

In general, a copy is made of the property contents and the caller is
responsible for freeing the memory by calling g_value_unset().

Note that g_object_get_property() is really intended for language
bindings, g_object_get() is much more convenient for C programming.
*/
func (self *TraitObject) GetProperty(property_name string, value *C.GValue) {
	__cgo__property_name := (*C.gchar)(unsafe.Pointer(C.CString(property_name)))
	C.g_object_get_property(self.CPointer, __cgo__property_name, value)
	C.free(unsafe.Pointer(__cgo__property_name))
	return
}

/*
This function gets back user data pointers stored via
g_object_set_qdata().
*/
func (self *TraitObject) GetQdata(quark C.GQuark) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_object_get_qdata(self.CPointer, quark)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

// g_object_get_valist is not generated due to varargs

/*
Checks whether @object has a [floating][floating-ref] reference.
*/
func (self *TraitObject) IsFloating() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_object_is_floating((C.gpointer)(unsafe.Pointer(self.CPointer)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Emits a "notify" signal for the property @property_name on @object.

When possible, eg. when signaling a property change from within the class
that registered the property, you should use g_object_notify_by_pspec()
instead.

Note that emission of the notify signal may be blocked with
g_object_freeze_notify(). In this case, the signal emissions are queued
and will be emitted (in reverse order) when g_object_thaw_notify() is
called.
*/
func (self *TraitObject) Notify(property_name string) {
	__cgo__property_name := (*C.gchar)(unsafe.Pointer(C.CString(property_name)))
	C.g_object_notify(self.CPointer, __cgo__property_name)
	C.free(unsafe.Pointer(__cgo__property_name))
	return
}

/*
Emits a "notify" signal for the property specified by @pspec on @object.

This function omits the property name lookup, hence it is faster than
g_object_notify().

One way to avoid using g_object_notify() from within the
class that registered the properties, and using g_object_notify_by_pspec()
instead, is to store the GParamSpec used with
g_object_class_install_property() inside a static array, e.g.:

|[<!-- language="C" -->
  enum
  {
    PROP_0,
    PROP_FOO,
    PROP_LAST
  };

  static GParamSpec *properties[PROP_LAST];

  static void
  my_object_class_init (MyObjectClass *klass)
  {
    properties[PROP_FOO] = g_param_spec_int ("foo", "Foo", "The foo",
                                             0, 100,
                                             50,
                                             G_PARAM_READWRITE);
    g_object_class_install_property (gobject_class,
                                     PROP_FOO,
                                     properties[PROP_FOO]);
  }
]|

and then notify a change on the "foo" property with:

|[<!-- language="C" -->
  g_object_notify_by_pspec (self, properties[PROP_FOO]);
]|
*/
func (self *TraitObject) NotifyByPspec(pspec IsParamSpec) {
	var __cgo__pspec *C.GParamSpec
	if pspec != nil {
		__cgo__pspec = pspec.GetParamSpecPointer()
	}
	C.g_object_notify_by_pspec(self.CPointer, __cgo__pspec)
	return
}

/*
Increases the reference count of @object.
*/
func (self *TraitObject) Ref() (return__ *Object) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_object_ref((C.gpointer)(unsafe.Pointer(self.CPointer)))
	return__ = NewObjectFromCPointer(unsafe.Pointer(__cgo__return__))
	return
}

/*
Increase the reference count of @object, and possibly remove the
[floating][floating-ref] reference, if @object has a floating reference.

In other words, if the object is floating, then this call "assumes
ownership" of the floating reference, converting it to a normal
reference by clearing the floating flag while leaving the reference
count unchanged.  If the object is not floating, then this call
adds a new normal reference increasing the reference count by one.
*/
func (self *TraitObject) RefSink() (return__ *Object) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_object_ref_sink((C.gpointer)(unsafe.Pointer(self.CPointer)))
	return__ = NewObjectFromCPointer(unsafe.Pointer(__cgo__return__))
	return
}

/*
Removes a reference added with g_object_add_toggle_ref(). The
reference count of the object is decreased by one.
*/
func (self *TraitObject) RemoveToggleRef(notify C.GToggleNotify, data unsafe.Pointer) {
	C.g_object_remove_toggle_ref(self.CPointer, notify, (C.gpointer)(data))
	return
}

// g_object_remove_weak_pointer is not generated due to inout param

/*
Compares the user data for the key @key on @object with
@oldval, and if they are the same, replaces @oldval with
@newval.

This is like a typical atomic compare-and-exchange
operation, for user data on an object.

If the previous value was replaced then ownership of the
old value (@oldval) is passed to the caller, including
the registered destroy notify for it (passed out in @old_destroy).
Its up to the caller to free this as he wishes, which may
or may not include using @old_destroy as sometimes replacement
should not destroy the object in the normal way.
*/
func (self *TraitObject) ReplaceData(key string, oldval unsafe.Pointer, newval unsafe.Pointer, destroy C.GDestroyNotify, old_destroy *C.GDestroyNotify) (return__ bool) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_object_replace_data(self.CPointer, __cgo__key, (C.gpointer)(oldval), (C.gpointer)(newval), destroy, old_destroy)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Compares the user data for the key @quark on @object with
@oldval, and if they are the same, replaces @oldval with
@newval.

This is like a typical atomic compare-and-exchange
operation, for user data on an object.

If the previous value was replaced then ownership of the
old value (@oldval) is passed to the caller, including
the registered destroy notify for it (passed out in @old_destroy).
Its up to the caller to free this as he wishes, which may
or may not include using @old_destroy as sometimes replacement
should not destroy the object in the normal way.
*/
func (self *TraitObject) ReplaceQdata(quark C.GQuark, oldval unsafe.Pointer, newval unsafe.Pointer, destroy C.GDestroyNotify, old_destroy *C.GDestroyNotify) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_object_replace_qdata(self.CPointer, quark, (C.gpointer)(oldval), (C.gpointer)(newval), destroy, old_destroy)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Releases all references to other objects. This can be used to break
reference cycles.

This functions should only be called from object system implementations.
*/
func (self *TraitObject) RunDispose() {
	C.g_object_run_dispose(self.CPointer)
	return
}

/*
Each object carries around a table of associations from
strings to pointers.  This function lets you set an association.

If the object already had an association with that name,
the old association will be destroyed.
*/
func (self *TraitObject) SetData(key string, data unsafe.Pointer) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	C.g_object_set_data(self.CPointer, __cgo__key, (C.gpointer)(data))
	C.free(unsafe.Pointer(__cgo__key))
	return
}

/*
Like g_object_set_data() except it adds notification
for when the association is destroyed, either by setting it
to a different value or when the object is destroyed.

Note that the @destroy callback is not called if @data is %NULL.
*/
func (self *TraitObject) SetDataFull(key string, data unsafe.Pointer, destroy C.GDestroyNotify) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	C.g_object_set_data_full(self.CPointer, __cgo__key, (C.gpointer)(data), destroy)
	C.free(unsafe.Pointer(__cgo__key))
	return
}

/*
Sets a property on an object.
*/
func (self *TraitObject) SetProperty(property_name string, value *C.GValue) {
	__cgo__property_name := (*C.gchar)(unsafe.Pointer(C.CString(property_name)))
	C.g_object_set_property(self.CPointer, __cgo__property_name, value)
	C.free(unsafe.Pointer(__cgo__property_name))
	return
}

/*
This sets an opaque, named pointer on an object.
The name is specified through a #GQuark (retrived e.g. via
g_quark_from_static_string()), and the pointer
can be gotten back from the @object with g_object_get_qdata()
until the @object is finalized.
Setting a previously set user data pointer, overrides (frees)
the old pointer set, using #NULL as pointer essentially
removes the data stored.
*/
func (self *TraitObject) SetQdata(quark C.GQuark, data unsafe.Pointer) {
	C.g_object_set_qdata(self.CPointer, quark, (C.gpointer)(data))
	return
}

/*
This function works like g_object_set_qdata(), but in addition,
a void (*destroy) (gpointer) function may be specified which is
called with @data as argument when the @object is finalized, or
the data is being overwritten by a call to g_object_set_qdata()
with the same @quark.
*/
func (self *TraitObject) SetQdataFull(quark C.GQuark, data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.g_object_set_qdata_full(self.CPointer, quark, (C.gpointer)(data), destroy)
	return
}

// g_object_set_valist is not generated due to varargs

/*
Remove a specified datum from the object's data associations,
without invoking the association's destroy handler.
*/
func (self *TraitObject) StealData(key string) (return__ unsafe.Pointer) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_object_steal_data(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
This function gets back user data pointers stored via
g_object_set_qdata() and removes the @data from object
without invoking its destroy() function (if any was
set).
Usually, calling this function is only required to update
user data pointers with a destroy notifier, for example:
|[<!-- language="C" -->
void
object_add_to_user_list (GObject     *object,
                         const gchar *new_string)
{
  // the quark, naming the object data
  GQuark quark_string_list = g_quark_from_static_string ("my-string-list");
  // retrive the old string list
  GList *list = g_object_steal_qdata (object, quark_string_list);

  // prepend new string
  list = g_list_prepend (list, g_strdup (new_string));
  // this changed 'list', so we need to set it again
  g_object_set_qdata_full (object, quark_string_list, list, free_string_list);
}
static void
free_string_list (gpointer data)
{
  GList *node, *list = data;

  for (node = list; node; node = node->next)
    g_free (node->data);
  g_list_free (list);
}
]|
Using g_object_get_qdata() in the above example, instead of
g_object_steal_qdata() would have left the destroy function set,
and thus the partial string list would have been freed upon
g_object_set_qdata_full().
*/
func (self *TraitObject) StealQdata(quark C.GQuark) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_object_steal_qdata(self.CPointer, quark)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Reverts the effect of a previous call to
g_object_freeze_notify(). The freeze count is decreased on @object
and when it reaches zero, queued "notify" signals are emitted.

Duplicate notifications for each property are squashed so that at most one
#GObject::notify signal is emitted for each property, in the reverse order
in which they have been queued.

It is an error to call this function when the freeze count is zero.
*/
func (self *TraitObject) ThawNotify() {
	C.g_object_thaw_notify(self.CPointer)
	return
}

/*
Decreases the reference count of @object. When its reference count
drops to 0, the object is finalized (i.e. its memory is freed).
*/
func (self *TraitObject) Unref() {
	C.g_object_unref((C.gpointer)(unsafe.Pointer(self.CPointer)))
	return
}

/*
This function essentially limits the life time of the @closure to
the life time of the object. That is, when the object is finalized,
the @closure is invalidated by calling g_closure_invalidate() on
it, in order to prevent invocations of the closure with a finalized
(nonexisting) object. Also, g_object_ref() and g_object_unref() are
added as marshal guards to the @closure, to ensure that an extra
reference count is held on @object during invocation of the
@closure.  Usually, this function will be called on closures that
use this @object as closure data.
*/
func (self *TraitObject) WatchClosure(closure *C.GClosure) {
	C.g_object_watch_closure(self.CPointer, closure)
	return
}

/*
Adds a weak reference callback to an object. Weak references are
used for notification when an object is finalized. They are called
"weak references" because they allow you to safely hold a pointer
to an object without calling g_object_ref() (g_object_ref() adds a
strong reference, that is, forces the object to stay alive).

Note that the weak references created by this method are not
thread-safe: they cannot safely be used in one thread if the
object's last g_object_unref() might happen in another thread.
Use #GWeakRef if thread-safety is required.
*/
func (self *TraitObject) WeakRef(notify C.GWeakNotify, data unsafe.Pointer) {
	C.g_object_weak_ref(self.CPointer, notify, (C.gpointer)(data))
	return
}

/*
Removes a weak reference callback to an object.
*/
func (self *TraitObject) WeakUnref(notify C.GWeakNotify, data unsafe.Pointer) {
	C.g_object_weak_unref(self.CPointer, notify, (C.gpointer)(data))
	return
}

type TraitParamSpec struct{ CPointer *C.GParamSpec }
type IsParamSpec interface {
	GetParamSpecPointer() *C.GParamSpec
}

func (self *TraitParamSpec) GetParamSpecPointer() *C.GParamSpec {
	return self.CPointer
}
func NewTraitParamSpec(p unsafe.Pointer) *TraitParamSpec {
	return &TraitParamSpec{(*C.GParamSpec)(p)}
}

/*
Get the short description of a #GParamSpec.
*/
func (self *TraitParamSpec) GetBlurb() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_param_spec_get_blurb(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

func (self *TraitParamSpec) GetDefaultValue() (return__ *C.GValue) {
	return__ = C.g_param_spec_get_default_value(self.CPointer)
	return
}

/*
Get the name of a #GParamSpec.

The name is always an "interned" string (as per g_intern_string()).
This allows for pointer-value comparisons.
*/
func (self *TraitParamSpec) GetName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_param_spec_get_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Get the nickname of a #GParamSpec.
*/
func (self *TraitParamSpec) GetNick() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_param_spec_get_nick(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets back user data pointers stored via g_param_spec_set_qdata().
*/
func (self *TraitParamSpec) GetQdata(quark C.GQuark) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_param_spec_get_qdata(self.CPointer, quark)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
If the paramspec redirects operations to another paramspec,
returns that paramspec. Redirect is used typically for
providing a new implementation of a property in a derived
type while preserving all the properties from the parent
type. Redirection is established by creating a property
of type #GParamSpecOverride. See g_object_class_override_property()
for an example of the use of this capability.
*/
func (self *TraitParamSpec) GetRedirectTarget() (return__ *ParamSpec) {
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_get_redirect_target(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Increments the reference count of @pspec.
*/
func (self *TraitParamSpec) Ref() (return__ *ParamSpec) {
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_ref(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Convenience function to ref and sink a #GParamSpec.
*/
func (self *TraitParamSpec) RefSink() (return__ *ParamSpec) {
	var __cgo__return__ *C.GParamSpec
	__cgo__return__ = C.g_param_spec_ref_sink(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewParamSpecFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Sets an opaque, named pointer on a #GParamSpec. The name is
specified through a #GQuark (retrieved e.g. via
g_quark_from_static_string()), and the pointer can be gotten back
from the @pspec with g_param_spec_get_qdata().  Setting a
previously set user data pointer, overrides (frees) the old pointer
set, using %NULL as pointer essentially removes the data stored.
*/
func (self *TraitParamSpec) SetQdata(quark C.GQuark, data unsafe.Pointer) {
	C.g_param_spec_set_qdata(self.CPointer, quark, (C.gpointer)(data))
	return
}

/*
This function works like g_param_spec_set_qdata(), but in addition,
a `void (*destroy) (gpointer)` function may be
specified which is called with @data as argument when the @pspec is
finalized, or the data is being overwritten by a call to
g_param_spec_set_qdata() with the same @quark.
*/
func (self *TraitParamSpec) SetQdataFull(quark C.GQuark, data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.g_param_spec_set_qdata_full(self.CPointer, quark, (C.gpointer)(data), destroy)
	return
}

/*
The initial reference count of a newly created #GParamSpec is 1,
even though no one has explicitly called g_param_spec_ref() on it
yet. So the initial reference count is flagged as "floating", until
someone calls `g_param_spec_ref (pspec); g_param_spec_sink
(pspec);` in sequence on it, taking over the initial
reference count (thus ending up with a @pspec that has a reference
count of 1 still, but is not flagged "floating" anymore).
*/
func (self *TraitParamSpec) Sink() {
	C.g_param_spec_sink(self.CPointer)
	return
}

/*
Gets back user data pointers stored via g_param_spec_set_qdata()
and removes the @data from @pspec without invoking its destroy()
function (if any was set).  Usually, calling this function is only
required to update user data pointers with a destroy notifier.
*/
func (self *TraitParamSpec) StealQdata(quark C.GQuark) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_param_spec_steal_qdata(self.CPointer, quark)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Decrements the reference count of a @pspec.
*/
func (self *TraitParamSpec) Unref() {
	C.g_param_spec_unref(self.CPointer)
	return
}

type TraitParamSpecBoolean struct{ CPointer *C.GParamSpecBoolean }
type IsParamSpecBoolean interface {
	GetParamSpecBooleanPointer() *C.GParamSpecBoolean
}

func (self *TraitParamSpecBoolean) GetParamSpecBooleanPointer() *C.GParamSpecBoolean {
	return self.CPointer
}
func NewTraitParamSpecBoolean(p unsafe.Pointer) *TraitParamSpecBoolean {
	return &TraitParamSpecBoolean{(*C.GParamSpecBoolean)(p)}
}

type TraitParamSpecBoxed struct{ CPointer *C.GParamSpecBoxed }
type IsParamSpecBoxed interface {
	GetParamSpecBoxedPointer() *C.GParamSpecBoxed
}

func (self *TraitParamSpecBoxed) GetParamSpecBoxedPointer() *C.GParamSpecBoxed {
	return self.CPointer
}
func NewTraitParamSpecBoxed(p unsafe.Pointer) *TraitParamSpecBoxed {
	return &TraitParamSpecBoxed{(*C.GParamSpecBoxed)(p)}
}

type TraitParamSpecChar struct{ CPointer *C.GParamSpecChar }
type IsParamSpecChar interface {
	GetParamSpecCharPointer() *C.GParamSpecChar
}

func (self *TraitParamSpecChar) GetParamSpecCharPointer() *C.GParamSpecChar {
	return self.CPointer
}
func NewTraitParamSpecChar(p unsafe.Pointer) *TraitParamSpecChar {
	return &TraitParamSpecChar{(*C.GParamSpecChar)(p)}
}

type TraitParamSpecDouble struct{ CPointer *C.GParamSpecDouble }
type IsParamSpecDouble interface {
	GetParamSpecDoublePointer() *C.GParamSpecDouble
}

func (self *TraitParamSpecDouble) GetParamSpecDoublePointer() *C.GParamSpecDouble {
	return self.CPointer
}
func NewTraitParamSpecDouble(p unsafe.Pointer) *TraitParamSpecDouble {
	return &TraitParamSpecDouble{(*C.GParamSpecDouble)(p)}
}

type TraitParamSpecEnum struct{ CPointer *C.GParamSpecEnum }
type IsParamSpecEnum interface {
	GetParamSpecEnumPointer() *C.GParamSpecEnum
}

func (self *TraitParamSpecEnum) GetParamSpecEnumPointer() *C.GParamSpecEnum {
	return self.CPointer
}
func NewTraitParamSpecEnum(p unsafe.Pointer) *TraitParamSpecEnum {
	return &TraitParamSpecEnum{(*C.GParamSpecEnum)(p)}
}

type TraitParamSpecFlags struct{ CPointer *C.GParamSpecFlags }
type IsParamSpecFlags interface {
	GetParamSpecFlagsPointer() *C.GParamSpecFlags
}

func (self *TraitParamSpecFlags) GetParamSpecFlagsPointer() *C.GParamSpecFlags {
	return self.CPointer
}
func NewTraitParamSpecFlags(p unsafe.Pointer) *TraitParamSpecFlags {
	return &TraitParamSpecFlags{(*C.GParamSpecFlags)(p)}
}

type TraitParamSpecFloat struct{ CPointer *C.GParamSpecFloat }
type IsParamSpecFloat interface {
	GetParamSpecFloatPointer() *C.GParamSpecFloat
}

func (self *TraitParamSpecFloat) GetParamSpecFloatPointer() *C.GParamSpecFloat {
	return self.CPointer
}
func NewTraitParamSpecFloat(p unsafe.Pointer) *TraitParamSpecFloat {
	return &TraitParamSpecFloat{(*C.GParamSpecFloat)(p)}
}

type TraitParamSpecGType struct{ CPointer *C.GParamSpecGType }
type IsParamSpecGType interface {
	GetParamSpecGTypePointer() *C.GParamSpecGType
}

func (self *TraitParamSpecGType) GetParamSpecGTypePointer() *C.GParamSpecGType {
	return self.CPointer
}
func NewTraitParamSpecGType(p unsafe.Pointer) *TraitParamSpecGType {
	return &TraitParamSpecGType{(*C.GParamSpecGType)(p)}
}

type TraitParamSpecInt struct{ CPointer *C.GParamSpecInt }
type IsParamSpecInt interface {
	GetParamSpecIntPointer() *C.GParamSpecInt
}

func (self *TraitParamSpecInt) GetParamSpecIntPointer() *C.GParamSpecInt {
	return self.CPointer
}
func NewTraitParamSpecInt(p unsafe.Pointer) *TraitParamSpecInt {
	return &TraitParamSpecInt{(*C.GParamSpecInt)(p)}
}

type TraitParamSpecInt64 struct{ CPointer *C.GParamSpecInt64 }
type IsParamSpecInt64 interface {
	GetParamSpecInt64Pointer() *C.GParamSpecInt64
}

func (self *TraitParamSpecInt64) GetParamSpecInt64Pointer() *C.GParamSpecInt64 {
	return self.CPointer
}
func NewTraitParamSpecInt64(p unsafe.Pointer) *TraitParamSpecInt64 {
	return &TraitParamSpecInt64{(*C.GParamSpecInt64)(p)}
}

type TraitParamSpecLong struct{ CPointer *C.GParamSpecLong }
type IsParamSpecLong interface {
	GetParamSpecLongPointer() *C.GParamSpecLong
}

func (self *TraitParamSpecLong) GetParamSpecLongPointer() *C.GParamSpecLong {
	return self.CPointer
}
func NewTraitParamSpecLong(p unsafe.Pointer) *TraitParamSpecLong {
	return &TraitParamSpecLong{(*C.GParamSpecLong)(p)}
}

type TraitParamSpecObject struct{ CPointer *C.GParamSpecObject }
type IsParamSpecObject interface {
	GetParamSpecObjectPointer() *C.GParamSpecObject
}

func (self *TraitParamSpecObject) GetParamSpecObjectPointer() *C.GParamSpecObject {
	return self.CPointer
}
func NewTraitParamSpecObject(p unsafe.Pointer) *TraitParamSpecObject {
	return &TraitParamSpecObject{(*C.GParamSpecObject)(p)}
}

type TraitParamSpecOverride struct{ CPointer *C.GParamSpecOverride }
type IsParamSpecOverride interface {
	GetParamSpecOverridePointer() *C.GParamSpecOverride
}

func (self *TraitParamSpecOverride) GetParamSpecOverridePointer() *C.GParamSpecOverride {
	return self.CPointer
}
func NewTraitParamSpecOverride(p unsafe.Pointer) *TraitParamSpecOverride {
	return &TraitParamSpecOverride{(*C.GParamSpecOverride)(p)}
}

type TraitParamSpecParam struct{ CPointer *C.GParamSpecParam }
type IsParamSpecParam interface {
	GetParamSpecParamPointer() *C.GParamSpecParam
}

func (self *TraitParamSpecParam) GetParamSpecParamPointer() *C.GParamSpecParam {
	return self.CPointer
}
func NewTraitParamSpecParam(p unsafe.Pointer) *TraitParamSpecParam {
	return &TraitParamSpecParam{(*C.GParamSpecParam)(p)}
}

type TraitParamSpecPointer struct{ CPointer *C.GParamSpecPointer }
type IsParamSpecPointer interface {
	GetParamSpecPointerPointer() *C.GParamSpecPointer
}

func (self *TraitParamSpecPointer) GetParamSpecPointerPointer() *C.GParamSpecPointer {
	return self.CPointer
}
func NewTraitParamSpecPointer(p unsafe.Pointer) *TraitParamSpecPointer {
	return &TraitParamSpecPointer{(*C.GParamSpecPointer)(p)}
}

type TraitParamSpecString struct{ CPointer *C.GParamSpecString }
type IsParamSpecString interface {
	GetParamSpecStringPointer() *C.GParamSpecString
}

func (self *TraitParamSpecString) GetParamSpecStringPointer() *C.GParamSpecString {
	return self.CPointer
}
func NewTraitParamSpecString(p unsafe.Pointer) *TraitParamSpecString {
	return &TraitParamSpecString{(*C.GParamSpecString)(p)}
}

type TraitParamSpecUChar struct{ CPointer *C.GParamSpecUChar }
type IsParamSpecUChar interface {
	GetParamSpecUCharPointer() *C.GParamSpecUChar
}

func (self *TraitParamSpecUChar) GetParamSpecUCharPointer() *C.GParamSpecUChar {
	return self.CPointer
}
func NewTraitParamSpecUChar(p unsafe.Pointer) *TraitParamSpecUChar {
	return &TraitParamSpecUChar{(*C.GParamSpecUChar)(p)}
}

type TraitParamSpecUInt struct{ CPointer *C.GParamSpecUInt }
type IsParamSpecUInt interface {
	GetParamSpecUIntPointer() *C.GParamSpecUInt
}

func (self *TraitParamSpecUInt) GetParamSpecUIntPointer() *C.GParamSpecUInt {
	return self.CPointer
}
func NewTraitParamSpecUInt(p unsafe.Pointer) *TraitParamSpecUInt {
	return &TraitParamSpecUInt{(*C.GParamSpecUInt)(p)}
}

type TraitParamSpecUInt64 struct{ CPointer *C.GParamSpecUInt64 }
type IsParamSpecUInt64 interface {
	GetParamSpecUInt64Pointer() *C.GParamSpecUInt64
}

func (self *TraitParamSpecUInt64) GetParamSpecUInt64Pointer() *C.GParamSpecUInt64 {
	return self.CPointer
}
func NewTraitParamSpecUInt64(p unsafe.Pointer) *TraitParamSpecUInt64 {
	return &TraitParamSpecUInt64{(*C.GParamSpecUInt64)(p)}
}

type TraitParamSpecULong struct{ CPointer *C.GParamSpecULong }
type IsParamSpecULong interface {
	GetParamSpecULongPointer() *C.GParamSpecULong
}

func (self *TraitParamSpecULong) GetParamSpecULongPointer() *C.GParamSpecULong {
	return self.CPointer
}
func NewTraitParamSpecULong(p unsafe.Pointer) *TraitParamSpecULong {
	return &TraitParamSpecULong{(*C.GParamSpecULong)(p)}
}

type TraitParamSpecUnichar struct{ CPointer *C.GParamSpecUnichar }
type IsParamSpecUnichar interface {
	GetParamSpecUnicharPointer() *C.GParamSpecUnichar
}

func (self *TraitParamSpecUnichar) GetParamSpecUnicharPointer() *C.GParamSpecUnichar {
	return self.CPointer
}
func NewTraitParamSpecUnichar(p unsafe.Pointer) *TraitParamSpecUnichar {
	return &TraitParamSpecUnichar{(*C.GParamSpecUnichar)(p)}
}

type TraitParamSpecValueArray struct{ CPointer *C.GParamSpecValueArray }
type IsParamSpecValueArray interface {
	GetParamSpecValueArrayPointer() *C.GParamSpecValueArray
}

func (self *TraitParamSpecValueArray) GetParamSpecValueArrayPointer() *C.GParamSpecValueArray {
	return self.CPointer
}
func NewTraitParamSpecValueArray(p unsafe.Pointer) *TraitParamSpecValueArray {
	return &TraitParamSpecValueArray{(*C.GParamSpecValueArray)(p)}
}

type TraitParamSpecVariant struct{ CPointer *C.GParamSpecVariant }
type IsParamSpecVariant interface {
	GetParamSpecVariantPointer() *C.GParamSpecVariant
}

func (self *TraitParamSpecVariant) GetParamSpecVariantPointer() *C.GParamSpecVariant {
	return self.CPointer
}
func NewTraitParamSpecVariant(p unsafe.Pointer) *TraitParamSpecVariant {
	return &TraitParamSpecVariant{(*C.GParamSpecVariant)(p)}
}

type TraitTypeModule struct{ CPointer *C.GTypeModule }
type IsTypeModule interface {
	GetTypeModulePointer() *C.GTypeModule
}

func (self *TraitTypeModule) GetTypeModulePointer() *C.GTypeModule {
	return self.CPointer
}
func NewTraitTypeModule(p unsafe.Pointer) *TraitTypeModule {
	return &TraitTypeModule{(*C.GTypeModule)(p)}
}

/*
Registers an additional interface for a type, whose interface lives
in the given type plugin. If the interface was already registered
for the type in this plugin, nothing will be done.

As long as any instances of the type exist, the type plugin will
not be unloaded.
*/
func (self *TraitTypeModule) AddInterface(instance_type C.GType, interface_type C.GType, interface_info *C.GInterfaceInfo) {
	C.g_type_module_add_interface(self.CPointer, instance_type, interface_type, interface_info)
	return
}

/*
Looks up or registers an enumeration that is implemented with a particular
type plugin. If a type with name @type_name was previously registered,
the #GType identifier for the type is returned, otherwise the type
is newly registered, and the resulting #GType identifier returned.

As long as any instances of the type exist, the type plugin will
not be unloaded.
*/
func (self *TraitTypeModule) RegisterEnum(name string, const_static_values *C.GEnumValue) (return__ C.GType) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	return__ = C.g_type_module_register_enum(self.CPointer, __cgo__name, const_static_values)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Looks up or registers a flags type that is implemented with a particular
type plugin. If a type with name @type_name was previously registered,
the #GType identifier for the type is returned, otherwise the type
is newly registered, and the resulting #GType identifier returned.

As long as any instances of the type exist, the type plugin will
not be unloaded.
*/
func (self *TraitTypeModule) RegisterFlags(name string, const_static_values *C.GFlagsValue) (return__ C.GType) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	return__ = C.g_type_module_register_flags(self.CPointer, __cgo__name, const_static_values)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Looks up or registers a type that is implemented with a particular
type plugin. If a type with name @type_name was previously registered,
the #GType identifier for the type is returned, otherwise the type
is newly registered, and the resulting #GType identifier returned.

When reregistering a type (typically because a module is unloaded
then reloaded, and reinitialized), @module and @parent_type must
be the same as they were previously.

As long as any instances of the type exist, the type plugin will
not be unloaded.
*/
func (self *TraitTypeModule) RegisterType(parent_type C.GType, type_name string, type_info *C.GTypeInfo, flags C.GTypeFlags) (return__ C.GType) {
	__cgo__type_name := (*C.gchar)(unsafe.Pointer(C.CString(type_name)))
	return__ = C.g_type_module_register_type(self.CPointer, parent_type, __cgo__type_name, type_info, flags)
	C.free(unsafe.Pointer(__cgo__type_name))
	return
}

/*
Sets the name for a #GTypeModule
*/
func (self *TraitTypeModule) SetName(name string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.g_type_module_set_name(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Decreases the use count of a #GTypeModule by one. If the
result is zero, the module will be unloaded. (However, the
#GTypeModule will not be freed, and types associated with the
#GTypeModule are not unregistered. Once a #GTypeModule is
initialized, it must exist forever.)
*/
func (self *TraitTypeModule) Unuse() {
	C.g_type_module_unuse(self.CPointer)
	return
}

/*
Increases the use count of a #GTypeModule by one. If the
use count was zero before, the plugin will be loaded.
If loading the plugin fails, the use count is reset to
its prior value.
*/
func (self *TraitTypeModule) Use() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_type_module_use(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}
