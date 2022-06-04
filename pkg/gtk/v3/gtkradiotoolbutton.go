// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkradiotoolbutton.go.
var GTypeRadioToolButton = coreglib.Type(C.gtk_radio_tool_button_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeRadioToolButton, F: marshalRadioToolButton},
	})
}

// RadioToolButtonOverrider contains methods that are overridable.
type RadioToolButtonOverrider interface {
}

// RadioToolButton is a ToolItem that contains a radio button, that is, a button
// that is part of a group of toggle buttons where only one button can be active
// at a time.
//
// Use gtk_radio_tool_button_new() to create a new GtkRadioToolButton. Use
// gtk_radio_tool_button_new_from_widget() to create a new GtkRadioToolButton
// that is part of the same group as an existing GtkRadioToolButton.
//
//
// CSS nodes
//
// GtkRadioToolButton has a single CSS node with name toolbutton.
type RadioToolButton struct {
	_ [0]func() // equal guard
	ToggleToolButton
}

var (
	_ coreglib.Objector = (*RadioToolButton)(nil)
	_ Binner            = (*RadioToolButton)(nil)
)

func classInitRadioToolButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRadioToolButton(obj *coreglib.Object) *RadioToolButton {
	return &RadioToolButton{
		ToggleToolButton: ToggleToolButton{
			ToolButton: ToolButton{
				ToolItem: ToolItem{
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
					Activatable: Activatable{
						Object: obj,
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
			},
		},
	}
}

func marshalRadioToolButton(p uintptr) (interface{}, error) {
	return wrapRadioToolButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewRadioToolButton creates a new RadioToolButton, adding it to group.
//
// The function takes the following parameters:
//
//    - group (optional): an existing radio button group, or NULL if you are
//      creating a new group.
//
// The function returns the following values:
//
//    - radioToolButton: new RadioToolButton.
//
func NewRadioToolButton(group []*RadioButton) *RadioToolButton {
	var _args [1]girepository.Argument

	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.void // out
			*(**C.void)(unsafe.Pointer(&dst)) = (*C.void)(unsafe.Pointer(coreglib.InternObject(src).Native()))
			*(**C.void)(unsafe.Pointer(&_args[0])) = C.g_slist_prepend(*(**C.void)(unsafe.Pointer(&_args[0])), C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_args[0])
	}

	_gret := girepository.MustFind("Gtk", "RadioToolButton").InvokeMethod("new_RadioToolButton", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(group)

	var _radioToolButton *RadioToolButton // out

	_radioToolButton = wrapRadioToolButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _radioToolButton
}

// NewRadioToolButtonFromStock creates a new RadioToolButton, adding it to
// group. The new RadioToolButton will contain an icon and label from the stock
// item indicated by stock_id.
//
// Deprecated: Use gtk_radio_tool_button_new() instead.
//
// The function takes the following parameters:
//
//    - group (optional): existing radio button group, or NULL if you are
//      creating a new group.
//    - stockId: name of a stock item.
//
// The function returns the following values:
//
//    - radioToolButton: new RadioToolButton.
//
func NewRadioToolButtonFromStock(group []*RadioButton, stockId string) *RadioToolButton {
	var _args [2]girepository.Argument

	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.void // out
			*(**C.void)(unsafe.Pointer(&dst)) = (*C.void)(unsafe.Pointer(coreglib.InternObject(src).Native()))
			*(**C.void)(unsafe.Pointer(&_args[0])) = C.g_slist_prepend(*(**C.void)(unsafe.Pointer(&_args[0])), C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_args[0])
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_args[1]))

	_gret := girepository.MustFind("Gtk", "RadioToolButton").InvokeMethod("new_RadioToolButton_from_stock", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(group)
	runtime.KeepAlive(stockId)

	var _radioToolButton *RadioToolButton // out

	_radioToolButton = wrapRadioToolButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _radioToolButton
}

// NewRadioToolButtonFromWidget creates a new RadioToolButton adding it to the
// same group as gruup.
//
// The function takes the following parameters:
//
//    - group (optional): existing RadioToolButton, or NULL.
//
// The function returns the following values:
//
//    - radioToolButton: new RadioToolButton.
//
func NewRadioToolButtonFromWidget(group *RadioToolButton) *RadioToolButton {
	var _args [1]girepository.Argument

	if group != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(group).Native()))
	}

	_gret := girepository.MustFind("Gtk", "RadioToolButton").InvokeMethod("new_RadioToolButton_from_widget", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(group)

	var _radioToolButton *RadioToolButton // out

	_radioToolButton = wrapRadioToolButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _radioToolButton
}

// NewRadioToolButtonWithStockFromWidget creates a new RadioToolButton adding it
// to the same group as group. The new RadioToolButton will contain an icon and
// label from the stock item indicated by stock_id.
//
// Deprecated: gtk_radio_tool_button_new_from_widget.
//
// The function takes the following parameters:
//
//    - group (optional): existing RadioToolButton.
//    - stockId: name of a stock item.
//
// The function returns the following values:
//
//    - radioToolButton: new RadioToolButton.
//
func NewRadioToolButtonWithStockFromWidget(group *RadioToolButton, stockId string) *RadioToolButton {
	var _args [2]girepository.Argument

	if group != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(group).Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_args[1]))

	_gret := girepository.MustFind("Gtk", "RadioToolButton").InvokeMethod("new_RadioToolButton_with_stock_from_widget", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(group)
	runtime.KeepAlive(stockId)

	var _radioToolButton *RadioToolButton // out

	_radioToolButton = wrapRadioToolButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _radioToolButton
}

// Group returns the radio button group button belongs to.
//
// The function returns the following values:
//
//    - sList: group button belongs to.
//
func (button *RadioToolButton) Group() []*RadioButton {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "RadioToolButton").InvokeMethod("get_group", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _sList []*RadioButton // out

	_sList = make([]*RadioButton, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst *RadioButton // out
		dst = wrapRadioButton(coreglib.Take(unsafe.Pointer(src)))
		_sList = append(_sList, dst)
	})

	return _sList
}

// SetGroup adds button to group, removing it from the group it belonged to
// before.
//
// The function takes the following parameters:
//
//    - group (optional): existing radio button group, or NULL.
//
func (button *RadioToolButton) SetGroup(group []*RadioButton) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.void // out
			*(**C.void)(unsafe.Pointer(&dst)) = (*C.void)(unsafe.Pointer(coreglib.InternObject(src).Native()))
			*(**C.void)(unsafe.Pointer(&_args[1])) = C.g_slist_prepend(*(**C.void)(unsafe.Pointer(&_args[1])), C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_args[1])
	}

	girepository.MustFind("Gtk", "RadioToolButton").InvokeMethod("set_group", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(group)
}
