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
*/
import "C"
import "unsafe"
import "reflect"
import "errors"
import "github.com/reusee/ggir/gobject"
import "github.com/reusee/ggir/glib"

func init() {
	_ = unsafe.Pointer(nil)
	_ = reflect.ValueOf(nil)
	_ = errors.New("")
}

var UnusedFix_ bool
var _ = gobject.UnusedFix_
var _ = glib.UnusedFix_

/*
Creates a new application launch context. This is not normally used,
instead you instantiate a subclass of this, such as #GdkAppLaunchContext.
*/
func AppLaunchContextNew() (return__ *AppLaunchContext) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_app_launch_context_new()
	if __cgo__return__ != nil {
		return__ = NewAppLaunchContextFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GApplication instance.

If non-%NULL, the application id must be valid.  See
g_application_id_is_valid().

If no application ID is given then some features of #GApplication
(most notably application uniqueness) will be disabled.
*/
func ApplicationNew(application_id string, flags C.GApplicationFlags) (return__ *Application) {
	__cgo__application_id := (*C.gchar)(unsafe.Pointer(C.CString(application_id)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_application_new(__cgo__application_id, flags)
	C.free(unsafe.Pointer(__cgo__application_id))
	if __cgo__return__ != nil {
		return__ = NewApplicationFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GInputStream from the given @base_stream, with
a buffer set to the default size (4 kilobytes).
*/
func BufferedInputStreamNew(base_stream IsInputStream) (return__ *BufferedInputStream) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_buffered_input_stream_new(base_stream.GetInputStreamPointer())
	if __cgo__return__ != nil {
		return__ = NewBufferedInputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GBufferedInputStream from the given @base_stream,
with a buffer set to @size.
*/
func BufferedInputStreamNewSized(base_stream IsInputStream, size int64) (return__ *BufferedInputStream) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_buffered_input_stream_new_sized(base_stream.GetInputStreamPointer(), C.gsize(size))
	if __cgo__return__ != nil {
		return__ = NewBufferedInputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new buffered output stream for a base stream.
*/
func BufferedOutputStreamNew(base_stream IsOutputStream) (return__ *BufferedOutputStream) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_buffered_output_stream_new(base_stream.GetOutputStreamPointer())
	if __cgo__return__ != nil {
		return__ = NewBufferedOutputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new buffered output stream with a given buffer size.
*/
func BufferedOutputStreamNewSized(base_stream IsOutputStream, size int64) (return__ *BufferedOutputStream) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_buffered_output_stream_new_sized(base_stream.GetOutputStreamPointer(), C.gsize(size))
	if __cgo__return__ != nil {
		return__ = NewBufferedOutputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new icon for a bytes.
*/
func BytesIconNew(bytes *C.GBytes) (return__ *BytesIcon) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_bytes_icon_new(bytes)
	if __cgo__return__ != nil {
		return__ = NewBytesIconFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GCancellable object.

Applications that want to start one or more operations
that should be cancellable should create a #GCancellable
and pass it to the operations.

One #GCancellable can be used in multiple consecutive
operations or in multiple concurrent operations.
*/
func CancellableNew() (return__ *Cancellable) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_cancellable_new()
	if __cgo__return__ != nil {
		return__ = NewCancellableFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GCharsetConverter.
*/
func CharsetConverterNew(to_charset string, from_charset string) (return__ *CharsetConverter, __err__ error) {
	__cgo__to_charset := (*C.gchar)(unsafe.Pointer(C.CString(to_charset)))
	__cgo__from_charset := (*C.gchar)(unsafe.Pointer(C.CString(from_charset)))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_charset_converter_new(__cgo__to_charset, __cgo__from_charset, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__to_charset))
	C.free(unsafe.Pointer(__cgo__from_charset))
	if __cgo__return__ != nil {
		return__ = NewCharsetConverterFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new converter input stream for the @base_stream.
*/
func ConverterInputStreamNew(base_stream IsInputStream, converter *C.GConverter) (return__ *ConverterInputStream) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_converter_input_stream_new(base_stream.GetInputStreamPointer(), converter)
	if __cgo__return__ != nil {
		return__ = NewConverterInputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new converter output stream for the @base_stream.
*/
func ConverterOutputStreamNew(base_stream IsOutputStream, converter *C.GConverter) (return__ *ConverterOutputStream) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_converter_output_stream_new(base_stream.GetOutputStreamPointer(), converter)
	if __cgo__return__ != nil {
		return__ = NewConverterOutputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GCredentials object with credentials matching the
the current process.
*/
func CredentialsNew() (return__ *Credentials) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_credentials_new()
	if __cgo__return__ != nil {
		return__ = NewCredentialsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GDBusAuthObserver object.
*/
func DBusAuthObserverNew() (return__ *DBusAuthObserver) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_auth_observer_new()
	if __cgo__return__ != nil {
		return__ = NewDBusAuthObserverFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Finishes an operation started with g_dbus_connection_new().
*/
func DBusConnectionNewFinish(res *C.GAsyncResult) (return__ *DBusConnection, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_connection_new_finish(res, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewDBusConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Finishes an operation started with g_dbus_connection_new_for_address().
*/
func DBusConnectionNewForAddressFinish(res *C.GAsyncResult) (return__ *DBusConnection, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_connection_new_for_address_finish(res, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewDBusConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Synchronously connects and sets up a D-Bus client connection for
exchanging D-Bus messages with an endpoint specified by @address
which must be in the D-Bus address format.

This constructor can only be used to initiate client-side
connections - use g_dbus_connection_new_sync() if you need to act
as the server. In particular, @flags cannot contain the
%G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER or
%G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS flags.

This is a synchronous failable constructor. See
g_dbus_connection_new_for_address() for the asynchronous version.

If @observer is not %NULL it may be used to control the
authentication process.
*/
func DBusConnectionNewForAddressSync(address string, flags C.GDBusConnectionFlags, observer IsDBusAuthObserver, cancellable IsCancellable) (return__ *DBusConnection, __err__ error) {
	__cgo__address := (*C.gchar)(unsafe.Pointer(C.CString(address)))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_connection_new_for_address_sync(__cgo__address, flags, observer.GetDBusAuthObserverPointer(), cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__address))
	if __cgo__return__ != nil {
		return__ = NewDBusConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Synchronously sets up a D-Bus connection for exchanging D-Bus messages
with the end represented by @stream.

If @stream is a #GSocketConnection, then the corresponding #GSocket
will be put into non-blocking mode.

The D-Bus connection will interact with @stream from a worker thread.
As a result, the caller should not interact with @stream after this
method has been called, except by calling g_object_unref() on it.

If @observer is not %NULL it may be used to control the
authentication process.

This is a synchronous failable constructor. See
g_dbus_connection_new() for the asynchronous version.
*/
func DBusConnectionNewSync(stream IsIOStream, guid string, flags C.GDBusConnectionFlags, observer IsDBusAuthObserver, cancellable IsCancellable) (return__ *DBusConnection, __err__ error) {
	__cgo__guid := (*C.gchar)(unsafe.Pointer(C.CString(guid)))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_connection_new_sync(stream.GetIOStreamPointer(), __cgo__guid, flags, observer.GetDBusAuthObserverPointer(), cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__guid))
	if __cgo__return__ != nil {
		return__ = NewDBusConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new empty #GDBusMessage.
*/
func DBusMessageNew() (return__ *DBusMessage) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_message_new()
	if __cgo__return__ != nil {
		return__ = NewDBusMessageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GDBusMessage from the data stored at @blob. The byte
order that the message was in can be retrieved using
g_dbus_message_get_byte_order().
*/
func DBusMessageNewFromBlob(blob []byte, blob_len int64, capabilities C.GDBusCapabilityFlags) (return__ *DBusMessage, __err__ error) {
	__header__blob := (*reflect.SliceHeader)(unsafe.Pointer(&blob))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_message_new_from_blob((*C.guchar)(unsafe.Pointer(__header__blob.Data)), C.gsize(blob_len), capabilities, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewDBusMessageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new #GDBusMessage for a method call.
*/
func DBusMessageNewMethodCall(name string, path string, interface_ string, method string) (return__ *DBusMessage) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	__cgo__interface_ := (*C.gchar)(unsafe.Pointer(C.CString(interface_)))
	__cgo__method := (*C.gchar)(unsafe.Pointer(C.CString(method)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_message_new_method_call(__cgo__name, __cgo__path, __cgo__interface_, __cgo__method)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__path))
	C.free(unsafe.Pointer(__cgo__interface_))
	C.free(unsafe.Pointer(__cgo__method))
	if __cgo__return__ != nil {
		return__ = NewDBusMessageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GDBusMessage for a signal emission.
*/
func DBusMessageNewSignal(path string, interface_ string, signal string) (return__ *DBusMessage) {
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	__cgo__interface_ := (*C.gchar)(unsafe.Pointer(C.CString(interface_)))
	__cgo__signal := (*C.gchar)(unsafe.Pointer(C.CString(signal)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_message_new_signal(__cgo__path, __cgo__interface_, __cgo__signal)
	C.free(unsafe.Pointer(__cgo__path))
	C.free(unsafe.Pointer(__cgo__interface_))
	C.free(unsafe.Pointer(__cgo__signal))
	if __cgo__return__ != nil {
		return__ = NewDBusMessageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Finishes an operation started with g_dbus_object_manager_client_new().
*/
func DBusObjectManagerClientNewFinish(res *C.GAsyncResult) (return__ *DBusObjectManagerClient, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_object_manager_client_new_finish(res, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewDBusObjectManagerClientFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Finishes an operation started with g_dbus_object_manager_client_new_for_bus().
*/
func DBusObjectManagerClientNewForBusFinish(res *C.GAsyncResult) (return__ *DBusObjectManagerClient, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_object_manager_client_new_for_bus_finish(res, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewDBusObjectManagerClientFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Like g_dbus_object_manager_client_new_sync() but takes a #GBusType instead
of a #GDBusConnection.

This is a synchronous failable constructor - the calling thread is
blocked until a reply is received. See g_dbus_object_manager_client_new_for_bus()
for the asynchronous version.
*/
func DBusObjectManagerClientNewForBusSync(bus_type C.GBusType, flags C.GDBusObjectManagerClientFlags, name string, object_path string, get_proxy_type_func C.GDBusProxyTypeFunc, get_proxy_type_user_data unsafe.Pointer, get_proxy_type_destroy_notify C.GDestroyNotify, cancellable IsCancellable) (return__ *DBusObjectManagerClient, __err__ error) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_object_manager_client_new_for_bus_sync(bus_type, flags, __cgo__name, __cgo__object_path, get_proxy_type_func, (C.gpointer)(get_proxy_type_user_data), get_proxy_type_destroy_notify, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__object_path))
	if __cgo__return__ != nil {
		return__ = NewDBusObjectManagerClientFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new #GDBusObjectManagerClient object.

This is a synchronous failable constructor - the calling thread is
blocked until a reply is received. See g_dbus_object_manager_client_new()
for the asynchronous version.
*/
func DBusObjectManagerClientNewSync(connection IsDBusConnection, flags C.GDBusObjectManagerClientFlags, name string, object_path string, get_proxy_type_func C.GDBusProxyTypeFunc, get_proxy_type_user_data unsafe.Pointer, get_proxy_type_destroy_notify C.GDestroyNotify, cancellable IsCancellable) (return__ *DBusObjectManagerClient, __err__ error) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_object_manager_client_new_sync(connection.GetDBusConnectionPointer(), flags, __cgo__name, __cgo__object_path, get_proxy_type_func, (C.gpointer)(get_proxy_type_user_data), get_proxy_type_destroy_notify, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__object_path))
	if __cgo__return__ != nil {
		return__ = NewDBusObjectManagerClientFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new #GDBusObjectManagerServer object.

The returned server isn't yet exported on any connection. To do so,
use g_dbus_object_manager_server_set_connection(). Normally you
want to export all of your objects before doing so to avoid <ulink
url="http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-objectmanager">InterfacesAdded</ulink>
signals being emitted.
*/
func DBusObjectManagerServerNew(object_path string) (return__ *DBusObjectManagerServer) {
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_object_manager_server_new(__cgo__object_path)
	C.free(unsafe.Pointer(__cgo__object_path))
	if __cgo__return__ != nil {
		return__ = NewDBusObjectManagerServerFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GDBusObjectProxy for the given connection and
object path.
*/
func DBusObjectProxyNew(connection IsDBusConnection, object_path string) (return__ *DBusObjectProxy) {
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_object_proxy_new(connection.GetDBusConnectionPointer(), __cgo__object_path)
	C.free(unsafe.Pointer(__cgo__object_path))
	if __cgo__return__ != nil {
		return__ = NewDBusObjectProxyFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GDBusObjectSkeleton.
*/
func DBusObjectSkeletonNew(object_path string) (return__ *DBusObjectSkeleton) {
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_object_skeleton_new(__cgo__object_path)
	C.free(unsafe.Pointer(__cgo__object_path))
	if __cgo__return__ != nil {
		return__ = NewDBusObjectSkeletonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Finishes creating a #GDBusProxy.
*/
func DBusProxyNewFinish(res *C.GAsyncResult) (return__ *DBusProxy, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_proxy_new_finish(res, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewDBusProxyFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Finishes creating a #GDBusProxy.
*/
func DBusProxyNewForBusFinish(res *C.GAsyncResult) (return__ *DBusProxy, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_proxy_new_for_bus_finish(res, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewDBusProxyFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Like g_dbus_proxy_new_sync() but takes a #GBusType instead of a #GDBusConnection.

#GDBusProxy is used in this [example][gdbus-wellknown-proxy].
*/
func DBusProxyNewForBusSync(bus_type C.GBusType, flags C.GDBusProxyFlags, info *C.GDBusInterfaceInfo, name string, object_path string, interface_name string, cancellable IsCancellable) (return__ *DBusProxy, __err__ error) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	__cgo__interface_name := (*C.gchar)(unsafe.Pointer(C.CString(interface_name)))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_proxy_new_for_bus_sync(bus_type, flags, info, __cgo__name, __cgo__object_path, __cgo__interface_name, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__object_path))
	C.free(unsafe.Pointer(__cgo__interface_name))
	if __cgo__return__ != nil {
		return__ = NewDBusProxyFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a proxy for accessing @interface_name on the remote object
at @object_path owned by @name at @connection and synchronously
loads D-Bus properties unless the
%G_DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES flag is used.

If the %G_DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS flag is not set, also sets up
match rules for signals. Connect to the #GDBusProxy::g-signal signal
to handle signals from the remote object.

If @name is a well-known name and the
%G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START and %G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START_AT_CONSTRUCTION
flags aren't set and no name owner currently exists, the message bus
will be requested to launch a name owner for the name.

This is a synchronous failable constructor. See g_dbus_proxy_new()
and g_dbus_proxy_new_finish() for the asynchronous version.

#GDBusProxy is used in this [example][gdbus-wellknown-proxy].
*/
func DBusProxyNewSync(connection IsDBusConnection, flags C.GDBusProxyFlags, info *C.GDBusInterfaceInfo, name string, object_path string, interface_name string, cancellable IsCancellable) (return__ *DBusProxy, __err__ error) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	__cgo__interface_name := (*C.gchar)(unsafe.Pointer(C.CString(interface_name)))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_proxy_new_sync(connection.GetDBusConnectionPointer(), flags, info, __cgo__name, __cgo__object_path, __cgo__interface_name, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__object_path))
	C.free(unsafe.Pointer(__cgo__interface_name))
	if __cgo__return__ != nil {
		return__ = NewDBusProxyFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new D-Bus server that listens on the first address in
@address that works.

Once constructed, you can use g_dbus_server_get_client_address() to
get a D-Bus address string that clients can use to connect.

Connect to the #GDBusServer::new-connection signal to handle
incoming connections.

The returned #GDBusServer isn't active - you have to start it with
g_dbus_server_start().

#GDBusServer is used in this [example][gdbus-peer-to-peer].

This is a synchronous failable constructor. See
g_dbus_server_new() for the asynchronous version.
*/
func DBusServerNewSync(address string, flags C.GDBusServerFlags, guid string, observer IsDBusAuthObserver, cancellable IsCancellable) (return__ *DBusServer, __err__ error) {
	__cgo__address := (*C.gchar)(unsafe.Pointer(C.CString(address)))
	__cgo__guid := (*C.gchar)(unsafe.Pointer(C.CString(guid)))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_dbus_server_new_sync(__cgo__address, flags, __cgo__guid, observer.GetDBusAuthObserverPointer(), cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__address))
	C.free(unsafe.Pointer(__cgo__guid))
	if __cgo__return__ != nil {
		return__ = NewDBusServerFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new data input stream for the @base_stream.
*/
func DataInputStreamNew(base_stream IsInputStream) (return__ *DataInputStream) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_data_input_stream_new(base_stream.GetInputStreamPointer())
	if __cgo__return__ != nil {
		return__ = NewDataInputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new data output stream for @base_stream.
*/
func DataOutputStreamNew(base_stream IsOutputStream) (return__ *DataOutputStream) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_data_output_stream_new(base_stream.GetOutputStreamPointer())
	if __cgo__return__ != nil {
		return__ = NewDataOutputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GDesktopAppInfo based on a desktop file id.

A desktop file id is the basename of the desktop file, including the
.desktop extension. GIO is looking for a desktop file with this name
in the `applications` subdirectories of the XDG
data directories (i.e. the directories specified in the `XDG_DATA_HOME`
and `XDG_DATA_DIRS` environment variables). GIO also supports the
prefix-to-subdirectory mapping that is described in the
[Menu Spec](http://standards.freedesktop.org/menu-spec/latest/)
(i.e. a desktop id of kde-foo.desktop will match
`/usr/share/applications/kde/foo.desktop`).
*/
func DesktopAppInfoNew(desktop_id string) (return__ *DesktopAppInfo) {
	__cgo__desktop_id := C.CString(desktop_id)
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_desktop_app_info_new(__cgo__desktop_id)
	C.free(unsafe.Pointer(__cgo__desktop_id))
	if __cgo__return__ != nil {
		return__ = NewDesktopAppInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GDesktopAppInfo.
*/
func DesktopAppInfoNewFromFilename(filename string) (return__ *DesktopAppInfo) {
	__cgo__filename := C.CString(filename)
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_desktop_app_info_new_from_filename(__cgo__filename)
	C.free(unsafe.Pointer(__cgo__filename))
	if __cgo__return__ != nil {
		return__ = NewDesktopAppInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GDesktopAppInfo.
*/
func DesktopAppInfoNewFromKeyfile(key_file *C.GKeyFile) (return__ *DesktopAppInfo) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_desktop_app_info_new_from_keyfile(key_file)
	if __cgo__return__ != nil {
		return__ = NewDesktopAppInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new emblem for @icon.
*/
func EmblemNew(icon *C.GIcon) (return__ *Emblem) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_emblem_new(icon)
	if __cgo__return__ != nil {
		return__ = NewEmblemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new emblem for @icon.
*/
func EmblemNewWithOrigin(icon *C.GIcon, origin C.GEmblemOrigin) (return__ *Emblem) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_emblem_new_with_origin(icon, origin)
	if __cgo__return__ != nil {
		return__ = NewEmblemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new emblemed icon for @icon with the emblem @emblem.
*/
func EmblemedIconNew(icon *C.GIcon, emblem IsEmblem) (return__ *EmblemedIcon) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_emblemed_icon_new(icon, emblem.GetEmblemPointer())
	if __cgo__return__ != nil {
		return__ = NewEmblemedIconFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new icon for a file.
*/
func FileIconNew(file *C.GFile) (return__ *FileIcon) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_file_icon_new(file)
	if __cgo__return__ != nil {
		return__ = NewFileIconFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new file info structure.
*/
func FileInfoNew() (return__ *FileInfo) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_file_info_new()
	if __cgo__return__ != nil {
		return__ = NewFileInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new filename completer.
*/
func FilenameCompleterNew() (return__ *FilenameCompleter) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_filename_completer_new()
	if __cgo__return__ != nil {
		return__ = NewFilenameCompleterFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new GIOModule that will load the specific
shared library when in use.
*/
func IOModuleNew(filename string) (return__ *IOModule) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_io_module_new(__cgo__filename)
	C.free(unsafe.Pointer(__cgo__filename))
	if __cgo__return__ != nil {
		return__ = NewIOModuleFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GInetAddress for the "any" address (unassigned/"don't
care") for @family.
*/
func InetAddressNewAny(family C.GSocketFamily) (return__ *InetAddress) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_inet_address_new_any(family)
	if __cgo__return__ != nil {
		return__ = NewInetAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GInetAddress from the given @family and @bytes.
@bytes should be 4 bytes for %G_SOCKET_FAMILY_IPV4 and 16 bytes for
%G_SOCKET_FAMILY_IPV6.
*/
func InetAddressNewFromBytes(bytes []uint8, family C.GSocketFamily) (return__ *InetAddress) {
	__header__bytes := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_inet_address_new_from_bytes((*C.guint8)(unsafe.Pointer(__header__bytes.Data)), family)
	if __cgo__return__ != nil {
		return__ = NewInetAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Parses @string as an IP address and creates a new #GInetAddress.
*/
func InetAddressNewFromString(string_ string) (return__ *InetAddress) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_inet_address_new_from_string(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	if __cgo__return__ != nil {
		return__ = NewInetAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GInetAddress for the loopback address for @family.
*/
func InetAddressNewLoopback(family C.GSocketFamily) (return__ *InetAddress) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_inet_address_new_loopback(family)
	if __cgo__return__ != nil {
		return__ = NewInetAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GInetAddressMask representing all addresses whose
first @length bits match @addr.
*/
func InetAddressMaskNew(addr IsInetAddress, length uint) (return__ *InetAddressMask, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_inet_address_mask_new(addr.GetInetAddressPointer(), C.guint(length), &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewInetAddressMaskFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Parses @mask_string as an IP address and (optional) length, and
creates a new #GInetAddressMask. The length, if present, is
delimited by a "/". If it is not present, then the length is
assumed to be the full length of the address.
*/
func InetAddressMaskNewFromString(mask_string string) (return__ *InetAddressMask, __err__ error) {
	__cgo__mask_string := (*C.gchar)(unsafe.Pointer(C.CString(mask_string)))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_inet_address_mask_new_from_string(__cgo__mask_string, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__mask_string))
	if __cgo__return__ != nil {
		return__ = NewInetAddressMaskFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new #GInetSocketAddress for @address and @port.
*/
func InetSocketAddressNew(address IsInetAddress, port uint16) (return__ *InetSocketAddress) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_inet_socket_address_new(address.GetInetAddressPointer(), C.guint16(port))
	if __cgo__return__ != nil {
		return__ = NewInetSocketAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GInetSocketAddress for @address and @port.

If @address is an IPv6 address, it can also contain a scope ID
(separated from the address by a "<literal>%</literal>").
*/
func InetSocketAddressNewFromString(address string, port uint) (return__ *InetSocketAddress) {
	__cgo__address := C.CString(address)
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_inet_socket_address_new_from_string(__cgo__address, C.guint(port))
	C.free(unsafe.Pointer(__cgo__address))
	if __cgo__return__ != nil {
		return__ = NewInetSocketAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new empty #GMemoryInputStream.
*/
func MemoryInputStreamNew() (return__ *MemoryInputStream) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_memory_input_stream_new()
	if __cgo__return__ != nil {
		return__ = NewMemoryInputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GMemoryInputStream with data from the given @bytes.
*/
func MemoryInputStreamNewFromBytes(bytes *C.GBytes) (return__ *MemoryInputStream) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_memory_input_stream_new_from_bytes(bytes)
	if __cgo__return__ != nil {
		return__ = NewMemoryInputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GMemoryInputStream with data in memory of a given size.
*/
func MemoryInputStreamNewFromData(data []byte, len_ int64, destroy C.GDestroyNotify) (return__ *MemoryInputStream) {
	__header__data := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_memory_input_stream_new_from_data((unsafe.Pointer)(unsafe.Pointer(__header__data.Data)), C.gssize(len_), destroy)
	if __cgo__return__ != nil {
		return__ = NewMemoryInputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GMemoryOutputStream.

In most cases this is not the function you want.  See
g_memory_output_stream_new_resizable() instead.

If @data is non-%NULL, the stream will use that for its internal storage.

If @realloc_fn is non-%NULL, it will be used for resizing the internal
storage when necessary and the stream will be considered resizable.
In that case, the stream will start out being (conceptually) empty.
@size is used only as a hint for how big @data is.  Specifically,
seeking to the end of a newly-created stream will seek to zero, not
@size.  Seeking past the end of the stream and then writing will
introduce a zero-filled gap.

If @realloc_fn is %NULL then the stream is fixed-sized.  Seeking to
the end will seek to @size exactly.  Writing past the end will give
an 'out of space' error.  Attempting to seek past the end will fail.
Unlike the resizable case, seeking to an offset within the stream and
writing will preserve the bytes passed in as @data before that point
and will return them as part of g_memory_output_stream_steal_data().
If you intend to seek you should probably therefore ensure that @data
is properly initialised.

It is probably only meaningful to provide @data and @size in the case
that you want a fixed-sized stream.  Put another way: if @realloc_fn
is non-%NULL then it makes most sense to give @data as %NULL and
@size as 0 (allowing #GMemoryOutputStream to do the initial
allocation for itself).

|[<!-- language="C" -->
// a stream that can grow
stream = g_memory_output_stream_new (NULL, 0, realloc, free);

// another stream that can grow
stream2 = g_memory_output_stream_new (NULL, 0, g_realloc, g_free);

// a fixed-size stream
data = malloc (200);
stream3 = g_memory_output_stream_new (data, 200, NULL, free);
]|
*/
func MemoryOutputStreamNew(data unsafe.Pointer, size int64, realloc_function C.GReallocFunc, destroy_function C.GDestroyNotify) (return__ *MemoryOutputStream) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_memory_output_stream_new((C.gpointer)(data), C.gsize(size), realloc_function, destroy_function)
	if __cgo__return__ != nil {
		return__ = NewMemoryOutputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GMemoryOutputStream, using g_realloc() and g_free()
for memory allocation.
*/
func MemoryOutputStreamNewResizable() (return__ *MemoryOutputStream) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_memory_output_stream_new_resizable()
	if __cgo__return__ != nil {
		return__ = NewMemoryOutputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GMenu.

The new menu has no items.
*/
func MenuNew() (return__ *Menu) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_menu_new()
	if __cgo__return__ != nil {
		return__ = NewMenuFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GMenuItem.

If @label is non-%NULL it is used to set the "label" attribute of the
new item.

If @detailed_action is non-%NULL it is used to set the "action" and
possibly the "target" attribute of the new item.  See
g_menu_item_set_detailed_action() for more information.
*/
func MenuItemNew(label string, detailed_action string) (return__ *MenuItem) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	__cgo__detailed_action := (*C.gchar)(unsafe.Pointer(C.CString(detailed_action)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_menu_item_new(__cgo__label, __cgo__detailed_action)
	C.free(unsafe.Pointer(__cgo__label))
	C.free(unsafe.Pointer(__cgo__detailed_action))
	if __cgo__return__ != nil {
		return__ = NewMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GMenuItem as an exact copy of an existing menu item in a
#GMenuModel.

@item_index must be valid (ie: be sure to call
g_menu_model_get_n_items() first).
*/
func MenuItemNewFromModel(model IsMenuModel, item_index int) (return__ *MenuItem) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_menu_item_new_from_model(model.GetMenuModelPointer(), C.gint(item_index))
	if __cgo__return__ != nil {
		return__ = NewMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GMenuItem representing a section.

This is a convenience API around g_menu_item_new() and
g_menu_item_set_section().

The effect of having one menu appear as a section of another is
exactly as it sounds: the items from @section become a direct part of
the menu that @menu_item is added to.

Visual separation is typically displayed between two non-empty
sections.  If @label is non-%NULL then it will be encorporated into
this visual indication.  This allows for labeled subsections of a
menu.

As a simple example, consider a typical "Edit" menu from a simple
program.  It probably contains an "Undo" and "Redo" item, followed by
a separator, followed by "Cut", "Copy" and "Paste".

This would be accomplished by creating three #GMenu instances.  The
first would be populated with the "Undo" and "Redo" items, and the
second with the "Cut", "Copy" and "Paste" items.  The first and
second menus would then be added as submenus of the third.  In XML
format, this would look something like the following:
|[
<menu id='edit-menu'>
  <section>
    <item label='Undo'/>
    <item label='Redo'/>
  </section>
  <section>
    <item label='Cut'/>
    <item label='Copy'/>
    <item label='Paste'/>
  </section>
</menu>
]|

The following example is exactly equivalent.  It is more illustrative
of the exact relationship between the menus and items (keeping in
mind that the 'link' element defines a new menu that is linked to the
containing one).  The style of the second example is more verbose and
difficult to read (and therefore not recommended except for the
purpose of understanding what is really going on).
|[
<menu id='edit-menu'>
  <item>
    <link name='section'>
      <item label='Undo'/>
      <item label='Redo'/>
    </link>
  </item>
  <item>
    <link name='section'>
      <item label='Cut'/>
      <item label='Copy'/>
      <item label='Paste'/>
    </link>
  </item>
</menu>
]|
*/
func MenuItemNewSection(label string, section IsMenuModel) (return__ *MenuItem) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_menu_item_new_section(__cgo__label, section.GetMenuModelPointer())
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GMenuItem representing a submenu.

This is a convenience API around g_menu_item_new() and
g_menu_item_set_submenu().
*/
func MenuItemNewSubmenu(label string, submenu IsMenuModel) (return__ *MenuItem) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_menu_item_new_submenu(__cgo__label, submenu.GetMenuModelPointer())
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new mount operation.
*/
func MountOperationNew() (return__ *MountOperation) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_mount_operation_new()
	if __cgo__return__ != nil {
		return__ = NewMountOperationFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GSocketConnectable for connecting to the given
@hostname and @port.
*/
func NetworkAddressNew(hostname string, port uint16) (return__ *NetworkAddress) {
	__cgo__hostname := (*C.gchar)(unsafe.Pointer(C.CString(hostname)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_network_address_new(__cgo__hostname, C.guint16(port))
	C.free(unsafe.Pointer(__cgo__hostname))
	if __cgo__return__ != nil {
		return__ = NewNetworkAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GNetworkService representing the given @service,
@protocol, and @domain. This will initially be unresolved; use the
#GSocketConnectable interface to resolve it.
*/
func NetworkServiceNew(service string, protocol string, domain string) (return__ *NetworkService) {
	__cgo__service := (*C.gchar)(unsafe.Pointer(C.CString(service)))
	__cgo__protocol := (*C.gchar)(unsafe.Pointer(C.CString(protocol)))
	__cgo__domain := (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_network_service_new(__cgo__service, __cgo__protocol, __cgo__domain)
	C.free(unsafe.Pointer(__cgo__service))
	C.free(unsafe.Pointer(__cgo__protocol))
	C.free(unsafe.Pointer(__cgo__domain))
	if __cgo__return__ != nil {
		return__ = NewNetworkServiceFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GNotification with @title as its title.

After populating @notification with more details, it can be sent to
the desktop shell with g_application_send_notification(). Changing
any properties after this call will not have any effect until
resending @notification.
*/
func NotificationNew(title string) (return__ *Notification) {
	__cgo__title := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_notification_new(__cgo__title)
	C.free(unsafe.Pointer(__cgo__title))
	if __cgo__return__ != nil {
		return__ = NewNotificationFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GAction corresponding to the value of property
@property_name on @object.

The property must be existent and readable and writable (and not
construct-only).

This function takes a reference on @object and doesn't release it
until the action is destroyed.
*/
func PropertyActionNew(name string, object unsafe.Pointer, property_name string) (return__ *PropertyAction) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__property_name := (*C.gchar)(unsafe.Pointer(C.CString(property_name)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_property_action_new(__cgo__name, (C.gpointer)(object), __cgo__property_name)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__property_name))
	if __cgo__return__ != nil {
		return__ = NewPropertyActionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GProxyAddress for @inetaddr with @protocol that should
tunnel through @dest_hostname and @dest_port.

(Note that this method doesn't set the #GProxyAddress:uri or
#GProxyAddress:destination-protocol fields; use g_object_new()
directly if you want to set those.)
*/
func ProxyAddressNew(inetaddr IsInetAddress, port uint16, protocol string, dest_hostname string, dest_port uint16, username string, password string) (return__ *ProxyAddress) {
	__cgo__protocol := (*C.gchar)(unsafe.Pointer(C.CString(protocol)))
	__cgo__dest_hostname := (*C.gchar)(unsafe.Pointer(C.CString(dest_hostname)))
	__cgo__username := (*C.gchar)(unsafe.Pointer(C.CString(username)))
	__cgo__password := (*C.gchar)(unsafe.Pointer(C.CString(password)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_proxy_address_new(inetaddr.GetInetAddressPointer(), C.guint16(port), __cgo__protocol, __cgo__dest_hostname, C.guint16(dest_port), __cgo__username, __cgo__password)
	C.free(unsafe.Pointer(__cgo__protocol))
	C.free(unsafe.Pointer(__cgo__dest_hostname))
	C.free(unsafe.Pointer(__cgo__username))
	C.free(unsafe.Pointer(__cgo__password))
	if __cgo__return__ != nil {
		return__ = NewProxyAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GSettings object with the schema specified by
@schema_id.

Signals on the newly created #GSettings object will be dispatched
via the thread-default #GMainContext in effect at the time of the
call to g_settings_new().  The new #GSettings will hold a reference
on the context.  See g_main_context_push_thread_default().
*/
func SettingsNew(schema_id string) (return__ *Settings) {
	__cgo__schema_id := (*C.gchar)(unsafe.Pointer(C.CString(schema_id)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_settings_new(__cgo__schema_id)
	C.free(unsafe.Pointer(__cgo__schema_id))
	if __cgo__return__ != nil {
		return__ = NewSettingsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GSettings object with a given schema, backend and
path.

It should be extremely rare that you ever want to use this function.
It is made available for advanced use-cases (such as plugin systems
that want to provide access to schemas loaded from custom locations,
etc).

At the most basic level, a #GSettings object is a pure composition of
4 things: a #GSettingsSchema, a #GSettingsBackend, a path within that
backend, and a #GMainContext to which signals are dispatched.

This constructor therefore gives you full control over constructing
#GSettings instances.  The first 4 parameters are given directly as
@schema, @backend and @path, and the main context is taken from the
thread-default (as per g_settings_new()).

If @backend is %NULL then the default backend is used.

If @path is %NULL then the path from the schema is used.  It is an
error f @path is %NULL and the schema has no path of its own or if
@path is non-%NULL and not equal to the path that the schema does
have.
*/
func SettingsNewFull(schema *C.GSettingsSchema, backend *C.GSettingsBackend, path string) (return__ *Settings) {
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_settings_new_full(schema, backend, __cgo__path)
	C.free(unsafe.Pointer(__cgo__path))
	if __cgo__return__ != nil {
		return__ = NewSettingsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GSettings object with the schema specified by
@schema_id and a given #GSettingsBackend.

Creating a #GSettings object with a different backend allows accessing
settings from a database other than the usual one. For example, it may make
sense to pass a backend corresponding to the "defaults" settings database on
the system to get a settings object that modifies the system default
settings instead of the settings for this user.
*/
func SettingsNewWithBackend(schema_id string, backend *C.GSettingsBackend) (return__ *Settings) {
	__cgo__schema_id := (*C.gchar)(unsafe.Pointer(C.CString(schema_id)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_settings_new_with_backend(__cgo__schema_id, backend)
	C.free(unsafe.Pointer(__cgo__schema_id))
	if __cgo__return__ != nil {
		return__ = NewSettingsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GSettings object with the schema specified by
@schema_id and a given #GSettingsBackend and path.

This is a mix of g_settings_new_with_backend() and
g_settings_new_with_path().
*/
func SettingsNewWithBackendAndPath(schema_id string, backend *C.GSettingsBackend, path string) (return__ *Settings) {
	__cgo__schema_id := (*C.gchar)(unsafe.Pointer(C.CString(schema_id)))
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_settings_new_with_backend_and_path(__cgo__schema_id, backend, __cgo__path)
	C.free(unsafe.Pointer(__cgo__schema_id))
	C.free(unsafe.Pointer(__cgo__path))
	if __cgo__return__ != nil {
		return__ = NewSettingsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GSettings object with the relocatable schema specified
by @schema_id and a given path.

You only need to do this if you want to directly create a settings
object with a schema that doesn't have a specified path of its own.
That's quite rare.

It is a programmer error to call this function for a schema that
has an explicitly specified path.

It is a programmer error if @path is not a valid path.  A valid path
begins and ends with '/' and does not contain two consecutive '/'
characters.
*/
func SettingsNewWithPath(schema_id string, path string) (return__ *Settings) {
	__cgo__schema_id := (*C.gchar)(unsafe.Pointer(C.CString(schema_id)))
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_settings_new_with_path(__cgo__schema_id, __cgo__path)
	C.free(unsafe.Pointer(__cgo__schema_id))
	C.free(unsafe.Pointer(__cgo__path))
	if __cgo__return__ != nil {
		return__ = NewSettingsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new action.

The created action is stateless.  See g_simple_action_new_stateful().
*/
func SimpleActionNew(name string, parameter_type *C.GVariantType) (return__ *SimpleAction) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_simple_action_new(__cgo__name, parameter_type)
	C.free(unsafe.Pointer(__cgo__name))
	if __cgo__return__ != nil {
		return__ = NewSimpleActionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new stateful action.

@state is the initial state of the action.  All future state values
must have the same #GVariantType as the initial state.

If the @state GVariant is floating, it is consumed.
*/
func SimpleActionNewStateful(name string, parameter_type *C.GVariantType, state *C.GVariant) (return__ *SimpleAction) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_simple_action_new_stateful(__cgo__name, parameter_type, state)
	C.free(unsafe.Pointer(__cgo__name))
	if __cgo__return__ != nil {
		return__ = NewSimpleActionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new, empty, #GSimpleActionGroup.
*/
func SimpleActionGroupNew() (return__ *SimpleActionGroup) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_simple_action_group_new()
	if __cgo__return__ != nil {
		return__ = NewSimpleActionGroupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GSimpleAsyncResult.

The common convention is to create the #GSimpleAsyncResult in the
function that starts the asynchronous operation and use that same
function as the @source_tag.

If your operation supports cancellation with #GCancellable (which it
probably should) then you should provide the user's cancellable to
g_simple_async_result_set_check_cancellable() immediately after
this function returns.
*/
func SimpleAsyncResultNew(source_object *C.GObject, callback C.GAsyncReadyCallback, user_data unsafe.Pointer, source_tag unsafe.Pointer) (return__ *SimpleAsyncResult) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_simple_async_result_new(source_object, callback, (C.gpointer)(user_data), (C.gpointer)(source_tag))
	if __cgo__return__ != nil {
		return__ = NewSimpleAsyncResultFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// g_simple_async_result_new_error is not generated due to varargs

/*
Creates a #GSimpleAsyncResult from an error condition.
*/
func SimpleAsyncResultNewFromError(source_object *C.GObject, callback C.GAsyncReadyCallback, user_data unsafe.Pointer, error_ *C.GError) (return__ *SimpleAsyncResult) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_simple_async_result_new_from_error(source_object, callback, (C.gpointer)(user_data), error_)
	if __cgo__return__ != nil {
		return__ = NewSimpleAsyncResultFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GSimpleAsyncResult from an error condition, and takes over the
caller's ownership of @error, so the caller does not need to free it anymore.
*/
func SimpleAsyncResultNewTakeError(source_object *C.GObject, callback C.GAsyncReadyCallback, user_data unsafe.Pointer, error_ *C.GError) (return__ *SimpleAsyncResult) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_simple_async_result_new_take_error(source_object, callback, (C.gpointer)(user_data), error_)
	if __cgo__return__ != nil {
		return__ = NewSimpleAsyncResultFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GPermission instance that represents an action that is
either always or never allowed.
*/
func SimplePermissionNew(allowed bool) (return__ *SimplePermission) {
	__cgo__allowed := C.gboolean(0)
	if allowed {
		__cgo__allowed = C.gboolean(1)
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_simple_permission_new(__cgo__allowed)
	if __cgo__return__ != nil {
		return__ = NewSimplePermissionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GSocket with the defined family, type and protocol.
If @protocol is 0 (%G_SOCKET_PROTOCOL_DEFAULT) the default protocol type
for the family and type is used.

The @protocol is a family and type specific int that specifies what
kind of protocol to use. #GSocketProtocol lists several common ones.
Many families only support one protocol, and use 0 for this, others
support several and using 0 means to use the default protocol for
the family and type.

The protocol id is passed directly to the operating
system, so you can use protocols not listed in #GSocketProtocol if you
know the protocol number used for it.
*/
func SocketNew(family C.GSocketFamily, type_ C.GSocketType, protocol C.GSocketProtocol) (return__ *Socket, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_socket_new(family, type_, protocol, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new #GSocket from a native file descriptor
or winsock SOCKET handle.

This reads all the settings from the file descriptor so that
all properties should work. Note that the file descriptor
will be set to non-blocking mode, independent on the blocking
mode of the #GSocket.
*/
func SocketNewFromFd(fd int) (return__ *Socket, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_socket_new_from_fd(C.gint(fd), &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a #GSocketAddress subclass corresponding to the native
struct sockaddr @native.
*/
func SocketAddressNewFromNative(native unsafe.Pointer, len_ int64) (return__ *SocketAddress) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_socket_address_new_from_native((C.gpointer)(native), C.gsize(len_))
	if __cgo__return__ != nil {
		return__ = NewSocketAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GSocketClient with the default options.
*/
func SocketClientNew() (return__ *SocketClient) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_socket_client_new()
	if __cgo__return__ != nil {
		return__ = NewSocketClientFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GSocketListener with no sockets to listen for.
New listeners can be added with e.g. g_socket_listener_add_address()
or g_socket_listener_add_inet_port().
*/
func SocketListenerNew() (return__ *SocketListener) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_socket_listener_new()
	if __cgo__return__ != nil {
		return__ = NewSocketListenerFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GSocketService with no sockets to listen for.
New listeners can be added with e.g. g_socket_listener_add_address()
or g_socket_listener_add_inet_port().
*/
func SocketServiceNew() (return__ *SocketService) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_socket_service_new()
	if __cgo__return__ != nil {
		return__ = NewSocketServiceFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// g_subprocess_new is not generated due to varargs

/*
Create a new process with the given flags and argument list.

The argument list is expected to be %NULL-terminated.
*/
func SubprocessNewv(argv []string, flags C.GSubprocessFlags) (return__ *Subprocess, __err__ error) {
	__header__argv := (*reflect.SliceHeader)(unsafe.Pointer(&argv))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_subprocess_newv((**C.gchar)(unsafe.Pointer(__header__argv.Data)), flags, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSubprocessFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new #GSubprocessLauncher.

The launcher is created with the default options.  A copy of the
environment of the calling process is made at the time of this call
and will be used as the environment that the process is launched in.
*/
func SubprocessLauncherNew(flags C.GSubprocessFlags) (return__ *SubprocessLauncher) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_subprocess_launcher_new(flags)
	if __cgo__return__ != nil {
		return__ = NewSubprocessLauncherFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GTask acting on @source_object, which will eventually be
used to invoke @callback in the current
[thread-default main context][g-main-context-push-thread-default].

Call this in the "start" method of your asynchronous method, and
pass the #GTask around throughout the asynchronous operation. You
can use g_task_set_task_data() to attach task-specific data to the
object, which you can retrieve later via g_task_get_task_data().

By default, if @cancellable is cancelled, then the return value of
the task will always be %G_IO_ERROR_CANCELLED, even if the task had
already completed before the cancellation. This allows for
simplified handling in cases where cancellation may imply that
other objects that the task depends on have been destroyed. If you
do not want this behavior, you can use
g_task_set_check_cancellable() to change it.
*/
func TaskNew(source_object gobject.IsObject, cancellable IsCancellable, callback C.GAsyncReadyCallback, callback_data unsafe.Pointer) (return__ *Task) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_task_new(C.gpointer(unsafe.Pointer(reflect.ValueOf(source_object.GetObjectPointer()).Pointer())), cancellable.GetCancellablePointer(), callback, (C.gpointer)(callback_data))
	if __cgo__return__ != nil {
		return__ = NewTaskFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Wraps @base_io_stream and @socket together as a #GSocketConnection.
*/
func TcpWrapperConnectionNew(base_io_stream IsIOStream, socket IsSocket) (return__ *TcpWrapperConnection) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_tcp_wrapper_connection_new(base_io_stream.GetIOStreamPointer(), socket.GetSocketPointer())
	if __cgo__return__ != nil {
		return__ = NewTcpWrapperConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Create a new #GTestDBus object.
*/
func TestDBusNew(flags C.GTestDBusFlags) (return__ *TestDBus) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_test_dbus_new(flags)
	if __cgo__return__ != nil {
		return__ = NewTestDBusFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new themed icon for @iconname.
*/
func ThemedIconNew(iconname string) (return__ *ThemedIcon) {
	__cgo__iconname := C.CString(iconname)
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_themed_icon_new(__cgo__iconname)
	C.free(unsafe.Pointer(__cgo__iconname))
	if __cgo__return__ != nil {
		return__ = NewThemedIconFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new themed icon for @iconnames.
*/
func ThemedIconNewFromNames(iconnames []string, len_ int) (return__ *ThemedIcon) {
	__header__iconnames := (*reflect.SliceHeader)(unsafe.Pointer(&iconnames))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_themed_icon_new_from_names((**C.char)(unsafe.Pointer(__header__iconnames.Data)), C.int(len_))
	if __cgo__return__ != nil {
		return__ = NewThemedIconFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new themed icon for @iconname, and all the names
that can be created by shortening @iconname at '-' characters.

In the following example, @icon1 and @icon2 are equivalent:
|[<!-- language="C" -->
const char *names[] = {
  "gnome-dev-cdrom-audio",
  "gnome-dev-cdrom",
  "gnome-dev",
  "gnome"
};

icon1 = g_themed_icon_new_from_names (names, 4);
icon2 = g_themed_icon_new_with_default_fallbacks ("gnome-dev-cdrom-audio");
]|
*/
func ThemedIconNewWithDefaultFallbacks(iconname string) (return__ *ThemedIcon) {
	__cgo__iconname := C.CString(iconname)
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_themed_icon_new_with_default_fallbacks(__cgo__iconname)
	C.free(unsafe.Pointer(__cgo__iconname))
	if __cgo__return__ != nil {
		return__ = NewThemedIconFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GThreadedSocketService with no listeners. Listeners
must be added with one of the #GSocketListener "add" methods.
*/
func ThreadedSocketServiceNew(max_threads int) (return__ *ThreadedSocketService) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_threaded_socket_service_new(C.int(max_threads))
	if __cgo__return__ != nil {
		return__ = NewThreadedSocketServiceFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GTlsCertificate from the PEM-encoded data in @file. If
@file cannot be read or parsed, the function will return %NULL and
set @error. Otherwise, this behaves like
g_tls_certificate_new_from_pem().
*/
func TlsCertificateNewFromFile(file string) (return__ *TlsCertificate, __err__ error) {
	__cgo__file := (*C.gchar)(unsafe.Pointer(C.CString(file)))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_tls_certificate_new_from_file(__cgo__file, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__file))
	if __cgo__return__ != nil {
		return__ = NewTlsCertificateFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a #GTlsCertificate from the PEM-encoded data in @cert_file
and @key_file. If either file cannot be read or parsed, the
function will return %NULL and set @error. Otherwise, this behaves
like g_tls_certificate_new_from_pem().
*/
func TlsCertificateNewFromFiles(cert_file string, key_file string) (return__ *TlsCertificate, __err__ error) {
	__cgo__cert_file := (*C.gchar)(unsafe.Pointer(C.CString(cert_file)))
	__cgo__key_file := (*C.gchar)(unsafe.Pointer(C.CString(key_file)))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_tls_certificate_new_from_files(__cgo__cert_file, __cgo__key_file, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__cert_file))
	C.free(unsafe.Pointer(__cgo__key_file))
	if __cgo__return__ != nil {
		return__ = NewTlsCertificateFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new #GTlsCertificate from the PEM-encoded data in @data.
If @data includes both a certificate and a private key, then the
returned certificate will include the private key data as well. (See
the #GTlsCertificate:private-key-pem property for information about
supported formats.)

If @data includes multiple certificates, only the first one will be
parsed.
*/
func TlsCertificateNewFromPem(data string, length int64) (return__ *TlsCertificate, __err__ error) {
	__cgo__data := (*C.gchar)(unsafe.Pointer(C.CString(data)))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_tls_certificate_new_from_pem(__cgo__data, C.gssize(length), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__data))
	if __cgo__return__ != nil {
		return__ = NewTlsCertificateFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Create a new #GTlsPassword object.
*/
func TlsPasswordNew(flags C.GTlsPasswordFlags, description string) (return__ *TlsPassword) {
	__cgo__description := (*C.gchar)(unsafe.Pointer(C.CString(description)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_tls_password_new(flags, __cgo__description)
	C.free(unsafe.Pointer(__cgo__description))
	if __cgo__return__ != nil {
		return__ = NewTlsPasswordFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GUnixCredentialsMessage with credentials matching the current processes.
*/
func UnixCredentialsMessageNew() (return__ *UnixCredentialsMessage) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_unix_credentials_message_new()
	if __cgo__return__ != nil {
		return__ = NewUnixCredentialsMessageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GUnixCredentialsMessage holding @credentials.
*/
func UnixCredentialsMessageNewWithCredentials(credentials IsCredentials) (return__ *UnixCredentialsMessage) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_unix_credentials_message_new_with_credentials(credentials.GetCredentialsPointer())
	if __cgo__return__ != nil {
		return__ = NewUnixCredentialsMessageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GUnixFDList containing no file descriptors.
*/
func UnixFDListNew() (return__ *UnixFDList) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_unix_fd_list_new()
	if __cgo__return__ != nil {
		return__ = NewUnixFDListFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GUnixFDList containing the file descriptors given in
@fds.  The file descriptors become the property of the new list and
may no longer be used by the caller.  The array itself is owned by
the caller.

Each file descriptor in the array should be set to close-on-exec.

If @n_fds is -1 then @fds must be terminated with -1.
*/
func UnixFDListNewFromArray(fds []int, n_fds int) (return__ *UnixFDList) {
	__header__fds := (*reflect.SliceHeader)(unsafe.Pointer(&fds))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_unix_fd_list_new_from_array((*C.gint)(unsafe.Pointer(__header__fds.Data)), C.gint(n_fds))
	if __cgo__return__ != nil {
		return__ = NewUnixFDListFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GUnixFDMessage containing an empty file descriptor
list.
*/
func UnixFDMessageNew() (return__ *UnixFDMessage) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_unix_fd_message_new()
	if __cgo__return__ != nil {
		return__ = NewUnixFDMessageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GUnixFDMessage containing @list.
*/
func UnixFDMessageNewWithFdList(fd_list IsUnixFDList) (return__ *UnixFDMessage) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_unix_fd_message_new_with_fd_list(fd_list.GetUnixFDListPointer())
	if __cgo__return__ != nil {
		return__ = NewUnixFDMessageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GUnixInputStream for the given @fd.

If @close_fd is %TRUE, the file descriptor will be closed
when the stream is closed.
*/
func UnixInputStreamNew(fd int, close_fd bool) (return__ *UnixInputStream) {
	__cgo__close_fd := C.gboolean(0)
	if close_fd {
		__cgo__close_fd = C.gboolean(1)
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_unix_input_stream_new(C.gint(fd), __cgo__close_fd)
	if __cgo__return__ != nil {
		return__ = NewUnixInputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets a new #GUnixMountMonitor. The default rate limit for which the
monitor will report consecutive changes for the mount and mount
point entry files is the default for a #GFileMonitor. Use
g_unix_mount_monitor_set_rate_limit() to change this.
*/
func UnixMountMonitorNew() (return__ *UnixMountMonitor) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_unix_mount_monitor_new()
	if __cgo__return__ != nil {
		return__ = NewUnixMountMonitorFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GUnixOutputStream for the given @fd.

If @close_fd, is %TRUE, the file descriptor will be closed when
the output stream is destroyed.
*/
func UnixOutputStreamNew(fd int, close_fd bool) (return__ *UnixOutputStream) {
	__cgo__close_fd := C.gboolean(0)
	if close_fd {
		__cgo__close_fd = C.gboolean(1)
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_unix_output_stream_new(C.gint(fd), __cgo__close_fd)
	if __cgo__return__ != nil {
		return__ = NewUnixOutputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GUnixSocketAddress for @path.

To create abstract socket addresses, on systems that support that,
use g_unix_socket_address_new_abstract().
*/
func UnixSocketAddressNew(path string) (return__ *UnixSocketAddress) {
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_unix_socket_address_new(__cgo__path)
	C.free(unsafe.Pointer(__cgo__path))
	if __cgo__return__ != nil {
		return__ = NewUnixSocketAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// g_unix_socket_address_new_abstract is not generated due to deprecation attr

/*
Creates a new #GUnixSocketAddress of type @type with name @path.

If @type is %G_UNIX_SOCKET_ADDRESS_PATH, this is equivalent to
calling g_unix_socket_address_new().

If @path_type is %G_UNIX_SOCKET_ADDRESS_ABSTRACT, then @path_len
bytes of @path will be copied to the socket's path, and only those
bytes will be considered part of the name. (If @path_len is -1,
then @path is assumed to be NUL-terminated.) For example, if @path
was "test", then calling g_socket_address_get_native_size() on the
returned socket would return 7 (2 bytes of overhead, 1 byte for the
abstract-socket indicator byte, and 4 bytes for the name "test").

If @path_type is %G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED, then
@path_len bytes of @path will be copied to the socket's path, the
rest of the path will be padded with 0 bytes, and the entire
zero-padded buffer will be considered the name. (As above, if
@path_len is -1, then @path is assumed to be NUL-terminated.) In
this case, g_socket_address_get_native_size() will always return
the full size of a `struct sockaddr_un`, although
g_unix_socket_address_get_path_len() will still return just the
length of @path.

%G_UNIX_SOCKET_ADDRESS_ABSTRACT is preferred over
%G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED for new programs. Of course,
when connecting to a server created by another process, you must
use the appropriate type corresponding to how that process created
its listening socket.
*/
func UnixSocketAddressNewWithType(path string, path_len int, type_ C.GUnixSocketAddressType) (return__ *UnixSocketAddress) {
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_unix_socket_address_new_with_type(__cgo__path, C.gint(path_len), type_)
	C.free(unsafe.Pointer(__cgo__path))
	if __cgo__return__ != nil {
		return__ = NewUnixSocketAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GZlibCompressor.
*/
func ZlibCompressorNew(format C.GZlibCompressorFormat, level int) (return__ *ZlibCompressor) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_zlib_compressor_new(format, C.int(level))
	if __cgo__return__ != nil {
		return__ = NewZlibCompressorFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GZlibDecompressor.
*/
func ZlibDecompressorNew(format C.GZlibCompressorFormat) (return__ *ZlibDecompressor) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.g_zlib_decompressor_new(format)
	if __cgo__return__ != nil {
		return__ = NewZlibDecompressorFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}
