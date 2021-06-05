// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_text_child_anchor_get_type()), F: marshalTextChildAnchor},
	})
}

// TextChildAnchor: a TextChildAnchor is a spot in the buffer where child
// widgets can be “anchored” (inserted inline, as if they were characters). The
// anchor can have multiple widgets anchored, to allow for multiple views.
type TextChildAnchor interface {
	gextras.Objector

	// Deleted determines whether a child anchor has been deleted from the
	// buffer. Keep in mind that the child anchor will be unreferenced when
	// removed from the buffer, so you need to hold your own reference (with
	// g_object_ref()) if you plan to use this function — otherwise all deleted
	// child anchors will also be finalized.
	Deleted() bool
	// Widgets gets a list of all widgets anchored at this child anchor. The
	// returned list should be freed with g_list_free().
	Widgets() *glib.List
}

// textChildAnchor implements the TextChildAnchor interface.
type textChildAnchor struct {
	gextras.Objector
}

var _ TextChildAnchor = (*textChildAnchor)(nil)

// WrapTextChildAnchor wraps a GObject to the right type. It is
// primarily used internally.
func WrapTextChildAnchor(obj *externglib.Object) TextChildAnchor {
	return TextChildAnchor{
		Objector: obj,
	}
}

func marshalTextChildAnchor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTextChildAnchor(obj), nil
}

// NewTextChildAnchor constructs a class TextChildAnchor.
func NewTextChildAnchor() TextChildAnchor {
	var cret C.GtkTextChildAnchor
	var ret1 TextChildAnchor

	cret = C.gtk_text_child_anchor_new()

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(TextChildAnchor)

	return ret1
}

// Deleted determines whether a child anchor has been deleted from the
// buffer. Keep in mind that the child anchor will be unreferenced when
// removed from the buffer, so you need to hold your own reference (with
// g_object_ref()) if you plan to use this function — otherwise all deleted
// child anchors will also be finalized.
func (a textChildAnchor) Deleted() bool {
	var arg0 *C.GtkTextChildAnchor

	arg0 = (*C.GtkTextChildAnchor)(unsafe.Pointer(a.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_text_child_anchor_get_deleted(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Widgets gets a list of all widgets anchored at this child anchor. The
// returned list should be freed with g_list_free().
func (a textChildAnchor) Widgets() *glib.List {
	var arg0 *C.GtkTextChildAnchor

	arg0 = (*C.GtkTextChildAnchor)(unsafe.Pointer(a.Native()))

	var cret *C.GList
	var ret1 *glib.List

	cret = C.gtk_text_child_anchor_get_widgets(arg0)

	ret1 = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}
