package gtk

/*
#include <gtk/gtk.h>
#include <gtk/gtkx.h>
#include <glib.h>
#include <stdlib.h>
#cgo pkg-config: gtk+-3.0
*/
import "C"

import "github.com/reusee/ggir/gobject"
import "github.com/reusee/ggir/gdk"
import "github.com/reusee/ggir/cairo"
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
var _ = gdk.UnusedFix_
var _ = cairo.UnusedFix_

/*
Finds the first accelerator in any #GtkAccelGroup attached
to @object that matches @accel_key and @accel_mods, and
activates that accelerator.
*/
func AccelGroupsActivate(object *C.GObject, accel_key uint, accel_mods C.GdkModifierType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_accel_groups_activate(object, C.guint(accel_key), accel_mods)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets a list of all accel groups which are attached to @object.
*/
func AccelGroupsFromObject(object *C.GObject) (return__ *C.GSList) {
	return__ = C.gtk_accel_groups_from_object(object)
	return
}

/*
Gets the value set by gtk_accelerator_set_default_mod_mask().
*/
func AcceleratorGetDefaultModMask() (return__ C.GdkModifierType) {
	return__ = C.gtk_accelerator_get_default_mod_mask()
	return
}

/*
Converts an accelerator keyval and modifier mask into a string
which can be used to represent the accelerator to the user.
*/
func AcceleratorGetLabel(accelerator_key uint, accelerator_mods C.GdkModifierType) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_accelerator_get_label(C.guint(accelerator_key), accelerator_mods)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts an accelerator keyval and modifier mask
into a (possibly translated) string that can be displayed to
a user, similarly to gtk_accelerator_get_label(), but handling
keycodes.

This is only useful for system-level components, applications
should use gtk_accelerator_parse() instead.
*/
func AcceleratorGetLabelWithKeycode(display *C.GdkDisplay, accelerator_key uint, keycode uint, accelerator_mods C.GdkModifierType) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_accelerator_get_label_with_keycode(display, C.guint(accelerator_key), C.guint(keycode), accelerator_mods)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts an accelerator keyval and modifier mask into a string
parseable by gtk_accelerator_parse(). For example, if you pass in
#GDK_KEY_q and #GDK_CONTROL_MASK, this function returns “<Control>q”.

If you need to display accelerators in the user interface,
see gtk_accelerator_get_label().
*/
func AcceleratorName(accelerator_key uint, accelerator_mods C.GdkModifierType) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_accelerator_name(C.guint(accelerator_key), accelerator_mods)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts an accelerator keyval and modifier mask
into a string parseable by gtk_accelerator_parse_with_keycode(),
similarly to gtk_accelerator_name() but handling keycodes.
This is only useful for system-level components, applications
should use gtk_accelerator_parse() instead.
*/
func AcceleratorNameWithKeycode(display *C.GdkDisplay, accelerator_key uint, keycode uint, accelerator_mods C.GdkModifierType) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_accelerator_name_with_keycode(display, C.guint(accelerator_key), C.guint(keycode), accelerator_mods)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Parses a string representing an accelerator. The format looks like
“<Control>a” or “<Shift><Alt>F1” or “<Release>z” (the last one is
for key release).

The parser is fairly liberal and allows lower or upper case, and also
abbreviations such as “<Ctl>” and “<Ctrl>”. Key names are parsed using
gdk_keyval_from_name(). For character keys the name is not the symbol,
but the lowercase name, e.g. one would use “<Ctrl>minus” instead of
“<Ctrl>-”.

If the parse fails, @accelerator_key and @accelerator_mods will
be set to 0 (zero).
*/
func AcceleratorParse(accelerator string) (accelerator_key uint, accelerator_mods C.GdkModifierType) {
	__cgo__accelerator := (*C.gchar)(unsafe.Pointer(C.CString(accelerator)))
	var __cgo__accelerator_key C.guint
	C.gtk_accelerator_parse(__cgo__accelerator, &__cgo__accelerator_key, &accelerator_mods)
	C.free(unsafe.Pointer(__cgo__accelerator))
	accelerator_key = uint(__cgo__accelerator_key)
	return
}

/*
Parses a string representing an accelerator, similarly to
gtk_accelerator_parse() but handles keycodes as well. This is only
useful for system-level components, applications should use
gtk_accelerator_parse() instead.

If @accelerator_codes is given and the result stored in it is non-%NULL,
the result must be freed with g_free().

If a keycode is present in the accelerator and no @accelerator_codes
is given, the parse will fail.

If the parse fails, @accelerator_key, @accelerator_mods and
@accelerator_codes will be set to 0 (zero).
*/
func AcceleratorParseWithKeycode(accelerator string) (accelerator_key uint, accelerator_codes *C.guint, accelerator_mods C.GdkModifierType) {
	__cgo__accelerator := (*C.gchar)(unsafe.Pointer(C.CString(accelerator)))
	var __cgo__accelerator_key C.guint
	C.gtk_accelerator_parse_with_keycode(__cgo__accelerator, &__cgo__accelerator_key, &accelerator_codes, &accelerator_mods)
	C.free(unsafe.Pointer(__cgo__accelerator))
	accelerator_key = uint(__cgo__accelerator_key)
	return
}

/*
Sets the modifiers that will be considered significant for keyboard
accelerators. The default mod mask is #GDK_CONTROL_MASK |
#GDK_SHIFT_MASK | #GDK_MOD1_MASK | #GDK_SUPER_MASK |
#GDK_HYPER_MASK | #GDK_META_MASK, that is, Control, Shift, Alt,
Super, Hyper and Meta. Other modifiers will by default be ignored
by #GtkAccelGroup.
You must include at least the three modifiers Control, Shift
and Alt in any value you pass to this function.

The default mod mask should be changed on application startup,
before using any accelerator groups.
*/
func AcceleratorSetDefaultModMask(default_mod_mask C.GdkModifierType) {
	C.gtk_accelerator_set_default_mod_mask(default_mod_mask)
	return
}

/*
Determines whether a given keyval and modifier mask constitute
a valid keyboard accelerator. For example, the #GDK_KEY_a keyval
plus #GDK_CONTROL_MASK is valid - this is a “Ctrl+a” accelerator.
But, you can't, for instance, use the #GDK_KEY_Control_L keyval
as an accelerator.
*/
func AcceleratorValid(keyval uint, modifiers C.GdkModifierType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_accelerator_valid(C.guint(keyval), modifiers)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_alternative_dialog_button_order is not generated due to deprecation attr

// gtk_binding_entry_add_signal_from_string is not generated due to explicit ignore

// gtk_binding_entry_add_signall is not generated due to explicit ignore

// gtk_binding_entry_remove is not generated due to explicit ignore

// gtk_binding_entry_skip is not generated due to explicit ignore

// gtk_binding_set_by_class is not generated due to explicit ignore

// gtk_binding_set_find is not generated due to explicit ignore

// gtk_binding_set_new is not generated due to explicit ignore

// gtk_bindings_activate is not generated due to explicit ignore

// gtk_bindings_activate_event is not generated due to explicit ignore

func BuilderErrorQuark() (return__ C.GQuark) {
	return__ = C.gtk_builder_error_quark()
	return
}

/*
This function is supposed to be called in #GtkWidget::draw
implementations for widgets that support multiple windows.
@cr must be untransformed from invoking of the draw function.
This function will return %TRUE if the contents of the given
@window are supposed to be drawn and %FALSE otherwise. Note
that when the drawing was not initiated by the windowing
system this function will return %TRUE for all windows, so
you need to draw the bottommost window first. Also, do not
use “else if” statements to check which window should be drawn.
*/
func CairoShouldDrawWindow(cr *C.cairo_t, window *C.GdkWindow) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_cairo_should_draw_window(cr, window)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Transforms the given cairo context @cr that from @widget-relative
coordinates to @window-relative coordinates.
If the @widget’s window is not an ancestor of @window, no
modification will be applied.

This is the inverse to the transformation GTK applies when
preparing an expose event to be emitted with the #GtkWidget::draw
signal. It is intended to help porting multiwindow widgets from
GTK+ 2 to the rendering architecture of GTK+ 3.
*/
func CairoTransformToWindow(cr *C.cairo_t, widget IsWidget, window *C.GdkWindow) {
	C.gtk_cairo_transform_to_window(cr, widget.GetWidgetPointer(), window)
	return
}

/*
Checks that the GTK+ library in use is compatible with the
given version. Generally you would pass in the constants
#GTK_MAJOR_VERSION, #GTK_MINOR_VERSION, #GTK_MICRO_VERSION
as the three arguments to this function; that produces
a check that the library in use is compatible with
the version of GTK+ the application or module was compiled
against.

Compatibility is defined by two things: first the version
of the running library is newer than the version
@required_major.required_minor.@required_micro. Second
the running library must be binary compatible with the
version @required_major.required_minor.@required_micro
(same major version.)

This function is primarily for GTK+ modules; the module
can call this function to check that it wasn’t loaded
into an incompatible version of GTK+. However, such a
check isn’t completely reliable, since the module may be
linked against an old version of GTK+ and calling the
old version of gtk_check_version(), but still get loaded
into an application using a newer version of GTK+.
*/
func CheckVersion(required_major uint, required_minor uint, required_micro uint) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_check_version(C.guint(required_major), C.guint(required_minor), C.guint(required_micro))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

func CssProviderErrorQuark() (return__ C.GQuark) {
	return__ = C.gtk_css_provider_error_quark()
	return
}

/*
Adds a GTK+ grab on @device, so all the events on @device and its
associated pointer or keyboard (if any) are delivered to @widget.
If the @block_others parameter is %TRUE, any other devices will be
unable to interact with @widget during the grab.
*/
func DeviceGrabAdd(widget IsWidget, device *C.GdkDevice, block_others bool) {
	__cgo__block_others := C.gboolean(0)
	if block_others {
		__cgo__block_others = C.gboolean(1)
	}
	C.gtk_device_grab_add(widget.GetWidgetPointer(), device, __cgo__block_others)
	return
}

/*
Removes a device grab from the given widget.

You have to pair calls to gtk_device_grab_add() and
gtk_device_grab_remove().
*/
func DeviceGrabRemove(widget IsWidget, device *C.GdkDevice) {
	C.gtk_device_grab_remove(widget.GetWidgetPointer(), device)
	return
}

/*
Prevents gtk_init(), gtk_init_check(), gtk_init_with_args() and
gtk_parse_args() from automatically
calling `setlocale (LC_ALL, "")`. You would
want to use this function if you wanted to set the locale for
your program to something other than the user’s locale, or if
you wanted to set different values for different locale categories.

Most programs should not need to call this function.
*/
func DisableSetlocale() {
	C.gtk_disable_setlocale()
	return
}

/*
Distributes @extra_space to child @sizes by bringing smaller
children up to natural size first.

The remaining space will be added to the @minimum_size member of the
GtkRequestedSize struct. If all sizes reach their natural size then
the remaining space is returned.
*/
func DistributeNaturalAllocation(extra_space int, n_requested_sizes uint, sizes *C.GtkRequestedSize) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_distribute_natural_allocation(C.gint(extra_space), C.guint(n_requested_sizes), sizes)
	return__ = int(__cgo__return__)
	return
}

/*
Informs the drag source that the drop is finished, and
that the data of the drag will no longer be required.
*/
func DragFinish(context *C.GdkDragContext, success bool, del bool, time_ uint32) {
	__cgo__success := C.gboolean(0)
	if success {
		__cgo__success = C.gboolean(1)
	}
	__cgo__del := C.gboolean(0)
	if del {
		__cgo__del = C.gboolean(1)
	}
	C.gtk_drag_finish(context, __cgo__success, __cgo__del, C.guint32(time_))
	return
}

/*
Determines the source widget for a drag.
*/
func DragGetSourceWidget(context *C.GdkDragContext) (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_drag_get_source_widget(context)
	if __cgo__return__ != nil {
		return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Sets the icon for a particular drag to the default
icon.
*/
func DragSetIconDefault(context *C.GdkDragContext) {
	C.gtk_drag_set_icon_default(context)
	return
}

/*
Sets the icon for a given drag from the given @icon.
See the documentation for gtk_drag_set_icon_name()
for more details about using icons in drag and drop.
*/
func DragSetIconGicon(context *C.GdkDragContext, icon *C.GIcon, hot_x int, hot_y int) {
	C.gtk_drag_set_icon_gicon(context, icon, C.gint(hot_x), C.gint(hot_y))
	return
}

/*
Sets the icon for a given drag from a named themed icon. See
the docs for #GtkIconTheme for more details. Note that the
size of the icon depends on the icon theme (the icon is
loaded at the symbolic size #GTK_ICON_SIZE_DND), thus
@hot_x and @hot_y have to be used with care.
*/
func DragSetIconName(context *C.GdkDragContext, icon_name string, hot_x int, hot_y int) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	C.gtk_drag_set_icon_name(context, __cgo__icon_name, C.gint(hot_x), C.gint(hot_y))
	C.free(unsafe.Pointer(__cgo__icon_name))
	return
}

/*
Sets @pixbuf as the icon for a given drag.
*/
func DragSetIconPixbuf(context *C.GdkDragContext, pixbuf *C.GdkPixbuf, hot_x int, hot_y int) {
	C.gtk_drag_set_icon_pixbuf(context, pixbuf, C.gint(hot_x), C.gint(hot_y))
	return
}

// gtk_drag_set_icon_stock is not generated due to deprecation attr

/*
Sets @surface as the icon for a given drag. GTK+ retains
references for the arguments, and will release them when
they are no longer needed.

To position the surface relative to the mouse, use
cairo_surface_set_device_offset() on @surface. The mouse
cursor will be positioned at the (0,0) coordinate of the
surface.
*/
func DragSetIconSurface(context *C.GdkDragContext, surface *C.cairo_surface_t) {
	C.gtk_drag_set_icon_surface(context, surface)
	return
}

/*
Changes the icon for a widget to a given widget.
GTK+ will not destroy the icon, so if you don’t want
it to persist, you should connect to the “drag-end”
signal and destroy it yourself.

GTK+ will, however, change the opacity and position of
the window as part of the drag animation. If you want
to reuse the window, you have to restore these to
the values you need after each drag operation.
*/
func DragSetIconWidget(context *C.GdkDragContext, widget IsWidget, hot_x int, hot_y int) {
	C.gtk_drag_set_icon_widget(context, widget.GetWidgetPointer(), C.gint(hot_x), C.gint(hot_y))
	return
}

// gtk_draw_insertion_cursor is not generated due to deprecation attr

/*
Checks if any events are pending.

This can be used to update the UI and invoke timeouts etc.
while doing some time intensive computation.

## Updating the UI during a long computation

|[<!-- language="C" -->
 // computation going on...

 while (gtk_events_pending ())
   gtk_main_iteration ();

 // ...computation continued
]|
*/
func EventsPending() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_events_pending()
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Analogical to gtk_true(), this function does nothing
but always returns %FALSE.
*/
func False() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_false()
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Registers an error quark for #GtkFileChooser if necessary.
*/
func FileChooserErrorQuark() (return__ C.GQuark) {
	return__ = C.gtk_file_chooser_error_quark()
	return
}

/*
Returns the binary age as passed to `libtool`
when building the GTK+ library the process is running against.
If `libtool` means nothing to you, don't
worry about it.
*/
func GetBinaryAge() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_get_binary_age()
	return__ = uint(__cgo__return__)
	return
}

/*
Obtains a copy of the event currently being processed by GTK+.

For example, if you are handling a #GtkButton::clicked signal,
the current event will be the #GdkEventButton that triggered
the ::clicked signal.
*/
func GetCurrentEvent() (return__ *C.GdkEvent) {
	return__ = C.gtk_get_current_event()
	return
}

/*
If there is a current event and it has a device, return that
device, otherwise return %NULL.
*/
func GetCurrentEventDevice() (return__ *C.GdkDevice) {
	return__ = C.gtk_get_current_event_device()
	return
}

/*
If there is a current event and it has a state field, place
that state field in @state and return %TRUE, otherwise return
%FALSE.
*/
func GetCurrentEventState() (state C.GdkModifierType, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_get_current_event_state(&state)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
If there is a current event and it has a timestamp,
return that timestamp, otherwise return %GDK_CURRENT_TIME.
*/
func GetCurrentEventTime() (return__ uint32) {
	var __cgo__return__ C.guint32
	__cgo__return__ = C.gtk_get_current_event_time()
	return__ = uint32(__cgo__return__)
	return
}

/*
Returns the GTK+ debug flags.

This function is intended for GTK+ modules that want
to adjust their debug output based on GTK+ debug flags.
*/
func GetDebugFlags() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_get_debug_flags()
	return__ = uint(__cgo__return__)
	return
}

/*
Returns the #PangoLanguage for the default language currently in
effect. (Note that this can change over the life of an
application.) The default language is derived from the current
locale. It determines, for example, whether GTK+ uses the
right-to-left or left-to-right text direction.

This function is equivalent to pango_language_get_default().
See that function for details.
*/
func GetDefaultLanguage() (return__ *C.PangoLanguage) {
	return__ = C.gtk_get_default_language()
	return
}

/*
If @event is %NULL or the event was not associated with any widget,
returns %NULL, otherwise returns the widget that received the event
originally.
*/
func GetEventWidget(event *C.GdkEvent) (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_get_event_widget(event)
	if __cgo__return__ != nil {
		return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the interface age as passed to `libtool`
when building the GTK+ library the process is running against.
If `libtool` means nothing to you, don't
worry about it.
*/
func GetInterfaceAge() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_get_interface_age()
	return__ = uint(__cgo__return__)
	return
}

/*
Get the direction of the current locale. This is the expected
reading direction for text and UI.

This function depends on the current locale being set with
setlocale() and will default to setting the %GTK_TEXT_DIR_LTR
direction otherwise. %GTK_TEXT_DIR_NONE will never be returned.

GTK+ sets the default text direction according to the locale
during gtk_init(), and you should normally use
gtk_widget_get_direction() or gtk_widget_get_default_direction()
to obtain the current direcion.

This function is only needed rare cases when the locale is
changed after GTK+ has already been initialized. In this case,
you can use it to update the default text direction as follows:

|[<!-- language="C" -->
setlocale (LC_ALL, new_locale);
direction = gtk_get_locale_direction ();
gtk_widget_set_default_direction (direction);
]|
*/
func GetLocaleDirection() (return__ C.GtkTextDirection) {
	return__ = C.gtk_get_locale_direction()
	return
}

/*
Returns the major version number of the GTK+ library.
(e.g. in GTK+ version 3.1.5 this is 3.)

This function is in the library, so it represents the GTK+ library
your code is running against. Contrast with the #GTK_MAJOR_VERSION
macro, which represents the major version of the GTK+ headers you
have included when compiling your code.
*/
func GetMajorVersion() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_get_major_version()
	return__ = uint(__cgo__return__)
	return
}

/*
Returns the micro version number of the GTK+ library.
(e.g. in GTK+ version 3.1.5 this is 5.)

This function is in the library, so it represents the GTK+ library
your code is are running against. Contrast with the
#GTK_MICRO_VERSION macro, which represents the micro version of the
GTK+ headers you have included when compiling your code.
*/
func GetMicroVersion() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_get_micro_version()
	return__ = uint(__cgo__return__)
	return
}

/*
Returns the minor version number of the GTK+ library.
(e.g. in GTK+ version 3.1.5 this is 1.)

This function is in the library, so it represents the GTK+ library
your code is are running against. Contrast with the
#GTK_MINOR_VERSION macro, which represents the minor version of the
GTK+ headers you have included when compiling your code.
*/
func GetMinorVersion() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_get_minor_version()
	return__ = uint(__cgo__return__)
	return
}

/*
Returns a #GOptionGroup for the commandline arguments recognized
by GTK+ and GDK.

You should add this group to your #GOptionContext
with g_option_context_add_group(), if you are using
g_option_context_parse() to parse your commandline arguments.
*/
func GetOptionGroup(open_default_display bool) (return__ *C.GOptionGroup) {
	__cgo__open_default_display := C.gboolean(0)
	if open_default_display {
		__cgo__open_default_display = C.gboolean(1)
	}
	return__ = C.gtk_get_option_group(__cgo__open_default_display)
	return
}

/*
Queries the current grab of the default window group.
*/
func GrabGetCurrent() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_grab_get_current()
	if __cgo__return__ != nil {
		return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_icon_size_from_name is not generated due to deprecation attr

// gtk_icon_size_get_name is not generated due to deprecation attr

/*
Obtains the pixel size of a semantic icon size @size:
#GTK_ICON_SIZE_MENU, #GTK_ICON_SIZE_BUTTON, etc.  This function
isn’t normally needed, gtk_icon_theme_load_icon() is the usual
way to get an icon for rendering, then just look at the size of
the rendered pixbuf. The rendered pixbuf may not even correspond to
the width/height returned by gtk_icon_size_lookup(), because themes
are free to render the pixbuf however they like, including changing
the usual size.
*/
func IconSizeLookup(size C.GtkIconSize) (width int, height int, return__ bool) {
	var __cgo__width C.gint
	var __cgo__height C.gint
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_icon_size_lookup(size, &__cgo__width, &__cgo__height)
	width = int(__cgo__width)
	height = int(__cgo__height)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_icon_size_lookup_for_settings is not generated due to deprecation attr

// gtk_icon_size_register is not generated due to deprecation attr

// gtk_icon_size_register_alias is not generated due to deprecation attr

func IconThemeErrorQuark() (return__ C.GQuark) {
	return__ = C.gtk_icon_theme_error_quark()
	return
}

/*
Call this function before using any other GTK+ functions in your GUI
applications.  It will initialize everything needed to operate the
toolkit and parses some standard command line options.

Although you are expected to pass the @argc, @argv parameters from main() to
this function, it is possible to pass %NULL if @argv is not available or
commandline handling is not required.

@argc and @argv are adjusted accordingly so your own code will
never see those standard arguments.

Note that there are some alternative ways to initialize GTK+:
if you are calling gtk_parse_args(), gtk_init_check(),
gtk_init_with_args() or g_option_context_parse() with
the option group returned by gtk_get_option_group(),
you don’t have to call gtk_init().

And if you are using #GtkApplication, you don't have to call any of the
initialization functions either; the #GtkApplication::startup handler
does it for you.

This function will terminate your program if it was unable to
initialize the windowing system for some reason. If you want
your program to fall back to a textual interface you want to
call gtk_init_check() instead.

Since 2.18, GTK+ calls `signal (SIGPIPE, SIG_IGN)`
during initialization, to ignore SIGPIPE signals, since these are
almost never wanted in graphical applications. If you do need to
handle SIGPIPE for some reason, reset the handler after gtk_init(),
but notice that other libraries (e.g. libdbus or gvfs) might do
similar things.
*/
func Init(argc *C.int, argv ***C.char) {
	C.gtk_init(argc, argv)
	return
}

/*
This function does the same work as gtk_init() with only a single
change: It does not terminate the program if the windowing system
can’t be initialized. Instead it returns %FALSE on failure.

This way the application can fall back to some other means of
communication with the user - for example a curses or command line
interface.
*/
func InitCheck(argc *C.int, argv ***C.char) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_init_check(argc, argv)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This function does the same work as gtk_init_check().
Additionally, it allows you to add your own commandline options,
and it automatically generates nicely formatted
`--help` output. Note that your program will
be terminated after writing out the help output.
*/
func InitWithArgs(argc *C.gint, argv ***C.gchar, parameter_string string, entries *C.GOptionEntry, translation_domain string) (return__ bool, __err__ error) {
	__cgo__parameter_string := (*C.gchar)(unsafe.Pointer(C.CString(parameter_string)))
	__cgo__translation_domain := (*C.gchar)(unsafe.Pointer(C.CString(translation_domain)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_init_with_args(argc, argv, __cgo__parameter_string, entries, __cgo__translation_domain, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__parameter_string))
	C.free(unsafe.Pointer(__cgo__translation_domain))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

// gtk_key_snooper_install is not generated due to deprecation attr

// gtk_key_snooper_remove is not generated due to deprecation attr

/*
Runs the main loop until gtk_main_quit() is called.

You can nest calls to gtk_main(). In that case gtk_main_quit()
will make the innermost invocation of the main loop return.
*/
func Main() {
	C.gtk_main()
	return
}

/*
Processes a single GDK event.

This is public only to allow filtering of events between GDK and GTK+.
You will not usually need to call this function directly.

While you should not call this function directly, you might want to
know how exactly events are handled. So here is what this function
does with the event:

1. Compress enter/leave notify events. If the event passed build an
   enter/leave pair together with the next event (peeked from GDK), both
   events are thrown away. This is to avoid a backlog of (de-)highlighting
   widgets crossed by the pointer.

2. Find the widget which got the event. If the widget can’t be determined
   the event is thrown away unless it belongs to a INCR transaction.

3. Then the event is pushed onto a stack so you can query the currently
   handled event with gtk_get_current_event().

4. The event is sent to a widget. If a grab is active all events for widgets
   that are not in the contained in the grab widget are sent to the latter
   with a few exceptions:
   - Deletion and destruction events are still sent to the event widget for
     obvious reasons.
   - Events which directly relate to the visual representation of the event
     widget.
   - Leave events are delivered to the event widget if there was an enter
     event delivered to it before without the paired leave event.
   - Drag events are not redirected because it is unclear what the semantics
     of that would be.
   Another point of interest might be that all key events are first passed
   through the key snooper functions if there are any. Read the description
   of gtk_key_snooper_install() if you need this feature.

5. After finishing the delivery the event is popped from the event stack.
*/
func MainDoEvent(event *C.GdkEvent) {
	C.gtk_main_do_event(event)
	return
}

/*
Runs a single iteration of the mainloop.

If no events are waiting to be processed GTK+ will block
until the next event is noticed. If you don’t want to block
look at gtk_main_iteration_do() or check if any events are
pending with gtk_events_pending() first.
*/
func MainIteration() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_main_iteration()
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Runs a single iteration of the mainloop.
If no events are available either return or block depending on
the value of @blocking.
*/
func MainIterationDo(blocking bool) (return__ bool) {
	__cgo__blocking := C.gboolean(0)
	if blocking {
		__cgo__blocking = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_main_iteration_do(__cgo__blocking)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Asks for the current nesting level of the main loop.
*/
func MainLevel() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_main_level()
	return__ = uint(__cgo__return__)
	return
}

/*
Makes the innermost invocation of the main loop return
when it regains control.
*/
func MainQuit() {
	C.gtk_main_quit()
	return
}

// gtk_paint_arrow is not generated due to deprecation attr

// gtk_paint_box is not generated due to deprecation attr

// gtk_paint_box_gap is not generated due to deprecation attr

// gtk_paint_check is not generated due to deprecation attr

// gtk_paint_diamond is not generated due to deprecation attr

// gtk_paint_expander is not generated due to deprecation attr

// gtk_paint_extension is not generated due to deprecation attr

// gtk_paint_flat_box is not generated due to deprecation attr

// gtk_paint_focus is not generated due to deprecation attr

// gtk_paint_handle is not generated due to deprecation attr

// gtk_paint_hline is not generated due to deprecation attr

// gtk_paint_layout is not generated due to deprecation attr

// gtk_paint_option is not generated due to deprecation attr

// gtk_paint_resize_grip is not generated due to deprecation attr

// gtk_paint_shadow is not generated due to deprecation attr

// gtk_paint_shadow_gap is not generated due to deprecation attr

// gtk_paint_slider is not generated due to deprecation attr

// gtk_paint_spinner is not generated due to deprecation attr

// gtk_paint_tab is not generated due to deprecation attr

// gtk_paint_vline is not generated due to deprecation attr

/*
Returns the name of the default paper size, which
depends on the current locale.
*/
func PaperSizeGetDefault() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_paper_size_get_default()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Creates a list of known paper sizes.
*/
func PaperSizeGetPaperSizes(include_custom bool) (return__ *C.GList) {
	__cgo__include_custom := C.gboolean(0)
	if include_custom {
		__cgo__include_custom = C.gboolean(1)
	}
	return__ = C.gtk_paper_size_get_paper_sizes(__cgo__include_custom)
	return
}

/*
Parses command line arguments, and initializes global
attributes of GTK+, but does not actually open a connection
to a display. (See gdk_display_open(), gdk_get_display_arg_name())

Any arguments used by GTK+ or GDK are removed from the array and
@argc and @argv are updated accordingly.

There is no need to call this function explicitly if you are using
gtk_init(), or gtk_init_check().
*/
func ParseArgs(argc *C.int, argv ***C.char) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_parse_args(argc, argv)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Registers an error quark for #GtkPrintOperation if necessary.
*/
func PrintErrorQuark() (return__ C.GQuark) {
	return__ = C.gtk_print_error_quark()
	return
}

/*
Runs a page setup dialog, letting the user modify the values from
@page_setup. If the user cancels the dialog, the returned #GtkPageSetup
is identical to the passed in @page_setup, otherwise it contains the
modifications done in the dialog.

Note that this function may use a recursive mainloop to show the page
setup dialog. See gtk_print_run_page_setup_dialog_async() if this is
a problem.
*/
func PrintRunPageSetupDialog(parent IsWindow, page_setup IsPageSetup, settings IsPrintSettings) (return__ *PageSetup) {
	var __cgo__parent *C.GtkWindow
	if parent != nil {
		__cgo__parent = parent.GetWindowPointer()
	}
	var __cgo__page_setup *C.GtkPageSetup
	if page_setup != nil {
		__cgo__page_setup = page_setup.GetPageSetupPointer()
	}
	var __cgo__return__ *C.GtkPageSetup
	__cgo__return__ = C.gtk_print_run_page_setup_dialog(__cgo__parent, __cgo__page_setup, settings.GetPrintSettingsPointer())
	if __cgo__return__ != nil {
		return__ = NewPageSetupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Runs a page setup dialog, letting the user modify the values from @page_setup.

In contrast to gtk_print_run_page_setup_dialog(), this function  returns after
showing the page setup dialog on platforms that support this, and calls @done_cb
from a signal handler for the ::response signal of the dialog.
*/
func PrintRunPageSetupDialogAsync(parent IsWindow, page_setup IsPageSetup, settings IsPrintSettings, done_cb C.GtkPageSetupDoneFunc, data unsafe.Pointer) {
	var __cgo__parent *C.GtkWindow
	if parent != nil {
		__cgo__parent = parent.GetWindowPointer()
	}
	var __cgo__page_setup *C.GtkPageSetup
	if page_setup != nil {
		__cgo__page_setup = page_setup.GetPageSetupPointer()
	}
	C.gtk_print_run_page_setup_dialog_async(__cgo__parent, __cgo__page_setup, settings.GetPrintSettingsPointer(), done_cb, (C.gpointer)(data))
	return
}

/*
Sends an event to a widget, propagating the event to parent widgets
if the event remains unhandled.

Events received by GTK+ from GDK normally begin in gtk_main_do_event().
Depending on the type of event, existence of modal dialogs, grabs, etc.,
the event may be propagated; if so, this function is used.

gtk_propagate_event() calls gtk_widget_event() on each widget it
decides to send the event to. So gtk_widget_event() is the lowest-level
function; it simply emits the #GtkWidget::event and possibly an
event-specific signal on a widget. gtk_propagate_event() is a bit
higher-level, and gtk_main_do_event() is the highest level.

All that said, you most likely don’t want to use any of these
functions; synthesizing events is rarely needed. There are almost
certainly better ways to achieve your goals. For example, use
gdk_window_invalidate_rect() or gtk_widget_queue_draw() instead
of making up expose events.
*/
func PropagateEvent(widget IsWidget, event *C.GdkEvent) {
	C.gtk_propagate_event(widget.GetWidgetPointer(), event)
	return
}

// gtk_rc_add_default_file is not generated due to deprecation attr

// gtk_rc_find_module_in_path is not generated due to deprecation attr

// gtk_rc_find_pixmap_in_path is not generated due to deprecation attr

// gtk_rc_get_default_files is not generated due to deprecation attr

// gtk_rc_get_im_module_file is not generated due to deprecation attr

// gtk_rc_get_im_module_path is not generated due to deprecation attr

// gtk_rc_get_module_dir is not generated due to deprecation attr

// gtk_rc_get_style is not generated due to deprecation attr

// gtk_rc_get_style_by_paths is not generated due to deprecation attr

// gtk_rc_get_theme_dir is not generated due to deprecation attr

// gtk_rc_parse is not generated due to deprecation attr

// gtk_rc_parse_color is not generated due to deprecation attr

// gtk_rc_parse_color_full is not generated due to deprecation attr

// gtk_rc_parse_priority is not generated due to deprecation attr

// gtk_rc_parse_state is not generated due to deprecation attr

// gtk_rc_parse_string is not generated due to deprecation attr

/*
A #GtkRcPropertyParser for use with gtk_settings_install_property_parser()
or gtk_widget_class_install_style_property_parser() which parses
borders in the form
`"{ left, right, top, bottom }"` for integers
left, right, top and bottom.
*/
func RcPropertyParseBorder(pspec *C.GParamSpec, gstring *C.GString, property_value *C.GValue) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_rc_property_parse_border(pspec, gstring, property_value)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
A #GtkRcPropertyParser for use with gtk_settings_install_property_parser()
or gtk_widget_class_install_style_property_parser() which parses a
color given either by its name or in the form
`{ red, green, blue }` where red, green and
blue are integers between 0 and 65535 or floating-point numbers
between 0 and 1.
*/
func RcPropertyParseColor(pspec *C.GParamSpec, gstring *C.GString, property_value *C.GValue) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_rc_property_parse_color(pspec, gstring, property_value)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
A #GtkRcPropertyParser for use with gtk_settings_install_property_parser()
or gtk_widget_class_install_style_property_parser() which parses a single
enumeration value.

The enumeration value can be specified by its name, its nickname or
its numeric value. For consistency with flags parsing, the value
may be surrounded by parentheses.
*/
func RcPropertyParseEnum(pspec *C.GParamSpec, gstring *C.GString, property_value *C.GValue) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_rc_property_parse_enum(pspec, gstring, property_value)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
A #GtkRcPropertyParser for use with gtk_settings_install_property_parser()
or gtk_widget_class_install_style_property_parser() which parses flags.

Flags can be specified by their name, their nickname or
numerically. Multiple flags can be specified in the form
`"( flag1 | flag2 | ... )"`.
*/
func RcPropertyParseFlags(pspec *C.GParamSpec, gstring *C.GString, property_value *C.GValue) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_rc_property_parse_flags(pspec, gstring, property_value)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
A #GtkRcPropertyParser for use with gtk_settings_install_property_parser()
or gtk_widget_class_install_style_property_parser() which parses a
requisition in the form
`"{ width, height }"` for integers %width and %height.
*/
func RcPropertyParseRequisition(pspec *C.GParamSpec, gstring *C.GString, property_value *C.GValue) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_rc_property_parse_requisition(pspec, gstring, property_value)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_rc_reparse_all is not generated due to deprecation attr

// gtk_rc_reparse_all_for_settings is not generated due to deprecation attr

// gtk_rc_reset_styles is not generated due to deprecation attr

// gtk_rc_scanner_new is not generated due to deprecation attr

// gtk_rc_set_default_files is not generated due to deprecation attr

func RecentChooserErrorQuark() (return__ C.GQuark) {
	return__ = C.gtk_recent_chooser_error_quark()
	return
}

func RecentManagerErrorQuark() (return__ C.GQuark) {
	return__ = C.gtk_recent_manager_error_quark()
	return
}

/*
Renders an activity indicator (such as in #GtkSpinner).
The state %GTK_STATE_FLAG_ACTIVE determines whether there is
activity going on.
*/
func RenderActivity(context IsStyleContext, cr *C.cairo_t, x float64, y float64, width float64, height float64) {
	C.gtk_render_activity(context.GetStyleContextPointer(), cr, C.gdouble(x), C.gdouble(y), C.gdouble(width), C.gdouble(height))
	return
}

/*
Renders an arrow pointing to @angle.

Typical arrow rendering at 0, 1&solidus;2 &pi;, &pi; and 3&solidus;2 &pi;:

![](arrows.png)
*/
func RenderArrow(context IsStyleContext, cr *C.cairo_t, angle float64, x float64, y float64, size float64) {
	C.gtk_render_arrow(context.GetStyleContextPointer(), cr, C.gdouble(angle), C.gdouble(x), C.gdouble(y), C.gdouble(size))
	return
}

/*
Renders the background of an element.

Typical background rendering, showing the effect of
`background-image`, `border-width` and `border-radius`:

![](background.png)
*/
func RenderBackground(context IsStyleContext, cr *C.cairo_t, x float64, y float64, width float64, height float64) {
	C.gtk_render_background(context.GetStyleContextPointer(), cr, C.gdouble(x), C.gdouble(y), C.gdouble(width), C.gdouble(height))
	return
}

/*
Renders a checkmark (as in a #GtkCheckButton).

The %GTK_STATE_FLAG_ACTIVE state determines whether the check is
on or off, and %GTK_STATE_FLAG_INCONSISTENT determines whether it
should be marked as undefined.

Typical checkmark rendering:

![](checks.png)
*/
func RenderCheck(context IsStyleContext, cr *C.cairo_t, x float64, y float64, width float64, height float64) {
	C.gtk_render_check(context.GetStyleContextPointer(), cr, C.gdouble(x), C.gdouble(y), C.gdouble(width), C.gdouble(height))
	return
}

/*
Renders an expander (as used in #GtkTreeView and #GtkExpander) in the area
defined by @x, @y, @width, @height. The state %GTK_STATE_FLAG_ACTIVE
determines whether the expander is collapsed or expanded.

Typical expander rendering:

![](expanders.png)
*/
func RenderExpander(context IsStyleContext, cr *C.cairo_t, x float64, y float64, width float64, height float64) {
	C.gtk_render_expander(context.GetStyleContextPointer(), cr, C.gdouble(x), C.gdouble(y), C.gdouble(width), C.gdouble(height))
	return
}

/*
Renders a extension (as in a #GtkNotebook tab) in the rectangle
defined by @x, @y, @width, @height. The side where the extension
connects to is defined by @gap_side.

Typical extension rendering:

![](extensions.png)
*/
func RenderExtension(context IsStyleContext, cr *C.cairo_t, x float64, y float64, width float64, height float64, gap_side C.GtkPositionType) {
	C.gtk_render_extension(context.GetStyleContextPointer(), cr, C.gdouble(x), C.gdouble(y), C.gdouble(width), C.gdouble(height), gap_side)
	return
}

/*
Renders a focus indicator on the rectangle determined by @x, @y, @width, @height.

Typical focus rendering:

![](focus.png)
*/
func RenderFocus(context IsStyleContext, cr *C.cairo_t, x float64, y float64, width float64, height float64) {
	C.gtk_render_focus(context.GetStyleContextPointer(), cr, C.gdouble(x), C.gdouble(y), C.gdouble(width), C.gdouble(height))
	return
}

/*
Renders a frame around the rectangle defined by @x, @y, @width, @height.

Examples of frame rendering, showing the effect of `border-image`,
`border-color`, `border-width`, `border-radius` and junctions:

![](frames.png)
*/
func RenderFrame(context IsStyleContext, cr *C.cairo_t, x float64, y float64, width float64, height float64) {
	C.gtk_render_frame(context.GetStyleContextPointer(), cr, C.gdouble(x), C.gdouble(y), C.gdouble(width), C.gdouble(height))
	return
}

/*
Renders a frame around the rectangle defined by (@x, @y, @width, @height),
leaving a gap on one side. @xy0_gap and @xy1_gap will mean X coordinates
for %GTK_POS_TOP and %GTK_POS_BOTTOM gap sides, and Y coordinates for
%GTK_POS_LEFT and %GTK_POS_RIGHT.

Typical rendering of a frame with a gap:

![](frame-gap.png)
*/
func RenderFrameGap(context IsStyleContext, cr *C.cairo_t, x float64, y float64, width float64, height float64, gap_side C.GtkPositionType, xy0_gap float64, xy1_gap float64) {
	C.gtk_render_frame_gap(context.GetStyleContextPointer(), cr, C.gdouble(x), C.gdouble(y), C.gdouble(width), C.gdouble(height), gap_side, C.gdouble(xy0_gap), C.gdouble(xy1_gap))
	return
}

/*
Renders a handle (as in #GtkHandleBox, #GtkPaned and
#GtkWindow’s resize grip), in the rectangle
determined by @x, @y, @width, @height.

Handles rendered for the paned and grip classes:

![](handles.png)
*/
func RenderHandle(context IsStyleContext, cr *C.cairo_t, x float64, y float64, width float64, height float64) {
	C.gtk_render_handle(context.GetStyleContextPointer(), cr, C.gdouble(x), C.gdouble(y), C.gdouble(width), C.gdouble(height))
	return
}

/*
Renders the icon in @pixbuf at the specified @x and @y coordinates.
*/
func RenderIcon(context IsStyleContext, cr *C.cairo_t, pixbuf *C.GdkPixbuf, x float64, y float64) {
	C.gtk_render_icon(context.GetStyleContextPointer(), cr, pixbuf, C.gdouble(x), C.gdouble(y))
	return
}

// gtk_render_icon_pixbuf is not generated due to deprecation attr

/*
Renders the icon in @surface at the specified @x and @y coordinates.
*/
func RenderIconSurface(context IsStyleContext, cr *C.cairo_t, surface *C.cairo_surface_t, x float64, y float64) {
	C.gtk_render_icon_surface(context.GetStyleContextPointer(), cr, surface, C.gdouble(x), C.gdouble(y))
	return
}

/*
Draws a text caret on @cr at the specified index of @layout.
*/
func RenderInsertionCursor(context IsStyleContext, cr *C.cairo_t, x float64, y float64, layout *C.PangoLayout, index int, direction C.PangoDirection) {
	C.gtk_render_insertion_cursor(context.GetStyleContextPointer(), cr, C.gdouble(x), C.gdouble(y), layout, C.int(index), direction)
	return
}

/*
Renders @layout on the coordinates @x, @y
*/
func RenderLayout(context IsStyleContext, cr *C.cairo_t, x float64, y float64, layout *C.PangoLayout) {
	C.gtk_render_layout(context.GetStyleContextPointer(), cr, C.gdouble(x), C.gdouble(y), layout)
	return
}

/*
Renders a line from (x0, y0) to (x1, y1).
*/
func RenderLine(context IsStyleContext, cr *C.cairo_t, x0 float64, y0 float64, x1 float64, y1 float64) {
	C.gtk_render_line(context.GetStyleContextPointer(), cr, C.gdouble(x0), C.gdouble(y0), C.gdouble(x1), C.gdouble(y1))
	return
}

/*
Renders an option mark (as in a #GtkRadioButton), the %GTK_STATE_FLAG_ACTIVE
state will determine whether the option is on or off, and
%GTK_STATE_FLAG_INCONSISTENT whether it should be marked as undefined.

Typical option mark rendering:

![](options.png)
*/
func RenderOption(context IsStyleContext, cr *C.cairo_t, x float64, y float64, width float64, height float64) {
	C.gtk_render_option(context.GetStyleContextPointer(), cr, C.gdouble(x), C.gdouble(y), C.gdouble(width), C.gdouble(height))
	return
}

/*
Renders a slider (as in #GtkScale) in the rectangle defined by @x, @y,
@width, @height. @orientation defines whether the slider is vertical
or horizontal.

Typical slider rendering:

![](sliders.png)
*/
func RenderSlider(context IsStyleContext, cr *C.cairo_t, x float64, y float64, width float64, height float64, orientation C.GtkOrientation) {
	C.gtk_render_slider(context.GetStyleContextPointer(), cr, C.gdouble(x), C.gdouble(y), C.gdouble(width), C.gdouble(height), orientation)
	return
}

/*
Converts a color from RGB space to HSV.

Input values must be in the [0.0, 1.0] range;
output values will be in the same range.
*/
func RgbToHsv(r float64, g float64, b float64) (h float64, s float64, v float64) {
	var __cgo__h C.gdouble
	var __cgo__s C.gdouble
	var __cgo__v C.gdouble
	C.gtk_rgb_to_hsv(C.gdouble(r), C.gdouble(g), C.gdouble(b), &__cgo__h, &__cgo__s, &__cgo__v)
	h = float64(__cgo__h)
	s = float64(__cgo__s)
	v = float64(__cgo__v)
	return
}

/*
Appends a specified target to the list of supported targets for a
given widget and selection.
*/
func SelectionAddTarget(widget IsWidget, selection C.GdkAtom, target C.GdkAtom, info uint) {
	C.gtk_selection_add_target(widget.GetWidgetPointer(), selection, target, C.guint(info))
	return
}

/*
Prepends a table of targets to the list of supported targets
for a given widget and selection.
*/
func SelectionAddTargets(widget IsWidget, selection C.GdkAtom, targets *C.GtkTargetEntry, ntargets uint) {
	C.gtk_selection_add_targets(widget.GetWidgetPointer(), selection, targets, C.guint(ntargets))
	return
}

/*
Remove all targets registered for the given selection for the
widget.
*/
func SelectionClearTargets(widget IsWidget, selection C.GdkAtom) {
	C.gtk_selection_clear_targets(widget.GetWidgetPointer(), selection)
	return
}

/*
Requests the contents of a selection. When received,
a “selection-received” signal will be generated.
*/
func SelectionConvert(widget IsWidget, selection C.GdkAtom, target C.GdkAtom, time_ uint32) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_selection_convert(widget.GetWidgetPointer(), selection, target, C.guint32(time_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Claims ownership of a given selection for a particular widget,
or, if @widget is %NULL, release ownership of the selection.
*/
func SelectionOwnerSet(widget IsWidget, selection C.GdkAtom, time_ uint32) (return__ bool) {
	var __cgo__widget *C.GtkWidget
	if widget != nil {
		__cgo__widget = widget.GetWidgetPointer()
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_selection_owner_set(__cgo__widget, selection, C.guint32(time_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Claim ownership of a given selection for a particular widget, or,
if @widget is %NULL, release ownership of the selection.
*/
func SelectionOwnerSetForDisplay(display *C.GdkDisplay, widget IsWidget, selection C.GdkAtom, time_ uint32) (return__ bool) {
	var __cgo__widget *C.GtkWidget
	if widget != nil {
		__cgo__widget = widget.GetWidgetPointer()
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_selection_owner_set_for_display(display, __cgo__widget, selection, C.guint32(time_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Removes all handlers and unsets ownership of all
selections for a widget. Called when widget is being
destroyed. This function will not generally be
called by applications.
*/
func SelectionRemoveAll(widget IsWidget) {
	C.gtk_selection_remove_all(widget.GetWidgetPointer())
	return
}

/*
Sets the GTK+ debug flags.
*/
func SetDebugFlags(flags uint) {
	C.gtk_set_debug_flags(C.guint(flags))
	return
}

// gtk_show_about_dialog is not generated due to varargs

/*
This is a convenience function for launching the default application
to show the uri. The uri must be of a form understood by GIO (i.e. you
need to install gvfs to get support for uri schemes such as http://
or ftp://, as only local files are handled by GIO itself).
Typical examples are
- `file:///home/gnome/pict.jpg`
- `http://www.gnome.org`
- `mailto:me@gnome.org`

Ideally the timestamp is taken from the event triggering
the gtk_show_uri() call. If timestamp is not known you can take
%GDK_CURRENT_TIME.
*/
func ShowUri(screen *C.GdkScreen, uri string, timestamp uint32) (return__ bool, __err__ error) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_show_uri(screen, __cgo__uri, C.guint32(timestamp), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__uri))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

// gtk_stock_add is not generated due to deprecation attr

// gtk_stock_add_static is not generated due to deprecation attr

// gtk_stock_list_ids is not generated due to deprecation attr

// gtk_stock_lookup is not generated due to deprecation attr

// gtk_stock_set_translate_func is not generated due to deprecation attr

/*
This function frees a target table as returned by
gtk_target_table_new_from_list()
*/
func TargetTableFree(targets *C.GtkTargetEntry, n_targets int) {
	C.gtk_target_table_free(targets, C.gint(n_targets))
	return
}

/*
This function creates an #GtkTargetEntry array that contains the
same targets as the passed %list. The returned table is newly
allocated and should be freed using gtk_target_table_free() when no
longer needed.
*/
func TargetTableNewFromList(list *C.GtkTargetList) (n_targets int, return__ *C.GtkTargetEntry) {
	var __cgo__n_targets C.gint
	return__ = C.gtk_target_table_new_from_list(list, &__cgo__n_targets)
	n_targets = int(__cgo__n_targets)
	return
}

/*
Determines if any of the targets in @targets can be used to
provide a #GdkPixbuf.
*/
func TargetsIncludeImage(targets *C.GdkAtom, n_targets int, writable bool) (return__ bool) {
	__cgo__writable := C.gboolean(0)
	if writable {
		__cgo__writable = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_targets_include_image(targets, C.gint(n_targets), __cgo__writable)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines if any of the targets in @targets can be used to
provide rich text.
*/
func TargetsIncludeRichText(targets *C.GdkAtom, n_targets int, buffer IsTextBuffer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_targets_include_rich_text(targets, C.gint(n_targets), buffer.GetTextBufferPointer())
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines if any of the targets in @targets can be used to
provide text.
*/
func TargetsIncludeText(targets *C.GdkAtom, n_targets int) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_targets_include_text(targets, C.gint(n_targets))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines if any of the targets in @targets can be used to
provide an uri list.
*/
func TargetsIncludeUri(targets *C.GdkAtom, n_targets int) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_targets_include_uri(targets, C.gint(n_targets))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Create a simple window with window title @window_title and
text contents @dialog_text.
The window will quit any running gtk_main()-loop when destroyed, and it
will automatically be destroyed upon test function teardown.
*/
func TestCreateSimpleWindow(window_title string, dialog_text string) (return__ *Widget) {
	__cgo__window_title := (*C.gchar)(unsafe.Pointer(C.CString(window_title)))
	__cgo__dialog_text := (*C.gchar)(unsafe.Pointer(C.CString(dialog_text)))
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_test_create_simple_window(__cgo__window_title, __cgo__dialog_text)
	C.free(unsafe.Pointer(__cgo__window_title))
	C.free(unsafe.Pointer(__cgo__dialog_text))
	if __cgo__return__ != nil {
		return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_test_create_widget is not generated due to varargs

// gtk_test_display_button_window is not generated due to varargs

/*
This function will search @widget and all its descendants for a GtkLabel
widget with a text string matching @label_pattern.
The @label_pattern may contain asterisks “*” and question marks “?” as
placeholders, g_pattern_match() is used for the matching.
Note that locales other than "C“ tend to alter (translate” label strings,
so this function is genrally only useful in test programs with
predetermined locales, see gtk_test_init() for more details.
*/
func TestFindLabel(widget IsWidget, label_pattern string) (return__ *Widget) {
	__cgo__label_pattern := (*C.gchar)(unsafe.Pointer(C.CString(label_pattern)))
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_test_find_label(widget.GetWidgetPointer(), __cgo__label_pattern)
	C.free(unsafe.Pointer(__cgo__label_pattern))
	if __cgo__return__ != nil {
		return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
This function will search siblings of @base_widget and siblings of its
ancestors for all widgets matching @widget_type.
Of the matching widgets, the one that is geometrically closest to
@base_widget will be returned.
The general purpose of this function is to find the most likely “action”
widget, relative to another labeling widget. Such as finding a
button or text entry widget, given its corresponding label widget.
*/
func TestFindSibling(base_widget IsWidget, widget_type C.GType) (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_test_find_sibling(base_widget.GetWidgetPointer(), widget_type)
	if __cgo__return__ != nil {
		return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
This function will search the descendants of @widget for a widget
of type @widget_type that has a label matching @label_pattern next
to it. This is most useful for automated GUI testing, e.g. to find
the “OK” button in a dialog and synthesize clicks on it.
However see gtk_test_find_label(), gtk_test_find_sibling() and
gtk_test_widget_click() for possible caveats involving the search of
such widgets and synthesizing widget events.
*/
func TestFindWidget(widget IsWidget, label_pattern string, widget_type C.GType) (return__ *Widget) {
	__cgo__label_pattern := (*C.gchar)(unsafe.Pointer(C.CString(label_pattern)))
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_test_find_widget(widget.GetWidgetPointer(), __cgo__label_pattern, widget_type)
	C.free(unsafe.Pointer(__cgo__label_pattern))
	if __cgo__return__ != nil {
		return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_test_init is not generated due to varargs

/*
Return the type ids that have been registered after
calling gtk_test_register_all_types().
*/
func TestListAllTypes() (n_types uint, return__ *C.GType) {
	var __cgo__n_types C.guint
	return__ = C.gtk_test_list_all_types(&__cgo__n_types)
	n_types = uint(__cgo__n_types)
	return
}

/*
Force registration of all core Gtk+ and Gdk object types.
This allowes to refer to any of those object types via
g_type_from_name() after calling this function.
*/
func TestRegisterAllTypes() {
	C.gtk_test_register_all_types()
	return
}

/*
Retrive the literal adjustment value for GtkRange based
widgets and spin buttons. Note that the value returned by
this function is anything between the lower and upper bounds
of the adjustment belonging to @widget, and is not a percentage
as passed in to gtk_test_slider_set_perc().
*/
func TestSliderGetValue(widget IsWidget) (return__ float64) {
	var __cgo__return__ C.double
	__cgo__return__ = C.gtk_test_slider_get_value(widget.GetWidgetPointer())
	return__ = float64(__cgo__return__)
	return
}

/*
This function will adjust the slider position of all GtkRange
based widgets, such as scrollbars or scales, it’ll also adjust
spin buttons. The adjustment value of these widgets is set to
a value between the lower and upper limits, according to the
@percentage argument.
*/
func TestSliderSetPerc(widget IsWidget, percentage float64) {
	C.gtk_test_slider_set_perc(widget.GetWidgetPointer(), C.double(percentage))
	return
}

/*
This function will generate a @button click in the upwards or downwards
spin button arrow areas, usually leading to an increase or decrease of
spin button’s value.
*/
func TestSpinButtonClick(spinner IsSpinButton, button uint, upwards bool) (return__ bool) {
	__cgo__upwards := C.gboolean(0)
	if upwards {
		__cgo__upwards = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_test_spin_button_click(spinner.GetSpinButtonPointer(), C.guint(button), __cgo__upwards)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrive the text string of @widget if it is a GtkLabel,
GtkEditable (entry and text widgets) or GtkTextView.
*/
func TestTextGet(widget IsWidget) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_test_text_get(widget.GetWidgetPointer())
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Set the text string of @widget to @string if it is a GtkLabel,
GtkEditable (entry and text widgets) or GtkTextView.
*/
func TestTextSet(widget IsWidget, string_ string) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	C.gtk_test_text_set(widget.GetWidgetPointer(), __cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return
}

/*
This function will generate a @button click (button press and button
release event) in the middle of the first GdkWindow found that belongs
to @widget.
For %GTK_NO_WINDOW widgets like GtkButton, this will often be an
input-only event window. For other widgets, this is usually widget->window.
Certain caveats should be considered when using this function, in
particular because the mouse pointer is warped to the button click
location, see gdk_test_simulate_button() for details.
*/
func TestWidgetClick(widget IsWidget, button uint, modifiers C.GdkModifierType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_test_widget_click(widget.GetWidgetPointer(), C.guint(button), modifiers)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This function will generate keyboard press and release events in
the middle of the first GdkWindow found that belongs to @widget.
For %GTK_NO_WINDOW widgets like GtkButton, this will often be an
input-only event window. For other widgets, this is usually widget->window.
Certain caveats should be considered when using this function, in
particular because the mouse pointer is warped to the key press
location, see gdk_test_simulate_key() for details.
*/
func TestWidgetSendKey(widget IsWidget, keyval uint, modifiers C.GdkModifierType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_test_widget_send_key(widget.GetWidgetPointer(), C.guint(keyval), modifiers)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Enters the main loop and waits for @widget to be “drawn”. In this
context that means it waits for the frame clock of @widget to have
run a full styling, layout and drawing cycle.

This function is intended to be used for syncing with actions that
depend on @widget relayouting or on interaction with the display
server.
*/
func TestWidgetWaitForDraw(widget IsWidget) {
	C.gtk_test_widget_wait_for_draw(widget.GetWidgetPointer())
	return
}

/*
Obtains a @tree_model and @path from selection data of target type
%GTK_TREE_MODEL_ROW. Normally called from a drag_data_received handler.
This function can only be used if @selection_data originates from the same
process that’s calling this function, because a pointer to the tree model
is being passed around. If you aren’t in the same process, then you'll
get memory corruption. In the #GtkTreeDragDest drag_data_received handler,
you can assume that selection data of type %GTK_TREE_MODEL_ROW is
in from the current process. The returned path must be freed with
gtk_tree_path_free().
*/
func TreeGetRowDragData(selection_data *C.GtkSelectionData) (tree_model *C.GtkTreeModel, path *C.GtkTreePath, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_get_row_drag_data(selection_data, &tree_model, &path)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Lets a set of row reference created by
gtk_tree_row_reference_new_proxy() know that the
model emitted the #GtkTreeModel::row-deleted signal.
*/
func TreeRowReferenceDeleted(proxy *C.GObject, path *C.GtkTreePath) {
	C.gtk_tree_row_reference_deleted(proxy, path)
	return
}

/*
Lets a set of row reference created by
gtk_tree_row_reference_new_proxy() know that the
model emitted the #GtkTreeModel::row-inserted signal.
*/
func TreeRowReferenceInserted(proxy *C.GObject, path *C.GtkTreePath) {
	C.gtk_tree_row_reference_inserted(proxy, path)
	return
}

/*
Lets a set of row reference created by
gtk_tree_row_reference_new_proxy() know that the
model emitted the #GtkTreeModel::rows-reordered signal.
*/
func TreeRowReferenceReordered(proxy *C.GObject, path *C.GtkTreePath, iter *C.GtkTreeIter, new_order []int) {
	__header__new_order := (*reflect.SliceHeader)(unsafe.Pointer(&new_order))
	C.gtk_tree_row_reference_reordered(proxy, path, iter, (*C.gint)(unsafe.Pointer(__header__new_order.Data)))
	return
}

/*
Sets selection data of target type %GTK_TREE_MODEL_ROW. Normally used
in a drag_data_get handler.
*/
func TreeSetRowDragData(selection_data *C.GtkSelectionData, tree_model *C.GtkTreeModel, path *C.GtkTreePath) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_set_row_drag_data(selection_data, tree_model, path)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
All this function does it to return %TRUE.

This can be useful for example if you want to inhibit the deletion
of a window. Of course you should not do this as the user expects
a reaction from clicking the close icon of the window...

## A persistent window

|[<!-- language="C" -->
#include <gtk/gtk.h>

int
main (int argc, char **argv)
{
  GtkWidget *win, *but;
  const char *text = "Close yourself. I mean it!";

  gtk_init (&argc, &argv);

  win = gtk_window_new (GTK_WINDOW_TOPLEVEL);
  g_signal_connect (win,
                    "delete-event",
                    G_CALLBACK (gtk_true),
                    NULL);
  g_signal_connect (win, "destroy",
                    G_CALLBACK (gtk_main_quit),
                    NULL);

  but = gtk_button_new_with_label (text);
  g_signal_connect_swapped (but, "clicked",
                            G_CALLBACK (gtk_object_destroy),
                            win);
  gtk_container_add (GTK_CONTAINER (win), but);

  gtk_widget_show_all (win);

  gtk_main ();

  return 0;
}
]|
*/
func True() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_true()
	return__ = __cgo__return__ == C.gboolean(1)
	return
}
