// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_im_multicontext_get_type()), F: marshalIMMulticontext},
	})
}

type IMMulticontextPrivate struct {
	native C.GtkIMMulticontextPrivate
}

// WrapIMMulticontextPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapIMMulticontextPrivate(ptr unsafe.Pointer) *IMMulticontextPrivate {
	if ptr == nil {
		return nil
	}

	return (*IMMulticontextPrivate)(ptr)
}

func marshalIMMulticontextPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapIMMulticontextPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (i *IMMulticontextPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}

// IMMulticontext: `GtkIMMulticontext` is input method supporting multiple,
// switchable input methods.
//
// Text widgets such as `GtkText` or `GtkTextView` use a `GtkIMMultiContext` to
// implement their `im-module` property for switching between different input
// methods.
type IMMulticontext interface {
	IMContext

	// ContextID gets the id of the currently active delegate of the @context.
	ContextID() string
	// SetContextID sets the context id for @context.
	//
	// This causes the currently active delegate of @context to be replaced by
	// the delegate corresponding to the new context id.
	SetContextID(contextID string)
}

// imMulticontext implements the IMMulticontext interface.
type imMulticontext struct {
	IMContext
}

var _ IMMulticontext = (*imMulticontext)(nil)

// WrapIMMulticontext wraps a GObject to the right type. It is
// primarily used internally.
func WrapIMMulticontext(obj *externglib.Object) IMMulticontext {
	return IMMulticontext{
		IMContext: WrapIMContext(obj),
	}
}

func marshalIMMulticontext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapIMMulticontext(obj), nil
}

// NewIMMulticontext constructs a class IMMulticontext.
func NewIMMulticontext() IMMulticontext {
	ret := C.gtk_im_multicontext_new()

	var ret0 IMMulticontext

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(IMMulticontext)

	return ret0
}

// ContextID gets the id of the currently active delegate of the @context.
func (c imMulticontext) ContextID() string {
	var arg0 *C.GtkIMMulticontext

	arg0 = (*C.GtkIMMulticontext)(c.Native())

	ret := C.gtk_im_multicontext_get_context_id(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// SetContextID sets the context id for @context.
//
// This causes the currently active delegate of @context to be replaced by
// the delegate corresponding to the new context id.
func (c imMulticontext) SetContextID(contextID string) {
	var arg0 *C.GtkIMMulticontext
	var arg1 *C.char

	arg0 = (*C.GtkIMMulticontext)(c.Native())
	arg1 = (*C.gchar)(C.CString(contextID))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_im_multicontext_set_context_id(arg0, arg1)
}
