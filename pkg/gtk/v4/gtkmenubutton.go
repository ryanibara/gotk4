// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void callbackDelete(gpointer);
// void _gotk4_gtk4_MenuButtonCreatePopupFunc(GtkMenuButton*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_menu_button_get_type()), F: marshalMenuButtonner},
	})
}

// MenuButtonCreatePopupFunc: user-provided callback function to create a popup
// for a GtkMenuButton on demand.
//
// This function is called when the popup of menu_button is shown, but none has
// been provided via gtk.MenuButton.SetPopover() or
// gtk.MenuButton.SetMenuModel().
type MenuButtonCreatePopupFunc func(menuButton *MenuButton)

//export _gotk4_gtk4_MenuButtonCreatePopupFunc
func _gotk4_gtk4_MenuButtonCreatePopupFunc(arg0 *C.GtkMenuButton, arg1 C.gpointer) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var menuButton *MenuButton // out

	menuButton = wrapMenuButton(externglib.Take(unsafe.Pointer(arg0)))

	fn := v.(MenuButtonCreatePopupFunc)
	fn(menuButton)
}

// MenuButton: GtkMenuButton widget is used to display a popup when clicked.
//
// !An example GtkMenuButton (menu-button.png)
//
// This popup can be provided either as a GtkPopover or as an abstract
// GMenuModel.
//
// The GtkMenuButton widget can show either an icon (set with the
// gtk.MenuButton:icon-name property) or a label (set with the
// gtk.MenuButton:label property). If neither is explicitly set, a gtk.Image is
// automatically created, using an arrow image oriented according to
// gtk.MenuButton:direction or the generic “open-menu-symbolic” icon if the
// direction is not set.
//
// The positioning of the popup is determined by the gtk.MenuButton:direction
// property of the menu button.
//
// For menus, the gtk.Widget:halign and gtk.Widget:valign properties of the menu
// are also taken into account. For example, when the direction is
// GTK_ARROW_DOWN and the horizontal alignment is GTK_ALIGN_START, the menu will
// be positioned below the button, with the starting edge (depending on the text
// direction) of the menu aligned with the starting edge of the button. If there
// is not enough space below the button, the menu is popped up above the button
// instead. If the alignment would move part of the menu offscreen, it is
// “pushed in”.
//
// | | start | center | end | | - | --- | --- | --- | | **down** | !
// (down-start.png) | ! (down-center.png) | ! (down-end.png) | | **up** | !
// (up-start.png) | ! (up-center.png) | ! (up-end.png) | | **left** | !
// (left-start.png) | ! (left-center.png) | ! (left-end.png) | | **right** | !
// (right-start.png) | ! (right-center.png) | ! (right-end.png) |
//
// CSS nodes
//
//    menubutton
//    ╰── button.toggle
//        ╰── <content>
//             ╰── [arrow]
//
//
// GtkMenuButton has a single CSS node with name menubutton which contains a
// button node with a .toggle style class.
//
// Inside the toggle button content, there is an arrow node for the indicator,
// which will carry one of the .none, .up, .down, .left or .right style classes
// to indicate the direction that the menu will appear in. The CSS is expected
// to provide a suitable image for each of these cases using the
// -gtk-icon-source property.
//
// Optionally, the menubutton node can carry the .circular style class to
// request a round appearance.
//
//
// Accessibility
//
// GtkMenuButton uses the K_ACCESSIBLE_ROLE_BUTTON role.
type MenuButton struct {
	Widget
}

func wrapMenuButton(obj *externglib.Object) *MenuButton {
	return &MenuButton{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalMenuButtonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMenuButton(obj), nil
}

// NewMenuButton creates a new GtkMenuButton widget with downwards-pointing
// arrow as the only child.
//
// You can replace the child widget with another GtkWidget should you wish to.
func NewMenuButton() *MenuButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_menu_button_new()

	var _menuButton *MenuButton // out

	_menuButton = wrapMenuButton(externglib.Take(unsafe.Pointer(_cret)))

	return _menuButton
}

// Direction returns the direction the popup will be pointing at when popped up.
func (menuButton *MenuButton) Direction() ArrowType {
	var _arg0 *C.GtkMenuButton // out
	var _cret C.GtkArrowType   // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	_cret = C.gtk_menu_button_get_direction(_arg0)
	runtime.KeepAlive(menuButton)

	var _arrowType ArrowType // out

	_arrowType = ArrowType(_cret)

	return _arrowType
}

// HasFrame returns whether the button has a frame.
func (menuButton *MenuButton) HasFrame() bool {
	var _arg0 *C.GtkMenuButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	_cret = C.gtk_menu_button_get_has_frame(_arg0)
	runtime.KeepAlive(menuButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IconName gets the name of the icon shown in the button.
func (menuButton *MenuButton) IconName() string {
	var _arg0 *C.GtkMenuButton // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	_cret = C.gtk_menu_button_get_icon_name(_arg0)
	runtime.KeepAlive(menuButton)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Label gets the label shown in the button
func (menuButton *MenuButton) Label() string {
	var _arg0 *C.GtkMenuButton // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	_cret = C.gtk_menu_button_get_label(_arg0)
	runtime.KeepAlive(menuButton)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// MenuModel returns the GMenuModel used to generate the popup.
func (menuButton *MenuButton) MenuModel() gio.MenuModeller {
	var _arg0 *C.GtkMenuButton // out
	var _cret *C.GMenuModel    // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	_cret = C.gtk_menu_button_get_menu_model(_arg0)
	runtime.KeepAlive(menuButton)

	var _menuModel gio.MenuModeller // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(gio.MenuModeller)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.MenuModeller")
			}
			_menuModel = rv
		}
	}

	return _menuModel
}

// Popover returns the GtkPopover that pops out of the button.
//
// If the button is not using a GtkPopover, this function returns NULL.
func (menuButton *MenuButton) Popover() *Popover {
	var _arg0 *C.GtkMenuButton // out
	var _cret *C.GtkPopover    // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	_cret = C.gtk_menu_button_get_popover(_arg0)
	runtime.KeepAlive(menuButton)

	var _popover *Popover // out

	if _cret != nil {
		_popover = wrapPopover(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _popover
}

// UseUnderline returns whether an embedded underline in the text indicates a
// mnemonic.
func (menuButton *MenuButton) UseUnderline() bool {
	var _arg0 *C.GtkMenuButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	_cret = C.gtk_menu_button_get_use_underline(_arg0)
	runtime.KeepAlive(menuButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Popdown dismiss the menu.
func (menuButton *MenuButton) Popdown() {
	var _arg0 *C.GtkMenuButton // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	C.gtk_menu_button_popdown(_arg0)
	runtime.KeepAlive(menuButton)
}

// Popup: pop up the menu.
func (menuButton *MenuButton) Popup() {
	var _arg0 *C.GtkMenuButton // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	C.gtk_menu_button_popup(_arg0)
	runtime.KeepAlive(menuButton)
}

// SetCreatePopupFunc sets func to be called when a popup is about to be shown.
//
// func should use one of
//
//    - gtk.MenuButton.SetPopover()
//    - gtk.MenuButton.SetMenuModel()
//
// to set a popup for menu_button. If func is non-NULL, menu_button will always
// be sensitive.
//
// Using this function will not reset the menu widget attached to menu_button.
// Instead, this can be done manually in func.
func (menuButton *MenuButton) SetCreatePopupFunc(fn MenuButtonCreatePopupFunc) {
	var _arg0 *C.GtkMenuButton               // out
	var _arg1 C.GtkMenuButtonCreatePopupFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))
	if fn != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk4_MenuButtonCreatePopupFunc)
		_arg2 = C.gpointer(gbox.Assign(fn))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_menu_button_set_create_popup_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(fn)
}

// SetDirection sets the direction in which the popup will be popped up.
//
// If the button is automatically populated with an arrow icon, its direction
// will be changed to match.
//
// If the does not fit in the available space in the given direction, GTK will
// its best to keep it inside the screen and fully visible.
//
// If you pass GTK_ARROW_NONE for a direction, the popup will behave as if you
// passed GTK_ARROW_DOWN (although you won’t see any arrows).
func (menuButton *MenuButton) SetDirection(direction ArrowType) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 C.GtkArrowType   // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))
	_arg1 = C.GtkArrowType(direction)

	C.gtk_menu_button_set_direction(_arg0, _arg1)
	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(direction)
}

// SetHasFrame sets the style of the button.
func (menuButton *MenuButton) SetHasFrame(hasFrame bool) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))
	if hasFrame {
		_arg1 = C.TRUE
	}

	C.gtk_menu_button_set_has_frame(_arg0, _arg1)
	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(hasFrame)
}

// SetIconName sets the name of an icon to show inside the menu button.
func (menuButton *MenuButton) SetIconName(iconName string) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_button_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(iconName)
}

// SetLabel sets the label to show inside the menu button.
func (menuButton *MenuButton) SetLabel(label string) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_button_set_label(_arg0, _arg1)
	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(label)
}

// SetMenuModel sets the GMenuModel from which the popup will be constructed.
//
// If menu_model is NULL, the button is disabled.
//
// A gtk.Popover will be created from the menu model with
// gtk.PopoverMenu.NewFromModel. Actions will be connected as documented for
// this function.
//
// If gtk.MenuButton:popover is already set, it will be dissociated from the
// menu_button, and the property is set to NULL.
func (menuButton *MenuButton) SetMenuModel(menuModel gio.MenuModeller) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 *C.GMenuModel    // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))
	if menuModel != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(menuModel.Native()))
	}

	C.gtk_menu_button_set_menu_model(_arg0, _arg1)
	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(menuModel)
}

// SetPopover sets the GtkPopover that will be popped up when the menu_button is
// clicked.
//
// If popover is NULL, the button is disabled.
//
// If gtk.MenuButton:menu-model is set, the menu model is dissociated from the
// menu_button, and the property is set to NULL.
func (menuButton *MenuButton) SetPopover(popover Widgetter) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))
	if popover != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(popover.Native()))
	}

	C.gtk_menu_button_set_popover(_arg0, _arg1)
	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(popover)
}

// SetUseUnderline: if true, an underline in the text indicates a mnemonic.
func (menuButton *MenuButton) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.gtk_menu_button_set_use_underline(_arg0, _arg1)
	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(useUnderline)
}
