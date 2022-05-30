// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for atkutil.go.
var (
	GTypeCoordType    = coreglib.Type(C.atk_coord_type_get_type())
	GTypeKeyEventType = coreglib.Type(C.atk_key_event_type_get_type())
	GTypeUtil         = coreglib.Type(C.atk_util_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeCoordType, F: marshalCoordType},
		{T: GTypeKeyEventType, F: marshalKeyEventType},
		{T: GTypeUtil, F: marshalUtil},
	})
}

// CoordType specifies how xy coordinates are to be interpreted. Used by
// functions such as atk_component_get_position() and
// atk_text_get_character_extents().
type CoordType C.gint

const (
	// XYScreen specifies xy coordinates relative to the screen.
	XYScreen CoordType = iota
	// XYWindow specifies xy coordinates relative to the widget's top-level
	// window.
	XYWindow
	// XYParent specifies xy coordinates relative to the widget's immediate
	// parent. Since: 2.30.
	XYParent
)

func marshalCoordType(p uintptr) (interface{}, error) {
	return CoordType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CoordType.
func (c CoordType) String() string {
	switch c {
	case XYScreen:
		return "Screen"
	case XYWindow:
		return "Window"
	case XYParent:
		return "Parent"
	default:
		return fmt.Sprintf("CoordType(%d)", c)
	}
}

// KeyEventType specifies the type of a keyboard evemt.
type KeyEventType C.gint

const (
	// KeyEventPress specifies a key press event.
	KeyEventPress KeyEventType = iota
	// KeyEventRelease specifies a key release event.
	KeyEventRelease
	// KeyEventLastDefined: not a valid value; specifies end of enumeration.
	KeyEventLastDefined
)

func marshalKeyEventType(p uintptr) (interface{}, error) {
	return KeyEventType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for KeyEventType.
func (k KeyEventType) String() string {
	switch k {
	case KeyEventPress:
		return "Press"
	case KeyEventRelease:
		return "Release"
	case KeyEventLastDefined:
		return "LastDefined"
	default:
		return fmt.Sprintf("KeyEventType(%d)", k)
	}
}

// KeySnoopFunc is a type of callback which is called whenever a key event
// occurs, if registered via atk_add_key_event_listener. It allows for
// pre-emptive interception of key events via the return code as described
// below.
type KeySnoopFunc func(event *KeyEventStruct) (gint int)

//export _gotk4_atk1_KeySnoopFunc
func _gotk4_atk1_KeySnoopFunc(arg1 *C.AtkKeyEventStruct, arg2 C.gpointer) (cret C.gint) {
	var fn KeySnoopFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(KeySnoopFunc)
	}

	var _event *KeyEventStruct // out

	_event = (*KeyEventStruct)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	gint := fn(_event)

	cret = C.gint(gint)

	return cret
}

// FocusTrackerNotify: cause the focus tracker functions which have been
// specified to be executed for the object.
//
// Deprecated: Focus tracking has been dropped as a feature to be implemented by
// ATK itself. As Object::focus-event was deprecated in favor of a
// Object::state-change signal, in order to notify a focus change on your
// implementation, you can use atk_object_notify_state_change() instead.
//
// The function takes the following parameters:
//
//    - object: Object.
//
func FocusTrackerNotify(object *ObjectClass) {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(object).Native()))
	*(**ObjectClass)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Atk", "focus_tracker_notify").Invoke(args[:], nil)

	runtime.KeepAlive(object)
}

// GetFocusObject gets the currently focused object.
//
// The function returns the following values:
//
//    - object: currently focused object for the current application.
//
func GetFocusObject() *ObjectClass {
	var _cret *C.void // in

	_gret := girepository.MustFind("Atk", "get_focus_object").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _object *ObjectClass // out

	_object = wrapObject(coreglib.Take(unsafe.Pointer(_cret)))

	return _object
}

// GetRoot gets the root accessible container for the current application.
//
// The function returns the following values:
//
//    - object: root accessible container for the current application.
//
func GetRoot() *ObjectClass {
	var _cret *C.void // in

	_gret := girepository.MustFind("Atk", "get_root").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _object *ObjectClass // out

	_object = wrapObject(coreglib.Take(unsafe.Pointer(_cret)))

	return _object
}

// GetToolkitName gets name string for the GUI toolkit implementing ATK for this
// application.
//
// The function returns the following values:
//
//    - utf8: name string for the GUI toolkit implementing ATK for this
//      application.
//
func GetToolkitName() string {
	var _cret *C.void // in

	_gret := girepository.MustFind("Atk", "get_toolkit_name").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// GetToolkitVersion gets version string for the GUI toolkit implementing ATK
// for this application.
//
// The function returns the following values:
//
//    - utf8: version string for the GUI toolkit implementing ATK for this
//      application.
//
func GetToolkitVersion() string {
	var _cret *C.void // in

	_gret := girepository.MustFind("Atk", "get_toolkit_version").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// GetVersion gets the current version for ATK.
//
// The function returns the following values:
//
//    - utf8: version string for ATK.
//
func GetVersion() string {
	var _cret *C.void // in

	_gret := girepository.MustFind("Atk", "get_version").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// RemoveFocusTracker removes the specified focus tracker from the list of
// functions to be called when any object receives focus.
//
// Deprecated: Focus tracking has been dropped as a feature to be implemented by
// ATK itself. If you need focus tracking on your implementation, subscribe to
// the Object::state-change "focused" signal.
//
// The function takes the following parameters:
//
//    - trackerId: id of the focus tracker to remove.
//
func RemoveFocusTracker(trackerId uint) {
	var args [1]girepository.Argument
	var _arg0 C.guint // out

	_arg0 = C.guint(trackerId)
	*(*uint)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Atk", "remove_focus_tracker").Invoke(args[:], nil)

	runtime.KeepAlive(trackerId)
}

// RemoveGlobalEventListener: listener_id is the value returned by
// #atk_add_global_event_listener when you registered that event listener.
//
// Toolkit implementor note: ATK provides a default implementation for this
// virtual method. ATK implementors are discouraged from reimplementing this
// method.
//
// Toolkit implementor note: this method is not intended to be used by ATK
// implementors but by ATK consumers.
//
// Removes the specified event listener.
//
// The function takes the following parameters:
//
//    - listenerId: id of the event listener to remove.
//
func RemoveGlobalEventListener(listenerId uint) {
	var args [1]girepository.Argument
	var _arg0 C.guint // out

	_arg0 = C.guint(listenerId)
	*(*uint)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Atk", "remove_global_event_listener").Invoke(args[:], nil)

	runtime.KeepAlive(listenerId)
}

// RemoveKeyEventListener: listener_id is the value returned by
// #atk_add_key_event_listener when you registered that event listener.
//
// Removes the specified event listener.
//
// The function takes the following parameters:
//
//    - listenerId: id of the event listener to remove.
//
func RemoveKeyEventListener(listenerId uint) {
	var args [1]girepository.Argument
	var _arg0 C.guint // out

	_arg0 = C.guint(listenerId)
	*(*uint)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Atk", "remove_key_event_listener").Invoke(args[:], nil)

	runtime.KeepAlive(listenerId)
}

// UtilOverrider contains methods that are overridable.
type UtilOverrider interface {
}

// Util: set of ATK utility functions which are used to support event
// registration of various types, and obtaining the 'root' accessible of a
// process and information about the current ATK implementation and toolkit
// version.
type Util struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Util)(nil)
)

func classInitUtiller(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapUtil(obj *coreglib.Object) *Util {
	return &Util{
		Object: obj,
	}
}

func marshalUtil(p uintptr) (interface{}, error) {
	return wrapUtil(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// KeyEventStruct encapsulates information about a key event.
//
// An instance of this type is always passed by reference.
type KeyEventStruct struct {
	*keyEventStruct
}

// keyEventStruct is the struct that's finalized.
type keyEventStruct struct {
	native *C.AtkKeyEventStruct
}

// Type: atkKeyEventType, generally one of ATK_KEY_EVENT_PRESS or
// ATK_KEY_EVENT_RELEASE.
func (k *KeyEventStruct) Type() int {
	var v int // out
	v = int(k.native._type)
	return v
}

// State: bitmask representing the state of the modifier keys immediately after
// the event takes place. The meaning of the bits is currently defined to match
// the bitmask used by GDK in GdkEventType.state, see
// http://developer.gnome.org/doc/API/2.0/gdk/gdk-Event-Structures.htmlEventKey.
func (k *KeyEventStruct) State() uint {
	var v uint // out
	v = uint(k.native.state)
	return v
}

// Keyval: guint representing a keysym value corresponding to those used by GDK
// and X11: see /usr/X11/include/keysymdef.h.
func (k *KeyEventStruct) Keyval() uint {
	var v uint // out
	v = uint(k.native.keyval)
	return v
}

// Length: length of member #string.
func (k *KeyEventStruct) Length() int {
	var v int // out
	v = int(k.native.length)
	return v
}

// String: string containing one of the following: either a string approximating
// the text that would result from this keypress, if the key is a control or
// graphic character, or a symbolic name for this keypress. Alphanumeric and
// printable keys will have the symbolic key name in this string member, for
// instance "A". "0", "semicolon", "aacute". Keypad keys have the prefix "KP".
func (k *KeyEventStruct) String() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(k.native.string)))
	return v
}

// Keycode: raw hardware code that generated the key event. This field is raraly
// useful.
func (k *KeyEventStruct) Keycode() uint16 {
	var v uint16 // out
	v = uint16(k.native.keycode)
	return v
}

// Timestamp: timestamp in milliseconds indicating when the event occurred.
// These timestamps are relative to a starting point which should be considered
// arbitrary, and only used to compare the dispatch times of events to one
// another.
func (k *KeyEventStruct) Timestamp() uint32 {
	var v uint32 // out
	v = uint32(k.native.timestamp)
	return v
}
