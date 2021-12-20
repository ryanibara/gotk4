// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_recent_manager_error_get_type()), F: marshalRecentManagerError},
		{T: externglib.Type(C.gtk_recent_manager_get_type()), F: marshalRecentManagerer},
		{T: externglib.Type(C.gtk_recent_info_get_type()), F: marshalRecentInfo},
	})
}

// RecentManagerError: error codes for RecentManager operations.
type RecentManagerError C.gint

const (
	// RecentManagerErrorNotFound: URI specified does not exists in the recently
	// used resources list.
	RecentManagerErrorNotFound RecentManagerError = iota
	// RecentManagerErrorInvalidURI: URI specified is not valid.
	RecentManagerErrorInvalidURI
	// RecentManagerErrorInvalidEncoding: supplied string is not UTF-8 encoded.
	RecentManagerErrorInvalidEncoding
	// RecentManagerErrorNotRegistered: no application has registered the
	// specified item.
	RecentManagerErrorNotRegistered
	// RecentManagerErrorRead: failure while reading the recently used resources
	// file.
	RecentManagerErrorRead
	// RecentManagerErrorWrite: failure while writing the recently used
	// resources file.
	RecentManagerErrorWrite
	// RecentManagerErrorUnknown: unspecified error.
	RecentManagerErrorUnknown
)

func marshalRecentManagerError(p uintptr) (interface{}, error) {
	return RecentManagerError(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for RecentManagerError.
func (r RecentManagerError) String() string {
	switch r {
	case RecentManagerErrorNotFound:
		return "NotFound"
	case RecentManagerErrorInvalidURI:
		return "InvalidURI"
	case RecentManagerErrorInvalidEncoding:
		return "InvalidEncoding"
	case RecentManagerErrorNotRegistered:
		return "NotRegistered"
	case RecentManagerErrorRead:
		return "Read"
	case RecentManagerErrorWrite:
		return "Write"
	case RecentManagerErrorUnknown:
		return "Unknown"
	default:
		return fmt.Sprintf("RecentManagerError(%d)", r)
	}
}

// RecentManagerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type RecentManagerOverrider interface {
	Changed()
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
type RecentManager struct {
	*externglib.Object

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ externglib.Objector = (*RecentManager)(nil)
)

func wrapRecentManager(obj *externglib.Object) *RecentManager {
	return &RecentManager{
		Object: obj,
	}
}

func marshalRecentManagerer(p uintptr) (interface{}, error) {
	return wrapRecentManager(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChanged: emitted when the current recently used resources manager
// changes its contents, either by calling gtk_recent_manager_add_item() or by
// another application.
func (manager *RecentManager) ConnectChanged(f func()) externglib.SignalHandle {
	return manager.Connect("changed", f)
}

// NewRecentManager creates a new recent manager object. Recent manager objects
// are used to handle the list of recently used resources. A RecentManager
// object monitors the recently used resources list, and emits the “changed”
// signal each time something inside the list changes.
//
// RecentManager objects are expensive: be sure to create them only when needed.
// You should use gtk_recent_manager_get_default() instead.
//
// The function returns the following values:
//
//    - recentManager: newly created RecentManager object.
//
func NewRecentManager() *RecentManager {
	var _cret *C.GtkRecentManager // in

	_cret = C.gtk_recent_manager_new()

	var _recentManager *RecentManager // out

	_recentManager = wrapRecentManager(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _recentManager
}

// AddFull adds a new resource, pointed by uri, into the recently used resources
// list, using the metadata specified inside the RecentData-struct passed in
// recent_data.
//
// The passed URI will be used to identify this resource inside the list.
//
// In order to register the new recently used resource, metadata about the
// resource must be passed as well as the URI; the metadata is stored in a
// RecentData-struct, which must contain the MIME type of the resource pointed
// by the URI; the name of the application that is registering the item, and a
// command line to be used when launching the item.
//
// Optionally, a RecentData-struct might contain a UTF-8 string to be used when
// viewing the item instead of the last component of the URI; a short
// description of the item; whether the item should be considered private - that
// is, should be displayed only by the applications that have registered it.
//
// The function takes the following parameters:
//
//    - uri: valid URI.
//    - recentData: metadata of the resource.
//
// The function returns the following values:
//
//    - ok: TRUE if the new item was successfully added to the recently used
//      resources list, FALSE otherwise.
//
func (manager *RecentManager) AddFull(uri string, recentData *RecentData) bool {
	var _arg0 *C.GtkRecentManager // out
	var _arg1 *C.gchar            // out
	var _arg2 *C.GtkRecentData    // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkRecentData)(gextras.StructNative(unsafe.Pointer(recentData)))

	_cret = C.gtk_recent_manager_add_full(_arg0, _arg1, _arg2)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(recentData)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AddItem adds a new resource, pointed by uri, into the recently used resources
// list.
//
// This function automatically retrieves some of the needed metadata and setting
// other metadata to common default values; it then feeds the data to
// gtk_recent_manager_add_full().
//
// See gtk_recent_manager_add_full() if you want to explicitly define the
// metadata for the resource pointed by uri.
//
// The function takes the following parameters:
//
//    - uri: valid URI.
//
// The function returns the following values:
//
//    - ok: TRUE if the new item was successfully added to the recently used
//      resources list.
//
func (manager *RecentManager) AddItem(uri string) bool {
	var _arg0 *C.GtkRecentManager // out
	var _arg1 *C.gchar            // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_recent_manager_add_item(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(uri)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Items gets the list of recently used resources.
//
// The function returns the following values:
//
//    - list of newly allocated RecentInfo objects. Use gtk_recent_info_unref()
//      on each item inside the list, and then free the list itself using
//      g_list_free().
//
func (manager *RecentManager) Items() []*RecentInfo {
	var _arg0 *C.GtkRecentManager // out
	var _cret *C.GList            // in

	_arg0 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_recent_manager_get_items(_arg0)
	runtime.KeepAlive(manager)

	var _list []*RecentInfo // out

	_list = make([]*RecentInfo, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkRecentInfo)(v)
		var dst *RecentInfo // out
		dst = (*RecentInfo)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gtk_recent_info_unref((*C.GtkRecentInfo)(intern.C))
			},
		)
		_list = append(_list, dst)
	})

	return _list
}

// HasItem checks whether there is a recently used resource registered with uri
// inside the recent manager.
//
// The function takes the following parameters:
//
//    - uri: URI.
//
// The function returns the following values:
//
//    - ok: TRUE if the resource was found, FALSE otherwise.
//
func (manager *RecentManager) HasItem(uri string) bool {
	var _arg0 *C.GtkRecentManager // out
	var _arg1 *C.gchar            // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_recent_manager_has_item(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(uri)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LookupItem searches for a URI inside the recently used resources list, and
// returns a RecentInfo-struct containing informations about the resource like
// its MIME type, or its display name.
//
// The function takes the following parameters:
//
//    - uri: URI.
//
// The function returns the following values:
//
//    - recentInfo (optional) containing information about the resource pointed
//      by uri, or NULL if the URI was not registered in the recently used
//      resources list. Free with gtk_recent_info_unref().
//
func (manager *RecentManager) LookupItem(uri string) (*RecentInfo, error) {
	var _arg0 *C.GtkRecentManager // out
	var _arg1 *C.gchar            // out
	var _cret *C.GtkRecentInfo    // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_recent_manager_lookup_item(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(uri)

	var _recentInfo *RecentInfo // out
	var _goerr error            // out

	if _cret != nil {
		_recentInfo = (*RecentInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_recentInfo)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gtk_recent_info_unref((*C.GtkRecentInfo)(intern.C))
			},
		)
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _recentInfo, _goerr
}

// MoveItem changes the location of a recently used resource from uri to
// new_uri.
//
// Please note that this function will not affect the resource pointed by the
// URIs, but only the URI used in the recently used resources list.
//
// The function takes the following parameters:
//
//    - uri: URI of a recently used resource.
//    - newUri (optional): new URI of the recently used resource, or NULL to
//      remove the item pointed by uri in the list.
//
func (manager *RecentManager) MoveItem(uri, newUri string) error {
	var _arg0 *C.GtkRecentManager // out
	var _arg1 *C.gchar            // out
	var _arg2 *C.gchar            // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	if newUri != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(newUri)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	C.gtk_recent_manager_move_item(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(newUri)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PurgeItems purges every item from the recently used resources list.
//
// The function returns the following values:
//
//    - gint: number of items that have been removed from the recently used
//      resources list.
//
func (manager *RecentManager) PurgeItems() (int, error) {
	var _arg0 *C.GtkRecentManager // out
	var _cret C.gint              // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_recent_manager_purge_items(_arg0, &_cerr)
	runtime.KeepAlive(manager)

	var _gint int    // out
	var _goerr error // out

	_gint = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gint, _goerr
}

// RemoveItem removes a resource pointed by uri from the recently used resources
// list handled by a recent manager.
//
// The function takes the following parameters:
//
//    - uri: URI of the item you wish to remove.
//
func (manager *RecentManager) RemoveItem(uri string) error {
	var _arg0 *C.GtkRecentManager // out
	var _arg1 *C.gchar            // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_manager_remove_item(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(uri)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// RecentManagerGetDefault gets a unique instance of RecentManager, that you can
// share in your application without caring about memory management.
//
// The function returns the following values:
//
//    - recentManager: unique RecentManager. Do not ref or unref it.
//
func RecentManagerGetDefault() *RecentManager {
	var _cret *C.GtkRecentManager // in

	_cret = C.gtk_recent_manager_get_default()

	var _recentManager *RecentManager // out

	_recentManager = wrapRecentManager(externglib.Take(unsafe.Pointer(_cret)))

	return _recentManager
}

// RecentData: meta-data to be passed to gtk_recent_manager_add_full() when
// registering a recently used resource.
//
// An instance of this type is always passed by reference.
type RecentData struct {
	*recentData
}

// recentData is the struct that's finalized.
type recentData struct {
	native *C.GtkRecentData
}

// DisplayName: UTF-8 encoded string, containing the name of the recently used
// resource to be displayed, or NULL;.
func (r *RecentData) DisplayName() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(r.native.display_name)))
	return v
}

// Description: UTF-8 encoded string, containing a short description of the
// resource, or NULL;.
func (r *RecentData) Description() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(r.native.description)))
	return v
}

// MIMEType: MIME type of the resource;.
func (r *RecentData) MIMEType() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(r.native.mime_type)))
	return v
}

// AppName: name of the application that is registering this recently used
// resource;.
func (r *RecentData) AppName() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(r.native.app_name)))
	return v
}

// AppExec: command line used to launch this resource; may contain the “\f” and
// “\u” escape characters which will be expanded to the resource file path and
// URI respectively when the command line is retrieved;.
func (r *RecentData) AppExec() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(r.native.app_exec)))
	return v
}

// Groups: vector of strings containing groups names;.
func (r *RecentData) Groups() []string {
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

// IsPrivate: whether this resource should be displayed only by the applications
// that have registered it or not.
func (r *RecentData) IsPrivate() bool {
	var v bool // out
	if r.native.is_private != 0 {
		v = true
	}
	return v
}

// RecentInfo contains private data only, and should be accessed using the
// provided API.
//
// RecentInfo constains all the meta-data associated with an entry in the
// recently used files list.
//
// An instance of this type is always passed by reference.
type RecentInfo struct {
	*recentInfo
}

// recentInfo is the struct that's finalized.
type recentInfo struct {
	native *C.GtkRecentInfo
}

func marshalRecentInfo(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &RecentInfo{&recentInfo{(*C.GtkRecentInfo)(b)}}, nil
}

// CreateAppInfo creates a Info for the specified RecentInfo.
//
// The function takes the following parameters:
//
//    - appName (optional): name of the application that should be mapped to a
//      Info; if NULL is used then the default application for the MIME type is
//      used.
//
// The function returns the following values:
//
//    - appInfo (optional): newly created Info, or NULL. In case of error, error
//      will be set either with a GTK_RECENT_MANAGER_ERROR or a G_IO_ERROR.
//
func (info *RecentInfo) CreateAppInfo(appName string) (gio.AppInfor, error) {
	var _arg0 *C.GtkRecentInfo // out
	var _arg1 *C.gchar         // out
	var _cret *C.GAppInfo      // in
	var _cerr *C.GError        // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))
	if appName != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(appName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_recent_info_create_app_info(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(info)
	runtime.KeepAlive(appName)

	var _appInfo gio.AppInfor // out
	var _goerr error          // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			casted := object.Cast()
			rv, ok := casted.(gio.AppInfor)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.AppInfor")
			}
			_appInfo = rv
		}
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _appInfo, _goerr
}

// Exists checks whether the resource pointed by info still exists. At the
// moment this check is done only on resources pointing to local files.
//
// The function returns the following values:
//
//    - ok: TRUE if the resource exists.
//
func (info *RecentInfo) Exists() bool {
	var _arg0 *C.GtkRecentInfo // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_exists(_arg0)
	runtime.KeepAlive(info)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Added gets the timestamp (seconds from system’s Epoch) when the resource was
// added to the recently used resources list.
//
// The function returns the following values:
//
//    - glong: number of seconds elapsed from system’s Epoch when the resource
//      was added to the list, or -1 on failure.
//
func (info *RecentInfo) Added() int32 {
	var _arg0 *C.GtkRecentInfo // out
	var _cret C.time_t         // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_get_added(_arg0)
	runtime.KeepAlive(info)

	var _glong int32 // out

	_glong = int32(_cret)

	return _glong
}

// Age gets the number of days elapsed since the last update of the resource
// pointed by info.
//
// The function returns the following values:
//
//    - gint: positive integer containing the number of days elapsed since the
//      time this resource was last modified.
//
func (info *RecentInfo) Age() int {
	var _arg0 *C.GtkRecentInfo // out
	var _cret C.gint           // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_get_age(_arg0)
	runtime.KeepAlive(info)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ApplicationInfo gets the data regarding the application that has registered
// the resource pointed by info.
//
// If the command line contains any escape characters defined inside the storage
// specification, they will be expanded.
//
// The function takes the following parameters:
//
//    - appName: name of the application that has registered this item.
//
// The function returns the following values:
//
//    - appExec: return location for the string containing the command line.
//    - count: return location for the number of times this item was registered.
//    - time_: return location for the timestamp this item was last registered
//      for this application.
//    - ok: TRUE if an application with app_name has registered this resource
//      inside the recently used list, or FALSE otherwise. The app_exec string is
//      owned by the RecentInfo and should not be modified or freed.
//
func (info *RecentInfo) ApplicationInfo(appName string) (string, uint, int32, bool) {
	var _arg0 *C.GtkRecentInfo // out
	var _arg1 *C.gchar         // out
	var _arg2 *C.gchar         // in
	var _arg3 C.guint          // in
	var _arg4 C.time_t         // in
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(appName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_recent_info_get_application_info(_arg0, _arg1, &_arg2, &_arg3, &_arg4)
	runtime.KeepAlive(info)
	runtime.KeepAlive(appName)

	var _appExec string // out
	var _count uint     // out
	var _time_ int32    // out
	var _ok bool        // out

	_appExec = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	_count = uint(_arg3)
	_time_ = int32(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _appExec, _count, _time_, _ok
}

// Applications retrieves the list of applications that have registered this
// resource.
//
// The function returns the following values:
//
//    - utf8s: a newly allocated NULL-terminated array of strings. Use
//      g_strfreev() to free it.
//
func (info *RecentInfo) Applications() []string {
	var _arg0 *C.GtkRecentInfo // out
	var _cret **C.gchar        // in
	var _arg1 C.gsize          // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_get_applications(_arg0, &_arg1)
	runtime.KeepAlive(info)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		src := unsafe.Slice(_cret, _arg1)
		_utf8s = make([]string, _arg1)
		for i := 0; i < int(_arg1); i++ {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// Description gets the (short) description of the resource.
//
// The function returns the following values:
//
//    - utf8: description of the resource. The returned string is owned by the
//      recent manager, and should not be freed.
//
func (info *RecentInfo) Description() string {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_get_description(_arg0)
	runtime.KeepAlive(info)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// DisplayName gets the name of the resource. If none has been defined, the
// basename of the resource is obtained.
//
// The function returns the following values:
//
//    - utf8: display name of the resource. The returned string is owned by the
//      recent manager, and should not be freed.
//
func (info *RecentInfo) DisplayName() string {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_get_display_name(_arg0)
	runtime.KeepAlive(info)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// GIcon retrieves the icon associated to the resource MIME type.
//
// The function returns the following values:
//
//    - icon (optional) containing the icon, or NULL. Use g_object_unref() when
//      finished using the icon.
//
func (info *RecentInfo) GIcon() gio.Iconner {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.GIcon         // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_get_gicon(_arg0)
	runtime.KeepAlive(info)

	var _icon gio.Iconner // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			casted := object.Cast()
			rv, ok := casted.(gio.Iconner)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.Iconner")
			}
			_icon = rv
		}
	}

	return _icon
}

// Groups returns all groups registered for the recently used item info. The
// array of returned group names will be NULL terminated, so length might
// optionally be NULL.
//
// The function returns the following values:
//
//    - utf8s: a newly allocated NULL terminated array of strings. Use
//      g_strfreev() to free it.
//
func (info *RecentInfo) Groups() []string {
	var _arg0 *C.GtkRecentInfo // out
	var _cret **C.gchar        // in
	var _arg1 C.gsize          // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_get_groups(_arg0, &_arg1)
	runtime.KeepAlive(info)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		src := unsafe.Slice(_cret, _arg1)
		_utf8s = make([]string, _arg1)
		for i := 0; i < int(_arg1); i++ {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// Icon retrieves the icon of size size associated to the resource MIME type.
//
// The function takes the following parameters:
//
//    - size of the icon in pixels.
//
// The function returns the following values:
//
//    - pixbuf (optional) containing the icon, or NULL. Use g_object_unref() when
//      finished using the icon.
//
func (info *RecentInfo) Icon(size int) *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkRecentInfo // out
	var _arg1 C.gint           // out
	var _cret *C.GdkPixbuf     // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = C.gint(size)

	_cret = C.gtk_recent_info_get_icon(_arg0, _arg1)
	runtime.KeepAlive(info)
	runtime.KeepAlive(size)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if _cret != nil {
		{
			obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
			_pixbuf = &gdkpixbuf.Pixbuf{
				Object: obj,
				LoadableIcon: gio.LoadableIcon{
					Icon: gio.Icon{
						Object: obj,
					},
				},
			}
		}
	}

	return _pixbuf
}

// MIMEType gets the MIME type of the resource.
//
// The function returns the following values:
//
//    - utf8: MIME type of the resource. The returned string is owned by the
//      recent manager, and should not be freed.
//
func (info *RecentInfo) MIMEType() string {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_get_mime_type(_arg0)
	runtime.KeepAlive(info)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Modified gets the timestamp (seconds from system’s Epoch) when the meta-data
// for the resource was last modified.
//
// The function returns the following values:
//
//    - glong: number of seconds elapsed from system’s Epoch when the resource
//      was last modified, or -1 on failure.
//
func (info *RecentInfo) Modified() int32 {
	var _arg0 *C.GtkRecentInfo // out
	var _cret C.time_t         // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_get_modified(_arg0)
	runtime.KeepAlive(info)

	var _glong int32 // out

	_glong = int32(_cret)

	return _glong
}

// PrivateHint gets the value of the “private” flag. Resources in the recently
// used list that have this flag set to TRUE should only be displayed by the
// applications that have registered them.
//
// The function returns the following values:
//
//    - ok: TRUE if the private flag was found, FALSE otherwise.
//
func (info *RecentInfo) PrivateHint() bool {
	var _arg0 *C.GtkRecentInfo // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_get_private_hint(_arg0)
	runtime.KeepAlive(info)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShortName computes a valid UTF-8 string that can be used as the name of the
// item in a menu or list. For example, calling this function on an item that
// refers to “file:///foo/bar.txt” will yield “bar.txt”.
//
// The function returns the following values:
//
//    - utf8: newly-allocated string in UTF-8 encoding free it with g_free().
//
func (info *RecentInfo) ShortName() string {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_get_short_name(_arg0)
	runtime.KeepAlive(info)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// URI gets the URI of the resource.
//
// The function returns the following values:
//
//    - utf8: URI of the resource. The returned string is owned by the recent
//      manager, and should not be freed.
//
func (info *RecentInfo) URI() string {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_get_uri(_arg0)
	runtime.KeepAlive(info)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// URIDisplay gets a displayable version of the resource’s URI. If the resource
// is local, it returns a local path; if the resource is not local, it returns
// the UTF-8 encoded content of gtk_recent_info_get_uri().
//
// The function returns the following values:
//
//    - utf8 (optional): newly allocated UTF-8 string containing the resource’s
//      URI or NULL. Use g_free() when done using it.
//
func (info *RecentInfo) URIDisplay() string {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_get_uri_display(_arg0)
	runtime.KeepAlive(info)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Visited gets the timestamp (seconds from system’s Epoch) when the meta-data
// for the resource was last visited.
//
// The function returns the following values:
//
//    - glong: number of seconds elapsed from system’s Epoch when the resource
//      was last visited, or -1 on failure.
//
func (info *RecentInfo) Visited() int32 {
	var _arg0 *C.GtkRecentInfo // out
	var _cret C.time_t         // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_get_visited(_arg0)
	runtime.KeepAlive(info)

	var _glong int32 // out

	_glong = int32(_cret)

	return _glong
}

// HasApplication checks whether an application registered this resource using
// app_name.
//
// The function takes the following parameters:
//
//    - appName: string containing an application name.
//
// The function returns the following values:
//
//    - ok: TRUE if an application with name app_name was found, FALSE otherwise.
//
func (info *RecentInfo) HasApplication(appName string) bool {
	var _arg0 *C.GtkRecentInfo // out
	var _arg1 *C.gchar         // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(appName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_recent_info_has_application(_arg0, _arg1)
	runtime.KeepAlive(info)
	runtime.KeepAlive(appName)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasGroup checks whether group_name appears inside the groups registered for
// the recently used item info.
//
// The function takes the following parameters:
//
//    - groupName: name of a group.
//
// The function returns the following values:
//
//    - ok: TRUE if the group was found.
//
func (info *RecentInfo) HasGroup(groupName string) bool {
	var _arg0 *C.GtkRecentInfo // out
	var _arg1 *C.gchar         // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_recent_info_has_group(_arg0, _arg1)
	runtime.KeepAlive(info)
	runtime.KeepAlive(groupName)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsLocal checks whether the resource is local or not by looking at the scheme
// of its URI.
//
// The function returns the following values:
//
//    - ok: TRUE if the resource is local.
//
func (info *RecentInfo) IsLocal() bool {
	var _arg0 *C.GtkRecentInfo // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_is_local(_arg0)
	runtime.KeepAlive(info)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LastApplication gets the name of the last application that have registered
// the recently used resource represented by info.
//
// The function returns the following values:
//
//    - utf8: application name. Use g_free() to free it.
//
func (info *RecentInfo) LastApplication() string {
	var _arg0 *C.GtkRecentInfo // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gtk_recent_info_last_application(_arg0)
	runtime.KeepAlive(info)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Match checks whether two RecentInfo-struct point to the same resource.
//
// The function takes the following parameters:
//
//    - infoB: RecentInfo.
//
// The function returns the following values:
//
//    - ok: TRUE if both RecentInfo-struct point to the same resource, FALSE
//      otherwise.
//
func (infoA *RecentInfo) Match(infoB *RecentInfo) bool {
	var _arg0 *C.GtkRecentInfo // out
	var _arg1 *C.GtkRecentInfo // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(infoA)))
	_arg1 = (*C.GtkRecentInfo)(gextras.StructNative(unsafe.Pointer(infoB)))

	_cret = C.gtk_recent_info_match(_arg0, _arg1)
	runtime.KeepAlive(infoA)
	runtime.KeepAlive(infoB)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
