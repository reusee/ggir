package cairo

/*
#include <cairo.h>
#include <stdlib.h>
*/
import "C"

var (
	// Content
	CONTENT_COLOR       = C.cairo_content_t(C.CAIRO_CONTENT_COLOR)
	CONTENT_ALPHA       = C.cairo_content_t(C.CAIRO_CONTENT_ALPHA)
	CONTENT_COLOR_ALPHA = C.cairo_content_t(C.CAIRO_CONTENT_COLOR_ALPHA)
)
