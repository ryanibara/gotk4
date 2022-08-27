// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypePrintSettings = coreglib.Type(C.gtk_print_settings_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePrintSettings, F: marshalPrintSettings},
	})
}

const PRINT_SETTINGS_COLLATE = "collate"
const PRINT_SETTINGS_DEFAULT_SOURCE = "default-source"
const PRINT_SETTINGS_DITHER = "dither"
const PRINT_SETTINGS_DUPLEX = "duplex"
const PRINT_SETTINGS_FINISHINGS = "finishings"
const PRINT_SETTINGS_MEDIA_TYPE = "media-type"
const PRINT_SETTINGS_NUMBER_UP = "number-up"
const PRINT_SETTINGS_NUMBER_UP_LAYOUT = "number-up-layout"
const PRINT_SETTINGS_N_COPIES = "n-copies"
const PRINT_SETTINGS_ORIENTATION = "orientation"
const PRINT_SETTINGS_OUTPUT_BIN = "output-bin"

// PRINT_SETTINGS_OUTPUT_FILE_FORMAT: key used by the “Print to file” printer to
// store the format of the output. The supported values are “PS” and “PDF”.
const PRINT_SETTINGS_OUTPUT_FILE_FORMAT = "output-file-format"

// PRINT_SETTINGS_OUTPUT_URI: key used by the “Print to file” printer to store
// the URI to which the output should be written. GTK+ itself supports only
// “file://” URIs.
const PRINT_SETTINGS_OUTPUT_URI = "output-uri"
const PRINT_SETTINGS_PAGE_RANGES = "page-ranges"
const PRINT_SETTINGS_PAGE_SET = "page-set"
const PRINT_SETTINGS_PAPER_FORMAT = "paper-format"
const PRINT_SETTINGS_PAPER_HEIGHT = "paper-height"
const PRINT_SETTINGS_PAPER_WIDTH = "paper-width"
const PRINT_SETTINGS_PRINTER = "printer"
const PRINT_SETTINGS_PRINTER_LPI = "printer-lpi"
const PRINT_SETTINGS_PRINT_PAGES = "print-pages"
const PRINT_SETTINGS_QUALITY = "quality"
const PRINT_SETTINGS_RESOLUTION = "resolution"
const PRINT_SETTINGS_RESOLUTION_X = "resolution-x"
const PRINT_SETTINGS_RESOLUTION_Y = "resolution-y"
const PRINT_SETTINGS_REVERSE = "reverse"
const PRINT_SETTINGS_SCALE = "scale"
const PRINT_SETTINGS_USE_COLOR = "use-color"
const PRINT_SETTINGS_WIN32_DRIVER_EXTRA = "win32-driver-extra"
const PRINT_SETTINGS_WIN32_DRIVER_VERSION = "win32-driver-version"

type PrintSettingsFunc func(key, value string)

// PrintSettings object represents the settings of a print dialog in a
// system-independent way. The main use for this object is that once you’ve
// printed you can get a settings object that represents the settings the user
// chose, and the next time you print you can pass that object in so that the
// user doesn’t have to re-set all his settings.
//
// Its also possible to enumerate the settings so that you can easily save the
// settings for the next time your app runs, or even store them in a document.
// The predefined keys try to use shared values as much as possible so that
// moving such a document between systems still works.
//
// Printing support was added in GTK+ 2.10.
type PrintSettings struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*PrintSettings)(nil)
)

func wrapPrintSettings(obj *coreglib.Object) *PrintSettings {
	return &PrintSettings{
		Object: obj,
	}
}

func marshalPrintSettings(p uintptr) (interface{}, error) {
	return wrapPrintSettings(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// PageRange: see also gtk_print_settings_set_page_ranges().
//
// An instance of this type is always passed by reference.
type PageRange struct {
	*pageRange
}

// pageRange is the struct that's finalized.
type pageRange struct {
	native *C.GtkPageRange
}

// NewPageRange creates a new PageRange instance from the given
// fields. Beware that this function allocates on the Go heap; be careful
// when using it!
func NewPageRange(start, end int) PageRange {
	var f0 C.gint // out
	f0 = C.gint(start)
	var f1 C.gint // out
	f1 = C.gint(end)

	v := C.GtkPageRange{
		start: f0,
		end:   f1,
	}

	return *(*PageRange)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

// Start: start of page range.
func (p *PageRange) Start() int {
	valptr := &p.native.start
	var _v int // out
	_v = int(*valptr)
	return _v
}

// End: end of page range.
func (p *PageRange) End() int {
	valptr := &p.native.end
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Start: start of page range.
func (p *PageRange) SetStart(start int) {
	valptr := &p.native.start
	*valptr = C.gint(start)
}

// End: end of page range.
func (p *PageRange) SetEnd(end int) {
	valptr := &p.native.end
	*valptr = C.gint(end)
}
