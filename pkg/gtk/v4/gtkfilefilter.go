// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_file_filter_get_type()), F: marshalFileFilterer},
	})
}

// FileFilter: GtkFileFilter filters files by name or mime type.
//
// GtkFileFilter can be used to restrict the files being shown in a
// GtkFileChooser. Files can be filtered based on their name (with
// gtk.FileFilter.AddPattern()) or on their mime type (with
// gtk.FileFilter.AddMIMEType()).
//
// Filtering by mime types handles aliasing and subclassing of mime types; e.g.
// a filter for text/plain also matches a file with mime type application/rtf,
// since application/rtf is a subclass of text/plain. Note that GtkFileFilter
// allows wildcards for the subtype of a mime type, so you can e.g. filter for
// image/\*.
//
// Normally, file filters are used by adding them to a GtkFileChooser (see
// gtk.FileChooser.AddFilter()), but it is also possible to manually use a file
// filter on any gtk.FilterListModel containing GFileInfo objects.
//
//
// GtkFileFilter as GtkBuildable
//
// The GtkFileFilter implementation of the GtkBuildable interface supports
// adding rules using the <mime-types> and <patterns> elements and listing the
// rules within. Specifying a <mime-type> or <pattern> has the same effect as as
// calling gtk.FileFilter.AddMIMEType() or gtk.FileFilter.AddPattern().
//
// An example of a UI definition fragment specifying GtkFileFilter rules:
//
//    <object class="GtkFileFilter">
//      <property name="name" translatable="yes">Text and Images</property>
//      <mime-types>
//        <mime-type>text/plain</mime-type>
//        <mime-type>image/ *</mime-type>
//      </mime-types>
//      <patterns>
//        <pattern>*.txt</pattern>
//        <pattern>*.png</pattern>
//      </patterns>
//    </object>.
type FileFilter struct {
	Filter

	*externglib.Object
	Buildable
}

var (
	_ externglib.Objector = (*FileFilter)(nil)
)

func wrapFileFilter(obj *externglib.Object) *FileFilter {
	return &FileFilter{
		Filter: Filter{
			Object: obj,
		},
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalFileFilterer(p uintptr) (interface{}, error) {
	return wrapFileFilter(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFileFilter creates a new GtkFileFilter with no rules added to it.
//
// Such a filter doesn’t accept any files, so is not particularly useful until
// you add rules with gtk.FileFilter.AddMIMEType(), gtk.FileFilter.AddPattern(),
// or gtk.FileFilter.AddPixbufFormats().
//
// To create a filter that accepts any file, use:
//
//    GtkFileFilter *filter = gtk_file_filter_new ();
//    gtk_file_filter_add_pattern (filter, "*");.
//
// The function returns the following values:
//
//    - fileFilter: new GtkFileFilter.
//
func NewFileFilter() *FileFilter {
	var _cret *C.GtkFileFilter // in

	_cret = C.gtk_file_filter_new()

	var _fileFilter *FileFilter // out

	_fileFilter = wrapFileFilter(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _fileFilter
}

// NewFileFilterFromGVariant: deserialize a file filter from a GVariant.
//
// The variant must be in the format produced by gtk.FileFilter.ToGVariant().
//
// The function takes the following parameters:
//
//    - variant: a{sv} GVariant.
//
// The function returns the following values:
//
//    - fileFilter: new GtkFileFilter object.
//
func NewFileFilterFromGVariant(variant *glib.Variant) *FileFilter {
	var _arg1 *C.GVariant      // out
	var _cret *C.GtkFileFilter // in

	_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(variant)))

	_cret = C.gtk_file_filter_new_from_gvariant(_arg1)
	runtime.KeepAlive(variant)

	var _fileFilter *FileFilter // out

	_fileFilter = wrapFileFilter(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _fileFilter
}

// AddMIMEType adds a rule allowing a given mime type to filter.
//
// The function takes the following parameters:
//
//    - mimeType: name of a MIME type.
//
func (filter *FileFilter) AddMIMEType(mimeType string) {
	var _arg0 *C.GtkFileFilter // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_filter_add_mime_type(_arg0, _arg1)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(mimeType)
}

// AddPattern adds a rule allowing a shell style glob to a filter.
//
// The function takes the following parameters:
//
//    - pattern: shell style glob.
//
func (filter *FileFilter) AddPattern(pattern string) {
	var _arg0 *C.GtkFileFilter // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(pattern)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_filter_add_pattern(_arg0, _arg1)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(pattern)
}

// AddPixbufFormats adds a rule allowing image files in the formats supported by
// GdkPixbuf.
//
// This is equivalent to calling gtk.FileFilter.AddMIMEType() for all the
// supported mime types.
func (filter *FileFilter) AddPixbufFormats() {
	var _arg0 *C.GtkFileFilter // out

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_file_filter_add_pixbuf_formats(_arg0)
	runtime.KeepAlive(filter)
}

// Attributes gets the attributes that need to be filled in for the GFileInfo
// passed to this filter.
//
// This function will not typically be used by applications; it is intended
// principally for use in the implementation of GtkFileChooser.
//
// The function returns the following values:
//
//    - utf8s: attributes.
//
func (filter *FileFilter) Attributes() []string {
	var _arg0 *C.GtkFileFilter // out
	var _cret **C.char         // in

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(filter.Native()))

	_cret = C.gtk_file_filter_get_attributes(_arg0)
	runtime.KeepAlive(filter)

	var _utf8s []string // out

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// Name gets the human-readable name for the filter.
//
// See gtk.FileFilter.SetName().
//
// The function returns the following values:
//
//    - utf8 (optional): human-readable name of the filter, or NULL. This value
//      is owned by GTK and must not be modified or freed.
//
func (filter *FileFilter) Name() string {
	var _arg0 *C.GtkFileFilter // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(filter.Native()))

	_cret = C.gtk_file_filter_get_name(_arg0)
	runtime.KeepAlive(filter)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SetName sets a human-readable name of the filter.
//
// This is the string that will be displayed in the file chooser if there is a
// selectable list of filters.
//
// The function takes the following parameters:
//
//    - name (optional) for the filter, or NULL to remove any existing name.
//
func (filter *FileFilter) SetName(name string) {
	var _arg0 *C.GtkFileFilter // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(filter.Native()))
	if name != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_file_filter_set_name(_arg0, _arg1)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(name)
}

// ToGVariant: serialize a file filter to an a{sv} variant.
//
// The function returns the following values:
//
//    - variant: new, floating, GVariant.
//
func (filter *FileFilter) ToGVariant() *glib.Variant {
	var _arg0 *C.GtkFileFilter // out
	var _cret *C.GVariant      // in

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(filter.Native()))

	_cret = C.gtk_file_filter_to_gvariant(_arg0)
	runtime.KeepAlive(filter)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	return _variant
}
