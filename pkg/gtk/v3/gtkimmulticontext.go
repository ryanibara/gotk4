// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtkimmulticontext.go.
var GTypeIMMulticontext = externglib.Type(C.gtk_im_multicontext_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeIMMulticontext, F: marshalIMMulticontext},
	})
}

// IMMulticontextOverrider contains methods that are overridable.
type IMMulticontextOverrider interface {
}

type IMMulticontext struct {
	_ [0]func() // equal guard
	IMContext
}

var (
	_ IMContexter = (*IMMulticontext)(nil)
)

func classInitIMMulticontexter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapIMMulticontext(obj *externglib.Object) *IMMulticontext {
	return &IMMulticontext{
		IMContext: IMContext{
			Object: obj,
		},
	}
}

func marshalIMMulticontext(p uintptr) (interface{}, error) {
	return wrapIMMulticontext(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewIMMulticontext creates a new IMMulticontext.
//
// The function returns the following values:
//
//    - imMulticontext: new IMMulticontext.
//
func NewIMMulticontext() *IMMulticontext {
	var _cret *C.GtkIMContext // in

	_cret = C.gtk_im_multicontext_new()

	var _imMulticontext *IMMulticontext // out

	_imMulticontext = wrapIMMulticontext(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _imMulticontext
}

// AppendMenuitems: add menuitems for various available input methods to a menu;
// the menuitems, when selected, will switch the input method for the context
// and the global default input method.
//
// Deprecated: It is better to use the system-wide input method framework for
// changing input methods. Modern desktop shells offer on-screen displays for
// this that can triggered with a keyboard shortcut, e.g. Super-Space.
//
// The function takes the following parameters:
//
//    - menushell: MenuShell.
//
func (context *IMMulticontext) AppendMenuitems(menushell MenuSheller) {
	var _arg0 *C.GtkIMMulticontext // out
	var _arg1 *C.GtkMenuShell      // out

	_arg0 = (*C.GtkIMMulticontext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.GtkMenuShell)(unsafe.Pointer(menushell.Native()))

	C.gtk_im_multicontext_append_menuitems(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(menushell)
}

// ContextID gets the id of the currently active slave of the context.
//
// The function returns the following values:
//
//    - utf8: id of the currently active slave.
//
func (context *IMMulticontext) ContextID() string {
	var _arg0 *C.GtkIMMulticontext // out
	var _cret *C.char              // in

	_arg0 = (*C.GtkIMMulticontext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_im_multicontext_get_context_id(_arg0)
	runtime.KeepAlive(context)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetContextID sets the context id for context.
//
// This causes the currently active slave of context to be replaced by the slave
// corresponding to the new context id.
//
// The function takes the following parameters:
//
//    - contextId: id to use.
//
func (context *IMMulticontext) SetContextID(contextId string) {
	var _arg0 *C.GtkIMMulticontext // out
	var _arg1 *C.char              // out

	_arg0 = (*C.GtkIMMulticontext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(contextId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_im_multicontext_set_context_id(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(contextId)
}
