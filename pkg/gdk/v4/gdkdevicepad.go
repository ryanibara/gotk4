// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_device_pad_feature_get_type()), F: marshalDevicePadFeature},
		{T: externglib.Type(C.gdk_device_pad_get_type()), F: marshalDevicePadder},
	})
}

// DevicePadFeature: pad feature.
type DevicePadFeature int

const (
	// Button: button
	DevicePadFeatureButton DevicePadFeature = iota
	// Ring: ring-shaped interactive area
	DevicePadFeatureRing
	// Strip: straight interactive area
	DevicePadFeatureStrip
)

func marshalDevicePadFeature(p uintptr) (interface{}, error) {
	return DevicePadFeature(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// DevicePadder describes DevicePad's methods.
type DevicePadder interface {
	// GroupNModes returns the number of modes that @group may have.
	GroupNModes(groupIdx int) int
	// NGroups returns the number of groups this pad device has.
	NGroups() int
}

// DevicePad: `GdkDevicePad` is an interface implemented by devices of type
// GDK_SOURCE_TABLET_PAD
//
// It allows querying the features provided by the pad device.
//
// Tablet pads may contain one or more groups, each containing a subset of the
// buttons/rings/strips available. [method@Gdk.DevicePad.get_n_groups] can be
// used to obtain the number of groups, [method@Gdk.DevicePad.get_n_features]
// and [method@Gdk.DevicePad.get_feature_group] can be combined to find out the
// number of buttons/rings/strips the device has, and how are they grouped.
//
// Each of those groups have different modes, which may be used to map each
// individual pad feature to multiple actions. Only one mode is effective
// (current) for each given group, different groups may have different current
// modes. The number of available modes in a group can be found out through
// [method@Gdk.DevicePad.get_group_n_modes], and the current mode for a given
// group will be notified through events of type K_PAD_GROUP_MODE.
type DevicePad struct {
	Device
}

var (
	_ DevicePadder    = (*DevicePad)(nil)
	_ gextras.Nativer = (*DevicePad)(nil)
)

func wrapDevicePad(obj *externglib.Object) DevicePadder {
	return &DevicePad{
		Device: Device{
			Object: obj,
		},
	}
}

func marshalDevicePadder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDevicePad(obj), nil
}

// GroupNModes returns the number of modes that @group may have.
func (pad *DevicePad) GroupNModes(groupIdx int) int {
	var _arg0 *C.GdkDevicePad // out
	var _arg1 C.int           // out
	var _cret C.int           // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(pad.Native()))
	_arg1 = C.int(groupIdx)

	_cret = C.gdk_device_pad_get_group_n_modes(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NGroups returns the number of groups this pad device has.
//
// Pads have at least one group. A pad group is a subcollection of
// buttons/strip/rings that is affected collectively by a same current mode.
func (pad *DevicePad) NGroups() int {
	var _arg0 *C.GdkDevicePad // out
	var _cret C.int           // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(pad.Native()))

	_cret = C.gdk_device_pad_get_n_groups(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
