// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_ui_manager_get_type()), F: marshalUIManager},
	})
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
type UIManager interface {
	gextras.Objector
	Buildable

	// AddUi adds a UI element to the current contents of @manager.
	//
	// If @type is GTK_UI_MANAGER_AUTO, GTK+ inserts a menuitem, toolitem or
	// separator if such an element can be inserted at the place determined by
	// @path. Otherwise @type must indicate an element that can be inserted at
	// the place determined by @path.
	//
	// If @path points to a menuitem or toolitem, the new element will be
	// inserted before or after this item, depending on @top.
	AddUi(mergeID uint, path string, name string, action string, typ UIManagerItemType, top bool)
	// AddUiFromFile parses a file containing a [UI definition][XML-UI] and
	// merges it with the current contents of @manager.
	AddUiFromFile(filename string) (guint uint, err error)
	// AddUiFromResource parses a resource file containing a [UI
	// definition][XML-UI] and merges it with the current contents of @manager.
	AddUiFromResource(resourcePath string) (guint uint, err error)
	// AddUiFromString parses a string containing a [UI definition][XML-UI] and
	// merges it with the current contents of @manager. An enclosing <ui>
	// element is added if it is missing.
	AddUiFromString(buffer string, length int) (guint uint, err error)
	// EnsureUpdate makes sure that all pending updates to the UI have been
	// completed.
	//
	// This may occasionally be necessary, since UIManager updates the UI in an
	// idle function. A typical example where this function is useful is to
	// enforce that the menubar and toolbar have been added to the main window
	// before showing it:
	//
	//    gtk_container_add (GTK_CONTAINER (window), vbox);
	//    g_signal_connect (merge, "add-widget",
	//                      G_CALLBACK (add_widget), vbox);
	//    gtk_ui_manager_add_ui_from_file (merge, "my-menus");
	//    gtk_ui_manager_add_ui_from_file (merge, "my-toolbars");
	//    gtk_ui_manager_ensure_update (merge);
	//    gtk_widget_show (window);
	EnsureUpdate()
	// AccelGroup returns the AccelGroup associated with @manager.
	AccelGroup() AccelGroup
	// Action looks up an action by following a path. See
	// gtk_ui_manager_get_widget() for more information about paths.
	Action(path string) Action
	// ActionGroups returns the list of action groups associated with @manager.
	ActionGroups() *glib.List
	// AddTearoffs returns whether menus generated by this UIManager will have
	// tearoff menu items.
	AddTearoffs() bool
	// Toplevels obtains a list of all toplevel widgets of the requested types.
	Toplevels(types UIManagerItemType) *glib.SList
	// Ui creates a [UI definition][XML-UI] of the merged UI.
	Ui() string
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
	Widget(path string) Widget
	// InsertActionGroup inserts an action group into the list of action groups
	// associated with @manager. Actions in earlier groups hide actions with the
	// same name in later groups.
	//
	// If @pos is larger than the number of action groups in @manager, or
	// negative, @action_group will be inserted at the end of the internal list.
	InsertActionGroup(actionGroup ActionGroup, pos int)
	// NewMergeID returns an unused merge id, suitable for use with
	// gtk_ui_manager_add_ui().
	NewMergeID() uint
	// RemoveActionGroup removes an action group from the list of action groups
	// associated with @manager.
	RemoveActionGroup(actionGroup ActionGroup)
	// RemoveUi unmerges the part of @manager's content identified by @merge_id.
	RemoveUi(mergeID uint)
	// SetAddTearoffs sets the “add_tearoffs” property, which controls whether
	// menus generated by this UIManager will have tearoff menu items.
	//
	// Note that this only affects regular menus. Generated popup menus never
	// have tearoff menu items.
	SetAddTearoffs(addTearoffs bool)
}

// uiManager implements the UIManager interface.
type uiManager struct {
	gextras.Objector
	Buildable
}

var _ UIManager = (*uiManager)(nil)

// WrapUIManager wraps a GObject to the right type. It is
// primarily used internally.
func WrapUIManager(obj *externglib.Object) UIManager {
	return UIManager{
		Objector:  obj,
		Buildable: WrapBuildable(obj),
	}
}

func marshalUIManager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapUIManager(obj), nil
}

// NewUIManager constructs a class UIManager.
func NewUIManager() UIManager {
	ret := C.gtk_ui_manager_new()

	var ret0 UIManager

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(UIManager)

	return ret0
}

// AddUi adds a UI element to the current contents of @manager.
//
// If @type is GTK_UI_MANAGER_AUTO, GTK+ inserts a menuitem, toolitem or
// separator if such an element can be inserted at the place determined by
// @path. Otherwise @type must indicate an element that can be inserted at
// the place determined by @path.
//
// If @path points to a menuitem or toolitem, the new element will be
// inserted before or after this item, depending on @top.
func (m uiManager) AddUi(mergeID uint, path string, name string, action string, typ UIManagerItemType, top bool) {
	var arg0 *C.GtkUIManager
	var arg1 C.guint
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 *C.gchar
	var arg5 C.GtkUIManagerItemType
	var arg6 C.gboolean

	arg0 = (*C.GtkUIManager)(m.Native())
	arg1 = C.guint(mergeID)
	arg2 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = (*C.gchar)(C.CString(action))
	defer C.free(unsafe.Pointer(arg4))
	arg5 = (C.GtkUIManagerItemType)(typ)
	if top {
		arg6 = C.TRUE
	}

	C.gtk_ui_manager_add_ui(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// AddUiFromFile parses a file containing a [UI definition][XML-UI] and
// merges it with the current contents of @manager.
func (m uiManager) AddUiFromFile(filename string) (guint uint, err error) {
	var arg0 *C.GtkUIManager
	var arg1 *C.gchar
	var gError *C.GError

	arg0 = (*C.GtkUIManager)(m.Native())
	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.gtk_ui_manager_add_ui_from_file(arg0, arg1, &gError)

	var ret0 uint
	var goError error

	ret0 = uint(ret)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// AddUiFromResource parses a resource file containing a [UI
// definition][XML-UI] and merges it with the current contents of @manager.
func (m uiManager) AddUiFromResource(resourcePath string) (guint uint, err error) {
	var arg0 *C.GtkUIManager
	var arg1 *C.gchar
	var gError *C.GError

	arg0 = (*C.GtkUIManager)(m.Native())
	arg1 = (*C.gchar)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.gtk_ui_manager_add_ui_from_resource(arg0, arg1, &gError)

	var ret0 uint
	var goError error

	ret0 = uint(ret)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// AddUiFromString parses a string containing a [UI definition][XML-UI] and
// merges it with the current contents of @manager. An enclosing <ui>
// element is added if it is missing.
func (m uiManager) AddUiFromString(buffer string, length int) (guint uint, err error) {
	var arg0 *C.GtkUIManager
	var arg1 *C.gchar
	var arg2 C.gssize
	var gError *C.GError

	arg0 = (*C.GtkUIManager)(m.Native())
	arg1 = (*C.gchar)(C.CString(buffer))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gssize(length)

	ret := C.gtk_ui_manager_add_ui_from_string(arg0, arg1, arg2, &gError)

	var ret0 uint
	var goError error

	ret0 = uint(ret)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// EnsureUpdate makes sure that all pending updates to the UI have been
// completed.
//
// This may occasionally be necessary, since UIManager updates the UI in an
// idle function. A typical example where this function is useful is to
// enforce that the menubar and toolbar have been added to the main window
// before showing it:
//
//    gtk_container_add (GTK_CONTAINER (window), vbox);
//    g_signal_connect (merge, "add-widget",
//                      G_CALLBACK (add_widget), vbox);
//    gtk_ui_manager_add_ui_from_file (merge, "my-menus");
//    gtk_ui_manager_add_ui_from_file (merge, "my-toolbars");
//    gtk_ui_manager_ensure_update (merge);
//    gtk_widget_show (window);
func (m uiManager) EnsureUpdate() {
	var arg0 *C.GtkUIManager

	arg0 = (*C.GtkUIManager)(m.Native())

	C.gtk_ui_manager_ensure_update(arg0)
}

// AccelGroup returns the AccelGroup associated with @manager.
func (m uiManager) AccelGroup() AccelGroup {
	var arg0 *C.GtkUIManager

	arg0 = (*C.GtkUIManager)(m.Native())

	ret := C.gtk_ui_manager_get_accel_group(arg0)

	var ret0 AccelGroup

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(AccelGroup)

	return ret0
}

// Action looks up an action by following a path. See
// gtk_ui_manager_get_widget() for more information about paths.
func (m uiManager) Action(path string) Action {
	var arg0 *C.GtkUIManager
	var arg1 *C.gchar

	arg0 = (*C.GtkUIManager)(m.Native())
	arg1 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.gtk_ui_manager_get_action(arg0, arg1)

	var ret0 Action

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Action)

	return ret0
}

// ActionGroups returns the list of action groups associated with @manager.
func (m uiManager) ActionGroups() *glib.List {
	var arg0 *C.GtkUIManager

	arg0 = (*C.GtkUIManager)(m.Native())

	ret := C.gtk_ui_manager_get_action_groups(arg0)

	var ret0 *glib.List

	{
		ret0 = glib.WrapList(unsafe.Pointer(ret))
	}

	return ret0
}

// AddTearoffs returns whether menus generated by this UIManager will have
// tearoff menu items.
func (m uiManager) AddTearoffs() bool {
	var arg0 *C.GtkUIManager

	arg0 = (*C.GtkUIManager)(m.Native())

	ret := C.gtk_ui_manager_get_add_tearoffs(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != C.false

	return ret0
}

// Toplevels obtains a list of all toplevel widgets of the requested types.
func (m uiManager) Toplevels(types UIManagerItemType) *glib.SList {
	var arg0 *C.GtkUIManager
	var arg1 C.GtkUIManagerItemType

	arg0 = (*C.GtkUIManager)(m.Native())
	arg1 = (C.GtkUIManagerItemType)(types)

	ret := C.gtk_ui_manager_get_toplevels(arg0, arg1)

	var ret0 *glib.SList

	{
		ret0 = glib.WrapSList(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *glib.SList) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// Ui creates a [UI definition][XML-UI] of the merged UI.
func (m uiManager) Ui() string {
	var arg0 *C.GtkUIManager

	arg0 = (*C.GtkUIManager)(m.Native())

	ret := C.gtk_ui_manager_get_ui(arg0)

	var ret0 string

	ret0 = C.GoString(ret)
	C.free(unsafe.Pointer(ret))

	return ret0
}

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
func (m uiManager) Widget(path string) Widget {
	var arg0 *C.GtkUIManager
	var arg1 *C.gchar

	arg0 = (*C.GtkUIManager)(m.Native())
	arg1 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.gtk_ui_manager_get_widget(arg0, arg1)

	var ret0 Widget

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Widget)

	return ret0
}

// InsertActionGroup inserts an action group into the list of action groups
// associated with @manager. Actions in earlier groups hide actions with the
// same name in later groups.
//
// If @pos is larger than the number of action groups in @manager, or
// negative, @action_group will be inserted at the end of the internal list.
func (m uiManager) InsertActionGroup(actionGroup ActionGroup, pos int) {
	var arg0 *C.GtkUIManager
	var arg1 *C.GtkActionGroup
	var arg2 C.gint

	arg0 = (*C.GtkUIManager)(m.Native())
	arg1 = (*C.GtkActionGroup)(actionGroup.Native())
	arg2 = C.gint(pos)

	C.gtk_ui_manager_insert_action_group(arg0, arg1, arg2)
}

// NewMergeID returns an unused merge id, suitable for use with
// gtk_ui_manager_add_ui().
func (m uiManager) NewMergeID() uint {
	var arg0 *C.GtkUIManager

	arg0 = (*C.GtkUIManager)(m.Native())

	ret := C.gtk_ui_manager_new_merge_id(arg0)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// RemoveActionGroup removes an action group from the list of action groups
// associated with @manager.
func (m uiManager) RemoveActionGroup(actionGroup ActionGroup) {
	var arg0 *C.GtkUIManager
	var arg1 *C.GtkActionGroup

	arg0 = (*C.GtkUIManager)(m.Native())
	arg1 = (*C.GtkActionGroup)(actionGroup.Native())

	C.gtk_ui_manager_remove_action_group(arg0, arg1)
}

// RemoveUi unmerges the part of @manager's content identified by @merge_id.
func (m uiManager) RemoveUi(mergeID uint) {
	var arg0 *C.GtkUIManager
	var arg1 C.guint

	arg0 = (*C.GtkUIManager)(m.Native())
	arg1 = C.guint(mergeID)

	C.gtk_ui_manager_remove_ui(arg0, arg1)
}

// SetAddTearoffs sets the “add_tearoffs” property, which controls whether
// menus generated by this UIManager will have tearoff menu items.
//
// Note that this only affects regular menus. Generated popup menus never
// have tearoff menu items.
func (m uiManager) SetAddTearoffs(addTearoffs bool) {
	var arg0 *C.GtkUIManager
	var arg1 C.gboolean

	arg0 = (*C.GtkUIManager)(m.Native())
	if addTearoffs {
		arg1 = C.TRUE
	}

	C.gtk_ui_manager_set_add_tearoffs(arg0, arg1)
}

type UIManagerPrivate struct {
	native C.GtkUIManagerPrivate
}

// WrapUIManagerPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapUIManagerPrivate(ptr unsafe.Pointer) *UIManagerPrivate {
	if ptr == nil {
		return nil
	}

	return (*UIManagerPrivate)(ptr)
}

func marshalUIManagerPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapUIManagerPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (u *UIManagerPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&u.native)
}
