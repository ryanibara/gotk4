// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_menu_shell_get_type()), F: marshalMenuSheller},
	})
}

// MenuShellOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type MenuShellOverrider interface {
	// The function takes the following parameters:
	//
	ActivateCurrent(forceHide bool)
	// Cancel cancels the selection within the menu shell.
	Cancel()
	// Deactivate deactivates the menu shell.
	//
	// Typically this results in the menu shell being erased from the screen.
	Deactivate()
	// The function returns the following values:
	//
	PopupDelay() int
	// Insert adds a new MenuItem to the menu shell’s item list at the position
	// indicated by position.
	//
	// The function takes the following parameters:
	//
	//    - child to add.
	//    - position in the item list where child is added. Positions are
	//      numbered from 0 to n-1.
	//
	Insert(child Widgetter, position int)
	// The function takes the following parameters:
	//
	MoveCurrent(direction MenuDirectionType)
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	MoveSelected(distance int) bool
	// SelectItem selects the menu item from the menu shell.
	//
	// The function takes the following parameters:
	//
	//    - menuItem to select.
	//
	SelectItem(menuItem Widgetter)
	SelectionDone()
}

// MenuShell is the abstract base class used to derive the Menu and MenuBar
// subclasses.
//
// A MenuShell is a container of MenuItem objects arranged in a list which can
// be navigated, selected, and activated by the user to perform application
// functions. A MenuItem can have a submenu associated with it, allowing for
// nested hierarchical menus.
//
//
// Terminology
//
// A menu item can be “selected”, this means that it is displayed in the
// prelight state, and if it has a submenu, that submenu will be popped up.
//
// A menu is “active” when it is visible onscreen and the user is selecting from
// it. A menubar is not active until the user clicks on one of its menuitems.
// When a menu is active, passing the mouse over a submenu will pop it up.
//
// There is also is a concept of the current menu and a current menu item. The
// current menu item is the selected menu item that is furthest down in the
// hierarchy. (Every active menu shell does not necessarily contain a selected
// menu item, but if it does, then the parent menu shell must also contain a
// selected menu item.) The current menu is the menu that contains the current
// menu item. It will always have a GTK grab and receive all key presses.
type MenuShell struct {
	Container

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ Containerer = (*MenuShell)(nil)
)

// MenuSheller describes types inherited from class MenuShell.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type MenuSheller interface {
	externglib.Objector
	baseMenuShell() *MenuShell
}

var _ MenuSheller = (*MenuShell)(nil)

func wrapMenuShell(obj *externglib.Object) *MenuShell {
	return &MenuShell{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
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

func marshalMenuSheller(p uintptr) (interface{}, error) {
	return wrapMenuShell(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (menuShell *MenuShell) baseMenuShell() *MenuShell {
	return menuShell
}

// BaseMenuShell returns the underlying base object.
func BaseMenuShell(obj MenuSheller) *MenuShell {
	return obj.baseMenuShell()
}

// ConnectActivateCurrent: action signal that activates the current menu item
// within the menu shell.
func (menuShell *MenuShell) ConnectActivateCurrent(f func(forceHide bool)) externglib.SignalHandle {
	return menuShell.Connect("activate-current", f)
}

// ConnectCancel: action signal which cancels the selection within the menu
// shell. Causes the MenuShell::selection-done signal to be emitted.
func (menuShell *MenuShell) ConnectCancel(f func()) externglib.SignalHandle {
	return menuShell.Connect("cancel", f)
}

// ConnectCycleFocus: keybinding signal which moves the focus in the given
// direction.
func (menuShell *MenuShell) ConnectCycleFocus(f func(direction DirectionType)) externglib.SignalHandle {
	return menuShell.Connect("cycle-focus", f)
}

// ConnectDeactivate: this signal is emitted when a menu shell is deactivated.
func (menuShell *MenuShell) ConnectDeactivate(f func()) externglib.SignalHandle {
	return menuShell.Connect("deactivate", f)
}

// ConnectInsert signal is emitted when a new MenuItem is added to a MenuShell.
// A separate signal is used instead of GtkContainer::add because of the need
// for an additional position parameter.
//
// The inverse of this signal is the GtkContainer::removed signal.
func (menuShell *MenuShell) ConnectInsert(f func(child Widgetter, position int)) externglib.SignalHandle {
	return menuShell.Connect("insert", f)
}

// ConnectMoveCurrent: keybinding signal which moves the current menu item in
// the direction specified by direction.
func (menuShell *MenuShell) ConnectMoveCurrent(f func(direction MenuDirectionType)) externglib.SignalHandle {
	return menuShell.Connect("move-current", f)
}

// ConnectMoveSelected signal is emitted to move the selection to another item.
func (menuShell *MenuShell) ConnectMoveSelected(f func(distance int) bool) externglib.SignalHandle {
	return menuShell.Connect("move-selected", f)
}

// ConnectSelectionDone: this signal is emitted when a selection has been
// completed within a menu shell.
func (menuShell *MenuShell) ConnectSelectionDone(f func()) externglib.SignalHandle {
	return menuShell.Connect("selection-done", f)
}

// ActivateItem activates the menu item within the menu shell.
//
// The function takes the following parameters:
//
//    - menuItem to activate.
//    - forceDeactivate: if TRUE, force the deactivation of the menu shell after
//      the menu item is activated.
//
func (menuShell *MenuShell) ActivateItem(menuItem Widgetter, forceDeactivate bool) {
	var _arg0 *C.GtkMenuShell // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(menuShell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(menuItem.Native()))
	if forceDeactivate {
		_arg2 = C.TRUE
	}

	C.gtk_menu_shell_activate_item(_arg0, _arg1, _arg2)
	runtime.KeepAlive(menuShell)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(forceDeactivate)
}

// Append adds a new MenuItem to the end of the menu shell's item list.
//
// The function takes the following parameters:
//
//    - child to add.
//
func (menuShell *MenuShell) Append(child *MenuItem) {
	var _arg0 *C.GtkMenuShell // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(menuShell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_menu_shell_append(_arg0, _arg1)
	runtime.KeepAlive(menuShell)
	runtime.KeepAlive(child)
}

// BindModel establishes a binding between a MenuShell and a Model.
//
// The contents of shell are removed and then refilled with menu items according
// to model. When model changes, shell is updated. Calling this function twice
// on shell with different model will cause the first binding to be replaced
// with a binding to the new model. If model is NULL then any previous binding
// is undone and all children are removed.
//
// with_separators determines if toplevel items (eg: sections) have separators
// inserted between them. This is typically desired for menus but doesn’t make
// sense for menubars.
//
// If action_namespace is non-NULL then the effect is as if all actions
// mentioned in the model have their names prefixed with the namespace, plus a
// dot. For example, if the action “quit” is mentioned and action_namespace is
// “app” then the effective action name is “app.quit”.
//
// This function uses Actionable to define the action name and target values on
// the created menu items. If you want to use an action group other than “app”
// and “win”, or if you want to use a MenuShell outside of a ApplicationWindow,
// then you will need to attach your own action group to the widget hierarchy
// using gtk_widget_insert_action_group(). As an example, if you created a group
// with a “quit” action and inserted it with the name “mygroup” then you would
// use the action name “mygroup.quit” in your Model.
//
// For most cases you are probably better off using gtk_menu_new_from_model() or
// gtk_menu_bar_new_from_model() or just directly passing the Model to
// gtk_application_set_app_menu() or gtk_application_set_menubar().
//
// The function takes the following parameters:
//
//    - model (optional) to bind to or NULL to remove binding.
//    - actionNamespace (optional): namespace for actions in model.
//    - withSeparators: TRUE if toplevel items in shell should have separators
//      between them.
//
func (menuShell *MenuShell) BindModel(model gio.MenuModeller, actionNamespace string, withSeparators bool) {
	var _arg0 *C.GtkMenuShell // out
	var _arg1 *C.GMenuModel   // out
	var _arg2 *C.gchar        // out
	var _arg3 C.gboolean      // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(menuShell.Native()))
	if model != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))
	}
	if actionNamespace != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(actionNamespace)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	if withSeparators {
		_arg3 = C.TRUE
	}

	C.gtk_menu_shell_bind_model(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(menuShell)
	runtime.KeepAlive(model)
	runtime.KeepAlive(actionNamespace)
	runtime.KeepAlive(withSeparators)
}

// Cancel cancels the selection within the menu shell.
func (menuShell *MenuShell) Cancel() {
	var _arg0 *C.GtkMenuShell // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(menuShell.Native()))

	C.gtk_menu_shell_cancel(_arg0)
	runtime.KeepAlive(menuShell)
}

// Deactivate deactivates the menu shell.
//
// Typically this results in the menu shell being erased from the screen.
func (menuShell *MenuShell) Deactivate() {
	var _arg0 *C.GtkMenuShell // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(menuShell.Native()))

	C.gtk_menu_shell_deactivate(_arg0)
	runtime.KeepAlive(menuShell)
}

// Deselect deselects the currently selected item from the menu shell, if any.
func (menuShell *MenuShell) Deselect() {
	var _arg0 *C.GtkMenuShell // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(menuShell.Native()))

	C.gtk_menu_shell_deselect(_arg0)
	runtime.KeepAlive(menuShell)
}

// ParentShell gets the parent menu shell.
//
// The parent menu shell of a submenu is the Menu or MenuBar from which it was
// opened up.
//
// The function returns the following values:
//
//    - widget: parent MenuShell.
//
func (menuShell *MenuShell) ParentShell() Widgetter {
	var _arg0 *C.GtkMenuShell // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(menuShell.Native()))

	_cret = C.gtk_menu_shell_get_parent_shell(_arg0)
	runtime.KeepAlive(menuShell)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// SelectedItem gets the currently selected item.
//
// The function returns the following values:
//
//    - widget: currently selected item.
//
func (menuShell *MenuShell) SelectedItem() Widgetter {
	var _arg0 *C.GtkMenuShell // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(menuShell.Native()))

	_cret = C.gtk_menu_shell_get_selected_item(_arg0)
	runtime.KeepAlive(menuShell)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// TakeFocus returns TRUE if the menu shell will take the keyboard focus on
// popup.
//
// The function returns the following values:
//
//    - ok: TRUE if the menu shell will take the keyboard focus on popup.
//
func (menuShell *MenuShell) TakeFocus() bool {
	var _arg0 *C.GtkMenuShell // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(menuShell.Native()))

	_cret = C.gtk_menu_shell_get_take_focus(_arg0)
	runtime.KeepAlive(menuShell)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Insert adds a new MenuItem to the menu shell’s item list at the position
// indicated by position.
//
// The function takes the following parameters:
//
//    - child to add.
//    - position in the item list where child is added. Positions are numbered
//      from 0 to n-1.
//
func (menuShell *MenuShell) Insert(child Widgetter, position int) {
	var _arg0 *C.GtkMenuShell // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gint          // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(menuShell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.gint(position)

	C.gtk_menu_shell_insert(_arg0, _arg1, _arg2)
	runtime.KeepAlive(menuShell)
	runtime.KeepAlive(child)
	runtime.KeepAlive(position)
}

// Prepend adds a new MenuItem to the beginning of the menu shell's item list.
//
// The function takes the following parameters:
//
//    - child to add.
//
func (menuShell *MenuShell) Prepend(child Widgetter) {
	var _arg0 *C.GtkMenuShell // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(menuShell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_menu_shell_prepend(_arg0, _arg1)
	runtime.KeepAlive(menuShell)
	runtime.KeepAlive(child)
}

// SelectFirst: select the first visible or selectable child of the menu shell;
// don’t select tearoff items unless the only item is a tearoff item.
//
// The function takes the following parameters:
//
//    - searchSensitive: if TRUE, search for the first selectable menu item,
//      otherwise select nothing if the first item isn’t sensitive. This should
//      be FALSE if the menu is being popped up initially.
//
func (menuShell *MenuShell) SelectFirst(searchSensitive bool) {
	var _arg0 *C.GtkMenuShell // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(menuShell.Native()))
	if searchSensitive {
		_arg1 = C.TRUE
	}

	C.gtk_menu_shell_select_first(_arg0, _arg1)
	runtime.KeepAlive(menuShell)
	runtime.KeepAlive(searchSensitive)
}

// SelectItem selects the menu item from the menu shell.
//
// The function takes the following parameters:
//
//    - menuItem to select.
//
func (menuShell *MenuShell) SelectItem(menuItem Widgetter) {
	var _arg0 *C.GtkMenuShell // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(menuShell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(menuItem.Native()))

	C.gtk_menu_shell_select_item(_arg0, _arg1)
	runtime.KeepAlive(menuShell)
	runtime.KeepAlive(menuItem)
}

// SetTakeFocus: if take_focus is TRUE (the default) the menu shell will take
// the keyboard focus so that it will receive all keyboard events which is
// needed to enable keyboard navigation in menus.
//
// Setting take_focus to FALSE is useful only for special applications like
// virtual keyboard implementations which should not take keyboard focus.
//
// The take_focus state of a menu or menu bar is automatically propagated to
// submenus whenever a submenu is popped up, so you don’t have to worry about
// recursively setting it for your entire menu hierarchy. Only when
// programmatically picking a submenu and popping it up manually, the take_focus
// property of the submenu needs to be set explicitly.
//
// Note that setting it to FALSE has side-effects:
//
// If the focus is in some other app, it keeps the focus and keynav in the menu
// doesn’t work. Consequently, keynav on the menu will only work if the focus is
// on some toplevel owned by the onscreen keyboard.
//
// To avoid confusing the user, menus with take_focus set to FALSE should not
// display mnemonics or accelerators, since it cannot be guaranteed that they
// will work.
//
// See also gdk_keyboard_grab().
//
// The function takes the following parameters:
//
//    - takeFocus: TRUE if the menu shell should take the keyboard focus on
//      popup.
//
func (menuShell *MenuShell) SetTakeFocus(takeFocus bool) {
	var _arg0 *C.GtkMenuShell // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(menuShell.Native()))
	if takeFocus {
		_arg1 = C.TRUE
	}

	C.gtk_menu_shell_set_take_focus(_arg0, _arg1)
	runtime.KeepAlive(menuShell)
	runtime.KeepAlive(takeFocus)
}
