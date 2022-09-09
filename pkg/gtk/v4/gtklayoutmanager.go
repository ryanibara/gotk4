// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_LayoutManagerClass_unroot(GtkLayoutManager*);
// extern void _gotk4_gtk4_LayoutManagerClass_root(GtkLayoutManager*);
// extern void _gotk4_gtk4_LayoutManagerClass_measure(GtkLayoutManager*, GtkWidget*, GtkOrientation, int, int*, int*, int*, int*);
// extern void _gotk4_gtk4_LayoutManagerClass_allocate(GtkLayoutManager*, GtkWidget*, int, int, int);
// extern GtkSizeRequestMode _gotk4_gtk4_LayoutManagerClass_get_request_mode(GtkLayoutManager*, GtkWidget*);
// extern GtkLayoutChild* _gotk4_gtk4_LayoutManagerClass_create_layout_child(GtkLayoutManager*, GtkWidget*, GtkWidget*);
// GtkLayoutChild* _gotk4_gtk4_LayoutManager_virtual_create_layout_child(void* fnptr, GtkLayoutManager* arg0, GtkWidget* arg1, GtkWidget* arg2) {
//   return ((GtkLayoutChild* (*)(GtkLayoutManager*, GtkWidget*, GtkWidget*))(fnptr))(arg0, arg1, arg2);
// };
// GtkSizeRequestMode _gotk4_gtk4_LayoutManager_virtual_get_request_mode(void* fnptr, GtkLayoutManager* arg0, GtkWidget* arg1) {
//   return ((GtkSizeRequestMode (*)(GtkLayoutManager*, GtkWidget*))(fnptr))(arg0, arg1);
// };
// void _gotk4_gtk4_LayoutManager_virtual_allocate(void* fnptr, GtkLayoutManager* arg0, GtkWidget* arg1, int arg2, int arg3, int arg4) {
//   ((void (*)(GtkLayoutManager*, GtkWidget*, int, int, int))(fnptr))(arg0, arg1, arg2, arg3, arg4);
// };
// void _gotk4_gtk4_LayoutManager_virtual_measure(void* fnptr, GtkLayoutManager* arg0, GtkWidget* arg1, GtkOrientation arg2, int arg3, int* arg4, int* arg5, int* arg6, int* arg7) {
//   ((void (*)(GtkLayoutManager*, GtkWidget*, GtkOrientation, int, int*, int*, int*, int*))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7);
// };
// void _gotk4_gtk4_LayoutManager_virtual_root(void* fnptr, GtkLayoutManager* arg0) {
//   ((void (*)(GtkLayoutManager*))(fnptr))(arg0);
// };
// void _gotk4_gtk4_LayoutManager_virtual_unroot(void* fnptr, GtkLayoutManager* arg0) {
//   ((void (*)(GtkLayoutManager*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeLayoutManager = coreglib.Type(C.gtk_layout_manager_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeLayoutManager, F: marshalLayoutManager},
	})
}

// LayoutManagerOverrides contains methods that are overridable.
type LayoutManagerOverrides struct {
	// Allocate assigns the given width, height, and baseline to a widget, and
	// computes the position and sizes of the children of the widget using the
	// layout management policy of manager.
	//
	// The function takes the following parameters:
	//
	//    - widget: GtkWidget using manager.
	//    - width: new width of the widget.
	//    - height: new height of the widget.
	//    - baseline position of the widget, or -1.
	//
	Allocate func(widget Widgetter, width, height, baseline int)
	// CreateLayoutChild: create a LayoutChild instance for the given for_child
	// widget.
	//
	// The function takes the following parameters:
	//
	//    - widget using the manager.
	//    - forChild: child of widget.
	//
	// The function returns the following values:
	//
	//    - layoutChild: LayoutChild.
	//
	CreateLayoutChild func(widget, forChild Widgetter) LayoutChilder
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	RequestMode func(widget Widgetter) SizeRequestMode
	// Measure measures the size of the widget using manager, for the given
	// orientation and size.
	//
	// See the gtk.Widget documentation on layout management for more details.
	//
	// The function takes the following parameters:
	//
	//    - widget: GtkWidget using manager.
	//    - orientation to measure.
	//    - forSize: size for the opposite of orientation; for instance, if the
	//      orientation is GTK_ORIENTATION_HORIZONTAL, this is the height of the
	//      widget; if the orientation is GTK_ORIENTATION_VERTICAL, this is the
	//      width of the widget. This allows to measure the height for the given
	//      width, and the width for the given height. Use -1 if the size is not
	//      known.
	//
	// The function returns the following values:
	//
	//    - minimum (optional) size for the given size and orientation.
	//    - natural (optional): natural, or preferred size for the given size and
	//      orientation.
	//    - minimumBaseline (optional): baseline position for the minimum size.
	//    - naturalBaseline (optional): baseline position for the natural size.
	//
	Measure func(widget Widgetter, orientation Orientation, forSize int) (minimum, natural, minimumBaseline, naturalBaseline int)
	Root    func()
	Unroot  func()
}

func defaultLayoutManagerOverrides(v *LayoutManager) LayoutManagerOverrides {
	return LayoutManagerOverrides{
		Allocate:          v.allocate,
		CreateLayoutChild: v.createLayoutChild,
		RequestMode:       v.requestMode,
		Measure:           v.measure,
		Root:              v.root,
		Unroot:            v.unroot,
	}
}

// LayoutManager: layout managers are delegate classes that handle the preferred
// size and the allocation of a widget.
//
// You typically subclass GtkLayoutManager if you want to implement a layout
// policy for the children of a widget, or if you want to determine the size of
// a widget depending on its contents.
//
// Each GtkWidget can only have a GtkLayoutManager instance associated to it at
// any given time; it is possible, though, to replace the layout manager
// instance using gtk.Widget.SetLayoutManager().
//
//
// Layout properties
//
// A layout manager can expose properties for controlling the layout of each
// child, by creating an object type derived from gtk.LayoutChild and installing
// the properties on it as normal GObject properties.
//
// Each GtkLayoutChild instance storing the layout properties for a specific
// child is created through the gtk.LayoutManager.GetLayoutChild() method; a
// GtkLayoutManager controls the creation of its GtkLayoutChild instances by
// overriding the GtkLayoutManagerClass.create_layout_child() virtual function.
// The typical implementation should look like:
//
//    static GtkLayoutChild *
//    create_layout_child (GtkLayoutManager *manager,
//                         GtkWidget        *container,
//                         GtkWidget        *child)
//    {
//      return g_object_new (your_layout_child_get_type (),
//                           "layout-manager", manager,
//                           "child-widget", child,
//                           NULL);
//    }
//
//
// The gtk.LayoutChild:layout-manager and gtk.LayoutChild:child-widget
// properties on the newly created GtkLayoutChild instance are mandatory. The
// GtkLayoutManager will cache the newly created GtkLayoutChild instance until
// the widget is removed from its parent, or the parent removes the layout
// manager.
//
// Each GtkLayoutManager instance creating a GtkLayoutChild should use
// gtk.LayoutManager.GetLayoutChild() every time it needs to query the layout
// properties; each GtkLayoutChild instance should call
// gtk.LayoutManager.LayoutChanged() every time a property is updated, in order
// to queue a new size measuring and allocation.
type LayoutManager struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*LayoutManager)(nil)
)

// LayoutManagerer describes types inherited from class LayoutManager.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type LayoutManagerer interface {
	coreglib.Objector
	baseLayoutManager() *LayoutManager
}

var _ LayoutManagerer = (*LayoutManager)(nil)

func init() {
	coreglib.RegisterClassInfo[*LayoutManager, *LayoutManagerClass, LayoutManagerOverrides](
		GTypeLayoutManager,
		initLayoutManagerClass,
		wrapLayoutManager,
		defaultLayoutManagerOverrides,
	)
}

func initLayoutManagerClass(gclass unsafe.Pointer, overrides LayoutManagerOverrides, classInitFunc func(*LayoutManagerClass)) {
	pclass := (*C.GtkLayoutManagerClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeLayoutManager))))

	if overrides.Allocate != nil {
		pclass.allocate = (*[0]byte)(C._gotk4_gtk4_LayoutManagerClass_allocate)
	}

	if overrides.CreateLayoutChild != nil {
		pclass.create_layout_child = (*[0]byte)(C._gotk4_gtk4_LayoutManagerClass_create_layout_child)
	}

	if overrides.RequestMode != nil {
		pclass.get_request_mode = (*[0]byte)(C._gotk4_gtk4_LayoutManagerClass_get_request_mode)
	}

	if overrides.Measure != nil {
		pclass.measure = (*[0]byte)(C._gotk4_gtk4_LayoutManagerClass_measure)
	}

	if overrides.Root != nil {
		pclass.root = (*[0]byte)(C._gotk4_gtk4_LayoutManagerClass_root)
	}

	if overrides.Unroot != nil {
		pclass.unroot = (*[0]byte)(C._gotk4_gtk4_LayoutManagerClass_unroot)
	}

	if classInitFunc != nil {
		class := (*LayoutManagerClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapLayoutManager(obj *coreglib.Object) *LayoutManager {
	return &LayoutManager{
		Object: obj,
	}
}

func marshalLayoutManager(p uintptr) (interface{}, error) {
	return wrapLayoutManager(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (manager *LayoutManager) baseLayoutManager() *LayoutManager {
	return manager
}

// BaseLayoutManager returns the underlying base object.
func BaseLayoutManager(obj LayoutManagerer) *LayoutManager {
	return obj.baseLayoutManager()
}

// Allocate assigns the given width, height, and baseline to a widget, and
// computes the position and sizes of the children of the widget using the
// layout management policy of manager.
//
// The function takes the following parameters:
//
//    - widget: GtkWidget using manager.
//    - width: new width of the widget.
//    - height: new height of the widget.
//    - baseline position of the widget, or -1.
//
func (manager *LayoutManager) Allocate(widget Widgetter, width, height, baseline int) {
	var _arg0 *C.GtkLayoutManager // out
	var _arg1 *C.GtkWidget        // out
	var _arg2 C.int               // out
	var _arg3 C.int               // out
	var _arg4 C.int               // out

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg2 = C.int(width)
	_arg3 = C.int(height)
	_arg4 = C.int(baseline)

	C.gtk_layout_manager_allocate(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(baseline)
}

// LayoutChild retrieves a GtkLayoutChild instance for the GtkLayoutManager,
// creating one if necessary.
//
// The child widget must be a child of the widget using manager.
//
// The GtkLayoutChild instance is owned by the GtkLayoutManager, and is
// guaranteed to exist as long as child is a child of the GtkWidget using the
// given GtkLayoutManager.
//
// The function takes the following parameters:
//
//    - child: GtkWidget.
//
// The function returns the following values:
//
//    - layoutChild: GtkLayoutChild.
//
func (manager *LayoutManager) LayoutChild(child Widgetter) LayoutChilder {
	var _arg0 *C.GtkLayoutManager // out
	var _arg1 *C.GtkWidget        // out
	var _cret *C.GtkLayoutChild   // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_cret = C.gtk_layout_manager_get_layout_child(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(child)

	var _layoutChild LayoutChilder // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.LayoutChilder is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(LayoutChilder)
			return ok
		})
		rv, ok := casted.(LayoutChilder)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.LayoutChilder")
		}
		_layoutChild = rv
	}

	return _layoutChild
}

// RequestMode retrieves the request mode of manager.
//
// The function returns the following values:
//
//    - sizeRequestMode: GtkSizeRequestMode.
//
func (manager *LayoutManager) RequestMode() SizeRequestMode {
	var _arg0 *C.GtkLayoutManager  // out
	var _cret C.GtkSizeRequestMode // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	_cret = C.gtk_layout_manager_get_request_mode(_arg0)
	runtime.KeepAlive(manager)

	var _sizeRequestMode SizeRequestMode // out

	_sizeRequestMode = SizeRequestMode(_cret)

	return _sizeRequestMode
}

// Widget retrieves the GtkWidget using the given GtkLayoutManager.
//
// The function returns the following values:
//
//    - widget (optional): GtkWidget.
//
func (manager *LayoutManager) Widget() Widgetter {
	var _arg0 *C.GtkLayoutManager // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	_cret = C.gtk_layout_manager_get_widget(_arg0)
	runtime.KeepAlive(manager)

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

// LayoutChanged queues a resize on the GtkWidget using manager, if any.
//
// This function should be called by subclasses of GtkLayoutManager in response
// to changes to their layout management policies.
func (manager *LayoutManager) LayoutChanged() {
	var _arg0 *C.GtkLayoutManager // out

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	C.gtk_layout_manager_layout_changed(_arg0)
	runtime.KeepAlive(manager)
}

// Measure measures the size of the widget using manager, for the given
// orientation and size.
//
// See the gtk.Widget documentation on layout management for more details.
//
// The function takes the following parameters:
//
//    - widget: GtkWidget using manager.
//    - orientation to measure.
//    - forSize: size for the opposite of orientation; for instance, if the
//      orientation is GTK_ORIENTATION_HORIZONTAL, this is the height of the
//      widget; if the orientation is GTK_ORIENTATION_VERTICAL, this is the width
//      of the widget. This allows to measure the height for the given width, and
//      the width for the given height. Use -1 if the size is not known.
//
// The function returns the following values:
//
//    - minimum (optional) size for the given size and orientation.
//    - natural (optional): natural, or preferred size for the given size and
//      orientation.
//    - minimumBaseline (optional): baseline position for the minimum size.
//    - naturalBaseline (optional): baseline position for the natural size.
//
func (manager *LayoutManager) Measure(widget Widgetter, orientation Orientation, forSize int) (minimum, natural, minimumBaseline, naturalBaseline int) {
	var _arg0 *C.GtkLayoutManager // out
	var _arg1 *C.GtkWidget        // out
	var _arg2 C.GtkOrientation    // out
	var _arg3 C.int               // out
	var _arg4 C.int               // in
	var _arg5 C.int               // in
	var _arg6 C.int               // in
	var _arg7 C.int               // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg2 = C.GtkOrientation(orientation)
	_arg3 = C.int(forSize)

	C.gtk_layout_manager_measure(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5, &_arg6, &_arg7)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(orientation)
	runtime.KeepAlive(forSize)

	var _minimum int         // out
	var _natural int         // out
	var _minimumBaseline int // out
	var _naturalBaseline int // out

	_minimum = int(_arg4)
	_natural = int(_arg5)
	_minimumBaseline = int(_arg6)
	_naturalBaseline = int(_arg7)

	return _minimum, _natural, _minimumBaseline, _naturalBaseline
}

// Allocate assigns the given width, height, and baseline to a widget, and
// computes the position and sizes of the children of the widget using the
// layout management policy of manager.
//
// The function takes the following parameters:
//
//    - widget: GtkWidget using manager.
//    - width: new width of the widget.
//    - height: new height of the widget.
//    - baseline position of the widget, or -1.
//
func (manager *LayoutManager) allocate(widget Widgetter, width, height, baseline int) {
	gclass := (*C.GtkLayoutManagerClass)(coreglib.PeekParentClass(manager))
	fnarg := gclass.allocate

	var _arg0 *C.GtkLayoutManager // out
	var _arg1 *C.GtkWidget        // out
	var _arg2 C.int               // out
	var _arg3 C.int               // out
	var _arg4 C.int               // out

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg2 = C.int(width)
	_arg3 = C.int(height)
	_arg4 = C.int(baseline)

	C._gotk4_gtk4_LayoutManager_virtual_allocate(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(baseline)
}

// createLayoutChild: create a LayoutChild instance for the given for_child
// widget.
//
// The function takes the following parameters:
//
//    - widget using the manager.
//    - forChild: child of widget.
//
// The function returns the following values:
//
//    - layoutChild: LayoutChild.
//
func (manager *LayoutManager) createLayoutChild(widget, forChild Widgetter) LayoutChilder {
	gclass := (*C.GtkLayoutManagerClass)(coreglib.PeekParentClass(manager))
	fnarg := gclass.create_layout_child

	var _arg0 *C.GtkLayoutManager // out
	var _arg1 *C.GtkWidget        // out
	var _arg2 *C.GtkWidget        // out
	var _cret *C.GtkLayoutChild   // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(forChild).Native()))

	_cret = C._gotk4_gtk4_LayoutManager_virtual_create_layout_child(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(forChild)

	var _layoutChild LayoutChilder // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.LayoutChilder is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(LayoutChilder)
			return ok
		})
		rv, ok := casted.(LayoutChilder)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.LayoutChilder")
		}
		_layoutChild = rv
	}

	return _layoutChild
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (manager *LayoutManager) requestMode(widget Widgetter) SizeRequestMode {
	gclass := (*C.GtkLayoutManagerClass)(coreglib.PeekParentClass(manager))
	fnarg := gclass.get_request_mode

	var _arg0 *C.GtkLayoutManager  // out
	var _arg1 *C.GtkWidget         // out
	var _cret C.GtkSizeRequestMode // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_cret = C._gotk4_gtk4_LayoutManager_virtual_get_request_mode(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(widget)

	var _sizeRequestMode SizeRequestMode // out

	_sizeRequestMode = SizeRequestMode(_cret)

	return _sizeRequestMode
}

// Measure measures the size of the widget using manager, for the given
// orientation and size.
//
// See the gtk.Widget documentation on layout management for more details.
//
// The function takes the following parameters:
//
//    - widget: GtkWidget using manager.
//    - orientation to measure.
//    - forSize: size for the opposite of orientation; for instance, if the
//      orientation is GTK_ORIENTATION_HORIZONTAL, this is the height of the
//      widget; if the orientation is GTK_ORIENTATION_VERTICAL, this is the width
//      of the widget. This allows to measure the height for the given width, and
//      the width for the given height. Use -1 if the size is not known.
//
// The function returns the following values:
//
//    - minimum (optional) size for the given size and orientation.
//    - natural (optional): natural, or preferred size for the given size and
//      orientation.
//    - minimumBaseline (optional): baseline position for the minimum size.
//    - naturalBaseline (optional): baseline position for the natural size.
//
func (manager *LayoutManager) measure(widget Widgetter, orientation Orientation, forSize int) (minimum, natural, minimumBaseline, naturalBaseline int) {
	gclass := (*C.GtkLayoutManagerClass)(coreglib.PeekParentClass(manager))
	fnarg := gclass.measure

	var _arg0 *C.GtkLayoutManager // out
	var _arg1 *C.GtkWidget        // out
	var _arg2 C.GtkOrientation    // out
	var _arg3 C.int               // out
	var _arg4 C.int               // in
	var _arg5 C.int               // in
	var _arg6 C.int               // in
	var _arg7 C.int               // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg2 = C.GtkOrientation(orientation)
	_arg3 = C.int(forSize)

	C._gotk4_gtk4_LayoutManager_virtual_measure(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5, &_arg6, &_arg7)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(orientation)
	runtime.KeepAlive(forSize)

	var _minimum int         // out
	var _natural int         // out
	var _minimumBaseline int // out
	var _naturalBaseline int // out

	_minimum = int(_arg4)
	_natural = int(_arg5)
	_minimumBaseline = int(_arg6)
	_naturalBaseline = int(_arg7)

	return _minimum, _natural, _minimumBaseline, _naturalBaseline
}

func (manager *LayoutManager) root() {
	gclass := (*C.GtkLayoutManagerClass)(coreglib.PeekParentClass(manager))
	fnarg := gclass.root

	var _arg0 *C.GtkLayoutManager // out

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	C._gotk4_gtk4_LayoutManager_virtual_root(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(manager)
}

func (manager *LayoutManager) unroot() {
	gclass := (*C.GtkLayoutManagerClass)(coreglib.PeekParentClass(manager))
	fnarg := gclass.unroot

	var _arg0 *C.GtkLayoutManager // out

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	C._gotk4_gtk4_LayoutManager_virtual_unroot(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(manager)
}

// LayoutManagerClass: GtkLayoutManagerClass structure contains only private
// data, and should only be accessed through the provided API, or when
// subclassing LayoutManager.
//
// An instance of this type is always passed by reference.
type LayoutManagerClass struct {
	*layoutManagerClass
}

// layoutManagerClass is the struct that's finalized.
type layoutManagerClass struct {
	native *C.GtkLayoutManagerClass
}

// LayoutChildType: type of LayoutChild used by this layout manager.
func (l *LayoutManagerClass) LayoutChildType() coreglib.Type {
	valptr := &l.native.layout_child_type
	var _v coreglib.Type // out
	_v = coreglib.Type(*valptr)
	return _v
}
