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

type ActionEntry C.GActionEntry
type DBusAnnotationInfo C.GDBusAnnotationInfo
type DBusArgInfo C.GDBusArgInfo
type DBusErrorEntry C.GDBusErrorEntry
type DBusInterfaceInfo C.GDBusInterfaceInfo
type DBusInterfaceVTable C.GDBusInterfaceVTable
type DBusMethodInfo C.GDBusMethodInfo
type DBusNodeInfo C.GDBusNodeInfo
type DBusPropertyInfo C.GDBusPropertyInfo
type DBusSignalInfo C.GDBusSignalInfo
type DBusSubtreeVTable C.GDBusSubtreeVTable
type FileAttributeInfo C.GFileAttributeInfo
type FileAttributeInfoList C.GFileAttributeInfoList
type FileAttributeMatcher C.GFileAttributeMatcher
type IOExtension C.GIOExtension
type IOExtensionPoint C.GIOExtensionPoint
type IOModuleScope C.GIOModuleScope
type IOSchedulerJob C.GIOSchedulerJob
type IOStreamAdapter C.GIOStreamAdapter
type InputVector C.GInputVector
type OutputVector C.GOutputVector
type Resource C.GResource
type SettingsBackend C.GSettingsBackend
type SettingsSchema C.GSettingsSchema
type SettingsSchemaKey C.GSettingsSchemaKey
type SettingsSchemaSource C.GSettingsSchemaSource
type SrvTarget C.GSrvTarget
type StaticResource C.GStaticResource
type UnixMountEntry C.GUnixMountEntry
type UnixMountPoint C.GUnixMountPoint
