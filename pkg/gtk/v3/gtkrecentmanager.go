// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_recent_manager_get_type()), F: marshalRecentManager},
		{T: externglib.Type(C.gtk_recent_info_get_type()), F: marshalRecentInfo},
	})
}

// RecentManager provides a facility for adding, removing and looking up
// recently used files. Each recently used file is identified by its URI, and
// has meta-data associated to it, like the names and command lines of the
// applications that have registered it, the number of time each application has
// registered the same file, the mime type of the file and whether the file
// should be displayed only by the applications that have registered it.
//
// The recently used files list is per user.
//
// The RecentManager acts like a database of all the recently used files. You
// can create new RecentManager objects, but it is more efficient to use the
// default manager created by GTK+.
//
// Adding a new recently used file is as simple as:
//
//    GtkRecentManager *manager;
//    GtkRecentInfo *info;
//    GError *error = NULL;
//
//    manager = gtk_recent_manager_get_default ();
//    info = gtk_recent_manager_lookup_item (manager, file_uri, &error);
//    if (error)
//      {
//        g_warning ("Could not find the file: s", error->message);
//        g_error_free (error);
//      }
//    else
//     {
//       // Use the info object
//       gtk_recent_info_unref (info);
//     }
//
// In order to retrieve the list of recently used files, you can use
// gtk_recent_manager_get_items(), which returns a list of RecentInfo-structs.
//
// A RecentManager is the model used to populate the contents of one, or more
// RecentChooser implementations.
//
// Note that the maximum age of the recently used files list is controllable
// through the Settings:gtk-recent-files-max-age property.
//
// Recently used files are supported since GTK+ 2.10.
type RecentManager interface {
	gextras.Objector

	// AddFull adds a new resource, pointed by @uri, into the recently used
	// resources list, using the metadata specified inside the RecentData-struct
	// passed in @recent_data.
	//
	// The passed URI will be used to identify this resource inside the list.
	//
	// In order to register the new recently used resource, metadata about the
	// resource must be passed as well as the URI; the metadata is stored in a
	// RecentData-struct, which must contain the MIME type of the resource
	// pointed by the URI; the name of the application that is registering the
	// item, and a command line to be used when launching the item.
	//
	// Optionally, a RecentData-struct might contain a UTF-8 string to be used
	// when viewing the item instead of the last component of the URI; a short
	// description of the item; whether the item should be considered private -
	// that is, should be displayed only by the applications that have
	// registered it.
	AddFull(uri string, recentData *RecentData) bool
	// AddItem adds a new resource, pointed by @uri, into the recently used
	// resources list.
	//
	// This function automatically retrieves some of the needed metadata and
	// setting other metadata to common default values; it then feeds the data
	// to gtk_recent_manager_add_full().
	//
	// See gtk_recent_manager_add_full() if you want to explicitly define the
	// metadata for the resource pointed by @uri.
	AddItem(uri string) bool
	// Items gets the list of recently used resources.
	Items() *glib.List
	// HasItem checks whether there is a recently used resource registered with
	// @uri inside the recent manager.
	HasItem(uri string) bool
	// LookupItem searches for a URI inside the recently used resources list,
	// and returns a RecentInfo-struct containing informations about the
	// resource like its MIME type, or its display name.
	LookupItem(uri string) (recentInfo *RecentInfo, err error)
	// MoveItem changes the location of a recently used resource from @uri to
	// @new_uri.
	//
	// Please note that this function will not affect the resource pointed by
	// the URIs, but only the URI used in the recently used resources list.
	MoveItem(uri string, newURI string) error
	// PurgeItems purges every item from the recently used resources list.
	PurgeItems() (gint int, err error)
	// RemoveItem removes a resource pointed by @uri from the recently used
	// resources list handled by a recent manager.
	RemoveItem(uri string) error
}

// recentManager implements the RecentManager interface.
type recentManager struct {
	gextras.Objector
}

var _ RecentManager = (*recentManager)(nil)

// WrapRecentManager wraps a GObject to the right type. It is
// primarily used internally.
func WrapRecentManager(obj *externglib.Object) RecentManager {
	return RecentManager{
		Objector: obj,
	}
}

func marshalRecentManager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRecentManager(obj), nil
}

// NewRecentManager constructs a class RecentManager.
func NewRecentManager() RecentManager {
	var cret C.GtkRecentManager
	var ret1 RecentManager

	cret = C.gtk_recent_manager_new()

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(RecentManager)

	return ret1
}

// AddFull adds a new resource, pointed by @uri, into the recently used
// resources list, using the metadata specified inside the RecentData-struct
// passed in @recent_data.
//
// The passed URI will be used to identify this resource inside the list.
//
// In order to register the new recently used resource, metadata about the
// resource must be passed as well as the URI; the metadata is stored in a
// RecentData-struct, which must contain the MIME type of the resource
// pointed by the URI; the name of the application that is registering the
// item, and a command line to be used when launching the item.
//
// Optionally, a RecentData-struct might contain a UTF-8 string to be used
// when viewing the item instead of the last component of the URI; a short
// description of the item; whether the item should be considered private -
// that is, should be displayed only by the applications that have
// registered it.
func (m recentManager) AddFull(uri string, recentData *RecentData) bool {
	var arg0 *C.GtkRecentManager
	var arg1 *C.gchar
	var arg2 *C.GtkRecentData

	arg0 = (*C.GtkRecentManager)(unsafe.Pointer(m.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GtkRecentData)(unsafe.Pointer(recentData.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_recent_manager_add_full(arg0, uri, recentData)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// AddItem adds a new resource, pointed by @uri, into the recently used
// resources list.
//
// This function automatically retrieves some of the needed metadata and
// setting other metadata to common default values; it then feeds the data
// to gtk_recent_manager_add_full().
//
// See gtk_recent_manager_add_full() if you want to explicitly define the
// metadata for the resource pointed by @uri.
func (m recentManager) AddItem(uri string) bool {
	var arg0 *C.GtkRecentManager
	var arg1 *C.gchar

	arg0 = (*C.GtkRecentManager)(unsafe.Pointer(m.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_recent_manager_add_item(arg0, uri)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Items gets the list of recently used resources.
func (m recentManager) Items() *glib.List {
	var arg0 *C.GtkRecentManager

	arg0 = (*C.GtkRecentManager)(unsafe.Pointer(m.Native()))

	var cret *C.GList
	var ret1 *glib.List

	cret = C.gtk_recent_manager_get_items(arg0)

	ret1 = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// HasItem checks whether there is a recently used resource registered with
// @uri inside the recent manager.
func (m recentManager) HasItem(uri string) bool {
	var arg0 *C.GtkRecentManager
	var arg1 *C.gchar

	arg0 = (*C.GtkRecentManager)(unsafe.Pointer(m.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_recent_manager_has_item(arg0, uri)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// LookupItem searches for a URI inside the recently used resources list,
// and returns a RecentInfo-struct containing informations about the
// resource like its MIME type, or its display name.
func (m recentManager) LookupItem(uri string) (recentInfo *RecentInfo, err error) {
	var arg0 *C.GtkRecentManager
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GtkRecentManager)(unsafe.Pointer(m.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GtkRecentInfo
	var ret1 *RecentInfo
	var goerr error

	cret = C.gtk_recent_manager_lookup_item(arg0, uri, &errout)

	ret1 = WrapRecentInfo(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *RecentInfo) {
		C.free(unsafe.Pointer(v.Native()))
	})
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return ret1, goerr
}

// MoveItem changes the location of a recently used resource from @uri to
// @new_uri.
//
// Please note that this function will not affect the resource pointed by
// the URIs, but only the URI used in the recently used resources list.
func (m recentManager) MoveItem(uri string, newURI string) error {
	var arg0 *C.GtkRecentManager
	var arg1 *C.gchar
	var arg2 *C.gchar
	var errout *C.GError

	arg0 = (*C.GtkRecentManager)(unsafe.Pointer(m.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(newURI))
	defer C.free(unsafe.Pointer(arg2))

	var goerr error

	C.gtk_recent_manager_move_item(arg0, uri, newURI, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// PurgeItems purges every item from the recently used resources list.
func (m recentManager) PurgeItems() (gint int, err error) {
	var arg0 *C.GtkRecentManager
	var errout *C.GError

	arg0 = (*C.GtkRecentManager)(unsafe.Pointer(m.Native()))

	var cret C.gint
	var ret1 int
	var goerr error

	cret = C.gtk_recent_manager_purge_items(arg0, &errout)

	ret1 = C.gint(cret)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return ret1, goerr
}

// RemoveItem removes a resource pointed by @uri from the recently used
// resources list handled by a recent manager.
func (m recentManager) RemoveItem(uri string) error {
	var arg0 *C.GtkRecentManager
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GtkRecentManager)(unsafe.Pointer(m.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var goerr error

	C.gtk_recent_manager_remove_item(arg0, uri, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// RecentData: meta-data to be passed to gtk_recent_manager_add_full() when
// registering a recently used resource.
type RecentData struct {
	native C.GtkRecentData
}

// WrapRecentData wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRecentData(ptr unsafe.Pointer) *RecentData {
	if ptr == nil {
		return nil
	}

	return (*RecentData)(ptr)
}

func marshalRecentData(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRecentData(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RecentData) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// DisplayName gets the field inside the struct.
func (r *RecentData) DisplayName() string {
	v = C.GoString(r.native.display_name)
}

// Description gets the field inside the struct.
func (r *RecentData) Description() string {
	v = C.GoString(r.native.description)
}

// MIMEType gets the field inside the struct.
func (r *RecentData) MIMEType() string {
	v = C.GoString(r.native.mime_type)
}

// AppName gets the field inside the struct.
func (r *RecentData) AppName() string {
	v = C.GoString(r.native.app_name)
}

// AppExec gets the field inside the struct.
func (r *RecentData) AppExec() string {
	v = C.GoString(r.native.app_exec)
}

// Groups gets the field inside the struct.
func (r *RecentData) Groups() []string {
	{
		var length int
		for p := r.native.groups; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		v = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.gchar)(ptr.Add(unsafe.Pointer(r.native.groups), i))
			v[i] = C.GoString(src)
		}
	}
}

// IsPrivate gets the field inside the struct.
func (r *RecentData) IsPrivate() bool {
	v = C.bool(r.native.is_private) != C.false
}

// RecentInfo contains private data only, and should be accessed using the
// provided API.
//
// RecentInfo constains all the meta-data associated with an entry in the
// recently used files list.
type RecentInfo struct {
	native C.GtkRecentInfo
}

// WrapRecentInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRecentInfo(ptr unsafe.Pointer) *RecentInfo {
	if ptr == nil {
		return nil
	}

	return (*RecentInfo)(ptr)
}

func marshalRecentInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRecentInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RecentInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// CreateAppInfo creates a Info for the specified RecentInfo
func (i *RecentInfo) CreateAppInfo(appName string) (appInfo gio.AppInfo, err error) {
	var arg0 *C.GtkRecentInfo
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))
	arg1 = (*C.gchar)(C.CString(appName))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GAppInfo
	var ret1 gio.AppInfo
	var goerr error

	cret = C.gtk_recent_info_create_app_info(arg0, appName, &errout)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(gio.AppInfo)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return ret1, goerr
}

// Exists checks whether the resource pointed by @info still exists. At the
// moment this check is done only on resources pointing to local files.
func (i *RecentInfo) Exists() bool {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_recent_info_exists(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Added gets the timestamp (seconds from system’s Epoch) when the resource was
// added to the recently used resources list.
func (i *RecentInfo) Added() int32 {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret C.time_t
	var ret1 int32

	cret = C.gtk_recent_info_get_added(arg0)

	ret1 = C.time_t(cret)

	return ret1
}

// Age gets the number of days elapsed since the last update of the resource
// pointed by @info.
func (i *RecentInfo) Age() int {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret C.gint
	var ret1 int

	cret = C.gtk_recent_info_get_age(arg0)

	ret1 = C.gint(cret)

	return ret1
}

// ApplicationInfo gets the data regarding the application that has registered
// the resource pointed by @info.
//
// If the command line contains any escape characters defined inside the storage
// specification, they will be expanded.
func (i *RecentInfo) ApplicationInfo(appName string) (appExec string, count uint, time_ int32, ok bool) {
	var arg0 *C.GtkRecentInfo
	var arg1 *C.gchar

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))
	arg1 = (*C.gchar)(C.CString(appName))
	defer C.free(unsafe.Pointer(arg1))

	var arg2 *C.gchar
	var ret2 string
	var arg3 C.guint
	var ret3 uint
	var arg4 C.time_t
	var ret4 int32
	var cret C.gboolean
	var ret4 bool

	cret = C.gtk_recent_info_get_application_info(arg0, appName, &arg2, &arg3, &arg4)

	ret2 = C.GoString(arg2)
	ret3 = C.guint(arg3)
	ret4 = C.time_t(arg4)
	ret4 = C.bool(cret) != C.false

	return ret2, ret3, ret4, ret4
}

// Applications retrieves the list of applications that have registered this
// resource.
func (i *RecentInfo) Applications() (length uint, utf8s []string) {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret **C.gchar
	var arg1 *C.gsize
	var ret2 []string

	cret = C.gtk_recent_info_get_applications(arg0, &arg1)

	ret2 = make([]string, arg1)
	for i := 0; i < uintptr(arg1); i++ {
		src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
		ret2[i] = C.GoString(src)
		defer C.free(unsafe.Pointer(src))
	}

	return ret1, ret2
}

// Description gets the (short) description of the resource.
func (i *RecentInfo) Description() string {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_recent_info_get_description(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// DisplayName gets the name of the resource. If none has been defined, the
// basename of the resource is obtained.
func (i *RecentInfo) DisplayName() string {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_recent_info_get_display_name(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// GIcon retrieves the icon associated to the resource MIME type.
func (i *RecentInfo) GIcon() gio.Icon {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret *C.GIcon
	var ret1 gio.Icon

	cret = C.gtk_recent_info_get_gicon(arg0)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(gio.Icon)

	return ret1
}

// Groups returns all groups registered for the recently used item @info. The
// array of returned group names will be nil terminated, so length might
// optionally be nil.
func (i *RecentInfo) Groups() (length uint, utf8s []string) {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret **C.gchar
	var arg1 *C.gsize
	var ret2 []string

	cret = C.gtk_recent_info_get_groups(arg0, &arg1)

	ret2 = make([]string, arg1)
	for i := 0; i < uintptr(arg1); i++ {
		src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
		ret2[i] = C.GoString(src)
		defer C.free(unsafe.Pointer(src))
	}

	return ret1, ret2
}

// Icon retrieves the icon of size @size associated to the resource MIME type.
func (i *RecentInfo) Icon(size int) gdkpixbuf.Pixbuf {
	var arg0 *C.GtkRecentInfo
	var arg1 C.gint

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))
	arg1 = C.gint(size)

	var cret *C.GdkPixbuf
	var ret1 gdkpixbuf.Pixbuf

	cret = C.gtk_recent_info_get_icon(arg0, size)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(gdkpixbuf.Pixbuf)

	return ret1
}

// MIMEType gets the MIME type of the resource.
func (i *RecentInfo) MIMEType() string {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_recent_info_get_mime_type(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Modified gets the timestamp (seconds from system’s Epoch) when the meta-data
// for the resource was last modified.
func (i *RecentInfo) Modified() int32 {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret C.time_t
	var ret1 int32

	cret = C.gtk_recent_info_get_modified(arg0)

	ret1 = C.time_t(cret)

	return ret1
}

// PrivateHint gets the value of the “private” flag. Resources in the recently
// used list that have this flag set to true should only be displayed by the
// applications that have registered them.
func (i *RecentInfo) PrivateHint() bool {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_recent_info_get_private_hint(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// ShortName computes a valid UTF-8 string that can be used as the name of the
// item in a menu or list. For example, calling this function on an item that
// refers to “file:///foo/bar.txt” will yield “bar.txt”.
func (i *RecentInfo) ShortName() string {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_recent_info_get_short_name(arg0)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// URI gets the URI of the resource.
func (i *RecentInfo) URI() string {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_recent_info_get_uri(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// URIDisplay gets a displayable version of the resource’s URI. If the resource
// is local, it returns a local path; if the resource is not local, it returns
// the UTF-8 encoded content of gtk_recent_info_get_uri().
func (i *RecentInfo) URIDisplay() string {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_recent_info_get_uri_display(arg0)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// Visited gets the timestamp (seconds from system’s Epoch) when the meta-data
// for the resource was last visited.
func (i *RecentInfo) Visited() int32 {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret C.time_t
	var ret1 int32

	cret = C.gtk_recent_info_get_visited(arg0)

	ret1 = C.time_t(cret)

	return ret1
}

// HasApplication checks whether an application registered this resource using
// @app_name.
func (i *RecentInfo) HasApplication(appName string) bool {
	var arg0 *C.GtkRecentInfo
	var arg1 *C.gchar

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))
	arg1 = (*C.gchar)(C.CString(appName))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_recent_info_has_application(arg0, appName)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// HasGroup checks whether @group_name appears inside the groups registered for
// the recently used item @info.
func (i *RecentInfo) HasGroup(groupName string) bool {
	var arg0 *C.GtkRecentInfo
	var arg1 *C.gchar

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))
	arg1 = (*C.gchar)(C.CString(groupName))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_recent_info_has_group(arg0, groupName)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// IsLocal checks whether the resource is local or not by looking at the scheme
// of its URI.
func (i *RecentInfo) IsLocal() bool {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_recent_info_is_local(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// LastApplication gets the name of the last application that have registered
// the recently used resource represented by @info.
func (i *RecentInfo) LastApplication() string {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_recent_info_last_application(arg0)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// Match checks whether two RecentInfo-struct point to the same resource.
func (i *RecentInfo) Match(infoB *RecentInfo) bool {
	var arg0 *C.GtkRecentInfo
	var arg1 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))
	arg1 = (*C.GtkRecentInfo)(unsafe.Pointer(infoB.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_recent_info_match(arg0, infoB)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Ref increases the reference count of @recent_info by one.
func (i *RecentInfo) Ref() *RecentInfo {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	var cret *C.GtkRecentInfo
	var ret1 *RecentInfo

	cret = C.gtk_recent_info_ref(arg0)

	ret1 = WrapRecentInfo(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *RecentInfo) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Unref decreases the reference count of @info by one. If the reference count
// reaches zero, @info is deallocated, and the memory freed.
func (i *RecentInfo) Unref() {
	var arg0 *C.GtkRecentInfo

	arg0 = (*C.GtkRecentInfo)(unsafe.Pointer(i.Native()))

	C.gtk_recent_info_unref(arg0)
}

type RecentManagerPrivate struct {
	native C.GtkRecentManagerPrivate
}

// WrapRecentManagerPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRecentManagerPrivate(ptr unsafe.Pointer) *RecentManagerPrivate {
	if ptr == nil {
		return nil
	}

	return (*RecentManagerPrivate)(ptr)
}

func marshalRecentManagerPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRecentManagerPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RecentManagerPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}
