// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: graphene-gobject-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_quad_get_type()), F: marshalQuad},
	})
}

// Quad: a 4 vertex quadrilateral, as represented by four #graphene_point_t.
//
// The contents of a #graphene_quad_t are private and should never be accessed
// directly.
type Quad struct {
	native C.graphene_quad_t
}

// WrapQuad wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapQuad(ptr unsafe.Pointer) *Quad {
	if ptr == nil {
		return nil
	}

	return (*Quad)(ptr)
}

func marshalQuad(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapQuad(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (q *Quad) Native() unsafe.Pointer {
	return unsafe.Pointer(&q.native)
}

// NewQuadAlloc constructs a struct Quad.
func NewQuadAlloc() *Quad {
	ret := C.graphene_quad_alloc()

	var ret0 *Quad

	{
		ret0 = WrapQuad(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *Quad) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// Bounds computes the bounding rectangle of @q and places it into @r.
func (q *Quad) Bounds() Rect {
	var arg0 *C.graphene_quad_t
	var arg1 *C.graphene_rect_t // out

	arg0 = (*C.graphene_quad_t)(q.Native())

	C.graphene_quad_bounds(arg0, &arg1)

	var ret0 *Rect

	{
		ret0 = WrapRect(unsafe.Pointer(arg1))
	}

	return ret0
}

// Contains checks if the given #graphene_quad_t contains the given
// #graphene_point_t.
func (q *Quad) Contains(p *Point) bool {
	var arg0 *C.graphene_quad_t
	var arg1 *C.graphene_point_t

	arg0 = (*C.graphene_quad_t)(q.Native())
	arg1 = (*C.graphene_point_t)(p.Native())

	ret := C.graphene_quad_contains(arg0, arg1)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Free frees the resources allocated by graphene_quad_alloc()
func (q *Quad) Free() {
	var arg0 *C.graphene_quad_t

	arg0 = (*C.graphene_quad_t)(q.Native())

	C.graphene_quad_free(arg0)
}

// Point retrieves the point of a #graphene_quad_t at the given index.
func (q *Quad) Point(index_ uint) *Point {
	var arg0 *C.graphene_quad_t
	var arg1 C.uint

	arg0 = (*C.graphene_quad_t)(q.Native())
	arg1 = C.uint(index_)

	ret := C.graphene_quad_get_point(arg0, arg1)

	var ret0 *Point

	{
		ret0 = WrapPoint(unsafe.Pointer(ret))
	}

	return ret0
}

// Init initializes a #graphene_quad_t with the given points.
func (q *Quad) Init(p1 *Point, p2 *Point, p3 *Point, p4 *Point) *Quad {
	var arg0 *C.graphene_quad_t
	var arg1 *C.graphene_point_t
	var arg2 *C.graphene_point_t
	var arg3 *C.graphene_point_t
	var arg4 *C.graphene_point_t

	arg0 = (*C.graphene_quad_t)(q.Native())
	arg1 = (*C.graphene_point_t)(p1.Native())
	arg2 = (*C.graphene_point_t)(p2.Native())
	arg3 = (*C.graphene_point_t)(p3.Native())
	arg4 = (*C.graphene_point_t)(p4.Native())

	ret := C.graphene_quad_init(arg0, arg1, arg2, arg3, arg4)

	var ret0 *Quad

	{
		ret0 = WrapQuad(unsafe.Pointer(ret))
	}

	return ret0
}

// InitFromPoints initializes a #graphene_quad_t using an array of points.
func (q *Quad) InitFromPoints(points [4]Point) *Quad {
	var arg0 *C.graphene_quad_t
	var arg1 *C.graphene_point_t

	arg0 = (*C.graphene_quad_t)(q.Native())
	arg1 = (*C.graphene_point_t)(&points)
	defer runtime.KeepAlive(&points)

	ret := C.graphene_quad_init_from_points(arg0, arg1)

	var ret0 *Quad

	{
		ret0 = WrapQuad(unsafe.Pointer(ret))
	}

	return ret0
}

// InitFromRect initializes a #graphene_quad_t using the four corners of the
// given #graphene_rect_t.
func (q *Quad) InitFromRect(r *Rect) *Quad {
	var arg0 *C.graphene_quad_t
	var arg1 *C.graphene_rect_t

	arg0 = (*C.graphene_quad_t)(q.Native())
	arg1 = (*C.graphene_rect_t)(r.Native())

	ret := C.graphene_quad_init_from_rect(arg0, arg1)

	var ret0 *Quad

	{
		ret0 = WrapQuad(unsafe.Pointer(ret))
	}

	return ret0
}
