// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeVPaned = coreglib.Type(C.gtk_vpaned_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVPaned, F: marshalVPaned},
	})
}

// VPanedOverrides contains methods that are overridable.
type VPanedOverrides struct {
}

func defaultVPanedOverrides(v *VPaned) VPanedOverrides {
	return VPanedOverrides{}
}

// VPaned widget is a container widget with two children arranged vertically.
// The division between the two panes is adjustable by the user by dragging a
// handle. See Paned for details.
//
// GtkVPaned has been deprecated, use Paned instead.
type VPaned struct {
	_ [0]func() // equal guard
	Paned
}

var (
	_ Containerer       = (*VPaned)(nil)
	_ coreglib.Objector = (*VPaned)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*VPaned, *VPanedClass, VPanedOverrides](
		GTypeVPaned,
		initVPanedClass,
		wrapVPaned,
		defaultVPanedOverrides,
	)
}

func initVPanedClass(gclass unsafe.Pointer, overrides VPanedOverrides, classInitFunc func(*VPanedClass)) {
	if classInitFunc != nil {
		class := (*VPanedClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapVPaned(obj *coreglib.Object) *VPaned {
	return &VPaned{
		Paned: Paned{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalVPaned(p uintptr) (interface{}, error) {
	return wrapVPaned(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewVPaned: create a new VPaned
//
// Deprecated: Use gtk_paned_new() with GTK_ORIENTATION_VERTICAL instead.
//
// The function returns the following values:
//
//   - vPaned: new VPaned.
//
func NewVPaned() *VPaned {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_vpaned_new()

	var _vPaned *VPaned // out

	_vPaned = wrapVPaned(coreglib.Take(unsafe.Pointer(_cret)))

	return _vPaned
}

// VPanedClass: instance of this type is always passed by reference.
type VPanedClass struct {
	*vPanedClass
}

// vPanedClass is the struct that's finalized.
type vPanedClass struct {
	native *C.GtkVPanedClass
}

func (v *VPanedClass) ParentClass() *PanedClass {
	valptr := &v.native.parent_class
	var _v *PanedClass // out
	_v = (*PanedClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
