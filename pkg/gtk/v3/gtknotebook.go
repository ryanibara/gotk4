// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_notebook_get_type()), F: marshalNotebook},
	})
}

// Notebook: the Notebook widget is a Container whose children are pages that
// can be switched between using tab labels along one edge.
//
// There are many configuration options for GtkNotebook. Among other things, you
// can choose on which edge the tabs appear (see gtk_notebook_set_tab_pos()),
// whether, if there are too many tabs to fit the notebook should be made bigger
// or scrolling arrows added (see gtk_notebook_set_scrollable()), and whether
// there will be a popup menu allowing the users to switch pages. (see
// gtk_notebook_popup_enable(), gtk_notebook_popup_disable())
//
//
// GtkNotebook as GtkBuildable
//
// The GtkNotebook implementation of the Buildable interface supports placing
// children into tabs by specifying “tab” as the “type” attribute of a <child>
// element. Note that the content of the tab must be created before the tab can
// be filled. A tab child can be specified without specifying a <child> type
// attribute.
//
// To add a child widget in the notebooks action area, specify "action-start" or
// “action-end” as the “type” attribute of the <child> element.
//
// An example of a UI definition fragment with GtkNotebook:
//
//    <object class="GtkNotebook">
//      <child>
//        <object class="GtkLabel" id="notebook-content">
//          <property name="label">Content</property>
//        </object>
//      </child>
//      <child type="tab">
//        <object class="GtkLabel" id="notebook-tab">
//          <property name="label">Tab</property>
//        </object>
//      </child>
//    </object>
//
// CSS nodes
//
//    notebook
//    ├── header.top
//    │   ├── [<action widget>]
//    │   ├── tabs
//    │   │   ├── [arrow]
//    │   │   ├── tab
//    │   │   │   ╰── <tab label>
//    ┊   ┊   ┊
//    │   │   ├── tab[.reorderable-page]
//    │   │   │   ╰── <tab label>
//    │   │   ╰── [arrow]
//    │   ╰── [<action widget>]
//    │
//    ╰── stack
//        ├── <child>
//        ┊
//        ╰── <child>
//
// GtkNotebook has a main CSS node with name notebook, a subnode with name
// header and below that a subnode with name tabs which contains one subnode per
// tab with name tab.
//
// If action widgets are present, their CSS nodes are placed next to the tabs
// node. If the notebook is scrollable, CSS nodes with name arrow are placed as
// first and last child of the tabs node.
//
// The main node gets the .frame style class when the notebook has a border (see
// gtk_notebook_set_show_border()).
//
// The header node gets one of the style class .top, .bottom, .left or .right,
// depending on where the tabs are placed. For reorderable pages, the tab node
// gets the .reorderable-page class.
//
// A tab node gets the .dnd style class while it is moved with drag-and-drop.
//
// The nodes are always arranged from left-to-right, regarldess of text
// direction.
type Notebook interface {
	Container

	AppendPageNotebook(child Widget, tabLabel Widget) int

	AppendPageMenuNotebook(child Widget, tabLabel Widget, menuLabel Widget) int

	DetachTabNotebook(child Widget)

	ActionWidget(packType PackType) Widget

	CurrentPage() int

	GroupName() string

	MenuLabel(child Widget) Widget

	MenuLabelText(child Widget) string

	NPages() int

	NthPage(pageNum int) Widget

	Scrollable() bool

	ShowBorder() bool

	ShowTabs() bool

	TabDetachable(child Widget) bool

	TabHborder() uint16

	TabLabel(child Widget) Widget

	TabLabelText(child Widget) string

	TabPos() PositionType

	TabReorderable(child Widget) bool

	TabVborder() uint16

	InsertPageNotebook(child Widget, tabLabel Widget, position int) int

	InsertPageMenuNotebook(child Widget, tabLabel Widget, menuLabel Widget, position int) int

	NextPageNotebook()

	PageNumNotebook(child Widget) int

	PopupDisableNotebook()

	PopupEnableNotebook()

	PrependPageNotebook(child Widget, tabLabel Widget) int

	PrependPageMenuNotebook(child Widget, tabLabel Widget, menuLabel Widget) int

	PrevPageNotebook()

	RemovePageNotebook(pageNum int)

	ReorderChildNotebook(child Widget, position int)

	SetActionWidgetNotebook(widget Widget, packType PackType)

	SetCurrentPageNotebook(pageNum int)

	SetGroupNameNotebook(groupName string)

	SetMenuLabelNotebook(child Widget, menuLabel Widget)

	SetMenuLabelTextNotebook(child Widget, menuText string)

	SetScrollableNotebook(scrollable bool)

	SetShowBorderNotebook(showBorder bool)

	SetShowTabsNotebook(showTabs bool)

	SetTabDetachableNotebook(child Widget, detachable bool)

	SetTabLabelNotebook(child Widget, tabLabel Widget)

	SetTabLabelTextNotebook(child Widget, tabText string)

	SetTabPosNotebook(pos PositionType)

	SetTabReorderableNotebook(child Widget, reorderable bool)
}

// notebook implements the Notebook class.
type notebook struct {
	Container
}

// WrapNotebook wraps a GObject to the right type. It is
// primarily used internally.
func WrapNotebook(obj *externglib.Object) Notebook {
	return notebook{
		Container: WrapContainer(obj),
	}
}

func marshalNotebook(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNotebook(obj), nil
}

func NewNotebook() Notebook {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_notebook_new()

	var _notebook Notebook // out

	_notebook = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Notebook)

	return _notebook
}

func (n notebook) AppendPageNotebook(child Widget, tabLabel Widget) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))

	_cret = C.gtk_notebook_append_page(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (n notebook) AppendPageMenuNotebook(child Widget, tabLabel Widget, menuLabel Widget) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _arg3 *C.GtkWidget   // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	_arg3 = (*C.GtkWidget)(unsafe.Pointer(menuLabel.Native()))

	_cret = C.gtk_notebook_append_page_menu(_arg0, _arg1, _arg2, _arg3)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (n notebook) DetachTabNotebook(child Widget) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_notebook_detach_tab(_arg0, _arg1)
}

func (n notebook) ActionWidget(packType PackType) Widget {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.GtkPackType  // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = C.GtkPackType(packType)

	_cret = C.gtk_notebook_get_action_widget(_arg0, _arg1)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (n notebook) CurrentPage() int {
	var _arg0 *C.GtkNotebook // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	_cret = C.gtk_notebook_get_current_page(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (n notebook) GroupName() string {
	var _arg0 *C.GtkNotebook // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	_cret = C.gtk_notebook_get_group_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (n notebook) MenuLabel(child Widget) Widget {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_menu_label(_arg0, _arg1)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (n notebook) MenuLabelText(child Widget) string {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_menu_label_text(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (n notebook) NPages() int {
	var _arg0 *C.GtkNotebook // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	_cret = C.gtk_notebook_get_n_pages(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (n notebook) NthPage(pageNum int) Widget {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gint         // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = C.gint(pageNum)

	_cret = C.gtk_notebook_get_nth_page(_arg0, _arg1)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (n notebook) Scrollable() bool {
	var _arg0 *C.GtkNotebook // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	_cret = C.gtk_notebook_get_scrollable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (n notebook) ShowBorder() bool {
	var _arg0 *C.GtkNotebook // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	_cret = C.gtk_notebook_get_show_border(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (n notebook) ShowTabs() bool {
	var _arg0 *C.GtkNotebook // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	_cret = C.gtk_notebook_get_show_tabs(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (n notebook) TabDetachable(child Widget) bool {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_tab_detachable(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (n notebook) TabHborder() uint16 {
	var _arg0 *C.GtkNotebook // out
	var _cret C.guint16      // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	_cret = C.gtk_notebook_get_tab_hborder(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

func (n notebook) TabLabel(child Widget) Widget {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_tab_label(_arg0, _arg1)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (n notebook) TabLabelText(child Widget) string {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_tab_label_text(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (n notebook) TabPos() PositionType {
	var _arg0 *C.GtkNotebook    // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	_cret = C.gtk_notebook_get_tab_pos(_arg0)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

func (n notebook) TabReorderable(child Widget) bool {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_tab_reorderable(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (n notebook) TabVborder() uint16 {
	var _arg0 *C.GtkNotebook // out
	var _cret C.guint16      // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	_cret = C.gtk_notebook_get_tab_vborder(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

func (n notebook) InsertPageNotebook(child Widget, tabLabel Widget, position int) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _arg3 C.gint         // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	_arg3 = C.gint(position)

	_cret = C.gtk_notebook_insert_page(_arg0, _arg1, _arg2, _arg3)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (n notebook) InsertPageMenuNotebook(child Widget, tabLabel Widget, menuLabel Widget, position int) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _arg3 *C.GtkWidget   // out
	var _arg4 C.gint         // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	_arg3 = (*C.GtkWidget)(unsafe.Pointer(menuLabel.Native()))
	_arg4 = C.gint(position)

	_cret = C.gtk_notebook_insert_page_menu(_arg0, _arg1, _arg2, _arg3, _arg4)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (n notebook) NextPageNotebook() {
	var _arg0 *C.GtkNotebook // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	C.gtk_notebook_next_page(_arg0)
}

func (n notebook) PageNumNotebook(child Widget) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_page_num(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (n notebook) PopupDisableNotebook() {
	var _arg0 *C.GtkNotebook // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	C.gtk_notebook_popup_disable(_arg0)
}

func (n notebook) PopupEnableNotebook() {
	var _arg0 *C.GtkNotebook // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	C.gtk_notebook_popup_enable(_arg0)
}

func (n notebook) PrependPageNotebook(child Widget, tabLabel Widget) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))

	_cret = C.gtk_notebook_prepend_page(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (n notebook) PrependPageMenuNotebook(child Widget, tabLabel Widget, menuLabel Widget) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _arg3 *C.GtkWidget   // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	_arg3 = (*C.GtkWidget)(unsafe.Pointer(menuLabel.Native()))

	_cret = C.gtk_notebook_prepend_page_menu(_arg0, _arg1, _arg2, _arg3)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (n notebook) PrevPageNotebook() {
	var _arg0 *C.GtkNotebook // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))

	C.gtk_notebook_prev_page(_arg0)
}

func (n notebook) RemovePageNotebook(pageNum int) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = C.gint(pageNum)

	C.gtk_notebook_remove_page(_arg0, _arg1)
}

func (n notebook) ReorderChildNotebook(child Widget, position int) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.gint         // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.gint(position)

	C.gtk_notebook_reorder_child(_arg0, _arg1, _arg2)
}

func (n notebook) SetActionWidgetNotebook(widget Widget, packType PackType) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.GtkPackType  // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.GtkPackType(packType)

	C.gtk_notebook_set_action_widget(_arg0, _arg1, _arg2)
}

func (n notebook) SetCurrentPageNotebook(pageNum int) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = C.gint(pageNum)

	C.gtk_notebook_set_current_page(_arg0, _arg1)
}

func (n notebook) SetGroupNameNotebook(groupName string) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.gchar)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_notebook_set_group_name(_arg0, _arg1)
}

func (n notebook) SetMenuLabelNotebook(child Widget, menuLabel Widget) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(menuLabel.Native()))

	C.gtk_notebook_set_menu_label(_arg0, _arg1, _arg2)
}

func (n notebook) SetMenuLabelTextNotebook(child Widget, menuText string) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.gchar       // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.gchar)(C.CString(menuText))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_notebook_set_menu_label_text(_arg0, _arg1, _arg2)
}

func (n notebook) SetScrollableNotebook(scrollable bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	if scrollable {
		_arg1 = C.TRUE
	}

	C.gtk_notebook_set_scrollable(_arg0, _arg1)
}

func (n notebook) SetShowBorderNotebook(showBorder bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	if showBorder {
		_arg1 = C.TRUE
	}

	C.gtk_notebook_set_show_border(_arg0, _arg1)
}

func (n notebook) SetShowTabsNotebook(showTabs bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	if showTabs {
		_arg1 = C.TRUE
	}

	C.gtk_notebook_set_show_tabs(_arg0, _arg1)
}

func (n notebook) SetTabDetachableNotebook(child Widget, detachable bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if detachable {
		_arg2 = C.TRUE
	}

	C.gtk_notebook_set_tab_detachable(_arg0, _arg1, _arg2)
}

func (n notebook) SetTabLabelNotebook(child Widget, tabLabel Widget) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))

	C.gtk_notebook_set_tab_label(_arg0, _arg1, _arg2)
}

func (n notebook) SetTabLabelTextNotebook(child Widget, tabText string) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.gchar       // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.gchar)(C.CString(tabText))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_notebook_set_tab_label_text(_arg0, _arg1, _arg2)
}

func (n notebook) SetTabPosNotebook(pos PositionType) {
	var _arg0 *C.GtkNotebook    // out
	var _arg1 C.GtkPositionType // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = C.GtkPositionType(pos)

	C.gtk_notebook_set_tab_pos(_arg0, _arg1)
}

func (n notebook) SetTabReorderableNotebook(child Widget, reorderable bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if reorderable {
		_arg2 = C.TRUE
	}

	C.gtk_notebook_set_tab_reorderable(_arg0, _arg1, _arg2)
}

func (b notebook) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b notebook) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b notebook) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b notebook) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b notebook) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b notebook) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b notebook) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b notebook) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b notebook) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b notebook) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
