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
		{T: externglib.Type(C.graphene_quad_get_type()), F: marshalQuad},
	})
}

// Quad: 4 vertex quadrilateral, as represented by four #graphene_point_t.
//
// The contents of a #graphene_quad_t are private and should never be accessed
// directly.
type Quad struct {
	native C.graphene_quad_t
}

func marshalQuad(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Quad)(unsafe.Pointer(b)), nil
}

// NewQuadAlloc constructs a struct Quad.
func NewQuadAlloc() *Quad {
	var _cret *C.graphene_quad_t // in

	_cret = C.graphene_quad_alloc()

	var _quad *Quad // out

	_quad = (*Quad)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_quad, func(v *Quad) {
		C.graphene_quad_free((*C.graphene_quad_t)(unsafe.Pointer(v)))
	})

	return _quad
}

// Native returns the underlying C source pointer.
func (q *Quad) Native() unsafe.Pointer {
	return unsafe.Pointer(&q.native)
}

// Bounds computes the bounding rectangle of q and places it into r.
func (q *Quad) Bounds() Rect {
	var _arg0 *C.graphene_quad_t // out
	var _r Rect

	_arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q))

	C.graphene_quad_bounds(_arg0, (*C.graphene_rect_t)(unsafe.Pointer(&_r)))

	return _r
}

// Contains checks if the given #graphene_quad_t contains the given
// #graphene_point_t.
func (q *Quad) Contains(p *Point) bool {
	var _arg0 *C.graphene_quad_t  // out
	var _arg1 *C.graphene_point_t // out
	var _cret C._Bool             // in

	_arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q))
	_arg1 = (*C.graphene_point_t)(unsafe.Pointer(p))

	_cret = C.graphene_quad_contains(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Free frees the resources allocated by graphene_quad_alloc()
func (q *Quad) free() {
	var _arg0 *C.graphene_quad_t // out

	_arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q))

	C.graphene_quad_free(_arg0)
}

// Point retrieves the point of a #graphene_quad_t at the given index.
func (q *Quad) Point(index_ uint) *Point {
	var _arg0 *C.graphene_quad_t  // out
	var _arg1 C.uint              // out
	var _cret *C.graphene_point_t // in

	_arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q))
	_arg1 = C.uint(index_)

	_cret = C.graphene_quad_get_point(_arg0, _arg1)

	var _point *Point // out

	_point = (*Point)(unsafe.Pointer(_cret))

	return _point
}

// Init initializes a #graphene_quad_t with the given points.
func (q *Quad) Init(p1 *Point, p2 *Point, p3 *Point, p4 *Point) *Quad {
	var _arg0 *C.graphene_quad_t  // out
	var _arg1 *C.graphene_point_t // out
	var _arg2 *C.graphene_point_t // out
	var _arg3 *C.graphene_point_t // out
	var _arg4 *C.graphene_point_t // out
	var _cret *C.graphene_quad_t  // in

	_arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q))
	_arg1 = (*C.graphene_point_t)(unsafe.Pointer(p1))
	_arg2 = (*C.graphene_point_t)(unsafe.Pointer(p2))
	_arg3 = (*C.graphene_point_t)(unsafe.Pointer(p3))
	_arg4 = (*C.graphene_point_t)(unsafe.Pointer(p4))

	_cret = C.graphene_quad_init(_arg0, _arg1, _arg2, _arg3, _arg4)

	var _quad *Quad // out

	_quad = (*Quad)(unsafe.Pointer(_cret))

	return _quad
}

// InitFromPoints initializes a #graphene_quad_t using an array of points.
func (q *Quad) InitFromPoints(points [4]Point) *Quad {
	var _arg0 *C.graphene_quad_t // out
	var _arg1 *C.graphene_point_t
	var _cret *C.graphene_quad_t // in

	_arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q))
	_arg1 = (*C.graphene_point_t)(unsafe.Pointer(&points))

	_cret = C.graphene_quad_init_from_points(_arg0, _arg1)

	var _quad *Quad // out

	_quad = (*Quad)(unsafe.Pointer(_cret))

	return _quad
}

// InitFromRect initializes a #graphene_quad_t using the four corners of the
// given #graphene_rect_t.
func (q *Quad) InitFromRect(r *Rect) *Quad {
	var _arg0 *C.graphene_quad_t // out
	var _arg1 *C.graphene_rect_t // out
	var _cret *C.graphene_quad_t // in

	_arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(r))

	_cret = C.graphene_quad_init_from_rect(_arg0, _arg1)

	var _quad *Quad // out

	_quad = (*Quad)(unsafe.Pointer(_cret))

	return _quad
}
