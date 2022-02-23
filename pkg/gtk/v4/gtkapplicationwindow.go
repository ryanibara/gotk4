// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// glib.Type values for gtkapplicationwindow.go.
var GTypeApplicationWindow = externglib.Type(C.gtk_application_window_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeApplicationWindow, F: marshalApplicationWindow},
	})
}

// ApplicationWindowOverrider contains methods that are overridable.
type ApplicationWindowOverrider interface {
}

// ApplicationWindow: GtkApplicationWindow is a GtkWindow subclass that
// integrates with GtkApplication.
//
// Notably, GtkApplicationWindow can handle an application menubar.
//
// This class implements the GActionGroup and GActionMap interfaces, to let you
// add window-specific actions that will be exported by the associated
// gtk.Application, together with its application-wide actions. Window-specific
// actions are prefixed with the “win.” prefix and application-wide actions are
// prefixed with the “app.” prefix. Actions must be addressed with the prefixed
// name when referring to them from a GMenuModel.
//
// Note that widgets that are placed inside a GtkApplicationWindow can also
// activate these actions, if they implement the gtk.Actionable interface.
//
// The settings gtk.Settings:gtk-shell-shows-app-menu and
// gtk.Settings:gtk-shell-shows-menubar tell GTK whether the desktop environment
// is showing the application menu and menubar models outside the application as
// part of the desktop shell. For instance, on OS X, both menus will be
// displayed remotely; on Windows neither will be.
//
// If the desktop environment does not display the menubar, then
// GtkApplicationWindow will automatically show a menubar for it. This behaviour
// can be overridden with the gtk.ApplicationWindow:show-menubar property. If
// the desktop environment does not display the application menu, then it will
// automatically be included in the menubar or in the windows client-side
// decorations.
//
// See gtk.PopoverMenu for information about the XML language used by GtkBuilder
// for menu models.
//
// See also: gtk.Application.SetMenubar().
//
//
// A GtkApplicationWindow with a menubar
//
// The code sample below shows how to set up a GtkApplicationWindow with a menu
// bar defined on the gtk.Application:
//
//    GtkApplication *app = gtk_application_new ("org.gtk.test", 0);
//
//    GtkBuilder *builder = gtk_builder_new_from_string (
//        "<interface>"
//        "  <menu id='menubar'>"
//        "    <submenu>"
//        "      <attribute name='label' translatable='yes'>_Edit</attribute>"
//        "      <item>"
//        "        <attribute name='label' translatable='yes'>_Copy</attribute>"
//        "        <attribute name='action'>win.copy</attribute>"
//        "      </item>"
//        "      <item>"
//        "        <attribute name='label' translatable='yes'>_Paste</attribute>"
//        "        <attribute name='action'>win.paste</attribute>"
//        "      </item>"
//        "    </submenu>"
//        "  </menu>"
//        "</interface>",
//        -1);
//
//    GMenuModel *menubar = G_MENU_MODEL (gtk_builder_get_object (builder, "menubar"));
//    gtk_application_set_menubar (GTK_APPLICATION (app), menubar);
//    g_object_unref (builder);
//
//    // ...
//
//    GtkWidget *window = gtk_application_window_new (app);.
type ApplicationWindow struct {
	_ [0]func() // equal guard
	Window

	*externglib.Object
	gio.ActionGroup
	gio.ActionMap
}

var (
	_ externglib.Objector = (*ApplicationWindow)(nil)
	_ Widgetter           = (*ApplicationWindow)(nil)
)

func classInitApplicationWindower(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapApplicationWindow(obj *externglib.Object) *ApplicationWindow {
	return &ApplicationWindow{
		Window: Window{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
			Object: obj,
			Root: Root{
				NativeSurface: NativeSurface{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						Object: obj,
						Accessible: Accessible{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
						ConstraintTarget: ConstraintTarget{
							Object: obj,
						},
					},
				},
			},
			ShortcutManager: ShortcutManager{
				Object: obj,
			},
		},
		Object: obj,
		ActionGroup: gio.ActionGroup{
			Object: obj,
		},
		ActionMap: gio.ActionMap{
			Object: obj,
		},
	}
}

func marshalApplicationWindow(p uintptr) (interface{}, error) {
	return wrapApplicationWindow(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewApplicationWindow creates a new GtkApplicationWindow.
//
// The function takes the following parameters:
//
//    - application: GtkApplication.
//
// The function returns the following values:
//
//    - applicationWindow: newly created GtkApplicationWindow.
//
func NewApplicationWindow(application *Application) *ApplicationWindow {
	var _arg1 *C.GtkApplication // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))

	_cret = C.gtk_application_window_new(_arg1)
	runtime.KeepAlive(application)

	var _applicationWindow *ApplicationWindow // out

	_applicationWindow = wrapApplicationWindow(externglib.Take(unsafe.Pointer(_cret)))

	return _applicationWindow
}

// HelpOverlay gets the GtkShortcutsWindow that is associated with window.
//
// See gtk.ApplicationWindow.SetHelpOverlay().
//
// The function returns the following values:
//
//    - shortcutsWindow (optional): help overlay associated with window, or NULL.
//
func (window *ApplicationWindow) HelpOverlay() *ShortcutsWindow {
	var _arg0 *C.GtkApplicationWindow // out
	var _cret *C.GtkShortcutsWindow   // in

	_arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gtk_application_window_get_help_overlay(_arg0)
	runtime.KeepAlive(window)

	var _shortcutsWindow *ShortcutsWindow // out

	if _cret != nil {
		_shortcutsWindow = wrapShortcutsWindow(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _shortcutsWindow
}

// ID returns the unique ID of the window.
//
//    If the window has not yet been added to a GtkApplication, returns 0.
//
// The function returns the following values:
//
//    - guint: unique ID for window, or 0 if the window has not yet been added to
//      a GtkApplication.
//
func (window *ApplicationWindow) ID() uint {
	var _arg0 *C.GtkApplicationWindow // out
	var _cret C.guint                 // in

	_arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gtk_application_window_get_id(_arg0)
	runtime.KeepAlive(window)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// ShowMenubar returns whether the window will display a menubar for the app
// menu and menubar as needed.
//
// The function returns the following values:
//
//    - ok: TRUE if window will display a menubar when needed.
//
func (window *ApplicationWindow) ShowMenubar() bool {
	var _arg0 *C.GtkApplicationWindow // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gtk_application_window_get_show_menubar(_arg0)
	runtime.KeepAlive(window)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetHelpOverlay associates a shortcuts window with the application window.
//
// Additionally, sets up an action with the name win.show-help-overlay to
// present it.
//
// window takes responsibility for destroying help_overlay.
//
// The function takes the following parameters:
//
//    - helpOverlay (optional): GtkShortcutsWindow.
//
func (window *ApplicationWindow) SetHelpOverlay(helpOverlay *ShortcutsWindow) {
	var _arg0 *C.GtkApplicationWindow // out
	var _arg1 *C.GtkShortcutsWindow   // out

	_arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(window.Native()))
	if helpOverlay != nil {
		_arg1 = (*C.GtkShortcutsWindow)(unsafe.Pointer(helpOverlay.Native()))
	}

	C.gtk_application_window_set_help_overlay(_arg0, _arg1)
	runtime.KeepAlive(window)
	runtime.KeepAlive(helpOverlay)
}

// SetShowMenubar sets whether the window will display a menubar for the app
// menu and menubar as needed.
//
// The function takes the following parameters:
//
//    - showMenubar: whether to show a menubar when needed.
//
func (window *ApplicationWindow) SetShowMenubar(showMenubar bool) {
	var _arg0 *C.GtkApplicationWindow // out
	var _arg1 C.gboolean              // out

	_arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(window.Native()))
	if showMenubar {
		_arg1 = C.TRUE
	}

	C.gtk_application_window_set_show_menubar(_arg0, _arg1)
	runtime.KeepAlive(window)
	runtime.KeepAlive(showMenubar)
}
