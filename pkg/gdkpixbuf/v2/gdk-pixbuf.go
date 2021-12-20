// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gdk-pixbuf-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_pixbuf_get_type()), F: marshalPixbuffer},
	})
}

// Pixbuf: pixel buffer.
//
// GdkPixbuf contains information about an image's pixel data, its color space,
// bits per sample, width and height, and the rowstride (the number of bytes
// between the start of one row and the start of the next).
//
//
// Creating new GdkPixbuf
//
// The most basic way to create a pixbuf is to wrap an existing pixel buffer
// with a gdkpixbuf.Pixbuf instance. You can use the
// gdkpixbuf.Pixbuf.NewFromData function to do this.
//
// Every time you create a new GdkPixbuf instance for some data, you will need
// to specify the destroy notification function that will be called when the
// data buffer needs to be freed; this will happen when a GdkPixbuf is finalized
// by the reference counting functions. If you have a chunk of static data
// compiled into your application, you can pass in NULL as the destroy
// notification function so that the data will not be freed.
//
// The gdkpixbuf.Pixbuf.New constructor function can be used as a convenience to
// create a pixbuf with an empty buffer; this is equivalent to allocating a data
// buffer using malloc() and then wrapping it with gdk_pixbuf_new_from_data().
// The gdk_pixbuf_new() function will compute an optimal rowstride so that
// rendering can be performed with an efficient algorithm.
//
// As a special case, you can use the gdkpixbuf.Pixbuf.NewFromXPMData function
// to create a pixbuf from inline XPM image data.
//
// You can also copy an existing pixbuf with the pixbuf.Copy function. This is
// not the same as just acquiring a reference to the old pixbuf instance: the
// copy function will actually duplicate the pixel data in memory and create a
// new pixbuf instance for it.
//
//
// Reference counting
//
// GdkPixbuf structures are reference counted. This means that an application
// can share a single pixbuf among many parts of the code. When a piece of the
// program needs to use a pixbuf, it should acquire a reference to it by calling
// g_object_ref(); when it no longer needs the pixbuf, it should release the
// reference it acquired by calling g_object_unref(). The resources associated
// with a GdkPixbuf will be freed when its reference count drops to zero.
// Newly-created GdkPixbuf instances start with a reference count of one.
//
//
// Image Data
//
// Image data in a pixbuf is stored in memory in an uncompressed, packed format.
// Rows in the image are stored top to bottom, and in each row pixels are stored
// from left to right.
//
// There may be padding at the end of a row.
//
// The "rowstride" value of a pixbuf, as returned by
// gdkpixbuf.Pixbuf.GetRowstride(), indicates the number of bytes between rows.
//
// **NOTE**: If you are copying raw pixbuf data with memcpy() note that the last
// row in the pixbuf may not be as wide as the full rowstride, but rather just
// as wide as the pixel data needs to be; that is: it is unsafe to do memcpy
// (dest, pixels, rowstride * height) to copy a whole pixbuf. Use
// gdkpixbuf.Pixbuf.Copy() instead, or compute the width in bytes of the last
// row as:
//
//    last_row = width * ((n_channels * bits_per_sample + 7) / 8);
//
//
// The same rule applies when iterating over each row of a GdkPixbuf pixels
// array.
//
// The following code illustrates a simple put_pixel() function for RGB pixbufs
// with 8 bits per channel with an alpha channel.
//
//    static void
//    put_pixel (GdkPixbuf *pixbuf,
//               int x,
//    	   int y,
//    	   guchar red,
//    	   guchar green,
//    	   guchar blue,
//    	   guchar alpha)
//    {
//      int n_channels = gdk_pixbuf_get_n_channels (pixbuf);
//
//      // Ensure that the pixbuf is valid
//      g_assert (gdk_pixbuf_get_colorspace (pixbuf) == GDK_COLORSPACE_RGB);
//      g_assert (gdk_pixbuf_get_bits_per_sample (pixbuf) == 8);
//      g_assert (gdk_pixbuf_get_has_alpha (pixbuf));
//      g_assert (n_channels == 4);
//
//      int width = gdk_pixbuf_get_width (pixbuf);
//      int height = gdk_pixbuf_get_height (pixbuf);
//
//      // Ensure that the coordinates are in a valid range
//      g_assert (x >= 0 && x < width);
//      g_assert (y >= 0 && y < height);
//
//      int rowstride = gdk_pixbuf_get_rowstride (pixbuf);
//
//      // The pixel buffer in the GdkPixbuf instance
//      guchar *pixels = gdk_pixbuf_get_pixels (pixbuf);
//
//      // The pixel we wish to modify
//      guchar *p = pixels + y * rowstride + x * n_channels;
//      p[0] = red;
//      p[1] = green;
//      p[2] = blue;
//      p[3] = alpha;
//    }
//
//
//
// Loading images
//
// The GdkPixBuf class provides a simple mechanism for loading an image from a
// file in synchronous and asynchronous fashion.
//
// For GUI applications, it is recommended to use the asynchronous stream API to
// avoid blocking the control flow of the application.
//
// Additionally, GdkPixbuf provides the gdkpixbuf.PixbufLoader` API for
// progressive image loading.
//
//
// Saving images
//
// The GdkPixbuf class provides methods for saving image data in a number of
// file formats. The formatted data can be written to a file or to a memory
// buffer. GdkPixbuf can also call a user-defined callback on the data, which
// allows to e.g. write the image to a socket or store it in a database.
type Pixbuf struct {
	*externglib.Object

	gio.LoadableIcon

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ externglib.Objector = (*Pixbuf)(nil)
)

func wrapPixbuf(obj *externglib.Object) *Pixbuf {
	return &Pixbuf{
		Object: obj,
		LoadableIcon: gio.LoadableIcon{
			Icon: gio.Icon{
				Object: obj,
			},
		},
	}
}

func marshalPixbuffer(p uintptr) (interface{}, error) {
	return wrapPixbuf(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
