// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_statusbar_get_type()), F: marshalStatusbarrer},
	})
}

// StatusbarOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type StatusbarOverrider interface {
	// The function takes the following parameters:
	//
	//    - contextId
	//    - text
	//
	TextPopped(contextId uint, text string)
	// The function takes the following parameters:
	//
	//    - contextId
	//    - text
	//
	TextPushed(contextId uint, text string)
}

// Statusbar is usually placed along the bottom of an application's main Window.
// It may provide a regular commentary of the application's status (as is
// usually the case in a web browser, for example), or may be used to simply
// output a message when the status changes, (when an upload is complete in an
// FTP client, for example).
//
// Status bars in GTK+ maintain a stack of messages. The message at the top of
// the each bar’s stack is the one that will currently be displayed.
//
// Any messages added to a statusbar’s stack must specify a context id that is
// used to uniquely identify the source of a message. This context id can be
// generated by gtk_statusbar_get_context_id(), given a message and the
// statusbar that it will be added to. Note that messages are stored in a stack,
// and when choosing which message to display, the stack structure is adhered
// to, regardless of the context identifier of a message.
//
// One could say that a statusbar maintains one stack of messages for display
// purposes, but allows multiple message producers to maintain sub-stacks of the
// messages they produced (via context ids).
//
// Status bars are created using gtk_statusbar_new().
//
// Messages are added to the bar’s stack with gtk_statusbar_push().
//
// The message at the top of the stack can be removed using gtk_statusbar_pop().
// A message can be removed from anywhere in the stack if its message id was
// recorded at the time it was added. This is done using gtk_statusbar_remove().
//
//
// CSS node
//
// GtkStatusbar has a single CSS node with name statusbar.
type Statusbar struct {
	Box

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ Containerer         = (*Statusbar)(nil)
	_ externglib.Objector = (*Statusbar)(nil)
)

func wrapStatusbar(obj *externglib.Object) *Statusbar {
	return &Statusbar{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalStatusbarrer(p uintptr) (interface{}, error) {
	return wrapStatusbar(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectTextPopped is emitted whenever a new message is popped off a
// statusbar's stack.
func (statusbar *Statusbar) ConnectTextPopped(f func(contextId uint, text string)) externglib.SignalHandle {
	return statusbar.Connect("text-popped", f)
}

// ConnectTextPushed is emitted whenever a new message gets pushed onto a
// statusbar's stack.
func (statusbar *Statusbar) ConnectTextPushed(f func(contextId uint, text string)) externglib.SignalHandle {
	return statusbar.Connect("text-pushed", f)
}

// NewStatusbar creates a new Statusbar ready for messages.
//
// The function returns the following values:
//
//    - statusbar: new Statusbar.
//
func NewStatusbar() *Statusbar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_statusbar_new()

	var _statusbar *Statusbar // out

	_statusbar = wrapStatusbar(externglib.Take(unsafe.Pointer(_cret)))

	return _statusbar
}

// ContextID returns a new context identifier, given a description of the actual
// context. Note that the description is not shown in the UI.
//
// The function takes the following parameters:
//
//    - contextDescription: textual description of what context the new message
//      is being used in.
//
// The function returns the following values:
//
//    - guint: integer id.
//
func (statusbar *Statusbar) ContextID(contextDescription string) uint {
	var _arg0 *C.GtkStatusbar // out
	var _arg1 *C.gchar        // out
	var _cret C.guint         // in

	_arg0 = (*C.GtkStatusbar)(unsafe.Pointer(statusbar.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(contextDescription)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_statusbar_get_context_id(_arg0, _arg1)
	runtime.KeepAlive(statusbar)
	runtime.KeepAlive(contextDescription)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// MessageArea retrieves the box containing the label widget.
//
// The function returns the following values:
//
//    - box: Box.
//
func (statusbar *Statusbar) MessageArea() *Box {
	var _arg0 *C.GtkStatusbar // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkStatusbar)(unsafe.Pointer(statusbar.Native()))

	_cret = C.gtk_statusbar_get_message_area(_arg0)
	runtime.KeepAlive(statusbar)

	var _box *Box // out

	_box = wrapBox(externglib.Take(unsafe.Pointer(_cret)))

	return _box
}

// Pop removes the first message in the Statusbar’s stack with the given context
// id.
//
// Note that this may not change the displayed message, if the message at the
// top of the stack has a different context id.
//
// The function takes the following parameters:
//
//    - contextId: context identifier.
//
func (statusbar *Statusbar) Pop(contextId uint) {
	var _arg0 *C.GtkStatusbar // out
	var _arg1 C.guint         // out

	_arg0 = (*C.GtkStatusbar)(unsafe.Pointer(statusbar.Native()))
	_arg1 = C.guint(contextId)

	C.gtk_statusbar_pop(_arg0, _arg1)
	runtime.KeepAlive(statusbar)
	runtime.KeepAlive(contextId)
}

// Push pushes a new message onto a statusbar’s stack.
//
// The function takes the following parameters:
//
//    - contextId message’s context id, as returned by
//      gtk_statusbar_get_context_id().
//    - text: message to add to the statusbar.
//
// The function returns the following values:
//
//    - guint: message id that can be used with gtk_statusbar_remove().
//
func (statusbar *Statusbar) Push(contextId uint, text string) uint {
	var _arg0 *C.GtkStatusbar // out
	var _arg1 C.guint         // out
	var _arg2 *C.gchar        // out
	var _cret C.guint         // in

	_arg0 = (*C.GtkStatusbar)(unsafe.Pointer(statusbar.Native()))
	_arg1 = C.guint(contextId)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_statusbar_push(_arg0, _arg1, _arg2)
	runtime.KeepAlive(statusbar)
	runtime.KeepAlive(contextId)
	runtime.KeepAlive(text)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Remove forces the removal of a message from a statusbar’s stack. The exact
// context_id and message_id must be specified.
//
// The function takes the following parameters:
//
//    - contextId: context identifier.
//    - messageId: message identifier, as returned by gtk_statusbar_push().
//
func (statusbar *Statusbar) Remove(contextId, messageId uint) {
	var _arg0 *C.GtkStatusbar // out
	var _arg1 C.guint         // out
	var _arg2 C.guint         // out

	_arg0 = (*C.GtkStatusbar)(unsafe.Pointer(statusbar.Native()))
	_arg1 = C.guint(contextId)
	_arg2 = C.guint(messageId)

	C.gtk_statusbar_remove(_arg0, _arg1, _arg2)
	runtime.KeepAlive(statusbar)
	runtime.KeepAlive(contextId)
	runtime.KeepAlive(messageId)
}

// RemoveAll forces the removal of all messages from a statusbar's stack with
// the exact context_id.
//
// The function takes the following parameters:
//
//    - contextId: context identifier.
//
func (statusbar *Statusbar) RemoveAll(contextId uint) {
	var _arg0 *C.GtkStatusbar // out
	var _arg1 C.guint         // out

	_arg0 = (*C.GtkStatusbar)(unsafe.Pointer(statusbar.Native()))
	_arg1 = C.guint(contextId)

	C.gtk_statusbar_remove_all(_arg0, _arg1)
	runtime.KeepAlive(statusbar)
	runtime.KeepAlive(contextId)
}
