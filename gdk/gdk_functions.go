package gdk

/*
#include <gdk/gdk.h>
#include <gdk/gdkprivate.h>
#include <glib.h>
#include <stdlib.h>
#cgo pkg-config: gdk-3.0
*/
import "C"

import "github.com/reusee/ggir/gobject"
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

/*
Appends gdk option entries to the passed in option group. This is
not public API and must not be used by applications.
*/
func AddOptionEntriesLibgtkOnly(group *C.GOptionGroup) {
	C.gdk_add_option_entries_libgtk_only(group)
	return
}

/*
Finds or creates an atom corresponding to a given string.
*/
func AtomIntern(atom_name string, only_if_exists bool) (return__ C.GdkAtom) {
	__cgo__atom_name := (*C.gchar)(unsafe.Pointer(C.CString(atom_name)))
	__cgo__only_if_exists := C.gboolean(0)
	if only_if_exists {
		__cgo__only_if_exists = C.gboolean(1)
	}
	return__ = C.gdk_atom_intern(__cgo__atom_name, __cgo__only_if_exists)
	C.free(unsafe.Pointer(__cgo__atom_name))
	return
}

/*
Finds or creates an atom corresponding to a given string.

Note that this function is identical to gdk_atom_intern() except
that if a new #GdkAtom is created the string itself is used rather
than a copy. This saves memory, but can only be used if the string
will always exist. It can be used with statically
allocated strings in the main program, but not with statically
allocated memory in dynamically loaded modules, if you expect to
ever unload the module again (e.g. do not use this function in
GTK+ theme engines).
*/
func AtomInternStaticString(atom_name string) (return__ C.GdkAtom) {
	__cgo__atom_name := (*C.gchar)(unsafe.Pointer(C.CString(atom_name)))
	return__ = C.gdk_atom_intern_static_string(__cgo__atom_name)
	C.free(unsafe.Pointer(__cgo__atom_name))
	return
}

/*
Emits a short beep on the default display.
*/
func Beep() {
	C.gdk_beep()
	return
}

/*
Creates a Cairo context for drawing to @window.

Note that calling cairo_reset_clip() on the resulting #cairo_t will
produce undefined results, so avoid it at all costs.
*/
func CairoCreate(window IsWindow) (return__ *C.cairo_t) {
	var __cgo__window *C.GdkWindow
	if window != nil {
		__cgo__window = window.GetWindowPointer()
	}
	return__ = C.gdk_cairo_create(__cgo__window)
	return
}

/*
This is a convenience function around cairo_clip_extents().
It rounds the clip extents to integer coordinates and returns
a boolean indicating if a clip area exists.
*/
func CairoGetClipRectangle(cr *C.cairo_t) (rect C.GdkRectangle, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_cairo_get_clip_rectangle(cr, &rect)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Adds the given rectangle to the current path of @cr.
*/
func CairoRectangle(cr *C.cairo_t, rectangle *C.GdkRectangle) {
	C.gdk_cairo_rectangle(cr, rectangle)
	return
}

/*
Adds the given region to the current path of @cr.
*/
func CairoRegion(cr *C.cairo_t, region *C.cairo_region_t) {
	C.gdk_cairo_region(cr, region)
	return
}

/*
Creates region that describes covers the area where the given
@surface is more than 50% opaque.

This function takes into account device offsets that might be
set with cairo_surface_set_device_offset().
*/
func CairoRegionCreateFromSurface(surface *C.cairo_surface_t) (return__ *C.cairo_region_t) {
	return__ = C.gdk_cairo_region_create_from_surface(surface)
	return
}

// gdk_cairo_set_source_color is not generated due to deprecation attr

/*
Sets the given pixbuf as the source pattern for @cr.

The pattern has an extend mode of %CAIRO_EXTEND_NONE and is aligned
so that the origin of @pixbuf is @pixbuf_x, @pixbuf_y.
*/
func CairoSetSourcePixbuf(cr *C.cairo_t, pixbuf *C.GdkPixbuf, pixbuf_x float64, pixbuf_y float64) {
	C.gdk_cairo_set_source_pixbuf(cr, pixbuf, C.gdouble(pixbuf_x), C.gdouble(pixbuf_y))
	return
}

/*
Sets the specified #GdkRGBA as the source color of @cr.
*/
func CairoSetSourceRgba(cr *C.cairo_t, rgba *C.GdkRGBA) {
	C.gdk_cairo_set_source_rgba(cr, rgba)
	return
}

/*
Sets the given window as the source pattern for @cr.

The pattern has an extend mode of %CAIRO_EXTEND_NONE and is aligned
so that the origin of @window is @x, @y. The window contains all its
subwindows when rendering.

Note that the contents of @window are undefined outside of the
visible part of @window, so use this function with care.
*/
func CairoSetSourceWindow(cr *C.cairo_t, window IsWindow, x float64, y float64) {
	var __cgo__window *C.GdkWindow
	if window != nil {
		__cgo__window = window.GetWindowPointer()
	}
	C.gdk_cairo_set_source_window(cr, __cgo__window, C.gdouble(x), C.gdouble(y))
	return
}

/*
Creates an image surface with the same contents as
the pixbuf.
*/
func CairoSurfaceCreateFromPixbuf(pixbuf *C.GdkPixbuf, scale int, for_window IsWindow) (return__ *C.cairo_surface_t) {
	var __cgo__for_window *C.GdkWindow
	if for_window != nil {
		__cgo__for_window = for_window.GetWindowPointer()
	}
	return__ = C.gdk_cairo_surface_create_from_pixbuf(pixbuf, C.int(scale), __cgo__for_window)
	return
}

// gdk_color_parse is not generated due to deprecation attr

/*
Disables multidevice support in GDK. This call must happen prior
to gdk_display_open(), gtk_init(), gtk_init_with_args() or
gtk_init_check() in order to take effect.

Most common GTK+ applications won’t ever need to call this. Only
applications that do mixed GDK/Xlib calls could want to disable
multidevice support if such Xlib code deals with input devices in
any way and doesn’t observe the presence of XInput 2.
*/
func DisableMultidevice() {
	C.gdk_disable_multidevice()
	return
}

/*
Aborts a drag without dropping.

This function is called by the drag source.
*/
func DragAbort(context IsDragContext, time_ uint32) {
	var __cgo__context *C.GdkDragContext
	if context != nil {
		__cgo__context = context.GetDragContextPointer()
	}
	C.gdk_drag_abort(__cgo__context, C.guint32(time_))
	return
}

/*
Starts a drag and creates a new drag context for it.
This function assumes that the drag is controlled by the
client pointer device, use gdk_drag_begin_for_device() to
begin a drag with a different device.

This function is called by the drag source.
*/
func DragBegin(window IsWindow, targets *C.GList) (return__ *DragContext) {
	var __cgo__window *C.GdkWindow
	if window != nil {
		__cgo__window = window.GetWindowPointer()
	}
	var __cgo__return__ *C.GdkDragContext
	__cgo__return__ = C.gdk_drag_begin(__cgo__window, targets)
	if __cgo__return__ != nil {
		return__ = NewDragContextFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Starts a drag and creates a new drag context for it.

This function is called by the drag source.
*/
func DragBeginForDevice(window IsWindow, device IsDevice, targets *C.GList) (return__ *DragContext) {
	var __cgo__window *C.GdkWindow
	if window != nil {
		__cgo__window = window.GetWindowPointer()
	}
	var __cgo__device *C.GdkDevice
	if device != nil {
		__cgo__device = device.GetDevicePointer()
	}
	var __cgo__return__ *C.GdkDragContext
	__cgo__return__ = C.gdk_drag_begin_for_device(__cgo__window, __cgo__device, targets)
	if __cgo__return__ != nil {
		return__ = NewDragContextFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Drops on the current destination.

This function is called by the drag source.
*/
func DragDrop(context IsDragContext, time_ uint32) {
	var __cgo__context *C.GdkDragContext
	if context != nil {
		__cgo__context = context.GetDragContextPointer()
	}
	C.gdk_drag_drop(__cgo__context, C.guint32(time_))
	return
}

/*
Returns whether the dropped data has been successfully
transferred. This function is intended to be used while
handling a %GDK_DROP_FINISHED event, its return value is
meaningless at other times.
*/
func DragDropSucceeded(context IsDragContext) (return__ bool) {
	var __cgo__context *C.GdkDragContext
	if context != nil {
		__cgo__context = context.GetDragContextPointer()
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_drag_drop_succeeded(__cgo__context)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Finds the destination window and DND protocol to use at the
given pointer position.

This function is called by the drag source to obtain the
@dest_window and @protocol parameters for gdk_drag_motion().
*/
func DragFindWindowForScreen(context IsDragContext, drag_window IsWindow, screen IsScreen, x_root int, y_root int) (dest_window *Window, protocol C.GdkDragProtocol) {
	var __cgo__context *C.GdkDragContext
	if context != nil {
		__cgo__context = context.GetDragContextPointer()
	}
	var __cgo__drag_window *C.GdkWindow
	if drag_window != nil {
		__cgo__drag_window = drag_window.GetWindowPointer()
	}
	var __cgo__screen *C.GdkScreen
	if screen != nil {
		__cgo__screen = screen.GetScreenPointer()
	}
	var __cgo__dest_window *C.GdkWindow
	C.gdk_drag_find_window_for_screen(__cgo__context, __cgo__drag_window, __cgo__screen, C.gint(x_root), C.gint(y_root), &__cgo__dest_window, &protocol)
	if __cgo__dest_window != nil {
		dest_window = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__dest_window).Pointer()))
	}
	return
}

/*
Returns the selection atom for the current source window.
*/
func DragGetSelection(context IsDragContext) (return__ C.GdkAtom) {
	var __cgo__context *C.GdkDragContext
	if context != nil {
		__cgo__context = context.GetDragContextPointer()
	}
	return__ = C.gdk_drag_get_selection(__cgo__context)
	return
}

/*
Updates the drag context when the pointer moves or the
set of actions changes.

This function is called by the drag source.
*/
func DragMotion(context IsDragContext, dest_window IsWindow, protocol C.GdkDragProtocol, x_root int, y_root int, suggested_action C.GdkDragAction, possible_actions C.GdkDragAction, time_ uint32) (return__ bool) {
	var __cgo__context *C.GdkDragContext
	if context != nil {
		__cgo__context = context.GetDragContextPointer()
	}
	var __cgo__dest_window *C.GdkWindow
	if dest_window != nil {
		__cgo__dest_window = dest_window.GetWindowPointer()
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_drag_motion(__cgo__context, __cgo__dest_window, protocol, C.gint(x_root), C.gint(y_root), suggested_action, possible_actions, C.guint32(time_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Selects one of the actions offered by the drag source.

This function is called by the drag destination in response to
gdk_drag_motion() called by the drag source.
*/
func DragStatus(context IsDragContext, action C.GdkDragAction, time_ uint32) {
	var __cgo__context *C.GdkDragContext
	if context != nil {
		__cgo__context = context.GetDragContextPointer()
	}
	C.gdk_drag_status(__cgo__context, action, C.guint32(time_))
	return
}

/*
Ends the drag operation after a drop.

This function is called by the drag destination.
*/
func DropFinish(context IsDragContext, success bool, time_ uint32) {
	var __cgo__context *C.GdkDragContext
	if context != nil {
		__cgo__context = context.GetDragContextPointer()
	}
	__cgo__success := C.gboolean(0)
	if success {
		__cgo__success = C.gboolean(1)
	}
	C.gdk_drop_finish(__cgo__context, __cgo__success, C.guint32(time_))
	return
}

/*
Accepts or rejects a drop.

This function is called by the drag destination in response
to a drop initiated by the drag source.
*/
func DropReply(context IsDragContext, accepted bool, time_ uint32) {
	var __cgo__context *C.GdkDragContext
	if context != nil {
		__cgo__context = context.GetDragContextPointer()
	}
	__cgo__accepted := C.gboolean(0)
	if accepted {
		__cgo__accepted = C.gboolean(1)
	}
	C.gdk_drop_reply(__cgo__context, __cgo__accepted, C.guint32(time_))
	return
}

/*
Removes an error trap pushed with gdk_error_trap_push().
May block until an error has been definitively received
or not received from the X server. gdk_error_trap_pop_ignored()
is preferred if you don’t need to know whether an error
occurred, because it never has to block. If you don't
need the return value of gdk_error_trap_pop(), use
gdk_error_trap_pop_ignored().

Prior to GDK 3.0, this function would not automatically
sync for you, so you had to gdk_flush() if your last
call to Xlib was not a blocking round trip.
*/
func ErrorTrapPop() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_error_trap_pop()
	return__ = int(__cgo__return__)
	return
}

/*
Removes an error trap pushed with gdk_error_trap_push(), but
without bothering to wait and see whether an error occurred.  If an
error arrives later asynchronously that was triggered while the
trap was pushed, that error will be ignored.
*/
func ErrorTrapPopIgnored() {
	C.gdk_error_trap_pop_ignored()
	return
}

/*
This function allows X errors to be trapped instead of the normal
behavior of exiting the application. It should only be used if it
is not possible to avoid the X error in any other way. Errors are
ignored on all #GdkDisplay currently known to the
#GdkDisplayManager. If you don’t care which error happens and just
want to ignore everything, pop with gdk_error_trap_pop_ignored().
If you need the error code, use gdk_error_trap_pop() which may have
to block and wait for the error to arrive from the X server.

This API exists on all platforms but only does anything on X.

You can use gdk_x11_display_error_trap_push() to ignore errors
on only a single display.

## Trapping an X error

|[<!-- language="C" -->
gdk_error_trap_push ();

 // ... Call the X function which may cause an error here ...


if (gdk_error_trap_pop ())
 {
   // ... Handle the error here ...
 }
]|
*/
func ErrorTrapPush() {
	C.gdk_error_trap_push()
	return
}

/*
Checks all open displays for a #GdkEvent to process,to be processed
on, fetching events from the windowing system if necessary.
See gdk_display_get_event().
*/
func EventGet() (return__ *C.GdkEvent) {
	return__ = C.gdk_event_get()
	return
}

/*
Sets the function to call to handle all events from GDK.

Note that GTK+ uses this to install its own event handler, so it is
usually not useful for GTK+ applications. (Although an application
can call this function then call gtk_main_do_event() to pass
events to GTK+.)
*/
func EventHandlerSet(func_ C.GdkEventFunc, data unsafe.Pointer, notify C.GDestroyNotify) {
	C.gdk_event_handler_set(func_, (C.gpointer)(data), notify)
	return
}

/*
If there is an event waiting in the event queue of some open
display, returns a copy of it. See gdk_display_peek_event().
*/
func EventPeek() (return__ *C.GdkEvent) {
	return__ = C.gdk_event_peek()
	return
}

/*
Request more motion notifies if @event is a motion notify hint event.

This function should be used instead of gdk_window_get_pointer() to
request further motion notifies, because it also works for extension
events where motion notifies are provided for devices other than the
core pointer. Coordinate extraction, processing and requesting more
motion events from a %GDK_MOTION_NOTIFY event usually works like this:

|[<!-- language="C" -->
{
  // motion_event handler
  x = motion_event->x;
  y = motion_event->y;
  // handle (x,y) motion
  gdk_event_request_motions (motion_event); // handles is_hint events
}
]|
*/
func EventRequestMotions(event *C.GdkEventMotion) {
	C.gdk_event_request_motions(event)
	return
}

/*
If both events contain X/Y information, this function will return %TRUE
and return in @angle the relative angle from @event1 to @event2. The rotation
direction for positive angles is from the positive X axis towards the positive
Y axis.
*/
func EventsGetAngle(event1 *C.GdkEvent, event2 *C.GdkEvent) (angle float64, return__ bool) {
	var __cgo__angle C.gdouble
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_events_get_angle(event1, event2, &__cgo__angle)
	angle = float64(__cgo__angle)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
If both events contain X/Y information, the center of both coordinates
will be returned in @x and @y.
*/
func EventsGetCenter(event1 *C.GdkEvent, event2 *C.GdkEvent) (x float64, y float64, return__ bool) {
	var __cgo__x C.gdouble
	var __cgo__y C.gdouble
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_events_get_center(event1, event2, &__cgo__x, &__cgo__y)
	x = float64(__cgo__x)
	y = float64(__cgo__y)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
If both events have X/Y information, the distance between both coordinates
(as in a straight line going from @event1 to @event2) will be returned.
*/
func EventsGetDistance(event1 *C.GdkEvent, event2 *C.GdkEvent) (distance float64, return__ bool) {
	var __cgo__distance C.gdouble
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_events_get_distance(event1, event2, &__cgo__distance)
	distance = float64(__cgo__distance)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if any events are ready to be processed for any display.
*/
func EventsPending() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_events_pending()
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Flushes the output buffers of all display connections and waits
until all requests have been processed.
This is rarely needed by applications.
*/
func Flush() {
	C.gdk_flush()
	return
}

/*
Obtains the root window (parent all other windows are inside)
for the default display and screen.
*/
func GetDefaultRootWindow() (return__ *Window) {
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_get_default_root_window()
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gdk_get_display is not generated due to deprecation attr

/*
Gets the display name specified in the command line arguments passed
to gdk_init() or gdk_parse_args(), if any.
*/
func GetDisplayArgName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gdk_get_display_arg_name()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the program class. Unless the program class has explicitly
been set with gdk_set_program_class() or with the `--class`
commandline option, the default value is the program name (determined
with g_get_prgname()) with the first character converted to uppercase.
*/
func GetProgramClass() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gdk_get_program_class()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets whether event debugging output is enabled.
*/
func GetShowEvents() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_get_show_events()
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gdk_init is not generated due to inout param

// gdk_init_check is not generated due to inout param

// gdk_keyboard_grab is not generated due to deprecation attr

// gdk_keyboard_ungrab is not generated due to deprecation attr

/*
Obtains the upper- and lower-case versions of the keyval @symbol.
Examples of keyvals are #GDK_KEY_a, #GDK_KEY_Enter, #GDK_KEY_F1, etc.
*/
func KeyvalConvertCase(symbol uint) (lower uint, upper uint) {
	var __cgo__lower C.guint
	var __cgo__upper C.guint
	C.gdk_keyval_convert_case(C.guint(symbol), &__cgo__lower, &__cgo__upper)
	lower = uint(__cgo__lower)
	upper = uint(__cgo__upper)
	return
}

/*
Converts a key name to a key value.

The names are the same as those in the
`gdk/gdkkeysyms.h` header file
but without the leading “GDK_KEY_”.
*/
func KeyvalFromName(keyval_name string) (return__ uint) {
	__cgo__keyval_name := (*C.gchar)(unsafe.Pointer(C.CString(keyval_name)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.gdk_keyval_from_name(__cgo__keyval_name)
	C.free(unsafe.Pointer(__cgo__keyval_name))
	return__ = uint(__cgo__return__)
	return
}

/*
Returns %TRUE if the given key value is in lower case.
*/
func KeyvalIsLower(keyval uint) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_keyval_is_lower(C.guint(keyval))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if the given key value is in upper case.
*/
func KeyvalIsUpper(keyval uint) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_keyval_is_upper(C.guint(keyval))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Converts a key value into a symbolic name.

The names are the same as those in the
`gdk/gdkkeysyms.h` header file
but without the leading “GDK_KEY_”.
*/
func KeyvalName(keyval uint) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gdk_keyval_name(C.guint(keyval))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts a key value to lower case, if applicable.
*/
func KeyvalToLower(keyval uint) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gdk_keyval_to_lower(C.guint(keyval))
	return__ = uint(__cgo__return__)
	return
}

/*
Convert from a GDK key symbol to the corresponding ISO10646 (Unicode)
character.
*/
func KeyvalToUnicode(keyval uint) (return__ uint32) {
	var __cgo__return__ C.guint32
	__cgo__return__ = C.gdk_keyval_to_unicode(C.guint(keyval))
	return__ = uint32(__cgo__return__)
	return
}

/*
Converts a key value to upper case, if applicable.
*/
func KeyvalToUpper(keyval uint) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gdk_keyval_to_upper(C.guint(keyval))
	return__ = uint(__cgo__return__)
	return
}

/*
Lists the available visuals for the default screen.
(See gdk_screen_list_visuals())
A visual describes a hardware image data format.
For example, a visual might support 24-bit color, or 8-bit color,
and might expect pixels to be in a certain format.

Call g_list_free() on the return value when you’re finished with it.
*/
func ListVisuals() (return__ *C.GList) {
	return__ = C.gdk_list_visuals()
	return
}

/*
Indicates to the GUI environment that the application has finished
loading. If the applications opens windows, this function is
normally called after opening the application’s initial set of
windows.

GTK+ will call this function automatically after opening the first
#GtkWindow unless gtk_window_set_auto_startup_notification() is called
to disable that feature.
*/
func NotifyStartupComplete() {
	C.gdk_notify_startup_complete()
	return
}

/*
Indicates to the GUI environment that the application has
finished loading, using a given identifier.

GTK+ will call this function automatically for #GtkWindow
with custom startup-notification identifier unless
gtk_window_set_auto_startup_notification() is called to
disable that feature.
*/
func NotifyStartupCompleteWithId(startup_id string) {
	__cgo__startup_id := (*C.gchar)(unsafe.Pointer(C.CString(startup_id)))
	C.gdk_notify_startup_complete_with_id(__cgo__startup_id)
	C.free(unsafe.Pointer(__cgo__startup_id))
	return
}

/*
Gets the window that @window is embedded in.
*/
func OffscreenWindowGetEmbedder(window IsWindow) (return__ *Window) {
	var __cgo__window *C.GdkWindow
	if window != nil {
		__cgo__window = window.GetWindowPointer()
	}
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_offscreen_window_get_embedder(__cgo__window)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the offscreen surface that an offscreen window renders into.
If you need to keep this around over window resizes, you need to
add a reference to it.
*/
func OffscreenWindowGetSurface(window IsWindow) (return__ *C.cairo_surface_t) {
	var __cgo__window *C.GdkWindow
	if window != nil {
		__cgo__window = window.GetWindowPointer()
	}
	return__ = C.gdk_offscreen_window_get_surface(__cgo__window)
	return
}

/*
Sets @window to be embedded in @embedder.

To fully embed an offscreen window, in addition to calling this
function, it is also necessary to handle the #GdkWindow::pick-embedded-child
signal on the @embedder and the #GdkWindow::to-embedder and
#GdkWindow::from-embedder signals on @window.
*/
func OffscreenWindowSetEmbedder(window IsWindow, embedder IsWindow) {
	var __cgo__window *C.GdkWindow
	if window != nil {
		__cgo__window = window.GetWindowPointer()
	}
	var __cgo__embedder *C.GdkWindow
	if embedder != nil {
		__cgo__embedder = embedder.GetWindowPointer()
	}
	C.gdk_offscreen_window_set_embedder(__cgo__window, __cgo__embedder)
	return
}

/*
Creates a #PangoContext for the default GDK screen.

The context must be freed when you’re finished with it.

When using GTK+, normally you should use gtk_widget_get_pango_context()
instead of this function, to get the appropriate context for
the widget you intend to render text onto.

The newly created context will have the default font options (see
#cairo_font_options_t) for the default screen; if these options
change it will not be updated. Using gtk_widget_get_pango_context()
is more convenient if you want to keep a context around and track
changes to the screen’s font rendering settings.
*/
func PangoContextGet() (return__ *C.PangoContext) {
	return__ = C.gdk_pango_context_get()
	return
}

/*
Creates a #PangoContext for @screen.

The context must be freed when you’re finished with it.

When using GTK+, normally you should use gtk_widget_get_pango_context()
instead of this function, to get the appropriate context for
the widget you intend to render text onto.

The newly created context will have the default font options
(see #cairo_font_options_t) for the screen; if these options
change it will not be updated. Using gtk_widget_get_pango_context()
is more convenient if you want to keep a context around and track
changes to the screen’s font rendering settings.
*/
func PangoContextGetForScreen(screen IsScreen) (return__ *C.PangoContext) {
	var __cgo__screen *C.GdkScreen
	if screen != nil {
		__cgo__screen = screen.GetScreenPointer()
	}
	return__ = C.gdk_pango_context_get_for_screen(__cgo__screen)
	return
}

/*
Obtains a clip region which contains the areas where the given ranges
of text would be drawn. @x_origin and @y_origin are the top left point
to center the layout. @index_ranges should contain
ranges of bytes in the layout’s text.

Note that the regions returned correspond to logical extents of the text
ranges, not ink extents. So the drawn layout may in fact touch areas out of
the clip region.  The clip region is mainly useful for highlightling parts
of text, such as when text is selected.
*/
func PangoLayoutGetClipRegion(layout *C.PangoLayout, x_origin int, y_origin int, index_ranges *C.gint, n_ranges int) (return__ *C.cairo_region_t) {
	return__ = C.gdk_pango_layout_get_clip_region(layout, C.gint(x_origin), C.gint(y_origin), index_ranges, C.gint(n_ranges))
	return
}

/*
Obtains a clip region which contains the areas where the given
ranges of text would be drawn. @x_origin and @y_origin are the top left
position of the layout. @index_ranges
should contain ranges of bytes in the layout’s text. The clip
region will include space to the left or right of the line (to the
layout bounding box) if you have indexes above or below the indexes
contained inside the line. This is to draw the selection all the way
to the side of the layout. However, the clip region is in line coordinates,
not layout coordinates.

Note that the regions returned correspond to logical extents of the text
ranges, not ink extents. So the drawn line may in fact touch areas out of
the clip region.  The clip region is mainly useful for highlightling parts
of text, such as when text is selected.
*/
func PangoLayoutLineGetClipRegion(line *C.PangoLayoutLine, x_origin int, y_origin int, index_ranges []int, n_ranges int) (return__ *C.cairo_region_t) {
	__header__index_ranges := (*reflect.SliceHeader)(unsafe.Pointer(&index_ranges))
	return__ = C.gdk_pango_layout_line_get_clip_region(line, C.gint(x_origin), C.gint(y_origin), (*C.gint)(unsafe.Pointer(__header__index_ranges.Data)), C.gint(n_ranges))
	return
}

// gdk_parse_args is not generated due to inout param

/*
Transfers image data from a #cairo_surface_t and converts it to an RGB(A)
representation inside a #GdkPixbuf. This allows you to efficiently read
individual pixels from cairo surfaces. For #GdkWindows, use
gdk_pixbuf_get_from_window() instead.

This function will create an RGB pixbuf with 8 bits per channel.
The pixbuf will contain an alpha channel if the @surface contains one.
*/
func PixbufGetFromSurface(surface *C.cairo_surface_t, src_x int, src_y int, width int, height int) (return__ *C.GdkPixbuf) {
	return__ = C.gdk_pixbuf_get_from_surface(surface, C.gint(src_x), C.gint(src_y), C.gint(width), C.gint(height))
	return
}

/*
Transfers image data from a #GdkWindow and converts it to an RGB(A)
representation inside a #GdkPixbuf. In other words, copies
image data from a server-side drawable to a client-side RGB(A) buffer.
This allows you to efficiently read individual pixels on the client side.

This function will create an RGB pixbuf with 8 bits per channel with
the same size specified by the @width and @height arguments. The pixbuf
will contain an alpha channel if the @window contains one.

If the window is off the screen, then there is no image data in the
obscured/offscreen regions to be placed in the pixbuf. The contents of
portions of the pixbuf corresponding to the offscreen region are undefined.

If the window you’re obtaining data from is partially obscured by
other windows, then the contents of the pixbuf areas corresponding
to the obscured regions are undefined.

If the window is not mapped (typically because it’s iconified/minimized
or not on the current workspace), then %NULL will be returned.

If memory can’t be allocated for the return value, %NULL will be returned
instead.

(In short, there are several ways this function can fail, and if it fails
 it returns %NULL; so check the return value.)
*/
func PixbufGetFromWindow(window IsWindow, src_x int, src_y int, width int, height int) (return__ *C.GdkPixbuf) {
	var __cgo__window *C.GdkWindow
	if window != nil {
		__cgo__window = window.GetWindowPointer()
	}
	return__ = C.gdk_pixbuf_get_from_window(__cgo__window, C.gint(src_x), C.gint(src_y), C.gint(width), C.gint(height))
	return
}

// gdk_pointer_grab is not generated due to deprecation attr

// gdk_pointer_is_grabbed is not generated due to deprecation attr

// gdk_pointer_ungrab is not generated due to deprecation attr

func PreParseLibgtkOnly() {
	C.gdk_pre_parse_libgtk_only()
	return
}

/*
Changes the contents of a property on a window.
*/
func PropertyChange(window IsWindow, property C.GdkAtom, type_ C.GdkAtom, format int, mode C.GdkPropMode, data *C.guchar, nelements int) {
	var __cgo__window *C.GdkWindow
	if window != nil {
		__cgo__window = window.GetWindowPointer()
	}
	C.gdk_property_change(__cgo__window, property, type_, C.gint(format), mode, data, C.gint(nelements))
	return
}

/*
Deletes a property from a window.
*/
func PropertyDelete(window IsWindow, property C.GdkAtom) {
	var __cgo__window *C.GdkWindow
	if window != nil {
		__cgo__window = window.GetWindowPointer()
	}
	C.gdk_property_delete(__cgo__window, property)
	return
}

/*
Retrieves a portion of the contents of a property. If the
property does not exist, then the function returns %FALSE,
and %GDK_NONE will be stored in @actual_property_type.

The XGetWindowProperty() function that gdk_property_get()
uses has a very confusing and complicated set of semantics.
Unfortunately, gdk_property_get() makes the situation
worse instead of better (the semantics should be considered
undefined), and also prints warnings to stderr in cases where it
should return a useful error to the program. You are advised to use
XGetWindowProperty() directly until a replacement function for
gdk_property_get() is provided.
*/
func PropertyGet(window IsWindow, property C.GdkAtom, type_ C.GdkAtom, offset uint64, length uint64, pdelete int) (actual_property_type C.GdkAtom, actual_format int, actual_length int, data string, return__ bool) {
	var __cgo__window *C.GdkWindow
	if window != nil {
		__cgo__window = window.GetWindowPointer()
	}
	var __cgo__actual_format C.gint
	var __cgo__actual_length C.gint
	var __cgo__data *C.guchar
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_property_get(__cgo__window, property, type_, C.gulong(offset), C.gulong(length), C.gint(pdelete), &actual_property_type, &__cgo__actual_format, &__cgo__actual_length, &__cgo__data)
	actual_format = int(__cgo__actual_format)
	actual_length = int(__cgo__actual_length)
	data = C.GoString((*C.char)(unsafe.Pointer(__cgo__data)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This function returns the available bit depths for the default
screen. It’s equivalent to listing the visuals
(gdk_list_visuals()) and then looking at the depth field in each
visual, removing duplicates.

The array returned by this function should not be freed.
*/
func QueryDepths() (depths []int, count int) {
	var __cgo__depths *C.gint
	var __cgo__count C.gint
	C.gdk_query_depths(&__cgo__depths, &__cgo__count)
	defer func() {
		__header__depths := (*reflect.SliceHeader)(unsafe.Pointer(&depths))
		__header__depths.Len = int(count)
		__header__depths.Cap = int(count)
		__header__depths.Data = uintptr(unsafe.Pointer(__cgo__depths))
	}()
	count = int(__cgo__count)
	return
}

/*
This function returns the available visual types for the default
screen. It’s equivalent to listing the visuals
(gdk_list_visuals()) and then looking at the type field in each
visual, removing duplicates.

The array returned by this function should not be freed.
*/
func QueryVisualTypes() (visual_types *C.GdkVisualType, count int) {
	var __cgo__count C.gint
	C.gdk_query_visual_types(&visual_types, &__cgo__count)
	count = int(__cgo__count)
	return
}

func RectangleGetType() (return__ C.GType) {
	return__ = C.gdk_rectangle_get_type()
	return
}

/*
Calculates the intersection of two rectangles. It is allowed for
@dest to be the same as either @src1 or @src2. If the rectangles
do not intersect, @dest’s width and height is set to 0 and its x
and y values are undefined. If you are only interested in whether
the rectangles intersect, but not in the intersecting area itself,
pass %NULL for @dest.
*/
func RectangleIntersect(src1 *C.GdkRectangle, src2 *C.GdkRectangle) (dest C.GdkRectangle, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_rectangle_intersect(src1, src2, &dest)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Calculates the union of two rectangles.
The union of rectangles @src1 and @src2 is the smallest rectangle which
includes both @src1 and @src2 within it.
It is allowed for @dest to be the same as either @src1 or @src2.
*/
func RectangleUnion(src1 *C.GdkRectangle, src2 *C.GdkRectangle) (dest C.GdkRectangle) {
	C.gdk_rectangle_union(src1, src2, &dest)
	return
}

/*
Retrieves the contents of a selection in a given
form.
*/
func SelectionConvert(requestor IsWindow, selection C.GdkAtom, target C.GdkAtom, time_ uint32) {
	var __cgo__requestor *C.GdkWindow
	if requestor != nil {
		__cgo__requestor = requestor.GetWindowPointer()
	}
	C.gdk_selection_convert(__cgo__requestor, selection, target, C.guint32(time_))
	return
}

/*
Determines the owner of the given selection.
*/
func SelectionOwnerGet(selection C.GdkAtom) (return__ *Window) {
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_selection_owner_get(selection)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Determine the owner of the given selection.

Note that the return value may be owned by a different
process if a foreign window was previously created for that
window, but a new foreign window will never be created by this call.
*/
func SelectionOwnerGetForDisplay(display IsDisplay, selection C.GdkAtom) (return__ *Window) {
	var __cgo__display *C.GdkDisplay
	if display != nil {
		__cgo__display = display.GetDisplayPointer()
	}
	var __cgo__return__ *C.GdkWindow
	__cgo__return__ = C.gdk_selection_owner_get_for_display(__cgo__display, selection)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Sets the owner of the given selection.
*/
func SelectionOwnerSet(owner IsWindow, selection C.GdkAtom, time_ uint32, send_event bool) (return__ bool) {
	var __cgo__owner *C.GdkWindow
	if owner != nil {
		__cgo__owner = owner.GetWindowPointer()
	}
	__cgo__send_event := C.gboolean(0)
	if send_event {
		__cgo__send_event = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_selection_owner_set(__cgo__owner, selection, C.guint32(time_), __cgo__send_event)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the #GdkWindow @owner as the current owner of the selection @selection.
*/
func SelectionOwnerSetForDisplay(display IsDisplay, owner IsWindow, selection C.GdkAtom, time_ uint32, send_event bool) (return__ bool) {
	var __cgo__display *C.GdkDisplay
	if display != nil {
		__cgo__display = display.GetDisplayPointer()
	}
	var __cgo__owner *C.GdkWindow
	if owner != nil {
		__cgo__owner = owner.GetWindowPointer()
	}
	__cgo__send_event := C.gboolean(0)
	if send_event {
		__cgo__send_event = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_selection_owner_set_for_display(__cgo__display, __cgo__owner, selection, C.guint32(time_), __cgo__send_event)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves selection data that was stored by the selection
data in response to a call to gdk_selection_convert(). This function
will not be used by applications, who should use the #GtkClipboard
API instead.
*/
func SelectionPropertyGet(requestor IsWindow, data []string, prop_type *C.GdkAtom, prop_format *C.gint) (return__ int) {
	var __cgo__requestor *C.GdkWindow
	if requestor != nil {
		__cgo__requestor = requestor.GetWindowPointer()
	}
	__header__data := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_selection_property_get(__cgo__requestor, (**C.guchar)(unsafe.Pointer(__header__data.Data)), prop_type, prop_format)
	return__ = int(__cgo__return__)
	return
}

/*
Sends a response to SelectionRequest event.
*/
func SelectionSendNotify(requestor IsWindow, selection C.GdkAtom, target C.GdkAtom, property C.GdkAtom, time_ uint32) {
	var __cgo__requestor *C.GdkWindow
	if requestor != nil {
		__cgo__requestor = requestor.GetWindowPointer()
	}
	C.gdk_selection_send_notify(__cgo__requestor, selection, target, property, C.guint32(time_))
	return
}

/*
Send a response to SelectionRequest event.
*/
func SelectionSendNotifyForDisplay(display IsDisplay, requestor IsWindow, selection C.GdkAtom, target C.GdkAtom, property C.GdkAtom, time_ uint32) {
	var __cgo__display *C.GdkDisplay
	if display != nil {
		__cgo__display = display.GetDisplayPointer()
	}
	var __cgo__requestor *C.GdkWindow
	if requestor != nil {
		__cgo__requestor = requestor.GetWindowPointer()
	}
	C.gdk_selection_send_notify_for_display(__cgo__display, __cgo__requestor, selection, target, property, C.guint32(time_))
	return
}

/*
Sets a list of backends that GDK should try to use.

This can be be useful if your application does not
work with certain GDK backends.

By default, GDK tries all included backends.

For example,
|[<!-- language="C" -->
gdk_set_allowed_backends ("wayland,quartz,*");
]|
instructs GDK to try the Wayland backend first,
followed by the Quartz backend, and then all
others.

If the `GDK_BACKEND` environment variable
is set, it determines what backends are tried in what
order, while still respecting the set of allowed backends
that are specified by this function.

The possible backend names are x11, win32, quartz,
broadway, wayland. You can also include a * in the
list to try all remaining backends.

This call must happen prior to gdk_display_open(),
gtk_init(), gtk_init_with_args() or gtk_init_check()
in order to take effect.
*/
func SetAllowedBackends(backends string) {
	__cgo__backends := (*C.gchar)(unsafe.Pointer(C.CString(backends)))
	C.gdk_set_allowed_backends(__cgo__backends)
	C.free(unsafe.Pointer(__cgo__backends))
	return
}

/*
Set the double click time for the default display. See
gdk_display_set_double_click_time().
See also gdk_display_set_double_click_distance().
Applications should not set this, it is a
global user-configured setting.
*/
func SetDoubleClickTime(msec uint) {
	C.gdk_set_double_click_time(C.guint(msec))
	return
}

/*
Sets the program class. The X11 backend uses the program class to set
the class name part of the `WM_CLASS` property on
toplevel windows; see the ICCCM.
*/
func SetProgramClass(program_class string) {
	__cgo__program_class := (*C.gchar)(unsafe.Pointer(C.CString(program_class)))
	C.gdk_set_program_class(__cgo__program_class)
	C.free(unsafe.Pointer(__cgo__program_class))
	return
}

/*
Sets whether a trace of received events is output.
Note that GTK+ must be compiled with debugging (that is,
configured using the `--enable-debug` option)
to use this option.
*/
func SetShowEvents(show_events bool) {
	__cgo__show_events := C.gboolean(0)
	if show_events {
		__cgo__show_events = C.gboolean(1)
	}
	C.gdk_set_show_events(__cgo__show_events)
	return
}

/*
Obtains a desktop-wide setting, such as the double-click time,
for the default screen. See gdk_screen_get_setting().
*/
func SettingGet(name string, value *C.GValue) (return__ bool) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_setting_get(__cgo__name, value)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

func SynthesizeWindowState(window IsWindow, unset_flags C.GdkWindowState, set_flags C.GdkWindowState) {
	var __cgo__window *C.GdkWindow
	if window != nil {
		__cgo__window = window.GetWindowPointer()
	}
	C.gdk_synthesize_window_state(__cgo__window, unset_flags, set_flags)
	return
}

/*
Retrieves a pixel from @window to force the windowing
system to carry out any pending rendering commands.

This function is intended to be used to synchronize with rendering
pipelines, to benchmark windowing system rendering operations.
*/
func TestRenderSync(window IsWindow) {
	var __cgo__window *C.GdkWindow
	if window != nil {
		__cgo__window = window.GetWindowPointer()
	}
	C.gdk_test_render_sync(__cgo__window)
	return
}

/*
This function is intended to be used in GTK+ test programs.
It will warp the mouse pointer to the given (@x,@y) coordinates
within @window and simulate a button press or release event.
Because the mouse pointer needs to be warped to the target
location, use of this function outside of test programs that
run in their own virtual windowing system (e.g. Xvfb) is not
recommended.

Also, gdk_test_simulate_button() is a fairly low level function,
for most testing purposes, gtk_test_widget_click() is the right
function to call which will generate a button press event followed
by its accompanying button release event.
*/
func TestSimulateButton(window IsWindow, x int, y int, button uint, modifiers C.GdkModifierType, button_pressrelease C.GdkEventType) (return__ bool) {
	var __cgo__window *C.GdkWindow
	if window != nil {
		__cgo__window = window.GetWindowPointer()
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_test_simulate_button(__cgo__window, C.gint(x), C.gint(y), C.guint(button), modifiers, button_pressrelease)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This function is intended to be used in GTK+ test programs.
If (@x,@y) are > (-1,-1), it will warp the mouse pointer to
the given (@x,@y) coordinates within @window and simulate a
key press or release event.

When the mouse pointer is warped to the target location, use
of this function outside of test programs that run in their
own virtual windowing system (e.g. Xvfb) is not recommended.
If (@x,@y) are passed as (-1,-1), the mouse pointer will not
be warped and @window origin will be used as mouse pointer
location for the event.

Also, gdk_test_simulate_key() is a fairly low level function,
for most testing purposes, gtk_test_widget_send_key() is the
right function to call which will generate a key press event
followed by its accompanying key release event.
*/
func TestSimulateKey(window IsWindow, x int, y int, keyval uint, modifiers C.GdkModifierType, key_pressrelease C.GdkEventType) (return__ bool) {
	var __cgo__window *C.GdkWindow
	if window != nil {
		__cgo__window = window.GetWindowPointer()
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_test_simulate_key(__cgo__window, C.gint(x), C.gint(y), C.guint(keyval), modifiers, key_pressrelease)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Converts a text property in the given encoding to
a list of UTF-8 strings.
*/
func TextPropertyToUtf8ListForDisplay(display IsDisplay, encoding C.GdkAtom, format int, text []byte, length int) (list **C.gchar, return__ int) {
	var __cgo__display *C.GdkDisplay
	if display != nil {
		__cgo__display = display.GetDisplayPointer()
	}
	__header__text := (*reflect.SliceHeader)(unsafe.Pointer(&text))
	var __cgo__return__ C.gint
	__cgo__return__ = C.gdk_text_property_to_utf8_list_for_display(__cgo__display, encoding, C.gint(format), (*C.guchar)(unsafe.Pointer(__header__text.Data)), C.gint(length), &list)
	return__ = int(__cgo__return__)
	return
}

/*
A wrapper for the common usage of gdk_threads_add_idle_full()
assigning the default priority, #G_PRIORITY_DEFAULT_IDLE.

See gdk_threads_add_idle_full().
*/
func ThreadsAddIdle(function C.GSourceFunc, data unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gdk_threads_add_idle(function, (C.gpointer)(data))
	return__ = uint(__cgo__return__)
	return
}

/*
Adds a function to be called whenever there are no higher priority
events pending.  If the function returns %FALSE it is automatically
removed from the list of event sources and will not be called again.

This variant of g_idle_add_full() calls @function with the GDK lock
held. It can be thought of a MT-safe version for GTK+ widgets for the
following use case, where you have to worry about idle_callback()
running in thread A and accessing @self after it has been finalized
in thread B:

|[<!-- language="C" -->
static gboolean
idle_callback (gpointer data)
{
   // gdk_threads_enter(); would be needed for g_idle_add()

   SomeWidget *self = data;
   // do stuff with self

   self->idle_id = 0;

   // gdk_threads_leave(); would be needed for g_idle_add()
   return FALSE;
}

static void
some_widget_do_stuff_later (SomeWidget *self)
{
   self->idle_id = gdk_threads_add_idle (idle_callback, self)
   // using g_idle_add() here would require thread protection in the callback
}

static void
some_widget_finalize (GObject *object)
{
   SomeWidget *self = SOME_WIDGET (object);
   if (self->idle_id)
     g_source_remove (self->idle_id);
   G_OBJECT_CLASS (parent_class)->finalize (object);
}
]|
*/
func ThreadsAddIdleFull(priority int, function C.GSourceFunc, data unsafe.Pointer, notify C.GDestroyNotify) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gdk_threads_add_idle_full(C.gint(priority), function, (C.gpointer)(data), notify)
	return__ = uint(__cgo__return__)
	return
}

/*
A wrapper for the common usage of gdk_threads_add_timeout_full()
assigning the default priority, #G_PRIORITY_DEFAULT.

See gdk_threads_add_timeout_full().
*/
func ThreadsAddTimeout(interval uint, function C.GSourceFunc, data unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gdk_threads_add_timeout(C.guint(interval), function, (C.gpointer)(data))
	return__ = uint(__cgo__return__)
	return
}

/*
Sets a function to be called at regular intervals holding the GDK lock,
with the given priority.  The function is called repeatedly until it
returns %FALSE, at which point the timeout is automatically destroyed
and the function will not be called again.  The @notify function is
called when the timeout is destroyed.  The first call to the
function will be at the end of the first @interval.

Note that timeout functions may be delayed, due to the processing of other
event sources. Thus they should not be relied on for precise timing.
After each call to the timeout function, the time of the next
timeout is recalculated based on the current time and the given interval
(it does not try to “catch up” time lost in delays).

This variant of g_timeout_add_full() can be thought of a MT-safe version
for GTK+ widgets for the following use case:

|[<!-- language="C" -->
static gboolean timeout_callback (gpointer data)
{
   SomeWidget *self = data;

   // do stuff with self

   self->timeout_id = 0;

   return G_SOURCE_REMOVE;
}

static void some_widget_do_stuff_later (SomeWidget *self)
{
   self->timeout_id = g_timeout_add (timeout_callback, self)
}

static void some_widget_finalize (GObject *object)
{
   SomeWidget *self = SOME_WIDGET (object);

   if (self->timeout_id)
     g_source_remove (self->timeout_id);

   G_OBJECT_CLASS (parent_class)->finalize (object);
}
]|
*/
func ThreadsAddTimeoutFull(priority int, interval uint, function C.GSourceFunc, data unsafe.Pointer, notify C.GDestroyNotify) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gdk_threads_add_timeout_full(C.gint(priority), C.guint(interval), function, (C.gpointer)(data), notify)
	return__ = uint(__cgo__return__)
	return
}

/*
A wrapper for the common usage of gdk_threads_add_timeout_seconds_full()
assigning the default priority, #G_PRIORITY_DEFAULT.

For details, see gdk_threads_add_timeout_full().
*/
func ThreadsAddTimeoutSeconds(interval uint, function C.GSourceFunc, data unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gdk_threads_add_timeout_seconds(C.guint(interval), function, (C.gpointer)(data))
	return__ = uint(__cgo__return__)
	return
}

/*
A variant of gdk_threads_add_timeout_full() with second-granularity.
See g_timeout_add_seconds_full() for a discussion of why it is
a good idea to use this function if you don’t need finer granularity.
*/
func ThreadsAddTimeoutSecondsFull(priority int, interval uint, function C.GSourceFunc, data unsafe.Pointer, notify C.GDestroyNotify) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gdk_threads_add_timeout_seconds_full(C.gint(priority), C.guint(interval), function, (C.gpointer)(data), notify)
	return__ = uint(__cgo__return__)
	return
}

// gdk_threads_enter is not generated due to deprecation attr

// gdk_threads_init is not generated due to deprecation attr

// gdk_threads_leave is not generated due to deprecation attr

// gdk_threads_set_lock_functions is not generated due to deprecation attr

/*
Convert from a ISO10646 character to a key symbol.
*/
func UnicodeToKeyval(wc uint32) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gdk_unicode_to_keyval(C.guint32(wc))
	return__ = uint(__cgo__return__)
	return
}

/*
Converts an UTF-8 string into the best possible representation
as a STRING. The representation of characters not in STRING
is not specified; it may be as pseudo-escape sequences
\x{ABCD}, or it may be in some other form of approximation.
*/
func Utf8ToStringTarget(str string) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gdk_utf8_to_string_target(__cgo__str)
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}
