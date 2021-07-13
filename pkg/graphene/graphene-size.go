// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_size_get_type()), F: marshalSize},
	})
}

// Size: size.
type Size struct {
	native C.graphene_size_t
}

func marshalSize(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Size)(unsafe.Pointer(b)), nil
}

// NewSizeAlloc constructs a struct Size.
func NewSizeAlloc() *Size {
	var _cret *C.graphene_size_t // in

	_cret = C.graphene_size_alloc()

	var _size *Size // out

	_size = (*Size)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_size, func(v *Size) {
		C.graphene_size_free((*C.graphene_size_t)(unsafe.Pointer(v)))
	})

	return _size
}

// Native returns the underlying C source pointer.
func (s *Size) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// Equal checks whether the two give #graphene_size_t are equal.
func (a *Size) Equal(b *Size) bool {
	var _arg0 *C.graphene_size_t // out
	var _arg1 *C.graphene_size_t // out
	var _cret C._Bool            // in

	_arg0 = (*C.graphene_size_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_size_t)(unsafe.Pointer(b))

	_cret = C.graphene_size_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Free frees the resources allocated by graphene_size_alloc().
func (s *Size) free() {
	var _arg0 *C.graphene_size_t // out

	_arg0 = (*C.graphene_size_t)(unsafe.Pointer(s))

	C.graphene_size_free(_arg0)
}

// Init initializes a #graphene_size_t using the given width and height.
func (s *Size) Init(width float32, height float32) *Size {
	var _arg0 *C.graphene_size_t // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out
	var _cret *C.graphene_size_t // in

	_arg0 = (*C.graphene_size_t)(unsafe.Pointer(s))
	_arg1 = C.float(width)
	_arg2 = C.float(height)

	_cret = C.graphene_size_init(_arg0, _arg1, _arg2)

	var _size *Size // out

	_size = (*Size)(unsafe.Pointer(_cret))

	return _size
}

// InitFromSize initializes a #graphene_size_t using the width and height of the
// given src.
func (s *Size) InitFromSize(src *Size) *Size {
	var _arg0 *C.graphene_size_t // out
	var _arg1 *C.graphene_size_t // out
	var _cret *C.graphene_size_t // in

	_arg0 = (*C.graphene_size_t)(unsafe.Pointer(s))
	_arg1 = (*C.graphene_size_t)(unsafe.Pointer(src))

	_cret = C.graphene_size_init_from_size(_arg0, _arg1)

	var _size *Size // out

	_size = (*Size)(unsafe.Pointer(_cret))

	return _size
}

// Interpolate: linearly interpolates the two given #graphene_size_t using the
// given interpolation factor.
func (a *Size) Interpolate(b *Size, factor float64) Size {
	var _arg0 *C.graphene_size_t // out
	var _arg1 *C.graphene_size_t // out
	var _arg2 C.double           // out
	var _res Size

	_arg0 = (*C.graphene_size_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_size_t)(unsafe.Pointer(b))
	_arg2 = C.double(factor)

	C.graphene_size_interpolate(_arg0, _arg1, _arg2, (*C.graphene_size_t)(unsafe.Pointer(&_res)))

	return _res
}

// Scale scales the components of a #graphene_size_t using the given factor.
func (s *Size) Scale(factor float32) Size {
	var _arg0 *C.graphene_size_t // out
	var _arg1 C.float            // out
	var _res Size

	_arg0 = (*C.graphene_size_t)(unsafe.Pointer(s))
	_arg1 = C.float(factor)

	C.graphene_size_scale(_arg0, _arg1, (*C.graphene_size_t)(unsafe.Pointer(&_res)))

	return _res
}

// SizeZero: constant pointer to a zero #graphene_size_t, useful for equality
// checks and interpolations.
func SizeZero() *Size {
	var _cret *C.graphene_size_t // in

	_cret = C.graphene_size_zero()

	var _size *Size // out

	_size = (*Size)(unsafe.Pointer(_cret))

	return _size
}
