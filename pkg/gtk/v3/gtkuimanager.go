// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
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
		{T: externglib.Type(C.gtk_ui_manager_item_type_get_type()), F: marshalUIManagerItemType},
		{T: externglib.Type(C.gtk_ui_manager_get_type()), F: marshalUIManagerer},
	})
}

// UIManagerItemType: these enumeration values are used by
// gtk_ui_manager_add_ui() to determine what UI element to create.
//
// Deprecated: since version 3.10.
type UIManagerItemType int

const (
	// UIManagerItemTypeAuto: pick the type of the UI element according to
	// context.
	UIManagerItemTypeAuto UIManagerItemType = 0b0
	// UIManagerItemTypeMenubar: create a menubar.
	UIManagerItemTypeMenubar UIManagerItemType = 0b1
	// UIManagerItemTypeMenu: create a menu.
	UIManagerItemTypeMenu UIManagerItemType = 0b10
	// UIManagerItemTypeToolbar: create a toolbar.
	UIManagerItemTypeToolbar UIManagerItemType = 0b100
	// UIManagerItemTypePlaceholder: insert a placeholder.
	UIManagerItemTypePlaceholder UIManagerItemType = 0b1000
	// UIManagerItemTypePopup: create a popup menu.
	UIManagerItemTypePopup UIManagerItemType = 0b10000
	// UIManagerItemTypeMenuitem: create a menuitem.
	UIManagerItemTypeMenuitem UIManagerItemType = 0b100000
	// UIManagerItemTypeToolitem: create a toolitem.
	UIManagerItemTypeToolitem UIManagerItemType = 0b1000000
	// UIManagerItemTypeSeparator: create a separator.
	UIManagerItemTypeSeparator UIManagerItemType = 0b10000000
	// UIManagerItemTypeAccelerator: install an accelerator.
	UIManagerItemTypeAccelerator UIManagerItemType = 0b100000000
	// UIManagerItemTypePopupWithAccels: same as GTK_UI_MANAGER_POPUP, but the
	// actions’ accelerators are shown.
	UIManagerItemTypePopupWithAccels UIManagerItemType = 0b1000000000
)

func marshalUIManagerItemType(p uintptr) (interface{}, error) {
	return UIManagerItemType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// UIManagerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type UIManagerOverrider interface {
	ActionsChanged()
	AddWidget(widget Widgeter)
	ConnectProxy(action Actioner, proxy Widgeter)
	DisconnectProxy(action Actioner, proxy Widgeter)
	// Action looks up an action by following a path. See
	// gtk_ui_manager_get_widget() for more information about paths.
	//
	// Deprecated: since version 3.10.
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
	Widget(path string) *Widget
	PostActivate(action Actioner)
	PreActivate(action Actioner)
}

// UIManagerer describes UIManager's methods.
type UIManagerer interface {
	// AddUi adds a UI element to the current contents of @manager.
	AddUi(mergeId uint, path string, name string, action string, typ UIManagerItemType, top bool)
	// AddUiFromFile parses a file containing a [UI definition][XML-UI] and
	// merges it with the current contents of @manager.
	AddUiFromFile(filename string) (uint, error)
	// AddUiFromResource parses a resource file containing a [UI
	// definition][XML-UI] and merges it with the current contents of @manager.
	AddUiFromResource(resourcePath string) (uint, error)
	// AddUiFromString parses a string containing a [UI definition][XML-UI] and
	// merges it with the current contents of @manager.
	AddUiFromString(buffer string, length int) (uint, error)
	// EnsureUpdate makes sure that all pending updates to the UI have been
	// completed.
	EnsureUpdate()
	// AccelGroup returns the AccelGroup associated with @manager.
	AccelGroup() *AccelGroup
	// Action looks up an action by following a path.
	Action(path string) *Action
	// AddTearoffs returns whether menus generated by this UIManager will have
	// tearoff menu items.
	AddTearoffs() bool
	// Ui creates a [UI definition][XML-UI] of the merged UI.
	Ui() string
	// Widget looks up a widget by following a path.
	Widget(path string) *Widget
	// InsertActionGroup inserts an action group into the list of action groups
	// associated with @manager.
	InsertActionGroup(actionGroup ActionGrouper, pos int)
	// NewMergeID returns an unused merge id, suitable for use with
	// gtk_ui_manager_add_ui().
	NewMergeID() uint
	// RemoveActionGroup removes an action group from the list of action groups
	// associated with @manager.
	RemoveActionGroup(actionGroup ActionGrouper)
	// RemoveUi unmerges the part of @manager's content identified by @merge_id.
	RemoveUi(mergeId uint)
	// SetAddTearoffs sets the “add_tearoffs” property, which controls whether
	// menus generated by this UIManager will have tearoff menu items.
	SetAddTearoffs(addTearoffs bool)
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
// the path `/ui/menubar/JustifyMenu/Left` and the toolitem with the same name
// has path `/ui/toolbar1/JustifyToolItems/Left`.
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
// For separators in toolbars, you can set `expand="true"` to turn them from a
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
//    </object>
type UIManager struct {
	*externglib.Object

	Buildable
}

var (
	_ UIManagerer     = (*UIManager)(nil)
	_ gextras.Nativer = (*UIManager)(nil)
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
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapUIManager(obj), nil
}

// NewUIManager creates a new ui manager object.
//
// Deprecated: since version 3.10.
func NewUIManager() *UIManager {
	var _cret *C.GtkUIManager // in

	_cret = C.gtk_ui_manager_new()

	var _uiManager *UIManager // out

	_uiManager = wrapUIManager(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _uiManager
}

// AddUi adds a UI element to the current contents of @manager.
//
// If @type is GTK_UI_MANAGER_AUTO, GTK+ inserts a menuitem, toolitem or
// separator if such an element can be inserted at the place determined by
// @path. Otherwise @type must indicate an element that can be inserted at the
// place determined by @path.
//
// If @path points to a menuitem or toolitem, the new element will be inserted
// before or after this item, depending on @top.
//
// Deprecated: since version 3.10.
func (manager *UIManager) AddUi(mergeId uint, path string, name string, action string, typ UIManagerItemType, top bool) {
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
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(action)))
	_arg5 = C.GtkUIManagerItemType(typ)
	if top {
		_arg6 = C.TRUE
	}

	C.gtk_ui_manager_add_ui(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// AddUiFromFile parses a file containing a [UI definition][XML-UI] and merges
// it with the current contents of @manager.
//
// Deprecated: since version 3.10.
func (manager *UIManager) AddUiFromFile(filename string) (uint, error) {
	var _arg0 *C.GtkUIManager // out
	var _arg1 *C.gchar        // out
	var _cret C.guint         // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))

	_cret = C.gtk_ui_manager_add_ui_from_file(_arg0, _arg1, &_cerr)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _guint, _goerr
}

// AddUiFromResource parses a resource file containing a [UI definition][XML-UI]
// and merges it with the current contents of @manager.
//
// Deprecated: since version 3.10.
func (manager *UIManager) AddUiFromResource(resourcePath string) (uint, error) {
	var _arg0 *C.GtkUIManager // out
	var _arg1 *C.gchar        // out
	var _cret C.guint         // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(resourcePath)))

	_cret = C.gtk_ui_manager_add_ui_from_resource(_arg0, _arg1, &_cerr)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _guint, _goerr
}

// AddUiFromString parses a string containing a [UI definition][XML-UI] and
// merges it with the current contents of @manager. An enclosing <ui> element is
// added if it is missing.
//
// Deprecated: since version 3.10.
func (manager *UIManager) AddUiFromString(buffer string, length int) (uint, error) {
	var _arg0 *C.GtkUIManager // out
	var _arg1 *C.gchar        // out
	var _arg2 C.gssize        // out
	var _cret C.guint         // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(buffer)))
	_arg2 = C.gssize(length)

	_cret = C.gtk_ui_manager_add_ui_from_string(_arg0, _arg1, _arg2, &_cerr)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

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
}

// AccelGroup returns the AccelGroup associated with @manager.
//
// Deprecated: since version 3.10.
func (manager *UIManager) AccelGroup() *AccelGroup {
	var _arg0 *C.GtkUIManager  // out
	var _cret *C.GtkAccelGroup // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_ui_manager_get_accel_group(_arg0)

	var _accelGroup *AccelGroup // out

	_accelGroup = wrapAccelGroup(externglib.Take(unsafe.Pointer(_cret)))

	return _accelGroup
}

// Action looks up an action by following a path. See
// gtk_ui_manager_get_widget() for more information about paths.
//
// Deprecated: since version 3.10.
func (manager *UIManager) Action(path string) *Action {
	var _arg0 *C.GtkUIManager // out
	var _arg1 *C.gchar        // out
	var _cret *C.GtkAction    // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(path)))

	_cret = C.gtk_ui_manager_get_action(_arg0, _arg1)

	var _action *Action // out

	_action = wrapAction(externglib.Take(unsafe.Pointer(_cret)))

	return _action
}

// AddTearoffs returns whether menus generated by this UIManager will have
// tearoff menu items.
//
// Deprecated: Tearoff menus are deprecated and should not be used in newly
// written code.
func (manager *UIManager) AddTearoffs() bool {
	var _arg0 *C.GtkUIManager // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_ui_manager_get_add_tearoffs(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Ui creates a [UI definition][XML-UI] of the merged UI.
//
// Deprecated: since version 3.10.
func (manager *UIManager) Ui() string {
	var _arg0 *C.GtkUIManager // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_ui_manager_get_ui(_arg0)

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
func (manager *UIManager) Widget(path string) *Widget {
	var _arg0 *C.GtkUIManager // out
	var _arg1 *C.gchar        // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(path)))

	_cret = C.gtk_ui_manager_get_widget(_arg0, _arg1)

	var _widget *Widget // out

	_widget = wrapWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _widget
}

// InsertActionGroup inserts an action group into the list of action groups
// associated with @manager. Actions in earlier groups hide actions with the
// same name in later groups.
//
// If @pos is larger than the number of action groups in @manager, or negative,
// @action_group will be inserted at the end of the internal list.
//
// Deprecated: since version 3.10.
func (manager *UIManager) InsertActionGroup(actionGroup ActionGrouper, pos int) {
	var _arg0 *C.GtkUIManager   // out
	var _arg1 *C.GtkActionGroup // out
	var _arg2 C.gint            // out

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GtkActionGroup)(unsafe.Pointer((actionGroup).(gextras.Nativer).Native()))
	_arg2 = C.gint(pos)

	C.gtk_ui_manager_insert_action_group(_arg0, _arg1, _arg2)
}

// NewMergeID returns an unused merge id, suitable for use with
// gtk_ui_manager_add_ui().
//
// Deprecated: since version 3.10.
func (manager *UIManager) NewMergeID() uint {
	var _arg0 *C.GtkUIManager // out
	var _cret C.guint         // in

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_ui_manager_new_merge_id(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// RemoveActionGroup removes an action group from the list of action groups
// associated with @manager.
//
// Deprecated: since version 3.10.
func (manager *UIManager) RemoveActionGroup(actionGroup ActionGrouper) {
	var _arg0 *C.GtkUIManager   // out
	var _arg1 *C.GtkActionGroup // out

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GtkActionGroup)(unsafe.Pointer((actionGroup).(gextras.Nativer).Native()))

	C.gtk_ui_manager_remove_action_group(_arg0, _arg1)
}

// RemoveUi unmerges the part of @manager's content identified by @merge_id.
//
// Deprecated: since version 3.10.
func (manager *UIManager) RemoveUi(mergeId uint) {
	var _arg0 *C.GtkUIManager // out
	var _arg1 C.guint         // out

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	_arg1 = C.guint(mergeId)

	C.gtk_ui_manager_remove_ui(_arg0, _arg1)
}

// SetAddTearoffs sets the “add_tearoffs” property, which controls whether menus
// generated by this UIManager will have tearoff menu items.
//
// Note that this only affects regular menus. Generated popup menus never have
// tearoff menu items.
//
// Deprecated: Tearoff menus are deprecated and should not be used in newly
// written code.
func (manager *UIManager) SetAddTearoffs(addTearoffs bool) {
	var _arg0 *C.GtkUIManager // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkUIManager)(unsafe.Pointer(manager.Native()))
	if addTearoffs {
		_arg1 = C.TRUE
	}

	C.gtk_ui_manager_set_add_tearoffs(_arg0, _arg1)
}
