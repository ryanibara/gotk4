// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// NewAttrBackgroundAlpha: create a new background alpha attribute.
//
// The function takes the following parameters:
//
//   - alpha value, between 1 and 65536.
//
// The function returns the following values:
//
//   - attribute: newly allocated PangoAttribute, which should be freed with
//     pango.Attribute.Destroy().
//
func NewAttrBackgroundAlpha(alpha uint16) *Attribute {
	var _arg1 C.guint16         // out
	var _cret *C.PangoAttribute // in

	_arg1 = C.guint16(alpha)

	_cret = C.pango_attr_background_alpha_new(_arg1)
	runtime.KeepAlive(alpha)

	var _attribute *Attribute // out

	_attribute = (*Attribute)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_attribute)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_attribute_destroy((*C.PangoAttribute)(intern.C))
		},
	)

	return _attribute
}

// NewAttrForegroundAlpha: create a new foreground alpha attribute.
//
// The function takes the following parameters:
//
//   - alpha value, between 1 and 65536.
//
// The function returns the following values:
//
//   - attribute: newly allocated PangoAttribute, which should be freed with
//     pango.Attribute.Destroy().
//
func NewAttrForegroundAlpha(alpha uint16) *Attribute {
	var _arg1 C.guint16         // out
	var _cret *C.PangoAttribute // in

	_arg1 = C.guint16(alpha)

	_cret = C.pango_attr_foreground_alpha_new(_arg1)
	runtime.KeepAlive(alpha)

	var _attribute *Attribute // out

	_attribute = (*Attribute)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_attribute)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_attribute_destroy((*C.PangoAttribute)(intern.C))
		},
	)

	return _attribute
}

// AttrFontFeatures: PangoAttrFontFeatures structure is used to represent
// OpenType font features as an attribute.
//
// An instance of this type is always passed by reference.
type AttrFontFeatures struct {
	*attrFontFeatures
}

// attrFontFeatures is the struct that's finalized.
type attrFontFeatures struct {
	native *C.PangoAttrFontFeatures
}

// Attr: common portion of the attribute.
func (a *AttrFontFeatures) Attr() *Attribute {
	valptr := &a.native.attr
	var _v *Attribute // out
	_v = (*Attribute)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

// Features: featues, as a string in CSS syntax.
func (a *AttrFontFeatures) Features() string {
	valptr := &a.native.features
	var _v string // out
	_v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return _v
}

// NewAttrFontFeatures: create a new font features tag attribute.
//
// The function takes the following parameters:
//
//   - features: string with OpenType font features, in CSS syntax.
//
// The function returns the following values:
//
//   - attribute: newly allocated PangoAttribute, which should be freed with
//     pango.Attribute.Destroy().
//
func NewAttrFontFeatures(features string) *Attribute {
	var _arg1 *C.gchar          // out
	var _cret *C.PangoAttribute // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(features)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.pango_attr_font_features_new(_arg1)
	runtime.KeepAlive(features)

	var _attribute *Attribute // out

	_attribute = (*Attribute)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_attribute)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_attribute_destroy((*C.PangoAttribute)(intern.C))
		},
	)

	return _attribute
}
