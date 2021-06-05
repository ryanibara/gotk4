// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.gtk_tool_button_get_type()), F: marshalToolButton},
	})
}

// ToolButton ToolButtons are ToolItems containing buttons.
//
// Use gtk_tool_button_new() to create a new ToolButton.
//
// The label of a ToolButton is determined by the properties
// ToolButton:label-widget, ToolButton:label, and ToolButton:stock-id. If
// ToolButton:label-widget is non-nil, then that widget is used as the label.
// Otherwise, if ToolButton:label is non-nil, that string is used as the label.
// Otherwise, if ToolButton:stock-id is non-nil, the label is determined by the
// stock item. Otherwise, the button does not have a label.
//
// The icon of a ToolButton is determined by the properties
// ToolButton:icon-widget and ToolButton:stock-id. If ToolButton:icon-widget is
// non-nil, then that widget is used as the icon. Otherwise, if
// ToolButton:stock-id is non-nil, the icon is determined by the stock item.
// Otherwise, the button does not have a icon.
//
//
// CSS nodes
//
// GtkToolButton has a single CSS node with name toolbutton.
type ToolButton interface {
	ToolItem
	Actionable
	Activatable
	Buildable

	// IconName returns the name of the themed icon for the tool button, see
	// gtk_tool_button_set_icon_name().
	IconName() string
	// IconWidget: return the widget used as icon widget on @button. See
	// gtk_tool_button_set_icon_widget().
	IconWidget() Widget
	// Label returns the label used by the tool button, or nil if the tool
	// button doesn’t have a label. or uses a the label from a stock item. The
	// returned string is owned by GTK+, and must not be modified or freed.
	Label() string
	// LabelWidget returns the widget used as label on @button. See
	// gtk_tool_button_set_label_widget().
	LabelWidget() Widget
	// StockID returns the name of the stock item. See
	// gtk_tool_button_set_stock_id(). The returned string is owned by GTK+ and
	// must not be freed or modifed.
	StockID() string
	// UseUnderline returns whether underscores in the label property are used
	// as mnemonics on menu items on the overflow menu. See
	// gtk_tool_button_set_use_underline().
	UseUnderline() bool
	// SetIconName sets the icon for the tool button from a named themed icon.
	// See the docs for IconTheme for more details. The ToolButton:icon-name
	// property only has an effect if not overridden by non-nil
	// ToolButton:label-widget, ToolButton:icon-widget and ToolButton:stock-id
	// properties.
	SetIconName(iconName string)
	// SetIconWidget sets @icon as the widget used as icon on @button. If
	// @icon_widget is nil the icon is determined by the ToolButton:stock-id
	// property. If the ToolButton:stock-id property is also nil, @button will
	// not have an icon.
	SetIconWidget(iconWidget Widget)
	// SetLabel sets @label as the label used for the tool button. The
	// ToolButton:label property only has an effect if not overridden by a
	// non-nil ToolButton:label-widget property. If both the
	// ToolButton:label-widget and ToolButton:label properties are nil, the
	// label is determined by the ToolButton:stock-id property. If the
	// ToolButton:stock-id property is also nil, @button will not have a label.
	SetLabel(label string)
	// SetLabelWidget sets @label_widget as the widget that will be used as the
	// label for @button. If @label_widget is nil the ToolButton:label property
	// is used as label. If ToolButton:label is also nil, the label in the stock
	// item determined by the ToolButton:stock-id property is used as label. If
	// ToolButton:stock-id is also nil, @button does not have a label.
	SetLabelWidget(labelWidget Widget)
	// SetStockID sets the name of the stock item. See
	// gtk_tool_button_new_from_stock(). The stock_id property only has an
	// effect if not overridden by non-nil ToolButton:label-widget and
	// ToolButton:icon-widget properties.
	SetStockID(stockID string)
	// SetUseUnderline: if set, an underline in the label property indicates
	// that the next character should be used for the mnemonic accelerator key
	// in the overflow menu. For example, if the label property is “_Open” and
	// @use_underline is true, the label on the tool button will be “Open” and
	// the item on the overflow menu will have an underlined “O”.
	//
	// Labels shown on tool buttons never have mnemonics on them; this property
	// only affects the menu item on the overflow menu.
	SetUseUnderline(useUnderline bool)
}

// toolButton implements the ToolButton interface.
type toolButton struct {
	ToolItem
	Actionable
	Activatable
	Buildable
}

var _ ToolButton = (*toolButton)(nil)

// WrapToolButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapToolButton(obj *externglib.Object) ToolButton {
	return ToolButton{
		ToolItem:    WrapToolItem(obj),
		Actionable:  WrapActionable(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalToolButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToolButton(obj), nil
}

// NewToolButton constructs a class ToolButton.
func NewToolButton(iconWidget Widget, label string) ToolButton {
	var arg1 *C.GtkWidget
	var arg2 *C.gchar

	arg1 = (*C.GtkWidget)(unsafe.Pointer(iconWidget.Native()))
	arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.GtkToolButton
	var ret1 ToolButton

	cret = C.gtk_tool_button_new(iconWidget, label)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ToolButton)

	return ret1
}

// NewToolButtonFromStock constructs a class ToolButton.
func NewToolButtonFromStock(stockID string) ToolButton {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(stockID))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkToolButton
	var ret1 ToolButton

	cret = C.gtk_tool_button_new_from_stock(stockID)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ToolButton)

	return ret1
}

// IconName returns the name of the themed icon for the tool button, see
// gtk_tool_button_set_icon_name().
func (b toolButton) IconName() string {
	var arg0 *C.GtkToolButton

	arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_tool_button_get_icon_name(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// IconWidget: return the widget used as icon widget on @button. See
// gtk_tool_button_set_icon_widget().
func (b toolButton) IconWidget() Widget {
	var arg0 *C.GtkToolButton

	arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_tool_button_get_icon_widget(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// Label returns the label used by the tool button, or nil if the tool
// button doesn’t have a label. or uses a the label from a stock item. The
// returned string is owned by GTK+, and must not be modified or freed.
func (b toolButton) Label() string {
	var arg0 *C.GtkToolButton

	arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_tool_button_get_label(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// LabelWidget returns the widget used as label on @button. See
// gtk_tool_button_set_label_widget().
func (b toolButton) LabelWidget() Widget {
	var arg0 *C.GtkToolButton

	arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_tool_button_get_label_widget(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// StockID returns the name of the stock item. See
// gtk_tool_button_set_stock_id(). The returned string is owned by GTK+ and
// must not be freed or modifed.
func (b toolButton) StockID() string {
	var arg0 *C.GtkToolButton

	arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_tool_button_get_stock_id(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// UseUnderline returns whether underscores in the label property are used
// as mnemonics on menu items on the overflow menu. See
// gtk_tool_button_set_use_underline().
func (b toolButton) UseUnderline() bool {
	var arg0 *C.GtkToolButton

	arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_tool_button_get_use_underline(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// SetIconName sets the icon for the tool button from a named themed icon.
// See the docs for IconTheme for more details. The ToolButton:icon-name
// property only has an effect if not overridden by non-nil
// ToolButton:label-widget, ToolButton:icon-widget and ToolButton:stock-id
// properties.
func (b toolButton) SetIconName(iconName string) {
	var arg0 *C.GtkToolButton
	var arg1 *C.gchar

	arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_tool_button_set_icon_name(arg0, iconName)
}

// SetIconWidget sets @icon as the widget used as icon on @button. If
// @icon_widget is nil the icon is determined by the ToolButton:stock-id
// property. If the ToolButton:stock-id property is also nil, @button will
// not have an icon.
func (b toolButton) SetIconWidget(iconWidget Widget) {
	var arg0 *C.GtkToolButton
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(iconWidget.Native()))

	C.gtk_tool_button_set_icon_widget(arg0, iconWidget)
}

// SetLabel sets @label as the label used for the tool button. The
// ToolButton:label property only has an effect if not overridden by a
// non-nil ToolButton:label-widget property. If both the
// ToolButton:label-widget and ToolButton:label properties are nil, the
// label is determined by the ToolButton:stock-id property. If the
// ToolButton:stock-id property is also nil, @button will not have a label.
func (b toolButton) SetLabel(label string) {
	var arg0 *C.GtkToolButton
	var arg1 *C.gchar

	arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_tool_button_set_label(arg0, label)
}

// SetLabelWidget sets @label_widget as the widget that will be used as the
// label for @button. If @label_widget is nil the ToolButton:label property
// is used as label. If ToolButton:label is also nil, the label in the stock
// item determined by the ToolButton:stock-id property is used as label. If
// ToolButton:stock-id is also nil, @button does not have a label.
func (b toolButton) SetLabelWidget(labelWidget Widget) {
	var arg0 *C.GtkToolButton
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))

	C.gtk_tool_button_set_label_widget(arg0, labelWidget)
}

// SetStockID sets the name of the stock item. See
// gtk_tool_button_new_from_stock(). The stock_id property only has an
// effect if not overridden by non-nil ToolButton:label-widget and
// ToolButton:icon-widget properties.
func (b toolButton) SetStockID(stockID string) {
	var arg0 *C.GtkToolButton
	var arg1 *C.gchar

	arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(stockID))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_tool_button_set_stock_id(arg0, stockID)
}

// SetUseUnderline: if set, an underline in the label property indicates
// that the next character should be used for the mnemonic accelerator key
// in the overflow menu. For example, if the label property is “_Open” and
// @use_underline is true, the label on the tool button will be “Open” and
// the item on the overflow menu will have an underlined “O”.
//
// Labels shown on tool buttons never have mnemonics on them; this property
// only affects the menu item on the overflow menu.
func (b toolButton) SetUseUnderline(useUnderline bool) {
	var arg0 *C.GtkToolButton
	var arg1 C.gboolean

	arg0 = (*C.GtkToolButton)(unsafe.Pointer(b.Native()))
	if useUnderline {
		arg1 = C.gboolean(1)
	}

	C.gtk_tool_button_set_use_underline(arg0, useUnderline)
}

type ToolButtonPrivate struct {
	native C.GtkToolButtonPrivate
}

// WrapToolButtonPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapToolButtonPrivate(ptr unsafe.Pointer) *ToolButtonPrivate {
	if ptr == nil {
		return nil
	}

	return (*ToolButtonPrivate)(ptr)
}

func marshalToolButtonPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapToolButtonPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *ToolButtonPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}
