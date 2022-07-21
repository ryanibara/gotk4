// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GTypeRevealerTransitionType returns the GType for the type RevealerTransitionType.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeRevealerTransitionType() coreglib.Type {
	gtype := coreglib.Type(C.gtk_revealer_transition_type_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalRevealerTransitionType)
	return gtype
}

// GTypeRevealer returns the GType for the type Revealer.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeRevealer() coreglib.Type {
	gtype := coreglib.Type(C.gtk_revealer_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalRevealer)
	return gtype
}

// RevealerTransitionType: these enumeration values describe the possible
// transitions when the child of a Revealer widget is shown or hidden.
type RevealerTransitionType C.gint

const (
	// RevealerTransitionTypeNone: no transition.
	RevealerTransitionTypeNone RevealerTransitionType = iota
	// RevealerTransitionTypeCrossfade: fade in.
	RevealerTransitionTypeCrossfade
	// RevealerTransitionTypeSlideRight: slide in from the left.
	RevealerTransitionTypeSlideRight
	// RevealerTransitionTypeSlideLeft: slide in from the right.
	RevealerTransitionTypeSlideLeft
	// RevealerTransitionTypeSlideUp: slide in from the bottom.
	RevealerTransitionTypeSlideUp
	// RevealerTransitionTypeSlideDown: slide in from the top.
	RevealerTransitionTypeSlideDown
)

func marshalRevealerTransitionType(p uintptr) (interface{}, error) {
	return RevealerTransitionType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for RevealerTransitionType.
func (r RevealerTransitionType) String() string {
	switch r {
	case RevealerTransitionTypeNone:
		return "None"
	case RevealerTransitionTypeCrossfade:
		return "Crossfade"
	case RevealerTransitionTypeSlideRight:
		return "SlideRight"
	case RevealerTransitionTypeSlideLeft:
		return "SlideLeft"
	case RevealerTransitionTypeSlideUp:
		return "SlideUp"
	case RevealerTransitionTypeSlideDown:
		return "SlideDown"
	default:
		return fmt.Sprintf("RevealerTransitionType(%d)", r)
	}
}

// RevealerOverrider contains methods that are overridable.
type RevealerOverrider interface {
}

// Revealer widget is a container which animates the transition of its child
// from invisible to visible.
//
// The style of transition can be controlled with
// gtk_revealer_set_transition_type().
//
// These animations respect the Settings:gtk-enable-animations setting.
//
//
// CSS nodes
//
// GtkRevealer has a single CSS node with name revealer.
//
// The GtkRevealer widget was added in GTK+ 3.10.
type Revealer struct {
	_ [0]func() // equal guard
	Bin
}

var (
	_ Binner = (*Revealer)(nil)
)

func classInitRevealerer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRevealer(obj *coreglib.Object) *Revealer {
	return &Revealer{
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

func marshalRevealer(p uintptr) (interface{}, error) {
	return wrapRevealer(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewRevealer creates a new Revealer.
//
// The function returns the following values:
//
//    - revealer: newly created Revealer.
//
func NewRevealer() *Revealer {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_revealer_new()

	var _revealer *Revealer // out

	_revealer = wrapRevealer(coreglib.Take(unsafe.Pointer(_cret)))

	return _revealer
}

// ChildRevealed returns whether the child is fully revealed, in other words
// whether the transition to the revealed state is completed.
//
// The function returns the following values:
//
//    - ok: TRUE if the child is fully revealed.
//
func (revealer *Revealer) ChildRevealed() bool {
	var _arg0 *C.GtkRevealer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(coreglib.InternObject(revealer).Native()))

	_cret = C.gtk_revealer_get_child_revealed(_arg0)
	runtime.KeepAlive(revealer)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RevealChild returns whether the child is currently revealed. See
// gtk_revealer_set_reveal_child().
//
// This function returns TRUE as soon as the transition is to the revealed state
// is started. To learn whether the child is fully revealed (ie the transition
// is completed), use gtk_revealer_get_child_revealed().
//
// The function returns the following values:
//
//    - ok: TRUE if the child is revealed.
//
func (revealer *Revealer) RevealChild() bool {
	var _arg0 *C.GtkRevealer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(coreglib.InternObject(revealer).Native()))

	_cret = C.gtk_revealer_get_reveal_child(_arg0)
	runtime.KeepAlive(revealer)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TransitionDuration returns the amount of time (in milliseconds) that
// transitions will take.
//
// The function returns the following values:
//
//    - guint: transition duration.
//
func (revealer *Revealer) TransitionDuration() uint {
	var _arg0 *C.GtkRevealer // out
	var _cret C.guint        // in

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(coreglib.InternObject(revealer).Native()))

	_cret = C.gtk_revealer_get_transition_duration(_arg0)
	runtime.KeepAlive(revealer)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// TransitionType gets the type of animation that will be used for transitions
// in revealer.
//
// The function returns the following values:
//
//    - revealerTransitionType: current transition type of revealer.
//
func (revealer *Revealer) TransitionType() RevealerTransitionType {
	var _arg0 *C.GtkRevealer              // out
	var _cret C.GtkRevealerTransitionType // in

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(coreglib.InternObject(revealer).Native()))

	_cret = C.gtk_revealer_get_transition_type(_arg0)
	runtime.KeepAlive(revealer)

	var _revealerTransitionType RevealerTransitionType // out

	_revealerTransitionType = RevealerTransitionType(_cret)

	return _revealerTransitionType
}

// SetRevealChild tells the Revealer to reveal or conceal its child.
//
// The transition will be animated with the current transition type of revealer.
//
// The function takes the following parameters:
//
//    - revealChild: TRUE to reveal the child.
//
func (revealer *Revealer) SetRevealChild(revealChild bool) {
	var _arg0 *C.GtkRevealer // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(coreglib.InternObject(revealer).Native()))
	if revealChild {
		_arg1 = C.TRUE
	}

	C.gtk_revealer_set_reveal_child(_arg0, _arg1)
	runtime.KeepAlive(revealer)
	runtime.KeepAlive(revealChild)
}

// SetTransitionDuration sets the duration that transitions will take.
//
// The function takes the following parameters:
//
//    - duration: new duration, in milliseconds.
//
func (revealer *Revealer) SetTransitionDuration(duration uint) {
	var _arg0 *C.GtkRevealer // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(coreglib.InternObject(revealer).Native()))
	_arg1 = C.guint(duration)

	C.gtk_revealer_set_transition_duration(_arg0, _arg1)
	runtime.KeepAlive(revealer)
	runtime.KeepAlive(duration)
}

// SetTransitionType sets the type of animation that will be used for
// transitions in revealer. Available types include various kinds of fades and
// slides.
//
// The function takes the following parameters:
//
//    - transition: new transition type.
//
func (revealer *Revealer) SetTransitionType(transition RevealerTransitionType) {
	var _arg0 *C.GtkRevealer              // out
	var _arg1 C.GtkRevealerTransitionType // out

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(coreglib.InternObject(revealer).Native()))
	_arg1 = C.GtkRevealerTransitionType(transition)

	C.gtk_revealer_set_transition_type(_arg0, _arg1)
	runtime.KeepAlive(revealer)
	runtime.KeepAlive(transition)
}
