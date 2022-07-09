// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypePicture returns the GType for the type Picture.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypePicture() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Picture").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalPicture)
	return gtype
}

// PictureOverrider contains methods that are overridable.
type PictureOverrider interface {
}

// Picture: GtkPicture widget displays a GdkPaintable.
//
// !An example GtkPicture (picture.png)
//
// Many convenience functions are provided to make pictures simple to use. For
// example, if you want to load an image from a file, and then display it,
// there’s a convenience function to do this:
//
//    GtkWidget *widget = gtk_picture_new_for_filename ("myfile.png");
//
//
// If the file isn’t loaded successfully, the picture will contain a “broken
// image” icon similar to that used in many web browsers. If you want to handle
// errors in loading the file yourself, for example by displaying an error
// message, then load the image with gdk.Texture.NewFromFile, then create the
// GtkPicture with gtk.Picture.NewForPaintable.
//
// Sometimes an application will want to avoid depending on external data files,
// such as image files. See the documentation of GResource for details. In this
// case, gtk.Picture.NewForResource and gtk.Picture.SetResource() should be
// used.
//
// GtkPicture displays an image at its natural size. See gtk.Image if you want
// to display a fixed-size image, such as an icon.
//
//
// Sizing the paintable
//
// You can influence how the paintable is displayed inside the GtkPicture. By
// turning off gtk.Picture:keep-aspect-ratio you can allow the paintable to get
// stretched. gtk.Picture:can-shrink can be unset to make sure that paintables
// are never made smaller than their ideal size - but be careful if you do not
// know the size of the paintable in use (like when displaying user-loaded
// images). This can easily cause the picture to grow larger than the screen.
// And gtkwidget:halign and gtkwidget:valign can be used to make sure the
// paintable doesn't fill all available space but is instead displayed at its
// original size.
//
//
// CSS nodes
//
// GtkPicture has a single CSS node with the name picture.
//
//
// Accessibility
//
// GtkPicture uses the GTK_ACCESSIBLE_ROLE_IMG role.
type Picture struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Picture)(nil)
)

func classInitPicturer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapPicture(obj *coreglib.Object) *Picture {
	return &Picture{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalPicture(p uintptr) (interface{}, error) {
	return wrapPicture(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPicture creates a new empty GtkPicture widget.
//
// The function returns the following values:
//
//    - picture: newly created GtkPicture widget.
//
func NewPicture() *Picture {
	_gret := girepository.MustFind("Gtk", "Picture").InvokeMethod("new_Picture", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _picture *Picture // out

	_picture = wrapPicture(coreglib.Take(unsafe.Pointer(_cret)))

	return _picture
}

// NewPictureForFile creates a new GtkPicture displaying the given file.
//
// If the file isn’t found or can’t be loaded, the resulting GtkPicture is
// empty.
//
// If you need to detect failures to load the file, use gdk.Texture.NewFromFile
// to load the file yourself, then create the GtkPicture from the texture.
//
// The function takes the following parameters:
//
//    - file (optional): GFile.
//
// The function returns the following values:
//
//    - picture: new GtkPicture.
//
func NewPictureForFile(file gio.Filer) *Picture {
	var _args [1]girepository.Argument

	if file != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(file).Native()))
	}

	_gret := girepository.MustFind("Gtk", "Picture").InvokeMethod("new_Picture_for_file", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(file)

	var _picture *Picture // out

	_picture = wrapPicture(coreglib.Take(unsafe.Pointer(_cret)))

	return _picture
}

// NewPictureForFilename creates a new GtkPicture displaying the file filename.
//
// This is a utility function that calls gtk.Picture.NewForFile. See that
// function for details.
//
// The function takes the following parameters:
//
//    - filename (optional): filename.
//
// The function returns the following values:
//
//    - picture: new GtkPicture.
//
func NewPictureForFilename(filename string) *Picture {
	var _args [1]girepository.Argument

	if filename != "" {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(filename)))
		defer C.free(unsafe.Pointer(_args[0]))
	}

	_gret := girepository.MustFind("Gtk", "Picture").InvokeMethod("new_Picture_for_filename", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(filename)

	var _picture *Picture // out

	_picture = wrapPicture(coreglib.Take(unsafe.Pointer(_cret)))

	return _picture
}

// NewPictureForPaintable creates a new GtkPicture displaying paintable.
//
// The GtkPicture will track changes to the paintable and update its size and
// contents in response to it.
//
// The function takes the following parameters:
//
//    - paintable (optional): GdkPaintable, or NULL.
//
// The function returns the following values:
//
//    - picture: new GtkPicture.
//
func NewPictureForPaintable(paintable gdk.Paintabler) *Picture {
	var _args [1]girepository.Argument

	if paintable != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))
	}

	_gret := girepository.MustFind("Gtk", "Picture").InvokeMethod("new_Picture_for_paintable", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(paintable)

	var _picture *Picture // out

	_picture = wrapPicture(coreglib.Take(unsafe.Pointer(_cret)))

	return _picture
}

// NewPictureForPixbuf creates a new GtkPicture displaying pixbuf.
//
// This is a utility function that calls gtk.Picture.NewForPaintable, See that
// function for details.
//
// The pixbuf must not be modified after passing it to this function.
//
// The function takes the following parameters:
//
//    - pixbuf (optional): GdkPixbuf, or NULL.
//
// The function returns the following values:
//
//    - picture: new GtkPicture.
//
func NewPictureForPixbuf(pixbuf *gdkpixbuf.Pixbuf) *Picture {
	var _args [1]girepository.Argument

	if pixbuf != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))
	}

	_gret := girepository.MustFind("Gtk", "Picture").InvokeMethod("new_Picture_for_pixbuf", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(pixbuf)

	var _picture *Picture // out

	_picture = wrapPicture(coreglib.Take(unsafe.Pointer(_cret)))

	return _picture
}

// NewPictureForResource creates a new GtkPicture displaying the resource at
// resource_path.
//
// This is a utility function that calls gtk.Picture.NewForFile. See that
// function for details.
//
// The function takes the following parameters:
//
//    - resourcePath (optional): resource path to play back.
//
// The function returns the following values:
//
//    - picture: new GtkPicture.
//
func NewPictureForResource(resourcePath string) *Picture {
	var _args [1]girepository.Argument

	if resourcePath != "" {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(resourcePath)))
		defer C.free(unsafe.Pointer(_args[0]))
	}

	_gret := girepository.MustFind("Gtk", "Picture").InvokeMethod("new_Picture_for_resource", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(resourcePath)

	var _picture *Picture // out

	_picture = wrapPicture(coreglib.Take(unsafe.Pointer(_cret)))

	return _picture
}

// AlternativeText gets the alternative textual description of the picture.
//
// The returned string will be NULL if the picture cannot be described
// textually.
//
// The function returns the following values:
//
//    - utf8 (optional): alternative textual description of self.
//
func (self *Picture) AlternativeText() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "Picture").InvokeMethod("get_alternative_text", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// CanShrink returns whether the GtkPicture respects its contents size.
//
// The function returns the following values:
//
//    - ok: TRUE if the picture can be made smaller than its contents.
//
func (self *Picture) CanShrink() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "Picture").InvokeMethod("get_can_shrink", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// File gets the GFile currently displayed if self is displaying a file.
//
// If self is not displaying a file, for example when gtk.Picture.SetPaintable()
// was used, then NULL is returned.
//
// The function returns the following values:
//
//    - file (optional): GFile displayed by self.
//
func (self *Picture) File() *gio.File {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "Picture").InvokeMethod("get_file", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _file *gio.File // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_file = &gio.File{
				Object: obj,
			}
		}
	}

	return _file
}

// KeepAspectRatio returns whether the GtkPicture preserves its contents aspect
// ratio.
//
// The function returns the following values:
//
//    - ok: TRUE if the self tries to keep the contents' aspect ratio.
//
func (self *Picture) KeepAspectRatio() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "Picture").InvokeMethod("get_keep_aspect_ratio", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Paintable gets the GdkPaintable being displayed by the GtkPicture.
//
// The function returns the following values:
//
//    - paintable (optional): displayed paintable, or NULL if the picture is
//      empty.
//
func (self *Picture) Paintable() *gdk.Paintable {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "Picture").InvokeMethod("get_paintable", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _paintable *gdk.Paintable // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_paintable = &gdk.Paintable{
				Object: obj,
			}
		}
	}

	return _paintable
}

// SetAlternativeText sets an alternative textual description for the picture
// contents.
//
// It is equivalent to the "alt" attribute for images on websites.
//
// This text will be made available to accessibility tools.
//
// If the picture cannot be described textually, set this property to NULL.
//
// The function takes the following parameters:
//
//    - alternativeText (optional): textual description of the contents.
//
func (self *Picture) SetAlternativeText(alternativeText string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if alternativeText != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(alternativeText)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "Picture").InvokeMethod("set_alternative_text", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(alternativeText)
}

// SetCanShrink: if set to TRUE, the self can be made smaller than its contents.
//
// The contents will then be scaled down when rendering.
//
// If you want to still force a minimum size manually, consider using
// gtk.Widget.SetSizeRequest().
//
// Also of note is that a similar function for growing does not exist because
// the grow behavior can be controlled via gtk.Widget.SetHAlign() and
// gtk.Widget.SetVAlign().
//
// The function takes the following parameters:
//
//    - canShrink: if self can be made smaller than its contents.
//
func (self *Picture) SetCanShrink(canShrink bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if canShrink {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Picture").InvokeMethod("set_can_shrink", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(canShrink)
}

// SetFile makes self load and display file.
//
// See gtk.Picture.NewForFile for details.
//
// The function takes the following parameters:
//
//    - file (optional): GFile or NULL.
//
func (self *Picture) SetFile(file gio.Filer) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if file != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(file).Native()))
	}

	girepository.MustFind("Gtk", "Picture").InvokeMethod("set_file", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(file)
}

// SetFilename makes self load and display the given filename.
//
// This is a utility function that calls gtk.Picture.SetFile().
//
// The function takes the following parameters:
//
//    - filename (optional) to play.
//
func (self *Picture) SetFilename(filename string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if filename != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(filename)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "Picture").InvokeMethod("set_filename", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(filename)
}

// SetKeepAspectRatio: if set to TRUE, the self will render its contents
// according to their aspect ratio.
//
// That means that empty space may show up at the top/bottom or left/right of
// self.
//
// If set to FALSE or if the contents provide no aspect ratio, the contents will
// be stretched over the picture's whole area.
//
// The function takes the following parameters:
//
//    - keepAspectRatio: whether to keep aspect ratio.
//
func (self *Picture) SetKeepAspectRatio(keepAspectRatio bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if keepAspectRatio {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Picture").InvokeMethod("set_keep_aspect_ratio", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(keepAspectRatio)
}

// SetPaintable makes self display the given paintable.
//
// If paintable is NULL, nothing will be displayed.
//
// See gtk.Picture.NewForPaintable for details.
//
// The function takes the following parameters:
//
//    - paintable (optional): GdkPaintable or NULL.
//
func (self *Picture) SetPaintable(paintable gdk.Paintabler) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if paintable != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))
	}

	girepository.MustFind("Gtk", "Picture").InvokeMethod("set_paintable", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(paintable)
}

// SetPixbuf sets a GtkPicture to show a GdkPixbuf.
//
// See gtk.Picture.NewForPixbuf for details.
//
// This is a utility function that calls gtk.Picture.SetPaintable().
//
// The function takes the following parameters:
//
//    - pixbuf (optional): GdkPixbuf or NULL.
//
func (self *Picture) SetPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if pixbuf != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))
	}

	girepository.MustFind("Gtk", "Picture").InvokeMethod("set_pixbuf", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(pixbuf)
}

// SetResource makes self load and display the resource at the given
// resource_path.
//
// This is a utility function that calls gtk.Picture.SetFile().
//
// The function takes the following parameters:
//
//    - resourcePath (optional): resource to set.
//
func (self *Picture) SetResource(resourcePath string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if resourcePath != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(resourcePath)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "Picture").InvokeMethod("set_resource", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(resourcePath)
}
