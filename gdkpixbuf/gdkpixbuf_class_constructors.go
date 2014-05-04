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

func init() {
	_ = unsafe.Pointer(nil)
	_ = reflect.ValueOf(nil)
	_ = errors.New("")
}

/*
Creates a new #GdkPixbuf structure and allocates a buffer for it.  The
buffer has an optimal rowstride.  Note that the buffer is not cleared;
you will have to fill it completely yourself.
*/
func PixbufNew(colorspace C.GdkColorspace, has_alpha bool, bits_per_sample int, width int, height int) (return__ *Pixbuf) {
	__cgo__has_alpha := C.gboolean(0)
	if has_alpha {
		__cgo__has_alpha = C.gboolean(1)
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_new(colorspace, __cgo__has_alpha, C.int(bits_per_sample), C.int(width), C.int(height))
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #GdkPixbuf out of in-memory image data.  Currently only RGB
images with 8 bits per sample are supported.
*/
func PixbufNewFromData(data []byte, colorspace C.GdkColorspace, has_alpha bool, bits_per_sample int, width int, height int, rowstride int, destroy_fn C.GdkPixbufDestroyNotify, destroy_fn_data unsafe.Pointer) (return__ *Pixbuf) {
	__header__data := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	__cgo__has_alpha := C.gboolean(0)
	if has_alpha {
		__cgo__has_alpha = C.gboolean(1)
	}
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_new_from_data((*C.guchar)(unsafe.Pointer(__header__data.Data)), colorspace, __cgo__has_alpha, C.int(bits_per_sample), C.int(width), C.int(height), C.int(rowstride), destroy_fn, (C.gpointer)(destroy_fn_data))
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new pixbuf by loading an image from a file.  The file format is
detected automatically. If %NULL is returned, then @error will be set.
Possible errors are in the #GDK_PIXBUF_ERROR and #G_FILE_ERROR domains.
*/
func PixbufNewFromFile(filename string) (return__ *Pixbuf, __err__ error) {
	__cgo__filename := C.CString(filename)
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_new_from_file(__cgo__filename, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__filename))
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new pixbuf by loading an image from a file.  The file format is
detected automatically. If %NULL is returned, then @error will be set.
Possible errors are in the #GDK_PIXBUF_ERROR and #G_FILE_ERROR domains.
The image will be scaled to fit in the requested size, optionally preserving
the image's aspect ratio.

When preserving the aspect ratio, a @width of -1 will cause the image
to be scaled to the exact given height, and a @height of -1 will cause
the image to be scaled to the exact given width. When not preserving
aspect ratio, a @width or @height of -1 means to not scale the image
at all in that dimension. Negative values for @width and @height are
allowed since 2.8.
*/
func PixbufNewFromFileAtScale(filename string, width int, height int, preserve_aspect_ratio bool) (return__ *Pixbuf, __err__ error) {
	__cgo__filename := C.CString(filename)
	__cgo__preserve_aspect_ratio := C.gboolean(0)
	if preserve_aspect_ratio {
		__cgo__preserve_aspect_ratio = C.gboolean(1)
	}
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_new_from_file_at_scale(__cgo__filename, C.int(width), C.int(height), __cgo__preserve_aspect_ratio, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__filename))
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new pixbuf by loading an image from a file.
The file format is detected automatically. If %NULL is returned, then
@error will be set. Possible errors are in the #GDK_PIXBUF_ERROR and
#G_FILE_ERROR domains.

The image will be scaled to fit in the requested size, preserving
the image's aspect ratio. Note that the returned pixbuf may be smaller
than @width x @height, if the aspect ratio requires it. To load
and image at the requested size, regardless of aspect ratio, use
gdk_pixbuf_new_from_file_at_scale().
*/
func PixbufNewFromFileAtSize(filename string, width int, height int) (return__ *Pixbuf, __err__ error) {
	__cgo__filename := C.CString(filename)
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_new_from_file_at_size(__cgo__filename, C.int(width), C.int(height), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__filename))
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Create a #GdkPixbuf from a flat representation that is suitable for
storing as inline data in a program. This is useful if you want to
ship a program with images, but don't want to depend on any
external files.

gdk-pixbuf ships with a program called [gdk-pixbuf-csource][gdk-pixbuf-csource],
which allows for conversion of #GdkPixbufs into such a inline representation.
In almost all cases, you should pass the `--raw` option to
`gdk-pixbuf-csource`. A sample invocation would be:

|[
 gdk-pixbuf-csource --raw --name=myimage_inline myimage.png
]|

For the typical case where the inline pixbuf is read-only static data,
you don't need to copy the pixel data unless you intend to write to
it, so you can pass %FALSE for @copy_pixels.  (If you pass `--rle` to
`gdk-pixbuf-csource`, a copy will be made even if @copy_pixels is %FALSE,
so using this option is generally a bad idea.)

If you create a pixbuf from const inline data compiled into your
program, it's probably safe to ignore errors and disable length checks,
since things will always succeed:
|[
pixbuf = gdk_pixbuf_new_from_inline (-1, myimage_inline, FALSE, NULL);
]|

For non-const inline data, you could get out of memory. For untrusted
inline data located at runtime, you could have corrupt inline data in
addition.
*/
func PixbufNewFromInline(data_length int, data []byte, copy_pixels bool) (return__ *Pixbuf, __err__ error) {
	__header__data := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	__cgo__copy_pixels := C.gboolean(0)
	if copy_pixels {
		__cgo__copy_pixels = C.gboolean(1)
	}
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_new_from_inline(C.gint(data_length), (*C.guint8)(unsafe.Pointer(__header__data.Data)), __cgo__copy_pixels, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new pixbuf by loading an image from an resource.

The file format is detected automatically. If %NULL is returned, then
@error will be set.
*/
func PixbufNewFromResource(resource_path string) (return__ *Pixbuf, __err__ error) {
	__cgo__resource_path := C.CString(resource_path)
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_new_from_resource(__cgo__resource_path, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__resource_path))
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new pixbuf by loading an image from an resource.

The file format is detected automatically. If %NULL is returned, then
@error will be set.

The image will be scaled to fit in the requested size, optionally
preserving the image's aspect ratio. When preserving the aspect ratio,
a @width of -1 will cause the image to be scaled to the exact given
height, and a @height of -1 will cause the image to be scaled to the
exact given width. When not preserving aspect ratio, a @width or
@height of -1 means to not scale the image at all in that dimension.

The stream is not closed.
*/
func PixbufNewFromResourceAtScale(resource_path string, width int, height int, preserve_aspect_ratio bool) (return__ *Pixbuf, __err__ error) {
	__cgo__resource_path := C.CString(resource_path)
	__cgo__preserve_aspect_ratio := C.gboolean(0)
	if preserve_aspect_ratio {
		__cgo__preserve_aspect_ratio = C.gboolean(1)
	}
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_new_from_resource_at_scale(__cgo__resource_path, C.int(width), C.int(height), __cgo__preserve_aspect_ratio, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__resource_path))
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new pixbuf by loading an image from an input stream.

The file format is detected automatically. If %NULL is returned, then
@error will be set. The @cancellable can be used to abort the operation
from another thread. If the operation was cancelled, the error
%G_IO_ERROR_CANCELLED will be returned. Other possible errors are in
the #GDK_PIXBUF_ERROR and %G_IO_ERROR domains.

The stream is not closed.
*/
func PixbufNewFromStream(stream *C.GInputStream, cancellable *C.GCancellable) (return__ *Pixbuf, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_new_from_stream(stream, cancellable, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new pixbuf by loading an image from an input stream.

The file format is detected automatically. If %NULL is returned, then
@error will be set. The @cancellable can be used to abort the operation
from another thread. If the operation was cancelled, the error
%G_IO_ERROR_CANCELLED will be returned. Other possible errors are in
the #GDK_PIXBUF_ERROR and %G_IO_ERROR domains.

The image will be scaled to fit in the requested size, optionally
preserving the image's aspect ratio. When preserving the aspect ratio,
a @width of -1 will cause the image to be scaled to the exact given
height, and a @height of -1 will cause the image to be scaled to the
exact given width. When not preserving aspect ratio, a @width or
@height of -1 means to not scale the image at all in that dimension.

The stream is not closed.
*/
func PixbufNewFromStreamAtScale(stream *C.GInputStream, width int, height int, preserve_aspect_ratio bool, cancellable *C.GCancellable) (return__ *Pixbuf, __err__ error) {
	__cgo__preserve_aspect_ratio := C.gboolean(0)
	if preserve_aspect_ratio {
		__cgo__preserve_aspect_ratio = C.gboolean(1)
	}
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_new_from_stream_at_scale(stream, C.gint(width), C.gint(height), __cgo__preserve_aspect_ratio, cancellable, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Finishes an asynchronous pixbuf creation operation started with
gdk_pixbuf_new_from_stream_async().
*/
func PixbufNewFromStreamFinish(async_result *C.GAsyncResult) (return__ *Pixbuf, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_new_from_stream_finish(async_result, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new pixbuf by parsing XPM data in memory.  This data is commonly
the result of including an XPM file into a program's C source.
*/
func PixbufNewFromXpmData(data []string) (return__ *Pixbuf) {
	__header__data := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_new_from_xpm_data((**C.char)(unsafe.Pointer(__header__data.Data)))
	if __cgo__return__ != nil {
		return__ = NewPixbufFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new animation by loading it from a file.  The file format is
detected automatically.  If the file's format does not support multi-frame
images, then an animation with a single frame will be created. Possible errors
are in the #GDK_PIXBUF_ERROR and #G_FILE_ERROR domains.
*/
func PixbufAnimationNewFromFile(filename string) (return__ *PixbufAnimation, __err__ error) {
	__cgo__filename := C.CString(filename)
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_animation_new_from_file(__cgo__filename, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__filename))
	if __cgo__return__ != nil {
		return__ = NewPixbufAnimationFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new pixbuf animation by loading an image from an resource.

The file format is detected automatically. If %NULL is returned, then
@error will be set.
*/
func PixbufAnimationNewFromResource(resource_path string) (return__ *PixbufAnimation, __err__ error) {
	__cgo__resource_path := C.CString(resource_path)
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_animation_new_from_resource(__cgo__resource_path, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__resource_path))
	if __cgo__return__ != nil {
		return__ = NewPixbufAnimationFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new animation by loading it from an input stream.

The file format is detected automatically. If %NULL is returned, then
@error will be set. The @cancellable can be used to abort the operation
from another thread. If the operation was cancelled, the error
%G_IO_ERROR_CANCELLED will be returned. Other possible errors are in
the #GDK_PIXBUF_ERROR and %G_IO_ERROR domains.

The stream is not closed.
*/
func PixbufAnimationNewFromStream(stream *C.GInputStream, cancellable *C.GCancellable) (return__ *PixbufAnimation, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_animation_new_from_stream(stream, cancellable, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewPixbufAnimationFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Finishes an asynchronous pixbuf animation creation operation started with
gdk_pixbuf_animation_new_from_stream_async().
*/
func PixbufAnimationNewFromStreamFinish(async_result *C.GAsyncResult) (return__ *PixbufAnimation, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_animation_new_from_stream_finish(async_result, &__cgo_error__)
	if __cgo__return__ != nil {
		return__ = NewPixbufAnimationFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new pixbuf loader object.
*/
func PixbufLoaderNew() (return__ *PixbufLoader) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_loader_new()
	if __cgo__return__ != nil {
		return__ = NewPixbufLoaderFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new pixbuf loader object that always attempts to parse
image data as if it were an image of mime type @mime_type, instead of
identifying the type automatically. Useful if you want an error if
the image isn't the expected mime type, for loading image formats
that can't be reliably identified by looking at the data, or if
the user manually forces a specific mime type.

The list of supported mime types depends on what image loaders
are installed, but typically "image/png", "image/jpeg", "image/gif",
"image/tiff" and "image/x-xpixmap" are among the supported mime types.
To obtain the full list of supported mime types, call
gdk_pixbuf_format_get_mime_types() on each of the #GdkPixbufFormat
structs returned by gdk_pixbuf_get_formats().
*/
func PixbufLoaderNewWithMimeType(mime_type string) (return__ *PixbufLoader, __err__ error) {
	__cgo__mime_type := C.CString(mime_type)
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_loader_new_with_mime_type(__cgo__mime_type, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__mime_type))
	if __cgo__return__ != nil {
		return__ = NewPixbufLoaderFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new pixbuf loader object that always attempts to parse
image data as if it were an image of type @image_type, instead of
identifying the type automatically. Useful if you want an error if
the image isn't the expected type, for loading image formats
that can't be reliably identified by looking at the data, or if
the user manually forces a specific type.

The list of supported image formats depends on what image loaders
are installed, but typically "png", "jpeg", "gif", "tiff" and
"xpm" are among the supported formats. To obtain the full list of
supported image formats, call gdk_pixbuf_format_get_name() on each
of the #GdkPixbufFormat structs returned by gdk_pixbuf_get_formats().
*/
func PixbufLoaderNewWithType(image_type string) (return__ *PixbufLoader, __err__ error) {
	__cgo__image_type := C.CString(image_type)
	var __cgo_error__ *C.GError
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_loader_new_with_type(__cgo__image_type, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__image_type))
	if __cgo__return__ != nil {
		return__ = NewPixbufLoaderFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Creates a new, empty animation.
*/
func PixbufSimpleAnimNew(width int, height int, rate float32) (return__ *PixbufSimpleAnim) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.gdk_pixbuf_simple_anim_new(C.gint(width), C.gint(height), C.gfloat(rate))
	if __cgo__return__ != nil {
		return__ = NewPixbufSimpleAnimFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}
