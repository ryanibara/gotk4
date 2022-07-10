// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk3_ButtonClass_activate(void*);
// extern void _gotk4_gtk3_ButtonClass_clicked(void*);
// extern void _gotk4_gtk3_ButtonClass_enter(void*);
// extern void _gotk4_gtk3_ButtonClass_leave(void*);
// extern void _gotk4_gtk3_ButtonClass_pressed(void*);
// extern void _gotk4_gtk3_ButtonClass_released(void*);
// extern void _gotk4_gtk3_Button_ConnectActivate(gpointer, guintptr);
// extern void _gotk4_gtk3_Button_ConnectClicked(gpointer, guintptr);
// extern void _gotk4_gtk3_Button_ConnectEnter(gpointer, guintptr);
// extern void _gotk4_gtk3_Button_ConnectLeave(gpointer, guintptr);
// extern void _gotk4_gtk3_Button_ConnectPressed(gpointer, guintptr);
// extern void _gotk4_gtk3_Button_ConnectReleased(gpointer, guintptr);
import "C"

// GTypeButton returns the GType for the type Button.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeButton() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Button").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalButton)
	return gtype
}

// ButtonOverrider contains methods that are overridable.
type ButtonOverrider interface {
	Activate()
	// Clicked emits a Button::clicked signal to the given Button.
	Clicked()
	// Enter emits a Button::enter signal to the given Button.
	//
	// Deprecated: Use the Widget::enter-notify-event signal.
	Enter()
	// Leave emits a Button::leave signal to the given Button.
	//
	// Deprecated: Use the Widget::leave-notify-event signal.
	Leave()
	// Pressed emits a Button::pressed signal to the given Button.
	//
	// Deprecated: Use the Widget::button-press-event signal.
	Pressed()
	// Released emits a Button::released signal to the given Button.
	//
	// Deprecated: Use the Widget::button-release-event signal.
	Released()
}

// Button widget is generally used to trigger a callback function that is called
// when the button is pressed. The various signals and how to use them are
// outlined below.
//
// The Button widget can hold any valid child widget. That is, it can hold
// almost any other standard Widget. The most commonly used child is the Label.
//
//
// CSS nodes
//
// GtkButton has a single CSS node with name button. The node will get the style
// classes .image-button or .text-button, if the content is just an image or
// label, respectively. It may also receive the .flat style class.
//
// Other style classes that are commonly used with GtkButton include
// .suggested-action and .destructive-action. In special cases, buttons can be
// made round by adding the .circular style class.
//
// Button-like widgets like ToggleButton, MenuButton, VolumeButton, LockButton,
// ColorButton, FontButton or FileChooserButton use style classes such as
// .toggle, .popup, .scale, .lock, .color, .font, .file to differentiate
// themselves from a plain GtkButton.
type Button struct {
	_ [0]func() // equal guard
	Bin

	*coreglib.Object
	Actionable
	Activatable
}

var (
	_ Binner            = (*Button)(nil)
	_ coreglib.Objector = (*Button)(nil)
)

func classInitButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "ButtonClass")

	if _, ok := goval.(interface{ Activate() }); ok {
		o := pclass.StructFieldOffset("activate")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ButtonClass_activate)
	}

	if _, ok := goval.(interface{ Clicked() }); ok {
		o := pclass.StructFieldOffset("clicked")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ButtonClass_clicked)
	}

	if _, ok := goval.(interface{ Enter() }); ok {
		o := pclass.StructFieldOffset("enter")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ButtonClass_enter)
	}

	if _, ok := goval.(interface{ Leave() }); ok {
		o := pclass.StructFieldOffset("leave")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ButtonClass_leave)
	}

	if _, ok := goval.(interface{ Pressed() }); ok {
		o := pclass.StructFieldOffset("pressed")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ButtonClass_pressed)
	}

	if _, ok := goval.(interface{ Released() }); ok {
		o := pclass.StructFieldOffset("released")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ButtonClass_released)
	}
}

//export _gotk4_gtk3_ButtonClass_activate
func _gotk4_gtk3_ButtonClass_activate(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Activate() })

	iface.Activate()
}

//export _gotk4_gtk3_ButtonClass_clicked
func _gotk4_gtk3_ButtonClass_clicked(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Clicked() })

	iface.Clicked()
}

//export _gotk4_gtk3_ButtonClass_enter
func _gotk4_gtk3_ButtonClass_enter(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Enter() })

	iface.Enter()
}

//export _gotk4_gtk3_ButtonClass_leave
func _gotk4_gtk3_ButtonClass_leave(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Leave() })

	iface.Leave()
}

//export _gotk4_gtk3_ButtonClass_pressed
func _gotk4_gtk3_ButtonClass_pressed(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Pressed() })

	iface.Pressed()
}

//export _gotk4_gtk3_ButtonClass_released
func _gotk4_gtk3_ButtonClass_released(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Released() })

	iface.Released()
}

func wrapButton(obj *coreglib.Object) *Button {
	return &Button{
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
	}
}

func marshalButton(p uintptr) (interface{}, error) {
	return wrapButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Button_ConnectActivate
func _gotk4_gtk3_Button_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectActivate signal on GtkButton is an action signal and emitting it
// causes the button to animate press then release. Applications should never
// connect to this signal, but use the Button::clicked signal.
func (button *Button) ConnectActivate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "activate", false, unsafe.Pointer(C._gotk4_gtk3_Button_ConnectActivate), f)
}

//export _gotk4_gtk3_Button_ConnectClicked
func _gotk4_gtk3_Button_ConnectClicked(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectClicked is emitted when the button has been activated (pressed and
// released).
func (button *Button) ConnectClicked(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "clicked", false, unsafe.Pointer(C._gotk4_gtk3_Button_ConnectClicked), f)
}

//export _gotk4_gtk3_Button_ConnectEnter
func _gotk4_gtk3_Button_ConnectEnter(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectEnter is emitted when the pointer enters the button.
func (button *Button) ConnectEnter(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "enter", false, unsafe.Pointer(C._gotk4_gtk3_Button_ConnectEnter), f)
}

//export _gotk4_gtk3_Button_ConnectLeave
func _gotk4_gtk3_Button_ConnectLeave(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectLeave is emitted when the pointer leaves the button.
func (button *Button) ConnectLeave(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "leave", false, unsafe.Pointer(C._gotk4_gtk3_Button_ConnectLeave), f)
}

//export _gotk4_gtk3_Button_ConnectPressed
func _gotk4_gtk3_Button_ConnectPressed(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectPressed is emitted when the button is pressed.
func (button *Button) ConnectPressed(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "pressed", false, unsafe.Pointer(C._gotk4_gtk3_Button_ConnectPressed), f)
}

//export _gotk4_gtk3_Button_ConnectReleased
func _gotk4_gtk3_Button_ConnectReleased(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectReleased is emitted when the button is released.
func (button *Button) ConnectReleased(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "released", false, unsafe.Pointer(C._gotk4_gtk3_Button_ConnectReleased), f)
}

// NewButton creates a new Button widget. To add a child widget to the button,
// use gtk_container_add().
//
// The function returns the following values:
//
//    - button: newly created Button widget.
//
func NewButton() *Button {
	_info := girepository.MustFind("Gtk", "Button")
	_gret := _info.InvokeClassMethod("new_Button", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _button *Button // out

	_button = wrapButton(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _button
}

// NewButtonFromStock creates a new Button containing the image and text from a
// [stock item][gtkstock]. Some stock ids have preprocessor macros like
// K_STOCK_OK and K_STOCK_APPLY.
//
// If stock_id is unknown, then it will be treated as a mnemonic label (as for
// gtk_button_new_with_mnemonic()).
//
// Deprecated: Stock items are deprecated. Use gtk_button_new_with_label()
// instead.
//
// The function takes the following parameters:
//
//    - stockId: name of the stock item.
//
// The function returns the following values:
//
//    - button: new Button.
//
func NewButtonFromStock(stockId string) *Button {
	var _args [1]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "Button")
	_gret := _info.InvokeClassMethod("new_Button_from_stock", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stockId)

	var _button *Button // out

	_button = wrapButton(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _button
}

// NewButtonWithLabel creates a Button widget with a Label child containing the
// given text.
//
// The function takes the following parameters:
//
//    - label: text you want the Label to hold.
//
// The function returns the following values:
//
//    - button: newly created Button widget.
//
func NewButtonWithLabel(label string) *Button {
	var _args [1]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "Button")
	_gret := _info.InvokeClassMethod("new_Button_with_label", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(label)

	var _button *Button // out

	_button = wrapButton(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _button
}

// NewButtonWithMnemonic creates a new Button containing a label. If characters
// in label are preceded by an underscore, they are underlined. If you need a
// literal underscore character in a label, use “__” (two underscores). The
// first underlined character represents a keyboard accelerator called a
// mnemonic. Pressing Alt and that key activates the button.
//
// The function takes the following parameters:
//
//    - label: text of the button, with an underscore in front of the mnemonic
//      character.
//
// The function returns the following values:
//
//    - button: new Button.
//
func NewButtonWithMnemonic(label string) *Button {
	var _args [1]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "Button")
	_gret := _info.InvokeClassMethod("new_Button_with_mnemonic", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(label)

	var _button *Button // out

	_button = wrapButton(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _button
}

// Clicked emits a Button::clicked signal to the given Button.
func (button *Button) Clicked() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_info := girepository.MustFind("Gtk", "Button")
	_info.InvokeClassMethod("clicked", _args[:], nil)

	runtime.KeepAlive(button)
}

// Enter emits a Button::enter signal to the given Button.
//
// Deprecated: Use the Widget::enter-notify-event signal.
func (button *Button) Enter() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_info := girepository.MustFind("Gtk", "Button")
	_info.InvokeClassMethod("enter", _args[:], nil)

	runtime.KeepAlive(button)
}

// Alignment gets the alignment of the child in the button.
//
// Deprecated: Access the child widget directly if you need to control its
// alignment.
//
// The function returns the following values:
//
//    - xalign: return location for horizontal alignment.
//    - yalign: return location for vertical alignment.
//
func (button *Button) Alignment() (xalign, yalign float32) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_info := girepository.MustFind("Gtk", "Button")
	_info.InvokeClassMethod("get_alignment", _args[:], _outs[:])

	runtime.KeepAlive(button)

	var _xalign float32 // out
	var _yalign float32 // out

	_xalign = float32(*(*C.gfloat)(unsafe.Pointer(&_outs[0])))
	_yalign = float32(*(*C.gfloat)(unsafe.Pointer(&_outs[1])))

	return _xalign, _yalign
}

// AlwaysShowImage returns whether the button will ignore the
// Settings:gtk-button-images setting and always show the image, if available.
//
// The function returns the following values:
//
//    - ok: TRUE if the button will always show the image.
//
func (button *Button) AlwaysShowImage() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_info := girepository.MustFind("Gtk", "Button")
	_gret := _info.InvokeClassMethod("get_always_show_image", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// EventWindow returns the button’s event window if it is realized, NULL
// otherwise. This function should be rarely needed.
//
// The function returns the following values:
//
//    - window button’s event window.
//
func (button *Button) EventWindow() gdk.Windower {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_info := girepository.MustFind("Gtk", "Button")
	_gret := _info.InvokeClassMethod("get_event_window", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _window gdk.Windower // out

	{
		objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.Windower)
			return ok
		})
		rv, ok := casted.(gdk.Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_window = rv
	}

	return _window
}

// FocusOnClick returns whether the button grabs focus when it is clicked with
// the mouse. See gtk_button_set_focus_on_click().
//
// Deprecated: Use gtk_widget_get_focus_on_click() instead.
//
// The function returns the following values:
//
//    - ok: TRUE if the button grabs focus when it is clicked with the mouse.
//
func (button *Button) FocusOnClick() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_info := girepository.MustFind("Gtk", "Button")
	_gret := _info.InvokeClassMethod("get_focus_on_click", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Image gets the widget that is currenty set as the image of button. This may
// have been explicitly set by gtk_button_set_image() or constructed by
// gtk_button_new_from_stock().
//
// The function returns the following values:
//
//    - widget (optional) or NULL in case there is no image.
//
func (button *Button) Image() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_info := girepository.MustFind("Gtk", "Button")
	_gret := _info.InvokeClassMethod("get_image", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))

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

// Label fetches the text from the label of the button, as set by
// gtk_button_set_label(). If the label text has not been set the return value
// will be NULL. This will be the case if you create an empty button with
// gtk_button_new() to use as a container.
//
// The function returns the following values:
//
//    - utf8: text of the label widget. This string is owned by the widget and
//      must not be modified or freed.
//
func (button *Button) Label() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_info := girepository.MustFind("Gtk", "Button")
	_gret := _info.InvokeClassMethod("get_label", _args[:], nil)
	_cret := *(**C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_cret)))))

	return _utf8
}

// UseStock returns whether the button label is a stock item.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - ok: TRUE if the button label is used to select a stock item instead of
//      being used directly as the label text.
//
func (button *Button) UseStock() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_info := girepository.MustFind("Gtk", "Button")
	_gret := _info.InvokeClassMethod("get_use_stock", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// UseUnderline returns whether an embedded underline in the button label
// indicates a mnemonic. See gtk_button_set_use_underline ().
//
// The function returns the following values:
//
//    - ok: TRUE if an embedded underline in the button label indicates the
//      mnemonic accelerator keys.
//
func (button *Button) UseUnderline() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_info := girepository.MustFind("Gtk", "Button")
	_gret := _info.InvokeClassMethod("get_use_underline", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(button)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Leave emits a Button::leave signal to the given Button.
//
// Deprecated: Use the Widget::leave-notify-event signal.
func (button *Button) Leave() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_info := girepository.MustFind("Gtk", "Button")
	_info.InvokeClassMethod("leave", _args[:], nil)

	runtime.KeepAlive(button)
}

// Pressed emits a Button::pressed signal to the given Button.
//
// Deprecated: Use the Widget::button-press-event signal.
func (button *Button) Pressed() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_info := girepository.MustFind("Gtk", "Button")
	_info.InvokeClassMethod("pressed", _args[:], nil)

	runtime.KeepAlive(button)
}

// Released emits a Button::released signal to the given Button.
//
// Deprecated: Use the Widget::button-release-event signal.
func (button *Button) Released() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_info := girepository.MustFind("Gtk", "Button")
	_info.InvokeClassMethod("released", _args[:], nil)

	runtime.KeepAlive(button)
}

// SetAlignment sets the alignment of the child. This property has no effect
// unless the child is a Misc or a Alignment.
//
// Deprecated: Access the child widget directly if you need to control its
// alignment.
//
// The function takes the following parameters:
//
//    - xalign: horizontal position of the child, 0.0 is left aligned, 1.0 is
//      right aligned.
//    - yalign: vertical position of the child, 0.0 is top aligned, 1.0 is bottom
//      aligned.
//
func (button *Button) SetAlignment(xalign, yalign float32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	*(*C.gfloat)(unsafe.Pointer(&_args[1])) = C.gfloat(xalign)
	*(*C.gfloat)(unsafe.Pointer(&_args[2])) = C.gfloat(yalign)

	_info := girepository.MustFind("Gtk", "Button")
	_info.InvokeClassMethod("set_alignment", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(xalign)
	runtime.KeepAlive(yalign)
}

// SetAlwaysShowImage: if TRUE, the button will ignore the
// Settings:gtk-button-images setting and always show the image, if available.
//
// Use this property if the button would be useless or hard to use without the
// image.
//
// The function takes the following parameters:
//
//    - alwaysShow: TRUE if the menuitem should always show the image.
//
func (button *Button) SetAlwaysShowImage(alwaysShow bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if alwaysShow {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Button")
	_info.InvokeClassMethod("set_always_show_image", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(alwaysShow)
}

// SetFocusOnClick sets whether the button will grab focus when it is clicked
// with the mouse. Making mouse clicks not grab focus is useful in places like
// toolbars where you don’t want the keyboard focus removed from the main area
// of the application.
//
// Deprecated: Use gtk_widget_set_focus_on_click() instead.
//
// The function takes the following parameters:
//
//    - focusOnClick: whether the button grabs focus when clicked with the mouse.
//
func (button *Button) SetFocusOnClick(focusOnClick bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if focusOnClick {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Button")
	_info.InvokeClassMethod("set_focus_on_click", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(focusOnClick)
}

// SetImage: set the image of button to the given widget. The image will be
// displayed if the label text is NULL or if Button:always-show-image is TRUE.
// You don’t have to call gtk_widget_show() on image yourself.
//
// The function takes the following parameters:
//
//    - image (optional): widget to set as the image for the button, or NULL to
//      unset.
//
func (button *Button) SetImage(image Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if image != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	}

	_info := girepository.MustFind("Gtk", "Button")
	_info.InvokeClassMethod("set_image", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(image)
}

// SetLabel sets the text of the label of the button to str. This text is also
// used to select the stock item if gtk_button_set_use_stock() is used.
//
// This will also clear any previously set labels.
//
// The function takes the following parameters:
//
//    - label: string.
//
func (button *Button) SetLabel(label string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gtk", "Button")
	_info.InvokeClassMethod("set_label", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(label)
}

// SetUseStock: if TRUE, the label set on the button is used as a stock id to
// select the stock item for the button.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - useStock: TRUE if the button should use a stock item.
//
func (button *Button) SetUseStock(useStock bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if useStock {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Button")
	_info.InvokeClassMethod("set_use_stock", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(useStock)
}

// SetUseUnderline: if true, an underline in the text of the button label
// indicates the next character should be used for the mnemonic accelerator key.
//
// The function takes the following parameters:
//
//    - useUnderline: TRUE if underlines in the text indicate mnemonics.
//
func (button *Button) SetUseUnderline(useUnderline bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if useUnderline {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Button")
	_info.InvokeClassMethod("set_use_underline", _args[:], nil)

	runtime.KeepAlive(button)
	runtime.KeepAlive(useUnderline)
}
