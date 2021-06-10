// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_menu_button_get_type()), F: marshalMenuButton},
	})
}

// MenuButton: the MenuButton widget is used to display a popup when clicked on.
// This popup can be provided either as a Menu, a Popover or an abstract Model.
//
// The MenuButton widget can hold any valid child widget. That is, it can hold
// almost any other standard Widget. The most commonly used child is Image. If
// no widget is explicitely added to the MenuButton, a Image is automatically
// created, using an arrow image oriented according to MenuButton:direction or
// the generic “open-menu-symbolic” icon if the direction is not set.
//
// The positioning of the popup is determined by the MenuButton:direction
// property of the menu button.
//
// For menus, the Widget:halign and Widget:valign properties of the menu are
// also taken into account. For example, when the direction is GTK_ARROW_DOWN
// and the horizontal alignment is GTK_ALIGN_START, the menu will be positioned
// below the button, with the starting edge (depending on the text direction) of
// the menu aligned with the starting edge of the button. If there is not enough
// space below the button, the menu is popped up above the button instead. If
// the alignment would move part of the menu offscreen, it is “pushed in”.
//
// Direction = Down
//
// - halign = start
//
//    ! (down-start.png)
//
// - halign = center
//
//    ! (down-center.png)
//
// - halign = end
//
//    ! (down-end.png)
//
// Direction = Up
//
// - halign = start
//
//    ! (up-start.png)
//
// - halign = center
//
//    ! (up-center.png)
//
// - halign = end
//
//    ! (up-end.png)
//
// Direction = Left
//
// - valign = start
//
//    ! (left-start.png)
//
// - valign = center
//
//    ! (left-center.png)
//
// - valign = end
//
//    ! (left-end.png)
//
// Direction = Right
//
// - valign = start
//
//    ! (right-start.png)
//
// - valign = center
//
//    ! (right-center.png)
//
// - valign = end
//
//    ! (right-end.png)
//
//
// CSS nodes
//
// GtkMenuButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .popup style class.
type MenuButton interface {
	ToggleButton
	Actionable
	Activatable
	Buildable

	// UsePopover returns whether a Popover or a Menu will be constructed from
	// the menu model.
	UsePopover() bool
	// SetAlignWidget sets the Widget to use to line the menu with when popped
	// up. Note that the @align_widget must contain the MenuButton itself.
	//
	// Setting it to nil means that the menu will be aligned with the button
	// itself.
	//
	// Note that this property is only used with menus currently, and not for
	// popovers.
	SetAlignWidget(alignWidget Widget)
	// SetDirection sets the direction in which the popup will be popped up, as
	// well as changing the arrow’s direction. The child will not be changed to
	// an arrow if it was customized.
	//
	// If the does not fit in the available space in the given direction, GTK+
	// will its best to keep it inside the screen and fully visible.
	//
	// If you pass GTK_ARROW_NONE for a @direction, the popup will behave as if
	// you passed GTK_ARROW_DOWN (although you won’t see any arrows).
	SetDirection(direction ArrowType)
	// SetMenuModel sets the Model from which the popup will be constructed, or
	// nil to dissociate any existing menu model and disable the button.
	//
	// Depending on the value of MenuButton:use-popover, either a Menu will be
	// created with gtk_menu_new_from_model(), or a Popover with
	// gtk_popover_new_from_model(). In either case, actions will be connected
	// as documented for these functions.
	//
	// If MenuButton:popup or MenuButton:popover are already set, those widgets
	// are dissociated from the @menu_button, and those properties are set to
	// nil.
	SetMenuModel(menuModel gio.MenuModel)
	// SetPopover sets the Popover that will be popped up when the @menu_button
	// is clicked, or nil to dissociate any existing popover and disable the
	// button.
	//
	// If MenuButton:menu-model or MenuButton:popup are set, those objects are
	// dissociated from the @menu_button, and those properties are set to nil.
	SetPopover(popover Widget)
	// SetPopup sets the Menu that will be popped up when the @menu_button is
	// clicked, or nil to dissociate any existing menu and disable the button.
	//
	// If MenuButton:menu-model or MenuButton:popover are set, those objects are
	// dissociated from the @menu_button, and those properties are set to nil.
	SetPopup(menu Widget)
	// SetUsePopover sets whether to construct a Popover instead of Menu when
	// gtk_menu_button_set_menu_model() is called. Note that this property is
	// only consulted when a new menu model is set.
	SetUsePopover(usePopover bool)
}

// menuButton implements the MenuButton interface.
type menuButton struct {
	ToggleButton
	Actionable
	Activatable
	Buildable
}

var _ MenuButton = (*menuButton)(nil)

// WrapMenuButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapMenuButton(obj *externglib.Object) MenuButton {
	return MenuButton{
		ToggleButton: WrapToggleButton(obj),
		Actionable:   WrapActionable(obj),
		Activatable:  WrapActivatable(obj),
		Buildable:    WrapBuildable(obj),
	}
}

func marshalMenuButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMenuButton(obj), nil
}

// UsePopover returns whether a Popover or a Menu will be constructed from
// the menu model.
func (m menuButton) UsePopover() bool {
	var _arg0 *C.GtkMenuButton

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(m.Native()))

	var _cret C.gboolean

	_cret = C.gtk_menu_button_get_use_popover(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetAlignWidget sets the Widget to use to line the menu with when popped
// up. Note that the @align_widget must contain the MenuButton itself.
//
// Setting it to nil means that the menu will be aligned with the button
// itself.
//
// Note that this property is only used with menus currently, and not for
// popovers.
func (m menuButton) SetAlignWidget(alignWidget Widget) {
	var _arg0 *C.GtkMenuButton
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(alignWidget.Native()))

	C.gtk_menu_button_set_align_widget(_arg0, _arg1)
}

// SetDirection sets the direction in which the popup will be popped up, as
// well as changing the arrow’s direction. The child will not be changed to
// an arrow if it was customized.
//
// If the does not fit in the available space in the given direction, GTK+
// will its best to keep it inside the screen and fully visible.
//
// If you pass GTK_ARROW_NONE for a @direction, the popup will behave as if
// you passed GTK_ARROW_DOWN (although you won’t see any arrows).
func (m menuButton) SetDirection(direction ArrowType) {
	var _arg0 *C.GtkMenuButton
	var _arg1 C.GtkArrowType

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(m.Native()))
	_arg1 = (C.GtkArrowType)(direction)

	C.gtk_menu_button_set_direction(_arg0, _arg1)
}

// SetMenuModel sets the Model from which the popup will be constructed, or
// nil to dissociate any existing menu model and disable the button.
//
// Depending on the value of MenuButton:use-popover, either a Menu will be
// created with gtk_menu_new_from_model(), or a Popover with
// gtk_popover_new_from_model(). In either case, actions will be connected
// as documented for these functions.
//
// If MenuButton:popup or MenuButton:popover are already set, those widgets
// are dissociated from the @menu_button, and those properties are set to
// nil.
func (m menuButton) SetMenuModel(menuModel gio.MenuModel) {
	var _arg0 *C.GtkMenuButton
	var _arg1 *C.GMenuModel

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(menuModel.Native()))

	C.gtk_menu_button_set_menu_model(_arg0, _arg1)
}

// SetPopover sets the Popover that will be popped up when the @menu_button
// is clicked, or nil to dissociate any existing popover and disable the
// button.
//
// If MenuButton:menu-model or MenuButton:popup are set, those objects are
// dissociated from the @menu_button, and those properties are set to nil.
func (m menuButton) SetPopover(popover Widget) {
	var _arg0 *C.GtkMenuButton
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(popover.Native()))

	C.gtk_menu_button_set_popover(_arg0, _arg1)
}

// SetPopup sets the Menu that will be popped up when the @menu_button is
// clicked, or nil to dissociate any existing menu and disable the button.
//
// If MenuButton:menu-model or MenuButton:popover are set, those objects are
// dissociated from the @menu_button, and those properties are set to nil.
func (m menuButton) SetPopup(menu Widget) {
	var _arg0 *C.GtkMenuButton
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(menu.Native()))

	C.gtk_menu_button_set_popup(_arg0, _arg1)
}

// SetUsePopover sets whether to construct a Popover instead of Menu when
// gtk_menu_button_set_menu_model() is called. Note that this property is
// only consulted when a new menu model is set.
func (m menuButton) SetUsePopover(usePopover bool) {
	var _arg0 *C.GtkMenuButton
	var _arg1 C.gboolean

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(m.Native()))
	if usePopover {
		_arg1 = C.gboolean(1)
	}

	C.gtk_menu_button_set_use_popover(_arg0, _arg1)
}
