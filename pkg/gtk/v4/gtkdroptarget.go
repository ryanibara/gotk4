// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_drop_target_get_type()), F: marshalDropTarget},
	})
}

// DropTarget: `GtkDropTarget` is an event controller to receive Drag-and-Drop
// operations.
//
// The most basic way to use a `GtkDropTarget` to receive drops on a widget is
// to create it via [ctor@Gtk.DropTarget.new], passing in the `GType` of the
// data you want to receive and connect to the [signal@Gtk.DropTarget::drop]
// signal to receive the data:
//
// “`c static gboolean on_drop (GtkDropTarget *target, const GValue *value,
// double x, double y, gpointer data) { MyWidget *self = data;
//
//    // Call the appropriate setter depending on the type of data
//    // that we received
//    if (G_VALUE_HOLDS (value, G_TYPE_FILE))
//      my_widget_set_file (self, g_value_get_object (value));
//    else if (G_VALUE_HOLDS (value, GDK_TYPE_PIXBUF))
//      my_widget_set_pixbuf (self, g_value_get_object (value));
//    else
//      return FALSE;
//
//    return TRUE;
//
// }
//
// static void my_widget_init (MyWidget *self) { GtkDropTarget *target =
// gtk_drop_target_new (G_TYPE_INVALID, GDK_ACTION_COPY);
//
//    // This widget accepts two types of drop types: GFile objects
//    // and GdkPixbuf objects
//    gtk_drop_target_set_gtypes (target, (GTypes [2]) {
//      G_TYPE_FILE,
//      GDK_TYPE_PIXBUF,
//    }, 2);
//
//    gtk_widget_add_controller (GTK_WIDGET (self), GTK_EVENT_CONTROLLER (target));
//
// } “`
//
// `GtkDropTarget` supports more options, such as:
//
//    * rejecting potential drops via the [signal@Gtk.DropTarget::accept] signal
//      and the [method@Gtk.DropTarget.reject] function to let other drop
//      targets handle the drop
//    * tracking an ongoing drag operation before the drop via the
//      [signal@Gtk.DropTarget::enter], [signal@Gtk.DropTarget::motion] and
//      [signal@Gtk.DropTarget::leave] signals
//    * configuring how to receive data by setting the
//      [property@Gtk.DropTarget:preload] property and listening for its
//      availability via the [property@Gtk.DropTarget:value] property
//
// However, `GtkDropTarget` is ultimately modeled in a synchronous way and only
// supports data transferred via `GType`. If you want full control over an
// ongoing drop, the [class@Gtk.DropTargetAsync] object gives you this ability.
//
// While a pointer is dragged over the drop target's widget and the drop has not
// been rejected, that widget will receive the GTK_STATE_FLAG_DROP_ACTIVE state,
// which can be used to style the widget.
//
// If you are not interested in receiving the drop, but just want to update UI
// state during a Drag-and-Drop operation (e.g. switching tabs), you can use
// [class@Gtk.DropControllerMotion].
type DropTarget interface {
	EventController

	// GTypes gets the list of supported `GTypes` for @self.
	//
	// If no type have been set, nil will be returned.
	GTypes() []externglib.Type
	// Preload gets whether data should be preloaded on hover.
	Preload() bool
	// Value gets the current drop data, as a `GValue`.
	Value() **externglib.Value
	// Reject rejects the ongoing drop operation.
	//
	// If no drop operation is ongoing, i.e when [property@Gtk.DropTarget:drop]
	// is nil, this function does nothing.
	//
	// This function should be used when delaying the decision on whether to
	// accept a drag or not until after reading the data.
	Reject()
	// SetActions sets the actions that this drop target supports.
	SetActions(actions gdk.DragAction)
	// SetGTypes sets the supported `GTypes` for this drop target.
	SetGTypes(types []externglib.Type)
	// SetPreload sets whether data should be preloaded on hover.
	SetPreload(preload bool)
}

// dropTarget implements the DropTarget interface.
type dropTarget struct {
	EventController
}

var _ DropTarget = (*dropTarget)(nil)

// WrapDropTarget wraps a GObject to the right type. It is
// primarily used internally.
func WrapDropTarget(obj *externglib.Object) DropTarget {
	return DropTarget{
		EventController: WrapEventController(obj),
	}
}

func marshalDropTarget(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDropTarget(obj), nil
}

// GTypes gets the list of supported `GTypes` for @self.
//
// If no type have been set, nil will be returned.
func (s dropTarget) GTypes() []externglib.Type {
	var _arg0 *C.GtkDropTarget

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	var _cret *C.GType
	var _arg1 *C.gsize

	_cret = C.gtk_drop_target_get_gtypes(_arg0, &_arg1)

	var _gTypes []externglib.Type

	{
		var src []C.GType
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(_arg1))

		_gTypes = make([]externglib.Type, _arg1)
		for i := 0; i < uintptr(_arg1); i++ {
			_gTypes = externglib.Type(_cret)
		}
	}

	return _gTypes
}

// Preload gets whether data should be preloaded on hover.
func (s dropTarget) Preload() bool {
	var _arg0 *C.GtkDropTarget

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	_cret = C.gtk_drop_target_get_preload(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Value gets the current drop data, as a `GValue`.
func (s dropTarget) Value() **externglib.Value {
	var _arg0 *C.GtkDropTarget

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	var _cret *C.GValue

	_cret = C.gtk_drop_target_get_value(_arg0)

	var _value **externglib.Value

	_value = externglib.ValueFromNative(unsafe.Pointer(_cret))

	return _value
}

// Reject rejects the ongoing drop operation.
//
// If no drop operation is ongoing, i.e when [property@Gtk.DropTarget:drop]
// is nil, this function does nothing.
//
// This function should be used when delaying the decision on whether to
// accept a drag or not until after reading the data.
func (s dropTarget) Reject() {
	var _arg0 *C.GtkDropTarget

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	C.gtk_drop_target_reject(_arg0)
}

// SetActions sets the actions that this drop target supports.
func (s dropTarget) SetActions(actions gdk.DragAction) {
	var _arg0 *C.GtkDropTarget
	var _arg1 C.GdkDragAction

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GdkDragAction)(actions)

	C.gtk_drop_target_set_actions(_arg0, _arg1)
}

// SetGTypes sets the supported `GTypes` for this drop target.
func (s dropTarget) SetGTypes(types []externglib.Type) {
	var _arg0 *C.GtkDropTarget
	var _arg1 *C.GType
	var _arg2 C.gsize

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))
	_arg2 = C.gsize(len(types))
	_arg1 = (*C.GType)(C.malloc(len(types) * C.sizeof_GType))
	defer C.free(unsafe.Pointer(_arg1))

	{
		var out []C.GType
		ptr.SetSlice(unsafe.Pointer(&out), unsafe.Pointer(_arg1), int(len(types)))

		for i := range types {
			_arg1 = C.GType(types)
		}
	}

	C.gtk_drop_target_set_gtypes(_arg0, _arg1, _arg2)
}

// SetPreload sets whether data should be preloaded on hover.
func (s dropTarget) SetPreload(preload bool) {
	var _arg0 *C.GtkDropTarget
	var _arg1 C.gboolean

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))
	if preload {
		_arg1 = C.gboolean(1)
	}

	C.gtk_drop_target_set_preload(_arg0, _arg1)
}
