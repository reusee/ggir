package gdk

/*
#include <gdk/gdk.h>
#include <gdk/gdkprivate.h>
#include <glib.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "reflect"
import "errors"
import "github.com/reusee/ggir/gobject"

func init() {
	_ = unsafe.Pointer(nil)
	_ = reflect.ValueOf(nil)
	_ = errors.New("")
}

var UnusedFix_ bool
var _ = gobject.UnusedFix_

// gdk_app_launch_context_new is not generated due to deprecation attr

/*
Creates a new cursor from the set of builtin cursors for the default display.
See gdk_cursor_new_for_display().

To make the cursor invisible, use %GDK_BLANK_CURSOR.
*/
func CursorNew(cursor_type C.GdkCursorType) (return__ *Cursor) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_cursor_new(cursor_type)
	if __cgo__return__ != nil {
		return__ = NewCursorFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new cursor from the set of builtin cursors.
Some useful ones are:
- ![](right_ptr.png) #GDK_RIGHT_PTR (right-facing arrow)
- ![](crosshair.png) #GDK_CROSSHAIR (crosshair)
- ![](xterm.png) #GDK_XTERM (I-beam)
- ![](watch.png) #GDK_WATCH (busy)
- ![](fleur.png) #GDK_FLEUR (for moving objects)
- ![](hand1.png) #GDK_HAND1 (a right-pointing hand)
- ![](hand2.png) #GDK_HAND2 (a left-pointing hand)
- ![](left_side.png) #GDK_LEFT_SIDE (resize left side)
- ![](right_side.png) #GDK_RIGHT_SIDE (resize right side)
- ![](top_left_corner.png) #GDK_TOP_LEFT_CORNER (resize northwest corner)
- ![](top_right_corner.png) #GDK_TOP_RIGHT_CORNER (resize northeast corner)
- ![](bottom_left_corner.png) #GDK_BOTTOM_LEFT_CORNER (resize southwest corner)
- ![](bottom_right_corner.png) #GDK_BOTTOM_RIGHT_CORNER (resize southeast corner)
- ![](top_side.png) #GDK_TOP_SIDE (resize top side)
- ![](bottom_side.png) #GDK_BOTTOM_SIDE (resize bottom side)
- ![](sb_h_double_arrow.png) #GDK_SB_H_DOUBLE_ARROW (move vertical splitter)
- ![](sb_v_double_arrow.png) #GDK_SB_V_DOUBLE_ARROW (move horizontal splitter)
- #GDK_BLANK_CURSOR (Blank cursor). Since 2.16
*/
func CursorNewForDisplay(display IsDisplay, cursor_type C.GdkCursorType) (return__ *Cursor) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_cursor_new_for_display(display.GetDisplayPointer(), cursor_type)
	if __cgo__return__ != nil {
		return__ = NewCursorFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new cursor by looking up @name in the current cursor
theme.
*/
func CursorNewFromName(display IsDisplay, name string) (return__ *Cursor) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_cursor_new_from_name(display.GetDisplayPointer(), __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	if __cgo__return__ != nil {
		return__ = NewCursorFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new cursor from a pixbuf.

Not all GDK backends support RGBA cursors. If they are not
supported, a monochrome approximation will be displayed.
The functions gdk_display_supports_cursor_alpha() and
gdk_display_supports_cursor_color() can be used to determine
whether RGBA cursors are supported;
gdk_display_get_default_cursor_size() and
gdk_display_get_maximal_cursor_size() give information about
cursor sizes.

If @x or @y are `-1`, the pixbuf must have
options named “x_hot” and “y_hot”, resp., containing
integer values between `0` and the width resp. height of
the pixbuf. (Since: 3.0)

On the X backend, support for RGBA cursors requires a
sufficently new version of the X Render extension.
*/
func CursorNewFromPixbuf(display IsDisplay, pixbuf *C.GdkPixbuf, x int, y int) (return__ *Cursor) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_cursor_new_from_pixbuf(display.GetDisplayPointer(), pixbuf, C.gint(x), C.gint(y))
	if __cgo__return__ != nil {
		return__ = NewCursorFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new cursor from a cairo image surface.

Not all GDK backends support RGBA cursors. If they are not
supported, a monochrome approximation will be displayed.
The functions gdk_display_supports_cursor_alpha() and
gdk_display_supports_cursor_color() can be used to determine
whether RGBA cursors are supported;
gdk_display_get_default_cursor_size() and
gdk_display_get_maximal_cursor_size() give information about
cursor sizes.

On the X backend, support for RGBA cursors requires a
sufficently new version of the X Render extension.
*/
func CursorNewFromSurface(display IsDisplay, surface *C.cairo_surface_t, x float64, y float64) (return__ *Cursor) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_cursor_new_from_surface(display.GetDisplayPointer(), surface, C.gdouble(x), C.gdouble(y))
	if __cgo__return__ != nil {
		return__ = NewCursorFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GdkWindow using the attributes from
@attributes. See #GdkWindowAttr and #GdkWindowAttributesType for
more details.  Note: to use this on displays other than the default
display, @parent must be specified.
*/
func WindowNew(parent IsWindow, attributes *C.GdkWindowAttr, attributes_mask C.gint) (return__ *Window) {
	var __cgo__parent *C.GdkWindow
	if parent != nil {
		__cgo__parent = parent.GetWindowPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_window_new(__cgo__parent, attributes, attributes_mask)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}
