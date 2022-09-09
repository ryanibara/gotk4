// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeTextChildAnchor = coreglib.Type(C.gtk_text_child_anchor_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTextChildAnchor, F: marshalTextChildAnchor},
	})
}

// TextChildAnchorOverrides contains methods that are overridable.
type TextChildAnchorOverrides struct {
}

func defaultTextChildAnchorOverrides(v *TextChildAnchor) TextChildAnchorOverrides {
	return TextChildAnchorOverrides{}
}

// TextChildAnchor: GtkTextChildAnchor is a spot in a GtkTextBuffer where child
// widgets can be “anchored”.
//
// The anchor can have multiple widgets anchored, to allow for multiple views.
type TextChildAnchor struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TextChildAnchor)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*TextChildAnchor, *TextChildAnchorClass, TextChildAnchorOverrides](
		GTypeTextChildAnchor,
		initTextChildAnchorClass,
		wrapTextChildAnchor,
		defaultTextChildAnchorOverrides,
	)
}

func initTextChildAnchorClass(gclass unsafe.Pointer, overrides TextChildAnchorOverrides, classInitFunc func(*TextChildAnchorClass)) {
	if classInitFunc != nil {
		class := (*TextChildAnchorClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapTextChildAnchor(obj *coreglib.Object) *TextChildAnchor {
	return &TextChildAnchor{
		Object: obj,
	}
}

func marshalTextChildAnchor(p uintptr) (interface{}, error) {
	return wrapTextChildAnchor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTextChildAnchor creates a new GtkTextChildAnchor.
//
// Usually you would then insert it into a GtkTextBuffer with
// gtk.TextBuffer.InsertChildAnchor(). To perform the creation and insertion in
// one step, use the convenience function gtk.TextBuffer.CreateChildAnchor().
//
// The function returns the following values:
//
//    - textChildAnchor: new GtkTextChildAnchor.
//
func NewTextChildAnchor() *TextChildAnchor {
	var _cret *C.GtkTextChildAnchor // in

	_cret = C.gtk_text_child_anchor_new()

	var _textChildAnchor *TextChildAnchor // out

	_textChildAnchor = wrapTextChildAnchor(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _textChildAnchor
}

// Deleted determines whether a child anchor has been deleted from the buffer.
//
// Keep in mind that the child anchor will be unreferenced when removed from the
// buffer, so you need to hold your own reference (with g_object_ref()) if you
// plan to use this function — otherwise all deleted child anchors will also be
// finalized.
//
// The function returns the following values:
//
//    - ok: TRUE if the child anchor has been deleted from its buffer.
//
func (anchor *TextChildAnchor) Deleted() bool {
	var _arg0 *C.GtkTextChildAnchor // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkTextChildAnchor)(unsafe.Pointer(coreglib.InternObject(anchor).Native()))

	_cret = C.gtk_text_child_anchor_get_deleted(_arg0)
	runtime.KeepAlive(anchor)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Widgets gets a list of all widgets anchored at this child anchor.
//
// The order in which the widgets are returned is not defined.
//
// The function returns the following values:
//
//    - widgets: an array of widgets anchored at anchor.
//
func (anchor *TextChildAnchor) Widgets() []Widgetter {
	var _arg0 *C.GtkTextChildAnchor // out
	var _cret **C.GtkWidget         // in
	var _arg1 C.guint               // in

	_arg0 = (*C.GtkTextChildAnchor)(unsafe.Pointer(coreglib.InternObject(anchor).Native()))

	_cret = C.gtk_text_child_anchor_get_widgets(_arg0, &_arg1)
	runtime.KeepAlive(anchor)

	var _widgets []Widgetter // out

	defer C.free(unsafe.Pointer(_cret))
	{
		src := unsafe.Slice((**C.GtkWidget)(_cret), _arg1)
		_widgets = make([]Widgetter, _arg1)
		for i := 0; i < int(_arg1); i++ {
			{
				objptr := unsafe.Pointer(src[i])
				if objptr == nil {
					panic("object of type gtk.Widgetter is nil")
				}

				object := coreglib.Take(objptr)
				casted := object.WalkCast(func(obj coreglib.Objector) bool {
					_, ok := obj.(Widgetter)
					return ok
				})
				rv, ok := casted.(Widgetter)
				if !ok {
					panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
				}
				_widgets[i] = rv
			}
		}
	}

	return _widgets
}

// TextChildAnchorClass: instance of this type is always passed by reference.
type TextChildAnchorClass struct {
	*textChildAnchorClass
}

// textChildAnchorClass is the struct that's finalized.
type textChildAnchorClass struct {
	native *C.GtkTextChildAnchorClass
}
