// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk-pixbuf/gdk-pixbuf.h>
import "C"

// ApplyEmbeddedOrientation takes an existing pixbuf and checks for the presence
// of an associated "orientation" option.
//
// The orientation option may be provided by the JPEG loader (which reads the
// exif orientation tag) or the TIFF loader (which reads the TIFF orientation
// tag, and compensates it for the partial transforms performed by libtiff).
//
// If an orientation option/tag is present, the appropriate transform will be
// performed so that the pixbuf is oriented correctly.
//
// The function returns the following values:
//
//    - pixbuf (optional): newly-created pixbuf.
//
func (src *Pixbuf) ApplyEmbeddedOrientation() *Pixbuf {
	var _arg0 *C.GdkPixbuf // out
	var _cret *C.GdkPixbuf // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(coreglib.InternObject(src).Native()))

	_cret = C.gdk_pixbuf_apply_embedded_orientation(_arg0)
	runtime.KeepAlive(src)

	var _pixbuf *Pixbuf // out

	if _cret != nil {
		_pixbuf = wrapPixbuf(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _pixbuf
}
