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

var _ = gobject.UnusedFix_
var _ = glib.UnusedFix_

type TraitAppInfoMonitor struct{ CPointer *C.GAppInfoMonitor }
type IsAppInfoMonitor interface {
	GetAppInfoMonitorPointer() *C.GAppInfoMonitor
}

func (self *TraitAppInfoMonitor) GetAppInfoMonitorPointer() *C.GAppInfoMonitor {
	return self.CPointer
}
func NewTraitAppInfoMonitor(p unsafe.Pointer) *TraitAppInfoMonitor {
	return &TraitAppInfoMonitor{(*C.GAppInfoMonitor)(p)}
}

type TraitAppLaunchContext struct{ CPointer *C.GAppLaunchContext }
type IsAppLaunchContext interface {
	GetAppLaunchContextPointer() *C.GAppLaunchContext
}

func (self *TraitAppLaunchContext) GetAppLaunchContextPointer() *C.GAppLaunchContext {
	return self.CPointer
}
func NewTraitAppLaunchContext(p unsafe.Pointer) *TraitAppLaunchContext {
	return &TraitAppLaunchContext{(*C.GAppLaunchContext)(p)}
}

/*
Gets the display string for the @context. This is used to ensure new
applications are started on the same display as the launching
application, by setting the `DISPLAY` environment variable.
*/
func (self *TraitAppLaunchContext) GetDisplay(info *C.GAppInfo, files *C.GList) (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_app_launch_context_get_display(self.CPointer, info, files)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Gets the complete environment variable list to be passed to
the child process when @context is used to launch an application.
This is a %NULL-terminated array of strings, where each string has
the form `KEY=VALUE`.
*/
func (self *TraitAppLaunchContext) GetEnvironment() (return__ **C.char) {
	return__ = C.g_app_launch_context_get_environment(self.CPointer)
	return
}

/*
Initiates startup notification for the application and returns the
`DESKTOP_STARTUP_ID` for the launched operation, if supported.

Startup notification IDs are defined in the
[FreeDesktop.Org Startup Notifications standard](http://standards.freedesktop.org/startup-notification-spec/startup-notification-latest.txt").
*/
func (self *TraitAppLaunchContext) GetStartupNotifyId(info *C.GAppInfo, files *C.GList) (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_app_launch_context_get_startup_notify_id(self.CPointer, info, files)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Called when an application has failed to launch, so that it can cancel
the application startup notification started in g_app_launch_context_get_startup_notify_id().
*/
func (self *TraitAppLaunchContext) LaunchFailed(startup_notify_id string) {
	__cgo__startup_notify_id := C.CString(startup_notify_id)
	C.g_app_launch_context_launch_failed(self.CPointer, __cgo__startup_notify_id)
	C.free(unsafe.Pointer(__cgo__startup_notify_id))
	return
}

/*
Arranges for @variable to be set to @value in the child's
environment when @context is used to launch an application.
*/
func (self *TraitAppLaunchContext) Setenv(variable string, value string) {
	__cgo__variable := C.CString(variable)
	__cgo__value := C.CString(value)
	C.g_app_launch_context_setenv(self.CPointer, __cgo__variable, __cgo__value)
	C.free(unsafe.Pointer(__cgo__variable))
	C.free(unsafe.Pointer(__cgo__value))
	return
}

/*
Arranges for @variable to be unset in the child's environment
when @context is used to launch an application.
*/
func (self *TraitAppLaunchContext) Unsetenv(variable string) {
	__cgo__variable := C.CString(variable)
	C.g_app_launch_context_unsetenv(self.CPointer, __cgo__variable)
	C.free(unsafe.Pointer(__cgo__variable))
	return
}

type TraitApplication struct{ CPointer *C.GApplication }
type IsApplication interface {
	GetApplicationPointer() *C.GApplication
}

func (self *TraitApplication) GetApplicationPointer() *C.GApplication {
	return self.CPointer
}
func NewTraitApplication(p unsafe.Pointer) *TraitApplication {
	return &TraitApplication{(*C.GApplication)(p)}
}

/*
Activates the application.

In essence, this results in the #GApplication::activate signal being
emitted in the primary instance.

The application must be registered before calling this function.
*/
func (self *TraitApplication) Activate() {
	C.g_application_activate(self.CPointer)
	return
}

/*
Adds main option entries to be handled by @application.

This function is comparable to g_option_context_add_main_entries().

After the commandline arguments are parsed, the
#GApplication::handle-local-options signal will be emitted.  At this
point, the application can inspect the values pointed to by @arg_data
in the given #GOptionEntrys.

Unlike #GOptionContext, #GApplication supports giving a %NULL
@arg_data for a non-callback #GOptionEntry.  This results in the
argument in question being packed into a #GVariantDict which is also
passed to #GApplication::handle-local-options, where it can be
inspected and modified.  If %G_APPLICATION_HANDLES_COMMAND_LINE is
set, then the resulting dictionary is sent to the primary instance,
where g_application_command_line_get_options_dict() will return it.
This "packing" is done according to the type of the argument --
booleans for normal flags, strings for strings, bytestrings for
filenames, etc.  The packing only occurs if the flag is given (ie: we
do not pack a "false" #GVariant in the case that a flag is missing).

In general, it is recommended that all commandline arguments are
parsed locally.  The options dictionary should then be used to
transmit the result of the parsing to the primary instance, where
g_variant_dict_lookup() can be used.  For local options, it is
possible to either use @arg_data in the usual way, or to consult (and
potentially remove) the option from the options dictionary.

This function is new in GLib 2.40.  Before then, the only real choice
was to send all of the commandline arguments (options and all) to the
primary instance for handling.  #GApplication ignored them completely
on the local side.  Calling this function "opts in" to the new
behaviour, and in particular, means that unrecognised options will be
treated as errors.  Unrecognised options have never been ignored when
%G_APPLICATION_HANDLES_COMMAND_LINE is unset.

If #GApplication::handle-local-options needs to see the list of
filenames, then the use of %G_OPTION_REMAINING is recommended.  If
@arg_data is %NULL then %G_OPTION_REMAINING can be used as a key into
the options dictionary.  If you do use %G_OPTION_REMAINING then you
need to handle these arguments for yourself because once they are
consumed, they will no longer be visible to the default handling
(which treats them as filenames to be opened).
*/
func (self *TraitApplication) AddMainOptionEntries(entries *C.GOptionEntry) {
	C.g_application_add_main_option_entries(self.CPointer, entries)
	return
}

/*
Adds a #GOptionGroup to the commandline handling of @application.

This function is comparable to g_option_context_add_group().

Unlike g_application_add_main_option_entries(), this function does
not deal with %NULL @arg_data and never transmits options to the
primary instance.

The reason for that is because, by the time the options arrive at the
primary instance, it is typically too late to do anything with them.
Taking the GTK option group as an example: GTK will already have been
initialised by the time the #GApplication::command-line handler runs.
In the case that this is not the first-running instance of the
application, the existing instance may already have been running for
a very long time.

This means that the options from #GOptionGroup are only really usable
in the case that the instance of the application being run is the
first instance.  Passing options like `--display=` or `--gdk-debug=`
on future runs will have no effect on the existing primary instance.

Calling this function will cause the options in the supplied option
group to be parsed, but it does not cause you to be "opted in" to the
new functionality whereby unrecognised options are rejected even if
%G_APPLICATION_HANDLES_COMMAND_LINE was given.
*/
func (self *TraitApplication) AddOptionGroup(group *C.GOptionGroup) {
	C.g_application_add_option_group(self.CPointer, group)
	return
}

/*
Gets the unique identifier for @application.
*/
func (self *TraitApplication) GetApplicationId() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_application_get_application_id(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the #GDBusConnection being used by the application, or %NULL.

If #GApplication is using its D-Bus backend then this function will
return the #GDBusConnection being used for uniqueness and
communication with the desktop environment and other instances of the
application.

If #GApplication is not using D-Bus then this function will return
%NULL.  This includes the situation where the D-Bus backend would
normally be in use but we were unable to connect to the bus.

This function must not be called before the application has been
registered.  See g_application_get_is_registered().
*/
func (self *TraitApplication) GetDbusConnection() (return__ *DBusConnection) {
	var __cgo__return__ *C.GDBusConnection
	__cgo__return__ = C.g_application_get_dbus_connection(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDBusConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the D-Bus object path being used by the application, or %NULL.

If #GApplication is using its D-Bus backend then this function will
return the D-Bus object path that #GApplication is using.  If the
application is the primary instance then there is an object published
at this path.  If the application is not the primary instance then
the result of this function is undefined.

If #GApplication is not using D-Bus then this function will return
%NULL.  This includes the situation where the D-Bus backend would
normally be in use but we were unable to connect to the bus.

This function must not be called before the application has been
registered.  See g_application_get_is_registered().
*/
func (self *TraitApplication) GetDbusObjectPath() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_application_get_dbus_object_path(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the flags for @application.

See #GApplicationFlags.
*/
func (self *TraitApplication) GetFlags() (return__ C.GApplicationFlags) {
	return__ = C.g_application_get_flags(self.CPointer)
	return
}

/*
Gets the current inactivity timeout for the application.

This is the amount of time (in milliseconds) after the last call to
g_application_release() before the application stops running.
*/
func (self *TraitApplication) GetInactivityTimeout() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_application_get_inactivity_timeout(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Checks if @application is registered.

An application is registered if g_application_register() has been
successfully called.
*/
func (self *TraitApplication) GetIsRegistered() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_application_get_is_registered(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if @application is remote.

If @application is remote then it means that another instance of
application already exists (the 'primary' instance).  Calls to
perform actions on @application will result in the actions being
performed by the primary instance.

The value of this property cannot be accessed before
g_application_register() has been called.  See
g_application_get_is_registered().
*/
func (self *TraitApplication) GetIsRemote() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_application_get_is_remote(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Increases the use count of @application.

Use this function to indicate that the application has a reason to
continue to run.  For example, g_application_hold() is called by GTK+
when a toplevel window is on the screen.

To cancel the hold, call g_application_release().
*/
func (self *TraitApplication) Hold() {
	C.g_application_hold(self.CPointer)
	return
}

/*
Increases the busy count of @application.

Use this function to indicate that the application is busy, for instance
while a long running operation is pending.

The busy state will be exposed to other processes, so a session shell will
use that information to indicate the state to the user (e.g. with a
spinner).

To cancel the busy indication, use g_application_unmark_busy().
*/
func (self *TraitApplication) MarkBusy() {
	C.g_application_mark_busy(self.CPointer)
	return
}

/*
Opens the given files.

In essence, this results in the #GApplication::open signal being emitted
in the primary instance.

@n_files must be greater than zero.

@hint is simply passed through to the ::open signal.  It is
intended to be used by applications that have multiple modes for
opening files (eg: "view" vs "edit", etc).  Unless you have a need
for this functionality, you should use "".

The application must be registered before calling this function
and it must have the %G_APPLICATION_HANDLES_OPEN flag set.
*/
func (self *TraitApplication) Open(files **C.GFile, n_files int, hint string) {
	__cgo__hint := (*C.gchar)(unsafe.Pointer(C.CString(hint)))
	C.g_application_open(self.CPointer, files, C.gint(n_files), __cgo__hint)
	C.free(unsafe.Pointer(__cgo__hint))
	return
}

/*
Immediately quits the application.

Upon return to the mainloop, g_application_run() will return,
calling only the 'shutdown' function before doing so.

The hold count is ignored.

The result of calling g_application_run() again after it returns is
unspecified.
*/
func (self *TraitApplication) Quit() {
	C.g_application_quit(self.CPointer)
	return
}

/*
Attempts registration of the application.

This is the point at which the application discovers if it is the
primary instance or merely acting as a remote for an already-existing
primary instance.  This is implemented by attempting to acquire the
application identifier as a unique bus name on the session bus using
GDBus.

If there is no application ID or if %G_APPLICATION_NON_UNIQUE was
given, then this process will always become the primary instance.

Due to the internal architecture of GDBus, method calls can be
dispatched at any time (even if a main loop is not running).  For
this reason, you must ensure that any object paths that you wish to
register are registered before calling this function.

If the application has already been registered then %TRUE is
returned with no work performed.

The #GApplication::startup signal is emitted if registration succeeds
and @application is the primary instance (including the non-unique
case).

In the event of an error (such as @cancellable being cancelled, or a
failure to connect to the session bus), %FALSE is returned and @error
is set appropriately.

Note: the return value of this function is not an indicator that this
instance is or is not the primary instance of the application.  See
g_application_get_is_remote() for that.
*/
func (self *TraitApplication) Register(cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_application_register(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Decrease the use count of @application.

When the use count reaches zero, the application will stop running.

Never call this function except to cancel the effect of a previous
call to g_application_hold().
*/
func (self *TraitApplication) Release() {
	C.g_application_release(self.CPointer)
	return
}

/*
Runs the application.

This function is intended to be run from main() and its return value
is intended to be returned by main(). Although you are expected to pass
the @argc, @argv parameters from main() to this function, it is possible
to pass %NULL if @argv is not available or commandline handling is not
required.  Note that on Windows, @argc and @argv are ignored, and
g_win32_get_command_line() is called internally (for proper support
of Unicode commandline arguments).

#GApplication will attempt to parse the commandline arguments.  You
can add commandline flags to the list of recognised options by way of
g_application_add_main_option_entries().  After this, the
#GApplication::handle-local-options signal is emitted, from which the
application can inspect the values of its #GOptionEntrys.

#GApplication::handle-local-options is a good place to handle options
such as `--version`, where an immediate reply from the local process is
desired (instead of communicating with an already-running instance).
A #GApplication::handle-local-options handler can stop further processing
by returning a non-negative value, which then becomes the exit status of
the process.

What happens next depends on the flags: if
%G_APPLICATION_HANDLES_COMMAND_LINE was specified then the remaining
commandline arguments are sent to the primary instance, where a
#GApplication::command-line signal is emitted.  Otherwise, the
remaining commandline arguments are assumed to be a list of files.
If there are no files listed, the application is activated via the
#GApplication::activate signal.  If there are one or more files, and
%G_APPLICATION_HANDLES_OPEN was specified then the files are opened
via the #GApplication::open signal.

If you are interested in doing more complicated local handling of the
commandline then you should implement your own #GApplication subclass
and override local_command_line(). In this case, you most likely want
to return %TRUE from your local_command_line() implementation to
suppress the default handling. See
[gapplication-example-cmdline2.c][gapplication-example-cmdline2]
for an example.

If, after the above is done, the use count of the application is zero
then the exit status is returned immediately.  If the use count is
non-zero then the default main context is iterated until the use count
falls to zero, at which point 0 is returned.

If the %G_APPLICATION_IS_SERVICE flag is set, then the service will
run for as much as 10 seconds with a use count of zero while waiting
for the message that caused the activation to arrive.  After that,
if the use count falls to zero the application will exit immediately,
except in the case that g_application_set_inactivity_timeout() is in
use.

This function sets the prgname (g_set_prgname()), if not already set,
to the basename of argv[0].  Since 2.38, if %G_APPLICATION_IS_SERVICE
is specified, the prgname is set to the application ID.  The main
impact of this is is that the wmclass of windows created by Gtk+ will
be set accordingly, which helps the window manager determine which
application is showing the window.

Since 2.40, applications that are not explicitly flagged as services
or launchers (ie: neither %G_APPLICATION_IS_SERVICE or
%G_APPLICATION_IS_LAUNCHER are given as flags) will check (from the
default handler for local_command_line) if "--gapplication-service"
was given in the command line.  If this flag is present then normal
commandline processing is interrupted and the
%G_APPLICATION_IS_SERVICE flag is set.  This provides a "compromise"
solution whereby running an application directly from the commandline
will invoke it in the normal way (which can be useful for debugging)
while still allowing applications to be D-Bus activated in service
mode.  The D-Bus service file should invoke the executable with
"--gapplication-service" as the sole commandline argument.  This
approach is suitable for use by most graphical applications but
should not be used from applications like editors that need precise
control over when processes invoked via the commandline will exit and
what their exit status will be.
*/
func (self *TraitApplication) Run(argc int, argv []string) (return__ int) {
	__header__argv := (*reflect.SliceHeader)(unsafe.Pointer(&argv))
	var __cgo__return__ C.int
	__cgo__return__ = C.g_application_run(self.CPointer, C.int(argc), (**C.char)(unsafe.Pointer(__header__argv.Data)))
	return__ = int(__cgo__return__)
	return
}

/*
Sends a notification on behalf of @application to the desktop shell.
There is no guarantee that the notification is displayed immediately,
or even at all.

Notifications may persist after the application exits. It will be
D-Bus-activated when the notification or one of its actions is
activated.

Modifying @notification after this call has no effect. However, the
object can be reused for a later call to this function.

@id may be any string that uniquely identifies the event for the
application. It does not need to be in any special format. For
example, "new-message" might be appropriate for a notification about
new messages.

If a previous notification was sent with the same @id, it will be
replaced with @notification and shown again as if it was a new
notification. This works even for notifications sent from a previous
execution of the application, as long as @id is the same string.

@id may be %NULL, but it is impossible to replace or withdraw
notifications without an id.

If @notification is no longer relevant, it can be withdrawn with
g_application_withdraw_notification().
*/
func (self *TraitApplication) SendNotification(id string, notification IsNotification) {
	__cgo__id := (*C.gchar)(unsafe.Pointer(C.CString(id)))
	C.g_application_send_notification(self.CPointer, __cgo__id, notification.GetNotificationPointer())
	C.free(unsafe.Pointer(__cgo__id))
	return
}

// g_application_set_action_group is not generated due to deprecation attr

/*
Sets the unique identifier for @application.

The application id can only be modified if @application has not yet
been registered.

If non-%NULL, the application id must be valid.  See
g_application_id_is_valid().
*/
func (self *TraitApplication) SetApplicationId(application_id string) {
	__cgo__application_id := (*C.gchar)(unsafe.Pointer(C.CString(application_id)))
	C.g_application_set_application_id(self.CPointer, __cgo__application_id)
	C.free(unsafe.Pointer(__cgo__application_id))
	return
}

/*
Sets or unsets the default application for the process, as returned
by g_application_get_default().

This function does not take its own reference on @application.  If
@application is destroyed then the default application will revert
back to %NULL.
*/
func (self *TraitApplication) SetDefault() {
	C.g_application_set_default(self.CPointer)
	return
}

/*
Sets the flags for @application.

The flags can only be modified if @application has not yet been
registered.

See #GApplicationFlags.
*/
func (self *TraitApplication) SetFlags(flags C.GApplicationFlags) {
	C.g_application_set_flags(self.CPointer, flags)
	return
}

/*
Sets the current inactivity timeout for the application.

This is the amount of time (in milliseconds) after the last call to
g_application_release() before the application stops running.

This call has no side effects of its own.  The value set here is only
used for next time g_application_release() drops the use count to
zero.  Any timeouts currently in progress are not impacted.
*/
func (self *TraitApplication) SetInactivityTimeout(inactivity_timeout uint) {
	C.g_application_set_inactivity_timeout(self.CPointer, C.guint(inactivity_timeout))
	return
}

/*
Decreases the busy count of @application.

When the busy count reaches zero, the new state will be propagated
to other processes.

This function must only be called to cancel the effect of a previous
call to g_application_mark_busy().
*/
func (self *TraitApplication) UnmarkBusy() {
	C.g_application_unmark_busy(self.CPointer)
	return
}

/*
Withdraws a notification that was sent with
g_application_send_notification().

This call does nothing if a notification with @id doesn't exist or
the notification was never sent.

This function works even for notifications sent in previous
executions of this application, as long @id is the same as it was for
the sent notification.

Note that notifications are dismissed when the user clicks on one
of the buttons in a notification or triggers its default action, so
there is no need to explicitly withdraw the notification in that case.
*/
func (self *TraitApplication) WithdrawNotification(id string) {
	__cgo__id := (*C.gchar)(unsafe.Pointer(C.CString(id)))
	C.g_application_withdraw_notification(self.CPointer, __cgo__id)
	C.free(unsafe.Pointer(__cgo__id))
	return
}

type TraitApplicationCommandLine struct{ CPointer *C.GApplicationCommandLine }
type IsApplicationCommandLine interface {
	GetApplicationCommandLinePointer() *C.GApplicationCommandLine
}

func (self *TraitApplicationCommandLine) GetApplicationCommandLinePointer() *C.GApplicationCommandLine {
	return self.CPointer
}
func NewTraitApplicationCommandLine(p unsafe.Pointer) *TraitApplicationCommandLine {
	return &TraitApplicationCommandLine{(*C.GApplicationCommandLine)(p)}
}

/*
Creates a #GFile corresponding to a filename that was given as part
of the invocation of @cmdline.

This differs from g_file_new_for_commandline_arg() in that it
resolves relative pathnames using the current working directory of
the invoking process rather than the local process.
*/
func (self *TraitApplicationCommandLine) CreateFileForArg(arg string) (return__ *C.GFile) {
	__cgo__arg := (*C.gchar)(unsafe.Pointer(C.CString(arg)))
	return__ = C.g_application_command_line_create_file_for_arg(self.CPointer, __cgo__arg)
	C.free(unsafe.Pointer(__cgo__arg))
	return
}

/*
Gets the list of arguments that was passed on the command line.

The strings in the array may contain non-UTF-8 data on UNIX (such as
filenames or arguments given in the system locale) but are always in
UTF-8 on Windows.

If you wish to use the return value with #GOptionContext, you must
use g_option_context_parse_strv().

The return value is %NULL-terminated and should be freed using
g_strfreev().
*/
func (self *TraitApplicationCommandLine) GetArguments() (argc int, return__ []string) {
	var __cgo__argc C.int
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.g_application_command_line_get_arguments(self.CPointer, &__cgo__argc)
	argc = int(__cgo__argc)
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Gets the working directory of the command line invocation.
The string may contain non-utf8 data.

It is possible that the remote application did not send a working
directory, so this may be %NULL.

The return value should not be modified or freed and is valid for as
long as @cmdline exists.
*/
func (self *TraitApplicationCommandLine) GetCwd() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_application_command_line_get_cwd(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the contents of the 'environ' variable of the command line
invocation, as would be returned by g_get_environ(), ie as a
%NULL-terminated list of strings in the form 'NAME=VALUE'.
The strings may contain non-utf8 data.

The remote application usually does not send an environment.  Use
%G_APPLICATION_SEND_ENVIRONMENT to affect that.  Even with this flag
set it is possible that the environment is still not available (due
to invocation messages from other applications).

The return value should not be modified or freed and is valid for as
long as @cmdline exists.

See g_application_command_line_getenv() if you are only interested
in the value of a single environment variable.
*/
func (self *TraitApplicationCommandLine) GetEnviron() (return__ **C.gchar) {
	return__ = C.g_application_command_line_get_environ(self.CPointer)
	return
}

/*
Gets the exit status of @cmdline.  See
g_application_command_line_set_exit_status() for more information.
*/
func (self *TraitApplicationCommandLine) GetExitStatus() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.g_application_command_line_get_exit_status(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Determines if @cmdline represents a remote invocation.
*/
func (self *TraitApplicationCommandLine) GetIsRemote() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_application_command_line_get_is_remote(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the options there were passed to g_application_command_line().

If you did not override local_command_line() then these are the same
options that were parsed according to the #GOptionEntrys added to the
application with g_application_add_main_option_entries() and possibly
modified from your GApplication::handle-local-options handler.

If no options were sent then an empty dictionary is returned so that
you don't need to check for %NULL.
*/
func (self *TraitApplicationCommandLine) GetOptionsDict() (return__ *C.GVariantDict) {
	return__ = C.g_application_command_line_get_options_dict(self.CPointer)
	return
}

/*
Gets the platform data associated with the invocation of @cmdline.

This is a #GVariant dictionary containing information about the
context in which the invocation occurred.  It typically contains
information like the current working directory and the startup
notification ID.

For local invocation, it will be %NULL.
*/
func (self *TraitApplicationCommandLine) GetPlatformData() (return__ *C.GVariant) {
	return__ = C.g_application_command_line_get_platform_data(self.CPointer)
	return
}

/*
Gets the stdin of the invoking process.

The #GInputStream can be used to read data passed to the standard
input of the invoking process.
This doesn't work on all platforms.  Presently, it is only available
on UNIX when using a DBus daemon capable of passing file descriptors.
If stdin is not available then %NULL will be returned.  In the
future, support may be expanded to other platforms.

You must only call this function once per commandline invocation.
*/
func (self *TraitApplicationCommandLine) GetStdin() (return__ *InputStream) {
	var __cgo__return__ *C.GInputStream
	__cgo__return__ = C.g_application_command_line_get_stdin(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewInputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the value of a particular environment variable of the command
line invocation, as would be returned by g_getenv().  The strings may
contain non-utf8 data.

The remote application usually does not send an environment.  Use
%G_APPLICATION_SEND_ENVIRONMENT to affect that.  Even with this flag
set it is possible that the environment is still not available (due
to invocation messages from other applications).

The return value should not be modified or freed and is valid for as
long as @cmdline exists.
*/
func (self *TraitApplicationCommandLine) Getenv(name string) (return__ string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_application_command_line_getenv(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

// g_application_command_line_print is not generated due to varargs

// g_application_command_line_printerr is not generated due to varargs

/*
Sets the exit status that will be used when the invoking process
exits.

The return value of the #GApplication::command-line signal is
passed to this function when the handler returns.  This is the usual
way of setting the exit status.

In the event that you want the remote invocation to continue running
and want to decide on the exit status in the future, you can use this
call.  For the case of a remote invocation, the remote process will
typically exit when the last reference is dropped on @cmdline.  The
exit status of the remote process will be equal to the last value
that was set with this function.

In the case that the commandline invocation is local, the situation
is slightly more complicated.  If the commandline invocation results
in the mainloop running (ie: because the use-count of the application
increased to a non-zero value) then the application is considered to
have been 'successful' in a certain sense, and the exit status is
always zero.  If the application use count is zero, though, the exit
status of the local #GApplicationCommandLine is used.
*/
func (self *TraitApplicationCommandLine) SetExitStatus(exit_status int) {
	C.g_application_command_line_set_exit_status(self.CPointer, C.int(exit_status))
	return
}

type TraitBufferedInputStream struct{ CPointer *C.GBufferedInputStream }
type IsBufferedInputStream interface {
	GetBufferedInputStreamPointer() *C.GBufferedInputStream
}

func (self *TraitBufferedInputStream) GetBufferedInputStreamPointer() *C.GBufferedInputStream {
	return self.CPointer
}
func NewTraitBufferedInputStream(p unsafe.Pointer) *TraitBufferedInputStream {
	return &TraitBufferedInputStream{(*C.GBufferedInputStream)(p)}
}

/*
Tries to read @count bytes from the stream into the buffer.
Will block during this read.

If @count is zero, returns zero and does nothing. A value of @count
larger than %G_MAXSSIZE will cause a %G_IO_ERROR_INVALID_ARGUMENT error.

On success, the number of bytes read into the buffer is returned.
It is not an error if this is not the same as the requested size, as it
can happen e.g. near the end of a file. Zero is returned on end of file
(or if @count is zero),  but never otherwise.

If @count is -1 then the attempted read size is equal to the number of
bytes that are required to fill the buffer.

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
operation was partially finished when the operation was cancelled the
partial result will be returned, without an error.

On error -1 is returned and @error is set accordingly.

For the asynchronous, non-blocking, version of this function, see
g_buffered_input_stream_fill_async().
*/
func (self *TraitBufferedInputStream) Fill(count int64, cancellable IsCancellable) (return__ int64, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_buffered_input_stream_fill(self.CPointer, C.gssize(count), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Reads data into @stream's buffer asynchronously, up to @count size.
@io_priority can be used to prioritize reads. For the synchronous
version of this function, see g_buffered_input_stream_fill().

If @count is -1 then the attempted read size is equal to the number
of bytes that are required to fill the buffer.
*/
func (self *TraitBufferedInputStream) FillAsync(count int64, io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_buffered_input_stream_fill_async(self.CPointer, C.gssize(count), C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes an asynchronous read.
*/
func (self *TraitBufferedInputStream) FillFinish(result *C.GAsyncResult) (return__ int64, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_buffered_input_stream_fill_finish(self.CPointer, result, &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the size of the available data within the stream.
*/
func (self *TraitBufferedInputStream) GetAvailable() (return__ int64) {
	var __cgo__return__ C.gsize
	__cgo__return__ = C.g_buffered_input_stream_get_available(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

/*
Gets the size of the input buffer.
*/
func (self *TraitBufferedInputStream) GetBufferSize() (return__ int64) {
	var __cgo__return__ C.gsize
	__cgo__return__ = C.g_buffered_input_stream_get_buffer_size(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

/*
Peeks in the buffer, copying data of size @count into @buffer,
offset @offset bytes.
*/
func (self *TraitBufferedInputStream) Peek(buffer []byte, offset int64, count int64) (return__ int64) {
	__header__buffer := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	var __cgo__return__ C.gsize
	__cgo__return__ = C.g_buffered_input_stream_peek(self.CPointer, (unsafe.Pointer)(unsafe.Pointer(__header__buffer.Data)), C.gsize(offset), C.gsize(count))
	return__ = int64(__cgo__return__)
	return
}

/*
Returns the buffer with the currently available bytes. The returned
buffer must not be modified and will become invalid when reading from
the stream or filling the buffer.
*/
func (self *TraitBufferedInputStream) PeekBuffer() (count int64, return__ []byte) {
	var __cgo__count C.gsize
	var __cgo__return__ unsafe.Pointer
	__cgo__return__ = C.g_buffered_input_stream_peek_buffer(self.CPointer, &__cgo__count)
	count = int64(__cgo__count)
	defer func() { return__ = C.GoBytes(unsafe.Pointer(__cgo__return__), C.int(count)) }()
	return
}

/*
Tries to read a single byte from the stream or the buffer. Will block
during this read.

On success, the byte read from the stream is returned. On end of stream
-1 is returned but it's not an exceptional error and @error is not set.

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
operation was partially finished when the operation was cancelled the
partial result will be returned, without an error.

On error -1 is returned and @error is set accordingly.
*/
func (self *TraitBufferedInputStream) ReadByte(cancellable IsCancellable) (return__ int, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.int
	__cgo__return__ = C.g_buffered_input_stream_read_byte(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Sets the size of the internal buffer of @stream to @size, or to the
size of the contents of the buffer. The buffer can never be resized
smaller than its current contents.
*/
func (self *TraitBufferedInputStream) SetBufferSize(size int64) {
	C.g_buffered_input_stream_set_buffer_size(self.CPointer, C.gsize(size))
	return
}

type TraitBufferedOutputStream struct{ CPointer *C.GBufferedOutputStream }
type IsBufferedOutputStream interface {
	GetBufferedOutputStreamPointer() *C.GBufferedOutputStream
}

func (self *TraitBufferedOutputStream) GetBufferedOutputStreamPointer() *C.GBufferedOutputStream {
	return self.CPointer
}
func NewTraitBufferedOutputStream(p unsafe.Pointer) *TraitBufferedOutputStream {
	return &TraitBufferedOutputStream{(*C.GBufferedOutputStream)(p)}
}

/*
Checks if the buffer automatically grows as data is added.
*/
func (self *TraitBufferedOutputStream) GetAutoGrow() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_buffered_output_stream_get_auto_grow(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the size of the buffer in the @stream.
*/
func (self *TraitBufferedOutputStream) GetBufferSize() (return__ int64) {
	var __cgo__return__ C.gsize
	__cgo__return__ = C.g_buffered_output_stream_get_buffer_size(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

/*
Sets whether or not the @stream's buffer should automatically grow.
If @auto_grow is true, then each write will just make the buffer
larger, and you must manually flush the buffer to actually write out
the data to the underlying stream.
*/
func (self *TraitBufferedOutputStream) SetAutoGrow(auto_grow bool) {
	__cgo__auto_grow := C.gboolean(0)
	if auto_grow {
		__cgo__auto_grow = C.gboolean(1)
	}
	C.g_buffered_output_stream_set_auto_grow(self.CPointer, __cgo__auto_grow)
	return
}

/*
Sets the size of the internal buffer to @size.
*/
func (self *TraitBufferedOutputStream) SetBufferSize(size int64) {
	C.g_buffered_output_stream_set_buffer_size(self.CPointer, C.gsize(size))
	return
}

type TraitBytesIcon struct{ CPointer *C.GBytesIcon }
type IsBytesIcon interface {
	GetBytesIconPointer() *C.GBytesIcon
}

func (self *TraitBytesIcon) GetBytesIconPointer() *C.GBytesIcon {
	return self.CPointer
}
func NewTraitBytesIcon(p unsafe.Pointer) *TraitBytesIcon {
	return &TraitBytesIcon{(*C.GBytesIcon)(p)}
}

/*
Gets the #GBytes associated with the given @icon.
*/
func (self *TraitBytesIcon) GetBytes() (return__ *C.GBytes) {
	return__ = C.g_bytes_icon_get_bytes(self.CPointer)
	return
}

type TraitCancellable struct{ CPointer *C.GCancellable }
type IsCancellable interface {
	GetCancellablePointer() *C.GCancellable
}

func (self *TraitCancellable) GetCancellablePointer() *C.GCancellable {
	return self.CPointer
}
func NewTraitCancellable(p unsafe.Pointer) *TraitCancellable {
	return &TraitCancellable{(*C.GCancellable)(p)}
}

/*
Will set @cancellable to cancelled, and will emit the
#GCancellable::cancelled signal. (However, see the warning about
race conditions in the documentation for that signal if you are
planning to connect to it.)

This function is thread-safe. In other words, you can safely call
it from a thread other than the one running the operation that was
passed the @cancellable.

The convention within gio is that cancelling an asynchronous
operation causes it to complete asynchronously. That is, if you
cancel the operation from the same thread in which it is running,
then the operation's #GAsyncReadyCallback will not be invoked until
the application returns to the main loop.
*/
func (self *TraitCancellable) Cancel() {
	C.g_cancellable_cancel(self.CPointer)
	return
}

/*
Convenience function to connect to the #GCancellable::cancelled
signal. Also handles the race condition that may happen
if the cancellable is cancelled right before connecting.

@callback is called at most once, either directly at the
time of the connect if @cancellable is already cancelled,
or when @cancellable is cancelled in some thread.

@data_destroy_func will be called when the handler is
disconnected, or immediately if the cancellable is already
cancelled.

See #GCancellable::cancelled for details on how to use this.

Since GLib 2.40, the lock protecting @cancellable is not held when
@callback is invoked.  This lifts a restriction in place for
earlier GLib versions which now makes it easier to write cleanup
code that unconditionally invokes e.g. g_cancellable_cancel().
*/
func (self *TraitCancellable) Connect(callback C.GCallback, data unsafe.Pointer, data_destroy_func C.GDestroyNotify) (return__ uint64) {
	var __cgo__return__ C.gulong
	__cgo__return__ = C.g_cancellable_connect(self.CPointer, callback, (C.gpointer)(data), data_destroy_func)
	return__ = uint64(__cgo__return__)
	return
}

/*
Disconnects a handler from a cancellable instance similar to
g_signal_handler_disconnect().  Additionally, in the event that a
signal handler is currently running, this call will block until the
handler has finished.  Calling this function from a
#GCancellable::cancelled signal handler will therefore result in a
deadlock.

This avoids a race condition where a thread cancels at the
same time as the cancellable operation is finished and the
signal handler is removed. See #GCancellable::cancelled for
details on how to use this.

If @cancellable is %NULL or @handler_id is %0 this function does
nothing.
*/
func (self *TraitCancellable) Disconnect(handler_id uint64) {
	C.g_cancellable_disconnect(self.CPointer, C.gulong(handler_id))
	return
}

/*
Gets the file descriptor for a cancellable job. This can be used to
implement cancellable operations on Unix systems. The returned fd will
turn readable when @cancellable is cancelled.

You are not supposed to read from the fd yourself, just check for
readable status. Reading to unset the readable status is done
with g_cancellable_reset().

After a successful return from this function, you should use
g_cancellable_release_fd() to free up resources allocated for
the returned file descriptor.

See also g_cancellable_make_pollfd().
*/
func (self *TraitCancellable) GetFd() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.g_cancellable_get_fd(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Checks if a cancellable job has been cancelled.
*/
func (self *TraitCancellable) IsCancelled() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_cancellable_is_cancelled(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Creates a #GPollFD corresponding to @cancellable; this can be passed
to g_poll() and used to poll for cancellation. This is useful both
for unix systems without a native poll and for portability to
windows.

When this function returns %TRUE, you should use
g_cancellable_release_fd() to free up resources allocated for the
@pollfd. After a %FALSE return, do not call g_cancellable_release_fd().

If this function returns %FALSE, either no @cancellable was given or
resource limits prevent this function from allocating the necessary
structures for polling. (On Linux, you will likely have reached
the maximum number of file descriptors.) The suggested way to handle
these cases is to ignore the @cancellable.

You are not supposed to read from the fd yourself, just check for
readable status. Reading to unset the readable status is done
with g_cancellable_reset().
*/
func (self *TraitCancellable) MakePollfd(pollfd *C.GPollFD) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_cancellable_make_pollfd(self.CPointer, pollfd)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Pops @cancellable off the cancellable stack (verifying that @cancellable
is on the top of the stack).
*/
func (self *TraitCancellable) PopCurrent() {
	C.g_cancellable_pop_current(self.CPointer)
	return
}

/*
Pushes @cancellable onto the cancellable stack. The current
cancellable can then be received using g_cancellable_get_current().

This is useful when implementing cancellable operations in
code that does not allow you to pass down the cancellable object.

This is typically called automatically by e.g. #GFile operations,
so you rarely have to call this yourself.
*/
func (self *TraitCancellable) PushCurrent() {
	C.g_cancellable_push_current(self.CPointer)
	return
}

/*
Releases a resources previously allocated by g_cancellable_get_fd()
or g_cancellable_make_pollfd().

For compatibility reasons with older releases, calling this function
is not strictly required, the resources will be automatically freed
when the @cancellable is finalized. However, the @cancellable will
block scarce file descriptors until it is finalized if this function
is not called. This can cause the application to run out of file
descriptors when many #GCancellables are used at the same time.
*/
func (self *TraitCancellable) ReleaseFd() {
	C.g_cancellable_release_fd(self.CPointer)
	return
}

/*
Resets @cancellable to its uncancelled state.

If cancellable is currently in use by any cancellable operation
then the behavior of this function is undefined.
*/
func (self *TraitCancellable) Reset() {
	C.g_cancellable_reset(self.CPointer)
	return
}

/*
If the @cancellable is cancelled, sets the error to notify
that the operation was cancelled.
*/
func (self *TraitCancellable) SetErrorIfCancelled() (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_cancellable_set_error_if_cancelled(self.CPointer, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a source that triggers if @cancellable is cancelled and
calls its callback of type #GCancellableSourceFunc. This is
primarily useful for attaching to another (non-cancellable) source
with g_source_add_child_source() to add cancellability to it.

For convenience, you can call this with a %NULL #GCancellable,
in which case the source will never trigger.
*/
func (self *TraitCancellable) SourceNew() (return__ *C.GSource) {
	return__ = C.g_cancellable_source_new(self.CPointer)
	return
}

type TraitCharsetConverter struct{ CPointer *C.GCharsetConverter }
type IsCharsetConverter interface {
	GetCharsetConverterPointer() *C.GCharsetConverter
}

func (self *TraitCharsetConverter) GetCharsetConverterPointer() *C.GCharsetConverter {
	return self.CPointer
}
func NewTraitCharsetConverter(p unsafe.Pointer) *TraitCharsetConverter {
	return &TraitCharsetConverter{(*C.GCharsetConverter)(p)}
}

/*
Gets the number of fallbacks that @converter has applied so far.
*/
func (self *TraitCharsetConverter) GetNumFallbacks() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_charset_converter_get_num_fallbacks(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Gets the #GCharsetConverter:use-fallback property.
*/
func (self *TraitCharsetConverter) GetUseFallback() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_charset_converter_get_use_fallback(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the #GCharsetConverter:use-fallback property.
*/
func (self *TraitCharsetConverter) SetUseFallback(use_fallback bool) {
	__cgo__use_fallback := C.gboolean(0)
	if use_fallback {
		__cgo__use_fallback = C.gboolean(1)
	}
	C.g_charset_converter_set_use_fallback(self.CPointer, __cgo__use_fallback)
	return
}

type TraitConverterInputStream struct{ CPointer *C.GConverterInputStream }
type IsConverterInputStream interface {
	GetConverterInputStreamPointer() *C.GConverterInputStream
}

func (self *TraitConverterInputStream) GetConverterInputStreamPointer() *C.GConverterInputStream {
	return self.CPointer
}
func NewTraitConverterInputStream(p unsafe.Pointer) *TraitConverterInputStream {
	return &TraitConverterInputStream{(*C.GConverterInputStream)(p)}
}

/*
Gets the #GConverter that is used by @converter_stream.
*/
func (self *TraitConverterInputStream) GetConverter() (return__ *C.GConverter) {
	return__ = C.g_converter_input_stream_get_converter(self.CPointer)
	return
}

type TraitConverterOutputStream struct{ CPointer *C.GConverterOutputStream }
type IsConverterOutputStream interface {
	GetConverterOutputStreamPointer() *C.GConverterOutputStream
}

func (self *TraitConverterOutputStream) GetConverterOutputStreamPointer() *C.GConverterOutputStream {
	return self.CPointer
}
func NewTraitConverterOutputStream(p unsafe.Pointer) *TraitConverterOutputStream {
	return &TraitConverterOutputStream{(*C.GConverterOutputStream)(p)}
}

/*
Gets the #GConverter that is used by @converter_stream.
*/
func (self *TraitConverterOutputStream) GetConverter() (return__ *C.GConverter) {
	return__ = C.g_converter_output_stream_get_converter(self.CPointer)
	return
}

type TraitCredentials struct{ CPointer *C.GCredentials }
type IsCredentials interface {
	GetCredentialsPointer() *C.GCredentials
}

func (self *TraitCredentials) GetCredentialsPointer() *C.GCredentials {
	return self.CPointer
}
func NewTraitCredentials(p unsafe.Pointer) *TraitCredentials {
	return &TraitCredentials{(*C.GCredentials)(p)}
}

/*
Gets a pointer to native credentials of type @native_type from
@credentials.

It is a programming error (which will cause an warning to be
logged) to use this method if there is no #GCredentials support for
the OS or if @native_type isn't supported by the OS.
*/
func (self *TraitCredentials) GetNative(native_type C.GCredentialsType) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_credentials_get_native(self.CPointer, native_type)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Tries to get the UNIX process identifier from @credentials. This
method is only available on UNIX platforms.

This operation can fail if #GCredentials is not supported on the
OS or if the native credentials type does not contain information
about the UNIX process ID.
*/
func (self *TraitCredentials) GetUnixPid() (return__ int, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.pid_t
	__cgo__return__ = C.g_credentials_get_unix_pid(self.CPointer, &__cgo_error__)
	return__ = int(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Tries to get the UNIX user identifier from @credentials. This
method is only available on UNIX platforms.

This operation can fail if #GCredentials is not supported on the
OS or if the native credentials type does not contain information
about the UNIX user.
*/
func (self *TraitCredentials) GetUnixUser() (return__ uint, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.uid_t
	__cgo__return__ = C.g_credentials_get_unix_user(self.CPointer, &__cgo_error__)
	return__ = uint(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Checks if @credentials and @other_credentials is the same user.

This operation can fail if #GCredentials is not supported on the
the OS.
*/
func (self *TraitCredentials) IsSameUser(other_credentials IsCredentials) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_credentials_is_same_user(self.CPointer, other_credentials.GetCredentialsPointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Copies the native credentials of type @native_type from @native
into @credentials.

It is a programming error (which will cause an warning to be
logged) to use this method if there is no #GCredentials support for
the OS or if @native_type isn't supported by the OS.
*/
func (self *TraitCredentials) SetNative(native_type C.GCredentialsType, native unsafe.Pointer) {
	C.g_credentials_set_native(self.CPointer, native_type, (C.gpointer)(native))
	return
}

/*
Tries to set the UNIX user identifier on @credentials. This method
is only available on UNIX platforms.

This operation can fail if #GCredentials is not supported on the
OS or if the native credentials type does not contain information
about the UNIX user. It can also fail if the OS does not allow the
use of "spoofed" credentials.
*/
func (self *TraitCredentials) SetUnixUser(uid uint) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_credentials_set_unix_user(self.CPointer, C.uid_t(uid), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a human-readable textual representation of @credentials
that can be used in logging and debug messages. The format of the
returned string may change in future GLib release.
*/
func (self *TraitCredentials) ToString() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_credentials_to_string(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

type TraitDBusActionGroup struct{ CPointer *C.GDBusActionGroup }
type IsDBusActionGroup interface {
	GetDBusActionGroupPointer() *C.GDBusActionGroup
}

func (self *TraitDBusActionGroup) GetDBusActionGroupPointer() *C.GDBusActionGroup {
	return self.CPointer
}
func NewTraitDBusActionGroup(p unsafe.Pointer) *TraitDBusActionGroup {
	return &TraitDBusActionGroup{(*C.GDBusActionGroup)(p)}
}

type TraitDBusAuthObserver struct{ CPointer *C.GDBusAuthObserver }
type IsDBusAuthObserver interface {
	GetDBusAuthObserverPointer() *C.GDBusAuthObserver
}

func (self *TraitDBusAuthObserver) GetDBusAuthObserverPointer() *C.GDBusAuthObserver {
	return self.CPointer
}
func NewTraitDBusAuthObserver(p unsafe.Pointer) *TraitDBusAuthObserver {
	return &TraitDBusAuthObserver{(*C.GDBusAuthObserver)(p)}
}

/*
Emits the #GDBusAuthObserver::allow-mechanism signal on @observer.
*/
func (self *TraitDBusAuthObserver) AllowMechanism(mechanism string) (return__ bool) {
	__cgo__mechanism := (*C.gchar)(unsafe.Pointer(C.CString(mechanism)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_auth_observer_allow_mechanism(self.CPointer, __cgo__mechanism)
	C.free(unsafe.Pointer(__cgo__mechanism))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Emits the #GDBusAuthObserver::authorize-authenticated-peer signal on @observer.
*/
func (self *TraitDBusAuthObserver) AuthorizeAuthenticatedPeer(stream IsIOStream, credentials IsCredentials) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_auth_observer_authorize_authenticated_peer(self.CPointer, stream.GetIOStreamPointer(), credentials.GetCredentialsPointer())
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

type TraitDBusConnection struct{ CPointer *C.GDBusConnection }
type IsDBusConnection interface {
	GetDBusConnectionPointer() *C.GDBusConnection
}

func (self *TraitDBusConnection) GetDBusConnectionPointer() *C.GDBusConnection {
	return self.CPointer
}
func NewTraitDBusConnection(p unsafe.Pointer) *TraitDBusConnection {
	return &TraitDBusConnection{(*C.GDBusConnection)(p)}
}

/*
Adds a message filter. Filters are handlers that are run on all
incoming and outgoing messages, prior to standard dispatch. Filters
are run in the order that they were added.  The same handler can be
added as a filter more than once, in which case it will be run more
than once.  Filters added during a filter callback won't be run on
the message being processed. Filter functions are allowed to modify
and even drop messages.

Note that filters are run in a dedicated message handling thread so
they can't block and, generally, can't do anything but signal a
worker thread. Also note that filters are rarely needed - use API
such as g_dbus_connection_send_message_with_reply(),
g_dbus_connection_signal_subscribe() or g_dbus_connection_call() instead.

If a filter consumes an incoming message the message is not
dispatched anywhere else - not even the standard dispatch machinery
(that API such as g_dbus_connection_signal_subscribe() and
g_dbus_connection_send_message_with_reply() relies on) will see the
message. Similary, if a filter consumes an outgoing message, the
message will not be sent to the other peer.
*/
func (self *TraitDBusConnection) AddFilter(filter_function C.GDBusMessageFilterFunction, user_data unsafe.Pointer, user_data_free_func C.GDestroyNotify) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_dbus_connection_add_filter(self.CPointer, filter_function, (C.gpointer)(user_data), user_data_free_func)
	return__ = uint(__cgo__return__)
	return
}

/*
Asynchronously invokes the @method_name method on the
@interface_name D-Bus interface on the remote object at
@object_path owned by @bus_name.

If @connection is closed then the operation will fail with
%G_IO_ERROR_CLOSED. If @cancellable is canceled, the operation will
fail with %G_IO_ERROR_CANCELLED. If @parameters contains a value
not compatible with the D-Bus protocol, the operation fails with
%G_IO_ERROR_INVALID_ARGUMENT.

If @reply_type is non-%NULL then the reply will be checked for having this type and an
error will be raised if it does not match.  Said another way, if you give a @reply_type
then any non-%NULL return value will be of this type.

If the @parameters #GVariant is floating, it is consumed. This allows
convenient 'inline' use of g_variant_new(), e.g.:
|[<!-- language="C" -->
 g_dbus_connection_call (connection,
                         "org.freedesktop.StringThings",
                         "/org/freedesktop/StringThings",
                         "org.freedesktop.StringThings",
                         "TwoStrings",
                         g_variant_new ("(ss)",
                                        "Thing One",
                                        "Thing Two"),
                         NULL,
                         G_DBUS_CALL_FLAGS_NONE,
                         -1,
                         NULL,
                         (GAsyncReadyCallback) two_strings_done,
                         NULL);
]|

This is an asynchronous method. When the operation is finished,
@callback will be invoked in the
[thread-default main context][g-main-context-push-thread-default]
of the thread you are calling this method from. You can then call
g_dbus_connection_call_finish() to get the result of the operation.
See g_dbus_connection_call_sync() for the synchronous version of this
function.

If @callback is %NULL then the D-Bus method call message will be sent with
the %G_DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED flag set.
*/
func (self *TraitDBusConnection) Call(bus_name string, object_path string, interface_name string, method_name string, parameters *C.GVariant, reply_type *C.GVariantType, flags C.GDBusCallFlags, timeout_msec int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__bus_name := (*C.gchar)(unsafe.Pointer(C.CString(bus_name)))
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	__cgo__interface_name := (*C.gchar)(unsafe.Pointer(C.CString(interface_name)))
	__cgo__method_name := (*C.gchar)(unsafe.Pointer(C.CString(method_name)))
	C.g_dbus_connection_call(self.CPointer, __cgo__bus_name, __cgo__object_path, __cgo__interface_name, __cgo__method_name, parameters, reply_type, flags, C.gint(timeout_msec), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__bus_name))
	C.free(unsafe.Pointer(__cgo__object_path))
	C.free(unsafe.Pointer(__cgo__interface_name))
	C.free(unsafe.Pointer(__cgo__method_name))
	return
}

/*
Finishes an operation started with g_dbus_connection_call().
*/
func (self *TraitDBusConnection) CallFinish(res *C.GAsyncResult) (return__ *C.GVariant, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_dbus_connection_call_finish(self.CPointer, res, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Synchronously invokes the @method_name method on the
@interface_name D-Bus interface on the remote object at
@object_path owned by @bus_name.

If @connection is closed then the operation will fail with
%G_IO_ERROR_CLOSED. If @cancellable is canceled, the
operation will fail with %G_IO_ERROR_CANCELLED. If @parameters
contains a value not compatible with the D-Bus protocol, the operation
fails with %G_IO_ERROR_INVALID_ARGUMENT.

If @reply_type is non-%NULL then the reply will be checked for having
this type and an error will be raised if it does not match.  Said
another way, if you give a @reply_type then any non-%NULL return
value will be of this type.

If the @parameters #GVariant is floating, it is consumed.
This allows convenient 'inline' use of g_variant_new(), e.g.:
|[<!-- language="C" -->
 g_dbus_connection_call_sync (connection,
                              "org.freedesktop.StringThings",
                              "/org/freedesktop/StringThings",
                              "org.freedesktop.StringThings",
                              "TwoStrings",
                              g_variant_new ("(ss)",
                                             "Thing One",
                                             "Thing Two"),
                              NULL,
                              G_DBUS_CALL_FLAGS_NONE,
                              -1,
                              NULL,
                              &error);
]|

The calling thread is blocked until a reply is received. See
g_dbus_connection_call() for the asynchronous version of
this method.
*/
func (self *TraitDBusConnection) CallSync(bus_name string, object_path string, interface_name string, method_name string, parameters *C.GVariant, reply_type *C.GVariantType, flags C.GDBusCallFlags, timeout_msec int, cancellable IsCancellable) (return__ *C.GVariant, __err__ error) {
	__cgo__bus_name := (*C.gchar)(unsafe.Pointer(C.CString(bus_name)))
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	__cgo__interface_name := (*C.gchar)(unsafe.Pointer(C.CString(interface_name)))
	__cgo__method_name := (*C.gchar)(unsafe.Pointer(C.CString(method_name)))
	var __cgo_error__ *C.GError
	return__ = C.g_dbus_connection_call_sync(self.CPointer, __cgo__bus_name, __cgo__object_path, __cgo__interface_name, __cgo__method_name, parameters, reply_type, flags, C.gint(timeout_msec), cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__bus_name))
	C.free(unsafe.Pointer(__cgo__object_path))
	C.free(unsafe.Pointer(__cgo__interface_name))
	C.free(unsafe.Pointer(__cgo__method_name))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Like g_dbus_connection_call() but also takes a #GUnixFDList object.

This method is only available on UNIX.
*/
func (self *TraitDBusConnection) CallWithUnixFdList(bus_name string, object_path string, interface_name string, method_name string, parameters *C.GVariant, reply_type *C.GVariantType, flags C.GDBusCallFlags, timeout_msec int, fd_list IsUnixFDList, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__bus_name := (*C.gchar)(unsafe.Pointer(C.CString(bus_name)))
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	__cgo__interface_name := (*C.gchar)(unsafe.Pointer(C.CString(interface_name)))
	__cgo__method_name := (*C.gchar)(unsafe.Pointer(C.CString(method_name)))
	C.g_dbus_connection_call_with_unix_fd_list(self.CPointer, __cgo__bus_name, __cgo__object_path, __cgo__interface_name, __cgo__method_name, parameters, reply_type, flags, C.gint(timeout_msec), fd_list.GetUnixFDListPointer(), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__bus_name))
	C.free(unsafe.Pointer(__cgo__object_path))
	C.free(unsafe.Pointer(__cgo__interface_name))
	C.free(unsafe.Pointer(__cgo__method_name))
	return
}

/*
Finishes an operation started with g_dbus_connection_call_with_unix_fd_list().
*/
func (self *TraitDBusConnection) CallWithUnixFdListFinish(res *C.GAsyncResult) (out_fd_list *UnixFDList, return__ *C.GVariant, __err__ error) {
	var __cgo__out_fd_list *C.GUnixFDList
	var __cgo_error__ *C.GError
	return__ = C.g_dbus_connection_call_with_unix_fd_list_finish(self.CPointer, &__cgo__out_fd_list, res, &__cgo_error__)
	if __cgo__out_fd_list != nil {
		out_fd_list = NewUnixFDListFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__out_fd_list).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Like g_dbus_connection_call_sync() but also takes and returns #GUnixFDList objects.

This method is only available on UNIX.
*/
func (self *TraitDBusConnection) CallWithUnixFdListSync(bus_name string, object_path string, interface_name string, method_name string, parameters *C.GVariant, reply_type *C.GVariantType, flags C.GDBusCallFlags, timeout_msec int, fd_list IsUnixFDList, cancellable IsCancellable) (out_fd_list *UnixFDList, return__ *C.GVariant, __err__ error) {
	__cgo__bus_name := (*C.gchar)(unsafe.Pointer(C.CString(bus_name)))
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	__cgo__interface_name := (*C.gchar)(unsafe.Pointer(C.CString(interface_name)))
	__cgo__method_name := (*C.gchar)(unsafe.Pointer(C.CString(method_name)))
	var __cgo__out_fd_list *C.GUnixFDList
	var __cgo_error__ *C.GError
	return__ = C.g_dbus_connection_call_with_unix_fd_list_sync(self.CPointer, __cgo__bus_name, __cgo__object_path, __cgo__interface_name, __cgo__method_name, parameters, reply_type, flags, C.gint(timeout_msec), fd_list.GetUnixFDListPointer(), &__cgo__out_fd_list, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__bus_name))
	C.free(unsafe.Pointer(__cgo__object_path))
	C.free(unsafe.Pointer(__cgo__interface_name))
	C.free(unsafe.Pointer(__cgo__method_name))
	if __cgo__out_fd_list != nil {
		out_fd_list = NewUnixFDListFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__out_fd_list).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Closes @connection. Note that this never causes the process to
exit (this might only happen if the other end of a shared message
bus connection disconnects, see #GDBusConnection:exit-on-close).

Once the connection is closed, operations such as sending a message
will return with the error %G_IO_ERROR_CLOSED. Closing a connection
will not automatically flush the connection so queued messages may
be lost. Use g_dbus_connection_flush() if you need such guarantees.

If @connection is already closed, this method fails with
%G_IO_ERROR_CLOSED.

When @connection has been closed, the #GDBusConnection::closed
signal is emitted in the
[thread-default main context][g-main-context-push-thread-default]
of the thread that @connection was constructed in.

This is an asynchronous method. When the operation is finished,
@callback will be invoked in the
[thread-default main context][g-main-context-push-thread-default]
of the thread you are calling this method from. You can
then call g_dbus_connection_close_finish() to get the result of the
operation. See g_dbus_connection_close_sync() for the synchronous
version.
*/
func (self *TraitDBusConnection) Close(cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_dbus_connection_close(self.CPointer, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes an operation started with g_dbus_connection_close().
*/
func (self *TraitDBusConnection) CloseFinish(res *C.GAsyncResult) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_connection_close_finish(self.CPointer, res, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Synchronously closees @connection. The calling thread is blocked
until this is done. See g_dbus_connection_close() for the
asynchronous version of this method and more details about what it
does.
*/
func (self *TraitDBusConnection) CloseSync(cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_connection_close_sync(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Emits a signal.

If the parameters GVariant is floating, it is consumed.

This can only fail if @parameters is not compatible with the D-Bus protocol.
*/
func (self *TraitDBusConnection) EmitSignal(destination_bus_name string, object_path string, interface_name string, signal_name string, parameters *C.GVariant) (return__ bool, __err__ error) {
	__cgo__destination_bus_name := (*C.gchar)(unsafe.Pointer(C.CString(destination_bus_name)))
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	__cgo__interface_name := (*C.gchar)(unsafe.Pointer(C.CString(interface_name)))
	__cgo__signal_name := (*C.gchar)(unsafe.Pointer(C.CString(signal_name)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_connection_emit_signal(self.CPointer, __cgo__destination_bus_name, __cgo__object_path, __cgo__interface_name, __cgo__signal_name, parameters, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__destination_bus_name))
	C.free(unsafe.Pointer(__cgo__object_path))
	C.free(unsafe.Pointer(__cgo__interface_name))
	C.free(unsafe.Pointer(__cgo__signal_name))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Exports @action_group on @connection at @object_path.

The implemented D-Bus API should be considered private.  It is
subject to change in the future.

A given object path can only have one action group exported on it.
If this constraint is violated, the export will fail and 0 will be
returned (with @error set accordingly).

You can unexport the action group using
g_dbus_connection_unexport_action_group() with the return value of
this function.

The thread default main context is taken at the time of this call.
All incoming action activations and state change requests are
reported from this context.  Any changes on the action group that
cause it to emit signals must also come from this same context.
Since incoming action activations and state change requests are
rather likely to cause changes on the action group, this effectively
limits a given action group to being exported from only one main
context.
*/
func (self *TraitDBusConnection) ExportActionGroup(object_path string, action_group *C.GActionGroup) (return__ uint, __err__ error) {
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_dbus_connection_export_action_group(self.CPointer, __cgo__object_path, action_group, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__object_path))
	return__ = uint(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Exports @menu on @connection at @object_path.

The implemented D-Bus API should be considered private.
It is subject to change in the future.

An object path can only have one menu model exported on it. If this
constraint is violated, the export will fail and 0 will be
returned (with @error set accordingly).

You can unexport the menu model using
g_dbus_connection_unexport_menu_model() with the return value of
this function.
*/
func (self *TraitDBusConnection) ExportMenuModel(object_path string, menu IsMenuModel) (return__ uint, __err__ error) {
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_dbus_connection_export_menu_model(self.CPointer, __cgo__object_path, menu.GetMenuModelPointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__object_path))
	return__ = uint(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously flushes @connection, that is, writes all queued
outgoing message to the transport and then flushes the transport
(using g_output_stream_flush_async()). This is useful in programs
that wants to emit a D-Bus signal and then exit immediately. Without
flushing the connection, there is no guaranteed that the message has
been sent to the networking buffers in the OS kernel.

This is an asynchronous method. When the operation is finished,
@callback will be invoked in the
[thread-default main context][g-main-context-push-thread-default]
of the thread you are calling this method from. You can
then call g_dbus_connection_flush_finish() to get the result of the
operation. See g_dbus_connection_flush_sync() for the synchronous
version.
*/
func (self *TraitDBusConnection) Flush(cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_dbus_connection_flush(self.CPointer, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes an operation started with g_dbus_connection_flush().
*/
func (self *TraitDBusConnection) FlushFinish(res *C.GAsyncResult) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_connection_flush_finish(self.CPointer, res, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Synchronously flushes @connection. The calling thread is blocked
until this is done. See g_dbus_connection_flush() for the
asynchronous version of this method and more details about what it
does.
*/
func (self *TraitDBusConnection) FlushSync(cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_connection_flush_sync(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the capabilities negotiated with the remote peer
*/
func (self *TraitDBusConnection) GetCapabilities() (return__ C.GDBusCapabilityFlags) {
	return__ = C.g_dbus_connection_get_capabilities(self.CPointer)
	return
}

/*
Gets whether the process is terminated when @connection is
closed by the remote peer. See
#GDBusConnection:exit-on-close for more details.
*/
func (self *TraitDBusConnection) GetExitOnClose() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_connection_get_exit_on_close(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
The GUID of the peer performing the role of server when
authenticating. See #GDBusConnection:guid for more details.
*/
func (self *TraitDBusConnection) GetGuid() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_connection_get_guid(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Retrieves the last serial number assigned to a #GDBusMessage on
the current thread. This includes messages sent via both low-level
API such as g_dbus_connection_send_message() as well as
high-level API such as g_dbus_connection_emit_signal(),
g_dbus_connection_call() or g_dbus_proxy_call().
*/
func (self *TraitDBusConnection) GetLastSerial() (return__ uint32) {
	var __cgo__return__ C.guint32
	__cgo__return__ = C.g_dbus_connection_get_last_serial(self.CPointer)
	return__ = uint32(__cgo__return__)
	return
}

/*
Gets the credentials of the authenticated peer. This will always
return %NULL unless @connection acted as a server
(e.g. %G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER was passed)
when set up and the client passed credentials as part of the
authentication process.

In a message bus setup, the message bus is always the server and
each application is a client. So this method will always return
%NULL for message bus clients.
*/
func (self *TraitDBusConnection) GetPeerCredentials() (return__ *Credentials) {
	var __cgo__return__ *C.GCredentials
	__cgo__return__ = C.g_dbus_connection_get_peer_credentials(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewCredentialsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the underlying stream used for IO.

While the #GDBusConnection is active, it will interact with this
stream from a worker thread, so it is not safe to interact with
the stream directly.
*/
func (self *TraitDBusConnection) GetStream() (return__ *IOStream) {
	var __cgo__return__ *C.GIOStream
	__cgo__return__ = C.g_dbus_connection_get_stream(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewIOStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the unique name of @connection as assigned by the message
bus. This can also be used to figure out if @connection is a
message bus connection.
*/
func (self *TraitDBusConnection) GetUniqueName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_connection_get_unique_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets whether @connection is closed.
*/
func (self *TraitDBusConnection) IsClosed() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_connection_is_closed(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Registers callbacks for exported objects at @object_path with the
D-Bus interface that is described in @interface_info.

Calls to functions in @vtable (and @user_data_free_func) will happen
in the
[thread-default main context][g-main-context-push-thread-default]
of the thread you are calling this method from.

Note that all #GVariant values passed to functions in @vtable will match
the signature given in @interface_info - if a remote caller passes
incorrect values, the `org.freedesktop.DBus.Error.InvalidArgs`
is returned to the remote caller.

Additionally, if the remote caller attempts to invoke methods or
access properties not mentioned in @interface_info the
`org.freedesktop.DBus.Error.UnknownMethod` resp.
`org.freedesktop.DBus.Error.InvalidArgs` errors
are returned to the caller.

It is considered a programming error if the
#GDBusInterfaceGetPropertyFunc function in @vtable returns a
#GVariant of incorrect type.

If an existing callback is already registered at @object_path and
@interface_name, then @error is set to #G_IO_ERROR_EXISTS.

GDBus automatically implements the standard D-Bus interfaces
org.freedesktop.DBus.Properties, org.freedesktop.DBus.Introspectable
and org.freedesktop.Peer, so you don't have to implement those for the
objects you export. You can implement org.freedesktop.DBus.Properties
yourself, e.g. to handle getting and setting of properties asynchronously.

Note that the reference count on @interface_info will be
incremented by 1 (unless allocated statically, e.g. if the
reference count is -1, see g_dbus_interface_info_ref()) for as long
as the object is exported. Also note that @vtable will be copied.

See this [server][gdbus-server] for an example of how to use this method.
*/
func (self *TraitDBusConnection) RegisterObject(object_path string, interface_info *C.GDBusInterfaceInfo, vtable *C.GDBusInterfaceVTable, user_data unsafe.Pointer, user_data_free_func C.GDestroyNotify) (return__ uint, __err__ error) {
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_dbus_connection_register_object(self.CPointer, __cgo__object_path, interface_info, vtable, (C.gpointer)(user_data), user_data_free_func, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__object_path))
	return__ = uint(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Registers a whole subtree of dynamic objects.

The @enumerate and @introspection functions in @vtable are used to
convey, to remote callers, what nodes exist in the subtree rooted
by @object_path.

When handling remote calls into any node in the subtree, first the
@enumerate function is used to check if the node exists. If the node exists
or the #G_DBUS_SUBTREE_FLAGS_DISPATCH_TO_UNENUMERATED_NODES flag is set
the @introspection function is used to check if the node supports the
requested method. If so, the @dispatch function is used to determine
where to dispatch the call. The collected #GDBusInterfaceVTable and
#gpointer will be used to call into the interface vtable for processing
the request.

All calls into user-provided code will be invoked in the
[thread-default main context][g-main-context-push-thread-default]
of the thread you are calling this method from.

If an existing subtree is already registered at @object_path or
then @error is set to #G_IO_ERROR_EXISTS.

Note that it is valid to register regular objects (using
g_dbus_connection_register_object()) in a subtree registered with
g_dbus_connection_register_subtree() - if so, the subtree handler
is tried as the last resort. One way to think about a subtree
handler is to consider it a fallback handler for object paths not
registered via g_dbus_connection_register_object() or other bindings.

Note that @vtable will be copied so you cannot change it after
registration.

See this [server][gdbus-subtree-server] for an example of how to use
this method.
*/
func (self *TraitDBusConnection) RegisterSubtree(object_path string, vtable *C.GDBusSubtreeVTable, flags C.GDBusSubtreeFlags, user_data unsafe.Pointer, user_data_free_func C.GDestroyNotify) (return__ uint, __err__ error) {
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_dbus_connection_register_subtree(self.CPointer, __cgo__object_path, vtable, flags, (C.gpointer)(user_data), user_data_free_func, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__object_path))
	return__ = uint(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Removes a filter.
*/
func (self *TraitDBusConnection) RemoveFilter(filter_id uint) {
	C.g_dbus_connection_remove_filter(self.CPointer, C.guint(filter_id))
	return
}

/*
Asynchronously sends @message to the peer represented by @connection.

Unless @flags contain the
%G_DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL flag, the serial number
will be assigned by @connection and set on @message via
g_dbus_message_set_serial(). If @out_serial is not %NULL, then the
serial number used will be written to this location prior to
submitting the message to the underlying transport.

If @connection is closed then the operation will fail with
%G_IO_ERROR_CLOSED. If @message is not well-formed,
the operation fails with %G_IO_ERROR_INVALID_ARGUMENT.

See this [server][gdbus-server] and [client][gdbus-unix-fd-client]
for an example of how to use this low-level API to send and receive
UNIX file descriptors.

Note that @message must be unlocked, unless @flags contain the
%G_DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL flag.
*/
func (self *TraitDBusConnection) SendMessage(message IsDBusMessage, flags C.GDBusSendMessageFlags) (out_serial uint32, return__ bool, __err__ error) {
	var __cgo__out_serial C.guint32
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_connection_send_message(self.CPointer, message.GetDBusMessagePointer(), flags, &__cgo__out_serial, &__cgo_error__)
	out_serial = uint32(__cgo__out_serial)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously sends @message to the peer represented by @connection.

Unless @flags contain the
%G_DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL flag, the serial number
will be assigned by @connection and set on @message via
g_dbus_message_set_serial(). If @out_serial is not %NULL, then the
serial number used will be written to this location prior to
submitting the message to the underlying transport.

If @connection is closed then the operation will fail with
%G_IO_ERROR_CLOSED. If @cancellable is canceled, the operation will
fail with %G_IO_ERROR_CANCELLED. If @message is not well-formed,
the operation fails with %G_IO_ERROR_INVALID_ARGUMENT.

This is an asynchronous method. When the operation is finished, @callback
will be invoked in the
[thread-default main context][g-main-context-push-thread-default]
of the thread you are calling this method from. You can then call
g_dbus_connection_send_message_with_reply_finish() to get the result of the operation.
See g_dbus_connection_send_message_with_reply_sync() for the synchronous version.

Note that @message must be unlocked, unless @flags contain the
%G_DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL flag.

See this [server][gdbus-server] and [client][gdbus-unix-fd-client]
for an example of how to use this low-level API to send and receive
UNIX file descriptors.
*/
func (self *TraitDBusConnection) SendMessageWithReply(message IsDBusMessage, flags C.GDBusSendMessageFlags, timeout_msec int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) (out_serial uint32) {
	var __cgo__out_serial C.guint32
	C.g_dbus_connection_send_message_with_reply(self.CPointer, message.GetDBusMessagePointer(), flags, C.gint(timeout_msec), &__cgo__out_serial, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	out_serial = uint32(__cgo__out_serial)
	return
}

/*
Finishes an operation started with g_dbus_connection_send_message_with_reply().

Note that @error is only set if a local in-process error
occurred. That is to say that the returned #GDBusMessage object may
be of type %G_DBUS_MESSAGE_TYPE_ERROR. Use
g_dbus_message_to_gerror() to transcode this to a #GError.

See this [server][gdbus-server] and [client][gdbus-unix-fd-client]
for an example of how to use this low-level API to send and receive
UNIX file descriptors.
*/
func (self *TraitDBusConnection) SendMessageWithReplyFinish(res *C.GAsyncResult) (return__ *DBusMessage, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GDBusMessage
	__cgo__return__ = C.g_dbus_connection_send_message_with_reply_finish(self.CPointer, res, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewDBusMessageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Synchronously sends @message to the peer represented by @connection
and blocks the calling thread until a reply is received or the
timeout is reached. See g_dbus_connection_send_message_with_reply()
for the asynchronous version of this method.

Unless @flags contain the
%G_DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL flag, the serial number
will be assigned by @connection and set on @message via
g_dbus_message_set_serial(). If @out_serial is not %NULL, then the
serial number used will be written to this location prior to
submitting the message to the underlying transport.

If @connection is closed then the operation will fail with
%G_IO_ERROR_CLOSED. If @cancellable is canceled, the operation will
fail with %G_IO_ERROR_CANCELLED. If @message is not well-formed,
the operation fails with %G_IO_ERROR_INVALID_ARGUMENT.

Note that @error is only set if a local in-process error
occurred. That is to say that the returned #GDBusMessage object may
be of type %G_DBUS_MESSAGE_TYPE_ERROR. Use
g_dbus_message_to_gerror() to transcode this to a #GError.

See this [server][gdbus-server] and [client][gdbus-unix-fd-client]
for an example of how to use this low-level API to send and receive
UNIX file descriptors.

Note that @message must be unlocked, unless @flags contain the
%G_DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL flag.
*/
func (self *TraitDBusConnection) SendMessageWithReplySync(message IsDBusMessage, flags C.GDBusSendMessageFlags, timeout_msec int, cancellable IsCancellable) (out_serial uint32, return__ *DBusMessage, __err__ error) {
	var __cgo__out_serial C.guint32
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GDBusMessage
	__cgo__return__ = C.g_dbus_connection_send_message_with_reply_sync(self.CPointer, message.GetDBusMessagePointer(), flags, C.gint(timeout_msec), &__cgo__out_serial, cancellable.GetCancellablePointer(), &__cgo_error__)
	out_serial = uint32(__cgo__out_serial)
	if __cgo__return__ != nil {
		return__ = NewDBusMessageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Sets whether the process should be terminated when @connection is
closed by the remote peer. See #GDBusConnection:exit-on-close for
more details.

Note that this function should be used with care. Most modern UNIX
desktops tie the notion of a user session the session bus, and expect
all of a users applications to quit when their bus connection goes away.
If you are setting @exit_on_close to %FALSE for the shared session
bus connection, you should make sure that your application exits
when the user session ends.
*/
func (self *TraitDBusConnection) SetExitOnClose(exit_on_close bool) {
	__cgo__exit_on_close := C.gboolean(0)
	if exit_on_close {
		__cgo__exit_on_close = C.gboolean(1)
	}
	C.g_dbus_connection_set_exit_on_close(self.CPointer, __cgo__exit_on_close)
	return
}

/*
Subscribes to signals on @connection and invokes @callback with a whenever
the signal is received. Note that @callback will be invoked in the
[thread-default main context][g-main-context-push-thread-default]
of the thread you are calling this method from.

If @connection is not a message bus connection, @sender must be
%NULL.

If @sender is a well-known name note that @callback is invoked with
the unique name for the owner of @sender, not the well-known name
as one would expect. This is because the message bus rewrites the
name. As such, to avoid certain race conditions, users should be
tracking the name owner of the well-known name and use that when
processing the received signal.

If one of %G_DBUS_SIGNAL_FLAGS_MATCH_ARG0_NAMESPACE or
%G_DBUS_SIGNAL_FLAGS_MATCH_ARG0_PATH are given, @arg0 is
interpreted as part of a namespace or path.  The first argument
of a signal is matched against that part as specified by D-Bus.
*/
func (self *TraitDBusConnection) SignalSubscribe(sender string, interface_name string, member string, object_path string, arg0 string, flags C.GDBusSignalFlags, callback C.GDBusSignalCallback, user_data unsafe.Pointer, user_data_free_func C.GDestroyNotify) (return__ uint) {
	__cgo__sender := (*C.gchar)(unsafe.Pointer(C.CString(sender)))
	__cgo__interface_name := (*C.gchar)(unsafe.Pointer(C.CString(interface_name)))
	__cgo__member := (*C.gchar)(unsafe.Pointer(C.CString(member)))
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	__cgo__arg0 := (*C.gchar)(unsafe.Pointer(C.CString(arg0)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_dbus_connection_signal_subscribe(self.CPointer, __cgo__sender, __cgo__interface_name, __cgo__member, __cgo__object_path, __cgo__arg0, flags, callback, (C.gpointer)(user_data), user_data_free_func)
	C.free(unsafe.Pointer(__cgo__sender))
	C.free(unsafe.Pointer(__cgo__interface_name))
	C.free(unsafe.Pointer(__cgo__member))
	C.free(unsafe.Pointer(__cgo__object_path))
	C.free(unsafe.Pointer(__cgo__arg0))
	return__ = uint(__cgo__return__)
	return
}

/*
Unsubscribes from signals.
*/
func (self *TraitDBusConnection) SignalUnsubscribe(subscription_id uint) {
	C.g_dbus_connection_signal_unsubscribe(self.CPointer, C.guint(subscription_id))
	return
}

/*
If @connection was created with
%G_DBUS_CONNECTION_FLAGS_DELAY_MESSAGE_PROCESSING, this method
starts processing messages. Does nothing on if @connection wasn't
created with this flag or if the method has already been called.
*/
func (self *TraitDBusConnection) StartMessageProcessing() {
	C.g_dbus_connection_start_message_processing(self.CPointer)
	return
}

/*
Reverses the effect of a previous call to
g_dbus_connection_export_action_group().

It is an error to call this function with an ID that wasn't returned
from g_dbus_connection_export_action_group() or to call it with the
same ID more than once.
*/
func (self *TraitDBusConnection) UnexportActionGroup(export_id uint) {
	C.g_dbus_connection_unexport_action_group(self.CPointer, C.guint(export_id))
	return
}

/*
Reverses the effect of a previous call to
g_dbus_connection_export_menu_model().

It is an error to call this function with an ID that wasn't returned
from g_dbus_connection_export_menu_model() or to call it with the
same ID more than once.
*/
func (self *TraitDBusConnection) UnexportMenuModel(export_id uint) {
	C.g_dbus_connection_unexport_menu_model(self.CPointer, C.guint(export_id))
	return
}

/*
Unregisters an object.
*/
func (self *TraitDBusConnection) UnregisterObject(registration_id uint) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_connection_unregister_object(self.CPointer, C.guint(registration_id))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Unregisters a subtree.
*/
func (self *TraitDBusConnection) UnregisterSubtree(registration_id uint) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_connection_unregister_subtree(self.CPointer, C.guint(registration_id))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

type TraitDBusInterfaceSkeleton struct{ CPointer *C.GDBusInterfaceSkeleton }
type IsDBusInterfaceSkeleton interface {
	GetDBusInterfaceSkeletonPointer() *C.GDBusInterfaceSkeleton
}

func (self *TraitDBusInterfaceSkeleton) GetDBusInterfaceSkeletonPointer() *C.GDBusInterfaceSkeleton {
	return self.CPointer
}
func NewTraitDBusInterfaceSkeleton(p unsafe.Pointer) *TraitDBusInterfaceSkeleton {
	return &TraitDBusInterfaceSkeleton{(*C.GDBusInterfaceSkeleton)(p)}
}

/*
Exports @interface_ at @object_path on @connection.

This can be called multiple times to export the same @interface_
onto multiple connections however the @object_path provided must be
the same for all connections.

Use g_dbus_interface_skeleton_unexport() to unexport the object.
*/
func (self *TraitDBusInterfaceSkeleton) Export(connection IsDBusConnection, object_path string) (return__ bool, __err__ error) {
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_interface_skeleton_export(self.CPointer, connection.GetDBusConnectionPointer(), __cgo__object_path, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__object_path))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
If @interface_ has outstanding changes, request for these changes to be
emitted immediately.

For example, an exported D-Bus interface may queue up property
changes and emit the
`org.freedesktop.DBus.Properties::Propert``
signal later (e.g. in an idle handler). This technique is useful
for collapsing multiple property changes into one.
*/
func (self *TraitDBusInterfaceSkeleton) Flush() {
	C.g_dbus_interface_skeleton_flush(self.CPointer)
	return
}

/*
Gets the first connection that @interface_ is exported on, if any.
*/
func (self *TraitDBusInterfaceSkeleton) GetConnection() (return__ *DBusConnection) {
	var __cgo__return__ *C.GDBusConnection
	__cgo__return__ = C.g_dbus_interface_skeleton_get_connection(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDBusConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets a list of the connections that @interface_ is exported on.
*/
func (self *TraitDBusInterfaceSkeleton) GetConnections() (return__ *C.GList) {
	return__ = C.g_dbus_interface_skeleton_get_connections(self.CPointer)
	return
}

/*
Gets the #GDBusInterfaceSkeletonFlags that describes what the behavior
of @interface_
*/
func (self *TraitDBusInterfaceSkeleton) GetFlags() (return__ C.GDBusInterfaceSkeletonFlags) {
	return__ = C.g_dbus_interface_skeleton_get_flags(self.CPointer)
	return
}

/*
Gets D-Bus introspection information for the D-Bus interface
implemented by @interface_.
*/
func (self *TraitDBusInterfaceSkeleton) GetInfo() (return__ *C.GDBusInterfaceInfo) {
	return__ = C.g_dbus_interface_skeleton_get_info(self.CPointer)
	return
}

/*
Gets the object path that @interface_ is exported on, if any.
*/
func (self *TraitDBusInterfaceSkeleton) GetObjectPath() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_interface_skeleton_get_object_path(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets all D-Bus properties for @interface_.
*/
func (self *TraitDBusInterfaceSkeleton) GetProperties() (return__ *C.GVariant) {
	return__ = C.g_dbus_interface_skeleton_get_properties(self.CPointer)
	return
}

/*
Gets the interface vtable for the D-Bus interface implemented by
@interface_. The returned function pointers should expect @interface_
itself to be passed as @user_data.
*/
func (self *TraitDBusInterfaceSkeleton) GetVtable() (return__ *C.GDBusInterfaceVTable) {
	return__ = C.g_dbus_interface_skeleton_get_vtable(self.CPointer)
	return
}

/*
Checks if @interface_ is exported on @connection.
*/
func (self *TraitDBusInterfaceSkeleton) HasConnection(connection IsDBusConnection) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_interface_skeleton_has_connection(self.CPointer, connection.GetDBusConnectionPointer())
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets flags describing what the behavior of @skeleton should be.
*/
func (self *TraitDBusInterfaceSkeleton) SetFlags(flags C.GDBusInterfaceSkeletonFlags) {
	C.g_dbus_interface_skeleton_set_flags(self.CPointer, flags)
	return
}

/*
Stops exporting @interface_ on all connections it is exported on.

To unexport @interface_ from only a single connection, use
g_dbus_interface_skeleton_unexport_from_connection()
*/
func (self *TraitDBusInterfaceSkeleton) Unexport() {
	C.g_dbus_interface_skeleton_unexport(self.CPointer)
	return
}

/*
Stops exporting @interface_ on @connection.

To stop exporting on all connections the interface is exported on,
use g_dbus_interface_skeleton_unexport().
*/
func (self *TraitDBusInterfaceSkeleton) UnexportFromConnection(connection IsDBusConnection) {
	C.g_dbus_interface_skeleton_unexport_from_connection(self.CPointer, connection.GetDBusConnectionPointer())
	return
}

type TraitDBusMenuModel struct{ CPointer *C.GDBusMenuModel }
type IsDBusMenuModel interface {
	GetDBusMenuModelPointer() *C.GDBusMenuModel
}

func (self *TraitDBusMenuModel) GetDBusMenuModelPointer() *C.GDBusMenuModel {
	return self.CPointer
}
func NewTraitDBusMenuModel(p unsafe.Pointer) *TraitDBusMenuModel {
	return &TraitDBusMenuModel{(*C.GDBusMenuModel)(p)}
}

type TraitDBusMessage struct{ CPointer *C.GDBusMessage }
type IsDBusMessage interface {
	GetDBusMessagePointer() *C.GDBusMessage
}

func (self *TraitDBusMessage) GetDBusMessagePointer() *C.GDBusMessage {
	return self.CPointer
}
func NewTraitDBusMessage(p unsafe.Pointer) *TraitDBusMessage {
	return &TraitDBusMessage{(*C.GDBusMessage)(p)}
}

/*
Copies @message. The copy is a deep copy and the returned
#GDBusMessage is completely identical except that it is guaranteed
to not be locked.

This operation can fail if e.g. @message contains file descriptors
and the per-process or system-wide open files limit is reached.
*/
func (self *TraitDBusMessage) Copy() (return__ *DBusMessage, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GDBusMessage
	__cgo__return__ = C.g_dbus_message_copy(self.CPointer, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewDBusMessageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Convenience to get the first item in the body of @message.
*/
func (self *TraitDBusMessage) GetArg0() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_message_get_arg0(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the body of a message.
*/
func (self *TraitDBusMessage) GetBody() (return__ *C.GVariant) {
	return__ = C.g_dbus_message_get_body(self.CPointer)
	return
}

/*
Gets the byte order of @message.
*/
func (self *TraitDBusMessage) GetByteOrder() (return__ C.GDBusMessageByteOrder) {
	return__ = C.g_dbus_message_get_byte_order(self.CPointer)
	return
}

/*
Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_DESTINATION header field.
*/
func (self *TraitDBusMessage) GetDestination() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_message_get_destination(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME header field.
*/
func (self *TraitDBusMessage) GetErrorName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_message_get_error_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the flags for @message.
*/
func (self *TraitDBusMessage) GetFlags() (return__ C.GDBusMessageFlags) {
	return__ = C.g_dbus_message_get_flags(self.CPointer)
	return
}

/*
Gets a header field on @message.
*/
func (self *TraitDBusMessage) GetHeader(header_field C.GDBusMessageHeaderField) (return__ *C.GVariant) {
	return__ = C.g_dbus_message_get_header(self.CPointer, header_field)
	return
}

/*
Gets an array of all header fields on @message that are set.
*/
func (self *TraitDBusMessage) GetHeaderFields() (return__ *C.guchar) {
	return__ = C.g_dbus_message_get_header_fields(self.CPointer)
	return
}

/*
Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_INTERFACE header field.
*/
func (self *TraitDBusMessage) GetInterface() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_message_get_interface(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Checks whether @message is locked. To monitor changes to this
value, conncet to the #GObject::notify signal to listen for changes
on the #GDBusMessage:locked property.
*/
func (self *TraitDBusMessage) GetLocked() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_message_get_locked(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_MEMBER header field.
*/
func (self *TraitDBusMessage) GetMember() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_message_get_member(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the type of @message.
*/
func (self *TraitDBusMessage) GetMessageType() (return__ C.GDBusMessageType) {
	return__ = C.g_dbus_message_get_message_type(self.CPointer)
	return
}

/*
Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS header field.
*/
func (self *TraitDBusMessage) GetNumUnixFds() (return__ uint32) {
	var __cgo__return__ C.guint32
	__cgo__return__ = C.g_dbus_message_get_num_unix_fds(self.CPointer)
	return__ = uint32(__cgo__return__)
	return
}

/*
Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_PATH header field.
*/
func (self *TraitDBusMessage) GetPath() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_message_get_path(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL header field.
*/
func (self *TraitDBusMessage) GetReplySerial() (return__ uint32) {
	var __cgo__return__ C.guint32
	__cgo__return__ = C.g_dbus_message_get_reply_serial(self.CPointer)
	return__ = uint32(__cgo__return__)
	return
}

/*
Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_SENDER header field.
*/
func (self *TraitDBusMessage) GetSender() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_message_get_sender(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the serial for @message.
*/
func (self *TraitDBusMessage) GetSerial() (return__ uint32) {
	var __cgo__return__ C.guint32
	__cgo__return__ = C.g_dbus_message_get_serial(self.CPointer)
	return__ = uint32(__cgo__return__)
	return
}

/*
Convenience getter for the %G_DBUS_MESSAGE_HEADER_FIELD_SIGNATURE header field.
*/
func (self *TraitDBusMessage) GetSignature() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_message_get_signature(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the UNIX file descriptors associated with @message, if any.

This method is only available on UNIX.
*/
func (self *TraitDBusMessage) GetUnixFdList() (return__ *UnixFDList) {
	var __cgo__return__ *C.GUnixFDList
	__cgo__return__ = C.g_dbus_message_get_unix_fd_list(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewUnixFDListFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
If @message is locked, does nothing. Otherwise locks the message.
*/
func (self *TraitDBusMessage) Lock() {
	C.g_dbus_message_lock(self.CPointer)
	return
}

// g_dbus_message_new_method_error is not generated due to varargs

/*
Creates a new #GDBusMessage that is an error reply to @method_call_message.
*/
func (self *TraitDBusMessage) NewMethodErrorLiteral(error_name string, error_message string) (return__ *DBusMessage) {
	__cgo__error_name := (*C.gchar)(unsafe.Pointer(C.CString(error_name)))
	__cgo__error_message := (*C.gchar)(unsafe.Pointer(C.CString(error_message)))
	var __cgo__return__ *C.GDBusMessage
	__cgo__return__ = C.g_dbus_message_new_method_error_literal(self.CPointer, __cgo__error_name, __cgo__error_message)
	C.free(unsafe.Pointer(__cgo__error_name))
	C.free(unsafe.Pointer(__cgo__error_message))
	if __cgo__return__ != nil {
		return__ = NewDBusMessageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// g_dbus_message_new_method_error_valist is not generated due to varargs

/*
Creates a new #GDBusMessage that is a reply to @method_call_message.
*/
func (self *TraitDBusMessage) NewMethodReply() (return__ *DBusMessage) {
	var __cgo__return__ *C.GDBusMessage
	__cgo__return__ = C.g_dbus_message_new_method_reply(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDBusMessageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Produces a human-readable multi-line description of @message.

The contents of the description has no ABI guarantees, the contents
and formatting is subject to change at any time. Typical output
looks something like this:
|[
Flags:   none
Version: 0
Serial:  4
Headers:
  path -> objectpath '/org/gtk/GDBus/TestObject'
  interface -> 'org.gtk.GDBus.TestInterface'
  member -> 'GimmeStdout'
  destination -> ':1.146'
Body: ()
UNIX File Descriptors:
  (none)
]|
or
|[
Flags:   no-reply-expected
Version: 0
Serial:  477
Headers:
  reply-serial -> uint32 4
  destination -> ':1.159'
  sender -> ':1.146'
  num-unix-fds -> uint32 1
Body: ()
UNIX File Descriptors:
  fd 12: dev=0:10,mode=020620,ino=5,uid=500,gid=5,rdev=136:2,size=0,atime=1273085037,mtime=1273085851,ctime=1272982635
]|
*/
func (self *TraitDBusMessage) Print(indent uint) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_message_print(self.CPointer, C.guint(indent))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Sets the body @message. As a side-effect the
%G_DBUS_MESSAGE_HEADER_FIELD_SIGNATURE header field is set to the
type string of @body (or cleared if @body is %NULL).

If @body is floating, @message assumes ownership of @body.
*/
func (self *TraitDBusMessage) SetBody(body *C.GVariant) {
	C.g_dbus_message_set_body(self.CPointer, body)
	return
}

/*
Sets the byte order of @message.
*/
func (self *TraitDBusMessage) SetByteOrder(byte_order C.GDBusMessageByteOrder) {
	C.g_dbus_message_set_byte_order(self.CPointer, byte_order)
	return
}

/*
Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_DESTINATION header field.
*/
func (self *TraitDBusMessage) SetDestination(value string) {
	__cgo__value := (*C.gchar)(unsafe.Pointer(C.CString(value)))
	C.g_dbus_message_set_destination(self.CPointer, __cgo__value)
	C.free(unsafe.Pointer(__cgo__value))
	return
}

/*
Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME header field.
*/
func (self *TraitDBusMessage) SetErrorName(value string) {
	__cgo__value := (*C.gchar)(unsafe.Pointer(C.CString(value)))
	C.g_dbus_message_set_error_name(self.CPointer, __cgo__value)
	C.free(unsafe.Pointer(__cgo__value))
	return
}

/*
Sets the flags to set on @message.
*/
func (self *TraitDBusMessage) SetFlags(flags C.GDBusMessageFlags) {
	C.g_dbus_message_set_flags(self.CPointer, flags)
	return
}

/*
Sets a header field on @message.

If @value is floating, @message assumes ownership of @value.
*/
func (self *TraitDBusMessage) SetHeader(header_field C.GDBusMessageHeaderField, value *C.GVariant) {
	C.g_dbus_message_set_header(self.CPointer, header_field, value)
	return
}

/*
Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_INTERFACE header field.
*/
func (self *TraitDBusMessage) SetInterface(value string) {
	__cgo__value := (*C.gchar)(unsafe.Pointer(C.CString(value)))
	C.g_dbus_message_set_interface(self.CPointer, __cgo__value)
	C.free(unsafe.Pointer(__cgo__value))
	return
}

/*
Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_MEMBER header field.
*/
func (self *TraitDBusMessage) SetMember(value string) {
	__cgo__value := (*C.gchar)(unsafe.Pointer(C.CString(value)))
	C.g_dbus_message_set_member(self.CPointer, __cgo__value)
	C.free(unsafe.Pointer(__cgo__value))
	return
}

/*
Sets @message to be of @type.
*/
func (self *TraitDBusMessage) SetMessageType(type_ C.GDBusMessageType) {
	C.g_dbus_message_set_message_type(self.CPointer, type_)
	return
}

/*
Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS header field.
*/
func (self *TraitDBusMessage) SetNumUnixFds(value uint32) {
	C.g_dbus_message_set_num_unix_fds(self.CPointer, C.guint32(value))
	return
}

/*
Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_PATH header field.
*/
func (self *TraitDBusMessage) SetPath(value string) {
	__cgo__value := (*C.gchar)(unsafe.Pointer(C.CString(value)))
	C.g_dbus_message_set_path(self.CPointer, __cgo__value)
	C.free(unsafe.Pointer(__cgo__value))
	return
}

/*
Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL header field.
*/
func (self *TraitDBusMessage) SetReplySerial(value uint32) {
	C.g_dbus_message_set_reply_serial(self.CPointer, C.guint32(value))
	return
}

/*
Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_SENDER header field.
*/
func (self *TraitDBusMessage) SetSender(value string) {
	__cgo__value := (*C.gchar)(unsafe.Pointer(C.CString(value)))
	C.g_dbus_message_set_sender(self.CPointer, __cgo__value)
	C.free(unsafe.Pointer(__cgo__value))
	return
}

/*
Sets the serial for @message.
*/
func (self *TraitDBusMessage) SetSerial(serial uint32) {
	C.g_dbus_message_set_serial(self.CPointer, C.guint32(serial))
	return
}

/*
Convenience setter for the %G_DBUS_MESSAGE_HEADER_FIELD_SIGNATURE header field.
*/
func (self *TraitDBusMessage) SetSignature(value string) {
	__cgo__value := (*C.gchar)(unsafe.Pointer(C.CString(value)))
	C.g_dbus_message_set_signature(self.CPointer, __cgo__value)
	C.free(unsafe.Pointer(__cgo__value))
	return
}

/*
Sets the UNIX file descriptors associated with @message. As a
side-effect the %G_DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS header
field is set to the number of fds in @fd_list (or cleared if
@fd_list is %NULL).

This method is only available on UNIX.
*/
func (self *TraitDBusMessage) SetUnixFdList(fd_list IsUnixFDList) {
	C.g_dbus_message_set_unix_fd_list(self.CPointer, fd_list.GetUnixFDListPointer())
	return
}

/*
Serializes @message to a blob. The byte order returned by
g_dbus_message_get_byte_order() will be used.
*/
func (self *TraitDBusMessage) ToBlob(capabilities C.GDBusCapabilityFlags) (out_size int64, return__ []byte, __err__ error) {
	var __cgo__out_size C.gsize
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.guchar
	__cgo__return__ = C.g_dbus_message_to_blob(self.CPointer, &__cgo__out_size, capabilities, &__cgo_error__)
	out_size = int64(__cgo__out_size)
	defer func() { return__ = C.GoBytes(unsafe.Pointer(__cgo__return__), C.int(out_size)) }()
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
If @message is not of type %G_DBUS_MESSAGE_TYPE_ERROR does
nothing and returns %FALSE.

Otherwise this method encodes the error in @message as a #GError
using g_dbus_error_set_dbus_error() using the information in the
%G_DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME header field of @message as
well as the first string item in @message's body.
*/
func (self *TraitDBusMessage) ToGerror() (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_message_to_gerror(self.CPointer, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

type TraitDBusMethodInvocation struct{ CPointer *C.GDBusMethodInvocation }
type IsDBusMethodInvocation interface {
	GetDBusMethodInvocationPointer() *C.GDBusMethodInvocation
}

func (self *TraitDBusMethodInvocation) GetDBusMethodInvocationPointer() *C.GDBusMethodInvocation {
	return self.CPointer
}
func NewTraitDBusMethodInvocation(p unsafe.Pointer) *TraitDBusMethodInvocation {
	return &TraitDBusMethodInvocation{(*C.GDBusMethodInvocation)(p)}
}

/*
Gets the #GDBusConnection the method was invoked on.
*/
func (self *TraitDBusMethodInvocation) GetConnection() (return__ *DBusConnection) {
	var __cgo__return__ *C.GDBusConnection
	__cgo__return__ = C.g_dbus_method_invocation_get_connection(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDBusConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the name of the D-Bus interface the method was invoked on.

If this method call is a property Get, Set or GetAll call that has
been redirected to the method call handler then
"org.freedesktop.DBus.Properties" will be returned.  See
#GDBusInterfaceVTable for more information.
*/
func (self *TraitDBusMethodInvocation) GetInterfaceName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_method_invocation_get_interface_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the #GDBusMessage for the method invocation. This is useful if
you need to use low-level protocol features, such as UNIX file
descriptor passing, that cannot be properly expressed in the
#GVariant API.

See this [server][gdbus-server] and [client][gdbus-unix-fd-client]
for an example of how to use this low-level API to send and receive
UNIX file descriptors.
*/
func (self *TraitDBusMethodInvocation) GetMessage() (return__ *DBusMessage) {
	var __cgo__return__ *C.GDBusMessage
	__cgo__return__ = C.g_dbus_method_invocation_get_message(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDBusMessageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets information about the method call, if any.

If this method invocation is a property Get, Set or GetAll call that
has been redirected to the method call handler then %NULL will be
returned.  See g_dbus_method_invocation_get_property_info() and
#GDBusInterfaceVTable for more information.
*/
func (self *TraitDBusMethodInvocation) GetMethodInfo() (return__ *C.GDBusMethodInfo) {
	return__ = C.g_dbus_method_invocation_get_method_info(self.CPointer)
	return
}

/*
Gets the name of the method that was invoked.
*/
func (self *TraitDBusMethodInvocation) GetMethodName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_method_invocation_get_method_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the object path the method was invoked on.
*/
func (self *TraitDBusMethodInvocation) GetObjectPath() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_method_invocation_get_object_path(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the parameters of the method invocation. If there are no input
parameters then this will return a GVariant with 0 children rather than NULL.
*/
func (self *TraitDBusMethodInvocation) GetParameters() (return__ *C.GVariant) {
	return__ = C.g_dbus_method_invocation_get_parameters(self.CPointer)
	return
}

/*
Gets information about the property that this method call is for, if
any.

This will only be set in the case of an invocation in response to a
property Get or Set call that has been directed to the method call
handler for an object on account of its property_get() or
property_set() vtable pointers being unset.

See #GDBusInterfaceVTable for more information.

If the call was GetAll, %NULL will be returned.
*/
func (self *TraitDBusMethodInvocation) GetPropertyInfo() (return__ *C.GDBusPropertyInfo) {
	return__ = C.g_dbus_method_invocation_get_property_info(self.CPointer)
	return
}

/*
Gets the bus name that invoked the method.
*/
func (self *TraitDBusMethodInvocation) GetSender() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_method_invocation_get_sender(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the @user_data #gpointer passed to g_dbus_connection_register_object().
*/
func (self *TraitDBusMethodInvocation) GetUserData() (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_dbus_method_invocation_get_user_data(self.CPointer)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Finishes handling a D-Bus method call by returning an error.

This method will free @invocation, you cannot use it afterwards.
*/
func (self *TraitDBusMethodInvocation) ReturnDbusError(error_name string, error_message string) {
	__cgo__error_name := (*C.gchar)(unsafe.Pointer(C.CString(error_name)))
	__cgo__error_message := (*C.gchar)(unsafe.Pointer(C.CString(error_message)))
	C.g_dbus_method_invocation_return_dbus_error(self.CPointer, __cgo__error_name, __cgo__error_message)
	C.free(unsafe.Pointer(__cgo__error_name))
	C.free(unsafe.Pointer(__cgo__error_message))
	return
}

// g_dbus_method_invocation_return_error is not generated due to varargs

/*
Like g_dbus_method_invocation_return_error() but without printf()-style formatting.

This method will free @invocation, you cannot use it afterwards.
*/
func (self *TraitDBusMethodInvocation) ReturnErrorLiteral(domain C.GQuark, code int, message string) {
	__cgo__message := (*C.gchar)(unsafe.Pointer(C.CString(message)))
	C.g_dbus_method_invocation_return_error_literal(self.CPointer, domain, C.gint(code), __cgo__message)
	C.free(unsafe.Pointer(__cgo__message))
	return
}

// g_dbus_method_invocation_return_error_valist is not generated due to varargs

/*
Like g_dbus_method_invocation_return_error() but takes a #GError
instead of the error domain, error code and message.

This method will free @invocation, you cannot use it afterwards.
*/
func (self *TraitDBusMethodInvocation) ReturnGerror(error_ *C.GError) {
	C.g_dbus_method_invocation_return_gerror(self.CPointer, error_)
	return
}

/*
Finishes handling a D-Bus method call by returning @parameters.
If the @parameters GVariant is floating, it is consumed.

It is an error if @parameters is not of the right format.

This method will free @invocation, you cannot use it afterwards.
*/
func (self *TraitDBusMethodInvocation) ReturnValue(parameters *C.GVariant) {
	C.g_dbus_method_invocation_return_value(self.CPointer, parameters)
	return
}

/*
Like g_dbus_method_invocation_return_value() but also takes a #GUnixFDList.

This method is only available on UNIX.

This method will free @invocation, you cannot use it afterwards.
*/
func (self *TraitDBusMethodInvocation) ReturnValueWithUnixFdList(parameters *C.GVariant, fd_list IsUnixFDList) {
	C.g_dbus_method_invocation_return_value_with_unix_fd_list(self.CPointer, parameters, fd_list.GetUnixFDListPointer())
	return
}

/*
Like g_dbus_method_invocation_return_gerror() but takes ownership
of @error so the caller does not need to free it.

This method will free @invocation, you cannot use it afterwards.
*/
func (self *TraitDBusMethodInvocation) TakeError(error_ *C.GError) {
	C.g_dbus_method_invocation_take_error(self.CPointer, error_)
	return
}

type TraitDBusObjectManagerClient struct{ CPointer *C.GDBusObjectManagerClient }
type IsDBusObjectManagerClient interface {
	GetDBusObjectManagerClientPointer() *C.GDBusObjectManagerClient
}

func (self *TraitDBusObjectManagerClient) GetDBusObjectManagerClientPointer() *C.GDBusObjectManagerClient {
	return self.CPointer
}
func NewTraitDBusObjectManagerClient(p unsafe.Pointer) *TraitDBusObjectManagerClient {
	return &TraitDBusObjectManagerClient{(*C.GDBusObjectManagerClient)(p)}
}

/*
Gets the #GDBusConnection used by @manager.
*/
func (self *TraitDBusObjectManagerClient) GetConnection() (return__ *DBusConnection) {
	var __cgo__return__ *C.GDBusConnection
	__cgo__return__ = C.g_dbus_object_manager_client_get_connection(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDBusConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the flags that @manager was constructed with.
*/
func (self *TraitDBusObjectManagerClient) GetFlags() (return__ C.GDBusObjectManagerClientFlags) {
	return__ = C.g_dbus_object_manager_client_get_flags(self.CPointer)
	return
}

/*
Gets the name that @manager is for, or %NULL if not a message bus
connection.
*/
func (self *TraitDBusObjectManagerClient) GetName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_object_manager_client_get_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
The unique name that owns the name that @manager is for or %NULL if
no-one currently owns that name. You can connect to the
#GObject::notify signal to track changes to the
#GDBusObjectManagerClient:name-owner property.
*/
func (self *TraitDBusObjectManagerClient) GetNameOwner() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_object_manager_client_get_name_owner(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

type TraitDBusObjectManagerServer struct{ CPointer *C.GDBusObjectManagerServer }
type IsDBusObjectManagerServer interface {
	GetDBusObjectManagerServerPointer() *C.GDBusObjectManagerServer
}

func (self *TraitDBusObjectManagerServer) GetDBusObjectManagerServerPointer() *C.GDBusObjectManagerServer {
	return self.CPointer
}
func NewTraitDBusObjectManagerServer(p unsafe.Pointer) *TraitDBusObjectManagerServer {
	return &TraitDBusObjectManagerServer{(*C.GDBusObjectManagerServer)(p)}
}

/*
Exports @object on @manager.

If there is already a #GDBusObject exported at the object path,
then the old object is removed.

The object path for @object must be in the hierarchy rooted by the
object path for @manager.

Note that @manager will take a reference on @object for as long as
it is exported.
*/
func (self *TraitDBusObjectManagerServer) Export(object IsDBusObjectSkeleton) {
	C.g_dbus_object_manager_server_export(self.CPointer, object.GetDBusObjectSkeletonPointer())
	return
}

/*
Like g_dbus_object_manager_server_export() but appends a string of
the form _N (with N being a natural number) to @object's object path
if an object with the given path already exists. As such, the
#GDBusObjectProxy:g-object-path property of @object may be modified.
*/
func (self *TraitDBusObjectManagerServer) ExportUniquely(object IsDBusObjectSkeleton) {
	C.g_dbus_object_manager_server_export_uniquely(self.CPointer, object.GetDBusObjectSkeletonPointer())
	return
}

/*
Gets the #GDBusConnection used by @manager.
*/
func (self *TraitDBusObjectManagerServer) GetConnection() (return__ *DBusConnection) {
	var __cgo__return__ *C.GDBusConnection
	__cgo__return__ = C.g_dbus_object_manager_server_get_connection(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDBusConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns whether @object is currently exported on @manager.
*/
func (self *TraitDBusObjectManagerServer) IsExported(object IsDBusObjectSkeleton) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_object_manager_server_is_exported(self.CPointer, object.GetDBusObjectSkeletonPointer())
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Exports all objects managed by @manager on @connection. If
@connection is %NULL, stops exporting objects.
*/
func (self *TraitDBusObjectManagerServer) SetConnection(connection IsDBusConnection) {
	C.g_dbus_object_manager_server_set_connection(self.CPointer, connection.GetDBusConnectionPointer())
	return
}

/*
If @manager has an object at @path, removes the object. Otherwise
does nothing.

Note that @object_path must be in the hierarchy rooted by the
object path for @manager.
*/
func (self *TraitDBusObjectManagerServer) Unexport(object_path string) (return__ bool) {
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_object_manager_server_unexport(self.CPointer, __cgo__object_path)
	C.free(unsafe.Pointer(__cgo__object_path))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

type TraitDBusObjectProxy struct{ CPointer *C.GDBusObjectProxy }
type IsDBusObjectProxy interface {
	GetDBusObjectProxyPointer() *C.GDBusObjectProxy
}

func (self *TraitDBusObjectProxy) GetDBusObjectProxyPointer() *C.GDBusObjectProxy {
	return self.CPointer
}
func NewTraitDBusObjectProxy(p unsafe.Pointer) *TraitDBusObjectProxy {
	return &TraitDBusObjectProxy{(*C.GDBusObjectProxy)(p)}
}

/*
Gets the connection that @proxy is for.
*/
func (self *TraitDBusObjectProxy) GetConnection() (return__ *DBusConnection) {
	var __cgo__return__ *C.GDBusConnection
	__cgo__return__ = C.g_dbus_object_proxy_get_connection(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDBusConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

type TraitDBusObjectSkeleton struct{ CPointer *C.GDBusObjectSkeleton }
type IsDBusObjectSkeleton interface {
	GetDBusObjectSkeletonPointer() *C.GDBusObjectSkeleton
}

func (self *TraitDBusObjectSkeleton) GetDBusObjectSkeletonPointer() *C.GDBusObjectSkeleton {
	return self.CPointer
}
func NewTraitDBusObjectSkeleton(p unsafe.Pointer) *TraitDBusObjectSkeleton {
	return &TraitDBusObjectSkeleton{(*C.GDBusObjectSkeleton)(p)}
}

/*
Adds @interface_ to @object.

If @object already contains a #GDBusInterfaceSkeleton with the same
interface name, it is removed before @interface_ is added.

Note that @object takes its own reference on @interface_ and holds
it until removed.
*/
func (self *TraitDBusObjectSkeleton) AddInterface(interface_ IsDBusInterfaceSkeleton) {
	C.g_dbus_object_skeleton_add_interface(self.CPointer, interface_.GetDBusInterfaceSkeletonPointer())
	return
}

/*
This method simply calls g_dbus_interface_skeleton_flush() on all
interfaces belonging to @object. See that method for when flushing
is useful.
*/
func (self *TraitDBusObjectSkeleton) Flush() {
	C.g_dbus_object_skeleton_flush(self.CPointer)
	return
}

/*
Removes @interface_ from @object.
*/
func (self *TraitDBusObjectSkeleton) RemoveInterface(interface_ IsDBusInterfaceSkeleton) {
	C.g_dbus_object_skeleton_remove_interface(self.CPointer, interface_.GetDBusInterfaceSkeletonPointer())
	return
}

/*
Removes the #GDBusInterface with @interface_name from @object.

If no D-Bus interface of the given interface exists, this function
does nothing.
*/
func (self *TraitDBusObjectSkeleton) RemoveInterfaceByName(interface_name string) {
	__cgo__interface_name := (*C.gchar)(unsafe.Pointer(C.CString(interface_name)))
	C.g_dbus_object_skeleton_remove_interface_by_name(self.CPointer, __cgo__interface_name)
	C.free(unsafe.Pointer(__cgo__interface_name))
	return
}

/*
Sets the object path for @object.
*/
func (self *TraitDBusObjectSkeleton) SetObjectPath(object_path string) {
	__cgo__object_path := (*C.gchar)(unsafe.Pointer(C.CString(object_path)))
	C.g_dbus_object_skeleton_set_object_path(self.CPointer, __cgo__object_path)
	C.free(unsafe.Pointer(__cgo__object_path))
	return
}

type TraitDBusProxy struct{ CPointer *C.GDBusProxy }
type IsDBusProxy interface {
	GetDBusProxyPointer() *C.GDBusProxy
}

func (self *TraitDBusProxy) GetDBusProxyPointer() *C.GDBusProxy {
	return self.CPointer
}
func NewTraitDBusProxy(p unsafe.Pointer) *TraitDBusProxy {
	return &TraitDBusProxy{(*C.GDBusProxy)(p)}
}

/*
Asynchronously invokes the @method_name method on @proxy.

If @method_name contains any dots, then @name is split into interface and
method name parts. This allows using @proxy for invoking methods on
other interfaces.

If the #GDBusConnection associated with @proxy is closed then
the operation will fail with %G_IO_ERROR_CLOSED. If
@cancellable is canceled, the operation will fail with
%G_IO_ERROR_CANCELLED. If @parameters contains a value not
compatible with the D-Bus protocol, the operation fails with
%G_IO_ERROR_INVALID_ARGUMENT.

If the @parameters #GVariant is floating, it is consumed. This allows
convenient 'inline' use of g_variant_new(), e.g.:
|[<!-- language="C" -->
 g_dbus_proxy_call (proxy,
                    "TwoStrings",
                    g_variant_new ("(ss)",
                                   "Thing One",
                                   "Thing Two"),
                    G_DBUS_CALL_FLAGS_NONE,
                    -1,
                    NULL,
                    (GAsyncReadyCallback) two_strings_done,
                    &data);
]|

If @proxy has an expected interface (see
#GDBusProxy:g-interface-info) and @method_name is referenced by it,
then the return value is checked against the return type.

This is an asynchronous method. When the operation is finished,
@callback will be invoked in the
[thread-default main context][g-main-context-push-thread-default]
of the thread you are calling this method from.
You can then call g_dbus_proxy_call_finish() to get the result of
the operation. See g_dbus_proxy_call_sync() for the synchronous
version of this method.

If @callback is %NULL then the D-Bus method call message will be sent with
the %G_DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED flag set.
*/
func (self *TraitDBusProxy) Call(method_name string, parameters *C.GVariant, flags C.GDBusCallFlags, timeout_msec int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__method_name := (*C.gchar)(unsafe.Pointer(C.CString(method_name)))
	C.g_dbus_proxy_call(self.CPointer, __cgo__method_name, parameters, flags, C.gint(timeout_msec), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__method_name))
	return
}

/*
Finishes an operation started with g_dbus_proxy_call().
*/
func (self *TraitDBusProxy) CallFinish(res *C.GAsyncResult) (return__ *C.GVariant, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_dbus_proxy_call_finish(self.CPointer, res, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Synchronously invokes the @method_name method on @proxy.

If @method_name contains any dots, then @name is split into interface and
method name parts. This allows using @proxy for invoking methods on
other interfaces.

If the #GDBusConnection associated with @proxy is disconnected then
the operation will fail with %G_IO_ERROR_CLOSED. If
@cancellable is canceled, the operation will fail with
%G_IO_ERROR_CANCELLED. If @parameters contains a value not
compatible with the D-Bus protocol, the operation fails with
%G_IO_ERROR_INVALID_ARGUMENT.

If the @parameters #GVariant is floating, it is consumed. This allows
convenient 'inline' use of g_variant_new(), e.g.:
|[<!-- language="C" -->
 g_dbus_proxy_call_sync (proxy,
                         "TwoStrings",
                         g_variant_new ("(ss)",
                                        "Thing One",
                                        "Thing Two"),
                         G_DBUS_CALL_FLAGS_NONE,
                         -1,
                         NULL,
                         &error);
]|

The calling thread is blocked until a reply is received. See
g_dbus_proxy_call() for the asynchronous version of this
method.

If @proxy has an expected interface (see
#GDBusProxy:g-interface-info) and @method_name is referenced by it,
then the return value is checked against the return type.
*/
func (self *TraitDBusProxy) CallSync(method_name string, parameters *C.GVariant, flags C.GDBusCallFlags, timeout_msec int, cancellable IsCancellable) (return__ *C.GVariant, __err__ error) {
	__cgo__method_name := (*C.gchar)(unsafe.Pointer(C.CString(method_name)))
	var __cgo_error__ *C.GError
	return__ = C.g_dbus_proxy_call_sync(self.CPointer, __cgo__method_name, parameters, flags, C.gint(timeout_msec), cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__method_name))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Like g_dbus_proxy_call() but also takes a #GUnixFDList object.

This method is only available on UNIX.
*/
func (self *TraitDBusProxy) CallWithUnixFdList(method_name string, parameters *C.GVariant, flags C.GDBusCallFlags, timeout_msec int, fd_list IsUnixFDList, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__method_name := (*C.gchar)(unsafe.Pointer(C.CString(method_name)))
	C.g_dbus_proxy_call_with_unix_fd_list(self.CPointer, __cgo__method_name, parameters, flags, C.gint(timeout_msec), fd_list.GetUnixFDListPointer(), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__method_name))
	return
}

/*
Finishes an operation started with g_dbus_proxy_call_with_unix_fd_list().
*/
func (self *TraitDBusProxy) CallWithUnixFdListFinish(res *C.GAsyncResult) (out_fd_list *UnixFDList, return__ *C.GVariant, __err__ error) {
	var __cgo__out_fd_list *C.GUnixFDList
	var __cgo_error__ *C.GError
	return__ = C.g_dbus_proxy_call_with_unix_fd_list_finish(self.CPointer, &__cgo__out_fd_list, res, &__cgo_error__)
	if __cgo__out_fd_list != nil {
		out_fd_list = NewUnixFDListFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__out_fd_list).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Like g_dbus_proxy_call_sync() but also takes and returns #GUnixFDList objects.

This method is only available on UNIX.
*/
func (self *TraitDBusProxy) CallWithUnixFdListSync(method_name string, parameters *C.GVariant, flags C.GDBusCallFlags, timeout_msec int, fd_list IsUnixFDList, cancellable IsCancellable) (out_fd_list *UnixFDList, return__ *C.GVariant, __err__ error) {
	__cgo__method_name := (*C.gchar)(unsafe.Pointer(C.CString(method_name)))
	var __cgo__out_fd_list *C.GUnixFDList
	var __cgo_error__ *C.GError
	return__ = C.g_dbus_proxy_call_with_unix_fd_list_sync(self.CPointer, __cgo__method_name, parameters, flags, C.gint(timeout_msec), fd_list.GetUnixFDListPointer(), &__cgo__out_fd_list, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__method_name))
	if __cgo__out_fd_list != nil {
		out_fd_list = NewUnixFDListFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__out_fd_list).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Looks up the value for a property from the cache. This call does no
blocking IO.

If @proxy has an expected interface (see
#GDBusProxy:g-interface-info) and @property_name is referenced by
it, then @value is checked against the type of the property.
*/
func (self *TraitDBusProxy) GetCachedProperty(property_name string) (return__ *C.GVariant) {
	__cgo__property_name := (*C.gchar)(unsafe.Pointer(C.CString(property_name)))
	return__ = C.g_dbus_proxy_get_cached_property(self.CPointer, __cgo__property_name)
	C.free(unsafe.Pointer(__cgo__property_name))
	return
}

/*
Gets the names of all cached properties on @proxy.
*/
func (self *TraitDBusProxy) GetCachedPropertyNames() (return__ **C.gchar) {
	return__ = C.g_dbus_proxy_get_cached_property_names(self.CPointer)
	return
}

/*
Gets the connection @proxy is for.
*/
func (self *TraitDBusProxy) GetConnection() (return__ *DBusConnection) {
	var __cgo__return__ *C.GDBusConnection
	__cgo__return__ = C.g_dbus_proxy_get_connection(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDBusConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the timeout to use if -1 (specifying default timeout) is
passed as @timeout_msec in the g_dbus_proxy_call() and
g_dbus_proxy_call_sync() functions.

See the #GDBusProxy:g-default-timeout property for more details.
*/
func (self *TraitDBusProxy) GetDefaultTimeout() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_dbus_proxy_get_default_timeout(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the flags that @proxy was constructed with.
*/
func (self *TraitDBusProxy) GetFlags() (return__ C.GDBusProxyFlags) {
	return__ = C.g_dbus_proxy_get_flags(self.CPointer)
	return
}

/*
Returns the #GDBusInterfaceInfo, if any, specifying the interface
that @proxy conforms to. See the #GDBusProxy:g-interface-info
property for more details.
*/
func (self *TraitDBusProxy) GetInterfaceInfo() (return__ *C.GDBusInterfaceInfo) {
	return__ = C.g_dbus_proxy_get_interface_info(self.CPointer)
	return
}

/*
Gets the D-Bus interface name @proxy is for.
*/
func (self *TraitDBusProxy) GetInterfaceName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_proxy_get_interface_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the name that @proxy was constructed for.
*/
func (self *TraitDBusProxy) GetName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_proxy_get_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
The unique name that owns the name that @proxy is for or %NULL if
no-one currently owns that name. You may connect to the
#GObject::notify signal to track changes to the
#GDBusProxy:g-name-owner property.
*/
func (self *TraitDBusProxy) GetNameOwner() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_proxy_get_name_owner(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the object path @proxy is for.
*/
func (self *TraitDBusProxy) GetObjectPath() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_proxy_get_object_path(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
If @value is not %NULL, sets the cached value for the property with
name @property_name to the value in @value.

If @value is %NULL, then the cached value is removed from the
property cache.

If @proxy has an expected interface (see
#GDBusProxy:g-interface-info) and @property_name is referenced by
it, then @value is checked against the type of the property.

If the @value #GVariant is floating, it is consumed. This allows
convenient 'inline' use of g_variant_new(), e.g.
|[<!-- language="C" -->
 g_dbus_proxy_set_cached_property (proxy,
                                   "SomeProperty",
                                   g_variant_new ("(si)",
                                                 "A String",
                                                 42));
]|

Normally you will not need to use this method since @proxy
is tracking changes using the
`org.freedesktop.DBus.Properties.PropertiesChanged`
D-Bus signal. However, for performance reasons an object may
decide to not use this signal for some properties and instead
use a proprietary out-of-band mechanism to transmit changes.

As a concrete example, consider an object with a property
`ChatroomParticipants` which is an array of strings. Instead of
transmitting the same (long) array every time the property changes,
it is more efficient to only transmit the delta using e.g. signals
`ChatroomParticipantJoined(String name)` and
`ChatroomParticipantParted(String name)`.
*/
func (self *TraitDBusProxy) SetCachedProperty(property_name string, value *C.GVariant) {
	__cgo__property_name := (*C.gchar)(unsafe.Pointer(C.CString(property_name)))
	C.g_dbus_proxy_set_cached_property(self.CPointer, __cgo__property_name, value)
	C.free(unsafe.Pointer(__cgo__property_name))
	return
}

/*
Sets the timeout to use if -1 (specifying default timeout) is
passed as @timeout_msec in the g_dbus_proxy_call() and
g_dbus_proxy_call_sync() functions.

See the #GDBusProxy:g-default-timeout property for more details.
*/
func (self *TraitDBusProxy) SetDefaultTimeout(timeout_msec int) {
	C.g_dbus_proxy_set_default_timeout(self.CPointer, C.gint(timeout_msec))
	return
}

/*
Ensure that interactions with @proxy conform to the given
interface. See the #GDBusProxy:g-interface-info property for more
details.
*/
func (self *TraitDBusProxy) SetInterfaceInfo(info *C.GDBusInterfaceInfo) {
	C.g_dbus_proxy_set_interface_info(self.CPointer, info)
	return
}

type TraitDBusServer struct{ CPointer *C.GDBusServer }
type IsDBusServer interface {
	GetDBusServerPointer() *C.GDBusServer
}

func (self *TraitDBusServer) GetDBusServerPointer() *C.GDBusServer {
	return self.CPointer
}
func NewTraitDBusServer(p unsafe.Pointer) *TraitDBusServer {
	return &TraitDBusServer{(*C.GDBusServer)(p)}
}

/*
Gets a D-Bus address string that can be used by clients to connect
to @server.
*/
func (self *TraitDBusServer) GetClientAddress() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_server_get_client_address(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the flags for @server.
*/
func (self *TraitDBusServer) GetFlags() (return__ C.GDBusServerFlags) {
	return__ = C.g_dbus_server_get_flags(self.CPointer)
	return
}

/*
Gets the GUID for @server.
*/
func (self *TraitDBusServer) GetGuid() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dbus_server_get_guid(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets whether @server is active.
*/
func (self *TraitDBusServer) IsActive() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_dbus_server_is_active(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Starts @server.
*/
func (self *TraitDBusServer) Start() {
	C.g_dbus_server_start(self.CPointer)
	return
}

/*
Stops @server.
*/
func (self *TraitDBusServer) Stop() {
	C.g_dbus_server_stop(self.CPointer)
	return
}

type TraitDataInputStream struct{ CPointer *C.GDataInputStream }
type IsDataInputStream interface {
	GetDataInputStreamPointer() *C.GDataInputStream
}

func (self *TraitDataInputStream) GetDataInputStreamPointer() *C.GDataInputStream {
	return self.CPointer
}
func NewTraitDataInputStream(p unsafe.Pointer) *TraitDataInputStream {
	return &TraitDataInputStream{(*C.GDataInputStream)(p)}
}

/*
Gets the byte order for the data input stream.
*/
func (self *TraitDataInputStream) GetByteOrder() (return__ C.GDataStreamByteOrder) {
	return__ = C.g_data_input_stream_get_byte_order(self.CPointer)
	return
}

/*
Gets the current newline type for the @stream.
*/
func (self *TraitDataInputStream) GetNewlineType() (return__ C.GDataStreamNewlineType) {
	return__ = C.g_data_input_stream_get_newline_type(self.CPointer)
	return
}

/*
Reads an unsigned 8-bit/1-byte value from @stream.
*/
func (self *TraitDataInputStream) ReadByte(cancellable IsCancellable) (return__ uint8, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.guchar
	__cgo__return__ = C.g_data_input_stream_read_byte(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = uint8(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Reads a 16-bit/2-byte value from @stream.

In order to get the correct byte order for this read operation,
see g_data_input_stream_get_byte_order() and g_data_input_stream_set_byte_order().
*/
func (self *TraitDataInputStream) ReadInt16(cancellable IsCancellable) (return__ int16, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gint16
	__cgo__return__ = C.g_data_input_stream_read_int16(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int16(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Reads a signed 32-bit/4-byte value from @stream.

In order to get the correct byte order for this read operation,
see g_data_input_stream_get_byte_order() and g_data_input_stream_set_byte_order().

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
*/
func (self *TraitDataInputStream) ReadInt32(cancellable IsCancellable) (return__ int32, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gint32
	__cgo__return__ = C.g_data_input_stream_read_int32(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int32(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Reads a 64-bit/8-byte value from @stream.

In order to get the correct byte order for this read operation,
see g_data_input_stream_get_byte_order() and g_data_input_stream_set_byte_order().

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
*/
func (self *TraitDataInputStream) ReadInt64(cancellable IsCancellable) (return__ int64, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gint64
	__cgo__return__ = C.g_data_input_stream_read_int64(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Reads a line from the data input stream.  Note that no encoding
checks or conversion is performed; the input is not guaranteed to
be UTF-8, and may in fact have embedded NUL characters.

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
*/
func (self *TraitDataInputStream) ReadLine(cancellable IsCancellable) (length int64, return__ *C.char, __err__ error) {
	var __cgo__length C.gsize
	var __cgo_error__ *C.GError
	return__ = C.g_data_input_stream_read_line(self.CPointer, &__cgo__length, cancellable.GetCancellablePointer(), &__cgo_error__)
	length = int64(__cgo__length)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
The asynchronous version of g_data_input_stream_read_line().  It is
an error to have two outstanding calls to this function.

When the operation is finished, @callback will be called. You
can then call g_data_input_stream_read_line_finish() to get
the result of the operation.
*/
func (self *TraitDataInputStream) ReadLineAsync(io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_data_input_stream_read_line_async(self.CPointer, C.gint(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finish an asynchronous call started by
g_data_input_stream_read_line_async().  Note the warning about
string encoding in g_data_input_stream_read_line() applies here as
well.
*/
func (self *TraitDataInputStream) ReadLineFinish(result *C.GAsyncResult) (length int64, return__ *C.char, __err__ error) {
	var __cgo__length C.gsize
	var __cgo_error__ *C.GError
	return__ = C.g_data_input_stream_read_line_finish(self.CPointer, result, &__cgo__length, &__cgo_error__)
	length = int64(__cgo__length)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Finish an asynchronous call started by
g_data_input_stream_read_line_async().
*/
func (self *TraitDataInputStream) ReadLineFinishUtf8(result *C.GAsyncResult) (length int64, return__ string, __err__ error) {
	var __cgo__length C.gsize
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_data_input_stream_read_line_finish_utf8(self.CPointer, result, &__cgo__length, &__cgo_error__)
	length = int64(__cgo__length)
	return__ = C.GoString(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Reads a UTF-8 encoded line from the data input stream.

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
*/
func (self *TraitDataInputStream) ReadLineUtf8(cancellable IsCancellable) (length int64, return__ string, __err__ error) {
	var __cgo__length C.gsize
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_data_input_stream_read_line_utf8(self.CPointer, &__cgo__length, cancellable.GetCancellablePointer(), &__cgo_error__)
	length = int64(__cgo__length)
	return__ = C.GoString(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Reads an unsigned 16-bit/2-byte value from @stream.

In order to get the correct byte order for this read operation,
see g_data_input_stream_get_byte_order() and g_data_input_stream_set_byte_order().
*/
func (self *TraitDataInputStream) ReadUint16(cancellable IsCancellable) (return__ uint16, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.guint16
	__cgo__return__ = C.g_data_input_stream_read_uint16(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = uint16(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Reads an unsigned 32-bit/4-byte value from @stream.

In order to get the correct byte order for this read operation,
see g_data_input_stream_get_byte_order() and g_data_input_stream_set_byte_order().

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
*/
func (self *TraitDataInputStream) ReadUint32(cancellable IsCancellable) (return__ uint32, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.guint32
	__cgo__return__ = C.g_data_input_stream_read_uint32(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = uint32(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Reads an unsigned 64-bit/8-byte value from @stream.

In order to get the correct byte order for this read operation,
see g_data_input_stream_get_byte_order().

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
*/
func (self *TraitDataInputStream) ReadUint64(cancellable IsCancellable) (return__ uint64, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.guint64
	__cgo__return__ = C.g_data_input_stream_read_uint64(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = uint64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Reads a string from the data input stream, up to the first
occurrence of any of the stop characters.

Note that, in contrast to g_data_input_stream_read_until_async(),
this function consumes the stop character that it finds.

Don't use this function in new code.  Its functionality is
inconsistent with g_data_input_stream_read_until_async().  Both
functions will be marked as deprecated in a future release.  Use
g_data_input_stream_read_upto() instead, but note that that function
does not consume the stop character.
*/
func (self *TraitDataInputStream) ReadUntil(stop_chars string, cancellable IsCancellable) (length int64, return__ string, __err__ error) {
	__cgo__stop_chars := (*C.gchar)(unsafe.Pointer(C.CString(stop_chars)))
	var __cgo__length C.gsize
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_data_input_stream_read_until(self.CPointer, __cgo__stop_chars, &__cgo__length, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__stop_chars))
	length = int64(__cgo__length)
	return__ = C.GoString(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
The asynchronous version of g_data_input_stream_read_until().
It is an error to have two outstanding calls to this function.

Note that, in contrast to g_data_input_stream_read_until(),
this function does not consume the stop character that it finds.  You
must read it for yourself.

When the operation is finished, @callback will be called. You
can then call g_data_input_stream_read_until_finish() to get
the result of the operation.

Don't use this function in new code.  Its functionality is
inconsistent with g_data_input_stream_read_until().  Both functions
will be marked as deprecated in a future release.  Use
g_data_input_stream_read_upto_async() instead.
*/
func (self *TraitDataInputStream) ReadUntilAsync(stop_chars string, io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__stop_chars := (*C.gchar)(unsafe.Pointer(C.CString(stop_chars)))
	C.g_data_input_stream_read_until_async(self.CPointer, __cgo__stop_chars, C.gint(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__stop_chars))
	return
}

/*
Finish an asynchronous call started by
g_data_input_stream_read_until_async().
*/
func (self *TraitDataInputStream) ReadUntilFinish(result *C.GAsyncResult) (length int64, return__ string, __err__ error) {
	var __cgo__length C.gsize
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_data_input_stream_read_until_finish(self.CPointer, result, &__cgo__length, &__cgo_error__)
	length = int64(__cgo__length)
	return__ = C.GoString(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Reads a string from the data input stream, up to the first
occurrence of any of the stop characters.

In contrast to g_data_input_stream_read_until(), this function
does not consume the stop character. You have to use
g_data_input_stream_read_byte() to get it before calling
g_data_input_stream_read_upto() again.

Note that @stop_chars may contain '\0' if @stop_chars_len is
specified.
*/
func (self *TraitDataInputStream) ReadUpto(stop_chars string, stop_chars_len int64, cancellable IsCancellable) (length int64, return__ string, __err__ error) {
	__cgo__stop_chars := (*C.gchar)(unsafe.Pointer(C.CString(stop_chars)))
	var __cgo__length C.gsize
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_data_input_stream_read_upto(self.CPointer, __cgo__stop_chars, C.gssize(stop_chars_len), &__cgo__length, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__stop_chars))
	length = int64(__cgo__length)
	return__ = C.GoString(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
The asynchronous version of g_data_input_stream_read_upto().
It is an error to have two outstanding calls to this function.

In contrast to g_data_input_stream_read_until(), this function
does not consume the stop character. You have to use
g_data_input_stream_read_byte() to get it before calling
g_data_input_stream_read_upto() again.

Note that @stop_chars may contain '\0' if @stop_chars_len is
specified.

When the operation is finished, @callback will be called. You
can then call g_data_input_stream_read_upto_finish() to get
the result of the operation.
*/
func (self *TraitDataInputStream) ReadUptoAsync(stop_chars string, stop_chars_len int64, io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__stop_chars := (*C.gchar)(unsafe.Pointer(C.CString(stop_chars)))
	C.g_data_input_stream_read_upto_async(self.CPointer, __cgo__stop_chars, C.gssize(stop_chars_len), C.gint(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__stop_chars))
	return
}

/*
Finish an asynchronous call started by
g_data_input_stream_read_upto_async().

Note that this function does not consume the stop character. You
have to use g_data_input_stream_read_byte() to get it before calling
g_data_input_stream_read_upto_async() again.
*/
func (self *TraitDataInputStream) ReadUptoFinish(result *C.GAsyncResult) (length int64, return__ string, __err__ error) {
	var __cgo__length C.gsize
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_data_input_stream_read_upto_finish(self.CPointer, result, &__cgo__length, &__cgo_error__)
	length = int64(__cgo__length)
	return__ = C.GoString(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This function sets the byte order for the given @stream. All subsequent
reads from the @stream will be read in the given @order.
*/
func (self *TraitDataInputStream) SetByteOrder(order C.GDataStreamByteOrder) {
	C.g_data_input_stream_set_byte_order(self.CPointer, order)
	return
}

/*
Sets the newline type for the @stream.

Note that using G_DATA_STREAM_NEWLINE_TYPE_ANY is slightly unsafe. If a read
chunk ends in "CR" we must read an additional byte to know if this is "CR" or
"CR LF", and this might block if there is no more data available.
*/
func (self *TraitDataInputStream) SetNewlineType(type_ C.GDataStreamNewlineType) {
	C.g_data_input_stream_set_newline_type(self.CPointer, type_)
	return
}

type TraitDataOutputStream struct{ CPointer *C.GDataOutputStream }
type IsDataOutputStream interface {
	GetDataOutputStreamPointer() *C.GDataOutputStream
}

func (self *TraitDataOutputStream) GetDataOutputStreamPointer() *C.GDataOutputStream {
	return self.CPointer
}
func NewTraitDataOutputStream(p unsafe.Pointer) *TraitDataOutputStream {
	return &TraitDataOutputStream{(*C.GDataOutputStream)(p)}
}

/*
Gets the byte order for the stream.
*/
func (self *TraitDataOutputStream) GetByteOrder() (return__ C.GDataStreamByteOrder) {
	return__ = C.g_data_output_stream_get_byte_order(self.CPointer)
	return
}

/*
Puts a byte into the output stream.
*/
func (self *TraitDataOutputStream) PutByte(data uint8, cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_data_output_stream_put_byte(self.CPointer, C.guchar(data), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Puts a signed 16-bit integer into the output stream.
*/
func (self *TraitDataOutputStream) PutInt16(data int16, cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_data_output_stream_put_int16(self.CPointer, C.gint16(data), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Puts a signed 32-bit integer into the output stream.
*/
func (self *TraitDataOutputStream) PutInt32(data int32, cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_data_output_stream_put_int32(self.CPointer, C.gint32(data), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Puts a signed 64-bit integer into the stream.
*/
func (self *TraitDataOutputStream) PutInt64(data int64, cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_data_output_stream_put_int64(self.CPointer, C.gint64(data), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Puts a string into the output stream.
*/
func (self *TraitDataOutputStream) PutString(str string, cancellable IsCancellable) (return__ bool, __err__ error) {
	__cgo__str := C.CString(str)
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_data_output_stream_put_string(self.CPointer, __cgo__str, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__str))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Puts an unsigned 16-bit integer into the output stream.
*/
func (self *TraitDataOutputStream) PutUint16(data uint16, cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_data_output_stream_put_uint16(self.CPointer, C.guint16(data), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Puts an unsigned 32-bit integer into the stream.
*/
func (self *TraitDataOutputStream) PutUint32(data uint32, cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_data_output_stream_put_uint32(self.CPointer, C.guint32(data), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Puts an unsigned 64-bit integer into the stream.
*/
func (self *TraitDataOutputStream) PutUint64(data uint64, cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_data_output_stream_put_uint64(self.CPointer, C.guint64(data), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Sets the byte order of the data output stream to @order.
*/
func (self *TraitDataOutputStream) SetByteOrder(order C.GDataStreamByteOrder) {
	C.g_data_output_stream_set_byte_order(self.CPointer, order)
	return
}

type TraitDesktopAppInfo struct{ CPointer *C.GDesktopAppInfo }
type IsDesktopAppInfo interface {
	GetDesktopAppInfoPointer() *C.GDesktopAppInfo
}

func (self *TraitDesktopAppInfo) GetDesktopAppInfoPointer() *C.GDesktopAppInfo {
	return self.CPointer
}
func NewTraitDesktopAppInfo(p unsafe.Pointer) *TraitDesktopAppInfo {
	return &TraitDesktopAppInfo{(*C.GDesktopAppInfo)(p)}
}

/*
Gets the user-visible display name of the "additional application
action" specified by @action_name.

This corresponds to the "Name" key within the keyfile group for the
action.
*/
func (self *TraitDesktopAppInfo) GetActionName(action_name string) (return__ string) {
	__cgo__action_name := (*C.gchar)(unsafe.Pointer(C.CString(action_name)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_desktop_app_info_get_action_name(self.CPointer, __cgo__action_name)
	C.free(unsafe.Pointer(__cgo__action_name))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Looks up a boolean value in the keyfile backing @info.

The @key is looked up in the "Desktop Entry" group.
*/
func (self *TraitDesktopAppInfo) GetBoolean(key string) (return__ bool) {
	__cgo__key := C.CString(key)
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_desktop_app_info_get_boolean(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the categories from the desktop file.
*/
func (self *TraitDesktopAppInfo) GetCategories() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_desktop_app_info_get_categories(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
When @info was created from a known filename, return it.  In some
situations such as the #GDesktopAppInfo returned from
g_desktop_app_info_new_from_keyfile(), this function will return %NULL.
*/
func (self *TraitDesktopAppInfo) GetFilename() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_desktop_app_info_get_filename(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Gets the generic name from the destkop file.
*/
func (self *TraitDesktopAppInfo) GetGenericName() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_desktop_app_info_get_generic_name(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
A desktop file is hidden if the Hidden key in it is
set to True.
*/
func (self *TraitDesktopAppInfo) GetIsHidden() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_desktop_app_info_get_is_hidden(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the keywords from the desktop file.
*/
func (self *TraitDesktopAppInfo) GetKeywords() (return__ **C.char) {
	return__ = C.g_desktop_app_info_get_keywords(self.CPointer)
	return
}

/*
Gets the value of the NoDisplay key, which helps determine if the
application info should be shown in menus. See
#G_KEY_FILE_DESKTOP_KEY_NO_DISPLAY and g_app_info_should_show().
*/
func (self *TraitDesktopAppInfo) GetNodisplay() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_desktop_app_info_get_nodisplay(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if the application info should be shown in menus that list available
applications for a specific name of the desktop, based on the
`OnlyShowIn` and `NotShowIn` keys.

If @desktop_env is %NULL, then the name of the desktop set with
g_desktop_app_info_set_desktop_env() is used.

Note that g_app_info_should_show() for @info will include this check (with
%NULL for @desktop_env) as well as additional checks.
*/
func (self *TraitDesktopAppInfo) GetShowIn(desktop_env string) (return__ bool) {
	__cgo__desktop_env := (*C.gchar)(unsafe.Pointer(C.CString(desktop_env)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_desktop_app_info_get_show_in(self.CPointer, __cgo__desktop_env)
	C.free(unsafe.Pointer(__cgo__desktop_env))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves the StartupWMClass field from @info. This represents the
WM_CLASS property of the main window of the application, if launched
through @info.
*/
func (self *TraitDesktopAppInfo) GetStartupWmClass() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_desktop_app_info_get_startup_wm_class(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Looks up a string value in the keyfile backing @info.

The @key is looked up in the "Desktop Entry" group.
*/
func (self *TraitDesktopAppInfo) GetString(key string) (return__ string) {
	__cgo__key := C.CString(key)
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_desktop_app_info_get_string(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Returns whether @key exists in the "Desktop Entry" group
of the keyfile backing @info.
*/
func (self *TraitDesktopAppInfo) HasKey(key string) (return__ bool) {
	__cgo__key := C.CString(key)
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_desktop_app_info_has_key(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Activates the named application action.

You may only call this function on action names that were
returned from g_desktop_app_info_list_actions().

Note that if the main entry of the desktop file indicates that the
application supports startup notification, and @launch_context is
non-%NULL, then startup notification will be used when activating the
action (and as such, invocation of the action on the receiving side
must signal the end of startup notification when it is completed).
This is the expected behaviour of applications declaring additional
actions, as per the desktop file specification.

As with g_app_info_launch() there is no way to detect failures that
occur while using this function.
*/
func (self *TraitDesktopAppInfo) LaunchAction(action_name string, launch_context IsAppLaunchContext) {
	__cgo__action_name := (*C.gchar)(unsafe.Pointer(C.CString(action_name)))
	C.g_desktop_app_info_launch_action(self.CPointer, __cgo__action_name, launch_context.GetAppLaunchContextPointer())
	C.free(unsafe.Pointer(__cgo__action_name))
	return
}

/*
This function performs the equivalent of g_app_info_launch_uris(),
but is intended primarily for operating system components that
launch applications.  Ordinary applications should use
g_app_info_launch_uris().

If the application is launched via traditional UNIX fork()/exec()
then @spawn_flags, @user_setup and @user_setup_data are used for the
call to g_spawn_async().  Additionally, @pid_callback (with
@pid_callback_data) will be called to inform about the PID of the
created process.

If application launching occurs via some other mechanism (eg: D-Bus
activation) then @spawn_flags, @user_setup, @user_setup_data,
@pid_callback and @pid_callback_data are ignored.
*/
func (self *TraitDesktopAppInfo) LaunchUrisAsManager(uris *C.GList, launch_context IsAppLaunchContext, spawn_flags C.GSpawnFlags, user_setup C.GSpawnChildSetupFunc, user_setup_data unsafe.Pointer, pid_callback C.GDesktopAppLaunchCallback, pid_callback_data unsafe.Pointer) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_desktop_app_info_launch_uris_as_manager(self.CPointer, uris, launch_context.GetAppLaunchContextPointer(), spawn_flags, user_setup, (C.gpointer)(user_setup_data), pid_callback, (C.gpointer)(pid_callback_data), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Returns the list of "additional application actions" supported on the
desktop file, as per the desktop file specification.

As per the specification, this is the list of actions that are
explicitly listed in the "Actions" key of the [Desktop Entry] group.
*/
func (self *TraitDesktopAppInfo) ListActions() (return__ **C.gchar) {
	return__ = C.g_desktop_app_info_list_actions(self.CPointer)
	return
}

type TraitEmblem struct{ CPointer *C.GEmblem }
type IsEmblem interface {
	GetEmblemPointer() *C.GEmblem
}

func (self *TraitEmblem) GetEmblemPointer() *C.GEmblem {
	return self.CPointer
}
func NewTraitEmblem(p unsafe.Pointer) *TraitEmblem {
	return &TraitEmblem{(*C.GEmblem)(p)}
}

/*
Gives back the icon from @emblem.
*/
func (self *TraitEmblem) GetIcon() (return__ *C.GIcon) {
	return__ = C.g_emblem_get_icon(self.CPointer)
	return
}

/*
Gets the origin of the emblem.
*/
func (self *TraitEmblem) GetOrigin() (return__ C.GEmblemOrigin) {
	return__ = C.g_emblem_get_origin(self.CPointer)
	return
}

type TraitEmblemedIcon struct{ CPointer *C.GEmblemedIcon }
type IsEmblemedIcon interface {
	GetEmblemedIconPointer() *C.GEmblemedIcon
}

func (self *TraitEmblemedIcon) GetEmblemedIconPointer() *C.GEmblemedIcon {
	return self.CPointer
}
func NewTraitEmblemedIcon(p unsafe.Pointer) *TraitEmblemedIcon {
	return &TraitEmblemedIcon{(*C.GEmblemedIcon)(p)}
}

/*
Adds @emblem to the #GList of #GEmblems.
*/
func (self *TraitEmblemedIcon) AddEmblem(emblem IsEmblem) {
	C.g_emblemed_icon_add_emblem(self.CPointer, emblem.GetEmblemPointer())
	return
}

/*
Removes all the emblems from @icon.
*/
func (self *TraitEmblemedIcon) ClearEmblems() {
	C.g_emblemed_icon_clear_emblems(self.CPointer)
	return
}

/*
Gets the list of emblems for the @icon.
*/
func (self *TraitEmblemedIcon) GetEmblems() (return__ *C.GList) {
	return__ = C.g_emblemed_icon_get_emblems(self.CPointer)
	return
}

/*
Gets the main icon for @emblemed.
*/
func (self *TraitEmblemedIcon) GetIcon() (return__ *C.GIcon) {
	return__ = C.g_emblemed_icon_get_icon(self.CPointer)
	return
}

type TraitFileEnumerator struct{ CPointer *C.GFileEnumerator }
type IsFileEnumerator interface {
	GetFileEnumeratorPointer() *C.GFileEnumerator
}

func (self *TraitFileEnumerator) GetFileEnumeratorPointer() *C.GFileEnumerator {
	return self.CPointer
}
func NewTraitFileEnumerator(p unsafe.Pointer) *TraitFileEnumerator {
	return &TraitFileEnumerator{(*C.GFileEnumerator)(p)}
}

/*
Releases all resources used by this enumerator, making the
enumerator return %G_IO_ERROR_CLOSED on all calls.

This will be automatically called when the last reference
is dropped, but you might want to call this function to make
sure resources are released as early as possible.
*/
func (self *TraitFileEnumerator) Close(cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_enumerator_close(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously closes the file enumerator.

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned in
g_file_enumerator_close_finish().
*/
func (self *TraitFileEnumerator) CloseAsync(io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_file_enumerator_close_async(self.CPointer, C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes closing a file enumerator, started from g_file_enumerator_close_async().

If the file enumerator was already closed when g_file_enumerator_close_async()
was called, then this function will report %G_IO_ERROR_CLOSED in @error, and
return %FALSE. If the file enumerator had pending operation when the close
operation was started, then this function will report %G_IO_ERROR_PENDING, and
return %FALSE.  If @cancellable was not %NULL, then the operation may have been
cancelled by triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be set, and %FALSE will be
returned.
*/
func (self *TraitFileEnumerator) CloseFinish(result *C.GAsyncResult) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_enumerator_close_finish(self.CPointer, result, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Return a new #GFile which refers to the file named by @info in the source
directory of @enumerator.  This function is primarily intended to be used
inside loops with g_file_enumerator_next_file().

This is a convenience method that's equivalent to:
|[<!-- language="C" -->
  gchar *name = g_file_info_get_name (info);
  GFile *child = g_file_get_child (g_file_enumerator_get_container (enumr),
                                   name);
]|
*/
func (self *TraitFileEnumerator) GetChild(info IsFileInfo) (return__ *C.GFile) {
	return__ = C.g_file_enumerator_get_child(self.CPointer, info.GetFileInfoPointer())
	return
}

/*
Get the #GFile container which is being enumerated.
*/
func (self *TraitFileEnumerator) GetContainer() (return__ *C.GFile) {
	return__ = C.g_file_enumerator_get_container(self.CPointer)
	return
}

/*
Checks if the file enumerator has pending operations.
*/
func (self *TraitFileEnumerator) HasPending() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_enumerator_has_pending(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if the file enumerator has been closed.
*/
func (self *TraitFileEnumerator) IsClosed() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_enumerator_is_closed(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns information for the next file in the enumerated object.
Will block until the information is available. The #GFileInfo
returned from this function will contain attributes that match the
attribute string that was passed when the #GFileEnumerator was created.

See the documentation of #GFileEnumerator for information about the
order of returned files.

On error, returns %NULL and sets @error to the error. If the
enumerator is at the end, %NULL will be returned and @error will
be unset.
*/
func (self *TraitFileEnumerator) NextFile(cancellable IsCancellable) (return__ *FileInfo, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GFileInfo
	__cgo__return__ = C.g_file_enumerator_next_file(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewFileInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Request information for a number of files from the enumerator asynchronously.
When all i/o for the operation is finished the @callback will be called with
the requested information.

See the documentation of #GFileEnumerator for information about the
order of returned files.

The callback can be called with less than @num_files files in case of error
or at the end of the enumerator. In case of a partial error the callback will
be called with any succeeding items and no error, and on the next request the
error will be reported. If a request is cancelled the callback will be called
with %G_IO_ERROR_CANCELLED.

During an async request no other sync and async calls are allowed, and will
result in %G_IO_ERROR_PENDING errors.

Any outstanding i/o request with higher priority (lower numerical value) will
be executed before an outstanding request with lower priority. Default
priority is %G_PRIORITY_DEFAULT.
*/
func (self *TraitFileEnumerator) NextFilesAsync(num_files int, io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_file_enumerator_next_files_async(self.CPointer, C.int(num_files), C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes the asynchronous operation started with g_file_enumerator_next_files_async().
*/
func (self *TraitFileEnumerator) NextFilesFinish(result *C.GAsyncResult) (return__ *C.GList, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_file_enumerator_next_files_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Sets the file enumerator as having pending operations.
*/
func (self *TraitFileEnumerator) SetPending(pending bool) {
	__cgo__pending := C.gboolean(0)
	if pending {
		__cgo__pending = C.gboolean(1)
	}
	C.g_file_enumerator_set_pending(self.CPointer, __cgo__pending)
	return
}

type TraitFileIOStream struct{ CPointer *C.GFileIOStream }
type IsFileIOStream interface {
	GetFileIOStreamPointer() *C.GFileIOStream
}

func (self *TraitFileIOStream) GetFileIOStreamPointer() *C.GFileIOStream {
	return self.CPointer
}
func NewTraitFileIOStream(p unsafe.Pointer) *TraitFileIOStream {
	return &TraitFileIOStream{(*C.GFileIOStream)(p)}
}

/*
Gets the entity tag for the file when it has been written.
This must be called after the stream has been written
and closed, as the etag can change while writing.
*/
func (self *TraitFileIOStream) GetEtag() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_file_io_stream_get_etag(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Queries a file io stream for the given @attributes.
This function blocks while querying the stream. For the asynchronous
version of this function, see g_file_io_stream_query_info_async().
While the stream is blocked, the stream will set the pending flag
internally, and any other operations on the stream will fail with
%G_IO_ERROR_PENDING.

Can fail if the stream was already closed (with @error being set to
%G_IO_ERROR_CLOSED), the stream has pending operations (with @error being
set to %G_IO_ERROR_PENDING), or if querying info is not supported for
the stream's interface (with @error being set to %G_IO_ERROR_NOT_SUPPORTED). I
all cases of failure, %NULL will be returned.

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be set, and %NULL will
be returned.
*/
func (self *TraitFileIOStream) QueryInfo(attributes string, cancellable IsCancellable) (return__ *FileInfo, __err__ error) {
	__cgo__attributes := C.CString(attributes)
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GFileInfo
	__cgo__return__ = C.g_file_io_stream_query_info(self.CPointer, __cgo__attributes, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__attributes))
	if __cgo__return__ != nil {
		return__ = NewFileInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously queries the @stream for a #GFileInfo. When completed,
@callback will be called with a #GAsyncResult which can be used to
finish the operation with g_file_io_stream_query_info_finish().

For the synchronous version of this function, see
g_file_io_stream_query_info().
*/
func (self *TraitFileIOStream) QueryInfoAsync(attributes string, io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__attributes := C.CString(attributes)
	C.g_file_io_stream_query_info_async(self.CPointer, __cgo__attributes, C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__attributes))
	return
}

/*
Finalizes the asynchronous query started
by g_file_io_stream_query_info_async().
*/
func (self *TraitFileIOStream) QueryInfoFinish(result *C.GAsyncResult) (return__ *FileInfo, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GFileInfo
	__cgo__return__ = C.g_file_io_stream_query_info_finish(self.CPointer, result, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewFileInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

type TraitFileIcon struct{ CPointer *C.GFileIcon }
type IsFileIcon interface {
	GetFileIconPointer() *C.GFileIcon
}

func (self *TraitFileIcon) GetFileIconPointer() *C.GFileIcon {
	return self.CPointer
}
func NewTraitFileIcon(p unsafe.Pointer) *TraitFileIcon {
	return &TraitFileIcon{(*C.GFileIcon)(p)}
}

/*
Gets the #GFile associated with the given @icon.
*/
func (self *TraitFileIcon) GetFile() (return__ *C.GFile) {
	return__ = C.g_file_icon_get_file(self.CPointer)
	return
}

type TraitFileInfo struct{ CPointer *C.GFileInfo }
type IsFileInfo interface {
	GetFileInfoPointer() *C.GFileInfo
}

func (self *TraitFileInfo) GetFileInfoPointer() *C.GFileInfo {
	return self.CPointer
}
func NewTraitFileInfo(p unsafe.Pointer) *TraitFileInfo {
	return &TraitFileInfo{(*C.GFileInfo)(p)}
}

/*
Clears the status information from @info.
*/
func (self *TraitFileInfo) ClearStatus() {
	C.g_file_info_clear_status(self.CPointer)
	return
}

/*
Copies all of the [GFileAttribute][gio-GFileAttribute]
from @src_info to @dest_info.
*/
func (self *TraitFileInfo) CopyInto(dest_info IsFileInfo) {
	C.g_file_info_copy_into(self.CPointer, dest_info.GetFileInfoPointer())
	return
}

/*
Duplicates a file info structure.
*/
func (self *TraitFileInfo) Dup() (return__ *FileInfo) {
	var __cgo__return__ *C.GFileInfo
	__cgo__return__ = C.g_file_info_dup(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewFileInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the value of a attribute, formated as a string.
This escapes things as needed to make the string valid
utf8.
*/
func (self *TraitFileInfo) GetAttributeAsString(attribute string) (return__ string) {
	__cgo__attribute := C.CString(attribute)
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_file_info_get_attribute_as_string(self.CPointer, __cgo__attribute)
	C.free(unsafe.Pointer(__cgo__attribute))
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Gets the value of a boolean attribute. If the attribute does not
contain a boolean value, %FALSE will be returned.
*/
func (self *TraitFileInfo) GetAttributeBoolean(attribute string) (return__ bool) {
	__cgo__attribute := C.CString(attribute)
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_info_get_attribute_boolean(self.CPointer, __cgo__attribute)
	C.free(unsafe.Pointer(__cgo__attribute))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value of a byte string attribute. If the attribute does
not contain a byte string, %NULL will be returned.
*/
func (self *TraitFileInfo) GetAttributeByteString(attribute string) (return__ string) {
	__cgo__attribute := C.CString(attribute)
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_file_info_get_attribute_byte_string(self.CPointer, __cgo__attribute)
	C.free(unsafe.Pointer(__cgo__attribute))
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Gets the attribute type, value and status for an attribute key.
*/
func (self *TraitFileInfo) GetAttributeData(attribute string) (type_ C.GFileAttributeType, value_pp unsafe.Pointer, status C.GFileAttributeStatus, return__ bool) {
	__cgo__attribute := C.CString(attribute)
	var __cgo__value_pp C.gpointer
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_info_get_attribute_data(self.CPointer, __cgo__attribute, &type_, &__cgo__value_pp, &status)
	C.free(unsafe.Pointer(__cgo__attribute))
	value_pp = unsafe.Pointer(__cgo__value_pp)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets a signed 32-bit integer contained within the attribute. If the
attribute does not contain a signed 32-bit integer, or is invalid,
0 will be returned.
*/
func (self *TraitFileInfo) GetAttributeInt32(attribute string) (return__ int32) {
	__cgo__attribute := C.CString(attribute)
	var __cgo__return__ C.gint32
	__cgo__return__ = C.g_file_info_get_attribute_int32(self.CPointer, __cgo__attribute)
	C.free(unsafe.Pointer(__cgo__attribute))
	return__ = int32(__cgo__return__)
	return
}

/*
Gets a signed 64-bit integer contained within the attribute. If the
attribute does not contain an signed 64-bit integer, or is invalid,
0 will be returned.
*/
func (self *TraitFileInfo) GetAttributeInt64(attribute string) (return__ int64) {
	__cgo__attribute := C.CString(attribute)
	var __cgo__return__ C.gint64
	__cgo__return__ = C.g_file_info_get_attribute_int64(self.CPointer, __cgo__attribute)
	C.free(unsafe.Pointer(__cgo__attribute))
	return__ = int64(__cgo__return__)
	return
}

/*
Gets the value of a #GObject attribute. If the attribute does
not contain a #GObject, %NULL will be returned.
*/
func (self *TraitFileInfo) GetAttributeObject(attribute string) (return__ *C.GObject) {
	__cgo__attribute := C.CString(attribute)
	return__ = C.g_file_info_get_attribute_object(self.CPointer, __cgo__attribute)
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Gets the attribute status for an attribute key.
*/
func (self *TraitFileInfo) GetAttributeStatus(attribute string) (return__ C.GFileAttributeStatus) {
	__cgo__attribute := C.CString(attribute)
	return__ = C.g_file_info_get_attribute_status(self.CPointer, __cgo__attribute)
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Gets the value of a string attribute. If the attribute does
not contain a string, %NULL will be returned.
*/
func (self *TraitFileInfo) GetAttributeString(attribute string) (return__ string) {
	__cgo__attribute := C.CString(attribute)
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_file_info_get_attribute_string(self.CPointer, __cgo__attribute)
	C.free(unsafe.Pointer(__cgo__attribute))
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Gets the value of a stringv attribute. If the attribute does
not contain a stringv, %NULL will be returned.
*/
func (self *TraitFileInfo) GetAttributeStringv(attribute string) (return__ **C.char) {
	__cgo__attribute := C.CString(attribute)
	return__ = C.g_file_info_get_attribute_stringv(self.CPointer, __cgo__attribute)
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Gets the attribute type for an attribute key.
*/
func (self *TraitFileInfo) GetAttributeType(attribute string) (return__ C.GFileAttributeType) {
	__cgo__attribute := C.CString(attribute)
	return__ = C.g_file_info_get_attribute_type(self.CPointer, __cgo__attribute)
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Gets an unsigned 32-bit integer contained within the attribute. If the
attribute does not contain an unsigned 32-bit integer, or is invalid,
0 will be returned.
*/
func (self *TraitFileInfo) GetAttributeUint32(attribute string) (return__ uint32) {
	__cgo__attribute := C.CString(attribute)
	var __cgo__return__ C.guint32
	__cgo__return__ = C.g_file_info_get_attribute_uint32(self.CPointer, __cgo__attribute)
	C.free(unsafe.Pointer(__cgo__attribute))
	return__ = uint32(__cgo__return__)
	return
}

/*
Gets a unsigned 64-bit integer contained within the attribute. If the
attribute does not contain an unsigned 64-bit integer, or is invalid,
0 will be returned.
*/
func (self *TraitFileInfo) GetAttributeUint64(attribute string) (return__ uint64) {
	__cgo__attribute := C.CString(attribute)
	var __cgo__return__ C.guint64
	__cgo__return__ = C.g_file_info_get_attribute_uint64(self.CPointer, __cgo__attribute)
	C.free(unsafe.Pointer(__cgo__attribute))
	return__ = uint64(__cgo__return__)
	return
}

/*
Gets the file's content type.
*/
func (self *TraitFileInfo) GetContentType() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_file_info_get_content_type(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Returns the #GDateTime representing the deletion date of the file, as
available in G_FILE_ATTRIBUTE_TRASH_DELETION_DATE. If the
G_FILE_ATTRIBUTE_TRASH_DELETION_DATE attribute is unset, %NULL is returned.
*/
func (self *TraitFileInfo) GetDeletionDate() (return__ *C.GDateTime) {
	return__ = C.g_file_info_get_deletion_date(self.CPointer)
	return
}

/*
Gets a display name for a file.
*/
func (self *TraitFileInfo) GetDisplayName() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_file_info_get_display_name(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Gets the edit name for a file.
*/
func (self *TraitFileInfo) GetEditName() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_file_info_get_edit_name(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Gets the [entity tag][gfile-etag] for a given
#GFileInfo. See %G_FILE_ATTRIBUTE_ETAG_VALUE.
*/
func (self *TraitFileInfo) GetEtag() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_file_info_get_etag(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Gets a file's type (whether it is a regular file, symlink, etc).
This is different from the file's content type, see g_file_info_get_content_type().
*/
func (self *TraitFileInfo) GetFileType() (return__ C.GFileType) {
	return__ = C.g_file_info_get_file_type(self.CPointer)
	return
}

/*
Gets the icon for a file.
*/
func (self *TraitFileInfo) GetIcon() (return__ *C.GIcon) {
	return__ = C.g_file_info_get_icon(self.CPointer)
	return
}

/*
Checks if a file is a backup file.
*/
func (self *TraitFileInfo) GetIsBackup() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_info_get_is_backup(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if a file is hidden.
*/
func (self *TraitFileInfo) GetIsHidden() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_info_get_is_hidden(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if a file is a symlink.
*/
func (self *TraitFileInfo) GetIsSymlink() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_info_get_is_symlink(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the modification time of the current @info and sets it
in @result.
*/
func (self *TraitFileInfo) GetModificationTime() (result C.GTimeVal) {
	C.g_file_info_get_modification_time(self.CPointer, &result)
	return
}

/*
Gets the name for a file.
*/
func (self *TraitFileInfo) GetName() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_file_info_get_name(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Gets the file's size.
*/
func (self *TraitFileInfo) GetSize() (return__ int64) {
	var __cgo__return__ C.goffset
	__cgo__return__ = C.g_file_info_get_size(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

/*
Gets the value of the sort_order attribute from the #GFileInfo.
See %G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER.
*/
func (self *TraitFileInfo) GetSortOrder() (return__ int32) {
	var __cgo__return__ C.gint32
	__cgo__return__ = C.g_file_info_get_sort_order(self.CPointer)
	return__ = int32(__cgo__return__)
	return
}

/*
Gets the symbolic icon for a file.
*/
func (self *TraitFileInfo) GetSymbolicIcon() (return__ *C.GIcon) {
	return__ = C.g_file_info_get_symbolic_icon(self.CPointer)
	return
}

/*
Gets the symlink target for a given #GFileInfo.
*/
func (self *TraitFileInfo) GetSymlinkTarget() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_file_info_get_symlink_target(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Checks if a file info structure has an attribute named @attribute.
*/
func (self *TraitFileInfo) HasAttribute(attribute string) (return__ bool) {
	__cgo__attribute := C.CString(attribute)
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_info_has_attribute(self.CPointer, __cgo__attribute)
	C.free(unsafe.Pointer(__cgo__attribute))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if a file info structure has an attribute in the
specified @name_space.
*/
func (self *TraitFileInfo) HasNamespace(name_space string) (return__ bool) {
	__cgo__name_space := C.CString(name_space)
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_info_has_namespace(self.CPointer, __cgo__name_space)
	C.free(unsafe.Pointer(__cgo__name_space))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Lists the file info structure's attributes.
*/
func (self *TraitFileInfo) ListAttributes(name_space string) (return__ **C.char) {
	__cgo__name_space := C.CString(name_space)
	return__ = C.g_file_info_list_attributes(self.CPointer, __cgo__name_space)
	C.free(unsafe.Pointer(__cgo__name_space))
	return
}

/*
Removes all cases of @attribute from @info if it exists.
*/
func (self *TraitFileInfo) RemoveAttribute(attribute string) {
	__cgo__attribute := C.CString(attribute)
	C.g_file_info_remove_attribute(self.CPointer, __cgo__attribute)
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Sets the @attribute to contain the given value, if possible. To unset the
attribute, use %G_ATTRIBUTE_TYPE_INVALID for @type.
*/
func (self *TraitFileInfo) SetAttribute(attribute string, type_ C.GFileAttributeType, value_p unsafe.Pointer) {
	__cgo__attribute := C.CString(attribute)
	C.g_file_info_set_attribute(self.CPointer, __cgo__attribute, type_, (C.gpointer)(value_p))
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Sets the @attribute to contain the given @attr_value,
if possible.
*/
func (self *TraitFileInfo) SetAttributeBoolean(attribute string, attr_value bool) {
	__cgo__attribute := C.CString(attribute)
	__cgo__attr_value := C.gboolean(0)
	if attr_value {
		__cgo__attr_value = C.gboolean(1)
	}
	C.g_file_info_set_attribute_boolean(self.CPointer, __cgo__attribute, __cgo__attr_value)
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Sets the @attribute to contain the given @attr_value,
if possible.
*/
func (self *TraitFileInfo) SetAttributeByteString(attribute string, attr_value string) {
	__cgo__attribute := C.CString(attribute)
	__cgo__attr_value := C.CString(attr_value)
	C.g_file_info_set_attribute_byte_string(self.CPointer, __cgo__attribute, __cgo__attr_value)
	C.free(unsafe.Pointer(__cgo__attribute))
	C.free(unsafe.Pointer(__cgo__attr_value))
	return
}

/*
Sets the @attribute to contain the given @attr_value,
if possible.
*/
func (self *TraitFileInfo) SetAttributeInt32(attribute string, attr_value int32) {
	__cgo__attribute := C.CString(attribute)
	C.g_file_info_set_attribute_int32(self.CPointer, __cgo__attribute, C.gint32(attr_value))
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Sets the @attribute to contain the given @attr_value,
if possible.
*/
func (self *TraitFileInfo) SetAttributeInt64(attribute string, attr_value int64) {
	__cgo__attribute := C.CString(attribute)
	C.g_file_info_set_attribute_int64(self.CPointer, __cgo__attribute, C.gint64(attr_value))
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Sets @mask on @info to match specific attribute types.
*/
func (self *TraitFileInfo) SetAttributeMask(mask *C.GFileAttributeMatcher) {
	C.g_file_info_set_attribute_mask(self.CPointer, mask)
	return
}

/*
Sets the @attribute to contain the given @attr_value,
if possible.
*/
func (self *TraitFileInfo) SetAttributeObject(attribute string, attr_value *C.GObject) {
	__cgo__attribute := C.CString(attribute)
	C.g_file_info_set_attribute_object(self.CPointer, __cgo__attribute, attr_value)
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Sets the attribute status for an attribute key. This is only
needed by external code that implement g_file_set_attributes_from_info()
or similar functions.

The attribute must exist in @info for this to work. Otherwise %FALSE
is returned and @info is unchanged.
*/
func (self *TraitFileInfo) SetAttributeStatus(attribute string, status C.GFileAttributeStatus) (return__ bool) {
	__cgo__attribute := C.CString(attribute)
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_info_set_attribute_status(self.CPointer, __cgo__attribute, status)
	C.free(unsafe.Pointer(__cgo__attribute))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the @attribute to contain the given @attr_value,
if possible.
*/
func (self *TraitFileInfo) SetAttributeString(attribute string, attr_value string) {
	__cgo__attribute := C.CString(attribute)
	__cgo__attr_value := C.CString(attr_value)
	C.g_file_info_set_attribute_string(self.CPointer, __cgo__attribute, __cgo__attr_value)
	C.free(unsafe.Pointer(__cgo__attribute))
	C.free(unsafe.Pointer(__cgo__attr_value))
	return
}

/*
Sets the @attribute to contain the given @attr_value,
if possible.

Sinze: 2.22
*/
func (self *TraitFileInfo) SetAttributeStringv(attribute string, attr_value []string) {
	__cgo__attribute := C.CString(attribute)
	__header__attr_value := (*reflect.SliceHeader)(unsafe.Pointer(&attr_value))
	C.g_file_info_set_attribute_stringv(self.CPointer, __cgo__attribute, (**C.char)(unsafe.Pointer(__header__attr_value.Data)))
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Sets the @attribute to contain the given @attr_value,
if possible.
*/
func (self *TraitFileInfo) SetAttributeUint32(attribute string, attr_value uint32) {
	__cgo__attribute := C.CString(attribute)
	C.g_file_info_set_attribute_uint32(self.CPointer, __cgo__attribute, C.guint32(attr_value))
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Sets the @attribute to contain the given @attr_value,
if possible.
*/
func (self *TraitFileInfo) SetAttributeUint64(attribute string, attr_value uint64) {
	__cgo__attribute := C.CString(attribute)
	C.g_file_info_set_attribute_uint64(self.CPointer, __cgo__attribute, C.guint64(attr_value))
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Sets the content type attribute for a given #GFileInfo.
See %G_FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE.
*/
func (self *TraitFileInfo) SetContentType(content_type string) {
	__cgo__content_type := C.CString(content_type)
	C.g_file_info_set_content_type(self.CPointer, __cgo__content_type)
	C.free(unsafe.Pointer(__cgo__content_type))
	return
}

/*
Sets the display name for the current #GFileInfo.
See %G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME.
*/
func (self *TraitFileInfo) SetDisplayName(display_name string) {
	__cgo__display_name := C.CString(display_name)
	C.g_file_info_set_display_name(self.CPointer, __cgo__display_name)
	C.free(unsafe.Pointer(__cgo__display_name))
	return
}

/*
Sets the edit name for the current file.
See %G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME.
*/
func (self *TraitFileInfo) SetEditName(edit_name string) {
	__cgo__edit_name := C.CString(edit_name)
	C.g_file_info_set_edit_name(self.CPointer, __cgo__edit_name)
	C.free(unsafe.Pointer(__cgo__edit_name))
	return
}

/*
Sets the file type in a #GFileInfo to @type.
See %G_FILE_ATTRIBUTE_STANDARD_TYPE.
*/
func (self *TraitFileInfo) SetFileType(type_ C.GFileType) {
	C.g_file_info_set_file_type(self.CPointer, type_)
	return
}

/*
Sets the icon for a given #GFileInfo.
See %G_FILE_ATTRIBUTE_STANDARD_ICON.
*/
func (self *TraitFileInfo) SetIcon(icon *C.GIcon) {
	C.g_file_info_set_icon(self.CPointer, icon)
	return
}

/*
Sets the "is_hidden" attribute in a #GFileInfo according to @is_hidden.
See %G_FILE_ATTRIBUTE_STANDARD_IS_HIDDEN.
*/
func (self *TraitFileInfo) SetIsHidden(is_hidden bool) {
	__cgo__is_hidden := C.gboolean(0)
	if is_hidden {
		__cgo__is_hidden = C.gboolean(1)
	}
	C.g_file_info_set_is_hidden(self.CPointer, __cgo__is_hidden)
	return
}

/*
Sets the "is_symlink" attribute in a #GFileInfo according to @is_symlink.
See %G_FILE_ATTRIBUTE_STANDARD_IS_SYMLINK.
*/
func (self *TraitFileInfo) SetIsSymlink(is_symlink bool) {
	__cgo__is_symlink := C.gboolean(0)
	if is_symlink {
		__cgo__is_symlink = C.gboolean(1)
	}
	C.g_file_info_set_is_symlink(self.CPointer, __cgo__is_symlink)
	return
}

/*
Sets the %G_FILE_ATTRIBUTE_TIME_MODIFIED attribute in the file
info to the given time value.
*/
func (self *TraitFileInfo) SetModificationTime(mtime *C.GTimeVal) {
	C.g_file_info_set_modification_time(self.CPointer, mtime)
	return
}

/*
Sets the name attribute for the current #GFileInfo.
See %G_FILE_ATTRIBUTE_STANDARD_NAME.
*/
func (self *TraitFileInfo) SetName(name string) {
	__cgo__name := C.CString(name)
	C.g_file_info_set_name(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Sets the %G_FILE_ATTRIBUTE_STANDARD_SIZE attribute in the file info
to the given size.
*/
func (self *TraitFileInfo) SetSize(size int64) {
	C.g_file_info_set_size(self.CPointer, C.goffset(size))
	return
}

/*
Sets the sort order attribute in the file info structure. See
%G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER.
*/
func (self *TraitFileInfo) SetSortOrder(sort_order int32) {
	C.g_file_info_set_sort_order(self.CPointer, C.gint32(sort_order))
	return
}

/*
Sets the symbolic icon for a given #GFileInfo.
See %G_FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON.
*/
func (self *TraitFileInfo) SetSymbolicIcon(icon *C.GIcon) {
	C.g_file_info_set_symbolic_icon(self.CPointer, icon)
	return
}

/*
Sets the %G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET attribute in the file info
to the given symlink target.
*/
func (self *TraitFileInfo) SetSymlinkTarget(symlink_target string) {
	__cgo__symlink_target := C.CString(symlink_target)
	C.g_file_info_set_symlink_target(self.CPointer, __cgo__symlink_target)
	C.free(unsafe.Pointer(__cgo__symlink_target))
	return
}

/*
Unsets a mask set by g_file_info_set_attribute_mask(), if one
is set.
*/
func (self *TraitFileInfo) UnsetAttributeMask() {
	C.g_file_info_unset_attribute_mask(self.CPointer)
	return
}

type TraitFileInputStream struct{ CPointer *C.GFileInputStream }
type IsFileInputStream interface {
	GetFileInputStreamPointer() *C.GFileInputStream
}

func (self *TraitFileInputStream) GetFileInputStreamPointer() *C.GFileInputStream {
	return self.CPointer
}
func NewTraitFileInputStream(p unsafe.Pointer) *TraitFileInputStream {
	return &TraitFileInputStream{(*C.GFileInputStream)(p)}
}

/*
Queries a file input stream the given @attributes. This function blocks
while querying the stream. For the asynchronous (non-blocking) version
of this function, see g_file_input_stream_query_info_async(). While the
stream is blocked, the stream will set the pending flag internally, and
any other operations on the stream will fail with %G_IO_ERROR_PENDING.
*/
func (self *TraitFileInputStream) QueryInfo(attributes string, cancellable IsCancellable) (return__ *FileInfo, __err__ error) {
	__cgo__attributes := C.CString(attributes)
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GFileInfo
	__cgo__return__ = C.g_file_input_stream_query_info(self.CPointer, __cgo__attributes, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__attributes))
	if __cgo__return__ != nil {
		return__ = NewFileInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Queries the stream information asynchronously.
When the operation is finished @callback will be called.
You can then call g_file_input_stream_query_info_finish()
to get the result of the operation.

For the synchronous version of this function,
see g_file_input_stream_query_info().

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be set
*/
func (self *TraitFileInputStream) QueryInfoAsync(attributes string, io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__attributes := C.CString(attributes)
	C.g_file_input_stream_query_info_async(self.CPointer, __cgo__attributes, C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__attributes))
	return
}

/*
Finishes an asynchronous info query operation.
*/
func (self *TraitFileInputStream) QueryInfoFinish(result *C.GAsyncResult) (return__ *FileInfo, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GFileInfo
	__cgo__return__ = C.g_file_input_stream_query_info_finish(self.CPointer, result, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewFileInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

type TraitFileMonitor struct{ CPointer *C.GFileMonitor }
type IsFileMonitor interface {
	GetFileMonitorPointer() *C.GFileMonitor
}

func (self *TraitFileMonitor) GetFileMonitorPointer() *C.GFileMonitor {
	return self.CPointer
}
func NewTraitFileMonitor(p unsafe.Pointer) *TraitFileMonitor {
	return &TraitFileMonitor{(*C.GFileMonitor)(p)}
}

/*
Cancels a file monitor.
*/
func (self *TraitFileMonitor) Cancel() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_monitor_cancel(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Emits the #GFileMonitor::changed signal if a change
has taken place. Should be called from file monitor
implementations only.

The signal will be emitted from an idle handler (in the
[thread-default main context][g-main-context-push-thread-default]).
*/
func (self *TraitFileMonitor) EmitEvent(child *C.GFile, other_file *C.GFile, event_type C.GFileMonitorEvent) {
	C.g_file_monitor_emit_event(self.CPointer, child, other_file, event_type)
	return
}

/*
Returns whether the monitor is canceled.
*/
func (self *TraitFileMonitor) IsCancelled() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_monitor_is_cancelled(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the rate limit to which the @monitor will report
consecutive change events to the same file.
*/
func (self *TraitFileMonitor) SetRateLimit(limit_msecs int) {
	C.g_file_monitor_set_rate_limit(self.CPointer, C.gint(limit_msecs))
	return
}

type TraitFileOutputStream struct{ CPointer *C.GFileOutputStream }
type IsFileOutputStream interface {
	GetFileOutputStreamPointer() *C.GFileOutputStream
}

func (self *TraitFileOutputStream) GetFileOutputStreamPointer() *C.GFileOutputStream {
	return self.CPointer
}
func NewTraitFileOutputStream(p unsafe.Pointer) *TraitFileOutputStream {
	return &TraitFileOutputStream{(*C.GFileOutputStream)(p)}
}

/*
Gets the entity tag for the file when it has been written.
This must be called after the stream has been written
and closed, as the etag can change while writing.
*/
func (self *TraitFileOutputStream) GetEtag() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_file_output_stream_get_etag(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Queries a file output stream for the given @attributes.
This function blocks while querying the stream. For the asynchronous
version of this function, see g_file_output_stream_query_info_async().
While the stream is blocked, the stream will set the pending flag
internally, and any other operations on the stream will fail with
%G_IO_ERROR_PENDING.

Can fail if the stream was already closed (with @error being set to
%G_IO_ERROR_CLOSED), the stream has pending operations (with @error being
set to %G_IO_ERROR_PENDING), or if querying info is not supported for
the stream's interface (with @error being set to %G_IO_ERROR_NOT_SUPPORTED). In
all cases of failure, %NULL will be returned.

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be set, and %NULL will
be returned.
*/
func (self *TraitFileOutputStream) QueryInfo(attributes string, cancellable IsCancellable) (return__ *FileInfo, __err__ error) {
	__cgo__attributes := C.CString(attributes)
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GFileInfo
	__cgo__return__ = C.g_file_output_stream_query_info(self.CPointer, __cgo__attributes, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__attributes))
	if __cgo__return__ != nil {
		return__ = NewFileInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously queries the @stream for a #GFileInfo. When completed,
@callback will be called with a #GAsyncResult which can be used to
finish the operation with g_file_output_stream_query_info_finish().

For the synchronous version of this function, see
g_file_output_stream_query_info().
*/
func (self *TraitFileOutputStream) QueryInfoAsync(attributes string, io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__attributes := C.CString(attributes)
	C.g_file_output_stream_query_info_async(self.CPointer, __cgo__attributes, C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__attributes))
	return
}

/*
Finalizes the asynchronous query started
by g_file_output_stream_query_info_async().
*/
func (self *TraitFileOutputStream) QueryInfoFinish(result *C.GAsyncResult) (return__ *FileInfo, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GFileInfo
	__cgo__return__ = C.g_file_output_stream_query_info_finish(self.CPointer, result, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewFileInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

type TraitFilenameCompleter struct{ CPointer *C.GFilenameCompleter }
type IsFilenameCompleter interface {
	GetFilenameCompleterPointer() *C.GFilenameCompleter
}

func (self *TraitFilenameCompleter) GetFilenameCompleterPointer() *C.GFilenameCompleter {
	return self.CPointer
}
func NewTraitFilenameCompleter(p unsafe.Pointer) *TraitFilenameCompleter {
	return &TraitFilenameCompleter{(*C.GFilenameCompleter)(p)}
}

/*
Obtains a completion for @initial_text from @completer.
*/
func (self *TraitFilenameCompleter) GetCompletionSuffix(initial_text string) (return__ string) {
	__cgo__initial_text := C.CString(initial_text)
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_filename_completer_get_completion_suffix(self.CPointer, __cgo__initial_text)
	C.free(unsafe.Pointer(__cgo__initial_text))
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Gets an array of completion strings for a given initial text.
*/
func (self *TraitFilenameCompleter) GetCompletions(initial_text string) (return__ **C.char) {
	__cgo__initial_text := C.CString(initial_text)
	return__ = C.g_filename_completer_get_completions(self.CPointer, __cgo__initial_text)
	C.free(unsafe.Pointer(__cgo__initial_text))
	return
}

/*
If @dirs_only is %TRUE, @completer will only
complete directory names, and not file names.
*/
func (self *TraitFilenameCompleter) SetDirsOnly(dirs_only bool) {
	__cgo__dirs_only := C.gboolean(0)
	if dirs_only {
		__cgo__dirs_only = C.gboolean(1)
	}
	C.g_filename_completer_set_dirs_only(self.CPointer, __cgo__dirs_only)
	return
}

type TraitFilterInputStream struct{ CPointer *C.GFilterInputStream }
type IsFilterInputStream interface {
	GetFilterInputStreamPointer() *C.GFilterInputStream
}

func (self *TraitFilterInputStream) GetFilterInputStreamPointer() *C.GFilterInputStream {
	return self.CPointer
}
func NewTraitFilterInputStream(p unsafe.Pointer) *TraitFilterInputStream {
	return &TraitFilterInputStream{(*C.GFilterInputStream)(p)}
}

/*
Gets the base stream for the filter stream.
*/
func (self *TraitFilterInputStream) GetBaseStream() (return__ *InputStream) {
	var __cgo__return__ *C.GInputStream
	__cgo__return__ = C.g_filter_input_stream_get_base_stream(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewInputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns whether the base stream will be closed when @stream is
closed.
*/
func (self *TraitFilterInputStream) GetCloseBaseStream() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_filter_input_stream_get_close_base_stream(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets whether the base stream will be closed when @stream is closed.
*/
func (self *TraitFilterInputStream) SetCloseBaseStream(close_base bool) {
	__cgo__close_base := C.gboolean(0)
	if close_base {
		__cgo__close_base = C.gboolean(1)
	}
	C.g_filter_input_stream_set_close_base_stream(self.CPointer, __cgo__close_base)
	return
}

type TraitFilterOutputStream struct{ CPointer *C.GFilterOutputStream }
type IsFilterOutputStream interface {
	GetFilterOutputStreamPointer() *C.GFilterOutputStream
}

func (self *TraitFilterOutputStream) GetFilterOutputStreamPointer() *C.GFilterOutputStream {
	return self.CPointer
}
func NewTraitFilterOutputStream(p unsafe.Pointer) *TraitFilterOutputStream {
	return &TraitFilterOutputStream{(*C.GFilterOutputStream)(p)}
}

/*
Gets the base stream for the filter stream.
*/
func (self *TraitFilterOutputStream) GetBaseStream() (return__ *OutputStream) {
	var __cgo__return__ *C.GOutputStream
	__cgo__return__ = C.g_filter_output_stream_get_base_stream(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewOutputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns whether the base stream will be closed when @stream is
closed.
*/
func (self *TraitFilterOutputStream) GetCloseBaseStream() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_filter_output_stream_get_close_base_stream(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets whether the base stream will be closed when @stream is closed.
*/
func (self *TraitFilterOutputStream) SetCloseBaseStream(close_base bool) {
	__cgo__close_base := C.gboolean(0)
	if close_base {
		__cgo__close_base = C.gboolean(1)
	}
	C.g_filter_output_stream_set_close_base_stream(self.CPointer, __cgo__close_base)
	return
}

type TraitIOModule struct{ CPointer *C.GIOModule }
type IsIOModule interface {
	GetIOModulePointer() *C.GIOModule
}

func (self *TraitIOModule) GetIOModulePointer() *C.GIOModule {
	return self.CPointer
}
func NewTraitIOModule(p unsafe.Pointer) *TraitIOModule {
	return &TraitIOModule{(*C.GIOModule)(p)}
}

// g_io_module_load is not generated due to explicit ignore

// g_io_module_unload is not generated due to explicit ignore

type TraitIOStream struct{ CPointer *C.GIOStream }
type IsIOStream interface {
	GetIOStreamPointer() *C.GIOStream
}

func (self *TraitIOStream) GetIOStreamPointer() *C.GIOStream {
	return self.CPointer
}
func NewTraitIOStream(p unsafe.Pointer) *TraitIOStream {
	return &TraitIOStream{(*C.GIOStream)(p)}
}

/*
Clears the pending flag on @stream.
*/
func (self *TraitIOStream) ClearPending() {
	C.g_io_stream_clear_pending(self.CPointer)
	return
}

/*
Closes the stream, releasing resources related to it. This will also
closes the individual input and output streams, if they are not already
closed.

Once the stream is closed, all other operations will return
%G_IO_ERROR_CLOSED. Closing a stream multiple times will not
return an error.

Closing a stream will automatically flush any outstanding buffers
in the stream.

Streams will be automatically closed when the last reference
is dropped, but you might want to call this function to make sure
resources are released as early as possible.

Some streams might keep the backing store of the stream (e.g. a file
descriptor) open after the stream is closed. See the documentation for
the individual stream for details.

On failure the first error that happened will be reported, but the
close operation will finish as much as possible. A stream that failed
to close will still return %G_IO_ERROR_CLOSED for all operations.
Still, it is important to check and report the error to the user,
otherwise there might be a loss of data as all data might not be written.

If @cancellable is not NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
Cancelling a close will still leave the stream closed, but some streams
can use a faster close that doesn't block to e.g. check errors.

The default implementation of this method just calls close on the
individual input/output streams.
*/
func (self *TraitIOStream) Close(cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_io_stream_close(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Requests an asynchronous close of the stream, releasing resources
related to it. When the operation is finished @callback will be
called. You can then call g_io_stream_close_finish() to get
the result of the operation.

For behaviour details see g_io_stream_close().

The asynchronous methods have a default fallback that uses threads
to implement asynchronicity, so they are optional for inheriting
classes. However, if you override one you must override all.
*/
func (self *TraitIOStream) CloseAsync(io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_io_stream_close_async(self.CPointer, C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Closes a stream.
*/
func (self *TraitIOStream) CloseFinish(result *C.GAsyncResult) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_io_stream_close_finish(self.CPointer, result, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the input stream for this object. This is used
for reading.
*/
func (self *TraitIOStream) GetInputStream() (return__ *InputStream) {
	var __cgo__return__ *C.GInputStream
	__cgo__return__ = C.g_io_stream_get_input_stream(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewInputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the output stream for this object. This is used for
writing.
*/
func (self *TraitIOStream) GetOutputStream() (return__ *OutputStream) {
	var __cgo__return__ *C.GOutputStream
	__cgo__return__ = C.g_io_stream_get_output_stream(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewOutputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Checks if a stream has pending actions.
*/
func (self *TraitIOStream) HasPending() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_io_stream_has_pending(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if a stream is closed.
*/
func (self *TraitIOStream) IsClosed() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_io_stream_is_closed(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets @stream to have actions pending. If the pending flag is
already set or @stream is closed, it will return %FALSE and set
@error.
*/
func (self *TraitIOStream) SetPending() (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_io_stream_set_pending(self.CPointer, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asyncronously splice the output stream of @stream1 to the input stream of
@stream2, and splice the output stream of @stream2 to the input stream of
@stream1.

When the operation is finished @callback will be called.
You can then call g_io_stream_splice_finish() to get the
result of the operation.
*/
func (self *TraitIOStream) SpliceAsync(stream2 IsIOStream, flags C.GIOStreamSpliceFlags, io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_io_stream_splice_async(self.CPointer, stream2.GetIOStreamPointer(), flags, C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

type TraitInetAddress struct{ CPointer *C.GInetAddress }
type IsInetAddress interface {
	GetInetAddressPointer() *C.GInetAddress
}

func (self *TraitInetAddress) GetInetAddressPointer() *C.GInetAddress {
	return self.CPointer
}
func NewTraitInetAddress(p unsafe.Pointer) *TraitInetAddress {
	return &TraitInetAddress{(*C.GInetAddress)(p)}
}

/*
Checks if two #GInetAddress instances are equal, e.g. the same address.
*/
func (self *TraitInetAddress) Equal(other_address IsInetAddress) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_inet_address_equal(self.CPointer, other_address.GetInetAddressPointer())
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets @address's family
*/
func (self *TraitInetAddress) GetFamily() (return__ C.GSocketFamily) {
	return__ = C.g_inet_address_get_family(self.CPointer)
	return
}

/*
Tests whether @address is the "any" address for its family.
*/
func (self *TraitInetAddress) GetIsAny() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_inet_address_get_is_any(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Tests whether @address is a link-local address (that is, if it
identifies a host on a local network that is not connected to the
Internet).
*/
func (self *TraitInetAddress) GetIsLinkLocal() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_inet_address_get_is_link_local(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Tests whether @address is the loopback address for its family.
*/
func (self *TraitInetAddress) GetIsLoopback() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_inet_address_get_is_loopback(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Tests whether @address is a global multicast address.
*/
func (self *TraitInetAddress) GetIsMcGlobal() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_inet_address_get_is_mc_global(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Tests whether @address is a link-local multicast address.
*/
func (self *TraitInetAddress) GetIsMcLinkLocal() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_inet_address_get_is_mc_link_local(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Tests whether @address is a node-local multicast address.
*/
func (self *TraitInetAddress) GetIsMcNodeLocal() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_inet_address_get_is_mc_node_local(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Tests whether @address is an organization-local multicast address.
*/
func (self *TraitInetAddress) GetIsMcOrgLocal() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_inet_address_get_is_mc_org_local(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Tests whether @address is a site-local multicast address.
*/
func (self *TraitInetAddress) GetIsMcSiteLocal() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_inet_address_get_is_mc_site_local(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Tests whether @address is a multicast address.
*/
func (self *TraitInetAddress) GetIsMulticast() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_inet_address_get_is_multicast(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Tests whether @address is a site-local address such as 10.0.0.1
(that is, the address identifies a host on a local network that can
not be reached directly from the Internet, but which may have
outgoing Internet connectivity via a NAT or firewall).
*/
func (self *TraitInetAddress) GetIsSiteLocal() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_inet_address_get_is_site_local(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the size of the native raw binary address for @address. This
is the size of the data that you get from g_inet_address_to_bytes().
*/
func (self *TraitInetAddress) GetNativeSize() (return__ int64) {
	var __cgo__return__ C.gsize
	__cgo__return__ = C.g_inet_address_get_native_size(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

/*
Gets the raw binary address data from @address.
*/
func (self *TraitInetAddress) ToBytes() (return__ *C.guint8) {
	return__ = C.g_inet_address_to_bytes(self.CPointer)
	return
}

/*
Converts @address to string form.
*/
func (self *TraitInetAddress) ToString() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_inet_address_to_string(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

type TraitInetAddressMask struct{ CPointer *C.GInetAddressMask }
type IsInetAddressMask interface {
	GetInetAddressMaskPointer() *C.GInetAddressMask
}

func (self *TraitInetAddressMask) GetInetAddressMaskPointer() *C.GInetAddressMask {
	return self.CPointer
}
func NewTraitInetAddressMask(p unsafe.Pointer) *TraitInetAddressMask {
	return &TraitInetAddressMask{(*C.GInetAddressMask)(p)}
}

/*
Tests if @mask and @mask2 are the same mask.
*/
func (self *TraitInetAddressMask) Equal(mask2 IsInetAddressMask) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_inet_address_mask_equal(self.CPointer, mask2.GetInetAddressMaskPointer())
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets @mask's base address
*/
func (self *TraitInetAddressMask) GetAddress() (return__ *InetAddress) {
	var __cgo__return__ *C.GInetAddress
	__cgo__return__ = C.g_inet_address_mask_get_address(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewInetAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the #GSocketFamily of @mask's address
*/
func (self *TraitInetAddressMask) GetFamily() (return__ C.GSocketFamily) {
	return__ = C.g_inet_address_mask_get_family(self.CPointer)
	return
}

/*
Gets @mask's length
*/
func (self *TraitInetAddressMask) GetLength() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_inet_address_mask_get_length(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Tests if @address falls within the range described by @mask.
*/
func (self *TraitInetAddressMask) Matches(address IsInetAddress) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_inet_address_mask_matches(self.CPointer, address.GetInetAddressPointer())
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Converts @mask back to its corresponding string form.
*/
func (self *TraitInetAddressMask) ToString() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_inet_address_mask_to_string(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

type TraitInetSocketAddress struct{ CPointer *C.GInetSocketAddress }
type IsInetSocketAddress interface {
	GetInetSocketAddressPointer() *C.GInetSocketAddress
}

func (self *TraitInetSocketAddress) GetInetSocketAddressPointer() *C.GInetSocketAddress {
	return self.CPointer
}
func NewTraitInetSocketAddress(p unsafe.Pointer) *TraitInetSocketAddress {
	return &TraitInetSocketAddress{(*C.GInetSocketAddress)(p)}
}

/*
Gets @address's #GInetAddress.
*/
func (self *TraitInetSocketAddress) GetAddress() (return__ *InetAddress) {
	var __cgo__return__ *C.GInetAddress
	__cgo__return__ = C.g_inet_socket_address_get_address(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewInetAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the `sin6_flowinfo` field from @address,
which must be an IPv6 address.
*/
func (self *TraitInetSocketAddress) GetFlowinfo() (return__ uint32) {
	var __cgo__return__ C.guint32
	__cgo__return__ = C.g_inet_socket_address_get_flowinfo(self.CPointer)
	return__ = uint32(__cgo__return__)
	return
}

/*
Gets @address's port.
*/
func (self *TraitInetSocketAddress) GetPort() (return__ uint16) {
	var __cgo__return__ C.guint16
	__cgo__return__ = C.g_inet_socket_address_get_port(self.CPointer)
	return__ = uint16(__cgo__return__)
	return
}

/*
Gets the `sin6_scope_id` field from @address,
which must be an IPv6 address.
*/
func (self *TraitInetSocketAddress) GetScopeId() (return__ uint32) {
	var __cgo__return__ C.guint32
	__cgo__return__ = C.g_inet_socket_address_get_scope_id(self.CPointer)
	return__ = uint32(__cgo__return__)
	return
}

type TraitInputStream struct{ CPointer *C.GInputStream }
type IsInputStream interface {
	GetInputStreamPointer() *C.GInputStream
}

func (self *TraitInputStream) GetInputStreamPointer() *C.GInputStream {
	return self.CPointer
}
func NewTraitInputStream(p unsafe.Pointer) *TraitInputStream {
	return &TraitInputStream{(*C.GInputStream)(p)}
}

/*
Clears the pending flag on @stream.
*/
func (self *TraitInputStream) ClearPending() {
	C.g_input_stream_clear_pending(self.CPointer)
	return
}

/*
Closes the stream, releasing resources related to it.

Once the stream is closed, all other operations will return %G_IO_ERROR_CLOSED.
Closing a stream multiple times will not return an error.

Streams will be automatically closed when the last reference
is dropped, but you might want to call this function to make sure
resources are released as early as possible.

Some streams might keep the backing store of the stream (e.g. a file descriptor)
open after the stream is closed. See the documentation for the individual
stream for details.

On failure the first error that happened will be reported, but the close
operation will finish as much as possible. A stream that failed to
close will still return %G_IO_ERROR_CLOSED for all operations. Still, it
is important to check and report the error to the user.

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
Cancelling a close will still leave the stream closed, but some streams
can use a faster close that doesn't block to e.g. check errors.
*/
func (self *TraitInputStream) Close(cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_input_stream_close(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Requests an asynchronous closes of the stream, releasing resources related to it.
When the operation is finished @callback will be called.
You can then call g_input_stream_close_finish() to get the result of the
operation.

For behaviour details see g_input_stream_close().

The asyncronous methods have a default fallback that uses threads to implement
asynchronicity, so they are optional for inheriting classes. However, if you
override one you must override all.
*/
func (self *TraitInputStream) CloseAsync(io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_input_stream_close_async(self.CPointer, C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes closing a stream asynchronously, started from g_input_stream_close_async().
*/
func (self *TraitInputStream) CloseFinish(result *C.GAsyncResult) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_input_stream_close_finish(self.CPointer, result, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Checks if an input stream has pending actions.
*/
func (self *TraitInputStream) HasPending() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_input_stream_has_pending(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if an input stream is closed.
*/
func (self *TraitInputStream) IsClosed() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_input_stream_is_closed(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Tries to read @count bytes from the stream into the buffer starting at
@buffer. Will block during this read.

If count is zero returns zero and does nothing. A value of @count
larger than %G_MAXSSIZE will cause a %G_IO_ERROR_INVALID_ARGUMENT error.

On success, the number of bytes read into the buffer is returned.
It is not an error if this is not the same as the requested size, as it
can happen e.g. near the end of a file. Zero is returned on end of file
(or if @count is zero),  but never otherwise.

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
operation was partially finished when the operation was cancelled the
partial result will be returned, without an error.

On error -1 is returned and @error is set accordingly.
*/
func (self *TraitInputStream) Read(buffer []byte, count int64, cancellable IsCancellable) (return__ int64, __err__ error) {
	__header__buffer := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_input_stream_read(self.CPointer, (unsafe.Pointer)(unsafe.Pointer(__header__buffer.Data)), C.gsize(count), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Tries to read @count bytes from the stream into the buffer starting at
@buffer. Will block during this read.

This function is similar to g_input_stream_read(), except it tries to
read as many bytes as requested, only stopping on an error or end of stream.

On a successful read of @count bytes, or if we reached the end of the
stream,  %TRUE is returned, and @bytes_read is set to the number of bytes
read into @buffer.

If there is an error during the operation %FALSE is returned and @error
is set to indicate the error status, @bytes_read is updated to contain
the number of bytes read into @buffer before the error occurred.
*/
func (self *TraitInputStream) ReadAll(buffer []byte, count int64, cancellable IsCancellable) (bytes_read int64, return__ bool, __err__ error) {
	__header__buffer := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	var __cgo__bytes_read C.gsize
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_input_stream_read_all(self.CPointer, (unsafe.Pointer)(unsafe.Pointer(__header__buffer.Data)), C.gsize(count), &__cgo__bytes_read, cancellable.GetCancellablePointer(), &__cgo_error__)
	bytes_read = int64(__cgo__bytes_read)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Request an asynchronous read of @count bytes from the stream into the buffer
starting at @buffer. When the operation is finished @callback will be called.
You can then call g_input_stream_read_finish() to get the result of the
operation.

During an async request no other sync and async calls are allowed on @stream, and will
result in %G_IO_ERROR_PENDING errors.

A value of @count larger than %G_MAXSSIZE will cause a %G_IO_ERROR_INVALID_ARGUMENT error.

On success, the number of bytes read into the buffer will be passed to the
callback. It is not an error if this is not the same as the requested size, as it
can happen e.g. near the end of a file, but generally we try to read
as many bytes as requested. Zero is returned on end of file
(or if @count is zero),  but never otherwise.

Any outstanding i/o request with higher priority (lower numerical value) will
be executed before an outstanding request with lower priority. Default
priority is %G_PRIORITY_DEFAULT.

The asyncronous methods have a default fallback that uses threads to implement
asynchronicity, so they are optional for inheriting classes. However, if you
override one you must override all.
*/
func (self *TraitInputStream) ReadAsync(buffer []byte, count int64, io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__header__buffer := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	C.g_input_stream_read_async(self.CPointer, (unsafe.Pointer)(unsafe.Pointer(__header__buffer.Data)), C.gsize(count), C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Like g_input_stream_read(), this tries to read @count bytes from
the stream in a blocking fashion. However, rather than reading into
a user-supplied buffer, this will create a new #GBytes containing
the data that was read. This may be easier to use from language
bindings.

If count is zero, returns a zero-length #GBytes and does nothing. A
value of @count larger than %G_MAXSSIZE will cause a
%G_IO_ERROR_INVALID_ARGUMENT error.

On success, a new #GBytes is returned. It is not an error if the
size of this object is not the same as the requested size, as it
can happen e.g. near the end of a file. A zero-length #GBytes is
returned on end of file (or if @count is zero), but never
otherwise.

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
operation was partially finished when the operation was cancelled the
partial result will be returned, without an error.

On error %NULL is returned and @error is set accordingly.
*/
func (self *TraitInputStream) ReadBytes(count int64, cancellable IsCancellable) (return__ *C.GBytes, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_input_stream_read_bytes(self.CPointer, C.gsize(count), cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Request an asynchronous read of @count bytes from the stream into a
new #GBytes. When the operation is finished @callback will be
called. You can then call g_input_stream_read_bytes_finish() to get the
result of the operation.

During an async request no other sync and async calls are allowed
on @stream, and will result in %G_IO_ERROR_PENDING errors.

A value of @count larger than %G_MAXSSIZE will cause a
%G_IO_ERROR_INVALID_ARGUMENT error.

On success, the new #GBytes will be passed to the callback. It is
not an error if this is smaller than the requested size, as it can
happen e.g. near the end of a file, but generally we try to read as
many bytes as requested. Zero is returned on end of file (or if
@count is zero), but never otherwise.

Any outstanding I/O request with higher priority (lower numerical
value) will be executed before an outstanding request with lower
priority. Default priority is %G_PRIORITY_DEFAULT.
*/
func (self *TraitInputStream) ReadBytesAsync(count int64, io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_input_stream_read_bytes_async(self.CPointer, C.gsize(count), C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes an asynchronous stream read-into-#GBytes operation.
*/
func (self *TraitInputStream) ReadBytesFinish(result *C.GAsyncResult) (return__ *C.GBytes, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_input_stream_read_bytes_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Finishes an asynchronous stream read operation.
*/
func (self *TraitInputStream) ReadFinish(result *C.GAsyncResult) (return__ int64, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_input_stream_read_finish(self.CPointer, result, &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Sets @stream to have actions pending. If the pending flag is
already set or @stream is closed, it will return %FALSE and set
@error.
*/
func (self *TraitInputStream) SetPending() (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_input_stream_set_pending(self.CPointer, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Tries to skip @count bytes from the stream. Will block during the operation.

This is identical to g_input_stream_read(), from a behaviour standpoint,
but the bytes that are skipped are not returned to the user. Some
streams have an implementation that is more efficient than reading the data.

This function is optional for inherited classes, as the default implementation
emulates it using read.

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
operation was partially finished when the operation was cancelled the
partial result will be returned, without an error.
*/
func (self *TraitInputStream) Skip(count int64, cancellable IsCancellable) (return__ int64, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_input_stream_skip(self.CPointer, C.gsize(count), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Request an asynchronous skip of @count bytes from the stream.
When the operation is finished @callback will be called.
You can then call g_input_stream_skip_finish() to get the result
of the operation.

During an async request no other sync and async calls are allowed,
and will result in %G_IO_ERROR_PENDING errors.

A value of @count larger than %G_MAXSSIZE will cause a %G_IO_ERROR_INVALID_ARGUMENT error.

On success, the number of bytes skipped will be passed to the callback.
It is not an error if this is not the same as the requested size, as it
can happen e.g. near the end of a file, but generally we try to skip
as many bytes as requested. Zero is returned on end of file
(or if @count is zero), but never otherwise.

Any outstanding i/o request with higher priority (lower numerical value)
will be executed before an outstanding request with lower priority.
Default priority is %G_PRIORITY_DEFAULT.

The asynchronous methods have a default fallback that uses threads to
implement asynchronicity, so they are optional for inheriting classes.
However, if you override one, you must override all.
*/
func (self *TraitInputStream) SkipAsync(count int64, io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_input_stream_skip_async(self.CPointer, C.gsize(count), C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes a stream skip operation.
*/
func (self *TraitInputStream) SkipFinish(result *C.GAsyncResult) (return__ int64, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_input_stream_skip_finish(self.CPointer, result, &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

type TraitMemoryInputStream struct{ CPointer *C.GMemoryInputStream }
type IsMemoryInputStream interface {
	GetMemoryInputStreamPointer() *C.GMemoryInputStream
}

func (self *TraitMemoryInputStream) GetMemoryInputStreamPointer() *C.GMemoryInputStream {
	return self.CPointer
}
func NewTraitMemoryInputStream(p unsafe.Pointer) *TraitMemoryInputStream {
	return &TraitMemoryInputStream{(*C.GMemoryInputStream)(p)}
}

/*
Appends @bytes to data that can be read from the input stream.
*/
func (self *TraitMemoryInputStream) AddBytes(bytes *C.GBytes) {
	C.g_memory_input_stream_add_bytes(self.CPointer, bytes)
	return
}

/*
Appends @data to data that can be read from the input stream
*/
func (self *TraitMemoryInputStream) AddData(data []byte, len_ int64, destroy C.GDestroyNotify) {
	__header__data := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	C.g_memory_input_stream_add_data(self.CPointer, (unsafe.Pointer)(unsafe.Pointer(__header__data.Data)), C.gssize(len_), destroy)
	return
}

type TraitMemoryOutputStream struct{ CPointer *C.GMemoryOutputStream }
type IsMemoryOutputStream interface {
	GetMemoryOutputStreamPointer() *C.GMemoryOutputStream
}

func (self *TraitMemoryOutputStream) GetMemoryOutputStreamPointer() *C.GMemoryOutputStream {
	return self.CPointer
}
func NewTraitMemoryOutputStream(p unsafe.Pointer) *TraitMemoryOutputStream {
	return &TraitMemoryOutputStream{(*C.GMemoryOutputStream)(p)}
}

/*
Gets any loaded data from the @ostream.

Note that the returned pointer may become invalid on the next
write or truncate operation on the stream.
*/
func (self *TraitMemoryOutputStream) GetData() (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_memory_output_stream_get_data(self.CPointer)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Returns the number of bytes from the start up to including the last
byte written in the stream that has not been truncated away.
*/
func (self *TraitMemoryOutputStream) GetDataSize() (return__ int64) {
	var __cgo__return__ C.gsize
	__cgo__return__ = C.g_memory_output_stream_get_data_size(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

/*
Gets the size of the currently allocated data area (available from
g_memory_output_stream_get_data()).

You probably don't want to use this function on resizable streams.
See g_memory_output_stream_get_data_size() instead.  For resizable
streams the size returned by this function is an implementation
detail and may be change at any time in response to operations on the
stream.

If the stream is fixed-sized (ie: no realloc was passed to
g_memory_output_stream_new()) then this is the maximum size of the
stream and further writes will return %G_IO_ERROR_NO_SPACE.

In any case, if you want the number of bytes currently written to the
stream, use g_memory_output_stream_get_data_size().
*/
func (self *TraitMemoryOutputStream) GetSize() (return__ int64) {
	var __cgo__return__ C.gsize
	__cgo__return__ = C.g_memory_output_stream_get_size(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

/*
Returns data from the @ostream as a #GBytes. @ostream must be
closed before calling this function.
*/
func (self *TraitMemoryOutputStream) StealAsBytes() (return__ *C.GBytes) {
	return__ = C.g_memory_output_stream_steal_as_bytes(self.CPointer)
	return
}

/*
Gets any loaded data from the @ostream. Ownership of the data
is transferred to the caller; when no longer needed it must be
freed using the free function set in @ostream's
#GMemoryOutputStream:destroy-function property.

@ostream must be closed before calling this function.
*/
func (self *TraitMemoryOutputStream) StealData() (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_memory_output_stream_steal_data(self.CPointer)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

type TraitMenu struct{ CPointer *C.GMenu }
type IsMenu interface {
	GetMenuPointer() *C.GMenu
}

func (self *TraitMenu) GetMenuPointer() *C.GMenu {
	return self.CPointer
}
func NewTraitMenu(p unsafe.Pointer) *TraitMenu {
	return &TraitMenu{(*C.GMenu)(p)}
}

/*
Convenience function for appending a normal menu item to the end of
@menu.  Combine g_menu_item_new() and g_menu_insert_item() for a more
flexible alternative.
*/
func (self *TraitMenu) Append(label string, detailed_action string) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	__cgo__detailed_action := (*C.gchar)(unsafe.Pointer(C.CString(detailed_action)))
	C.g_menu_append(self.CPointer, __cgo__label, __cgo__detailed_action)
	C.free(unsafe.Pointer(__cgo__label))
	C.free(unsafe.Pointer(__cgo__detailed_action))
	return
}

/*
Appends @item to the end of @menu.

See g_menu_insert_item() for more information.
*/
func (self *TraitMenu) AppendItem(item IsMenuItem) {
	C.g_menu_append_item(self.CPointer, item.GetMenuItemPointer())
	return
}

/*
Convenience function for appending a section menu item to the end of
@menu.  Combine g_menu_item_new_section() and g_menu_insert_item() for a
more flexible alternative.
*/
func (self *TraitMenu) AppendSection(label string, section IsMenuModel) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	C.g_menu_append_section(self.CPointer, __cgo__label, section.GetMenuModelPointer())
	C.free(unsafe.Pointer(__cgo__label))
	return
}

/*
Convenience function for appending a submenu menu item to the end of
@menu.  Combine g_menu_item_new_submenu() and g_menu_insert_item() for a
more flexible alternative.
*/
func (self *TraitMenu) AppendSubmenu(label string, submenu IsMenuModel) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	C.g_menu_append_submenu(self.CPointer, __cgo__label, submenu.GetMenuModelPointer())
	C.free(unsafe.Pointer(__cgo__label))
	return
}

/*
Marks @menu as frozen.

After the menu is frozen, it is an error to attempt to make any
changes to it.  In effect this means that the #GMenu API must no
longer be used.

This function causes g_menu_model_is_mutable() to begin returning
%FALSE, which has some positive performance implications.
*/
func (self *TraitMenu) Freeze() {
	C.g_menu_freeze(self.CPointer)
	return
}

/*
Convenience function for inserting a normal menu item into @menu.
Combine g_menu_item_new() and g_menu_insert_item() for a more flexible
alternative.
*/
func (self *TraitMenu) Insert(position int, label string, detailed_action string) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	__cgo__detailed_action := (*C.gchar)(unsafe.Pointer(C.CString(detailed_action)))
	C.g_menu_insert(self.CPointer, C.gint(position), __cgo__label, __cgo__detailed_action)
	C.free(unsafe.Pointer(__cgo__label))
	C.free(unsafe.Pointer(__cgo__detailed_action))
	return
}

/*
Inserts @item into @menu.

The "insertion" is actually done by copying all of the attribute and
link values of @item and using them to form a new item within @menu.
As such, @item itself is not really inserted, but rather, a menu item
that is exactly the same as the one presently described by @item.

This means that @item is essentially useless after the insertion
occurs.  Any changes you make to it are ignored unless it is inserted
again (at which point its updated values will be copied).

You should probably just free @item once you're done.

There are many convenience functions to take care of common cases.
See g_menu_insert(), g_menu_insert_section() and
g_menu_insert_submenu() as well as "prepend" and "append" variants of
each of these functions.
*/
func (self *TraitMenu) InsertItem(position int, item IsMenuItem) {
	C.g_menu_insert_item(self.CPointer, C.gint(position), item.GetMenuItemPointer())
	return
}

/*
Convenience function for inserting a section menu item into @menu.
Combine g_menu_item_new_section() and g_menu_insert_item() for a more
flexible alternative.
*/
func (self *TraitMenu) InsertSection(position int, label string, section IsMenuModel) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	C.g_menu_insert_section(self.CPointer, C.gint(position), __cgo__label, section.GetMenuModelPointer())
	C.free(unsafe.Pointer(__cgo__label))
	return
}

/*
Convenience function for inserting a submenu menu item into @menu.
Combine g_menu_item_new_submenu() and g_menu_insert_item() for a more
flexible alternative.
*/
func (self *TraitMenu) InsertSubmenu(position int, label string, submenu IsMenuModel) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	C.g_menu_insert_submenu(self.CPointer, C.gint(position), __cgo__label, submenu.GetMenuModelPointer())
	C.free(unsafe.Pointer(__cgo__label))
	return
}

/*
Convenience function for prepending a normal menu item to the start
of @menu.  Combine g_menu_item_new() and g_menu_insert_item() for a more
flexible alternative.
*/
func (self *TraitMenu) Prepend(label string, detailed_action string) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	__cgo__detailed_action := (*C.gchar)(unsafe.Pointer(C.CString(detailed_action)))
	C.g_menu_prepend(self.CPointer, __cgo__label, __cgo__detailed_action)
	C.free(unsafe.Pointer(__cgo__label))
	C.free(unsafe.Pointer(__cgo__detailed_action))
	return
}

/*
Prepends @item to the start of @menu.

See g_menu_insert_item() for more information.
*/
func (self *TraitMenu) PrependItem(item IsMenuItem) {
	C.g_menu_prepend_item(self.CPointer, item.GetMenuItemPointer())
	return
}

/*
Convenience function for prepending a section menu item to the start
of @menu.  Combine g_menu_item_new_section() and g_menu_insert_item() for
a more flexible alternative.
*/
func (self *TraitMenu) PrependSection(label string, section IsMenuModel) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	C.g_menu_prepend_section(self.CPointer, __cgo__label, section.GetMenuModelPointer())
	C.free(unsafe.Pointer(__cgo__label))
	return
}

/*
Convenience function for prepending a submenu menu item to the start
of @menu.  Combine g_menu_item_new_submenu() and g_menu_insert_item() for
a more flexible alternative.
*/
func (self *TraitMenu) PrependSubmenu(label string, submenu IsMenuModel) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	C.g_menu_prepend_submenu(self.CPointer, __cgo__label, submenu.GetMenuModelPointer())
	C.free(unsafe.Pointer(__cgo__label))
	return
}

/*
Removes an item from the menu.

@position gives the index of the item to remove.

It is an error if position is not in range the range from 0 to one
less than the number of items in the menu.

It is not possible to remove items by identity since items are added
to the menu simply by copying their links and attributes (ie:
identity of the item itself is not preserved).
*/
func (self *TraitMenu) Remove(position int) {
	C.g_menu_remove(self.CPointer, C.gint(position))
	return
}

/*
Removes all items in the menu.
*/
func (self *TraitMenu) RemoveAll() {
	C.g_menu_remove_all(self.CPointer)
	return
}

type TraitMenuAttributeIter struct{ CPointer *C.GMenuAttributeIter }
type IsMenuAttributeIter interface {
	GetMenuAttributeIterPointer() *C.GMenuAttributeIter
}

func (self *TraitMenuAttributeIter) GetMenuAttributeIterPointer() *C.GMenuAttributeIter {
	return self.CPointer
}
func NewTraitMenuAttributeIter(p unsafe.Pointer) *TraitMenuAttributeIter {
	return &TraitMenuAttributeIter{(*C.GMenuAttributeIter)(p)}
}

/*
Gets the name of the attribute at the current iterator position, as
a string.

The iterator is not advanced.
*/
func (self *TraitMenuAttributeIter) GetName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_menu_attribute_iter_get_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
This function combines g_menu_attribute_iter_next() with
g_menu_attribute_iter_get_name() and g_menu_attribute_iter_get_value().

First the iterator is advanced to the next (possibly first) attribute.
If that fails, then %FALSE is returned and there are no other
effects.

If successful, @name and @value are set to the name and value of the
attribute that has just been advanced to.  At this point,
g_menu_attribute_iter_get_name() and g_menu_attribute_iter_get_value() will
return the same values again.

The value returned in @name remains valid for as long as the iterator
remains at the current position.  The value returned in @value must
be unreffed using g_variant_unref() when it is no longer in use.
*/
func (self *TraitMenuAttributeIter) GetNext() (out_name string, value *C.GVariant, return__ bool) {
	var __cgo__out_name *C.gchar
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_menu_attribute_iter_get_next(self.CPointer, &__cgo__out_name, &value)
	out_name = C.GoString((*C.char)(unsafe.Pointer(__cgo__out_name)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value of the attribute at the current iterator position.

The iterator is not advanced.
*/
func (self *TraitMenuAttributeIter) GetValue() (return__ *C.GVariant) {
	return__ = C.g_menu_attribute_iter_get_value(self.CPointer)
	return
}

/*
Attempts to advance the iterator to the next (possibly first)
attribute.

%TRUE is returned on success, or %FALSE if there are no more
attributes.

You must call this function when you first acquire the iterator
to advance it to the first attribute (and determine if the first
attribute exists at all).
*/
func (self *TraitMenuAttributeIter) Next() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_menu_attribute_iter_next(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

type TraitMenuItem struct{ CPointer *C.GMenuItem }
type IsMenuItem interface {
	GetMenuItemPointer() *C.GMenuItem
}

func (self *TraitMenuItem) GetMenuItemPointer() *C.GMenuItem {
	return self.CPointer
}
func NewTraitMenuItem(p unsafe.Pointer) *TraitMenuItem {
	return &TraitMenuItem{(*C.GMenuItem)(p)}
}

// g_menu_item_get_attribute is not generated due to varargs

/*
Queries the named @attribute on @menu_item.

If @expected_type is specified and the attribute does not have this
type, %NULL is returned.  %NULL is also returned if the attribute
simply does not exist.
*/
func (self *TraitMenuItem) GetAttributeValue(attribute string, expected_type *C.GVariantType) (return__ *C.GVariant) {
	__cgo__attribute := (*C.gchar)(unsafe.Pointer(C.CString(attribute)))
	return__ = C.g_menu_item_get_attribute_value(self.CPointer, __cgo__attribute, expected_type)
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Queries the named @link on @menu_item.
*/
func (self *TraitMenuItem) GetLink(link string) (return__ *MenuModel) {
	__cgo__link := (*C.gchar)(unsafe.Pointer(C.CString(link)))
	var __cgo__return__ *C.GMenuModel
	__cgo__return__ = C.g_menu_item_get_link(self.CPointer, __cgo__link)
	C.free(unsafe.Pointer(__cgo__link))
	if __cgo__return__ != nil {
		return__ = NewMenuModelFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// g_menu_item_set_action_and_target is not generated due to varargs

/*
Sets or unsets the "action" and "target" attributes of @menu_item.

If @action is %NULL then both the "action" and "target" attributes
are unset (and @target_value is ignored).

If @action is non-%NULL then the "action" attribute is set.  The
"target" attribute is then set to the value of @target_value if it is
non-%NULL or unset otherwise.

Normal menu items (ie: not submenu, section or other custom item
types) are expected to have the "action" attribute set to identify
the action that they are associated with.  The state type of the
action help to determine the disposition of the menu item.  See
#GAction and #GActionGroup for an overview of actions.

In general, clicking on the menu item will result in activation of
the named action with the "target" attribute given as the parameter
to the action invocation.  If the "target" attribute is not set then
the action is invoked with no parameter.

If the action has no state then the menu item is usually drawn as a
plain menu item (ie: with no additional decoration).

If the action has a boolean state then the menu item is usually drawn
as a toggle menu item (ie: with a checkmark or equivalent
indication).  The item should be marked as 'toggled' or 'checked'
when the boolean state is %TRUE.

If the action has a string state then the menu item is usually drawn
as a radio menu item (ie: with a radio bullet or equivalent
indication).  The item should be marked as 'selected' when the string
state is equal to the value of the @target property.

See g_menu_item_set_action_and_target() or
g_menu_item_set_detailed_action() for two equivalent calls that are
probably more convenient for most uses.
*/
func (self *TraitMenuItem) SetActionAndTargetValue(action string, target_value *C.GVariant) {
	__cgo__action := (*C.gchar)(unsafe.Pointer(C.CString(action)))
	C.g_menu_item_set_action_and_target_value(self.CPointer, __cgo__action, target_value)
	C.free(unsafe.Pointer(__cgo__action))
	return
}

// g_menu_item_set_attribute is not generated due to varargs

/*
Sets or unsets an attribute on @menu_item.

The attribute to set or unset is specified by @attribute. This
can be one of the standard attribute names %G_MENU_ATTRIBUTE_LABEL,
%G_MENU_ATTRIBUTE_ACTION, %G_MENU_ATTRIBUTE_TARGET, or a custom
attribute name.
Attribute names are restricted to lowercase characters, numbers
and '-'. Furthermore, the names must begin with a lowercase character,
must not end with a '-', and must not contain consecutive dashes.

must consist only of lowercase
ASCII characters, digits and '-'.

If @value is non-%NULL then it is used as the new value for the
attribute.  If @value is %NULL then the attribute is unset. If
the @value #GVariant is floating, it is consumed.

See also g_menu_item_set_attribute() for a more convenient way to do
the same.
*/
func (self *TraitMenuItem) SetAttributeValue(attribute string, value *C.GVariant) {
	__cgo__attribute := (*C.gchar)(unsafe.Pointer(C.CString(attribute)))
	C.g_menu_item_set_attribute_value(self.CPointer, __cgo__attribute, value)
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Sets the "action" and possibly the "target" attribute of @menu_item.

The format of @detailed_action is the same format parsed by
g_action_parse_detailed_name().

See g_menu_item_set_action_and_target() or
g_menu_item_set_action_and_target_value() for more flexible (but
slightly less convenient) alternatives.

See also g_menu_item_set_action_and_target_value() for a description of
the semantics of the action and target attributes.
*/
func (self *TraitMenuItem) SetDetailedAction(detailed_action string) {
	__cgo__detailed_action := (*C.gchar)(unsafe.Pointer(C.CString(detailed_action)))
	C.g_menu_item_set_detailed_action(self.CPointer, __cgo__detailed_action)
	C.free(unsafe.Pointer(__cgo__detailed_action))
	return
}

/*
Sets (or unsets) the icon on @menu_item.

This call is the same as calling g_icon_serialize() and using the
result as the value to g_menu_item_set_attribute_value() for
%G_MENU_ATTRIBUTE_ICON.

This API is only intended for use with "noun" menu items; things like
bookmarks or applications in an "Open With" menu.  Don't use it on
menu items corresponding to verbs (eg: stock icons for 'Save' or
'Quit').

If @icon is %NULL then the icon is unset.
*/
func (self *TraitMenuItem) SetIcon(icon *C.GIcon) {
	C.g_menu_item_set_icon(self.CPointer, icon)
	return
}

/*
Sets or unsets the "label" attribute of @menu_item.

If @label is non-%NULL it is used as the label for the menu item.  If
it is %NULL then the label attribute is unset.
*/
func (self *TraitMenuItem) SetLabel(label string) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	C.g_menu_item_set_label(self.CPointer, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	return
}

/*
Creates a link from @menu_item to @model if non-%NULL, or unsets it.

Links are used to establish a relationship between a particular menu
item and another menu.  For example, %G_MENU_LINK_SUBMENU is used to
associate a submenu with a particular menu item, and %G_MENU_LINK_SECTION
is used to create a section. Other types of link can be used, but there
is no guarantee that clients will be able to make sense of them.
Link types are restricted to lowercase characters, numbers
and '-'. Furthermore, the names must begin with a lowercase character,
must not end with a '-', and must not contain consecutive dashes.
*/
func (self *TraitMenuItem) SetLink(link string, model IsMenuModel) {
	__cgo__link := (*C.gchar)(unsafe.Pointer(C.CString(link)))
	C.g_menu_item_set_link(self.CPointer, __cgo__link, model.GetMenuModelPointer())
	C.free(unsafe.Pointer(__cgo__link))
	return
}

/*
Sets or unsets the "section" link of @menu_item to @section.

The effect of having one menu appear as a section of another is
exactly as it sounds: the items from @section become a direct part of
the menu that @menu_item is added to.  See g_menu_item_new_section()
for more information about what it means for a menu item to be a
section.
*/
func (self *TraitMenuItem) SetSection(section IsMenuModel) {
	C.g_menu_item_set_section(self.CPointer, section.GetMenuModelPointer())
	return
}

/*
Sets or unsets the "submenu" link of @menu_item to @submenu.

If @submenu is non-%NULL, it is linked to.  If it is %NULL then the
link is unset.

The effect of having one menu appear as a submenu of another is
exactly as it sounds.
*/
func (self *TraitMenuItem) SetSubmenu(submenu IsMenuModel) {
	C.g_menu_item_set_submenu(self.CPointer, submenu.GetMenuModelPointer())
	return
}

type TraitMenuLinkIter struct{ CPointer *C.GMenuLinkIter }
type IsMenuLinkIter interface {
	GetMenuLinkIterPointer() *C.GMenuLinkIter
}

func (self *TraitMenuLinkIter) GetMenuLinkIterPointer() *C.GMenuLinkIter {
	return self.CPointer
}
func NewTraitMenuLinkIter(p unsafe.Pointer) *TraitMenuLinkIter {
	return &TraitMenuLinkIter{(*C.GMenuLinkIter)(p)}
}

/*
Gets the name of the link at the current iterator position.

The iterator is not advanced.
*/
func (self *TraitMenuLinkIter) GetName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_menu_link_iter_get_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
This function combines g_menu_link_iter_next() with
g_menu_link_iter_get_name() and g_menu_link_iter_get_value().

First the iterator is advanced to the next (possibly first) link.
If that fails, then %FALSE is returned and there are no other effects.

If successful, @out_link and @value are set to the name and #GMenuModel
of the link that has just been advanced to.  At this point,
g_menu_link_iter_get_name() and g_menu_link_iter_get_value() will return the
same values again.

The value returned in @out_link remains valid for as long as the iterator
remains at the current position.  The value returned in @value must
be unreffed using g_object_unref() when it is no longer in use.
*/
func (self *TraitMenuLinkIter) GetNext() (out_link string, value *MenuModel, return__ bool) {
	var __cgo__out_link *C.gchar
	var __cgo__value *C.GMenuModel
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_menu_link_iter_get_next(self.CPointer, &__cgo__out_link, &__cgo__value)
	out_link = C.GoString((*C.char)(unsafe.Pointer(__cgo__out_link)))
	if __cgo__value != nil {
		value = NewMenuModelFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__value).Pointer()))
	}
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the linked #GMenuModel at the current iterator position.

The iterator is not advanced.
*/
func (self *TraitMenuLinkIter) GetValue() (return__ *MenuModel) {
	var __cgo__return__ *C.GMenuModel
	__cgo__return__ = C.g_menu_link_iter_get_value(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewMenuModelFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Attempts to advance the iterator to the next (possibly first)
link.

%TRUE is returned on success, or %FALSE if there are no more links.

You must call this function when you first acquire the iterator to
advance it to the first link (and determine if the first link exists
at all).
*/
func (self *TraitMenuLinkIter) Next() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_menu_link_iter_next(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

type TraitMenuModel struct{ CPointer *C.GMenuModel }
type IsMenuModel interface {
	GetMenuModelPointer() *C.GMenuModel
}

func (self *TraitMenuModel) GetMenuModelPointer() *C.GMenuModel {
	return self.CPointer
}
func NewTraitMenuModel(p unsafe.Pointer) *TraitMenuModel {
	return &TraitMenuModel{(*C.GMenuModel)(p)}
}

// g_menu_model_get_item_attribute is not generated due to varargs

/*
Queries the item at position @item_index in @model for the attribute
specified by @attribute.

If @expected_type is non-%NULL then it specifies the expected type of
the attribute.  If it is %NULL then any type will be accepted.

If the attribute exists and matches @expected_type (or if the
expected type is unspecified) then the value is returned.

If the attribute does not exist, or does not match the expected type
then %NULL is returned.
*/
func (self *TraitMenuModel) GetItemAttributeValue(item_index int, attribute string, expected_type *C.GVariantType) (return__ *C.GVariant) {
	__cgo__attribute := (*C.gchar)(unsafe.Pointer(C.CString(attribute)))
	return__ = C.g_menu_model_get_item_attribute_value(self.CPointer, C.gint(item_index), __cgo__attribute, expected_type)
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Queries the item at position @item_index in @model for the link
specified by @link.

If the link exists, the linked #GMenuModel is returned.  If the link
does not exist, %NULL is returned.
*/
func (self *TraitMenuModel) GetItemLink(item_index int, link string) (return__ *MenuModel) {
	__cgo__link := (*C.gchar)(unsafe.Pointer(C.CString(link)))
	var __cgo__return__ *C.GMenuModel
	__cgo__return__ = C.g_menu_model_get_item_link(self.CPointer, C.gint(item_index), __cgo__link)
	C.free(unsafe.Pointer(__cgo__link))
	if __cgo__return__ != nil {
		return__ = NewMenuModelFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Query the number of items in @model.
*/
func (self *TraitMenuModel) GetNItems() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_menu_model_get_n_items(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Queries if @model is mutable.

An immutable #GMenuModel will never emit the #GMenuModel::items-changed
signal. Consumers of the model may make optimisations accordingly.
*/
func (self *TraitMenuModel) IsMutable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_menu_model_is_mutable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Requests emission of the #GMenuModel::items-changed signal on @model.

This function should never be called except by #GMenuModel
subclasses.  Any other calls to this function will very likely lead
to a violation of the interface of the model.

The implementation should update its internal representation of the
menu before emitting the signal.  The implementation should further
expect to receive queries about the new state of the menu (and
particularly added menu items) while signal handlers are running.

The implementation must dispatch this call directly from a mainloop
entry and not in response to calls -- particularly those from the
#GMenuModel API.  Said another way: the menu must not change while
user code is running without returning to the mainloop.
*/
func (self *TraitMenuModel) ItemsChanged(position int, removed int, added int) {
	C.g_menu_model_items_changed(self.CPointer, C.gint(position), C.gint(removed), C.gint(added))
	return
}

/*
Creates a #GMenuAttributeIter to iterate over the attributes of
the item at position @item_index in @model.

You must free the iterator with g_object_unref() when you are done.
*/
func (self *TraitMenuModel) IterateItemAttributes(item_index int) (return__ *MenuAttributeIter) {
	var __cgo__return__ *C.GMenuAttributeIter
	__cgo__return__ = C.g_menu_model_iterate_item_attributes(self.CPointer, C.gint(item_index))
	if __cgo__return__ != nil {
		return__ = NewMenuAttributeIterFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GMenuLinkIter to iterate over the links of the item at
position @item_index in @model.

You must free the iterator with g_object_unref() when you are done.
*/
func (self *TraitMenuModel) IterateItemLinks(item_index int) (return__ *MenuLinkIter) {
	var __cgo__return__ *C.GMenuLinkIter
	__cgo__return__ = C.g_menu_model_iterate_item_links(self.CPointer, C.gint(item_index))
	if __cgo__return__ != nil {
		return__ = NewMenuLinkIterFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

type TraitMountOperation struct{ CPointer *C.GMountOperation }
type IsMountOperation interface {
	GetMountOperationPointer() *C.GMountOperation
}

func (self *TraitMountOperation) GetMountOperationPointer() *C.GMountOperation {
	return self.CPointer
}
func NewTraitMountOperation(p unsafe.Pointer) *TraitMountOperation {
	return &TraitMountOperation{(*C.GMountOperation)(p)}
}

/*
Check to see whether the mount operation is being used
for an anonymous user.
*/
func (self *TraitMountOperation) GetAnonymous() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_mount_operation_get_anonymous(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets a choice from the mount operation.
*/
func (self *TraitMountOperation) GetChoice() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.g_mount_operation_get_choice(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the domain of the mount operation.
*/
func (self *TraitMountOperation) GetDomain() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_mount_operation_get_domain(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Gets a password from the mount operation.
*/
func (self *TraitMountOperation) GetPassword() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_mount_operation_get_password(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Gets the state of saving passwords for the mount operation.
*/
func (self *TraitMountOperation) GetPasswordSave() (return__ C.GPasswordSave) {
	return__ = C.g_mount_operation_get_password_save(self.CPointer)
	return
}

/*
Get the user name from the mount operation.
*/
func (self *TraitMountOperation) GetUsername() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_mount_operation_get_username(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Emits the #GMountOperation::reply signal.
*/
func (self *TraitMountOperation) Reply(result C.GMountOperationResult) {
	C.g_mount_operation_reply(self.CPointer, result)
	return
}

/*
Sets the mount operation to use an anonymous user if @anonymous is %TRUE.
*/
func (self *TraitMountOperation) SetAnonymous(anonymous bool) {
	__cgo__anonymous := C.gboolean(0)
	if anonymous {
		__cgo__anonymous = C.gboolean(1)
	}
	C.g_mount_operation_set_anonymous(self.CPointer, __cgo__anonymous)
	return
}

/*
Sets a default choice for the mount operation.
*/
func (self *TraitMountOperation) SetChoice(choice int) {
	C.g_mount_operation_set_choice(self.CPointer, C.int(choice))
	return
}

/*
Sets the mount operation's domain.
*/
func (self *TraitMountOperation) SetDomain(domain string) {
	__cgo__domain := C.CString(domain)
	C.g_mount_operation_set_domain(self.CPointer, __cgo__domain)
	C.free(unsafe.Pointer(__cgo__domain))
	return
}

/*
Sets the mount operation's password to @password.
*/
func (self *TraitMountOperation) SetPassword(password string) {
	__cgo__password := C.CString(password)
	C.g_mount_operation_set_password(self.CPointer, __cgo__password)
	C.free(unsafe.Pointer(__cgo__password))
	return
}

/*
Sets the state of saving passwords for the mount operation.
*/
func (self *TraitMountOperation) SetPasswordSave(save C.GPasswordSave) {
	C.g_mount_operation_set_password_save(self.CPointer, save)
	return
}

/*
Sets the user name within @op to @username.
*/
func (self *TraitMountOperation) SetUsername(username string) {
	__cgo__username := C.CString(username)
	C.g_mount_operation_set_username(self.CPointer, __cgo__username)
	C.free(unsafe.Pointer(__cgo__username))
	return
}

type TraitNativeVolumeMonitor struct{ CPointer *C.GNativeVolumeMonitor }
type IsNativeVolumeMonitor interface {
	GetNativeVolumeMonitorPointer() *C.GNativeVolumeMonitor
}

func (self *TraitNativeVolumeMonitor) GetNativeVolumeMonitorPointer() *C.GNativeVolumeMonitor {
	return self.CPointer
}
func NewTraitNativeVolumeMonitor(p unsafe.Pointer) *TraitNativeVolumeMonitor {
	return &TraitNativeVolumeMonitor{(*C.GNativeVolumeMonitor)(p)}
}

type TraitNetworkAddress struct{ CPointer *C.GNetworkAddress }
type IsNetworkAddress interface {
	GetNetworkAddressPointer() *C.GNetworkAddress
}

func (self *TraitNetworkAddress) GetNetworkAddressPointer() *C.GNetworkAddress {
	return self.CPointer
}
func NewTraitNetworkAddress(p unsafe.Pointer) *TraitNetworkAddress {
	return &TraitNetworkAddress{(*C.GNetworkAddress)(p)}
}

/*
Gets @addr's hostname. This might be either UTF-8 or ASCII-encoded,
depending on what @addr was created with.
*/
func (self *TraitNetworkAddress) GetHostname() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_network_address_get_hostname(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets @addr's port number
*/
func (self *TraitNetworkAddress) GetPort() (return__ uint16) {
	var __cgo__return__ C.guint16
	__cgo__return__ = C.g_network_address_get_port(self.CPointer)
	return__ = uint16(__cgo__return__)
	return
}

/*
Gets @addr's scheme
*/
func (self *TraitNetworkAddress) GetScheme() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_network_address_get_scheme(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

type TraitNetworkService struct{ CPointer *C.GNetworkService }
type IsNetworkService interface {
	GetNetworkServicePointer() *C.GNetworkService
}

func (self *TraitNetworkService) GetNetworkServicePointer() *C.GNetworkService {
	return self.CPointer
}
func NewTraitNetworkService(p unsafe.Pointer) *TraitNetworkService {
	return &TraitNetworkService{(*C.GNetworkService)(p)}
}

/*
Gets the domain that @srv serves. This might be either UTF-8 or
ASCII-encoded, depending on what @srv was created with.
*/
func (self *TraitNetworkService) GetDomain() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_network_service_get_domain(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets @srv's protocol name (eg, "tcp").
*/
func (self *TraitNetworkService) GetProtocol() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_network_service_get_protocol(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Get's the URI scheme used to resolve proxies. By default, the service name
is used as scheme.
*/
func (self *TraitNetworkService) GetScheme() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_network_service_get_scheme(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets @srv's service name (eg, "ldap").
*/
func (self *TraitNetworkService) GetService() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_network_service_get_service(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Set's the URI scheme used to resolve proxies. By default, the service name
is used as scheme.
*/
func (self *TraitNetworkService) SetScheme(scheme string) {
	__cgo__scheme := (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	C.g_network_service_set_scheme(self.CPointer, __cgo__scheme)
	C.free(unsafe.Pointer(__cgo__scheme))
	return
}

type TraitNotification struct{ CPointer *C.GNotification }
type IsNotification interface {
	GetNotificationPointer() *C.GNotification
}

func (self *TraitNotification) GetNotificationPointer() *C.GNotification {
	return self.CPointer
}
func NewTraitNotification(p unsafe.Pointer) *TraitNotification {
	return &TraitNotification{(*C.GNotification)(p)}
}

/*
Adds a button to @notification that activates the action in
@detailed_action when clicked. That action must be an
application-wide action (starting with "app."). If @detailed_action
contains a target, the action will be activated with that target as
its parameter.

See g_action_parse_detailed_name() for a description of the format
for @detailed_action.
*/
func (self *TraitNotification) AddButton(label string, detailed_action string) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	__cgo__detailed_action := (*C.gchar)(unsafe.Pointer(C.CString(detailed_action)))
	C.g_notification_add_button(self.CPointer, __cgo__label, __cgo__detailed_action)
	C.free(unsafe.Pointer(__cgo__label))
	C.free(unsafe.Pointer(__cgo__detailed_action))
	return
}

// g_notification_add_button_with_target is not generated due to varargs

/*
Adds a button to @notification that activates @action when clicked.
@action must be an application-wide action (it must start with "app.").

If @target is non-%NULL, @action will be activated with @target as
its parameter.
*/
func (self *TraitNotification) AddButtonWithTargetValue(label string, action string, target *C.GVariant) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	__cgo__action := (*C.gchar)(unsafe.Pointer(C.CString(action)))
	C.g_notification_add_button_with_target_value(self.CPointer, __cgo__label, __cgo__action, target)
	C.free(unsafe.Pointer(__cgo__label))
	C.free(unsafe.Pointer(__cgo__action))
	return
}

/*
Sets the body of @notification to @body.
*/
func (self *TraitNotification) SetBody(body string) {
	__cgo__body := (*C.gchar)(unsafe.Pointer(C.CString(body)))
	C.g_notification_set_body(self.CPointer, __cgo__body)
	C.free(unsafe.Pointer(__cgo__body))
	return
}

/*
Sets the default action of @notification to @detailed_action. This
action is activated when the notification is clicked on.

The action in @detailed_action must be an application-wide action (it
must start with "app."). If @detailed_action contains a target, the
given action will be activated with that target as its parameter.
See g_action_parse_detailed_name() for a description of the format
for @detailed_action.

When no default action is set, the application that the notification
was sent on is activated.
*/
func (self *TraitNotification) SetDefaultAction(detailed_action string) {
	__cgo__detailed_action := (*C.gchar)(unsafe.Pointer(C.CString(detailed_action)))
	C.g_notification_set_default_action(self.CPointer, __cgo__detailed_action)
	C.free(unsafe.Pointer(__cgo__detailed_action))
	return
}

// g_notification_set_default_action_and_target is not generated due to varargs

/*
Sets the default action of @notification to @action. This action is
activated when the notification is clicked on. It must be an
application-wide action (start with "app.").

If @target_format is given, it is used to collect remaining
positional parameters into a GVariant instance, similar to
g_variant_new().

If @target is non-%NULL, @action will be activated with @target as
its parameter.

When no default action is set, the application that the notification
was sent on is activated.
*/
func (self *TraitNotification) SetDefaultActionAndTargetValue(action string, target *C.GVariant) {
	__cgo__action := (*C.gchar)(unsafe.Pointer(C.CString(action)))
	C.g_notification_set_default_action_and_target_value(self.CPointer, __cgo__action, target)
	C.free(unsafe.Pointer(__cgo__action))
	return
}

/*
Sets the icon of @notification to @icon.
*/
func (self *TraitNotification) SetIcon(icon *C.GIcon) {
	C.g_notification_set_icon(self.CPointer, icon)
	return
}

/*
Sets the title of @notification to @title.
*/
func (self *TraitNotification) SetTitle(title string) {
	__cgo__title := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	C.g_notification_set_title(self.CPointer, __cgo__title)
	C.free(unsafe.Pointer(__cgo__title))
	return
}

/*
Sets or unsets whether @notification is marked as urgent.
*/
func (self *TraitNotification) SetUrgent(urgent bool) {
	__cgo__urgent := C.gboolean(0)
	if urgent {
		__cgo__urgent = C.gboolean(1)
	}
	C.g_notification_set_urgent(self.CPointer, __cgo__urgent)
	return
}

type TraitOutputStream struct{ CPointer *C.GOutputStream }
type IsOutputStream interface {
	GetOutputStreamPointer() *C.GOutputStream
}

func (self *TraitOutputStream) GetOutputStreamPointer() *C.GOutputStream {
	return self.CPointer
}
func NewTraitOutputStream(p unsafe.Pointer) *TraitOutputStream {
	return &TraitOutputStream{(*C.GOutputStream)(p)}
}

/*
Clears the pending flag on @stream.
*/
func (self *TraitOutputStream) ClearPending() {
	C.g_output_stream_clear_pending(self.CPointer)
	return
}

/*
Closes the stream, releasing resources related to it.

Once the stream is closed, all other operations will return %G_IO_ERROR_CLOSED.
Closing a stream multiple times will not return an error.

Closing a stream will automatically flush any outstanding buffers in the
stream.

Streams will be automatically closed when the last reference
is dropped, but you might want to call this function to make sure
resources are released as early as possible.

Some streams might keep the backing store of the stream (e.g. a file descriptor)
open after the stream is closed. See the documentation for the individual
stream for details.

On failure the first error that happened will be reported, but the close
operation will finish as much as possible. A stream that failed to
close will still return %G_IO_ERROR_CLOSED for all operations. Still, it
is important to check and report the error to the user, otherwise
there might be a loss of data as all data might not be written.

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
Cancelling a close will still leave the stream closed, but there some streams
can use a faster close that doesn't block to e.g. check errors. On
cancellation (as with any error) there is no guarantee that all written
data will reach the target.
*/
func (self *TraitOutputStream) Close(cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_output_stream_close(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Requests an asynchronous close of the stream, releasing resources
related to it. When the operation is finished @callback will be
called. You can then call g_output_stream_close_finish() to get
the result of the operation.

For behaviour details see g_output_stream_close().

The asyncronous methods have a default fallback that uses threads
to implement asynchronicity, so they are optional for inheriting
classes. However, if you override one you must override all.
*/
func (self *TraitOutputStream) CloseAsync(io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_output_stream_close_async(self.CPointer, C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Closes an output stream.
*/
func (self *TraitOutputStream) CloseFinish(result *C.GAsyncResult) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_output_stream_close_finish(self.CPointer, result, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Forces a write of all user-space buffered data for the given
@stream. Will block during the operation. Closing the stream will
implicitly cause a flush.

This function is optional for inherited classes.

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
*/
func (self *TraitOutputStream) Flush(cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_output_stream_flush(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Forces an asynchronous write of all user-space buffered data for
the given @stream.
For behaviour details see g_output_stream_flush().

When the operation is finished @callback will be
called. You can then call g_output_stream_flush_finish() to get the
result of the operation.
*/
func (self *TraitOutputStream) FlushAsync(io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_output_stream_flush_async(self.CPointer, C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes flushing an output stream.
*/
func (self *TraitOutputStream) FlushFinish(result *C.GAsyncResult) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_output_stream_flush_finish(self.CPointer, result, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Checks if an ouput stream has pending actions.
*/
func (self *TraitOutputStream) HasPending() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_output_stream_has_pending(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if an output stream has already been closed.
*/
func (self *TraitOutputStream) IsClosed() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_output_stream_is_closed(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if an output stream is being closed. This can be
used inside e.g. a flush implementation to see if the
flush (or other i/o operation) is called from within
the closing operation.
*/
func (self *TraitOutputStream) IsClosing() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_output_stream_is_closing(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// g_output_stream_printf is not generated due to varargs

/*
Sets @stream to have actions pending. If the pending flag is
already set or @stream is closed, it will return %FALSE and set
@error.
*/
func (self *TraitOutputStream) SetPending() (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_output_stream_set_pending(self.CPointer, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Splices an input stream into an output stream.
*/
func (self *TraitOutputStream) Splice(source IsInputStream, flags C.GOutputStreamSpliceFlags, cancellable IsCancellable) (return__ int64, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_output_stream_splice(self.CPointer, source.GetInputStreamPointer(), flags, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Splices a stream asynchronously.
When the operation is finished @callback will be called.
You can then call g_output_stream_splice_finish() to get the
result of the operation.

For the synchronous, blocking version of this function, see
g_output_stream_splice().
*/
func (self *TraitOutputStream) SpliceAsync(source IsInputStream, flags C.GOutputStreamSpliceFlags, io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_output_stream_splice_async(self.CPointer, source.GetInputStreamPointer(), flags, C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes an asynchronous stream splice operation.
*/
func (self *TraitOutputStream) SpliceFinish(result *C.GAsyncResult) (return__ int64, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_output_stream_splice_finish(self.CPointer, result, &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

// g_output_stream_vprintf is not generated due to varargs

/*
Tries to write @count bytes from @buffer into the stream. Will block
during the operation.

If count is 0, returns 0 and does nothing. A value of @count
larger than %G_MAXSSIZE will cause a %G_IO_ERROR_INVALID_ARGUMENT error.

On success, the number of bytes written to the stream is returned.
It is not an error if this is not the same as the requested size, as it
can happen e.g. on a partial I/O error, or if there is not enough
storage in the stream. All writes block until at least one byte
is written or an error occurs; 0 is never returned (unless
@count is 0).

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
operation was partially finished when the operation was cancelled the
partial result will be returned, without an error.

On error -1 is returned and @error is set accordingly.
*/
func (self *TraitOutputStream) Write(buffer []byte, count int64, cancellable IsCancellable) (return__ int64, __err__ error) {
	__header__buffer := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_output_stream_write(self.CPointer, (unsafe.Pointer)(unsafe.Pointer(__header__buffer.Data)), C.gsize(count), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Tries to write @count bytes from @buffer into the stream. Will block
during the operation.

This function is similar to g_output_stream_write(), except it tries to
write as many bytes as requested, only stopping on an error.

On a successful write of @count bytes, %TRUE is returned, and @bytes_written
is set to @count.

If there is an error during the operation %FALSE is returned and @error
is set to indicate the error status, @bytes_written is updated to contain
the number of bytes written into the stream before the error occurred.
*/
func (self *TraitOutputStream) WriteAll(buffer []byte, count int64, cancellable IsCancellable) (bytes_written int64, return__ bool, __err__ error) {
	__header__buffer := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	var __cgo__bytes_written C.gsize
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_output_stream_write_all(self.CPointer, (unsafe.Pointer)(unsafe.Pointer(__header__buffer.Data)), C.gsize(count), &__cgo__bytes_written, cancellable.GetCancellablePointer(), &__cgo_error__)
	bytes_written = int64(__cgo__bytes_written)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Request an asynchronous write of @count bytes from @buffer into
the stream. When the operation is finished @callback will be called.
You can then call g_output_stream_write_finish() to get the result of the
operation.

During an async request no other sync and async calls are allowed,
and will result in %G_IO_ERROR_PENDING errors.

A value of @count larger than %G_MAXSSIZE will cause a
%G_IO_ERROR_INVALID_ARGUMENT error.

On success, the number of bytes written will be passed to the
@callback. It is not an error if this is not the same as the
requested size, as it can happen e.g. on a partial I/O error,
but generally we try to write as many bytes as requested.

You are guaranteed that this method will never fail with
%G_IO_ERROR_WOULD_BLOCK - if @stream can't accept more data, the
method will just wait until this changes.

Any outstanding I/O request with higher priority (lower numerical
value) will be executed before an outstanding request with lower
priority. Default priority is %G_PRIORITY_DEFAULT.

The asyncronous methods have a default fallback that uses threads
to implement asynchronicity, so they are optional for inheriting
classes. However, if you override one you must override all.

For the synchronous, blocking version of this function, see
g_output_stream_write().

Note that no copy of @buffer will be made, so it must stay valid
until @callback is called. See g_output_stream_write_bytes_async()
for a #GBytes version that will automatically hold a reference to
the contents (without copying) for the duration of the call.
*/
func (self *TraitOutputStream) WriteAsync(buffer []byte, count int64, io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__header__buffer := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	C.g_output_stream_write_async(self.CPointer, (unsafe.Pointer)(unsafe.Pointer(__header__buffer.Data)), C.gsize(count), C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
A wrapper function for g_output_stream_write() which takes a
#GBytes as input.  This can be more convenient for use by language
bindings or in other cases where the refcounted nature of #GBytes
is helpful over a bare pointer interface.

However, note that this function may still perform partial writes,
just like g_output_stream_write().  If that occurs, to continue
writing, you will need to create a new #GBytes containing just the
remaining bytes, using g_bytes_new_from_bytes(). Passing the same
#GBytes instance multiple times potentially can result in duplicated
data in the output stream.
*/
func (self *TraitOutputStream) WriteBytes(bytes *C.GBytes, cancellable IsCancellable) (return__ int64, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_output_stream_write_bytes(self.CPointer, bytes, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This function is similar to g_output_stream_write_async(), but
takes a #GBytes as input.  Due to the refcounted nature of #GBytes,
this allows the stream to avoid taking a copy of the data.

However, note that this function may still perform partial writes,
just like g_output_stream_write_async(). If that occurs, to continue
writing, you will need to create a new #GBytes containing just the
remaining bytes, using g_bytes_new_from_bytes(). Passing the same
#GBytes instance multiple times potentially can result in duplicated
data in the output stream.

For the synchronous, blocking version of this function, see
g_output_stream_write_bytes().
*/
func (self *TraitOutputStream) WriteBytesAsync(bytes *C.GBytes, io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_output_stream_write_bytes_async(self.CPointer, bytes, C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes a stream write-from-#GBytes operation.
*/
func (self *TraitOutputStream) WriteBytesFinish(result *C.GAsyncResult) (return__ int64, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_output_stream_write_bytes_finish(self.CPointer, result, &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Finishes a stream write operation.
*/
func (self *TraitOutputStream) WriteFinish(result *C.GAsyncResult) (return__ int64, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_output_stream_write_finish(self.CPointer, result, &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

type TraitPermission struct{ CPointer *C.GPermission }
type IsPermission interface {
	GetPermissionPointer() *C.GPermission
}

func (self *TraitPermission) GetPermissionPointer() *C.GPermission {
	return self.CPointer
}
func NewTraitPermission(p unsafe.Pointer) *TraitPermission {
	return &TraitPermission{(*C.GPermission)(p)}
}

/*
Attempts to acquire the permission represented by @permission.

The precise method by which this happens depends on the permission
and the underlying authentication mechanism.  A simple example is
that a dialog may appear asking the user to enter their password.

You should check with g_permission_get_can_acquire() before calling
this function.

If the permission is acquired then %TRUE is returned.  Otherwise,
%FALSE is returned and @error is set appropriately.

This call is blocking, likely for a very long time (in the case that
user interaction is required).  See g_permission_acquire_async() for
the non-blocking version.
*/
func (self *TraitPermission) Acquire(cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_permission_acquire(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Attempts to acquire the permission represented by @permission.

This is the first half of the asynchronous version of
g_permission_acquire().
*/
func (self *TraitPermission) AcquireAsync(cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_permission_acquire_async(self.CPointer, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Collects the result of attempting to acquire the permission
represented by @permission.

This is the second half of the asynchronous version of
g_permission_acquire().
*/
func (self *TraitPermission) AcquireFinish(result *C.GAsyncResult) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_permission_acquire_finish(self.CPointer, result, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the value of the 'allowed' property.  This property is %TRUE if
the caller currently has permission to perform the action that
@permission represents the permission to perform.
*/
func (self *TraitPermission) GetAllowed() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_permission_get_allowed(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value of the 'can-acquire' property.  This property is %TRUE
if it is generally possible to acquire the permission by calling
g_permission_acquire().
*/
func (self *TraitPermission) GetCanAcquire() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_permission_get_can_acquire(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value of the 'can-release' property.  This property is %TRUE
if it is generally possible to release the permission by calling
g_permission_release().
*/
func (self *TraitPermission) GetCanRelease() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_permission_get_can_release(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This function is called by the #GPermission implementation to update
the properties of the permission.  You should never call this
function except from a #GPermission implementation.

GObject notify signals are generated, as appropriate.
*/
func (self *TraitPermission) ImplUpdate(allowed bool, can_acquire bool, can_release bool) {
	__cgo__allowed := C.gboolean(0)
	if allowed {
		__cgo__allowed = C.gboolean(1)
	}
	__cgo__can_acquire := C.gboolean(0)
	if can_acquire {
		__cgo__can_acquire = C.gboolean(1)
	}
	__cgo__can_release := C.gboolean(0)
	if can_release {
		__cgo__can_release = C.gboolean(1)
	}
	C.g_permission_impl_update(self.CPointer, __cgo__allowed, __cgo__can_acquire, __cgo__can_release)
	return
}

/*
Attempts to release the permission represented by @permission.

The precise method by which this happens depends on the permission
and the underlying authentication mechanism.  In most cases the
permission will be dropped immediately without further action.

You should check with g_permission_get_can_release() before calling
this function.

If the permission is released then %TRUE is returned.  Otherwise,
%FALSE is returned and @error is set appropriately.

This call is blocking, likely for a very long time (in the case that
user interaction is required).  See g_permission_release_async() for
the non-blocking version.
*/
func (self *TraitPermission) Release(cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_permission_release(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Attempts to release the permission represented by @permission.

This is the first half of the asynchronous version of
g_permission_release().
*/
func (self *TraitPermission) ReleaseAsync(cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_permission_release_async(self.CPointer, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Collects the result of attempting to release the permission
represented by @permission.

This is the second half of the asynchronous version of
g_permission_release().
*/
func (self *TraitPermission) ReleaseFinish(result *C.GAsyncResult) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_permission_release_finish(self.CPointer, result, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

type TraitPropertyAction struct{ CPointer *C.GPropertyAction }
type IsPropertyAction interface {
	GetPropertyActionPointer() *C.GPropertyAction
}

func (self *TraitPropertyAction) GetPropertyActionPointer() *C.GPropertyAction {
	return self.CPointer
}
func NewTraitPropertyAction(p unsafe.Pointer) *TraitPropertyAction {
	return &TraitPropertyAction{(*C.GPropertyAction)(p)}
}

type TraitProxyAddress struct{ CPointer *C.GProxyAddress }
type IsProxyAddress interface {
	GetProxyAddressPointer() *C.GProxyAddress
}

func (self *TraitProxyAddress) GetProxyAddressPointer() *C.GProxyAddress {
	return self.CPointer
}
func NewTraitProxyAddress(p unsafe.Pointer) *TraitProxyAddress {
	return &TraitProxyAddress{(*C.GProxyAddress)(p)}
}

/*
Gets @proxy's destination hostname; that is, the name of the host
that will be connected to via the proxy, not the name of the proxy
itself.
*/
func (self *TraitProxyAddress) GetDestinationHostname() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_proxy_address_get_destination_hostname(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets @proxy's destination port; that is, the port on the
destination host that will be connected to via the proxy, not the
port number of the proxy itself.
*/
func (self *TraitProxyAddress) GetDestinationPort() (return__ uint16) {
	var __cgo__return__ C.guint16
	__cgo__return__ = C.g_proxy_address_get_destination_port(self.CPointer)
	return__ = uint16(__cgo__return__)
	return
}

/*
Gets the protocol that is being spoken to the destination
server; eg, "http" or "ftp".
*/
func (self *TraitProxyAddress) GetDestinationProtocol() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_proxy_address_get_destination_protocol(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets @proxy's password.
*/
func (self *TraitProxyAddress) GetPassword() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_proxy_address_get_password(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets @proxy's protocol. eg, "socks" or "http"
*/
func (self *TraitProxyAddress) GetProtocol() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_proxy_address_get_protocol(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the proxy URI that @proxy was constructed from.
*/
func (self *TraitProxyAddress) GetUri() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_proxy_address_get_uri(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets @proxy's username.
*/
func (self *TraitProxyAddress) GetUsername() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_proxy_address_get_username(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

type TraitProxyAddressEnumerator struct{ CPointer *C.GProxyAddressEnumerator }
type IsProxyAddressEnumerator interface {
	GetProxyAddressEnumeratorPointer() *C.GProxyAddressEnumerator
}

func (self *TraitProxyAddressEnumerator) GetProxyAddressEnumeratorPointer() *C.GProxyAddressEnumerator {
	return self.CPointer
}
func NewTraitProxyAddressEnumerator(p unsafe.Pointer) *TraitProxyAddressEnumerator {
	return &TraitProxyAddressEnumerator{(*C.GProxyAddressEnumerator)(p)}
}

type TraitResolver struct{ CPointer *C.GResolver }
type IsResolver interface {
	GetResolverPointer() *C.GResolver
}

func (self *TraitResolver) GetResolverPointer() *C.GResolver {
	return self.CPointer
}
func NewTraitResolver(p unsafe.Pointer) *TraitResolver {
	return &TraitResolver{(*C.GResolver)(p)}
}

/*
Synchronously reverse-resolves @address to determine its
associated hostname.

If the DNS resolution fails, @error (if non-%NULL) will be set to
a value from #GResolverError.

If @cancellable is non-%NULL, it can be used to cancel the
operation, in which case @error (if non-%NULL) will be set to
%G_IO_ERROR_CANCELLED.
*/
func (self *TraitResolver) LookupByAddress(address IsInetAddress, cancellable IsCancellable) (return__ string, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_resolver_lookup_by_address(self.CPointer, address.GetInetAddressPointer(), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Begins asynchronously reverse-resolving @address to determine its
associated hostname, and eventually calls @callback, which must
call g_resolver_lookup_by_address_finish() to get the final result.
*/
func (self *TraitResolver) LookupByAddressAsync(address IsInetAddress, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_resolver_lookup_by_address_async(self.CPointer, address.GetInetAddressPointer(), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Retrieves the result of a previous call to
g_resolver_lookup_by_address_async().

If the DNS resolution failed, @error (if non-%NULL) will be set to
a value from #GResolverError. If the operation was cancelled,
@error will be set to %G_IO_ERROR_CANCELLED.
*/
func (self *TraitResolver) LookupByAddressFinish(result *C.GAsyncResult) (return__ string, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_resolver_lookup_by_address_finish(self.CPointer, result, &__cgo_error__)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Synchronously resolves @hostname to determine its associated IP
address(es). @hostname may be an ASCII-only or UTF-8 hostname, or
the textual form of an IP address (in which case this just becomes
a wrapper around g_inet_address_new_from_string()).

On success, g_resolver_lookup_by_name() will return a #GList of
#GInetAddress, sorted in order of preference and guaranteed to not
contain duplicates. That is, if using the result to connect to
@hostname, you should attempt to connect to the first address
first, then the second if the first fails, etc. If you are using
the result to listen on a socket, it is appropriate to add each
result using e.g. g_socket_listener_add_address().

If the DNS resolution fails, @error (if non-%NULL) will be set to a
value from #GResolverError.

If @cancellable is non-%NULL, it can be used to cancel the
operation, in which case @error (if non-%NULL) will be set to
%G_IO_ERROR_CANCELLED.

If you are planning to connect to a socket on the resolved IP
address, it may be easier to create a #GNetworkAddress and use its
#GSocketConnectable interface.
*/
func (self *TraitResolver) LookupByName(hostname string, cancellable IsCancellable) (return__ *C.GList, __err__ error) {
	__cgo__hostname := (*C.gchar)(unsafe.Pointer(C.CString(hostname)))
	var __cgo_error__ *C.GError
	return__ = C.g_resolver_lookup_by_name(self.CPointer, __cgo__hostname, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__hostname))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Begins asynchronously resolving @hostname to determine its
associated IP address(es), and eventually calls @callback, which
must call g_resolver_lookup_by_name_finish() to get the result.
See g_resolver_lookup_by_name() for more details.
*/
func (self *TraitResolver) LookupByNameAsync(hostname string, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__hostname := (*C.gchar)(unsafe.Pointer(C.CString(hostname)))
	C.g_resolver_lookup_by_name_async(self.CPointer, __cgo__hostname, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__hostname))
	return
}

/*
Retrieves the result of a call to
g_resolver_lookup_by_name_async().

If the DNS resolution failed, @error (if non-%NULL) will be set to
a value from #GResolverError. If the operation was cancelled,
@error will be set to %G_IO_ERROR_CANCELLED.
*/
func (self *TraitResolver) LookupByNameFinish(result *C.GAsyncResult) (return__ *C.GList, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_resolver_lookup_by_name_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Synchronously performs a DNS record lookup for the given @rrname and returns
a list of records as #GVariant tuples. See #GResolverRecordType for
information on what the records contain for each @record_type.

If the DNS resolution fails, @error (if non-%NULL) will be set to
a value from #GResolverError.

If @cancellable is non-%NULL, it can be used to cancel the
operation, in which case @error (if non-%NULL) will be set to
%G_IO_ERROR_CANCELLED.
*/
func (self *TraitResolver) LookupRecords(rrname string, record_type C.GResolverRecordType, cancellable IsCancellable) (return__ *C.GList, __err__ error) {
	__cgo__rrname := (*C.gchar)(unsafe.Pointer(C.CString(rrname)))
	var __cgo_error__ *C.GError
	return__ = C.g_resolver_lookup_records(self.CPointer, __cgo__rrname, record_type, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__rrname))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Begins asynchronously performing a DNS lookup for the given
@rrname, and eventually calls @callback, which must call
g_resolver_lookup_records_finish() to get the final result. See
g_resolver_lookup_records() for more details.
*/
func (self *TraitResolver) LookupRecordsAsync(rrname string, record_type C.GResolverRecordType, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__rrname := (*C.gchar)(unsafe.Pointer(C.CString(rrname)))
	C.g_resolver_lookup_records_async(self.CPointer, __cgo__rrname, record_type, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__rrname))
	return
}

/*
Retrieves the result of a previous call to
g_resolver_lookup_records_async(). Returns a list of records as #GVariant
tuples. See #GResolverRecordType for information on what the records contain.

If the DNS resolution failed, @error (if non-%NULL) will be set to
a value from #GResolverError. If the operation was cancelled,
@error will be set to %G_IO_ERROR_CANCELLED.
*/
func (self *TraitResolver) LookupRecordsFinish(result *C.GAsyncResult) (return__ *C.GList, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_resolver_lookup_records_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Synchronously performs a DNS SRV lookup for the given @service and
@protocol in the given @domain and returns an array of #GSrvTarget.
@domain may be an ASCII-only or UTF-8 hostname. Note also that the
@service and @protocol arguments do not include the leading underscore
that appears in the actual DNS entry.

On success, g_resolver_lookup_service() will return a #GList of
#GSrvTarget, sorted in order of preference. (That is, you should
attempt to connect to the first target first, then the second if
the first fails, etc.)

If the DNS resolution fails, @error (if non-%NULL) will be set to
a value from #GResolverError.

If @cancellable is non-%NULL, it can be used to cancel the
operation, in which case @error (if non-%NULL) will be set to
%G_IO_ERROR_CANCELLED.

If you are planning to connect to the service, it is usually easier
to create a #GNetworkService and use its #GSocketConnectable
interface.
*/
func (self *TraitResolver) LookupService(service string, protocol string, domain string, cancellable IsCancellable) (return__ *C.GList, __err__ error) {
	__cgo__service := (*C.gchar)(unsafe.Pointer(C.CString(service)))
	__cgo__protocol := (*C.gchar)(unsafe.Pointer(C.CString(protocol)))
	__cgo__domain := (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	var __cgo_error__ *C.GError
	return__ = C.g_resolver_lookup_service(self.CPointer, __cgo__service, __cgo__protocol, __cgo__domain, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__service))
	C.free(unsafe.Pointer(__cgo__protocol))
	C.free(unsafe.Pointer(__cgo__domain))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Begins asynchronously performing a DNS SRV lookup for the given
@service and @protocol in the given @domain, and eventually calls
@callback, which must call g_resolver_lookup_service_finish() to
get the final result. See g_resolver_lookup_service() for more
details.
*/
func (self *TraitResolver) LookupServiceAsync(service string, protocol string, domain string, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__service := (*C.gchar)(unsafe.Pointer(C.CString(service)))
	__cgo__protocol := (*C.gchar)(unsafe.Pointer(C.CString(protocol)))
	__cgo__domain := (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	C.g_resolver_lookup_service_async(self.CPointer, __cgo__service, __cgo__protocol, __cgo__domain, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__service))
	C.free(unsafe.Pointer(__cgo__protocol))
	C.free(unsafe.Pointer(__cgo__domain))
	return
}

/*
Retrieves the result of a previous call to
g_resolver_lookup_service_async().

If the DNS resolution failed, @error (if non-%NULL) will be set to
a value from #GResolverError. If the operation was cancelled,
@error will be set to %G_IO_ERROR_CANCELLED.
*/
func (self *TraitResolver) LookupServiceFinish(result *C.GAsyncResult) (return__ *C.GList, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_resolver_lookup_service_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Sets @resolver to be the application's default resolver (reffing
@resolver, and unreffing the previous default resolver, if any).
Future calls to g_resolver_get_default() will return this resolver.

This can be used if an application wants to perform any sort of DNS
caching or "pinning"; it can implement its own #GResolver that
calls the original default resolver for DNS operations, and
implements its own cache policies on top of that, and then set
itself as the default resolver for all later code to use.
*/
func (self *TraitResolver) SetDefault() {
	C.g_resolver_set_default(self.CPointer)
	return
}

type TraitSettings struct{ CPointer *C.GSettings }
type IsSettings interface {
	GetSettingsPointer() *C.GSettings
}

func (self *TraitSettings) GetSettingsPointer() *C.GSettings {
	return self.CPointer
}
func NewTraitSettings(p unsafe.Pointer) *TraitSettings {
	return &TraitSettings{(*C.GSettings)(p)}
}

/*
Applies any changes that have been made to the settings.  This
function does nothing unless @settings is in 'delay-apply' mode;
see g_settings_delay().  In the normal case settings are always
applied immediately.
*/
func (self *TraitSettings) Apply() {
	C.g_settings_apply(self.CPointer)
	return
}

/*
Create a binding between the @key in the @settings object
and the property @property of @object.

The binding uses the default GIO mapping functions to map
between the settings and property values. These functions
handle booleans, numeric types and string types in a
straightforward way. Use g_settings_bind_with_mapping() if
you need a custom mapping, or map between types that are not
supported by the default mapping functions.

Unless the @flags include %G_SETTINGS_BIND_NO_SENSITIVITY, this
function also establishes a binding between the writability of
@key and the "sensitive" property of @object (if @object has
a boolean property by that name). See g_settings_bind_writable()
for more details about writable bindings.

Note that the lifecycle of the binding is tied to the object,
and that you can have only one binding per object property.
If you bind the same property twice on the same object, the second
binding overrides the first one.
*/
func (self *TraitSettings) Bind(key string, object gobject.IsObject, property string, flags C.GSettingsBindFlags) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	__cgo__property := (*C.gchar)(unsafe.Pointer(C.CString(property)))
	C.g_settings_bind(self.CPointer, __cgo__key, C.gpointer(unsafe.Pointer(reflect.ValueOf(object.GetObjectPointer()).Pointer())), __cgo__property, flags)
	C.free(unsafe.Pointer(__cgo__key))
	C.free(unsafe.Pointer(__cgo__property))
	return
}

/*
Create a binding between the @key in the @settings object
and the property @property of @object.

The binding uses the provided mapping functions to map between
settings and property values.

Note that the lifecycle of the binding is tied to the object,
and that you can have only one binding per object property.
If you bind the same property twice on the same object, the second
binding overrides the first one.
*/
func (self *TraitSettings) BindWithMapping(key string, object gobject.IsObject, property string, flags C.GSettingsBindFlags, get_mapping C.GSettingsBindGetMapping, set_mapping C.GSettingsBindSetMapping, user_data unsafe.Pointer, destroy C.GDestroyNotify) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	__cgo__property := (*C.gchar)(unsafe.Pointer(C.CString(property)))
	C.g_settings_bind_with_mapping(self.CPointer, __cgo__key, C.gpointer(unsafe.Pointer(reflect.ValueOf(object.GetObjectPointer()).Pointer())), __cgo__property, flags, get_mapping, set_mapping, (C.gpointer)(user_data), destroy)
	C.free(unsafe.Pointer(__cgo__key))
	C.free(unsafe.Pointer(__cgo__property))
	return
}

/*
Create a binding between the writability of @key in the
@settings object and the property @property of @object.
The property must be boolean; "sensitive" or "visible"
properties of widgets are the most likely candidates.

Writable bindings are always uni-directional; changes of the
writability of the setting will be propagated to the object
property, not the other way.

When the @inverted argument is %TRUE, the binding inverts the
value as it passes from the setting to the object, i.e. @property
will be set to %TRUE if the key is not writable.

Note that the lifecycle of the binding is tied to the object,
and that you can have only one binding per object property.
If you bind the same property twice on the same object, the second
binding overrides the first one.
*/
func (self *TraitSettings) BindWritable(key string, object gobject.IsObject, property string, inverted bool) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	__cgo__property := (*C.gchar)(unsafe.Pointer(C.CString(property)))
	__cgo__inverted := C.gboolean(0)
	if inverted {
		__cgo__inverted = C.gboolean(1)
	}
	C.g_settings_bind_writable(self.CPointer, __cgo__key, C.gpointer(unsafe.Pointer(reflect.ValueOf(object.GetObjectPointer()).Pointer())), __cgo__property, __cgo__inverted)
	C.free(unsafe.Pointer(__cgo__key))
	C.free(unsafe.Pointer(__cgo__property))
	return
}

/*
Creates a #GAction corresponding to a given #GSettings key.

The action has the same name as the key.

The value of the key becomes the state of the action and the action
is enabled when the key is writable.  Changing the state of the
action results in the key being written to.  Changes to the value or
writability of the key cause appropriate change notifications to be
emitted for the action.

For boolean-valued keys, action activations take no parameter and
result in the toggling of the value.  For all other types,
activations take the new value for the key (which must have the
correct type).
*/
func (self *TraitSettings) CreateAction(key string) (return__ *C.GAction) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	return__ = C.g_settings_create_action(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return
}

/*
Changes the #GSettings object into 'delay-apply' mode. In this
mode, changes to @settings are not immediately propagated to the
backend, but kept locally until g_settings_apply() is called.
*/
func (self *TraitSettings) Delay() {
	C.g_settings_delay(self.CPointer)
	return
}

// g_settings_get is not generated due to varargs

/*
Gets the value that is stored at @key in @settings.

A convenience variant of g_settings_get() for booleans.

It is a programmer error to give a @key that isn't specified as
having a boolean type in the schema for @settings.
*/
func (self *TraitSettings) GetBoolean(key string) (return__ bool) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_settings_get_boolean(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Creates a child settings object which has a base path of
`base-path/@name`, where `base-path` is the base path of
@settings.

The schema for the child settings object must have been declared
in the schema of @settings using a <child> element.
*/
func (self *TraitSettings) GetChild(name string) (return__ *Settings) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ *C.GSettings
	__cgo__return__ = C.g_settings_get_child(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	if __cgo__return__ != nil {
		return__ = NewSettingsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the "default value" of a key.

This is the value that would be read if g_settings_reset() were to be
called on the key.

Note that this may be a different value than returned by
g_settings_schema_key_get_default_value() if the system administrator
has provided a default value.

Comparing the return values of g_settings_get_default_value() and
g_settings_get_value() is not sufficient for determining if a value
has been set because the user may have explicitly set the value to
something that happens to be equal to the default.  The difference
here is that if the default changes in the future, the user's key
will still be set.

This function may be useful for adding an indication to a UI of what
the default value was before the user set it.

It is a programmer error to give a @key that isn't contained in the
schema for @settings.
*/
func (self *TraitSettings) GetDefaultValue(key string) (return__ *C.GVariant) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	return__ = C.g_settings_get_default_value(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return
}

/*
Gets the value that is stored at @key in @settings.

A convenience variant of g_settings_get() for doubles.

It is a programmer error to give a @key that isn't specified as
having a 'double' type in the schema for @settings.
*/
func (self *TraitSettings) GetDouble(key string) (return__ float64) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.g_settings_get_double(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the value that is stored in @settings for @key and converts it
to the enum value that it represents.

In order to use this function the type of the value must be a string
and it must be marked in the schema file as an enumerated type.

It is a programmer error to give a @key that isn't contained in the
schema for @settings or is not marked as an enumerated type.

If the value stored in the configuration database is not a valid
value for the enumerated type then this function will return the
default value.
*/
func (self *TraitSettings) GetEnum(key string) (return__ int) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_settings_get_enum(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = int(__cgo__return__)
	return
}

/*
Gets the value that is stored in @settings for @key and converts it
to the flags value that it represents.

In order to use this function the type of the value must be an array
of strings and it must be marked in the schema file as an flags type.

It is a programmer error to give a @key that isn't contained in the
schema for @settings or is not marked as a flags type.

If the value stored in the configuration database is not a valid
value for the flags type then this function will return the default
value.
*/
func (self *TraitSettings) GetFlags(key string) (return__ uint) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_settings_get_flags(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = uint(__cgo__return__)
	return
}

/*
Returns whether the #GSettings object has any unapplied
changes.  This can only be the case if it is in 'delayed-apply' mode.
*/
func (self *TraitSettings) GetHasUnapplied() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_settings_get_has_unapplied(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value that is stored at @key in @settings.

A convenience variant of g_settings_get() for 32-bit integers.

It is a programmer error to give a @key that isn't specified as
having a int32 type in the schema for @settings.
*/
func (self *TraitSettings) GetInt(key string) (return__ int) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_settings_get_int(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = int(__cgo__return__)
	return
}

/*
Gets the value that is stored at @key in @settings, subject to
application-level validation/mapping.

You should use this function when the application needs to perform
some processing on the value of the key (for example, parsing).  The
@mapping function performs that processing.  If the function
indicates that the processing was unsuccessful (due to a parse error,
for example) then the mapping is tried again with another value.

This allows a robust 'fall back to defaults' behaviour to be
implemented somewhat automatically.

The first value that is tried is the user's setting for the key.  If
the mapping function fails to map this value, other values may be
tried in an unspecified order (system or site defaults, translated
schema default values, untranslated schema default values, etc).

If the mapping function fails for all possible values, one additional
attempt is made: the mapping function is called with a %NULL value.
If the mapping function still indicates failure at this point then
the application will be aborted.

The result parameter for the @mapping function is pointed to a
#gpointer which is initially set to %NULL.  The same pointer is given
to each invocation of @mapping.  The final value of that #gpointer is
what is returned by this function.  %NULL is valid; it is returned
just as any other value would be.
*/
func (self *TraitSettings) GetMapped(key string, mapping C.GSettingsGetMapping, user_data unsafe.Pointer) (return__ unsafe.Pointer) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_settings_get_mapped(self.CPointer, __cgo__key, mapping, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__key))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

// g_settings_get_range is not generated due to deprecation attr

/*
Gets the value that is stored at @key in @settings.

A convenience variant of g_settings_get() for strings.

It is a programmer error to give a @key that isn't specified as
having a string type in the schema for @settings.
*/
func (self *TraitSettings) GetString(key string) (return__ string) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_settings_get_string(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
A convenience variant of g_settings_get() for string arrays.

It is a programmer error to give a @key that isn't specified as
having an array of strings type in the schema for @settings.
*/
func (self *TraitSettings) GetStrv(key string) (return__ **C.gchar) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	return__ = C.g_settings_get_strv(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return
}

/*
Gets the value that is stored at @key in @settings.

A convenience variant of g_settings_get() for 32-bit unsigned
integers.

It is a programmer error to give a @key that isn't specified as
having a uint32 type in the schema for @settings.
*/
func (self *TraitSettings) GetUint(key string) (return__ uint) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_settings_get_uint(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = uint(__cgo__return__)
	return
}

/*
Checks the "user value" of a key, if there is one.

The user value of a key is the last value that was set by the user.

After calling g_settings_reset() this function should always return
%NULL (assuming something is not wrong with the system
configuration).

It is possible that g_settings_get_value() will return a different
value than this function.  This can happen in the case that the user
set a value for a key that was subsequently locked down by the system
administrator -- this function will return the user's old value.

This function may be useful for adding a "reset" option to a UI or
for providing indication that a particular value has been changed.

It is a programmer error to give a @key that isn't contained in the
schema for @settings.
*/
func (self *TraitSettings) GetUserValue(key string) (return__ *C.GVariant) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	return__ = C.g_settings_get_user_value(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return
}

/*
Gets the value that is stored in @settings for @key.

It is a programmer error to give a @key that isn't contained in the
schema for @settings.
*/
func (self *TraitSettings) GetValue(key string) (return__ *C.GVariant) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	return__ = C.g_settings_get_value(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return
}

/*
Finds out if a key can be written or not
*/
func (self *TraitSettings) IsWritable(name string) (return__ bool) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_settings_is_writable(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the list of children on @settings.

The list is exactly the list of strings for which it is not an error
to call g_settings_get_child().

For GSettings objects that are lists, this value can change at any
time and you should connect to the "children-changed" signal to watch
for those changes.  Note that there is a race condition here: you may
request a child after listing it only for it to have been destroyed
in the meantime.  For this reason, g_settings_get_child() may return
%NULL even for a child that was listed by this function.

For GSettings objects that are not lists, you should probably not be
calling this function from "normal" code (since you should already
know what children are in your schema).  This function may still be
useful there for introspection reasons, however.

You should free the return value with g_strfreev() when you are done
with it.
*/
func (self *TraitSettings) ListChildren() (return__ **C.gchar) {
	return__ = C.g_settings_list_children(self.CPointer)
	return
}

/*
Introspects the list of keys on @settings.

You should probably not be calling this function from "normal" code
(since you should already know what keys are in your schema).  This
function is intended for introspection reasons.

You should free the return value with g_strfreev() when you are done
with it.
*/
func (self *TraitSettings) ListKeys() (return__ **C.gchar) {
	return__ = C.g_settings_list_keys(self.CPointer)
	return
}

// g_settings_range_check is not generated due to deprecation attr

/*
Resets @key to its default value.

This call resets the key, as much as possible, to its default value.
That might the value specified in the schema or the one set by the
administrator.
*/
func (self *TraitSettings) Reset(key string) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	C.g_settings_reset(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return
}

/*
Reverts all non-applied changes to the settings.  This function
does nothing unless @settings is in 'delay-apply' mode; see
g_settings_delay().  In the normal case settings are always applied
immediately.

Change notifications will be emitted for affected keys.
*/
func (self *TraitSettings) Revert() {
	C.g_settings_revert(self.CPointer)
	return
}

// g_settings_set is not generated due to varargs

/*
Sets @key in @settings to @value.

A convenience variant of g_settings_set() for booleans.

It is a programmer error to give a @key that isn't specified as
having a boolean type in the schema for @settings.
*/
func (self *TraitSettings) SetBoolean(key string, value bool) (return__ bool) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	__cgo__value := C.gboolean(0)
	if value {
		__cgo__value = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_settings_set_boolean(self.CPointer, __cgo__key, __cgo__value)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets @key in @settings to @value.

A convenience variant of g_settings_set() for doubles.

It is a programmer error to give a @key that isn't specified as
having a 'double' type in the schema for @settings.
*/
func (self *TraitSettings) SetDouble(key string, value float64) (return__ bool) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_settings_set_double(self.CPointer, __cgo__key, C.gdouble(value))
	C.free(unsafe.Pointer(__cgo__key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Looks up the enumerated type nick for @value and writes it to @key,
within @settings.

It is a programmer error to give a @key that isn't contained in the
schema for @settings or is not marked as an enumerated type, or for
@value not to be a valid value for the named type.

After performing the write, accessing @key directly with
g_settings_get_string() will return the 'nick' associated with
@value.
*/
func (self *TraitSettings) SetEnum(key string, value int) (return__ bool) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_settings_set_enum(self.CPointer, __cgo__key, C.gint(value))
	C.free(unsafe.Pointer(__cgo__key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Looks up the flags type nicks for the bits specified by @value, puts
them in an array of strings and writes the array to @key, within
@settings.

It is a programmer error to give a @key that isn't contained in the
schema for @settings or is not marked as a flags type, or for @value
to contain any bits that are not value for the named type.

After performing the write, accessing @key directly with
g_settings_get_strv() will return an array of 'nicks'; one for each
bit in @value.
*/
func (self *TraitSettings) SetFlags(key string, value uint) (return__ bool) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_settings_set_flags(self.CPointer, __cgo__key, C.guint(value))
	C.free(unsafe.Pointer(__cgo__key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets @key in @settings to @value.

A convenience variant of g_settings_set() for 32-bit integers.

It is a programmer error to give a @key that isn't specified as
having a int32 type in the schema for @settings.
*/
func (self *TraitSettings) SetInt(key string, value int) (return__ bool) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_settings_set_int(self.CPointer, __cgo__key, C.gint(value))
	C.free(unsafe.Pointer(__cgo__key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets @key in @settings to @value.

A convenience variant of g_settings_set() for strings.

It is a programmer error to give a @key that isn't specified as
having a string type in the schema for @settings.
*/
func (self *TraitSettings) SetString(key string, value string) (return__ bool) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	__cgo__value := (*C.gchar)(unsafe.Pointer(C.CString(value)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_settings_set_string(self.CPointer, __cgo__key, __cgo__value)
	C.free(unsafe.Pointer(__cgo__key))
	C.free(unsafe.Pointer(__cgo__value))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets @key in @settings to @value.

A convenience variant of g_settings_set() for string arrays.  If
@value is %NULL, then @key is set to be the empty array.

It is a programmer error to give a @key that isn't specified as
having an array of strings type in the schema for @settings.
*/
func (self *TraitSettings) SetStrv(key string, value []string) (return__ bool) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	__header__value := (*reflect.SliceHeader)(unsafe.Pointer(&value))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_settings_set_strv(self.CPointer, __cgo__key, (**C.gchar)(unsafe.Pointer(__header__value.Data)))
	C.free(unsafe.Pointer(__cgo__key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets @key in @settings to @value.

A convenience variant of g_settings_set() for 32-bit unsigned
integers.

It is a programmer error to give a @key that isn't specified as
having a uint32 type in the schema for @settings.
*/
func (self *TraitSettings) SetUint(key string, value uint) (return__ bool) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_settings_set_uint(self.CPointer, __cgo__key, C.guint(value))
	C.free(unsafe.Pointer(__cgo__key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets @key in @settings to @value.

It is a programmer error to give a @key that isn't contained in the
schema for @settings or for @value to have the incorrect type, per
the schema.

If @value is floating then this function consumes the reference.
*/
func (self *TraitSettings) SetValue(key string, value *C.GVariant) (return__ bool) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_settings_set_value(self.CPointer, __cgo__key, value)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

type TraitSimpleAction struct{ CPointer *C.GSimpleAction }
type IsSimpleAction interface {
	GetSimpleActionPointer() *C.GSimpleAction
}

func (self *TraitSimpleAction) GetSimpleActionPointer() *C.GSimpleAction {
	return self.CPointer
}
func NewTraitSimpleAction(p unsafe.Pointer) *TraitSimpleAction {
	return &TraitSimpleAction{(*C.GSimpleAction)(p)}
}

/*
Sets the action as enabled or not.

An action must be enabled in order to be activated or in order to
have its state changed from outside callers.

This should only be called by the implementor of the action.  Users
of the action should not attempt to modify its enabled flag.
*/
func (self *TraitSimpleAction) SetEnabled(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.g_simple_action_set_enabled(self.CPointer, __cgo__enabled)
	return
}

/*
Sets the state of the action.

This directly updates the 'state' property to the given value.

This should only be called by the implementor of the action.  Users
of the action should not attempt to directly modify the 'state'
property.  Instead, they should call g_action_change_state() to
request the change.

If the @value GVariant is floating, it is consumed.
*/
func (self *TraitSimpleAction) SetState(value *C.GVariant) {
	C.g_simple_action_set_state(self.CPointer, value)
	return
}

type TraitSimpleActionGroup struct{ CPointer *C.GSimpleActionGroup }
type IsSimpleActionGroup interface {
	GetSimpleActionGroupPointer() *C.GSimpleActionGroup
}

func (self *TraitSimpleActionGroup) GetSimpleActionGroupPointer() *C.GSimpleActionGroup {
	return self.CPointer
}
func NewTraitSimpleActionGroup(p unsafe.Pointer) *TraitSimpleActionGroup {
	return &TraitSimpleActionGroup{(*C.GSimpleActionGroup)(p)}
}

// g_simple_action_group_add_entries is not generated due to deprecation attr

// g_simple_action_group_insert is not generated due to deprecation attr

// g_simple_action_group_lookup is not generated due to deprecation attr

// g_simple_action_group_remove is not generated due to deprecation attr

type TraitSimpleAsyncResult struct{ CPointer *C.GSimpleAsyncResult }
type IsSimpleAsyncResult interface {
	GetSimpleAsyncResultPointer() *C.GSimpleAsyncResult
}

func (self *TraitSimpleAsyncResult) GetSimpleAsyncResultPointer() *C.GSimpleAsyncResult {
	return self.CPointer
}
func NewTraitSimpleAsyncResult(p unsafe.Pointer) *TraitSimpleAsyncResult {
	return &TraitSimpleAsyncResult{(*C.GSimpleAsyncResult)(p)}
}

/*
Completes an asynchronous I/O job immediately. Must be called in
the thread where the asynchronous result was to be delivered, as it
invokes the callback directly. If you are in a different thread use
g_simple_async_result_complete_in_idle().

Calling this function takes a reference to @simple for as long as
is needed to complete the call.
*/
func (self *TraitSimpleAsyncResult) Complete() {
	C.g_simple_async_result_complete(self.CPointer)
	return
}

/*
Completes an asynchronous function in an idle handler in the
[thread-default main context][g-main-context-push-thread-default]
of the thread that @simple was initially created in
(and re-pushes that context around the invocation of the callback).

Calling this function takes a reference to @simple for as long as
is needed to complete the call.
*/
func (self *TraitSimpleAsyncResult) CompleteInIdle() {
	C.g_simple_async_result_complete_in_idle(self.CPointer)
	return
}

/*
Gets the operation result boolean from within the asynchronous result.
*/
func (self *TraitSimpleAsyncResult) GetOpResGboolean() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_simple_async_result_get_op_res_gboolean(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets a pointer result as returned by the asynchronous function.
*/
func (self *TraitSimpleAsyncResult) GetOpResGpointer() (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_simple_async_result_get_op_res_gpointer(self.CPointer)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Gets a gssize from the asynchronous result.
*/
func (self *TraitSimpleAsyncResult) GetOpResGssize() (return__ int64) {
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_simple_async_result_get_op_res_gssize(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

/*
Gets the source tag for the #GSimpleAsyncResult.
*/
func (self *TraitSimpleAsyncResult) GetSourceTag() (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_simple_async_result_get_source_tag(self.CPointer)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Propagates an error from within the simple asynchronous result to
a given destination.

If the #GCancellable given to a prior call to
g_simple_async_result_set_check_cancellable() is cancelled then this
function will return %TRUE with @dest set appropriately.
*/
func (self *TraitSimpleAsyncResult) PropagateError() (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_simple_async_result_propagate_error(self.CPointer, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Runs the asynchronous job in a separate thread and then calls
g_simple_async_result_complete_in_idle() on @simple to return
the result to the appropriate main loop.

Calling this function takes a reference to @simple for as long as
is needed to run the job and report its completion.
*/
func (self *TraitSimpleAsyncResult) RunInThread(func_ C.GSimpleAsyncThreadFunc, io_priority int, cancellable IsCancellable) {
	C.g_simple_async_result_run_in_thread(self.CPointer, func_, C.int(io_priority), cancellable.GetCancellablePointer())
	return
}

/*
Sets a #GCancellable to check before dispatching results.

This function has one very specific purpose: the provided cancellable
is checked at the time of g_simple_async_result_propagate_error() If
it is cancelled, these functions will return an "Operation was
cancelled" error (%G_IO_ERROR_CANCELLED).

Implementors of cancellable asynchronous functions should use this in
order to provide a guarantee to their callers that cancelling an
async operation will reliably result in an error being returned for
that operation (even if a positive result for the operation has
already been sent as an idle to the main context to be dispatched).

The checking described above is done regardless of any call to the
unrelated g_simple_async_result_set_handle_cancellation() function.
*/
func (self *TraitSimpleAsyncResult) SetCheckCancellable(check_cancellable IsCancellable) {
	C.g_simple_async_result_set_check_cancellable(self.CPointer, check_cancellable.GetCancellablePointer())
	return
}

// g_simple_async_result_set_error is not generated due to varargs

// g_simple_async_result_set_error_va is not generated due to varargs

/*
Sets the result from a #GError.
*/
func (self *TraitSimpleAsyncResult) SetFromError(error_ *C.GError) {
	C.g_simple_async_result_set_from_error(self.CPointer, error_)
	return
}

/*
Sets whether to handle cancellation within the asynchronous operation.

This function has nothing to do with
g_simple_async_result_set_check_cancellable().  It only refers to the
#GCancellable passed to g_simple_async_result_run_in_thread().
*/
func (self *TraitSimpleAsyncResult) SetHandleCancellation(handle_cancellation bool) {
	__cgo__handle_cancellation := C.gboolean(0)
	if handle_cancellation {
		__cgo__handle_cancellation = C.gboolean(1)
	}
	C.g_simple_async_result_set_handle_cancellation(self.CPointer, __cgo__handle_cancellation)
	return
}

/*
Sets the operation result to a boolean within the asynchronous result.
*/
func (self *TraitSimpleAsyncResult) SetOpResGboolean(op_res bool) {
	__cgo__op_res := C.gboolean(0)
	if op_res {
		__cgo__op_res = C.gboolean(1)
	}
	C.g_simple_async_result_set_op_res_gboolean(self.CPointer, __cgo__op_res)
	return
}

/*
Sets the operation result within the asynchronous result to a pointer.
*/
func (self *TraitSimpleAsyncResult) SetOpResGpointer(op_res unsafe.Pointer, destroy_op_res C.GDestroyNotify) {
	C.g_simple_async_result_set_op_res_gpointer(self.CPointer, (C.gpointer)(op_res), destroy_op_res)
	return
}

/*
Sets the operation result within the asynchronous result to
the given @op_res.
*/
func (self *TraitSimpleAsyncResult) SetOpResGssize(op_res int64) {
	C.g_simple_async_result_set_op_res_gssize(self.CPointer, C.gssize(op_res))
	return
}

/*
Sets the result from @error, and takes over the caller's ownership
of @error, so the caller does not need to free it any more.
*/
func (self *TraitSimpleAsyncResult) TakeError(error_ *C.GError) {
	C.g_simple_async_result_take_error(self.CPointer, error_)
	return
}

type TraitSimplePermission struct{ CPointer *C.GSimplePermission }
type IsSimplePermission interface {
	GetSimplePermissionPointer() *C.GSimplePermission
}

func (self *TraitSimplePermission) GetSimplePermissionPointer() *C.GSimplePermission {
	return self.CPointer
}
func NewTraitSimplePermission(p unsafe.Pointer) *TraitSimplePermission {
	return &TraitSimplePermission{(*C.GSimplePermission)(p)}
}

type TraitSimpleProxyResolver struct{ CPointer *C.GSimpleProxyResolver }
type IsSimpleProxyResolver interface {
	GetSimpleProxyResolverPointer() *C.GSimpleProxyResolver
}

func (self *TraitSimpleProxyResolver) GetSimpleProxyResolverPointer() *C.GSimpleProxyResolver {
	return self.CPointer
}
func NewTraitSimpleProxyResolver(p unsafe.Pointer) *TraitSimpleProxyResolver {
	return &TraitSimpleProxyResolver{(*C.GSimpleProxyResolver)(p)}
}

/*
Sets the default proxy on @resolver, to be used for any URIs that
don't match #GSimpleProxyResolver:ignore-hosts or a proxy set
via g_simple_proxy_resolver_set_uri_proxy().

If @default_proxy starts with "socks://",
#GSimpleProxyResolver will treat it as referring to all three of
the socks5, socks4a, and socks4 proxy types.
*/
func (self *TraitSimpleProxyResolver) SetDefaultProxy(default_proxy string) {
	__cgo__default_proxy := (*C.gchar)(unsafe.Pointer(C.CString(default_proxy)))
	C.g_simple_proxy_resolver_set_default_proxy(self.CPointer, __cgo__default_proxy)
	C.free(unsafe.Pointer(__cgo__default_proxy))
	return
}

/*
Sets the list of ignored hosts.

See #GSimpleProxyResolver:ignore-hosts for more details on how the
@ignore_hosts argument is interpreted.
*/
func (self *TraitSimpleProxyResolver) SetIgnoreHosts(ignore_hosts []string) {
	__header__ignore_hosts := (*reflect.SliceHeader)(unsafe.Pointer(&ignore_hosts))
	C.g_simple_proxy_resolver_set_ignore_hosts(self.CPointer, (**C.gchar)(unsafe.Pointer(__header__ignore_hosts.Data)))
	return
}

/*
Adds a URI-scheme-specific proxy to @resolver; URIs whose scheme
matches @uri_scheme (and which don't match
#GSimpleProxyResolver:ignore-hosts) will be proxied via @proxy.

As with #GSimpleProxyResolver:default-proxy, if @proxy starts with
"socks://", #GSimpleProxyResolver will treat it
as referring to all three of the socks5, socks4a, and socks4 proxy
types.
*/
func (self *TraitSimpleProxyResolver) SetUriProxy(uri_scheme string, proxy string) {
	__cgo__uri_scheme := (*C.gchar)(unsafe.Pointer(C.CString(uri_scheme)))
	__cgo__proxy := (*C.gchar)(unsafe.Pointer(C.CString(proxy)))
	C.g_simple_proxy_resolver_set_uri_proxy(self.CPointer, __cgo__uri_scheme, __cgo__proxy)
	C.free(unsafe.Pointer(__cgo__uri_scheme))
	C.free(unsafe.Pointer(__cgo__proxy))
	return
}

type TraitSocket struct{ CPointer *C.GSocket }
type IsSocket interface {
	GetSocketPointer() *C.GSocket
}

func (self *TraitSocket) GetSocketPointer() *C.GSocket {
	return self.CPointer
}
func NewTraitSocket(p unsafe.Pointer) *TraitSocket {
	return &TraitSocket{(*C.GSocket)(p)}
}

/*
Accept incoming connections on a connection-based socket. This removes
the first outstanding connection request from the listening socket and
creates a #GSocket object for it.

The @socket must be bound to a local address with g_socket_bind() and
must be listening for incoming connections (g_socket_listen()).

If there are no outstanding connections then the operation will block
or return %G_IO_ERROR_WOULD_BLOCK if non-blocking I/O is enabled.
To be notified of an incoming connection, wait for the %G_IO_IN condition.
*/
func (self *TraitSocket) Accept(cancellable IsCancellable) (return__ *Socket, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocket
	__cgo__return__ = C.g_socket_accept(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
When a socket is created it is attached to an address family, but it
doesn't have an address in this family. g_socket_bind() assigns the
address (sometimes called name) of the socket.

It is generally required to bind to a local address before you can
receive connections. (See g_socket_listen() and g_socket_accept() ).
In certain situations, you may also want to bind a socket that will be
used to initiate connections, though this is not normally required.

If @socket is a TCP socket, then @allow_reuse controls the setting
of the `SO_REUSEADDR` socket option; normally it should be %TRUE for
server sockets (sockets that you will eventually call
g_socket_accept() on), and %FALSE for client sockets. (Failing to
set this flag on a server socket may cause g_socket_bind() to return
%G_IO_ERROR_ADDRESS_IN_USE if the server program is stopped and then
immediately restarted.)

If @socket is a UDP socket, then @allow_reuse determines whether or
not other UDP sockets can be bound to the same address at the same
time. In particular, you can have several UDP sockets bound to the
same address, and they will all receive all of the multicast and
broadcast packets sent to that address. (The behavior of unicast
UDP packets to an address with multiple listeners is not defined.)
*/
func (self *TraitSocket) Bind(address IsSocketAddress, allow_reuse bool) (return__ bool, __err__ error) {
	__cgo__allow_reuse := C.gboolean(0)
	if allow_reuse {
		__cgo__allow_reuse = C.gboolean(1)
	}
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_bind(self.CPointer, address.GetSocketAddressPointer(), __cgo__allow_reuse, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Checks and resets the pending connect error for the socket.
This is used to check for errors when g_socket_connect() is
used in non-blocking mode.
*/
func (self *TraitSocket) CheckConnectResult() (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_check_connect_result(self.CPointer, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Closes the socket, shutting down any active connection.

Closing a socket does not wait for all outstanding I/O operations
to finish, so the caller should not rely on them to be guaranteed
to complete even if the close returns with no error.

Once the socket is closed, all other operations will return
%G_IO_ERROR_CLOSED. Closing a socket multiple times will not
return an error.

Sockets will be automatically closed when the last reference
is dropped, but you might want to call this function to make sure
resources are released as early as possible.

Beware that due to the way that TCP works, it is possible for
recently-sent data to be lost if either you close a socket while the
%G_IO_IN condition is set, or else if the remote connection tries to
send something to you after you close the socket but before it has
finished reading all of the data you sent. There is no easy generic
way to avoid this problem; the easiest fix is to design the network
protocol such that the client will never send data "out of turn".
Another solution is for the server to half-close the connection by
calling g_socket_shutdown() with only the @shutdown_write flag set,
and then wait for the client to notice this and close its side of the
connection, after which the server can safely call g_socket_close().
(This is what #GTcpConnection does if you call
g_tcp_connection_set_graceful_disconnect(). But of course, this
only works if the client will close its connection after the server
does.)
*/
func (self *TraitSocket) Close() (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_close(self.CPointer, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Checks on the readiness of @socket to perform operations.
The operations specified in @condition are checked for and masked
against the currently-satisfied conditions on @socket. The result
is returned.

Note that on Windows, it is possible for an operation to return
%G_IO_ERROR_WOULD_BLOCK even immediately after
g_socket_condition_check() has claimed that the socket is ready for
writing. Rather than calling g_socket_condition_check() and then
writing to the socket if it succeeds, it is generally better to
simply try writing to the socket right away, and try again later if
the initial attempt returns %G_IO_ERROR_WOULD_BLOCK.

It is meaningless to specify %G_IO_ERR or %G_IO_HUP in condition;
these conditions will always be set in the output if they are true.

This call never blocks.
*/
func (self *TraitSocket) ConditionCheck(condition C.GIOCondition) (return__ C.GIOCondition) {
	return__ = C.g_socket_condition_check(self.CPointer, condition)
	return
}

/*
Waits for up to @timeout microseconds for @condition to become true
on @socket. If the condition is met, %TRUE is returned.

If @cancellable is cancelled before the condition is met, or if
@timeout (or the socket's #GSocket:timeout) is reached before the
condition is met, then %FALSE is returned and @error, if non-%NULL,
is set to the appropriate value (%G_IO_ERROR_CANCELLED or
%G_IO_ERROR_TIMED_OUT).

If you don't want a timeout, use g_socket_condition_wait().
(Alternatively, you can pass -1 for @timeout.)

Note that although @timeout is in microseconds for consistency with
other GLib APIs, this function actually only has millisecond
resolution, and the behavior is undefined if @timeout is not an
exact number of milliseconds.
*/
func (self *TraitSocket) ConditionTimedWait(condition C.GIOCondition, timeout int64, cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_condition_timed_wait(self.CPointer, condition, C.gint64(timeout), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Waits for @condition to become true on @socket. When the condition
is met, %TRUE is returned.

If @cancellable is cancelled before the condition is met, or if the
socket has a timeout set and it is reached before the condition is
met, then %FALSE is returned and @error, if non-%NULL, is set to
the appropriate value (%G_IO_ERROR_CANCELLED or
%G_IO_ERROR_TIMED_OUT).

See also g_socket_condition_timed_wait().
*/
func (self *TraitSocket) ConditionWait(condition C.GIOCondition, cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_condition_wait(self.CPointer, condition, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Connect the socket to the specified remote address.

For connection oriented socket this generally means we attempt to make
a connection to the @address. For a connection-less socket it sets
the default address for g_socket_send() and discards all incoming datagrams
from other sources.

Generally connection oriented sockets can only connect once, but
connection-less sockets can connect multiple times to change the
default address.

If the connect call needs to do network I/O it will block, unless
non-blocking I/O is enabled. Then %G_IO_ERROR_PENDING is returned
and the user can be notified of the connection finishing by waiting
for the G_IO_OUT condition. The result of the connection must then be
checked with g_socket_check_connect_result().
*/
func (self *TraitSocket) Connect(address IsSocketAddress, cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_connect(self.CPointer, address.GetSocketAddressPointer(), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a #GSocketConnection subclass of the right type for
@socket.
*/
func (self *TraitSocket) ConnectionFactoryCreateConnection() (return__ *SocketConnection) {
	var __cgo__return__ *C.GSocketConnection
	__cgo__return__ = C.g_socket_connection_factory_create_connection(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewSocketConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a %GSource that can be attached to a %GMainContext to monitor
for the availability of the specified @condition on the socket.

The callback on the source is of the #GSocketSourceFunc type.

It is meaningless to specify %G_IO_ERR or %G_IO_HUP in @condition;
these conditions will always be reported output if they are true.

@cancellable if not %NULL can be used to cancel the source, which will
cause the source to trigger, reporting the current condition (which
is likely 0 unless cancellation happened at the same time as a
condition change). You can check for this in the callback using
g_cancellable_is_cancelled().

If @socket has a timeout set, and it is reached before @condition
occurs, the source will then trigger anyway, reporting %G_IO_IN or
%G_IO_OUT depending on @condition. However, @socket will have been
marked as having had a timeout, and so the next #GSocket I/O method
you call will then fail with a %G_IO_ERROR_TIMED_OUT.
*/
func (self *TraitSocket) CreateSource(condition C.GIOCondition, cancellable IsCancellable) (return__ *C.GSource) {
	return__ = C.g_socket_create_source(self.CPointer, condition, cancellable.GetCancellablePointer())
	return
}

/*
Get the amount of data pending in the OS input buffer.

If @socket is a UDP or SCTP socket, this will return the size of
just the next packet, even if additional packets are buffered after
that one.

Note that on Windows, this function is rather inefficient in the
UDP case, and so if you know any plausible upper bound on the size
of the incoming packet, it is better to just do a
g_socket_receive() with a buffer of that size, rather than calling
g_socket_get_available_bytes() first and then doing a receive of
exactly the right size.
*/
func (self *TraitSocket) GetAvailableBytes() (return__ int64) {
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_socket_get_available_bytes(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

/*
Gets the blocking mode of the socket. For details on blocking I/O,
see g_socket_set_blocking().
*/
func (self *TraitSocket) GetBlocking() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_get_blocking(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the broadcast setting on @socket; if %TRUE,
it is possible to send packets to broadcast
addresses.
*/
func (self *TraitSocket) GetBroadcast() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_get_broadcast(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the credentials of the foreign process connected to this
socket, if any (e.g. it is only supported for %G_SOCKET_FAMILY_UNIX
sockets).

If this operation isn't supported on the OS, the method fails with
the %G_IO_ERROR_NOT_SUPPORTED error. On Linux this is implemented
by reading the %SO_PEERCRED option on the underlying socket.

Other ways to obtain credentials from a foreign peer includes the
#GUnixCredentialsMessage type and
g_unix_connection_send_credentials() /
g_unix_connection_receive_credentials() functions.
*/
func (self *TraitSocket) GetCredentials() (return__ *Credentials, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GCredentials
	__cgo__return__ = C.g_socket_get_credentials(self.CPointer, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewCredentialsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the socket family of the socket.
*/
func (self *TraitSocket) GetFamily() (return__ C.GSocketFamily) {
	return__ = C.g_socket_get_family(self.CPointer)
	return
}

/*
Returns the underlying OS socket object. On unix this
is a socket file descriptor, and on Windows this is
a Winsock2 SOCKET handle. This may be useful for
doing platform specific or otherwise unusual operations
on the socket.
*/
func (self *TraitSocket) GetFd() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.g_socket_get_fd(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the keepalive mode of the socket. For details on this,
see g_socket_set_keepalive().
*/
func (self *TraitSocket) GetKeepalive() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_get_keepalive(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the listen backlog setting of the socket. For details on this,
see g_socket_set_listen_backlog().
*/
func (self *TraitSocket) GetListenBacklog() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_socket_get_listen_backlog(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Try to get the local address of a bound socket. This is only
useful if the socket has been bound to a local address,
either explicitly or implicitly when connecting.
*/
func (self *TraitSocket) GetLocalAddress() (return__ *SocketAddress, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocketAddress
	__cgo__return__ = C.g_socket_get_local_address(self.CPointer, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the multicast loopback setting on @socket; if %TRUE (the
default), outgoing multicast packets will be looped back to
multicast listeners on the same host.
*/
func (self *TraitSocket) GetMulticastLoopback() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_get_multicast_loopback(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the multicast time-to-live setting on @socket; see
g_socket_set_multicast_ttl() for more details.
*/
func (self *TraitSocket) GetMulticastTtl() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_socket_get_multicast_ttl(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Gets the value of an integer-valued option on @socket, as with
getsockopt(). (If you need to fetch a  non-integer-valued option,
you will need to call getsockopt() directly.)

The [<gio/gnetworking.h>][gio-gnetworking.h]
header pulls in system headers that will define most of the
standard/portable socket options. For unusual socket protocols or
platform-dependent options, you may need to include additional
headers.

Note that even for socket options that are a single byte in size,
@value is still a pointer to a #gint variable, not a #guchar;
g_socket_get_option() will handle the conversion internally.
*/
func (self *TraitSocket) GetOption(level int, optname int) (value int, return__ bool, __err__ error) {
	var __cgo__value C.gint
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_get_option(self.CPointer, C.gint(level), C.gint(optname), &__cgo__value, &__cgo_error__)
	value = int(__cgo__value)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the socket protocol id the socket was created with.
In case the protocol is unknown, -1 is returned.
*/
func (self *TraitSocket) GetProtocol() (return__ C.GSocketProtocol) {
	return__ = C.g_socket_get_protocol(self.CPointer)
	return
}

/*
Try to get the remove address of a connected socket. This is only
useful for connection oriented sockets that have been connected.
*/
func (self *TraitSocket) GetRemoteAddress() (return__ *SocketAddress, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocketAddress
	__cgo__return__ = C.g_socket_get_remote_address(self.CPointer, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the socket type of the socket.
*/
func (self *TraitSocket) GetSocketType() (return__ C.GSocketType) {
	return__ = C.g_socket_get_socket_type(self.CPointer)
	return
}

/*
Gets the timeout setting of the socket. For details on this, see
g_socket_set_timeout().
*/
func (self *TraitSocket) GetTimeout() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_socket_get_timeout(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Gets the unicast time-to-live setting on @socket; see
g_socket_set_ttl() for more details.
*/
func (self *TraitSocket) GetTtl() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_socket_get_ttl(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Checks whether a socket is closed.
*/
func (self *TraitSocket) IsClosed() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_is_closed(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Check whether the socket is connected. This is only useful for
connection-oriented sockets.
*/
func (self *TraitSocket) IsConnected() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_is_connected(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Registers @socket to receive multicast messages sent to @group.
@socket must be a %G_SOCKET_TYPE_DATAGRAM socket, and must have
been bound to an appropriate interface and port with
g_socket_bind().

If @iface is %NULL, the system will automatically pick an interface
to bind to based on @group.

If @source_specific is %TRUE, source-specific multicast as defined
in RFC 4604 is used. Note that on older platforms this may fail
with a %G_IO_ERROR_NOT_SUPPORTED error.
*/
func (self *TraitSocket) JoinMulticastGroup(group IsInetAddress, source_specific bool, iface string) (return__ bool, __err__ error) {
	__cgo__source_specific := C.gboolean(0)
	if source_specific {
		__cgo__source_specific = C.gboolean(1)
	}
	__cgo__iface := (*C.gchar)(unsafe.Pointer(C.CString(iface)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_join_multicast_group(self.CPointer, group.GetInetAddressPointer(), __cgo__source_specific, __cgo__iface, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__iface))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Removes @socket from the multicast group defined by @group, @iface,
and @source_specific (which must all have the same values they had
when you joined the group).

@socket remains bound to its address and port, and can still receive
unicast messages after calling this.
*/
func (self *TraitSocket) LeaveMulticastGroup(group IsInetAddress, source_specific bool, iface string) (return__ bool, __err__ error) {
	__cgo__source_specific := C.gboolean(0)
	if source_specific {
		__cgo__source_specific = C.gboolean(1)
	}
	__cgo__iface := (*C.gchar)(unsafe.Pointer(C.CString(iface)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_leave_multicast_group(self.CPointer, group.GetInetAddressPointer(), __cgo__source_specific, __cgo__iface, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__iface))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Marks the socket as a server socket, i.e. a socket that is used
to accept incoming requests using g_socket_accept().

Before calling this the socket must be bound to a local address using
g_socket_bind().

To set the maximum amount of outstanding clients, use
g_socket_set_listen_backlog().
*/
func (self *TraitSocket) Listen() (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_listen(self.CPointer, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Receive data (up to @size bytes) from a socket. This is mainly used by
connection-oriented sockets; it is identical to g_socket_receive_from()
with @address set to %NULL.

For %G_SOCKET_TYPE_DATAGRAM and %G_SOCKET_TYPE_SEQPACKET sockets,
g_socket_receive() will always read either 0 or 1 complete messages from
the socket. If the received message is too large to fit in @buffer, then
the data beyond @size bytes will be discarded, without any explicit
indication that this has occurred.

For %G_SOCKET_TYPE_STREAM sockets, g_socket_receive() can return any
number of bytes, up to @size. If more than @size bytes have been
received, the additional data will be returned in future calls to
g_socket_receive().

If the socket is in blocking mode the call will block until there
is some data to receive, the connection is closed, or there is an
error. If there is no data available and the socket is in
non-blocking mode, a %G_IO_ERROR_WOULD_BLOCK error will be
returned. To be notified when data is available, wait for the
%G_IO_IN condition.

On error -1 is returned and @error is set accordingly.
*/
func (self *TraitSocket) Receive(buffer []byte, size int64, cancellable IsCancellable) (return__ int64, __err__ error) {
	__header__buffer := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_socket_receive(self.CPointer, (*C.gchar)(unsafe.Pointer(__header__buffer.Data)), C.gsize(size), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Receive data (up to @size bytes) from a socket.

If @address is non-%NULL then @address will be set equal to the
source address of the received packet.
@address is owned by the caller.

See g_socket_receive() for additional information.
*/
func (self *TraitSocket) ReceiveFrom(buffer []byte, size int64, cancellable IsCancellable) (address *SocketAddress, return__ int64, __err__ error) {
	var __cgo__address *C.GSocketAddress
	__header__buffer := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_socket_receive_from(self.CPointer, &__cgo__address, (*C.gchar)(unsafe.Pointer(__header__buffer.Data)), C.gsize(size), cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo__address != nil {
		address = NewSocketAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__address).Pointer()))
	}
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Receive data from a socket.  This is the most complicated and
fully-featured version of this call. For easier use, see
g_socket_receive() and g_socket_receive_from().

If @address is non-%NULL then @address will be set equal to the
source address of the received packet.
@address is owned by the caller.

@vector must point to an array of #GInputVector structs and
@num_vectors must be the length of this array.  These structs
describe the buffers that received data will be scattered into.
If @num_vectors is -1, then @vectors is assumed to be terminated
by a #GInputVector with a %NULL buffer pointer.

As a special case, if @num_vectors is 0 (in which case, @vectors
may of course be %NULL), then a single byte is received and
discarded. This is to facilitate the common practice of sending a
single '\0' byte for the purposes of transferring ancillary data.

@messages, if non-%NULL, will be set to point to a newly-allocated
array of #GSocketControlMessage instances or %NULL if no such
messages was received. These correspond to the control messages
received from the kernel, one #GSocketControlMessage per message
from the kernel. This array is %NULL-terminated and must be freed
by the caller using g_free() after calling g_object_unref() on each
element. If @messages is %NULL, any control messages received will
be discarded.

@num_messages, if non-%NULL, will be set to the number of control
messages received.

If both @messages and @num_messages are non-%NULL, then
@num_messages gives the number of #GSocketControlMessage instances
in @messages (ie: not including the %NULL terminator).

@flags is an in/out parameter. The commonly available arguments
for this are available in the #GSocketMsgFlags enum, but the
values there are the same as the system values, and the flags
are passed in as-is, so you can pass in system-specific flags too
(and g_socket_receive_message() may pass system-specific flags out).

As with g_socket_receive(), data may be discarded if @socket is
%G_SOCKET_TYPE_DATAGRAM or %G_SOCKET_TYPE_SEQPACKET and you do not
provide enough buffer space to read a complete message. You can pass
%G_SOCKET_MSG_PEEK in @flags to peek at the current message without
removing it from the receive queue, but there is no portable way to find
out the length of the message other than by reading it into a
sufficiently-large buffer.

If the socket is in blocking mode the call will block until there
is some data to receive, the connection is closed, or there is an
error. If there is no data available and the socket is in
non-blocking mode, a %G_IO_ERROR_WOULD_BLOCK error will be
returned. To be notified when data is available, wait for the
%G_IO_IN condition.

On error -1 is returned and @error is set accordingly.
*/
func (self *TraitSocket) ReceiveMessage(vectors *C.GInputVector, num_vectors int, messages ***C.GSocketControlMessage, num_messages *C.gint, flags *C.gint, cancellable IsCancellable) (address *SocketAddress, return__ int64, __err__ error) {
	var __cgo__address *C.GSocketAddress
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_socket_receive_message(self.CPointer, &__cgo__address, vectors, C.gint(num_vectors), messages, num_messages, flags, cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo__address != nil {
		address = NewSocketAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__address).Pointer()))
	}
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This behaves exactly the same as g_socket_receive(), except that
the choice of blocking or non-blocking behavior is determined by
the @blocking argument rather than by @socket's properties.
*/
func (self *TraitSocket) ReceiveWithBlocking(buffer []byte, size int64, blocking bool, cancellable IsCancellable) (return__ int64, __err__ error) {
	__header__buffer := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	__cgo__blocking := C.gboolean(0)
	if blocking {
		__cgo__blocking = C.gboolean(1)
	}
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_socket_receive_with_blocking(self.CPointer, (*C.gchar)(unsafe.Pointer(__header__buffer.Data)), C.gsize(size), __cgo__blocking, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Tries to send @size bytes from @buffer on the socket. This is
mainly used by connection-oriented sockets; it is identical to
g_socket_send_to() with @address set to %NULL.

If the socket is in blocking mode the call will block until there is
space for the data in the socket queue. If there is no space available
and the socket is in non-blocking mode a %G_IO_ERROR_WOULD_BLOCK error
will be returned. To be notified when space is available, wait for the
%G_IO_OUT condition. Note though that you may still receive
%G_IO_ERROR_WOULD_BLOCK from g_socket_send() even if you were previously
notified of a %G_IO_OUT condition. (On Windows in particular, this is
very common due to the way the underlying APIs work.)

On error -1 is returned and @error is set accordingly.
*/
func (self *TraitSocket) Send(buffer []byte, size int64, cancellable IsCancellable) (return__ int64, __err__ error) {
	__header__buffer := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_socket_send(self.CPointer, (*C.gchar)(unsafe.Pointer(__header__buffer.Data)), C.gsize(size), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Send data to @address on @socket.  This is the most complicated and
fully-featured version of this call. For easier use, see
g_socket_send() and g_socket_send_to().

If @address is %NULL then the message is sent to the default receiver
(set by g_socket_connect()).

@vectors must point to an array of #GOutputVector structs and
@num_vectors must be the length of this array. (If @num_vectors is -1,
then @vectors is assumed to be terminated by a #GOutputVector with a
%NULL buffer pointer.) The #GOutputVector structs describe the buffers
that the sent data will be gathered from. Using multiple
#GOutputVectors is more memory-efficient than manually copying
data from multiple sources into a single buffer, and more
network-efficient than making multiple calls to g_socket_send().

@messages, if non-%NULL, is taken to point to an array of @num_messages
#GSocketControlMessage instances. These correspond to the control
messages to be sent on the socket.
If @num_messages is -1 then @messages is treated as a %NULL-terminated
array.

@flags modify how the message is sent. The commonly available arguments
for this are available in the #GSocketMsgFlags enum, but the
values there are the same as the system values, and the flags
are passed in as-is, so you can pass in system-specific flags too.

If the socket is in blocking mode the call will block until there is
space for the data in the socket queue. If there is no space available
and the socket is in non-blocking mode a %G_IO_ERROR_WOULD_BLOCK error
will be returned. To be notified when space is available, wait for the
%G_IO_OUT condition. Note though that you may still receive
%G_IO_ERROR_WOULD_BLOCK from g_socket_send() even if you were previously
notified of a %G_IO_OUT condition. (On Windows in particular, this is
very common due to the way the underlying APIs work.)

On error -1 is returned and @error is set accordingly.
*/
func (self *TraitSocket) SendMessage(address IsSocketAddress, vectors *C.GOutputVector, num_vectors int, messages **C.GSocketControlMessage, num_messages int, flags int, cancellable IsCancellable) (return__ int64, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_socket_send_message(self.CPointer, address.GetSocketAddressPointer(), vectors, C.gint(num_vectors), messages, C.gint(num_messages), C.gint(flags), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Tries to send @size bytes from @buffer to @address. If @address is
%NULL then the message is sent to the default receiver (set by
g_socket_connect()).

See g_socket_send() for additional information.
*/
func (self *TraitSocket) SendTo(address IsSocketAddress, buffer []byte, size int64, cancellable IsCancellable) (return__ int64, __err__ error) {
	__header__buffer := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_socket_send_to(self.CPointer, address.GetSocketAddressPointer(), (*C.gchar)(unsafe.Pointer(__header__buffer.Data)), C.gsize(size), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This behaves exactly the same as g_socket_send(), except that
the choice of blocking or non-blocking behavior is determined by
the @blocking argument rather than by @socket's properties.
*/
func (self *TraitSocket) SendWithBlocking(buffer []byte, size int64, blocking bool, cancellable IsCancellable) (return__ int64, __err__ error) {
	__header__buffer := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	__cgo__blocking := C.gboolean(0)
	if blocking {
		__cgo__blocking = C.gboolean(1)
	}
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_socket_send_with_blocking(self.CPointer, (*C.gchar)(unsafe.Pointer(__header__buffer.Data)), C.gsize(size), __cgo__blocking, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Sets the blocking mode of the socket. In blocking mode
all operations block until they succeed or there is an error. In
non-blocking mode all functions return results immediately or
with a %G_IO_ERROR_WOULD_BLOCK error.

All sockets are created in blocking mode. However, note that the
platform level socket is always non-blocking, and blocking mode
is a GSocket level feature.
*/
func (self *TraitSocket) SetBlocking(blocking bool) {
	__cgo__blocking := C.gboolean(0)
	if blocking {
		__cgo__blocking = C.gboolean(1)
	}
	C.g_socket_set_blocking(self.CPointer, __cgo__blocking)
	return
}

/*
Sets whether @socket should allow sending to broadcast addresses.
This is %FALSE by default.
*/
func (self *TraitSocket) SetBroadcast(broadcast bool) {
	__cgo__broadcast := C.gboolean(0)
	if broadcast {
		__cgo__broadcast = C.gboolean(1)
	}
	C.g_socket_set_broadcast(self.CPointer, __cgo__broadcast)
	return
}

/*
Sets or unsets the %SO_KEEPALIVE flag on the underlying socket. When
this flag is set on a socket, the system will attempt to verify that the
remote socket endpoint is still present if a sufficiently long period of
time passes with no data being exchanged. If the system is unable to
verify the presence of the remote endpoint, it will automatically close
the connection.

This option is only functional on certain kinds of sockets. (Notably,
%G_SOCKET_PROTOCOL_TCP sockets.)

The exact time between pings is system- and protocol-dependent, but will
normally be at least two hours. Most commonly, you would set this flag
on a server socket if you want to allow clients to remain idle for long
periods of time, but also want to ensure that connections are eventually
garbage-collected if clients crash or become unreachable.
*/
func (self *TraitSocket) SetKeepalive(keepalive bool) {
	__cgo__keepalive := C.gboolean(0)
	if keepalive {
		__cgo__keepalive = C.gboolean(1)
	}
	C.g_socket_set_keepalive(self.CPointer, __cgo__keepalive)
	return
}

/*
Sets the maximum number of outstanding connections allowed
when listening on this socket. If more clients than this are
connecting to the socket and the application is not handling them
on time then the new connections will be refused.

Note that this must be called before g_socket_listen() and has no
effect if called after that.
*/
func (self *TraitSocket) SetListenBacklog(backlog int) {
	C.g_socket_set_listen_backlog(self.CPointer, C.gint(backlog))
	return
}

/*
Sets whether outgoing multicast packets will be received by sockets
listening on that multicast address on the same host. This is %TRUE
by default.
*/
func (self *TraitSocket) SetMulticastLoopback(loopback bool) {
	__cgo__loopback := C.gboolean(0)
	if loopback {
		__cgo__loopback = C.gboolean(1)
	}
	C.g_socket_set_multicast_loopback(self.CPointer, __cgo__loopback)
	return
}

/*
Sets the time-to-live for outgoing multicast datagrams on @socket.
By default, this is 1, meaning that multicast packets will not leave
the local network.
*/
func (self *TraitSocket) SetMulticastTtl(ttl uint) {
	C.g_socket_set_multicast_ttl(self.CPointer, C.guint(ttl))
	return
}

/*
Sets the value of an integer-valued option on @socket, as with
setsockopt(). (If you need to set a non-integer-valued option,
you will need to call setsockopt() directly.)

The [<gio/gnetworking.h>][gio-gnetworking.h]
header pulls in system headers that will define most of the
standard/portable socket options. For unusual socket protocols or
platform-dependent options, you may need to include additional
headers.
*/
func (self *TraitSocket) SetOption(level int, optname int, value int) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_set_option(self.CPointer, C.gint(level), C.gint(optname), C.gint(value), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Sets the time in seconds after which I/O operations on @socket will
time out if they have not yet completed.

On a blocking socket, this means that any blocking #GSocket
operation will time out after @timeout seconds of inactivity,
returning %G_IO_ERROR_TIMED_OUT.

On a non-blocking socket, calls to g_socket_condition_wait() will
also fail with %G_IO_ERROR_TIMED_OUT after the given time. Sources
created with g_socket_create_source() will trigger after
@timeout seconds of inactivity, with the requested condition
set, at which point calling g_socket_receive(), g_socket_send(),
g_socket_check_connect_result(), etc, will fail with
%G_IO_ERROR_TIMED_OUT.

If @timeout is 0 (the default), operations will never time out
on their own.

Note that if an I/O operation is interrupted by a signal, this may
cause the timeout to be reset.
*/
func (self *TraitSocket) SetTimeout(timeout uint) {
	C.g_socket_set_timeout(self.CPointer, C.guint(timeout))
	return
}

/*
Sets the time-to-live for outgoing unicast packets on @socket.
By default the platform-specific default value is used.
*/
func (self *TraitSocket) SetTtl(ttl uint) {
	C.g_socket_set_ttl(self.CPointer, C.guint(ttl))
	return
}

/*
Shut down part of a full-duplex connection.

If @shutdown_read is %TRUE then the receiving side of the connection
is shut down, and further reading is disallowed.

If @shutdown_write is %TRUE then the sending side of the connection
is shut down, and further writing is disallowed.

It is allowed for both @shutdown_read and @shutdown_write to be %TRUE.

One example where this is used is graceful disconnect for TCP connections
where you close the sending side, then wait for the other side to close
the connection, thus ensuring that the other side saw all sent data.
*/
func (self *TraitSocket) Shutdown(shutdown_read bool, shutdown_write bool) (return__ bool, __err__ error) {
	__cgo__shutdown_read := C.gboolean(0)
	if shutdown_read {
		__cgo__shutdown_read = C.gboolean(1)
	}
	__cgo__shutdown_write := C.gboolean(0)
	if shutdown_write {
		__cgo__shutdown_write = C.gboolean(1)
	}
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_shutdown(self.CPointer, __cgo__shutdown_read, __cgo__shutdown_write, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Checks if a socket is capable of speaking IPv4.

IPv4 sockets are capable of speaking IPv4.  On some operating systems
and under some combinations of circumstances IPv6 sockets are also
capable of speaking IPv4.  See RFC 3493 section 3.7 for more
information.

No other types of sockets are currently considered as being capable
of speaking IPv4.
*/
func (self *TraitSocket) SpeaksIpv4() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_speaks_ipv4(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

type TraitSocketAddress struct{ CPointer *C.GSocketAddress }
type IsSocketAddress interface {
	GetSocketAddressPointer() *C.GSocketAddress
}

func (self *TraitSocketAddress) GetSocketAddressPointer() *C.GSocketAddress {
	return self.CPointer
}
func NewTraitSocketAddress(p unsafe.Pointer) *TraitSocketAddress {
	return &TraitSocketAddress{(*C.GSocketAddress)(p)}
}

/*
Gets the socket family type of @address.
*/
func (self *TraitSocketAddress) GetFamily() (return__ C.GSocketFamily) {
	return__ = C.g_socket_address_get_family(self.CPointer)
	return
}

/*
Gets the size of @address's native struct sockaddr.
You can use this to allocate memory to pass to
g_socket_address_to_native().
*/
func (self *TraitSocketAddress) GetNativeSize() (return__ int64) {
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_socket_address_get_native_size(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

/*
Converts a #GSocketAddress to a native struct sockaddr, which can
be passed to low-level functions like connect() or bind().

If not enough space is available, a %G_IO_ERROR_NO_SPACE error
is returned. If the address type is not known on the system
then a %G_IO_ERROR_NOT_SUPPORTED error is returned.
*/
func (self *TraitSocketAddress) ToNative(dest unsafe.Pointer, destlen int64) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_address_to_native(self.CPointer, (C.gpointer)(dest), C.gsize(destlen), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

type TraitSocketAddressEnumerator struct{ CPointer *C.GSocketAddressEnumerator }
type IsSocketAddressEnumerator interface {
	GetSocketAddressEnumeratorPointer() *C.GSocketAddressEnumerator
}

func (self *TraitSocketAddressEnumerator) GetSocketAddressEnumeratorPointer() *C.GSocketAddressEnumerator {
	return self.CPointer
}
func NewTraitSocketAddressEnumerator(p unsafe.Pointer) *TraitSocketAddressEnumerator {
	return &TraitSocketAddressEnumerator{(*C.GSocketAddressEnumerator)(p)}
}

/*
Retrieves the next #GSocketAddress from @enumerator. Note that this
may block for some amount of time. (Eg, a #GNetworkAddress may need
to do a DNS lookup before it can return an address.) Use
g_socket_address_enumerator_next_async() if you need to avoid
blocking.

If @enumerator is expected to yield addresses, but for some reason
is unable to (eg, because of a DNS error), then the first call to
g_socket_address_enumerator_next() will return an appropriate error
in *@error. However, if the first call to
g_socket_address_enumerator_next() succeeds, then any further
internal errors (other than @cancellable being triggered) will be
ignored.
*/
func (self *TraitSocketAddressEnumerator) Next(cancellable IsCancellable) (return__ *SocketAddress, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocketAddress
	__cgo__return__ = C.g_socket_address_enumerator_next(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously retrieves the next #GSocketAddress from @enumerator
and then calls @callback, which must call
g_socket_address_enumerator_next_finish() to get the result.
*/
func (self *TraitSocketAddressEnumerator) NextAsync(cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_socket_address_enumerator_next_async(self.CPointer, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Retrieves the result of a completed call to
g_socket_address_enumerator_next_async(). See
g_socket_address_enumerator_next() for more information about
error handling.
*/
func (self *TraitSocketAddressEnumerator) NextFinish(result *C.GAsyncResult) (return__ *SocketAddress, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocketAddress
	__cgo__return__ = C.g_socket_address_enumerator_next_finish(self.CPointer, result, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

type TraitSocketClient struct{ CPointer *C.GSocketClient }
type IsSocketClient interface {
	GetSocketClientPointer() *C.GSocketClient
}

func (self *TraitSocketClient) GetSocketClientPointer() *C.GSocketClient {
	return self.CPointer
}
func NewTraitSocketClient(p unsafe.Pointer) *TraitSocketClient {
	return &TraitSocketClient{(*C.GSocketClient)(p)}
}

/*
Enable proxy protocols to be handled by the application. When the
indicated proxy protocol is returned by the #GProxyResolver,
#GSocketClient will consider this protocol as supported but will
not try to find a #GProxy instance to handle handshaking. The
application must check for this case by calling
g_socket_connection_get_remote_address() on the returned
#GSocketConnection, and seeing if it's a #GProxyAddress of the
appropriate type, to determine whether or not it needs to handle
the proxy handshaking itself.

This should be used for proxy protocols that are dialects of
another protocol such as HTTP proxy. It also allows cohabitation of
proxy protocols that are reused between protocols. A good example
is HTTP. It can be used to proxy HTTP, FTP and Gopher and can also
be use as generic socket proxy through the HTTP CONNECT method.

When the proxy is detected as being an application proxy, TLS handshake
will be skipped. This is required to let the application do the proxy
specific handshake.
*/
func (self *TraitSocketClient) AddApplicationProxy(protocol string) {
	__cgo__protocol := (*C.gchar)(unsafe.Pointer(C.CString(protocol)))
	C.g_socket_client_add_application_proxy(self.CPointer, __cgo__protocol)
	C.free(unsafe.Pointer(__cgo__protocol))
	return
}

/*
Tries to resolve the @connectable and make a network connection to it.

Upon a successful connection, a new #GSocketConnection is constructed
and returned.  The caller owns this new object and must drop their
reference to it when finished with it.

The type of the #GSocketConnection object returned depends on the type of
the underlying socket that is used. For instance, for a TCP/IP connection
it will be a #GTcpConnection.

The socket created will be the same family as the address that the
@connectable resolves to, unless family is set with g_socket_client_set_family()
or indirectly via g_socket_client_set_local_address(). The socket type
defaults to %G_SOCKET_TYPE_STREAM but can be set with
g_socket_client_set_socket_type().

If a local address is specified with g_socket_client_set_local_address() the
socket will be bound to this address before connecting.
*/
func (self *TraitSocketClient) Connect(connectable *C.GSocketConnectable, cancellable IsCancellable) (return__ *SocketConnection, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocketConnection
	__cgo__return__ = C.g_socket_client_connect(self.CPointer, connectable, cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This is the asynchronous version of g_socket_client_connect().

When the operation is finished @callback will be
called. You can then call g_socket_client_connect_finish() to get
the result of the operation.
*/
func (self *TraitSocketClient) ConnectAsync(connectable *C.GSocketConnectable, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_socket_client_connect_async(self.CPointer, connectable, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes an async connect operation. See g_socket_client_connect_async()
*/
func (self *TraitSocketClient) ConnectFinish(result *C.GAsyncResult) (return__ *SocketConnection, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocketConnection
	__cgo__return__ = C.g_socket_client_connect_finish(self.CPointer, result, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This is a helper function for g_socket_client_connect().

Attempts to create a TCP connection to the named host.

@host_and_port may be in any of a number of recognized formats; an IPv6
address, an IPv4 address, or a domain name (in which case a DNS
lookup is performed).  Quoting with [] is supported for all address
types.  A port override may be specified in the usual way with a
colon.  Ports may be given as decimal numbers or symbolic names (in
which case an /etc/services lookup is performed).

If no port override is given in @host_and_port then @default_port will be
used as the port number to connect to.

In general, @host_and_port is expected to be provided by the user (allowing
them to give the hostname, and a port override if necessary) and
@default_port is expected to be provided by the application.

In the case that an IP address is given, a single connection
attempt is made.  In the case that a name is given, multiple
connection attempts may be made, in turn and according to the
number of address records in DNS, until a connection succeeds.

Upon a successful connection, a new #GSocketConnection is constructed
and returned.  The caller owns this new object and must drop their
reference to it when finished with it.

In the event of any failure (DNS error, service not found, no hosts
connectable) %NULL is returned and @error (if non-%NULL) is set
accordingly.
*/
func (self *TraitSocketClient) ConnectToHost(host_and_port string, default_port uint16, cancellable IsCancellable) (return__ *SocketConnection, __err__ error) {
	__cgo__host_and_port := (*C.gchar)(unsafe.Pointer(C.CString(host_and_port)))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocketConnection
	__cgo__return__ = C.g_socket_client_connect_to_host(self.CPointer, __cgo__host_and_port, C.guint16(default_port), cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__host_and_port))
	if __cgo__return__ != nil {
		return__ = NewSocketConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This is the asynchronous version of g_socket_client_connect_to_host().

When the operation is finished @callback will be
called. You can then call g_socket_client_connect_to_host_finish() to get
the result of the operation.
*/
func (self *TraitSocketClient) ConnectToHostAsync(host_and_port string, default_port uint16, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__host_and_port := (*C.gchar)(unsafe.Pointer(C.CString(host_and_port)))
	C.g_socket_client_connect_to_host_async(self.CPointer, __cgo__host_and_port, C.guint16(default_port), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__host_and_port))
	return
}

/*
Finishes an async connect operation. See g_socket_client_connect_to_host_async()
*/
func (self *TraitSocketClient) ConnectToHostFinish(result *C.GAsyncResult) (return__ *SocketConnection, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocketConnection
	__cgo__return__ = C.g_socket_client_connect_to_host_finish(self.CPointer, result, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Attempts to create a TCP connection to a service.

This call looks up the SRV record for @service at @domain for the
"tcp" protocol.  It then attempts to connect, in turn, to each of
the hosts providing the service until either a connection succeeds
or there are no hosts remaining.

Upon a successful connection, a new #GSocketConnection is constructed
and returned.  The caller owns this new object and must drop their
reference to it when finished with it.

In the event of any failure (DNS error, service not found, no hosts
connectable) %NULL is returned and @error (if non-%NULL) is set
accordingly.
*/
func (self *TraitSocketClient) ConnectToService(domain string, service string, cancellable IsCancellable) (return__ *SocketConnection, __err__ error) {
	__cgo__domain := (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	__cgo__service := (*C.gchar)(unsafe.Pointer(C.CString(service)))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocketConnection
	__cgo__return__ = C.g_socket_client_connect_to_service(self.CPointer, __cgo__domain, __cgo__service, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__domain))
	C.free(unsafe.Pointer(__cgo__service))
	if __cgo__return__ != nil {
		return__ = NewSocketConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This is the asynchronous version of
g_socket_client_connect_to_service().
*/
func (self *TraitSocketClient) ConnectToServiceAsync(domain string, service string, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__domain := (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	__cgo__service := (*C.gchar)(unsafe.Pointer(C.CString(service)))
	C.g_socket_client_connect_to_service_async(self.CPointer, __cgo__domain, __cgo__service, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__domain))
	C.free(unsafe.Pointer(__cgo__service))
	return
}

/*
Finishes an async connect operation. See g_socket_client_connect_to_service_async()
*/
func (self *TraitSocketClient) ConnectToServiceFinish(result *C.GAsyncResult) (return__ *SocketConnection, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocketConnection
	__cgo__return__ = C.g_socket_client_connect_to_service_finish(self.CPointer, result, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This is a helper function for g_socket_client_connect().

Attempts to create a TCP connection with a network URI.

@uri may be any valid URI containing an "authority" (hostname/port)
component. If a port is not specified in the URI, @default_port
will be used. TLS will be negotiated if #GSocketClient:tls is %TRUE.
(#GSocketClient does not know to automatically assume TLS for
certain URI schemes.)

Using this rather than g_socket_client_connect() or
g_socket_client_connect_to_host() allows #GSocketClient to
determine when to use application-specific proxy protocols.

Upon a successful connection, a new #GSocketConnection is constructed
and returned.  The caller owns this new object and must drop their
reference to it when finished with it.

In the event of any failure (DNS error, service not found, no hosts
connectable) %NULL is returned and @error (if non-%NULL) is set
accordingly.
*/
func (self *TraitSocketClient) ConnectToUri(uri string, default_port uint16, cancellable IsCancellable) (return__ *SocketConnection, __err__ error) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocketConnection
	__cgo__return__ = C.g_socket_client_connect_to_uri(self.CPointer, __cgo__uri, C.guint16(default_port), cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__uri))
	if __cgo__return__ != nil {
		return__ = NewSocketConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This is the asynchronous version of g_socket_client_connect_to_uri().

When the operation is finished @callback will be
called. You can then call g_socket_client_connect_to_uri_finish() to get
the result of the operation.
*/
func (self *TraitSocketClient) ConnectToUriAsync(uri string, default_port uint16, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	C.g_socket_client_connect_to_uri_async(self.CPointer, __cgo__uri, C.guint16(default_port), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__uri))
	return
}

/*
Finishes an async connect operation. See g_socket_client_connect_to_uri_async()
*/
func (self *TraitSocketClient) ConnectToUriFinish(result *C.GAsyncResult) (return__ *SocketConnection, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocketConnection
	__cgo__return__ = C.g_socket_client_connect_to_uri_finish(self.CPointer, result, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the proxy enable state; see g_socket_client_set_enable_proxy()
*/
func (self *TraitSocketClient) GetEnableProxy() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_client_get_enable_proxy(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the socket family of the socket client.

See g_socket_client_set_family() for details.
*/
func (self *TraitSocketClient) GetFamily() (return__ C.GSocketFamily) {
	return__ = C.g_socket_client_get_family(self.CPointer)
	return
}

/*
Gets the local address of the socket client.

See g_socket_client_set_local_address() for details.
*/
func (self *TraitSocketClient) GetLocalAddress() (return__ *SocketAddress) {
	var __cgo__return__ *C.GSocketAddress
	__cgo__return__ = C.g_socket_client_get_local_address(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewSocketAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the protocol name type of the socket client.

See g_socket_client_set_protocol() for details.
*/
func (self *TraitSocketClient) GetProtocol() (return__ C.GSocketProtocol) {
	return__ = C.g_socket_client_get_protocol(self.CPointer)
	return
}

/*
Gets the #GProxyResolver being used by @client. Normally, this will
be the resolver returned by g_proxy_resolver_get_default(), but you
can override it with g_socket_client_set_proxy_resolver().
*/
func (self *TraitSocketClient) GetProxyResolver() (return__ *C.GProxyResolver) {
	return__ = C.g_socket_client_get_proxy_resolver(self.CPointer)
	return
}

/*
Gets the socket type of the socket client.

See g_socket_client_set_socket_type() for details.
*/
func (self *TraitSocketClient) GetSocketType() (return__ C.GSocketType) {
	return__ = C.g_socket_client_get_socket_type(self.CPointer)
	return
}

/*
Gets the I/O timeout time for sockets created by @client.

See g_socket_client_set_timeout() for details.
*/
func (self *TraitSocketClient) GetTimeout() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_socket_client_get_timeout(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Gets whether @client creates TLS connections. See
g_socket_client_set_tls() for details.
*/
func (self *TraitSocketClient) GetTls() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_client_get_tls(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the TLS validation flags used creating TLS connections via
@client.
*/
func (self *TraitSocketClient) GetTlsValidationFlags() (return__ C.GTlsCertificateFlags) {
	return__ = C.g_socket_client_get_tls_validation_flags(self.CPointer)
	return
}

/*
Sets whether or not @client attempts to make connections via a
proxy server. When enabled (the default), #GSocketClient will use a
#GProxyResolver to determine if a proxy protocol such as SOCKS is
needed, and automatically do the necessary proxy negotiation.

See also g_socket_client_set_proxy_resolver().
*/
func (self *TraitSocketClient) SetEnableProxy(enable bool) {
	__cgo__enable := C.gboolean(0)
	if enable {
		__cgo__enable = C.gboolean(1)
	}
	C.g_socket_client_set_enable_proxy(self.CPointer, __cgo__enable)
	return
}

/*
Sets the socket family of the socket client.
If this is set to something other than %G_SOCKET_FAMILY_INVALID
then the sockets created by this object will be of the specified
family.

This might be useful for instance if you want to force the local
connection to be an ipv4 socket, even though the address might
be an ipv6 mapped to ipv4 address.
*/
func (self *TraitSocketClient) SetFamily(family C.GSocketFamily) {
	C.g_socket_client_set_family(self.CPointer, family)
	return
}

/*
Sets the local address of the socket client.
The sockets created by this object will bound to the
specified address (if not %NULL) before connecting.

This is useful if you want to ensure that the local
side of the connection is on a specific port, or on
a specific interface.
*/
func (self *TraitSocketClient) SetLocalAddress(address IsSocketAddress) {
	C.g_socket_client_set_local_address(self.CPointer, address.GetSocketAddressPointer())
	return
}

/*
Sets the protocol of the socket client.
The sockets created by this object will use of the specified
protocol.

If @protocol is %0 that means to use the default
protocol for the socket family and type.
*/
func (self *TraitSocketClient) SetProtocol(protocol C.GSocketProtocol) {
	C.g_socket_client_set_protocol(self.CPointer, protocol)
	return
}

/*
Overrides the #GProxyResolver used by @client. You can call this if
you want to use specific proxies, rather than using the system
default proxy settings.

Note that whether or not the proxy resolver is actually used
depends on the setting of #GSocketClient:enable-proxy, which is not
changed by this function (but which is %TRUE by default)
*/
func (self *TraitSocketClient) SetProxyResolver(proxy_resolver *C.GProxyResolver) {
	C.g_socket_client_set_proxy_resolver(self.CPointer, proxy_resolver)
	return
}

/*
Sets the socket type of the socket client.
The sockets created by this object will be of the specified
type.

It doesn't make sense to specify a type of %G_SOCKET_TYPE_DATAGRAM,
as GSocketClient is used for connection oriented services.
*/
func (self *TraitSocketClient) SetSocketType(type_ C.GSocketType) {
	C.g_socket_client_set_socket_type(self.CPointer, type_)
	return
}

/*
Sets the I/O timeout for sockets created by @client. @timeout is a
time in seconds, or 0 for no timeout (the default).

The timeout value affects the initial connection attempt as well,
so setting this may cause calls to g_socket_client_connect(), etc,
to fail with %G_IO_ERROR_TIMED_OUT.
*/
func (self *TraitSocketClient) SetTimeout(timeout uint) {
	C.g_socket_client_set_timeout(self.CPointer, C.guint(timeout))
	return
}

/*
Sets whether @client creates TLS (aka SSL) connections. If @tls is
%TRUE, @client will wrap its connections in a #GTlsClientConnection
and perform a TLS handshake when connecting.

Note that since #GSocketClient must return a #GSocketConnection,
but #GTlsClientConnection is not a #GSocketConnection, this
actually wraps the resulting #GTlsClientConnection in a
#GTcpWrapperConnection when returning it. You can use
g_tcp_wrapper_connection_get_base_io_stream() on the return value
to extract the #GTlsClientConnection.

If you need to modify the behavior of the TLS handshake (eg, by
setting a client-side certificate to use, or connecting to the
#GTlsConnection::accept-certificate signal), you can connect to
@client's #GSocketClient::event signal and wait for it to be
emitted with %G_SOCKET_CLIENT_TLS_HANDSHAKING, which will give you
a chance to see the #GTlsClientConnection before the handshake
starts.
*/
func (self *TraitSocketClient) SetTls(tls bool) {
	__cgo__tls := C.gboolean(0)
	if tls {
		__cgo__tls = C.gboolean(1)
	}
	C.g_socket_client_set_tls(self.CPointer, __cgo__tls)
	return
}

/*
Sets the TLS validation flags used when creating TLS connections
via @client. The default value is %G_TLS_CERTIFICATE_VALIDATE_ALL.
*/
func (self *TraitSocketClient) SetTlsValidationFlags(flags C.GTlsCertificateFlags) {
	C.g_socket_client_set_tls_validation_flags(self.CPointer, flags)
	return
}

type TraitSocketConnection struct{ CPointer *C.GSocketConnection }
type IsSocketConnection interface {
	GetSocketConnectionPointer() *C.GSocketConnection
}

func (self *TraitSocketConnection) GetSocketConnectionPointer() *C.GSocketConnection {
	return self.CPointer
}
func NewTraitSocketConnection(p unsafe.Pointer) *TraitSocketConnection {
	return &TraitSocketConnection{(*C.GSocketConnection)(p)}
}

/*
Connect @connection to the specified remote address.
*/
func (self *TraitSocketConnection) Connect(address IsSocketAddress, cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_connection_connect(self.CPointer, address.GetSocketAddressPointer(), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously connect @connection to the specified remote address.

This clears the #GSocket:blocking flag on @connection's underlying
socket if it is currently set.

Use g_socket_connection_connect_finish() to retrieve the result.
*/
func (self *TraitSocketConnection) ConnectAsync(address IsSocketAddress, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_socket_connection_connect_async(self.CPointer, address.GetSocketAddressPointer(), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Gets the result of a g_socket_connection_connect_async() call.
*/
func (self *TraitSocketConnection) ConnectFinish(result *C.GAsyncResult) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_connection_connect_finish(self.CPointer, result, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Try to get the local address of a socket connection.
*/
func (self *TraitSocketConnection) GetLocalAddress() (return__ *SocketAddress, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocketAddress
	__cgo__return__ = C.g_socket_connection_get_local_address(self.CPointer, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Try to get the remote address of a socket connection.

Since GLib 2.40, when used with g_socket_client_connect() or
g_socket_client_connect_async(), during emission of
%G_SOCKET_CLIENT_CONNECTING, this function will return the remote
address that will be used for the connection.  This allows
applications to print e.g. "Connecting to example.com
(10.42.77.3)...".
*/
func (self *TraitSocketConnection) GetRemoteAddress() (return__ *SocketAddress, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocketAddress
	__cgo__return__ = C.g_socket_connection_get_remote_address(self.CPointer, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the underlying #GSocket object of the connection.
This can be useful if you want to do something unusual on it
not supported by the #GSocketConnection APIs.
*/
func (self *TraitSocketConnection) GetSocket() (return__ *Socket) {
	var __cgo__return__ *C.GSocket
	__cgo__return__ = C.g_socket_connection_get_socket(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewSocketFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Checks if @connection is connected. This is equivalent to calling
g_socket_is_connected() on @connection's underlying #GSocket.
*/
func (self *TraitSocketConnection) IsConnected() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_connection_is_connected(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

type TraitSocketControlMessage struct{ CPointer *C.GSocketControlMessage }
type IsSocketControlMessage interface {
	GetSocketControlMessagePointer() *C.GSocketControlMessage
}

func (self *TraitSocketControlMessage) GetSocketControlMessagePointer() *C.GSocketControlMessage {
	return self.CPointer
}
func NewTraitSocketControlMessage(p unsafe.Pointer) *TraitSocketControlMessage {
	return &TraitSocketControlMessage{(*C.GSocketControlMessage)(p)}
}

/*
Returns the "level" (i.e. the originating protocol) of the control message.
This is often SOL_SOCKET.
*/
func (self *TraitSocketControlMessage) GetLevel() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.g_socket_control_message_get_level(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the protocol specific type of the control message.
For instance, for UNIX fd passing this would be SCM_RIGHTS.
*/
func (self *TraitSocketControlMessage) GetMsgType() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.g_socket_control_message_get_msg_type(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the space required for the control message, not including
headers or alignment.
*/
func (self *TraitSocketControlMessage) GetSize() (return__ int64) {
	var __cgo__return__ C.gsize
	__cgo__return__ = C.g_socket_control_message_get_size(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

/*
Converts the data in the message to bytes placed in the
message.

@data is guaranteed to have enough space to fit the size
returned by g_socket_control_message_get_size() on this
object.
*/
func (self *TraitSocketControlMessage) Serialize(data unsafe.Pointer) {
	C.g_socket_control_message_serialize(self.CPointer, (C.gpointer)(data))
	return
}

type TraitSocketListener struct{ CPointer *C.GSocketListener }
type IsSocketListener interface {
	GetSocketListenerPointer() *C.GSocketListener
}

func (self *TraitSocketListener) GetSocketListenerPointer() *C.GSocketListener {
	return self.CPointer
}
func NewTraitSocketListener(p unsafe.Pointer) *TraitSocketListener {
	return &TraitSocketListener{(*C.GSocketListener)(p)}
}

/*
Blocks waiting for a client to connect to any of the sockets added
to the listener. Returns a #GSocketConnection for the socket that was
accepted.

If @source_object is not %NULL it will be filled out with the source
object specified when the corresponding socket or address was added
to the listener.

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
*/
func (self *TraitSocketListener) Accept(cancellable IsCancellable) (source_object *C.GObject, return__ *SocketConnection, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocketConnection
	__cgo__return__ = C.g_socket_listener_accept(self.CPointer, &source_object, cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This is the asynchronous version of g_socket_listener_accept().

When the operation is finished @callback will be
called. You can then call g_socket_listener_accept_socket()
to get the result of the operation.
*/
func (self *TraitSocketListener) AcceptAsync(cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_socket_listener_accept_async(self.CPointer, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes an async accept operation. See g_socket_listener_accept_async()
*/
func (self *TraitSocketListener) AcceptFinish(result *C.GAsyncResult) (source_object *C.GObject, return__ *SocketConnection, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocketConnection
	__cgo__return__ = C.g_socket_listener_accept_finish(self.CPointer, result, &source_object, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketConnectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Blocks waiting for a client to connect to any of the sockets added
to the listener. Returns the #GSocket that was accepted.

If you want to accept the high-level #GSocketConnection, not a #GSocket,
which is often the case, then you should use g_socket_listener_accept()
instead.

If @source_object is not %NULL it will be filled out with the source
object specified when the corresponding socket or address was added
to the listener.

If @cancellable is not %NULL, then the operation can be cancelled by
triggering the cancellable object from another thread. If the operation
was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
*/
func (self *TraitSocketListener) AcceptSocket(cancellable IsCancellable) (source_object *C.GObject, return__ *Socket, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocket
	__cgo__return__ = C.g_socket_listener_accept_socket(self.CPointer, &source_object, cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This is the asynchronous version of g_socket_listener_accept_socket().

When the operation is finished @callback will be
called. You can then call g_socket_listener_accept_socket_finish()
to get the result of the operation.
*/
func (self *TraitSocketListener) AcceptSocketAsync(cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_socket_listener_accept_socket_async(self.CPointer, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes an async accept operation. See g_socket_listener_accept_socket_async()
*/
func (self *TraitSocketListener) AcceptSocketFinish(result *C.GAsyncResult) (source_object *C.GObject, return__ *Socket, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSocket
	__cgo__return__ = C.g_socket_listener_accept_socket_finish(self.CPointer, result, &source_object, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSocketFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a socket of type @type and protocol @protocol, binds
it to @address and adds it to the set of sockets we're accepting
sockets from.

Note that adding an IPv6 address, depending on the platform,
may or may not result in a listener that also accepts IPv4
connections.  For more deterministic behavior, see
g_socket_listener_add_inet_port().

@source_object will be passed out in the various calls
to accept to identify this particular source, which is
useful if you're listening on multiple addresses and do
different things depending on what address is connected to.

If successful and @effective_address is non-%NULL then it will
be set to the address that the binding actually occurred at.  This
is helpful for determining the port number that was used for when
requesting a binding to port 0 (ie: "any port").  This address, if
requested, belongs to the caller and must be freed.
*/
func (self *TraitSocketListener) AddAddress(address IsSocketAddress, type_ C.GSocketType, protocol C.GSocketProtocol, source_object *C.GObject) (effective_address *SocketAddress, return__ bool, __err__ error) {
	var __cgo__effective_address *C.GSocketAddress
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_listener_add_address(self.CPointer, address.GetSocketAddressPointer(), type_, protocol, source_object, &__cgo__effective_address, &__cgo_error__)
	if __cgo__effective_address != nil {
		effective_address = NewSocketAddressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__effective_address).Pointer()))
	}
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Listens for TCP connections on any available port number for both
IPv6 and IPv4 (if each is available).

This is useful if you need to have a socket for incoming connections
but don't care about the specific port number.

@source_object will be passed out in the various calls
to accept to identify this particular source, which is
useful if you're listening on multiple addresses and do
different things depending on what address is connected to.
*/
func (self *TraitSocketListener) AddAnyInetPort(source_object *C.GObject) (return__ uint16, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.guint16
	__cgo__return__ = C.g_socket_listener_add_any_inet_port(self.CPointer, source_object, &__cgo_error__)
	return__ = uint16(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Helper function for g_socket_listener_add_address() that
creates a TCP/IP socket listening on IPv4 and IPv6 (if
supported) on the specified port on all interfaces.

@source_object will be passed out in the various calls
to accept to identify this particular source, which is
useful if you're listening on multiple addresses and do
different things depending on what address is connected to.
*/
func (self *TraitSocketListener) AddInetPort(port uint16, source_object *C.GObject) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_listener_add_inet_port(self.CPointer, C.guint16(port), source_object, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Adds @socket to the set of sockets that we try to accept
new clients from. The socket must be bound to a local
address and listened to.

@source_object will be passed out in the various calls
to accept to identify this particular source, which is
useful if you're listening on multiple addresses and do
different things depending on what address is connected to.
*/
func (self *TraitSocketListener) AddSocket(socket IsSocket, source_object *C.GObject) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_listener_add_socket(self.CPointer, socket.GetSocketPointer(), source_object, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Closes all the sockets in the listener.
*/
func (self *TraitSocketListener) Close() {
	C.g_socket_listener_close(self.CPointer)
	return
}

/*
Sets the listen backlog on the sockets in the listener.

See g_socket_set_listen_backlog() for details
*/
func (self *TraitSocketListener) SetBacklog(listen_backlog int) {
	C.g_socket_listener_set_backlog(self.CPointer, C.int(listen_backlog))
	return
}

type TraitSocketService struct{ CPointer *C.GSocketService }
type IsSocketService interface {
	GetSocketServicePointer() *C.GSocketService
}

func (self *TraitSocketService) GetSocketServicePointer() *C.GSocketService {
	return self.CPointer
}
func NewTraitSocketService(p unsafe.Pointer) *TraitSocketService {
	return &TraitSocketService{(*C.GSocketService)(p)}
}

/*
Check whether the service is active or not. An active
service will accept new clients that connect, while
a non-active service will let connecting clients queue
up until the service is started.
*/
func (self *TraitSocketService) IsActive() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_socket_service_is_active(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Starts the service, i.e. start accepting connections
from the added sockets when the mainloop runs.

This call is thread-safe, so it may be called from a thread
handling an incoming client request.
*/
func (self *TraitSocketService) Start() {
	C.g_socket_service_start(self.CPointer)
	return
}

/*
Stops the service, i.e. stops accepting connections
from the added sockets when the mainloop runs.

This call is thread-safe, so it may be called from a thread
handling an incoming client request.

Note that this only stops accepting new connections; it does not
close the listening sockets, and you can call
g_socket_service_start() again later to begin listening again. To
close the listening sockets, call g_socket_listener_close(). (This
will happen automatically when the #GSocketService is finalized.)
*/
func (self *TraitSocketService) Stop() {
	C.g_socket_service_stop(self.CPointer)
	return
}

type TraitSubprocess struct{ CPointer *C.GSubprocess }
type IsSubprocess interface {
	GetSubprocessPointer() *C.GSubprocess
}

func (self *TraitSubprocess) GetSubprocessPointer() *C.GSubprocess {
	return self.CPointer
}
func NewTraitSubprocess(p unsafe.Pointer) *TraitSubprocess {
	return &TraitSubprocess{(*C.GSubprocess)(p)}
}

/*
Communicate with the subprocess until it terminates, and all input
and output has been completed.

If @stdin_buf is given, the subprocess must have been created with
%G_SUBPROCESS_FLAGS_STDIN_PIPE.  The given data is fed to the
stdin of the subprocess and the pipe is closed (ie: EOF).

At the same time (as not to cause blocking when dealing with large
amounts of data), if %G_SUBPROCESS_FLAGS_STDOUT_PIPE or
%G_SUBPROCESS_FLAGS_STDERR_PIPE were used, reads from those
streams.  The data that was read is returned in @stdout and/or
the @stderr.

If the subprocess was created with %G_SUBPROCESS_FLAGS_STDOUT_PIPE,
@stdout_buf will contain the data read from stdout.  Otherwise, for
subprocesses not created with %G_SUBPROCESS_FLAGS_STDOUT_PIPE,
@stdout_buf will be set to %NULL.  Similar provisions apply to
@stderr_buf and %G_SUBPROCESS_FLAGS_STDERR_PIPE.

As usual, any output variable may be given as %NULL to ignore it.

If you desire the stdout and stderr data to be interleaved, create
the subprocess with %G_SUBPROCESS_FLAGS_STDOUT_PIPE and
%G_SUBPROCESS_FLAGS_STDERR_MERGE.  The merged result will be returned
in @stdout_buf and @stderr_buf will be set to %NULL.

In case of any error (including cancellation), %FALSE will be
returned with @error set.  Some or all of the stdin data may have
been written.  Any stdout or stderr data that has been read will be
discarded. None of the out variables (aside from @error) will have
been set to anything in particular and should not be inspected.

In the case that %TRUE is returned, the subprocess has exited and the
exit status inspection APIs (eg: g_subprocess_get_if_exited(),
g_subprocess_get_exit_status()) may be used.

You should not attempt to use any of the subprocess pipes after
starting this function, since they may be left in strange states,
even if the operation was cancelled.  You should especially not
attempt to interact with the pipes while the operation is in progress
(either from another thread or if using the asynchronous version).
*/
func (self *TraitSubprocess) Communicate(stdin_buf *C.GBytes, cancellable IsCancellable) (stdout_buf *C.GBytes, stderr_buf *C.GBytes, return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_subprocess_communicate(self.CPointer, stdin_buf, cancellable.GetCancellablePointer(), &stdout_buf, &stderr_buf, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronous version of g_subprocess_communicate().  Complete
invocation with g_subprocess_communicate_finish().
*/
func (self *TraitSubprocess) CommunicateAsync(stdin_buf *C.GBytes, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_subprocess_communicate_async(self.CPointer, stdin_buf, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Complete an invocation of g_subprocess_communicate_async().
*/
func (self *TraitSubprocess) CommunicateFinish(result *C.GAsyncResult) (stdout_buf *C.GBytes, stderr_buf *C.GBytes, return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_subprocess_communicate_finish(self.CPointer, result, &stdout_buf, &stderr_buf, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Like g_subprocess_communicate(), but validates the output of the
process as UTF-8, and returns it as a regular NUL terminated string.
*/
func (self *TraitSubprocess) CommunicateUtf8(stdin_buf string, cancellable IsCancellable) (stdout_buf string, stderr_buf string, return__ bool, __err__ error) {
	__cgo__stdin_buf := C.CString(stdin_buf)
	var __cgo__stdout_buf *C.char
	var __cgo__stderr_buf *C.char
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_subprocess_communicate_utf8(self.CPointer, __cgo__stdin_buf, cancellable.GetCancellablePointer(), &__cgo__stdout_buf, &__cgo__stderr_buf, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__stdin_buf))
	stdout_buf = C.GoString(__cgo__stdout_buf)
	stderr_buf = C.GoString(__cgo__stderr_buf)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronous version of g_subprocess_communicate_utf().  Complete
invocation with g_subprocess_communicate_utf8_finish().
*/
func (self *TraitSubprocess) CommunicateUtf8Async(stdin_buf string, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__stdin_buf := C.CString(stdin_buf)
	C.g_subprocess_communicate_utf8_async(self.CPointer, __cgo__stdin_buf, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__stdin_buf))
	return
}

/*
Complete an invocation of g_subprocess_communicate_utf8_async().
*/
func (self *TraitSubprocess) CommunicateUtf8Finish(result *C.GAsyncResult) (stdout_buf string, stderr_buf string, return__ bool, __err__ error) {
	var __cgo__stdout_buf *C.char
	var __cgo__stderr_buf *C.char
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_subprocess_communicate_utf8_finish(self.CPointer, result, &__cgo__stdout_buf, &__cgo__stderr_buf, &__cgo_error__)
	stdout_buf = C.GoString(__cgo__stdout_buf)
	stderr_buf = C.GoString(__cgo__stderr_buf)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Use an operating-system specific method to attempt an immediate,
forceful termination of the process.  There is no mechanism to
determine whether or not the request itself was successful;
however, you can use g_subprocess_wait() to monitor the status of
the process after calling this function.

On Unix, this function sends %SIGKILL.
*/
func (self *TraitSubprocess) ForceExit() {
	C.g_subprocess_force_exit(self.CPointer)
	return
}

/*
Check the exit status of the subprocess, given that it exited
normally.  This is the value passed to the exit() system call or the
return value from main.

This is equivalent to the system WEXITSTATUS macro.

It is an error to call this function before g_subprocess_wait() and
unless g_subprocess_get_if_exited() returned %TRUE.
*/
func (self *TraitSubprocess) GetExitStatus() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_subprocess_get_exit_status(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
On UNIX, returns the process ID as a decimal string.
On Windows, returns the result of GetProcessId() also as a string.
*/
func (self *TraitSubprocess) GetIdentifier() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_subprocess_get_identifier(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Check if the given subprocess exited normally (ie: by way of exit()
or return from main()).

This is equivalent to the system WIFEXITED macro.

It is an error to call this function before g_subprocess_wait() has
returned.
*/
func (self *TraitSubprocess) GetIfExited() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_subprocess_get_if_exited(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Check if the given subprocess terminated in response to a signal.

This is equivalent to the system WIFSIGNALED macro.

It is an error to call this function before g_subprocess_wait() has
returned.
*/
func (self *TraitSubprocess) GetIfSignaled() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_subprocess_get_if_signaled(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the raw status code of the process, as from waitpid().

This value has no particular meaning, but it can be used with the
macros defined by the system headers such as WIFEXITED.  It can also
be used with g_spawn_check_exit_status().

It is more likely that you want to use g_subprocess_get_if_exited()
followed by g_subprocess_get_exit_status().

It is an error to call this function before g_subprocess_wait() has
returned.
*/
func (self *TraitSubprocess) GetStatus() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_subprocess_get_status(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the #GInputStream from which to read the stderr output of
@subprocess.

The process must have been created with
%G_SUBPROCESS_FLAGS_STDERR_PIPE.
*/
func (self *TraitSubprocess) GetStderrPipe() (return__ *InputStream) {
	var __cgo__return__ *C.GInputStream
	__cgo__return__ = C.g_subprocess_get_stderr_pipe(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewInputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the #GOutputStream that you can write to in order to give data
to the stdin of @subprocess.

The process must have been created with
%G_SUBPROCESS_FLAGS_STDIN_PIPE.
*/
func (self *TraitSubprocess) GetStdinPipe() (return__ *OutputStream) {
	var __cgo__return__ *C.GOutputStream
	__cgo__return__ = C.g_subprocess_get_stdin_pipe(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewOutputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the #GInputStream from which to read the stdout output of
@subprocess.

The process must have been created with
%G_SUBPROCESS_FLAGS_STDOUT_PIPE.
*/
func (self *TraitSubprocess) GetStdoutPipe() (return__ *InputStream) {
	var __cgo__return__ *C.GInputStream
	__cgo__return__ = C.g_subprocess_get_stdout_pipe(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewInputStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Checks if the process was "successful".  A process is considered
successful if it exited cleanly with an exit status of 0, either by
way of the exit() system call or return from main().

It is an error to call this function before g_subprocess_wait() has
returned.
*/
func (self *TraitSubprocess) GetSuccessful() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_subprocess_get_successful(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the signal number that caused the subprocess to terminate, given
that it terminated due to a signal.

This is equivalent to the system WTERMSIG macro.

It is an error to call this function before g_subprocess_wait() and
unless g_subprocess_get_if_signaled() returned %TRUE.
*/
func (self *TraitSubprocess) GetTermSig() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_subprocess_get_term_sig(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Sends the UNIX signal @signal_num to the subprocess, if it is still
running.

This API is race-free.  If the subprocess has terminated, it will not
be signalled.

This API is not available on Windows.
*/
func (self *TraitSubprocess) SendSignal(signal_num int) {
	C.g_subprocess_send_signal(self.CPointer, C.gint(signal_num))
	return
}

/*
Synchronously wait for the subprocess to terminate.

After the process terminates you can query its exit status with
functions such as g_subprocess_get_if_exited() and
g_subprocess_get_exit_status().

This function does not fail in the case of the subprocess having
abnormal termination.  See g_subprocess_wait_check() for that.
*/
func (self *TraitSubprocess) Wait(cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_subprocess_wait(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Wait for the subprocess to terminate.

This is the asynchronous version of g_subprocess_wait().
*/
func (self *TraitSubprocess) WaitAsync(cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_subprocess_wait_async(self.CPointer, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Combines g_subprocess_wait() with g_spawn_check_exit_status().
*/
func (self *TraitSubprocess) WaitCheck(cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_subprocess_wait_check(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Combines g_subprocess_wait_async() with g_spawn_check_exit_status().

This is the asynchronous version of g_subprocess_wait_check().
*/
func (self *TraitSubprocess) WaitCheckAsync(cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_subprocess_wait_check_async(self.CPointer, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Collects the result of a previous call to
g_subprocess_wait_check_async().
*/
func (self *TraitSubprocess) WaitCheckFinish(result *C.GAsyncResult) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_subprocess_wait_check_finish(self.CPointer, result, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Collects the result of a previous call to
g_subprocess_wait_async().
*/
func (self *TraitSubprocess) WaitFinish(result *C.GAsyncResult) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_subprocess_wait_finish(self.CPointer, result, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

type TraitSubprocessLauncher struct{ CPointer *C.GSubprocessLauncher }
type IsSubprocessLauncher interface {
	GetSubprocessLauncherPointer() *C.GSubprocessLauncher
}

func (self *TraitSubprocessLauncher) GetSubprocessLauncherPointer() *C.GSubprocessLauncher {
	return self.CPointer
}
func NewTraitSubprocessLauncher(p unsafe.Pointer) *TraitSubprocessLauncher {
	return &TraitSubprocessLauncher{(*C.GSubprocessLauncher)(p)}
}

/*
Returns the value of the environment variable @variable in the
environment of processes launched from this launcher.

The returned string is in the GLib file name encoding.  On UNIX, this
means that it can be an arbitrary byte string.  On Windows, it will
be UTF-8.
*/
func (self *TraitSubprocessLauncher) Getenv(variable string) (return__ string) {
	__cgo__variable := (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_subprocess_launcher_getenv(self.CPointer, __cgo__variable)
	C.free(unsafe.Pointer(__cgo__variable))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Sets up a child setup function.

The child setup function will be called after fork() but before
exec() on the child's side.

@destroy_notify will not be automatically called on the child's side
of the fork().  It will only be called when the last reference on the
#GSubprocessLauncher is dropped or when a new child setup function is
given.

%NULL can be given as @child_setup to disable the functionality.

Child setup functions are only available on UNIX.
*/
func (self *TraitSubprocessLauncher) SetChildSetup(child_setup C.GSpawnChildSetupFunc, user_data unsafe.Pointer, destroy_notify C.GDestroyNotify) {
	C.g_subprocess_launcher_set_child_setup(self.CPointer, child_setup, (C.gpointer)(user_data), destroy_notify)
	return
}

/*
Sets the current working directory that processes will be launched
with.

By default processes are launched with the current working directory
of the launching process at the time of launch.
*/
func (self *TraitSubprocessLauncher) SetCwd(cwd string) {
	__cgo__cwd := (*C.gchar)(unsafe.Pointer(C.CString(cwd)))
	C.g_subprocess_launcher_set_cwd(self.CPointer, __cgo__cwd)
	C.free(unsafe.Pointer(__cgo__cwd))
	return
}

/*
Replace the entire environment of processes launched from this
launcher with the given 'environ' variable.

Typically you will build this variable by using g_listenv() to copy
the process 'environ' and using the functions g_environ_setenv(),
g_environ_unsetenv(), etc.

As an alternative, you can use g_subprocess_launcher_setenv(),
g_subprocess_launcher_unsetenv(), etc.

All strings in this array are expected to be in the GLib file name
encoding.  On UNIX, this means that they can be arbitrary byte
strings.  On Windows, they should be in UTF-8.
*/
func (self *TraitSubprocessLauncher) SetEnviron(env []string) {
	__header__env := (*reflect.SliceHeader)(unsafe.Pointer(&env))
	C.g_subprocess_launcher_set_environ(self.CPointer, (**C.gchar)(unsafe.Pointer(__header__env.Data)))
	return
}

/*
Sets the flags on the launcher.

The default flags are %G_SUBPROCESS_FLAGS_NONE.

You may not set flags that specify conflicting options for how to
handle a particular stdio stream (eg: specifying both
%G_SUBPROCESS_FLAGS_STDIN_PIPE and
%G_SUBPROCESS_FLAGS_STDIN_INHERIT).

You may also not set a flag that conflicts with a previous call to a
function like g_subprocess_launcher_set_stdin_file_path() or
g_subprocess_launcher_take_stdout_fd().
*/
func (self *TraitSubprocessLauncher) SetFlags(flags C.GSubprocessFlags) {
	C.g_subprocess_launcher_set_flags(self.CPointer, flags)
	return
}

/*
Sets the file path to use as the stderr for spawned processes.

If @path is %NULL then any previously given path is unset.

The file will be created or truncated when the process is spawned, as
would be the case if using '2>' at the shell.

If you want to send both stdout and stderr to the same file then use
%G_SUBPROCESS_FLAGS_STDERR_MERGE.

You may not set a stderr file path if a stderr fd is already set or
if the launcher flags contain any flags directing stderr elsewhere.

This feature is only available on UNIX.
*/
func (self *TraitSubprocessLauncher) SetStderrFilePath(path string) {
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	C.g_subprocess_launcher_set_stderr_file_path(self.CPointer, __cgo__path)
	C.free(unsafe.Pointer(__cgo__path))
	return
}

/*
Sets the file path to use as the stdin for spawned processes.

If @path is %NULL then any previously given path is unset.

The file must exist or spawning the process will fail.

You may not set a stdin file path if a stdin fd is already set or if
the launcher flags contain any flags directing stdin elsewhere.

This feature is only available on UNIX.
*/
func (self *TraitSubprocessLauncher) SetStdinFilePath(path string) {
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	C.g_subprocess_launcher_set_stdin_file_path(self.CPointer, __cgo__path)
	C.free(unsafe.Pointer(__cgo__path))
	return
}

/*
Sets the file path to use as the stdout for spawned processes.

If @path is %NULL then any previously given path is unset.

The file will be created or truncated when the process is spawned, as
would be the case if using '>' at the shell.

You may not set a stdout file path if a stdout fd is already set or
if the launcher flags contain any flags directing stdout elsewhere.

This feature is only available on UNIX.
*/
func (self *TraitSubprocessLauncher) SetStdoutFilePath(path string) {
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	C.g_subprocess_launcher_set_stdout_file_path(self.CPointer, __cgo__path)
	C.free(unsafe.Pointer(__cgo__path))
	return
}

/*
Sets the environment variable @variable in the environment of
processes launched from this launcher.

Both the variable's name and value should be in the GLib file name
encoding. On UNIX, this means that they can be arbitrary byte
strings. On Windows, they should be in UTF-8.
*/
func (self *TraitSubprocessLauncher) Setenv(variable string, value string, overwrite bool) {
	__cgo__variable := (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	__cgo__value := (*C.gchar)(unsafe.Pointer(C.CString(value)))
	__cgo__overwrite := C.gboolean(0)
	if overwrite {
		__cgo__overwrite = C.gboolean(1)
	}
	C.g_subprocess_launcher_setenv(self.CPointer, __cgo__variable, __cgo__value, __cgo__overwrite)
	C.free(unsafe.Pointer(__cgo__variable))
	C.free(unsafe.Pointer(__cgo__value))
	return
}

// g_subprocess_launcher_spawn is not generated due to varargs

/*
A convenience helper for creating a #GSubprocess given a provided
array of arguments.
*/
func (self *TraitSubprocessLauncher) Spawnv(argv []string) (return__ *Subprocess, __err__ error) {
	__header__argv := (*reflect.SliceHeader)(unsafe.Pointer(&argv))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GSubprocess
	__cgo__return__ = C.g_subprocess_launcher_spawnv(self.CPointer, (**C.gchar)(unsafe.Pointer(__header__argv.Data)), &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewSubprocessFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Transfer an arbitrary file descriptor from parent process to the
child.  This function takes "ownership" of the fd; it will be closed
in the parent when @self is freed.

By default, all file descriptors from the parent will be closed.
This function allows you to create (for example) a custom pipe() or
socketpair() before launching the process, and choose the target
descriptor in the child.

An example use case is GNUPG, which has a command line argument
--passphrase-fd providing a file descriptor number where it expects
the passphrase to be written.
*/
func (self *TraitSubprocessLauncher) TakeFd(source_fd int, target_fd int) {
	C.g_subprocess_launcher_take_fd(self.CPointer, C.gint(source_fd), C.gint(target_fd))
	return
}

/*
Sets the file descriptor to use as the stderr for spawned processes.

If @fd is -1 then any previously given fd is unset.

Note that the default behaviour is to pass stderr through to the
stderr of the parent process.

The passed @fd belongs to the #GSubprocessLauncher.  It will be
automatically closed when the launcher is finalized.  The file
descriptor will also be closed on the child side when executing the
spawned process.

You may not set a stderr fd if a stderr file path is already set or
if the launcher flags contain any flags directing stderr elsewhere.

This feature is only available on UNIX.
*/
func (self *TraitSubprocessLauncher) TakeStderrFd(fd int) {
	C.g_subprocess_launcher_take_stderr_fd(self.CPointer, C.gint(fd))
	return
}

/*
Sets the file descriptor to use as the stdin for spawned processes.

If @fd is -1 then any previously given fd is unset.

Note that if your intention is to have the stdin of the calling
process inherited by the child then %G_SUBPROCESS_FLAGS_STDIN_INHERIT
is a better way to go about doing that.

The passed @fd is noted but will not be touched in the current
process.  It is therefore necessary that it be kept open by the
caller until the subprocess is spawned.  The file descriptor will
also not be explicitly closed on the child side, so it must be marked
O_CLOEXEC if that's what you want.

You may not set a stdin fd if a stdin file path is already set or if
the launcher flags contain any flags directing stdin elsewhere.

This feature is only available on UNIX.
*/
func (self *TraitSubprocessLauncher) TakeStdinFd(fd int) {
	C.g_subprocess_launcher_take_stdin_fd(self.CPointer, C.gint(fd))
	return
}

/*
Sets the file descriptor to use as the stdout for spawned processes.

If @fd is -1 then any previously given fd is unset.

Note that the default behaviour is to pass stdout through to the
stdout of the parent process.

The passed @fd is noted but will not be touched in the current
process.  It is therefore necessary that it be kept open by the
caller until the subprocess is spawned.  The file descriptor will
also not be explicitly closed on the child side, so it must be marked
O_CLOEXEC if that's what you want.

You may not set a stdout fd if a stdout file path is already set or
if the launcher flags contain any flags directing stdout elsewhere.

This feature is only available on UNIX.
*/
func (self *TraitSubprocessLauncher) TakeStdoutFd(fd int) {
	C.g_subprocess_launcher_take_stdout_fd(self.CPointer, C.gint(fd))
	return
}

/*
Removes the environment variable @variable from the environment of
processes launched from this launcher.

The variable name should be in the GLib file name encoding.  On UNIX,
this means that they can be arbitrary byte strings.  On Windows, they
should be in UTF-8.
*/
func (self *TraitSubprocessLauncher) Unsetenv(variable string) {
	__cgo__variable := (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	C.g_subprocess_launcher_unsetenv(self.CPointer, __cgo__variable)
	C.free(unsafe.Pointer(__cgo__variable))
	return
}

type TraitTask struct{ CPointer *C.GTask }
type IsTask interface {
	GetTaskPointer() *C.GTask
}

func (self *TraitTask) GetTaskPointer() *C.GTask {
	return self.CPointer
}
func NewTraitTask(p unsafe.Pointer) *TraitTask {
	return &TraitTask{(*C.GTask)(p)}
}

/*
A utility function for dealing with async operations where you need
to wait for a #GSource to trigger. Attaches @source to @task's
#GMainContext with @task's [priority][io-priority], and sets @source's
callback to @callback, with @task as the callback's `user_data`.

This takes a reference on @task until @source is destroyed.
*/
func (self *TraitTask) AttachSource(source *C.GSource, callback C.GSourceFunc) {
	C.g_task_attach_source(self.CPointer, source, callback)
	return
}

/*
Gets @task's #GCancellable
*/
func (self *TraitTask) GetCancellable() (return__ *Cancellable) {
	var __cgo__return__ *C.GCancellable
	__cgo__return__ = C.g_task_get_cancellable(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewCancellableFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets @task's check-cancellable flag. See
g_task_set_check_cancellable() for more details.
*/
func (self *TraitTask) GetCheckCancellable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_task_get_check_cancellable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the #GMainContext that @task will return its result in (that
is, the context that was the
[thread-default main context][g-main-context-push-thread-default]
at the point when @task was created).

This will always return a non-%NULL value, even if the task's
context is the default #GMainContext.
*/
func (self *TraitTask) GetContext() (return__ *C.GMainContext) {
	return__ = C.g_task_get_context(self.CPointer)
	return
}

/*
Gets @task's priority
*/
func (self *TraitTask) GetPriority() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_task_get_priority(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets @task's return-on-cancel flag. See
g_task_set_return_on_cancel() for more details.
*/
func (self *TraitTask) GetReturnOnCancel() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_task_get_return_on_cancel(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the source object from @task. Like
g_async_result_get_source_object(), but does not ref the object.
*/
func (self *TraitTask) GetSourceObject() (return__ *gobject.Object) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_task_get_source_object(self.CPointer)
	return__ = gobject.NewObjectFromCPointer(unsafe.Pointer(__cgo__return__))
	return
}

/*
Gets @task's source tag. See g_task_set_source_tag().
*/
func (self *TraitTask) GetSourceTag() (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_task_get_source_tag(self.CPointer)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Gets @task's `task_data`.
*/
func (self *TraitTask) GetTaskData() (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_task_get_task_data(self.CPointer)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Tests if @task resulted in an error.
*/
func (self *TraitTask) HadError() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_task_had_error(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the result of @task as a #gboolean.

If the task resulted in an error, or was cancelled, then this will
instead return %FALSE and set @error.

Since this method transfers ownership of the return value (or
error) to the caller, you may only call it once.
*/
func (self *TraitTask) PropagateBoolean() (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_task_propagate_boolean(self.CPointer, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the result of @task as an integer (#gssize).

If the task resulted in an error, or was cancelled, then this will
instead return -1 and set @error.

Since this method transfers ownership of the return value (or
error) to the caller, you may only call it once.
*/
func (self *TraitTask) PropagateInt() (return__ int64, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_task_propagate_int(self.CPointer, &__cgo_error__)
	return__ = int64(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the result of @task as a pointer, and transfers ownership
of that value to the caller.

If the task resulted in an error, or was cancelled, then this will
instead return %NULL and set @error.

Since this method transfers ownership of the return value (or
error) to the caller, you may only call it once.
*/
func (self *TraitTask) PropagatePointer() (return__ unsafe.Pointer, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_task_propagate_pointer(self.CPointer, &__cgo_error__)
	return__ = unsafe.Pointer(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Sets @task's result to @result and completes the task (see
g_task_return_pointer() for more discussion of exactly what this
means).
*/
func (self *TraitTask) ReturnBoolean(result bool) {
	__cgo__result := C.gboolean(0)
	if result {
		__cgo__result = C.gboolean(1)
	}
	C.g_task_return_boolean(self.CPointer, __cgo__result)
	return
}

/*
Sets @task's result to @error (which @task assumes ownership of)
and completes the task (see g_task_return_pointer() for more
discussion of exactly what this means).

Note that since the task takes ownership of @error, and since the
task may be completed before returning from g_task_return_error(),
you cannot assume that @error is still valid after calling this.
Call g_error_copy() on the error if you need to keep a local copy
as well.

See also g_task_return_new_error().
*/
func (self *TraitTask) ReturnError(error_ *C.GError) {
	C.g_task_return_error(self.CPointer, error_)
	return
}

/*
Checks if @task's #GCancellable has been cancelled, and if so, sets
@task's error accordingly and completes the task (see
g_task_return_pointer() for more discussion of exactly what this
means).
*/
func (self *TraitTask) ReturnErrorIfCancelled() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_task_return_error_if_cancelled(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets @task's result to @result and completes the task (see
g_task_return_pointer() for more discussion of exactly what this
means).
*/
func (self *TraitTask) ReturnInt(result int64) {
	C.g_task_return_int(self.CPointer, C.gssize(result))
	return
}

// g_task_return_new_error is not generated due to varargs

/*
Sets @task's result to @result and completes the task. If @result
is not %NULL, then @result_destroy will be used to free @result if
the caller does not take ownership of it with
g_task_propagate_pointer().

"Completes the task" means that for an ordinary asynchronous task
it will either invoke the task's callback, or else queue that
callback to be invoked in the proper #GMainContext, or in the next
iteration of the current #GMainContext. For a task run via
g_task_run_in_thread() or g_task_run_in_thread_sync(), calling this
method will save @result to be returned to the caller later, but
the task will not actually be completed until the #GTaskThreadFunc
exits.

Note that since the task may be completed before returning from
g_task_return_pointer(), you cannot assume that @result is still
valid after calling this, unless you are still holding another
reference on it.
*/
func (self *TraitTask) ReturnPointer(result unsafe.Pointer, result_destroy C.GDestroyNotify) {
	C.g_task_return_pointer(self.CPointer, (C.gpointer)(result), result_destroy)
	return
}

/*
Runs @task_func in another thread. When @task_func returns, @task's
#GAsyncReadyCallback will be invoked in @task's #GMainContext.

This takes a ref on @task until the task completes.

See #GTaskThreadFunc for more details about how @task_func is handled.
*/
func (self *TraitTask) RunInThread(task_func C.GTaskThreadFunc) {
	C.g_task_run_in_thread(self.CPointer, task_func)
	return
}

/*
Runs @task_func in another thread, and waits for it to return or be
cancelled. You can use g_task_propagate_pointer(), etc, afterward
to get the result of @task_func.

See #GTaskThreadFunc for more details about how @task_func is handled.

Normally this is used with tasks created with a %NULL
`callback`, but note that even if the task does
have a callback, it will not be invoked when @task_func returns.
*/
func (self *TraitTask) RunInThreadSync(task_func C.GTaskThreadFunc) {
	C.g_task_run_in_thread_sync(self.CPointer, task_func)
	return
}

/*
Sets or clears @task's check-cancellable flag. If this is %TRUE
(the default), then g_task_propagate_pointer(), etc, and
g_task_had_error() will check the task's #GCancellable first, and
if it has been cancelled, then they will consider the task to have
returned an "Operation was cancelled" error
(%G_IO_ERROR_CANCELLED), regardless of any other error or return
value the task may have had.

If @check_cancellable is %FALSE, then the #GTask will not check the
cancellable itself, and it is up to @task's owner to do this (eg,
via g_task_return_error_if_cancelled()).

If you are using g_task_set_return_on_cancel() as well, then
you must leave check-cancellable set %TRUE.
*/
func (self *TraitTask) SetCheckCancellable(check_cancellable bool) {
	__cgo__check_cancellable := C.gboolean(0)
	if check_cancellable {
		__cgo__check_cancellable = C.gboolean(1)
	}
	C.g_task_set_check_cancellable(self.CPointer, __cgo__check_cancellable)
	return
}

/*
Sets @task's priority. If you do not call this, it will default to
%G_PRIORITY_DEFAULT.

This will affect the priority of #GSources created with
g_task_attach_source() and the scheduling of tasks run in threads,
and can also be explicitly retrieved later via
g_task_get_priority().
*/
func (self *TraitTask) SetPriority(priority int) {
	C.g_task_set_priority(self.CPointer, C.gint(priority))
	return
}

/*
Sets or clears @task's return-on-cancel flag. This is only
meaningful for tasks run via g_task_run_in_thread() or
g_task_run_in_thread_sync().

If @return_on_cancel is %TRUE, then cancelling @task's
#GCancellable will immediately cause it to return, as though the
task's #GTaskThreadFunc had called
g_task_return_error_if_cancelled() and then returned.

This allows you to create a cancellable wrapper around an
uninterruptable function. The #GTaskThreadFunc just needs to be
careful that it does not modify any externally-visible state after
it has been cancelled. To do that, the thread should call
g_task_set_return_on_cancel() again to (atomically) set
return-on-cancel %FALSE before making externally-visible changes;
if the task gets cancelled before the return-on-cancel flag could
be changed, g_task_set_return_on_cancel() will indicate this by
returning %FALSE.

You can disable and re-enable this flag multiple times if you wish.
If the task's #GCancellable is cancelled while return-on-cancel is
%FALSE, then calling g_task_set_return_on_cancel() to set it %TRUE
again will cause the task to be cancelled at that point.

If the task's #GCancellable is already cancelled before you call
g_task_run_in_thread()/g_task_run_in_thread_sync(), then the
#GTaskThreadFunc will still be run (for consistency), but the task
will also be completed right away.
*/
func (self *TraitTask) SetReturnOnCancel(return_on_cancel bool) (return__ bool) {
	__cgo__return_on_cancel := C.gboolean(0)
	if return_on_cancel {
		__cgo__return_on_cancel = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_task_set_return_on_cancel(self.CPointer, __cgo__return_on_cancel)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets @task's source tag. You can use this to tag a task return
value with a particular pointer (usually a pointer to the function
doing the tagging) and then later check it using
g_task_get_source_tag() (or g_async_result_is_tagged()) in the
task's "finish" function, to figure out if the response came from a
particular place.
*/
func (self *TraitTask) SetSourceTag(source_tag unsafe.Pointer) {
	C.g_task_set_source_tag(self.CPointer, (C.gpointer)(source_tag))
	return
}

/*
Sets @task's task data (freeing the existing task data, if any).
*/
func (self *TraitTask) SetTaskData(task_data unsafe.Pointer, task_data_destroy C.GDestroyNotify) {
	C.g_task_set_task_data(self.CPointer, (C.gpointer)(task_data), task_data_destroy)
	return
}

type TraitTcpConnection struct{ CPointer *C.GTcpConnection }
type IsTcpConnection interface {
	GetTcpConnectionPointer() *C.GTcpConnection
}

func (self *TraitTcpConnection) GetTcpConnectionPointer() *C.GTcpConnection {
	return self.CPointer
}
func NewTraitTcpConnection(p unsafe.Pointer) *TraitTcpConnection {
	return &TraitTcpConnection{(*C.GTcpConnection)(p)}
}

/*
Checks if graceful disconnects are used. See
g_tcp_connection_set_graceful_disconnect().
*/
func (self *TraitTcpConnection) GetGracefulDisconnect() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_tcp_connection_get_graceful_disconnect(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This enabled graceful disconnects on close. A graceful disconnect
means that we signal the receiving end that the connection is terminated
and wait for it to close the connection before closing the connection.

A graceful disconnect means that we can be sure that we successfully sent
all the outstanding data to the other end, or get an error reported.
However, it also means we have to wait for all the data to reach the
other side and for it to acknowledge this by closing the socket, which may
take a while. For this reason it is disabled by default.
*/
func (self *TraitTcpConnection) SetGracefulDisconnect(graceful_disconnect bool) {
	__cgo__graceful_disconnect := C.gboolean(0)
	if graceful_disconnect {
		__cgo__graceful_disconnect = C.gboolean(1)
	}
	C.g_tcp_connection_set_graceful_disconnect(self.CPointer, __cgo__graceful_disconnect)
	return
}

type TraitTcpWrapperConnection struct{ CPointer *C.GTcpWrapperConnection }
type IsTcpWrapperConnection interface {
	GetTcpWrapperConnectionPointer() *C.GTcpWrapperConnection
}

func (self *TraitTcpWrapperConnection) GetTcpWrapperConnectionPointer() *C.GTcpWrapperConnection {
	return self.CPointer
}
func NewTraitTcpWrapperConnection(p unsafe.Pointer) *TraitTcpWrapperConnection {
	return &TraitTcpWrapperConnection{(*C.GTcpWrapperConnection)(p)}
}

/*
Get's @conn's base #GIOStream
*/
func (self *TraitTcpWrapperConnection) GetBaseIoStream() (return__ *IOStream) {
	var __cgo__return__ *C.GIOStream
	__cgo__return__ = C.g_tcp_wrapper_connection_get_base_io_stream(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewIOStreamFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

type TraitTestDBus struct{ CPointer *C.GTestDBus }
type IsTestDBus interface {
	GetTestDBusPointer() *C.GTestDBus
}

func (self *TraitTestDBus) GetTestDBusPointer() *C.GTestDBus {
	return self.CPointer
}
func NewTraitTestDBus(p unsafe.Pointer) *TraitTestDBus {
	return &TraitTestDBus{(*C.GTestDBus)(p)}
}

/*
Add a path where dbus-daemon will look up .service files. This can't be
called after g_test_dbus_up().
*/
func (self *TraitTestDBus) AddServiceDir(path string) {
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	C.g_test_dbus_add_service_dir(self.CPointer, __cgo__path)
	C.free(unsafe.Pointer(__cgo__path))
	return
}

/*
Stop the session bus started by g_test_dbus_up().

This will wait for the singleton returned by g_bus_get() or g_bus_get_sync()
is destroyed. This is done to ensure that the next unit test won't get a
leaked singleton from this test.
*/
func (self *TraitTestDBus) Down() {
	C.g_test_dbus_down(self.CPointer)
	return
}

/*
Get the address on which dbus-daemon is running. If g_test_dbus_up() has not
been called yet, %NULL is returned. This can be used with
g_dbus_connection_new_for_address().
*/
func (self *TraitTestDBus) GetBusAddress() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_test_dbus_get_bus_address(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Get the flags of the #GTestDBus object.
*/
func (self *TraitTestDBus) GetFlags() (return__ C.GTestDBusFlags) {
	return__ = C.g_test_dbus_get_flags(self.CPointer)
	return
}

/*
Stop the session bus started by g_test_dbus_up().

Unlike g_test_dbus_down(), this won't verify the #GDBusConnection
singleton returned by g_bus_get() or g_bus_get_sync() is destroyed. Unit
tests wanting to verify behaviour after the session bus has been stopped
can use this function but should still call g_test_dbus_down() when done.
*/
func (self *TraitTestDBus) Stop() {
	C.g_test_dbus_stop(self.CPointer)
	return
}

/*
Start a dbus-daemon instance and set DBUS_SESSION_BUS_ADDRESS. After this
call, it is safe for unit tests to start sending messages on the session bus.

If this function is called from setup callback of g_test_add(),
g_test_dbus_down() must be called in its teardown callback.

If this function is called from unit test's main(), then g_test_dbus_down()
must be called after g_test_run().
*/
func (self *TraitTestDBus) Up() {
	C.g_test_dbus_up(self.CPointer)
	return
}

type TraitThemedIcon struct{ CPointer *C.GThemedIcon }
type IsThemedIcon interface {
	GetThemedIconPointer() *C.GThemedIcon
}

func (self *TraitThemedIcon) GetThemedIconPointer() *C.GThemedIcon {
	return self.CPointer
}
func NewTraitThemedIcon(p unsafe.Pointer) *TraitThemedIcon {
	return &TraitThemedIcon{(*C.GThemedIcon)(p)}
}

/*
Append a name to the list of icons from within @icon.

Note that doing so invalidates the hash computed by prior calls
to g_icon_hash().
*/
func (self *TraitThemedIcon) AppendName(iconname string) {
	__cgo__iconname := C.CString(iconname)
	C.g_themed_icon_append_name(self.CPointer, __cgo__iconname)
	C.free(unsafe.Pointer(__cgo__iconname))
	return
}

/*
Gets the names of icons from within @icon.
*/
func (self *TraitThemedIcon) GetNames() (return__ **C.gchar) {
	return__ = C.g_themed_icon_get_names(self.CPointer)
	return
}

/*
Prepend a name to the list of icons from within @icon.

Note that doing so invalidates the hash computed by prior calls
to g_icon_hash().
*/
func (self *TraitThemedIcon) PrependName(iconname string) {
	__cgo__iconname := C.CString(iconname)
	C.g_themed_icon_prepend_name(self.CPointer, __cgo__iconname)
	C.free(unsafe.Pointer(__cgo__iconname))
	return
}

type TraitThreadedSocketService struct{ CPointer *C.GThreadedSocketService }
type IsThreadedSocketService interface {
	GetThreadedSocketServicePointer() *C.GThreadedSocketService
}

func (self *TraitThreadedSocketService) GetThreadedSocketServicePointer() *C.GThreadedSocketService {
	return self.CPointer
}
func NewTraitThreadedSocketService(p unsafe.Pointer) *TraitThreadedSocketService {
	return &TraitThreadedSocketService{(*C.GThreadedSocketService)(p)}
}

type TraitTlsCertificate struct{ CPointer *C.GTlsCertificate }
type IsTlsCertificate interface {
	GetTlsCertificatePointer() *C.GTlsCertificate
}

func (self *TraitTlsCertificate) GetTlsCertificatePointer() *C.GTlsCertificate {
	return self.CPointer
}
func NewTraitTlsCertificate(p unsafe.Pointer) *TraitTlsCertificate {
	return &TraitTlsCertificate{(*C.GTlsCertificate)(p)}
}

/*
Gets the #GTlsCertificate representing @cert's issuer, if known
*/
func (self *TraitTlsCertificate) GetIssuer() (return__ *TlsCertificate) {
	var __cgo__return__ *C.GTlsCertificate
	__cgo__return__ = C.g_tls_certificate_get_issuer(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewTlsCertificateFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Check if two #GTlsCertificate objects represent the same certificate.
The raw DER byte data of the two certificates are checked for equality.
This has the effect that two certificates may compare equal even if
their #GTlsCertificate:issuer, #GTlsCertificate:private-key, or
#GTlsCertificate:private-key-pem properties differ.
*/
func (self *TraitTlsCertificate) IsSame(cert_two IsTlsCertificate) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_tls_certificate_is_same(self.CPointer, cert_two.GetTlsCertificatePointer())
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This verifies @cert and returns a set of #GTlsCertificateFlags
indicating any problems found with it. This can be used to verify a
certificate outside the context of making a connection, or to
check a certificate against a CA that is not part of the system
CA database.

If @identity is not %NULL, @cert's name(s) will be compared against
it, and %G_TLS_CERTIFICATE_BAD_IDENTITY will be set in the return
value if it does not match. If @identity is %NULL, that bit will
never be set in the return value.

If @trusted_ca is not %NULL, then @cert (or one of the certificates
in its chain) must be signed by it, or else
%G_TLS_CERTIFICATE_UNKNOWN_CA will be set in the return value. If
@trusted_ca is %NULL, that bit will never be set in the return
value.

(All other #GTlsCertificateFlags values will always be set or unset
as appropriate.)
*/
func (self *TraitTlsCertificate) Verify(identity *C.GSocketConnectable, trusted_ca IsTlsCertificate) (return__ C.GTlsCertificateFlags) {
	return__ = C.g_tls_certificate_verify(self.CPointer, identity, trusted_ca.GetTlsCertificatePointer())
	return
}

type TraitTlsConnection struct{ CPointer *C.GTlsConnection }
type IsTlsConnection interface {
	GetTlsConnectionPointer() *C.GTlsConnection
}

func (self *TraitTlsConnection) GetTlsConnectionPointer() *C.GTlsConnection {
	return self.CPointer
}
func NewTraitTlsConnection(p unsafe.Pointer) *TraitTlsConnection {
	return &TraitTlsConnection{(*C.GTlsConnection)(p)}
}

/*
Used by #GTlsConnection implementations to emit the
#GTlsConnection::accept-certificate signal.
*/
func (self *TraitTlsConnection) EmitAcceptCertificate(peer_cert IsTlsCertificate, errors C.GTlsCertificateFlags) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_tls_connection_emit_accept_certificate(self.CPointer, peer_cert.GetTlsCertificatePointer(), errors)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets @conn's certificate, as set by
g_tls_connection_set_certificate().
*/
func (self *TraitTlsConnection) GetCertificate() (return__ *TlsCertificate) {
	var __cgo__return__ *C.GTlsCertificate
	__cgo__return__ = C.g_tls_connection_get_certificate(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewTlsCertificateFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the certificate database that @conn uses to verify
peer certificates. See g_tls_connection_set_database().
*/
func (self *TraitTlsConnection) GetDatabase() (return__ *TlsDatabase) {
	var __cgo__return__ *C.GTlsDatabase
	__cgo__return__ = C.g_tls_connection_get_database(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewTlsDatabaseFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Get the object that will be used to interact with the user. It will be used
for things like prompting the user for passwords. If %NULL is returned, then
no user interaction will occur for this connection.
*/
func (self *TraitTlsConnection) GetInteraction() (return__ *TlsInteraction) {
	var __cgo__return__ *C.GTlsInteraction
	__cgo__return__ = C.g_tls_connection_get_interaction(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewTlsInteractionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets @conn's peer's certificate after the handshake has completed.
(It is not set during the emission of
#GTlsConnection::accept-certificate.)
*/
func (self *TraitTlsConnection) GetPeerCertificate() (return__ *TlsCertificate) {
	var __cgo__return__ *C.GTlsCertificate
	__cgo__return__ = C.g_tls_connection_get_peer_certificate(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewTlsCertificateFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the errors associated with validating @conn's peer's
certificate, after the handshake has completed. (It is not set
during the emission of #GTlsConnection::accept-certificate.)
*/
func (self *TraitTlsConnection) GetPeerCertificateErrors() (return__ C.GTlsCertificateFlags) {
	return__ = C.g_tls_connection_get_peer_certificate_errors(self.CPointer)
	return
}

/*
Gets @conn rehandshaking mode. See
g_tls_connection_set_rehandshake_mode() for details.
*/
func (self *TraitTlsConnection) GetRehandshakeMode() (return__ C.GTlsRehandshakeMode) {
	return__ = C.g_tls_connection_get_rehandshake_mode(self.CPointer)
	return
}

/*
Tests whether or not @conn expects a proper TLS close notification
when the connection is closed. See
g_tls_connection_set_require_close_notify() for details.
*/
func (self *TraitTlsConnection) GetRequireCloseNotify() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_tls_connection_get_require_close_notify(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// g_tls_connection_get_use_system_certdb is not generated due to deprecation attr

/*
Attempts a TLS handshake on @conn.

On the client side, it is never necessary to call this method;
although the connection needs to perform a handshake after
connecting (or after sending a "STARTTLS"-type command) and may
need to rehandshake later if the server requests it,
#GTlsConnection will handle this for you automatically when you try
to send or receive data on the connection. However, you can call
g_tls_connection_handshake() manually if you want to know for sure
whether the initial handshake succeeded or failed (as opposed to
just immediately trying to write to @conn's output stream, in which
case if it fails, it may not be possible to tell if it failed
before or after completing the handshake).

Likewise, on the server side, although a handshake is necessary at
the beginning of the communication, you do not need to call this
function explicitly unless you want clearer error reporting.
However, you may call g_tls_connection_handshake() later on to
renegotiate parameters (encryption methods, etc) with the client.

#GTlsConnection::accept_certificate may be emitted during the
handshake.
*/
func (self *TraitTlsConnection) Handshake(cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_tls_connection_handshake(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously performs a TLS handshake on @conn. See
g_tls_connection_handshake() for more information.
*/
func (self *TraitTlsConnection) HandshakeAsync(io_priority int, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_tls_connection_handshake_async(self.CPointer, C.int(io_priority), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finish an asynchronous TLS handshake operation. See
g_tls_connection_handshake() for more information.
*/
func (self *TraitTlsConnection) HandshakeFinish(result *C.GAsyncResult) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_tls_connection_handshake_finish(self.CPointer, result, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This sets the certificate that @conn will present to its peer
during the TLS handshake. For a #GTlsServerConnection, it is
mandatory to set this, and that will normally be done at construct
time.

For a #GTlsClientConnection, this is optional. If a handshake fails
with %G_TLS_ERROR_CERTIFICATE_REQUIRED, that means that the server
requires a certificate, and if you try connecting again, you should
call this method first. You can call
g_tls_client_connection_get_accepted_cas() on the failed connection
to get a list of Certificate Authorities that the server will
accept certificates from.

(It is also possible that a server will allow the connection with
or without a certificate; in that case, if you don't provide a
certificate, you can tell that the server requested one by the fact
that g_tls_client_connection_get_accepted_cas() will return
non-%NULL.)
*/
func (self *TraitTlsConnection) SetCertificate(certificate IsTlsCertificate) {
	C.g_tls_connection_set_certificate(self.CPointer, certificate.GetTlsCertificatePointer())
	return
}

/*
Sets the certificate database that is used to verify peer certificates.
This is set to the default database by default. See
g_tls_backend_get_default_database(). If set to %NULL, then
peer certificate validation will always set the
%G_TLS_CERTIFICATE_UNKNOWN_CA error (meaning
#GTlsConnection::accept-certificate will always be emitted on
client-side connections, unless that bit is not set in
#GTlsClientConnection:validation-flags).
*/
func (self *TraitTlsConnection) SetDatabase(database IsTlsDatabase) {
	C.g_tls_connection_set_database(self.CPointer, database.GetTlsDatabasePointer())
	return
}

/*
Set the object that will be used to interact with the user. It will be used
for things like prompting the user for passwords.

The @interaction argument will normally be a derived subclass of
#GTlsInteraction. %NULL can also be provided if no user interaction
should occur for this connection.
*/
func (self *TraitTlsConnection) SetInteraction(interaction IsTlsInteraction) {
	C.g_tls_connection_set_interaction(self.CPointer, interaction.GetTlsInteractionPointer())
	return
}

/*
Sets how @conn behaves with respect to rehandshaking requests.

%G_TLS_REHANDSHAKE_NEVER means that it will never agree to
rehandshake after the initial handshake is complete. (For a client,
this means it will refuse rehandshake requests from the server, and
for a server, this means it will close the connection with an error
if the client attempts to rehandshake.)

%G_TLS_REHANDSHAKE_SAFELY means that the connection will allow a
rehandshake only if the other end of the connection supports the
TLS `renegotiation_info` extension. This is the default behavior,
but means that rehandshaking will not work against older
implementations that do not support that extension.

%G_TLS_REHANDSHAKE_UNSAFELY means that the connection will allow
rehandshaking even without the `renegotiation_info` extension. On
the server side in particular, this is not recommended, since it
leaves the server open to certain attacks. However, this mode is
necessary if you need to allow renegotiation with older client
software.
*/
func (self *TraitTlsConnection) SetRehandshakeMode(mode C.GTlsRehandshakeMode) {
	C.g_tls_connection_set_rehandshake_mode(self.CPointer, mode)
	return
}

/*
Sets whether or not @conn expects a proper TLS close notification
before the connection is closed. If this is %TRUE (the default),
then @conn will expect to receive a TLS close notification from its
peer before the connection is closed, and will return a
%G_TLS_ERROR_EOF error if the connection is closed without proper
notification (since this may indicate a network error, or
man-in-the-middle attack).

In some protocols, the application will know whether or not the
connection was closed cleanly based on application-level data
(because the application-level data includes a length field, or is
somehow self-delimiting); in this case, the close notify is
redundant and sometimes omitted. (TLS 1.1 explicitly allows this;
in TLS 1.0 it is technically an error, but often done anyway.) You
can use g_tls_connection_set_require_close_notify() to tell @conn
to allow an "unannounced" connection close, in which case the close
will show up as a 0-length read, as in a non-TLS
#GSocketConnection, and it is up to the application to check that
the data has been fully received.

Note that this only affects the behavior when the peer closes the
connection; when the application calls g_io_stream_close() itself
on @conn, this will send a close notification regardless of the
setting of this property. If you explicitly want to do an unclean
close, you can close @conn's #GTlsConnection:base-io-stream rather
than closing @conn itself.
*/
func (self *TraitTlsConnection) SetRequireCloseNotify(require_close_notify bool) {
	__cgo__require_close_notify := C.gboolean(0)
	if require_close_notify {
		__cgo__require_close_notify = C.gboolean(1)
	}
	C.g_tls_connection_set_require_close_notify(self.CPointer, __cgo__require_close_notify)
	return
}

// g_tls_connection_set_use_system_certdb is not generated due to deprecation attr

type TraitTlsDatabase struct{ CPointer *C.GTlsDatabase }
type IsTlsDatabase interface {
	GetTlsDatabasePointer() *C.GTlsDatabase
}

func (self *TraitTlsDatabase) GetTlsDatabasePointer() *C.GTlsDatabase {
	return self.CPointer
}
func NewTraitTlsDatabase(p unsafe.Pointer) *TraitTlsDatabase {
	return &TraitTlsDatabase{(*C.GTlsDatabase)(p)}
}

/*
Create a handle string for the certificate. The database will only be able
to create a handle for certificates that originate from the database. In
cases where the database cannot create a handle for a certificate, %NULL
will be returned.

This handle should be stable across various instances of the application,
and between applications. If a certificate is modified in the database,
then it is not guaranteed that this handle will continue to point to it.
*/
func (self *TraitTlsDatabase) CreateCertificateHandle(certificate IsTlsCertificate) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_tls_database_create_certificate_handle(self.CPointer, certificate.GetTlsCertificatePointer())
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Lookup a certificate by its handle.

The handle should have been created by calling
g_tls_database_create_certificate_handle() on a #GTlsDatabase object of
the same TLS backend. The handle is designed to remain valid across
instantiations of the database.

If the handle is no longer valid, or does not point to a certificate in
this database, then %NULL will be returned.

This function can block, use g_tls_database_lookup_certificate_for_handle_async() to perform
the lookup operation asynchronously.
*/
func (self *TraitTlsDatabase) LookupCertificateForHandle(handle string, interaction IsTlsInteraction, flags C.GTlsDatabaseLookupFlags, cancellable IsCancellable) (return__ *TlsCertificate, __err__ error) {
	__cgo__handle := (*C.gchar)(unsafe.Pointer(C.CString(handle)))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GTlsCertificate
	__cgo__return__ = C.g_tls_database_lookup_certificate_for_handle(self.CPointer, __cgo__handle, interaction.GetTlsInteractionPointer(), flags, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__handle))
	if __cgo__return__ != nil {
		return__ = NewTlsCertificateFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously lookup a certificate by its handle in the database. See
g_tls_database_lookup_certificate_for_handle() for more information.
*/
func (self *TraitTlsDatabase) LookupCertificateForHandleAsync(handle string, interaction IsTlsInteraction, flags C.GTlsDatabaseLookupFlags, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__handle := (*C.gchar)(unsafe.Pointer(C.CString(handle)))
	C.g_tls_database_lookup_certificate_for_handle_async(self.CPointer, __cgo__handle, interaction.GetTlsInteractionPointer(), flags, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__handle))
	return
}

/*
Finish an asynchronous lookup of a certificate by its handle. See
g_tls_database_lookup_certificate_handle() for more information.

If the handle is no longer valid, or does not point to a certificate in
this database, then %NULL will be returned.
*/
func (self *TraitTlsDatabase) LookupCertificateForHandleFinish(result *C.GAsyncResult) (return__ *TlsCertificate, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GTlsCertificate
	__cgo__return__ = C.g_tls_database_lookup_certificate_for_handle_finish(self.CPointer, result, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewTlsCertificateFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Lookup the issuer of @certificate in the database.

The %issuer property
of @certificate is not modified, and the two certificates are not hooked
into a chain.

This function can block, use g_tls_database_lookup_certificate_issuer_async() to perform
the lookup operation asynchronously.
*/
func (self *TraitTlsDatabase) LookupCertificateIssuer(certificate IsTlsCertificate, interaction IsTlsInteraction, flags C.GTlsDatabaseLookupFlags, cancellable IsCancellable) (return__ *TlsCertificate, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GTlsCertificate
	__cgo__return__ = C.g_tls_database_lookup_certificate_issuer(self.CPointer, certificate.GetTlsCertificatePointer(), interaction.GetTlsInteractionPointer(), flags, cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewTlsCertificateFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously lookup the issuer of @certificate in the database. See
g_tls_database_lookup_certificate_issuer() for more information.
*/
func (self *TraitTlsDatabase) LookupCertificateIssuerAsync(certificate IsTlsCertificate, interaction IsTlsInteraction, flags C.GTlsDatabaseLookupFlags, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_tls_database_lookup_certificate_issuer_async(self.CPointer, certificate.GetTlsCertificatePointer(), interaction.GetTlsInteractionPointer(), flags, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finish an asynchronous lookup issuer operation. See
g_tls_database_lookup_certificate_issuer() for more information.
*/
func (self *TraitTlsDatabase) LookupCertificateIssuerFinish(result *C.GAsyncResult) (return__ *TlsCertificate, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GTlsCertificate
	__cgo__return__ = C.g_tls_database_lookup_certificate_issuer_finish(self.CPointer, result, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewTlsCertificateFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Lookup certificates issued by this issuer in the database.

This function can block, use g_tls_database_lookup_certificates_issued_by_async() to perform
the lookup operation asynchronously.
*/
func (self *TraitTlsDatabase) LookupCertificatesIssuedBy(issuer_raw_dn *C.GByteArray, interaction IsTlsInteraction, flags C.GTlsDatabaseLookupFlags, cancellable IsCancellable) (return__ *C.GList, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_tls_database_lookup_certificates_issued_by(self.CPointer, issuer_raw_dn, interaction.GetTlsInteractionPointer(), flags, cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously lookup certificates issued by this issuer in the database. See
g_tls_database_lookup_certificates_issued_by() for more information.

The database may choose to hold a reference to the issuer byte array for the duration
of of this asynchronous operation. The byte array should not be modified during
this time.
*/
func (self *TraitTlsDatabase) LookupCertificatesIssuedByAsync(issuer_raw_dn *C.GByteArray, interaction IsTlsInteraction, flags C.GTlsDatabaseLookupFlags, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_tls_database_lookup_certificates_issued_by_async(self.CPointer, issuer_raw_dn, interaction.GetTlsInteractionPointer(), flags, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finish an asynchronous lookup of certificates. See
g_tls_database_lookup_certificates_issued_by() for more information.
*/
func (self *TraitTlsDatabase) LookupCertificatesIssuedByFinish(result *C.GAsyncResult) (return__ *C.GList, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_tls_database_lookup_certificates_issued_by_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Verify's a certificate chain after looking up and adding any missing
certificates to the chain.

@chain is a chain of #GTlsCertificate objects each pointing to the next
certificate in the chain by its %issuer property. The chain may initially
consist of one or more certificates. After the verification process is
complete, @chain may be modified by adding missing certificates, or removing
extra certificates. If a certificate anchor was found, then it is added to
the @chain.

@purpose describes the purpose (or usage) for which the certificate
is being used. Typically @purpose will be set to #G_TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER
which means that the certificate is being used to authenticate a server
(and we are acting as the client).

The @identity is used to check for pinned certificates (trust exceptions)
in the database. These will override the normal verification process on a
host by host basis.

Currently there are no @flags, and %G_TLS_DATABASE_VERIFY_NONE should be
used.

This function can block, use g_tls_database_verify_chain_async() to perform
the verification operation asynchronously.
*/
func (self *TraitTlsDatabase) VerifyChain(chain IsTlsCertificate, purpose string, identity *C.GSocketConnectable, interaction IsTlsInteraction, flags C.GTlsDatabaseVerifyFlags, cancellable IsCancellable) (return__ C.GTlsCertificateFlags, __err__ error) {
	__cgo__purpose := (*C.gchar)(unsafe.Pointer(C.CString(purpose)))
	var __cgo_error__ *C.GError
	return__ = C.g_tls_database_verify_chain(self.CPointer, chain.GetTlsCertificatePointer(), __cgo__purpose, identity, interaction.GetTlsInteractionPointer(), flags, cancellable.GetCancellablePointer(), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__purpose))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously verify's a certificate chain after looking up and adding
any missing certificates to the chain. See g_tls_database_verify_chain()
for more information.
*/
func (self *TraitTlsDatabase) VerifyChainAsync(chain IsTlsCertificate, purpose string, identity *C.GSocketConnectable, interaction IsTlsInteraction, flags C.GTlsDatabaseVerifyFlags, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__purpose := (*C.gchar)(unsafe.Pointer(C.CString(purpose)))
	C.g_tls_database_verify_chain_async(self.CPointer, chain.GetTlsCertificatePointer(), __cgo__purpose, identity, interaction.GetTlsInteractionPointer(), flags, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__purpose))
	return
}

/*
Finish an asynchronous verify chain operation. See
g_tls_database_verify_chain() for more information. *
*/
func (self *TraitTlsDatabase) VerifyChainFinish(result *C.GAsyncResult) (return__ C.GTlsCertificateFlags, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_tls_database_verify_chain_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

type TraitTlsInteraction struct{ CPointer *C.GTlsInteraction }
type IsTlsInteraction interface {
	GetTlsInteractionPointer() *C.GTlsInteraction
}

func (self *TraitTlsInteraction) GetTlsInteractionPointer() *C.GTlsInteraction {
	return self.CPointer
}
func NewTraitTlsInteraction(p unsafe.Pointer) *TraitTlsInteraction {
	return &TraitTlsInteraction{(*C.GTlsInteraction)(p)}
}

/*
Run synchronous interaction to ask the user for a password. In general,
g_tls_interaction_invoke_ask_password() should be used instead of this
function.

Derived subclasses usually implement a password prompt, although they may
also choose to provide a password from elsewhere. The @password value will
be filled in and then @callback will be called. Alternatively the user may
abort this password request, which will usually abort the TLS connection.

If the interaction is cancelled by the cancellation object, or by the
user then %G_TLS_INTERACTION_FAILED will be returned with an error that
contains a %G_IO_ERROR_CANCELLED error code. Certain implementations may
not support immediate cancellation.
*/
func (self *TraitTlsInteraction) AskPassword(password IsTlsPassword, cancellable IsCancellable) (return__ C.GTlsInteractionResult, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_tls_interaction_ask_password(self.CPointer, password.GetTlsPasswordPointer(), cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Run asynchronous interaction to ask the user for a password. In general,
g_tls_interaction_invoke_ask_password() should be used instead of this
function.

Derived subclasses usually implement a password prompt, although they may
also choose to provide a password from elsewhere. The @password value will
be filled in and then @callback will be called. Alternatively the user may
abort this password request, which will usually abort the TLS connection.

If the interaction is cancelled by the cancellation object, or by the
user then %G_TLS_INTERACTION_FAILED will be returned with an error that
contains a %G_IO_ERROR_CANCELLED error code. Certain implementations may
not support immediate cancellation.

Certain implementations may not support immediate cancellation.
*/
func (self *TraitTlsInteraction) AskPasswordAsync(password IsTlsPassword, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_tls_interaction_ask_password_async(self.CPointer, password.GetTlsPasswordPointer(), cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Complete an ask password user interaction request. This should be once
the g_tls_interaction_ask_password_async() completion callback is called.

If %G_TLS_INTERACTION_HANDLED is returned, then the #GTlsPassword passed
to g_tls_interaction_ask_password() will have its password filled in.

If the interaction is cancelled by the cancellation object, or by the
user then %G_TLS_INTERACTION_FAILED will be returned with an error that
contains a %G_IO_ERROR_CANCELLED error code.
*/
func (self *TraitTlsInteraction) AskPasswordFinish(result *C.GAsyncResult) (return__ C.GTlsInteractionResult, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_tls_interaction_ask_password_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Invoke the interaction to ask the user for a password. It invokes this
interaction in the main loop, specifically the #GMainContext returned by
g_main_context_get_thread_default() when the interaction is created. This
is called by called by #GTlsConnection or #GTlsDatabase to ask the user
for a password.

Derived subclasses usually implement a password prompt, although they may
also choose to provide a password from elsewhere. The @password value will
be filled in and then @callback will be called. Alternatively the user may
abort this password request, which will usually abort the TLS connection.

The implementation can either be a synchronous (eg: modal dialog) or an
asynchronous one (eg: modeless dialog). This function will take care of
calling which ever one correctly.

If the interaction is cancelled by the cancellation object, or by the
user then %G_TLS_INTERACTION_FAILED will be returned with an error that
contains a %G_IO_ERROR_CANCELLED error code. Certain implementations may
not support immediate cancellation.
*/
func (self *TraitTlsInteraction) InvokeAskPassword(password IsTlsPassword, cancellable IsCancellable) (return__ C.GTlsInteractionResult, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_tls_interaction_invoke_ask_password(self.CPointer, password.GetTlsPasswordPointer(), cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Invoke the interaction to ask the user to choose a certificate to
use with the connection. It invokes this interaction in the main
loop, specifically the #GMainContext returned by
g_main_context_get_thread_default() when the interaction is
created. This is called by called by #GTlsConnection when the peer
requests a certificate during the handshake.

Derived subclasses usually implement a certificate selector,
although they may also choose to provide a certificate from
elsewhere. Alternatively the user may abort this certificate
request, which may or may not abort the TLS connection.

The implementation can either be a synchronous (eg: modal dialog) or an
asynchronous one (eg: modeless dialog). This function will take care of
calling which ever one correctly.

If the interaction is cancelled by the cancellation object, or by the
user then %G_TLS_INTERACTION_FAILED will be returned with an error that
contains a %G_IO_ERROR_CANCELLED error code. Certain implementations may
not support immediate cancellation.
*/
func (self *TraitTlsInteraction) InvokeRequestCertificate(connection IsTlsConnection, flags C.GTlsCertificateRequestFlags, cancellable IsCancellable) (return__ C.GTlsInteractionResult, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_tls_interaction_invoke_request_certificate(self.CPointer, connection.GetTlsConnectionPointer(), flags, cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Run synchronous interaction to ask the user to choose a certificate to use
with the connection. In general, g_tls_interaction_invoke_request_certificate()
should be used instead of this function.

Derived subclasses usually implement a certificate selector, although they may
also choose to provide a certificate from elsewhere. Alternatively the user may
abort this certificate request, which will usually abort the TLS connection.

If %G_TLS_INTERACTION_HANDLED is returned, then the #GTlsConnection
passed to g_tls_interaction_request_certificate() will have had its
#GTlsConnection:certificate filled in.

If the interaction is cancelled by the cancellation object, or by the
user then %G_TLS_INTERACTION_FAILED will be returned with an error that
contains a %G_IO_ERROR_CANCELLED error code. Certain implementations may
not support immediate cancellation.
*/
func (self *TraitTlsInteraction) RequestCertificate(connection IsTlsConnection, flags C.GTlsCertificateRequestFlags, cancellable IsCancellable) (return__ C.GTlsInteractionResult, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_tls_interaction_request_certificate(self.CPointer, connection.GetTlsConnectionPointer(), flags, cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Run asynchronous interaction to ask the user for a certificate to use with
the connection. In general, g_tls_interaction_invoke_request_certificate() should
be used instead of this function.

Derived subclasses usually implement a certificate selector, although they may
also choose to provide a certificate from elsewhere. @callback will be called
when the operation completes. Alternatively the user may abort this certificate
request, which will usually abort the TLS connection.
*/
func (self *TraitTlsInteraction) RequestCertificateAsync(connection IsTlsConnection, flags C.GTlsCertificateRequestFlags, cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_tls_interaction_request_certificate_async(self.CPointer, connection.GetTlsConnectionPointer(), flags, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Complete an request certificate user interaction request. This should be once
the g_tls_interaction_request_certificate_async() completion callback is called.

If %G_TLS_INTERACTION_HANDLED is returned, then the #GTlsConnection
passed to g_tls_interaction_request_certificate_async() will have had its
#GTlsConnection:certificate filled in.

If the interaction is cancelled by the cancellation object, or by the
user then %G_TLS_INTERACTION_FAILED will be returned with an error that
contains a %G_IO_ERROR_CANCELLED error code.
*/
func (self *TraitTlsInteraction) RequestCertificateFinish(result *C.GAsyncResult) (return__ C.GTlsInteractionResult, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_tls_interaction_request_certificate_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

type TraitTlsPassword struct{ CPointer *C.GTlsPassword }
type IsTlsPassword interface {
	GetTlsPasswordPointer() *C.GTlsPassword
}

func (self *TraitTlsPassword) GetTlsPasswordPointer() *C.GTlsPassword {
	return self.CPointer
}
func NewTraitTlsPassword(p unsafe.Pointer) *TraitTlsPassword {
	return &TraitTlsPassword{(*C.GTlsPassword)(p)}
}

/*
Get a description string about what the password will be used for.
*/
func (self *TraitTlsPassword) GetDescription() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_tls_password_get_description(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Get flags about the password.
*/
func (self *TraitTlsPassword) GetFlags() (return__ C.GTlsPasswordFlags) {
	return__ = C.g_tls_password_get_flags(self.CPointer)
	return
}

/*
Get the password value. If @length is not %NULL then it will be
filled in with the length of the password value. (Note that the
password value is not nul-terminated, so you can only pass %NULL
for @length in contexts where you know the password will have a
certain fixed length.)
*/
func (self *TraitTlsPassword) GetValue(length *C.gsize) (return__ *C.guchar) {
	return__ = C.g_tls_password_get_value(self.CPointer, length)
	return
}

/*
Get a user readable translated warning. Usually this warning is a
representation of the password flags returned from
g_tls_password_get_flags().
*/
func (self *TraitTlsPassword) GetWarning() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_tls_password_get_warning(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Set a description string about what the password will be used for.
*/
func (self *TraitTlsPassword) SetDescription(description string) {
	__cgo__description := (*C.gchar)(unsafe.Pointer(C.CString(description)))
	C.g_tls_password_set_description(self.CPointer, __cgo__description)
	C.free(unsafe.Pointer(__cgo__description))
	return
}

/*
Set flags about the password.
*/
func (self *TraitTlsPassword) SetFlags(flags C.GTlsPasswordFlags) {
	C.g_tls_password_set_flags(self.CPointer, flags)
	return
}

/*
Set the value for this password. The @value will be copied by the password
object.

Specify the @length, for a non-nul-terminated password. Pass -1 as
@length if using a nul-terminated password, and @length will be
calculated automatically. (Note that the terminating nul is not
considered part of the password in this case.)
*/
func (self *TraitTlsPassword) SetValue(value *C.guchar, length int64) {
	C.g_tls_password_set_value(self.CPointer, value, C.gssize(length))
	return
}

/*
Provide the value for this password.

The @value will be owned by the password object, and later freed using
the @destroy function callback.

Specify the @length, for a non-nul-terminated password. Pass -1 as
@length if using a nul-terminated password, and @length will be
calculated automatically. (Note that the terminating nul is not
considered part of the password in this case.)
*/
func (self *TraitTlsPassword) SetValueFull(value *C.guchar, length int64, destroy C.GDestroyNotify) {
	C.g_tls_password_set_value_full(self.CPointer, value, C.gssize(length), destroy)
	return
}

/*
Set a user readable translated warning. Usually this warning is a
representation of the password flags returned from
g_tls_password_get_flags().
*/
func (self *TraitTlsPassword) SetWarning(warning string) {
	__cgo__warning := (*C.gchar)(unsafe.Pointer(C.CString(warning)))
	C.g_tls_password_set_warning(self.CPointer, __cgo__warning)
	C.free(unsafe.Pointer(__cgo__warning))
	return
}

type TraitUnixConnection struct{ CPointer *C.GUnixConnection }
type IsUnixConnection interface {
	GetUnixConnectionPointer() *C.GUnixConnection
}

func (self *TraitUnixConnection) GetUnixConnectionPointer() *C.GUnixConnection {
	return self.CPointer
}
func NewTraitUnixConnection(p unsafe.Pointer) *TraitUnixConnection {
	return &TraitUnixConnection{(*C.GUnixConnection)(p)}
}

/*
Receives credentials from the sending end of the connection.  The
sending end has to call g_unix_connection_send_credentials() (or
similar) for this to work.

As well as reading the credentials this also reads (and discards) a
single byte from the stream, as this is required for credentials
passing to work on some implementations.

Other ways to exchange credentials with a foreign peer includes the
#GUnixCredentialsMessage type and g_socket_get_credentials() function.
*/
func (self *TraitUnixConnection) ReceiveCredentials(cancellable IsCancellable) (return__ *Credentials, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GCredentials
	__cgo__return__ = C.g_unix_connection_receive_credentials(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewCredentialsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously receive credentials.

For more details, see g_unix_connection_receive_credentials() which is
the synchronous version of this call.

When the operation is finished, @callback will be called. You can then call
g_unix_connection_receive_credentials_finish() to get the result of the operation.
*/
func (self *TraitUnixConnection) ReceiveCredentialsAsync(cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_unix_connection_receive_credentials_async(self.CPointer, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes an asynchronous receive credentials operation started with
g_unix_connection_receive_credentials_async().
*/
func (self *TraitUnixConnection) ReceiveCredentialsFinish(result *C.GAsyncResult) (return__ *Credentials, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GCredentials
	__cgo__return__ = C.g_unix_connection_receive_credentials_finish(self.CPointer, result, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewCredentialsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Receives a file descriptor from the sending end of the connection.
The sending end has to call g_unix_connection_send_fd() for this
to work.

As well as reading the fd this also reads a single byte from the
stream, as this is required for fd passing to work on some
implementations.
*/
func (self *TraitUnixConnection) ReceiveFd(cancellable IsCancellable) (return__ int, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_unix_connection_receive_fd(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = int(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Passes the credentials of the current user the receiving side
of the connection. The receiving end has to call
g_unix_connection_receive_credentials() (or similar) to accept the
credentials.

As well as sending the credentials this also writes a single NUL
byte to the stream, as this is required for credentials passing to
work on some implementations.

Other ways to exchange credentials with a foreign peer includes the
#GUnixCredentialsMessage type and g_socket_get_credentials() function.
*/
func (self *TraitUnixConnection) SendCredentials(cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unix_connection_send_credentials(self.CPointer, cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously send credentials.

For more details, see g_unix_connection_send_credentials() which is
the synchronous version of this call.

When the operation is finished, @callback will be called. You can then call
g_unix_connection_send_credentials_finish() to get the result of the operation.
*/
func (self *TraitUnixConnection) SendCredentialsAsync(cancellable IsCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.g_unix_connection_send_credentials_async(self.CPointer, cancellable.GetCancellablePointer(), callback, (C.gpointer)(user_data))
	return
}

/*
Finishes an asynchronous send credentials operation started with
g_unix_connection_send_credentials_async().
*/
func (self *TraitUnixConnection) SendCredentialsFinish(result *C.GAsyncResult) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unix_connection_send_credentials_finish(self.CPointer, result, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Passes a file descriptor to the receiving side of the
connection. The receiving end has to call g_unix_connection_receive_fd()
to accept the file descriptor.

As well as sending the fd this also writes a single byte to the
stream, as this is required for fd passing to work on some
implementations.
*/
func (self *TraitUnixConnection) SendFd(fd int, cancellable IsCancellable) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unix_connection_send_fd(self.CPointer, C.gint(fd), cancellable.GetCancellablePointer(), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

type TraitUnixCredentialsMessage struct{ CPointer *C.GUnixCredentialsMessage }
type IsUnixCredentialsMessage interface {
	GetUnixCredentialsMessagePointer() *C.GUnixCredentialsMessage
}

func (self *TraitUnixCredentialsMessage) GetUnixCredentialsMessagePointer() *C.GUnixCredentialsMessage {
	return self.CPointer
}
func NewTraitUnixCredentialsMessage(p unsafe.Pointer) *TraitUnixCredentialsMessage {
	return &TraitUnixCredentialsMessage{(*C.GUnixCredentialsMessage)(p)}
}

/*
Gets the credentials stored in @message.
*/
func (self *TraitUnixCredentialsMessage) GetCredentials() (return__ *Credentials) {
	var __cgo__return__ *C.GCredentials
	__cgo__return__ = C.g_unix_credentials_message_get_credentials(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewCredentialsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

type TraitUnixFDList struct{ CPointer *C.GUnixFDList }
type IsUnixFDList interface {
	GetUnixFDListPointer() *C.GUnixFDList
}

func (self *TraitUnixFDList) GetUnixFDListPointer() *C.GUnixFDList {
	return self.CPointer
}
func NewTraitUnixFDList(p unsafe.Pointer) *TraitUnixFDList {
	return &TraitUnixFDList{(*C.GUnixFDList)(p)}
}

/*
Adds a file descriptor to @list.

The file descriptor is duplicated using dup(). You keep your copy
of the descriptor and the copy contained in @list will be closed
when @list is finalized.

A possible cause of failure is exceeding the per-process or
system-wide file descriptor limit.

The index of the file descriptor in the list is returned.  If you use
this index with g_unix_fd_list_get() then you will receive back a
duplicated copy of the same file descriptor.
*/
func (self *TraitUnixFDList) Append(fd int) (return__ int, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_unix_fd_list_append(self.CPointer, C.gint(fd), &__cgo_error__)
	return__ = int(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets a file descriptor out of @list.

@index_ specifies the index of the file descriptor to get.  It is a
programmer error for @index_ to be out of range; see
g_unix_fd_list_get_length().

The file descriptor is duplicated using dup() and set as
close-on-exec before being returned.  You must call close() on it
when you are done.

A possible cause of failure is exceeding the per-process or
system-wide file descriptor limit.
*/
func (self *TraitUnixFDList) Get(index_ int) (return__ int, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_unix_fd_list_get(self.CPointer, C.gint(index_), &__cgo_error__)
	return__ = int(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the length of @list (ie: the number of file descriptors
contained within).
*/
func (self *TraitUnixFDList) GetLength() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_unix_fd_list_get_length(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the array of file descriptors that is contained in this
object.

After this call, the descriptors remain the property of @list.  The
caller must not close them and must not free the array.  The array is
valid only until @list is changed in any way.

If @length is non-%NULL then it is set to the number of file
descriptors in the returned array. The returned array is also
terminated with -1.

This function never returns %NULL. In case there are no file
descriptors contained in @list, an empty array is returned.
*/
func (self *TraitUnixFDList) PeekFds() (length int, return__ []int) {
	var __cgo__length C.gint
	var __cgo__return__ *C.gint
	__cgo__return__ = C.g_unix_fd_list_peek_fds(self.CPointer, &__cgo__length)
	length = int(__cgo__length)
	defer func() {
		__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&return__))
		__header__return__.Len = int(length)
		__header__return__.Cap = int(length)
		__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	}()
	return
}

/*
Returns the array of file descriptors that is contained in this
object.

After this call, the descriptors are no longer contained in
@list. Further calls will return an empty list (unless more
descriptors have been added).

The return result of this function must be freed with g_free().
The caller is also responsible for closing all of the file
descriptors.  The file descriptors in the array are set to
close-on-exec.

If @length is non-%NULL then it is set to the number of file
descriptors in the returned array. The returned array is also
terminated with -1.

This function never returns %NULL. In case there are no file
descriptors contained in @list, an empty array is returned.
*/
func (self *TraitUnixFDList) StealFds() (length int, return__ []int) {
	var __cgo__length C.gint
	var __cgo__return__ *C.gint
	__cgo__return__ = C.g_unix_fd_list_steal_fds(self.CPointer, &__cgo__length)
	length = int(__cgo__length)
	defer func() {
		__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&return__))
		__header__return__.Len = int(length)
		__header__return__.Cap = int(length)
		__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	}()
	return
}

type TraitUnixFDMessage struct{ CPointer *C.GUnixFDMessage }
type IsUnixFDMessage interface {
	GetUnixFDMessagePointer() *C.GUnixFDMessage
}

func (self *TraitUnixFDMessage) GetUnixFDMessagePointer() *C.GUnixFDMessage {
	return self.CPointer
}
func NewTraitUnixFDMessage(p unsafe.Pointer) *TraitUnixFDMessage {
	return &TraitUnixFDMessage{(*C.GUnixFDMessage)(p)}
}

/*
Adds a file descriptor to @message.

The file descriptor is duplicated using dup(). You keep your copy
of the descriptor and the copy contained in @message will be closed
when @message is finalized.

A possible cause of failure is exceeding the per-process or
system-wide file descriptor limit.
*/
func (self *TraitUnixFDMessage) AppendFd(fd int) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unix_fd_message_append_fd(self.CPointer, C.gint(fd), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the #GUnixFDList contained in @message.  This function does not
return a reference to the caller, but the returned list is valid for
the lifetime of @message.
*/
func (self *TraitUnixFDMessage) GetFdList() (return__ *UnixFDList) {
	var __cgo__return__ *C.GUnixFDList
	__cgo__return__ = C.g_unix_fd_message_get_fd_list(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewUnixFDListFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the array of file descriptors that is contained in this
object.

After this call, the descriptors are no longer contained in
@message. Further calls will return an empty list (unless more
descriptors have been added).

The return result of this function must be freed with g_free().
The caller is also responsible for closing all of the file
descriptors.

If @length is non-%NULL then it is set to the number of file
descriptors in the returned array. The returned array is also
terminated with -1.

This function never returns %NULL. In case there are no file
descriptors contained in @message, an empty array is returned.
*/
func (self *TraitUnixFDMessage) StealFds() (length int, return__ []int) {
	var __cgo__length C.gint
	var __cgo__return__ *C.gint
	__cgo__return__ = C.g_unix_fd_message_steal_fds(self.CPointer, &__cgo__length)
	length = int(__cgo__length)
	defer func() {
		__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&return__))
		__header__return__.Len = int(length)
		__header__return__.Cap = int(length)
		__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	}()
	return
}

type TraitUnixInputStream struct{ CPointer *C.GUnixInputStream }
type IsUnixInputStream interface {
	GetUnixInputStreamPointer() *C.GUnixInputStream
}

func (self *TraitUnixInputStream) GetUnixInputStreamPointer() *C.GUnixInputStream {
	return self.CPointer
}
func NewTraitUnixInputStream(p unsafe.Pointer) *TraitUnixInputStream {
	return &TraitUnixInputStream{(*C.GUnixInputStream)(p)}
}

/*
Returns whether the file descriptor of @stream will be
closed when the stream is closed.
*/
func (self *TraitUnixInputStream) GetCloseFd() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unix_input_stream_get_close_fd(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Return the UNIX file descriptor that the stream reads from.
*/
func (self *TraitUnixInputStream) GetFd() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_unix_input_stream_get_fd(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Sets whether the file descriptor of @stream shall be closed
when the stream is closed.
*/
func (self *TraitUnixInputStream) SetCloseFd(close_fd bool) {
	__cgo__close_fd := C.gboolean(0)
	if close_fd {
		__cgo__close_fd = C.gboolean(1)
	}
	C.g_unix_input_stream_set_close_fd(self.CPointer, __cgo__close_fd)
	return
}

type TraitUnixMountMonitor struct{ CPointer *C.GUnixMountMonitor }
type IsUnixMountMonitor interface {
	GetUnixMountMonitorPointer() *C.GUnixMountMonitor
}

func (self *TraitUnixMountMonitor) GetUnixMountMonitorPointer() *C.GUnixMountMonitor {
	return self.CPointer
}
func NewTraitUnixMountMonitor(p unsafe.Pointer) *TraitUnixMountMonitor {
	return &TraitUnixMountMonitor{(*C.GUnixMountMonitor)(p)}
}

/*
Sets the rate limit to which the @mount_monitor will report
consecutive change events to the mount and mount point entry files.
*/
func (self *TraitUnixMountMonitor) SetRateLimit(limit_msec int) {
	C.g_unix_mount_monitor_set_rate_limit(self.CPointer, C.int(limit_msec))
	return
}

type TraitUnixOutputStream struct{ CPointer *C.GUnixOutputStream }
type IsUnixOutputStream interface {
	GetUnixOutputStreamPointer() *C.GUnixOutputStream
}

func (self *TraitUnixOutputStream) GetUnixOutputStreamPointer() *C.GUnixOutputStream {
	return self.CPointer
}
func NewTraitUnixOutputStream(p unsafe.Pointer) *TraitUnixOutputStream {
	return &TraitUnixOutputStream{(*C.GUnixOutputStream)(p)}
}

/*
Returns whether the file descriptor of @stream will be
closed when the stream is closed.
*/
func (self *TraitUnixOutputStream) GetCloseFd() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unix_output_stream_get_close_fd(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Return the UNIX file descriptor that the stream writes to.
*/
func (self *TraitUnixOutputStream) GetFd() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_unix_output_stream_get_fd(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Sets whether the file descriptor of @stream shall be closed
when the stream is closed.
*/
func (self *TraitUnixOutputStream) SetCloseFd(close_fd bool) {
	__cgo__close_fd := C.gboolean(0)
	if close_fd {
		__cgo__close_fd = C.gboolean(1)
	}
	C.g_unix_output_stream_set_close_fd(self.CPointer, __cgo__close_fd)
	return
}

type TraitUnixSocketAddress struct{ CPointer *C.GUnixSocketAddress }
type IsUnixSocketAddress interface {
	GetUnixSocketAddressPointer() *C.GUnixSocketAddress
}

func (self *TraitUnixSocketAddress) GetUnixSocketAddressPointer() *C.GUnixSocketAddress {
	return self.CPointer
}
func NewTraitUnixSocketAddress(p unsafe.Pointer) *TraitUnixSocketAddress {
	return &TraitUnixSocketAddress{(*C.GUnixSocketAddress)(p)}
}

/*
Gets @address's type.
*/
func (self *TraitUnixSocketAddress) GetAddressType() (return__ C.GUnixSocketAddressType) {
	return__ = C.g_unix_socket_address_get_address_type(self.CPointer)
	return
}

// g_unix_socket_address_get_is_abstract is not generated due to deprecation attr

/*
Gets @address's path, or for abstract sockets the "name".

Guaranteed to be zero-terminated, but an abstract socket
may contain embedded zeros, and thus you should use
g_unix_socket_address_get_path_len() to get the true length
of this string.
*/
func (self *TraitUnixSocketAddress) GetPath() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_unix_socket_address_get_path(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Gets the length of @address's path.

For details, see g_unix_socket_address_get_path().
*/
func (self *TraitUnixSocketAddress) GetPathLen() (return__ int64) {
	var __cgo__return__ C.gsize
	__cgo__return__ = C.g_unix_socket_address_get_path_len(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

type TraitVfs struct{ CPointer *C.GVfs }
type IsVfs interface {
	GetVfsPointer() *C.GVfs
}

func (self *TraitVfs) GetVfsPointer() *C.GVfs {
	return self.CPointer
}
func NewTraitVfs(p unsafe.Pointer) *TraitVfs {
	return &TraitVfs{(*C.GVfs)(p)}
}

/*
Gets a #GFile for @path.
*/
func (self *TraitVfs) GetFileForPath(path string) (return__ *C.GFile) {
	__cgo__path := C.CString(path)
	return__ = C.g_vfs_get_file_for_path(self.CPointer, __cgo__path)
	C.free(unsafe.Pointer(__cgo__path))
	return
}

/*
Gets a #GFile for @uri.

This operation never fails, but the returned object
might not support any I/O operation if the URI
is malformed or if the URI scheme is not supported.
*/
func (self *TraitVfs) GetFileForUri(uri string) (return__ *C.GFile) {
	__cgo__uri := C.CString(uri)
	return__ = C.g_vfs_get_file_for_uri(self.CPointer, __cgo__uri)
	C.free(unsafe.Pointer(__cgo__uri))
	return
}

/*
Gets a list of URI schemes supported by @vfs.
*/
func (self *TraitVfs) GetSupportedUriSchemes() (return__ **C.gchar) {
	return__ = C.g_vfs_get_supported_uri_schemes(self.CPointer)
	return
}

/*
Checks if the VFS is active.
*/
func (self *TraitVfs) IsActive() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_vfs_is_active(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This operation never fails, but the returned object might
not support any I/O operations if the @parse_name cannot
be parsed by the #GVfs module.
*/
func (self *TraitVfs) ParseName(parse_name string) (return__ *C.GFile) {
	__cgo__parse_name := C.CString(parse_name)
	return__ = C.g_vfs_parse_name(self.CPointer, __cgo__parse_name)
	C.free(unsafe.Pointer(__cgo__parse_name))
	return
}

type TraitVolumeMonitor struct{ CPointer *C.GVolumeMonitor }
type IsVolumeMonitor interface {
	GetVolumeMonitorPointer() *C.GVolumeMonitor
}

func (self *TraitVolumeMonitor) GetVolumeMonitorPointer() *C.GVolumeMonitor {
	return self.CPointer
}
func NewTraitVolumeMonitor(p unsafe.Pointer) *TraitVolumeMonitor {
	return &TraitVolumeMonitor{(*C.GVolumeMonitor)(p)}
}

/*
Gets a list of drives connected to the system.

The returned list should be freed with g_list_free(), after
its elements have been unreffed with g_object_unref().
*/
func (self *TraitVolumeMonitor) GetConnectedDrives() (return__ *C.GList) {
	return__ = C.g_volume_monitor_get_connected_drives(self.CPointer)
	return
}

/*
Finds a #GMount object by its UUID (see g_mount_get_uuid())
*/
func (self *TraitVolumeMonitor) GetMountForUuid(uuid string) (return__ *C.GMount) {
	__cgo__uuid := C.CString(uuid)
	return__ = C.g_volume_monitor_get_mount_for_uuid(self.CPointer, __cgo__uuid)
	C.free(unsafe.Pointer(__cgo__uuid))
	return
}

/*
Gets a list of the mounts on the system.

The returned list should be freed with g_list_free(), after
its elements have been unreffed with g_object_unref().
*/
func (self *TraitVolumeMonitor) GetMounts() (return__ *C.GList) {
	return__ = C.g_volume_monitor_get_mounts(self.CPointer)
	return
}

/*
Finds a #GVolume object by its UUID (see g_volume_get_uuid())
*/
func (self *TraitVolumeMonitor) GetVolumeForUuid(uuid string) (return__ *C.GVolume) {
	__cgo__uuid := C.CString(uuid)
	return__ = C.g_volume_monitor_get_volume_for_uuid(self.CPointer, __cgo__uuid)
	C.free(unsafe.Pointer(__cgo__uuid))
	return
}

/*
Gets a list of the volumes on the system.

The returned list should be freed with g_list_free(), after
its elements have been unreffed with g_object_unref().
*/
func (self *TraitVolumeMonitor) GetVolumes() (return__ *C.GList) {
	return__ = C.g_volume_monitor_get_volumes(self.CPointer)
	return
}

type TraitZlibCompressor struct{ CPointer *C.GZlibCompressor }
type IsZlibCompressor interface {
	GetZlibCompressorPointer() *C.GZlibCompressor
}

func (self *TraitZlibCompressor) GetZlibCompressorPointer() *C.GZlibCompressor {
	return self.CPointer
}
func NewTraitZlibCompressor(p unsafe.Pointer) *TraitZlibCompressor {
	return &TraitZlibCompressor{(*C.GZlibCompressor)(p)}
}

/*
Returns the #GZlibCompressor:file-info property.
*/
func (self *TraitZlibCompressor) GetFileInfo() (return__ *FileInfo) {
	var __cgo__return__ *C.GFileInfo
	__cgo__return__ = C.g_zlib_compressor_get_file_info(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewFileInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Sets @file_info in @compressor. If non-%NULL, and @compressor's
#GZlibCompressor:format property is %G_ZLIB_COMPRESSOR_FORMAT_GZIP,
it will be used to set the file name and modification time in
the GZIP header of the compressed data.

Note: it is an error to call this function while a compression is in
progress; it may only be called immediately after creation of @compressor,
or after resetting it with g_converter_reset().
*/
func (self *TraitZlibCompressor) SetFileInfo(file_info IsFileInfo) {
	C.g_zlib_compressor_set_file_info(self.CPointer, file_info.GetFileInfoPointer())
	return
}

type TraitZlibDecompressor struct{ CPointer *C.GZlibDecompressor }
type IsZlibDecompressor interface {
	GetZlibDecompressorPointer() *C.GZlibDecompressor
}

func (self *TraitZlibDecompressor) GetZlibDecompressorPointer() *C.GZlibDecompressor {
	return self.CPointer
}
func NewTraitZlibDecompressor(p unsafe.Pointer) *TraitZlibDecompressor {
	return &TraitZlibDecompressor{(*C.GZlibDecompressor)(p)}
}

/*
Retrieves the #GFileInfo constructed from the GZIP header data
of compressed data processed by @compressor, or %NULL if @decompressor's
#GZlibDecompressor:format property is not %G_ZLIB_COMPRESSOR_FORMAT_GZIP,
or the header data was not fully processed yet, or it not present in the
data stream at all.
*/
func (self *TraitZlibDecompressor) GetFileInfo() (return__ *FileInfo) {
	var __cgo__return__ *C.GFileInfo
	__cgo__return__ = C.g_zlib_decompressor_get_file_info(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewFileInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}
