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
		{T: externglib.Type(C.gtk_paned_get_type()), F: marshalPaned},
	})
}

// Paned: `GtkPaned` has two panes, arranged either horizontally or vertically.
//
// !An example GtkPaned (panes.png)
//
// The division between the two panes is adjustable by the user by dragging a
// handle.
//
// Child widgets are added to the panes of the widget with
// [method@Gtk.Paned.set_start_child] and [method@Gtk.Paned.set_end_child]. The
// division between the two children is set by default from the size requests of
// the children, but it can be adjusted by the user.
//
// A paned widget draws a separator between the two child widgets and a small
// handle that the user can drag to adjust the division. It does not draw any
// relief around the children or around the separator. (The space in which the
// separator is called the gutter.) Often, it is useful to put each child inside
// a [class@Gtk.Frame] so that the gutter appears as a ridge. No separator is
// drawn if one of the children is missing.
//
// Each child has two options that can be set, @resize and @shrink. If @resize
// is true, then when the `GtkPaned` is resized, that child will expand or
// shrink along with the paned widget. If @shrink is true, then that child can
// be made smaller than its requisition by the user. Setting @shrink to false
// allows the application to set a minimum size. If @resize is false for both
// children, then this is treated as if @resize is true for both children.
//
// The application can set the position of the slider as if it were set by the
// user, by calling [method@Gtk.Paned.set_position].
//
//
// CSS nodes
//
// “` paned ├── <child> ├── separator[.wide] ╰── <child> “`
//
// `GtkPaned` has a main CSS node with name paned, and a subnode for the
// separator with name separator. The subnode gets a .wide style class when the
// paned is supposed to be wide.
//
// In horizontal orientation, the nodes are arranged based on the text
// direction, so in left-to-right mode, :first-child will select the leftmost
// child, while it will select the rightmost child in RTL layouts.
//
// Creating a paned widget with minimum sizes.
//
// “`c GtkWidget *hpaned = gtk_paned_new (GTK_ORIENTATION_HORIZONTAL); GtkWidget
// *frame1 = gtk_frame_new (NULL); GtkWidget *frame2 = gtk_frame_new (NULL);
//
// gtk_widget_set_size_request (hpaned, 200, -1);
//
// gtk_paned_set_start_child (GTK_PANED (hpaned), frame1);
// gtk_paned_set_start_child_resize (GTK_PANED (hpaned), TRUE);
// gtk_paned_set_start_child_shrink (GTK_PANED (hpaned), FALSE);
// gtk_widget_set_size_request (frame1, 50, -1);
//
// gtk_paned_set_end_child (GTK_PANED (hpaned), frame2);
// gtk_paned_set_end_child_resize (GTK_PANED (hpaned), FALSE);
// gtk_paned_set_end_child_shrink (GTK_PANED (hpaned), FALSE);
// gtk_widget_set_size_request (frame2, 50, -1); “`
type Paned interface {
	Widget

	// AsAccessible casts the class to the Accessible interface.
	AsAccessible() Accessible
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsConstraintTarget casts the class to the ConstraintTarget interface.
	AsConstraintTarget() ConstraintTarget
	// AsOrientable casts the class to the Orientable interface.
	AsOrientable() Orientable

	// EndChild retrieves the end child of the given `GtkPaned`.
	//
	// See also: `GtkPaned`:end-child
	EndChild() Widget
	// Position obtains the position of the divider between the two panes.
	Position() int
	// ResizeEndChild returns whether the end child can be resized.
	ResizeEndChild() bool
	// ResizeStartChild returns whether the start child can be resized.
	ResizeStartChild() bool
	// ShrinkEndChild returns whether the end child can be shrunk.
	ShrinkEndChild() bool
	// ShrinkStartChild returns whether the start child can be shrunk.
	ShrinkStartChild() bool
	// StartChild retrieves the start child of the given `GtkPaned`.
	//
	// See also: `GtkPaned`:start-child
	StartChild() Widget
	// WideHandle gets whether the separator should be wide.
	WideHandle() bool
	// SetEndChildPaned sets the end child of @paned to @child.
	SetEndChildPaned(child Widget)
	// SetPositionPaned sets the position of the divider between the two panes.
	SetPositionPaned(position int)
	// SetResizeEndChildPaned sets the `GtkPaned`:resize-end-child property
	SetResizeEndChildPaned(resize bool)
	// SetResizeStartChildPaned sets the `GtkPaned`:resize-start-child property
	SetResizeStartChildPaned(resize bool)
	// SetShrinkEndChildPaned sets the `GtkPaned`:shrink-end-child property
	SetShrinkEndChildPaned(resize bool)
	// SetShrinkStartChildPaned sets the `GtkPaned`:shrink-start-child property
	SetShrinkStartChildPaned(resize bool)
	// SetStartChildPaned sets the start child of @paned to @child.
	SetStartChildPaned(child Widget)
	// SetWideHandlePaned sets whether the separator should be wide.
	SetWideHandlePaned(wide bool)
}

// paned implements the Paned class.
type paned struct {
	Widget
}

// WrapPaned wraps a GObject to the right type. It is
// primarily used internally.
func WrapPaned(obj *externglib.Object) Paned {
	return paned{
		Widget: WrapWidget(obj),
	}
}

func marshalPaned(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPaned(obj), nil
}

// NewPaned creates a new `GtkPaned` widget.
func NewPaned(orientation Orientation) Paned {
	var _arg1 C.GtkOrientation // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)

	_cret = C.gtk_paned_new(_arg1)

	var _paned Paned // out

	_paned = WrapPaned(externglib.Take(unsafe.Pointer(_cret)))

	return _paned
}

func (p paned) EndChild() Widget {
	var _arg0 *C.GtkPaned  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_paned_get_end_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (p paned) Position() int {
	var _arg0 *C.GtkPaned // out
	var _cret C.int       // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_paned_get_position(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (p paned) ResizeEndChild() bool {
	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_paned_get_resize_end_child(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p paned) ResizeStartChild() bool {
	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_paned_get_resize_start_child(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p paned) ShrinkEndChild() bool {
	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_paned_get_shrink_end_child(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p paned) ShrinkStartChild() bool {
	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_paned_get_shrink_start_child(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p paned) StartChild() Widget {
	var _arg0 *C.GtkPaned  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_paned_get_start_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (p paned) WideHandle() bool {
	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_paned_get_wide_handle(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p paned) SetEndChildPaned(child Widget) {
	var _arg0 *C.GtkPaned  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_paned_set_end_child(_arg0, _arg1)
}

func (p paned) SetPositionPaned(position int) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	_arg1 = C.int(position)

	C.gtk_paned_set_position(_arg0, _arg1)
}

func (p paned) SetResizeEndChildPaned(resize bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	if resize {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_resize_end_child(_arg0, _arg1)
}

func (p paned) SetResizeStartChildPaned(resize bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	if resize {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_resize_start_child(_arg0, _arg1)
}

func (p paned) SetShrinkEndChildPaned(resize bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	if resize {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_shrink_end_child(_arg0, _arg1)
}

func (p paned) SetShrinkStartChildPaned(resize bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	if resize {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_shrink_start_child(_arg0, _arg1)
}

func (p paned) SetStartChildPaned(child Widget) {
	var _arg0 *C.GtkPaned  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_paned_set_start_child(_arg0, _arg1)
}

func (p paned) SetWideHandlePaned(wide bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	if wide {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_wide_handle(_arg0, _arg1)
}

func (p paned) AsAccessible() Accessible {
	return WrapAccessible(gextras.InternObject(p))
}

func (p paned) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(p))
}

func (p paned) AsConstraintTarget() ConstraintTarget {
	return WrapConstraintTarget(gextras.InternObject(p))
}

func (p paned) AsOrientable() Orientable {
	return WrapOrientable(gextras.InternObject(p))
}
