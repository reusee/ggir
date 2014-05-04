package gdk

/*
#include <gdk/gdk.h>
#include <gdk/gdkprivate.h>
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

type TraitAppLaunchContext struct{ CPointer *C.GdkAppLaunchContext }
type IsAppLaunchContext interface {
	GetAppLaunchContextPointer() *C.GdkAppLaunchContext
}

func (self *TraitAppLaunchContext) GetAppLaunchContextPointer() *C.GdkAppLaunchContext {
	return self.CPointer
}
func NewTraitAppLaunchContext(p unsafe.Pointer) *TraitAppLaunchContext {
	return &TraitAppLaunchContext{(*C.GdkAppLaunchContext)(p)}
}

/*
Sets the workspace on which applications will be launched when
using this context when running under a window manager that
supports multiple workspaces, as described in the
[Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec).

When the workspace is not specified or @desktop is set to -1,
it is up to the window manager to pick one, typically it will
be the current workspace.
*/
func (self *TraitAppLaunchContext) SetDesktop(desktop int) {
	C.gdk_app_launch_context_set_desktop(self.CPointer, C.gint(desktop))
	return
}

// gdk_app_launch_context_set_display is not generated due to deprecation attr

/*
Sets the icon for applications that are launched with this
context.

Window Managers can use this information when displaying startup
notification.

See also gdk_app_launch_context_set_icon_name().
*/
func (self *TraitAppLaunchContext) SetIcon(icon *C.GIcon) {
	C.gdk_app_launch_context_set_icon(self.CPointer, icon)
	return
}

/*
Sets the icon for applications that are launched with this context.
The @icon_name will be interpreted in the same way as the Icon field
in desktop files. See also gdk_app_launch_context_set_icon().

If both @icon and @icon_name are set, the @icon_name takes priority.
If neither @icon or @icon_name is set, the icon is taken from either
the file that is passed to launched application or from the #GAppInfo
for the launched application itself.
*/
func (self *TraitAppLaunchContext) SetIconName(icon_name string) {
	__cgo__icon_name := C.CString(icon_name)
	C.gdk_app_launch_context_set_icon_name(self.CPointer, __cgo__icon_name)
	C.free(unsafe.Pointer(__cgo__icon_name))
	return
}

/*
Sets the screen on which applications will be launched when
using this context. See also gdk_app_launch_context_set_display().

If both @screen and @display are set, the @screen takes priority.
If neither @screen or @display are set, the default screen and
display are used.
*/
func (self *TraitAppLaunchContext) SetScreen(screen IsScreen) {
	C.gdk_app_launch_context_set_screen(self.CPointer, screen.GetScreenPointer())
	return
}

/*
Sets the timestamp of @context. The timestamp should ideally
be taken from the event that triggered the launch.

Window managers can use this information to avoid moving the
focus to the newly launched application when the user is busy
typing in another window. This is also known as 'focus stealing
prevention'.
*/
func (self *TraitAppLaunchContext) SetTimestamp(timestamp uint32) {
	C.gdk_app_launch_context_set_timestamp(self.CPointer, C.guint32(timestamp))
	return
}

type TraitCursor struct{ CPointer *C.GdkCursor }
type IsCursor interface {
	GetCursorPointer() *C.GdkCursor
}

func (self *TraitCursor) GetCursorPointer() *C.GdkCursor {
	return self.CPointer
}
func NewTraitCursor(p unsafe.Pointer) *TraitCursor {
	return &TraitCursor{(*C.GdkCursor)(p)}
}

/*
Returns the cursor type for this cursor.
*/
func (self *TraitCursor) GetCursorType() (return__ C.GdkCursorType) {
	return__ = C.gdk_cursor_get_cursor_type(self.CPointer)
	return
}

/*
Returns the display on which the #GdkCursor is defined.
*/
func (self *TraitCursor) GetDisplay() (return__ *Display) {
	var __cgo__return__ *C.GdkDisplay
	__cgo__return__ = C.gdk_cursor_get_display(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDisplayFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns a #GdkPixbuf with the image used to display the cursor.

Note that depending on the capabilities of the windowing system and
on the cursor, GDK may not be able to obtain the image data. In this
case, %NULL is returned.
*/
func (self *TraitCursor) GetImage() (return__ *C.GdkPixbuf) {
	return__ = C.gdk_cursor_get_image(self.CPointer)
	return
}

/*
Returns a cairo image surface with the image used to display the cursor.

Note that depending on the capabilities of the windowing system and
on the cursor, GDK may not be able to obtain the image data. In this
case, %NULL is returned.
*/
func (self *TraitCursor) GetSurface(x_hot *C.gdouble, y_hot *C.gdouble) (return__ *C.cairo_surface_t) {
	return__ = C.gdk_cursor_get_surface(self.CPointer, x_hot, y_hot)
	return
}

// gdk_cursor_ref is not generated due to deprecation attr

// gdk_cursor_unref is not generated due to deprecation attr

type TraitDevice struct{ CPointer *C.GdkDevice }
type IsDevice interface {
	GetDevicePointer() *C.GdkDevice
}

func (self *TraitDevice) GetDevicePointer() *C.GdkDevice {
	return self.CPointer
}
func NewTraitDevice(p unsafe.Pointer) *TraitDevice {
	return &TraitDevice{(*C.GdkDevice)(p)}
}

/*
Returns the associated device to @device, if @device is of type
%GDK_DEVICE_TYPE_MASTER, it will return the paired pointer or
keyboard.

If @device is of type %GDK_DEVICE_TYPE_SLAVE, it will return
the master device to which @device is attached to.

If @device is of type %GDK_DEVICE_TYPE_FLOATING, %NULL will be
returned, as there is no associated device.
*/
func (self *TraitDevice) GetAssociatedDevice() (return__ *Device) {
	var __cgo__return__ *C.GdkDevice
	__cgo__return__ = C.gdk_device_get_associated_device(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDeviceFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Interprets an array of double as axis values for a given device,
and locates the value in the array for a given axis use.
*/
func (self *TraitDevice) GetAxis(axes *C.gdouble, use C.GdkAxisUse) (value float64, return__ bool) {
	var __cgo__value C.gdouble
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_device_get_axis(self.CPointer, axes, use, &__cgo__value)
	value = float64(__cgo__value)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the axis use for @index_.
*/
func (self *TraitDevice) GetAxisUse(index_ uint) (return__ C.GdkAxisUse) {
	return__ = C.gdk_device_get_axis_use(self.CPointer, C.guint(index_))
	return
}

/*
Interprets an array of double as axis values for a given device,
and locates the value in the array for a given axis label, as returned
by gdk_device_list_axes()
*/
func (self *TraitDevice) GetAxisValue(axes *C.gdouble, axis_label C.GdkAtom, value *C.gdouble) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_device_get_axis_value(self.CPointer, axes, axis_label, value)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the device type for @device.
*/
func (self *TraitDevice) GetDeviceType() (return__ C.GdkDeviceType) {
	return__ = C.gdk_device_get_device_type(self.CPointer)
	return
}

/*
Returns the #GdkDisplay to which @device pertains.
*/
func (self *TraitDevice) GetDisplay() (return__ *Display) {
	var __cgo__return__ *C.GdkDisplay
	__cgo__return__ = C.gdk_device_get_display(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDisplayFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Determines whether the pointer follows device motion.
*/
func (self *TraitDevice) GetHasCursor() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_device_get_has_cursor(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Obtains the motion history for a pointer device; given a starting and
ending timestamp, return all events in the motion history for
the device in the given range of time. Some windowing systems
do not support motion history, in which case, %FALSE will
be returned. (This is not distinguishable from the case where
motion history is supported and no events were found.)

Note that there is also gdk_window_set_event_compression() to get
more motion events delivered directly, independent of the windowing
system.
*/
func (self *TraitDevice) GetHistory(window IsWindow, start uint32, stop uint32) (events **C.GdkTimeCoord, n_events int, return__ bool) {
	var __cgo__n_events C.gint
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_device_get_history(self.CPointer, window.GetWindowPointer(), C.guint32(start), C.guint32(stop), &events, &__cgo__n_events)
	n_events = int(__cgo__n_events)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
If @index_ has a valid keyval, this function will return %TRUE
and fill in @keyval and @modifiers with the keyval settings.
*/
func (self *TraitDevice) GetKey(index_ uint) (keyval uint, modifiers C.GdkModifierType, return__ bool) {
	var __cgo__keyval C.guint
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_device_get_key(self.CPointer, C.guint(index_), &__cgo__keyval, &modifiers)
	keyval = uint(__cgo__keyval)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets information about which window the given pointer device is in, based on
that have been received so far from the display server. If another application
has a pointer grab, or this application has a grab with owner_events = %FALSE,
%NULL may be returned even if the pointer is physically over one of this
application's windows.
*/
func (self *TraitDevice) GetLastEventWindow() (return__ *Window) {
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_device_get_last_event_window(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Determines the mode of the device.
*/
func (self *TraitDevice) GetMode() (return__ C.GdkInputMode) {
	return__ = C.gdk_device_get_mode(self.CPointer)
	return
}

/*
Returns the number of axes the device currently has.
*/
func (self *TraitDevice) GetNAxes() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_device_get_n_axes(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the number of keys the device currently has.
*/
func (self *TraitDevice) GetNKeys() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_device_get_n_keys(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Determines the name of the device.
*/
func (self *TraitDevice) GetName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gdk_device_get_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the current location of @device. As a slave device
coordinates are those of its master pointer, This function
may not be called on devices of type %GDK_DEVICE_TYPE_SLAVE,
unless there is an ongoing grab on them, see gdk_device_grab().
*/
func (self *TraitDevice) GetPosition() (screen *Screen, x int, y int) {
	var __cgo__screen *C.GdkScreen
	var __cgo__x C.gint
	var __cgo__y C.gint
	C.gdk_device_get_position(self.CPointer, &__cgo__screen, &__cgo__x, &__cgo__y)
	if __cgo__screen != nil {
		screen = NewScreenFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__screen).Pointer()))
	}
	x = int(__cgo__x)
	y = int(__cgo__y)
	return
}

/*
Gets the current location of @device in double precision. As a slave device's
coordinates are those of its master pointer, this function
may not be called on devices of type %GDK_DEVICE_TYPE_SLAVE,
unless there is an ongoing grab on them. See gdk_device_grab().
*/
func (self *TraitDevice) GetPositionDouble() (screen *Screen, x float64, y float64) {
	var __cgo__screen *C.GdkScreen
	var __cgo__x C.gdouble
	var __cgo__y C.gdouble
	C.gdk_device_get_position_double(self.CPointer, &__cgo__screen, &__cgo__x, &__cgo__y)
	if __cgo__screen != nil {
		screen = NewScreenFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__screen).Pointer()))
	}
	x = float64(__cgo__x)
	y = float64(__cgo__y)
	return
}

/*
Determines the type of the device.
*/
func (self *TraitDevice) GetSource() (return__ C.GdkInputSource) {
	return__ = C.gdk_device_get_source(self.CPointer)
	return
}

/*
Gets the current state of a pointer device relative to @window. As a slave
device’s coordinates are those of its master pointer, this
function may not be called on devices of type %GDK_DEVICE_TYPE_SLAVE,
unless there is an ongoing grab on them. See gdk_device_grab().
*/
func (self *TraitDevice) GetState(window IsWindow, axes *C.gdouble, mask *C.GdkModifierType) {
	C.gdk_device_get_state(self.CPointer, window.GetWindowPointer(), axes, mask)
	return
}

/*
Obtains the window underneath @device, returning the location of the device in @win_x and @win_y. Returns
%NULL if the window tree under @device is not known to GDK (for example, belongs to another application).

As a slave device coordinates are those of its master pointer, This
function may not be called on devices of type %GDK_DEVICE_TYPE_SLAVE,
unless there is an ongoing grab on them, see gdk_device_grab().
*/
func (self *TraitDevice) GetWindowAtPosition() (win_x int, win_y int, return__ *Window) {
	var __cgo__win_x C.gint
	var __cgo__win_y C.gint
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_device_get_window_at_position(self.CPointer, &__cgo__win_x, &__cgo__win_y)
	win_x = int(__cgo__win_x)
	win_y = int(__cgo__win_y)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Obtains the window underneath @device, returning the location of the device in @win_x and @win_y in
double precision. Returns %NULL if the window tree under @device is not known to GDK (for example,
belongs to another application).

As a slave device coordinates are those of its master pointer, This
function may not be called on devices of type %GDK_DEVICE_TYPE_SLAVE,
unless there is an ongoing grab on them, see gdk_device_grab().
*/
func (self *TraitDevice) GetWindowAtPositionDouble() (win_x float64, win_y float64, return__ *Window) {
	var __cgo__win_x C.gdouble
	var __cgo__win_y C.gdouble
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_device_get_window_at_position_double(self.CPointer, &__cgo__win_x, &__cgo__win_y)
	win_x = float64(__cgo__win_x)
	win_y = float64(__cgo__win_y)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Grabs the device so that all events coming from this device are passed to
this application until the device is ungrabbed with gdk_device_ungrab(),
or the window becomes unviewable. This overrides any previous grab on the device
by this client.

Device grabs are used for operations which need complete control over the
given device events (either pointer or keyboard). For example in GTK+ this
is used for Drag and Drop operations, popup menus and such.

Note that if the event mask of an X window has selected both button press
and button release events, then a button press event will cause an automatic
pointer grab until the button is released. X does this automatically since
most applications expect to receive button press and release events in pairs.
It is equivalent to a pointer grab on the window with @owner_events set to
%TRUE.

If you set up anything at the time you take the grab that needs to be
cleaned up when the grab ends, you should handle the #GdkEventGrabBroken
events that are emitted when the grab ends unvoluntarily.
*/
func (self *TraitDevice) Grab(window IsWindow, grab_ownership C.GdkGrabOwnership, owner_events bool, event_mask C.GdkEventMask, cursor IsCursor, time_ uint32) (return__ C.GdkGrabStatus) {
	__cgo__owner_events := C.gboolean(0)
	if owner_events {
		__cgo__owner_events = C.gboolean(1)
	}
	return__ = C.gdk_device_grab(self.CPointer, window.GetWindowPointer(), grab_ownership, __cgo__owner_events, event_mask, cursor.GetCursorPointer(), C.guint32(time_))
	return
}

/*
Returns a #GList of #GdkAtoms, containing the labels for
the axes that @device currently has.
*/
func (self *TraitDevice) ListAxes() (return__ *C.GList) {
	return__ = C.gdk_device_list_axes(self.CPointer)
	return
}

/*
If the device if of type %GDK_DEVICE_TYPE_MASTER, it will return
the list of slave devices attached to it, otherwise it will return
%NULL
*/
func (self *TraitDevice) ListSlaveDevices() (return__ *C.GList) {
	return__ = C.gdk_device_list_slave_devices(self.CPointer)
	return
}

/*
Specifies how an axis of a device is used.
*/
func (self *TraitDevice) SetAxisUse(index_ uint, use C.GdkAxisUse) {
	C.gdk_device_set_axis_use(self.CPointer, C.guint(index_), use)
	return
}

/*
Specifies the X key event to generate when a macro button of a device
is pressed.
*/
func (self *TraitDevice) SetKey(index_ uint, keyval uint, modifiers C.GdkModifierType) {
	C.gdk_device_set_key(self.CPointer, C.guint(index_), C.guint(keyval), modifiers)
	return
}

/*
Sets a the mode of an input device. The mode controls if the
device is active and whether the device’s range is mapped to the
entire screen or to a single window.
*/
func (self *TraitDevice) SetMode(mode C.GdkInputMode) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_device_set_mode(self.CPointer, mode)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Release any grab on @device.
*/
func (self *TraitDevice) Ungrab(time_ uint32) {
	C.gdk_device_ungrab(self.CPointer, C.guint32(time_))
	return
}

/*
Warps @device in @display to the point @x,@y on
the screen @screen, unless the device is confined
to a window by a grab, in which case it will be moved
as far as allowed by the grab. Warping the pointer
creates events as if the user had moved the mouse
instantaneously to the destination.

Note that the pointer should normally be under the
control of the user. This function was added to cover
some rare use cases like keyboard navigation support
for the color picker in the #GtkColorSelectionDialog.
*/
func (self *TraitDevice) Warp(screen IsScreen, x int, y int) {
	C.gdk_device_warp(self.CPointer, screen.GetScreenPointer(), C.gint(x), C.gint(y))
	return
}

type TraitDeviceManager struct{ CPointer *C.GdkDeviceManager }
type IsDeviceManager interface {
	GetDeviceManagerPointer() *C.GdkDeviceManager
}

func (self *TraitDeviceManager) GetDeviceManagerPointer() *C.GdkDeviceManager {
	return self.CPointer
}
func NewTraitDeviceManager(p unsafe.Pointer) *TraitDeviceManager {
	return &TraitDeviceManager{(*C.GdkDeviceManager)(p)}
}

/*
Returns the client pointer, that is, the master pointer that acts as the core pointer
for this application. In X11, window managers may change this depending on the interaction
pattern under the presence of several pointers.

You should use this function seldomly, only in code that isn’t triggered by a #GdkEvent
and there aren’t other means to get a meaningful #GdkDevice to operate on.
*/
func (self *TraitDeviceManager) GetClientPointer() (return__ *Device) {
	var __cgo__return__ *C.GdkDevice
	__cgo__return__ = C.gdk_device_manager_get_client_pointer(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDeviceFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the #GdkDisplay associated to @device_manager.
*/
func (self *TraitDeviceManager) GetDisplay() (return__ *Display) {
	var __cgo__return__ *C.GdkDisplay
	__cgo__return__ = C.gdk_device_manager_get_display(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDisplayFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the list of devices of type @type currently attached to
@device_manager.
*/
func (self *TraitDeviceManager) ListDevices(type_ C.GdkDeviceType) (return__ *C.GList) {
	return__ = C.gdk_device_manager_list_devices(self.CPointer, type_)
	return
}

type TraitDisplay struct{ CPointer *C.GdkDisplay }
type IsDisplay interface {
	GetDisplayPointer() *C.GdkDisplay
}

func (self *TraitDisplay) GetDisplayPointer() *C.GdkDisplay {
	return self.CPointer
}
func NewTraitDisplay(p unsafe.Pointer) *TraitDisplay {
	return &TraitDisplay{(*C.GdkDisplay)(p)}
}

/*
Emits a short beep on @display
*/
func (self *TraitDisplay) Beep() {
	C.gdk_display_beep(self.CPointer)
	return
}

/*
Closes the connection to the windowing system for the given display,
and cleans up associated resources.
*/
func (self *TraitDisplay) Close() {
	C.gdk_display_close(self.CPointer)
	return
}

/*
Returns %TRUE if there is an ongoing grab on @device for @display.
*/
func (self *TraitDisplay) DeviceIsGrabbed(device IsDevice) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_display_device_is_grabbed(self.CPointer, device.GetDevicePointer())
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Flushes any requests queued for the windowing system; this happens automatically
when the main loop blocks waiting for new events, but if your application
is drawing without returning control to the main loop, you may need
to call this function explicitly. A common case where this function
needs to be called is when an application is executing drawing commands
from a thread other than the thread where the main loop is running.

This is most useful for X11. On windowing systems where requests are
handled synchronously, this function will do nothing.
*/
func (self *TraitDisplay) Flush() {
	C.gdk_display_flush(self.CPointer)
	return
}

/*
Returns a #GdkAppLaunchContext suitable for launching
applications on the given display.
*/
func (self *TraitDisplay) GetAppLaunchContext() (return__ *AppLaunchContext) {
	var __cgo__return__ *C.GdkAppLaunchContext
	__cgo__return__ = C.gdk_display_get_app_launch_context(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewAppLaunchContextFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the default size to use for cursors on @display.
*/
func (self *TraitDisplay) GetDefaultCursorSize() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gdk_display_get_default_cursor_size(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Returns the default group leader window for all toplevel windows
on @display. This window is implicitly created by GDK.
See gdk_window_set_group().
*/
func (self *TraitDisplay) GetDefaultGroup() (return__ *Window) {
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_display_get_default_group(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Get the default #GdkScreen for @display.
*/
func (self *TraitDisplay) GetDefaultScreen() (return__ *Screen) {
	var __cgo__return__ *C.GdkScreen
	__cgo__return__ = C.gdk_display_get_default_screen(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewScreenFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the #GdkDeviceManager associated to @display.
*/
func (self *TraitDisplay) GetDeviceManager() (return__ *DeviceManager) {
	var __cgo__return__ *C.GdkDeviceManager
	__cgo__return__ = C.gdk_display_get_device_manager(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDeviceManagerFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the next #GdkEvent to be processed for @display, fetching events from the
windowing system if necessary.
*/
func (self *TraitDisplay) GetEvent() (return__ *C.GdkEvent) {
	return__ = C.gdk_display_get_event(self.CPointer)
	return
}

/*
Gets the maximal size to use for cursors on @display.
*/
func (self *TraitDisplay) GetMaximalCursorSize() (width uint, height uint) {
	var __cgo__width C.guint
	var __cgo__height C.guint
	C.gdk_display_get_maximal_cursor_size(self.CPointer, &__cgo__width, &__cgo__height)
	width = uint(__cgo__width)
	height = uint(__cgo__height)
	return
}

// gdk_display_get_n_screens is not generated due to deprecation attr

/*
Gets the name of the display.
*/
func (self *TraitDisplay) GetName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gdk_display_get_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

// gdk_display_get_pointer is not generated due to deprecation attr

/*
Returns a screen object for one of the screens of the display.
*/
func (self *TraitDisplay) GetScreen(screen_num int) (return__ *Screen) {
	var __cgo__return__ *C.GdkScreen
	__cgo__return__ = C.gdk_display_get_screen(self.CPointer, C.gint(screen_num))
	if __cgo__return__ != nil {
		return__ = NewScreenFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gdk_display_get_window_at_pointer is not generated due to deprecation attr

/*
Returns whether the display has events that are waiting
to be processed.
*/
func (self *TraitDisplay) HasPending() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_display_has_pending(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Finds out if the display has been closed.
*/
func (self *TraitDisplay) IsClosed() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_display_is_closed(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gdk_display_keyboard_ungrab is not generated due to deprecation attr

// gdk_display_list_devices is not generated due to deprecation attr

/*
Indicates to the GUI environment that the application has
finished loading, using a given identifier.

GTK+ will call this function automatically for #GtkWindow
with custom startup-notification identifier unless
gtk_window_set_auto_startup_notification() is called to
disable that feature.
*/
func (self *TraitDisplay) NotifyStartupComplete(startup_id string) {
	__cgo__startup_id := (*C.gchar)(unsafe.Pointer(C.CString(startup_id)))
	C.gdk_display_notify_startup_complete(self.CPointer, __cgo__startup_id)
	C.free(unsafe.Pointer(__cgo__startup_id))
	return
}

/*
Gets a copy of the first #GdkEvent in the @display’s event queue, without
removing the event from the queue.  (Note that this function will
not get more events from the windowing system.  It only checks the events
that have already been moved to the GDK event queue.)
*/
func (self *TraitDisplay) PeekEvent() (return__ *C.GdkEvent) {
	return__ = C.gdk_display_peek_event(self.CPointer)
	return
}

// gdk_display_pointer_is_grabbed is not generated due to deprecation attr

// gdk_display_pointer_ungrab is not generated due to deprecation attr

/*
Appends a copy of the given event onto the front of the event
queue for @display.
*/
func (self *TraitDisplay) PutEvent(event *C.GdkEvent) {
	C.gdk_display_put_event(self.CPointer, event)
	return
}

/*
Request #GdkEventOwnerChange events for ownership changes
of the selection named by the given atom.
*/
func (self *TraitDisplay) RequestSelectionNotification(selection C.GdkAtom) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_display_request_selection_notification(self.CPointer, selection)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the double click distance (two clicks within this distance
count as a double click and result in a #GDK_2BUTTON_PRESS event).
See also gdk_display_set_double_click_time().
Applications should not set this, it is a global
user-configured setting.
*/
func (self *TraitDisplay) SetDoubleClickDistance(distance uint) {
	C.gdk_display_set_double_click_distance(self.CPointer, C.guint(distance))
	return
}

/*
Sets the double click time (two clicks within this time interval
count as a double click and result in a #GDK_2BUTTON_PRESS event).
Applications should not set this, it is a global
user-configured setting.
*/
func (self *TraitDisplay) SetDoubleClickTime(msec uint) {
	C.gdk_display_set_double_click_time(self.CPointer, C.guint(msec))
	return
}

/*
Issues a request to the clipboard manager to store the
clipboard data. On X11, this is a special program that works
according to the
[FreeDesktop Clipboard Specification](http://www.freedesktop.org/Standards/clipboard-manager-spec).
*/
func (self *TraitDisplay) StoreClipboard(clipboard_window IsWindow, time_ uint32, targets *C.GdkAtom, n_targets int) {
	C.gdk_display_store_clipboard(self.CPointer, clipboard_window.GetWindowPointer(), C.guint32(time_), targets, C.gint(n_targets))
	return
}

/*
Returns whether the speicifed display supports clipboard
persistance; i.e. if it’s possible to store the clipboard data after an
application has quit. On X11 this checks if a clipboard daemon is
running.
*/
func (self *TraitDisplay) SupportsClipboardPersistence() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_display_supports_clipboard_persistence(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if gdk_window_set_composited() can be used
to redirect drawing on the window using compositing.

Currently this only works on X11 with XComposite and
XDamage extensions available.
*/
func (self *TraitDisplay) SupportsComposite() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_display_supports_composite(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if cursors can use an 8bit alpha channel
on @display. Otherwise, cursors are restricted to bilevel
alpha (i.e. a mask).
*/
func (self *TraitDisplay) SupportsCursorAlpha() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_display_supports_cursor_alpha(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if multicolored cursors are supported
on @display. Otherwise, cursors have only a forground
and a background color.
*/
func (self *TraitDisplay) SupportsCursorColor() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_display_supports_cursor_color(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if gdk_window_input_shape_combine_mask() can
be used to modify the input shape of windows on @display.
*/
func (self *TraitDisplay) SupportsInputShapes() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_display_supports_input_shapes(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether #GdkEventOwnerChange events will be
sent when the owner of a selection changes.
*/
func (self *TraitDisplay) SupportsSelectionNotification() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_display_supports_selection_notification(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if gdk_window_shape_combine_mask() can
be used to create shaped windows on @display.
*/
func (self *TraitDisplay) SupportsShapes() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_display_supports_shapes(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Flushes any requests queued for the windowing system and waits until all
requests have been handled. This is often used for making sure that the
display is synchronized with the current state of the program. Calling
gdk_display_sync() before gdk_error_trap_pop() makes sure that any errors
generated from earlier requests are handled before the error trap is
removed.

This is most useful for X11. On windowing systems where requests are
handled synchronously, this function will do nothing.
*/
func (self *TraitDisplay) Sync() {
	C.gdk_display_sync(self.CPointer)
	return
}

// gdk_display_warp_pointer is not generated due to deprecation attr

type TraitDisplayManager struct{ CPointer *C.GdkDisplayManager }
type IsDisplayManager interface {
	GetDisplayManagerPointer() *C.GdkDisplayManager
}

func (self *TraitDisplayManager) GetDisplayManagerPointer() *C.GdkDisplayManager {
	return self.CPointer
}
func NewTraitDisplayManager(p unsafe.Pointer) *TraitDisplayManager {
	return &TraitDisplayManager{(*C.GdkDisplayManager)(p)}
}

/*
Gets the default #GdkDisplay.
*/
func (self *TraitDisplayManager) GetDefaultDisplay() (return__ *Display) {
	var __cgo__return__ *C.GdkDisplay
	__cgo__return__ = C.gdk_display_manager_get_default_display(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDisplayFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
List all currently open displays.
*/
func (self *TraitDisplayManager) ListDisplays() (return__ *C.GSList) {
	return__ = C.gdk_display_manager_list_displays(self.CPointer)
	return
}

/*
Opens a display.
*/
func (self *TraitDisplayManager) OpenDisplay(name string) (return__ *Display) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ *C.GdkDisplay
	__cgo__return__ = C.gdk_display_manager_open_display(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	if __cgo__return__ != nil {
		return__ = NewDisplayFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Sets @display as the default display.
*/
func (self *TraitDisplayManager) SetDefaultDisplay(display IsDisplay) {
	C.gdk_display_manager_set_default_display(self.CPointer, display.GetDisplayPointer())
	return
}

type TraitDragContext struct{ CPointer *C.GdkDragContext }
type IsDragContext interface {
	GetDragContextPointer() *C.GdkDragContext
}

func (self *TraitDragContext) GetDragContextPointer() *C.GdkDragContext {
	return self.CPointer
}
func NewTraitDragContext(p unsafe.Pointer) *TraitDragContext {
	return &TraitDragContext{(*C.GdkDragContext)(p)}
}

/*
Determines the bitmask of actions proposed by the source if
gdk_drag_context_get_suggested_action() returns GDK_ACTION_ASK.
*/
func (self *TraitDragContext) GetActions() (return__ C.GdkDragAction) {
	return__ = C.gdk_drag_context_get_actions(self.CPointer)
	return
}

/*
Returns the destination windw for the DND operation.
*/
func (self *TraitDragContext) GetDestWindow() (return__ *Window) {
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_drag_context_get_dest_window(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the #GdkDevice associated to the drag context.
*/
func (self *TraitDragContext) GetDevice() (return__ *Device) {
	var __cgo__return__ *C.GdkDevice
	__cgo__return__ = C.gdk_drag_context_get_device(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDeviceFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the drag protocol thats used by this context.
*/
func (self *TraitDragContext) GetProtocol() (return__ C.GdkDragProtocol) {
	return__ = C.gdk_drag_context_get_protocol(self.CPointer)
	return
}

/*
Determines the action chosen by the drag destination.
*/
func (self *TraitDragContext) GetSelectedAction() (return__ C.GdkDragAction) {
	return__ = C.gdk_drag_context_get_selected_action(self.CPointer)
	return
}

/*
Returns the #GdkWindow where the DND operation started.
*/
func (self *TraitDragContext) GetSourceWindow() (return__ *Window) {
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_drag_context_get_source_window(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Determines the suggested drag action of the context.
*/
func (self *TraitDragContext) GetSuggestedAction() (return__ C.GdkDragAction) {
	return__ = C.gdk_drag_context_get_suggested_action(self.CPointer)
	return
}

/*
Retrieves the list of targets of the context.
*/
func (self *TraitDragContext) ListTargets() (return__ *C.GList) {
	return__ = C.gdk_drag_context_list_targets(self.CPointer)
	return
}

/*
Associates a #GdkDevice to @context, so all Drag and Drop events
for @context are emitted as if they came from this device.
*/
func (self *TraitDragContext) SetDevice(device IsDevice) {
	C.gdk_drag_context_set_device(self.CPointer, device.GetDevicePointer())
	return
}

type TraitFrameClock struct{ CPointer *C.GdkFrameClock }
type IsFrameClock interface {
	GetFrameClockPointer() *C.GdkFrameClock
}

func (self *TraitFrameClock) GetFrameClockPointer() *C.GdkFrameClock {
	return self.CPointer
}
func NewTraitFrameClock(p unsafe.Pointer) *TraitFrameClock {
	return &TraitFrameClock{(*C.GdkFrameClock)(p)}
}

/*
Starts updates for an animation. Until a matching call to
gdk_frame_clock_end_updating() is made, the frame clock will continually
request a new frame with the %GDK_FRAME_CLOCK_PHASE_UPDATE phase.
This function may be called multiple times and frames will be
requested until gdk_frame_clock_end_updating() is called the same
number of times.
*/
func (self *TraitFrameClock) BeginUpdating() {
	C.gdk_frame_clock_begin_updating(self.CPointer)
	return
}

/*
Stops updates for an animation. See the documentation for
gdk_frame_clock_begin_updating().
*/
func (self *TraitFrameClock) EndUpdating() {
	C.gdk_frame_clock_end_updating(self.CPointer)
	return
}

/*
Gets the frame timings for the current frame.
*/
func (self *TraitFrameClock) GetCurrentTimings() (return__ *C.GdkFrameTimings) {
	return__ = C.gdk_frame_clock_get_current_timings(self.CPointer)
	return
}

/*
A #GdkFrameClock maintains a 64-bit counter that increments for
each frame drawn.
*/
func (self *TraitFrameClock) GetFrameCounter() (return__ int64) {
	var __cgo__return__ C.gint64
	__cgo__return__ = C.gdk_frame_clock_get_frame_counter(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

/*
Gets the time that should currently be used for animations.  Inside
the processing of a frame, it’s the time used to compute the
animation position of everything in a frame. Outside of a frame, it's
the time of the conceptual “previous frame,” which may be either
the actual previous frame time, or if that’s too old, an updated
time.
*/
func (self *TraitFrameClock) GetFrameTime() (return__ int64) {
	var __cgo__return__ C.gint64
	__cgo__return__ = C.gdk_frame_clock_get_frame_time(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

/*
#GdkFrameClock internally keeps a history of #GdkFrameTimings
objects for recent frames that can be retrieved with
gdk_frame_clock_get_timings(). The set of stored frames
is the set from the counter values given by
gdk_frame_clock_get_history_start() and
gdk_frame_clock_get_frame_counter(), inclusive.
*/
func (self *TraitFrameClock) GetHistoryStart() (return__ int64) {
	var __cgo__return__ C.gint64
	__cgo__return__ = C.gdk_frame_clock_get_history_start(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

/*
Using the frame history stored in the frame clock, finds the last
known presentation time and refresh interval, and assuming that
presentation times are separated by the refresh interval,
predicts a presentation time that is a multiple of the refresh
interval after the last presentation time, and later than @base_time.
*/
func (self *TraitFrameClock) GetRefreshInfo(base_time int64, refresh_interval_return *C.gint64, presentation_time_return *C.gint64) {
	C.gdk_frame_clock_get_refresh_info(self.CPointer, C.gint64(base_time), refresh_interval_return, presentation_time_return)
	return
}

/*
Retrieves a #GdkFrameTimings object holding timing information
for the current frame or a recent frame. The #GdkFrameTimings
object may not yet be complete: see gdk_frame_timings_get_complete().
*/
func (self *TraitFrameClock) GetTimings(frame_counter int64) (return__ *C.GdkFrameTimings) {
	return__ = C.gdk_frame_clock_get_timings(self.CPointer, C.gint64(frame_counter))
	return
}

/*
Asks the frame clock to run a particular phase. The signal
corresponding the requested phase will be emitted the next
time the frame clock processes. Multiple calls to
gdk_frame_clock_request_phase() will be combined together
and only one frame processed. If you are displaying animated
content and want to continually request the
%GDK_FRAME_CLOCK_PHASE_UPDATE phase for a period of time,
you should use gdk_frame_clock_begin_updating() instead, since
this allows GTK+ to adjust system parameters to get maximally
smooth animations.
*/
func (self *TraitFrameClock) RequestPhase(phase C.GdkFrameClockPhase) {
	C.gdk_frame_clock_request_phase(self.CPointer, phase)
	return
}

type TraitKeymap struct{ CPointer *C.GdkKeymap }
type IsKeymap interface {
	GetKeymapPointer() *C.GdkKeymap
}

func (self *TraitKeymap) GetKeymapPointer() *C.GdkKeymap {
	return self.CPointer
}
func NewTraitKeymap(p unsafe.Pointer) *TraitKeymap {
	return &TraitKeymap{(*C.GdkKeymap)(p)}
}

// gdk_keymap_add_virtual_modifiers is not generated due to inout param

/*
Returns whether the Caps Lock modifer is locked.
*/
func (self *TraitKeymap) GetCapsLockState() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_keymap_get_caps_lock_state(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the direction of effective layout of the keymap.
*/
func (self *TraitKeymap) GetDirection() (return__ C.PangoDirection) {
	return__ = C.gdk_keymap_get_direction(self.CPointer)
	return
}

/*
Returns the keyvals bound to @hardware_keycode.
The Nth #GdkKeymapKey in @keys is bound to the Nth
keyval in @keyvals. Free the returned arrays with g_free().
When a keycode is pressed by the user, the keyval from
this list of entries is selected by considering the effective
keyboard group and level. See gdk_keymap_translate_keyboard_state().
*/
func (self *TraitKeymap) GetEntriesForKeycode(hardware_keycode uint) (keys *C.GdkKeymapKey, keyvals *C.guint, n_entries int, return__ bool) {
	var __cgo__n_entries C.gint
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_keymap_get_entries_for_keycode(self.CPointer, C.guint(hardware_keycode), &keys, &keyvals, &__cgo__n_entries)
	n_entries = int(__cgo__n_entries)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Obtains a list of keycode/group/level combinations that will
generate @keyval. Groups and levels are two kinds of keyboard mode;
in general, the level determines whether the top or bottom symbol
on a key is used, and the group determines whether the left or
right symbol is used. On US keyboards, the shift key changes the
keyboard level, and there are no groups. A group switch key might
convert a keyboard between Hebrew to English modes, for example.
#GdkEventKey contains a %group field that indicates the active
keyboard group. The level is computed from the modifier mask.
The returned array should be freed
with g_free().
*/
func (self *TraitKeymap) GetEntriesForKeyval(keyval uint) (keys *C.GdkKeymapKey, n_keys int, return__ bool) {
	var __cgo__n_keys C.gint
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_keymap_get_entries_for_keyval(self.CPointer, C.guint(keyval), &keys, &__cgo__n_keys)
	n_keys = int(__cgo__n_keys)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the modifier mask the @keymap’s windowing system backend
uses for a particular purpose.

Note that this function always returns real hardware modifiers, not
virtual ones (e.g. it will return #GDK_MOD1_MASK rather than
#GDK_META_MASK if the backend maps MOD1 to META), so there are use
cases where the return value of this function has to be transformed
by gdk_keymap_add_virtual_modifiers() in order to contain the
expected result.
*/
func (self *TraitKeymap) GetModifierMask(intent C.GdkModifierIntent) (return__ C.GdkModifierType) {
	return__ = C.gdk_keymap_get_modifier_mask(self.CPointer, intent)
	return
}

/*
Returns the current modifier state.
*/
func (self *TraitKeymap) GetModifierState() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gdk_keymap_get_modifier_state(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Returns whether the Num Lock modifer is locked.
*/
func (self *TraitKeymap) GetNumLockState() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_keymap_get_num_lock_state(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines if keyboard layouts for both right-to-left and left-to-right
languages are in use.
*/
func (self *TraitKeymap) HaveBidiLayouts() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_keymap_have_bidi_layouts(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Looks up the keyval mapped to a keycode/group/level triplet.
If no keyval is bound to @key, returns 0. For normal user input,
you want to use gdk_keymap_translate_keyboard_state() instead of
this function, since the effective group/level may not be
the same as the current keyboard state.
*/
func (self *TraitKeymap) LookupKey(key *C.GdkKeymapKey) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gdk_keymap_lookup_key(self.CPointer, key)
	return__ = uint(__cgo__return__)
	return
}

// gdk_keymap_map_virtual_modifiers is not generated due to inout param

/*
Translates the contents of a #GdkEventKey into a keyval, effective
group, and level. Modifiers that affected the translation and
are thus unavailable for application use are returned in
@consumed_modifiers.
See [Groups][key-group-explanation] for an explanation of
groups and levels. The @effective_group is the group that was
actually used for the translation; some keys such as Enter are not
affected by the active keyboard group. The @level is derived from
@state. For convenience, #GdkEventKey already contains the translated
keyval, so this function isn’t as useful as you might think.

@consumed_modifiers gives modifiers that should be masked outfrom @state
when comparing this key press to a hot key. For instance, on a US keyboard,
the `plus` symbol is shifted, so when comparing a key press to a
`<Control>plus` accelerator `<Shift>` should be masked out.

|[<!-- language="C" -->
// We want to ignore irrelevant modifiers like ScrollLock
#define ALL_ACCELS_MASK (GDK_CONTROL_MASK | GDK_SHIFT_MASK | GDK_MOD1_MASK)
gdk_keymap_translate_keyboard_state (keymap, event->hardware_keycode,
                                     event->state, event->group,
                                     &keyval, NULL, NULL, &consumed);
if (keyval == GDK_PLUS &&
    (event->state & ~consumed & ALL_ACCELS_MASK) == GDK_CONTROL_MASK)
  // Control was pressed
]|

An older interpretation @consumed_modifiers was that it contained
all modifiers that might affect the translation of the key;
this allowed accelerators to be stored with irrelevant consumed
modifiers, by doing:
|[<!-- language="C" -->
// XXX Don’t do this XXX
if (keyval == accel_keyval &&
    (event->state & ~consumed & ALL_ACCELS_MASK) == (accel_mods & ~consumed))
  // Accelerator was pressed
]|

However, this did not work if multi-modifier combinations were
used in the keymap, since, for instance, `<Control>` would be
masked out even if only `<Control><Alt>` was used in the keymap.
To support this usage as well as well as possible, all single
modifier combinations that could affect the key for any combination
of modifiers will be returned in @consumed_modifiers; multi-modifier
combinations are returned only when actually found in @state. When
you store accelerators, you should always store them with consumed
modifiers removed. Store `<Control>plus`, not `<Control><Shift>plus`,
*/
func (self *TraitKeymap) TranslateKeyboardState(hardware_keycode uint, state C.GdkModifierType, group int) (keyval uint, effective_group int, level int, consumed_modifiers C.GdkModifierType, return__ bool) {
	var __cgo__keyval C.guint
	var __cgo__effective_group C.gint
	var __cgo__level C.gint
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_keymap_translate_keyboard_state(self.CPointer, C.guint(hardware_keycode), state, C.gint(group), &__cgo__keyval, &__cgo__effective_group, &__cgo__level, &consumed_modifiers)
	keyval = uint(__cgo__keyval)
	effective_group = int(__cgo__effective_group)
	level = int(__cgo__level)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

type TraitScreen struct{ CPointer *C.GdkScreen }
type IsScreen interface {
	GetScreenPointer() *C.GdkScreen
}

func (self *TraitScreen) GetScreenPointer() *C.GdkScreen {
	return self.CPointer
}
func NewTraitScreen(p unsafe.Pointer) *TraitScreen {
	return &TraitScreen{(*C.GdkScreen)(p)}
}

/*
Returns the screen’s currently active window.

On X11, this is done by inspecting the _NET_ACTIVE_WINDOW property
on the root window, as described in the
[Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec).
If there is no currently currently active
window, or the window manager does not support the
_NET_ACTIVE_WINDOW hint, this function returns %NULL.

On other platforms, this function may return %NULL, depending on whether
it is implementable on that platform.

The returned window should be unrefed using g_object_unref() when
no longer needed.
*/
func (self *TraitScreen) GetActiveWindow() (return__ *Window) {
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_screen_get_active_window(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the display to which the @screen belongs.
*/
func (self *TraitScreen) GetDisplay() (return__ *Display) {
	var __cgo__return__ *C.GdkDisplay
	__cgo__return__ = C.gdk_screen_get_display(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDisplayFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets any options previously set with gdk_screen_set_font_options().
*/
func (self *TraitScreen) GetFontOptions() (return__ *C.cairo_font_options_t) {
	return__ = C.gdk_screen_get_font_options(self.CPointer)
	return
}

/*
Gets the height of @screen in pixels
*/
func (self *TraitScreen) GetHeight() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_screen_get_height(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the height of @screen in millimeters.
Note that on some X servers this value will not be correct.
*/
func (self *TraitScreen) GetHeightMm() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_screen_get_height_mm(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the monitor number in which the point (@x,@y) is located.
*/
func (self *TraitScreen) GetMonitorAtPoint(x int, y int) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_screen_get_monitor_at_point(self.CPointer, C.gint(x), C.gint(y))
	return__ = int(__cgo__return__)
	return
}

/*
Returns the number of the monitor in which the largest area of the
bounding rectangle of @window resides.
*/
func (self *TraitScreen) GetMonitorAtWindow(window IsWindow) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_screen_get_monitor_at_window(self.CPointer, window.GetWindowPointer())
	return__ = int(__cgo__return__)
	return
}

/*
Retrieves the #GdkRectangle representing the size and position of
the individual monitor within the entire screen area.

Monitor numbers start at 0. To obtain the number of monitors of
@screen, use gdk_screen_get_n_monitors().

Note that the size of the entire screen area can be retrieved via
gdk_screen_get_width() and gdk_screen_get_height().
*/
func (self *TraitScreen) GetMonitorGeometry(monitor_num int) (dest C.GdkRectangle) {
	C.gdk_screen_get_monitor_geometry(self.CPointer, C.gint(monitor_num), &dest)
	return
}

/*
Gets the height in millimeters of the specified monitor.
*/
func (self *TraitScreen) GetMonitorHeightMm(monitor_num int) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_screen_get_monitor_height_mm(self.CPointer, C.gint(monitor_num))
	return__ = int(__cgo__return__)
	return
}

/*
Returns the output name of the specified monitor.
Usually something like VGA, DVI, or TV, not the actual
product name of the display device.
*/
func (self *TraitScreen) GetMonitorPlugName(monitor_num int) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gdk_screen_get_monitor_plug_name(self.CPointer, C.gint(monitor_num))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the internal scale factor that maps from monitor coordiantes
to the actual device pixels. On traditional systems this is 1, but
on very high density outputs this can be a higher value (often 2).

This can be used if you want to create pixel based data for a
particula monitor, but most of the time you’re drawing to a window
where it is better to use gdk_window_get_scale_factor() instead.
*/
func (self *TraitScreen) GetMonitorScaleFactor(monitor_num int) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_screen_get_monitor_scale_factor(self.CPointer, C.gint(monitor_num))
	return__ = int(__cgo__return__)
	return
}

/*
Gets the width in millimeters of the specified monitor, if available.
*/
func (self *TraitScreen) GetMonitorWidthMm(monitor_num int) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_screen_get_monitor_width_mm(self.CPointer, C.gint(monitor_num))
	return__ = int(__cgo__return__)
	return
}

/*
Retrieves the #GdkRectangle representing the size and position of
the “work area” on a monitor within the entire screen area.

The work area should be considered when positioning menus and
similar popups, to avoid placing them below panels, docks or other
desktop components.

Monitor numbers start at 0. To obtain the number of monitors of
@screen, use gdk_screen_get_n_monitors().
*/
func (self *TraitScreen) GetMonitorWorkarea(monitor_num int) (dest C.GdkRectangle) {
	C.gdk_screen_get_monitor_workarea(self.CPointer, C.gint(monitor_num), &dest)
	return
}

/*
Returns the number of monitors which @screen consists of.
*/
func (self *TraitScreen) GetNMonitors() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_screen_get_n_monitors(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the index of @screen among the screens in the display
to which it belongs. (See gdk_screen_get_display())
*/
func (self *TraitScreen) GetNumber() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_screen_get_number(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the primary monitor for @screen.  The primary monitor
is considered the monitor where the “main desktop” lives.
While normal application windows typically allow the window
manager to place the windows, specialized desktop applications
such as panels should place themselves on the primary monitor.

If no primary monitor is configured by the user, the return value
will be 0, defaulting to the first monitor.
*/
func (self *TraitScreen) GetPrimaryMonitor() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_screen_get_primary_monitor(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the resolution for font handling on the screen; see
gdk_screen_set_resolution() for full details.
*/
func (self *TraitScreen) GetResolution() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gdk_screen_get_resolution(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets a visual to use for creating windows with an alpha channel.
The windowing system on which GTK+ is running
may not support this capability, in which case %NULL will
be returned. Even if a non-%NULL value is returned, its
possible that the window’s alpha channel won’t be honored
when displaying the window on the screen: in particular, for
X an appropriate windowing manager and compositing manager
must be running to provide appropriate display.

This functionality is not implemented in the Windows backend.

For setting an overall opacity for a top-level window, see
gdk_window_set_opacity().
*/
func (self *TraitScreen) GetRgbaVisual() (return__ *Visual) {
	var __cgo__return__ *C.GdkVisual
	__cgo__return__ = C.gdk_screen_get_rgba_visual(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewVisualFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the root window of @screen.
*/
func (self *TraitScreen) GetRootWindow() (return__ *Window) {
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_screen_get_root_window(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Retrieves a desktop-wide setting such as double-click time
for the #GdkScreen @screen.

FIXME needs a list of valid settings here, or a link to
more information.
*/
func (self *TraitScreen) GetSetting(name string, value *C.GValue) (return__ bool) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_screen_get_setting(self.CPointer, __cgo__name, value)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the system’s default visual for @screen.
This is the visual for the root window of the display.
The return value should not be freed.
*/
func (self *TraitScreen) GetSystemVisual() (return__ *Visual) {
	var __cgo__return__ *C.GdkVisual
	__cgo__return__ = C.gdk_screen_get_system_visual(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewVisualFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Obtains a list of all toplevel windows known to GDK on the screen @screen.
A toplevel window is a child of the root window (see
gdk_get_default_root_window()).

The returned list should be freed with g_list_free(), but
its elements need not be freed.
*/
func (self *TraitScreen) GetToplevelWindows() (return__ *C.GList) {
	return__ = C.gdk_screen_get_toplevel_windows(self.CPointer)
	return
}

/*
Gets the width of @screen in pixels
*/
func (self *TraitScreen) GetWidth() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_screen_get_width(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the width of @screen in millimeters.
Note that on some X servers this value will not be correct.
*/
func (self *TraitScreen) GetWidthMm() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_screen_get_width_mm(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns a #GList of #GdkWindows representing the current
window stack.

On X11, this is done by inspecting the _NET_CLIENT_LIST_STACKING
property on the root window, as described in the
[Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec).
If the window manager does not support the
_NET_CLIENT_LIST_STACKING hint, this function returns %NULL.

On other platforms, this function may return %NULL, depending on whether
it is implementable on that platform.

The returned list is newly allocated and owns references to the
windows it contains, so it should be freed using g_list_free() and
its windows unrefed using g_object_unref() when no longer needed.
*/
func (self *TraitScreen) GetWindowStack() (return__ *C.GList) {
	return__ = C.gdk_screen_get_window_stack(self.CPointer)
	return
}

/*
Returns whether windows with an RGBA visual can reasonably
be expected to have their alpha channel drawn correctly on
the screen.

On X11 this function returns whether a compositing manager is
compositing @screen.
*/
func (self *TraitScreen) IsComposited() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_screen_is_composited(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Lists the available visuals for the specified @screen.
A visual describes a hardware image data format.
For example, a visual might support 24-bit color, or 8-bit color,
and might expect pixels to be in a certain format.

Call g_list_free() on the return value when you’re finished with it.
*/
func (self *TraitScreen) ListVisuals() (return__ *C.GList) {
	return__ = C.gdk_screen_list_visuals(self.CPointer)
	return
}

/*
Determines the name to pass to gdk_display_open() to get
a #GdkDisplay with this screen as the default screen.
*/
func (self *TraitScreen) MakeDisplayName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gdk_screen_make_display_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Sets the default font options for the screen. These
options will be set on any #PangoContext’s newly created
with gdk_pango_context_get_for_screen(). Changing the
default set of font options does not affect contexts that
have already been created.
*/
func (self *TraitScreen) SetFontOptions(options *C.cairo_font_options_t) {
	C.gdk_screen_set_font_options(self.CPointer, options)
	return
}

/*
Sets the resolution for font handling on the screen. This is a
scale factor between points specified in a #PangoFontDescription
and cairo units. The default value is 96, meaning that a 10 point
font will be 13 units high. (10 * 96. / 72. = 13.3).
*/
func (self *TraitScreen) SetResolution(dpi float64) {
	C.gdk_screen_set_resolution(self.CPointer, C.gdouble(dpi))
	return
}

type TraitVisual struct{ CPointer *C.GdkVisual }
type IsVisual interface {
	GetVisualPointer() *C.GdkVisual
}

func (self *TraitVisual) GetVisualPointer() *C.GdkVisual {
	return self.CPointer
}
func NewTraitVisual(p unsafe.Pointer) *TraitVisual {
	return &TraitVisual{(*C.GdkVisual)(p)}
}

/*
Returns the number of significant bits per red, green and blue value.
*/
func (self *TraitVisual) GetBitsPerRgb() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_visual_get_bits_per_rgb(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Obtains values that are needed to calculate blue pixel values in TrueColor
and DirectColor. The “mask” is the significant bits within the pixel.
The “shift” is the number of bits left we must shift a primary for it
to be in position (according to the "mask"). Finally, "precision" refers
to how much precision the pixel value contains for a particular primary.
*/
func (self *TraitVisual) GetBluePixelDetails() (mask uint32, shift int, precision int) {
	var __cgo__mask C.guint32
	var __cgo__shift C.gint
	var __cgo__precision C.gint
	C.gdk_visual_get_blue_pixel_details(self.CPointer, &__cgo__mask, &__cgo__shift, &__cgo__precision)
	mask = uint32(__cgo__mask)
	shift = int(__cgo__shift)
	precision = int(__cgo__precision)
	return
}

/*
Returns the byte order of this visual.
*/
func (self *TraitVisual) GetByteOrder() (return__ C.GdkByteOrder) {
	return__ = C.gdk_visual_get_byte_order(self.CPointer)
	return
}

/*
Returns the size of a colormap for this visual.
*/
func (self *TraitVisual) GetColormapSize() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_visual_get_colormap_size(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the bit depth of this visual.
*/
func (self *TraitVisual) GetDepth() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_visual_get_depth(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Obtains values that are needed to calculate green pixel values in TrueColor
and DirectColor. The “mask” is the significant bits within the pixel.
The “shift” is the number of bits left we must shift a primary for it
to be in position (according to the "mask"). Finally, "precision" refers
to how much precision the pixel value contains for a particular primary.
*/
func (self *TraitVisual) GetGreenPixelDetails() (mask uint32, shift int, precision int) {
	var __cgo__mask C.guint32
	var __cgo__shift C.gint
	var __cgo__precision C.gint
	C.gdk_visual_get_green_pixel_details(self.CPointer, &__cgo__mask, &__cgo__shift, &__cgo__precision)
	mask = uint32(__cgo__mask)
	shift = int(__cgo__shift)
	precision = int(__cgo__precision)
	return
}

/*
Obtains values that are needed to calculate red pixel values in TrueColor
and DirectColor. The “mask” is the significant bits within the pixel.
The “shift” is the number of bits left we must shift a primary for it
to be in position (according to the "mask"). Finally, "precision" refers
to how much precision the pixel value contains for a particular primary.
*/
func (self *TraitVisual) GetRedPixelDetails() (mask uint32, shift int, precision int) {
	var __cgo__mask C.guint32
	var __cgo__shift C.gint
	var __cgo__precision C.gint
	C.gdk_visual_get_red_pixel_details(self.CPointer, &__cgo__mask, &__cgo__shift, &__cgo__precision)
	mask = uint32(__cgo__mask)
	shift = int(__cgo__shift)
	precision = int(__cgo__precision)
	return
}

/*
Gets the screen to which this visual belongs
*/
func (self *TraitVisual) GetScreen() (return__ *Screen) {
	var __cgo__return__ *C.GdkScreen
	__cgo__return__ = C.gdk_visual_get_screen(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewScreenFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the type of visual this is (PseudoColor, TrueColor, etc).
*/
func (self *TraitVisual) GetVisualType() (return__ C.GdkVisualType) {
	return__ = C.gdk_visual_get_visual_type(self.CPointer)
	return
}

type TraitWindow struct{ CPointer *C.GdkWindow }
type IsWindow interface {
	GetWindowPointer() *C.GdkWindow
}

func (self *TraitWindow) GetWindowPointer() *C.GdkWindow {
	return self.CPointer
}
func NewTraitWindow(p unsafe.Pointer) *TraitWindow {
	return &TraitWindow{(*C.GdkWindow)(p)}
}

/*
Adds an event filter to @window, allowing you to intercept events
before they reach GDK. This is a low-level operation and makes it
easy to break GDK and/or GTK+, so you have to know what you're
doing. Pass %NULL for @window to get all events for all windows,
instead of events for a specific window.

If you are interested in X GenericEvents, bear in mind that
XGetEventData() has been already called on the event, and
XFreeEventData() must not be called within @function.
*/
func (self *TraitWindow) AddFilter(function C.GdkFilterFunc, data unsafe.Pointer) {
	C.gdk_window_add_filter(self.CPointer, function, (C.gpointer)(data))
	return
}

/*
Emits a short beep associated to @window in the appropriate
display, if supported. Otherwise, emits a short beep on
the display just as gdk_display_beep().
*/
func (self *TraitWindow) Beep() {
	C.gdk_window_beep(self.CPointer)
	return
}

/*
Begins a window move operation (for a toplevel window).

This function assumes that the drag is controlled by the
client pointer device, use gdk_window_begin_move_drag_for_device()
to begin a drag with a different device.
*/
func (self *TraitWindow) BeginMoveDrag(button int, root_x int, root_y int, timestamp uint32) {
	C.gdk_window_begin_move_drag(self.CPointer, C.gint(button), C.gint(root_x), C.gint(root_y), C.guint32(timestamp))
	return
}

/*
Begins a window move operation (for a toplevel window).
You might use this function to implement a “window move grip,” for
example. The function works best with window managers that support the
[Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec)
but has a fallback implementation for other window managers.
*/
func (self *TraitWindow) BeginMoveDragForDevice(device IsDevice, button int, root_x int, root_y int, timestamp uint32) {
	C.gdk_window_begin_move_drag_for_device(self.CPointer, device.GetDevicePointer(), C.gint(button), C.gint(root_x), C.gint(root_y), C.guint32(timestamp))
	return
}

/*
A convenience wrapper around gdk_window_begin_paint_region() which
creates a rectangular region for you. See
gdk_window_begin_paint_region() for details.
*/
func (self *TraitWindow) BeginPaintRect(rectangle *C.GdkRectangle) {
	C.gdk_window_begin_paint_rect(self.CPointer, rectangle)
	return
}

/*
Indicates that you are beginning the process of redrawing @region.
A backing store (offscreen buffer) large enough to contain @region
will be created. The backing store will be initialized with the
background color or background surface for @window. Then, all
drawing operations performed on @window will be diverted to the
backing store.  When you call gdk_window_end_paint(), the backing
store will be copied to @window, making it visible onscreen. Only
the part of @window contained in @region will be modified; that is,
drawing operations are clipped to @region.

The net result of all this is to remove flicker, because the user
sees the finished product appear all at once when you call
gdk_window_end_paint(). If you draw to @window directly without
calling gdk_window_begin_paint_region(), the user may see flicker
as individual drawing operations are performed in sequence.  The
clipping and background-initializing features of
gdk_window_begin_paint_region() are conveniences for the
programmer, so you can avoid doing that work yourself.

When using GTK+, the widget system automatically places calls to
gdk_window_begin_paint_region() and gdk_window_end_paint() around
emissions of the expose_event signal. That is, if you’re writing an
expose event handler, you can assume that the exposed area in
#GdkEventExpose has already been cleared to the window background,
is already set as the clip region, and already has a backing store.
Therefore in most cases, application code need not call
gdk_window_begin_paint_region(). (You can disable the automatic
calls around expose events on a widget-by-widget basis by calling
gtk_widget_set_double_buffered().)

If you call this function multiple times before calling the
matching gdk_window_end_paint(), the backing stores are pushed onto
a stack. gdk_window_end_paint() copies the topmost backing store
onscreen, subtracts the topmost region from all other regions in
the stack, and pops the stack. All drawing operations affect only
the topmost backing store in the stack. One matching call to
gdk_window_end_paint() is required for each call to
gdk_window_begin_paint_region().
*/
func (self *TraitWindow) BeginPaintRegion(region *C.cairo_region_t) {
	C.gdk_window_begin_paint_region(self.CPointer, region)
	return
}

/*
Begins a window resize operation (for a toplevel window).

This function assumes that the drag is controlled by the
client pointer device, use gdk_window_begin_resize_drag_for_device()
to begin a drag with a different device.
*/
func (self *TraitWindow) BeginResizeDrag(edge C.GdkWindowEdge, button int, root_x int, root_y int, timestamp uint32) {
	C.gdk_window_begin_resize_drag(self.CPointer, edge, C.gint(button), C.gint(root_x), C.gint(root_y), C.guint32(timestamp))
	return
}

/*
Begins a window resize operation (for a toplevel window).
You might use this function to implement a “window resize grip,” for
example; in fact #GtkStatusbar uses it. The function works best
with window managers that support the
[Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec)
but has a fallback implementation for other window managers.
*/
func (self *TraitWindow) BeginResizeDragForDevice(edge C.GdkWindowEdge, device IsDevice, button int, root_x int, root_y int, timestamp uint32) {
	C.gdk_window_begin_resize_drag_for_device(self.CPointer, edge, device.GetDevicePointer(), C.gint(button), C.gint(root_x), C.gint(root_y), C.guint32(timestamp))
	return
}

// gdk_window_configure_finished is not generated due to deprecation attr

/*
Transforms window coordinates from a parent window to a child
window, where the parent window is the normal parent as returned by
gdk_window_get_parent() for normal windows, and the window's
embedder as returned by gdk_offscreen_window_get_embedder() for
offscreen windows.

For normal windows, calling this function is equivalent to subtracting
the return values of gdk_window_get_position() from the parent coordinates.
For offscreen windows however (which can be arbitrarily transformed),
this function calls the GdkWindow::from-embedder: signal to translate
the coordinates.

You should always use this function when writing generic code that
walks down a window hierarchy.

See also: gdk_window_coords_to_parent()
*/
func (self *TraitWindow) CoordsFromParent(parent_x float64, parent_y float64) (x float64, y float64) {
	var __cgo__x C.gdouble
	var __cgo__y C.gdouble
	C.gdk_window_coords_from_parent(self.CPointer, C.gdouble(parent_x), C.gdouble(parent_y), &__cgo__x, &__cgo__y)
	x = float64(__cgo__x)
	y = float64(__cgo__y)
	return
}

/*
Transforms window coordinates from a child window to its parent
window, where the parent window is the normal parent as returned by
gdk_window_get_parent() for normal windows, and the window's
embedder as returned by gdk_offscreen_window_get_embedder() for
offscreen windows.

For normal windows, calling this function is equivalent to adding
the return values of gdk_window_get_position() to the child coordinates.
For offscreen windows however (which can be arbitrarily transformed),
this function calls the GdkWindow::to-embedder: signal to translate
the coordinates.

You should always use this function when writing generic code that
walks up a window hierarchy.

See also: gdk_window_coords_from_parent()
*/
func (self *TraitWindow) CoordsToParent(x float64, y float64) (parent_x float64, parent_y float64) {
	var __cgo__parent_x C.gdouble
	var __cgo__parent_y C.gdouble
	C.gdk_window_coords_to_parent(self.CPointer, C.gdouble(x), C.gdouble(y), &__cgo__parent_x, &__cgo__parent_y)
	parent_x = float64(__cgo__parent_x)
	parent_y = float64(__cgo__parent_y)
	return
}

/*
Create a new image surface that is efficient to draw on the
given @window.

Initially the surface contents are all 0 (transparent if contents
have transparency, black otherwise.)
*/
func (self *TraitWindow) CreateSimilarImageSurface(format C.cairo_format_t, width int, height int, scale int) (return__ *C.cairo_surface_t) {
	return__ = C.gdk_window_create_similar_image_surface(self.CPointer, format, C.int(width), C.int(height), C.int(scale))
	return
}

/*
Create a new surface that is as compatible as possible with the
given @window. For example the new surface will have the same
fallback resolution and font options as @window. Generally, the new
surface will also use the same backend as @window, unless that is
not possible for some reason. The type of the returned surface may
be examined with cairo_surface_get_type().

Initially the surface contents are all 0 (transparent if contents
have transparency, black otherwise.)
*/
func (self *TraitWindow) CreateSimilarSurface(content C.cairo_content_t, width int, height int) (return__ *C.cairo_surface_t) {
	return__ = C.gdk_window_create_similar_surface(self.CPointer, content, C.int(width), C.int(height))
	return
}

/*
Attempt to deiconify (unminimize) @window. On X11 the window manager may
choose to ignore the request to deiconify. When using GTK+,
use gtk_window_deiconify() instead of the #GdkWindow variant. Or better yet,
you probably want to use gtk_window_present(), which raises the window, focuses it,
unminimizes it, and puts it on the current desktop.
*/
func (self *TraitWindow) Deiconify() {
	C.gdk_window_deiconify(self.CPointer)
	return
}

/*
Destroys the window system resources associated with @window and decrements @window's
reference count. The window system resources for all children of @window are also
destroyed, but the children’s reference counts are not decremented.

Note that a window will not be destroyed automatically when its reference count
reaches zero. You must call this function yourself before that happens.
*/
func (self *TraitWindow) Destroy() {
	C.gdk_window_destroy(self.CPointer)
	return
}

func (self *TraitWindow) DestroyNotify() {
	C.gdk_window_destroy_notify(self.CPointer)
	return
}

// gdk_window_enable_synchronized_configure is not generated due to deprecation attr

/*
Indicates that the backing store created by the most recent call to
gdk_window_begin_paint_region() should be copied onscreen and
deleted, leaving the next-most-recent backing store or no backing
store at all as the active paint region. See
gdk_window_begin_paint_region() for full details. It is an error to
call this function without a matching
gdk_window_begin_paint_region() first.
*/
func (self *TraitWindow) EndPaint() {
	C.gdk_window_end_paint(self.CPointer)
	return
}

/*
Tries to ensure that there is a window-system native window for this
GdkWindow. This may fail in some situations, returning %FALSE.

Offscreen window and children of them can never have native windows.

Some backends may not support native child windows.
*/
func (self *TraitWindow) EnsureNative() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_window_ensure_native(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Flush all outstanding cached operations on a window, leaving the
window in a state which reflects all that has been drawn before.

Gdk uses multiple kinds of caching to get better performance and
nicer drawing. For instance, during exposes all paints to a window
using double buffered rendering are keep on a surface until the last
window has been exposed.

Normally this should be completely invisible to applications, as
we automatically flush the windows when required, but this might
be needed if you for instance mix direct native drawing with
gdk drawing. For Gtk widgets that don’t use double buffering this
will be called automatically before sending the expose event.
*/
func (self *TraitWindow) Flush() {
	C.gdk_window_flush(self.CPointer)
	return
}

/*
Sets keyboard focus to @window. In most cases, gtk_window_present()
should be used on a #GtkWindow, rather than calling this function.
*/
func (self *TraitWindow) Focus(timestamp uint32) {
	C.gdk_window_focus(self.CPointer, C.guint32(timestamp))
	return
}

/*
Temporarily freezes a window and all its descendants such that it won't
receive expose events.  The window will begin receiving expose events
again when gdk_window_thaw_toplevel_updates_libgtk_only() is called. If
gdk_window_freeze_toplevel_updates_libgtk_only()
has been called more than once,
gdk_window_thaw_toplevel_updates_libgtk_only() must be called
an equal number of times to begin processing exposes.

This function is not part of the GDK public API and is only
for use by GTK+.
*/
func (self *TraitWindow) FreezeToplevelUpdatesLibgtkOnly() {
	C.gdk_window_freeze_toplevel_updates_libgtk_only(self.CPointer)
	return
}

/*
Temporarily freezes a window such that it won’t receive expose
events.  The window will begin receiving expose events again when
gdk_window_thaw_updates() is called. If gdk_window_freeze_updates()
has been called more than once, gdk_window_thaw_updates() must be called
an equal number of times to begin processing exposes.
*/
func (self *TraitWindow) FreezeUpdates() {
	C.gdk_window_freeze_updates(self.CPointer)
	return
}

/*
Moves the window into fullscreen mode. This means the
window covers the entire screen and is above any panels
or task bars.

If the window was already fullscreen, then this function does nothing.

On X11, asks the window manager to put @window in a fullscreen
state, if the window manager supports this operation. Not all
window managers support this, and some deliberately ignore it or
don’t have a concept of “fullscreen”; so you can’t rely on the
fullscreenification actually happening. But it will happen with
most standard window managers, and GDK makes a best effort to get
it to happen.
*/
func (self *TraitWindow) Fullscreen() {
	C.gdk_window_fullscreen(self.CPointer)
	return
}

/*
This function informs GDK that the geometry of an embedded
offscreen window has changed. This is necessary for GDK to keep
track of which offscreen window the pointer is in.
*/
func (self *TraitWindow) GeometryChanged() {
	C.gdk_window_geometry_changed(self.CPointer)
	return
}

/*
Determines whether or not the desktop environment shuld be hinted that
the window does not want to receive input focus.
*/
func (self *TraitWindow) GetAcceptFocus() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_window_get_accept_focus(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the pattern used to clear the background on @window. If @window
does not have its own background and reuses the parent's, %NULL is
returned and you’ll have to query it yourself.
*/
func (self *TraitWindow) GetBackgroundPattern() (return__ *C.cairo_pattern_t) {
	return__ = C.gdk_window_get_background_pattern(self.CPointer)
	return
}

/*
Gets the list of children of @window known to GDK.
This function only returns children created via GDK,
so for example it’s useless when used with the root window;
it only returns windows an application created itself.

The returned list must be freed, but the elements in the
list need not be.
*/
func (self *TraitWindow) GetChildren() (return__ *C.GList) {
	return__ = C.gdk_window_get_children(self.CPointer)
	return
}

/*
Gets the list of children of @window known to GDK with a
particular @user_data set on it.

The returned list must be freed, but the elements in the
list need not be.

The list is returned in (relative) stacking order, i.e. the
lowest window is first.
*/
func (self *TraitWindow) GetChildrenWithUserData(user_data unsafe.Pointer) (return__ *C.GList) {
	return__ = C.gdk_window_get_children_with_user_data(self.CPointer, (C.gpointer)(user_data))
	return
}

/*
Computes the region of a window that potentially can be written
to by drawing primitives. This region may not take into account
other factors such as if the window is obscured by other windows,
but no area outside of this region will be affected by drawing
primitives.
*/
func (self *TraitWindow) GetClipRegion() (return__ *C.cairo_region_t) {
	return__ = C.gdk_window_get_clip_region(self.CPointer)
	return
}

/*
Determines whether @window is composited.

See gdk_window_set_composited().
*/
func (self *TraitWindow) GetComposited() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_window_get_composited(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves a #GdkCursor pointer for the cursor currently set on the
specified #GdkWindow, or %NULL.  If the return value is %NULL then
there is no custom cursor set on the specified window, and it is
using the cursor for its parent window.
*/
func (self *TraitWindow) GetCursor() (return__ *Cursor) {
	var __cgo__return__ *C.GdkCursor
	__cgo__return__ = C.gdk_window_get_cursor(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewCursorFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the decorations set on the GdkWindow with
gdk_window_set_decorations().
*/
func (self *TraitWindow) GetDecorations() (decorations C.GdkWMDecoration, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_window_get_decorations(self.CPointer, &decorations)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves a #GdkCursor pointer for the @device currently set on the
specified #GdkWindow, or %NULL.  If the return value is %NULL then
there is no custom cursor set on the specified window, and it is
using the cursor for its parent window.
*/
func (self *TraitWindow) GetDeviceCursor(device IsDevice) (return__ *Cursor) {
	var __cgo__return__ *C.GdkCursor
	__cgo__return__ = C.gdk_window_get_device_cursor(self.CPointer, device.GetDevicePointer())
	if __cgo__return__ != nil {
		return__ = NewCursorFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the event mask for @window corresponding to an specific device.
*/
func (self *TraitWindow) GetDeviceEvents(device IsDevice) (return__ C.GdkEventMask) {
	return__ = C.gdk_window_get_device_events(self.CPointer, device.GetDevicePointer())
	return
}

/*
Obtains the current device position and modifier state.
The position is given in coordinates relative to the upper left
corner of @window.

Use gdk_window_get_device_position_double() if you need subpixel precision.
*/
func (self *TraitWindow) GetDevicePosition(device IsDevice) (x int, y int, mask C.GdkModifierType, return__ *Window) {
	var __cgo__x C.gint
	var __cgo__y C.gint
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_window_get_device_position(self.CPointer, device.GetDevicePointer(), &__cgo__x, &__cgo__y, &mask)
	x = int(__cgo__x)
	y = int(__cgo__y)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Obtains the current device position in doubles and modifier state.
The position is given in coordinates relative to the upper left
corner of @window.
*/
func (self *TraitWindow) GetDevicePositionDouble(device IsDevice) (x float64, y float64, mask C.GdkModifierType, return__ *Window) {
	var __cgo__x C.gdouble
	var __cgo__y C.gdouble
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_window_get_device_position_double(self.CPointer, device.GetDevicePointer(), &__cgo__x, &__cgo__y, &mask)
	x = float64(__cgo__x)
	y = float64(__cgo__y)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the #GdkDisplay associated with a #GdkWindow.
*/
func (self *TraitWindow) GetDisplay() (return__ *Display) {
	var __cgo__return__ *C.GdkDisplay
	__cgo__return__ = C.gdk_window_get_display(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewDisplayFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Finds out the DND protocol supported by a window.
*/
func (self *TraitWindow) GetDragProtocol() (target *Window, return__ C.GdkDragProtocol) {
	var __cgo__target *C.GdkWindow
	return__ = C.gdk_window_get_drag_protocol(self.CPointer, &__cgo__target)
	if __cgo__target != nil {
		target = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__target).Pointer()))
	}
	return
}

/*
Obtains the parent of @window, as known to GDK. Works like
gdk_window_get_parent() for normal windows, but returns the
window’s embedder for offscreen windows.

See also: gdk_offscreen_window_get_embedder()
*/
func (self *TraitWindow) GetEffectiveParent() (return__ *Window) {
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_window_get_effective_parent(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the toplevel window that’s an ancestor of @window.

Works like gdk_window_get_toplevel(), but treats an offscreen window's
embedder as its parent, using gdk_window_get_effective_parent().

See also: gdk_offscreen_window_get_embedder()
*/
func (self *TraitWindow) GetEffectiveToplevel() (return__ *Window) {
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_window_get_effective_toplevel(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Get the current event compression setting for this window.
*/
func (self *TraitWindow) GetEventCompression() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_window_get_event_compression(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the event mask for @window for all master input devices. See
gdk_window_set_events().
*/
func (self *TraitWindow) GetEvents() (return__ C.GdkEventMask) {
	return__ = C.gdk_window_get_events(self.CPointer)
	return
}

/*
Determines whether or not the desktop environment should be hinted that the
window does not want to receive input focus when it is mapped.
*/
func (self *TraitWindow) GetFocusOnMap() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_window_get_focus_on_map(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the frame clock for the window. The frame clock for a window
never changes unless the window is reparented to a new toplevel
window.
*/
func (self *TraitWindow) GetFrameClock() (return__ *FrameClock) {
	var __cgo__return__ *C.GdkFrameClock
	__cgo__return__ = C.gdk_window_get_frame_clock(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewFrameClockFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Obtains the bounding box of the window, including window manager
titlebar/borders if any. The frame position is given in root window
coordinates. To get the position of the window itself (rather than
the frame) in root window coordinates, use gdk_window_get_origin().
*/
func (self *TraitWindow) GetFrameExtents() (rect C.GdkRectangle) {
	C.gdk_window_get_frame_extents(self.CPointer, &rect)
	return
}

/*
Obtains the #GdkFullscreenMode of the @window.
*/
func (self *TraitWindow) GetFullscreenMode() (return__ C.GdkFullscreenMode) {
	return__ = C.gdk_window_get_fullscreen_mode(self.CPointer)
	return
}

/*
Any of the return location arguments to this function may be %NULL,
if you aren’t interested in getting the value of that field.

The X and Y coordinates returned are relative to the parent window
of @window, which for toplevels usually means relative to the
window decorations (titlebar, etc.) rather than relative to the
root window (screen-size background window).

On the X11 platform, the geometry is obtained from the X server,
so reflects the latest position of @window; this may be out-of-sync
with the position of @window delivered in the most-recently-processed
#GdkEventConfigure. gdk_window_get_position() in contrast gets the
position from the most recent configure event.

Note: If @window is not a toplevel, it is much better
to call gdk_window_get_position(), gdk_window_get_width() and
gdk_window_get_height() instead, because it avoids the roundtrip to
the X server and because these functions support the full 32-bit
coordinate space, whereas gdk_window_get_geometry() is restricted to
the 16-bit coordinates of X11.
*/
func (self *TraitWindow) GetGeometry() (x int, y int, width int, height int) {
	var __cgo__x C.gint
	var __cgo__y C.gint
	var __cgo__width C.gint
	var __cgo__height C.gint
	C.gdk_window_get_geometry(self.CPointer, &__cgo__x, &__cgo__y, &__cgo__width, &__cgo__height)
	x = int(__cgo__x)
	y = int(__cgo__y)
	width = int(__cgo__width)
	height = int(__cgo__height)
	return
}

/*
Returns the group leader window for @window. See gdk_window_set_group().
*/
func (self *TraitWindow) GetGroup() (return__ *Window) {
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_window_get_group(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the height of the given @window.

On the X11 platform the returned size is the size reported in the
most-recently-processed configure event, rather than the current
size on the X server.
*/
func (self *TraitWindow) GetHeight() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.gdk_window_get_height(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Determines whether or not the window manager is hinted that @window
has modal behaviour.
*/
func (self *TraitWindow) GetModalHint() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_window_get_modal_hint(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Obtains the position of a window in root window coordinates.
(Compare with gdk_window_get_position() and
gdk_window_get_geometry() which return the position of a window
relative to its parent window.)
*/
func (self *TraitWindow) GetOrigin() (x int, y int, return__ int) {
	var __cgo__x C.gint
	var __cgo__y C.gint
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_window_get_origin(self.CPointer, &__cgo__x, &__cgo__y)
	x = int(__cgo__x)
	y = int(__cgo__y)
	return__ = int(__cgo__return__)
	return
}

/*
Obtains the parent of @window, as known to GDK. Does not query the
X server; thus this returns the parent as passed to gdk_window_new(),
not the actual parent. This should never matter unless you’re using
Xlib calls mixed with GDK calls on the X11 platform. It may also
matter for toplevel windows, because the window manager may choose
to reparent them.

Note that you should use gdk_window_get_effective_parent() when
writing generic code that walks up a window hierarchy, because
gdk_window_get_parent() will most likely not do what you expect if
there are offscreen windows in the hierarchy.
*/
func (self *TraitWindow) GetParent() (return__ *Window) {
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_window_get_parent(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gdk_window_get_pointer is not generated due to deprecation attr

/*
Obtains the position of the window as reported in the
most-recently-processed #GdkEventConfigure. Contrast with
gdk_window_get_geometry() which queries the X server for the
current window position, regardless of which events have been
received or processed.

The position coordinates are relative to the window’s parent window.
*/
func (self *TraitWindow) GetPosition() (x int, y int) {
	var __cgo__x C.gint
	var __cgo__y C.gint
	C.gdk_window_get_position(self.CPointer, &__cgo__x, &__cgo__y)
	x = int(__cgo__x)
	y = int(__cgo__y)
	return
}

/*
Obtains the position of a window position in root
window coordinates. This is similar to
gdk_window_get_origin() but allows you go pass
in any position in the window, not just the origin.
*/
func (self *TraitWindow) GetRootCoords(x int, y int) (root_x int, root_y int) {
	var __cgo__root_x C.gint
	var __cgo__root_y C.gint
	C.gdk_window_get_root_coords(self.CPointer, C.gint(x), C.gint(y), &__cgo__root_x, &__cgo__root_y)
	root_x = int(__cgo__root_x)
	root_y = int(__cgo__root_y)
	return
}

/*
Obtains the top-left corner of the window manager frame in root
window coordinates.
*/
func (self *TraitWindow) GetRootOrigin() (x int, y int) {
	var __cgo__x C.gint
	var __cgo__y C.gint
	C.gdk_window_get_root_origin(self.CPointer, &__cgo__x, &__cgo__y)
	x = int(__cgo__x)
	y = int(__cgo__y)
	return
}

/*
Returns the internal scale factor that maps from window coordiantes
to the actual device pixels. On traditional systems this is 1, but
on very high density outputs this can be a higher value (often 2).

A higher value means that drawing is automatically scaled up to
a higher resolution, so any code doing drawing will automatically look
nicer. However, if you are supplying pixel-based data the scale
value can be used to determine whether to use a pixel resource
with higher resolution data.

The scale of a window may change during runtime, if this happens
a configure event will be sent to the toplevel window.
*/
func (self *TraitWindow) GetScaleFactor() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_window_get_scale_factor(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the #GdkScreen associated with a #GdkWindow.
*/
func (self *TraitWindow) GetScreen() (return__ *Screen) {
	var __cgo__return__ *C.GdkScreen
	__cgo__return__ = C.gdk_window_get_screen(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewScreenFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the event mask for @window corresponding to the device class specified
by @source.
*/
func (self *TraitWindow) GetSourceEvents(source C.GdkInputSource) (return__ C.GdkEventMask) {
	return__ = C.gdk_window_get_source_events(self.CPointer, source)
	return
}

/*
Gets the bitwise OR of the currently active window state flags,
from the #GdkWindowState enumeration.
*/
func (self *TraitWindow) GetState() (return__ C.GdkWindowState) {
	return__ = C.gdk_window_get_state(self.CPointer)
	return
}

/*
Returns %TRUE if the window is aware of the existence of multiple
devices.
*/
func (self *TraitWindow) GetSupportMultidevice() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_window_get_support_multidevice(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the toplevel window that’s an ancestor of @window.

Any window type but %GDK_WINDOW_CHILD is considered a
toplevel window, as is a %GDK_WINDOW_CHILD window that
has a root window as parent.

Note that you should use gdk_window_get_effective_toplevel() when
you want to get to a window’s toplevel as seen on screen, because
gdk_window_get_toplevel() will most likely not do what you expect
if there are offscreen windows in the hierarchy.
*/
func (self *TraitWindow) GetToplevel() (return__ *Window) {
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_window_get_toplevel(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
This function returns the type hint set for a window.
*/
func (self *TraitWindow) GetTypeHint() (return__ C.GdkWindowTypeHint) {
	return__ = C.gdk_window_get_type_hint(self.CPointer)
	return
}

/*
Transfers ownership of the update area from @window to the caller
of the function. That is, after calling this function, @window will
no longer have an invalid/dirty region; the update area is removed
from @window and handed to you. If a window has no update area,
gdk_window_get_update_area() returns %NULL. You are responsible for
calling cairo_region_destroy() on the returned region if it’s non-%NULL.
*/
func (self *TraitWindow) GetUpdateArea() (return__ *C.cairo_region_t) {
	return__ = C.gdk_window_get_update_area(self.CPointer)
	return
}

/*
Retrieves the user data for @window, which is normally the widget
that @window belongs to. See gdk_window_set_user_data().
*/
func (self *TraitWindow) GetUserData() (data unsafe.Pointer) {
	var __cgo__data C.gpointer
	C.gdk_window_get_user_data(self.CPointer, &__cgo__data)
	data = unsafe.Pointer(__cgo__data)
	return
}

/*
Computes the region of the @window that is potentially visible.
This does not necessarily take into account if the window is
obscured by other windows, but no area outside of this region
is visible.
*/
func (self *TraitWindow) GetVisibleRegion() (return__ *C.cairo_region_t) {
	return__ = C.gdk_window_get_visible_region(self.CPointer)
	return
}

/*
Gets the #GdkVisual describing the pixel format of @window.
*/
func (self *TraitWindow) GetVisual() (return__ *Visual) {
	var __cgo__return__ *C.GdkVisual
	__cgo__return__ = C.gdk_window_get_visual(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewVisualFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the width of the given @window.

On the X11 platform the returned size is the size reported in the
most-recently-processed configure event, rather than the current
size on the X server.
*/
func (self *TraitWindow) GetWidth() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.gdk_window_get_width(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the type of the window. See #GdkWindowType.
*/
func (self *TraitWindow) GetWindowType() (return__ C.GdkWindowType) {
	return__ = C.gdk_window_get_window_type(self.CPointer)
	return
}

/*
Checks whether the window has a native window or not. Note that
you can use gdk_window_ensure_native() if a native window is needed.
*/
func (self *TraitWindow) HasNative() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_window_has_native(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
For toplevel windows, withdraws them, so they will no longer be
known to the window manager; for all windows, unmaps them, so
they won’t be displayed. Normally done automatically as
part of gtk_widget_hide().
*/
func (self *TraitWindow) Hide() {
	C.gdk_window_hide(self.CPointer)
	return
}

/*
Asks to iconify (minimize) @window. The window manager may choose
to ignore the request, but normally will honor it. Using
gtk_window_iconify() is preferred, if you have a #GtkWindow widget.

This function only makes sense when @window is a toplevel window.
*/
func (self *TraitWindow) Iconify() {
	C.gdk_window_iconify(self.CPointer)
	return
}

/*
Like gdk_window_shape_combine_region(), but the shape applies
only to event handling. Mouse events which happen while
the pointer position corresponds to an unset bit in the
mask will be passed on the window below @window.

An input shape is typically used with RGBA windows.
The alpha channel of the window defines which pixels are
invisible and allows for nicely antialiased borders,
and the input shape controls where the window is
“clickable”.

On the X11 platform, this requires version 1.1 of the
shape extension.

On the Win32 platform, this functionality is not present and the
function does nothing.
*/
func (self *TraitWindow) InputShapeCombineRegion(shape_region *C.cairo_region_t, offset_x int, offset_y int) {
	C.gdk_window_input_shape_combine_region(self.CPointer, shape_region, C.gint(offset_x), C.gint(offset_y))
	return
}

/*
Adds @region to the update area for @window. The update area is the
region that needs to be redrawn, or “dirty region.” The call
gdk_window_process_updates() sends one or more expose events to the
window, which together cover the entire update area. An
application would normally redraw the contents of @window in
response to those expose events.

GDK will call gdk_window_process_all_updates() on your behalf
whenever your program returns to the main loop and becomes idle, so
normally there’s no need to do that manually, you just need to
invalidate regions that you know should be redrawn.

The @child_func parameter controls whether the region of
each child window that intersects @region will also be invalidated.
Only children for which @child_func returns TRUE will have the area
invalidated.
*/
func (self *TraitWindow) InvalidateMaybeRecurse(region *C.cairo_region_t, child_func C.GdkWindowChildFunc, user_data unsafe.Pointer) {
	C.gdk_window_invalidate_maybe_recurse(self.CPointer, region, child_func, (C.gpointer)(user_data))
	return
}

/*
A convenience wrapper around gdk_window_invalidate_region() which
invalidates a rectangular region. See
gdk_window_invalidate_region() for details.
*/
func (self *TraitWindow) InvalidateRect(rect *C.GdkRectangle, invalidate_children bool) {
	__cgo__invalidate_children := C.gboolean(0)
	if invalidate_children {
		__cgo__invalidate_children = C.gboolean(1)
	}
	C.gdk_window_invalidate_rect(self.CPointer, rect, __cgo__invalidate_children)
	return
}

/*
Adds @region to the update area for @window. The update area is the
region that needs to be redrawn, or “dirty region.” The call
gdk_window_process_updates() sends one or more expose events to the
window, which together cover the entire update area. An
application would normally redraw the contents of @window in
response to those expose events.

GDK will call gdk_window_process_all_updates() on your behalf
whenever your program returns to the main loop and becomes idle, so
normally there’s no need to do that manually, you just need to
invalidate regions that you know should be redrawn.

The @invalidate_children parameter controls whether the region of
each child window that intersects @region will also be invalidated.
If %FALSE, then the update area for child windows will remain
unaffected. See gdk_window_invalidate_maybe_recurse if you need
fine grained control over which children are invalidated.
*/
func (self *TraitWindow) InvalidateRegion(region *C.cairo_region_t, invalidate_children bool) {
	__cgo__invalidate_children := C.gboolean(0)
	if invalidate_children {
		__cgo__invalidate_children = C.gboolean(1)
	}
	C.gdk_window_invalidate_region(self.CPointer, region, __cgo__invalidate_children)
	return
}

/*
Check to see if a window is destroyed..
*/
func (self *TraitWindow) IsDestroyed() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_window_is_destroyed(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether or not the window is an input only window.
*/
func (self *TraitWindow) IsInputOnly() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_window_is_input_only(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether or not the window is shaped.
*/
func (self *TraitWindow) IsShaped() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_window_is_shaped(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Check if the window and all ancestors of the window are
mapped. (This is not necessarily "viewable" in the X sense, since
we only check as far as we have GDK window parents, not to the root
window.)
*/
func (self *TraitWindow) IsViewable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_window_is_viewable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks whether the window has been mapped (with gdk_window_show() or
gdk_window_show_unraised()).
*/
func (self *TraitWindow) IsVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_window_is_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Lowers @window to the bottom of the Z-order (stacking order), so that
other windows with the same parent window appear above @window.
This is true whether or not the other windows are visible.

If @window is a toplevel, the window manager may choose to deny the
request to move the window in the Z-order, gdk_window_lower() only
requests the restack, does not guarantee it.

Note that gdk_window_show() raises the window again, so don’t call this
function before gdk_window_show(). (Try gdk_window_show_unraised().)
*/
func (self *TraitWindow) Lower() {
	C.gdk_window_lower(self.CPointer)
	return
}

/*
Maximizes the window. If the window was already maximized, then
this function does nothing.

On X11, asks the window manager to maximize @window, if the window
manager supports this operation. Not all window managers support
this, and some deliberately ignore it or don’t have a concept of
“maximized”; so you can’t rely on the maximization actually
happening. But it will happen with most standard window managers,
and GDK makes a best effort to get it to happen.

On Windows, reliably maximizes the window.
*/
func (self *TraitWindow) Maximize() {
	C.gdk_window_maximize(self.CPointer)
	return
}

/*
Merges the input shape masks for any child windows into the
input shape mask for @window. i.e. the union of all input masks
for @window and its children will become the new input mask
for @window. See gdk_window_input_shape_combine_region().

This function is distinct from gdk_window_set_child_input_shapes()
because it includes @window’s input shape mask in the set of
shapes to be merged.
*/
func (self *TraitWindow) MergeChildInputShapes() {
	C.gdk_window_merge_child_input_shapes(self.CPointer)
	return
}

/*
Merges the shape masks for any child windows into the
shape mask for @window. i.e. the union of all masks
for @window and its children will become the new mask
for @window. See gdk_window_shape_combine_region().

This function is distinct from gdk_window_set_child_shapes()
because it includes @window’s shape mask in the set of shapes to
be merged.
*/
func (self *TraitWindow) MergeChildShapes() {
	C.gdk_window_merge_child_shapes(self.CPointer)
	return
}

/*
Repositions a window relative to its parent window.
For toplevel windows, window managers may ignore or modify the move;
you should probably use gtk_window_move() on a #GtkWindow widget
anyway, instead of using GDK functions. For child windows,
the move will reliably succeed.

If you’re also planning to resize the window, use gdk_window_move_resize()
to both move and resize simultaneously, for a nicer visual effect.
*/
func (self *TraitWindow) Move(x int, y int) {
	C.gdk_window_move(self.CPointer, C.gint(x), C.gint(y))
	return
}

/*
Move the part of @window indicated by @region by @dy pixels in the Y
direction and @dx pixels in the X direction. The portions of @region
that not covered by the new position of @region are invalidated.

Child windows are not moved.
*/
func (self *TraitWindow) MoveRegion(region *C.cairo_region_t, dx int, dy int) {
	C.gdk_window_move_region(self.CPointer, region, C.gint(dx), C.gint(dy))
	return
}

/*
Equivalent to calling gdk_window_move() and gdk_window_resize(),
except that both operations are performed at once, avoiding strange
visual effects. (i.e. the user may be able to see the window first
move, then resize, if you don’t use gdk_window_move_resize().)
*/
func (self *TraitWindow) MoveResize(x int, y int, width int, height int) {
	C.gdk_window_move_resize(self.CPointer, C.gint(x), C.gint(y), C.gint(width), C.gint(height))
	return
}

/*
Like gdk_window_get_children(), but does not copy the list of
children, so the list does not need to be freed.
*/
func (self *TraitWindow) PeekChildren() (return__ *C.GList) {
	return__ = C.gdk_window_peek_children(self.CPointer)
	return
}

/*
Sends one or more expose events to @window. The areas in each
expose event will cover the entire update area for the window (see
gdk_window_invalidate_region() for details). Normally GDK calls
gdk_window_process_all_updates() on your behalf, so there’s no
need to call this function unless you want to force expose events
to be delivered immediately and synchronously (vs. the usual
case, where GDK delivers them in an idle handler). Occasionally
this is useful to produce nicer scrolling behavior, for example.
*/
func (self *TraitWindow) ProcessUpdates(update_children bool) {
	__cgo__update_children := C.gboolean(0)
	if update_children {
		__cgo__update_children = C.gboolean(1)
	}
	C.gdk_window_process_updates(self.CPointer, __cgo__update_children)
	return
}

/*
Raises @window to the top of the Z-order (stacking order), so that
other windows with the same parent window appear below @window.
This is true whether or not the windows are visible.

If @window is a toplevel, the window manager may choose to deny the
request to move the window in the Z-order, gdk_window_raise() only
requests the restack, does not guarantee it.
*/
func (self *TraitWindow) Raise() {
	C.gdk_window_raise(self.CPointer)
	return
}

/*
Registers a window as a potential drop destination.
*/
func (self *TraitWindow) RegisterDnd() {
	C.gdk_window_register_dnd(self.CPointer)
	return
}

/*
Remove a filter previously added with gdk_window_add_filter().
*/
func (self *TraitWindow) RemoveFilter(function C.GdkFilterFunc, data unsafe.Pointer) {
	C.gdk_window_remove_filter(self.CPointer, function, (C.gpointer)(data))
	return
}

/*
Reparents @window into the given @new_parent. The window being
reparented will be unmapped as a side effect.
*/
func (self *TraitWindow) Reparent(new_parent IsWindow, x int, y int) {
	C.gdk_window_reparent(self.CPointer, new_parent.GetWindowPointer(), C.gint(x), C.gint(y))
	return
}

/*
Resizes @window; for toplevel windows, asks the window manager to resize
the window. The window manager may not allow the resize. When using GTK+,
use gtk_window_resize() instead of this low-level GDK function.

Windows may not be resized below 1x1.

If you’re also planning to move the window, use gdk_window_move_resize()
to both move and resize simultaneously, for a nicer visual effect.
*/
func (self *TraitWindow) Resize(width int, height int) {
	C.gdk_window_resize(self.CPointer, C.gint(width), C.gint(height))
	return
}

/*
Changes the position of  @window in the Z-order (stacking order), so that
it is above @sibling (if @above is %TRUE) or below @sibling (if @above is
%FALSE).

If @sibling is %NULL, then this either raises (if @above is %TRUE) or
lowers the window.

If @window is a toplevel, the window manager may choose to deny the
request to move the window in the Z-order, gdk_window_restack() only
requests the restack, does not guarantee it.
*/
func (self *TraitWindow) Restack(sibling IsWindow, above bool) {
	__cgo__above := C.gboolean(0)
	if above {
		__cgo__above = C.gboolean(1)
	}
	C.gdk_window_restack(self.CPointer, sibling.GetWindowPointer(), __cgo__above)
	return
}

/*
Scroll the contents of @window, both pixels and children, by the
given amount. @window itself does not move. Portions of the window
that the scroll operation brings in from offscreen areas are
invalidated. The invalidated region may be bigger than what would
strictly be necessary.

For X11, a minimum area will be invalidated if the window has no
subwindows, or if the edges of the window’s parent do not extend
beyond the edges of the window. In other cases, a multi-step process
is used to scroll the window which may produce temporary visual
artifacts and unnecessary invalidations.
*/
func (self *TraitWindow) Scroll(dx int, dy int) {
	C.gdk_window_scroll(self.CPointer, C.gint(dx), C.gint(dy))
	return
}

/*
Setting @accept_focus to %FALSE hints the desktop environment that the
window doesn’t want to receive input focus.

On X, it is the responsibility of the window manager to interpret this
hint. ICCCM-compliant window manager usually respect it.
*/
func (self *TraitWindow) SetAcceptFocus(accept_focus bool) {
	__cgo__accept_focus := C.gboolean(0)
	if accept_focus {
		__cgo__accept_focus = C.gboolean(1)
	}
	C.gdk_window_set_accept_focus(self.CPointer, __cgo__accept_focus)
	return
}

// gdk_window_set_background is not generated due to deprecation attr

/*
Sets the background of @window.

A background of %NULL means that the window will inherit its
background from its parent window.

The windowing system will normally fill a window with its background
when the window is obscured then exposed.
*/
func (self *TraitWindow) SetBackgroundPattern(pattern *C.cairo_pattern_t) {
	C.gdk_window_set_background_pattern(self.CPointer, pattern)
	return
}

/*
Sets the background color of @window.

See also gdk_window_set_background_pattern().
*/
func (self *TraitWindow) SetBackgroundRgba(rgba *C.GdkRGBA) {
	C.gdk_window_set_background_rgba(self.CPointer, rgba)
	return
}

/*
Sets the input shape mask of @window to the union of input shape masks
for all children of @window, ignoring the input shape mask of @window
itself. Contrast with gdk_window_merge_child_input_shapes() which includes
the input shape mask of @window in the masks to be merged.
*/
func (self *TraitWindow) SetChildInputShapes() {
	C.gdk_window_set_child_input_shapes(self.CPointer)
	return
}

/*
Sets the shape mask of @window to the union of shape masks
for all children of @window, ignoring the shape mask of @window
itself. Contrast with gdk_window_merge_child_shapes() which includes
the shape mask of @window in the masks to be merged.
*/
func (self *TraitWindow) SetChildShapes() {
	C.gdk_window_set_child_shapes(self.CPointer)
	return
}

/*
Sets a #GdkWindow as composited, or unsets it. Composited
windows do not automatically have their contents drawn to
the screen. Drawing is redirected to an offscreen buffer
and an expose event is emitted on the parent of the composited
window. It is the responsibility of the parent’s expose handler
to manually merge the off-screen content onto the screen in
whatever way it sees fit.

It only makes sense for child windows to be composited; see
gdk_window_set_opacity() if you need translucent toplevel
windows.

An additional effect of this call is that the area of this
window is no longer clipped from regions marked for
invalidation on its parent. Draws done on the parent
window are also no longer clipped by the child.

This call is only supported on some systems (currently,
only X11 with new enough Xcomposite and Xdamage extensions).
You must call gdk_display_supports_composite() to check if
setting a window as composited is supported before
attempting to do so.
*/
func (self *TraitWindow) SetComposited(composited bool) {
	__cgo__composited := C.gboolean(0)
	if composited {
		__cgo__composited = C.gboolean(1)
	}
	C.gdk_window_set_composited(self.CPointer, __cgo__composited)
	return
}

/*
Sets the default mouse pointer for a #GdkWindow. Use gdk_cursor_new_for_display()
or gdk_cursor_new_from_pixbuf() to create the cursor. To make the cursor
invisible, use %GDK_BLANK_CURSOR. Passing %NULL for the @cursor argument
to gdk_window_set_cursor() means that @window will use the cursor of its
parent window. Most windows should use this default.
*/
func (self *TraitWindow) SetCursor(cursor IsCursor) {
	C.gdk_window_set_cursor(self.CPointer, cursor.GetCursorPointer())
	return
}

/*
“Decorations” are the features the window manager adds to a toplevel #GdkWindow.
This function sets the traditional Motif window manager hints that tell the
window manager which decorations you would like your window to have.
Usually you should use gtk_window_set_decorated() on a #GtkWindow instead of
using the GDK function directly.

The @decorations argument is the logical OR of the fields in
the #GdkWMDecoration enumeration. If #GDK_DECOR_ALL is included in the
mask, the other bits indicate which decorations should be turned off.
If #GDK_DECOR_ALL is not included, then the other bits indicate
which decorations should be turned on.

Most window managers honor a decorations hint of 0 to disable all decorations,
but very few honor all possible combinations of bits.
*/
func (self *TraitWindow) SetDecorations(decorations C.GdkWMDecoration) {
	C.gdk_window_set_decorations(self.CPointer, decorations)
	return
}

/*
Sets a specific #GdkCursor for a given device when it gets inside @window.
Use gdk_cursor_new_for_display() or gdk_cursor_new_from_pixbuf() to create
the cursor. To make the cursor invisible, use %GDK_BLANK_CURSOR. Passing
%NULL for the @cursor argument to gdk_window_set_cursor() means that
@window will use the cursor of its parent window. Most windows should
use this default.
*/
func (self *TraitWindow) SetDeviceCursor(device IsDevice, cursor IsCursor) {
	C.gdk_window_set_device_cursor(self.CPointer, device.GetDevicePointer(), cursor.GetCursorPointer())
	return
}

/*
Sets the event mask for a given device (Normally a floating device, not
attached to any visible pointer) to @window. For example, an event mask
including #GDK_BUTTON_PRESS_MASK means the window should report button
press events. The event mask is the bitwise OR of values from the
#GdkEventMask enumeration.
*/
func (self *TraitWindow) SetDeviceEvents(device IsDevice, event_mask C.GdkEventMask) {
	C.gdk_window_set_device_events(self.CPointer, device.GetDevicePointer(), event_mask)
	return
}

/*
Determines whether or not extra unprocessed motion events in
the event queue can be discarded. If %TRUE only the most recent
event will be delivered.

Some types of applications, e.g. paint programs, need to see all
motion events and will benefit from turning off event compression.

By default, event compression is enabled.
*/
func (self *TraitWindow) SetEventCompression(event_compression bool) {
	__cgo__event_compression := C.gboolean(0)
	if event_compression {
		__cgo__event_compression = C.gboolean(1)
	}
	C.gdk_window_set_event_compression(self.CPointer, __cgo__event_compression)
	return
}

/*
The event mask for a window determines which events will be reported
for that window from all master input devices. For example, an event mask
including #GDK_BUTTON_PRESS_MASK means the window should report button
press events. The event mask is the bitwise OR of values from the
#GdkEventMask enumeration.
*/
func (self *TraitWindow) SetEvents(event_mask C.GdkEventMask) {
	C.gdk_window_set_events(self.CPointer, event_mask)
	return
}

/*
Setting @focus_on_map to %FALSE hints the desktop environment that the
window doesn’t want to receive input focus when it is mapped.
focus_on_map should be turned off for windows that aren’t triggered
interactively (such as popups from network activity).

On X, it is the responsibility of the window manager to interpret
this hint. Window managers following the freedesktop.org window
manager extension specification should respect it.
*/
func (self *TraitWindow) SetFocusOnMap(focus_on_map bool) {
	__cgo__focus_on_map := C.gboolean(0)
	if focus_on_map {
		__cgo__focus_on_map = C.gboolean(1)
	}
	C.gdk_window_set_focus_on_map(self.CPointer, __cgo__focus_on_map)
	return
}

/*
Specifies whether the @window should span over all monitors (in a multi-head
setup) or only the current monitor when in fullscreen mode.

The @mode argument is from the #GdkFullscreenMode enumeration.
If #GDK_FULLSCREEN_ON_ALL_MONITORS is specified, the fullscreen @window will
span over all monitors from the #GdkScreen.

On X11, searches through the list of monitors from the #GdkScreen the ones
which delimit the 4 edges of the entire #GdkScreen and will ask the window
manager to span the @window over these monitors.

If the XINERAMA extension is not available or not usable, this function
has no effect.

Not all window managers support this, so you can’t rely on the fullscreen
window to span over the multiple monitors when #GDK_FULLSCREEN_ON_ALL_MONITORS
is specified.
*/
func (self *TraitWindow) SetFullscreenMode(mode C.GdkFullscreenMode) {
	C.gdk_window_set_fullscreen_mode(self.CPointer, mode)
	return
}

/*
Sets hints about the window management functions to make available
via buttons on the window frame.

On the X backend, this function sets the traditional Motif window
manager hint for this purpose. However, few window managers do
anything reliable or interesting with this hint. Many ignore it
entirely.

The @functions argument is the logical OR of values from the
#GdkWMFunction enumeration. If the bitmask includes #GDK_FUNC_ALL,
then the other bits indicate which functions to disable; if
it doesn’t include #GDK_FUNC_ALL, it indicates which functions to
enable.
*/
func (self *TraitWindow) SetFunctions(functions C.GdkWMFunction) {
	C.gdk_window_set_functions(self.CPointer, functions)
	return
}

/*
Sets the geometry hints for @window. Hints flagged in @geom_mask
are set, hints not flagged in @geom_mask are unset.
To unset all hints, use a @geom_mask of 0 and a @geometry of %NULL.

This function provides hints to the windowing system about
acceptable sizes for a toplevel window. The purpose of
this is to constrain user resizing, but the windowing system
will typically  (but is not required to) also constrain the
current size of the window to the provided values and
constrain programatic resizing via gdk_window_resize() or
gdk_window_move_resize().

Note that on X11, this effect has no effect on windows
of type %GDK_WINDOW_TEMP or windows where override redirect
has been turned on via gdk_window_set_override_redirect()
since these windows are not resizable by the user.

Since you can’t count on the windowing system doing the
constraints for programmatic resizes, you should generally
call gdk_window_constrain_size() yourself to determine
appropriate sizes.
*/
func (self *TraitWindow) SetGeometryHints(geometry *C.GdkGeometry, geom_mask C.GdkWindowHints) {
	C.gdk_window_set_geometry_hints(self.CPointer, geometry, geom_mask)
	return
}

/*
Sets the group leader window for @window. By default,
GDK sets the group leader for all toplevel windows
to a global window implicitly created by GDK. With this function
you can override this default.

The group leader window allows the window manager to distinguish
all windows that belong to a single application. It may for example
allow users to minimize/unminimize all windows belonging to an
application at once. You should only set a non-default group window
if your application pretends to be multiple applications.
*/
func (self *TraitWindow) SetGroup(leader IsWindow) {
	C.gdk_window_set_group(self.CPointer, leader.GetWindowPointer())
	return
}

/*
Sets a list of icons for the window. One of these will be used
to represent the window when it has been iconified. The icon is
usually shown in an icon box or some sort of task bar. Which icon
size is shown depends on the window manager. The window manager
can scale the icon  but setting several size icons can give better
image quality since the window manager may only need to scale the
icon by a small amount or not at all.
*/
func (self *TraitWindow) SetIconList(pixbufs *C.GList) {
	C.gdk_window_set_icon_list(self.CPointer, pixbufs)
	return
}

/*
Windows may have a name used while minimized, distinct from the
name they display in their titlebar. Most of the time this is a bad
idea from a user interface standpoint. But you can set such a name
with this function, if you like.

After calling this with a non-%NULL @name, calls to gdk_window_set_title()
will not update the icon title.

Using %NULL for @name unsets the icon title; further calls to
gdk_window_set_title() will again update the icon title as well.
*/
func (self *TraitWindow) SetIconName(name string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gdk_window_set_icon_name(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Registers an invalidate handler for a specific window. This
will get called whenever a region in the window or its children
is invalidated.

This can be used to record the invalidated region, which is
useful if you are keeping an offscreen copy of some region
and want to keep it up to date. You can also modify the
invalidated region in case you’re doing some effect where
e.g. a child widget appears in multiple places.
*/
func (self *TraitWindow) SetInvalidateHandler(handler C.GdkWindowInvalidateHandlerFunc) {
	C.gdk_window_set_invalidate_handler(self.CPointer, handler)
	return
}

/*
Set if @window must be kept above other windows. If the
window was already above, then this function does nothing.

On X11, asks the window manager to keep @window above, if the window
manager supports this operation. Not all window managers support
this, and some deliberately ignore it or don’t have a concept of
“keep above”; so you can’t rely on the window being kept above.
But it will happen with most standard window managers,
and GDK makes a best effort to get it to happen.
*/
func (self *TraitWindow) SetKeepAbove(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gdk_window_set_keep_above(self.CPointer, __cgo__setting)
	return
}

/*
Set if @window must be kept below other windows. If the
window was already below, then this function does nothing.

On X11, asks the window manager to keep @window below, if the window
manager supports this operation. Not all window managers support
this, and some deliberately ignore it or don’t have a concept of
“keep below”; so you can’t rely on the window being kept below.
But it will happen with most standard window managers,
and GDK makes a best effort to get it to happen.
*/
func (self *TraitWindow) SetKeepBelow(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gdk_window_set_keep_below(self.CPointer, __cgo__setting)
	return
}

/*
The application can use this hint to tell the window manager
that a certain window has modal behaviour. The window manager
can use this information to handle modal windows in a special
way.

You should only use this on windows for which you have
previously called gdk_window_set_transient_for()
*/
func (self *TraitWindow) SetModalHint(modal bool) {
	__cgo__modal := C.gboolean(0)
	if modal {
		__cgo__modal = C.gboolean(1)
	}
	C.gdk_window_set_modal_hint(self.CPointer, __cgo__modal)
	return
}

/*
Set @window to render as partially transparent,
with opacity 0 being fully transparent and 1 fully opaque. (Values
of the opacity parameter are clamped to the [0,1] range.)

For toplevel windows this depends on support from the windowing system
that may not always be there. For instance, On X11, this works only on
X screens with a compositing manager running.

For child windows this function only works for non-native windows.

For setting up per-pixel alpha topelevels, see gdk_screen_get_rgba_visual(),
and for non-toplevels, see gdk_window_set_composited().

Support for non-toplevel windows was added in 3.8.
*/
func (self *TraitWindow) SetOpacity(opacity float64) {
	C.gdk_window_set_opacity(self.CPointer, C.gdouble(opacity))
	return
}

/*
For optimizization purposes, compositing window managers may
like to not draw obscured regions of windows, or turn off blending
during for these regions. With RGB windows with no transparency,
this is just the shape of the window, but with ARGB32 windows, the
compositor does not know what regions of the window are transparent
or not.

This function only works for toplevel windows.

GTK+ will automatically update this property automatically if
the @window background is opaque, as we know where the opaque regions
are. If your window background is not opaque, please update this
property in your #GtkWidget::style-updated handler.
*/
func (self *TraitWindow) SetOpaqueRegion(region *C.cairo_region_t) {
	C.gdk_window_set_opaque_region(self.CPointer, region)
	return
}

/*
An override redirect window is not under the control of the window manager.
This means it won’t have a titlebar, won’t be minimizable, etc. - it will
be entirely under the control of the application. The window manager
can’t see the override redirect window at all.

Override redirect should only be used for short-lived temporary
windows, such as popup menus. #GtkMenu uses an override redirect
window in its implementation, for example.
*/
func (self *TraitWindow) SetOverrideRedirect(override_redirect bool) {
	__cgo__override_redirect := C.gboolean(0)
	if override_redirect {
		__cgo__override_redirect = C.gboolean(1)
	}
	C.gdk_window_set_override_redirect(self.CPointer, __cgo__override_redirect)
	return
}

/*
When using GTK+, typically you should use gtk_window_set_role() instead
of this low-level function.

The window manager and session manager use a window’s role to
distinguish it from other kinds of window in the same application.
When an application is restarted after being saved in a previous
session, all windows with the same title and role are treated as
interchangeable.  So if you have two windows with the same title
that should be distinguished for session management purposes, you
should set the role on those windows. It doesn’t matter what string
you use for the role, as long as you have a different role for each
non-interchangeable kind of window.
*/
func (self *TraitWindow) SetRole(role string) {
	__cgo__role := (*C.gchar)(unsafe.Pointer(C.CString(role)))
	C.gdk_window_set_role(self.CPointer, __cgo__role)
	C.free(unsafe.Pointer(__cgo__role))
	return
}

/*
Newer GTK+ windows using client-side decorations use extra geometry
around their frames for effects like shadows and invisible borders.
Window managers that want to maximize windows or snap to edges need
to know where the extents of the actual frame lie, so that users
don’t feel like windows are snapping against random invisible edges.

Note that this property is automatically updated by GTK+, so this
function should only be used by applications which do not use GTK+
to create toplevel windows.
*/
func (self *TraitWindow) SetShadowWidth(left int, right int, top int, bottom int) {
	C.gdk_window_set_shadow_width(self.CPointer, C.gint(left), C.gint(right), C.gint(top), C.gint(bottom))
	return
}

/*
Toggles whether a window should appear in a pager (workspace
switcher, or other desktop utility program that displays a small
thumbnail representation of the windows on the desktop). If a
window’s semantic type as specified with gdk_window_set_type_hint()
already fully describes the window, this function should
not be called in addition, instead you should
allow the window to be treated according to standard policy for
its semantic type.
*/
func (self *TraitWindow) SetSkipPagerHint(skips_pager bool) {
	__cgo__skips_pager := C.gboolean(0)
	if skips_pager {
		__cgo__skips_pager = C.gboolean(1)
	}
	C.gdk_window_set_skip_pager_hint(self.CPointer, __cgo__skips_pager)
	return
}

/*
Toggles whether a window should appear in a task list or window
list. If a window’s semantic type as specified with
gdk_window_set_type_hint() already fully describes the window, this
function should not be called in addition,
instead you should allow the window to be treated according to
standard policy for its semantic type.
*/
func (self *TraitWindow) SetSkipTaskbarHint(skips_taskbar bool) {
	__cgo__skips_taskbar := C.gboolean(0)
	if skips_taskbar {
		__cgo__skips_taskbar = C.gboolean(1)
	}
	C.gdk_window_set_skip_taskbar_hint(self.CPointer, __cgo__skips_taskbar)
	return
}

/*
Sets the event mask for any floating device (i.e. not attached to any
visible pointer) that has the source defined as @source. This event
mask will be applied both to currently existing, newly added devices
after this call, and devices being attached/detached.
*/
func (self *TraitWindow) SetSourceEvents(source C.GdkInputSource, event_mask C.GdkEventMask) {
	C.gdk_window_set_source_events(self.CPointer, source, event_mask)
	return
}

/*
When using GTK+, typically you should use gtk_window_set_startup_id()
instead of this low-level function.
*/
func (self *TraitWindow) SetStartupId(startup_id string) {
	__cgo__startup_id := (*C.gchar)(unsafe.Pointer(C.CString(startup_id)))
	C.gdk_window_set_startup_id(self.CPointer, __cgo__startup_id)
	C.free(unsafe.Pointer(__cgo__startup_id))
	return
}

/*
Set the bit gravity of the given window to static, and flag it so
all children get static subwindow gravity. This is used if you are
implementing scary features that involve deep knowledge of the
windowing system. Don’t worry about it unless you have to.
*/
func (self *TraitWindow) SetStaticGravities(use_static bool) (return__ bool) {
	__cgo__use_static := C.gboolean(0)
	if use_static {
		__cgo__use_static = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_window_set_static_gravities(self.CPointer, __cgo__use_static)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This function will enable multidevice features in @window.

Multidevice aware windows will need to handle properly multiple,
per device enter/leave events, device grabs and grab ownerships.
*/
func (self *TraitWindow) SetSupportMultidevice(support_multidevice bool) {
	__cgo__support_multidevice := C.gboolean(0)
	if support_multidevice {
		__cgo__support_multidevice = C.gboolean(1)
	}
	C.gdk_window_set_support_multidevice(self.CPointer, __cgo__support_multidevice)
	return
}

/*
Sets the title of a toplevel window, to be displayed in the titlebar.
If you haven’t explicitly set the icon name for the window
(using gdk_window_set_icon_name()), the icon name will be set to
@title as well. @title must be in UTF-8 encoding (as with all
user-readable strings in GDK/GTK+). @title may not be %NULL.
*/
func (self *TraitWindow) SetTitle(title string) {
	__cgo__title := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	C.gdk_window_set_title(self.CPointer, __cgo__title)
	C.free(unsafe.Pointer(__cgo__title))
	return
}

/*
Indicates to the window manager that @window is a transient dialog
associated with the application window @parent. This allows the
window manager to do things like center @window on @parent and
keep @window above @parent.

See gtk_window_set_transient_for() if you’re using #GtkWindow or
#GtkDialog.
*/
func (self *TraitWindow) SetTransientFor(parent IsWindow) {
	C.gdk_window_set_transient_for(self.CPointer, parent.GetWindowPointer())
	return
}

/*
The application can use this call to provide a hint to the window
manager about the functionality of a window. The window manager
can use this information when determining the decoration and behaviour
of the window.

The hint must be set before the window is mapped.
*/
func (self *TraitWindow) SetTypeHint(hint C.GdkWindowTypeHint) {
	C.gdk_window_set_type_hint(self.CPointer, hint)
	return
}

/*
Toggles whether a window needs the user's
urgent attention.
*/
func (self *TraitWindow) SetUrgencyHint(urgent bool) {
	__cgo__urgent := C.gboolean(0)
	if urgent {
		__cgo__urgent = C.gboolean(1)
	}
	C.gdk_window_set_urgency_hint(self.CPointer, __cgo__urgent)
	return
}

/*
For most purposes this function is deprecated in favor of
g_object_set_data(). However, for historical reasons GTK+ stores
the #GtkWidget that owns a #GdkWindow as user data on the
#GdkWindow. So, custom widget implementations should use
this function for that. If GTK+ receives an event for a #GdkWindow,
and the user data for the window is non-%NULL, GTK+ will assume the
user data is a #GtkWidget, and forward the event to that widget.
*/
func (self *TraitWindow) SetUserData(user_data C.gpointer) {
	C.gdk_window_set_user_data(self.CPointer, user_data)
	return
}

/*
Makes pixels in @window outside @shape_region be transparent,
so that the window may be nonrectangular.

If @shape_region is %NULL, the shape will be unset, so the whole
window will be opaque again. @offset_x and @offset_y are ignored
if @shape_region is %NULL.

On the X11 platform, this uses an X server extension which is
widely available on most common platforms, but not available on
very old X servers, and occasionally the implementation will be
buggy. On servers without the shape extension, this function
will do nothing.

This function works on both toplevel and child windows.
*/
func (self *TraitWindow) ShapeCombineRegion(shape_region *C.cairo_region_t, offset_x int, offset_y int) {
	C.gdk_window_shape_combine_region(self.CPointer, shape_region, C.gint(offset_x), C.gint(offset_y))
	return
}

/*
Like gdk_window_show_unraised(), but also raises the window to the
top of the window stack (moves the window to the front of the
Z-order).

This function maps a window so it’s visible onscreen. Its opposite
is gdk_window_hide().

When implementing a #GtkWidget, you should call this function on the widget's
#GdkWindow as part of the “map” method.
*/
func (self *TraitWindow) Show() {
	C.gdk_window_show(self.CPointer)
	return
}

/*
Shows a #GdkWindow onscreen, but does not modify its stacking
order. In contrast, gdk_window_show() will raise the window
to the top of the window stack.

On the X11 platform, in Xlib terms, this function calls
XMapWindow() (it also updates some internal GDK state, which means
that you can’t really use XMapWindow() directly on a GDK window).
*/
func (self *TraitWindow) ShowUnraised() {
	C.gdk_window_show_unraised(self.CPointer)
	return
}

/*
“Pins” a window such that it’s on all workspaces and does not scroll
with viewports, for window managers that have scrollable viewports.
(When using #GtkWindow, gtk_window_stick() may be more useful.)

On the X11 platform, this function depends on window manager
support, so may have no effect with many window managers. However,
GDK will do the best it can to convince the window manager to stick
the window. For window managers that don’t support this operation,
there’s nothing you can do to force it to happen.
*/
func (self *TraitWindow) Stick() {
	C.gdk_window_stick(self.CPointer)
	return
}

/*
Thaws a window frozen with
gdk_window_freeze_toplevel_updates_libgtk_only().

This function is not part of the GDK public API and is only
for use by GTK+.
*/
func (self *TraitWindow) ThawToplevelUpdatesLibgtkOnly() {
	C.gdk_window_thaw_toplevel_updates_libgtk_only(self.CPointer)
	return
}

/*
Thaws a window frozen with gdk_window_freeze_updates().
*/
func (self *TraitWindow) ThawUpdates() {
	C.gdk_window_thaw_updates(self.CPointer)
	return
}

/*
Moves the window out of fullscreen mode. If the window was not
fullscreen, does nothing.

On X11, asks the window manager to move @window out of the fullscreen
state, if the window manager supports this operation. Not all
window managers support this, and some deliberately ignore it or
don’t have a concept of “fullscreen”; so you can’t rely on the
unfullscreenification actually happening. But it will happen with
most standard window managers, and GDK makes a best effort to get
it to happen.
*/
func (self *TraitWindow) Unfullscreen() {
	C.gdk_window_unfullscreen(self.CPointer)
	return
}

/*
Unmaximizes the window. If the window wasn’t maximized, then this
function does nothing.

On X11, asks the window manager to unmaximize @window, if the
window manager supports this operation. Not all window managers
support this, and some deliberately ignore it or don’t have a
concept of “maximized”; so you can’t rely on the unmaximization
actually happening. But it will happen with most standard window
managers, and GDK makes a best effort to get it to happen.

On Windows, reliably unmaximizes the window.
*/
func (self *TraitWindow) Unmaximize() {
	C.gdk_window_unmaximize(self.CPointer)
	return
}

/*
Reverse operation for gdk_window_stick(); see gdk_window_stick(),
and gtk_window_unstick().
*/
func (self *TraitWindow) Unstick() {
	C.gdk_window_unstick(self.CPointer)
	return
}

/*
Withdraws a window (unmaps it and asks the window manager to forget about it).
This function is not really useful as gdk_window_hide() automatically
withdraws toplevel windows before hiding them.
*/
func (self *TraitWindow) Withdraw() {
	C.gdk_window_withdraw(self.CPointer)
	return
}
