// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/ptr"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

// UserDirectory: these are logical ids for special directories which are
// defined depending on the platform used. You should use
// g_get_user_special_dir() to retrieve the full path associated to the logical
// id.
//
// The Directory enumeration can be extended at later date. Not every platform
// has a directory for every logical id in this enumeration.
type UserDirectory int

const (
	// UserDirectoryDirectoryDesktop: the user's Desktop directory
	UserDirectoryDirectoryDesktop UserDirectory = 0
	// UserDirectoryDirectoryDocuments: the user's Documents directory
	UserDirectoryDirectoryDocuments UserDirectory = 1
	// UserDirectoryDirectoryDownload: the user's Downloads directory
	UserDirectoryDirectoryDownload UserDirectory = 2
	// UserDirectoryDirectoryMusic: the user's Music directory
	UserDirectoryDirectoryMusic UserDirectory = 3
	// UserDirectoryDirectoryPictures: the user's Pictures directory
	UserDirectoryDirectoryPictures UserDirectory = 4
	// UserDirectoryDirectoryPublicShare: the user's shared directory
	UserDirectoryDirectoryPublicShare UserDirectory = 5
	// UserDirectoryDirectoryTemplates: the user's Templates directory
	UserDirectoryDirectoryTemplates UserDirectory = 6
	// UserDirectoryDirectoryVideos: the user's Movies directory
	UserDirectoryDirectoryVideos UserDirectory = 7
	// UserDirectoryNDirectories: the number of enum values
	UserDirectoryNDirectories UserDirectory = 8
)

// FormatSizeFlags flags to modify the format of the string returned by
// g_format_size_full().
type FormatSizeFlags int

const (
	// FormatSizeFlagsDefault: behave the same as g_format_size()
	FormatSizeFlagsDefault FormatSizeFlags = 0b0
	// FormatSizeFlagsLongFormat: include the exact number of bytes as part of
	// the returned string. For example, "45.6 kB (45,612 bytes)".
	FormatSizeFlagsLongFormat FormatSizeFlags = 0b1
	// FormatSizeFlagsIecUnits: use IEC (base 1024) units with "KiB"-style
	// suffixes. IEC units should only be used for reporting things with a
	// strong "power of 2" basis, like RAM sizes or RAID stripe sizes. Network
	// and storage sizes should be reported in the normal SI units.
	FormatSizeFlagsIecUnits FormatSizeFlags = 0b10
	// FormatSizeFlagsBits: set the size as a quantity in bits, rather than
	// bytes, and return units in bits. For example, ‘Mb’ rather than ‘MB’.
	FormatSizeFlagsBits FormatSizeFlags = 0b100
)

// BitNthLsf: find the position of the first bit set in @mask, searching from
// (but not including) @nth_bit upwards. Bits are numbered from 0 (least
// significant) to sizeof(#gulong) * 8 - 1 (31 or 63, usually). To start
// searching from the 0th bit, set @nth_bit to -1.
func BitNthLsf(mask uint32, nthBit int) int {
	var arg1 C.gulong
	var arg2 C.gint

	arg1 = C.gulong(mask)
	arg2 = C.gint(nthBit)

	var cret C.gint
	var ret1 int

	cret = C.g_bit_nth_lsf(mask, nthBit)

	ret1 = C.gint(cret)

	return ret1
}

// BitNthMsf: find the position of the first bit set in @mask, searching from
// (but not including) @nth_bit downwards. Bits are numbered from 0 (least
// significant) to sizeof(#gulong) * 8 - 1 (31 or 63, usually). To start
// searching from the last bit, set @nth_bit to -1 or GLIB_SIZEOF_LONG * 8.
func BitNthMsf(mask uint32, nthBit int) int {
	var arg1 C.gulong
	var arg2 C.gint

	arg1 = C.gulong(mask)
	arg2 = C.gint(nthBit)

	var cret C.gint
	var ret1 int

	cret = C.g_bit_nth_msf(mask, nthBit)

	ret1 = C.gint(cret)

	return ret1
}

// BitStorage gets the number of bits used to hold @number, e.g. if @number is
// 4, 3 bits are needed.
func BitStorage(number uint32) uint {
	var arg1 C.gulong

	arg1 = C.gulong(number)

	var cret C.guint
	var ret1 uint

	cret = C.g_bit_storage(number)

	ret1 = C.guint(cret)

	return ret1
}

// FindProgramInPath locates the first executable named @program in the user's
// path, in the same way that execvp() would locate it. Returns an allocated
// string with the absolute path name, or nil if the program is not found in the
// path. If @program is already an absolute path, returns a copy of @program if
// @program exists and is executable, and nil otherwise. On Windows, if @program
// does not have a file type suffix, tries with the suffixes .exe, .cmd, .bat
// and .com, and the suffixes in the `PATHEXT` environment variable.
//
// On Windows, it looks for the file in the same way as CreateProcess() would.
// This means first in the directory where the executing program was loaded
// from, then in the current directory, then in the Windows 32-bit system
// directory, then in the Windows directory, and finally in the directories in
// the `PATH` environment variable. If the program is found, the return value
// contains the full name including the type suffix.
func FindProgramInPath(program string) string {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(program))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var ret1 string

	cret = C.g_find_program_in_path(program)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// FormatSize formats a size (for example the size of a file) into a human
// readable string. Sizes are rounded to the nearest size prefix (kB, MB, GB)
// and are displayed rounded to the nearest tenth. E.g. the file size 3292528
// bytes will be converted into the string "3.2 MB". The returned string is
// UTF-8, and may use a non-breaking space to separate the number and units, to
// ensure they aren’t separated when line wrapped.
//
// The prefix units base is 1000 (i.e. 1 kB is 1000 bytes).
//
// This string should be freed with g_free() when not needed any longer.
//
// See g_format_size_full() for more options about how the size might be
// formatted.
func FormatSize(size uint64) string {
	var arg1 C.guint64

	arg1 = C.guint64(size)

	var cret *C.gchar
	var ret1 string

	cret = C.g_format_size(size)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// FormatSizeForDisplay formats a size (for example the size of a file) into a
// human readable string. Sizes are rounded to the nearest size prefix (KB, MB,
// GB) and are displayed rounded to the nearest tenth. E.g. the file size
// 3292528 bytes will be converted into the string "3.1 MB".
//
// The prefix units base is 1024 (i.e. 1 KB is 1024 bytes).
//
// This string should be freed with g_free() when not needed any longer.
func FormatSizeForDisplay(size int64) string {
	var arg1 C.goffset

	arg1 = C.goffset(size)

	var cret *C.gchar
	var ret1 string

	cret = C.g_format_size_for_display(size)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// FormatSizeFull formats a size.
//
// This function is similar to g_format_size() but allows for flags that modify
// the output. See SizeFlags.
func FormatSizeFull(size uint64, flags FormatSizeFlags) string {
	var arg1 C.guint64
	var arg2 C.GFormatSizeFlags

	arg1 = C.guint64(size)
	arg2 = (C.GFormatSizeFlags)(flags)

	var cret *C.gchar
	var ret1 string

	cret = C.g_format_size_full(size, flags)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// GetApplicationName gets a human-readable name for the application, as set by
// g_set_application_name(). This name should be localized if possible, and is
// intended for display to the user. Contrast with g_get_prgname(), which gets a
// non-localized name. If g_set_application_name() has not been called, returns
// the result of g_get_prgname() (which may be nil if g_set_prgname() has also
// not been called).
func GetApplicationName() string {
	var cret *C.gchar
	var ret1 string

	cret = C.g_get_application_name()

	ret1 = C.GoString(cret)

	return ret1
}

// GetHomeDir gets the current user's home directory.
//
// As with most UNIX tools, this function will return the value of the `HOME`
// environment variable if it is set to an existing absolute path name, falling
// back to the `passwd` file in the case that it is unset.
//
// If the path given in `HOME` is non-absolute, does not exist, or is not a
// directory, the result is undefined.
//
// Before version 2.36 this function would ignore the `HOME` environment
// variable, taking the value from the `passwd` database instead. This was
// changed to increase the compatibility of GLib with other programs (and the
// XDG basedir specification) and to increase testability of programs based on
// GLib (by making it easier to run them from test frameworks).
//
// If your program has a strong requirement for either the new or the old
// behaviour (and if you don't wish to increase your GLib dependency to ensure
// that the new behaviour is in effect) then you should either directly check
// the `HOME` environment variable yourself or unset it before calling any
// functions in GLib.
func GetHomeDir() string {
	var cret *C.gchar
	var ret1 string

	cret = C.g_get_home_dir()

	ret1 = C.GoString(cret)

	return ret1
}

// GetHostName: return a name for the machine.
//
// The returned name is not necessarily a fully-qualified domain name, or even
// present in DNS or some other name service at all. It need not even be unique
// on your local network or site, but usually it is. Callers should not rely on
// the return value having any specific properties like uniqueness for security
// purposes. Even if the name of the machine is changed while an application is
// running, the return value from this function does not change. The returned
// string is owned by GLib and should not be modified or freed. If no name can
// be determined, a default fixed string "localhost" is returned.
//
// The encoding of the returned string is UTF-8.
func GetHostName() string {
	var cret *C.gchar
	var ret1 string

	cret = C.g_get_host_name()

	ret1 = C.GoString(cret)

	return ret1
}

// GetOsInfo: get information about the operating system.
//
// On Linux this comes from the `/etc/os-release` file. On other systems, it may
// come from a variety of sources. You can either use the standard key names
// like G_OS_INFO_KEY_NAME or pass any UTF-8 string key name. For example,
// `/etc/os-release` provides a number of other less commonly used values that
// may be useful. No key is guaranteed to be provided, so the caller should
// always check if the result is nil.
func GetOsInfo(keyName string) string {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(keyName))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var ret1 string

	cret = C.g_get_os_info(keyName)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// GetPrgname gets the name of the program. This name should not be localized,
// in contrast to g_get_application_name().
//
// If you are using #GApplication the program name is set in
// g_application_run(). In case of GDK or GTK+ it is set in gdk_init(), which is
// called by gtk_init() and the Application::startup handler. The program name
// is found by taking the last component of @argv[0].
func GetPrgname() string {
	var cret *C.gchar
	var ret1 string

	cret = C.g_get_prgname()

	ret1 = C.GoString(cret)

	return ret1
}

// GetRealName gets the real name of the user. This usually comes from the
// user's entry in the `passwd` file. The encoding of the returned string is
// system-defined. (On Windows, it is, however, always UTF-8.) If the real user
// name cannot be determined, the string "Unknown" is returned.
func GetRealName() string {
	var cret *C.gchar
	var ret1 string

	cret = C.g_get_real_name()

	ret1 = C.GoString(cret)

	return ret1
}

// GetSystemConfigDirs returns an ordered list of base directories in which to
// access system-wide configuration information.
//
// On UNIX platforms this is determined using the mechanisms described in the
// XDG Base Directory Specification
// (http://www.freedesktop.org/Standards/basedir-spec). In this case the list of
// directories retrieved will be `XDG_CONFIG_DIRS`.
//
// On Windows it follows XDG Base Directory Specification if `XDG_CONFIG_DIRS`
// is defined. If `XDG_CONFIG_DIRS` is undefined, the directory that contains
// application data for all users is used instead. A typical path is
// `C:\Documents and Settings\All Users\Application Data`. This folder is used
// for application data that is not user specific. For example, an application
// can store a spell-check dictionary, a database of clip art, or a log file in
// the CSIDL_COMMON_APPDATA folder. This information will not roam and is
// available to anyone using the computer.
func GetSystemConfigDirs() []string {
	var cret **C.gchar
	var ret1 []string

	cret = C.g_get_system_config_dirs()

	{
		var length int
		for p := cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		ret1 = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
			ret1[i] = C.GoString(src)
		}
	}

	return ret1
}

// GetSystemDataDirs returns an ordered list of base directories in which to
// access system-wide application data.
//
// On UNIX platforms this is determined using the mechanisms described in the
// XDG Base Directory Specification
// (http://www.freedesktop.org/Standards/basedir-spec) In this case the list of
// directories retrieved will be `XDG_DATA_DIRS`.
//
// On Windows it follows XDG Base Directory Specification if `XDG_DATA_DIRS` is
// defined. If `XDG_DATA_DIRS` is undefined, the first elements in the list are
// the Application Data and Documents folders for All Users. (These can be
// determined only on Windows 2000 or later and are not present in the list on
// other Windows versions.) See documentation for CSIDL_COMMON_APPDATA and
// CSIDL_COMMON_DOCUMENTS.
//
// Then follows the "share" subfolder in the installation folder for the package
// containing the DLL that calls this function, if it can be determined.
//
// Finally the list contains the "share" subfolder in the installation folder
// for GLib, and in the installation folder for the package the application's
// .exe file belongs to.
//
// The installation folders above are determined by looking up the folder where
// the module (DLL or EXE) in question is located. If the folder's name is
// "bin", its parent is used, otherwise the folder itself.
//
// Note that on Windows the returned list can vary depending on where this
// function is called.
func GetSystemDataDirs() []string {
	var cret **C.gchar
	var ret1 []string

	cret = C.g_get_system_data_dirs()

	{
		var length int
		for p := cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		ret1 = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
			ret1[i] = C.GoString(src)
		}
	}

	return ret1
}

// GetTmpDir gets the directory to use for temporary files.
//
// On UNIX, this is taken from the `TMPDIR` environment variable. If the
// variable is not set, `P_tmpdir` is used, as defined by the system C library.
// Failing that, a hard-coded default of "/tmp" is returned.
//
// On Windows, the `TEMP` environment variable is used, with the root directory
// of the Windows installation (eg: "C:\") used as a default.
//
// The encoding of the returned string is system-defined. On Windows, it is
// always UTF-8. The return value is never nil or the empty string.
func GetTmpDir() string {
	var cret *C.gchar
	var ret1 string

	cret = C.g_get_tmp_dir()

	ret1 = C.GoString(cret)

	return ret1
}

// GetUserCacheDir returns a base directory in which to store non-essential,
// cached data specific to particular user.
//
// On UNIX platforms this is determined using the mechanisms described in the
// XDG Base Directory Specification
// (http://www.freedesktop.org/Standards/basedir-spec). In this case the
// directory retrieved will be `XDG_CACHE_HOME`.
//
// On Windows it follows XDG Base Directory Specification if `XDG_CACHE_HOME` is
// defined. If `XDG_CACHE_HOME` is undefined, the directory that serves as a
// common repository for temporary Internet files is used instead. A typical
// path is `C:\Documents and Settings\username\Local Settings\Temporary Internet
// Files`. See the documentation for `CSIDL_INTERNET_CACHE`
// (https://msdn.microsoft.com/en-us/library/windows/desktop/bb76249428v=vs.8529.aspx#csidl_internet_cache).
func GetUserCacheDir() string {
	var cret *C.gchar
	var ret1 string

	cret = C.g_get_user_cache_dir()

	ret1 = C.GoString(cret)

	return ret1
}

// GetUserConfigDir returns a base directory in which to store user-specific
// application configuration information such as user preferences and settings.
//
// On UNIX platforms this is determined using the mechanisms described in the
// XDG Base Directory Specification
// (http://www.freedesktop.org/Standards/basedir-spec). In this case the
// directory retrieved will be `XDG_CONFIG_HOME`.
//
// On Windows it follows XDG Base Directory Specification if `XDG_CONFIG_HOME`
// is defined. If `XDG_CONFIG_HOME` is undefined, the folder to use for local
// (as opposed to roaming) application data is used instead. See the
// documentation for `CSIDL_LOCAL_APPDATA`
// (https://msdn.microsoft.com/en-us/library/windows/desktop/bb76249428v=vs.8529.aspx#csidl_local_appdata).
// Note that in this case on Windows it will be the same as what
// g_get_user_data_dir() returns.
func GetUserConfigDir() string {
	var cret *C.gchar
	var ret1 string

	cret = C.g_get_user_config_dir()

	ret1 = C.GoString(cret)

	return ret1
}

// GetUserDataDir returns a base directory in which to access application data
// such as icons that is customized for a particular user.
//
// On UNIX platforms this is determined using the mechanisms described in the
// XDG Base Directory Specification
// (http://www.freedesktop.org/Standards/basedir-spec). In this case the
// directory retrieved will be `XDG_DATA_HOME`.
//
// On Windows it follows XDG Base Directory Specification if `XDG_DATA_HOME` is
// defined. If `XDG_DATA_HOME` is undefined, the folder to use for local (as
// opposed to roaming) application data is used instead. See the documentation
// for `CSIDL_LOCAL_APPDATA`
// (https://msdn.microsoft.com/en-us/library/windows/desktop/bb76249428v=vs.8529.aspx#csidl_local_appdata).
// Note that in this case on Windows it will be the same as what
// g_get_user_config_dir() returns.
func GetUserDataDir() string {
	var cret *C.gchar
	var ret1 string

	cret = C.g_get_user_data_dir()

	ret1 = C.GoString(cret)

	return ret1
}

// GetUserName gets the user name of the current user. The encoding of the
// returned string is system-defined. On UNIX, it might be the preferred file
// name encoding, or something else, and there is no guarantee that it is even
// consistent on a machine. On Windows, it is always UTF-8.
func GetUserName() string {
	var cret *C.gchar
	var ret1 string

	cret = C.g_get_user_name()

	ret1 = C.GoString(cret)

	return ret1
}

// GetUserRuntimeDir returns a directory that is unique to the current user on
// the local system.
//
// This is determined using the mechanisms described in the XDG Base Directory
// Specification (http://www.freedesktop.org/Standards/basedir-spec). This is
// the directory specified in the `XDG_RUNTIME_DIR` environment variable. In the
// case that this variable is not set, we return the value of
// g_get_user_cache_dir(), after verifying that it exists.
func GetUserRuntimeDir() string {
	var cret *C.gchar
	var ret1 string

	cret = C.g_get_user_runtime_dir()

	ret1 = C.GoString(cret)

	return ret1
}

// GetUserSpecialDir returns the full path of a special directory using its
// logical id.
//
// On UNIX this is done using the XDG special user directories. For
// compatibility with existing practise, G_USER_DIRECTORY_DESKTOP falls back to
// `$HOME/Desktop` when XDG special user directories have not been set up.
//
// Depending on the platform, the user might be able to change the path of the
// special directory without requiring the session to restart; GLib will not
// reflect any change once the special directories are loaded.
func GetUserSpecialDir(directory UserDirectory) string {
	var arg1 C.GUserDirectory

	arg1 = (C.GUserDirectory)(directory)

	var cret *C.gchar
	var ret1 string

	cret = C.g_get_user_special_dir(directory)

	ret1 = C.GoString(cret)

	return ret1
}

// NullifyPointer: set the pointer at the specified location to nil.
func NullifyPointer(nullifyLocation interface{}) {
	var arg1 *C.gpointer

	arg1 = *C.gpointer(nullifyLocation)

	C.g_nullify_pointer(nullifyLocation)
}

// ParseDebugString parses a string containing debugging options into a guint
// containing bit flags. This is used within GDK and GTK+ to parse the debug
// options passed on the command line or through environment variables.
//
// If @string is equal to "all", all flags are set. Any flags specified along
// with "all" in @string are inverted; thus, "all,foo,bar" or "foo,bar,all" sets
// all flags except those corresponding to "foo" and "bar".
//
// If @string is equal to "help", all the available keys in @keys are printed
// out to standard error.
func ParseDebugString(string string, keys []DebugKey) uint {

	var cret C.guint
	var ret1 uint

	cret = C.g_parse_debug_string(string, keys, nkeys)

	ret1 = C.guint(cret)

	return ret1
}

// ReloadUserSpecialDirsCache resets the cache used for
// g_get_user_special_dir(), so that the latest on-disk version is used. Call
// this only if you just changed the data on disk yourself.
//
// Due to thread safety issues this may cause leaking of strings that were
// previously returned from g_get_user_special_dir() that can't be freed. We
// ensure to only leak the data for the directories that actually changed value
// though.
func ReloadUserSpecialDirsCache() {
	C.g_reload_user_special_dirs_cache()
}

// SetApplicationName sets a human-readable name for the application. This name
// should be localized if possible, and is intended for display to the user.
// Contrast with g_set_prgname(), which sets a non-localized name.
// g_set_prgname() will be called automatically by gtk_init(), but
// g_set_application_name() will not.
//
// Note that for thread safety reasons, this function can only be called once.
//
// The application name will be used in contexts such as error messages, or when
// displaying an application's name in the task list.
func SetApplicationName(applicationName string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(applicationName))
	defer C.free(unsafe.Pointer(arg1))

	C.g_set_application_name(applicationName)
}

// SetPrgname sets the name of the program. This name should not be localized,
// in contrast to g_set_application_name().
//
// If you are using #GApplication the program name is set in
// g_application_run(). In case of GDK or GTK+ it is set in gdk_init(), which is
// called by gtk_init() and the Application::startup handler. The program name
// is found by taking the last component of @argv[0].
//
// Note that for thread-safety reasons this function can only be called once.
func SetPrgname(prgname string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(prgname))
	defer C.free(unsafe.Pointer(arg1))

	C.g_set_prgname(prgname)
}

// DebugKey associates a string with a bit flag. Used in g_parse_debug_string().
type DebugKey struct {
	native C.GDebugKey
}

// WrapDebugKey wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDebugKey(ptr unsafe.Pointer) *DebugKey {
	if ptr == nil {
		return nil
	}

	return (*DebugKey)(ptr)
}

func marshalDebugKey(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDebugKey(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DebugKey) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// Key gets the field inside the struct.
func (d *DebugKey) Key() string {
	v = C.GoString(d.native.key)
}

// Value gets the field inside the struct.
func (d *DebugKey) Value() uint {
	v = C.guint(d.native.value)
}
