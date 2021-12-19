// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
// extern gboolean _gotk4_atk1_ImageIface_set_image_description(AtkImage*, gchar*);
// extern gchar* _gotk4_atk1_ImageIface_get_image_description(AtkImage*);
// extern gchar* _gotk4_atk1_ImageIface_get_image_locale(AtkImage*);
// extern void _gotk4_atk1_ImageIface_get_image_position(AtkImage*, gint*, gint*, AtkCoordType);
// extern void _gotk4_atk1_ImageIface_get_image_size(AtkImage*, gint*, gint*);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_image_get_type()), F: marshalImager},
	})
}

// ImageOverrider contains methods that are overridable.
type ImageOverrider interface {
	// ImageDescription: get a textual description of this image.
	//
	// The function returns the following values:
	//
	//    - utf8: string representing the image description.
	//
	ImageDescription() string
	// ImageLocale retrieves the locale identifier associated to the Image.
	//
	// The function returns the following values:
	//
	//    - utf8 (optional): string corresponding to the POSIX LC_MESSAGES locale
	//      used by the image description, or NULL if the image does not specify
	//      a locale.
	//
	ImageLocale() string
	// ImagePosition gets the position of the image in the form of a point
	// specifying the images top-left corner.
	//
	// If the position can not be obtained (e.g. missing support), x and y are
	// set to -1.
	//
	// The function takes the following parameters:
	//
	//    - coordType specifies whether the coordinates are relative to the
	//      screen or to the components top level window.
	//
	// The function returns the following values:
	//
	//    - x (optional) address of #gint to put x coordinate position;
	//      otherwise, -1 if value cannot be obtained.
	//    - y (optional) address of #gint to put y coordinate position;
	//      otherwise, -1 if value cannot be obtained.
	//
	ImagePosition(coordType CoordType) (x int, y int)
	// ImageSize: get the width and height in pixels for the specified image.
	// The values of width and height are returned as -1 if the values cannot be
	// obtained (for instance, if the object is not onscreen).
	//
	// If the size can not be obtained (e.g. missing support), x and y are set
	// to -1.
	//
	// The function returns the following values:
	//
	//    - width (optional): filled with the image width, or -1 if the value
	//      cannot be obtained.
	//    - height (optional): filled with the image height, or -1 if the value
	//      cannot be obtained.
	//
	ImageSize() (width int, height int)
	// SetImageDescription sets the textual description for this image.
	//
	// The function takes the following parameters:
	//
	//    - description: string description to set for image.
	//
	// The function returns the following values:
	//
	//    - ok: boolean TRUE, or FALSE if operation could not be completed.
	//
	SetImageDescription(description string) bool
}

// Image should be implemented by Object subtypes on behalf of components which
// display image/pixmap information onscreen, and which provide information
// (other than just widget borders, etc.) via that image content. For instance,
// icons, buttons with icons, toolbar elements, and image viewing panes
// typically should implement Image.
//
// Image primarily provides two types of information: coordinate information
// (useful for screen review mode of screenreaders, and for use by onscreen
// magnifiers), and descriptive information. The descriptive information is
// provided for alternative, text-only presentation of the most significant
// information present in the image.
type Image struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Image)(nil)
)

// Imager describes Image's interface methods.
type Imager interface {
	externglib.Objector

	// ImageDescription: get a textual description of this image.
	ImageDescription() string
	// ImageLocale retrieves the locale identifier associated to the Image.
	ImageLocale() string
	// ImagePosition gets the position of the image in the form of a point
	// specifying the images top-left corner.
	ImagePosition(coordType CoordType) (x int, y int)
	// ImageSize: get the width and height in pixels for the specified image.
	ImageSize() (width int, height int)
	// SetImageDescription sets the textual description for this image.
	SetImageDescription(description string) bool
}

var _ Imager = (*Image)(nil)

func ifaceInitImager(gifacePtr, data C.gpointer) {
	iface := (*C.AtkImageIface)(unsafe.Pointer(gifacePtr))
	iface.get_image_description = (*[0]byte)(C._gotk4_atk1_ImageIface_get_image_description)
	iface.get_image_locale = (*[0]byte)(C._gotk4_atk1_ImageIface_get_image_locale)
	iface.get_image_position = (*[0]byte)(C._gotk4_atk1_ImageIface_get_image_position)
	iface.get_image_size = (*[0]byte)(C._gotk4_atk1_ImageIface_get_image_size)
	iface.set_image_description = (*[0]byte)(C._gotk4_atk1_ImageIface_set_image_description)
}

//export _gotk4_atk1_ImageIface_get_image_description
func _gotk4_atk1_ImageIface_get_image_description(arg0 *C.AtkImage) (cret *C.gchar) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ImageOverrider)

	utf8 := iface.ImageDescription()

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_atk1_ImageIface_get_image_locale
func _gotk4_atk1_ImageIface_get_image_locale(arg0 *C.AtkImage) (cret *C.gchar) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ImageOverrider)

	utf8 := iface.ImageLocale()

	if utf8 != "" {
		cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))
		defer C.free(unsafe.Pointer(cret))
	}

	return cret
}

//export _gotk4_atk1_ImageIface_get_image_position
func _gotk4_atk1_ImageIface_get_image_position(arg0 *C.AtkImage, arg1 *C.gint, arg2 *C.gint, arg3 C.AtkCoordType) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ImageOverrider)

	var _coordType CoordType // out

	_coordType = CoordType(arg3)

	x, y := iface.ImagePosition(_coordType)

	*arg1 = C.gint(x)
	*arg2 = C.gint(y)
}

//export _gotk4_atk1_ImageIface_get_image_size
func _gotk4_atk1_ImageIface_get_image_size(arg0 *C.AtkImage, arg1 *C.gint, arg2 *C.gint) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ImageOverrider)

	width, height := iface.ImageSize()

	*arg1 = C.gint(width)
	*arg2 = C.gint(height)
}

//export _gotk4_atk1_ImageIface_set_image_description
func _gotk4_atk1_ImageIface_set_image_description(arg0 *C.AtkImage, arg1 *C.gchar) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ImageOverrider)

	var _description string // out

	_description = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	ok := iface.SetImageDescription(_description)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapImage(obj *externglib.Object) *Image {
	return &Image{
		Object: obj,
	}
}

func marshalImager(p uintptr) (interface{}, error) {
	return wrapImage(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ImageDescription: get a textual description of this image.
//
// The function returns the following values:
//
//    - utf8: string representing the image description.
//
func (image *Image) ImageDescription() string {
	var _arg0 *C.AtkImage // out
	var _cret *C.gchar    // in

	_arg0 = (*C.AtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.atk_image_get_image_description(_arg0)
	runtime.KeepAlive(image)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ImageLocale retrieves the locale identifier associated to the Image.
//
// The function returns the following values:
//
//    - utf8 (optional): string corresponding to the POSIX LC_MESSAGES locale
//      used by the image description, or NULL if the image does not specify a
//      locale.
//
func (image *Image) ImageLocale() string {
	var _arg0 *C.AtkImage // out
	var _cret *C.gchar    // in

	_arg0 = (*C.AtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.atk_image_get_image_locale(_arg0)
	runtime.KeepAlive(image)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ImagePosition gets the position of the image in the form of a point
// specifying the images top-left corner.
//
// If the position can not be obtained (e.g. missing support), x and y are set
// to -1.
//
// The function takes the following parameters:
//
//    - coordType specifies whether the coordinates are relative to the screen or
//      to the components top level window.
//
// The function returns the following values:
//
//    - x (optional) address of #gint to put x coordinate position; otherwise, -1
//      if value cannot be obtained.
//    - y (optional) address of #gint to put y coordinate position; otherwise, -1
//      if value cannot be obtained.
//
func (image *Image) ImagePosition(coordType CoordType) (x int, y int) {
	var _arg0 *C.AtkImage    // out
	var _arg1 C.gint         // in
	var _arg2 C.gint         // in
	var _arg3 C.AtkCoordType // out

	_arg0 = (*C.AtkImage)(unsafe.Pointer(image.Native()))
	_arg3 = C.AtkCoordType(coordType)

	C.atk_image_get_image_position(_arg0, &_arg1, &_arg2, _arg3)
	runtime.KeepAlive(image)
	runtime.KeepAlive(coordType)

	var _x int // out
	var _y int // out

	_x = int(_arg1)
	_y = int(_arg2)

	return _x, _y
}

// ImageSize: get the width and height in pixels for the specified image. The
// values of width and height are returned as -1 if the values cannot be
// obtained (for instance, if the object is not onscreen).
//
// If the size can not be obtained (e.g. missing support), x and y are set to
// -1.
//
// The function returns the following values:
//
//    - width (optional): filled with the image width, or -1 if the value cannot
//      be obtained.
//    - height (optional): filled with the image height, or -1 if the value
//      cannot be obtained.
//
func (image *Image) ImageSize() (width int, height int) {
	var _arg0 *C.AtkImage // out
	var _arg1 C.gint      // in
	var _arg2 C.gint      // in

	_arg0 = (*C.AtkImage)(unsafe.Pointer(image.Native()))

	C.atk_image_get_image_size(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(image)

	var _width int  // out
	var _height int // out

	_width = int(_arg1)
	_height = int(_arg2)

	return _width, _height
}

// SetImageDescription sets the textual description for this image.
//
// The function takes the following parameters:
//
//    - description: string description to set for image.
//
// The function returns the following values:
//
//    - ok: boolean TRUE, or FALSE if operation could not be completed.
//
func (image *Image) SetImageDescription(description string) bool {
	var _arg0 *C.AtkImage // out
	var _arg1 *C.gchar    // out
	var _cret C.gboolean  // in

	_arg0 = (*C.AtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(description)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.atk_image_set_image_description(_arg0, _arg1)
	runtime.KeepAlive(image)
	runtime.KeepAlive(description)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
