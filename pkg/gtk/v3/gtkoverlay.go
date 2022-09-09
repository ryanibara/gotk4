// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern gboolean _gotk4_gtk3_Overlay_ConnectGetChildPosition(gpointer, GtkWidget*, GdkRectangle*, guintptr);
// extern gboolean _gotk4_gtk3_OverlayClass_get_child_position(GtkOverlay*, GtkWidget*, GtkAllocation*);
// gboolean _gotk4_gtk3_Overlay_virtual_get_child_position(void* fnptr, GtkOverlay* arg0, GtkWidget* arg1, GtkAllocation* arg2) {
//   return ((gboolean (*)(GtkOverlay*, GtkWidget*, GtkAllocation*))(fnptr))(arg0, arg1, arg2);
// };
import "C"

// GType values.
var (
	GTypeOverlay = coreglib.Type(C.gtk_overlay_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeOverlay, F: marshalOverlay},
	})
}

// OverlayOverrides contains methods that are overridable.
type OverlayOverrides struct {
	// The function takes the following parameters:
	//
	//    - widget
	//    - allocation
	//
	// The function returns the following values:
	//
	ChildPosition func(widget Widgetter, allocation *Allocation) bool
}

func defaultOverlayOverrides(v *Overlay) OverlayOverrides {
	return OverlayOverrides{
		ChildPosition: v.childPosition,
	}
}

// Overlay is a container which contains a single main child, on top of which it
// can place “overlay” widgets. The position of each overlay widget is
// determined by its Widget:halign and Widget:valign properties. E.g. a widget
// with both alignments set to GTK_ALIGN_START will be placed at the top left
// corner of the GtkOverlay container, whereas an overlay with halign set to
// GTK_ALIGN_CENTER and valign set to GTK_ALIGN_END will be placed a the bottom
// edge of the GtkOverlay, horizontally centered. The position can be adjusted
// by setting the margin properties of the child to non-zero values.
//
// More complicated placement of overlays is possible by connecting to the
// Overlay::get-child-position signal.
//
// An overlay’s minimum and natural sizes are those of its main child. The sizes
// of overlay children are not considered when measuring these preferred sizes.
//
//
// GtkOverlay as GtkBuildable
//
// The GtkOverlay implementation of the GtkBuildable interface supports placing
// a child as an overlay by specifying “overlay” as the “type” attribute of a
// <child> element.
//
//
// CSS nodes
//
// GtkOverlay has a single CSS node with the name “overlay”. Overlay children
// whose alignments cause them to be positioned at an edge get the style classes
// “.left”, “.right”, “.top”, and/or “.bottom” according to their position.
type Overlay struct {
	_ [0]func() // equal guard
	Bin
}

var (
	_ Binner = (*Overlay)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Overlay, *OverlayClass, OverlayOverrides](
		GTypeOverlay,
		initOverlayClass,
		wrapOverlay,
		defaultOverlayOverrides,
	)
}

func initOverlayClass(gclass unsafe.Pointer, overrides OverlayOverrides, classInitFunc func(*OverlayClass)) {
	pclass := (*C.GtkOverlayClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeOverlay))))

	if overrides.ChildPosition != nil {
		pclass.get_child_position = (*[0]byte)(C._gotk4_gtk3_OverlayClass_get_child_position)
	}

	if classInitFunc != nil {
		class := (*OverlayClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapOverlay(obj *coreglib.Object) *Overlay {
	return &Overlay{
		Bin: Bin{
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
		},
	}
}

func marshalOverlay(p uintptr) (interface{}, error) {
	return wrapOverlay(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectGetChildPosition signal is emitted to determine the position and size
// of any overlay child widgets. A handler for this signal should fill
// allocation with the desired position and size for widget, relative to the
// 'main' child of overlay.
//
// The default handler for this signal uses the widget's halign and valign
// properties to determine the position and gives the widget its natural size
// (except that an alignment of GTK_ALIGN_FILL will cause the overlay to be
// full-width/height). If the main child is a ScrolledWindow, the overlays are
// placed relative to its contents.
func (overlay *Overlay) ConnectGetChildPosition(f func(widget Widgetter) (allocation *gdk.Rectangle, ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(overlay, "get-child-position", false, unsafe.Pointer(C._gotk4_gtk3_Overlay_ConnectGetChildPosition), f)
}

// NewOverlay creates a new Overlay.
//
// The function returns the following values:
//
//    - overlay: new Overlay object.
//
func NewOverlay() *Overlay {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_overlay_new()

	var _overlay *Overlay // out

	_overlay = wrapOverlay(coreglib.Take(unsafe.Pointer(_cret)))

	return _overlay
}

// AddOverlay adds widget to overlay.
//
// The widget will be stacked on top of the main widget added with
// gtk_container_add().
//
// The position at which widget is placed is determined from its Widget:halign
// and Widget:valign properties.
//
// The function takes the following parameters:
//
//    - widget to be added to the container.
//
func (overlay *Overlay) AddOverlay(widget Widgetter) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(coreglib.InternObject(overlay).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	C.gtk_overlay_add_overlay(_arg0, _arg1)
	runtime.KeepAlive(overlay)
	runtime.KeepAlive(widget)
}

// OverlayPassThrough: convenience function to get the value of the
// Overlay:pass-through child property for widget.
//
// The function takes the following parameters:
//
//    - widget: overlay child of Overlay.
//
// The function returns the following values:
//
//    - ok: whether the widget is a pass through child.
//
func (overlay *Overlay) OverlayPassThrough(widget Widgetter) bool {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(coreglib.InternObject(overlay).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_cret = C.gtk_overlay_get_overlay_pass_through(_arg0, _arg1)
	runtime.KeepAlive(overlay)
	runtime.KeepAlive(widget)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ReorderOverlay moves child to a new index in the list of overlay children.
// The list contains overlays in the order that these were added to overlay by
// default. See also Overlay:index.
//
// A widget’s index in the overlay children list determines which order the
// children are drawn if they overlap. The first child is drawn at the bottom.
// It also affects the default focus chain order.
//
// The function takes the following parameters:
//
//    - child: overlaid Widget to move.
//    - index_: new index for child in the list of overlay children of overlay,
//      starting from 0. If negative, indicates the end of the list.
//
func (overlay *Overlay) ReorderOverlay(child Widgetter, index_ int) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.int         // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(coreglib.InternObject(overlay).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	_arg2 = C.int(index_)

	C.gtk_overlay_reorder_overlay(_arg0, _arg1, _arg2)
	runtime.KeepAlive(overlay)
	runtime.KeepAlive(child)
	runtime.KeepAlive(index_)
}

// SetOverlayPassThrough: convenience function to set the value of the
// Overlay:pass-through child property for widget.
//
// The function takes the following parameters:
//
//    - widget: overlay child of Overlay.
//    - passThrough: whether the child should pass the input through.
//
func (overlay *Overlay) SetOverlayPassThrough(widget Widgetter, passThrough bool) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gboolean    // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(coreglib.InternObject(overlay).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	if passThrough {
		_arg2 = C.TRUE
	}

	C.gtk_overlay_set_overlay_pass_through(_arg0, _arg1, _arg2)
	runtime.KeepAlive(overlay)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(passThrough)
}

// The function takes the following parameters:
//
//    - widget
//    - allocation
//
// The function returns the following values:
//
func (overlay *Overlay) childPosition(widget Widgetter, allocation *Allocation) bool {
	gclass := (*C.GtkOverlayClass)(coreglib.PeekParentClass(overlay))
	fnarg := gclass.get_child_position

	var _arg0 *C.GtkOverlay    // out
	var _arg1 *C.GtkWidget     // out
	var _arg2 *C.GtkAllocation // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(coreglib.InternObject(overlay).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg2 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(allocation)))
	type _ = *Allocation
	type _ = *gdk.Rectangle

	_cret = C._gotk4_gtk3_Overlay_virtual_get_child_position(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(overlay)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(allocation)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// OverlayClass: instance of this type is always passed by reference.
type OverlayClass struct {
	*overlayClass
}

// overlayClass is the struct that's finalized.
type overlayClass struct {
	native *C.GtkOverlayClass
}

// ParentClass: parent class.
func (o *OverlayClass) ParentClass() *BinClass {
	valptr := &o.native.parent_class
	var _v *BinClass // out
	_v = (*BinClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
