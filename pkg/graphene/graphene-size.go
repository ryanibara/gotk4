// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

// GTypeSize returns the GType for the type Size.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeSize() coreglib.Type {
	gtype := coreglib.Type(C.graphene_size_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalSize)
	return gtype
}

// Size: size.
//
// An instance of this type is always passed by reference.
type Size struct {
	*size
}

// size is the struct that's finalized.
type size struct {
	native *C.graphene_size_t
}

func marshalSize(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Size{&size{(*C.graphene_size_t)(b)}}, nil
}

// NewSizeAlloc constructs a struct Size.
func NewSizeAlloc() *Size {
	var _cret *C.graphene_size_t // in

	_cret = C.graphene_size_alloc()

	var _size *Size // out

	_size = (*Size)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_size)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.graphene_size_free((*C.graphene_size_t)(intern.C))
		},
	)

	return _size
}

// Width: width.
func (s *Size) Width() float32 {
	valptr := s.native.width
	var v float32 // out
	v = float32(valptr)
	return v
}

// Height: height.
func (s *Size) Height() float32 {
	valptr := s.native.height
	var v float32 // out
	v = float32(valptr)
	return v
}

// Equal checks whether the two give #graphene_size_t are equal.
//
// The function takes the following parameters:
//
//    - b: #graphene_size_t.
//
// The function returns the following values:
//
//    - ok: true if the sizes are equal.
//
func (a *Size) Equal(b *Size) bool {
	var _arg0 *C.graphene_size_t // out
	var _arg1 *C.graphene_size_t // out
	var _cret C._Bool            // in

	_arg0 = (*C.graphene_size_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_size_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_size_equal(_arg0, _arg1)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Init initializes a #graphene_size_t using the given width and height.
//
// The function takes the following parameters:
//
//    - width: width.
//    - height: height.
//
// The function returns the following values:
//
//    - size: initialized #graphene_size_t.
//
func (s *Size) Init(width float32, height float32) *Size {
	var _arg0 *C.graphene_size_t // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out
	var _cret *C.graphene_size_t // in

	_arg0 = (*C.graphene_size_t)(gextras.StructNative(unsafe.Pointer(s)))
	_arg1 = C.float(width)
	_arg2 = C.float(height)

	_cret = C.graphene_size_init(_arg0, _arg1, _arg2)
	runtime.KeepAlive(s)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _size *Size // out

	_size = (*Size)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _size
}

// InitFromSize initializes a #graphene_size_t using the width and height of the
// given src.
//
// The function takes the following parameters:
//
//    - src: #graphene_size_t.
//
// The function returns the following values:
//
//    - size: initialized #graphene_size_t.
//
func (s *Size) InitFromSize(src *Size) *Size {
	var _arg0 *C.graphene_size_t // out
	var _arg1 *C.graphene_size_t // out
	var _cret *C.graphene_size_t // in

	_arg0 = (*C.graphene_size_t)(gextras.StructNative(unsafe.Pointer(s)))
	_arg1 = (*C.graphene_size_t)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.graphene_size_init_from_size(_arg0, _arg1)
	runtime.KeepAlive(s)
	runtime.KeepAlive(src)

	var _size *Size // out

	_size = (*Size)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _size
}

// Interpolate: linearly interpolates the two given #graphene_size_t using the
// given interpolation factor.
//
// The function takes the following parameters:
//
//    - b: #graphene_size_t.
//    - factor: linear interpolation factor.
//
// The function returns the following values:
//
//    - res: return location for the interpolated size.
//
func (a *Size) Interpolate(b *Size, factor float64) *Size {
	var _arg0 *C.graphene_size_t // out
	var _arg1 *C.graphene_size_t // out
	var _arg2 C.double           // out
	var _arg3 C.graphene_size_t  // in

	_arg0 = (*C.graphene_size_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_size_t)(gextras.StructNative(unsafe.Pointer(b)))
	_arg2 = C.double(factor)

	C.graphene_size_interpolate(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
	runtime.KeepAlive(factor)

	var _res *Size // out

	_res = (*Size)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _res
}

// Scale scales the components of a #graphene_size_t using the given factor.
//
// The function takes the following parameters:
//
//    - factor: scaling factor.
//
// The function returns the following values:
//
//    - res: return location for the scaled size.
//
func (s *Size) Scale(factor float32) *Size {
	var _arg0 *C.graphene_size_t // out
	var _arg1 C.float            // out
	var _arg2 C.graphene_size_t  // in

	_arg0 = (*C.graphene_size_t)(gextras.StructNative(unsafe.Pointer(s)))
	_arg1 = C.float(factor)

	C.graphene_size_scale(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(s)
	runtime.KeepAlive(factor)

	var _res *Size // out

	_res = (*Size)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// SizeZero: constant pointer to a zero #graphene_size_t, useful for equality
// checks and interpolations.
//
// The function returns the following values:
//
//    - size: constant size.
//
func SizeZero() *Size {
	var _cret *C.graphene_size_t // in

	_cret = C.graphene_size_zero()

	var _size *Size // out

	_size = (*Size)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _size
}
