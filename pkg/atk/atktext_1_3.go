// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// AtkTextRange** _gotk4_atk1_Text_virtual_get_bounded_ranges(void* fnptr, AtkText* arg0, AtkTextRectangle* arg1, AtkCoordType arg2, AtkTextClipType arg3, AtkTextClipType arg4) {
//   return ((AtkTextRange** (*)(AtkText*, AtkTextRectangle*, AtkCoordType, AtkTextClipType, AtkTextClipType))(fnptr))(arg0, arg1, arg2, arg3, arg4);
// };
// void _gotk4_atk1_Text_virtual_get_range_extents(void* fnptr, AtkText* arg0, gint arg1, gint arg2, AtkCoordType arg3, AtkTextRectangle* arg4) {
//   ((void (*)(AtkText*, gint, gint, AtkCoordType, AtkTextRectangle*))(fnptr))(arg0, arg1, arg2, arg3, arg4);
// };
import "C"

// BoundedRanges: get the ranges of text in the specified bounding box.
//
// The function takes the following parameters:
//
//    - rect: atkTextRectangle giving the dimensions of the bounding box.
//    - coordType: specify whether coordinates are relative to the screen or
//      widget window.
//    - xClipType: specify the horizontal clip type.
//    - yClipType: specify the vertical clip type.
//
// The function returns the following values:
//
//    - textRanges: array of AtkTextRange. The last element of the array returned
//      by this function will be NULL.
//
func (text *Text) BoundedRanges(rect *TextRectangle, coordType CoordType, xClipType, yClipType TextClipType) []*TextRange {
	var _arg0 *C.AtkText          // out
	var _arg1 *C.AtkTextRectangle // out
	var _arg2 C.AtkCoordType      // out
	var _arg3 C.AtkTextClipType   // out
	var _arg4 C.AtkTextClipType   // out
	var _cret **C.AtkTextRange    // in

	_arg0 = (*C.AtkText)(unsafe.Pointer(coreglib.InternObject(text).Native()))
	_arg1 = (*C.AtkTextRectangle)(gextras.StructNative(unsafe.Pointer(rect)))
	_arg2 = C.AtkCoordType(coordType)
	_arg3 = C.AtkTextClipType(xClipType)
	_arg4 = C.AtkTextClipType(yClipType)

	_cret = C.atk_text_get_bounded_ranges(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(text)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(coordType)
	runtime.KeepAlive(xClipType)
	runtime.KeepAlive(yClipType)

	var _textRanges []*TextRange // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.AtkTextRange
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_textRanges = make([]*TextRange, i)
		for i := range src {
			_textRanges[i] = (*TextRange)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(_textRanges[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}

	return _textRanges
}

// RangeExtents: get the bounding box for text within the specified range.
//
// If the extents can not be obtained (e.g. or missing support), the rectangle
// fields are set to -1.
//
// The function takes the following parameters:
//
//    - startOffset: offset of the first text character for which boundary
//      information is required.
//    - endOffset: offset of the text character after the last character for
//      which boundary information is required.
//    - coordType: specify whether coordinates are relative to the screen or
//      widget window.
//
// The function returns the following values:
//
//    - rect: pointer to a AtkTextRectangle which is filled in by this function.
//
func (text *Text) RangeExtents(startOffset, endOffset int, coordType CoordType) *TextRectangle {
	var _arg0 *C.AtkText         // out
	var _arg1 C.gint             // out
	var _arg2 C.gint             // out
	var _arg3 C.AtkCoordType     // out
	var _arg4 C.AtkTextRectangle // in

	_arg0 = (*C.AtkText)(unsafe.Pointer(coreglib.InternObject(text).Native()))
	_arg1 = C.gint(startOffset)
	_arg2 = C.gint(endOffset)
	_arg3 = C.AtkCoordType(coordType)

	C.atk_text_get_range_extents(_arg0, _arg1, _arg2, _arg3, &_arg4)
	runtime.KeepAlive(text)
	runtime.KeepAlive(startOffset)
	runtime.KeepAlive(endOffset)
	runtime.KeepAlive(coordType)

	var _rect *TextRectangle // out

	_rect = (*TextRectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg4))))

	return _rect
}

// boundedRanges: get the ranges of text in the specified bounding box.
//
// The function takes the following parameters:
//
//    - rect: atkTextRectangle giving the dimensions of the bounding box.
//    - coordType: specify whether coordinates are relative to the screen or
//      widget window.
//    - xClipType: specify the horizontal clip type.
//    - yClipType: specify the vertical clip type.
//
// The function returns the following values:
//
//    - textRanges: array of AtkTextRange. The last element of the array returned
//      by this function will be NULL.
//
func (text *Text) boundedRanges(rect *TextRectangle, coordType CoordType, xClipType, yClipType TextClipType) []*TextRange {
	gclass := (*C.AtkTextIface)(coreglib.PeekParentClass(text))
	fnarg := gclass.get_bounded_ranges

	var _arg0 *C.AtkText          // out
	var _arg1 *C.AtkTextRectangle // out
	var _arg2 C.AtkCoordType      // out
	var _arg3 C.AtkTextClipType   // out
	var _arg4 C.AtkTextClipType   // out
	var _cret **C.AtkTextRange    // in

	_arg0 = (*C.AtkText)(unsafe.Pointer(coreglib.InternObject(text).Native()))
	_arg1 = (*C.AtkTextRectangle)(gextras.StructNative(unsafe.Pointer(rect)))
	_arg2 = C.AtkCoordType(coordType)
	_arg3 = C.AtkTextClipType(xClipType)
	_arg4 = C.AtkTextClipType(yClipType)

	_cret = C._gotk4_atk1_Text_virtual_get_bounded_ranges(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(text)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(coordType)
	runtime.KeepAlive(xClipType)
	runtime.KeepAlive(yClipType)

	var _textRanges []*TextRange // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.AtkTextRange
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_textRanges = make([]*TextRange, i)
		for i := range src {
			_textRanges[i] = (*TextRange)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(_textRanges[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}

	return _textRanges
}

// rangeExtents: get the bounding box for text within the specified range.
//
// If the extents can not be obtained (e.g. or missing support), the rectangle
// fields are set to -1.
//
// The function takes the following parameters:
//
//    - startOffset: offset of the first text character for which boundary
//      information is required.
//    - endOffset: offset of the text character after the last character for
//      which boundary information is required.
//    - coordType: specify whether coordinates are relative to the screen or
//      widget window.
//
// The function returns the following values:
//
//    - rect: pointer to a AtkTextRectangle which is filled in by this function.
//
func (text *Text) rangeExtents(startOffset, endOffset int, coordType CoordType) *TextRectangle {
	gclass := (*C.AtkTextIface)(coreglib.PeekParentClass(text))
	fnarg := gclass.get_range_extents

	var _arg0 *C.AtkText         // out
	var _arg1 C.gint             // out
	var _arg2 C.gint             // out
	var _arg3 C.AtkCoordType     // out
	var _arg4 C.AtkTextRectangle // in

	_arg0 = (*C.AtkText)(unsafe.Pointer(coreglib.InternObject(text).Native()))
	_arg1 = C.gint(startOffset)
	_arg2 = C.gint(endOffset)
	_arg3 = C.AtkCoordType(coordType)

	C._gotk4_atk1_Text_virtual_get_range_extents(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, &_arg4)
	runtime.KeepAlive(text)
	runtime.KeepAlive(startOffset)
	runtime.KeepAlive(endOffset)
	runtime.KeepAlive(coordType)

	var _rect *TextRectangle // out

	_rect = (*TextRectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg4))))

	return _rect
}
