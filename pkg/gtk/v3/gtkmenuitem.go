// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_menu_item_get_type()), F: marshalMenuItemmer},
	})
}

// MenuItemOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type MenuItemOverrider interface {
	// Activate emits the MenuItem::activate signal on the given item
	Activate()
	//
	ActivateItem()
	// Deselect emits the MenuItem::deselect signal on the given item.
	Deselect()
	// Label sets @text on the @menu_item label
	Label() string
	// Select emits the MenuItem::select signal on the given item.
	Select()
	// SetLabel sets @text on the @menu_item label
	SetLabel(label string)
	// ToggleSizeAllocate emits the MenuItem::toggle-size-allocate signal on the
	// given item.
	ToggleSizeAllocate(allocation int)
}

// MenuItemmer describes MenuItem's methods.
type MenuItemmer interface {
	// Activate emits the MenuItem::activate signal on the given item
	Activate()
	// Deselect emits the MenuItem::deselect signal on the given item.
	Deselect()
	// AccelPath: retrieve the accelerator path that was previously set on
	// @menu_item.
	AccelPath() string
	// Label sets @text on the @menu_item label
	Label() string
	// ReserveIndicator returns whether the @menu_item reserves space for the
	// submenu indicator, regardless if it has a submenu or not.
	ReserveIndicator() bool
	// RightJustified gets whether the menu item appears justified at the right
	// side of the menu bar.
	RightJustified() bool
	// Submenu gets the submenu underneath this menu item, if any.
	Submenu() *Widget
	// UseUnderline checks if an underline in the text indicates the next
	// character should be used for the mnemonic accelerator key.
	UseUnderline() bool
	// Select emits the MenuItem::select signal on the given item.
	Select()
	// SetAccelPath: set the accelerator path on @menu_item, through which
	// runtime changes of the menu item’s accelerator caused by the user can be
	// identified and saved to persistent storage (see gtk_accel_map_save() on
	// this).
	SetAccelPath(accelPath string)
	// SetLabel sets @text on the @menu_item label
	SetLabel(label string)
	// SetReserveIndicator sets whether the @menu_item should reserve space for
	// the submenu indicator, regardless if it actually has a submenu or not.
	SetReserveIndicator(reserve bool)
	// SetRightJustified sets whether the menu item appears justified at the
	// right side of a menu bar.
	SetRightJustified(rightJustified bool)
	// SetSubmenu sets or replaces the menu item’s submenu, or removes it when a
	// nil submenu is passed.
	SetSubmenu(submenu Menuer)
	// SetUseUnderline: if true, an underline in the text indicates the next
	// character should be used for the mnemonic accelerator key.
	SetUseUnderline(setting bool)
	// ToggleSizeAllocate emits the MenuItem::toggle-size-allocate signal on the
	// given item.
	ToggleSizeAllocate(allocation int)
}

// MenuItem widget and the derived widgets are the only valid children for
// menus. Their function is to correctly handle highlighting, alignment, events
// and submenus.
//
// As a GtkMenuItem derives from Bin it can hold any valid child widget,
// although only a few are really useful.
//
// By default, a GtkMenuItem sets a AccelLabel as its child. GtkMenuItem has
// direct functions to set the label and its mnemonic. For more advanced label
// settings, you can fetch the child widget from the GtkBin.
//
// An example for setting markup and accelerator on a MenuItem:
//
//    menuitem
//    ├── <child>
//    ╰── [arrow.right]
//
// GtkMenuItem has a single CSS node with name menuitem. If the menuitem has a
// submenu, it gets another CSS node with name arrow, which has the .left or
// .right style class.
type MenuItem struct {
	Bin

	Actionable
	Activatable
}

var (
	_ MenuItemmer     = (*MenuItem)(nil)
	_ gextras.Nativer = (*MenuItem)(nil)
)

func wrapMenuItem(obj *externglib.Object) MenuItemmer {
	return &MenuItem{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
		},
		Actionable: Actionable{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
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

func marshalMenuItemmer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMenuItem(obj), nil
}

// NewMenuItem creates a new MenuItem.
func NewMenuItem() *MenuItem {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_menu_item_new()

	var _menuItem *MenuItem // out

	_menuItem = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*MenuItem)

	return _menuItem
}

// NewMenuItemWithLabel creates a new MenuItem whose child is a Label.
func NewMenuItemWithLabel(label string) *MenuItem {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_menu_item_new_with_label(_arg1)

	var _menuItem *MenuItem // out

	_menuItem = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*MenuItem)

	return _menuItem
}

// NewMenuItemWithMnemonic creates a new MenuItem containing a label.
//
// The label will be created using gtk_label_new_with_mnemonic(), so underscores
// in @label indicate the mnemonic for the menu item.
func NewMenuItemWithMnemonic(label string) *MenuItem {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_menu_item_new_with_mnemonic(_arg1)

	var _menuItem *MenuItem // out

	_menuItem = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*MenuItem)

	return _menuItem
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *MenuItem) Native() uintptr {
	return v.Bin.Container.Widget.InitiallyUnowned.Object.Native()
}

// Activate emits the MenuItem::activate signal on the given item
func (menuItem *MenuItem) Activate() {
	var _arg0 *C.GtkMenuItem // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(menuItem.Native()))

	C.gtk_menu_item_activate(_arg0)
}

// Deselect emits the MenuItem::deselect signal on the given item.
func (menuItem *MenuItem) Deselect() {
	var _arg0 *C.GtkMenuItem // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(menuItem.Native()))

	C.gtk_menu_item_deselect(_arg0)
}

// AccelPath: retrieve the accelerator path that was previously set on
// @menu_item.
//
// See gtk_menu_item_set_accel_path() for details.
func (menuItem *MenuItem) AccelPath() string {
	var _arg0 *C.GtkMenuItem // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(menuItem.Native()))

	_cret = C.gtk_menu_item_get_accel_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Label sets @text on the @menu_item label
func (menuItem *MenuItem) Label() string {
	var _arg0 *C.GtkMenuItem // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(menuItem.Native()))

	_cret = C.gtk_menu_item_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ReserveIndicator returns whether the @menu_item reserves space for the
// submenu indicator, regardless if it has a submenu or not.
func (menuItem *MenuItem) ReserveIndicator() bool {
	var _arg0 *C.GtkMenuItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(menuItem.Native()))

	_cret = C.gtk_menu_item_get_reserve_indicator(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RightJustified gets whether the menu item appears justified at the right side
// of the menu bar.
//
// Deprecated: See gtk_menu_item_set_right_justified().
func (menuItem *MenuItem) RightJustified() bool {
	var _arg0 *C.GtkMenuItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(menuItem.Native()))

	_cret = C.gtk_menu_item_get_right_justified(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Submenu gets the submenu underneath this menu item, if any. See
// gtk_menu_item_set_submenu().
func (menuItem *MenuItem) Submenu() *Widget {
	var _arg0 *C.GtkMenuItem // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(menuItem.Native()))

	_cret = C.gtk_menu_item_get_submenu(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// UseUnderline checks if an underline in the text indicates the next character
// should be used for the mnemonic accelerator key.
func (menuItem *MenuItem) UseUnderline() bool {
	var _arg0 *C.GtkMenuItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(menuItem.Native()))

	_cret = C.gtk_menu_item_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Select emits the MenuItem::select signal on the given item.
func (menuItem *MenuItem) Select() {
	var _arg0 *C.GtkMenuItem // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(menuItem.Native()))

	C.gtk_menu_item_select(_arg0)
}

// SetAccelPath: set the accelerator path on @menu_item, through which runtime
// changes of the menu item’s accelerator caused by the user can be identified
// and saved to persistent storage (see gtk_accel_map_save() on this). To set up
// a default accelerator for this menu item, call gtk_accel_map_add_entry() with
// the same @accel_path. See also gtk_accel_map_add_entry() on the specifics of
// accelerator paths, and gtk_menu_set_accel_path() for a more convenient
// variant of this function.
//
// This function is basically a convenience wrapper that handles calling
// gtk_widget_set_accel_path() with the appropriate accelerator group for the
// menu item.
//
// Note that you do need to set an accelerator on the parent menu with
// gtk_menu_set_accel_group() for this to work.
//
// Note that @accel_path string will be stored in a #GQuark. Therefore, if you
// pass a static string, you can save some memory by interning it first with
// g_intern_static_string().
func (menuItem *MenuItem) SetAccelPath(accelPath string) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(menuItem.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelPath)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_item_set_accel_path(_arg0, _arg1)
}

// SetLabel sets @text on the @menu_item label
func (menuItem *MenuItem) SetLabel(label string) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(menuItem.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_item_set_label(_arg0, _arg1)
}

// SetReserveIndicator sets whether the @menu_item should reserve space for the
// submenu indicator, regardless if it actually has a submenu or not.
//
// There should be little need for applications to call this functions.
func (menuItem *MenuItem) SetReserveIndicator(reserve bool) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(menuItem.Native()))
	if reserve {
		_arg1 = C.TRUE
	}

	C.gtk_menu_item_set_reserve_indicator(_arg0, _arg1)
}

// SetRightJustified sets whether the menu item appears justified at the right
// side of a menu bar. This was traditionally done for “Help” menu items, but is
// now considered a bad idea. (If the widget layout is reversed for a
// right-to-left language like Hebrew or Arabic, right-justified-menu-items
// appear at the left.)
//
// Deprecated: If you insist on using it, use gtk_widget_set_hexpand() and
// gtk_widget_set_halign().
func (menuItem *MenuItem) SetRightJustified(rightJustified bool) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(menuItem.Native()))
	if rightJustified {
		_arg1 = C.TRUE
	}

	C.gtk_menu_item_set_right_justified(_arg0, _arg1)
}

// SetSubmenu sets or replaces the menu item’s submenu, or removes it when a nil
// submenu is passed.
func (menuItem *MenuItem) SetSubmenu(submenu Menuer) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(menuItem.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((submenu).(gextras.Nativer).Native()))

	C.gtk_menu_item_set_submenu(_arg0, _arg1)
}

// SetUseUnderline: if true, an underline in the text indicates the next
// character should be used for the mnemonic accelerator key.
func (menuItem *MenuItem) SetUseUnderline(setting bool) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(menuItem.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_menu_item_set_use_underline(_arg0, _arg1)
}

// ToggleSizeAllocate emits the MenuItem::toggle-size-allocate signal on the
// given item.
func (menuItem *MenuItem) ToggleSizeAllocate(allocation int) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(menuItem.Native()))
	_arg1 = C.gint(allocation)

	C.gtk_menu_item_toggle_size_allocate(_arg0, _arg1)
}
