// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gtk4_ScrollableInterface_get_border(void*, void*);
import "C"

// glib.Type values for gtkscrollable.go.
var GTypeScrollable = coreglib.Type(C.gtk_scrollable_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeScrollable, F: marshalScrollable},
	})
}

// ScrollableOverrider contains methods that are overridable.
type ScrollableOverrider interface {
	// Border returns the size of a non-scrolling border around the outside of
	// the scrollable.
	//
	// An example for this would be treeview headers. GTK can use this
	// information to display overlaid graphics, like the overshoot indication,
	// at the right position.
	//
	// The function returns the following values:
	//
	//    - border: return location for the results.
	//    - ok: TRUE if border has been set.
	//
	Border() (*Border, bool)
}

// Scrollable: GtkScrollable is an interface for widgets with native scrolling
// ability.
//
// To implement this interface you should override the
// gtk.Scrollable:hadjustment and gtk.Scrollable:vadjustment properties.
//
//
// Creating a scrollable widget
//
// All scrollable widgets should do the following.
//
// - When a parent widget sets the scrollable child widget’s adjustments, the
// widget should populate the adjustments’ gtk.Adjustment:lower,
// gtk.Adjustment:upper, gtk.Adjustment:step-increment,
// gtk.Adjustment:page-increment and gtk.Adjustment:page-size properties and
// connect to the gtk.Adjustment::value-changed signal.
//
// - Because its preferred size is the size for a fully expanded widget, the
// scrollable widget must be able to cope with underallocations. This means that
// it must accept any value passed to its GtkWidgetClass.size_allocate()
// function.
//
// - When the parent allocates space to the scrollable child widget, the widget
// should update the adjustments’ properties with new values.
//
// - When any of the adjustments emits the gtk.Adjustment::value-changed signal,
// the scrollable widget should scroll its contents.
//
// Scrollable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Scrollable struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Scrollable)(nil)
)

// Scrollabler describes Scrollable's interface methods.
type Scrollabler interface {
	coreglib.Objector

	// Border returns the size of a non-scrolling border around the outside of
	// the scrollable.
	Border() (*Border, bool)
	// HAdjustment retrieves the GtkAdjustment used for horizontal scrolling.
	HAdjustment() *Adjustment
	// VAdjustment retrieves the GtkAdjustment used for vertical scrolling.
	VAdjustment() *Adjustment
	// SetHAdjustment sets the horizontal adjustment of the GtkScrollable.
	SetHAdjustment(hadjustment *Adjustment)
	// SetVAdjustment sets the vertical adjustment of the GtkScrollable.
	SetVAdjustment(vadjustment *Adjustment)
}

var _ Scrollabler = (*Scrollable)(nil)

func ifaceInitScrollabler(gifacePtr, data C.gpointer) {
	iface := (*C.GtkScrollableInterface)(unsafe.Pointer(gifacePtr))
	iface.get_border = (*[0]byte)(C._gotk4_gtk4_ScrollableInterface_get_border)
}

//export _gotk4_gtk4_ScrollableInterface_get_border
func _gotk4_gtk4_ScrollableInterface_get_border(arg0 *C.void, arg1 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ScrollableOverrider)

	border, ok := iface.Border()

	*arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(border)))
	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapScrollable(obj *coreglib.Object) *Scrollable {
	return &Scrollable{
		Object: obj,
	}
}

func marshalScrollable(p uintptr) (interface{}, error) {
	return wrapScrollable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Border returns the size of a non-scrolling border around the outside of the
// scrollable.
//
// An example for this would be treeview headers. GTK can use this information
// to display overlaid graphics, like the overshoot indication, at the right
// position.
//
// The function returns the following values:
//
//    - border: return location for the results.
//    - ok: TRUE if border has been set.
//
func (scrollable *Scrollable) Border() (*Border, bool) {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrollable).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrollable)

	var _border *Border // out
	var _ok bool        // out

	_border = (*Border)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _border, _ok
}

// HAdjustment retrieves the GtkAdjustment used for horizontal scrolling.
//
// The function returns the following values:
//
//    - adjustment: horizontal GtkAdjustment.
//
func (scrollable *Scrollable) HAdjustment() *Adjustment {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrollable).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrollable)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// VAdjustment retrieves the GtkAdjustment used for vertical scrolling.
//
// The function returns the following values:
//
//    - adjustment: vertical GtkAdjustment.
//
func (scrollable *Scrollable) VAdjustment() *Adjustment {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrollable).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrollable)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// SetHAdjustment sets the horizontal adjustment of the GtkScrollable.
//
// The function takes the following parameters:
//
//    - hadjustment (optional): GtkAdjustment.
//
func (scrollable *Scrollable) SetHAdjustment(hadjustment *Adjustment) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrollable).Native()))
	if hadjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(hadjustment).Native()))
	}

	runtime.KeepAlive(scrollable)
	runtime.KeepAlive(hadjustment)
}

// SetVAdjustment sets the vertical adjustment of the GtkScrollable.
//
// The function takes the following parameters:
//
//    - vadjustment (optional): GtkAdjustment.
//
func (scrollable *Scrollable) SetVAdjustment(vadjustment *Adjustment) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrollable).Native()))
	if vadjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(vadjustment).Native()))
	}

	runtime.KeepAlive(scrollable)
	runtime.KeepAlive(vadjustment)
}
