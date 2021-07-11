// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_desktop_app_info_lookup_get_type()), F: marshalDesktopAppInfoLookupper},
		{T: externglib.Type(C.g_desktop_app_info_get_type()), F: marshalDesktopAppInfor},
	})
}

// DesktopAppInfoLookupOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DesktopAppInfoLookupOverrider interface {
	// DefaultForURIScheme gets the default application for launching
	// applications using this URI scheme for a particular AppInfoLookup
	// implementation.
	//
	// The AppInfoLookup interface and this function is used to implement
	// g_app_info_get_default_for_uri_scheme() backends in a GIO module. There
	// is no reason for applications to use it directly. Applications should use
	// g_app_info_get_default_for_uri_scheme().
	//
	// Deprecated: The AppInfoLookup interface is deprecated and unused by GIO.
	DefaultForURIScheme(uriScheme string) *AppInfo
}

// DesktopAppInfoLookupper describes DesktopAppInfoLookup's methods.
type DesktopAppInfoLookupper interface {
	// DefaultForURIScheme gets the default application for launching
	// applications using this URI scheme for a particular AppInfoLookup
	// implementation.
	DefaultForURIScheme(uriScheme string) *AppInfo
}

// DesktopAppInfoLookup is an opaque data structure and can only be accessed
// using the following functions.
//
// Deprecated: The AppInfoLookup interface is deprecated and unused by GIO.
type DesktopAppInfoLookup struct {
	*externglib.Object
}

var (
	_ DesktopAppInfoLookupper = (*DesktopAppInfoLookup)(nil)
	_ gextras.Nativer         = (*DesktopAppInfoLookup)(nil)
)

func wrapDesktopAppInfoLookup(obj *externglib.Object) DesktopAppInfoLookupper {
	return &DesktopAppInfoLookup{
		Object: obj,
	}
}

func marshalDesktopAppInfoLookupper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDesktopAppInfoLookup(obj), nil
}

// DefaultForURIScheme gets the default application for launching applications
// using this URI scheme for a particular AppInfoLookup implementation.
//
// The AppInfoLookup interface and this function is used to implement
// g_app_info_get_default_for_uri_scheme() backends in a GIO module. There is no
// reason for applications to use it directly. Applications should use
// g_app_info_get_default_for_uri_scheme().
//
// Deprecated: The AppInfoLookup interface is deprecated and unused by GIO.
func (lookup *DesktopAppInfoLookup) DefaultForURIScheme(uriScheme string) *AppInfo {
	var _arg0 *C.GDesktopAppInfoLookup // out
	var _arg1 *C.char                  // out
	var _cret *C.GAppInfo              // in

	_arg0 = (*C.GDesktopAppInfoLookup)(unsafe.Pointer(lookup.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uriScheme)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_lookup_get_default_for_uri_scheme(_arg0, _arg1)

	var _appInfo *AppInfo // out

	_appInfo = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*AppInfo)

	return _appInfo
}

// DesktopAppInfor describes DesktopAppInfo's methods.
type DesktopAppInfor interface {
	// ActionName gets the user-visible display name of the "additional
	// application action" specified by @action_name.
	ActionName(actionName string) string
	// Boolean looks up a boolean value in the keyfile backing @info.
	Boolean(key string) bool
	// Categories gets the categories from the desktop file.
	Categories() string
	// Filename: when @info was created from a known filename, return it.
	Filename() string
	// GenericName gets the generic name from the desktop file.
	GenericName() string
	// IsHidden: desktop file is hidden if the Hidden key in it is set to True.
	IsHidden() bool
	// Keywords gets the keywords from the desktop file.
	Keywords() []string
	// LocaleString looks up a localized string value in the keyfile backing
	// @info translated to the current locale.
	LocaleString(key string) string
	// Nodisplay gets the value of the NoDisplay key, which helps determine if
	// the application info should be shown in menus.
	Nodisplay() bool
	// ShowIn checks if the application info should be shown in menus that list
	// available applications for a specific name of the desktop, based on the
	// `OnlyShowIn` and `NotShowIn` keys.
	ShowIn(desktopEnv string) bool
	// StartupWmClass retrieves the StartupWMClass field from @info.
	StartupWmClass() string
	// String looks up a string value in the keyfile backing @info.
	String(key string) string
	// HasKey returns whether @key exists in the "Desktop Entry" group of the
	// keyfile backing @info.
	HasKey(key string) bool
	// LaunchAction activates the named application action.
	LaunchAction(actionName string, launchContext AppLaunchContexter)
	// ListActions returns the list of "additional application actions"
	// supported on the desktop file, as per the desktop file specification.
	ListActions() []string
}

// DesktopAppInfo is an implementation of Info based on desktop files.
//
// Note that `<gio/gdesktopappinfo.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file when
// using it.
type DesktopAppInfo struct {
	*externglib.Object

	AppInfo
}

var (
	_ DesktopAppInfor = (*DesktopAppInfo)(nil)
	_ gextras.Nativer = (*DesktopAppInfo)(nil)
)

func wrapDesktopAppInfo(obj *externglib.Object) DesktopAppInfor {
	return &DesktopAppInfo{
		Object: obj,
		AppInfo: AppInfo{
			Object: obj,
		},
	}
}

func marshalDesktopAppInfor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDesktopAppInfo(obj), nil
}

// NewDesktopAppInfo creates a new AppInfo based on a desktop file id.
//
// A desktop file id is the basename of the desktop file, including the .desktop
// extension. GIO is looking for a desktop file with this name in the
// `applications` subdirectories of the XDG data directories (i.e. the
// directories specified in the `XDG_DATA_HOME` and `XDG_DATA_DIRS` environment
// variables). GIO also supports the prefix-to-subdirectory mapping that is
// described in the Menu Spec
// (http://standards.freedesktop.org/menu-spec/latest/) (i.e. a desktop id of
// kde-foo.desktop will match `/usr/share/applications/kde/foo.desktop`).
func NewDesktopAppInfo(desktopId string) *DesktopAppInfo {
	var _arg1 *C.char            // out
	var _cret *C.GDesktopAppInfo // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(desktopId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_new(_arg1)

	var _desktopAppInfo *DesktopAppInfo // out

	_desktopAppInfo = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*DesktopAppInfo)

	return _desktopAppInfo
}

// NewDesktopAppInfoFromFilename creates a new AppInfo.
func NewDesktopAppInfoFromFilename(filename string) *DesktopAppInfo {
	var _arg1 *C.char            // out
	var _cret *C.GDesktopAppInfo // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_new_from_filename(_arg1)

	var _desktopAppInfo *DesktopAppInfo // out

	_desktopAppInfo = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*DesktopAppInfo)

	return _desktopAppInfo
}

// NewDesktopAppInfoFromKeyfile creates a new AppInfo.
func NewDesktopAppInfoFromKeyfile(keyFile *glib.KeyFile) *DesktopAppInfo {
	var _arg1 *C.GKeyFile        // out
	var _cret *C.GDesktopAppInfo // in

	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile))

	_cret = C.g_desktop_app_info_new_from_keyfile(_arg1)

	var _desktopAppInfo *DesktopAppInfo // out

	_desktopAppInfo = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*DesktopAppInfo)

	return _desktopAppInfo
}

// ActionName gets the user-visible display name of the "additional application
// action" specified by @action_name.
//
// This corresponds to the "Name" key within the keyfile group for the action.
func (info *DesktopAppInfo) ActionName(actionName string) string {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.gchar           // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_get_action_name(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Boolean looks up a boolean value in the keyfile backing @info.
//
// The @key is looked up in the "Desktop Entry" group.
func (info *DesktopAppInfo) Boolean(key string) bool {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.char            // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_get_boolean(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Categories gets the categories from the desktop file.
func (info *DesktopAppInfo) Categories() string {
	var _arg0 *C.GDesktopAppInfo // out
	var _cret *C.char            // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_desktop_app_info_get_categories(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Filename: when @info was created from a known filename, return it. In some
// situations such as the AppInfo returned from
// g_desktop_app_info_new_from_keyfile(), this function will return nil.
func (info *DesktopAppInfo) Filename() string {
	var _arg0 *C.GDesktopAppInfo // out
	var _cret *C.char            // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_desktop_app_info_get_filename(_arg0)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// GenericName gets the generic name from the desktop file.
func (info *DesktopAppInfo) GenericName() string {
	var _arg0 *C.GDesktopAppInfo // out
	var _cret *C.char            // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_desktop_app_info_get_generic_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// IsHidden: desktop file is hidden if the Hidden key in it is set to True.
func (info *DesktopAppInfo) IsHidden() bool {
	var _arg0 *C.GDesktopAppInfo // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_desktop_app_info_get_is_hidden(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Keywords gets the keywords from the desktop file.
func (info *DesktopAppInfo) Keywords() []string {
	var _arg0 *C.GDesktopAppInfo // out
	var _cret **C.char

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_desktop_app_info_get_keywords(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
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

// LocaleString looks up a localized string value in the keyfile backing @info
// translated to the current locale.
//
// The @key is looked up in the "Desktop Entry" group.
func (info *DesktopAppInfo) LocaleString(key string) string {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.char            // out
	var _cret *C.char            // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_get_locale_string(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Nodisplay gets the value of the NoDisplay key, which helps determine if the
// application info should be shown in menus. See
// KEY_FILE_DESKTOP_KEY_NO_DISPLAY and g_app_info_should_show().
func (info *DesktopAppInfo) Nodisplay() bool {
	var _arg0 *C.GDesktopAppInfo // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_desktop_app_info_get_nodisplay(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowIn checks if the application info should be shown in menus that list
// available applications for a specific name of the desktop, based on the
// `OnlyShowIn` and `NotShowIn` keys.
//
// @desktop_env should typically be given as nil, in which case the
// `XDG_CURRENT_DESKTOP` environment variable is consulted. If you want to
// override the default mechanism then you may specify @desktop_env, but this is
// not recommended.
//
// Note that g_app_info_should_show() for @info will include this check (with
// nil for @desktop_env) as well as additional checks.
func (info *DesktopAppInfo) ShowIn(desktopEnv string) bool {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.gchar           // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(desktopEnv)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_get_show_in(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StartupWmClass retrieves the StartupWMClass field from @info. This represents
// the WM_CLASS property of the main window of the application, if launched
// through @info.
func (info *DesktopAppInfo) StartupWmClass() string {
	var _arg0 *C.GDesktopAppInfo // out
	var _cret *C.char            // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_desktop_app_info_get_startup_wm_class(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// String looks up a string value in the keyfile backing @info.
//
// The @key is looked up in the "Desktop Entry" group.
func (info *DesktopAppInfo) String(key string) string {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.char            // out
	var _cret *C.char            // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_get_string(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// HasKey returns whether @key exists in the "Desktop Entry" group of the
// keyfile backing @info.
func (info *DesktopAppInfo) HasKey(key string) bool {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.char            // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_desktop_app_info_has_key(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LaunchAction activates the named application action.
//
// You may only call this function on action names that were returned from
// g_desktop_app_info_list_actions().
//
// Note that if the main entry of the desktop file indicates that the
// application supports startup notification, and @launch_context is non-nil,
// then startup notification will be used when activating the action (and as
// such, invocation of the action on the receiving side must signal the end of
// startup notification when it is completed). This is the expected behaviour of
// applications declaring additional actions, as per the desktop file
// specification.
//
// As with g_app_info_launch() there is no way to detect failures that occur
// while using this function.
func (info *DesktopAppInfo) LaunchAction(actionName string, launchContext AppLaunchContexter) {
	var _arg0 *C.GDesktopAppInfo   // out
	var _arg1 *C.gchar             // out
	var _arg2 *C.GAppLaunchContext // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GAppLaunchContext)(unsafe.Pointer((launchContext).(gextras.Nativer).Native()))

	C.g_desktop_app_info_launch_action(_arg0, _arg1, _arg2)
}

// ListActions returns the list of "additional application actions" supported on
// the desktop file, as per the desktop file specification.
//
// As per the specification, this is the list of actions that are explicitly
// listed in the "Actions" key of the [Desktop Entry] group.
func (info *DesktopAppInfo) ListActions() []string {
	var _arg0 *C.GDesktopAppInfo // out
	var _cret **C.gchar

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_desktop_app_info_list_actions(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
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
