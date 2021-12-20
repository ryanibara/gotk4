// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
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
		{T: externglib.Type(C.gtk_ui_manager_item_type_get_type()), F: marshalUIManagerItemType},
		{T: externglib.Type(C.gtk_ui_manager_get_type()), F: marshalUIManagerer},
	})
}

// UIManagerItemType: these enumeration values are used by
// gtk_ui_manager_add_ui() to determine what UI element to create.
//
// Deprecated: since version 3.10.
type UIManagerItemType C.guint

const (
	// UiManagerAuto: pick the type of the UI element according to context.
	UiManagerAuto UIManagerItemType = 0b0
	// UiManagerMenubar: create a menubar.
	UiManagerMenubar UIManagerItemType = 0b1
	// UiManagerMenu: create a menu.
	UiManagerMenu UIManagerItemType = 0b10
	// UiManagerToolbar: create a toolbar.
	UiManagerToolbar UIManagerItemType = 0b100
	// UiManagerPlaceholder: insert a placeholder.
	UiManagerPlaceholder UIManagerItemType = 0b1000
	// UiManagerPopup: create a popup menu.
	UiManagerPopup UIManagerItemType = 0b10000
	// UiManagerMenuitem: create a menuitem.
	UiManagerMenuitem UIManagerItemType = 0b100000
	// UiManagerToolitem: create a toolitem.
	UiManagerToolitem UIManagerItemType = 0b1000000
	// UiManagerSeparator: create a separator.
	UiManagerSeparator UIManagerItemType = 0b10000000
	// UiManagerAccelerator: install an accelerator.
	UiManagerAccelerator UIManagerItemType = 0b100000000
	// UiManagerPopupWithAccels: same as GTK_UI_MANAGER_POPUP, but the actions’
	// accelerators are shown.
	UiManagerPopupWithAccels UIManagerItemType = 0b1000000000
)

func marshalUIManagerItemType(p uintptr) (interface{}, error) {
	return UIManagerItemType(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for UIManagerItemType.
func (u UIManagerItemType) String() string {
	if u == 0 {
		return "UIManagerItemType(0)"
	}

	var builder strings.Builder
	builder.Grow(198)

	for u != 0 {
		next := u & (u - 1)
		bit := u - next

		switch bit {
		case UiManagerAuto:
			builder.WriteString("Auto|")
		case UiManagerMenubar:
			builder.WriteString("Menubar|")
		case UiManagerMenu:
			builder.WriteString("Menu|")
		case UiManagerToolbar:
			builder.WriteString("Toolbar|")
		case UiManagerPlaceholder:
			builder.WriteString("Placeholder|")
		case UiManagerPopup:
			builder.WriteString("Popup|")
		case UiManagerMenuitem:
			builder.WriteString("Menuitem|")
		case UiManagerToolitem:
			builder.WriteString("Toolitem|")
		case UiManagerSeparator:
			builder.WriteString("Separator|")
		case UiManagerAccelerator:
			builder.WriteString("Accelerator|")
		case UiManagerPopupWithAccels:
			builder.WriteString("PopupWithAccels|")
		default:
			builder.WriteString(fmt.Sprintf("UIManagerItemType(0b%b)|", bit))
		}

		u = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if u contains other.
func (u UIManagerItemType) Has(other UIManagerItemType) bool {
	return (u & other) == other
}

// UIManagerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type UIManagerOverrider interface {
	ActionsChanged()
	// The function takes the following parameters:
	//
	AddWidget(widget Widgetter)
	// The function takes the following parameters:
	//
	//    - action
	//    - proxy
	//
	ConnectProxy(action *Action, proxy Widgetter)
	// The function takes the following parameters:
	//
	//    - action
	//    - proxy
	//
	DisconnectProxy(action *Action, proxy Widgetter)
	// Action looks up an action by following a path. See
	// gtk_ui_manager_get_widget() for more information about paths.
	//
	// Deprecated: since version 3.10.
	//
	// The function takes the following parameters:
	//
	//    - path: path.
	//
	// The function returns the following values:
	//
	//    - action whose proxy widget is found by following the path, or NULL if
	//      no widget was found.
	//
	Action(path string) *Action
	// Widget looks up a widget by following a path. The path consists of the
	// names specified in the XML description of the UI. separated by “/”.
	// Elements which don’t have a name or action attribute in the XML (e.g.
	// <popup>) can be addressed by their XML element name (e.g. "popup"). The
	// root element ("/ui") can be omitted in the path.
	//
	// Note that the widget found by following a path that ends in a <menu>;
	// element is the menuitem to which the menu is attached, not the menu it
	// manages.
	//
	// Also note that the widgets constructed by a ui manager are not tied to
	// the lifecycle of the ui manager. If you add the widgets returned by this
	// function to some container or explicitly ref them, they will survive the
	// destruction of the ui manager.
	//
	// Deprecated: since version 3.10.
	//
	// The function takes the following parameters:
	//
	//    - path: path.
	//
	// The function returns the following values:
	//
	//    - widget found by following the path, or NULL if no widget was found.
	//
	Widget(path string) Widgetter
	// The function takes the following parameters:
	//
	PostActivate(action *Action)
	// The function takes the following parameters:
	//
	PreActivate(action *Action)
}

// UIManager: > GtkUIManager is deprecated since GTK+ 3.10. To construct user
// interfaces > from XML definitions, you should use Builder, Model, et al. To >
// work with actions, use #GAction, Actionable et al. These newer classes >
// support richer functionality and integration with various desktop shells. >
// It should be possible to migrate most/all functionality from GtkUIManager.
//
// A UIManager constructs a user interface (menus and toolbars) from one or more
// UI definitions, which reference actions from one or more action groups.
//
//
// UI Definitions
//
// The UI definitions are specified in an XML format which can be roughly
// described by the following DTD.
//
// > Do not confuse the GtkUIManager UI Definitions described here with > the
// similarly named [GtkBuilder UI Definitions][BUILDER-UI].
//
//    <!ELEMENT ui          (menubar|toolbar|popup|accelerator)* >
//    <!ELEMENT menubar     (menuitem|separator|placeholder|menu)* >
//    <!ELEMENT menu        (menuitem|separator|placeholder|menu)* >
//    <!ELEMENT popup       (menuitem|separator|placeholder|menu)* >
//    <!ELEMENT toolbar     (toolitem|separator|placeholder)* >
//    <!ELEMENT placeholder (menuitem|toolitem|separator|placeholder|menu)* >
//    <!ELEMENT menuitem     EMPTY >
//    <!ELEMENT toolitem     (menu?) >
//    <!ELEMENT separator    EMPTY >
//    <!ELEMENT accelerator  EMPTY >
//    <!ATTLIST menubar      name                      PLIED
//                           action                    PLIED >
//    <!ATTLIST toolbar      name                      PLIED
//                           action                    PLIED >
//    <!ATTLIST popup        name                      PLIED
//                           action                    PLIED
//                           accelerators (true|false) PLIED >
//    <!ATTLIST placeholder  name                      PLIED
//                           action                    PLIED >
//    <!ATTLIST separator    name                      PLIED
//                           action                    PLIED
//                           expand       (true|false) PLIED >
//    <!ATTLIST menu         name                      PLIED
//                           action                    QUIRED
//                           position     (top|bot)    PLIED >
//    <!ATTLIST menuitem     name                      PLIED
//                           action                    QUIRED
//                           position     (top|bot)    PLIED
//                           always-show-image (true|false) PLIED >
//    <!ATTLIST toolitem     name                      PLIED
//                           action                    QUIRED
//                           position     (top|bot)    PLIED >
//    <!ATTLIST accelerator  name                      PLIED
//                           action                    QUIRED >
//
// There are some additional restrictions beyond those specified in the DTD,
// e.g. every toolitem must have a toolbar in its anchestry and every menuitem
// must have a menubar or popup in its anchestry. Since a Parser is used to
// parse the UI description, it must not only be valid XML, but valid markup.
//
// If a name is not specified, it defaults to the action. If an action is not
// specified either, the element name is used. The name and action attributes
// must not contain “/” characters after parsing (since that would mess up path
// lookup) and must be usable as XML attributes when enclosed in doublequotes,
// thus they must not “"” characters or references to the &quot; entity.
//
// A UI definition #
//
//    <ui>
//      <menubar>
//        <menu name="FileMenu" action="FileMenuAction">
//          <menuitem name="New" action="New2Action" />
//          <placeholder name="FileMenuAdditions" />
//        </menu>
//        <menu name="JustifyMenu" action="JustifyMenuAction">
//          <menuitem name="Left" action="justify-left"/>
//          <menuitem name="Centre" action="justify-center"/>
//          <menuitem name="Right" action="justify-right"/>
//          <menuitem name="Fill" action="justify-fill"/>
//        </menu>
//      </menubar>
//      <toolbar action="toolbar1">
//        <placeholder name="JustifyToolItems">
//          <separator/>
//          <toolitem name="Left" action="justify-left"/>
//          <toolitem name="Centre" action="justify-center"/>
//          <toolitem name="Right" action="justify-right"/>
//          <toolitem name="Fill" action="justify-fill"/>
//          <separator/>
//        </placeholder>
//      </toolbar>
//    </ui>
//
// The constructed widget hierarchy is very similar to the element tree of the
// XML, with the exception that placeholders are merged into their parents. The
// correspondence of XML elements to widgets should be almost obvious:
//
// - menubar
//
//    a MenuBar
//
// - toolbar
//
//    a Toolbar
//
// - popup
//
//    a toplevel Menu
//
// - menu
//
//    a Menu attached to a menuitem
//
// - menuitem
//
//    a MenuItem subclass, the exact type depends on the action
//
// - toolitem
//
//    a ToolItem subclass, the exact type depends on the
//    action. Note that toolitem elements may contain a menu element,
//    but only if their associated action specifies a
//    MenuToolButton as proxy.
//
// - separator
//
//    a SeparatorMenuItem or SeparatorToolItem
//
// - accelerator
//
//    a keyboard accelerator
//
// The “position” attribute determines where a constructed widget is positioned
// wrt. to its siblings in the partially constructed tree. If it is “top”, the
// widget is prepended, otherwise it is appended.
//
//
// UI Merging
//
// The most remarkable feature of UIManager is that it can overlay a set of
// menuitems and toolitems over another one, and demerge them later.
//
// Merging is done based on the names of the XML elements. Each element is
// identified by a path which consists of the names of its anchestors, separated
// by slashes. For example, the menuitem named “Left” in the example above has
// the path /ui/menubar/JustifyMenu/Left and the toolitem with the same name has
// path /ui/toolbar1/JustifyToolItems/Left.
//
// Accelerators #
//
// Every action has an accelerator path. Accelerators are installed together
// with menuitem proxies, but they can also be explicitly added with
// <accelerator> elements in the UI definition. This makes it possible to have
// accelerators for actions even if they have no visible proxies.
//
//
// Smart Separators
//
// The separators created by UIManager are “smart”, i.e. they do not show up in
// the UI unless they end up between two visible menu or tool items. Separators
// which are located at the very beginning or end of the menu or toolbar
// containing them, or multiple separators next to each other, are hidden. This
// is a useful feature, since the merging of UI elements from multiple sources
// can make it hard or impossible to determine in advance whether a separator
// will end up in such an unfortunate position.
//
// For separators in toolbars, you can set expand="true" to turn them from a
// small, visible separator to an expanding, invisible one. Toolitems following
// an expanding separator are effectively right-aligned.
//
//
// Empty Menus
//
// Submenus pose similar problems to separators inconnection with merging. It is
// impossible to know in advance whether they will end up empty after merging.
// UIManager offers two ways to treat empty submenus:
//
// - make them disappear by hiding the menu item they’re attached to
//
// - add an insensitive “Empty” item
//
// The behaviour is chosen based on the “hide_if_empty” property of the action
// to which the submenu is associated.
//
//
// GtkUIManager as GtkBuildable
//
// The GtkUIManager implementation of the GtkBuildable interface accepts
// GtkActionGroup objects as <child> elements in UI definitions.
//
// A GtkUIManager UI definition as described above can be embedded in an
// GtkUIManager <object> element in a GtkBuilder UI definition.
//
// The widgets that are constructed by a GtkUIManager can be embedded in other
// parts of the constructed user interface with the help of the “constructor”
// attribute. See the example below.
//
// An embedded GtkUIManager UI definition
//
//    <object class="GtkUIManager" id="uiman">
//      <child>
//        <object class="GtkActionGroup" id="actiongroup">
//          <child>
//            <object class="GtkAction" id="file">
//              <property name="label">_File</property>
//            </object>
//          </child>
//        </object>
//      </child>
//      <ui>
//        <menubar name="menubar1">
//          <menu action="file">
//          </menu>
//        </menubar>
//      </ui>
//    </object>
//    <object class="GtkWindow" id="main-window">
//      <child>
//        <object class="GtkMenuBar" id="menubar1" constructor="uiman"/>
//      </child>
//    </object>.
type UIManager struct {
	*externglib.Object

	Buildable
}

var (
	_ externglib.Objector = (*UIManager)(nil)
)

func wrapUIManager(obj *externglib.Object) *UIManager {
	return &UIManager{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalUIManagerer(p uintptr) (interface{}, error) {
	return wrapUIManager(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectActionsChanged signal is emitted whenever the set of actions changes.
func (manager *UIManager) ConnectActionsChanged(f func()) externglib.SignalHandle {
	return manager.Connect("actions-changed", f)
}

// ConnectAddWidget signal is emitted for each generated menubar and toolbar. It
// is not emitted for generated popup menus, which can be obtained by
// gtk_ui_manager_get_widget().
func (manager *UIManager) ConnectAddWidget(f func(widget Widgetter)) externglib.SignalHandle {
	return manager.Connect("add-widget", f)
}

// ConnectConnectProxy signal is emitted after connecting a proxy to an action
// in the group.
//
// This is intended for simple customizations for which a custom action class
// would be too clumsy, e.g. showing tooltips for menuitems in the statusbar.
func (manager *UIManager) ConnectConnectProxy(f func(action Action, proxy Widgetter)) externglib.SignalHandle {
	return manager.Connect("connect-proxy", f)
}

// ConnectDisconnectProxy signal is emitted after disconnecting a proxy from an
// action in the group.
func (manager *UIManager) ConnectDisconnectProxy(f func(action Action, proxy Widgetter)) externglib.SignalHandle {
	return manager.Connect("disconnect-proxy", f)
}

// ConnectPostActivate signal is emitted just after the action is activated.
//
// This is intended for applications to get notification just after any action
// is activated.
func (manager *UIManager) ConnectPostActivate(f func(action Action)) externglib.SignalHandle {
	return manager.Connect("post-activate", f)
}

// ConnectPreActivate signal is emitted just before the action is activated.
//
// This is intended for applications to get notification just before any action
// is activated.
func (manager *UIManager) ConnectPreActivate(f func(action Action)) externglib.SignalHandle {
	return manager.Connect("pre-activate", f)
}

// NewUIManager creates a new ui manager object.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - uiManager: new ui manager object.
//
func NewUIManager() *UIManager {
	var _cret *C.GtkUIManager // in

	_cret = C.gtk_ui_manager_new()

	var _uiManager *UIManager // out

	_uiManager = wrapUIManager(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _uiManager
}

// AddUi adds a UI element to the current contents of manager.
//
// If type is GTK_UI_MANAGER_AUTO, GTK+ inserts a menuitem, toolitem or
// separator if such an element can be inserted at the place determined by path.
// Otherwise type must indicate an element that can be inserted at the place
// determined by path.
//
// If path points to a menuitem or toolitem, the new element will be inserted
// before or after this item, depending on top.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - mergeId: merge id for the merged UI, see gtk_ui_manager_new_merge_id().
//    - path: path.
//    - name for the added UI element.
//    - action (optional): name of the action to be proxied, or NULL to add a
//      separator.
//    - typ: type of UI element to add.
//    - top: if TRUE, the UI element is added before its siblings, otherwise it
//      is added after its siblings.
//
func (manager *UIManager) AddUi(mergeId uint, path, name, action string, typ UIManagerItemType, top bool) {
	var _arg0 *C.GtkUIManager        // out
	var _arg1 C.guint                // out
	var _arg2 *C.gchar               // out
	var _arg3 *C.gchar               // out
	var _arg4 *C.gchar               // out
	var _arg5 C.GtkUIManagerItemType // out
	var _arg6 C.gboolean             // out

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = C.guint(mergeId)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg3))
	if action != "" {
		_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(action)))
		defer C.free(unsafe.Pointer(_arg4))
	}
	_arg5 = C.GtkUIManagerItemType(typ)
	if top {
		_arg6 = C.TRUE
	}

	C.gtk_ui_manager_add_ui(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(mergeId)
	runtime.KeepAlive(path)
	runtime.KeepAlive(name)
	runtime.KeepAlive(action)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(top)
}

// AddUiFromFile parses a file containing a [UI definition][XML-UI] and merges
// it with the current contents of manager.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - filename: name of the file to parse.
//
// The function returns the following values:
//
//    - guint: merge id for the merged UI. The merge id can be used to unmerge
//      the UI with gtk_ui_manager_remove_ui(). If an error occurred, the return
//      value is 0.
//
func (manager *UIManager) AddUiFromFile(filename string) (uint, error) {
	var _arg0 *C.GtkUIManager // out
	var _arg1 *C.gchar        // out
	var _cret C.guint         // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_ui_manager_add_ui_from_file(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(filename)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _guint, _goerr
}

// AddUiFromResource parses a resource file containing a [UI definition][XML-UI]
// and merges it with the current contents of manager.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - resourcePath: resource path of the file to parse.
//
// The function returns the following values:
//
//    - guint: merge id for the merged UI. The merge id can be used to unmerge
//      the UI with gtk_ui_manager_remove_ui(). If an error occurred, the return
//      value is 0.
//
func (manager *UIManager) AddUiFromResource(resourcePath string) (uint, error) {
	var _arg0 *C.GtkUIManager // out
	var _arg1 *C.gchar        // out
	var _cret C.guint         // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_ui_manager_add_ui_from_resource(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(resourcePath)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _guint, _goerr
}

// AddUiFromString parses a string containing a [UI definition][XML-UI] and
// merges it with the current contents of manager. An enclosing <ui> element is
// added if it is missing.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - buffer: string to parse.
//    - length of buffer (may be -1 if buffer is nul-terminated).
//
// The function returns the following values:
//
//    - guint: merge id for the merged UI. The merge id can be used to unmerge
//      the UI with gtk_ui_manager_remove_ui(). If an error occurred, the return
//      value is 0.
//
func (manager *UIManager) AddUiFromString(buffer string, length int) (uint, error) {
	var _arg0 *C.GtkUIManager // out
	var _arg1 *C.gchar        // out
	var _arg2 C.gssize        // out
	var _cret C.guint         // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(buffer)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(length)

	_cret = C.gtk_ui_manager_add_ui_from_string(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(length)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _guint, _goerr
}

// EnsureUpdate makes sure that all pending updates to the UI have been
// completed.
//
// This may occasionally be necessary, since UIManager updates the UI in an idle
// function. A typical example where this function is useful is to enforce that
// the menubar and toolbar have been added to the main window before showing it:
//
//    gtk_container_add (GTK_CONTAINER (window), vbox);
//    g_signal_connect (merge, "add-widget",
//                      G_CALLBACK (add_widget), vbox);
//    gtk_ui_manager_add_ui_from_file (merge, "my-menus");
//    gtk_ui_manager_add_ui_from_file (merge, "my-toolbars");
//    gtk_ui_manager_ensure_update (merge);
//    gtk_widget_show (window);
//
// Deprecated: since version 3.10.
func (manager *UIManager) EnsureUpdate() {
	var _arg0 *C.GtkUIManager // out

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))

	C.gtk_ui_manager_ensure_update(_arg0)
	runtime.KeepAlive(manager)
}

// AccelGroup returns the AccelGroup associated with manager.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - accelGroup: AccelGroup.
//
func (manager *UIManager) AccelGroup() *AccelGroup {
	var _arg0 *C.GtkUIManager  // out
	var _cret *C.GtkAccelGroup // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_ui_manager_get_accel_group(_arg0)
	runtime.KeepAlive(manager)

	var _accelGroup *AccelGroup // out

	_accelGroup = wrapAccelGroup(externglib.Take(unsafe.Pointer(_cret)))

	return _accelGroup
}

// Action looks up an action by following a path. See
// gtk_ui_manager_get_widget() for more information about paths.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - path: path.
//
// The function returns the following values:
//
//    - action whose proxy widget is found by following the path, or NULL if no
//      widget was found.
//
func (manager *UIManager) Action(path string) *Action {
	var _arg0 *C.GtkUIManager // out
	var _arg1 *C.gchar        // out
	var _cret *C.GtkAction    // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_ui_manager_get_action(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(path)

	var _action *Action // out

	_action = wrapAction(externglib.Take(unsafe.Pointer(_cret)))

	return _action
}

// ActionGroups returns the list of action groups associated with manager.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - list of action groups. The list is owned by GTK+ and should not be
//      modified.
//
func (manager *UIManager) ActionGroups() []ActionGroup {
	var _arg0 *C.GtkUIManager // out
	var _cret *C.GList        // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_ui_manager_get_action_groups(_arg0)
	runtime.KeepAlive(manager)

	var _list []ActionGroup // out

	_list = make([]ActionGroup, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.GtkActionGroup)(v)
		var dst ActionGroup // out
		dst = *wrapActionGroup(externglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// AddTearoffs returns whether menus generated by this UIManager will have
// tearoff menu items.
//
// Deprecated: Tearoff menus are deprecated and should not be used in newly
// written code.
//
// The function returns the following values:
//
//    - ok: whether tearoff menu items are added.
//
func (manager *UIManager) AddTearoffs() bool {
	var _arg0 *C.GtkUIManager // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_ui_manager_get_add_tearoffs(_arg0)
	runtime.KeepAlive(manager)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Toplevels obtains a list of all toplevel widgets of the requested types.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - types specifies the types of toplevel widgets to include. Allowed types
//      are K_UI_MANAGER_MENUBAR, K_UI_MANAGER_TOOLBAR and K_UI_MANAGER_POPUP.
//
// The function returns the following values:
//
//    - sList: newly-allocated List of all toplevel widgets of the requested
//      types. Free the returned list with g_slist_free().
//
func (manager *UIManager) Toplevels(types UIManagerItemType) []Widgetter {
	var _arg0 *C.GtkUIManager        // out
	var _arg1 C.GtkUIManagerItemType // out
	var _cret *C.GSList              // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = C.GtkUIManagerItemType(types)

	_cret = C.gtk_ui_manager_get_toplevels(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(types)

	var _sList []Widgetter // out

	_sList = make([]Widgetter, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkWidget)(v)
		var dst Widgetter // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gtk.Widgetter is nil")
			}

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
			}
			dst = rv
		}
		_sList = append(_sList, dst)
	})

	return _sList
}

// Ui creates a [UI definition][XML-UI] of the merged UI.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - utf8: newly allocated string containing an XML representation of the
//      merged UI.
//
func (manager *UIManager) Ui() string {
	var _arg0 *C.GtkUIManager // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_ui_manager_get_ui(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Widget looks up a widget by following a path. The path consists of the names
// specified in the XML description of the UI. separated by “/”. Elements which
// don’t have a name or action attribute in the XML (e.g. <popup>) can be
// addressed by their XML element name (e.g. "popup"). The root element ("/ui")
// can be omitted in the path.
//
// Note that the widget found by following a path that ends in a <menu>; element
// is the menuitem to which the menu is attached, not the menu it manages.
//
// Also note that the widgets constructed by a ui manager are not tied to the
// lifecycle of the ui manager. If you add the widgets returned by this function
// to some container or explicitly ref them, they will survive the destruction
// of the ui manager.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - path: path.
//
// The function returns the following values:
//
//    - widget found by following the path, or NULL if no widget was found.
//
func (manager *UIManager) Widget(path string) Widgetter {
	var _arg0 *C.GtkUIManager // out
	var _arg1 *C.gchar        // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_ui_manager_get_widget(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(path)

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

// InsertActionGroup inserts an action group into the list of action groups
// associated with manager. Actions in earlier groups hide actions with the same
// name in later groups.
//
// If pos is larger than the number of action groups in manager, or negative,
// action_group will be inserted at the end of the internal list.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - actionGroup: action group to be inserted.
//    - pos: position at which the group will be inserted.
//
func (manager *UIManager) InsertActionGroup(actionGroup *ActionGroup, pos int) {
	var _arg0 *C.GtkUIManager   // out
	var _arg1 *C.GtkActionGroup // out
	var _arg2 C.gint            // out

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg2 = C.gint(pos)

	C.gtk_ui_manager_insert_action_group(_arg0, _arg1, _arg2)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(pos)
}

// NewMergeID returns an unused merge id, suitable for use with
// gtk_ui_manager_add_ui().
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - guint: unused merge id.
//
func (manager *UIManager) NewMergeID() uint {
	var _arg0 *C.GtkUIManager // out
	var _cret C.guint         // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_ui_manager_new_merge_id(_arg0)
	runtime.KeepAlive(manager)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// RemoveActionGroup removes an action group from the list of action groups
// associated with manager.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - actionGroup: action group to be removed.
//
func (manager *UIManager) RemoveActionGroup(actionGroup *ActionGroup) {
	var _arg0 *C.GtkUIManager   // out
	var _arg1 *C.GtkActionGroup // out

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))

	C.gtk_ui_manager_remove_action_group(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(actionGroup)
}

// RemoveUi unmerges the part of manager's content identified by merge_id.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - mergeId: merge id as returned by gtk_ui_manager_add_ui_from_string().
//
func (manager *UIManager) RemoveUi(mergeId uint) {
	var _arg0 *C.GtkUIManager // out
	var _arg1 C.guint         // out

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = C.guint(mergeId)

	C.gtk_ui_manager_remove_ui(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(mergeId)
}

// SetAddTearoffs sets the “add_tearoffs” property, which controls whether menus
// generated by this UIManager will have tearoff menu items.
//
// Note that this only affects regular menus. Generated popup menus never have
// tearoff menu items.
//
// Deprecated: Tearoff menus are deprecated and should not be used in newly
// written code.
//
// The function takes the following parameters:
//
//    - addTearoffs: whether tearoff menu items are added.
//
func (manager *UIManager) SetAddTearoffs(addTearoffs bool) {
	var _arg0 *C.GtkUIManager // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	if addTearoffs {
		_arg1 = C.TRUE
	}

	C.gtk_ui_manager_set_add_tearoffs(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(addTearoffs)
}
