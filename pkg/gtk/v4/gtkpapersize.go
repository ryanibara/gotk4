// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_paper_size_get_type()), F: marshalPaperSize},
	})
}

// PaperSize: `GtkPaperSize` handles paper sizes.
//
// It uses the standard called PWG 5101.1-2002 PWG: Standard for Media
// Standardized Names (http://www.pwg.org/standards.html) to name the paper
// sizes (and to get the data for the page sizes). In addition to standard paper
// sizes, `GtkPaperSize` allows to construct custom paper sizes with arbitrary
// dimensions.
//
// The `GtkPaperSize` object stores not only the dimensions (width and height)
// of a paper size and its name, it also provides default print margins.
type PaperSize struct {
	native C.GtkPaperSize
}

// WrapPaperSize wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPaperSize(ptr unsafe.Pointer) *PaperSize {
	return (*PaperSize)(ptr)
}

func marshalPaperSize(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*PaperSize)(unsafe.Pointer(b)), nil
}

// NewPaperSize constructs a struct PaperSize.
func NewPaperSize(name string) PaperSize {
	var _arg1 *C.char         // out
	var _cret *C.GtkPaperSize // in

	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_paper_size_new(_arg1)

	var _paperSize PaperSize // out

	_paperSize = (PaperSize)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_paperSize, func(v PaperSize) {
		C.gtk_paper_size_free((*C.GtkPaperSize)(unsafe.Pointer(v)))
	})

	return _paperSize
}

// NewPaperSizeCustom constructs a struct PaperSize.
func NewPaperSizeCustom(name string, displayName string, width float64, height float64, unit Unit) PaperSize {
	var _arg1 *C.char         // out
	var _arg2 *C.char         // out
	var _arg3 C.double        // out
	var _arg4 C.double        // out
	var _arg5 C.GtkUnit       // out
	var _cret *C.GtkPaperSize // in

	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(displayName))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.double(width)
	_arg4 = C.double(height)
	_arg5 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_new_custom(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _paperSize PaperSize // out

	_paperSize = (PaperSize)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_paperSize, func(v PaperSize) {
		C.gtk_paper_size_free((*C.GtkPaperSize)(unsafe.Pointer(v)))
	})

	return _paperSize
}

// NewPaperSizeFromGVariant constructs a struct PaperSize.
func NewPaperSizeFromGVariant(variant glib.Variant) PaperSize {
	var _arg1 *C.GVariant     // out
	var _cret *C.GtkPaperSize // in

	_arg1 = (*C.GVariant)(unsafe.Pointer(variant))

	_cret = C.gtk_paper_size_new_from_gvariant(_arg1)

	var _paperSize PaperSize // out

	_paperSize = (PaperSize)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_paperSize, func(v PaperSize) {
		C.gtk_paper_size_free((*C.GtkPaperSize)(unsafe.Pointer(v)))
	})

	return _paperSize
}

// NewPaperSizeFromIpp constructs a struct PaperSize.
func NewPaperSizeFromIpp(ippName string, width float64, height float64) PaperSize {
	var _arg1 *C.char         // out
	var _arg2 C.double        // out
	var _arg3 C.double        // out
	var _cret *C.GtkPaperSize // in

	_arg1 = (*C.char)(C.CString(ippName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.double(width)
	_arg3 = C.double(height)

	_cret = C.gtk_paper_size_new_from_ipp(_arg1, _arg2, _arg3)

	var _paperSize PaperSize // out

	_paperSize = (PaperSize)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_paperSize, func(v PaperSize) {
		C.gtk_paper_size_free((*C.GtkPaperSize)(unsafe.Pointer(v)))
	})

	return _paperSize
}

// NewPaperSizeFromKeyFile constructs a struct PaperSize.
func NewPaperSizeFromKeyFile(keyFile glib.KeyFile, groupName string) (PaperSize, error) {
	var _arg1 *C.GKeyFile     // out
	var _arg2 *C.char         // out
	var _cret *C.GtkPaperSize // in
	var _cerr **C.GError      // in

	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile))
	_arg2 = (*C.char)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_paper_size_new_from_key_file(_arg1, _arg2, &_cerr)

	var _paperSize PaperSize // out
	var _goerr error         // out

	_paperSize = (PaperSize)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_paperSize, func(v PaperSize) {
		C.gtk_paper_size_free((*C.GtkPaperSize)(unsafe.Pointer(v)))
	})
	{
		var refTmpIn *C.GError
		var refTmpOut error

		refTmpIn = *_cerr

		refTmpOut = gerror.Take(unsafe.Pointer(refTmpIn))

		_goerr = refTmpOut
	}

	return _paperSize, _goerr
}

// NewPaperSizeFromPpd constructs a struct PaperSize.
func NewPaperSizeFromPpd(ppdName string, ppdDisplayName string, width float64, height float64) PaperSize {
	var _arg1 *C.char         // out
	var _arg2 *C.char         // out
	var _arg3 C.double        // out
	var _arg4 C.double        // out
	var _cret *C.GtkPaperSize // in

	_arg1 = (*C.char)(C.CString(ppdName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(ppdDisplayName))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.double(width)
	_arg4 = C.double(height)

	_cret = C.gtk_paper_size_new_from_ppd(_arg1, _arg2, _arg3, _arg4)

	var _paperSize PaperSize // out

	_paperSize = (PaperSize)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_paperSize, func(v PaperSize) {
		C.gtk_paper_size_free((*C.GtkPaperSize)(unsafe.Pointer(v)))
	})

	return _paperSize
}

// Native returns the underlying C source pointer.
func (p *PaperSize) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// Copy copies an existing `GtkPaperSize`.
func (o *PaperSize) Copy() PaperSize {
	var _arg0 *C.GtkPaperSize // out
	var _cret *C.GtkPaperSize // in

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(o))

	_cret = C.gtk_paper_size_copy(_arg0)

	var _paperSize PaperSize // out

	_paperSize = (PaperSize)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_paperSize, func(v PaperSize) {
		C.gtk_paper_size_free((*C.GtkPaperSize)(unsafe.Pointer(v)))
	})

	return _paperSize
}

// Free the given `GtkPaperSize` object.
func (s *PaperSize) Free() {
	var _arg0 *C.GtkPaperSize // out

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s))

	C.gtk_paper_size_free(_arg0)
}

// DefaultBottomMargin gets the default bottom margin for the `GtkPaperSize`.
func (s *PaperSize) DefaultBottomMargin(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.double        // in

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_default_bottom_margin(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DefaultLeftMargin gets the default left margin for the `GtkPaperSize`.
func (s *PaperSize) DefaultLeftMargin(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.double        // in

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_default_left_margin(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DefaultRightMargin gets the default right margin for the `GtkPaperSize`.
func (s *PaperSize) DefaultRightMargin(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.double        // in

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_default_right_margin(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DefaultTopMargin gets the default top margin for the `GtkPaperSize`.
func (s *PaperSize) DefaultTopMargin(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.double        // in

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_default_top_margin(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DisplayName gets the human-readable name of the `GtkPaperSize`.
func (s *PaperSize) DisplayName() string {
	var _arg0 *C.GtkPaperSize // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s))

	_cret = C.gtk_paper_size_get_display_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Height gets the paper height of the `GtkPaperSize`, in units of @unit.
func (s *PaperSize) Height(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.double        // in

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_height(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Name gets the name of the `GtkPaperSize`.
func (s *PaperSize) Name() string {
	var _arg0 *C.GtkPaperSize // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s))

	_cret = C.gtk_paper_size_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// PpdName gets the PPD name of the `GtkPaperSize`, which may be nil.
func (s *PaperSize) PpdName() string {
	var _arg0 *C.GtkPaperSize // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s))

	_cret = C.gtk_paper_size_get_ppd_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Width gets the paper width of the `GtkPaperSize`, in units of @unit.
func (s *PaperSize) Width(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.double        // in

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_width(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// IsCustom returns true if @size is not a standard paper size.
func (s *PaperSize) IsCustom() bool {
	var _arg0 *C.GtkPaperSize // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s))

	_cret = C.gtk_paper_size_is_custom(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsEqual compares two `GtkPaperSize` objects.
func (s *PaperSize) IsEqual(size2 PaperSize) bool {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 *C.GtkPaperSize // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s))
	_arg1 = (*C.GtkPaperSize)(unsafe.Pointer(size2))

	_cret = C.gtk_paper_size_is_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsIpp returns true if @size is an IPP standard paper size.
func (s *PaperSize) IsIpp() bool {
	var _arg0 *C.GtkPaperSize // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s))

	_cret = C.gtk_paper_size_is_ipp(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetSize changes the dimensions of a @size to @width x @height.
func (s *PaperSize) SetSize(width float64, height float64, unit Unit) {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.double        // out
	var _arg2 C.double        // out
	var _arg3 C.GtkUnit       // out

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s))
	_arg1 = C.double(width)
	_arg2 = C.double(height)
	_arg3 = C.GtkUnit(unit)

	C.gtk_paper_size_set_size(_arg0, _arg1, _arg2, _arg3)
}

// ToGVariant: serialize a paper size to an `a{sv}` variant.
func (p *PaperSize) ToGVariant() glib.Variant {
	var _arg0 *C.GtkPaperSize // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(p))

	_cret = C.gtk_paper_size_to_gvariant(_arg0)

	var _variant glib.Variant // out

	_variant = (glib.Variant)(unsafe.Pointer(_cret))
	C.g_variant_ref(_cret)

	return _variant
}

// ToKeyFile: this function adds the paper size from @size to @key_file.
func (s *PaperSize) ToKeyFile(keyFile glib.KeyFile, groupName string) {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 *C.GKeyFile     // out
	var _arg2 *C.char         // out

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s))
	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile))
	_arg2 = (*C.char)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_paper_size_to_key_file(_arg0, _arg1, _arg2)
}
