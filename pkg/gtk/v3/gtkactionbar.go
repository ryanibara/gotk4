// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkactionbar.go.
var GTypeActionBar = coreglib.Type(C.gtk_action_bar_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeActionBar, F: marshalActionBar},
	})
}

// ActionBarOverrider contains methods that are overridable.
type ActionBarOverrider interface {
}

// ActionBar is designed to present contextual actions. It is expected to be
// displayed below the content and expand horizontally to fill the area.
//
// It allows placing children at the start or the end. In addition, it contains
// an internal centered box which is centered with respect to the full width of
// the box, even if the children at either side take up different amounts of
// space.
//
//
// CSS nodes
//
// GtkActionBar has a single CSS node with name actionbar.
type ActionBar struct {
	_ [0]func() // equal guard
	Bin
}

var (
	_ Binner = (*ActionBar)(nil)
)

func classInitActionBarrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapActionBar(obj *coreglib.Object) *ActionBar {
	return &ActionBar{
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
	}
}

func marshalActionBar(p uintptr) (interface{}, error) {
	return wrapActionBar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewActionBar creates a new ActionBar widget.
//
// The function returns the following values:
//
//    - actionBar: new ActionBar.
//
func NewActionBar() *ActionBar {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "ActionBar").InvokeMethod("new_ActionBar", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _actionBar *ActionBar // out

	_actionBar = wrapActionBar(coreglib.Take(unsafe.Pointer(_cret)))

	return _actionBar
}

// CenterWidget retrieves the center bar widget of the bar.
//
// The function returns the following values:
//
//    - widget (optional): center Widget or NULL.
//
func (actionBar *ActionBar) CenterWidget() Widgetter {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionBar).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ActionBar").InvokeMethod("get_center_widget", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(actionBar)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// PackEnd adds child to action_bar, packed with reference to the end of the
// action_bar.
//
// The function takes the following parameters:
//
//    - child to be added to action_bar.
//
func (actionBar *ActionBar) PackEnd(child Widgetter) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionBar).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "ActionBar").InvokeMethod("pack_end", _args[:], nil)

	runtime.KeepAlive(actionBar)
	runtime.KeepAlive(child)
}

// PackStart adds child to action_bar, packed with reference to the start of the
// action_bar.
//
// The function takes the following parameters:
//
//    - child to be added to action_bar.
//
func (actionBar *ActionBar) PackStart(child Widgetter) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionBar).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "ActionBar").InvokeMethod("pack_start", _args[:], nil)

	runtime.KeepAlive(actionBar)
	runtime.KeepAlive(child)
}

// SetCenterWidget sets the center widget for the ActionBar.
//
// The function takes the following parameters:
//
//    - centerWidget (optional): widget to use for the center.
//
func (actionBar *ActionBar) SetCenterWidget(centerWidget Widgetter) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(actionBar).Native()))
	if centerWidget != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(centerWidget).Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "ActionBar").InvokeMethod("set_center_widget", _args[:], nil)

	runtime.KeepAlive(actionBar)
	runtime.KeepAlive(centerWidget)
}
