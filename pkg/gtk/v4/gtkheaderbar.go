// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_header_bar_get_type()), F: marshalHeaderBarrer},
	})
}

// HeaderBar: GtkHeaderBar is a widget for creating custom title bars for
// windows.
//
// !An example GtkHeaderBar (headerbar.png)
//
// GtkHeaderBar is similar to a horizontal GtkCenterBox. It allows children to
// be placed at the start or the end. In addition, it allows the window title to
// be displayed. The title will be centered with respect to the width of the
// box, even if the children at either side take up different amounts of space.
//
// GtkHeaderBar can add typical window frame controls, such as minimize,
// maximize and close buttons, or the window icon.
//
// For these reasons, GtkHeaderBar is the natural choice for use as the custom
// titlebar widget of a `GtkWindow (see gtk.Window.SetTitlebar()), as it gives
// features typical of titlebars while allowing the addition of child widgets.
//
//
// GtkHeaderBar as GtkBuildable
//
// The GtkHeaderBar implementation of the GtkBuildable interface supports adding
// children at the start or end sides by specifying “start” or “end” as the
// “type” attribute of a <child> element, or setting the title widget by
// specifying “title” value.
//
// By default the GtkHeaderBar uses a GtkLabel displaying the title of the
// window it is contained in as the title widget, equivalent to the following UI
// definition:
//
//    <object class="GtkHeaderBar">
//      <property name="title-widget">
//        <object class="GtkLabel">
//          <property name="label" translatable="yes">Label</property>
//          <property name="single-line-mode">True</property>
//          <property name="ellipsize">end</property>
//          <property name="width-chars">5</property>
//          <style>
//            <class name="title"/>
//          </style>
//        </object>
//      </property>
//    </object>
//
//
// CSS nodes
//
//    headerbar
//    ╰── windowhandle
//        ╰── box
//            ├── box.start
//            │   ├── windowcontrols.start
//            │   ╰── [other children]
//            ├── [Title Widget]
//            ╰── box.end
//                ├── [other children]
//                ╰── windowcontrols.end
//
//
// A GtkHeaderBar's CSS node is called headerbar. It contains a windowhandle
// subnode, which contains a box subnode, which contains two box subnodes at the
// start and end of the header bar, as well as a center node that represents the
// title.
//
// Each of the boxes contains a windowcontrols subnode, see gtk.WindowControls
// for details, as well as other children.
//
//
// Accessibility
//
// GtkHeaderBar uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type HeaderBar struct {
	Widget
}

func wrapHeaderBar(obj *externglib.Object) *HeaderBar {
	return &HeaderBar{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalHeaderBarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapHeaderBar(obj), nil
}

// NewHeaderBar creates a new GtkHeaderBar widget.
func NewHeaderBar() *HeaderBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_header_bar_new()

	var _headerBar *HeaderBar // out

	_headerBar = wrapHeaderBar(externglib.Take(unsafe.Pointer(_cret)))

	return _headerBar
}

// DecorationLayout gets the decoration layout of the GtkHeaderBar.
func (bar *HeaderBar) DecorationLayout() string {
	var _arg0 *C.GtkHeaderBar // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(bar.Native()))

	_cret = C.gtk_header_bar_get_decoration_layout(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ShowTitleButtons returns whether this header bar shows the standard window
// title buttons.
func (bar *HeaderBar) ShowTitleButtons() bool {
	var _arg0 *C.GtkHeaderBar // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(bar.Native()))

	_cret = C.gtk_header_bar_get_show_title_buttons(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TitleWidget retrieves the title widget of the header.
//
// See gtk.HeaderBar.SetTitleWidget().
func (bar *HeaderBar) TitleWidget() Widgetter {
	var _arg0 *C.GtkHeaderBar // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(bar.Native()))

	_cret = C.gtk_header_bar_get_title_widget(_arg0)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// PackEnd adds child to bar, packed with reference to the end of the bar.
func (bar *HeaderBar) PackEnd(child Widgetter) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(bar.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_header_bar_pack_end(_arg0, _arg1)
}

// PackStart adds child to bar, packed with reference to the start of the bar.
func (bar *HeaderBar) PackStart(child Widgetter) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(bar.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_header_bar_pack_start(_arg0, _arg1)
}

// Remove removes a child from the GtkHeaderBar.
//
// The child must have been added with gtk.HeaderBar.PackStart(),
// gtk.HeaderBar.PackEnd() or gtk.HeaderBar.SetTitleWidget().
func (bar *HeaderBar) Remove(child Widgetter) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(bar.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_header_bar_remove(_arg0, _arg1)
}

// SetDecorationLayout sets the decoration layout for this header bar.
//
// This property overrides the gtk.Settings:gtk-decoration-layout setting.
//
// There can be valid reasons for overriding the setting, such as a header bar
// design that does not allow for buttons to take room on the right, or only
// offers room for a single close button. Split header bars are another example
// for overriding the setting.
//
// The format of the string is button names, separated by commas. A colon
// separates the buttons that should appear on the left from those on the right.
// Recognized button names are minimize, maximize, close and icon (the window
// icon).
//
// For example, “icon:minimize,maximize,close” specifies a icon on the left, and
// minimize, maximize and close buttons on the right.
func (bar *HeaderBar) SetDecorationLayout(layout string) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(bar.Native()))
	if layout != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(layout)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_header_bar_set_decoration_layout(_arg0, _arg1)
}

// SetShowTitleButtons sets whether this header bar shows the standard window
// title buttons.
func (bar *HeaderBar) SetShowTitleButtons(setting bool) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(bar.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_header_bar_set_show_title_buttons(_arg0, _arg1)
}

// SetTitleWidget sets the title for the GtkHeaderBar.
//
// When set to NULL, the headerbar will display the title of the window it is
// contained in.
//
// The title should help a user identify the current view. To achieve the same
// style as the builtin title, use the “title” style class.
//
// You should set the title widget to NULL, for the window title label to be
// visible again.
func (bar *HeaderBar) SetTitleWidget(titleWidget Widgetter) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(bar.Native()))
	if titleWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(titleWidget.Native()))
	}

	C.gtk_header_bar_set_title_widget(_arg0, _arg1)
}
