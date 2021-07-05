// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_application_inhibit_flags_get_type()), F: marshalApplicationInhibitFlags},
		{T: externglib.Type(C.gtk_application_get_type()), F: marshalApplication},
	})
}

// ApplicationInhibitFlags types of user actions that may be blocked by
// gtk_application_inhibit().
type ApplicationInhibitFlags int

const (
	// ApplicationInhibitFlagsLogout: inhibit ending the user session by logging
	// out or by shutting down the computer
	ApplicationInhibitFlagsLogout ApplicationInhibitFlags = 0b1
	// ApplicationInhibitFlagsSwitch: inhibit user switching
	ApplicationInhibitFlagsSwitch ApplicationInhibitFlags = 0b10
	// ApplicationInhibitFlagsSuspend: inhibit suspending the session or
	// computer
	ApplicationInhibitFlagsSuspend ApplicationInhibitFlags = 0b100
	// ApplicationInhibitFlagsIdle: inhibit the session being marked as idle
	// (and possibly locked)
	ApplicationInhibitFlagsIdle ApplicationInhibitFlags = 0b1000
)

func marshalApplicationInhibitFlags(p uintptr) (interface{}, error) {
	return ApplicationInhibitFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
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
// (https://developer.gnome.org/gtk3/stable/gtk-getting-started.html#id-1.2.3.3)
type Application interface {
	gio.Application

	// AsActionGroup casts the class to the gio.ActionGroup interface.
	AsActionGroup() gio.ActionGroup

	// AddAcceleratorApplication installs an accelerator that will cause the
	// named action to be activated when the key combination specificed by
	// @accelerator is pressed.
	//
	// @accelerator must be a string that can be parsed by
	// gtk_accelerator_parse(), e.g. "<Primary>q" or “<Control><Alt>p”.
	//
	// @action_name must be the name of an action as it would be used in the app
	// menu, i.e. actions that have been added to the application are referred
	// to with an “app.” prefix, and window-specific actions with a “win.”
	// prefix.
	//
	// GtkApplication also extracts accelerators out of “accel” attributes in
	// the Models passed to gtk_application_set_app_menu() and
	// gtk_application_set_menubar(), which is usually more convenient than
	// calling this function for each accelerator.
	//
	// Deprecated: since version 3.14.
	AddAcceleratorApplication(accelerator string, actionName string, parameter glib.Variant)
	// AddWindowApplication adds a window to @application.
	//
	// This call can only happen after the @application has started; typically,
	// you should add new application windows in response to the emission of the
	// #GApplication::activate signal.
	//
	// This call is equivalent to setting the Window:application property of
	// @window to @application.
	//
	// Normally, the connection between the application and the window will
	// remain until the window is destroyed, but you can explicitly remove it
	// with gtk_application_remove_window().
	//
	// GTK+ will keep the @application running as long as it has any windows.
	AddWindowApplication(window Window)
	// AccelsForAction gets the accelerators that are currently associated with
	// the given action.
	AccelsForAction(detailedActionName string) []string
	// ActionsForAccel returns the list of actions (possibly empty) that @accel
	// maps to. Each item in the list is a detailed action name in the usual
	// form.
	//
	// This might be useful to discover if an accel already exists in order to
	// prevent installation of a conflicting accelerator (from an accelerator
	// editor or a plugin system, for example). Note that having more than one
	// action per accelerator may not be a bad thing and might make sense in
	// cases where the actions never appear in the same context.
	//
	// In case there are no actions for a given accelerator, an empty array is
	// returned. nil is never returned.
	//
	// It is a programmer error to pass an invalid accelerator string. If you
	// are unsure, check it with gtk_accelerator_parse() first.
	ActionsForAccel(accel string) []string
	// ActiveWindow gets the “active” window for the application.
	//
	// The active window is the one that was most recently focused (within the
	// application). This window may not have the focus at the moment if another
	// application has it — this is just the most recently-focused window within
	// this application.
	ActiveWindow() Window
	// AppMenu returns the menu model that has been set with
	// gtk_application_set_app_menu().
	AppMenu() gio.MenuModel
	// MenuByID gets a menu from automatically loaded resources. See [Automatic
	// resources][automatic-resources] for more information.
	MenuByID(id string) gio.Menu
	// Menubar returns the menu model that has been set with
	// gtk_application_set_menubar().
	Menubar() gio.MenuModel
	// WindowByID returns the ApplicationWindow with the given ID.
	//
	// The ID of a ApplicationWindow can be retrieved with
	// gtk_application_window_get_id().
	WindowByID(id uint) Window
	// InhibitApplication: inform the session manager that certain types of
	// actions should be inhibited. This is not guaranteed to work on all
	// platforms and for all types of actions.
	//
	// Applications should invoke this method when they begin an operation that
	// should not be interrupted, such as creating a CD or DVD. The types of
	// actions that may be blocked are specified by the @flags parameter. When
	// the application completes the operation it should call
	// gtk_application_uninhibit() to remove the inhibitor. Note that an
	// application can have multiple inhibitors, and all of them must be
	// individually removed. Inhibitors are also cleared when the application
	// exits.
	//
	// Applications should not expect that they will always be able to block the
	// action. In most cases, users will be given the option to force the action
	// to take place.
	//
	// Reasons should be short and to the point.
	//
	// If @window is given, the session manager may point the user to this
	// window to find out more about why the action is inhibited.
	InhibitApplication(window Window, flags ApplicationInhibitFlags, reason string) uint
	// IsInhibitedApplication determines if any of the actions specified in
	// @flags are currently inhibited (possibly by another application).
	//
	// Note that this information may not be available (for example when the
	// application is running in a sandbox).
	IsInhibitedApplication(flags ApplicationInhibitFlags) bool
	// ListActionDescriptionsApplication lists the detailed action names which
	// have associated accelerators. See
	// gtk_application_set_accels_for_action().
	ListActionDescriptionsApplication() []string
	// PrefersAppMenuApplication determines if the desktop environment in which
	// the application is running would prefer an application menu be shown.
	//
	// If this function returns true then the application should call
	// gtk_application_set_app_menu() with the contents of an application menu,
	// which will be shown by the desktop environment. If it returns false then
	// you should consider using an alternate approach, such as a menubar.
	//
	// The value returned by this function is purely advisory and you are free
	// to ignore it. If you call gtk_application_set_app_menu() even if the
	// desktop environment doesn't support app menus, then a fallback will be
	// provided.
	//
	// Applications are similarly free not to set an app menu even if the
	// desktop environment wants to show one. In that case, a fallback will also
	// be created by the desktop environment (GNOME, for example, uses a menu
	// with only a "Quit" item in it).
	//
	// The value returned by this function never changes. Once it returns a
	// particular value, it is guaranteed to always return the same value.
	//
	// You may only call this function after the application has been registered
	// and after the base startup handler has run. You're most likely to want to
	// use this from your own startup handler. It may also make sense to consult
	// this function while constructing UI (in activate, open or an action
	// activation handler) in order to determine if you should show a gear menu
	// or not.
	//
	// This function will return false on Mac OS and a default app menu will be
	// created automatically with the "usual" contents of that menu typical to
	// most Mac OS applications. If you call gtk_application_set_app_menu()
	// anyway, then this menu will be replaced with your own.
	PrefersAppMenuApplication() bool
	// RemoveAcceleratorApplication removes an accelerator that has been
	// previously added with gtk_application_add_accelerator().
	//
	// Deprecated: since version 3.14.
	RemoveAcceleratorApplication(actionName string, parameter glib.Variant)
	// RemoveWindowApplication: remove a window from @application.
	//
	// If @window belongs to @application then this call is equivalent to
	// setting the Window:application property of @window to nil.
	//
	// The application may stop running as a result of a call to this function.
	RemoveWindowApplication(window Window)
	// SetAccelsForActionApplication sets zero or more keyboard accelerators
	// that will trigger the given action. The first item in @accels will be the
	// primary accelerator, which may be displayed in the UI.
	//
	// To remove all accelerators for an action, use an empty, zero-terminated
	// array for @accels.
	//
	// For the @detailed_action_name, see g_action_parse_detailed_name() and
	// g_action_print_detailed_name().
	SetAccelsForActionApplication(detailedActionName string, accels []string)
	// SetAppMenuApplication sets or unsets the application menu for
	// @application.
	//
	// This can only be done in the primary instance of the application, after
	// it has been registered. #GApplication::startup is a good place to call
	// this.
	//
	// The application menu is a single menu containing items that typically
	// impact the application as a whole, rather than acting on a specific
	// window or document. For example, you would expect to see “Preferences” or
	// “Quit” in an application menu, but not “Save” or “Print”.
	//
	// If supported, the application menu will be rendered by the desktop
	// environment.
	//
	// Use the base Map interface to add actions, to respond to the user
	// selecting these menu items.
	SetAppMenuApplication(appMenu gio.MenuModel)
	// SetMenubarApplication sets or unsets the menubar for windows of
	// @application.
	//
	// This is a menubar in the traditional sense.
	//
	// This can only be done in the primary instance of the application, after
	// it has been registered. #GApplication::startup is a good place to call
	// this.
	//
	// Depending on the desktop environment, this may appear at the top of each
	// window, or at the top of the screen. In some environments, if both the
	// application menu and the menubar are set, the application menu will be
	// presented as if it were the first item of the menubar. Other environments
	// treat the two as completely separate — for example, the application menu
	// may be rendered by the desktop shell while the menubar (if set) remains
	// in each individual window.
	//
	// Use the base Map interface to add actions, to respond to the user
	// selecting these menu items.
	SetMenubarApplication(menubar gio.MenuModel)
	// UninhibitApplication removes an inhibitor that has been established with
	// gtk_application_inhibit(). Inhibitors are also cleared when the
	// application exits.
	UninhibitApplication(cookie uint)
}

// application implements the Application class.
type application struct {
	gio.Application
}

// WrapApplication wraps a GObject to the right type. It is
// primarily used internally.
func WrapApplication(obj *externglib.Object) Application {
	return application{
		Application: gio.WrapApplication(obj),
	}
}

func marshalApplication(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapApplication(obj), nil
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
// achieved by setting suitable environment variables such as `G_DEBUG`, so this
// should not be a big problem. If you absolutely must support GTK+ commandline
// arguments, you can explicitly call gtk_init() before creating the application
// instance.
//
// If non-nil, the application ID must be valid. See
// g_application_id_is_valid().
//
// If no application ID is given then some features (most notably application
// uniqueness) will be disabled. A null application ID is only allowed with GTK+
// 3.6 or later.
func NewApplication(applicationId string, flags gio.ApplicationFlags) Application {
	var _arg1 *C.gchar            // out
	var _arg2 C.GApplicationFlags // out
	var _cret *C.GtkApplication   // in

	_arg1 = (*C.gchar)(C.CString(applicationId))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GApplicationFlags(flags)

	_cret = C.gtk_application_new(_arg1, _arg2)

	var _application Application // out

	_application = WrapApplication(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _application
}

func (a application) AddAcceleratorApplication(accelerator string, actionName string, parameter glib.Variant) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.gchar          // out
	var _arg2 *C.gchar          // out
	var _arg3 *C.GVariant       // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(accelerator))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GVariant)(unsafe.Pointer(parameter))

	C.gtk_application_add_accelerator(_arg0, _arg1, _arg2, _arg3)
}

func (a application) AddWindowApplication(window Window) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.GtkWindow      // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_application_add_window(_arg0, _arg1)
}

func (a application) AccelsForAction(detailedActionName string) []string {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.gchar          // out
	var _cret **C.gchar

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(detailedActionName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_application_get_accels_for_action(_arg0, _arg1)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

func (a application) ActionsForAccel(accel string) []string {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.gchar          // out
	var _cret **C.gchar

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(accel))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_application_get_actions_for_accel(_arg0, _arg1)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

func (a application) ActiveWindow() Window {
	var _arg0 *C.GtkApplication // out
	var _cret *C.GtkWindow      // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_application_get_active_window(_arg0)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _window
}

func (a application) AppMenu() gio.MenuModel {
	var _arg0 *C.GtkApplication // out
	var _cret *C.GMenuModel     // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_application_get_app_menu(_arg0)

	var _menuModel gio.MenuModel // out

	_menuModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gio.MenuModel)

	return _menuModel
}

func (a application) MenuByID(id string) gio.Menu {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.gchar          // out
	var _cret *C.GMenu          // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(id))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_application_get_menu_by_id(_arg0, _arg1)

	var _menu gio.Menu // out

	_menu = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gio.Menu)

	return _menu
}

func (a application) Menubar() gio.MenuModel {
	var _arg0 *C.GtkApplication // out
	var _cret *C.GMenuModel     // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_application_get_menubar(_arg0)

	var _menuModel gio.MenuModel // out

	_menuModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gio.MenuModel)

	return _menuModel
}

func (a application) WindowByID(id uint) Window {
	var _arg0 *C.GtkApplication // out
	var _arg1 C.guint           // out
	var _cret *C.GtkWindow      // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = C.guint(id)

	_cret = C.gtk_application_get_window_by_id(_arg0, _arg1)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _window
}

func (a application) InhibitApplication(window Window, flags ApplicationInhibitFlags, reason string) uint {
	var _arg0 *C.GtkApplication            // out
	var _arg1 *C.GtkWindow                 // out
	var _arg2 C.GtkApplicationInhibitFlags // out
	var _arg3 *C.gchar                     // out
	var _cret C.guint                      // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = C.GtkApplicationInhibitFlags(flags)
	_arg3 = (*C.gchar)(C.CString(reason))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.gtk_application_inhibit(_arg0, _arg1, _arg2, _arg3)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (a application) IsInhibitedApplication(flags ApplicationInhibitFlags) bool {
	var _arg0 *C.GtkApplication            // out
	var _arg1 C.GtkApplicationInhibitFlags // out
	var _cret C.gboolean                   // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = C.GtkApplicationInhibitFlags(flags)

	_cret = C.gtk_application_is_inhibited(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a application) ListActionDescriptionsApplication() []string {
	var _arg0 *C.GtkApplication // out
	var _cret **C.gchar

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_application_list_action_descriptions(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

func (a application) PrefersAppMenuApplication() bool {
	var _arg0 *C.GtkApplication // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_application_prefers_app_menu(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a application) RemoveAcceleratorApplication(actionName string, parameter glib.Variant) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.gchar          // out
	var _arg2 *C.GVariant       // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GVariant)(unsafe.Pointer(parameter))

	C.gtk_application_remove_accelerator(_arg0, _arg1, _arg2)
}

func (a application) RemoveWindowApplication(window Window) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.GtkWindow      // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_application_remove_window(_arg0, _arg1)
}

func (a application) SetAccelsForActionApplication(detailedActionName string, accels []string) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.gchar          // out
	var _arg2 **C.gchar

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(detailedActionName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (**C.gchar)(C.malloc(C.ulong(len(accels)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(accels))
		for i := range accels {
			out[i] = (*C.gchar)(C.CString(accels[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_application_set_accels_for_action(_arg0, _arg1, _arg2)
}

func (a application) SetAppMenuApplication(appMenu gio.MenuModel) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.GMenuModel     // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(appMenu.Native()))

	C.gtk_application_set_app_menu(_arg0, _arg1)
}

func (a application) SetMenubarApplication(menubar gio.MenuModel) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.GMenuModel     // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(menubar.Native()))

	C.gtk_application_set_menubar(_arg0, _arg1)
}

func (a application) UninhibitApplication(cookie uint) {
	var _arg0 *C.GtkApplication // out
	var _arg1 C.guint           // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = C.guint(cookie)

	C.gtk_application_uninhibit(_arg0, _arg1)
}

func (a application) AsActionGroup() gio.ActionGroup {
	return gio.WrapActionGroup(gextras.InternObject(a))
}
