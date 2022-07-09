// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeBox returns the GType for the type Box.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeBox() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Box").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalBox)
	return gtype
}

// BoxOverrider contains methods that are overridable.
type BoxOverrider interface {
}

// Box widget arranges child widgets into a single row or column, depending upon
// the value of its Orientable:orientation property. Within the other dimension,
// all children are allocated the same size. Of course, the Widget:halign and
// Widget:valign properties can be used on the children to influence their
// allocation.
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
type Box struct {
	_ [0]func() // equal guard
	Container

	*coreglib.Object
	Orientable
}

var (
	_ Containerer       = (*Box)(nil)
	_ coreglib.Objector = (*Box)(nil)
)

func classInitBoxer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapBox(obj *coreglib.Object) *Box {
	return &Box{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalBox(p uintptr) (interface{}, error) {
	return wrapBox(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CenterWidget retrieves the center widget of the box.
//
// The function returns the following values:
//
//    - widget (optional): center widget or NULL in case no center widget is set.
//
func (box *Box) CenterWidget() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))

	_gret := girepository.MustFind("Gtk", "Box").InvokeMethod("get_center_widget", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(box)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
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

// Homogeneous returns whether the box is homogeneous (all children are the same
// size). See gtk_box_set_homogeneous().
//
// The function returns the following values:
//
//    - ok: TRUE if the box is homogeneous.
//
func (box *Box) Homogeneous() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))

	_gret := girepository.MustFind("Gtk", "Box").InvokeMethod("get_homogeneous", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(box)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Spacing gets the value set by gtk_box_set_spacing().
//
// The function returns the following values:
//
//    - gint: spacing between children.
//
func (box *Box) Spacing() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))

	_gret := girepository.MustFind("Gtk", "Box").InvokeMethod("get_spacing", _args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(box)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// PackEnd adds child to box, packed with reference to the end of box. The child
// is packed after (away from end of) any other child packed with reference to
// the end of box.
//
// The function takes the following parameters:
//
//    - child to be added to box.
//    - expand: TRUE if the new child is to be given extra space allocated to
//      box. The extra space will be divided evenly between all children of box
//      that use this option.
//    - fill: TRUE if space given to child by the expand option is actually
//      allocated to child, rather than just padding it. This parameter has no
//      effect if expand is set to FALSE. A child is always allocated the full
//      height of a horizontal Box and the full width of a vertical Box. This
//      option affects the other dimension.
//    - padding: extra space in pixels to put between this child and its
//      neighbors, over and above the global amount specified by Box:spacing
//      property. If child is a widget at one of the reference ends of box, then
//      padding pixels are also put between child and the reference edge of box.
//
func (box *Box) PackEnd(child Widgetter, expand, fill bool, padding uint32) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	if expand {
		*(*C.gboolean)(unsafe.Pointer(&_args[2])) = C.TRUE
	}
	if fill {
		*(*C.gboolean)(unsafe.Pointer(&_args[3])) = C.TRUE
	}
	*(*C.guint)(unsafe.Pointer(&_args[4])) = C.guint(padding)

	girepository.MustFind("Gtk", "Box").InvokeMethod("pack_end", _args[:], nil)

	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
	runtime.KeepAlive(expand)
	runtime.KeepAlive(fill)
	runtime.KeepAlive(padding)
}

// PackStart adds child to box, packed with reference to the start of box. The
// child is packed after any other child packed with reference to the start of
// box.
//
// The function takes the following parameters:
//
//    - child to be added to box.
//    - expand: TRUE if the new child is to be given extra space allocated to
//      box. The extra space will be divided evenly between all children that use
//      this option.
//    - fill: TRUE if space given to child by the expand option is actually
//      allocated to child, rather than just padding it. This parameter has no
//      effect if expand is set to FALSE. A child is always allocated the full
//      height of a horizontal Box and the full width of a vertical Box. This
//      option affects the other dimension.
//    - padding: extra space in pixels to put between this child and its
//      neighbors, over and above the global amount specified by Box:spacing
//      property. If child is a widget at one of the reference ends of box, then
//      padding pixels are also put between child and the reference edge of box.
//
func (box *Box) PackStart(child Widgetter, expand, fill bool, padding uint32) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	if expand {
		*(*C.gboolean)(unsafe.Pointer(&_args[2])) = C.TRUE
	}
	if fill {
		*(*C.gboolean)(unsafe.Pointer(&_args[3])) = C.TRUE
	}
	*(*C.guint)(unsafe.Pointer(&_args[4])) = C.guint(padding)

	girepository.MustFind("Gtk", "Box").InvokeMethod("pack_start", _args[:], nil)

	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
	runtime.KeepAlive(expand)
	runtime.KeepAlive(fill)
	runtime.KeepAlive(padding)
}

// QueryChildPacking obtains information about how child is packed into box.
//
// The function takes the following parameters:
//
//    - child of the child to query.
//
// The function returns the following values:
//
//    - expand: pointer to return location for expand child property.
//    - fill: pointer to return location for fill child property.
//    - padding: pointer to return location for padding child property.
//    - packType: pointer to return location for pack-type child property.
//
func (box *Box) QueryChildPacking(child Widgetter) (expand, fill bool, padding uint32, packType PackType) {
	var _args [2]girepository.Argument
	var _outs [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	girepository.MustFind("Gtk", "Box").InvokeMethod("query_child_packing", _args[:], _outs[:])

	runtime.KeepAlive(box)
	runtime.KeepAlive(child)

	var _expand bool       // out
	var _fill bool         // out
	var _padding uint32    // out
	var _packType PackType // out

	if **(**C.void)(unsafe.Pointer(&_outs[0])) != 0 {
		_expand = true
	}
	if **(**C.void)(unsafe.Pointer(&_outs[1])) != 0 {
		_fill = true
	}
	_padding = *(*uint32)(unsafe.Pointer(_outs[2]))
	_packType = *(*PackType)(unsafe.Pointer(_outs[3]))

	return _expand, _fill, _padding, _packType
}

// ReorderChild moves child to a new position in the list of box children. The
// list contains widgets packed K_PACK_START as well as widgets packed
// K_PACK_END, in the order that these widgets were added to box.
//
// A widget’s position in the box children list determines where the widget is
// packed into box. A child widget at some position in the list will be packed
// just after all other widgets of the same packing type that appear earlier in
// the list.
//
// The function takes the following parameters:
//
//    - child to move.
//    - position: new position for child in the list of children of box, starting
//      from 0. If negative, indicates the end of the list.
//
func (box *Box) ReorderChild(child Widgetter, position int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(position)

	girepository.MustFind("Gtk", "Box").InvokeMethod("reorder_child", _args[:], nil)

	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
	runtime.KeepAlive(position)
}

// SetCenterWidget sets a center widget; that is a child widget that will be
// centered with respect to the full width of the box, even if the children at
// either side take up different amounts of space.
//
// The function takes the following parameters:
//
//    - widget (optional) to center.
//
func (box *Box) SetCenterWidget(widget Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	if widget != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	}

	girepository.MustFind("Gtk", "Box").InvokeMethod("set_center_widget", _args[:], nil)

	runtime.KeepAlive(box)
	runtime.KeepAlive(widget)
}

// SetHomogeneous sets the Box:homogeneous property of box, controlling whether
// or not all children of box are given equal space in the box.
//
// The function takes the following parameters:
//
//    - homogeneous: boolean value, TRUE to create equal allotments, FALSE for
//      variable allotments.
//
func (box *Box) SetHomogeneous(homogeneous bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	if homogeneous {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Box").InvokeMethod("set_homogeneous", _args[:], nil)

	runtime.KeepAlive(box)
	runtime.KeepAlive(homogeneous)
}

// SetSpacing sets the Box:spacing property of box, which is the number of
// pixels to place between children of box.
//
// The function takes the following parameters:
//
//    - spacing: number of pixels to put between children.
//
func (box *Box) SetSpacing(spacing int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(spacing)

	girepository.MustFind("Gtk", "Box").InvokeMethod("set_spacing", _args[:], nil)

	runtime.KeepAlive(box)
	runtime.KeepAlive(spacing)
}
