// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"fmt"
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

// glib.Type values for pango-coverage.go.
var (
	GTypeCoverageLevel = coreglib.Type(C.pango_coverage_level_get_type())
	GTypeCoverage      = coreglib.Type(C.pango_coverage_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeCoverageLevel, F: marshalCoverageLevel},
		{T: GTypeCoverage, F: marshalCoverage},
	})
}

// CoverageLevel: PangoCoverageLevel is used to indicate how well a font can
// represent a particular Unicode character for a particular script.
//
// Since 1.44, only PANGO_COVERAGE_NONE and PANGO_COVERAGE_EXACT will be
// returned.
type CoverageLevel C.gint

const (
	// CoverageNone: character is not representable with the font.
	CoverageNone CoverageLevel = iota
	// CoverageFallback: character is represented in a way that may be
	// comprehensible but is not the correct graphical form. For instance, a
	// Hangul character represented as a a sequence of Jamos, or a Latin
	// transliteration of a Cyrillic word.
	CoverageFallback
	// CoverageApproximate: character is represented as basically the correct
	// graphical form, but with a stylistic variant inappropriate for the
	// current script.
	CoverageApproximate
	// CoverageExact: character is represented as the correct graphical form.
	CoverageExact
)

func marshalCoverageLevel(p uintptr) (interface{}, error) {
	return CoverageLevel(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CoverageLevel.
func (c CoverageLevel) String() string {
	switch c {
	case CoverageNone:
		return "None"
	case CoverageFallback:
		return "Fallback"
	case CoverageApproximate:
		return "Approximate"
	case CoverageExact:
		return "Exact"
	default:
		return fmt.Sprintf("CoverageLevel(%d)", c)
	}
}

// Coverage structure is a map from Unicode characters to CoverageLevel values.
//
// It is often necessary in Pango to determine if a particular font can
// represent a particular character, and also how well it can represent that
// character. The Coverage is a data structure that is used to represent that
// information. It is an opaque structure with no public fields.
type Coverage struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Coverage)(nil)
)

func wrapCoverage(obj *coreglib.Object) *Coverage {
	return &Coverage{
		Object: obj,
	}
}

func marshalCoverage(p uintptr) (interface{}, error) {
	return wrapCoverage(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCoverage: create a new PangoCoverage.
//
// The function returns the following values:
//
//    - coverage: newly allocated PangoCoverage, initialized to
//      PANGO_COVERAGE_NONE with a reference count of one, which should be freed
//      with pango_coverage_unref().
//
func NewCoverage() *Coverage {
	var _cret *C.PangoCoverage // in

	_cret = C.pango_coverage_new()

	var _coverage *Coverage // out

	_coverage = wrapCoverage(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _coverage
}

// Copy an existing PangoCoverage.
//
// The function returns the following values:
//
//    - ret: newly allocated PangoCoverage, with a reference count of one, which
//      should be freed with pango_coverage_unref().
//
func (coverage *Coverage) Copy() *Coverage {
	var _arg0 *C.PangoCoverage // out
	var _cret *C.PangoCoverage // in

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(coreglib.InternObject(coverage).Native()))

	_cret = C.pango_coverage_copy(_arg0)
	runtime.KeepAlive(coverage)

	var _ret *Coverage // out

	_ret = wrapCoverage(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _ret
}

// Get: determine whether a particular index is covered by coverage.
//
// The function takes the following parameters:
//
//    - index_: index to check.
//
// The function returns the following values:
//
//    - coverageLevel: coverage level of coverage for character index_.
//
func (coverage *Coverage) Get(index_ int) CoverageLevel {
	var _arg0 *C.PangoCoverage     // out
	var _arg1 C.int                // out
	var _cret C.PangoCoverageLevel // in

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(coreglib.InternObject(coverage).Native()))
	_arg1 = C.int(index_)

	_cret = C.pango_coverage_get(_arg0, _arg1)
	runtime.KeepAlive(coverage)
	runtime.KeepAlive(index_)

	var _coverageLevel CoverageLevel // out

	_coverageLevel = CoverageLevel(_cret)

	return _coverageLevel
}

// Max: set the coverage for each index in coverage to be the max (better) value
// of the current coverage for the index and the coverage for the corresponding
// index in other.
//
// Deprecated: This function does nothing.
//
// The function takes the following parameters:
//
//    - other PangoCoverage.
//
func (coverage *Coverage) Max(other *Coverage) {
	var _arg0 *C.PangoCoverage // out
	var _arg1 *C.PangoCoverage // out

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(coreglib.InternObject(coverage).Native()))
	_arg1 = (*C.PangoCoverage)(unsafe.Pointer(coreglib.InternObject(other).Native()))

	C.pango_coverage_max(_arg0, _arg1)
	runtime.KeepAlive(coverage)
	runtime.KeepAlive(other)
}

// Set: modify a particular index within coverage.
//
// The function takes the following parameters:
//
//    - index_: index to modify.
//    - level: new level for index_.
//
func (coverage *Coverage) Set(index_ int, level CoverageLevel) {
	var _arg0 *C.PangoCoverage     // out
	var _arg1 C.int                // out
	var _arg2 C.PangoCoverageLevel // out

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(coreglib.InternObject(coverage).Native()))
	_arg1 = C.int(index_)
	_arg2 = C.PangoCoverageLevel(level)

	C.pango_coverage_set(_arg0, _arg1, _arg2)
	runtime.KeepAlive(coverage)
	runtime.KeepAlive(index_)
	runtime.KeepAlive(level)
}

// ToBytes: convert a PangoCoverage structure into a flat binary format.
//
// Deprecated: This returns NULL.
//
// The function returns the following values:
//
//    - bytes: location to store result (must be freed with g_free()).
//
func (coverage *Coverage) ToBytes() []byte {
	var _arg0 *C.PangoCoverage // out
	var _arg1 *C.guchar        // in
	var _arg2 C.int            // in

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(coreglib.InternObject(coverage).Native()))

	C.pango_coverage_to_bytes(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(coverage)

	var _bytes []byte // out

	defer C.free(unsafe.Pointer(_arg1))
	_bytes = make([]byte, _arg2)
	copy(_bytes, unsafe.Slice((*byte)(unsafe.Pointer(_arg1)), _arg2))

	return _bytes
}

// CoverageFromBytes: convert data generated from pango_coverage_to_bytes() back
// to a PangoCoverage.
//
// Deprecated: This returns NULL.
//
// The function takes the following parameters:
//
//    - bytes: binary data representing a PangoCoverage.
//
// The function returns the following values:
//
//    - coverage (optional): newly allocated PangoCoverage, or NULL if the data
//      was invalid.
//
func CoverageFromBytes(bytes []byte) *Coverage {
	var _arg1 *C.guchar // out
	var _arg2 C.int
	var _cret *C.PangoCoverage // in

	_arg2 = (C.int)(len(bytes))
	if len(bytes) > 0 {
		_arg1 = (*C.guchar)(unsafe.Pointer(&bytes[0]))
	}

	_cret = C.pango_coverage_from_bytes(_arg1, _arg2)
	runtime.KeepAlive(bytes)

	var _coverage *Coverage // out

	if _cret != nil {
		_coverage = wrapCoverage(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _coverage
}
