// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
type RecentFilterFunc func(filterInfo *RecentFilterInfo, ok bool)

//export gotk4_RecentFilterFunc
func _RecentFilterFunc(arg0 *C.GtkRecentFilterInfo, arg1 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var filterInfo *RecentFilterInfo // out

	filterInfo = (*RecentFilterInfo)(unsafe.Pointer(arg0))

	fn := v.(RecentFilterFunc)
	ok := fn(filterInfo)

	var cret C.gboolean // out

	if ok {
		cret = C.TRUE
	}

	return cret
}

// RecentFilter: a RecentFilter can be used to restrict the files being shown in
// a RecentChooser. Files can be filtered based on their name (with
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
	Buildable

	AddAgeRecentFilter(days int)

	AddApplicationRecentFilter(application string)

	AddGroupRecentFilter(group string)

	AddMIMETypeRecentFilter(mimeType string)

	AddPatternRecentFilter(pattern string)

	AddPixbufFormatsRecentFilter()

	FilterRecentFilter(filterInfo *RecentFilterInfo) bool

	GetName() string

	Needed() RecentFilterFlags

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

func NewRecentFilter() RecentFilter {
	var _cret *C.GtkRecentFilter // in

	_cret = C.gtk_recent_filter_new()

	var _recentFilter RecentFilter // out

	_recentFilter = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(RecentFilter)

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

func (f recentFilter) FilterRecentFilter(filterInfo *RecentFilterInfo) bool {
	var _arg0 *C.GtkRecentFilter     // out
	var _arg1 *C.GtkRecentFilterInfo // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.GtkRecentFilterInfo)(unsafe.Pointer(filterInfo.Native()))

	_cret = C.gtk_recent_filter_filter(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (f recentFilter) GetName() string {
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

func (b recentFilter) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b recentFilter) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b recentFilter) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b recentFilter) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b recentFilter) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b recentFilter) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b recentFilter) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b recentFilter) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b recentFilter) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b recentFilter) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

// RecentFilterInfo: a GtkRecentFilterInfo struct is used to pass information
// about the tested file to gtk_recent_filter_filter().
type RecentFilterInfo C.GtkRecentFilterInfo

// WrapRecentFilterInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRecentFilterInfo(ptr unsafe.Pointer) *RecentFilterInfo {
	return (*RecentFilterInfo)(ptr)
}

// Native returns the underlying C source pointer.
func (r *RecentFilterInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(r)
}
