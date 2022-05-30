// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	"github.com/diamondburned/gotk4/pkg/graphene"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// RoundedRect: rectangular region with rounded corners.
//
// Application code should normalize rectangles using
// gsk.RoundedRect.Normalize(); this function will ensure that the bounds of the
// rectangle are normalized and ensure that the corner values are positive and
// the corners do not overlap.
//
// All functions taking a GskRoundedRect as an argument will internally operate
// on a normalized copy; all functions returning a GskRoundedRect will always
// return a normalized one.
//
// The algorithm used for normalizing corner sizes is described in the CSS
// specification (https://drafts.csswg.org/css-backgrounds-3/#border-radius).
//
// An instance of this type is always passed by reference.
type RoundedRect struct {
	*roundedRect
}

// roundedRect is the struct that's finalized.
type roundedRect struct {
	native *C.GskRoundedRect
}

// ContainsPoint checks if the given point is inside the rounded rectangle.
//
// The function takes the following parameters:
//
//    - point to check.
//
// The function returns the following values:
//
//    - ok: TRUE if the point is inside the rounded rectangle.
//
func (self *RoundedRect) ContainsPoint(point *graphene.Point) bool {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(point)))
	*(**RoundedRect)(unsafe.Pointer(&args[1])) = _arg1

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(point)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContainsRect checks if the given rect is contained inside the rounded
// rectangle.
//
// The function takes the following parameters:
//
//    - rect: rectangle to check.
//
// The function returns the following values:
//
//    - ok: TRUE if the rect is fully contained inside the rounded rectangle.
//
func (self *RoundedRect) ContainsRect(rect *graphene.Rect) bool {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(rect)))
	*(**RoundedRect)(unsafe.Pointer(&args[1])) = _arg1

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(rect)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Init initializes the given GskRoundedRect with the given values.
//
// This function will implicitly normalize the GskRoundedRect before returning.
//
// The function takes the following parameters:
//
//    - bounds: graphene_rect_t describing the bounds.
//    - topLeft: rounding radius of the top left corner.
//    - topRight: rounding radius of the top right corner.
//    - bottomRight: rounding radius of the bottom right corner.
//    - bottomLeft: rounding radius of the bottom left corner.
//
// The function returns the following values:
//
//    - roundedRect: initialized rectangle.
//
func (self *RoundedRect) Init(bounds *graphene.Rect, topLeft *graphene.Size, topRight *graphene.Size, bottomRight *graphene.Size, bottomLeft *graphene.Size) *RoundedRect {
	var args [6]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out
	var _arg3 *C.void // out
	var _arg4 *C.void // out
	var _arg5 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(bounds)))
	_arg2 = (*C.void)(gextras.StructNative(unsafe.Pointer(topLeft)))
	_arg3 = (*C.void)(gextras.StructNative(unsafe.Pointer(topRight)))
	_arg4 = (*C.void)(gextras.StructNative(unsafe.Pointer(bottomRight)))
	_arg5 = (*C.void)(gextras.StructNative(unsafe.Pointer(bottomLeft)))
	*(**RoundedRect)(unsafe.Pointer(&args[1])) = _arg1
	*(**graphene.Rect)(unsafe.Pointer(&args[2])) = _arg2
	*(**graphene.Size)(unsafe.Pointer(&args[3])) = _arg3
	*(**graphene.Size)(unsafe.Pointer(&args[4])) = _arg4
	*(**graphene.Size)(unsafe.Pointer(&args[5])) = _arg5

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(bounds)
	runtime.KeepAlive(topLeft)
	runtime.KeepAlive(topRight)
	runtime.KeepAlive(bottomRight)
	runtime.KeepAlive(bottomLeft)

	var _roundedRect *RoundedRect // out

	_roundedRect = (*RoundedRect)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _roundedRect
}

// InitCopy initializes self using the given src rectangle.
//
// This function will not normalize the GskRoundedRect, so make sure the source
// is normalized.
//
// The function takes the following parameters:
//
//    - src: GskRoundedRect.
//
// The function returns the following values:
//
//    - roundedRect: initialized rectangle.
//
func (self *RoundedRect) InitCopy(src *RoundedRect) *RoundedRect {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(src)))
	*(**RoundedRect)(unsafe.Pointer(&args[1])) = _arg1

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(src)

	var _roundedRect *RoundedRect // out

	_roundedRect = (*RoundedRect)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _roundedRect
}

// IntersectsRect checks if part of the given rect is contained inside the
// rounded rectangle.
//
// The function takes the following parameters:
//
//    - rect: rectangle to check.
//
// The function returns the following values:
//
//    - ok: TRUE if the rect intersects with the rounded rectangle.
//
func (self *RoundedRect) IntersectsRect(rect *graphene.Rect) bool {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(rect)))
	*(**RoundedRect)(unsafe.Pointer(&args[1])) = _arg1

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(rect)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsRectilinear checks if all corners of self are right angles and the
// rectangle covers all of its bounds.
//
// This information can be used to decide if gsk.ClipNode.New or
// gsk.RoundedClipNode.New should be called.
//
// The function returns the following values:
//
//    - ok: TRUE if the rectangle is rectilinear.
//
func (self *RoundedRect) IsRectilinear() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(**RoundedRect)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Normalize normalizes the passed rectangle.
//
// This function will ensure that the bounds of the rectangle are normalized and
// ensure that the corner values are positive and the corners do not overlap.
//
// The function returns the following values:
//
//    - roundedRect: normalized rectangle.
//
func (self *RoundedRect) Normalize() *RoundedRect {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(self)))
	*(**RoundedRect)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _roundedRect *RoundedRect // out

	_roundedRect = (*RoundedRect)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _roundedRect
}
