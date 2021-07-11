// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_constraint_guide_get_type()), F: marshalConstraintGuider},
	})
}

// ConstraintGuider describes ConstraintGuide's methods.
type ConstraintGuider interface {
	// MaxSize gets the maximum size of @guide.
	MaxSize(width *int, height *int)
	// MinSize gets the minimum size of @guide.
	MinSize(width *int, height *int)
	// Name retrieves the name set using gtk_constraint_guide_set_name().
	Name() string
	// NatSize gets the natural size of @guide.
	NatSize(width *int, height *int)
	// Strength retrieves the strength set using
	// gtk_constraint_guide_set_strength().
	Strength() ConstraintStrength
	// SetMaxSize sets the maximum size of @guide.
	SetMaxSize(width int, height int)
	// SetMinSize sets the minimum size of @guide.
	SetMinSize(width int, height int)
	// SetName sets a name for the given `GtkConstraintGuide`.
	SetName(name string)
	// SetNatSize sets the natural size of @guide.
	SetNatSize(width int, height int)
}

// ConstraintGuide: `GtkConstraintGuide` is an invisible layout element in a
// `GtkConstraintLayout`.
//
// The `GtkConstraintLayout` treats guides like widgets. They can be used as the
// source or target of a `GtkConstraint`.
//
// Guides have a minimum, maximum and natural size. Depending on the constraints
// that are applied, they can act like a guideline that widgets can be aligned
// to, or like *flexible space*.
//
// Unlike a `GtkWidget`, a `GtkConstraintGuide` will not be drawn.
type ConstraintGuide struct {
	*externglib.Object

	ConstraintTarget
}

var (
	_ ConstraintGuider = (*ConstraintGuide)(nil)
	_ gextras.Nativer  = (*ConstraintGuide)(nil)
)

func wrapConstraintGuide(obj *externglib.Object) ConstraintGuider {
	return &ConstraintGuide{
		Object: obj,
		ConstraintTarget: ConstraintTarget{
			Object: obj,
		},
	}
}

func marshalConstraintGuider(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapConstraintGuide(obj), nil
}

// NewConstraintGuide creates a new `GtkConstraintGuide` object.
func NewConstraintGuide() *ConstraintGuide {
	var _cret *C.GtkConstraintGuide // in

	_cret = C.gtk_constraint_guide_new()

	var _constraintGuide *ConstraintGuide // out

	_constraintGuide = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*ConstraintGuide)

	return _constraintGuide
}

// MaxSize gets the maximum size of @guide.
func (guide *ConstraintGuide) MaxSize(width *int, height *int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 *C.int                // out
	var _arg2 *C.int                // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))
	_arg1 = (*C.int)(unsafe.Pointer(width))
	_arg2 = (*C.int)(unsafe.Pointer(height))

	C.gtk_constraint_guide_get_max_size(_arg0, _arg1, _arg2)
}

// MinSize gets the minimum size of @guide.
func (guide *ConstraintGuide) MinSize(width *int, height *int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 *C.int                // out
	var _arg2 *C.int                // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))
	_arg1 = (*C.int)(unsafe.Pointer(width))
	_arg2 = (*C.int)(unsafe.Pointer(height))

	C.gtk_constraint_guide_get_min_size(_arg0, _arg1, _arg2)
}

// Name retrieves the name set using gtk_constraint_guide_set_name().
func (guide *ConstraintGuide) Name() string {
	var _arg0 *C.GtkConstraintGuide // out
	var _cret *C.char               // in

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))

	_cret = C.gtk_constraint_guide_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// NatSize gets the natural size of @guide.
func (guide *ConstraintGuide) NatSize(width *int, height *int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 *C.int                // out
	var _arg2 *C.int                // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))
	_arg1 = (*C.int)(unsafe.Pointer(width))
	_arg2 = (*C.int)(unsafe.Pointer(height))

	C.gtk_constraint_guide_get_nat_size(_arg0, _arg1, _arg2)
}

// Strength retrieves the strength set using
// gtk_constraint_guide_set_strength().
func (guide *ConstraintGuide) Strength() ConstraintStrength {
	var _arg0 *C.GtkConstraintGuide   // out
	var _cret C.GtkConstraintStrength // in

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))

	_cret = C.gtk_constraint_guide_get_strength(_arg0)

	var _constraintStrength ConstraintStrength // out

	_constraintStrength = ConstraintStrength(_cret)

	return _constraintStrength
}

// SetMaxSize sets the maximum size of @guide.
//
// If @guide is attached to a `GtkConstraintLayout`, the constraints will be
// updated to reflect the new size.
func (guide *ConstraintGuide) SetMaxSize(width int, height int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 C.int                 // out
	var _arg2 C.int                 // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	C.gtk_constraint_guide_set_max_size(_arg0, _arg1, _arg2)
}

// SetMinSize sets the minimum size of @guide.
//
// If @guide is attached to a `GtkConstraintLayout`, the constraints will be
// updated to reflect the new size.
func (guide *ConstraintGuide) SetMinSize(width int, height int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 C.int                 // out
	var _arg2 C.int                 // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	C.gtk_constraint_guide_set_min_size(_arg0, _arg1, _arg2)
}

// SetName sets a name for the given `GtkConstraintGuide`.
//
// The name is useful for debugging purposes.
func (guide *ConstraintGuide) SetName(name string) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 *C.char               // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_constraint_guide_set_name(_arg0, _arg1)
}

// SetNatSize sets the natural size of @guide.
//
// If @guide is attached to a `GtkConstraintLayout`, the constraints will be
// updated to reflect the new size.
func (guide *ConstraintGuide) SetNatSize(width int, height int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 C.int                 // out
	var _arg2 C.int                 // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	C.gtk_constraint_guide_set_nat_size(_arg0, _arg1, _arg2)
}
