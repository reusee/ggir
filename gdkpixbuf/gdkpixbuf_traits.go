package gdkpixbuf

/*
#include <gdk-pixbuf/gdk-pixbuf.h>
#include <gdk-pixbuf/gdk-pixdata.h>
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

var _ = gobject.UnusedFix_

type TraitPixbuf struct{ CPointer *C.GdkPixbuf }
type IsPixbuf interface {
	GetPixbufPointer() *C.GdkPixbuf
}

func (self *TraitPixbuf) GetPixbufPointer() *C.GdkPixbuf {
	return self.CPointer
}
func NewTraitPixbuf(p unsafe.Pointer) *TraitPixbuf {
	return &TraitPixbuf{(*C.GdkPixbuf)(p)}
}

/*
Takes an existing pixbuf and adds an alpha channel to it.
If the existing pixbuf already had an alpha channel, the channel
values are copied from the original; otherwise, the alpha channel
is initialized to 255 (full opacity).

If @substitute_color is %TRUE, then the color specified by (@r, @g, @b) will be
assigned zero opacity. That is, if you pass (255, 255, 255) for the
substitute color, all white pixels will become fully transparent.
*/
func (self *TraitPixbuf) AddAlpha(substitute_color bool, r uint8, g uint8, b uint8) (return__ *Pixbuf) {
	__cgo__substitute_color := C.gboolean(0)
	if substitute_color {
		__cgo__substitute_color = C.gboolean(1)
	}
	var __cgo__return__ *C.GdkPixbuf
	__cgo__return__ = C.gdk_pixbuf_add_alpha(self.CPointer, __cgo__substitute_color, C.guchar(r), C.guchar(g), C.guchar(b))
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Takes an existing pixbuf and checks for the presence of an
associated "orientation" option, which may be provided by the
jpeg loader (which reads the exif orientation tag) or the
tiff loader (which reads the tiff orientation tag, and
compensates it for the partial transforms performed by
libtiff). If an orientation option/tag is present, the
appropriate transform will be performed so that the pixbuf
is oriented correctly.
*/
func (self *TraitPixbuf) ApplyEmbeddedOrientation() (return__ *Pixbuf) {
	var __cgo__return__ *C.GdkPixbuf
	__cgo__return__ = C.gdk_pixbuf_apply_embedded_orientation(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a transformation of the source image @src by scaling by
@scale_x and @scale_y then translating by @offset_x and @offset_y.
This gives an image in the coordinates of the destination pixbuf.
The rectangle (@dest_x, @dest_y, @dest_width, @dest_height)
is then composited onto the corresponding rectangle of the
original destination image.

When the destination rectangle contains parts not in the source
image, the data at the edges of the source image is replicated
to infinity.

![](composite.png)
*/
func (self *TraitPixbuf) Composite(dest IsPixbuf, dest_x int, dest_y int, dest_width int, dest_height int, offset_x float64, offset_y float64, scale_x float64, scale_y float64, interp_type C.GdkInterpType, overall_alpha int) {
	C.gdk_pixbuf_composite(self.CPointer, dest.GetPixbufPointer(), C.int(dest_x), C.int(dest_y), C.int(dest_width), C.int(dest_height), C.double(offset_x), C.double(offset_y), C.double(scale_x), C.double(scale_y), interp_type, C.int(overall_alpha))
	return
}

/*
Creates a transformation of the source image @src by scaling by
@scale_x and @scale_y then translating by @offset_x and @offset_y,
then composites the rectangle (@dest_x ,@dest_y, @dest_width,
@dest_height) of the resulting image with a checkboard of the
colors @color1 and @color2 and renders it onto the destination
image.

See gdk_pixbuf_composite_color_simple() for a simpler variant of this
function suitable for many tasks.
*/
func (self *TraitPixbuf) CompositeColor(dest IsPixbuf, dest_x int, dest_y int, dest_width int, dest_height int, offset_x float64, offset_y float64, scale_x float64, scale_y float64, interp_type C.GdkInterpType, overall_alpha int, check_x int, check_y int, check_size int, color1 uint32, color2 uint32) {
	C.gdk_pixbuf_composite_color(self.CPointer, dest.GetPixbufPointer(), C.int(dest_x), C.int(dest_y), C.int(dest_width), C.int(dest_height), C.double(offset_x), C.double(offset_y), C.double(scale_x), C.double(scale_y), interp_type, C.int(overall_alpha), C.int(check_x), C.int(check_y), C.int(check_size), C.guint32(color1), C.guint32(color2))
	return
}

/*
Creates a new #GdkPixbuf by scaling @src to @dest_width x
@dest_height and compositing the result with a checkboard of colors
@color1 and @color2.
*/
func (self *TraitPixbuf) CompositeColorSimple(dest_width int, dest_height int, interp_type C.GdkInterpType, overall_alpha int, check_size int, color1 uint32, color2 uint32) (return__ *Pixbuf) {
	var __cgo__return__ *C.GdkPixbuf
	__cgo__return__ = C.gdk_pixbuf_composite_color_simple(self.CPointer, C.int(dest_width), C.int(dest_height), interp_type, C.int(overall_alpha), C.int(check_size), C.guint32(color1), C.guint32(color2))
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GdkPixbuf with a copy of the information in the specified
@pixbuf.
*/
func (self *TraitPixbuf) Copy() (return__ *Pixbuf) {
	var __cgo__return__ *C.GdkPixbuf
	__cgo__return__ = C.gdk_pixbuf_copy(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Copies a rectangular area from @src_pixbuf to @dest_pixbuf.  Conversion of
pixbuf formats is done automatically.

If the source rectangle overlaps the destination rectangle on the
same pixbuf, it will be overwritten during the copy operation.
Therefore, you can not use this function to scroll a pixbuf.
*/
func (self *TraitPixbuf) CopyArea(src_x int, src_y int, width int, height int, dest_pixbuf IsPixbuf, dest_x int, dest_y int) {
	C.gdk_pixbuf_copy_area(self.CPointer, C.int(src_x), C.int(src_y), C.int(width), C.int(height), dest_pixbuf.GetPixbufPointer(), C.int(dest_x), C.int(dest_y))
	return
}

/*
Clears a pixbuf to the given RGBA value, converting the RGBA value into
the pixbuf's pixel format. The alpha will be ignored if the pixbuf
doesn't have an alpha channel.
*/
func (self *TraitPixbuf) Fill(pixel uint32) {
	C.gdk_pixbuf_fill(self.CPointer, C.guint32(pixel))
	return
}

/*
Flips a pixbuf horizontally or vertically and returns the
result in a new pixbuf.
*/
func (self *TraitPixbuf) Flip(horizontal bool) (return__ *Pixbuf) {
	__cgo__horizontal := C.gboolean(0)
	if horizontal {
		__cgo__horizontal = C.gboolean(1)
	}
	var __cgo__return__ *C.GdkPixbuf
	__cgo__return__ = C.gdk_pixbuf_flip(self.CPointer, __cgo__horizontal)
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Queries the number of bits per color sample in a pixbuf.
*/
func (self *TraitPixbuf) GetBitsPerSample() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.gdk_pixbuf_get_bits_per_sample(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the length of the pixel data, in bytes.
*/
func (self *TraitPixbuf) GetByteLength() (return__ int64) {
	var __cgo__return__ C.gsize
	__cgo__return__ = C.gdk_pixbuf_get_byte_length(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

/*
Queries the color space of a pixbuf.
*/
func (self *TraitPixbuf) GetColorspace() (return__ C.GdkColorspace) {
	return__ = C.gdk_pixbuf_get_colorspace(self.CPointer)
	return
}

/*
Queries whether a pixbuf has an alpha channel (opacity information).
*/
func (self *TraitPixbuf) GetHasAlpha() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_pixbuf_get_has_alpha(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Queries the height of a pixbuf.
*/
func (self *TraitPixbuf) GetHeight() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.gdk_pixbuf_get_height(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Queries the number of channels of a pixbuf.
*/
func (self *TraitPixbuf) GetNChannels() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.gdk_pixbuf_get_n_channels(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Looks up @key in the list of options that may have been attached to the
@pixbuf when it was loaded, or that may have been attached by another
function using gdk_pixbuf_set_option().

For instance, the ANI loader provides "Title" and "Artist" options.
The ICO, XBM, and XPM loaders provide "x_hot" and "y_hot" hot-spot
options for cursor definitions. The PNG loader provides the tEXt ancillary
chunk key/value pairs as options. Since 2.12, the TIFF and JPEG loaders
return an "orientation" option string that corresponds to the embedded
TIFF/Exif orientation tag (if present).
*/
func (self *TraitPixbuf) GetOption(key string) (return__ string) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gdk_pixbuf_get_option(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Queries a pointer to the pixel data of a pixbuf.
*/
func (self *TraitPixbuf) GetPixels() (return__ *C.guchar) {
	return__ = C.gdk_pixbuf_get_pixels(self.CPointer)
	return
}

/*
Queries a pointer to the pixel data of a pixbuf.
*/
func (self *TraitPixbuf) GetPixelsWithLength() (length uint, return__ []byte) {
	var __cgo__length C.guint
	var __cgo__return__ *C.guchar
	__cgo__return__ = C.gdk_pixbuf_get_pixels_with_length(self.CPointer, &__cgo__length)
	length = uint(__cgo__length)
	defer func() { return__ = C.GoBytes(unsafe.Pointer(__cgo__return__), C.int(length)) }()
	return
}

/*
Queries the rowstride of a pixbuf, which is the number of bytes between
the start of a row and the start of the next row.
*/
func (self *TraitPixbuf) GetRowstride() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.gdk_pixbuf_get_rowstride(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Queries the width of a pixbuf.
*/
func (self *TraitPixbuf) GetWidth() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.gdk_pixbuf_get_width(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Creates a new pixbuf which represents a sub-region of @src_pixbuf.
The new pixbuf shares its pixels with the original pixbuf, so
writing to one affects both.  The new pixbuf holds a reference to
@src_pixbuf, so @src_pixbuf will not be finalized until the new
pixbuf is finalized.

Note that if @src_pixbuf is read-only, this function will force it
to be mutable.
*/
func (self *TraitPixbuf) NewSubpixbuf(src_x int, src_y int, width int, height int) (return__ *Pixbuf) {
	var __cgo__return__ *C.GdkPixbuf
	__cgo__return__ = C.gdk_pixbuf_new_subpixbuf(self.CPointer, C.int(src_x), C.int(src_y), C.int(width), C.int(height))
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

func (self *TraitPixbuf) ReadPixelBytes() (return__ *C.GBytes) {
	return__ = C.gdk_pixbuf_read_pixel_bytes(self.CPointer)
	return
}

/*
Returns a read-only pointer to the raw pixel data; must not be
modified.  This function allows skipping the implicit copy that
must be made if gdk_pixbuf_get_pixels() is called on a read-only
pixbuf.
*/
func (self *TraitPixbuf) ReadPixels() (return__ *C.guint8) {
	return__ = C.gdk_pixbuf_read_pixels(self.CPointer)
	return
}

// gdk_pixbuf_ref is not generated due to deprecation attr

/*
Rotates a pixbuf by a multiple of 90 degrees, and returns the
result in a new pixbuf.
*/
func (self *TraitPixbuf) RotateSimple(angle C.GdkPixbufRotation) (return__ *Pixbuf) {
	var __cgo__return__ *C.GdkPixbuf
	__cgo__return__ = C.gdk_pixbuf_rotate_simple(self.CPointer, angle)
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Modifies saturation and optionally pixelates @src, placing the result in
@dest. @src and @dest may be the same pixbuf with no ill effects.  If
@saturation is 1.0 then saturation is not changed. If it's less than 1.0,
saturation is reduced (the image turns toward grayscale); if greater than
1.0, saturation is increased (the image gets more vivid colors). If @pixelate
is %TRUE, then pixels are faded in a checkerboard pattern to create a
pixelated image. @src and @dest must have the same image format, size, and
rowstride.
*/
func (self *TraitPixbuf) SaturateAndPixelate(dest IsPixbuf, saturation float32, pixelate bool) {
	__cgo__pixelate := C.gboolean(0)
	if pixelate {
		__cgo__pixelate = C.gboolean(1)
	}
	C.gdk_pixbuf_saturate_and_pixelate(self.CPointer, dest.GetPixbufPointer(), C.gfloat(saturation), __cgo__pixelate)
	return
}

// gdk_pixbuf_save is not generated due to varargs

// gdk_pixbuf_save_to_buffer is not generated due to varargs

/*
Saves pixbuf to a new buffer in format @type, which is currently "jpeg",
"tiff", "png", "ico" or "bmp".  See gdk_pixbuf_save_to_buffer()
for more details.
*/
func (self *TraitPixbuf) SaveToBufferv(type_ string, option_keys []string, option_values []string) (buffer []byte, buffer_size int64, return__ bool, __err__ error) {
	var __cgo__buffer *C.gchar
	var __cgo__buffer_size C.gsize
	__cgo__type_ := C.CString(type_)
	__header__option_keys := (*reflect.SliceHeader)(unsafe.Pointer(&option_keys))
	__header__option_values := (*reflect.SliceHeader)(unsafe.Pointer(&option_values))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_pixbuf_save_to_bufferv(self.CPointer, &__cgo__buffer, &__cgo__buffer_size, __cgo__type_, (**C.char)(unsafe.Pointer(__header__option_keys.Data)), (**C.char)(unsafe.Pointer(__header__option_values.Data)), &__cgo_error__)
	defer func() { buffer = C.GoBytes(unsafe.Pointer(__cgo__buffer), C.int(buffer_size)) }()
	buffer_size = int64(__cgo__buffer_size)
	C.free(unsafe.Pointer(__cgo__type_))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

// gdk_pixbuf_save_to_callback is not generated due to varargs

/*
Saves pixbuf to a callback in format @type, which is currently "jpeg",
"png", "tiff", "ico" or "bmp".  If @error is set, %FALSE will be returned. See
gdk_pixbuf_save_to_callback () for more details.
*/
func (self *TraitPixbuf) SaveToCallbackv(save_func C.GdkPixbufSaveFunc, user_data unsafe.Pointer, type_ string, option_keys []string, option_values []string) (return__ bool, __err__ error) {
	__cgo__type_ := C.CString(type_)
	__header__option_keys := (*reflect.SliceHeader)(unsafe.Pointer(&option_keys))
	__header__option_values := (*reflect.SliceHeader)(unsafe.Pointer(&option_values))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_pixbuf_save_to_callbackv(self.CPointer, save_func, (C.gpointer)(user_data), __cgo__type_, (**C.char)(unsafe.Pointer(__header__option_keys.Data)), (**C.char)(unsafe.Pointer(__header__option_values.Data)), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__type_))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

// gdk_pixbuf_save_to_stream is not generated due to varargs

// gdk_pixbuf_save_to_stream_async is not generated due to varargs

/*
Saves pixbuf to a file in @type, which is currently "jpeg", "png", "tiff", "ico" or "bmp".
If @error is set, %FALSE will be returned.
See gdk_pixbuf_save () for more details.
*/
func (self *TraitPixbuf) Savev(filename string, type_ string, option_keys []string, option_values []string) (return__ bool, __err__ error) {
	__cgo__filename := C.CString(filename)
	__cgo__type_ := C.CString(type_)
	__header__option_keys := (*reflect.SliceHeader)(unsafe.Pointer(&option_keys))
	__header__option_values := (*reflect.SliceHeader)(unsafe.Pointer(&option_values))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_pixbuf_savev(self.CPointer, __cgo__filename, __cgo__type_, (**C.char)(unsafe.Pointer(__header__option_keys.Data)), (**C.char)(unsafe.Pointer(__header__option_values.Data)), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__filename))
	C.free(unsafe.Pointer(__cgo__type_))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a transformation of the source image @src by scaling by
@scale_x and @scale_y then translating by @offset_x and @offset_y,
then renders the rectangle (@dest_x, @dest_y, @dest_width,
@dest_height) of the resulting image onto the destination image
replacing the previous contents.

Try to use gdk_pixbuf_scale_simple() first, this function is
the industrial-strength power tool you can fall back to if
gdk_pixbuf_scale_simple() isn't powerful enough.

If the source rectangle overlaps the destination rectangle on the
same pixbuf, it will be overwritten during the scaling which
results in rendering artifacts.
*/
func (self *TraitPixbuf) Scale(dest IsPixbuf, dest_x int, dest_y int, dest_width int, dest_height int, offset_x float64, offset_y float64, scale_x float64, scale_y float64, interp_type C.GdkInterpType) {
	C.gdk_pixbuf_scale(self.CPointer, dest.GetPixbufPointer(), C.int(dest_x), C.int(dest_y), C.int(dest_width), C.int(dest_height), C.double(offset_x), C.double(offset_y), C.double(scale_x), C.double(scale_y), interp_type)
	return
}

/*
Create a new #GdkPixbuf containing a copy of @src scaled to
@dest_width x @dest_height. Leaves @src unaffected.  @interp_type
should be #GDK_INTERP_NEAREST if you want maximum speed (but when
scaling down #GDK_INTERP_NEAREST is usually unusably ugly).  The
default @interp_type should be #GDK_INTERP_BILINEAR which offers
reasonable quality and speed.

You can scale a sub-portion of @src by creating a sub-pixbuf
pointing into @src; see gdk_pixbuf_new_subpixbuf().

For more complicated scaling/compositing see gdk_pixbuf_scale()
and gdk_pixbuf_composite().
*/
func (self *TraitPixbuf) ScaleSimple(dest_width int, dest_height int, interp_type C.GdkInterpType) (return__ *Pixbuf) {
	var __cgo__return__ *C.GdkPixbuf
	__cgo__return__ = C.gdk_pixbuf_scale_simple(self.CPointer, C.int(dest_width), C.int(dest_height), interp_type)
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// gdk_pixbuf_unref is not generated due to deprecation attr

type TraitPixbufAnimation struct{ CPointer *C.GdkPixbufAnimation }
type IsPixbufAnimation interface {
	GetPixbufAnimationPointer() *C.GdkPixbufAnimation
}

func (self *TraitPixbufAnimation) GetPixbufAnimationPointer() *C.GdkPixbufAnimation {
	return self.CPointer
}
func NewTraitPixbufAnimation(p unsafe.Pointer) *TraitPixbufAnimation {
	return &TraitPixbufAnimation{(*C.GdkPixbufAnimation)(p)}
}

/*
Queries the height of the bounding box of a pixbuf animation.
*/
func (self *TraitPixbufAnimation) GetHeight() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.gdk_pixbuf_animation_get_height(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Get an iterator for displaying an animation. The iterator provides
the frames that should be displayed at a given time.
It should be freed after use with g_object_unref().

@start_time would normally come from g_get_current_time(), and
marks the beginning of animation playback. After creating an
iterator, you should immediately display the pixbuf returned by
gdk_pixbuf_animation_iter_get_pixbuf(). Then, you should install a
timeout (with g_timeout_add()) or by some other mechanism ensure
that you'll update the image after
gdk_pixbuf_animation_iter_get_delay_time() milliseconds. Each time
the image is updated, you should reinstall the timeout with the new,
possibly-changed delay time.

As a shortcut, if @start_time is %NULL, the result of
g_get_current_time() will be used automatically.

To update the image (i.e. possibly change the result of
gdk_pixbuf_animation_iter_get_pixbuf() to a new frame of the animation),
call gdk_pixbuf_animation_iter_advance().

If you're using #GdkPixbufLoader, in addition to updating the image
after the delay time, you should also update it whenever you
receive the area_updated signal and
gdk_pixbuf_animation_iter_on_currently_loading_frame() returns
%TRUE. In this case, the frame currently being fed into the loader
has received new data, so needs to be refreshed. The delay time for
a frame may also be modified after an area_updated signal, for
example if the delay time for a frame is encoded in the data after
the frame itself. So your timeout should be reinstalled after any
area_updated signal.

A delay time of -1 is possible, indicating "infinite."
*/
func (self *TraitPixbufAnimation) GetIter(start_time *C.GTimeVal) (return__ *PixbufAnimationIter) {
	var __cgo__return__ *C.GdkPixbufAnimationIter
	__cgo__return__ = C.gdk_pixbuf_animation_get_iter(self.CPointer, start_time)
	if __cgo__return__ != nil {
		return__ = NewPixbufAnimationIterFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
If an animation is really just a plain image (has only one frame),
this function returns that image. If the animation is an animation,
this function returns a reasonable thing to display as a static
unanimated image, which might be the first frame, or something more
sophisticated. If an animation hasn't loaded any frames yet, this
function will return %NULL.
*/
func (self *TraitPixbufAnimation) GetStaticImage() (return__ *Pixbuf) {
	var __cgo__return__ *C.GdkPixbuf
	__cgo__return__ = C.gdk_pixbuf_animation_get_static_image(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Queries the width of the bounding box of a pixbuf animation.
*/
func (self *TraitPixbufAnimation) GetWidth() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.gdk_pixbuf_animation_get_width(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
If you load a file with gdk_pixbuf_animation_new_from_file() and it turns
out to be a plain, unanimated image, then this function will return
%TRUE. Use gdk_pixbuf_animation_get_static_image() to retrieve
the image.
*/
func (self *TraitPixbufAnimation) IsStaticImage() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_pixbuf_animation_is_static_image(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gdk_pixbuf_animation_ref is not generated due to deprecation attr

// gdk_pixbuf_animation_unref is not generated due to deprecation attr

type TraitPixbufAnimationIter struct{ CPointer *C.GdkPixbufAnimationIter }
type IsPixbufAnimationIter interface {
	GetPixbufAnimationIterPointer() *C.GdkPixbufAnimationIter
}

func (self *TraitPixbufAnimationIter) GetPixbufAnimationIterPointer() *C.GdkPixbufAnimationIter {
	return self.CPointer
}
func NewTraitPixbufAnimationIter(p unsafe.Pointer) *TraitPixbufAnimationIter {
	return &TraitPixbufAnimationIter{(*C.GdkPixbufAnimationIter)(p)}
}

/*
Possibly advances an animation to a new frame. Chooses the frame based
on the start time passed to gdk_pixbuf_animation_get_iter().

@current_time would normally come from g_get_current_time(), and
must be greater than or equal to the time passed to
gdk_pixbuf_animation_get_iter(), and must increase or remain
unchanged each time gdk_pixbuf_animation_iter_get_pixbuf() is
called. That is, you can't go backward in time; animations only
play forward.

As a shortcut, pass %NULL for the current time and g_get_current_time()
will be invoked on your behalf. So you only need to explicitly pass
@current_time if you're doing something odd like playing the animation
at double speed.

If this function returns %FALSE, there's no need to update the animation
display, assuming the display had been rendered prior to advancing;
if %TRUE, you need to call gdk_pixbuf_animation_iter_get_pixbuf()
and update the display with the new pixbuf.
*/
func (self *TraitPixbufAnimationIter) Advance(current_time *C.GTimeVal) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_pixbuf_animation_iter_advance(self.CPointer, current_time)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the number of milliseconds the current pixbuf should be displayed,
or -1 if the current pixbuf should be displayed forever. g_timeout_add()
conveniently takes a timeout in milliseconds, so you can use a timeout
to schedule the next update.
*/
func (self *TraitPixbufAnimationIter) GetDelayTime() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.gdk_pixbuf_animation_iter_get_delay_time(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the current pixbuf which should be displayed; the pixbuf might not
be the same size as the animation itself
(gdk_pixbuf_animation_get_width(), gdk_pixbuf_animation_get_height()).
This pixbuf should be displayed for
gdk_pixbuf_animation_iter_get_delay_time() milliseconds.  The caller
of this function does not own a reference to the returned pixbuf;
the returned pixbuf will become invalid when the iterator advances
to the next frame, which may happen anytime you call
gdk_pixbuf_animation_iter_advance(). Copy the pixbuf to keep it
(don't just add a reference), as it may get recycled as you advance
the iterator.
*/
func (self *TraitPixbufAnimationIter) GetPixbuf() (return__ *Pixbuf) {
	var __cgo__return__ *C.GdkPixbuf
	__cgo__return__ = C.gdk_pixbuf_animation_iter_get_pixbuf(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Used to determine how to respond to the area_updated signal on
#GdkPixbufLoader when loading an animation. area_updated is emitted
for an area of the frame currently streaming in to the loader. So if
you're on the currently loading frame, you need to redraw the screen for
the updated area.
*/
func (self *TraitPixbufAnimationIter) OnCurrentlyLoadingFrame() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_pixbuf_animation_iter_on_currently_loading_frame(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

type TraitPixbufLoader struct{ CPointer *C.GdkPixbufLoader }
type IsPixbufLoader interface {
	GetPixbufLoaderPointer() *C.GdkPixbufLoader
}

func (self *TraitPixbufLoader) GetPixbufLoaderPointer() *C.GdkPixbufLoader {
	return self.CPointer
}
func NewTraitPixbufLoader(p unsafe.Pointer) *TraitPixbufLoader {
	return &TraitPixbufLoader{(*C.GdkPixbufLoader)(p)}
}

/*
Informs a pixbuf loader that no further writes with
gdk_pixbuf_loader_write() will occur, so that it can free its
internal loading structures. Also, tries to parse any data that
hasn't yet been parsed; if the remaining data is partial or
corrupt, an error will be returned.  If %FALSE is returned, @error
will be set to an error from the #GDK_PIXBUF_ERROR or #G_FILE_ERROR
domains. If you're just cancelling a load rather than expecting it
to be finished, passing %NULL for @error to ignore it is
reasonable.

Remember that this does not unref the loader, so if you plan not to
use it anymore, please g_object_unref() it.
*/
func (self *TraitPixbufLoader) Close() (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_pixbuf_loader_close(self.CPointer, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Queries the #GdkPixbufAnimation that a pixbuf loader is currently creating.
In general it only makes sense to call this function after the "area-prepared"
signal has been emitted by the loader. If the loader doesn't have enough
bytes yet (hasn't emitted the "area-prepared" signal) this function will
return %NULL.
*/
func (self *TraitPixbufLoader) GetAnimation() (return__ *PixbufAnimation) {
	var __cgo__return__ *C.GdkPixbufAnimation
	__cgo__return__ = C.gdk_pixbuf_loader_get_animation(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewPixbufAnimationFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Obtains the available information about the format of the
currently loading image file.
*/
func (self *TraitPixbufLoader) GetFormat() (return__ *C.GdkPixbufFormat) {
	return__ = C.gdk_pixbuf_loader_get_format(self.CPointer)
	return
}

/*
Queries the #GdkPixbuf that a pixbuf loader is currently creating.
In general it only makes sense to call this function after the
"area-prepared" signal has been emitted by the loader; this means
that enough data has been read to know the size of the image that
will be allocated.  If the loader has not received enough data via
gdk_pixbuf_loader_write(), then this function returns %NULL.  The
returned pixbuf will be the same in all future calls to the loader,
so simply calling g_object_ref() should be sufficient to continue
using it.  Additionally, if the loader is an animation, it will
return the "static image" of the animation
(see gdk_pixbuf_animation_get_static_image()).
*/
func (self *TraitPixbufLoader) GetPixbuf() (return__ *Pixbuf) {
	var __cgo__return__ *C.GdkPixbuf
	__cgo__return__ = C.gdk_pixbuf_loader_get_pixbuf(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Causes the image to be scaled while it is loaded. The desired
image size can be determined relative to the original size of
the image by calling gdk_pixbuf_loader_set_size() from a
signal handler for the ::size-prepared signal.

Attempts to set the desired image size  are ignored after the
emission of the ::size-prepared signal.
*/
func (self *TraitPixbufLoader) SetSize(width int, height int) {
	C.gdk_pixbuf_loader_set_size(self.CPointer, C.int(width), C.int(height))
	return
}

/*
This will cause a pixbuf loader to parse the next @count bytes of
an image.  It will return %TRUE if the data was loaded successfully,
and %FALSE if an error occurred.  In the latter case, the loader
will be closed, and will not accept further writes. If %FALSE is
returned, @error will be set to an error from the #GDK_PIXBUF_ERROR
or #G_FILE_ERROR domains.
*/
func (self *TraitPixbufLoader) Write(buf []byte, count int64) (return__ bool, __err__ error) {
	__header__buf := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_pixbuf_loader_write(self.CPointer, (*C.guchar)(unsafe.Pointer(__header__buf.Data)), C.gsize(count), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This will cause a pixbuf loader to parse a buffer inside a #GBytes
for an image.  It will return %TRUE if the data was loaded successfully,
and %FALSE if an error occurred.  In the latter case, the loader
will be closed, and will not accept further writes. If %FALSE is
returned, @error will be set to an error from the #GDK_PIXBUF_ERROR
or #G_FILE_ERROR domains.

See also: gdk_pixbuf_loader_write()
*/
func (self *TraitPixbufLoader) WriteBytes(buffer *C.GBytes) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_pixbuf_loader_write_bytes(self.CPointer, buffer, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

type TraitPixbufSimpleAnim struct{ CPointer *C.GdkPixbufSimpleAnim }
type IsPixbufSimpleAnim interface {
	GetPixbufSimpleAnimPointer() *C.GdkPixbufSimpleAnim
}

func (self *TraitPixbufSimpleAnim) GetPixbufSimpleAnimPointer() *C.GdkPixbufSimpleAnim {
	return self.CPointer
}
func NewTraitPixbufSimpleAnim(p unsafe.Pointer) *TraitPixbufSimpleAnim {
	return &TraitPixbufSimpleAnim{(*C.GdkPixbufSimpleAnim)(p)}
}

/*
Adds a new frame to @animation. The @pixbuf must
have the dimensions specified when the animation
was constructed.
*/
func (self *TraitPixbufSimpleAnim) AddFrame(pixbuf IsPixbuf) {
	C.gdk_pixbuf_simple_anim_add_frame(self.CPointer, pixbuf.GetPixbufPointer())
	return
}

/*
Gets whether @animation should loop indefinitely when it reaches the end.
*/
func (self *TraitPixbufSimpleAnim) GetLoop() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gdk_pixbuf_simple_anim_get_loop(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets whether @animation should loop indefinitely when it reaches the end.
*/
func (self *TraitPixbufSimpleAnim) SetLoop(loop bool) {
	__cgo__loop := C.gboolean(0)
	if loop {
		__cgo__loop = C.gboolean(1)
	}
	C.gdk_pixbuf_simple_anim_set_loop(self.CPointer, __cgo__loop)
	return
}
