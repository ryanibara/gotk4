// Code generated by girgen. DO NOT EDIT.

package gdkpixdata

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdkpixbuf"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-pixbuf-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixdata.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{

		// Records
		// Skipped Pixdata.

	})
}

// PixdataDumpType: an enumeration which is used by gdk_pixdata_to_csource() to
// determine the form of C source to be generated. The three values
// @GDK_PIXDATA_DUMP_PIXDATA_STREAM, @GDK_PIXDATA_DUMP_PIXDATA_STRUCT and
// @GDK_PIXDATA_DUMP_MACROS are mutually exclusive, as are
// @GDK_PIXBUF_DUMP_GTYPES and @GDK_PIXBUF_DUMP_CTYPES. The remaining elements
// are optional flags that can be freely added.
type PixdataDumpType int

const (
	// PixdataDumpTypePixdataStream: generate pixbuf data stream (a single
	// string containing a serialized Pixdata structure in network byte order).
	PixdataDumpTypePixdataStream PixdataDumpType = 0b0
	// PixdataDumpTypePixdataStruct: generate Pixdata structure (needs the
	// Pixdata structure definition from gdk-pixdata.h).
	PixdataDumpTypePixdataStruct PixdataDumpType = 0b1
	// PixdataDumpTypeMacros: generate <function>*_ROWSTRIDE</function>,
	// <function>*_WIDTH</function>, <function>*_HEIGHT</function>,
	// <function>*_BYTES_PER_PIXEL</function> and
	// <function>*_RLE_PIXEL_DATA</function> or
	// <function>*_PIXEL_DATA</function> macro definitions for the image.
	PixdataDumpTypeMacros PixdataDumpType = 0b10
	// PixdataDumpTypeGtypes: generate GLib data types instead of standard C
	// data types.
	PixdataDumpTypeGtypes PixdataDumpType = 0b0
	// PixdataDumpTypeCtypes: generate standard C data types instead of GLib
	// data types.
	PixdataDumpTypeCtypes PixdataDumpType = 0b100000000
	// PixdataDumpTypeStatic: generate static symbols.
	PixdataDumpTypeStatic PixdataDumpType = 0b1000000000
	// PixdataDumpTypeConst: generate const symbols.
	PixdataDumpTypeConst PixdataDumpType = 0b10000000000
	// PixdataDumpTypeRleDecoder: provide a
	// <function>*_RUN_LENGTH_DECODE(image_buf, rle_data, size, bpp)</function>
	// macro definition to decode run-length encoded image data.
	PixdataDumpTypeRleDecoder PixdataDumpType = 0b10000000000000000
)

// PixdataType: an enumeration containing three sets of flags for a Pixdata
// struct: one for the used colorspace, one for the width of the samples and one
// for the encoding of the pixel data.
type PixdataType int

const (
	// PixdataTypeColorTypeRgb: each pixel has red, green and blue samples.
	PixdataTypeColorTypeRgb PixdataType = 0b1
	// PixdataTypeColorTypeRgba: each pixel has red, green and blue samples and
	// an alpha value.
	PixdataTypeColorTypeRgba PixdataType = 0b10
	// PixdataTypeColorTypeMask: mask for the colortype flags of the enum.
	PixdataTypeColorTypeMask PixdataType = 0b11111111
	// PixdataTypeSampleWidth8: each sample has 8 bits.
	PixdataTypeSampleWidth8 PixdataType = 0b10000000000000000
	// PixdataTypeSampleWidthMask: mask for the sample width flags of the enum.
	PixdataTypeSampleWidthMask PixdataType = 0b11110000000000000000
	// PixdataTypeEncodingRaw: the pixel data is in raw form.
	PixdataTypeEncodingRaw PixdataType = 0b1000000000000000000000000
	// PixdataTypeEncodingRle: the pixel data is run-length encoded. Runs may be
	// up to 127 bytes long; their length is stored in a single byte preceding
	// the pixel data for the run. If a run is constant, its length byte has the
	// high bit set and the pixel data consists of a single pixel which must be
	// repeated.
	PixdataTypeEncodingRle PixdataType = 0b10000000000000000000000000
	// PixdataTypeEncodingMask: mask for the encoding flags of the enum.
	PixdataTypeEncodingMask PixdataType = 0b1111000000000000000000000000
)

// PixbufFromPixdata: converts a Pixdata to a Pixbuf. If @copy_pixels is true or
// if the pixel data is run-length-encoded, the pixel data is copied into
// newly-allocated memory; otherwise it is reused.
func PixbufFromPixdata(pixdata *Pixdata, copyPixels bool) *gdkpixbuf.Pixbuf

// Pixdata: a Pixdata contains pixbuf information in a form suitable for
// serialization and streaming.
type Pixdata struct {
	// Magic: magic number. A valid Pixdata structure must have
	// K_PIXBUF_MAGIC_NUMBER here.
	Magic uint32
	// Length: less than 1 to disable length checks, otherwise
	// K_PIXDATA_HEADER_LENGTH + length of @pixel_data.
	Length int32
	// PixdataType: information about colorspace, sample width and encoding, in
	// a PixdataType.
	PixdataType uint32
	// Rowstride: distance in bytes between rows.
	Rowstride uint32
	// Width: width of the image in pixels.
	Width uint32
	// Height: height of the image in pixels.
	Height uint32
	// PixelData: @width x @height pixels, encoded according to @pixdata_type
	// and @rowstride.
	PixelData []uint8

	native *C.GdkPixdata
}

func wrapPixdata(p *C.GdkPixdata) *Pixdata {
	var v Pixdata

	v.Magic = uint32(p.magic)
	v.Length = int32(p.length)
	v.PixdataType = uint32(p.pixdata_type)
	v.Rowstride = uint32(p.rowstride)
	v.Width = uint32(p.width)
	v.Height = uint32(p.height)

	return &v
}

func marshalPixdata(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.GdkPixdata)(unsafe.Pointer(b))

	return wrapPixdata(c)
}

// Native returns the pointer to *C.GdkPixdata. The caller is expected to
// cast.
func (p *Pixdata) Native() unsafe.Pointer {
	return unsafe.Pointer(p.native)
}
