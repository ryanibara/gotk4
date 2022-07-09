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
// #include <glib-object.h>
// extern void _gotk4_gtk3_ToolButtonClass_clicked(void*);
// extern void _gotk4_gtk3_ToolButton_ConnectClicked(gpointer, guintptr);
import "C"

// GTypeToolButton returns the GType for the type ToolButton.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeToolButton() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ToolButton").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalToolButton)
	return gtype
}

// ToolButtonOverrider contains methods that are overridable.
type ToolButtonOverrider interface {
	Clicked()
}

// ToolButton are ToolItems containing buttons.
//
// Use gtk_tool_button_new() to create a new ToolButton.
//
// The label of a ToolButton is determined by the properties
// ToolButton:label-widget, ToolButton:label, and ToolButton:stock-id. If
// ToolButton:label-widget is non-NULL, then that widget is used as the label.
// Otherwise, if ToolButton:label is non-NULL, that string is used as the label.
// Otherwise, if ToolButton:stock-id is non-NULL, the label is determined by the
// stock item. Otherwise, the button does not have a label.
//
// The icon of a ToolButton is determined by the properties
// ToolButton:icon-widget and ToolButton:stock-id. If ToolButton:icon-widget is
// non-NULL, then that widget is used as the icon. Otherwise, if
// ToolButton:stock-id is non-NULL, the icon is determined by the stock item.
// Otherwise, the button does not have a icon.
//
//
// CSS nodes
//
// GtkToolButton has a single CSS node with name toolbutton.
type ToolButton struct {
	_ [0]func() // equal guard
	ToolItem

	*coreglib.Object
	Actionable
}

var (
	_ coreglib.Objector = (*ToolButton)(nil)
	_ Binner            = (*ToolButton)(nil)
)

func classInitToolButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "ToolButtonClass")

	if _, ok := goval.(interface{ Clicked() }); ok {
		o := pclass.StructFieldOffset("clicked")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ToolButtonClass_clicked)
	}
}

//export _gotk4_gtk3_ToolButtonClass_clicked
func _gotk4_gtk3_ToolButtonClass_clicked(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Clicked() })

	iface.Clicked()
}

func wrapToolButton(obj *coreglib.Object) *ToolButton {
	return &ToolButton{
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
	}
}

func marshalToolButton(p uintptr) (interface{}, error) {
	return wrapToolButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_ToolButton_ConnectClicked
func _gotk4_gtk3_ToolButton_ConnectClicked(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectClicked: this signal is emitted when the tool button is clicked with
// the mouse or activated with the keyboard.
func (button *ToolButton) ConnectClicked(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "clicked", false, unsafe.Pointer(C._gotk4_gtk3_ToolButton_ConnectClicked), f)
}

// NewToolButton creates a new ToolButton using icon_widget as contents and
// label as label.
//
// The function takes the following parameters:
//
//    - iconWidget (optional): widget that will be used as the button contents,
//      or NULL.
//    - label (optional): string that will be used as label, or NULL.
//
// The function returns the following values:
//
//    - toolButton: new ToolButton.
//
func NewToolButton(iconWidget Widgetter, label string) *ToolButton {
	var _args [2]girepository.Argument

	if iconWidget != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(iconWidget).Native()))
	}
	if label != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	_gret := girepository.MustFind("Gtk", "ToolButton").InvokeMethod("new_ToolButton", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(iconWidget)
	runtime.KeepAlive(label)

	var _toolButton *ToolButton // out

	_toolButton = wrapToolButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _toolButton
}

// NewToolButtonFromStock creates a new ToolButton containing the image and text
// from a stock item. Some stock ids have preprocessor macros like K_STOCK_OK
// and K_STOCK_APPLY.
//
// It is an error if stock_id is not a name of a stock item.
//
// Deprecated: Use gtk_tool_button_new() together with
// gtk_image_new_from_icon_name() instead.
//
// The function takes the following parameters:
//
//    - stockId: name of the stock item.
//
// The function returns the following values:
//
//    - toolButton: new ToolButton.
//
func NewToolButtonFromStock(stockId string) *ToolButton {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_args[0]))

	_gret := girepository.MustFind("Gtk", "ToolButton").InvokeMethod("new_ToolButton_from_stock", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stockId)

	var _toolButton *ToolButton // out

	_toolButton = wrapToolButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _toolButton
}

// IconName returns the name of the themed icon for the tool button, see
// gtk_tool_button_set_icon_name().
//
// The function returns the following values:
//
//    - utf8 (optional): icon name or NULL if the tool button has no themed icon.
//
func (button *ToolButton) IconName() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "ToolButton").InvokeMethod("get_icon_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// IconWidget: return the widget used as icon widget on button. See
// gtk_tool_button_set_icon_widget().
//
// The function returns the following values:
//
//    - widget (optional) used as icon on button, or NULL.
//
func (button *ToolButton) IconWidget() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "ToolButton").InvokeMethod("get_icon_widget", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
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

// Label returns the label used by the tool button, or NULL if the tool button
// doesn’t have a label. or uses a the label from a stock item. The returned
// string is owned by GTK+, and must not be modified or freed.
//
// The function returns the following values:
//
//    - utf8 (optional): label, or NULL.
//
func (button *ToolButton) Label() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "ToolButton").InvokeMethod("get_label", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// LabelWidget returns the widget used as label on button. See
// gtk_tool_button_set_label_widget().
//
// The function returns the following values:
//
//    - widget (optional) used as label on button, or NULL.
//
func (button *ToolButton) LabelWidget() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "ToolButton").InvokeMethod("get_label_widget", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
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

// StockID returns the name of the stock item. See
// gtk_tool_button_set_stock_id(). The returned string is owned by GTK+ and must
// not be freed or modifed.
//
// Deprecated: Use gtk_tool_button_get_icon_name() instead.
//
// The function returns the following values:
//
//    - utf8: name of the stock item for button.
//
func (button *ToolButton) StockID() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "ToolButton").InvokeMethod("get_stock_id", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UseUnderline returns whether underscores in the label property are used as
// mnemonics on menu items on the overflow menu. See
// gtk_tool_button_set_use_underline().
//
// The function returns the following values:
//
//    - ok: TRUE if underscores in the label property are used as mnemonics on
//      menu items on the overflow menu.
//
func (button *ToolButton) UseUnderline() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_gret := girepository.MustFind("Gtk", "ToolButton").InvokeMethod("get_use_underline", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetIconName sets the icon for the tool button from a named themed icon. See
// the docs for IconTheme for more details. The ToolButton:icon-name property
// only has an effect if not overridden by non-NULL ToolButton:label-widget,
// ToolButton:icon-widget and ToolButton:stock-id properties.
//
// The function takes the following parameters:
//
//    - iconName (optional): name of the themed icon.
//
func (button *ToolButton) SetIconName(iconName string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if iconName != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "ToolButton").InvokeMethod("set_icon_name", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(iconName)
}

// SetIconWidget sets icon as the widget used as icon on button. If icon_widget
// is NULL the icon is determined by the ToolButton:stock-id property. If the
// ToolButton:stock-id property is also NULL, button will not have an icon.
//
// The function takes the following parameters:
//
//    - iconWidget (optional): widget used as icon, or NULL.
//
func (button *ToolButton) SetIconWidget(iconWidget Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if iconWidget != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(iconWidget).Native()))
	}

	girepository.MustFind("Gtk", "ToolButton").InvokeMethod("set_icon_widget", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(iconWidget)
}

// SetLabel sets label as the label used for the tool button. The
// ToolButton:label property only has an effect if not overridden by a non-NULL
// ToolButton:label-widget property. If both the ToolButton:label-widget and
// ToolButton:label properties are NULL, the label is determined by the
// ToolButton:stock-id property. If the ToolButton:stock-id property is also
// NULL, button will not have a label.
//
// The function takes the following parameters:
//
//    - label (optional): string that will be used as label, or NULL.
//
func (button *ToolButton) SetLabel(label string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if label != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "ToolButton").InvokeMethod("set_label", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(label)
}

// SetLabelWidget sets label_widget as the widget that will be used as the label
// for button. If label_widget is NULL the ToolButton:label property is used as
// label. If ToolButton:label is also NULL, the label in the stock item
// determined by the ToolButton:stock-id property is used as label. If
// ToolButton:stock-id is also NULL, button does not have a label.
//
// The function takes the following parameters:
//
//    - labelWidget (optional): widget used as label, or NULL.
//
func (button *ToolButton) SetLabelWidget(labelWidget Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if labelWidget != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(labelWidget).Native()))
	}

	girepository.MustFind("Gtk", "ToolButton").InvokeMethod("set_label_widget", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(labelWidget)
}

// SetStockID sets the name of the stock item. See
// gtk_tool_button_new_from_stock(). The stock_id property only has an effect if
// not overridden by non-NULL ToolButton:label-widget and ToolButton:icon-widget
// properties.
//
// Deprecated: Use gtk_tool_button_set_icon_name() instead.
//
// The function takes the following parameters:
//
//    - stockId (optional): name of a stock item, or NULL.
//
func (button *ToolButton) SetStockID(stockId string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if stockId != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(stockId)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "ToolButton").InvokeMethod("set_stock_id", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(stockId)
}

// SetUseUnderline: if set, an underline in the label property indicates that
// the next character should be used for the mnemonic accelerator key in the
// overflow menu. For example, if the label property is “_Open” and
// use_underline is TRUE, the label on the tool button will be “Open” and the
// item on the overflow menu will have an underlined “O”.
//
// Labels shown on tool buttons never have mnemonics on them; this property only
// affects the menu item on the overflow menu.
//
// The function takes the following parameters:
//
//    - useUnderline: whether the button label has the form “_Open”.
//
func (button *ToolButton) SetUseUnderline(useUnderline bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if useUnderline {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "ToolButton").InvokeMethod("set_use_underline", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(useUnderline)
}
