// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gtk3_LinkButtonClass_activate_link(void*);
// extern gboolean _gotk4_gtk3_LinkButton_ConnectActivateLink(gpointer, guintptr);
import "C"

// glib.Type values for gtklinkbutton.go.
var GTypeLinkButton = coreglib.Type(C.gtk_link_button_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeLinkButton, F: marshalLinkButton},
	})
}

// LinkButtonOverrider contains methods that are overridable.
type LinkButtonOverrider interface {
	// The function returns the following values:
	//
	ActivateLink() bool
}

// LinkButton is a Button with a hyperlink, similar to the one used by web
// browsers, which triggers an action when clicked. It is useful to show quick
// links to resources.
//
// A link button is created by calling either gtk_link_button_new() or
// gtk_link_button_new_with_label(). If using the former, the URI you pass to
// the constructor is used as a label for the widget.
//
// The URI bound to a GtkLinkButton can be set specifically using
// gtk_link_button_set_uri(), and retrieved using gtk_link_button_get_uri().
//
// By default, GtkLinkButton calls gtk_show_uri_on_window() when the button is
// clicked. This behaviour can be overridden by connecting to the
// LinkButton::activate-link signal and returning TRUE from the signal handler.
//
//
// CSS nodes
//
// GtkLinkButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .link style class.
type LinkButton struct {
	_ [0]func() // equal guard
	Button
}

var (
	_ Binner            = (*LinkButton)(nil)
	_ coreglib.Objector = (*LinkButton)(nil)
)

func classInitLinkButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkLinkButtonClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkLinkButtonClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ ActivateLink() bool }); ok {
		pclass.activate_link = (*[0]byte)(C._gotk4_gtk3_LinkButtonClass_activate_link)
	}
}

//export _gotk4_gtk3_LinkButtonClass_activate_link
func _gotk4_gtk3_LinkButtonClass_activate_link(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ActivateLink() bool })

	ok := iface.ActivateLink()

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapLinkButton(obj *coreglib.Object) *LinkButton {
	return &LinkButton{
		Button: Button{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: coreglib.InitiallyUnowned{
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
			},
			Object: obj,
			Actionable: Actionable{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
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
			Activatable: Activatable{
				Object: obj,
			},
		},
	}
}

func marshalLinkButton(p uintptr) (interface{}, error) {
	return wrapLinkButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_LinkButton_ConnectActivateLink
func _gotk4_gtk3_LinkButton_ConnectActivateLink(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (ok bool))
	}

	ok := f()

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectActivateLink signal is emitted each time the LinkButton has been
// clicked.
//
// The default handler will call gtk_show_uri_on_window() with the URI stored
// inside the LinkButton:uri property.
//
// To override the default behavior, you can connect to the ::activate-link
// signal and stop the propagation of the signal by returning TRUE from your
// handler.
func (linkButton *LinkButton) ConnectActivateLink(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(linkButton, "activate-link", false, unsafe.Pointer(C._gotk4_gtk3_LinkButton_ConnectActivateLink), f)
}

// NewLinkButton creates a new LinkButton with the URI as its text.
//
// The function takes the following parameters:
//
//    - uri: valid URI.
//
// The function returns the following values:
//
//    - linkButton: new link button widget.
//
func NewLinkButton(uri string) *LinkButton {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg0))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "LinkButton").InvokeMethod("new_LinkButton", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(uri)

	var _linkButton *LinkButton // out

	_linkButton = wrapLinkButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _linkButton
}

// NewLinkButtonWithLabel creates a new LinkButton containing a label.
//
// The function takes the following parameters:
//
//    - uri: valid URI.
//    - label (optional): text of the button.
//
// The function returns the following values:
//
//    - linkButton: new link button widget.
//
func NewLinkButtonWithLabel(uri, label string) *LinkButton {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg0))
	if label != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "LinkButton").InvokeMethod("new_LinkButton_with_label", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(uri)
	runtime.KeepAlive(label)

	var _linkButton *LinkButton // out

	_linkButton = wrapLinkButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _linkButton
}

// URI retrieves the URI set using gtk_link_button_set_uri().
//
// The function returns the following values:
//
//    - utf8: valid URI. The returned string is owned by the link button and
//      should not be modified or freed.
//
func (linkButton *LinkButton) URI() string {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(linkButton).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "LinkButton").InvokeMethod("get_uri", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(linkButton)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Visited retrieves the “visited” state of the URI where the LinkButton points.
// The button becomes visited when it is clicked. If the URI is changed on the
// button, the “visited” state is unset again.
//
// The state may also be changed using gtk_link_button_set_visited().
//
// The function returns the following values:
//
//    - ok: TRUE if the link has been visited, FALSE otherwise.
//
func (linkButton *LinkButton) Visited() bool {
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(linkButton).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "LinkButton").InvokeMethod("get_visited", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(linkButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetURI sets uri as the URI where the LinkButton points. As a side-effect this
// unsets the “visited” state of the button.
//
// The function takes the following parameters:
//
//    - uri: valid URI.
//
func (linkButton *LinkButton) SetURI(uri string) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(linkButton).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "LinkButton").InvokeMethod("set_uri", _args[:], nil)

	runtime.KeepAlive(linkButton)
	runtime.KeepAlive(uri)
}

// SetVisited sets the “visited” state of the URI where the LinkButton points.
// See gtk_link_button_get_visited() for more details.
//
// The function takes the following parameters:
//
//    - visited: new “visited” state.
//
func (linkButton *LinkButton) SetVisited(visited bool) {
	var _args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(linkButton).Native()))
	if visited {
		_arg1 = C.TRUE
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gboolean)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "LinkButton").InvokeMethod("set_visited", _args[:], nil)

	runtime.KeepAlive(linkButton)
	runtime.KeepAlive(visited)
}
