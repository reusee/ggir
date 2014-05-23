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
import "github.com/reusee/ggir/gobject"
import "github.com/reusee/ggir/glib"
import "unsafe"
import "reflect"

func init() {
	_ = unsafe.Pointer(nil)
	_ = reflect.ValueOf(nil)
}

var _ = gobject.UnusedFix_
var _ = glib.UnusedFix_

func (self *TraitAppInfoMonitor) OnChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "changed", func(_self reflect.Value) {
	})
}

func (self *TraitAppLaunchContext) OnLaunchFailed(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "launch-failed", func(_self reflect.Value, object reflect.Value) {
	})
}

func (self *TraitApplication) OnActivate(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate", func(_self reflect.Value) {
	})
}

func (self *TraitApplication) OnCommandLine(f func(*ApplicationCommandLine)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "command-line", func(_self reflect.Value, command_line reflect.Value) {
	})
}

func (self *TraitApplication) OnHandleLocalOptions(f func(*glib.VariantDict)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "handle-local-options", func(_self reflect.Value, options reflect.Value) {
	})
}

func (self *TraitApplication) OnShutdown(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "shutdown", func(_self reflect.Value) {
	})
}

func (self *TraitApplication) OnStartup(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "startup", func(_self reflect.Value) {
	})
}

func (self *TraitCancellable) OnCancelled(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "cancelled", func(_self reflect.Value) {
	})
}

func (self *TraitDBusAuthObserver) OnAllowMechanism(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "allow-mechanism", func(_self reflect.Value, mechanism reflect.Value) {
	})
}

func (self *TraitDBusAuthObserver) OnAuthorizeAuthenticatedPeer(f func(*IOStream, *Credentials)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "authorize-authenticated-peer", func(_self reflect.Value, stream reflect.Value, credentials reflect.Value) {
	})
}

func (self *TraitDBusConnection) OnClosed(f func(bool, *glib.Error)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "closed", func(_self reflect.Value, remote_peer_vanished reflect.Value, error_ reflect.Value) {
	})
}

func (self *TraitDBusInterfaceSkeleton) OnGAuthorizeMethod(f func(*DBusMethodInvocation)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "g-authorize-method", func(_self reflect.Value, invocation reflect.Value) {
	})
}

func (self *TraitDBusObjectManagerClient) OnInterfaceProxyPropertiesChanged(f func(*DBusObjectProxy, *DBusProxy, *glib.Variant, []string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "interface-proxy-properties-changed", func(_self reflect.Value, object_proxy reflect.Value, interface_proxy reflect.Value, changed_properties reflect.Value, invalidated_properties reflect.Value) {
	})
}

func (self *TraitDBusObjectManagerClient) OnInterfaceProxySignal(f func(*DBusObjectProxy, *DBusProxy, string, string, *glib.Variant)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "interface-proxy-signal", func(_self reflect.Value, object_proxy reflect.Value, interface_proxy reflect.Value, sender_name reflect.Value, signal_name reflect.Value, parameters reflect.Value) {
	})
}

func (self *TraitDBusObjectSkeleton) OnAuthorizeMethod(f func(*DBusInterfaceSkeleton, *DBusMethodInvocation)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "authorize-method", func(_self reflect.Value, interface_ reflect.Value, invocation reflect.Value) {
	})
}

func (self *TraitDBusProxy) OnGPropertiesChanged(f func(*glib.Variant, []string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "g-properties-changed", func(_self reflect.Value, changed_properties reflect.Value, invalidated_properties reflect.Value) {
	})
}

func (self *TraitDBusProxy) OnGSignal(f func(string, string, *glib.Variant)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "g-signal", func(_self reflect.Value, sender_name reflect.Value, signal_name reflect.Value, parameters reflect.Value) {
	})
}

func (self *TraitDBusServer) OnNewConnection(f func(*DBusConnection)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "new-connection", func(_self reflect.Value, connection reflect.Value) {
	})
}

func (self *TraitFilenameCompleter) OnGotCompletionData(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "got-completion-data", func(_self reflect.Value) {
	})
}

func (self *TraitMenuModel) OnItemsChanged(f func(int, int, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "items-changed", func(_self reflect.Value, position reflect.Value, removed reflect.Value, added reflect.Value) {
	})
}

func (self *TraitMountOperation) OnAborted(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "aborted", func(_self reflect.Value) {
	})
}

func (self *TraitMountOperation) OnAskPassword(f func(string, string, string, C.GAskPasswordFlags)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "ask-password", func(_self reflect.Value, message reflect.Value, default_user reflect.Value, default_domain reflect.Value, flags reflect.Value) {
	})
}

func (self *TraitMountOperation) OnAskQuestion(f func(string, []string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "ask-question", func(_self reflect.Value, message reflect.Value, choices reflect.Value) {
	})
}

func (self *TraitMountOperation) OnReply(f func(C.GMountOperationResult)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "reply", func(_self reflect.Value, result reflect.Value) {
	})
}

func (self *TraitMountOperation) OnShowProcesses(f func(string, []int, []string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "show-processes", func(_self reflect.Value, message reflect.Value, processes reflect.Value, choices reflect.Value) {
	})
}

func (self *TraitMountOperation) OnShowUnmountProgress(f func(string, int64, int64)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "show-unmount-progress", func(_self reflect.Value, message reflect.Value, time_left reflect.Value, bytes_left reflect.Value) {
	})
}

func (self *TraitResolver) OnReload(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "reload", func(_self reflect.Value) {
	})
}

func (self *TraitSettings) OnChangeEvent(f func([]uint32, int)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "change-event", func(_self reflect.Value, keys reflect.Value, n_keys reflect.Value) {
	})
}

func (self *TraitSettings) OnChanged(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "changed", func(_self reflect.Value, key reflect.Value) {
	})
}

func (self *TraitSettings) OnWritableChangeEvent(f func(uint)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "writable-change-event", func(_self reflect.Value, key reflect.Value) {
	})
}

func (self *TraitSettings) OnWritableChanged(f func(string)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "writable-changed", func(_self reflect.Value, key reflect.Value) {
	})
}

func (self *TraitSimpleAction) OnActivate(f func(*glib.Variant)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "activate", func(_self reflect.Value, parameter reflect.Value) {
	})
}

func (self *TraitSimpleAction) OnChangeState(f func(*glib.Variant)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "change-state", func(_self reflect.Value, value reflect.Value) {
	})
}

func (self *TraitSocketService) OnIncoming(f func(*SocketConnection, *gobject.Object)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "incoming", func(_self reflect.Value, connection reflect.Value, source_object reflect.Value) {
	})
}

func (self *TraitThreadedSocketService) OnRun(f func(*SocketConnection, *gobject.Object)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "run", func(_self reflect.Value, connection reflect.Value, source_object reflect.Value) {
	})
}

func (self *TraitTlsConnection) OnAcceptCertificate(f func(*TlsCertificate, C.GTlsCertificateFlags)) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "accept-certificate", func(_self reflect.Value, peer_cert reflect.Value, errors reflect.Value) {
	})
}

func (self *TraitUnixMountMonitor) OnMountpointsChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "mountpoints-changed", func(_self reflect.Value) {
	})
}

func (self *TraitUnixMountMonitor) OnMountsChanged(f func()) {
	gobject.Connect(unsafe.Pointer(self.CPointer), "mounts-changed", func(_self reflect.Value) {
	})
}
