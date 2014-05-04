package cairo

/*
#include <cairo.h>
#include <stdlib.h>
*/
import "C"

type Context C.cairo_t
type Surface C.cairo_surface_t
type Matrix C.cairo_matrix_t
type Pattern C.cairo_pattern_t
type Region C.cairo_region_t
type FontOptions C.cairo_font_options_t
type FontType C.cairo_font_type_t
type FontFace C.cairo_font_face_t
type ScaledFont C.cairo_scaled_font_t
type Path C.cairo_path_t
type RectangleInt C.cairo_rectangle_int_t
