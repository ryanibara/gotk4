// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtkcellrendererprogress.go.
var GTypeCellRendererProgress = externglib.Type(C.gtk_cell_renderer_progress_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeCellRendererProgress, F: marshalCellRendererProgress},
	})
}

// CellRendererProgressOverrider contains methods that are overridable.
type CellRendererProgressOverrider interface {
	externglib.Objector
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

func classInitCellRendererProgresser(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapCellRendererProgress(obj *externglib.Object) *CellRendererProgress {
	return &CellRendererProgress{
		CellRenderer: CellRenderer{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
		},
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalCellRendererProgress(p uintptr) (interface{}, error) {
	return wrapCellRendererProgress(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
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

	_cellRendererProgress = wrapCellRendererProgress(externglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererProgress
}
