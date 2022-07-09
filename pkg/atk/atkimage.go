// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_atk1_ImageIface_set_image_description(void*, void*);
// extern gchar* _gotk4_atk1_ImageIface_get_image_description(void*);
// extern gchar* _gotk4_atk1_ImageIface_get_image_locale(void*);
// extern void _gotk4_atk1_ImageIface_get_image_size(void*, void*, void*);
import "C"

// GTypeImage returns the GType for the type Image.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeImage() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Atk", "Image").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalImage)
	return gtype
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
	ImageSize() (width, height int32)
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
//
// Image wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Image struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Image)(nil)
)

// Imager describes Image's interface methods.
type Imager interface {
	coreglib.Objector

	// ImageDescription: get a textual description of this image.
	ImageDescription() string
	// ImageLocale retrieves the locale identifier associated to the Image.
	ImageLocale() string
	// ImageSize: get the width and height in pixels for the specified image.
	ImageSize() (width, height int32)
	// SetImageDescription sets the textual description for this image.
	SetImageDescription(description string) bool
}

var _ Imager = (*Image)(nil)

func ifaceInitImager(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Atk", "ImageIface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_image_description"))) = unsafe.Pointer(C._gotk4_atk1_ImageIface_get_image_description)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_image_locale"))) = unsafe.Pointer(C._gotk4_atk1_ImageIface_get_image_locale)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_image_size"))) = unsafe.Pointer(C._gotk4_atk1_ImageIface_get_image_size)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("set_image_description"))) = unsafe.Pointer(C._gotk4_atk1_ImageIface_set_image_description)
}

//export _gotk4_atk1_ImageIface_get_image_description
func _gotk4_atk1_ImageIface_get_image_description(arg0 *C.void) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ImageOverrider)

	utf8 := iface.ImageDescription()

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_atk1_ImageIface_get_image_locale
func _gotk4_atk1_ImageIface_get_image_locale(arg0 *C.void) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ImageOverrider)

	utf8 := iface.ImageLocale()

	if utf8 != "" {
		cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
		defer C.free(unsafe.Pointer(cret))
	}

	return cret
}

//export _gotk4_atk1_ImageIface_get_image_size
func _gotk4_atk1_ImageIface_get_image_size(arg0 *C.void, arg1 *C.void, arg2 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ImageOverrider)

	width, height := iface.ImageSize()

	*arg1 = (*C.void)(unsafe.Pointer(width))
	*arg2 = (*C.void)(unsafe.Pointer(height))
}

//export _gotk4_atk1_ImageIface_set_image_description
func _gotk4_atk1_ImageIface_set_image_description(arg0 *C.void, arg1 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ImageOverrider)

	var _description string // out

	_description = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	ok := iface.SetImageDescription(_description)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapImage(obj *coreglib.Object) *Image {
	return &Image{
		Object: obj,
	}
}

func marshalImage(p uintptr) (interface{}, error) {
	return wrapImage(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ImageDescription: get a textual description of this image.
//
// The function returns the following values:
//
//    - utf8: string representing the image description.
//
func (image *Image) ImageDescription() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(image)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
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
func (image *Image) ImageSize() (width, height int32) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	runtime.KeepAlive(image)

	var _width int32  // out
	var _height int32 // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_width = *(*int32)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_height = *(*int32)(unsafe.Pointer(_outs[1]))
	}

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(description)))
	defer C.free(unsafe.Pointer(_args[1]))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(image)
	runtime.KeepAlive(description)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}
