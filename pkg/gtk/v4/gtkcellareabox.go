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
		{T: externglib.Type(C.gtk_cell_area_box_get_type()), F: marshalCellAreaBoxxer},
	})
}

// CellAreaBoxxer describes CellAreaBox's methods.
type CellAreaBoxxer interface {
	// Spacing gets the spacing added between cell renderers.
	Spacing() int
	// PackEnd adds @renderer to @box, packed with reference to the end of @box.
	PackEnd(renderer CellRendererer, expand bool, align bool, fixed bool)
	// PackStart adds @renderer to @box, packed with reference to the start of
	// @box.
	PackStart(renderer CellRendererer, expand bool, align bool, fixed bool)
	// SetSpacing sets the spacing to add between cell renderers in @box.
	SetSpacing(spacing int)
}

// CellAreaBox: cell area that renders GtkCellRenderers into a row or a column
//
// The CellAreaBox renders cell renderers into a row or a column depending on
// its Orientation.
//
// GtkCellAreaBox uses a notion of packing. Packing refers to adding cell
// renderers with reference to a particular position in a CellAreaBox. There are
// two reference positions: the start and the end of the box. When the
// CellAreaBox is oriented in the GTK_ORIENTATION_VERTICAL orientation, the
// start is defined as the top of the box and the end is defined as the bottom.
// In the GTK_ORIENTATION_HORIZONTAL orientation start is defined as the left
// side and the end is defined as the right side.
//
// Alignments of CellRenderers rendered in adjacent rows can be configured by
// configuring the CellAreaBox align child cell property with
// gtk_cell_area_cell_set_property() or by specifying the "align" argument to
// gtk_cell_area_box_pack_start() and gtk_cell_area_box_pack_end().
type CellAreaBox struct {
	CellArea

	Orientable
}

var (
	_ CellAreaBoxxer  = (*CellAreaBox)(nil)
	_ gextras.Nativer = (*CellAreaBox)(nil)
)

func wrapCellAreaBox(obj *externglib.Object) CellAreaBoxxer {
	return &CellAreaBox{
		CellArea: CellArea{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			CellLayout: CellLayout{
				Object: obj,
			},
		},
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalCellAreaBoxxer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCellAreaBox(obj), nil
}

// NewCellAreaBox creates a new CellAreaBox.
func NewCellAreaBox() *CellAreaBox {
	var _cret *C.GtkCellArea // in

	_cret = C.gtk_cell_area_box_new()

	var _cellAreaBox *CellAreaBox // out

	_cellAreaBox = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*CellAreaBox)

	return _cellAreaBox
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *CellAreaBox) Native() uintptr {
	return v.CellArea.InitiallyUnowned.Object.Native()
}

// Spacing gets the spacing added between cell renderers.
func (box *CellAreaBox) Spacing() int {
	var _arg0 *C.GtkCellAreaBox // out
	var _cret C.int             // in

	_arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_cell_area_box_get_spacing(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PackEnd adds @renderer to @box, packed with reference to the end of @box.
//
// The @renderer is packed after (away from end of) any other CellRenderer
// packed with reference to the end of @box.
func (box *CellAreaBox) PackEnd(renderer CellRendererer, expand bool, align bool, fixed bool) {
	var _arg0 *C.GtkCellAreaBox  // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.gboolean         // out
	var _arg3 C.gboolean         // out
	var _arg4 C.gboolean         // out

	_arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer((renderer).(gextras.Nativer).Native()))
	if expand {
		_arg2 = C.TRUE
	}
	if align {
		_arg3 = C.TRUE
	}
	if fixed {
		_arg4 = C.TRUE
	}

	C.gtk_cell_area_box_pack_end(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// PackStart adds @renderer to @box, packed with reference to the start of @box.
//
// The @renderer is packed after any other CellRenderer packed with reference to
// the start of @box.
func (box *CellAreaBox) PackStart(renderer CellRendererer, expand bool, align bool, fixed bool) {
	var _arg0 *C.GtkCellAreaBox  // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.gboolean         // out
	var _arg3 C.gboolean         // out
	var _arg4 C.gboolean         // out

	_arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer((renderer).(gextras.Nativer).Native()))
	if expand {
		_arg2 = C.TRUE
	}
	if align {
		_arg3 = C.TRUE
	}
	if fixed {
		_arg4 = C.TRUE
	}

	C.gtk_cell_area_box_pack_start(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// SetSpacing sets the spacing to add between cell renderers in @box.
func (box *CellAreaBox) SetSpacing(spacing int) {
	var _arg0 *C.GtkCellAreaBox // out
	var _arg1 C.int             // out

	_arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.int(spacing)

	C.gtk_cell_area_box_set_spacing(_arg0, _arg1)
}
