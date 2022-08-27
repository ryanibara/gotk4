// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <pango/pango.h>
// gboolean _gotk4_pango1_FontFamily_virtual_is_monospace(void* fnptr, PangoFontFamily* arg0) {
//   return ((gboolean (*)(PangoFontFamily*))(fnptr))(arg0);
// };
// void _gotk4_pango1_FontFace_virtual_list_sizes(void* fnptr, PangoFontFace* arg0, int** arg1, int* arg2) {
//   ((void (*)(PangoFontFace*, int**, int*))(fnptr))(arg0, arg1, arg2);
// };
import "C"

// ListSizes: list the available sizes for a font.
//
// This is only applicable to bitmap fonts. For scalable fonts, stores NULL at
// the location pointed to by sizes and 0 at the location pointed to by n_sizes.
// The sizes returned are in Pango units and are sorted in ascending order.
//
// The function returns the following values:
//
//    - sizes (optional): location to store a pointer to an array of int. This
//      array should be freed with g_free().
//
func (face *FontFace) ListSizes() []int {
	var _arg0 *C.PangoFontFace // out
	var _arg1 *C.int           // in
	var _arg2 C.int            // in

	_arg0 = (*C.PangoFontFace)(unsafe.Pointer(coreglib.InternObject(face).Native()))

	C.pango_font_face_list_sizes(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(face)

	var _sizes []int // out

	if _arg1 != nil {
		defer C.free(unsafe.Pointer(_arg1))
		{
			src := unsafe.Slice((*C.int)(_arg1), _arg2)
			_sizes = make([]int, _arg2)
			for i := 0; i < int(_arg2); i++ {
				_sizes[i] = int(src[i])
			}
		}
	}

	return _sizes
}

// listSizes: list the available sizes for a font.
//
// This is only applicable to bitmap fonts. For scalable fonts, stores NULL at
// the location pointed to by sizes and 0 at the location pointed to by n_sizes.
// The sizes returned are in Pango units and are sorted in ascending order.
//
// The function returns the following values:
//
//    - sizes (optional): location to store a pointer to an array of int. This
//      array should be freed with g_free().
//
func (face *FontFace) listSizes() []int {
	gclass := (*C.PangoFontFaceClass)(coreglib.PeekParentClass(face))
	fnarg := gclass.list_sizes

	var _arg0 *C.PangoFontFace // out
	var _arg1 *C.int           // in
	var _arg2 C.int            // in

	_arg0 = (*C.PangoFontFace)(unsafe.Pointer(coreglib.InternObject(face).Native()))

	C._gotk4_pango1_FontFace_virtual_list_sizes(unsafe.Pointer(fnarg), _arg0, &_arg1, &_arg2)
	runtime.KeepAlive(face)

	var _sizes []int // out

	if _arg1 != nil {
		defer C.free(unsafe.Pointer(_arg1))
		{
			src := unsafe.Slice((*C.int)(_arg1), _arg2)
			_sizes = make([]int, _arg2)
			for i := 0; i < int(_arg2); i++ {
				_sizes[i] = int(src[i])
			}
		}
	}

	return _sizes
}

// IsMonospace: monospace font is a font designed for text display where the the
// characters form a regular grid.
//
// For Western languages this would mean that the advance width of all
// characters are the same, but this categorization also includes Asian fonts
// which include double-width characters: characters that occupy two grid cells.
// g_unichar_iswide() returns a result that indicates whether a character is
// typically double-width in a monospace font.
//
// The best way to find out the grid-cell size is to call
// pango.FontMetrics.GetApproximateDigitWidth(), since the results of
// pango.FontMetrics.GetApproximateCharWidth() may be affected by double-width
// characters.
//
// The function returns the following values:
//
//    - ok: TRUE if the family is monospace.
//
func (family *FontFamily) IsMonospace() bool {
	var _arg0 *C.PangoFontFamily // out
	var _cret C.gboolean         // in

	_arg0 = (*C.PangoFontFamily)(unsafe.Pointer(coreglib.InternObject(family).Native()))

	_cret = C.pango_font_family_is_monospace(_arg0)
	runtime.KeepAlive(family)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// isMonospace: monospace font is a font designed for text display where the the
// characters form a regular grid.
//
// For Western languages this would mean that the advance width of all
// characters are the same, but this categorization also includes Asian fonts
// which include double-width characters: characters that occupy two grid cells.
// g_unichar_iswide() returns a result that indicates whether a character is
// typically double-width in a monospace font.
//
// The best way to find out the grid-cell size is to call
// pango.FontMetrics.GetApproximateDigitWidth(), since the results of
// pango.FontMetrics.GetApproximateCharWidth() may be affected by double-width
// characters.
//
// The function returns the following values:
//
//    - ok: TRUE if the family is monospace.
//
func (family *FontFamily) isMonospace() bool {
	gclass := (*C.PangoFontFamilyClass)(coreglib.PeekParentClass(family))
	fnarg := gclass.is_monospace

	var _arg0 *C.PangoFontFamily // out
	var _cret C.gboolean         // in

	_arg0 = (*C.PangoFontFamily)(unsafe.Pointer(coreglib.InternObject(family).Native()))

	_cret = C._gotk4_pango1_FontFamily_virtual_is_monospace(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(family)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
