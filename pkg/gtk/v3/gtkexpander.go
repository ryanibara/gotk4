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
		{T: externglib.Type(C.gtk_expander_get_type()), F: marshalExpander},
	})
}

// Expander: a Expander allows the user to hide or show its child by clicking on
// an expander triangle similar to the triangles used in a TreeView.
//
// Normally you use an expander as you would use any other descendant of Bin;
// you create the child widget and use gtk_container_add() to add it to the
// expander. When the expander is toggled, it will take care of showing and
// hiding the child automatically.
//
//
// Special Usage
//
// There are situations in which you may prefer to show and hide the expanded
// widget yourself, such as when you want to actually create the widget at
// expansion time. In this case, create a Expander but do not add a child to it.
// The expander widget has an Expander:expanded property which can be used to
// monitor its expansion state. You should watch this property with a signal
// connection as follows:
//
//    expander
//    ├── title
//    │   ├── arrow
//    │   ╰── <label widget>
//    ╰── <child>
//
// GtkExpander has three CSS nodes, the main node with the name expander, a
// subnode with name title and node below it with name arrow. The arrow of an
// expander that is showing its child gets the :checked pseudoclass added to it.
type Expander interface {
	Bin

	Expanded() bool

	Label() string

	LabelFill() bool

	LabelWidget() Widget

	ResizeToplevel() bool

	Spacing() int

	UseMarkup() bool

	UseUnderline() bool

	SetExpandedExpander(expanded bool)

	SetLabelExpander(label string)

	SetLabelFillExpander(labelFill bool)

	SetLabelWidgetExpander(labelWidget Widget)

	SetResizeToplevelExpander(resizeToplevel bool)

	SetSpacingExpander(spacing int)

	SetUseMarkupExpander(useMarkup bool)

	SetUseUnderlineExpander(useUnderline bool)
}

// expander implements the Expander class.
type expander struct {
	Bin
}

// WrapExpander wraps a GObject to the right type. It is
// primarily used internally.
func WrapExpander(obj *externglib.Object) Expander {
	return expander{
		Bin: WrapBin(obj),
	}
}

func marshalExpander(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapExpander(obj), nil
}

func NewExpander(label string) Expander {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_expander_new(_arg1)

	var _expander Expander // out

	_expander = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Expander)

	return _expander
}

func NewExpanderWithMnemonic(label string) Expander {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_expander_new_with_mnemonic(_arg1)

	var _expander Expander // out

	_expander = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Expander)

	return _expander
}

func (e expander) Expanded() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_expander_get_expanded(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e expander) Label() string {
	var _arg0 *C.GtkExpander // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_expander_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (e expander) LabelFill() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_expander_get_label_fill(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e expander) LabelWidget() Widget {
	var _arg0 *C.GtkExpander // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_expander_get_label_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (e expander) ResizeToplevel() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_expander_get_resize_toplevel(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e expander) Spacing() int {
	var _arg0 *C.GtkExpander // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_expander_get_spacing(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (e expander) UseMarkup() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_expander_get_use_markup(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e expander) UseUnderline() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_expander_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e expander) SetExpandedExpander(expanded bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if expanded {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_expanded(_arg0, _arg1)
}

func (e expander) SetLabelExpander(label string) {
	var _arg0 *C.GtkExpander // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_expander_set_label(_arg0, _arg1)
}

func (e expander) SetLabelFillExpander(labelFill bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if labelFill {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_label_fill(_arg0, _arg1)
}

func (e expander) SetLabelWidgetExpander(labelWidget Widget) {
	var _arg0 *C.GtkExpander // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))

	C.gtk_expander_set_label_widget(_arg0, _arg1)
}

func (e expander) SetResizeToplevelExpander(resizeToplevel bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if resizeToplevel {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_resize_toplevel(_arg0, _arg1)
}

func (e expander) SetSpacingExpander(spacing int) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	_arg1 = C.gint(spacing)

	C.gtk_expander_set_spacing(_arg0, _arg1)
}

func (e expander) SetUseMarkupExpander(useMarkup bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if useMarkup {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_use_markup(_arg0, _arg1)
}

func (e expander) SetUseUnderlineExpander(useUnderline bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_use_underline(_arg0, _arg1)
}

func (b expander) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b expander) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b expander) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b expander) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b expander) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b expander) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b expander) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b expander) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b expander) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b expander) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
