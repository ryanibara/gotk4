// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_header_bar_get_type()), F: marshalHeaderBar},
	})
}

// HeaderBar: `GtkHeaderBar` is a widget for creating custom title bars for
// windows.
//
// !An example GtkHeaderBar (headerbar.png)
//
// `GtkHeaderBar` is similar to a horizontal `GtkCenterBox`. It allows children
// to be placed at the start or the end. In addition, it allows the window title
// to be displayed. The title will be centered with respect to the width of the
// box, even if the children at either side take up different amounts of space.
//
// `GtkHeaderBar` can add typical window frame controls, such as minimize,
// maximize and close buttons, or the window icon.
//
// For these reasons, `GtkHeaderBar` is the natural choice for use as the custom
// titlebar widget of a `GtkWindow (see [method@Gtk.Window.set_titlebar]), as it
// gives features typical of titlebars while allowing the addition of child
// widgets.
//
//
// GtkHeaderBar as GtkBuildable
//
// The `GtkHeaderBar` implementation of the `GtkBuildable` interface supports
// adding children at the start or end sides by specifying “start” or “end” as
// the “type” attribute of a <child> element, or setting the title widget by
// specifying “title” value.
//
// By default the `GtkHeaderBar` uses a `GtkLabel` displaying the title of the
// window it is contained in as the title widget, equivalent to the following UI
// definition:
//
// “`xml <object class="GtkHeaderBar"> <property name="title-widget"> <object
// class="GtkLabel"> <property name="label" translatable="yes">Label</property>
// <property name="single-line-mode">True</property> <property
// name="ellipsize">end</property> <property name="width-chars">5</property>
// <style> <class name="title"/> </style> </object> </property> </object> “`
//
//
// CSS nodes
//
// “` headerbar ╰── windowhandle ╰── box ├── box.start │ ├──
// windowcontrols.start │ ╰── [other children] ├── [Title Widget] ╰── box.end
// ├── [other children] ╰── windowcontrols.end “`
//
// A `GtkHeaderBar`'s CSS node is called `headerbar`. It contains a
// `windowhandle` subnode, which contains a `box` subnode, which contains two
// `box` subnodes at the start and end of the header bar, as well as a center
// node that represents the title.
//
// Each of the boxes contains a `windowcontrols` subnode, see
// [class@Gtk.WindowControls] for details, as well as other children.
//
//
// Accessibility
//
// `GtkHeaderBar` uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type HeaderBar interface {
	Widget

	DecorationLayout() string

	ShowTitleButtons() bool

	TitleWidget() Widget

	PackEndHeaderBar(child Widget)

	PackStartHeaderBar(child Widget)

	RemoveHeaderBar(child Widget)

	SetDecorationLayoutHeaderBar(layout string)

	SetShowTitleButtonsHeaderBar(setting bool)

	SetTitleWidgetHeaderBar(titleWidget Widget)
}

// headerBar implements the HeaderBar class.
type headerBar struct {
	Widget
}

// WrapHeaderBar wraps a GObject to the right type. It is
// primarily used internally.
func WrapHeaderBar(obj *externglib.Object) HeaderBar {
	return headerBar{
		Widget: WrapWidget(obj),
	}
}

func marshalHeaderBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapHeaderBar(obj), nil
}

func NewHeaderBar() HeaderBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_header_bar_new()

	var _headerBar HeaderBar // out

	_headerBar = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(HeaderBar)

	return _headerBar
}

func (b headerBar) DecorationLayout() string {
	var _arg0 *C.GtkHeaderBar // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_header_bar_get_decoration_layout(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (b headerBar) ShowTitleButtons() bool {
	var _arg0 *C.GtkHeaderBar // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_header_bar_get_show_title_buttons(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (b headerBar) TitleWidget() Widget {
	var _arg0 *C.GtkHeaderBar // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_header_bar_get_title_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (b headerBar) PackEndHeaderBar(child Widget) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_header_bar_pack_end(_arg0, _arg1)
}

func (b headerBar) PackStartHeaderBar(child Widget) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_header_bar_pack_start(_arg0, _arg1)
}

func (b headerBar) RemoveHeaderBar(child Widget) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_header_bar_remove(_arg0, _arg1)
}

func (b headerBar) SetDecorationLayoutHeaderBar(layout string) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.char)(C.CString(layout))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_header_bar_set_decoration_layout(_arg0, _arg1)
}

func (b headerBar) SetShowTitleButtonsHeaderBar(setting bool) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_header_bar_set_show_title_buttons(_arg0, _arg1)
}

func (b headerBar) SetTitleWidgetHeaderBar(titleWidget Widget) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(titleWidget.Native()))

	C.gtk_header_bar_set_title_widget(_arg0, _arg1)
}

func (s headerBar) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s headerBar) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s headerBar) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s headerBar) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s headerBar) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s headerBar) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s headerBar) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b headerBar) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
