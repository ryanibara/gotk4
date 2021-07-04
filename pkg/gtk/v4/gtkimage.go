// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
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
		{T: externglib.Type(C.gtk_image_type_get_type()), F: marshalImageType},
		{T: externglib.Type(C.gtk_image_get_type()), F: marshalImage},
	})
}

// ImageType describes the image data representation used by a
// [class@Gtk.Image].
//
// If you want to get the image from the widget, you can only get the
// currently-stored representation; for instance, if the
// gtk_image_get_storage_type() returns GTK_IMAGE_PAINTABLE, then you can call
// gtk_image_get_paintable().
//
// For empty images, you can request any storage type (call any of the "get"
// functions), but they will all return nil values.
type ImageType int

const (
	// empty: there is no image displayed by the widget
	ImageTypeEmpty ImageType = 0
	// IconName: the widget contains a named icon
	ImageTypeIconName ImageType = 1
	// gicon: the widget contains a #GIcon
	ImageTypeGIcon ImageType = 2
	// paintable: the widget contains a Paintable
	ImageTypePaintable ImageType = 3
)

func marshalImageType(p uintptr) (interface{}, error) {
	return ImageType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Image: the `GtkImage` widget displays an image.
//
// !An example GtkImage (image.png)
//
// Various kinds of object can be displayed as an image; most typically, you
// would load a `GdkTexture` from a file, using the convenience function
// [ctor@Gtk.Image.new_from_file], for instance:
//
// “`c GtkWidget *image = gtk_image_new_from_file ("myfile.png"); “`
//
// If the file isn’t loaded successfully, the image will contain a “broken
// image” icon similar to that used in many web browsers.
//
// If you want to handle errors in loading the file yourself, for example by
// displaying an error message, then load the image with
// [ctor@Gdk.Texture.new_from_file], then create the `GtkImage` with
// [ctor@Gtk.Image.new_from_paintable].
//
// Sometimes an application will want to avoid depending on external data files,
// such as image files. See the documentation of `GResource` inside GIO, for
// details. In this case, [property@Gtk.Image:resource],
// [ctor@Gtk.Image.new_from_resource], and [method@Gtk.Image.set_from_resource]
// should be used.
//
// `GtkImage` displays its image as an icon, with a size that is determined by
// the application. See [class@Gtk.Picture] if you want to show an image at is
// actual size.
//
//
// CSS nodes
//
// `GtkImage` has a single CSS node with the name `image`. The style classes
// `.normal-icons` or `.large-icons` may appear, depending on the
// [property@Gtk.Image:icon-size] property.
//
//
// Accessibility
//
// `GtkImage` uses the `GTK_ACCESSIBLE_ROLE_IMG` role.
type Image interface {
	Widget

	ClearImage()

	IconName() string

	IconSize() IconSize

	PixelSize() int

	StorageType() ImageType

	SetFromFileImage(filename string)

	SetFromIconNameImage(iconName string)

	SetFromPixbufImage(pixbuf gdkpixbuf.Pixbuf)

	SetFromResourceImage(resourcePath string)

	SetIconSizeImage(iconSize IconSize)

	SetPixelSizeImage(pixelSize int)
}

// image implements the Image class.
type image struct {
	Widget
}

// WrapImage wraps a GObject to the right type. It is
// primarily used internally.
func WrapImage(obj *externglib.Object) Image {
	return image{
		Widget: WrapWidget(obj),
	}
}

func marshalImage(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapImage(obj), nil
}

func NewImage() Image {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_image_new()

	var _image Image // out

	_image = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Image)

	return _image
}

func NewImageFromFile(filename string) Image {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_image_new_from_file(_arg1)

	var _image Image // out

	_image = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Image)

	return _image
}

func NewImageFromIconName(iconName string) Image {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_image_new_from_icon_name(_arg1)

	var _image Image // out

	_image = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Image)

	return _image
}

func NewImageFromPixbuf(pixbuf gdkpixbuf.Pixbuf) Image {
	var _arg1 *C.GdkPixbuf // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	_cret = C.gtk_image_new_from_pixbuf(_arg1)

	var _image Image // out

	_image = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Image)

	return _image
}

func NewImageFromResource(resourcePath string) Image {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_image_new_from_resource(_arg1)

	var _image Image // out

	_image = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Image)

	return _image
}

func (i image) ClearImage() {
	var _arg0 *C.GtkImage // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	C.gtk_image_clear(_arg0)
}

func (i image) IconName() string {
	var _arg0 *C.GtkImage // out
	var _cret *C.char     // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_image_get_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (i image) IconSize() IconSize {
	var _arg0 *C.GtkImage   // out
	var _cret C.GtkIconSize // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_image_get_icon_size(_arg0)

	var _iconSize IconSize // out

	_iconSize = IconSize(_cret)

	return _iconSize
}

func (i image) PixelSize() int {
	var _arg0 *C.GtkImage // out
	var _cret C.int       // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_image_get_pixel_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (i image) StorageType() ImageType {
	var _arg0 *C.GtkImage    // out
	var _cret C.GtkImageType // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_image_get_storage_type(_arg0)

	var _imageType ImageType // out

	_imageType = ImageType(_cret)

	return _imageType
}

func (i image) SetFromFileImage(filename string) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_image_set_from_file(_arg0, _arg1)
}

func (i image) SetFromIconNameImage(iconName string) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_image_set_from_icon_name(_arg0, _arg1)
}

func (i image) SetFromPixbufImage(pixbuf gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkImage  // out
	var _arg1 *C.GdkPixbuf // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_image_set_from_pixbuf(_arg0, _arg1)
}

func (i image) SetFromResourceImage(resourcePath string) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_image_set_from_resource(_arg0, _arg1)
}

func (i image) SetIconSizeImage(iconSize IconSize) {
	var _arg0 *C.GtkImage   // out
	var _arg1 C.GtkIconSize // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	_arg1 = C.GtkIconSize(iconSize)

	C.gtk_image_set_icon_size(_arg0, _arg1)
}

func (i image) SetPixelSizeImage(pixelSize int) {
	var _arg0 *C.GtkImage // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(pixelSize)

	C.gtk_image_set_pixel_size(_arg0, _arg1)
}

func (s image) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s image) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s image) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s image) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s image) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s image) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s image) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b image) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
