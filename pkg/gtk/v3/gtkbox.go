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
		{T: externglib.Type(C.gtk_box_get_type()), F: marshalBox},
	})
}

// Box: the GtkBox widget arranges child widgets into a single row or column,
// depending upon the value of its Orientable:orientation property. Within the
// other dimension, all children are allocated the same size. Of course, the
// Widget:halign and Widget:valign properties can be used on the children to
// influence their allocation.
//
// GtkBox uses a notion of packing. Packing refers to adding widgets with
// reference to a particular position in a Container. For a GtkBox, there are
// two reference positions: the start and the end of the box. For a vertical
// Box, the start is defined as the top of the box and the end is defined as the
// bottom. For a horizontal Box the start is defined as the left side and the
// end is defined as the right side.
//
// Use repeated calls to gtk_box_pack_start() to pack widgets into a GtkBox from
// start to end. Use gtk_box_pack_end() to add widgets from end to start. You
// may intersperse these calls and add widgets from both ends of the same
// GtkBox.
//
// Because GtkBox is a Container, you may also use gtk_container_add() to insert
// widgets into the box, and they will be packed with the default values for
// expand and fill child properties. Use gtk_container_remove() to remove
// widgets from the GtkBox.
//
// Use gtk_box_set_homogeneous() to specify whether or not all children of the
// GtkBox are forced to get the same amount of space.
//
// Use gtk_box_set_spacing() to determine how much space will be minimally
// placed between all children in the GtkBox. Note that spacing is added between
// the children, while padding added by gtk_box_pack_start() or
// gtk_box_pack_end() is added on either side of the widget it belongs to.
//
// Use gtk_box_reorder_child() to move a GtkBox child to a different place in
// the box.
//
// Use gtk_box_set_child_packing() to reset the expand, fill and padding child
// properties. Use gtk_box_query_child_packing() to query these fields.
//
//
// CSS nodes
//
// GtkBox uses a single CSS node with name box.
//
// In horizontal orientation, the nodes of the children are always arranged from
// left to right. So :first-child will always select the leftmost child,
// regardless of text direction.
type Box interface {
	Container
	Orientable

	BaselinePosition() BaselinePosition

	CenterWidget() Widget

	Homogeneous() bool

	Spacing() int

	PackEndBox(child Widget, expand bool, fill bool, padding uint)

	PackStartBox(child Widget, expand bool, fill bool, padding uint)

	QueryChildPackingBox(child Widget) (expand bool, fill bool, padding uint, packType PackType)

	ReorderChildBox(child Widget, position int)

	SetBaselinePositionBox(position BaselinePosition)

	SetCenterWidgetBox(widget Widget)

	SetChildPackingBox(child Widget, expand bool, fill bool, padding uint, packType PackType)

	SetHomogeneousBox(homogeneous bool)

	SetSpacingBox(spacing int)
}

// box implements the Box class.
type box struct {
	Container
}

// WrapBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapBox(obj *externglib.Object) Box {
	return box{
		Container: WrapContainer(obj),
	}
}

func marshalBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBox(obj), nil
}

func NewBox(orientation Orientation, spacing int) Box {
	var _arg1 C.GtkOrientation // out
	var _arg2 C.gint           // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)
	_arg2 = C.gint(spacing)

	_cret = C.gtk_box_new(_arg1, _arg2)

	var _box Box // out

	_box = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Box)

	return _box
}

func (b box) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkBox             // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_box_get_baseline_position(_arg0)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

func (b box) CenterWidget() Widget {
	var _arg0 *C.GtkBox    // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_box_get_center_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (b box) Homogeneous() bool {
	var _arg0 *C.GtkBox  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_box_get_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (b box) Spacing() int {
	var _arg0 *C.GtkBox // out
	var _cret C.gint    // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_box_get_spacing(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (b box) PackEndBox(child Widget, expand bool, fill bool, padding uint) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gboolean   // out
	var _arg3 C.gboolean   // out
	var _arg4 C.guint      // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if expand {
		_arg2 = C.TRUE
	}
	if fill {
		_arg3 = C.TRUE
	}
	_arg4 = C.guint(padding)

	C.gtk_box_pack_end(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (b box) PackStartBox(child Widget, expand bool, fill bool, padding uint) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gboolean   // out
	var _arg3 C.gboolean   // out
	var _arg4 C.guint      // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if expand {
		_arg2 = C.TRUE
	}
	if fill {
		_arg3 = C.TRUE
	}
	_arg4 = C.guint(padding)

	C.gtk_box_pack_start(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (b box) QueryChildPackingBox(child Widget) (expand bool, fill bool, padding uint, packType PackType) {
	var _arg0 *C.GtkBox     // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gboolean    // in
	var _arg3 C.gboolean    // in
	var _arg4 C.guint       // in
	var _arg5 C.GtkPackType // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_box_query_child_packing(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)

	var _expand bool       // out
	var _fill bool         // out
	var _padding uint      // out
	var _packType PackType // out

	if _arg2 != 0 {
		_expand = true
	}
	if _arg3 != 0 {
		_fill = true
	}
	_padding = uint(_arg4)
	_packType = PackType(_arg5)

	return _expand, _fill, _padding, _packType
}

func (b box) ReorderChildBox(child Widget, position int) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.gint(position)

	C.gtk_box_reorder_child(_arg0, _arg1, _arg2)
}

func (b box) SetBaselinePositionBox(position BaselinePosition) {
	var _arg0 *C.GtkBox             // out
	var _arg1 C.GtkBaselinePosition // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.GtkBaselinePosition(position)

	C.gtk_box_set_baseline_position(_arg0, _arg1)
}

func (b box) SetCenterWidgetBox(widget Widget) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_box_set_center_widget(_arg0, _arg1)
}

func (b box) SetChildPackingBox(child Widget, expand bool, fill bool, padding uint, packType PackType) {
	var _arg0 *C.GtkBox     // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gboolean    // out
	var _arg3 C.gboolean    // out
	var _arg4 C.guint       // out
	var _arg5 C.GtkPackType // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if expand {
		_arg2 = C.TRUE
	}
	if fill {
		_arg3 = C.TRUE
	}
	_arg4 = C.guint(padding)
	_arg5 = C.GtkPackType(packType)

	C.gtk_box_set_child_packing(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (b box) SetHomogeneousBox(homogeneous bool) {
	var _arg0 *C.GtkBox  // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_box_set_homogeneous(_arg0, _arg1)
}

func (b box) SetSpacingBox(spacing int) {
	var _arg0 *C.GtkBox // out
	var _arg1 C.gint    // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.gint(spacing)

	C.gtk_box_set_spacing(_arg0, _arg1)
}

func (b box) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b box) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b box) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b box) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b box) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b box) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b box) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b box) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b box) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b box) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (o box) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o box) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
