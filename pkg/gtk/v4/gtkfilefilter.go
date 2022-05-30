// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkfilefilter.go.
var GTypeFileFilter = coreglib.Type(C.gtk_file_filter_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeFileFilter, F: marshalFileFilter},
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
	_ [0]func() // equal guard
	Filter

	*coreglib.Object
	Buildable
}

var (
	_ coreglib.Objector = (*FileFilter)(nil)
)

func wrapFileFilter(obj *coreglib.Object) *FileFilter {
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

func marshalFileFilter(p uintptr) (interface{}, error) {
	return wrapFileFilter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
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
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "FileFilter").InvokeMethod("new_FileFilter", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _fileFilter *FileFilter // out

	_fileFilter = wrapFileFilter(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

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
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(variant)))
	*(**glib.Variant)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "FileFilter").InvokeMethod("new_FileFilter_from_gvariant", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(variant)

	var _fileFilter *FileFilter // out

	_fileFilter = wrapFileFilter(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _fileFilter
}

// AddMIMEType adds a rule allowing a given mime type to filter.
//
// The function takes the following parameters:
//
//    - mimeType: name of a MIME type.
//
func (filter *FileFilter) AddMIMEType(mimeType string) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**FileFilter)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "FileFilter").InvokeMethod("add_mime_type", args[:], nil)

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
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(pattern)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**FileFilter)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "FileFilter").InvokeMethod("add_pattern", args[:], nil)

	runtime.KeepAlive(filter)
	runtime.KeepAlive(pattern)
}

// AddPixbufFormats adds a rule allowing image files in the formats supported by
// GdkPixbuf.
//
// This is equivalent to calling gtk.FileFilter.AddMIMEType() for all the
// supported mime types.
func (filter *FileFilter) AddPixbufFormats() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	*(**FileFilter)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "FileFilter").InvokeMethod("add_pixbuf_formats", args[:], nil)

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
	var args [1]girepository.Argument
	var _arg0 *C.void  // out
	var _cret **C.char // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	*(**FileFilter)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "FileFilter").InvokeMethod("get_attributes", args[:], nil)
	_cret = *(***C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(filter)

	var _utf8s []string // out

	{
		var i int
		var z *C.void
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
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	*(**FileFilter)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "FileFilter").InvokeMethod("get_name", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	if name != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	*(**FileFilter)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "FileFilter").InvokeMethod("set_name", args[:], nil)

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
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	*(**FileFilter)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "FileFilter").InvokeMethod("to_gvariant", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
