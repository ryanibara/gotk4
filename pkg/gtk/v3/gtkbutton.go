// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_button_get_type()), F: marshalButton},
	})
}

// Button: the Button widget is generally used to trigger a callback function
// that is called when the button is pressed. The various signals and how to use
// them are outlined below.
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
type Button interface {
	Bin
	Actionable
	Activatable
	Buildable

	// Clicked emits a Button::clicked signal to the given Button.
	Clicked()
	// Enter emits a Button::enter signal to the given Button.
	Enter()
	// Alignment gets the alignment of the child in the button.
	Alignment() (xalign float32, yalign float32)
	// AlwaysShowImage returns whether the button will ignore the
	// Settings:gtk-button-images setting and always show the image, if
	// available.
	AlwaysShowImage() bool
	// EventWindow returns the button’s event window if it is realized, nil
	// otherwise. This function should be rarely needed.
	EventWindow() gdk.Window
	// FocusOnClick returns whether the button grabs focus when it is clicked
	// with the mouse. See gtk_button_set_focus_on_click().
	FocusOnClick() bool
	// Image gets the widget that is currenty set as the image of @button. This
	// may have been explicitly set by gtk_button_set_image() or constructed by
	// gtk_button_new_from_stock().
	Image() Widget
	// ImagePosition gets the position of the image relative to the text inside
	// the button.
	ImagePosition() PositionType
	// Label fetches the text from the label of the button, as set by
	// gtk_button_set_label(). If the label text has not been set the return
	// value will be nil. This will be the case if you create an empty button
	// with gtk_button_new() to use as a container.
	Label() string
	// Relief returns the current relief style of the given Button.
	Relief() ReliefStyle
	// UseStock returns whether the button label is a stock item.
	UseStock() bool
	// UseUnderline returns whether an embedded underline in the button label
	// indicates a mnemonic. See gtk_button_set_use_underline ().
	UseUnderline() bool
	// Leave emits a Button::leave signal to the given Button.
	Leave()
	// Pressed emits a Button::pressed signal to the given Button.
	Pressed()
	// Released emits a Button::released signal to the given Button.
	Released()
	// SetAlignment sets the alignment of the child. This property has no effect
	// unless the child is a Misc or a Alignment.
	SetAlignment(xalign float32, yalign float32)
	// SetAlwaysShowImage: if true, the button will ignore the
	// Settings:gtk-button-images setting and always show the image, if
	// available.
	//
	// Use this property if the button would be useless or hard to use without
	// the image.
	SetAlwaysShowImage(alwaysShow bool)
	// SetFocusOnClick sets whether the button will grab focus when it is
	// clicked with the mouse. Making mouse clicks not grab focus is useful in
	// places like toolbars where you don’t want the keyboard focus removed from
	// the main area of the application.
	SetFocusOnClick(focusOnClick bool)
	// SetImage: set the image of @button to the given widget. The image will be
	// displayed if the label text is nil or if Button:always-show-image is
	// true. You don’t have to call gtk_widget_show() on @image yourself.
	SetImage(image Widget)
	// SetImagePosition sets the position of the image relative to the text
	// inside the button.
	SetImagePosition(position PositionType)
	// SetLabel sets the text of the label of the button to @str. This text is
	// also used to select the stock item if gtk_button_set_use_stock() is used.
	//
	// This will also clear any previously set labels.
	SetLabel(label string)
	// SetRelief sets the relief style of the edges of the given Button widget.
	// Two styles exist, GTK_RELIEF_NORMAL and GTK_RELIEF_NONE. The default
	// style is, as one can guess, GTK_RELIEF_NORMAL. The deprecated value
	// GTK_RELIEF_HALF behaves the same as GTK_RELIEF_NORMAL.
	SetRelief(relief ReliefStyle)
	// SetUseStock: if true, the label set on the button is used as a stock id
	// to select the stock item for the button.
	SetUseStock(useStock bool)
	// SetUseUnderline: if true, an underline in the text of the button label
	// indicates the next character should be used for the mnemonic accelerator
	// key.
	SetUseUnderline(useUnderline bool)
}

// button implements the Button interface.
type button struct {
	Bin
	Actionable
	Activatable
	Buildable
}

var _ Button = (*button)(nil)

// WrapButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapButton(obj *externglib.Object) Button {
	return Button{
		Bin:         WrapBin(obj),
		Actionable:  WrapActionable(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapButton(obj), nil
}

// NewButton constructs a class Button.
func NewButton() Button {
	var cret C.GtkButton
	var ret1 Button

	cret = C.gtk_button_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Button)

	return ret1
}

// NewButtonFromIconName constructs a class Button.
func NewButtonFromIconName(iconName string, size int) Button {
	var arg1 *C.gchar
	var arg2 C.GtkIconSize

	arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.GtkIconSize(size)

	var cret C.GtkButton
	var ret1 Button

	cret = C.gtk_button_new_from_icon_name(iconName, size)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Button)

	return ret1
}

// NewButtonFromStock constructs a class Button.
func NewButtonFromStock(stockID string) Button {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(stockID))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkButton
	var ret1 Button

	cret = C.gtk_button_new_from_stock(stockID)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Button)

	return ret1
}

// NewButtonWithLabel constructs a class Button.
func NewButtonWithLabel(label string) Button {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkButton
	var ret1 Button

	cret = C.gtk_button_new_with_label(label)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Button)

	return ret1
}

// NewButtonWithMnemonic constructs a class Button.
func NewButtonWithMnemonic(label string) Button {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkButton
	var ret1 Button

	cret = C.gtk_button_new_with_mnemonic(label)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Button)

	return ret1
}

// Clicked emits a Button::clicked signal to the given Button.
func (b button) Clicked() {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	C.gtk_button_clicked(arg0)
}

// Enter emits a Button::enter signal to the given Button.
func (b button) Enter() {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	C.gtk_button_enter(arg0)
}

// Alignment gets the alignment of the child in the button.
func (b button) Alignment() (xalign float32, yalign float32) {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var arg1 C.gfloat
	var ret1 float32
	var arg2 C.gfloat
	var ret2 float32

	C.gtk_button_get_alignment(arg0, &arg1, &arg2)

	*ret1 = C.gfloat(arg1)
	*ret2 = C.gfloat(arg2)

	return ret1, ret2
}

// AlwaysShowImage returns whether the button will ignore the
// Settings:gtk-button-images setting and always show the image, if
// available.
func (b button) AlwaysShowImage() bool {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_button_get_always_show_image(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// EventWindow returns the button’s event window if it is realized, nil
// otherwise. This function should be rarely needed.
func (b button) EventWindow() gdk.Window {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var cret *C.GdkWindow
	var ret1 gdk.Window

	cret = C.gtk_button_get_event_window(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Window)

	return ret1
}

// FocusOnClick returns whether the button grabs focus when it is clicked
// with the mouse. See gtk_button_set_focus_on_click().
func (b button) FocusOnClick() bool {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_button_get_focus_on_click(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Image gets the widget that is currenty set as the image of @button. This
// may have been explicitly set by gtk_button_set_image() or constructed by
// gtk_button_new_from_stock().
func (b button) Image() Widget {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_button_get_image(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// ImagePosition gets the position of the image relative to the text inside
// the button.
func (b button) ImagePosition() PositionType {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var cret C.GtkPositionType
	var ret1 PositionType

	cret = C.gtk_button_get_image_position(arg0)

	ret1 = PositionType(cret)

	return ret1
}

// Label fetches the text from the label of the button, as set by
// gtk_button_set_label(). If the label text has not been set the return
// value will be nil. This will be the case if you create an empty button
// with gtk_button_new() to use as a container.
func (b button) Label() string {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_button_get_label(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Relief returns the current relief style of the given Button.
func (b button) Relief() ReliefStyle {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var cret C.GtkReliefStyle
	var ret1 ReliefStyle

	cret = C.gtk_button_get_relief(arg0)

	ret1 = ReliefStyle(cret)

	return ret1
}

// UseStock returns whether the button label is a stock item.
func (b button) UseStock() bool {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_button_get_use_stock(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// UseUnderline returns whether an embedded underline in the button label
// indicates a mnemonic. See gtk_button_set_use_underline ().
func (b button) UseUnderline() bool {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_button_get_use_underline(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Leave emits a Button::leave signal to the given Button.
func (b button) Leave() {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	C.gtk_button_leave(arg0)
}

// Pressed emits a Button::pressed signal to the given Button.
func (b button) Pressed() {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	C.gtk_button_pressed(arg0)
}

// Released emits a Button::released signal to the given Button.
func (b button) Released() {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	C.gtk_button_released(arg0)
}

// SetAlignment sets the alignment of the child. This property has no effect
// unless the child is a Misc or a Alignment.
func (b button) SetAlignment(xalign float32, yalign float32) {
	var arg0 *C.GtkButton
	var arg1 C.gfloat
	var arg2 C.gfloat

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	arg1 = C.gfloat(xalign)
	arg2 = C.gfloat(yalign)

	C.gtk_button_set_alignment(arg0, xalign, yalign)
}

// SetAlwaysShowImage: if true, the button will ignore the
// Settings:gtk-button-images setting and always show the image, if
// available.
//
// Use this property if the button would be useless or hard to use without
// the image.
func (b button) SetAlwaysShowImage(alwaysShow bool) {
	var arg0 *C.GtkButton
	var arg1 C.gboolean

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	if alwaysShow {
		arg1 = C.gboolean(1)
	}

	C.gtk_button_set_always_show_image(arg0, alwaysShow)
}

// SetFocusOnClick sets whether the button will grab focus when it is
// clicked with the mouse. Making mouse clicks not grab focus is useful in
// places like toolbars where you don’t want the keyboard focus removed from
// the main area of the application.
func (b button) SetFocusOnClick(focusOnClick bool) {
	var arg0 *C.GtkButton
	var arg1 C.gboolean

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	if focusOnClick {
		arg1 = C.gboolean(1)
	}

	C.gtk_button_set_focus_on_click(arg0, focusOnClick)
}

// SetImage: set the image of @button to the given widget. The image will be
// displayed if the label text is nil or if Button:always-show-image is
// true. You don’t have to call gtk_widget_show() on @image yourself.
func (b button) SetImage(image Widget) {
	var arg0 *C.GtkButton
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(image.Native()))

	C.gtk_button_set_image(arg0, image)
}

// SetImagePosition sets the position of the image relative to the text
// inside the button.
func (b button) SetImagePosition(position PositionType) {
	var arg0 *C.GtkButton
	var arg1 C.GtkPositionType

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	arg1 = (C.GtkPositionType)(position)

	C.gtk_button_set_image_position(arg0, position)
}

// SetLabel sets the text of the label of the button to @str. This text is
// also used to select the stock item if gtk_button_set_use_stock() is used.
//
// This will also clear any previously set labels.
func (b button) SetLabel(label string) {
	var arg0 *C.GtkButton
	var arg1 *C.gchar

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_button_set_label(arg0, label)
}

// SetRelief sets the relief style of the edges of the given Button widget.
// Two styles exist, GTK_RELIEF_NORMAL and GTK_RELIEF_NONE. The default
// style is, as one can guess, GTK_RELIEF_NORMAL. The deprecated value
// GTK_RELIEF_HALF behaves the same as GTK_RELIEF_NORMAL.
func (b button) SetRelief(relief ReliefStyle) {
	var arg0 *C.GtkButton
	var arg1 C.GtkReliefStyle

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	arg1 = (C.GtkReliefStyle)(relief)

	C.gtk_button_set_relief(arg0, relief)
}

// SetUseStock: if true, the label set on the button is used as a stock id
// to select the stock item for the button.
func (b button) SetUseStock(useStock bool) {
	var arg0 *C.GtkButton
	var arg1 C.gboolean

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	if useStock {
		arg1 = C.gboolean(1)
	}

	C.gtk_button_set_use_stock(arg0, useStock)
}

// SetUseUnderline: if true, an underline in the text of the button label
// indicates the next character should be used for the mnemonic accelerator
// key.
func (b button) SetUseUnderline(useUnderline bool) {
	var arg0 *C.GtkButton
	var arg1 C.gboolean

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	if useUnderline {
		arg1 = C.gboolean(1)
	}

	C.gtk_button_set_use_underline(arg0, useUnderline)
}

type ButtonPrivate struct {
	native C.GtkButtonPrivate
}

// WrapButtonPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapButtonPrivate(ptr unsafe.Pointer) *ButtonPrivate {
	if ptr == nil {
		return nil
	}

	return (*ButtonPrivate)(ptr)
}

func marshalButtonPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapButtonPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (b *ButtonPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}
