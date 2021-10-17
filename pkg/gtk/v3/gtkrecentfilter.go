// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void callbackDelete(gpointer);
// gboolean _gotk4_gtk3_RecentFilterFunc(GtkRecentFilterInfo*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_recent_filter_flags_get_type()), F: marshalRecentFilterFlags},
		{T: externglib.Type(C.gtk_recent_filter_get_type()), F: marshalRecentFilterer},
	})
}

// RecentFilterFlags: these flags indicate what parts of a RecentFilterInfo
// struct are filled or need to be filled.
type RecentFilterFlags int

const (
	// RecentFilterURI: URI of the file being tested.
	RecentFilterURI RecentFilterFlags = 0b1
	// RecentFilterDisplayName: string that will be used to display the file in
	// the recent chooser.
	RecentFilterDisplayName RecentFilterFlags = 0b10
	// RecentFilterMIMEType: mime type of the file.
	RecentFilterMIMEType RecentFilterFlags = 0b100
	// RecentFilterApplication: list of applications that have registered the
	// file.
	RecentFilterApplication RecentFilterFlags = 0b1000
	// RecentFilterGroup groups to which the file belongs to.
	RecentFilterGroup RecentFilterFlags = 0b10000
	// RecentFilterAge: number of days elapsed since the file has been
	// registered.
	RecentFilterAge RecentFilterFlags = 0b100000
)

func marshalRecentFilterFlags(p uintptr) (interface{}, error) {
	return RecentFilterFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for RecentFilterFlags.
func (r RecentFilterFlags) String() string {
	if r == 0 {
		return "RecentFilterFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(118)

	for r != 0 {
		next := r & (r - 1)
		bit := r - next

		switch bit {
		case RecentFilterURI:
			builder.WriteString("URI|")
		case RecentFilterDisplayName:
			builder.WriteString("DisplayName|")
		case RecentFilterMIMEType:
			builder.WriteString("MIMEType|")
		case RecentFilterApplication:
			builder.WriteString("Application|")
		case RecentFilterGroup:
			builder.WriteString("Group|")
		case RecentFilterAge:
			builder.WriteString("Age|")
		default:
			builder.WriteString(fmt.Sprintf("RecentFilterFlags(0b%b)|", bit))
		}

		r = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if r contains other.
func (r RecentFilterFlags) Has(other RecentFilterFlags) bool {
	return (r & other) == other
}

// RecentFilterFunc: type of function that is used with custom filters, see
// gtk_recent_filter_add_custom().
type RecentFilterFunc func(filterInfo *RecentFilterInfo) (ok bool)

//export _gotk4_gtk3_RecentFilterFunc
func _gotk4_gtk3_RecentFilterFunc(arg0 *C.GtkRecentFilterInfo, arg1 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var filterInfo *RecentFilterInfo // out

	filterInfo = (*RecentFilterInfo)(gextras.NewStructNative(unsafe.Pointer(arg0)))

	fn := v.(RecentFilterFunc)
	ok := fn(filterInfo)

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
//    </object>.
type RecentFilter struct {
	externglib.InitiallyUnowned

	Buildable
	*externglib.Object
}

func wrapRecentFilter(obj *externglib.Object) *RecentFilter {
	return &RecentFilter{
		InitiallyUnowned: externglib.InitiallyUnowned{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalRecentFilterer(p uintptr) (interface{}, error) {
	return wrapRecentFilter(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewRecentFilter creates a new RecentFilter with no rules added to it. Such
// filter does not accept any recently used resources, so is not particularly
// useful until you add rules with gtk_recent_filter_add_pattern(),
// gtk_recent_filter_add_mime_type(), gtk_recent_filter_add_application(),
// gtk_recent_filter_add_age(). To create a filter that accepts any recently
// used resource, use:
//
//    GtkRecentFilter *filter = gtk_recent_filter_new ();
//    gtk_recent_filter_add_pattern (filter, "*");.
func NewRecentFilter() *RecentFilter {
	var _cret *C.GtkRecentFilter // in

	_cret = C.gtk_recent_filter_new()

	var _recentFilter *RecentFilter // out

	_recentFilter = wrapRecentFilter(externglib.Take(unsafe.Pointer(_cret)))

	return _recentFilter
}

// AddAge adds a rule that allows resources based on their age - that is, the
// number of days elapsed since they were last modified.
//
// The function takes the following parameters:
//
//    - days: number of days.
//
func (filter *RecentFilter) AddAge(days int) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 C.gint             // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = C.gint(days)

	C.gtk_recent_filter_add_age(_arg0, _arg1)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(days)
}

// AddApplication adds a rule that allows resources based on the name of the
// application that has registered them.
//
// The function takes the following parameters:
//
//    - application name.
//
func (filter *RecentFilter) AddApplication(application string) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(application)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_filter_add_application(_arg0, _arg1)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(application)
}

// AddCustom adds a rule to a filter that allows resources based on a custom
// callback function. The bitfield needed which is passed in provides
// information about what sorts of information that the filter function needs;
// this allows GTK+ to avoid retrieving expensive information when it isn’t
// needed by the filter.
//
// The function takes the following parameters:
//
//    - needed: bitfield of flags indicating the information that the custom
//    filter function needs.
//    - fn: callback function; if the function returns TRUE, then the file will
//    be displayed.
//
func (filter *RecentFilter) AddCustom(needed RecentFilterFlags, fn RecentFilterFunc) {
	var _arg0 *C.GtkRecentFilter     // out
	var _arg1 C.GtkRecentFilterFlags // out
	var _arg2 C.GtkRecentFilterFunc  // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = C.GtkRecentFilterFlags(needed)
	_arg2 = (*[0]byte)(C._gotk4_gtk3_RecentFilterFunc)
	_arg3 = C.gpointer(gbox.Assign(fn))
	_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_recent_filter_add_custom(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(needed)
	runtime.KeepAlive(fn)
}

// AddGroup adds a rule that allows resources based on the name of the group to
// which they belong.
//
// The function takes the following parameters:
//
//    - group name.
//
func (filter *RecentFilter) AddGroup(group string) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(group)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_filter_add_group(_arg0, _arg1)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(group)
}

// AddMIMEType adds a rule that allows resources based on their registered MIME
// type.
//
// The function takes the following parameters:
//
//    - mimeType: MIME type.
//
func (filter *RecentFilter) AddMIMEType(mimeType string) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_filter_add_mime_type(_arg0, _arg1)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(mimeType)
}

// AddPattern adds a rule that allows resources based on a pattern matching
// their display name.
//
// The function takes the following parameters:
//
//    - pattern: file pattern.
//
func (filter *RecentFilter) AddPattern(pattern string) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(pattern)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_filter_add_pattern(_arg0, _arg1)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(pattern)
}

// AddPixbufFormats adds a rule allowing image files in the formats supported by
// GdkPixbuf.
func (filter *RecentFilter) AddPixbufFormats() {
	var _arg0 *C.GtkRecentFilter // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_recent_filter_add_pixbuf_formats(_arg0)
	runtime.KeepAlive(filter)
}

// Filter tests whether a file should be displayed according to filter. The
// RecentFilterInfo filter_info should include the fields returned from
// gtk_recent_filter_get_needed(), and must set the RecentFilterInfo.contains
// field of filter_info to indicate which fields have been set.
//
// This function will not typically be used by applications; it is intended
// principally for use in the implementation of RecentChooser.
//
// The function takes the following parameters:
//
//    - filterInfo containing information about a recently used resource.
//
func (filter *RecentFilter) Filter(filterInfo *RecentFilterInfo) bool {
	var _arg0 *C.GtkRecentFilter     // out
	var _arg1 *C.GtkRecentFilterInfo // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = (*C.GtkRecentFilterInfo)(gextras.StructNative(unsafe.Pointer(filterInfo)))

	_cret = C.gtk_recent_filter_filter(_arg0, _arg1)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(filterInfo)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Name gets the human-readable name for the filter. See
// gtk_recent_filter_set_name().
func (filter *RecentFilter) Name() string {
	var _arg0 *C.GtkRecentFilter // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))

	_cret = C.gtk_recent_filter_get_name(_arg0)
	runtime.KeepAlive(filter)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Needed gets the fields that need to be filled in for the RecentFilterInfo
// passed to gtk_recent_filter_filter()
//
// This function will not typically be used by applications; it is intended
// principally for use in the implementation of RecentChooser.
func (filter *RecentFilter) Needed() RecentFilterFlags {
	var _arg0 *C.GtkRecentFilter     // out
	var _cret C.GtkRecentFilterFlags // in

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))

	_cret = C.gtk_recent_filter_get_needed(_arg0)
	runtime.KeepAlive(filter)

	var _recentFilterFlags RecentFilterFlags // out

	_recentFilterFlags = RecentFilterFlags(_cret)

	return _recentFilterFlags
}

// SetName sets the human-readable name of the filter; this is the string that
// will be displayed in the recently used resources selector user interface if
// there is a selectable list of filters.
//
// The function takes the following parameters:
//
//    - name: then human readable name of filter.
//
func (filter *RecentFilter) SetName(name string) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_filter_set_name(_arg0, _arg1)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(name)
}

// RecentFilterInfo struct is used to pass information about the tested file to
// gtk_recent_filter_filter().
//
// An instance of this type is always passed by reference.
type RecentFilterInfo struct {
	*recentFilterInfo
}

// recentFilterInfo is the struct that's finalized.
type recentFilterInfo struct {
	native *C.GtkRecentFilterInfo
}

// Contains to indicate which fields are set.
func (r *RecentFilterInfo) Contains() RecentFilterFlags {
	var v RecentFilterFlags // out
	v = RecentFilterFlags(r.native.contains)
	return v
}

// URI of the file being tested.
func (r *RecentFilterInfo) URI() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(r.native.uri)))
	return v
}

// DisplayName: string that will be used to display the file in the recent
// chooser.
func (r *RecentFilterInfo) DisplayName() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(r.native.display_name)))
	return v
}

// MIMEType: MIME type of the file.
func (r *RecentFilterInfo) MIMEType() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(r.native.mime_type)))
	return v
}

// Applications: list of applications that have registered the file.
func (r *RecentFilterInfo) Applications() []string {
	var v []string // out
	{
		var i int
		var z *C.gchar
		for p := r.native.applications; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(r.native.applications, i)
		v = make([]string, i)
		for i := range src {
			v[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}
	return v
}

// Groups groups to which the file belongs to.
func (r *RecentFilterInfo) Groups() []string {
	var v []string // out
	{
		var i int
		var z *C.gchar
		for p := r.native.groups; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(r.native.groups, i)
		v = make([]string, i)
		for i := range src {
			v[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}
	return v
}

// Age: number of days elapsed since the file has been registered.
func (r *RecentFilterInfo) Age() int {
	var v int // out
	v = int(r.native.age)
	return v
}
