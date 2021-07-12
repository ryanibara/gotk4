// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_device_pad_feature_get_type()), F: marshalDevicePadFeature},
		{T: externglib.Type(C.gdk_device_pad_get_type()), F: marshalDevicePader},
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

// DevicePader describes DevicePad's methods.
type DevicePader interface {
	// FeatureGroup returns the group the given @feature and @idx belong to, or
	// -1 if feature/index do not exist in @pad.
	FeatureGroup(feature DevicePadFeature, featureIdx int) int
	// GroupNModes returns the number of modes that @group may have.
	GroupNModes(groupIdx int) int
	// NFeatures returns the number of features a tablet pad has.
	NFeatures(feature DevicePadFeature) int
	// NGroups returns the number of groups this pad device has.
	NGroups() int
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
	Device
}

var (
	_ DevicePader     = (*DevicePad)(nil)
	_ gextras.Nativer = (*DevicePad)(nil)
)

func wrapDevicePad(obj *externglib.Object) *DevicePad {
	return &DevicePad{
		Device: Device{
			Object: obj,
		},
	}
}

func marshalDevicePader(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDevicePad(obj), nil
}

// FeatureGroup returns the group the given @feature and @idx belong to, or -1
// if feature/index do not exist in @pad.
func (pad *DevicePad) FeatureGroup(feature DevicePadFeature, featureIdx int) int {
	var _arg0 *C.GdkDevicePad       // out
	var _arg1 C.GdkDevicePadFeature // out
	var _arg2 C.gint                // out
	var _cret C.gint                // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(pad.Native()))
	_arg1 = C.GdkDevicePadFeature(feature)
	_arg2 = C.gint(featureIdx)

	_cret = C.gdk_device_pad_get_feature_group(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// GroupNModes returns the number of modes that @group may have.
func (pad *DevicePad) GroupNModes(groupIdx int) int {
	var _arg0 *C.GdkDevicePad // out
	var _arg1 C.gint          // out
	var _cret C.gint          // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(pad.Native()))
	_arg1 = C.gint(groupIdx)

	_cret = C.gdk_device_pad_get_group_n_modes(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NFeatures returns the number of features a tablet pad has.
func (pad *DevicePad) NFeatures(feature DevicePadFeature) int {
	var _arg0 *C.GdkDevicePad       // out
	var _arg1 C.GdkDevicePadFeature // out
	var _cret C.gint                // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(pad.Native()))
	_arg1 = C.GdkDevicePadFeature(feature)

	_cret = C.gdk_device_pad_get_n_features(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NGroups returns the number of groups this pad device has. Pads have at least
// one group. A pad group is a subcollection of buttons/strip/rings that is
// affected collectively by a same current mode.
func (pad *DevicePad) NGroups() int {
	var _arg0 *C.GdkDevicePad // out
	var _cret C.gint          // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(pad.Native()))

	_cret = C.gdk_device_pad_get_n_groups(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
