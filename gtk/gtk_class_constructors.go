package gtk

/*
#include <gtk/gtk.h>
#include <gtk/gtkx.h>
#include <glib.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "reflect"
import "errors"
import "github.com/reusee/ggir/gobject"
import "github.com/reusee/ggir/gdk"
import "github.com/reusee/ggir/cairo"

func init() {
	_ = unsafe.Pointer(nil)
	_ = reflect.ValueOf(nil)
	_ = errors.New("")
}

var UnusedFix_ bool
var _ = gobject.UnusedFix_
var _ = gdk.UnusedFix_
var _ = cairo.UnusedFix_

/*
Creates a new #GtkAboutDialog.
*/
func AboutDialogNew() (return__ *AboutDialog) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_about_dialog_new()
	if __cgo__return__ != nil {
		return__ = NewAboutDialogFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkAccelGroup.
*/
func AccelGroupNew() (return__ *AccelGroup) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_accel_group_new()
	if __cgo__return__ != nil {
		return__ = NewAccelGroupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkAccelLabel.
*/
func AccelLabelNew(string_ string) (return__ *AccelLabel) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_accel_label_new(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	if __cgo__return__ != nil {
		return__ = NewAccelLabelFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_action_new is not generated due to deprecation attr

/*
Creates a new #GtkActionBar widget.
*/
func ActionBarNew() (return__ *ActionBar) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_action_bar_new()
	if __cgo__return__ != nil {
		return__ = NewActionBarFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_action_group_new is not generated due to deprecation attr

/*
Creates a new #GtkAdjustment.
*/
func AdjustmentNew(value float64, lower float64, upper float64, step_increment float64, page_increment float64, page_size float64) (return__ *Adjustment) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_adjustment_new(C.gdouble(value), C.gdouble(lower), C.gdouble(upper), C.gdouble(step_increment), C.gdouble(page_increment), C.gdouble(page_size))
	if __cgo__return__ != nil {
		return__ = NewAdjustmentFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_alignment_new is not generated due to deprecation attr

/*
Creates a new #GtkAppChooserButton for applications
that can handle content of the given type.
*/
func AppChooserButtonNew(content_type string) (return__ *AppChooserButton) {
	__cgo__content_type := (*C.gchar)(unsafe.Pointer(C.CString(content_type)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_app_chooser_button_new(__cgo__content_type)
	C.free(unsafe.Pointer(__cgo__content_type))
	if __cgo__return__ != nil {
		return__ = NewAppChooserButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkAppChooserDialog for the provided #GFile,
to allow the user to select an application for it.
*/
func AppChooserDialogNew(parent IsWindow, flags C.GtkDialogFlags, file *C.GFile) (return__ *AppChooserDialog) {
	var __cgo__parent *C.GtkWindow
	if parent != nil {
		__cgo__parent = parent.GetWindowPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_app_chooser_dialog_new(__cgo__parent, flags, file)
	if __cgo__return__ != nil {
		return__ = NewAppChooserDialogFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkAppChooserDialog for the provided content type,
to allow the user to select an application for it.
*/
func AppChooserDialogNewForContentType(parent IsWindow, flags C.GtkDialogFlags, content_type string) (return__ *AppChooserDialog) {
	var __cgo__parent *C.GtkWindow
	if parent != nil {
		__cgo__parent = parent.GetWindowPointer()
	}
	__cgo__content_type := (*C.gchar)(unsafe.Pointer(C.CString(content_type)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_app_chooser_dialog_new_for_content_type(__cgo__parent, flags, __cgo__content_type)
	C.free(unsafe.Pointer(__cgo__content_type))
	if __cgo__return__ != nil {
		return__ = NewAppChooserDialogFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkAppChooserWidget for applications
that can handle content of the given type.
*/
func AppChooserWidgetNew(content_type string) (return__ *AppChooserWidget) {
	__cgo__content_type := (*C.gchar)(unsafe.Pointer(C.CString(content_type)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_app_chooser_widget_new(__cgo__content_type)
	C.free(unsafe.Pointer(__cgo__content_type))
	if __cgo__return__ != nil {
		return__ = NewAppChooserWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkApplication instance.

When using #GtkApplication, it is not necessary to call gtk_init()
manually. It is called as soon as the application gets registered as
the primary instance.

Concretely, gtk_init() is called in the default handler for the
#GApplication::startup signal. Therefore, #GtkApplication subclasses should
chain up in their #GApplication::startup handler before using any GTK+ API.

Note that commandline arguments are not passed to gtk_init().
All GTK+ functionality that is available via commandline arguments
can also be achieved by setting suitable environment variables
such as `G_DEBUG`, so this should not be a big
problem. If you absolutely must support GTK+ commandline arguments,
you can explicitly call gtk_init() before creating the application
instance.

If non-%NULL, the application ID must be valid.  See
g_application_id_is_valid().

If no application ID is given then some features (most notably application
uniqueness) will be disabled. A null application ID is only allowed with
GTK+ 3.6 or later.
*/
func ApplicationNew(application_id string, flags C.GApplicationFlags) (return__ *Application) {
	__cgo__application_id := (*C.gchar)(unsafe.Pointer(C.CString(application_id)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_application_new(__cgo__application_id, flags)
	C.free(unsafe.Pointer(__cgo__application_id))
	if __cgo__return__ != nil {
		return__ = NewApplicationFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkApplicationWindow.
*/
func ApplicationWindowNew(application IsApplication) (return__ *ApplicationWindow) {
	var __cgo__application *C.GtkApplication
	if application != nil {
		__cgo__application = application.GetApplicationPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_application_window_new(__cgo__application)
	if __cgo__return__ != nil {
		return__ = NewApplicationWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_arrow_new is not generated due to deprecation attr

/*
Create a new #GtkAspectFrame.
*/
func AspectFrameNew(label string, xalign float32, yalign float32, ratio float32, obey_child bool) (return__ *AspectFrame) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	__cgo__obey_child := C.gboolean(0)
	if obey_child {
		__cgo__obey_child = C.gboolean(1)
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_aspect_frame_new(__cgo__label, C.gfloat(xalign), C.gfloat(yalign), C.gfloat(ratio), __cgo__obey_child)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewAspectFrameFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkAssistant.
*/
func AssistantNew() (return__ *Assistant) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_assistant_new()
	if __cgo__return__ != nil {
		return__ = NewAssistantFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkBox.
*/
func BoxNew(orientation C.GtkOrientation, spacing int) (return__ *Box) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_box_new(orientation, C.gint(spacing))
	if __cgo__return__ != nil {
		return__ = NewBoxFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new empty builder object.

This function is only useful if you intend to make multiple calls to
gtk_builder_add_from_file(), gtk_builder_add_from_resource() or
gtk_builder_add_from_string() in order to merge multiple UI
descriptions into a single builder.

Most users will probably want to use gtk_builder_new_from_file(),
gtk_builder_new_from_resource() or gtk_builder_new_from_string().
*/
func BuilderNew() (return__ *Builder) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_builder_new()
	if __cgo__return__ != nil {
		return__ = NewBuilderFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Builds the [GtkBuilder UI definition][BUILDER-UI]
in the file @filename.

If there is an error opening the file or parsing the description then
the program will be aborted.  You should only ever attempt to parse
user interface descriptions that are shipped as part of your program.
*/
func BuilderNewFromFile(filename string) (return__ *Builder) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_builder_new_from_file(__cgo__filename)
	C.free(unsafe.Pointer(__cgo__filename))
	if __cgo__return__ != nil {
		return__ = NewBuilderFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Builds the [GtkBuilder UI definition][BUILDER-UI]
at @resource_path.

If there is an error locating the resurce or parsing the description
then the program will be aborted.
*/
func BuilderNewFromResource(resource_path string) (return__ *Builder) {
	__cgo__resource_path := (*C.gchar)(unsafe.Pointer(C.CString(resource_path)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_builder_new_from_resource(__cgo__resource_path)
	C.free(unsafe.Pointer(__cgo__resource_path))
	if __cgo__return__ != nil {
		return__ = NewBuilderFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Builds the user interface described by @string (in the
[GtkBuilder UI definition][BUILDER-UI] format).

If @string is %NULL-terminated then @length should be -1.  If @length
is not -1 then it is the length of @string.

If there is an error parsing @string then the program will be
aborted.  You should not attempt to parse user interface description
from untrusted sources.
*/
func BuilderNewFromString(string_ string, length int64) (return__ *Builder) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_builder_new_from_string(__cgo__string_, C.gssize(length))
	C.free(unsafe.Pointer(__cgo__string_))
	if __cgo__return__ != nil {
		return__ = NewBuilderFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkButton widget. To add a child widget to the button,
use gtk_container_add().
*/
func ButtonNew() (return__ *Button) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_button_new()
	if __cgo__return__ != nil {
		return__ = NewButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new button containing an icon from the current icon theme.

If the icon name isn’t known, a “broken image” icon will be
displayed instead. If the current icon theme is changed, the icon
will be updated appropriately.

This function is a convenience wrapper around gtk_button_new() and
gtk_button_set_image().
*/
func ButtonNewFromIconName(icon_name string, size C.GtkIconSize) (return__ *Button) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_button_new_from_icon_name(__cgo__icon_name, size)
	C.free(unsafe.Pointer(__cgo__icon_name))
	if __cgo__return__ != nil {
		return__ = NewButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_button_new_from_stock is not generated due to deprecation attr

/*
Creates a #GtkButton widget with a #GtkLabel child containing the given
text.
*/
func ButtonNewWithLabel(label string) (return__ *Button) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_button_new_with_label(__cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkButton containing a label.
If characters in @label are preceded by an underscore, they are underlined.
If you need a literal underscore character in a label, use “__” (two
underscores). The first underlined character represents a keyboard
accelerator called a mnemonic.
Pressing Alt and that key activates the button.
*/
func ButtonNewWithMnemonic(label string) (return__ *Button) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_button_new_with_mnemonic(__cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkButtonBox.
*/
func ButtonBoxNew(orientation C.GtkOrientation) (return__ *ButtonBox) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_button_box_new(orientation)
	if __cgo__return__ != nil {
		return__ = NewButtonBoxFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new calendar, with the current date being selected.
*/
func CalendarNew() (return__ *Calendar) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_calendar_new()
	if __cgo__return__ != nil {
		return__ = NewCalendarFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCellAreaBox.
*/
func CellAreaBoxNew() (return__ *CellAreaBox) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_cell_area_box_new()
	if __cgo__return__ != nil {
		return__ = NewCellAreaBoxFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCellRendererAccel.
*/
func CellRendererAccelNew() (return__ *CellRendererAccel) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_cell_renderer_accel_new()
	if __cgo__return__ != nil {
		return__ = NewCellRendererAccelFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCellRendererCombo.
Adjust how text is drawn using object properties.
Object properties can be set globally (with g_object_set()).
Also, with #GtkTreeViewColumn, you can bind a property to a value
in a #GtkTreeModel. For example, you can bind the “text” property
on the cell renderer to a string value in the model, thus rendering
a different string in each row of the #GtkTreeView.
*/
func CellRendererComboNew() (return__ *CellRendererCombo) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_cell_renderer_combo_new()
	if __cgo__return__ != nil {
		return__ = NewCellRendererComboFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCellRendererPixbuf. Adjust rendering
parameters using object properties. Object properties can be set
globally (with g_object_set()). Also, with #GtkTreeViewColumn, you
can bind a property to a value in a #GtkTreeModel. For example, you
can bind the “pixbuf” property on the cell renderer to a pixbuf value
in the model, thus rendering a different image in each row of the
#GtkTreeView.
*/
func CellRendererPixbufNew() (return__ *CellRendererPixbuf) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_cell_renderer_pixbuf_new()
	if __cgo__return__ != nil {
		return__ = NewCellRendererPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCellRendererProgress.
*/
func CellRendererProgressNew() (return__ *CellRendererProgress) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_cell_renderer_progress_new()
	if __cgo__return__ != nil {
		return__ = NewCellRendererProgressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCellRendererSpin.
*/
func CellRendererSpinNew() (return__ *CellRendererSpin) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_cell_renderer_spin_new()
	if __cgo__return__ != nil {
		return__ = NewCellRendererSpinFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns a new cell renderer which will show a spinner to indicate
activity.
*/
func CellRendererSpinnerNew() (return__ *CellRendererSpinner) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_cell_renderer_spinner_new()
	if __cgo__return__ != nil {
		return__ = NewCellRendererSpinnerFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCellRendererText. Adjust how text is drawn using
object properties. Object properties can be
set globally (with g_object_set()). Also, with #GtkTreeViewColumn,
you can bind a property to a value in a #GtkTreeModel. For example,
you can bind the “text” property on the cell renderer to a string
value in the model, thus rendering a different string in each row
of the #GtkTreeView
*/
func CellRendererTextNew() (return__ *CellRendererText) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_cell_renderer_text_new()
	if __cgo__return__ != nil {
		return__ = NewCellRendererTextFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCellRendererToggle. Adjust rendering
parameters using object properties. Object properties can be set
globally (with g_object_set()). Also, with #GtkTreeViewColumn, you
can bind a property to a value in a #GtkTreeModel. For example, you
can bind the “active” property on the cell renderer to a boolean value
in the model, thus causing the check button to reflect the state of
the model.
*/
func CellRendererToggleNew() (return__ *CellRendererToggle) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_cell_renderer_toggle_new()
	if __cgo__return__ != nil {
		return__ = NewCellRendererToggleFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCellView widget.
*/
func CellViewNew() (return__ *CellView) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_cell_view_new()
	if __cgo__return__ != nil {
		return__ = NewCellViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCellView widget with a specific #GtkCellArea
to layout cells and a specific #GtkCellAreaContext.

Specifying the same context for a handfull of cells lets
the underlying area synchronize the geometry for those cells,
in this way alignments with cellviews for other rows are
possible.
*/
func CellViewNewWithContext(area IsCellArea, context IsCellAreaContext) (return__ *CellView) {
	var __cgo__area *C.GtkCellArea
	if area != nil {
		__cgo__area = area.GetCellAreaPointer()
	}
	var __cgo__context *C.GtkCellAreaContext
	if context != nil {
		__cgo__context = context.GetCellAreaContextPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_cell_view_new_with_context(__cgo__area, __cgo__context)
	if __cgo__return__ != nil {
		return__ = NewCellViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCellView widget, adds a #GtkCellRendererText
to it, and makes it show @markup. The text can be
marked up with the [Pango text markup language][PangoMarkupFormat].
*/
func CellViewNewWithMarkup(markup string) (return__ *CellView) {
	__cgo__markup := (*C.gchar)(unsafe.Pointer(C.CString(markup)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_cell_view_new_with_markup(__cgo__markup)
	C.free(unsafe.Pointer(__cgo__markup))
	if __cgo__return__ != nil {
		return__ = NewCellViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCellView widget, adds a #GtkCellRendererPixbuf
to it, and makes it show @pixbuf.
*/
func CellViewNewWithPixbuf(pixbuf *C.GdkPixbuf) (return__ *CellView) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_cell_view_new_with_pixbuf(pixbuf)
	if __cgo__return__ != nil {
		return__ = NewCellViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCellView widget, adds a #GtkCellRendererText
to it, and makes it show @text.
*/
func CellViewNewWithText(text string) (return__ *CellView) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_cell_view_new_with_text(__cgo__text)
	C.free(unsafe.Pointer(__cgo__text))
	if __cgo__return__ != nil {
		return__ = NewCellViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCheckButton.
*/
func CheckButtonNew() (return__ *CheckButton) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_check_button_new()
	if __cgo__return__ != nil {
		return__ = NewCheckButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCheckButton with a #GtkLabel to the right of it.
*/
func CheckButtonNewWithLabel(label string) (return__ *CheckButton) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_check_button_new_with_label(__cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewCheckButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCheckButton containing a label. The label
will be created using gtk_label_new_with_mnemonic(), so underscores
in @label indicate the mnemonic for the check button.
*/
func CheckButtonNewWithMnemonic(label string) (return__ *CheckButton) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_check_button_new_with_mnemonic(__cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewCheckButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCheckMenuItem.
*/
func CheckMenuItemNew() (return__ *CheckMenuItem) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_check_menu_item_new()
	if __cgo__return__ != nil {
		return__ = NewCheckMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCheckMenuItem with a label.
*/
func CheckMenuItemNewWithLabel(label string) (return__ *CheckMenuItem) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_check_menu_item_new_with_label(__cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewCheckMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkCheckMenuItem containing a label. The label
will be created using gtk_label_new_with_mnemonic(), so underscores
in @label indicate the mnemonic for the menu item.
*/
func CheckMenuItemNewWithMnemonic(label string) (return__ *CheckMenuItem) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_check_menu_item_new_with_mnemonic(__cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewCheckMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new color button.

This returns a widget in the form of a small button containing
a swatch representing the current selected color. When the button
is clicked, a color-selection dialog will open, allowing the user
to select a color. The swatch will be updated to reflect the new
color when the user finishes.
*/
func ColorButtonNew() (return__ *ColorButton) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_color_button_new()
	if __cgo__return__ != nil {
		return__ = NewColorButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_color_button_new_with_color is not generated due to deprecation attr

/*
Creates a new color button.
*/
func ColorButtonNewWithRgba(rgba *C.GdkRGBA) (return__ *ColorButton) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_color_button_new_with_rgba(rgba)
	if __cgo__return__ != nil {
		return__ = NewColorButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkColorChooserDialog.
*/
func ColorChooserDialogNew(title string, parent IsWindow) (return__ *ColorChooserDialog) {
	__cgo__title := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	var __cgo__parent *C.GtkWindow
	if parent != nil {
		__cgo__parent = parent.GetWindowPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_color_chooser_dialog_new(__cgo__title, __cgo__parent)
	C.free(unsafe.Pointer(__cgo__title))
	if __cgo__return__ != nil {
		return__ = NewColorChooserDialogFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkColorChooserWidget.
*/
func ColorChooserWidgetNew() (return__ *ColorChooserWidget) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_color_chooser_widget_new()
	if __cgo__return__ != nil {
		return__ = NewColorChooserWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_color_selection_new is not generated due to explicit ignore

// gtk_color_selection_dialog_new is not generated due to explicit ignore

/*
Creates a new empty #GtkComboBox.
*/
func ComboBoxNew() (return__ *ComboBox) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_combo_box_new()
	if __cgo__return__ != nil {
		return__ = NewComboBoxFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new empty #GtkComboBox using @area to layout cells.
*/
func ComboBoxNewWithArea(area IsCellArea) (return__ *ComboBox) {
	var __cgo__area *C.GtkCellArea
	if area != nil {
		__cgo__area = area.GetCellAreaPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_combo_box_new_with_area(__cgo__area)
	if __cgo__return__ != nil {
		return__ = NewComboBoxFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new empty #GtkComboBox with an entry.

The new combo box will use @area to layout cells.
*/
func ComboBoxNewWithAreaAndEntry(area IsCellArea) (return__ *ComboBox) {
	var __cgo__area *C.GtkCellArea
	if area != nil {
		__cgo__area = area.GetCellAreaPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_combo_box_new_with_area_and_entry(__cgo__area)
	if __cgo__return__ != nil {
		return__ = NewComboBoxFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new empty #GtkComboBox with an entry.
*/
func ComboBoxNewWithEntry() (return__ *ComboBox) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_combo_box_new_with_entry()
	if __cgo__return__ != nil {
		return__ = NewComboBoxFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkComboBox with the model initialized to @model.
*/
func ComboBoxNewWithModel(model *C.GtkTreeModel) (return__ *ComboBox) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_combo_box_new_with_model(model)
	if __cgo__return__ != nil {
		return__ = NewComboBoxFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new empty #GtkComboBox with an entry
and with the model initialized to @model.
*/
func ComboBoxNewWithModelAndEntry(model *C.GtkTreeModel) (return__ *ComboBox) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_combo_box_new_with_model_and_entry(model)
	if __cgo__return__ != nil {
		return__ = NewComboBoxFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkComboBoxText, which is a #GtkComboBox just displaying
strings.
*/
func ComboBoxTextNew() (return__ *ComboBoxText) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_combo_box_text_new()
	if __cgo__return__ != nil {
		return__ = NewComboBoxTextFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkComboBoxText, which is a #GtkComboBox just displaying
strings. The combo box created by this function has an entry.
*/
func ComboBoxTextNewWithEntry() (return__ *ComboBoxText) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_combo_box_text_new_with_entry()
	if __cgo__return__ != nil {
		return__ = NewComboBoxTextFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

func ContainerCellAccessibleNew() (return__ *ContainerCellAccessible) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_container_cell_accessible_new()
	if __cgo__return__ != nil {
		return__ = NewContainerCellAccessibleFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns a newly created #GtkCssProvider.
*/
func CssProviderNew() (return__ *CssProvider) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_css_provider_new()
	if __cgo__return__ != nil {
		return__ = NewCssProviderFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new dialog box.

Widgets should not be packed into this #GtkWindow
directly, but into the @vbox and @action_area, as described above.
*/
func DialogNew() (return__ *Dialog) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_dialog_new()
	if __cgo__return__ != nil {
		return__ = NewDialogFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_dialog_new_with_buttons is not generated due to varargs

/*
Creates a new drawing area.
*/
func DrawingAreaNew() (return__ *DrawingArea) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_drawing_area_new()
	if __cgo__return__ != nil {
		return__ = NewDrawingAreaFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new entry.
*/
func EntryNew() (return__ *Entry) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_entry_new()
	if __cgo__return__ != nil {
		return__ = NewEntryFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new entry with the specified text buffer.
*/
func EntryNewWithBuffer(buffer IsEntryBuffer) (return__ *Entry) {
	var __cgo__buffer *C.GtkEntryBuffer
	if buffer != nil {
		__cgo__buffer = buffer.GetEntryBufferPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_entry_new_with_buffer(__cgo__buffer)
	if __cgo__return__ != nil {
		return__ = NewEntryFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Create a new GtkEntryBuffer object.

Optionally, specify initial text to set in the buffer.
*/
func EntryBufferNew(initial_chars string, n_initial_chars int) (return__ *EntryBuffer) {
	__cgo__initial_chars := (*C.gchar)(unsafe.Pointer(C.CString(initial_chars)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_entry_buffer_new(__cgo__initial_chars, C.gint(n_initial_chars))
	C.free(unsafe.Pointer(__cgo__initial_chars))
	if __cgo__return__ != nil {
		return__ = NewEntryBufferFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkEntryCompletion object.
*/
func EntryCompletionNew() (return__ *EntryCompletion) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_entry_completion_new()
	if __cgo__return__ != nil {
		return__ = NewEntryCompletionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkEntryCompletion object using the
specified @area to layout cells in the underlying
#GtkTreeViewColumn for the drop-down menu.
*/
func EntryCompletionNewWithArea(area IsCellArea) (return__ *EntryCompletion) {
	var __cgo__area *C.GtkCellArea
	if area != nil {
		__cgo__area = area.GetCellAreaPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_entry_completion_new_with_area(__cgo__area)
	if __cgo__return__ != nil {
		return__ = NewEntryCompletionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkEventBox.
*/
func EventBoxNew() (return__ *EventBox) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_event_box_new()
	if __cgo__return__ != nil {
		return__ = NewEventBoxFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new expander using @label as the text of the label.
*/
func ExpanderNew(label string) (return__ *Expander) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_expander_new(__cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewExpanderFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new expander using @label as the text of the label.
If characters in @label are preceded by an underscore, they are underlined.
If you need a literal underscore character in a label, use “__” (two
underscores). The first underlined character represents a keyboard
accelerator called a mnemonic.
Pressing Alt and that key activates the button.
*/
func ExpanderNewWithMnemonic(label string) (return__ *Expander) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_expander_new_with_mnemonic(__cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewExpanderFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new file-selecting button widget.
*/
func FileChooserButtonNew(title string, action C.GtkFileChooserAction) (return__ *FileChooserButton) {
	__cgo__title := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_file_chooser_button_new(__cgo__title, action)
	C.free(unsafe.Pointer(__cgo__title))
	if __cgo__return__ != nil {
		return__ = NewFileChooserButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GtkFileChooserButton widget which uses @dialog as its
file-picking window.

Note that @dialog must be a #GtkDialog (or subclass) which
implements the #GtkFileChooser interface and must not have
%GTK_DIALOG_DESTROY_WITH_PARENT set.

Also note that the dialog needs to have its confirmative button
added with response %GTK_RESPONSE_ACCEPT or %GTK_RESPONSE_OK in
order for the button to take over the file selected in the dialog.
*/
func FileChooserButtonNewWithDialog(dialog *C.GtkWidget) (return__ *FileChooserButton) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_file_chooser_button_new_with_dialog(dialog)
	if __cgo__return__ != nil {
		return__ = NewFileChooserButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_file_chooser_dialog_new is not generated due to varargs

/*
Creates a new #GtkFileChooserWidget.  This is a file chooser widget that can
be embedded in custom windows, and it is the same widget that is used by
#GtkFileChooserDialog.
*/
func FileChooserWidgetNew(action C.GtkFileChooserAction) (return__ *FileChooserWidget) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_file_chooser_widget_new(action)
	if __cgo__return__ != nil {
		return__ = NewFileChooserWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkFileFilter with no rules added to it.
Such a filter doesn’t accept any files, so is not
particularly useful until you add rules with
gtk_file_filter_add_mime_type(), gtk_file_filter_add_pattern(),
or gtk_file_filter_add_custom(). To create a filter
that accepts any file, use:
|[<!-- language="C" -->
GtkFileFilter *filter = gtk_file_filter_new ();
gtk_file_filter_add_pattern (filter, "*");
]|
*/
func FileFilterNew() (return__ *FileFilter) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_file_filter_new()
	if __cgo__return__ != nil {
		return__ = NewFileFilterFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkFixed.
*/
func FixedNew() (return__ *Fixed) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_fixed_new()
	if __cgo__return__ != nil {
		return__ = NewFixedFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a GtkFlowBox.
*/
func FlowBoxNew() (return__ *FlowBox) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_flow_box_new()
	if __cgo__return__ != nil {
		return__ = NewFlowBoxFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkFlowBoxChild, to be used as a child
of a #GtkFlowBox.
*/
func FlowBoxChildNew() (return__ *FlowBoxChild) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_flow_box_child_new()
	if __cgo__return__ != nil {
		return__ = NewFlowBoxChildFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new font picker widget.
*/
func FontButtonNew() (return__ *FontButton) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_font_button_new()
	if __cgo__return__ != nil {
		return__ = NewFontButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new font picker widget.
*/
func FontButtonNewWithFont(fontname string) (return__ *FontButton) {
	__cgo__fontname := (*C.gchar)(unsafe.Pointer(C.CString(fontname)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_font_button_new_with_font(__cgo__fontname)
	C.free(unsafe.Pointer(__cgo__fontname))
	if __cgo__return__ != nil {
		return__ = NewFontButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkFontChooserDialog.
*/
func FontChooserDialogNew(title string, parent IsWindow) (return__ *FontChooserDialog) {
	__cgo__title := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	var __cgo__parent *C.GtkWindow
	if parent != nil {
		__cgo__parent = parent.GetWindowPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_font_chooser_dialog_new(__cgo__title, __cgo__parent)
	C.free(unsafe.Pointer(__cgo__title))
	if __cgo__return__ != nil {
		return__ = NewFontChooserDialogFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkFontChooserWidget.
*/
func FontChooserWidgetNew() (return__ *FontChooserWidget) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_font_chooser_widget_new()
	if __cgo__return__ != nil {
		return__ = NewFontChooserWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_font_selection_new is not generated due to deprecation attr

// gtk_font_selection_dialog_new is not generated due to deprecation attr

/*
Creates a new #GtkFrame, with optional label @label.
If @label is %NULL, the label is omitted.
*/
func FrameNew(label string) (return__ *Frame) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_frame_new(__cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewFrameFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns a newly created #GtkGesture that recognizes drags.
*/
func GestureDragNew(widget IsWidget) (return__ *GestureDrag) {
	var __cgo__widget *C.GtkWidget
	if widget != nil {
		__cgo__widget = widget.GetWidgetPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_gesture_drag_new(__cgo__widget)
	if __cgo__return__ != nil {
		return__ = NewGestureDragFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns a newly created #GtkGesture that recognizes long presses.
*/
func GestureLongPressNew(widget IsWidget) (return__ *GestureLongPress) {
	var __cgo__widget *C.GtkWidget
	if widget != nil {
		__cgo__widget = widget.GetWidgetPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_gesture_long_press_new(__cgo__widget)
	if __cgo__return__ != nil {
		return__ = NewGestureLongPressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns a newly created #GtkGesture that recognizes single and multiple
presses.
*/
func GestureMultiPressNew(widget IsWidget) (return__ *GestureMultiPress) {
	var __cgo__widget *C.GtkWidget
	if widget != nil {
		__cgo__widget = widget.GetWidgetPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_gesture_multi_press_new(__cgo__widget)
	if __cgo__return__ != nil {
		return__ = NewGestureMultiPressFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns a newly created #GtkGesture that recognizes pan gestures.
*/
func GesturePanNew(widget IsWidget, orientation C.GtkOrientation) (return__ *GesturePan) {
	var __cgo__widget *C.GtkWidget
	if widget != nil {
		__cgo__widget = widget.GetWidgetPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_gesture_pan_new(__cgo__widget, orientation)
	if __cgo__return__ != nil {
		return__ = NewGesturePanFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns a newly created #GtkGesture that recognizes 2-touch
rotation gestures.
*/
func GestureRotateNew(widget IsWidget) (return__ *GestureRotate) {
	var __cgo__widget *C.GtkWidget
	if widget != nil {
		__cgo__widget = widget.GetWidgetPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_gesture_rotate_new(__cgo__widget)
	if __cgo__return__ != nil {
		return__ = NewGestureRotateFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns a newly created #GtkGesture that recognizes swipes.
*/
func GestureSwipeNew(widget IsWidget) (return__ *GestureSwipe) {
	var __cgo__widget *C.GtkWidget
	if widget != nil {
		__cgo__widget = widget.GetWidgetPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_gesture_swipe_new(__cgo__widget)
	if __cgo__return__ != nil {
		return__ = NewGestureSwipeFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns a newly created #GtkGesture that recognizes zoom
in/out gestures (usually known as pinch/zoom).
*/
func GestureZoomNew(widget IsWidget) (return__ *GestureZoom) {
	var __cgo__widget *C.GtkWidget
	if widget != nil {
		__cgo__widget = widget.GetWidgetPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_gesture_zoom_new(__cgo__widget)
	if __cgo__return__ != nil {
		return__ = NewGestureZoomFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new grid widget.
*/
func GridNew() (return__ *Grid) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_grid_new()
	if __cgo__return__ != nil {
		return__ = NewGridFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_hbox_new is not generated due to deprecation attr

// gtk_hbutton_box_new is not generated due to deprecation attr

// gtk_hpaned_new is not generated due to deprecation attr

// gtk_hsv_new is not generated due to explicit ignore

// gtk_hscale_new is not generated due to deprecation attr

// gtk_hscale_new_with_range is not generated due to deprecation attr

// gtk_hscrollbar_new is not generated due to deprecation attr

// gtk_hseparator_new is not generated due to deprecation attr

// gtk_handle_box_new is not generated due to deprecation attr

/*
Creates a new #GtkHeaderBar widget.
*/
func HeaderBarNew() (return__ *HeaderBar) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_header_bar_new()
	if __cgo__return__ != nil {
		return__ = NewHeaderBarFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkIMContextSimple.
*/
func IMContextSimpleNew() (return__ *IMContextSimple) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_im_context_simple_new()
	if __cgo__return__ != nil {
		return__ = NewIMContextSimpleFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkIMMulticontext.
*/
func IMMulticontextNew() (return__ *IMMulticontext) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_im_multicontext_new()
	if __cgo__return__ != nil {
		return__ = NewIMMulticontextFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_icon_factory_new is not generated due to deprecation attr

/*
Creates a #GtkIconInfo for a #GdkPixbuf.
*/
func IconInfoNewForPixbuf(icon_theme IsIconTheme, pixbuf *C.GdkPixbuf) (return__ *IconInfo) {
	var __cgo__icon_theme *C.GtkIconTheme
	if icon_theme != nil {
		__cgo__icon_theme = icon_theme.GetIconThemePointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_icon_info_new_for_pixbuf(__cgo__icon_theme, pixbuf)
	if __cgo__return__ != nil {
		return__ = NewIconInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new icon theme object. Icon theme objects are used
to lookup up an icon by name in a particular icon theme.
Usually, you’ll want to use gtk_icon_theme_get_default()
or gtk_icon_theme_get_for_screen() rather than creating
a new icon theme object for scratch.
*/
func IconThemeNew() (return__ *IconTheme) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_icon_theme_new()
	if __cgo__return__ != nil {
		return__ = NewIconThemeFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkIconView widget
*/
func IconViewNew() (return__ *IconView) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_icon_view_new()
	if __cgo__return__ != nil {
		return__ = NewIconViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkIconView widget using the
specified @area to layout cells inside the icons.
*/
func IconViewNewWithArea(area IsCellArea) (return__ *IconView) {
	var __cgo__area *C.GtkCellArea
	if area != nil {
		__cgo__area = area.GetCellAreaPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_icon_view_new_with_area(__cgo__area)
	if __cgo__return__ != nil {
		return__ = NewIconViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkIconView widget with the model @model.
*/
func IconViewNewWithModel(model *C.GtkTreeModel) (return__ *IconView) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_icon_view_new_with_model(model)
	if __cgo__return__ != nil {
		return__ = NewIconViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new empty #GtkImage widget.
*/
func ImageNew() (return__ *Image) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_image_new()
	if __cgo__return__ != nil {
		return__ = NewImageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GtkImage displaying the given animation.
The #GtkImage does not assume a reference to the
animation; you still need to unref it if you own references.
#GtkImage will add its own reference rather than adopting yours.

Note that the animation frames are shown using a timeout with
#G_PRIORITY_DEFAULT. When using animations to indicate busyness,
keep in mind that the animation will only be shown if the main loop
is not busy with something that has a higher priority.
*/
func ImageNewFromAnimation(animation *C.GdkPixbufAnimation) (return__ *Image) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_image_new_from_animation(animation)
	if __cgo__return__ != nil {
		return__ = NewImageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkImage displaying the file @filename. If the file
isn’t found or can’t be loaded, the resulting #GtkImage will
display a “broken image” icon. This function never returns %NULL,
it always returns a valid #GtkImage widget.

If the file contains an animation, the image will contain an
animation.

If you need to detect failures to load the file, use
gdk_pixbuf_new_from_file() to load the file yourself, then create
the #GtkImage from the pixbuf. (Or for animations, use
gdk_pixbuf_animation_new_from_file()).

The storage type (gtk_image_get_storage_type()) of the returned
image is not defined, it will be whatever is appropriate for
displaying the file.
*/
func ImageNewFromFile(filename string) (return__ *Image) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_image_new_from_file(__cgo__filename)
	C.free(unsafe.Pointer(__cgo__filename))
	if __cgo__return__ != nil {
		return__ = NewImageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GtkImage displaying an icon from the current icon theme.
If the icon name isn’t known, a “broken image” icon will be
displayed instead.  If the current icon theme is changed, the icon
will be updated appropriately.
*/
func ImageNewFromGicon(icon *C.GIcon, size C.GtkIconSize) (return__ *Image) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_image_new_from_gicon(icon, size)
	if __cgo__return__ != nil {
		return__ = NewImageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GtkImage displaying an icon from the current icon theme.
If the icon name isn’t known, a “broken image” icon will be
displayed instead.  If the current icon theme is changed, the icon
will be updated appropriately.
*/
func ImageNewFromIconName(icon_name string, size C.GtkIconSize) (return__ *Image) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_image_new_from_icon_name(__cgo__icon_name, size)
	C.free(unsafe.Pointer(__cgo__icon_name))
	if __cgo__return__ != nil {
		return__ = NewImageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_image_new_from_icon_set is not generated due to deprecation attr

/*
Creates a new #GtkImage displaying @pixbuf.
The #GtkImage does not assume a reference to the
pixbuf; you still need to unref it if you own references.
#GtkImage will add its own reference rather than adopting yours.

Note that this function just creates an #GtkImage from the pixbuf. The
#GtkImage created will not react to state changes. Should you want that,
you should use gtk_image_new_from_icon_name().
*/
func ImageNewFromPixbuf(pixbuf *C.GdkPixbuf) (return__ *Image) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_image_new_from_pixbuf(pixbuf)
	if __cgo__return__ != nil {
		return__ = NewImageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkImage displaying the resource file @resource_path. If the file
isn’t found or can’t be loaded, the resulting #GtkImage will
display a “broken image” icon. This function never returns %NULL,
it always returns a valid #GtkImage widget.

If the file contains an animation, the image will contain an
animation.

If you need to detect failures to load the file, use
gdk_pixbuf_new_from_file() to load the file yourself, then create
the #GtkImage from the pixbuf. (Or for animations, use
gdk_pixbuf_animation_new_from_file()).

The storage type (gtk_image_get_storage_type()) of the returned
image is not defined, it will be whatever is appropriate for
displaying the file.
*/
func ImageNewFromResource(resource_path string) (return__ *Image) {
	__cgo__resource_path := (*C.gchar)(unsafe.Pointer(C.CString(resource_path)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_image_new_from_resource(__cgo__resource_path)
	C.free(unsafe.Pointer(__cgo__resource_path))
	if __cgo__return__ != nil {
		return__ = NewImageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_image_new_from_stock is not generated due to deprecation attr

/*
Creates a new #GtkImage displaying @surface.
The #GtkImage does not assume a reference to the
surface; you still need to unref it if you own references.
#GtkImage will add its own reference rather than adopting yours.
*/
func ImageNewFromSurface(surface *C.cairo_surface_t) (return__ *Image) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_image_new_from_surface(surface)
	if __cgo__return__ != nil {
		return__ = NewImageFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_image_menu_item_new is not generated due to deprecation attr

// gtk_image_menu_item_new_from_stock is not generated due to deprecation attr

// gtk_image_menu_item_new_with_label is not generated due to deprecation attr

// gtk_image_menu_item_new_with_mnemonic is not generated due to deprecation attr

/*
Creates a new #GtkInfoBar object.
*/
func InfoBarNew() (return__ *InfoBar) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_info_bar_new()
	if __cgo__return__ != nil {
		return__ = NewInfoBarFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_info_bar_new_with_buttons is not generated due to varargs

/*
Creates a new #GtkInvisible.
*/
func InvisibleNew() (return__ *Invisible) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_invisible_new()
	if __cgo__return__ != nil {
		return__ = NewInvisibleFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkInvisible object for a specified screen
*/
func InvisibleNewForScreen(screen *C.GdkScreen) (return__ *Invisible) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_invisible_new_for_screen(screen)
	if __cgo__return__ != nil {
		return__ = NewInvisibleFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new label with the given text inside it. You can
pass %NULL to get an empty label widget.
*/
func LabelNew(str string) (return__ *Label) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_label_new(__cgo__str)
	C.free(unsafe.Pointer(__cgo__str))
	if __cgo__return__ != nil {
		return__ = NewLabelFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkLabel, containing the text in @str.

If characters in @str are preceded by an underscore, they are
underlined. If you need a literal underscore character in a label, use
'__' (two underscores). The first underlined character represents a
keyboard accelerator called a mnemonic. The mnemonic key can be used
to activate another widget, chosen automatically, or explicitly using
gtk_label_set_mnemonic_widget().

If gtk_label_set_mnemonic_widget() is not called, then the first
activatable ancestor of the #GtkLabel will be chosen as the mnemonic
widget. For instance, if the label is inside a button or menu item,
the button or menu item will automatically become the mnemonic widget
and be activated by the mnemonic.
*/
func LabelNewWithMnemonic(str string) (return__ *Label) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_label_new_with_mnemonic(__cgo__str)
	C.free(unsafe.Pointer(__cgo__str))
	if __cgo__return__ != nil {
		return__ = NewLabelFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkLayout. Unless you have a specific adjustment
you’d like the layout to use for scrolling, pass %NULL for
@hadjustment and @vadjustment.
*/
func LayoutNew(hadjustment IsAdjustment, vadjustment IsAdjustment) (return__ *Layout) {
	var __cgo__hadjustment *C.GtkAdjustment
	if hadjustment != nil {
		__cgo__hadjustment = hadjustment.GetAdjustmentPointer()
	}
	var __cgo__vadjustment *C.GtkAdjustment
	if vadjustment != nil {
		__cgo__vadjustment = vadjustment.GetAdjustmentPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_layout_new(__cgo__hadjustment, __cgo__vadjustment)
	if __cgo__return__ != nil {
		return__ = NewLayoutFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkLevelBar.
*/
func LevelBarNew() (return__ *LevelBar) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_level_bar_new()
	if __cgo__return__ != nil {
		return__ = NewLevelBarFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Utility constructor that creates a new #GtkLevelBar for the specified
interval.
*/
func LevelBarNewForInterval(min_value float64, max_value float64) (return__ *LevelBar) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_level_bar_new_for_interval(C.gdouble(min_value), C.gdouble(max_value))
	if __cgo__return__ != nil {
		return__ = NewLevelBarFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkLinkButton with the URI as its text.
*/
func LinkButtonNew(uri string) (return__ *LinkButton) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_link_button_new(__cgo__uri)
	C.free(unsafe.Pointer(__cgo__uri))
	if __cgo__return__ != nil {
		return__ = NewLinkButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkLinkButton containing a label.
*/
func LinkButtonNewWithLabel(uri string, label string) (return__ *LinkButton) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_link_button_new_with_label(__cgo__uri, __cgo__label)
	C.free(unsafe.Pointer(__cgo__uri))
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewLinkButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkListBox container.
*/
func ListBoxNew() (return__ *ListBox) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_list_box_new()
	if __cgo__return__ != nil {
		return__ = NewListBoxFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkListBoxRow, to be used as a child of a #GtkListBox.
*/
func ListBoxRowNew() (return__ *ListBoxRow) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_list_box_row_new()
	if __cgo__return__ != nil {
		return__ = NewListBoxRowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_list_store_new is not generated due to varargs

/*
Non-vararg creation function.  Used primarily by language bindings.
*/
func ListStoreNewv(n_columns int, types *C.GType) (return__ *ListStore) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_list_store_newv(C.gint(n_columns), types)
	if __cgo__return__ != nil {
		return__ = NewListStoreFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new lock button which reflects the @permission.
*/
func LockButtonNew(permission *C.GPermission) (return__ *LockButton) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_lock_button_new(permission)
	if __cgo__return__ != nil {
		return__ = NewLockButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkMenu
*/
func MenuNew() (return__ *Menu) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_menu_new()
	if __cgo__return__ != nil {
		return__ = NewMenuFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GtkMenu and populates it with menu items and
submenus according to @model.

The created menu items are connected to actions found in the
#GtkApplicationWindow to which the menu belongs - typically
by means of being attached to a widget (see gtk_menu_attach_to_widget())
that is contained within the #GtkApplicationWindows widget hierarchy.

Actions can also be added using gtk_widget_insert_action_group() on the menu's
attach widget or on any of its parent widgets.
*/
func MenuNewFromModel(model *C.GMenuModel) (return__ *Menu) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_menu_new_from_model(model)
	if __cgo__return__ != nil {
		return__ = NewMenuFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkMenuBar
*/
func MenuBarNew() (return__ *MenuBar) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_menu_bar_new()
	if __cgo__return__ != nil {
		return__ = NewMenuBarFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkMenuBar and populates it with menu items
and submenus according to @model.

The created menu items are connected to actions found in the
#GtkApplicationWindow to which the menu bar belongs - typically
by means of being contained within the #GtkApplicationWindows
widget hierarchy.
*/
func MenuBarNewFromModel(model *C.GMenuModel) (return__ *MenuBar) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_menu_bar_new_from_model(model)
	if __cgo__return__ != nil {
		return__ = NewMenuBarFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkMenuButton widget with downwards-pointing
arrow as the only child. You can replace the child widget
with another #GtkWidget should you wish to.
*/
func MenuButtonNew() (return__ *MenuButton) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_menu_button_new()
	if __cgo__return__ != nil {
		return__ = NewMenuButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkMenuItem.
*/
func MenuItemNew() (return__ *MenuItem) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_menu_item_new()
	if __cgo__return__ != nil {
		return__ = NewMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkMenuItem whose child is a #GtkLabel.
*/
func MenuItemNewWithLabel(label string) (return__ *MenuItem) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_menu_item_new_with_label(__cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkMenuItem containing a label.

The label will be created using gtk_label_new_with_mnemonic(),
so underscores in @label indicate the mnemonic for the menu item.
*/
func MenuItemNewWithMnemonic(label string) (return__ *MenuItem) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_menu_item_new_with_mnemonic(__cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkMenuToolButton using @icon_widget as icon and
@label as label.
*/
func MenuToolButtonNew(icon_widget IsWidget, label string) (return__ *MenuToolButton) {
	var __cgo__icon_widget *C.GtkWidget
	if icon_widget != nil {
		__cgo__icon_widget = icon_widget.GetWidgetPointer()
	}
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_menu_tool_button_new(__cgo__icon_widget, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewMenuToolButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_menu_tool_button_new_from_stock is not generated due to deprecation attr

// gtk_message_dialog_new is not generated due to varargs

// gtk_message_dialog_new_with_markup is not generated due to varargs

/*
Creates a new #GtkMountOperation
*/
func MountOperationNew(parent IsWindow) (return__ *MountOperation) {
	var __cgo__parent *C.GtkWindow
	if parent != nil {
		__cgo__parent = parent.GetWindowPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_mount_operation_new(__cgo__parent)
	if __cgo__return__ != nil {
		return__ = NewMountOperationFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkNotebook widget with no pages.
*/
func NotebookNew() (return__ *Notebook) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_notebook_new()
	if __cgo__return__ != nil {
		return__ = NewNotebookFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

func NotebookPageAccessibleNew(notebook IsNotebookAccessible, child IsWidget) (return__ *NotebookPageAccessible) {
	var __cgo__notebook *C.GtkNotebookAccessible
	if notebook != nil {
		__cgo__notebook = notebook.GetNotebookAccessiblePointer()
	}
	var __cgo__child *C.GtkWidget
	if child != nil {
		__cgo__child = child.GetWidgetPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_notebook_page_accessible_new(__cgo__notebook, __cgo__child)
	if __cgo__return__ != nil {
		return__ = NewNotebookPageAccessibleFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a toplevel container widget that is used to retrieve
snapshots of widgets without showing them on the screen.
*/
func OffscreenWindowNew() (return__ *OffscreenWindow) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_offscreen_window_new()
	if __cgo__return__ != nil {
		return__ = NewOffscreenWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkOverlay.
*/
func OverlayNew() (return__ *Overlay) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_overlay_new()
	if __cgo__return__ != nil {
		return__ = NewOverlayFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkPageSetup.
*/
func PageSetupNew() (return__ *PageSetup) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_page_setup_new()
	if __cgo__return__ != nil {
		return__ = NewPageSetupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Reads the page setup from the file @file_name. Returns a
new #GtkPageSetup object with the restored page setup,
or %NULL if an error occurred. See gtk_page_setup_to_file().
*/
func PageSetupNewFromFile(file_name string) (return__ *PageSetup, __err__ error) {
	__cgo__file_name := (*C.gchar)(unsafe.Pointer(C.CString(file_name)))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_page_setup_new_from_file(__cgo__file_name, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__file_name))
	if __cgo__return__ != nil {
		return__ = NewPageSetupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Reads the page setup from the group @group_name in the key file
@key_file. Returns a new #GtkPageSetup object with the restored
page setup, or %NULL if an error occurred.
*/
func PageSetupNewFromKeyFile(key_file *C.GKeyFile, group_name string) (return__ *PageSetup, __err__ error) {
	__cgo__group_name := (*C.gchar)(unsafe.Pointer(C.CString(group_name)))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_page_setup_new_from_key_file(key_file, __cgo__group_name, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__group_name))
	if __cgo__return__ != nil {
		return__ = NewPageSetupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new #GtkPaned widget.
*/
func PanedNew(orientation C.GtkOrientation) (return__ *Paned) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_paned_new(orientation)
	if __cgo__return__ != nil {
		return__ = NewPanedFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkPlacesSidebar widget.

The application should connect to at least the
#GtkPlacesSidebar::open-location signal to be notified
when the user makes a selection in the sidebar.
*/
func PlacesSidebarNew() (return__ *PlacesSidebar) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_places_sidebar_new()
	if __cgo__return__ != nil {
		return__ = NewPlacesSidebarFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new plug widget inside the #GtkSocket identified
by @socket_id. If @socket_id is 0, the plug is left “unplugged” and
can later be plugged into a #GtkSocket by  gtk_socket_add_id().
*/
func PlugNew(socket_id C.Window) (return__ *Plug) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_plug_new(socket_id)
	if __cgo__return__ != nil {
		return__ = NewPlugFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Create a new plug widget inside the #GtkSocket identified by socket_id.
*/
func PlugNewForDisplay(display *C.GdkDisplay, socket_id C.Window) (return__ *Plug) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_plug_new_for_display(display, socket_id)
	if __cgo__return__ != nil {
		return__ = NewPlugFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new popover to point to @relative_to
*/
func PopoverNew(relative_to IsWidget) (return__ *Popover) {
	var __cgo__relative_to *C.GtkWidget
	if relative_to != nil {
		__cgo__relative_to = relative_to.GetWidgetPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_popover_new(__cgo__relative_to)
	if __cgo__return__ != nil {
		return__ = NewPopoverFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GtkPopover and populates it according to
@model. The popover is pointed to the @relative_to widget.

The created buttons are connected to actions found in the
#GtkApplicationWindow to which the popover belongs - typically
by means of being attached to a widget that is contained within
the #GtkApplicationWindows widget hierarchy.

Actions can also be added using gtk_widget_insert_action_group()
on the menus attach widget or on any of its parent widgets.
*/
func PopoverNewFromModel(relative_to IsWidget, model *C.GMenuModel) (return__ *Popover) {
	var __cgo__relative_to *C.GtkWidget
	if relative_to != nil {
		__cgo__relative_to = relative_to.GetWidgetPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_popover_new_from_model(__cgo__relative_to, model)
	if __cgo__return__ != nil {
		return__ = NewPopoverFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkPrintOperation.
*/
func PrintOperationNew() (return__ *PrintOperation) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_print_operation_new()
	if __cgo__return__ != nil {
		return__ = NewPrintOperationFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkPrintSettings object.
*/
func PrintSettingsNew() (return__ *PrintSettings) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_print_settings_new()
	if __cgo__return__ != nil {
		return__ = NewPrintSettingsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Reads the print settings from @file_name. Returns a new #GtkPrintSettings
object with the restored settings, or %NULL if an error occurred. If the
file could not be loaded then error is set to either a #GFileError or
#GKeyFileError.  See gtk_print_settings_to_file().
*/
func PrintSettingsNewFromFile(file_name string) (return__ *PrintSettings, __err__ error) {
	__cgo__file_name := (*C.gchar)(unsafe.Pointer(C.CString(file_name)))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_print_settings_new_from_file(__cgo__file_name, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__file_name))
	if __cgo__return__ != nil {
		return__ = NewPrintSettingsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Reads the print settings from the group @group_name in @key_file.  Returns a
new #GtkPrintSettings object with the restored settings, or %NULL if an
error occurred. If the file could not be loaded then error is set to either
a #GFileError or #GKeyFileError.
*/
func PrintSettingsNewFromKeyFile(key_file *C.GKeyFile, group_name string) (return__ *PrintSettings, __err__ error) {
	__cgo__group_name := (*C.gchar)(unsafe.Pointer(C.CString(group_name)))
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_print_settings_new_from_key_file(key_file, __cgo__group_name, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__group_name))
	if __cgo__return__ != nil {
		return__ = NewPrintSettingsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new #GtkProgressBar.
*/
func ProgressBarNew() (return__ *ProgressBar) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_progress_bar_new()
	if __cgo__return__ != nil {
		return__ = NewProgressBarFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_radio_action_new is not generated due to deprecation attr

/*
Creates a new #GtkRadioButton. To be of any practical value, a widget should
then be packed into the radio button.
*/
func RadioButtonNew(group *C.GSList) (return__ *RadioButton) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_radio_button_new(group)
	if __cgo__return__ != nil {
		return__ = NewRadioButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkRadioButton, adding it to the same group as
@radio_group_member. As with gtk_radio_button_new(), a widget
should be packed into the radio button.
*/
func RadioButtonNewFromWidget(radio_group_member IsRadioButton) (return__ *RadioButton) {
	var __cgo__radio_group_member *C.GtkRadioButton
	if radio_group_member != nil {
		__cgo__radio_group_member = radio_group_member.GetRadioButtonPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_radio_button_new_from_widget(__cgo__radio_group_member)
	if __cgo__return__ != nil {
		return__ = NewRadioButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkRadioButton with a text label.
*/
func RadioButtonNewWithLabel(group *C.GSList, label string) (return__ *RadioButton) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_radio_button_new_with_label(group, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewRadioButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkRadioButton with a text label, adding it to
the same group as @radio_group_member.
*/
func RadioButtonNewWithLabelFromWidget(radio_group_member IsRadioButton, label string) (return__ *RadioButton) {
	var __cgo__radio_group_member *C.GtkRadioButton
	if radio_group_member != nil {
		__cgo__radio_group_member = radio_group_member.GetRadioButtonPointer()
	}
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_radio_button_new_with_label_from_widget(__cgo__radio_group_member, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewRadioButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkRadioButton containing a label, adding it to the same
group as @group. The label will be created using
gtk_label_new_with_mnemonic(), so underscores in @label indicate the
mnemonic for the button.
*/
func RadioButtonNewWithMnemonic(group *C.GSList, label string) (return__ *RadioButton) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_radio_button_new_with_mnemonic(group, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewRadioButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkRadioButton containing a label. The label
will be created using gtk_label_new_with_mnemonic(), so underscores
in @label indicate the mnemonic for the button.
*/
func RadioButtonNewWithMnemonicFromWidget(radio_group_member IsRadioButton, label string) (return__ *RadioButton) {
	var __cgo__radio_group_member *C.GtkRadioButton
	if radio_group_member != nil {
		__cgo__radio_group_member = radio_group_member.GetRadioButtonPointer()
	}
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_radio_button_new_with_mnemonic_from_widget(__cgo__radio_group_member, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewRadioButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkRadioMenuItem.
*/
func RadioMenuItemNew(group *C.GSList) (return__ *RadioMenuItem) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_radio_menu_item_new(group)
	if __cgo__return__ != nil {
		return__ = NewRadioMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkRadioMenuItem adding it to the same group as @group.
*/
func RadioMenuItemNewFromWidget(group IsRadioMenuItem) (return__ *RadioMenuItem) {
	var __cgo__group *C.GtkRadioMenuItem
	if group != nil {
		__cgo__group = group.GetRadioMenuItemPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_radio_menu_item_new_from_widget(__cgo__group)
	if __cgo__return__ != nil {
		return__ = NewRadioMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkRadioMenuItem whose child is a simple #GtkLabel.
*/
func RadioMenuItemNewWithLabel(group *C.GSList, label string) (return__ *RadioMenuItem) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_radio_menu_item_new_with_label(group, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewRadioMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new GtkRadioMenuItem whose child is a simple GtkLabel.
The new #GtkRadioMenuItem is added to the same group as @group.
*/
func RadioMenuItemNewWithLabelFromWidget(group IsRadioMenuItem, label string) (return__ *RadioMenuItem) {
	var __cgo__group *C.GtkRadioMenuItem
	if group != nil {
		__cgo__group = group.GetRadioMenuItemPointer()
	}
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_radio_menu_item_new_with_label_from_widget(__cgo__group, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewRadioMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkRadioMenuItem containing a label. The label
will be created using gtk_label_new_with_mnemonic(), so underscores
in @label indicate the mnemonic for the menu item.
*/
func RadioMenuItemNewWithMnemonic(group *C.GSList, label string) (return__ *RadioMenuItem) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_radio_menu_item_new_with_mnemonic(group, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewRadioMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new GtkRadioMenuItem containing a label. The label will be
created using gtk_label_new_with_mnemonic(), so underscores in label
indicate the mnemonic for the menu item.

The new #GtkRadioMenuItem is added to the same group as @group.
*/
func RadioMenuItemNewWithMnemonicFromWidget(group IsRadioMenuItem, label string) (return__ *RadioMenuItem) {
	var __cgo__group *C.GtkRadioMenuItem
	if group != nil {
		__cgo__group = group.GetRadioMenuItemPointer()
	}
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_radio_menu_item_new_with_mnemonic_from_widget(__cgo__group, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewRadioMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkRadioToolButton, adding it to @group.
*/
func RadioToolButtonNew(group *C.GSList) (return__ *RadioToolButton) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_radio_tool_button_new(group)
	if __cgo__return__ != nil {
		return__ = NewRadioToolButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_radio_tool_button_new_from_stock is not generated due to deprecation attr

/*
Creates a new #GtkRadioToolButton adding it to the same group as @gruup
*/
func RadioToolButtonNewFromWidget(group IsRadioToolButton) (return__ *RadioToolButton) {
	var __cgo__group *C.GtkRadioToolButton
	if group != nil {
		__cgo__group = group.GetRadioToolButtonPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_radio_tool_button_new_from_widget(__cgo__group)
	if __cgo__return__ != nil {
		return__ = NewRadioToolButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_radio_tool_button_new_with_stock_from_widget is not generated due to deprecation attr

// gtk_rc_style_new is not generated due to deprecation attr

// gtk_recent_action_new is not generated due to deprecation attr

// gtk_recent_action_new_for_manager is not generated due to deprecation attr

// gtk_recent_chooser_dialog_new is not generated due to varargs

// gtk_recent_chooser_dialog_new_for_manager is not generated due to varargs

/*
Creates a new #GtkRecentChooserMenu widget.

This kind of widget shows the list of recently used resources as
a menu, each item as a menu item.  Each item inside the menu might
have an icon, representing its MIME type, and a number, for mnemonic
access.

This widget implements the #GtkRecentChooser interface.

This widget creates its own #GtkRecentManager object.  See the
gtk_recent_chooser_menu_new_for_manager() function to know how to create
a #GtkRecentChooserMenu widget bound to another #GtkRecentManager object.
*/
func RecentChooserMenuNew() (return__ *RecentChooserMenu) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_recent_chooser_menu_new()
	if __cgo__return__ != nil {
		return__ = NewRecentChooserMenuFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkRecentChooserMenu widget using @manager as
the underlying recently used resources manager.

This is useful if you have implemented your own recent manager,
or if you have a customized instance of a #GtkRecentManager
object or if you wish to share a common #GtkRecentManager object
among multiple #GtkRecentChooser widgets.
*/
func RecentChooserMenuNewForManager(manager IsRecentManager) (return__ *RecentChooserMenu) {
	var __cgo__manager *C.GtkRecentManager
	if manager != nil {
		__cgo__manager = manager.GetRecentManagerPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_recent_chooser_menu_new_for_manager(__cgo__manager)
	if __cgo__return__ != nil {
		return__ = NewRecentChooserMenuFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkRecentChooserWidget object.  This is an embeddable widget
used to access the recently used resources list.
*/
func RecentChooserWidgetNew() (return__ *RecentChooserWidget) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_recent_chooser_widget_new()
	if __cgo__return__ != nil {
		return__ = NewRecentChooserWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkRecentChooserWidget with a specified recent manager.

This is useful if you have implemented your own recent manager, or if you
have a customized instance of a #GtkRecentManager object.
*/
func RecentChooserWidgetNewForManager(manager IsRecentManager) (return__ *RecentChooserWidget) {
	var __cgo__manager *C.GtkRecentManager
	if manager != nil {
		__cgo__manager = manager.GetRecentManagerPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_recent_chooser_widget_new_for_manager(__cgo__manager)
	if __cgo__return__ != nil {
		return__ = NewRecentChooserWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkRecentFilter with no rules added to it.
Such filter does not accept any recently used resources, so is not
particularly useful until you add rules with
gtk_recent_filter_add_pattern(), gtk_recent_filter_add_mime_type(),
gtk_recent_filter_add_application(), gtk_recent_filter_add_age().
To create a filter that accepts any recently used resource, use:
|[<!-- language="C" -->
GtkRecentFilter *filter = gtk_recent_filter_new ();
gtk_recent_filter_add_pattern (filter, "*");
]|
*/
func RecentFilterNew() (return__ *RecentFilter) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_recent_filter_new()
	if __cgo__return__ != nil {
		return__ = NewRecentFilterFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new recent manager object.  Recent manager objects are used to
handle the list of recently used resources.  A #GtkRecentManager object
monitors the recently used resources list, and emits the “changed” signal
each time something inside the list changes.

#GtkRecentManager objects are expensive: be sure to create them only when
needed. You should use gtk_recent_manager_get_default() instead.
*/
func RecentManagerNew() (return__ *RecentManager) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_recent_manager_new()
	if __cgo__return__ != nil {
		return__ = NewRecentManagerFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

func RendererCellAccessibleNew(renderer IsCellRenderer) (return__ *RendererCellAccessible) {
	var __cgo__renderer *C.GtkCellRenderer
	if renderer != nil {
		__cgo__renderer = renderer.GetCellRendererPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_renderer_cell_accessible_new(__cgo__renderer)
	if __cgo__return__ != nil {
		return__ = NewRendererCellAccessibleFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkRevealer.
*/
func RevealerNew() (return__ *Revealer) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_revealer_new()
	if __cgo__return__ != nil {
		return__ = NewRevealerFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkScale.
*/
func ScaleNew(orientation C.GtkOrientation, adjustment IsAdjustment) (return__ *Scale) {
	var __cgo__adjustment *C.GtkAdjustment
	if adjustment != nil {
		__cgo__adjustment = adjustment.GetAdjustmentPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_scale_new(orientation, __cgo__adjustment)
	if __cgo__return__ != nil {
		return__ = NewScaleFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new scale widget with the given orientation that lets the
user input a number between @min and @max (including @min and @max)
with the increment @step.  @step must be nonzero; it’s the distance
the slider moves when using the arrow keys to adjust the scale
value.

Note that the way in which the precision is derived works best if @step
is a power of ten. If the resulting precision is not suitable for your
needs, use gtk_scale_set_digits() to correct it.
*/
func ScaleNewWithRange(orientation C.GtkOrientation, min float64, max float64, step float64) (return__ *Scale) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_scale_new_with_range(orientation, C.gdouble(min), C.gdouble(max), C.gdouble(step))
	if __cgo__return__ != nil {
		return__ = NewScaleFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GtkScaleButton, with a range between @min and @max, with
a stepping of @step.
*/
func ScaleButtonNew(size C.GtkIconSize, min float64, max float64, step float64, icons []string) (return__ *ScaleButton) {
	__header__icons := (*reflect.SliceHeader)(unsafe.Pointer(&icons))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_scale_button_new(size, C.gdouble(min), C.gdouble(max), C.gdouble(step), (**C.gchar)(unsafe.Pointer(__header__icons.Data)))
	if __cgo__return__ != nil {
		return__ = NewScaleButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new scrollbar with the given orientation.
*/
func ScrollbarNew(orientation C.GtkOrientation, adjustment IsAdjustment) (return__ *Scrollbar) {
	var __cgo__adjustment *C.GtkAdjustment
	if adjustment != nil {
		__cgo__adjustment = adjustment.GetAdjustmentPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_scrollbar_new(orientation, __cgo__adjustment)
	if __cgo__return__ != nil {
		return__ = NewScrollbarFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new scrolled window.

The two arguments are the scrolled window’s adjustments; these will be
shared with the scrollbars and the child widget to keep the bars in sync
with the child. Usually you want to pass %NULL for the adjustments, which
will cause the scrolled window to create them for you.
*/
func ScrolledWindowNew(hadjustment IsAdjustment, vadjustment IsAdjustment) (return__ *ScrolledWindow) {
	var __cgo__hadjustment *C.GtkAdjustment
	if hadjustment != nil {
		__cgo__hadjustment = hadjustment.GetAdjustmentPointer()
	}
	var __cgo__vadjustment *C.GtkAdjustment
	if vadjustment != nil {
		__cgo__vadjustment = vadjustment.GetAdjustmentPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_scrolled_window_new(__cgo__hadjustment, __cgo__vadjustment)
	if __cgo__return__ != nil {
		return__ = NewScrolledWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GtkSearchBar. You will need to tell it about
which widget is going to be your text entry using
gtk_search_bar_connect_entry().
*/
func SearchBarNew() (return__ *SearchBar) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_search_bar_new()
	if __cgo__return__ != nil {
		return__ = NewSearchBarFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GtkSearchEntry, with a find icon when the search field is
empty, and a clear icon when it isn't.
*/
func SearchEntryNew() (return__ *SearchEntry) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_search_entry_new()
	if __cgo__return__ != nil {
		return__ = NewSearchEntryFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkSeparator with the given orientation.
*/
func SeparatorNew(orientation C.GtkOrientation) (return__ *Separator) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_separator_new(orientation)
	if __cgo__return__ != nil {
		return__ = NewSeparatorFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkSeparatorMenuItem.
*/
func SeparatorMenuItemNew() (return__ *SeparatorMenuItem) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_separator_menu_item_new()
	if __cgo__return__ != nil {
		return__ = NewSeparatorMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Create a new #GtkSeparatorToolItem
*/
func SeparatorToolItemNew() (return__ *SeparatorToolItem) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_separator_tool_item_new()
	if __cgo__return__ != nil {
		return__ = NewSeparatorToolItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Create a new #GtkSizeGroup.
*/
func SizeGroupNew(mode C.GtkSizeGroupMode) (return__ *SizeGroup) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_size_group_new(mode)
	if __cgo__return__ != nil {
		return__ = NewSizeGroupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Create a new empty #GtkSocket.
*/
func SocketNew() (return__ *Socket) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_socket_new()
	if __cgo__return__ != nil {
		return__ = NewSocketFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkSpinButton.
*/
func SpinButtonNew(adjustment IsAdjustment, climb_rate float64, digits uint) (return__ *SpinButton) {
	var __cgo__adjustment *C.GtkAdjustment
	if adjustment != nil {
		__cgo__adjustment = adjustment.GetAdjustmentPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_spin_button_new(__cgo__adjustment, C.gdouble(climb_rate), C.guint(digits))
	if __cgo__return__ != nil {
		return__ = NewSpinButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
This is a convenience constructor that allows creation of a numeric
#GtkSpinButton without manually creating an adjustment. The value is
initially set to the minimum value and a page increment of 10 * @step
is the default. The precision of the spin button is equivalent to the
precision of @step.

Note that the way in which the precision is derived works best if @step
is a power of ten. If the resulting precision is not suitable for your
needs, use gtk_spin_button_set_digits() to correct it.
*/
func SpinButtonNewWithRange(min float64, max float64, step float64) (return__ *SpinButton) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_spin_button_new_with_range(C.gdouble(min), C.gdouble(max), C.gdouble(step))
	if __cgo__return__ != nil {
		return__ = NewSpinButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns a new spinner widget. Not yet started.
*/
func SpinnerNew() (return__ *Spinner) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_spinner_new()
	if __cgo__return__ != nil {
		return__ = NewSpinnerFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkStack container.
*/
func StackNew() (return__ *Stack) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_stack_new()
	if __cgo__return__ != nil {
		return__ = NewStackFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Create a new #GtkStackSwitcher.
*/
func StackSwitcherNew() (return__ *StackSwitcher) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_stack_switcher_new()
	if __cgo__return__ != nil {
		return__ = NewStackSwitcherFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_status_icon_new is not generated due to deprecation attr

// gtk_status_icon_new_from_file is not generated due to deprecation attr

// gtk_status_icon_new_from_gicon is not generated due to deprecation attr

// gtk_status_icon_new_from_icon_name is not generated due to deprecation attr

// gtk_status_icon_new_from_pixbuf is not generated due to deprecation attr

// gtk_status_icon_new_from_stock is not generated due to deprecation attr

/*
Creates a new #GtkStatusbar ready for messages.
*/
func StatusbarNew() (return__ *Statusbar) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_statusbar_new()
	if __cgo__return__ != nil {
		return__ = NewStatusbarFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_style_new is not generated due to deprecation attr

/*
Creates a standalone #GtkStyleContext, this style context
won’t be attached to any widget, so you may want
to call gtk_style_context_set_path() yourself.

This function is only useful when using the theming layer
separated from GTK+, if you are using #GtkStyleContext to
theme #GtkWidgets, use gtk_widget_get_style_context()
in order to get a style context ready to theme the widget.
*/
func StyleContextNew() (return__ *StyleContext) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_style_context_new()
	if __cgo__return__ != nil {
		return__ = NewStyleContextFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns a newly created #GtkStyleProperties
*/
func StylePropertiesNew() (return__ *StyleProperties) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_style_properties_new()
	if __cgo__return__ != nil {
		return__ = NewStylePropertiesFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkSwitch widget.
*/
func SwitchNew() (return__ *Switch) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_switch_new()
	if __cgo__return__ != nil {
		return__ = NewSwitchFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_table_new is not generated due to deprecation attr

// gtk_tearoff_menu_item_new is not generated due to deprecation attr

/*
Creates a new text buffer.
*/
func TextBufferNew(table IsTextTagTable) (return__ *TextBuffer) {
	var __cgo__table *C.GtkTextTagTable
	if table != nil {
		__cgo__table = table.GetTextTagTablePointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_text_buffer_new(__cgo__table)
	if __cgo__return__ != nil {
		return__ = NewTextBufferFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkTextChildAnchor. Usually you would then insert
it into a #GtkTextBuffer with gtk_text_buffer_insert_child_anchor().
To perform the creation and insertion in one step, use the
convenience function gtk_text_buffer_create_child_anchor().
*/
func TextChildAnchorNew() (return__ *TextChildAnchor) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_text_child_anchor_new()
	if __cgo__return__ != nil {
		return__ = NewTextChildAnchorFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a text mark. Add it to a buffer using gtk_text_buffer_add_mark().
If @name is %NULL, the mark is anonymous; otherwise, the mark can be
retrieved by name using gtk_text_buffer_get_mark(). If a mark has left
gravity, and text is inserted at the mark’s current location, the mark
will be moved to the left of the newly-inserted text. If the mark has
right gravity (@left_gravity = %FALSE), the mark will end up on the
right of newly-inserted text. The standard left-to-right cursor is a
mark with right gravity (when you type, the cursor stays on the right
side of the text you’re typing).
*/
func TextMarkNew(name string, left_gravity bool) (return__ *TextMark) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__left_gravity := C.gboolean(0)
	if left_gravity {
		__cgo__left_gravity = C.gboolean(1)
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_text_mark_new(__cgo__name, __cgo__left_gravity)
	C.free(unsafe.Pointer(__cgo__name))
	if __cgo__return__ != nil {
		return__ = NewTextMarkFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GtkTextTag. Configure the tag using object arguments,
i.e. using g_object_set().
*/
func TextTagNew(name string) (return__ *TextTag) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_text_tag_new(__cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	if __cgo__return__ != nil {
		return__ = NewTextTagFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkTextTagTable. The table contains no tags by
default.
*/
func TextTagTableNew() (return__ *TextTagTable) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_text_tag_table_new()
	if __cgo__return__ != nil {
		return__ = NewTextTagTableFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkTextView. If you don’t call gtk_text_view_set_buffer()
before using the text view, an empty default buffer will be created
for you. Get the buffer with gtk_text_view_get_buffer(). If you want
to specify your own buffer, consider gtk_text_view_new_with_buffer().
*/
func TextViewNew() (return__ *TextView) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_text_view_new()
	if __cgo__return__ != nil {
		return__ = NewTextViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkTextView widget displaying the buffer
@buffer. One buffer can be shared among many widgets.
@buffer may be %NULL to create a default buffer, in which case
this function is equivalent to gtk_text_view_new(). The
text view adds its own reference count to the buffer; it does not
take over an existing reference.
*/
func TextViewNewWithBuffer(buffer IsTextBuffer) (return__ *TextView) {
	var __cgo__buffer *C.GtkTextBuffer
	if buffer != nil {
		__cgo__buffer = buffer.GetTextBufferPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_text_view_new_with_buffer(__cgo__buffer)
	if __cgo__return__ != nil {
		return__ = NewTextViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_toggle_action_new is not generated due to deprecation attr

/*
Creates a new toggle button. A widget should be packed into the button, as in gtk_button_new().
*/
func ToggleButtonNew() (return__ *ToggleButton) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_toggle_button_new()
	if __cgo__return__ != nil {
		return__ = NewToggleButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new toggle button with a text label.
*/
func ToggleButtonNewWithLabel(label string) (return__ *ToggleButton) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_toggle_button_new_with_label(__cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewToggleButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkToggleButton containing a label. The label
will be created using gtk_label_new_with_mnemonic(), so underscores
in @label indicate the mnemonic for the button.
*/
func ToggleButtonNewWithMnemonic(label string) (return__ *ToggleButton) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_toggle_button_new_with_mnemonic(__cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewToggleButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns a new #GtkToggleToolButton
*/
func ToggleToolButtonNew() (return__ *ToggleToolButton) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_toggle_tool_button_new()
	if __cgo__return__ != nil {
		return__ = NewToggleToolButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_toggle_tool_button_new_from_stock is not generated due to deprecation attr

/*
Creates a new %GtkToolButton using @icon_widget as contents and @label as
label.
*/
func ToolButtonNew(icon_widget IsWidget, label string) (return__ *ToolButton) {
	var __cgo__icon_widget *C.GtkWidget
	if icon_widget != nil {
		__cgo__icon_widget = icon_widget.GetWidgetPointer()
	}
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_tool_button_new(__cgo__icon_widget, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewToolButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_tool_button_new_from_stock is not generated due to deprecation attr

/*
Creates a new #GtkToolItem
*/
func ToolItemNew() (return__ *ToolItem) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_tool_item_new()
	if __cgo__return__ != nil {
		return__ = NewToolItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new tool item group with label @label.
*/
func ToolItemGroupNew(label string) (return__ *ToolItemGroup) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_tool_item_group_new(__cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewToolItemGroupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new tool palette.
*/
func ToolPaletteNew() (return__ *ToolPalette) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_tool_palette_new()
	if __cgo__return__ != nil {
		return__ = NewToolPaletteFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new toolbar.
*/
func ToolbarNew() (return__ *Toolbar) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_toolbar_new()
	if __cgo__return__ != nil {
		return__ = NewToolbarFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_tree_store_new is not generated due to varargs

/*
Non vararg creation function.  Used primarily by language bindings.
*/
func TreeStoreNewv(n_columns int, types *C.GType) (return__ *TreeStore) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_tree_store_newv(C.gint(n_columns), types)
	if __cgo__return__ != nil {
		return__ = NewTreeStoreFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkTreeView widget.
*/
func TreeViewNew() (return__ *TreeView) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_tree_view_new()
	if __cgo__return__ != nil {
		return__ = NewTreeViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkTreeView widget with the model initialized to @model.
*/
func TreeViewNewWithModel(model *C.GtkTreeModel) (return__ *TreeView) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_tree_view_new_with_model(model)
	if __cgo__return__ != nil {
		return__ = NewTreeViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkTreeViewColumn.
*/
func TreeViewColumnNew() (return__ *TreeViewColumn) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_tree_view_column_new()
	if __cgo__return__ != nil {
		return__ = NewTreeViewColumnFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkTreeViewColumn using @area to render its cells.
*/
func TreeViewColumnNewWithArea(area IsCellArea) (return__ *TreeViewColumn) {
	var __cgo__area *C.GtkCellArea
	if area != nil {
		__cgo__area = area.GetCellAreaPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_tree_view_column_new_with_area(__cgo__area)
	if __cgo__return__ != nil {
		return__ = NewTreeViewColumnFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_tree_view_column_new_with_attributes is not generated due to varargs

// gtk_ui_manager_new is not generated due to deprecation attr

// gtk_vbox_new is not generated due to deprecation attr

// gtk_vbutton_box_new is not generated due to deprecation attr

// gtk_vpaned_new is not generated due to deprecation attr

// gtk_vscale_new is not generated due to deprecation attr

// gtk_vscale_new_with_range is not generated due to deprecation attr

// gtk_vscrollbar_new is not generated due to deprecation attr

// gtk_vseparator_new is not generated due to deprecation attr

/*
Creates a new #GtkViewport with the given adjustments, or with default
adjustments if none are given.
*/
func ViewportNew(hadjustment IsAdjustment, vadjustment IsAdjustment) (return__ *Viewport) {
	var __cgo__hadjustment *C.GtkAdjustment
	if hadjustment != nil {
		__cgo__hadjustment = hadjustment.GetAdjustmentPointer()
	}
	var __cgo__vadjustment *C.GtkAdjustment
	if vadjustment != nil {
		__cgo__vadjustment = vadjustment.GetAdjustmentPointer()
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_viewport_new(__cgo__hadjustment, __cgo__vadjustment)
	if __cgo__return__ != nil {
		return__ = NewViewportFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a #GtkVolumeButton, with a range between 0.0 and 1.0, with
a stepping of 0.02. Volume values can be obtained and modified using
the functions from #GtkScaleButton.
*/
func VolumeButtonNew() (return__ *VolumeButton) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_volume_button_new()
	if __cgo__return__ != nil {
		return__ = NewVolumeButtonFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gtk_widget_new is not generated due to varargs

/*
Creates a new #GtkWindow, which is a toplevel window that can
contain other widgets. Nearly always, the type of the window should
be #GTK_WINDOW_TOPLEVEL. If you’re implementing something like a
popup menu from scratch (which is a bad idea, just use #GtkMenu),
you might use #GTK_WINDOW_POPUP. #GTK_WINDOW_POPUP is not for
dialogs, though in some other toolkits dialogs are called “popups”.
In GTK+, #GTK_WINDOW_POPUP means a pop-up menu or pop-up tooltip.
On X11, popup windows are not controlled by the
[window manager][gtk-X11-arch].

If you simply want an undecorated window (no window borders), use
gtk_window_set_decorated(), don’t use #GTK_WINDOW_POPUP.

All top-level windows created by gtk_window_new() are stored in
an internal top-level window list.  This list can be obtained from
gtk_window_list_toplevels().  Due to Gtk+ keeping a reference to
the window internally, gtk_window_new() does not return a reference
to the caller.

To delete a #GtkWindow, call gtk_widget_destroy().
*/
func WindowNew(type_ C.GtkWindowType) (return__ *Window) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_window_new(type_)
	if __cgo__return__ != nil {
		return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GtkWindowGroup object. Grabs added with
gtk_grab_add() only affect windows within the same #GtkWindowGroup.
*/
func WindowGroupNew() (return__ *WindowGroup) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gtk_window_group_new()
	if __cgo__return__ != nil {
		return__ = NewWindowGroupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}
