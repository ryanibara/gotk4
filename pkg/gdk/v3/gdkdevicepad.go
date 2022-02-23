// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gdkdevicepad.go.
var (
	GTypeDevicePadFeature = externglib.Type(C.gdk_device_pad_feature_get_type())
	GTypeDevicePad        = externglib.Type(C.gdk_device_pad_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
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
	return DevicePadFeature(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
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
type DevicePad struct {
	_ [0]func() // equal guard
	Device
}

var (
	_ Devicer = (*DevicePad)(nil)
)

// DevicePadder describes DevicePad's interface methods.
type DevicePadder interface {
	externglib.Objector

	// FeatureGroup returns the group the given feature and idx belong to, or -1
	// if feature/index do not exist in pad.
	FeatureGroup(feature DevicePadFeature, featureIdx int) int
	// GroupNModes returns the number of modes that group may have.
	GroupNModes(groupIdx int) int
	// NFeatures returns the number of features a tablet pad has.
	NFeatures(feature DevicePadFeature) int
	// NGroups returns the number of groups this pad device has.
	NGroups() int
}

var _ DevicePadder = (*DevicePad)(nil)

func ifaceInitDevicePadder(gifacePtr, data C.gpointer) {
}

func wrapDevicePad(obj *externglib.Object) *DevicePad {
	return &DevicePad{
		Device: Device{
			Object: obj,
		},
	}
}

func marshalDevicePad(p uintptr) (interface{}, error) {
	return wrapDevicePad(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// FeatureGroup returns the group the given feature and idx belong to, or -1 if
// feature/index do not exist in pad.
//
// The function takes the following parameters:
//
//    - feature type to get the group from.
//    - featureIdx: index of the feature to get the group from.
//
// The function returns the following values:
//
//    - gint: group number of the queried pad feature.
//
func (pad *DevicePad) FeatureGroup(feature DevicePadFeature, featureIdx int) int {
	var _arg0 *C.GdkDevicePad       // out
	var _arg1 C.GdkDevicePadFeature // out
	var _arg2 C.gint                // out
	var _cret C.gint                // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(pad.Native()))
	_arg1 = C.GdkDevicePadFeature(feature)
	_arg2 = C.gint(featureIdx)

	_cret = C.gdk_device_pad_get_feature_group(_arg0, _arg1, _arg2)
	runtime.KeepAlive(pad)
	runtime.KeepAlive(feature)
	runtime.KeepAlive(featureIdx)

	var _gint int // out

	_gint = int(_cret)

	return _gint
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
func (pad *DevicePad) GroupNModes(groupIdx int) int {
	var _arg0 *C.GdkDevicePad // out
	var _arg1 C.gint          // out
	var _cret C.gint          // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(pad.Native()))
	_arg1 = C.gint(groupIdx)

	_cret = C.gdk_device_pad_get_group_n_modes(_arg0, _arg1)
	runtime.KeepAlive(pad)
	runtime.KeepAlive(groupIdx)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NFeatures returns the number of features a tablet pad has.
//
// The function takes the following parameters:
//
//    - feature: pad feature.
//
// The function returns the following values:
//
//    - gint: amount of elements of type feature that this pad has.
//
func (pad *DevicePad) NFeatures(feature DevicePadFeature) int {
	var _arg0 *C.GdkDevicePad       // out
	var _arg1 C.GdkDevicePadFeature // out
	var _cret C.gint                // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(pad.Native()))
	_arg1 = C.GdkDevicePadFeature(feature)

	_cret = C.gdk_device_pad_get_n_features(_arg0, _arg1)
	runtime.KeepAlive(pad)
	runtime.KeepAlive(feature)

	var _gint int // out

	_gint = int(_cret)

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
func (pad *DevicePad) NGroups() int {
	var _arg0 *C.GdkDevicePad // out
	var _cret C.gint          // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(pad.Native()))

	_cret = C.gdk_device_pad_get_n_groups(_arg0)
	runtime.KeepAlive(pad)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
