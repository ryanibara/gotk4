// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
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
var (
	GTypeFileFilterFlags = coreglib.Type(C.gtk_file_filter_flags_get_type())
	GTypeFileFilter      = coreglib.Type(C.gtk_file_filter_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeFileFilterFlags, F: marshalFileFilterFlags},
		{T: GTypeFileFilter, F: marshalFileFilter},
	})
}

// FileFilterFlags: these flags indicate what parts of a FileFilterInfo struct
// are filled or need to be filled.
type FileFilterFlags C.guint

const (
	// FileFilterFilename: filename of the file being tested.
	FileFilterFilename FileFilterFlags = 0b1
	// FileFilterURI: URI for the file being tested.
	FileFilterURI FileFilterFlags = 0b10
	// FileFilterDisplayName: string that will be used to display the file in
	// the file chooser.
	FileFilterDisplayName FileFilterFlags = 0b100
	// FileFilterMIMEType: mime type of the file.
	FileFilterMIMEType FileFilterFlags = 0b1000
)

func marshalFileFilterFlags(p uintptr) (interface{}, error) {
	return FileFilterFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for FileFilterFlags.
func (f FileFilterFlags) String() string {
	if f == 0 {
		return "FileFilterFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(73)

	for f != 0 {
		next := f & (f - 1)
		bit := f - next

		switch bit {
		case FileFilterFilename:
			builder.WriteString("Filename|")
		case FileFilterURI:
			builder.WriteString("URI|")
		case FileFilterDisplayName:
			builder.WriteString("DisplayName|")
		case FileFilterMIMEType:
			builder.WriteString("MIMEType|")
		default:
			builder.WriteString(fmt.Sprintf("FileFilterFlags(0b%b)|", bit))
		}

		f = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if f contains other.
func (f FileFilterFlags) Has(other FileFilterFlags) bool {
	return (f & other) == other
}

// FileFilterFunc: type of function that is used with custom filters, see
// gtk_file_filter_add_custom().
type FileFilterFunc func(filterInfo *FileFilterInfo) (ok bool)

//export _gotk4_gtk3_FileFilterFunc
func _gotk4_gtk3_FileFilterFunc(arg1 *C.void, arg2 C.gpointer) (cret C.gboolean) {
	var fn FileFilterFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(FileFilterFunc)
	}

	var _filterInfo *FileFilterInfo // out

	_filterInfo = (*FileFilterInfo)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := fn(_filterInfo)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// FileFilter can be used to restrict the files being shown in a FileChooser.
// Files can be filtered based on their name (with
// gtk_file_filter_add_pattern()), on their mime type (with
// gtk_file_filter_add_mime_type()), or by a custom filter function (with
// gtk_file_filter_add_custom()).
//
// Filtering by mime types handles aliasing and subclassing of mime types; e.g.
// a filter for text/plain also matches a file with mime type application/rtf,
// since application/rtf is a subclass of text/plain. Note that FileFilter
// allows wildcards for the subtype of a mime type, so you can e.g. filter for
// image/\*.
//
// Normally, filters are used by adding them to a FileChooser, see
// gtk_file_chooser_add_filter(), but it is also possible to manually use a
// filter on a file with gtk_file_filter_filter().
//
//
// GtkFileFilter as GtkBuildable
//
// The GtkFileFilter implementation of the GtkBuildable interface supports
// adding rules using the <mime-types>, <patterns> and <applications> elements
// and listing the rules within. Specifying a <mime-type> or <pattern> has the
// same effect as as calling gtk_file_filter_add_mime_type() or
// gtk_file_filter_add_pattern().
//
// An example of a UI definition fragment specifying GtkFileFilter rules:
//
//    <object class="GtkFileFilter">
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
	coreglib.InitiallyUnowned

	*coreglib.Object
	Buildable
}

var (
	_ coreglib.Objector = (*FileFilter)(nil)
)

func wrapFileFilter(obj *coreglib.Object) *FileFilter {
	return &FileFilter{
		InitiallyUnowned: coreglib.InitiallyUnowned{
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

// NewFileFilter creates a new FileFilter with no rules added to it. Such a
// filter doesn’t accept any files, so is not particularly useful until you add
// rules with gtk_file_filter_add_mime_type(), gtk_file_filter_add_pattern(), or
// gtk_file_filter_add_custom(). To create a filter that accepts any file, use:
//
//    GtkFileFilter *filter = gtk_file_filter_new ();
//    gtk_file_filter_add_pattern (filter, "*");.
//
// The function returns the following values:
//
//    - fileFilter: new FileFilter.
//
func NewFileFilter() *FileFilter {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "FileFilter").InvokeMethod("new_FileFilter", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _fileFilter *FileFilter // out

	_fileFilter = wrapFileFilter(coreglib.Take(unsafe.Pointer(_cret)))

	return _fileFilter
}

// NewFileFilterFromGVariant: deserialize a file filter from an a{sv} variant in
// the format produced by gtk_file_filter_to_gvariant().
//
// The function takes the following parameters:
//
//    - variant: a{sv} #GVariant.
//
// The function returns the following values:
//
//    - fileFilter: new FileFilter object.
//
func NewFileFilterFromGVariant(variant *glib.Variant) *FileFilter {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(variant)))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "FileFilter").InvokeMethod("new_FileFilter_from_gvariant", _args[:], nil)
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
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg1))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "FileFilter").InvokeMethod("add_mime_type", _args[:], nil)

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
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(pattern)))
	defer C.free(unsafe.Pointer(_arg1))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "FileFilter").InvokeMethod("add_pattern", _args[:], nil)

	runtime.KeepAlive(filter)
	runtime.KeepAlive(pattern)
}

// AddPixbufFormats adds a rule allowing image files in the formats supported by
// GdkPixbuf.
func (filter *FileFilter) AddPixbufFormats() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "FileFilter").InvokeMethod("add_pixbuf_formats", _args[:], nil)

	runtime.KeepAlive(filter)
}

// Filter tests whether a file should be displayed according to filter. The
// FileFilterInfo filter_info should include the fields returned from
// gtk_file_filter_get_needed().
//
// This function will not typically be used by applications; it is intended
// principally for use in the implementation of FileChooser.
//
// The function takes the following parameters:
//
//    - filterInfo containing information about a file.
//
// The function returns the following values:
//
//    - ok: TRUE if the file should be displayed.
//
func (filter *FileFilter) Filter(filterInfo *FileFilterInfo) bool {
	var _args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(filterInfo)))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "FileFilter").InvokeMethod("filter", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(filter)
	runtime.KeepAlive(filterInfo)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Name gets the human-readable name for the filter. See
// gtk_file_filter_set_name().
//
// The function returns the following values:
//
//    - utf8 (optional): human-readable name of the filter, or NULL. This value
//      is owned by GTK+ and must not be modified or freed.
//
func (filter *FileFilter) Name() string {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "FileFilter").InvokeMethod("get_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(filter)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SetName sets the human-readable name of the filter; this is the string that
// will be displayed in the file selector user interface if there is a
// selectable list of filters.
//
// The function takes the following parameters:
//
//    - name (optional) for the filter, or NULL to remove any existing name.
//
func (filter *FileFilter) SetName(name string) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	if name != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "FileFilter").InvokeMethod("set_name", _args[:], nil)

	runtime.KeepAlive(filter)
	runtime.KeepAlive(name)
}

// ToGVariant: serialize a file filter to an a{sv} variant.
//
// The function returns the following values:
//
//    - variant: new, floating, #GVariant.
//
func (filter *FileFilter) ToGVariant() *glib.Variant {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "FileFilter").InvokeMethod("to_gvariant", _args[:], nil)
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

// FileFilterInfo is used to pass information about the tested file to
// gtk_file_filter_filter().
//
// An instance of this type is always passed by reference.
type FileFilterInfo struct {
	*fileFilterInfo
}

// fileFilterInfo is the struct that's finalized.
type fileFilterInfo struct {
	native *C.GtkFileFilterInfo
}
