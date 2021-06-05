// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_toggle_get_type()), F: marshalCellRendererToggle},
	})
}

// CellRendererToggle renders a toggle button in a cell. The button is drawn as
// a radio or a checkbutton, depending on the CellRendererToggle:radio property.
// When activated, it emits the CellRendererToggle::toggled signal.
type CellRendererToggle interface {
	CellRenderer

	// Activatable returns whether the cell renderer is activatable. See
	// gtk_cell_renderer_toggle_set_activatable().
	Activatable() bool
	// Active returns whether the cell renderer is active. See
	// gtk_cell_renderer_toggle_set_active().
	Active() bool
	// Radio returns whether we’re rendering radio toggles rather than
	// checkboxes.
	Radio() bool
	// SetActivatable makes the cell renderer activatable.
	SetActivatable(setting bool)
	// SetActive activates or deactivates a cell renderer.
	SetActive(setting bool)
	// SetRadio: if @radio is true, the cell renderer renders a radio toggle
	// (i.e. a toggle in a group of mutually-exclusive toggles). If false, it
	// renders a check toggle (a standalone boolean option). This can be set
	// globally for the cell renderer, or changed just before rendering each
	// cell in the model (for TreeView, you set up a per-row setting using
	// TreeViewColumn to associate model columns with cell renderer properties).
	SetRadio(radio bool)
}

// cellRendererToggle implements the CellRendererToggle interface.
type cellRendererToggle struct {
	CellRenderer
}

var _ CellRendererToggle = (*cellRendererToggle)(nil)

// WrapCellRendererToggle wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellRendererToggle(obj *externglib.Object) CellRendererToggle {
	return CellRendererToggle{
		CellRenderer: WrapCellRenderer(obj),
	}
}

func marshalCellRendererToggle(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellRendererToggle(obj), nil
}

// NewCellRendererToggle constructs a class CellRendererToggle.
func NewCellRendererToggle() CellRendererToggle {
	var cret C.GtkCellRendererToggle
	var ret1 CellRendererToggle

	cret = C.gtk_cell_renderer_toggle_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(CellRendererToggle)

	return ret1
}

// Activatable returns whether the cell renderer is activatable. See
// gtk_cell_renderer_toggle_set_activatable().
func (t cellRendererToggle) Activatable() bool {
	var arg0 *C.GtkCellRendererToggle

	arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(t.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_cell_renderer_toggle_get_activatable(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Active returns whether the cell renderer is active. See
// gtk_cell_renderer_toggle_set_active().
func (t cellRendererToggle) Active() bool {
	var arg0 *C.GtkCellRendererToggle

	arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(t.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_cell_renderer_toggle_get_active(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Radio returns whether we’re rendering radio toggles rather than
// checkboxes.
func (t cellRendererToggle) Radio() bool {
	var arg0 *C.GtkCellRendererToggle

	arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(t.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_cell_renderer_toggle_get_radio(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// SetActivatable makes the cell renderer activatable.
func (t cellRendererToggle) SetActivatable(setting bool) {
	var arg0 *C.GtkCellRendererToggle
	var arg1 C.gboolean

	arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(t.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_cell_renderer_toggle_set_activatable(arg0, setting)
}

// SetActive activates or deactivates a cell renderer.
func (t cellRendererToggle) SetActive(setting bool) {
	var arg0 *C.GtkCellRendererToggle
	var arg1 C.gboolean

	arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(t.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_cell_renderer_toggle_set_active(arg0, setting)
}

// SetRadio: if @radio is true, the cell renderer renders a radio toggle
// (i.e. a toggle in a group of mutually-exclusive toggles). If false, it
// renders a check toggle (a standalone boolean option). This can be set
// globally for the cell renderer, or changed just before rendering each
// cell in the model (for TreeView, you set up a per-row setting using
// TreeViewColumn to associate model columns with cell renderer properties).
func (t cellRendererToggle) SetRadio(radio bool) {
	var arg0 *C.GtkCellRendererToggle
	var arg1 C.gboolean

	arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(t.Native()))
	if radio {
		arg1 = C.gboolean(1)
	}

	C.gtk_cell_renderer_toggle_set_radio(arg0, radio)
}
