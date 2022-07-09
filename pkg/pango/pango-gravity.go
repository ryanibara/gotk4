// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeGravity returns the GType for the type Gravity.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeGravity() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Pango", "Gravity").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalGravity)
	return gtype
}

// GTypeGravityHint returns the GType for the type GravityHint.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeGravityHint() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Pango", "GravityHint").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalGravityHint)
	return gtype
}

// Gravity: PangoGravity represents the orientation of glyphs in a segment of
// text.
//
// This is useful when rendering vertical text layouts. In those situations, the
// layout is rotated using a non-identity pango.Matrix, and then glyph
// orientation is controlled using PangoGravity.
//
// Not every value in this enumeration makes sense for every usage of
// PangoGravity; for example, PANGO_GRAVITY_AUTO only can be passed to
// pango.Context.SetBaseGravity() and can only be returned by
// pango.Context.GetBaseGravity().
//
// See also: pango.GravityHint.
type Gravity C.gint

const (
	// GravitySouth glyphs stand upright (default).
	GravitySouth Gravity = iota
	// GravityEast glyphs are rotated 90 degrees clockwise.
	GravityEast
	// GravityNorth glyphs are upside-down.
	GravityNorth
	// GravityWest glyphs are rotated 90 degrees counter-clockwise.
	GravityWest
	// GravityAuto: gravity is resolved from the context matrix.
	GravityAuto
)

func marshalGravity(p uintptr) (interface{}, error) {
	return Gravity(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for Gravity.
func (g Gravity) String() string {
	switch g {
	case GravitySouth:
		return "South"
	case GravityEast:
		return "East"
	case GravityNorth:
		return "North"
	case GravityWest:
		return "West"
	case GravityAuto:
		return "Auto"
	default:
		return fmt.Sprintf("Gravity(%d)", g)
	}
}

// GravityHint: PangoGravityHint defines how horizontal scripts should behave in
// a vertical context.
//
// That is, English excerpts in a vertical paragraph for example.
//
// See also pango.Gravity.
type GravityHint C.gint

const (
	// GravityHintNatural scripts will take their natural gravity based on the
	// base gravity and the script. This is the default.
	GravityHintNatural GravityHint = iota
	// GravityHintStrong always use the base gravity set, regardless of the
	// script.
	GravityHintStrong
	// GravityHintLine: for scripts not in their natural direction (eg. Latin in
	// East gravity), choose per-script gravity such that every script respects
	// the line progression. This means, Latin and Arabic will take opposite
	// gravities and both flow top-to-bottom for example.
	GravityHintLine
)

func marshalGravityHint(p uintptr) (interface{}, error) {
	return GravityHint(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for GravityHint.
func (g GravityHint) String() string {
	switch g {
	case GravityHintNatural:
		return "Natural"
	case GravityHintStrong:
		return "Strong"
	case GravityHintLine:
		return "Line"
	default:
		return fmt.Sprintf("GravityHint(%d)", g)
	}
}
