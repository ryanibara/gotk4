// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_application_inhibit_flags_get_type()), F: marshalApplicationInhibitFlags},
		{T: externglib.Type(C.gtk_application_get_type()), F: marshalApplicationer},
	})
}

// ApplicationInhibitFlags types of user actions that may be blocked by
// gtk_application_inhibit().
type ApplicationInhibitFlags int

const (
	// ApplicationInhibitLogout: inhibit ending the user session by logging out
	// or by shutting down the computer.
	ApplicationInhibitLogout ApplicationInhibitFlags = 0b1
	// ApplicationInhibitSwitch: inhibit user switching.
	ApplicationInhibitSwitch ApplicationInhibitFlags = 0b10
	// ApplicationInhibitSuspend: inhibit suspending the session or computer.
	ApplicationInhibitSuspend ApplicationInhibitFlags = 0b100
	// ApplicationInhibitIdle: inhibit the session being marked as idle (and
	// possibly locked).
	ApplicationInhibitIdle ApplicationInhibitFlags = 0b1000
)

func marshalApplicationInhibitFlags(p uintptr) (interface{}, error) {
	return ApplicationInhibitFlags(C.g_value_get_flags((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the names in string for ApplicationInhibitFlags.
func (a ApplicationInhibitFlags) String() string {
	if a == 0 {
		return "ApplicationInhibitFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(98)

	for a != 0 {
		next := a & (a - 1)
		bit := a - next

		switch bit {
		case ApplicationInhibitLogout:
			builder.WriteString("Logout|")
		case ApplicationInhibitSwitch:
			builder.WriteString("Switch|")
		case ApplicationInhibitSuspend:
			builder.WriteString("Suspend|")
		case ApplicationInhibitIdle:
			builder.WriteString("Idle|")
		default:
			builder.WriteString(fmt.Sprintf("ApplicationInhibitFlags(0b%b)|", bit))
		}

		a = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if a contains other.
func (a ApplicationInhibitFlags) Has(other ApplicationInhibitFlags) bool {
	return (a & other) == other
}

// ApplicationOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ApplicationOverrider interface {
	WindowAdded(window *Window)
	WindowRemoved(window *Window)
}

// Application is a class that handles many important aspects of a GTK+
// application in a convenient fashion, without enforcing a one-size-fits-all
// application model.
//
// Currently, GtkApplication handles GTK+ initialization, application
// uniqueness, session management, provides some basic scriptability and desktop
// shell integration by exporting actions and menus and manages a list of
// toplevel windows whose life-cycle is automatically tied to the life-cycle of
// your application.
//
// While GtkApplication works fine with plain Windows, it is recommended to use
// it together with ApplicationWindow.
//
// When GDK threads are enabled, GtkApplication will acquire the GDK lock when
// invoking actions that arrive from other processes. The GDK lock is not
// touched for local action invocations. In order to have actions invoked in a
// predictable context it is therefore recommended that the GDK lock be held
// while invoking actions locally with g_action_group_activate_action(). The
// same applies to actions associated with ApplicationWindow and to the
// “activate” and “open” #GApplication methods.
//
//
// Automatic resources
//
// Application will automatically load menus from the Builder resource located
// at "gtk/menus.ui", relative to the application's resource base path (see
// g_application_set_resource_base_path()). The menu with the ID "app-menu" is
// taken as the application's app menu and the menu with the ID "menubar" is
// taken as the application's menubar. Additional menus (most interesting
// submenus) can be named and accessed via gtk_application_get_menu_by_id()
// which allows for dynamic population of a part of the menu structure.
//
// If the resources "gtk/menus-appmenu.ui" or "gtk/menus-traditional.ui" are
// present then these files will be used in preference, depending on the value
// of gtk_application_prefers_app_menu(). If the resource "gtk/menus-common.ui"
// is present it will be loaded as well. This is useful for storing items that
// are referenced from both "gtk/menus-appmenu.ui" and
// "gtk/menus-traditional.ui".
//
// It is also possible to provide the menus manually using
// gtk_application_set_app_menu() and gtk_application_set_menubar().
//
// Application will also automatically setup an icon search path for the default
// icon theme by appending "icons" to the resource base path. This allows your
// application to easily store its icons as resources. See
// gtk_icon_theme_add_resource_path() for more information.
//
// If there is a resource located at "gtk/help-overlay.ui" which defines a
// ShortcutsWindow with ID "help_overlay" then GtkApplication associates an
// instance of this shortcuts window with each ApplicationWindow and sets up
// keyboard accelerators (Control-F1 and Control-?) to open it. To create a menu
// item that displays the shortcuts window, associate the item with the action
// win.show-help-overlay.
//
//
// A simple application
//
// A simple example
// (https://git.gnome.org/browse/gtk+/tree/examples/bp/bloatpad.c)
//
// GtkApplication optionally registers with a session manager of the users
// session (if you set the Application:register-session property) and offers
// various functionality related to the session life-cycle.
//
// An application can block various ways to end the session with the
// gtk_application_inhibit() function. Typical use cases for this kind of
// inhibiting are long-running, uninterruptible operations, such as burning a CD
// or performing a disk backup. The session manager may not honor the inhibitor,
// but it can be expected to inform the user about the negative consequences of
// ending the session while inhibitors are present.
//
//
// See Also
//
// HowDoI: Using GtkApplication (https://wiki.gnome.org/HowDoI/GtkApplication),
// Getting Started with GTK+: Basics
// (https://developer.gnome.org/gtk3/stable/gtk-getting-started.html#id-1.2.3.3).
type Application struct {
	gio.Application
}

func wrapApplication(obj *externglib.Object) *Application {
	return &Application{
		Application: gio.Application{
			Object: obj,
			ActionGroup: gio.ActionGroup{
				Object: obj,
			},
			ActionMap: gio.ActionMap{
				Object: obj,
			},
		},
	}
}

func marshalApplicationer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapApplication(obj), nil
}

// NewApplication creates a new Application instance.
//
// When using Application, it is not necessary to call gtk_init() manually. It
// is called as soon as the application gets registered as the primary instance.
//
// Concretely, gtk_init() is called in the default handler for the
// #GApplication::startup signal. Therefore, Application subclasses should chain
// up in their #GApplication::startup handler before using any GTK+ API.
//
// Note that commandline arguments are not passed to gtk_init(). All GTK+
// functionality that is available via commandline arguments can also be
// achieved by setting suitable environment variables such as G_DEBUG, so this
// should not be a big problem. If you absolutely must support GTK+ commandline
// arguments, you can explicitly call gtk_init() before creating the application
// instance.
//
// If non-NULL, the application ID must be valid. See
// g_application_id_is_valid().
//
// If no application ID is given then some features (most notably application
// uniqueness) will be disabled. A null application ID is only allowed with GTK+
// 3.6 or later.
//
// The function takes the following parameters:
//
//    - applicationId: application ID.
//    - flags: application flags.
//
func NewApplication(applicationId string, flags gio.ApplicationFlags) *Application {
	var _arg1 *C.gchar            // out
	var _arg2 C.GApplicationFlags // out
	var _cret *C.GtkApplication   // in

	if applicationId != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(applicationId)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = C.GApplicationFlags(flags)

	_cret = C.gtk_application_new(_arg1, _arg2)
	runtime.KeepAlive(applicationId)
	runtime.KeepAlive(flags)

	var _application *Application // out

	_application = wrapApplication(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _application
}

// AddAccelerator installs an accelerator that will cause the named action to be
// activated when the key combination specificed by accelerator is pressed.
//
// accelerator must be a string that can be parsed by gtk_accelerator_parse(),
// e.g. "<Primary>q" or “<Control><Alt>p”.
//
// action_name must be the name of an action as it would be used in the app
// menu, i.e. actions that have been added to the application are referred to
// with an “app.” prefix, and window-specific actions with a “win.” prefix.
//
// GtkApplication also extracts accelerators out of “accel” attributes in the
// Models passed to gtk_application_set_app_menu() and
// gtk_application_set_menubar(), which is usually more convenient than calling
// this function for each accelerator.
//
// Deprecated: Use gtk_application_set_accels_for_action() instead.
//
// The function takes the following parameters:
//
//    - accelerator string.
//    - actionName: name of the action to activate.
//    - parameter to pass when activating the action, or NULL if the action
//    does not accept an activation parameter.
//
func (application *Application) AddAccelerator(accelerator, actionName string, parameter *glib.Variant) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.gchar          // out
	var _arg2 *C.gchar          // out
	var _arg3 *C.GVariant       // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelerator)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg2))
	if parameter != nil {
		_arg3 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(parameter)))
	}

	C.gtk_application_add_accelerator(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(application)
	runtime.KeepAlive(accelerator)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(parameter)
}

// AddWindow adds a window to application.
//
// This call can only happen after the application has started; typically, you
// should add new application windows in response to the emission of the
// #GApplication::activate signal.
//
// This call is equivalent to setting the Window:application property of window
// to application.
//
// Normally, the connection between the application and the window will remain
// until the window is destroyed, but you can explicitly remove it with
// gtk_application_remove_window().
//
// GTK+ will keep the application running as long as it has any windows.
//
// The function takes the following parameters:
//
//    - window: Window.
//
func (application *Application) AddWindow(window *Window) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.GtkWindow      // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_application_add_window(_arg0, _arg1)
	runtime.KeepAlive(application)
	runtime.KeepAlive(window)
}

// AccelsForAction gets the accelerators that are currently associated with the
// given action.
//
// The function takes the following parameters:
//
//    - detailedActionName: detailed action name, specifying an action and
//    target to obtain accelerators for.
//
func (application *Application) AccelsForAction(detailedActionName string) []string {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.gchar          // out
	var _cret **C.gchar         // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(detailedActionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_application_get_accels_for_action(_arg0, _arg1)
	runtime.KeepAlive(application)
	runtime.KeepAlive(detailedActionName)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// ActionsForAccel returns the list of actions (possibly empty) that accel maps
// to. Each item in the list is a detailed action name in the usual form.
//
// This might be useful to discover if an accel already exists in order to
// prevent installation of a conflicting accelerator (from an accelerator editor
// or a plugin system, for example). Note that having more than one action per
// accelerator may not be a bad thing and might make sense in cases where the
// actions never appear in the same context.
//
// In case there are no actions for a given accelerator, an empty array is
// returned. NULL is never returned.
//
// It is a programmer error to pass an invalid accelerator string. If you are
// unsure, check it with gtk_accelerator_parse() first.
//
// The function takes the following parameters:
//
//    - accel: accelerator that can be parsed by gtk_accelerator_parse().
//
func (application *Application) ActionsForAccel(accel string) []string {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.gchar          // out
	var _cret **C.gchar         // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accel)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_application_get_actions_for_accel(_arg0, _arg1)
	runtime.KeepAlive(application)
	runtime.KeepAlive(accel)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// ActiveWindow gets the “active” window for the application.
//
// The active window is the one that was most recently focused (within the
// application). This window may not have the focus at the moment if another
// application has it — this is just the most recently-focused window within
// this application.
func (application *Application) ActiveWindow() *Window {
	var _arg0 *C.GtkApplication // out
	var _cret *C.GtkWindow      // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))

	_cret = C.gtk_application_get_active_window(_arg0)
	runtime.KeepAlive(application)

	var _window *Window // out

	if _cret != nil {
		_window = wrapWindow(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _window
}

// AppMenu returns the menu model that has been set with
// gtk_application_set_app_menu().
func (application *Application) AppMenu() gio.MenuModeller {
	var _arg0 *C.GtkApplication // out
	var _cret *C.GMenuModel     // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))

	_cret = C.gtk_application_get_app_menu(_arg0)
	runtime.KeepAlive(application)

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

// MenuByID gets a menu from automatically loaded resources. See [Automatic
// resources][automatic-resources] for more information.
//
// The function takes the following parameters:
//
//    - id of the menu to look up.
//
func (application *Application) MenuByID(id string) *gio.Menu {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.gchar          // out
	var _cret *C.GMenu          // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_application_get_menu_by_id(_arg0, _arg1)
	runtime.KeepAlive(application)
	runtime.KeepAlive(id)

	var _menu *gio.Menu // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_menu = &gio.Menu{
			MenuModel: gio.MenuModel{
				Object: obj,
			},
		}
	}

	return _menu
}

// Menubar returns the menu model that has been set with
// gtk_application_set_menubar().
func (application *Application) Menubar() gio.MenuModeller {
	var _arg0 *C.GtkApplication // out
	var _cret *C.GMenuModel     // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))

	_cret = C.gtk_application_get_menubar(_arg0)
	runtime.KeepAlive(application)

	var _menuModel gio.MenuModeller // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.MenuModeller is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(gio.MenuModeller)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.MenuModeller")
		}
		_menuModel = rv
	}

	return _menuModel
}

// WindowByID returns the ApplicationWindow with the given ID.
//
// The ID of a ApplicationWindow can be retrieved with
// gtk_application_window_get_id().
//
// The function takes the following parameters:
//
//    - id: identifier number.
//
func (application *Application) WindowByID(id uint) *Window {
	var _arg0 *C.GtkApplication // out
	var _arg1 C.guint           // out
	var _cret *C.GtkWindow      // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))
	_arg1 = C.guint(id)

	_cret = C.gtk_application_get_window_by_id(_arg0, _arg1)
	runtime.KeepAlive(application)
	runtime.KeepAlive(id)

	var _window *Window // out

	if _cret != nil {
		_window = wrapWindow(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _window
}

// Windows gets a list of the Windows associated with application.
//
// The list is sorted by most recently focused window, such that the first
// element is the currently focused window. (Useful for choosing a parent for a
// transient window.)
//
// The list that is returned should not be modified in any way. It will only
// remain valid until the next focus change or window creation or deletion.
func (application *Application) Windows() []Window {
	var _arg0 *C.GtkApplication // out
	var _cret *C.GList          // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))

	_cret = C.gtk_application_get_windows(_arg0)
	runtime.KeepAlive(application)

	var _list []Window // out

	_list = make([]Window, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.GtkWindow)(v)
		var dst Window // out
		dst = *wrapWindow(externglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// Inhibit: inform the session manager that certain types of actions should be
// inhibited. This is not guaranteed to work on all platforms and for all types
// of actions.
//
// Applications should invoke this method when they begin an operation that
// should not be interrupted, such as creating a CD or DVD. The types of actions
// that may be blocked are specified by the flags parameter. When the
// application completes the operation it should call
// gtk_application_uninhibit() to remove the inhibitor. Note that an application
// can have multiple inhibitors, and all of them must be individually removed.
// Inhibitors are also cleared when the application exits.
//
// Applications should not expect that they will always be able to block the
// action. In most cases, users will be given the option to force the action to
// take place.
//
// Reasons should be short and to the point.
//
// If window is given, the session manager may point the user to this window to
// find out more about why the action is inhibited.
//
// The function takes the following parameters:
//
//    - window or NULL.
//    - flags: what types of actions should be inhibited.
//    - reason: short, human-readable string that explains why these operations
//    are inhibited.
//
func (application *Application) Inhibit(window *Window, flags ApplicationInhibitFlags, reason string) uint {
	var _arg0 *C.GtkApplication            // out
	var _arg1 *C.GtkWindow                 // out
	var _arg2 C.GtkApplicationInhibitFlags // out
	var _arg3 *C.gchar                     // out
	var _cret C.guint                      // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))
	if window != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))
	}
	_arg2 = C.GtkApplicationInhibitFlags(flags)
	if reason != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(reason)))
		defer C.free(unsafe.Pointer(_arg3))
	}

	_cret = C.gtk_application_inhibit(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(application)
	runtime.KeepAlive(window)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(reason)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// IsInhibited determines if any of the actions specified in flags are currently
// inhibited (possibly by another application).
//
// Note that this information may not be available (for example when the
// application is running in a sandbox).
//
// The function takes the following parameters:
//
//    - flags: what types of actions should be queried.
//
func (application *Application) IsInhibited(flags ApplicationInhibitFlags) bool {
	var _arg0 *C.GtkApplication            // out
	var _arg1 C.GtkApplicationInhibitFlags // out
	var _cret C.gboolean                   // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))
	_arg1 = C.GtkApplicationInhibitFlags(flags)

	_cret = C.gtk_application_is_inhibited(_arg0, _arg1)
	runtime.KeepAlive(application)
	runtime.KeepAlive(flags)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ListActionDescriptions lists the detailed action names which have associated
// accelerators. See gtk_application_set_accels_for_action().
func (application *Application) ListActionDescriptions() []string {
	var _arg0 *C.GtkApplication // out
	var _cret **C.gchar         // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))

	_cret = C.gtk_application_list_action_descriptions(_arg0)
	runtime.KeepAlive(application)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// PrefersAppMenu determines if the desktop environment in which the application
// is running would prefer an application menu be shown.
//
// If this function returns TRUE then the application should call
// gtk_application_set_app_menu() with the contents of an application menu,
// which will be shown by the desktop environment. If it returns FALSE then you
// should consider using an alternate approach, such as a menubar.
//
// The value returned by this function is purely advisory and you are free to
// ignore it. If you call gtk_application_set_app_menu() even if the desktop
// environment doesn't support app menus, then a fallback will be provided.
//
// Applications are similarly free not to set an app menu even if the desktop
// environment wants to show one. In that case, a fallback will also be created
// by the desktop environment (GNOME, for example, uses a menu with only a
// "Quit" item in it).
//
// The value returned by this function never changes. Once it returns a
// particular value, it is guaranteed to always return the same value.
//
// You may only call this function after the application has been registered and
// after the base startup handler has run. You're most likely to want to use
// this from your own startup handler. It may also make sense to consult this
// function while constructing UI (in activate, open or an action activation
// handler) in order to determine if you should show a gear menu or not.
//
// This function will return FALSE on Mac OS and a default app menu will be
// created automatically with the "usual" contents of that menu typical to most
// Mac OS applications. If you call gtk_application_set_app_menu() anyway, then
// this menu will be replaced with your own.
func (application *Application) PrefersAppMenu() bool {
	var _arg0 *C.GtkApplication // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))

	_cret = C.gtk_application_prefers_app_menu(_arg0)
	runtime.KeepAlive(application)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RemoveAccelerator removes an accelerator that has been previously added with
// gtk_application_add_accelerator().
//
// Deprecated: Use gtk_application_set_accels_for_action() instead.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to activate.
//    - parameter to pass when activating the action, or NULL if the action
//    does not accept an activation parameter.
//
func (application *Application) RemoveAccelerator(actionName string, parameter *glib.Variant) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.gchar          // out
	var _arg2 *C.GVariant       // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	if parameter != nil {
		_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(parameter)))
	}

	C.gtk_application_remove_accelerator(_arg0, _arg1, _arg2)
	runtime.KeepAlive(application)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(parameter)
}

// RemoveWindow: remove a window from application.
//
// If window belongs to application then this call is equivalent to setting the
// Window:application property of window to NULL.
//
// The application may stop running as a result of a call to this function.
//
// The function takes the following parameters:
//
//    - window: Window.
//
func (application *Application) RemoveWindow(window *Window) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.GtkWindow      // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_application_remove_window(_arg0, _arg1)
	runtime.KeepAlive(application)
	runtime.KeepAlive(window)
}

// SetAccelsForAction sets zero or more keyboard accelerators that will trigger
// the given action. The first item in accels will be the primary accelerator,
// which may be displayed in the UI.
//
// To remove all accelerators for an action, use an empty, zero-terminated array
// for accels.
//
// For the detailed_action_name, see g_action_parse_detailed_name() and
// g_action_print_detailed_name().
//
// The function takes the following parameters:
//
//    - detailedActionName: detailed action name, specifying an action and
//    target to associate accelerators with.
//    - accels: list of accelerators in the format understood by
//    gtk_accelerator_parse().
//
func (application *Application) SetAccelsForAction(detailedActionName string, accels []string) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.gchar          // out
	var _arg2 **C.gchar         // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(detailedActionName)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		_arg2 = (**C.gchar)(C.malloc(C.ulong(len(accels)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(accels)+1)
			var zero *C.gchar
			out[len(accels)] = zero
			for i := range accels {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(accels[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.gtk_application_set_accels_for_action(_arg0, _arg1, _arg2)
	runtime.KeepAlive(application)
	runtime.KeepAlive(detailedActionName)
	runtime.KeepAlive(accels)
}

// SetAppMenu sets or unsets the application menu for application.
//
// This can only be done in the primary instance of the application, after it
// has been registered. #GApplication::startup is a good place to call this.
//
// The application menu is a single menu containing items that typically impact
// the application as a whole, rather than acting on a specific window or
// document. For example, you would expect to see “Preferences” or “Quit” in an
// application menu, but not “Save” or “Print”.
//
// If supported, the application menu will be rendered by the desktop
// environment.
//
// Use the base Map interface to add actions, to respond to the user selecting
// these menu items.
//
// The function takes the following parameters:
//
//    - appMenu or NULL.
//
func (application *Application) SetAppMenu(appMenu gio.MenuModeller) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.GMenuModel     // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))
	if appMenu != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(appMenu.Native()))
	}

	C.gtk_application_set_app_menu(_arg0, _arg1)
	runtime.KeepAlive(application)
	runtime.KeepAlive(appMenu)
}

// SetMenubar sets or unsets the menubar for windows of application.
//
// This is a menubar in the traditional sense.
//
// This can only be done in the primary instance of the application, after it
// has been registered. #GApplication::startup is a good place to call this.
//
// Depending on the desktop environment, this may appear at the top of each
// window, or at the top of the screen. In some environments, if both the
// application menu and the menubar are set, the application menu will be
// presented as if it were the first item of the menubar. Other environments
// treat the two as completely separate — for example, the application menu may
// be rendered by the desktop shell while the menubar (if set) remains in each
// individual window.
//
// Use the base Map interface to add actions, to respond to the user selecting
// these menu items.
//
// The function takes the following parameters:
//
//    - menubar or NULL.
//
func (application *Application) SetMenubar(menubar gio.MenuModeller) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.GMenuModel     // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))
	if menubar != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(menubar.Native()))
	}

	C.gtk_application_set_menubar(_arg0, _arg1)
	runtime.KeepAlive(application)
	runtime.KeepAlive(menubar)
}

// Uninhibit removes an inhibitor that has been established with
// gtk_application_inhibit(). Inhibitors are also cleared when the application
// exits.
//
// The function takes the following parameters:
//
//    - cookie that was returned by gtk_application_inhibit().
//
func (application *Application) Uninhibit(cookie uint) {
	var _arg0 *C.GtkApplication // out
	var _arg1 C.guint           // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))
	_arg1 = C.guint(cookie)

	C.gtk_application_uninhibit(_arg0, _arg1)
	runtime.KeepAlive(application)
	runtime.KeepAlive(cookie)
}

// ConnectQueryEnd: emitted when the session manager is about to end the
// session, only if Application::register-session is TRUE. Applications can
// connect to this signal and call gtk_application_inhibit() with
// GTK_APPLICATION_INHIBIT_LOGOUT to delay the end of the session until state
// has been saved.
func (application *Application) ConnectQueryEnd(f func()) externglib.SignalHandle {
	return application.Connect("query-end", f)
}

// ConnectWindowAdded: emitted when a Window is added to application through
// gtk_application_add_window().
func (application *Application) ConnectWindowAdded(f func(window Window)) externglib.SignalHandle {
	return application.Connect("window-added", f)
}

// ConnectWindowRemoved: emitted when a Window is removed from application,
// either as a side-effect of being destroyed or explicitly through
// gtk_application_remove_window().
func (application *Application) ConnectWindowRemoved(f func(window Window)) externglib.SignalHandle {
	return application.Connect("window-removed", f)
}
