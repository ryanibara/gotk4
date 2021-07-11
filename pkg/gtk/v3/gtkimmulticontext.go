// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_im_multicontext_get_type()), F: marshalIMMulticontexter},
	})
}

// IMMulticontexter describes IMMulticontext's methods.
type IMMulticontexter interface {
	// AppendMenuitems: add menuitems for various available input methods to a
	// menu; the menuitems, when selected, will switch the input method for the
	// context and the global default input method.
	AppendMenuitems(menushell MenuSheller)
	// ContextID gets the id of the currently active slave of the @context.
	ContextID() string
	// SetContextID sets the context id for @context.
	SetContextID(contextId string)
}

//
type IMMulticontext struct {
	IMContext
}

var (
	_ IMMulticontexter = (*IMMulticontext)(nil)
	_ gextras.Nativer  = (*IMMulticontext)(nil)
)

func wrapIMMulticontext(obj *externglib.Object) IMMulticontexter {
	return &IMMulticontext{
		IMContext: IMContext{
			Object: obj,
		},
	}
}

func marshalIMMulticontexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapIMMulticontext(obj), nil
}

// NewIMMulticontext creates a new IMMulticontext.
func NewIMMulticontext() *IMMulticontext {
	var _cret *C.GtkIMContext // in

	_cret = C.gtk_im_multicontext_new()

	var _imMulticontext *IMMulticontext // out

	_imMulticontext = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*IMMulticontext)

	return _imMulticontext
}

// AppendMenuitems: add menuitems for various available input methods to a menu;
// the menuitems, when selected, will switch the input method for the context
// and the global default input method.
//
// Deprecated: It is better to use the system-wide input method framework for
// changing input methods. Modern desktop shells offer on-screen displays for
// this that can triggered with a keyboard shortcut, e.g. Super-Space.
func (context *IMMulticontext) AppendMenuitems(menushell MenuSheller) {
	var _arg0 *C.GtkIMMulticontext // out
	var _arg1 *C.GtkMenuShell      // out

	_arg0 = (*C.GtkIMMulticontext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.GtkMenuShell)(unsafe.Pointer((menushell).(gextras.Nativer).Native()))

	C.gtk_im_multicontext_append_menuitems(_arg0, _arg1)
}

// ContextID gets the id of the currently active slave of the @context.
func (context *IMMulticontext) ContextID() string {
	var _arg0 *C.GtkIMMulticontext // out
	var _cret *C.char              // in

	_arg0 = (*C.GtkIMMulticontext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_im_multicontext_get_context_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetContextID sets the context id for @context.
//
// This causes the currently active slave of @context to be replaced by the
// slave corresponding to the new context id.
func (context *IMMulticontext) SetContextID(contextId string) {
	var _arg0 *C.GtkIMMulticontext // out
	var _arg1 *C.char              // out

	_arg0 = (*C.GtkIMMulticontext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(contextId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_im_multicontext_set_context_id(_arg0, _arg1)
}
