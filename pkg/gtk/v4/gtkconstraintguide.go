// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_constraint_guide_get_type()), F: marshalConstraintGuider},
	})
}

// ConstraintGuide: GtkConstraintGuide is an invisible layout element in a
// GtkConstraintLayout.
//
// The GtkConstraintLayout treats guides like widgets. They can be used as the
// source or target of a GtkConstraint.
//
// Guides have a minimum, maximum and natural size. Depending on the constraints
// that are applied, they can act like a guideline that widgets can be aligned
// to, or like *flexible space*.
//
// Unlike a GtkWidget, a GtkConstraintGuide will not be drawn.
type ConstraintGuide struct {
	*externglib.Object

	ConstraintTarget
}

func wrapConstraintGuide(obj *externglib.Object) *ConstraintGuide {
	return &ConstraintGuide{
		Object: obj,
		ConstraintTarget: ConstraintTarget{
			Object: obj,
		},
	}
}

func marshalConstraintGuider(p uintptr) (interface{}, error) {
	return wrapConstraintGuide(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewConstraintGuide creates a new GtkConstraintGuide object.
func NewConstraintGuide() *ConstraintGuide {
	var _cret *C.GtkConstraintGuide // in

	_cret = C.gtk_constraint_guide_new()

	var _constraintGuide *ConstraintGuide // out

	_constraintGuide = wrapConstraintGuide(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _constraintGuide
}

// MaxSize gets the maximum size of guide.
//
// The function takes the following parameters:
//
//    - width: return location for the maximum width, or NULL.
//    - height: return location for the maximum height, or NULL.
//
func (guide *ConstraintGuide) MaxSize(width, height *int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 *C.int                // out
	var _arg2 *C.int                // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))
	if width != nil {
		_arg1 = (*C.int)(unsafe.Pointer(width))
	}
	if height != nil {
		_arg2 = (*C.int)(unsafe.Pointer(height))
	}

	C.gtk_constraint_guide_get_max_size(_arg0, _arg1, _arg2)
	runtime.KeepAlive(guide)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// MinSize gets the minimum size of guide.
//
// The function takes the following parameters:
//
//    - width: return location for the minimum width, or NULL.
//    - height: return location for the minimum height, or NULL.
//
func (guide *ConstraintGuide) MinSize(width, height *int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 *C.int                // out
	var _arg2 *C.int                // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))
	if width != nil {
		_arg1 = (*C.int)(unsafe.Pointer(width))
	}
	if height != nil {
		_arg2 = (*C.int)(unsafe.Pointer(height))
	}

	C.gtk_constraint_guide_get_min_size(_arg0, _arg1, _arg2)
	runtime.KeepAlive(guide)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// Name retrieves the name set using gtk_constraint_guide_set_name().
func (guide *ConstraintGuide) Name() string {
	var _arg0 *C.GtkConstraintGuide // out
	var _cret *C.char               // in

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))

	_cret = C.gtk_constraint_guide_get_name(_arg0)
	runtime.KeepAlive(guide)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// NatSize gets the natural size of guide.
//
// The function takes the following parameters:
//
//    - width: return location for the natural width, or NULL.
//    - height: return location for the natural height, or NULL.
//
func (guide *ConstraintGuide) NatSize(width, height *int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 *C.int                // out
	var _arg2 *C.int                // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))
	if width != nil {
		_arg1 = (*C.int)(unsafe.Pointer(width))
	}
	if height != nil {
		_arg2 = (*C.int)(unsafe.Pointer(height))
	}

	C.gtk_constraint_guide_get_nat_size(_arg0, _arg1, _arg2)
	runtime.KeepAlive(guide)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// Strength retrieves the strength set using
// gtk_constraint_guide_set_strength().
func (guide *ConstraintGuide) Strength() ConstraintStrength {
	var _arg0 *C.GtkConstraintGuide   // out
	var _cret C.GtkConstraintStrength // in

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))

	_cret = C.gtk_constraint_guide_get_strength(_arg0)
	runtime.KeepAlive(guide)

	var _constraintStrength ConstraintStrength // out

	_constraintStrength = ConstraintStrength(_cret)

	return _constraintStrength
}

// SetMaxSize sets the maximum size of guide.
//
// If guide is attached to a GtkConstraintLayout, the constraints will be
// updated to reflect the new size.
//
// The function takes the following parameters:
//
//    - width: new maximum width, or -1 to not change it.
//    - height: new maximum height, or -1 to not change it.
//
func (guide *ConstraintGuide) SetMaxSize(width, height int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 C.int                 // out
	var _arg2 C.int                 // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	C.gtk_constraint_guide_set_max_size(_arg0, _arg1, _arg2)
	runtime.KeepAlive(guide)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// SetMinSize sets the minimum size of guide.
//
// If guide is attached to a GtkConstraintLayout, the constraints will be
// updated to reflect the new size.
//
// The function takes the following parameters:
//
//    - width: new minimum width, or -1 to not change it.
//    - height: new minimum height, or -1 to not change it.
//
func (guide *ConstraintGuide) SetMinSize(width, height int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 C.int                 // out
	var _arg2 C.int                 // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	C.gtk_constraint_guide_set_min_size(_arg0, _arg1, _arg2)
	runtime.KeepAlive(guide)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// SetName sets a name for the given GtkConstraintGuide.
//
// The name is useful for debugging purposes.
//
// The function takes the following parameters:
//
//    - name for the guide.
//
func (guide *ConstraintGuide) SetName(name string) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 *C.char               // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))
	if name != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_constraint_guide_set_name(_arg0, _arg1)
	runtime.KeepAlive(guide)
	runtime.KeepAlive(name)
}

// SetNatSize sets the natural size of guide.
//
// If guide is attached to a GtkConstraintLayout, the constraints will be
// updated to reflect the new size.
//
// The function takes the following parameters:
//
//    - width: new natural width, or -1 to not change it.
//    - height: new natural height, or -1 to not change it.
//
func (guide *ConstraintGuide) SetNatSize(width, height int) {
	var _arg0 *C.GtkConstraintGuide // out
	var _arg1 C.int                 // out
	var _arg2 C.int                 // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	C.gtk_constraint_guide_set_nat_size(_arg0, _arg1, _arg2)
	runtime.KeepAlive(guide)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// SetStrength sets the strength of the constraint on the natural size of the
// given GtkConstraintGuide.
//
// The function takes the following parameters:
//
//    - strength of the constraint.
//
func (guide *ConstraintGuide) SetStrength(strength ConstraintStrength) {
	var _arg0 *C.GtkConstraintGuide   // out
	var _arg1 C.GtkConstraintStrength // out

	_arg0 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))
	_arg1 = C.GtkConstraintStrength(strength)

	C.gtk_constraint_guide_set_strength(_arg0, _arg1)
	runtime.KeepAlive(guide)
	runtime.KeepAlive(strength)
}
