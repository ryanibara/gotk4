// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"unsafe"

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
	GTypeCellRendererProgress = coreglib.Type(C.gtk_cell_renderer_progress_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCellRendererProgress, F: marshalCellRendererProgress},
	})
}

// CellRendererProgressOverrider contains methods that are overridable.
type CellRendererProgressOverrider interface {
}

// CellRendererProgress renders a numeric value as a progress par in a cell.
// Additionally, it can display a text on top of the progress bar.
//
// The CellRendererProgress cell renderer was added in GTK+ 2.6.
type CellRendererProgress struct {
	_ [0]func() // equal guard
	CellRenderer

	Orientable
}

var (
	_ CellRendererer = (*CellRendererProgress)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeCellRendererProgress,
		GoType:        reflect.TypeOf((*CellRendererProgress)(nil)),
		InitClass:     initClassCellRendererProgress,
		FinalizeClass: finalizeClassCellRendererProgress,
	})
}

func initClassCellRendererProgress(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		InitCellRendererProgress(*CellRendererProgressClass)
	}); ok {
		klass := (*CellRendererProgressClass)(gextras.NewStructNative(gclass))
		goval.InitCellRendererProgress(klass)
	}
}

func finalizeClassCellRendererProgress(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		FinalizeCellRendererProgress(*CellRendererProgressClass)
	}); ok {
		klass := (*CellRendererProgressClass)(gextras.NewStructNative(gclass))
		goval.FinalizeCellRendererProgress(klass)
	}
}

func wrapCellRendererProgress(obj *coreglib.Object) *CellRendererProgress {
	return &CellRendererProgress{
		CellRenderer: CellRenderer{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
		},
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalCellRendererProgress(p uintptr) (interface{}, error) {
	return wrapCellRendererProgress(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCellRendererProgress creates a new CellRendererProgress.
//
// The function returns the following values:
//
//    - cellRendererProgress: new cell renderer.
//
func NewCellRendererProgress() *CellRendererProgress {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_progress_new()

	var _cellRendererProgress *CellRendererProgress // out

	_cellRendererProgress = wrapCellRendererProgress(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererProgress
}

// CellRendererProgressClass: instance of this type is always passed by
// reference.
type CellRendererProgressClass struct {
	*cellRendererProgressClass
}

// cellRendererProgressClass is the struct that's finalized.
type cellRendererProgressClass struct {
	native *C.GtkCellRendererProgressClass
}

func (c *CellRendererProgressClass) ParentClass() *CellRendererClass {
	valptr := &c.native.parent_class
	var _v *CellRendererClass // out
	_v = (*CellRendererClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
