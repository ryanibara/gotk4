// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

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

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(coreglib.InternObject(menuShell).Native()))
	if model != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
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
