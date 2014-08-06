package gio

/*
#include <gio/gio.h>
#include <gio/gunixoutputstream.h>
#include <gio/gunixinputstream.h>
#include <gio/gunixcredentialsmessage.h>
#include <gio/gunixmounts.h>
#include <gio/gdesktopappinfo.h>
#include <gio/gunixfdlist.h>
#include <gio/gunixsocketaddress.h>
#include <gio/gunixfdmessage.h>
#include <gio/gnetworking.h>
#include <gio/gunixconnection.h>
#include <stdlib.h>
#cgo pkg-config: gio-2.0 gio-unix-2.0
*/
import "C"

import "github.com/reusee/ggir/gobject"
import "github.com/reusee/ggir/glib"
import (
	"errors"
	"reflect"
	"unsafe"
)

func init() {
	var _ unsafe.Pointer
	var _ reflect.SliceHeader
	_ = errors.New("")
}

var _ = gobject.UnusedFix_
var _ = glib.UnusedFix_

/*
Checks if @action_name is valid.

@action_name is valid if it consists only of alphanumeric characters,
plus '-' and '.'.  The empty string is not a valid action name.

It is an error to call this function with a non-utf8 @action_name.
@action_name must not be %NULL.
*/
func ActionNameIsValid(action_name string) (return__ bool) {
	__cgo__action_name := (*C.gchar)(unsafe.Pointer(C.CString(action_name)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_action_name_is_valid(__cgo__action_name)
	C.free(unsafe.Pointer(__cgo__action_name))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Parses a detailed action name into its separate name and target
components.

Detailed action names can have three formats.

The first format is used to represent an action name with no target
value and consists of just an action name containing no whitespace
nor the characters ':', '(' or ')'.  For example: "app.action".

The second format is used to represent an action with a target value
that is a non-empty string consisting only of alphanumerics, plus '-'
and '.'.  In that case, the action name and target value are
separated by a double colon ("::").  For example:
"app.action::target".

The third format is used to represent an action with any type of
target value, including strings.  The target value follows the action
name, surrounded in parens.  For example: "app.action(42)".  The
target value is parsed using g_variant_parse().  If a tuple-typed
value is desired, it must be specified in the same way, resulting in
two sets of parens, for example: "app.action((1,2,3))".  A string
target can be specified this way as well: "app.action('target')".
For strings, this third format must be used if * target value is
empty or contains characters other than alphanumerics, '-' and '.'.
*/
func ActionParseDetailedName(detailed_name string) (action_name string, target_value *C.GVariant, return__ bool, __err__ error) {
	__cgo__detailed_name := (*C.gchar)(unsafe.Pointer(C.CString(detailed_name)))
	var __cgo__action_name *C.gchar
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_action_parse_detailed_name(__cgo__detailed_name, &__cgo__action_name, &target_value, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__detailed_name))
	action_name = C.GoString((*C.char)(unsafe.Pointer(__cgo__action_name)))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Formats a detailed action name from @action_name and @target_value.

It is an error to call this function with an invalid action name.

This function is the opposite of
g_action_parse_detailed_action_name().  It will produce a string that
can be parsed back to the @action_name and @target_value by that
function.

See that function for the types of strings that will be printed by
this function.
*/
func ActionPrintDetailedName(action_name string, target_value *C.GVariant) (return__ string) {
	__cgo__action_name := (*C.gchar)(unsafe.Pointer(C.CString(action_name)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_action_print_detailed_name(__cgo__action_name, target_value)
	C.free(unsafe.Pointer(__cgo__action_name))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Creates a new #GAppInfo from the given information.

Note that for @commandline, the quoting rules of the Exec key of the
[freedesktop.org Desktop Entry Specification](http://freedesktop.org/Standards/desktop-entry-spec)
are applied. For example, if the @commandline contains
percent-encoded URIs, the percent-character must be doubled in order to prevent it from
being swallowed by Exec key unquoting. See the specification for exact quoting rules.
*/
func AppInfoCreateFromCommandline(commandline string, application_name string, flags C.GAppInfoCreateFlags) (return__ *C.GAppInfo, __err__ error) {
	__cgo__commandline := C.CString(commandline)
	__cgo__application_name := C.CString(application_name)
	var __cgo_error__ *C.GError
	return__ = C.g_app_info_create_from_commandline(__cgo__commandline, __cgo__application_name, flags, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__commandline))
	C.free(unsafe.Pointer(__cgo__application_name))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets a list of all of the applications currently registered
on this system.

For desktop files, this includes applications that have
`NoDisplay=true` set or are excluded from display by means
of `OnlyShowIn` or `NotShowIn`. See g_app_info_should_show().
The returned list does not include applications which have
the `Hidden` key set.
*/
func AppInfoGetAll() (return__ *C.GList) {
	return__ = C.g_app_info_get_all()
	return
}

/*
Gets a list of all #GAppInfos for a given content type,
including the recommended and fallback #GAppInfos. See
g_app_info_get_recommended_for_type() and
g_app_info_get_fallback_for_type().
*/
func AppInfoGetAllForType(content_type string) (return__ *C.GList) {
	__cgo__content_type := C.CString(content_type)
	return__ = C.g_app_info_get_all_for_type(__cgo__content_type)
	C.free(unsafe.Pointer(__cgo__content_type))
	return
}

/*
Gets the default #GAppInfo for a given content type.
*/
func AppInfoGetDefaultForType(content_type string, must_support_uris bool) (return__ *C.GAppInfo) {
	__cgo__content_type := C.CString(content_type)
	__cgo__must_support_uris := C.gboolean(0)
	if must_support_uris {
		__cgo__must_support_uris = C.gboolean(1)
	}
	return__ = C.g_app_info_get_default_for_type(__cgo__content_type, __cgo__must_support_uris)
	C.free(unsafe.Pointer(__cgo__content_type))
	return
}

/*
Gets the default application for handling URIs with
the given URI scheme. A URI scheme is the initial part
of the URI, up to but not including the ':', e.g. "http",
"ftp" or "sip".
*/
func AppInfoGetDefaultForUriScheme(uri_scheme string) (return__ *C.GAppInfo) {
	__cgo__uri_scheme := C.CString(uri_scheme)
	return__ = C.g_app_info_get_default_for_uri_scheme(__cgo__uri_scheme)
	C.free(unsafe.Pointer(__cgo__uri_scheme))
	return
}

/*
Gets a list of fallback #GAppInfos for a given content type, i.e.
those applications which claim to support the given content type
by MIME type subclassing and not directly.
*/
func AppInfoGetFallbackForType(content_type string) (return__ *C.GList) {
	__cgo__content_type := (*C.gchar)(unsafe.Pointer(C.CString(content_type)))
	return__ = C.g_app_info_get_fallback_for_type(__cgo__content_type)
	C.free(unsafe.Pointer(__cgo__content_type))
	return
}

/*
Gets a list of recommended #GAppInfos for a given content type, i.e.
those applications which claim to support the given content type exactly,
and not by MIME type subclassing.
Note that the first application of the list is the last used one, i.e.
the last one for which g_app_info_set_as_last_used_for_type() has been
called.
*/
func AppInfoGetRecommendedForType(content_type string) (return__ *C.GList) {
	__cgo__content_type := (*C.gchar)(unsafe.Pointer(C.CString(content_type)))
	return__ = C.g_app_info_get_recommended_for_type(__cgo__content_type)
	C.free(unsafe.Pointer(__cgo__content_type))
	return
}

/*
Utility function that launches the default application
registered to handle the specified uri. Synchronous I/O
is done on the uri to detect the type of the file if
required.
*/
func AppInfoLaunchDefaultForUri(uri string, launch_context IsAppLaunchContext) (return__ bool, __err__ error) {
	__cgo__uri := C.CString(uri)
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_app_info_launch_default_for_uri(__cgo__uri, launch_context.GetAppLaunchContextPointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__uri))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Removes all changes to the type associations done by
g_app_info_set_as_default_for_type(),
g_app_info_set_as_default_for_extension(),
g_app_info_add_supports_type() or
g_app_info_remove_supports_type().
*/
func AppInfoResetTypeAssociations(content_type string) {
	__cgo__content_type := C.CString(content_type)
	C.g_app_info_reset_type_associations(__cgo__content_type)
	C.free(unsafe.Pointer(__cgo__content_type))
	return
}

/*
Helper function for constructing #GAsyncInitable object. This is
similar to g_object_newv() but also initializes the object asynchronously.

When the initialization is finished, @callback will be called. You can
then call g_async_initable_new_finish() to get the new object and check
for any errors.
*/
func AsyncInitableNewvAsync(object_type C.GType, n_parameters uint, parameters *C.GParameter, io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_async_initable_newv_async(object_type, C.guint(n_parameters), parameters, C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Asynchronously connects to the message bus specified by @bus_type.

When the operation is finished, @callback will be invoked. You can
then call g_bus_get_finish() to get the result of the operation.

This is a asynchronous failable function. See g_bus_get_sync() for
the synchronous version.
*/
func BusGet(bus_type C.GBusType, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_bus_get(bus_type, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes an operation started with g_bus_get().

The returned object is a singleton, that is, shared with other
callers of g_bus_get() and g_bus_get_sync() for @bus_type. In the
event that you need a private message bus connection, use
g_dbus_address_get_for_bus_sync() and
g_dbus_connection_new_for_address().

Note that the returned #GDBusConnection object will (usually) have
the #GDBusConnection:exit-on-close property set to %TRUE.
*/
func BusGetFinish(res *C.GAsyncResult) (return__ *DBusConnection, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GDBusConnection
	__cgo__return__ = C.g_bus_get_finish(res, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewDBusConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Synchronously connects to the message bus specified by @bus_type.
Note that the returned object may shared with other callers,
e.g. if two separate parts of a process calls this function with
the same @bus_type, they will share the same object.

This is a synchronous failable function. See g_bus_get() and
g_bus_get_finish() for the asynchronous version.

The returned object is a singleton, that is, shared with other
callers of g_bus_get() and g_bus_get_sync() for @bus_type. In the
event that you need a private message bus connection, use
g_dbus_address_get_for_bus_sync() and
g_dbus_connection_new_for_address().

Note that the returned #GDBusConnection object will (usually) have
the #GDBusConnection:exit-on-close property set to %TRUE.
*/
func BusGetSync(bus_type C.GBusType, cancellable IsCancellable) (return__ *DBusConnection, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GDBusConnection
	__cgo__return__ = C.g_bus_get_sync(bus_type, cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewDBusConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Starts acquiring @name on the bus specified by @bus_type and calls
@name_acquired_handler and @name_lost_handler when the name is
acquired respectively lost. Callbacks will be invoked in the
[thread-default main context][g-main-context-push-thread-default]
of the thread you are calling this function from.

You are guaranteed that one of the @name_acquired_handler and @name_lost_handler
callbacks will be invoked after calling this function - there are three
possible cases:

- @name_lost_handler with a %NULL connection (if a connection to the bus
  can't be made).

- @bus_acquired_handler then @name_lost_handler (if the name can't be
  obtained)

- @bus_acquired_handler then @name_acquired_handler (if the name was
  obtained).

When you are done owning the name, just call g_bus_unown_name()
with the owner id this function returns.

If the name is acquired or lost (for example another application
could acquire the name if you allow replacement or the application
currently owning the name exits), the handlers are also invoked.
If the #GDBusConnection that is used for attempting to own the name
closes, then @name_lost_handler is invoked since it is no longer
possible for other processes to access the process.

You cannot use g_bus_own_name() several times for the same name (unless
interleaved with calls to g_bus_unown_name()) - only the first call
will work.

Another guarantee is that invocations of @name_acquired_handler
and @name_lost_handler are guaranteed to alternate; that
is, if @name_acquired_handler is invoked then you are
guaranteed that the next time one of the handlers is invoked, it
will be @name_lost_handler. The reverse is also true.

If you plan on exporting objects (using e.g.
g_dbus_connection_register_object()), note that it is generally too late
to export the objects in @name_acquired_handler. Instead, you can do this
in @bus_acquired_handler since you are guaranteed that this will run
before @name is requested from the bus.

This behavior makes it very simple to write applications that wants
to [own names][gdbus-owning-names] and export objects.
Simply register objects to be exported in @bus_acquired_handler and
unregister the objects (if any) in @name_lost_handler.
*/
func BusOwnName(bus_type C.GBusType, name string, flags C.GBusNameOwnerFlags, bus_acquired_handler C.GBusAcquiredCallback, name_acquired_handler C.GBusNameAcquiredCallback, name_lost_handler C.GBusNameLostCallback, user_data unsafe.Pointer, user_data_free_func C.GDestroyNotify) (return__ uint) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_bus_own_name(bus_type, __cgo__name, flags, bus_acquired_handler, name_acquired_handler, name_lost_handler, (C.gpointer)(user_data), user_data_free_func)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = uint(__cgo__return__)
	return
}

/*
Like g_bus_own_name() but takes a #GDBusConnection instead of a
#GBusType.
*/
func BusOwnNameOnConnection(connection IsDBusConnection, name string, flags C.GBusNameOwnerFlags, name_acquired_handler C.GBusNameAcquiredCallback, name_lost_handler C.GBusNameLostCallback, user_data unsafe.Pointer, user_data_free_func C.GDestroyNotify) (return__ uint) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_bus_own_name_on_connection(connection.GetDBusConnectionPointer(), __cgo__name, flags, name_acquired_handler, name_lost_handler, (C.gpointer)(user_data), user_data_free_func)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = uint(__cgo__return__)
	return
}

/*
Version of g_bus_own_name_on_connection() using closures instead of
callbacks for easier binding in other languages.
*/
func BusOwnNameOnConnectionWithClosures(connection IsDBusConnection, name string, flags C.GBusNameOwnerFlags, name_acquired_closure *C.GClosure, name_lost_closure *C.GClosure) (return__ uint) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_bus_own_name_on_connection_with_closures(connection.GetDBusConnectionPointer(), __cgo__name, flags, name_acquired_closure, name_lost_closure)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = uint(__cgo__return__)
	return
}

/*
Version of g_bus_own_name() using closures instead of callbacks for
easier binding in other languages.
*/
func BusOwnNameWithClosures(bus_type C.GBusType, name string, flags C.GBusNameOwnerFlags, bus_acquired_closure *C.GClosure, name_acquired_closure *C.GClosure, name_lost_closure *C.GClosure) (return__ uint) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_bus_own_name_with_closures(bus_type, __cgo__name, flags, bus_acquired_closure, name_acquired_closure, name_lost_closure)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = uint(__cgo__return__)
	return
}

/*
Stops owning a name.
*/
func BusUnownName(owner_id uint) {
	C.g_bus_unown_name(C.guint(owner_id))
	return
}

/*
Stops watching a name.
*/
func BusUnwatchName(watcher_id uint) {
	C.g_bus_unwatch_name(C.guint(watcher_id))
	return
}

/*
Starts watching @name on the bus specified by @bus_type and calls
@name_appeared_handler and @name_vanished_handler when the name is
known to have a owner respectively known to lose its
owner. Callbacks will be invoked in the
[thread-default main context][g-main-context-push-thread-default]
of the thread you are calling this function from.

You are guaranteed that one of the handlers will be invoked after
calling this function. When you are done watching the name, just
call g_bus_unwatch_name() with the watcher id this function
returns.

If the name vanishes or appears (for example the application owning
the name could restart), the handlers are also invoked. If the
#GDBusConnection that is used for watching the name disconnects, then
@name_vanished_handler is invoked since it is no longer
possible to access the name.

Another guarantee is that invocations of @name_appeared_handler
and @name_vanished_handler are guaranteed to alternate; that
is, if @name_appeared_handler is invoked then you are
guaranteed that the next time one of the handlers is invoked, it
will be @name_vanished_handler. The reverse is also true.

This behavior makes it very simple to write applications that want
to take action when a certain [name exists][gdbus-watching-names].
Basically, the application should create object proxies in
@name_appeared_handler and destroy them again (if any) in
@name_vanished_handler.
*/
func BusWatchName(bus_type C.GBusType, name string, flags C.GBusNameWatcherFlags, name_appeared_handler C.GBusNameAppearedCallback, name_vanished_handler C.GBusNameVanishedCallback, user_data unsafe.Pointer, user_data_free_func C.GDestroyNotify) (return__ uint) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_bus_watch_name(bus_type, __cgo__name, flags, name_appeared_handler, name_vanished_handler, (C.gpointer)(user_data), user_data_free_func)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = uint(__cgo__return__)
	return
}

/*
Like g_bus_watch_name() but takes a #GDBusConnection instead of a
#GBusType.
*/
func BusWatchNameOnConnection(connection IsDBusConnection, name string, flags C.GBusNameWatcherFlags, name_appeared_handler C.GBusNameAppearedCallback, name_vanished_handler C.GBusNameVanishedCallback, user_data unsafe.Pointer, user_data_free_func C.GDestroyNotify) (return__ uint) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_bus_watch_name_on_connection(connection.GetDBusConnectionPointer(), __cgo__name, flags, name_appeared_handler, name_vanished_handler, (C.gpointer)(user_data), user_data_free_func)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = uint(__cgo__return__)
	return
}

/*
Version of g_bus_watch_name_on_connection() using closures instead of callbacks for
easier binding in other languages.
*/
func BusWatchNameOnConnectionWithClosures(connection IsDBusConnection, name string, flags C.GBusNameWatcherFlags, name_appeared_closure *C.GClosure, name_vanished_closure *C.GClosure) (return__ uint) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_bus_watch_name_on_connection_with_closures(connection.GetDBusConnectionPointer(), __cgo__name, flags, name_appeared_closure, name_vanished_closure)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = uint(__cgo__return__)
	return
}

/*
Version of g_bus_watch_name() using closures instead of callbacks for
easier binding in other languages.
*/
func BusWatchNameWithClosures(bus_type C.GBusType, name string, flags C.GBusNameWatcherFlags, name_appeared_closure *C.GClosure, name_vanished_closure *C.GClosure) (return__ uint) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_bus_watch_name_with_closures(bus_type, __cgo__name, flags, name_appeared_closure, name_vanished_closure)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = uint(__cgo__return__)
	return
}

/*
Checks if a content type can be executable. Note that for instance
things like text files can be executables (i.e. scripts and batch files).
*/
func ContentTypeCanBeExecutable(type_ string) (return__ bool) {
	__cgo__type_ := (*C.gchar)(unsafe.Pointer(C.CString(type_)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_content_type_can_be_executable(__cgo__type_)
	C.free(unsafe.Pointer(__cgo__type_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Compares two content types for equality.
*/
func ContentTypeEquals(type1 string, type2 string) (return__ bool) {
	__cgo__type1 := (*C.gchar)(unsafe.Pointer(C.CString(type1)))
	__cgo__type2 := (*C.gchar)(unsafe.Pointer(C.CString(type2)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_content_type_equals(__cgo__type1, __cgo__type2)
	C.free(unsafe.Pointer(__cgo__type1))
	C.free(unsafe.Pointer(__cgo__type2))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Tries to find a content type based on the mime type name.
*/
func ContentTypeFromMimeType(mime_type string) (return__ string) {
	__cgo__mime_type := (*C.gchar)(unsafe.Pointer(C.CString(mime_type)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_content_type_from_mime_type(__cgo__mime_type)
	C.free(unsafe.Pointer(__cgo__mime_type))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the human readable description of the content type.
*/
func ContentTypeGetDescription(type_ string) (return__ string) {
	__cgo__type_ := (*C.gchar)(unsafe.Pointer(C.CString(type_)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_content_type_get_description(__cgo__type_)
	C.free(unsafe.Pointer(__cgo__type_))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the generic icon name for a content type.

See the
[shared-mime-info](http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
specification for more on the generic icon name.
*/
func ContentTypeGetGenericIconName(type_ string) (return__ string) {
	__cgo__type_ := (*C.gchar)(unsafe.Pointer(C.CString(type_)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_content_type_get_generic_icon_name(__cgo__type_)
	C.free(unsafe.Pointer(__cgo__type_))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the icon for a content type.
*/
func ContentTypeGetIcon(type_ string) (return__ *C.GIcon) {
	__cgo__type_ := (*C.gchar)(unsafe.Pointer(C.CString(type_)))
	return__ = C.g_content_type_get_icon(__cgo__type_)
	C.free(unsafe.Pointer(__cgo__type_))
	return
}

/*
Gets the mime type for the content type, if one is registered.
*/
func ContentTypeGetMimeType(type_ string) (return__ string) {
	__cgo__type_ := (*C.gchar)(unsafe.Pointer(C.CString(type_)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_content_type_get_mime_type(__cgo__type_)
	C.free(unsafe.Pointer(__cgo__type_))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the symbolic icon for a content type.
*/
func ContentTypeGetSymbolicIcon(type_ string) (return__ *C.GIcon) {
	__cgo__type_ := (*C.gchar)(unsafe.Pointer(C.CString(type_)))
	return__ = C.g_content_type_get_symbolic_icon(__cgo__type_)
	C.free(unsafe.Pointer(__cgo__type_))
	return
}

/*
Guesses the content type based on example data. If the function is
uncertain, @result_uncertain will be set to %TRUE. Either @filename
or @data may be %NULL, in which case the guess will be based solely
on the other argument.
*/
func ContentTypeGuess(filename string, data []byte, data_size int64) (result_uncertain bool, return__ string) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	__header__data := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	var __cgo__result_uncertain C.gboolean
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_content_type_guess(__cgo__filename, (*C.guchar)(unsafe.Pointer(__header__data.Data)), C.gsize(data_size), &__cgo__result_uncertain)
	C.free(unsafe.Pointer(__cgo__filename))
	result_uncertain = __cgo__result_uncertain == C.gboolean(1)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Tries to guess the type of the tree with root @root, by
looking at the files it contains. The result is an array
of content types, with the best guess coming first.

The types returned all have the form x-content/foo, e.g.
x-content/audio-cdda (for audio CDs) or x-content/image-dcf
(for a camera memory card). See the
[shared-mime-info](http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
specification for more on x-content types.

This function is useful in the implementation of
g_mount_guess_content_type().
*/
func ContentTypeGuessForTree(root *C.GFile) (return__ **C.gchar) {
	return__ = C.g_content_type_guess_for_tree(root)
	return
}

/*
Determines if @type is a subset of @supertype.
*/
func ContentTypeIsA(type_ string, supertype string) (return__ bool) {
	__cgo__type_ := (*C.gchar)(unsafe.Pointer(C.CString(type_)))
	__cgo__supertype := (*C.gchar)(unsafe.Pointer(C.CString(supertype)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_content_type_is_a(__cgo__type_, __cgo__supertype)
	C.free(unsafe.Pointer(__cgo__type_))
	C.free(unsafe.Pointer(__cgo__supertype))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if the content type is the generic "unknown" type.
On UNIX this is the "application/octet-stream" mimetype,
while on win32 it is "*".
*/
func ContentTypeIsUnknown(type_ string) (return__ bool) {
	__cgo__type_ := (*C.gchar)(unsafe.Pointer(C.CString(type_)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_content_type_is_unknown(__cgo__type_)
	C.free(unsafe.Pointer(__cgo__type_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets a list of strings containing all the registered content types
known to the system. The list and its data should be freed using
g_list_free_full (list, g_free).
*/
func ContentTypesGetRegistered() (return__ *C.GList) {
	return__ = C.g_content_types_get_registered()
	return
}

/*
Escape @string so it can appear in a D-Bus address as the value
part of a key-value pair.

For instance, if @string is "/run/bus-for-:0",
this function would return "/run/bus-for-%3A0",
which could be used in a D-Bus address like
"unix:nonce-tcp:host=127.0.0.1,port=42,noncefile=/run/bus-for-%3A0".
*/
func DbusAddressEscapeValue(string_ string) (return__ string) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_address_escape_value(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Synchronously looks up the D-Bus address for the well-known message
bus instance specified by @bus_type. This may involve using various
platform specific mechanisms.
*/
func DbusAddressGetForBusSync(bus_type C.GBusType, cancellable IsCancellable) (return__ string, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_address_get_for_bus_sync(bus_type, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously connects to an endpoint specified by @address and
sets up the connection so it is in a state to run the client-side
of the D-Bus authentication conversation.

When the operation is finished, @callback will be invoked. You can
then call g_dbus_address_get_stream_finish() to get the result of
the operation.

This is an asynchronous failable function. See
g_dbus_address_get_stream_sync() for the synchronous version.
*/
func DbusAddressGetStream(address string, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__address := (*C.gchar)(unsafe.Pointer(C.CString(address)))
	C.g_dbus_address_get_stream(__cgo__address, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__address))
	return
}

/*
Finishes an operation started with g_dbus_address_get_stream().
*/
func DbusAddressGetStreamFinish(res *C.GAsyncResult, out_guid []string) (return__ *IOStream, __err__ error) {
	__header__out_guid := (*reflect.SliceHeader)(unsafe.Pointer(&out_guid))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GIOStream
	__cgo__return__ = C.g_dbus_address_get_stream_finish(res, (**C.gchar)(unsafe.Pointer(__header__out_guid.Data)), &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewIOStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Synchronously connects to an endpoint specified by @address and
sets up the connection so it is in a state to run the client-side
of the D-Bus authentication conversation.

This is a synchronous failable function. See
g_dbus_address_get_stream() for the asynchronous version.
*/
func DbusAddressGetStreamSync(address string, out_guid []string, cancellable IsCancellable) (return__ *IOStream, __err__ error) {
	__cgo__address := (*C.gchar)(unsafe.Pointer(C.CString(address)))
	__header__out_guid := (*reflect.SliceHeader)(unsafe.Pointer(&out_guid))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GIOStream
	__cgo__return__ = C.g_dbus_address_get_stream_sync(__cgo__address, (**C.gchar)(unsafe.Pointer(__header__out_guid.Data)), cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__address))
	if __cgo__return__ != nil {
		return__ = NewIOStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Looks up the value of an annotation.

The cost of this function is O(n) in number of annotations.
*/
func DbusAnnotationInfoLookup(annotations **C.GDBusAnnotationInfo, name string) (return__ string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_annotation_info_lookup(annotations, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Creates a D-Bus error name to use for @error. If @error matches
a registered error (cf. g_dbus_error_register_error()), the corresponding
D-Bus error name will be returned.

Otherwise the a name of the form
`org.gtk.GDBus.UnmappedGError.Quark._ESCAPED_QUARK_NAME.Code_ERROR_CODE`
will be used. This allows other GDBus applications to map the error
on the wire back to a #GError using g_dbus_error_new_for_dbus_error().

This function is typically only used in object mappings to put a
#GError on the wire. Regular applications should not use it.
*/
func DbusErrorEncodeGerror(error_ *C.GError) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_error_encode_gerror(error_)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the D-Bus error name used for @error, if any.

This function is guaranteed to return a D-Bus error name for all
#GErrors returned from functions handling remote method calls
(e.g. g_dbus_connection_call_finish()) unless
g_dbus_error_strip_remote_error() has been used on @error.
*/
func DbusErrorGetRemoteError(error_ *C.GError) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_error_get_remote_error(error_)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Checks if @error represents an error received via D-Bus from a remote peer. If so,
use g_dbus_error_get_remote_error() to get the name of the error.
*/
func DbusErrorIsRemoteError(error_ *C.GError) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_error_is_remote_error(error_)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Creates a #GError based on the contents of @dbus_error_name and
@dbus_error_message.

Errors registered with g_dbus_error_register_error() will be looked
up using @dbus_error_name and if a match is found, the error domain
and code is used. Applications can use g_dbus_error_get_remote_error()
to recover @dbus_error_name.

If a match against a registered error is not found and the D-Bus
error name is in a form as returned by g_dbus_error_encode_gerror()
the error domain and code encoded in the name is used to
create the #GError. Also, @dbus_error_name is added to the error message
such that it can be recovered with g_dbus_error_get_remote_error().

Otherwise, a #GError with the error code %G_IO_ERROR_DBUS_ERROR
in the #G_IO_ERROR error domain is returned. Also, @dbus_error_name is
added to the error message such that it can be recovered with
g_dbus_error_get_remote_error().

In all three cases, @dbus_error_name can always be recovered from the
returned #GError using the g_dbus_error_get_remote_error() function
(unless g_dbus_error_strip_remote_error() hasn't been used on the returned error).

This function is typically only used in object mappings to prepare
#GError instances for applications. Regular applications should not use
it.
*/
func DbusErrorNewForDbusError(dbus_error_name string, dbus_error_message string) (return__ *C.GError) {
	__cgo__dbus_error_name := (*C.gchar)(unsafe.Pointer(C.CString(dbus_error_name)))
	__cgo__dbus_error_message := (*C.gchar)(unsafe.Pointer(C.CString(dbus_error_message)))
	return__ = C.g_dbus_error_new_for_dbus_error(__cgo__dbus_error_name, __cgo__dbus_error_message)
	C.free(unsafe.Pointer(__cgo__dbus_error_name))
	C.free(unsafe.Pointer(__cgo__dbus_error_message))
	return
}

func DbusErrorQuark() (return__ C.GQuark) {
	return__ = C.g_dbus_error_quark()
	return
}

/*
Creates an association to map between @dbus_error_name and
#GErrors specified by @error_domain and @error_code.

This is typically done in the routine that returns the #GQuark for
an error domain.
*/
func DbusErrorRegisterError(error_domain C.GQuark, error_code int, dbus_error_name string) (return__ bool) {
	__cgo__dbus_error_name := (*C.gchar)(unsafe.Pointer(C.CString(dbus_error_name)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_error_register_error(error_domain, C.gint(error_code), __cgo__dbus_error_name)
	C.free(unsafe.Pointer(__cgo__dbus_error_name))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Helper function for associating a #GError error domain with D-Bus error names.
*/
func DbusErrorRegisterErrorDomain(error_domain_quark_name string, quark_volatile *C.gsize, entries *C.GDBusErrorEntry, num_entries uint) {
	__cgo__error_domain_quark_name := (*C.gchar)(unsafe.Pointer(C.CString(error_domain_quark_name)))
	C.g_dbus_error_register_error_domain(__cgo__error_domain_quark_name, quark_volatile, entries, C.guint(num_entries))
	C.free(unsafe.Pointer(__cgo__error_domain_quark_name))
	return
}

/*
Looks for extra information in the error message used to recover
the D-Bus error name and strips it if found. If stripped, the
message field in @error will correspond exactly to what was
received on the wire.

This is typically used when presenting errors to the end user.
*/
func DbusErrorStripRemoteError(error_ *C.GError) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_error_strip_remote_error(error_)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Destroys an association previously set up with g_dbus_error_register_error().
*/
func DbusErrorUnregisterError(error_domain C.GQuark, error_code int, dbus_error_name string) (return__ bool) {
	__cgo__dbus_error_name := (*C.gchar)(unsafe.Pointer(C.CString(dbus_error_name)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_error_unregister_error(error_domain, C.gint(error_code), __cgo__dbus_error_name)
	C.free(unsafe.Pointer(__cgo__dbus_error_name))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Generate a D-Bus GUID that can be used with
e.g. g_dbus_connection_new().

See the D-Bus specification regarding what strings are valid D-Bus
GUID (for example, D-Bus GUIDs are not RFC-4122 compliant).
*/
func DbusGenerateGuid() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_generate_guid()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts a #GValue to a #GVariant of the type indicated by the @type
parameter.

The conversion is using the following rules:

- #G_TYPE_STRING: 's', 'o', 'g' or 'ay'
- #G_TYPE_STRV: 'as', 'ao' or 'aay'
- #G_TYPE_BOOLEAN: 'b'
- #G_TYPE_UCHAR: 'y'
- #G_TYPE_INT: 'i', 'n'
- #G_TYPE_UINT: 'u', 'q'
- #G_TYPE_INT64 'x'
- #G_TYPE_UINT64: 't'
- #G_TYPE_DOUBLE: 'd'
- #G_TYPE_VARIANT: Any #GVariantType

This can fail if e.g. @gvalue is of type #G_TYPE_STRING and @type
is ['i'][G-VARIANT-TYPE-INT32:CAPS]. It will also fail for any #GType
(including e.g. #G_TYPE_OBJECT and #G_TYPE_BOXED derived-types) not
in the table above.

Note that if @gvalue is of type #G_TYPE_VARIANT and its value is
%NULL, the empty #GVariant instance (never %NULL) for @type is
returned (e.g. 0 for scalar types, the empty string for string types,
'/' for object path types, the empty array for any array type and so on).

See the g_dbus_gvariant_to_gvalue() function for how to convert a
#GVariant to a #GValue.
*/
func DbusGvalueToGvariant(gvalue *C.GValue, type_ *C.GVariantType) (return__ *C.GVariant) {
	return__ = C.g_dbus_gvalue_to_gvariant(gvalue, type_)
	return
}

/*
Converts a #GVariant to a #GValue. If @value is floating, it is consumed.

The rules specified in the g_dbus_gvalue_to_gvariant() function are
used - this function is essentially its reverse form.

The conversion never fails - a valid #GValue is always returned in
@out_gvalue.
*/
func DbusGvariantToGvalue(value *C.GVariant) (out_gvalue C.GValue) {
	C.g_dbus_gvariant_to_gvalue(value, &out_gvalue)
	return
}

/*
Checks if @string is a D-Bus address.

This doesn't check if @string is actually supported by #GDBusServer
or #GDBusConnection - use g_dbus_is_supported_address() to do more
checks.
*/
func DbusIsAddress(string_ string) (return__ bool) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_is_address(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if @string is a D-Bus GUID.

See the D-Bus specification regarding what strings are valid D-Bus
GUID (for example, D-Bus GUIDs are not RFC-4122 compliant).
*/
func DbusIsGuid(string_ string) (return__ bool) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_is_guid(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if @string is a valid D-Bus interface name.
*/
func DbusIsInterfaceName(string_ string) (return__ bool) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_is_interface_name(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if @string is a valid D-Bus member (e.g. signal or method) name.
*/
func DbusIsMemberName(string_ string) (return__ bool) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_is_member_name(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if @string is a valid D-Bus bus name (either unique or well-known).
*/
func DbusIsName(string_ string) (return__ bool) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_is_name(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Like g_dbus_is_address() but also checks if the library suppors the
transports in @string and that key/value pairs for each transport
are valid.
*/
func DbusIsSupportedAddress(string_ string) (return__ bool, __err__ error) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_is_supported_address(__cgo__string_, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Checks if @string is a valid D-Bus unique bus name.
*/
func DbusIsUniqueName(string_ string) (return__ bool) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_is_unique_name(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Creates a #GFile with the given argument from the command line.
The value of @arg can be either a URI, an absolute path or a
relative path resolved relative to the current working directory.
This operation never fails, but the returned object might not
support any I/O operation if @arg points to a malformed path.

Note that on Windows, this function expects its argument to be in
UTF-8 -- not the system code page.  This means that you
should not use this function with string from argv as it is passed
to main().  g_win32_get_command_line() will return a UTF-8 version of
the commandline.  #GApplication also uses UTF-8 but
g_application_command_line_create_file_for_arg() may be more useful
for you there.  It is also always possible to use this function with
#GOptionContext arguments of type %G_OPTION_ARG_FILENAME.
*/
func FileNewForCommandlineArg(arg string) (return__ *C.GFile) {
	__cgo__arg := C.CString(arg)
	return__ = C.g_file_new_for_commandline_arg(__cgo__arg)
	C.free(unsafe.Pointer(__cgo__arg))
	return
}

/*
Creates a #GFile with the given argument from the command line.

This function is similar to g_file_new_for_commandline_arg() except
that it allows for passing the current working directory as an
argument instead of using the current working directory of the
process.

This is useful if the commandline argument was given in a context
other than the invocation of the current process.

See also g_application_command_line_create_file_for_arg().
*/
func FileNewForCommandlineArgAndCwd(arg string, cwd string) (return__ *C.GFile) {
	__cgo__arg := (*C.gchar)(unsafe.Pointer(C.CString(arg)))
	__cgo__cwd := (*C.gchar)(unsafe.Pointer(C.CString(cwd)))
	return__ = C.g_file_new_for_commandline_arg_and_cwd(__cgo__arg, __cgo__cwd)
	C.free(unsafe.Pointer(__cgo__arg))
	C.free(unsafe.Pointer(__cgo__cwd))
	return
}

/*
Constructs a #GFile for a given path. This operation never
fails, but the returned object might not support any I/O
operation if @path is malformed.
*/
func FileNewForPath(path string) (return__ *C.GFile) {
	__cgo__path := C.CString(path)
	return__ = C.g_file_new_for_path(__cgo__path)
	C.free(unsafe.Pointer(__cgo__path))
	return
}

/*
Constructs a #GFile for a given URI. This operation never
fails, but the returned object might not support any I/O
operation if @uri is malformed or if the uri type is
not supported.
*/
func FileNewForUri(uri string) (return__ *C.GFile) {
	__cgo__uri := C.CString(uri)
	return__ = C.g_file_new_for_uri(__cgo__uri)
	C.free(unsafe.Pointer(__cgo__uri))
	return
}

/*
Opens a file in the preferred directory for temporary files (as
returned by g_get_tmp_dir()) and returns a #GFile and
#GFileIOStream pointing to it.

@tmpl should be a string in the GLib file name encoding
containing a sequence of six 'X' characters, and containing no
directory components. If it is %NULL, a default template is used.

Unlike the other #GFile constructors, this will return %NULL if
a temporary file could not be created.
*/
func FileNewTmp(tmpl string) (iostream *FileIOStream, return__ *C.GFile, __err__ error) {
	__cgo__tmpl := C.CString(tmpl)
	var __cgo__iostream *C.GFileIOStream
	var __cgo_error__ *C.GError
	return__ = C.g_file_new_tmp(__cgo__tmpl, &__cgo__iostream, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__tmpl))
	if __cgo__iostream != nil {
		iostream = NewFileIOStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__iostream).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Constructs a #GFile with the given @parse_name (i.e. something
given by g_file_get_parse_name()). This operation never fails,
but the returned object might not support any I/O operation if
the @parse_name cannot be parsed.
*/
func FileParseName(parse_name string) (return__ *C.GFile) {
	__cgo__parse_name := C.CString(parse_name)
	return__ = C.g_file_parse_name(__cgo__parse_name)
	C.free(unsafe.Pointer(__cgo__parse_name))
	return
}

/*
Deserializes a #GIcon previously serialized using g_icon_serialize().
*/
func IconDeserialize(value *C.GVariant) (return__ *C.GIcon) {
	return__ = C.g_icon_deserialize(value)
	return
}

/*
Gets a hash for an icon.
*/
func IconHash(icon unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_icon_hash((C.gconstpointer)(icon))
	return__ = uint(__cgo__return__)
	return
}

/*
Generate a #GIcon instance from @str. This function can fail if
@str is not valid - see g_icon_to_string() for discussion.

If your application or library provides one or more #GIcon
implementations you need to ensure that each #GType is registered
with the type system prior to calling g_icon_new_for_string().
*/
func IconNewForString(str string) (return__ *C.GIcon, __err__ error) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo_error__ *C.GError
	return__ = C.g_icon_new_for_string(__cgo__str, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__str))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Helper function for constructing #GInitable object. This is
similar to g_object_newv() but also initializes the object
and returns %NULL, setting an error on failure.
*/
func InitableNewv(object_type C.GType, n_parameters uint, parameters *C.GParameter, cancellable IsCancellable) (return__ *gobject.Object, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_initable_newv(object_type, C.guint(n_parameters), parameters, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = gobject.NewObjectFromCPointer(unsafe.Pointer(__cgo__return__))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Converts errno.h error codes into GIO error codes.
*/
func IoErrorFromErrno(err_no int) (return__ C.GIOErrorEnum) {
	return__ = C.g_io_error_from_errno(C.gint(err_no))
	return
}

/*
Gets the GIO Error Quark.
*/
func IoErrorQuark() (return__ C.GQuark) {
	return__ = C.g_io_error_quark()
	return
}

/*
Registers @type as extension for the extension point with name
@extension_point_name.

If @type has already been registered as an extension for this
extension point, the existing #GIOExtension object is returned.
*/
func IoExtensionPointImplement(extension_point_name string, type_ C.GType, extension_name string, priority int) (return__ *C.GIOExtension) {
	__cgo__extension_point_name := C.CString(extension_point_name)
	__cgo__extension_name := C.CString(extension_name)
	return__ = C.g_io_extension_point_implement(__cgo__extension_point_name, type_, __cgo__extension_name, C.gint(priority))
	C.free(unsafe.Pointer(__cgo__extension_point_name))
	C.free(unsafe.Pointer(__cgo__extension_name))
	return
}

/*
Looks up an existing extension point.
*/
func IoExtensionPointLookup(name string) (return__ *C.GIOExtensionPoint) {
	__cgo__name := C.CString(name)
	return__ = C.g_io_extension_point_lookup(__cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Registers an extension point.
*/
func IoExtensionPointRegister(name string) (return__ *C.GIOExtensionPoint) {
	__cgo__name := C.CString(name)
	return__ = C.g_io_extension_point_register(__cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Loads all the modules in the specified directory.

If don't require all modules to be initialized (and thus registering
all gtypes) then you can use g_io_modules_scan_all_in_directory()
which allows delayed/lazy loading of modules.
*/
func IoModulesLoadAllInDirectory(dirname string) (return__ *C.GList) {
	__cgo__dirname := (*C.gchar)(unsafe.Pointer(C.CString(dirname)))
	return__ = C.g_io_modules_load_all_in_directory(__cgo__dirname)
	C.free(unsafe.Pointer(__cgo__dirname))
	return
}

/*
Loads all the modules in the specified directory.

If don't require all modules to be initialized (and thus registering
all gtypes) then you can use g_io_modules_scan_all_in_directory()
which allows delayed/lazy loading of modules.
*/
func IoModulesLoadAllInDirectoryWithScope(dirname string, scope *C.GIOModuleScope) (return__ *C.GList) {
	__cgo__dirname := (*C.gchar)(unsafe.Pointer(C.CString(dirname)))
	return__ = C.g_io_modules_load_all_in_directory_with_scope(__cgo__dirname, scope)
	C.free(unsafe.Pointer(__cgo__dirname))
	return
}

/*
Scans all the modules in the specified directory, ensuring that
any extension point implemented by a module is registered.

This may not actually load and initialize all the types in each
module, some modules may be lazily loaded and initialized when
an extension point it implementes is used with e.g.
g_io_extension_point_get_extensions() or
g_io_extension_point_get_extension_by_name().

If you need to guarantee that all types are loaded in all the modules,
use g_io_modules_load_all_in_directory().
*/
func IoModulesScanAllInDirectory(dirname string) {
	__cgo__dirname := C.CString(dirname)
	C.g_io_modules_scan_all_in_directory(__cgo__dirname)
	C.free(unsafe.Pointer(__cgo__dirname))
	return
}

/*
Scans all the modules in the specified directory, ensuring that
any extension point implemented by a module is registered.

This may not actually load and initialize all the types in each
module, some modules may be lazily loaded and initialized when
an extension point it implementes is used with e.g.
g_io_extension_point_get_extensions() or
g_io_extension_point_get_extension_by_name().

If you need to guarantee that all types are loaded in all the modules,
use g_io_modules_load_all_in_directory().
*/
func IoModulesScanAllInDirectoryWithScope(dirname string, scope *C.GIOModuleScope) {
	__cgo__dirname := (*C.gchar)(unsafe.Pointer(C.CString(dirname)))
	C.g_io_modules_scan_all_in_directory_with_scope(__cgo__dirname, scope)
	C.free(unsafe.Pointer(__cgo__dirname))
	return
}

// g_io_scheduler_cancel_all_jobs is not generated due to deprecation attr

// g_io_scheduler_push_job is not generated due to deprecation attr

/*
Gets the default #GNetworkMonitor for the system.
*/
func NetworkMonitorGetDefault() (return__ *C.GNetworkMonitor) {
	return__ = C.g_network_monitor_get_default()
	return
}

/*
Initializes the platform networking libraries (eg, on Windows, this
calls WSAStartup()). GLib will call this itself if it is needed, so
you only need to call it if you directly call system networking
functions (without calling any GLib networking functions first).
*/
func NetworkingInit() {
	C.g_networking_init()
	return
}

/*
Utility method for #GPollableInputStream and #GPollableOutputStream
implementations. Creates a new #GSource that expects a callback of
type #GPollableSourceFunc. The new source does not actually do
anything on its own; use g_source_add_child_source() to add other
sources to it to cause it to trigger.
*/
func PollableSourceNew(pollable_stream *C.GObject) (return__ *C.GSource) {
	return__ = C.g_pollable_source_new(pollable_stream)
	return
}

/*
Utility method for #GPollableInputStream and #GPollableOutputStream
implementations. Creates a new #GSource, as with
g_pollable_source_new(), but also attaching @child_source (with a
dummy callback), and @cancellable, if they are non-%NULL.
*/
func PollableSourceNewFull(pollable_stream gobject.IsObject, child_source *C.GSource, cancellable IsCancellable) (return__ *C.GSource) {
	return__ = C.g_pollable_source_new_full(C.gpointer(unsafe.Pointer(reflect.ValueOf(pollable_stream.GetObjectPointer()).Pointer())), child_source, cancellable.GetCancellablePointer())
	return
}

/*
Tries to read from @stream, as with g_input_stream_read() (if
@blocking is %TRUE) or g_pollable_input_stream_read_nonblocking()
(if @blocking is %FALSE). This can be used to more easily share
code between blocking and non-blocking implementations of a method.

If @blocking is %FALSE, then @stream must be a
#GPollableInputStream for which g_pollable_input_stream_can_poll()
returns %TRUE, or else the behavior is undefined. If @blocking is
%TRUE, then @stream does not need to be a #GPollableInputStream.
*/
func PollableStreamRead(stream IsInputStream, buffer unsafe.Pointer, count int64, blocking bool, cancellable IsCancellable) (return__ int64, __err__ error) {
	__cgo__blocking := C.gboolean(0)
	if blocking {
		__cgo__blocking = C.gboolean(1)
	}
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_pollable_stream_read(stream.GetInputStreamPointer(), buffer, C.gsize(count), __cgo__blocking, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Tries to write to @stream, as with g_output_stream_write() (if
@blocking is %TRUE) or g_pollable_output_stream_write_nonblocking()
(if @blocking is %FALSE). This can be used to more easily share
code between blocking and non-blocking implementations of a method.

If @blocking is %FALSE, then @stream must be a
#GPollableOutputStream for which
g_pollable_output_stream_can_poll() returns %TRUE or else the
behavior is undefined. If @blocking is %TRUE, then @stream does not
need to be a #GPollableOutputStream.
*/
func PollableStreamWrite(stream IsOutputStream, buffer []byte, count int64, blocking bool, cancellable IsCancellable) (return__ int64, __err__ error) {
	__header__buffer := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	__cgo__blocking := C.gboolean(0)
	if blocking {
		__cgo__blocking = C.gboolean(1)
	}
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_pollable_stream_write(stream.GetOutputStreamPointer(), (unsafe.Pointer)(unsafe.Pointer(__header__buffer.Data)), C.gsize(count), __cgo__blocking, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Tries to write @count bytes to @stream, as with
g_output_stream_write_all(), but using g_pollable_stream_write()
rather than g_output_stream_write().

On a successful write of @count bytes, %TRUE is returned, and
@bytes_written is set to @count.

If there is an error during the operation (including
%G_IO_ERROR_WOULD_BLOCK in the non-blocking case), %FALSE is
returned and @error is set to indicate the error status,
@bytes_written is updated to contain the number of bytes written
into the stream before the error occurred.

As with g_pollable_stream_write(), if @blocking is %FALSE, then
@stream must be a #GPollableOutputStream for which
g_pollable_output_stream_can_poll() returns %TRUE or else the
behavior is undefined. If @blocking is %TRUE, then @stream does not
need to be a #GPollableOutputStream.
*/
func PollableStreamWriteAll(stream IsOutputStream, buffer []byte, count int64, blocking bool, cancellable IsCancellable) (bytes_written int64, return__ bool, __err__ error) {
	__header__buffer := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	__cgo__blocking := C.gboolean(0)
	if blocking {
		__cgo__blocking = C.gboolean(1)
	}
	var __cgo__bytes_written C.gsize
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_pollable_stream_write_all(stream.GetOutputStreamPointer(), (unsafe.Pointer)(unsafe.Pointer(__header__buffer.Data)), C.gsize(count), __cgo__blocking, &__cgo__bytes_written, cancellable.GetCancellablePointer(), &__cgo_error__)
	bytes_written = int64(__cgo__bytes_written)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Lookup "gio-proxy" extension point for a proxy implementation that supports
specified protocol.
*/
func ProxyGetDefaultForProtocol(protocol string) (return__ *C.GProxy) {
	__cgo__protocol := (*C.gchar)(unsafe.Pointer(C.CString(protocol)))
	return__ = C.g_proxy_get_default_for_protocol(__cgo__protocol)
	C.free(unsafe.Pointer(__cgo__protocol))
	return
}

/*
Gets the default #GProxyResolver for the system.
*/
func ProxyResolverGetDefault() (return__ *C.GProxyResolver) {
	return__ = C.g_proxy_resolver_get_default()
	return
}

/*
Gets the #GResolver Error Quark.
*/
func ResolverErrorQuark() (return__ C.GQuark) {
	return__ = C.g_resolver_error_quark()
	return
}

/*
Gets the #GResource Error Quark.
*/
func ResourceErrorQuark() (return__ C.GQuark) {
	return__ = C.g_resource_error_quark()
	return
}

/*
Loads a binary resource bundle and creates a #GResource representation of it, allowing
you to query it for data.

If you want to use this resource in the global resource namespace you need
to register it with g_resources_register().
*/
func ResourceLoad(filename string) (return__ *C.GResource, __err__ error) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	var __cgo_error__ *C.GError
	return__ = C.g_resource_load(__cgo__filename, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__filename))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Returns all the names of children at the specified @path in the set of
globally registered resources.
The return result is a %NULL terminated list of strings which should
be released with g_strfreev().

@lookup_flags controls the behaviour of the lookup.
*/
func ResourcesEnumerateChildren(path string, lookup_flags C.GResourceLookupFlags) (return__ **C.char, __err__ error) {
	__cgo__path := C.CString(path)
	var __cgo_error__ *C.GError
	return__ = C.g_resources_enumerate_children(__cgo__path, lookup_flags, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__path))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Looks for a file at the specified @path in the set of
globally registered resources and if found returns information about it.

@lookup_flags controls the behaviour of the lookup.
*/
func ResourcesGetInfo(path string, lookup_flags C.GResourceLookupFlags) (size int64, flags uint32, return__ bool, __err__ error) {
	__cgo__path := C.CString(path)
	var __cgo__size C.gsize
	var __cgo__flags C.guint32
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_resources_get_info(__cgo__path, lookup_flags, &__cgo__size, &__cgo__flags, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__path))
	size = int64(__cgo__size)
	flags = uint32(__cgo__flags)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Looks for a file at the specified @path in the set of
globally registered resources and returns a #GBytes that
lets you directly access the data in memory.

The data is always followed by a zero byte, so you
can safely use the data as a C string. However, that byte
is not included in the size of the GBytes.

For uncompressed resource files this is a pointer directly into
the resource bundle, which is typically in some readonly data section
in the program binary. For compressed files we allocate memory on
the heap and automatically uncompress the data.

@lookup_flags controls the behaviour of the lookup.
*/
func ResourcesLookupData(path string, lookup_flags C.GResourceLookupFlags) (return__ *C.GBytes, __err__ error) {
	__cgo__path := C.CString(path)
	var __cgo_error__ *C.GError
	return__ = C.g_resources_lookup_data(__cgo__path, lookup_flags, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__path))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Looks for a file at the specified @path in the set of
globally registered resources and returns a #GInputStream
that lets you read the data.

@lookup_flags controls the behaviour of the lookup.
*/
func ResourcesOpenStream(path string, lookup_flags C.GResourceLookupFlags) (return__ *InputStream, __err__ error) {
	__cgo__path := C.CString(path)
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GInputStream
	__cgo__return__ = C.g_resources_open_stream(__cgo__path, lookup_flags, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__path))
	if __cgo__return__ != nil {
		return__ = NewInputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Registers the resource with the process-global set of resources.
Once a resource is registered the files in it can be accessed
with the global resource lookup functions like g_resources_lookup_data().
*/
func ResourcesRegister(resource *C.GResource) {
	C.g_resources_register(resource)
	return
}

/*
Unregisters the resource from the process-global set of resources.
*/
func ResourcesUnregister(resource *C.GResource) {
	C.g_resources_unregister(resource)
	return
}

/*
Gets the default system schema source.

This function is not required for normal uses of #GSettings but it
may be useful to authors of plugin management systems or to those who
want to introspect the content of schemas.

If no schemas are installed, %NULL will be returned.

The returned source may actually consist of multiple schema sources
from different directories, depending on which directories were given
in `XDG_DATA_DIRS` and `GSETTINGS_SCHEMA_DIR`. For this reason, all
lookups performed against the default source should probably be done
recursively.
*/
func SettingsSchemaSourceGetDefault() (return__ *C.GSettingsSchemaSource) {
	return__ = C.g_settings_schema_source_get_default()
	return
}

// g_simple_async_report_error_in_idle is not generated due to varargs

/*
Reports an error in an idle function. Similar to
g_simple_async_report_error_in_idle(), but takes a #GError rather
than building a new one.
*/
func SimpleAsyncReportGerrorInIdle(object *C.GObject, callback C.GAsyncReadyCallback, user_data unsafe.Pointer, error_ *C.GError) {
	C.g_simple_async_report_gerror_in_idle(object, callback, (C.gpointer)(user_data), error_)
	return
}

/*
Reports an error in an idle function. Similar to
g_simple_async_report_gerror_in_idle(), but takes over the caller's
ownership of @error, so the caller does not have to free it any more.
*/
func SimpleAsyncReportTakeGerrorInIdle(object *C.GObject, callback C.GAsyncReadyCallback, user_data unsafe.Pointer, error_ *C.GError) {
	C.g_simple_async_report_take_gerror_in_idle(object, callback, (C.gpointer)(user_data), error_)
	return
}

/*
Sorts @targets in place according to the algorithm in RFC 2782.
*/
func SrvTargetListSort(targets *C.GList) (return__ *C.GList) {
	return__ = C.g_srv_target_list_sort(targets)
	return
}

/*
Gets the default #GTlsBackend for the system.
*/
func TlsBackendGetDefault() (return__ *C.GTlsBackend) {
	return__ = C.g_tls_backend_get_default()
	return
}

/*
Creates a new #GTlsClientConnection wrapping @base_io_stream (which
must have pollable input and output streams) which is assumed to
communicate with the server identified by @server_identity.
*/
func TlsClientConnectionNew(base_io_stream IsIOStream, server_identity *C.GSocketConnectable) (return__ *C.GIOStream, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_tls_client_connection_new(base_io_stream.GetIOStreamPointer(), server_identity, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the TLS error quark.
*/
func TlsErrorQuark() (return__ C.GQuark) {
	return__ = C.g_tls_error_quark()
	return
}

/*
Creates a new #GTlsFileDatabase which uses anchor certificate authorities
in @anchors to verify certificate chains.

The certificates in @anchors must be PEM encoded.
*/
func TlsFileDatabaseNew(anchors string) (return__ *C.GTlsDatabase, __err__ error) {
	__cgo__anchors := (*C.gchar)(unsafe.Pointer(C.CString(anchors)))
	var __cgo_error__ *C.GError
	return__ = C.g_tls_file_database_new(__cgo__anchors, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__anchors))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new #GTlsServerConnection wrapping @base_io_stream (which
must have pollable input and output streams).
*/
func TlsServerConnectionNew(base_io_stream IsIOStream, certificate IsTlsCertificate) (return__ *C.GIOStream, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_tls_server_connection_new(base_io_stream.GetIOStreamPointer(), certificate.GetTlsCertificatePointer(), &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Determines if @mount_path is considered an implementation of the
OS. This is primarily used for hiding mountable and mounted volumes
that only are used in the OS and has little to no relevance to the
casual user.
*/
func UnixIsMountPathSystemInternal(mount_path string) (return__ bool) {
	__cgo__mount_path := C.CString(mount_path)
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unix_is_mount_path_system_internal(__cgo__mount_path)
	C.free(unsafe.Pointer(__cgo__mount_path))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets a #GUnixMountEntry for a given mount path. If @time_read
is set, it will be filled with a unix timestamp for checking
if the mounts have changed since with g_unix_mounts_changed_since().
*/
func UnixMountAt(mount_path string) (time_read uint64, return__ *C.GUnixMountEntry) {
	__cgo__mount_path := C.CString(mount_path)
	var __cgo__time_read C.guint64
	return__ = C.g_unix_mount_at(__cgo__mount_path, &__cgo__time_read)
	C.free(unsafe.Pointer(__cgo__mount_path))
	time_read = uint64(__cgo__time_read)
	return
}

/*
Compares two unix mounts.
*/
func UnixMountCompare(mount1 *C.GUnixMountEntry, mount2 *C.GUnixMountEntry) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_unix_mount_compare(mount1, mount2)
	return__ = int(__cgo__return__)
	return
}

/*
Frees a unix mount.
*/
func UnixMountFree(mount_entry *C.GUnixMountEntry) {
	C.g_unix_mount_free(mount_entry)
	return
}

/*
Gets the device path for a unix mount.
*/
func UnixMountGetDevicePath(mount_entry *C.GUnixMountEntry) (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_unix_mount_get_device_path(mount_entry)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Gets the filesystem type for the unix mount.
*/
func UnixMountGetFsType(mount_entry *C.GUnixMountEntry) (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_unix_mount_get_fs_type(mount_entry)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Gets the mount path for a unix mount.
*/
func UnixMountGetMountPath(mount_entry *C.GUnixMountEntry) (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_unix_mount_get_mount_path(mount_entry)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Guesses whether a Unix mount can be ejected.
*/
func UnixMountGuessCanEject(mount_entry *C.GUnixMountEntry) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unix_mount_guess_can_eject(mount_entry)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Guesses the icon of a Unix mount.
*/
func UnixMountGuessIcon(mount_entry *C.GUnixMountEntry) (return__ *C.GIcon) {
	return__ = C.g_unix_mount_guess_icon(mount_entry)
	return
}

/*
Guesses the name of a Unix mount.
The result is a translated string.
*/
func UnixMountGuessName(mount_entry *C.GUnixMountEntry) (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_unix_mount_guess_name(mount_entry)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Guesses whether a Unix mount should be displayed in the UI.
*/
func UnixMountGuessShouldDisplay(mount_entry *C.GUnixMountEntry) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unix_mount_guess_should_display(mount_entry)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Guesses the symbolic icon of a Unix mount.
*/
func UnixMountGuessSymbolicIcon(mount_entry *C.GUnixMountEntry) (return__ *C.GIcon) {
	return__ = C.g_unix_mount_guess_symbolic_icon(mount_entry)
	return
}

/*
Checks if a unix mount is mounted read only.
*/
func UnixMountIsReadonly(mount_entry *C.GUnixMountEntry) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unix_mount_is_readonly(mount_entry)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if a unix mount is a system path.
*/
func UnixMountIsSystemInternal(mount_entry *C.GUnixMountEntry) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unix_mount_is_system_internal(mount_entry)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if the unix mount points have changed since a given unix time.
*/
func UnixMountPointsChangedSince(time uint64) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unix_mount_points_changed_since(C.guint64(time))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets a #GList of #GUnixMountPoint containing the unix mount points.
If @time_read is set, it will be filled with the mount timestamp,
allowing for checking if the mounts have changed with
g_unix_mount_points_changed_since().
*/
func UnixMountPointsGet() (time_read uint64, return__ *C.GList) {
	var __cgo__time_read C.guint64
	return__ = C.g_unix_mount_points_get(&__cgo__time_read)
	time_read = uint64(__cgo__time_read)
	return
}

/*
Checks if the unix mounts have changed since a given unix time.
*/
func UnixMountsChangedSince(time uint64) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unix_mounts_changed_since(C.guint64(time))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets a #GList of #GUnixMountEntry containing the unix mounts.
If @time_read is set, it will be filled with the mount
timestamp, allowing for checking if the mounts have changed
with g_unix_mounts_changed_since().
*/
func UnixMountsGet() (time_read uint64, return__ *C.GList) {
	var __cgo__time_read C.guint64
	return__ = C.g_unix_mounts_get(&__cgo__time_read)
	time_read = uint64(__cgo__time_read)
	return
}
