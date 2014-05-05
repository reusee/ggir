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
#include <glib-object.h>
#cgo pkg-config: gobject-2.0
*/
import "C"
import "unsafe"
import "github.com/reusee/ggir/gobject"
import "github.com/reusee/ggir/glib"
import "runtime"

func init() {
	_ = unsafe.Pointer(nil)
	_ = runtime.Compiler
}

var _ = gobject.UnusedFix_
var _ = glib.UnusedFix_

type AppInfoMonitor struct {
	*TraitAppInfoMonitor
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAppInfoMonitorFromCPointer(p unsafe.Pointer) *AppInfoMonitor {
	ret := &AppInfoMonitor{
		NewTraitAppInfoMonitor(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *AppInfoMonitor) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type AppLaunchContext struct {
	*TraitAppLaunchContext
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAppLaunchContextFromCPointer(p unsafe.Pointer) *AppLaunchContext {
	ret := &AppLaunchContext{
		NewTraitAppLaunchContext(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *AppLaunchContext) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Application struct {
	*TraitApplication
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewApplicationFromCPointer(p unsafe.Pointer) *Application {
	ret := &Application{
		NewTraitApplication(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Application) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ApplicationCommandLine struct {
	*TraitApplicationCommandLine
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewApplicationCommandLineFromCPointer(p unsafe.Pointer) *ApplicationCommandLine {
	ret := &ApplicationCommandLine{
		NewTraitApplicationCommandLine(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ApplicationCommandLine) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type BufferedInputStream struct {
	*TraitBufferedInputStream
	*TraitFilterInputStream
	*TraitInputStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewBufferedInputStreamFromCPointer(p unsafe.Pointer) *BufferedInputStream {
	ret := &BufferedInputStream{
		NewTraitBufferedInputStream(p),
		NewTraitFilterInputStream(p),
		NewTraitInputStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *BufferedInputStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type BufferedOutputStream struct {
	*TraitBufferedOutputStream
	*TraitFilterOutputStream
	*TraitOutputStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewBufferedOutputStreamFromCPointer(p unsafe.Pointer) *BufferedOutputStream {
	ret := &BufferedOutputStream{
		NewTraitBufferedOutputStream(p),
		NewTraitFilterOutputStream(p),
		NewTraitOutputStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *BufferedOutputStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type BytesIcon struct {
	*TraitBytesIcon
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewBytesIconFromCPointer(p unsafe.Pointer) *BytesIcon {
	ret := &BytesIcon{
		NewTraitBytesIcon(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *BytesIcon) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Cancellable struct {
	*TraitCancellable
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCancellableFromCPointer(p unsafe.Pointer) *Cancellable {
	ret := &Cancellable{
		NewTraitCancellable(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Cancellable) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type CharsetConverter struct {
	*TraitCharsetConverter
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCharsetConverterFromCPointer(p unsafe.Pointer) *CharsetConverter {
	ret := &CharsetConverter{
		NewTraitCharsetConverter(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CharsetConverter) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ConverterInputStream struct {
	*TraitConverterInputStream
	*TraitFilterInputStream
	*TraitInputStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewConverterInputStreamFromCPointer(p unsafe.Pointer) *ConverterInputStream {
	ret := &ConverterInputStream{
		NewTraitConverterInputStream(p),
		NewTraitFilterInputStream(p),
		NewTraitInputStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ConverterInputStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ConverterOutputStream struct {
	*TraitConverterOutputStream
	*TraitFilterOutputStream
	*TraitOutputStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewConverterOutputStreamFromCPointer(p unsafe.Pointer) *ConverterOutputStream {
	ret := &ConverterOutputStream{
		NewTraitConverterOutputStream(p),
		NewTraitFilterOutputStream(p),
		NewTraitOutputStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ConverterOutputStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Credentials struct {
	*TraitCredentials
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCredentialsFromCPointer(p unsafe.Pointer) *Credentials {
	ret := &Credentials{
		NewTraitCredentials(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Credentials) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DBusActionGroup struct {
	*TraitDBusActionGroup
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDBusActionGroupFromCPointer(p unsafe.Pointer) *DBusActionGroup {
	ret := &DBusActionGroup{
		NewTraitDBusActionGroup(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DBusActionGroup) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DBusAuthObserver struct {
	*TraitDBusAuthObserver
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDBusAuthObserverFromCPointer(p unsafe.Pointer) *DBusAuthObserver {
	ret := &DBusAuthObserver{
		NewTraitDBusAuthObserver(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DBusAuthObserver) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DBusConnection struct {
	*TraitDBusConnection
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDBusConnectionFromCPointer(p unsafe.Pointer) *DBusConnection {
	ret := &DBusConnection{
		NewTraitDBusConnection(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DBusConnection) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DBusInterfaceSkeleton struct {
	*TraitDBusInterfaceSkeleton
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDBusInterfaceSkeletonFromCPointer(p unsafe.Pointer) *DBusInterfaceSkeleton {
	ret := &DBusInterfaceSkeleton{
		NewTraitDBusInterfaceSkeleton(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DBusInterfaceSkeleton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DBusMenuModel struct {
	*TraitDBusMenuModel
	*TraitMenuModel
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDBusMenuModelFromCPointer(p unsafe.Pointer) *DBusMenuModel {
	ret := &DBusMenuModel{
		NewTraitDBusMenuModel(p),
		NewTraitMenuModel(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DBusMenuModel) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DBusMessage struct {
	*TraitDBusMessage
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDBusMessageFromCPointer(p unsafe.Pointer) *DBusMessage {
	ret := &DBusMessage{
		NewTraitDBusMessage(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DBusMessage) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DBusMethodInvocation struct {
	*TraitDBusMethodInvocation
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDBusMethodInvocationFromCPointer(p unsafe.Pointer) *DBusMethodInvocation {
	ret := &DBusMethodInvocation{
		NewTraitDBusMethodInvocation(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DBusMethodInvocation) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DBusObjectManagerClient struct {
	*TraitDBusObjectManagerClient
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDBusObjectManagerClientFromCPointer(p unsafe.Pointer) *DBusObjectManagerClient {
	ret := &DBusObjectManagerClient{
		NewTraitDBusObjectManagerClient(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DBusObjectManagerClient) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DBusObjectManagerServer struct {
	*TraitDBusObjectManagerServer
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDBusObjectManagerServerFromCPointer(p unsafe.Pointer) *DBusObjectManagerServer {
	ret := &DBusObjectManagerServer{
		NewTraitDBusObjectManagerServer(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DBusObjectManagerServer) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DBusObjectProxy struct {
	*TraitDBusObjectProxy
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDBusObjectProxyFromCPointer(p unsafe.Pointer) *DBusObjectProxy {
	ret := &DBusObjectProxy{
		NewTraitDBusObjectProxy(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DBusObjectProxy) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DBusObjectSkeleton struct {
	*TraitDBusObjectSkeleton
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDBusObjectSkeletonFromCPointer(p unsafe.Pointer) *DBusObjectSkeleton {
	ret := &DBusObjectSkeleton{
		NewTraitDBusObjectSkeleton(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DBusObjectSkeleton) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DBusProxy struct {
	*TraitDBusProxy
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDBusProxyFromCPointer(p unsafe.Pointer) *DBusProxy {
	ret := &DBusProxy{
		NewTraitDBusProxy(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DBusProxy) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DBusServer struct {
	*TraitDBusServer
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDBusServerFromCPointer(p unsafe.Pointer) *DBusServer {
	ret := &DBusServer{
		NewTraitDBusServer(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DBusServer) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DataInputStream struct {
	*TraitDataInputStream
	*TraitBufferedInputStream
	*TraitFilterInputStream
	*TraitInputStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDataInputStreamFromCPointer(p unsafe.Pointer) *DataInputStream {
	ret := &DataInputStream{
		NewTraitDataInputStream(p),
		NewTraitBufferedInputStream(p),
		NewTraitFilterInputStream(p),
		NewTraitInputStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DataInputStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DataOutputStream struct {
	*TraitDataOutputStream
	*TraitFilterOutputStream
	*TraitOutputStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDataOutputStreamFromCPointer(p unsafe.Pointer) *DataOutputStream {
	ret := &DataOutputStream{
		NewTraitDataOutputStream(p),
		NewTraitFilterOutputStream(p),
		NewTraitOutputStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DataOutputStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type DesktopAppInfo struct {
	*TraitDesktopAppInfo
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDesktopAppInfoFromCPointer(p unsafe.Pointer) *DesktopAppInfo {
	ret := &DesktopAppInfo{
		NewTraitDesktopAppInfo(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *DesktopAppInfo) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Emblem struct {
	*TraitEmblem
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewEmblemFromCPointer(p unsafe.Pointer) *Emblem {
	ret := &Emblem{
		NewTraitEmblem(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Emblem) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type EmblemedIcon struct {
	*TraitEmblemedIcon
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewEmblemedIconFromCPointer(p unsafe.Pointer) *EmblemedIcon {
	ret := &EmblemedIcon{
		NewTraitEmblemedIcon(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *EmblemedIcon) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type FileEnumerator struct {
	*TraitFileEnumerator
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFileEnumeratorFromCPointer(p unsafe.Pointer) *FileEnumerator {
	ret := &FileEnumerator{
		NewTraitFileEnumerator(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FileEnumerator) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type FileIOStream struct {
	*TraitFileIOStream
	*TraitIOStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFileIOStreamFromCPointer(p unsafe.Pointer) *FileIOStream {
	ret := &FileIOStream{
		NewTraitFileIOStream(p),
		NewTraitIOStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FileIOStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type FileIcon struct {
	*TraitFileIcon
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFileIconFromCPointer(p unsafe.Pointer) *FileIcon {
	ret := &FileIcon{
		NewTraitFileIcon(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FileIcon) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type FileInfo struct {
	*TraitFileInfo
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFileInfoFromCPointer(p unsafe.Pointer) *FileInfo {
	ret := &FileInfo{
		NewTraitFileInfo(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FileInfo) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type FileInputStream struct {
	*TraitFileInputStream
	*TraitInputStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFileInputStreamFromCPointer(p unsafe.Pointer) *FileInputStream {
	ret := &FileInputStream{
		NewTraitFileInputStream(p),
		NewTraitInputStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FileInputStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type FileMonitor struct {
	*TraitFileMonitor
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFileMonitorFromCPointer(p unsafe.Pointer) *FileMonitor {
	ret := &FileMonitor{
		NewTraitFileMonitor(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FileMonitor) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type FileOutputStream struct {
	*TraitFileOutputStream
	*TraitOutputStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFileOutputStreamFromCPointer(p unsafe.Pointer) *FileOutputStream {
	ret := &FileOutputStream{
		NewTraitFileOutputStream(p),
		NewTraitOutputStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FileOutputStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type FilenameCompleter struct {
	*TraitFilenameCompleter
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFilenameCompleterFromCPointer(p unsafe.Pointer) *FilenameCompleter {
	ret := &FilenameCompleter{
		NewTraitFilenameCompleter(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FilenameCompleter) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type FilterInputStream struct {
	*TraitFilterInputStream
	*TraitInputStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFilterInputStreamFromCPointer(p unsafe.Pointer) *FilterInputStream {
	ret := &FilterInputStream{
		NewTraitFilterInputStream(p),
		NewTraitInputStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FilterInputStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type FilterOutputStream struct {
	*TraitFilterOutputStream
	*TraitOutputStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFilterOutputStreamFromCPointer(p unsafe.Pointer) *FilterOutputStream {
	ret := &FilterOutputStream{
		NewTraitFilterOutputStream(p),
		NewTraitOutputStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FilterOutputStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type IOModule struct {
	*TraitIOModule
	*gobject.TraitTypeModule
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewIOModuleFromCPointer(p unsafe.Pointer) *IOModule {
	ret := &IOModule{
		NewTraitIOModule(p),
		gobject.NewTraitTypeModule(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *IOModule) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type IOStream struct {
	*TraitIOStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewIOStreamFromCPointer(p unsafe.Pointer) *IOStream {
	ret := &IOStream{
		NewTraitIOStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *IOStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type InetAddress struct {
	*TraitInetAddress
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewInetAddressFromCPointer(p unsafe.Pointer) *InetAddress {
	ret := &InetAddress{
		NewTraitInetAddress(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *InetAddress) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type InetAddressMask struct {
	*TraitInetAddressMask
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewInetAddressMaskFromCPointer(p unsafe.Pointer) *InetAddressMask {
	ret := &InetAddressMask{
		NewTraitInetAddressMask(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *InetAddressMask) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type InetSocketAddress struct {
	*TraitInetSocketAddress
	*TraitSocketAddress
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewInetSocketAddressFromCPointer(p unsafe.Pointer) *InetSocketAddress {
	ret := &InetSocketAddress{
		NewTraitInetSocketAddress(p),
		NewTraitSocketAddress(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *InetSocketAddress) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type InputStream struct {
	*TraitInputStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewInputStreamFromCPointer(p unsafe.Pointer) *InputStream {
	ret := &InputStream{
		NewTraitInputStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *InputStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type MemoryInputStream struct {
	*TraitMemoryInputStream
	*TraitInputStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMemoryInputStreamFromCPointer(p unsafe.Pointer) *MemoryInputStream {
	ret := &MemoryInputStream{
		NewTraitMemoryInputStream(p),
		NewTraitInputStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *MemoryInputStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type MemoryOutputStream struct {
	*TraitMemoryOutputStream
	*TraitOutputStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMemoryOutputStreamFromCPointer(p unsafe.Pointer) *MemoryOutputStream {
	ret := &MemoryOutputStream{
		NewTraitMemoryOutputStream(p),
		NewTraitOutputStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *MemoryOutputStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Menu struct {
	*TraitMenu
	*TraitMenuModel
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMenuFromCPointer(p unsafe.Pointer) *Menu {
	ret := &Menu{
		NewTraitMenu(p),
		NewTraitMenuModel(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Menu) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type MenuAttributeIter struct {
	*TraitMenuAttributeIter
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMenuAttributeIterFromCPointer(p unsafe.Pointer) *MenuAttributeIter {
	ret := &MenuAttributeIter{
		NewTraitMenuAttributeIter(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *MenuAttributeIter) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type MenuItem struct {
	*TraitMenuItem
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMenuItemFromCPointer(p unsafe.Pointer) *MenuItem {
	ret := &MenuItem{
		NewTraitMenuItem(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *MenuItem) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type MenuLinkIter struct {
	*TraitMenuLinkIter
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMenuLinkIterFromCPointer(p unsafe.Pointer) *MenuLinkIter {
	ret := &MenuLinkIter{
		NewTraitMenuLinkIter(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *MenuLinkIter) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type MenuModel struct {
	*TraitMenuModel
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMenuModelFromCPointer(p unsafe.Pointer) *MenuModel {
	ret := &MenuModel{
		NewTraitMenuModel(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *MenuModel) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type MountOperation struct {
	*TraitMountOperation
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewMountOperationFromCPointer(p unsafe.Pointer) *MountOperation {
	ret := &MountOperation{
		NewTraitMountOperation(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *MountOperation) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type NativeVolumeMonitor struct {
	*TraitNativeVolumeMonitor
	*TraitVolumeMonitor
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewNativeVolumeMonitorFromCPointer(p unsafe.Pointer) *NativeVolumeMonitor {
	ret := &NativeVolumeMonitor{
		NewTraitNativeVolumeMonitor(p),
		NewTraitVolumeMonitor(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *NativeVolumeMonitor) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type NetworkAddress struct {
	*TraitNetworkAddress
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewNetworkAddressFromCPointer(p unsafe.Pointer) *NetworkAddress {
	ret := &NetworkAddress{
		NewTraitNetworkAddress(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *NetworkAddress) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type NetworkService struct {
	*TraitNetworkService
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewNetworkServiceFromCPointer(p unsafe.Pointer) *NetworkService {
	ret := &NetworkService{
		NewTraitNetworkService(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *NetworkService) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Notification struct {
	*TraitNotification
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewNotificationFromCPointer(p unsafe.Pointer) *Notification {
	ret := &Notification{
		NewTraitNotification(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Notification) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type OutputStream struct {
	*TraitOutputStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewOutputStreamFromCPointer(p unsafe.Pointer) *OutputStream {
	ret := &OutputStream{
		NewTraitOutputStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *OutputStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Permission struct {
	*TraitPermission
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPermissionFromCPointer(p unsafe.Pointer) *Permission {
	ret := &Permission{
		NewTraitPermission(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Permission) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type PropertyAction struct {
	*TraitPropertyAction
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPropertyActionFromCPointer(p unsafe.Pointer) *PropertyAction {
	ret := &PropertyAction{
		NewTraitPropertyAction(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *PropertyAction) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ProxyAddress struct {
	*TraitProxyAddress
	*TraitInetSocketAddress
	*TraitSocketAddress
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewProxyAddressFromCPointer(p unsafe.Pointer) *ProxyAddress {
	ret := &ProxyAddress{
		NewTraitProxyAddress(p),
		NewTraitInetSocketAddress(p),
		NewTraitSocketAddress(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ProxyAddress) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ProxyAddressEnumerator struct {
	*TraitProxyAddressEnumerator
	*TraitSocketAddressEnumerator
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewProxyAddressEnumeratorFromCPointer(p unsafe.Pointer) *ProxyAddressEnumerator {
	ret := &ProxyAddressEnumerator{
		NewTraitProxyAddressEnumerator(p),
		NewTraitSocketAddressEnumerator(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ProxyAddressEnumerator) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Resolver struct {
	*TraitResolver
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewResolverFromCPointer(p unsafe.Pointer) *Resolver {
	ret := &Resolver{
		NewTraitResolver(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Resolver) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Settings struct {
	*TraitSettings
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSettingsFromCPointer(p unsafe.Pointer) *Settings {
	ret := &Settings{
		NewTraitSettings(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Settings) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type SimpleAction struct {
	*TraitSimpleAction
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSimpleActionFromCPointer(p unsafe.Pointer) *SimpleAction {
	ret := &SimpleAction{
		NewTraitSimpleAction(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SimpleAction) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type SimpleActionGroup struct {
	*TraitSimpleActionGroup
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSimpleActionGroupFromCPointer(p unsafe.Pointer) *SimpleActionGroup {
	ret := &SimpleActionGroup{
		NewTraitSimpleActionGroup(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SimpleActionGroup) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type SimpleAsyncResult struct {
	*TraitSimpleAsyncResult
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSimpleAsyncResultFromCPointer(p unsafe.Pointer) *SimpleAsyncResult {
	ret := &SimpleAsyncResult{
		NewTraitSimpleAsyncResult(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SimpleAsyncResult) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type SimplePermission struct {
	*TraitSimplePermission
	*TraitPermission
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSimplePermissionFromCPointer(p unsafe.Pointer) *SimplePermission {
	ret := &SimplePermission{
		NewTraitSimplePermission(p),
		NewTraitPermission(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SimplePermission) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type SimpleProxyResolver struct {
	*TraitSimpleProxyResolver
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSimpleProxyResolverFromCPointer(p unsafe.Pointer) *SimpleProxyResolver {
	ret := &SimpleProxyResolver{
		NewTraitSimpleProxyResolver(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SimpleProxyResolver) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Socket struct {
	*TraitSocket
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSocketFromCPointer(p unsafe.Pointer) *Socket {
	ret := &Socket{
		NewTraitSocket(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Socket) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type SocketAddress struct {
	*TraitSocketAddress
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSocketAddressFromCPointer(p unsafe.Pointer) *SocketAddress {
	ret := &SocketAddress{
		NewTraitSocketAddress(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SocketAddress) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type SocketAddressEnumerator struct {
	*TraitSocketAddressEnumerator
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSocketAddressEnumeratorFromCPointer(p unsafe.Pointer) *SocketAddressEnumerator {
	ret := &SocketAddressEnumerator{
		NewTraitSocketAddressEnumerator(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SocketAddressEnumerator) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type SocketClient struct {
	*TraitSocketClient
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSocketClientFromCPointer(p unsafe.Pointer) *SocketClient {
	ret := &SocketClient{
		NewTraitSocketClient(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SocketClient) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type SocketConnection struct {
	*TraitSocketConnection
	*TraitIOStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSocketConnectionFromCPointer(p unsafe.Pointer) *SocketConnection {
	ret := &SocketConnection{
		NewTraitSocketConnection(p),
		NewTraitIOStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SocketConnection) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type SocketControlMessage struct {
	*TraitSocketControlMessage
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSocketControlMessageFromCPointer(p unsafe.Pointer) *SocketControlMessage {
	ret := &SocketControlMessage{
		NewTraitSocketControlMessage(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SocketControlMessage) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type SocketListener struct {
	*TraitSocketListener
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSocketListenerFromCPointer(p unsafe.Pointer) *SocketListener {
	ret := &SocketListener{
		NewTraitSocketListener(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SocketListener) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type SocketService struct {
	*TraitSocketService
	*TraitSocketListener
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSocketServiceFromCPointer(p unsafe.Pointer) *SocketService {
	ret := &SocketService{
		NewTraitSocketService(p),
		NewTraitSocketListener(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SocketService) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Subprocess struct {
	*TraitSubprocess
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSubprocessFromCPointer(p unsafe.Pointer) *Subprocess {
	ret := &Subprocess{
		NewTraitSubprocess(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Subprocess) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type SubprocessLauncher struct {
	*TraitSubprocessLauncher
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSubprocessLauncherFromCPointer(p unsafe.Pointer) *SubprocessLauncher {
	ret := &SubprocessLauncher{
		NewTraitSubprocessLauncher(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SubprocessLauncher) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Task struct {
	*TraitTask
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTaskFromCPointer(p unsafe.Pointer) *Task {
	ret := &Task{
		NewTraitTask(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Task) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TcpConnection struct {
	*TraitTcpConnection
	*TraitSocketConnection
	*TraitIOStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTcpConnectionFromCPointer(p unsafe.Pointer) *TcpConnection {
	ret := &TcpConnection{
		NewTraitTcpConnection(p),
		NewTraitSocketConnection(p),
		NewTraitIOStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TcpConnection) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TcpWrapperConnection struct {
	*TraitTcpWrapperConnection
	*TraitTcpConnection
	*TraitSocketConnection
	*TraitIOStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTcpWrapperConnectionFromCPointer(p unsafe.Pointer) *TcpWrapperConnection {
	ret := &TcpWrapperConnection{
		NewTraitTcpWrapperConnection(p),
		NewTraitTcpConnection(p),
		NewTraitSocketConnection(p),
		NewTraitIOStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TcpWrapperConnection) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TestDBus struct {
	*TraitTestDBus
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTestDBusFromCPointer(p unsafe.Pointer) *TestDBus {
	ret := &TestDBus{
		NewTraitTestDBus(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TestDBus) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ThemedIcon struct {
	*TraitThemedIcon
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewThemedIconFromCPointer(p unsafe.Pointer) *ThemedIcon {
	ret := &ThemedIcon{
		NewTraitThemedIcon(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ThemedIcon) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ThreadedSocketService struct {
	*TraitThreadedSocketService
	*TraitSocketService
	*TraitSocketListener
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewThreadedSocketServiceFromCPointer(p unsafe.Pointer) *ThreadedSocketService {
	ret := &ThreadedSocketService{
		NewTraitThreadedSocketService(p),
		NewTraitSocketService(p),
		NewTraitSocketListener(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ThreadedSocketService) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TlsCertificate struct {
	*TraitTlsCertificate
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTlsCertificateFromCPointer(p unsafe.Pointer) *TlsCertificate {
	ret := &TlsCertificate{
		NewTraitTlsCertificate(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TlsCertificate) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TlsConnection struct {
	*TraitTlsConnection
	*TraitIOStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTlsConnectionFromCPointer(p unsafe.Pointer) *TlsConnection {
	ret := &TlsConnection{
		NewTraitTlsConnection(p),
		NewTraitIOStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TlsConnection) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TlsDatabase struct {
	*TraitTlsDatabase
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTlsDatabaseFromCPointer(p unsafe.Pointer) *TlsDatabase {
	ret := &TlsDatabase{
		NewTraitTlsDatabase(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TlsDatabase) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TlsInteraction struct {
	*TraitTlsInteraction
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTlsInteractionFromCPointer(p unsafe.Pointer) *TlsInteraction {
	ret := &TlsInteraction{
		NewTraitTlsInteraction(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TlsInteraction) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type TlsPassword struct {
	*TraitTlsPassword
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewTlsPasswordFromCPointer(p unsafe.Pointer) *TlsPassword {
	ret := &TlsPassword{
		NewTraitTlsPassword(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *TlsPassword) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type UnixConnection struct {
	*TraitUnixConnection
	*TraitSocketConnection
	*TraitIOStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewUnixConnectionFromCPointer(p unsafe.Pointer) *UnixConnection {
	ret := &UnixConnection{
		NewTraitUnixConnection(p),
		NewTraitSocketConnection(p),
		NewTraitIOStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *UnixConnection) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type UnixCredentialsMessage struct {
	*TraitUnixCredentialsMessage
	*TraitSocketControlMessage
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewUnixCredentialsMessageFromCPointer(p unsafe.Pointer) *UnixCredentialsMessage {
	ret := &UnixCredentialsMessage{
		NewTraitUnixCredentialsMessage(p),
		NewTraitSocketControlMessage(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *UnixCredentialsMessage) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type UnixFDList struct {
	*TraitUnixFDList
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewUnixFDListFromCPointer(p unsafe.Pointer) *UnixFDList {
	ret := &UnixFDList{
		NewTraitUnixFDList(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *UnixFDList) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type UnixFDMessage struct {
	*TraitUnixFDMessage
	*TraitSocketControlMessage
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewUnixFDMessageFromCPointer(p unsafe.Pointer) *UnixFDMessage {
	ret := &UnixFDMessage{
		NewTraitUnixFDMessage(p),
		NewTraitSocketControlMessage(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *UnixFDMessage) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type UnixInputStream struct {
	*TraitUnixInputStream
	*TraitInputStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewUnixInputStreamFromCPointer(p unsafe.Pointer) *UnixInputStream {
	ret := &UnixInputStream{
		NewTraitUnixInputStream(p),
		NewTraitInputStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *UnixInputStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type UnixMountMonitor struct {
	*TraitUnixMountMonitor
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewUnixMountMonitorFromCPointer(p unsafe.Pointer) *UnixMountMonitor {
	ret := &UnixMountMonitor{
		NewTraitUnixMountMonitor(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *UnixMountMonitor) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type UnixOutputStream struct {
	*TraitUnixOutputStream
	*TraitOutputStream
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewUnixOutputStreamFromCPointer(p unsafe.Pointer) *UnixOutputStream {
	ret := &UnixOutputStream{
		NewTraitUnixOutputStream(p),
		NewTraitOutputStream(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *UnixOutputStream) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type UnixSocketAddress struct {
	*TraitUnixSocketAddress
	*TraitSocketAddress
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewUnixSocketAddressFromCPointer(p unsafe.Pointer) *UnixSocketAddress {
	ret := &UnixSocketAddress{
		NewTraitUnixSocketAddress(p),
		NewTraitSocketAddress(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *UnixSocketAddress) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Vfs struct {
	*TraitVfs
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewVfsFromCPointer(p unsafe.Pointer) *Vfs {
	ret := &Vfs{
		NewTraitVfs(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Vfs) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type VolumeMonitor struct {
	*TraitVolumeMonitor
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewVolumeMonitorFromCPointer(p unsafe.Pointer) *VolumeMonitor {
	ret := &VolumeMonitor{
		NewTraitVolumeMonitor(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *VolumeMonitor) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ZlibCompressor struct {
	*TraitZlibCompressor
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewZlibCompressorFromCPointer(p unsafe.Pointer) *ZlibCompressor {
	ret := &ZlibCompressor{
		NewTraitZlibCompressor(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ZlibCompressor) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ZlibDecompressor struct {
	*TraitZlibDecompressor
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewZlibDecompressorFromCPointer(p unsafe.Pointer) *ZlibDecompressor {
	ret := &ZlibDecompressor{
		NewTraitZlibDecompressor(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ZlibDecompressor) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}
