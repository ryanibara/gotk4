// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_popover_menu_flags_get_type()), F: marshalPopoverMenuFlags},
		{T: externglib.Type(C.gtk_popover_menu_get_type()), F: marshalPopoverMenu},
	})
}

// PopoverMenuFlags flags that affect how popover menus are created from a menu
// model.
type PopoverMenuFlags int

const (
	// PopoverMenuFlagsNested: create submenus as nested popovers. Without this
	// flag, submenus are created as sliding pages that replace the main menu.
	PopoverMenuFlagsNested PopoverMenuFlags = 0b1
)

func marshalPopoverMenuFlags(p uintptr) (interface{}, error) {
	return PopoverMenuFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PopoverMenu: `GtkPopoverMenu` is a subclass of `GtkPopover` that implements
// menu behavior.
//
// !An example GtkPopoverMenu (menu.png)
//
// `GtkPopoverMenu` treats its children like menus and allows switching between
// them. It can open submenus as traditional, nested submenus, or in a more
// touch-friendly sliding fashion.
//
// `GtkPopoverMenu` is meant to be used primarily with menu models, using
// [ctor@Gtk.PopoverMenu.new_from_model]. If you need to put other widgets such
// as a `GtkSpinButton` or a `GtkSwitch` into a popover, you can use
// [method@Gtk.PopoverMenu.add_child].
//
// For more dialog-like behavior, use a plain `GtkPopover`.
//
//
// Menu models
//
// The XML format understood by `GtkBuilder` for `GMenuModel` consists of a
// toplevel `<menu>` element, which contains one or more `<item>` elements. Each
// `<item>` element contains `<attribute>` and `<link>` elements with a
// mandatory name attribute. `<link>` elements have the same content model as
// `<menu>`. Instead of `<link name="submenu>` or `<link name="section">`, you
// can use `<submenu>` or `<section>` elements.
//
// “`xml <menu id='app-menu'> <section> <item> <attribute name='label'
// translatable='yes'>_New Window</attribute> <attribute
// name='action'>app.new</attribute> </item> <item> <attribute name='label'
// translatable='yes'>_About Sunny</attribute> <attribute
// name='action'>app.about</attribute> </item> <item> <attribute name='label'
// translatable='yes'>_Quit</attribute> <attribute
// name='action'>app.quit</attribute> </item> </section> </menu> “`
//
// Attribute values can be translated using gettext, like other `GtkBuilder`
// content. `<attribute>` elements can be marked for translation with a
// `translatable="yes"` attribute. It is also possible to specify message
// context and translator comments, using the context and comments attributes.
// To make use of this, the Builder must have been given the gettext domain to
// use.
//
// The following attributes are used when constructing menu items:
//
// - "label": a user-visible string to display - "action": the prefixed name of
// the action to trigger - "target": the parameter to use when activating the
// action - "icon" and "verb-icon": names of icons that may be displayed -
// "submenu-action": name of an action that may be used to determine if a
// submenu can be opened - "hidden-when": a string used to determine when the
// item will be hidden. Possible values include "action-disabled",
// "action-missing", "macos-menubar". This is mainly useful for exported menus,
// see [method@Gtk.Application.set_menubar]. - "custom": a string used to match
// against the ID of a custom child added with
// [method@Gtk.PopoverMenu.add_child], [method@Gtk.PopoverMenuBar.add_child], or
// in the ui file with `<child type="ID">`.
//
// The following attributes are used when constructing sections:
//
// - "label": a user-visible string to use as section heading - "display-hint":
// a string used to determine special formatting for the section. Possible
// values include "horizontal-buttons", "circular-buttons" and "inline-buttons".
// They all indicate that section should be displayed as a horizontal row of
// buttons. - "text-direction": a string used to determine the
// `GtkTextDirection` to use when "display-hint" is set to "horizontal-buttons".
// Possible values include "rtl", "ltr", and "none".
//
// The following attributes are used when constructing submenus:
//
// - "label": a user-visible string to display - "icon": icon name to display
//
// Menu items will also show accelerators, which are usually associated with
// actions via [method@Gtk.Application.set_accels_for_action],
// [id@gtk_widget_class_add_binding_action] or
// [method@Gtk.ShortcutController.add_shortcut].
//
//
// CSS Nodes
//
// `GtkPopoverMenu` is just a subclass of `GtkPopover` that adds custom content
// to it, therefore it has the same CSS nodes. It is one of the cases that add a
// .menu style class to the popover's main node.
//
//
// Accessibility
//
// `GtkPopoverMenu` uses the GTK_ACCESSIBLE_ROLE_MENU role, and its items use
// the GTK_ACCESSIBLE_ROLE_MENU_ITEM, GTK_ACCESSIBLE_ROLE_MENU_ITEM_CHECKBOX or
// GTK_ACCESSIBLE_ROLE_MENU_ITEM_RADIO roles, depending on the action they are
// connected to.
type PopoverMenu interface {
	Popover

	// AsAccessible casts the class to the Accessible interface.
	AsAccessible() Accessible
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsConstraintTarget casts the class to the ConstraintTarget interface.
	AsConstraintTarget() ConstraintTarget
	// AsNative casts the class to the Native interface.
	AsNative() Native
	// AsShortcutManager casts the class to the ShortcutManager interface.
	AsShortcutManager() ShortcutManager

	// AddChildPopoverMenu adds a custom widget to a generated menu.
	//
	// For this to work, the menu model of @popover must have an item with a
	// `custom` attribute that matches @id.
	AddChildPopoverMenu(child Widget, id string) bool
	// MenuModel returns the menu model used to populate the popover.
	MenuModel() gio.MenuModel
	// RemoveChildPopoverMenu removes a widget that has previously been added
	// with gtk_popover_menu_add_child().
	RemoveChildPopoverMenu(child Widget) bool
	// SetMenuModelPopoverMenu sets a new menu model on @popover.
	//
	// The existing contents of @popover are removed, and the @popover is
	// populated with new contents according to @model.
	SetMenuModelPopoverMenu(model gio.MenuModel)
}

// popoverMenu implements the PopoverMenu class.
type popoverMenu struct {
	Popover
}

// WrapPopoverMenu wraps a GObject to the right type. It is
// primarily used internally.
func WrapPopoverMenu(obj *externglib.Object) PopoverMenu {
	return popoverMenu{
		Popover: WrapPopover(obj),
	}
}

func marshalPopoverMenu(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPopoverMenu(obj), nil
}

// NewPopoverMenuFromModel creates a `GtkPopoverMenu` and populates it according
// to @model.
//
// The created buttons are connected to actions found in the
// `GtkApplicationWindow` to which the popover belongs - typically by means of
// being attached to a widget that is contained within the
// `GtkApplicationWindow`s widget hierarchy.
//
// Actions can also be added using [method@Gtk.Widget.insert_action_group] on
// the menus attach widget or on any of its parent widgets.
//
// This function creates menus with sliding submenus. See
// [ctor@Gtk.PopoverMenu.new_from_model_full] for a way to control this.
func NewPopoverMenuFromModel(model gio.MenuModel) PopoverMenu {
	var _arg1 *C.GMenuModel // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_popover_menu_new_from_model(_arg1)

	var _popoverMenu PopoverMenu // out

	_popoverMenu = WrapPopoverMenu(externglib.Take(unsafe.Pointer(_cret)))

	return _popoverMenu
}

// NewPopoverMenuFromModelFull creates a `GtkPopoverMenu` and populates it
// according to @model.
//
// The created buttons are connected to actions found in the action groups that
// are accessible from the parent widget. This includes the
// `GtkApplicationWindow` to which the popover belongs. Actions can also be
// added using [method@Gtk.Widget.insert_action_group] on the parent widget or
// on any of its parent widgets.
//
// The only flag that is supported currently is GTK_POPOVER_MENU_NESTED, which
// makes GTK create traditional, nested submenus instead of the default sliding
// submenus.
func NewPopoverMenuFromModelFull(model gio.MenuModel, flags PopoverMenuFlags) PopoverMenu {
	var _arg1 *C.GMenuModel         // out
	var _arg2 C.GtkPopoverMenuFlags // out
	var _cret *C.GtkWidget          // in

	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))
	_arg2 = C.GtkPopoverMenuFlags(flags)

	_cret = C.gtk_popover_menu_new_from_model_full(_arg1, _arg2)

	var _popoverMenu PopoverMenu // out

	_popoverMenu = WrapPopoverMenu(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _popoverMenu
}

func (p popoverMenu) AddChildPopoverMenu(child Widget, id string) bool {
	var _arg0 *C.GtkPopoverMenu // out
	var _arg1 *C.GtkWidget      // out
	var _arg2 *C.char           // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkPopoverMenu)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.char)(C.CString(id))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_popover_menu_add_child(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p popoverMenu) MenuModel() gio.MenuModel {
	var _arg0 *C.GtkPopoverMenu // out
	var _cret *C.GMenuModel     // in

	_arg0 = (*C.GtkPopoverMenu)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_menu_get_menu_model(_arg0)

	var _menuModel gio.MenuModel // out

	_menuModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gio.MenuModel)

	return _menuModel
}

func (p popoverMenu) RemoveChildPopoverMenu(child Widget) bool {
	var _arg0 *C.GtkPopoverMenu // out
	var _arg1 *C.GtkWidget      // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkPopoverMenu)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_popover_menu_remove_child(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p popoverMenu) SetMenuModelPopoverMenu(model gio.MenuModel) {
	var _arg0 *C.GtkPopoverMenu // out
	var _arg1 *C.GMenuModel     // out

	_arg0 = (*C.GtkPopoverMenu)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	C.gtk_popover_menu_set_menu_model(_arg0, _arg1)
}

func (p popoverMenu) AsAccessible() Accessible {
	return WrapAccessible(gextras.InternObject(p))
}

func (p popoverMenu) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(p))
}

func (p popoverMenu) AsConstraintTarget() ConstraintTarget {
	return WrapConstraintTarget(gextras.InternObject(p))
}

func (p popoverMenu) AsNative() Native {
	return WrapNative(gextras.InternObject(p))
}

func (p popoverMenu) AsShortcutManager() ShortcutManager {
	return WrapShortcutManager(gextras.InternObject(p))
}
