// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// Log2VisGetEmbeddingLevels: return the bidirectional embedding levels of the
// input paragraph.
//
// The bidirectional embedding levels are defined by the Unicode Bidirectional
// Algorithm available at:
//
//    http://www.unicode.org/reports/tr9/
//
// If the input base direction is a weak direction, the direction of the
// characters in the text will determine the final resolved direction.
//
// The function takes the following parameters:
//
//   - text to itemize.
//   - length: number of bytes (not characters) to process, or -1 if text is
//     nul-terminated and the length should be calculated.
//   - pbaseDir: input base direction, and output resolved direction.
//
// The function returns the following values:
//
//   - guint8: newly allocated array of embedding levels, one item per character
//     (not byte), that should be freed using g_free().
//
func Log2VisGetEmbeddingLevels(text string, length int, pbaseDir *Direction) *byte {
	var _arg1 *C.gchar          // out
	var _arg2 C.int             // out
	var _arg3 *C.PangoDirection // out
	var _cret *C.guint8         // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(length)
	_arg3 = (*C.PangoDirection)(unsafe.Pointer(pbaseDir))

	_cret = C.pango_log2vis_get_embedding_levels(_arg1, _arg2, _arg3)
	runtime.KeepAlive(text)
	runtime.KeepAlive(length)
	runtime.KeepAlive(pbaseDir)

	var _guint8 *byte // out

	_guint8 = (*byte)(unsafe.Pointer(_cret))

	return _guint8
}
