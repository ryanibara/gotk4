// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeHeaderBar = coreglib.Type(C.gtk_header_bar_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeHeaderBar, F: marshalHeaderBar},
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
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*HeaderBar)(nil)
)

func wrapHeaderBar(obj *coreglib.Object) *HeaderBar {
	return &HeaderBar{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
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
	}
}

func marshalHeaderBar(p uintptr) (interface{}, error) {
	return wrapHeaderBar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewHeaderBar creates a new GtkHeaderBar widget.
//
// The function returns the following values:
//
//    - headerBar: new GtkHeaderBar.
//
func NewHeaderBar() *HeaderBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_header_bar_new()

	var _headerBar *HeaderBar // out

	_headerBar = wrapHeaderBar(coreglib.Take(unsafe.Pointer(_cret)))

	return _headerBar
}

// DecorationLayout gets the decoration layout of the GtkHeaderBar.
//
// The function returns the following values:
//
//    - utf8 (optional): decoration layout.
//
func (bar *HeaderBar) DecorationLayout() string {
	var _arg0 *C.GtkHeaderBar // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_cret = C.gtk_header_bar_get_decoration_layout(_arg0)
	runtime.KeepAlive(bar)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ShowTitleButtons returns whether this header bar shows the standard window
// title buttons.
//
// The function returns the following values:
//
//    - ok: TRUE if title buttons are shown.
//
func (bar *HeaderBar) ShowTitleButtons() bool {
	var _arg0 *C.GtkHeaderBar // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_cret = C.gtk_header_bar_get_show_title_buttons(_arg0)
	runtime.KeepAlive(bar)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TitleWidget retrieves the title widget of the header.
//
// See gtk.HeaderBar.SetTitleWidget().
//
// The function returns the following values:
//
//    - widget (optional): title widget of the header, or NULL if none has been
//      set explicitly.
//
func (bar *HeaderBar) TitleWidget() Widgetter {
	var _arg0 *C.GtkHeaderBar // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_cret = C.gtk_header_bar_get_title_widget(_arg0)
	runtime.KeepAlive(bar)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// PackEnd adds child to bar, packed with reference to the end of the bar.
//
// The function takes the following parameters:
//
//    - child to be added to bar.
//
func (bar *HeaderBar) PackEnd(child Widgetter) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_header_bar_pack_end(_arg0, _arg1)
	runtime.KeepAlive(bar)
	runtime.KeepAlive(child)
}

// PackStart adds child to bar, packed with reference to the start of the bar.
//
// The function takes the following parameters:
//
//    - child to be added to bar.
//
func (bar *HeaderBar) PackStart(child Widgetter) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_header_bar_pack_start(_arg0, _arg1)
	runtime.KeepAlive(bar)
	runtime.KeepAlive(child)
}

// Remove removes a child from the GtkHeaderBar.
//
// The child must have been added with gtk.HeaderBar.PackStart(),
// gtk.HeaderBar.PackEnd() or gtk.HeaderBar.SetTitleWidget().
//
// The function takes the following parameters:
//
//    - child to remove.
//
func (bar *HeaderBar) Remove(child Widgetter) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_header_bar_remove(_arg0, _arg1)
	runtime.KeepAlive(bar)
	runtime.KeepAlive(child)
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
//
// The function takes the following parameters:
//
//    - layout (optional): decoration layout, or NULL to unset the layout.
//
func (bar *HeaderBar) SetDecorationLayout(layout string) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if layout != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(layout)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_header_bar_set_decoration_layout(_arg0, _arg1)
	runtime.KeepAlive(bar)
	runtime.KeepAlive(layout)
}

// SetShowTitleButtons sets whether this header bar shows the standard window
// title buttons.
//
// The function takes the following parameters:
//
//    - setting: TRUE to show standard title buttons.
//
func (bar *HeaderBar) SetShowTitleButtons(setting bool) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_header_bar_set_show_title_buttons(_arg0, _arg1)
	runtime.KeepAlive(bar)
	runtime.KeepAlive(setting)
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
//
// The function takes the following parameters:
//
//    - titleWidget (optional): widget to use for a title.
//
func (bar *HeaderBar) SetTitleWidget(titleWidget Widgetter) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if titleWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(titleWidget).Native()))
	}

	C.gtk_header_bar_set_title_widget(_arg0, _arg1)
	runtime.KeepAlive(bar)
	runtime.KeepAlive(titleWidget)
}
