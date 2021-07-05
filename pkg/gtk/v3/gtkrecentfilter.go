// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_recent_filter_flags_get_type()), F: marshalRecentFilterFlags},
		{T: externglib.Type(C.gtk_recent_filter_get_type()), F: marshalRecentFilter},
	})
}

// RecentFilterFlags: these flags indicate what parts of a RecentFilterInfo
// struct are filled or need to be filled.
type RecentFilterFlags int

const (
	// RecentFilterFlagsURI: the URI of the file being tested
	RecentFilterFlagsURI RecentFilterFlags = 0b1
	// RecentFilterFlagsDisplayName: the string that will be used to display the
	// file in the recent chooser
	RecentFilterFlagsDisplayName RecentFilterFlags = 0b10
	// RecentFilterFlagsMIMEType: the mime type of the file
	RecentFilterFlagsMIMEType RecentFilterFlags = 0b100
	// RecentFilterFlagsApplication: the list of applications that have
	// registered the file
	RecentFilterFlagsApplication RecentFilterFlags = 0b1000
	// RecentFilterFlagsGroup: the groups to which the file belongs to
	RecentFilterFlagsGroup RecentFilterFlags = 0b10000
	// RecentFilterFlagsAge: the number of days elapsed since the file has been
	// registered
	RecentFilterFlagsAge RecentFilterFlags = 0b100000
)

func marshalRecentFilterFlags(p uintptr) (interface{}, error) {
	return RecentFilterFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// RecentFilterFunc: the type of function that is used with custom filters, see
// gtk_recent_filter_add_custom().
type RecentFilterFunc func(filterInfo RecentFilterInfo) (ok bool)

//export gotk4_RecentFilterFunc
func gotk4_RecentFilterFunc(arg0 *C.GtkRecentFilterInfo, arg1 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var filterInfo RecentFilterInfo // out

	filterInfo = (RecentFilterInfo)(unsafe.Pointer(arg0))

	fn := v.(RecentFilterFunc)
	ok := fn(filterInfo)

	var cret C.gboolean // out

	if ok {
		cret = C.TRUE
	}

	return cret
}

// RecentFilter can be used to restrict the files being shown in a
// RecentChooser. Files can be filtered based on their name (with
// gtk_recent_filter_add_pattern()), on their mime type (with
// gtk_file_filter_add_mime_type()), on the application that has registered them
// (with gtk_recent_filter_add_application()), or by a custom filter function
// (with gtk_recent_filter_add_custom()).
//
// Filtering by mime type handles aliasing and subclassing of mime types; e.g. a
// filter for text/plain also matches a file with mime type application/rtf,
// since application/rtf is a subclass of text/plain. Note that RecentFilter
// allows wildcards for the subtype of a mime type, so you can e.g. filter for
// image/\*.
//
// Normally, filters are used by adding them to a RecentChooser, see
// gtk_recent_chooser_add_filter(), but it is also possible to manually use a
// filter on a file with gtk_recent_filter_filter().
//
// Recently used files are supported since GTK+ 2.10.
//
//
// GtkRecentFilter as GtkBuildable
//
// The GtkRecentFilter implementation of the GtkBuildable interface supports
// adding rules using the <mime-types>, <patterns> and <applications> elements
// and listing the rules within. Specifying a <mime-type>, <pattern> or
// <application> has the same effect as calling
// gtk_recent_filter_add_mime_type(), gtk_recent_filter_add_pattern() or
// gtk_recent_filter_add_application().
//
// An example of a UI definition fragment specifying GtkRecentFilter rules:
//
//    <object class="GtkRecentFilter">
//      <mime-types>
//        <mime-type>text/plain</mime-type>
//        <mime-type>image/png</mime-type>
//      </mime-types>
//      <patterns>
//        <pattern>*.txt</pattern>
//        <pattern>*.png</pattern>
//      </patterns>
//      <applications>
//        <application>gimp</application>
//        <application>gedit</application>
//        <application>glade</application>
//      </applications>
//    </object>
type RecentFilter interface {
	gextras.Objector

	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable

	// AddAgeRecentFilter adds a rule that allows resources based on their age -
	// that is, the number of days elapsed since they were last modified.
	AddAgeRecentFilter(days int)
	// AddApplicationRecentFilter adds a rule that allows resources based on the
	// name of the application that has registered them.
	AddApplicationRecentFilter(application string)
	// AddGroupRecentFilter adds a rule that allows resources based on the name
	// of the group to which they belong
	AddGroupRecentFilter(group string)
	// AddMIMETypeRecentFilter adds a rule that allows resources based on their
	// registered MIME type.
	AddMIMETypeRecentFilter(mimeType string)
	// AddPatternRecentFilter adds a rule that allows resources based on a
	// pattern matching their display name.
	AddPatternRecentFilter(pattern string)
	// AddPixbufFormatsRecentFilter adds a rule allowing image files in the
	// formats supported by GdkPixbuf.
	AddPixbufFormatsRecentFilter()
	// FilterRecentFilter tests whether a file should be displayed according to
	// @filter. The RecentFilterInfo @filter_info should include the fields
	// returned from gtk_recent_filter_get_needed(), and must set the
	// RecentFilterInfo.contains field of @filter_info to indicate which fields
	// have been set.
	//
	// This function will not typically be used by applications; it is intended
	// principally for use in the implementation of RecentChooser.
	FilterRecentFilter(filterInfo RecentFilterInfo) bool
	// Name gets the human-readable name for the filter. See
	// gtk_recent_filter_set_name().
	Name() string
	// Needed gets the fields that need to be filled in for the RecentFilterInfo
	// passed to gtk_recent_filter_filter()
	//
	// This function will not typically be used by applications; it is intended
	// principally for use in the implementation of RecentChooser.
	Needed() RecentFilterFlags
	// SetNameRecentFilter sets the human-readable name of the filter; this is
	// the string that will be displayed in the recently used resources selector
	// user interface if there is a selectable list of filters.
	SetNameRecentFilter(name string)
}

// recentFilter implements the RecentFilter class.
type recentFilter struct {
	gextras.Objector
}

// WrapRecentFilter wraps a GObject to the right type. It is
// primarily used internally.
func WrapRecentFilter(obj *externglib.Object) RecentFilter {
	return recentFilter{
		Objector: obj,
	}
}

func marshalRecentFilter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRecentFilter(obj), nil
}

// NewRecentFilter creates a new RecentFilter with no rules added to it. Such
// filter does not accept any recently used resources, so is not particularly
// useful until you add rules with gtk_recent_filter_add_pattern(),
// gtk_recent_filter_add_mime_type(), gtk_recent_filter_add_application(),
// gtk_recent_filter_add_age(). To create a filter that accepts any recently
// used resource, use:
//
//    GtkRecentFilter *filter = gtk_recent_filter_new ();
//    gtk_recent_filter_add_pattern (filter, "*");
func NewRecentFilter() RecentFilter {
	var _cret *C.GtkRecentFilter // in

	_cret = C.gtk_recent_filter_new()

	var _recentFilter RecentFilter // out

	_recentFilter = WrapRecentFilter(externglib.Take(unsafe.Pointer(_cret)))

	return _recentFilter
}

func (f recentFilter) AddAgeRecentFilter(days int) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 C.gint             // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))
	_arg1 = C.gint(days)

	C.gtk_recent_filter_add_age(_arg0, _arg1)
}

func (f recentFilter) AddApplicationRecentFilter(application string) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(application))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_filter_add_application(_arg0, _arg1)
}

func (f recentFilter) AddGroupRecentFilter(group string) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(group))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_filter_add_group(_arg0, _arg1)
}

func (f recentFilter) AddMIMETypeRecentFilter(mimeType string) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_filter_add_mime_type(_arg0, _arg1)
}

func (f recentFilter) AddPatternRecentFilter(pattern string) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(pattern))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_filter_add_pattern(_arg0, _arg1)
}

func (f recentFilter) AddPixbufFormatsRecentFilter() {
	var _arg0 *C.GtkRecentFilter // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))

	C.gtk_recent_filter_add_pixbuf_formats(_arg0)
}

func (f recentFilter) FilterRecentFilter(filterInfo RecentFilterInfo) bool {
	var _arg0 *C.GtkRecentFilter     // out
	var _arg1 *C.GtkRecentFilterInfo // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.GtkRecentFilterInfo)(unsafe.Pointer(filterInfo))

	_cret = C.gtk_recent_filter_filter(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (f recentFilter) Name() string {
	var _arg0 *C.GtkRecentFilter // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_recent_filter_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (f recentFilter) Needed() RecentFilterFlags {
	var _arg0 *C.GtkRecentFilter     // out
	var _cret C.GtkRecentFilterFlags // in

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_recent_filter_get_needed(_arg0)

	var _recentFilterFlags RecentFilterFlags // out

	_recentFilterFlags = RecentFilterFlags(_cret)

	return _recentFilterFlags
}

func (f recentFilter) SetNameRecentFilter(name string) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_filter_set_name(_arg0, _arg1)
}

func (r recentFilter) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(r))
}

// RecentFilterInfo struct is used to pass information about the tested file to
// gtk_recent_filter_filter().
type RecentFilterInfo struct {
	native C.GtkRecentFilterInfo
}

// WrapRecentFilterInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRecentFilterInfo(ptr unsafe.Pointer) *RecentFilterInfo {
	return (*RecentFilterInfo)(ptr)
}

// Native returns the underlying C source pointer.
func (r *RecentFilterInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}
