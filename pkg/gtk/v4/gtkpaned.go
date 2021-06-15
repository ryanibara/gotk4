// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
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
	Accessible
	Buildable
	ConstraintTarget
	Orientable

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
	// SetEndChild sets the end child of @paned to @child.
	SetEndChild(child Widget)
	// SetPosition sets the position of the divider between the two panes.
	SetPosition(position int)
	// SetResizeEndChild sets the `GtkPaned`:resize-end-child property
	SetResizeEndChild(resize bool)
	// SetResizeStartChild sets the `GtkPaned`:resize-start-child property
	SetResizeStartChild(resize bool)
	// SetShrinkEndChild sets the `GtkPaned`:shrink-end-child property
	SetShrinkEndChild(resize bool)
	// SetShrinkStartChild sets the `GtkPaned`:shrink-start-child property
	SetShrinkStartChild(resize bool)
	// SetStartChild sets the start child of @paned to @child.
	SetStartChild(child Widget)
	// SetWideHandle sets whether the separator should be wide.
	SetWideHandle(wide bool)
}

// paned implements the Paned class.
type paned struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Orientable
}

var _ Paned = (*paned)(nil)

// WrapPaned wraps a GObject to the right type. It is
// primarily used internally.
func WrapPaned(obj *externglib.Object) Paned {
	return paned{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Orientable:       WrapOrientable(obj),
	}
}

func marshalPaned(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPaned(obj), nil
}

// NewPaned constructs a class Paned.
func NewPaned(orientation Orientation) Paned {
	var _arg1 C.GtkOrientation // out
	var _cret C.GtkPaned       // in

	_arg1 = (C.GtkOrientation)(orientation)

	_cret = C.gtk_paned_new(_arg1)

	var _paned Paned // out

	_paned = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Paned)

	return _paned
}

// EndChild retrieves the end child of the given `GtkPaned`.
//
// See also: `GtkPaned`:end-child
func (p paned) EndChild() Widget {
	var _arg0 *C.GtkPaned  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_paned_get_end_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// Position obtains the position of the divider between the two panes.
func (p paned) Position() int {
	var _arg0 *C.GtkPaned // out
	var _cret C.int       // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_paned_get_position(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// ResizeEndChild returns whether the end child can be resized.
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

// ResizeStartChild returns whether the start child can be resized.
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

// ShrinkEndChild returns whether the end child can be shrunk.
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

// ShrinkStartChild returns whether the start child can be shrunk.
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

// StartChild retrieves the start child of the given `GtkPaned`.
//
// See also: `GtkPaned`:start-child
func (p paned) StartChild() Widget {
	var _arg0 *C.GtkPaned  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_paned_get_start_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// WideHandle gets whether the separator should be wide.
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

// SetEndChild sets the end child of @paned to @child.
func (p paned) SetEndChild(child Widget) {
	var _arg0 *C.GtkPaned  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_paned_set_end_child(_arg0, _arg1)
}

// SetPosition sets the position of the divider between the two panes.
func (p paned) SetPosition(position int) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	_arg1 = (C.int)(position)

	C.gtk_paned_set_position(_arg0, _arg1)
}

// SetResizeEndChild sets the `GtkPaned`:resize-end-child property
func (p paned) SetResizeEndChild(resize bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	if resize {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_resize_end_child(_arg0, _arg1)
}

// SetResizeStartChild sets the `GtkPaned`:resize-start-child property
func (p paned) SetResizeStartChild(resize bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	if resize {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_resize_start_child(_arg0, _arg1)
}

// SetShrinkEndChild sets the `GtkPaned`:shrink-end-child property
func (p paned) SetShrinkEndChild(resize bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	if resize {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_shrink_end_child(_arg0, _arg1)
}

// SetShrinkStartChild sets the `GtkPaned`:shrink-start-child property
func (p paned) SetShrinkStartChild(resize bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	if resize {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_shrink_start_child(_arg0, _arg1)
}

// SetStartChild sets the start child of @paned to @child.
func (p paned) SetStartChild(child Widget) {
	var _arg0 *C.GtkPaned  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_paned_set_start_child(_arg0, _arg1)
}

// SetWideHandle sets whether the separator should be wide.
func (p paned) SetWideHandle(wide bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	if wide {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_wide_handle(_arg0, _arg1)
}
