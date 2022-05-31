// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gdkdevicepad.go.
var (
	GTypeDevicePadFeature = coreglib.Type(C.gdk_device_pad_feature_get_type())
	GTypeDevicePad        = coreglib.Type(C.gdk_device_pad_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeDevicePadFeature, F: marshalDevicePadFeature},
		{T: GTypeDevicePad, F: marshalDevicePad},
	})
}

// DevicePadFeature: pad feature.
type DevicePadFeature C.gint

const (
	// DevicePadFeatureButton: button.
	DevicePadFeatureButton DevicePadFeature = iota
	// DevicePadFeatureRing: ring-shaped interactive area.
	DevicePadFeatureRing
	// DevicePadFeatureStrip: straight interactive area.
	DevicePadFeatureStrip
)

func marshalDevicePadFeature(p uintptr) (interface{}, error) {
	return DevicePadFeature(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for DevicePadFeature.
func (d DevicePadFeature) String() string {
	switch d {
	case DevicePadFeatureButton:
		return "Button"
	case DevicePadFeatureRing:
		return "Ring"
	case DevicePadFeatureStrip:
		return "Strip"
	default:
		return fmt.Sprintf("DevicePadFeature(%d)", d)
	}
}

// DevicePadOverrider contains methods that are overridable.
type DevicePadOverrider interface {
}

// DevicePad is an interface implemented by devices of type
// GDK_SOURCE_TABLET_PAD, it allows querying the features provided by the pad
// device.
//
// Tablet pads may contain one or more groups, each containing a subset of the
// buttons/rings/strips available. gdk_device_pad_get_n_groups() can be used to
// obtain the number of groups, gdk_device_pad_get_n_features() and
// gdk_device_pad_get_feature_group() can be combined to find out the number of
// buttons/rings/strips the device has, and how are they grouped.
//
// Each of those groups have different modes, which may be used to map each
// individual pad feature to multiple actions. Only one mode is effective
// (current) for each given group, different groups may have different current
// modes. The number of available modes in a group can be found out through
// gdk_device_pad_get_group_n_modes(), and the current mode for a given group
// will be notified through the EventPadGroupMode event.
//
// DevicePad wraps an interface. This means the user can get the
// underlying type by calling Cast().
type DevicePad struct {
	_ [0]func() // equal guard
	Device
}

var (
	_ Devicer = (*DevicePad)(nil)
)

// DevicePadder describes DevicePad's interface methods.
type DevicePadder interface {
	coreglib.Objector

	// GroupNModes returns the number of modes that group may have.
	GroupNModes(groupIdx int32) int32
	// NGroups returns the number of groups this pad device has.
	NGroups() int32
}

var _ DevicePadder = (*DevicePad)(nil)

func ifaceInitDevicePadder(gifacePtr, data C.gpointer) {
}

func wrapDevicePad(obj *coreglib.Object) *DevicePad {
	return &DevicePad{
		Device: Device{
			Object: obj,
		},
	}
}

func marshalDevicePad(p uintptr) (interface{}, error) {
	return wrapDevicePad(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// GroupNModes returns the number of modes that group may have.
//
// The function takes the following parameters:
//
//    - groupIdx: group to get the number of available modes from.
//
// The function returns the following values:
//
//    - gint: number of modes available in group.
//
func (pad *DevicePad) GroupNModes(groupIdx int32) int32 {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(pad).Native()))
	_arg1 = C.gint(groupIdx)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gint)(unsafe.Pointer(&_args[1])) = _arg1

	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(pad)
	runtime.KeepAlive(groupIdx)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// NGroups returns the number of groups this pad device has. Pads have at least
// one group. A pad group is a subcollection of buttons/strip/rings that is
// affected collectively by a same current mode.
//
// The function returns the following values:
//
//    - gint: number of button/ring/strip groups in the pad.
//
func (pad *DevicePad) NGroups() int32 {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(pad).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(pad)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}
