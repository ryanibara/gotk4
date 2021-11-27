// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib.h>
import "C"

// UserDirectory: these are logical ids for special directories which are
// defined depending on the platform used. You should use
// g_get_user_special_dir() to retrieve the full path associated to the logical
// id.
//
// The Directory enumeration can be extended at later date. Not every platform
// has a directory for every logical id in this enumeration.
type UserDirectory C.gint

const (
	// UserDirectoryDesktop user's Desktop directory.
	UserDirectoryDesktop UserDirectory = iota
	// UserDirectoryDocuments user's Documents directory.
	UserDirectoryDocuments
	// UserDirectoryDownload user's Downloads directory.
	UserDirectoryDownload
	// UserDirectoryMusic user's Music directory.
	UserDirectoryMusic
	// UserDirectoryPictures user's Pictures directory.
	UserDirectoryPictures
	// UserDirectoryPublicShare user's shared directory.
	UserDirectoryPublicShare
	// UserDirectoryTemplates user's Templates directory.
	UserDirectoryTemplates
	// UserDirectoryVideos user's Movies directory.
	UserDirectoryVideos
	// UserNDirectories: number of enum values.
	UserNDirectories
)

// String returns the name in string for UserDirectory.
func (u UserDirectory) String() string {
	switch u {
	case UserDirectoryDesktop:
		return "DirectoryDesktop"
	case UserDirectoryDocuments:
		return "DirectoryDocuments"
	case UserDirectoryDownload:
		return "DirectoryDownload"
	case UserDirectoryMusic:
		return "DirectoryMusic"
	case UserDirectoryPictures:
		return "DirectoryPictures"
	case UserDirectoryPublicShare:
		return "DirectoryPublicShare"
	case UserDirectoryTemplates:
		return "DirectoryTemplates"
	case UserDirectoryVideos:
		return "DirectoryVideos"
	case UserNDirectories:
		return "NDirectories"
	default:
		return fmt.Sprintf("UserDirectory(%d)", u)
	}
}

// FormatSizeFlags flags to modify the format of the string returned by
// g_format_size_full().
type FormatSizeFlags C.guint

const (
	// FormatSizeDefault: behave the same as g_format_size().
	FormatSizeDefault FormatSizeFlags = 0b0
	// FormatSizeLongFormat: include the exact number of bytes as part of the
	// returned string. For example, "45.6 kB (45,612 bytes)".
	FormatSizeLongFormat FormatSizeFlags = 0b1
	// FormatSizeIecUnits: use IEC (base 1024) units with "KiB"-style suffixes.
	// IEC units should only be used for reporting things with a strong "power
	// of 2" basis, like RAM sizes or RAID stripe sizes. Network and storage
	// sizes should be reported in the normal SI units.
	FormatSizeIecUnits FormatSizeFlags = 0b10
	// FormatSizeBits: set the size as a quantity in bits, rather than bytes,
	// and return units in bits. For example, ‘Mb’ rather than ‘MB’.
	FormatSizeBits FormatSizeFlags = 0b100
)

// String returns the names in string for FormatSizeFlags.
func (f FormatSizeFlags) String() string {
	if f == 0 {
		return "FormatSizeFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(72)

	for f != 0 {
		next := f & (f - 1)
		bit := f - next

		switch bit {
		case FormatSizeDefault:
			builder.WriteString("Default|")
		case FormatSizeLongFormat:
			builder.WriteString("LongFormat|")
		case FormatSizeIecUnits:
			builder.WriteString("IecUnits|")
		case FormatSizeBits:
			builder.WriteString("Bits|")
		default:
			builder.WriteString(fmt.Sprintf("FormatSizeFlags(0b%b)|", bit))
		}

		f = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if f contains other.
func (f FormatSizeFlags) Has(other FormatSizeFlags) bool {
	return (f & other) == other
}

// BitNthLSF: find the position of the first bit set in mask, searching from
// (but not including) nth_bit upwards. Bits are numbered from 0 (least
// significant) to sizeof(#gulong) * 8 - 1 (31 or 63, usually). To start
// searching from the 0th bit, set nth_bit to -1.
//
// The function takes the following parameters:
//
//    - mask containing flags.
//    - nthBit: index of the bit to start the search from.
//
func BitNthLSF(mask uint32, nthBit int) int {
	var _arg1 C.gulong // out
	var _arg2 C.gint   // out
	var _cret C.gint   // in

	_arg1 = C.gulong(mask)
	_arg2 = C.gint(nthBit)

	_cret = C.g_bit_nth_lsf(_arg1, _arg2)
	runtime.KeepAlive(mask)
	runtime.KeepAlive(nthBit)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// BitNthMSF: find the position of the first bit set in mask, searching from
// (but not including) nth_bit downwards. Bits are numbered from 0 (least
// significant) to sizeof(#gulong) * 8 - 1 (31 or 63, usually). To start
// searching from the last bit, set nth_bit to -1 or GLIB_SIZEOF_LONG * 8.
//
// The function takes the following parameters:
//
//    - mask containing flags.
//    - nthBit: index of the bit to start the search from.
//
func BitNthMSF(mask uint32, nthBit int) int {
	var _arg1 C.gulong // out
	var _arg2 C.gint   // out
	var _cret C.gint   // in

	_arg1 = C.gulong(mask)
	_arg2 = C.gint(nthBit)

	_cret = C.g_bit_nth_msf(_arg1, _arg2)
	runtime.KeepAlive(mask)
	runtime.KeepAlive(nthBit)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// BitStorage gets the number of bits used to hold number, e.g. if number is 4,
// 3 bits are needed.
//
// The function takes the following parameters:
//
//    - number: #guint.
//
func BitStorage(number uint32) uint {
	var _arg1 C.gulong // out
	var _cret C.guint  // in

	_arg1 = C.gulong(number)

	_cret = C.g_bit_storage(_arg1)
	runtime.KeepAlive(number)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// FindProgramInPath locates the first executable named program in the user's
// path, in the same way that execvp() would locate it. Returns an allocated
// string with the absolute path name, or NULL if the program is not found in
// the path. If program is already an absolute path, returns a copy of program
// if program exists and is executable, and NULL otherwise. On Windows, if
// program does not have a file type suffix, tries with the suffixes .exe, .cmd,
// .bat and .com, and the suffixes in the PATHEXT environment variable.
//
// On Windows, it looks for the file in the same way as CreateProcess() would.
// This means first in the directory where the executing program was loaded
// from, then in the current directory, then in the Windows 32-bit system
// directory, then in the Windows directory, and finally in the directories in
// the PATH environment variable. If the program is found, the return value
// contains the full name including the type suffix.
//
// The function takes the following parameters:
//
//    - program name in the GLib file name encoding.
//
func FindProgramInPath(program string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(program)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_find_program_in_path(_arg1)
	runtime.KeepAlive(program)

	var _filename string // out

	if _cret != nil {
		_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _filename
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
//
// The function takes the following parameters:
//
//    - size in bytes.
//
func FormatSize(size uint64) string {
	var _arg1 C.guint64 // out
	var _cret *C.gchar  // in

	_arg1 = C.guint64(size)

	_cret = C.g_format_size(_arg1)
	runtime.KeepAlive(size)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FormatSizeForDisplay formats a size (for example the size of a file) into a
// human readable string. Sizes are rounded to the nearest size prefix (KB, MB,
// GB) and are displayed rounded to the nearest tenth. E.g. the file size
// 3292528 bytes will be converted into the string "3.1 MB".
//
// The prefix units base is 1024 (i.e. 1 KB is 1024 bytes).
//
// This string should be freed with g_free() when not needed any longer.
//
// Deprecated: This function is broken due to its use of SI suffixes to denote
// IEC units. Use g_format_size() instead.
//
// The function takes the following parameters:
//
//    - size in bytes.
//
func FormatSizeForDisplay(size int64) string {
	var _arg1 C.goffset // out
	var _cret *C.gchar  // in

	_arg1 = C.goffset(size)

	_cret = C.g_format_size_for_display(_arg1)
	runtime.KeepAlive(size)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FormatSizeFull formats a size.
//
// This function is similar to g_format_size() but allows for flags that modify
// the output. See SizeFlags.
//
// The function takes the following parameters:
//
//    - size in bytes.
//    - flags to modify the output.
//
func FormatSizeFull(size uint64, flags FormatSizeFlags) string {
	var _arg1 C.guint64          // out
	var _arg2 C.GFormatSizeFlags // out
	var _cret *C.gchar           // in

	_arg1 = C.guint64(size)
	_arg2 = C.GFormatSizeFlags(flags)

	_cret = C.g_format_size_full(_arg1, _arg2)
	runtime.KeepAlive(size)
	runtime.KeepAlive(flags)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// GetApplicationName gets a human-readable name for the application, as set by
// g_set_application_name(). This name should be localized if possible, and is
// intended for display to the user. Contrast with g_get_prgname(), which gets a
// non-localized name. If g_set_application_name() has not been called, returns
// the result of g_get_prgname() (which may be NULL if g_set_prgname() has also
// not been called).
func GetApplicationName() string {
	var _cret *C.gchar // in

	_cret = C.g_get_application_name()

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// GetHomeDir gets the current user's home directory.
//
// As with most UNIX tools, this function will return the value of the HOME
// environment variable if it is set to an existing absolute path name, falling
// back to the passwd file in the case that it is unset.
//
// If the path given in HOME is non-absolute, does not exist, or is not a
// directory, the result is undefined.
//
// Before version 2.36 this function would ignore the HOME environment variable,
// taking the value from the passwd database instead. This was changed to
// increase the compatibility of GLib with other programs (and the XDG basedir
// specification) and to increase testability of programs based on GLib (by
// making it easier to run them from test frameworks).
//
// If your program has a strong requirement for either the new or the old
// behaviour (and if you don't wish to increase your GLib dependency to ensure
// that the new behaviour is in effect) then you should either directly check
// the HOME environment variable yourself or unset it before calling any
// functions in GLib.
func GetHomeDir() string {
	var _cret *C.gchar // in

	_cret = C.g_get_home_dir()

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
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
	var _cret *C.gchar // in

	_cret = C.g_get_host_name()

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// GetOsInfo: get information about the operating system.
//
// On Linux this comes from the /etc/os-release file. On other systems, it may
// come from a variety of sources. You can either use the standard key names
// like G_OS_INFO_KEY_NAME or pass any UTF-8 string key name. For example,
// /etc/os-release provides a number of other less commonly used values that may
// be useful. No key is guaranteed to be provided, so the caller should always
// check if the result is NULL.
//
// The function takes the following parameters:
//
//    - keyName: key for the OS info being requested, for example
//    G_OS_INFO_KEY_NAME.
//
func GetOsInfo(keyName string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(keyName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_get_os_info(_arg1)
	runtime.KeepAlive(keyName)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// GetPrgname gets the name of the program. This name should not be localized,
// in contrast to g_get_application_name().
//
// If you are using #GApplication the program name is set in
// g_application_run(). In case of GDK or GTK+ it is set in gdk_init(), which is
// called by gtk_init() and the Application::startup handler. The program name
// is found by taking the last component of argv[0].
func GetPrgname() string {
	var _cret *C.gchar // in

	_cret = C.g_get_prgname()

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// GetRealName gets the real name of the user. This usually comes from the
// user's entry in the passwd file. The encoding of the returned string is
// system-defined. (On Windows, it is, however, always UTF-8.) If the real user
// name cannot be determined, the string "Unknown" is returned.
func GetRealName() string {
	var _cret *C.gchar // in

	_cret = C.g_get_real_name()

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// GetSystemConfigDirs returns an ordered list of base directories in which to
// access system-wide configuration information.
//
// On UNIX platforms this is determined using the mechanisms described in the
// XDG Base Directory Specification
// (http://www.freedesktop.org/Standards/basedir-spec). In this case the list of
// directories retrieved will be XDG_CONFIG_DIRS.
//
// On Windows it follows XDG Base Directory Specification if XDG_CONFIG_DIRS is
// defined. If XDG_CONFIG_DIRS is undefined, the directory that contains
// application data for all users is used instead. A typical path is
// C:\Documents and Settings\All Users\Application Data. This folder is used for
// application data that is not user specific. For example, an application can
// store a spell-check dictionary, a database of clip art, or a log file in the
// CSIDL_COMMON_APPDATA folder. This information will not roam and is available
// to anyone using the computer.
//
// The return value is cached and modifying it at runtime is not supported, as
// it’s not thread-safe to modify environment variables at runtime.
func GetSystemConfigDirs() []string {
	var _cret **C.gchar // in

	_cret = C.g_get_system_config_dirs()

	var _filenames []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_filenames = make([]string, i)
		for i := range src {
			_filenames[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _filenames
}

// GetSystemDataDirs returns an ordered list of base directories in which to
// access system-wide application data.
//
// On UNIX platforms this is determined using the mechanisms described in the
// XDG Base Directory Specification
// (http://www.freedesktop.org/Standards/basedir-spec) In this case the list of
// directories retrieved will be XDG_DATA_DIRS.
//
// On Windows it follows XDG Base Directory Specification if XDG_DATA_DIRS is
// defined. If XDG_DATA_DIRS is undefined, the first elements in the list are
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
//
// The return value is cached and modifying it at runtime is not supported, as
// it’s not thread-safe to modify environment variables at runtime.
func GetSystemDataDirs() []string {
	var _cret **C.gchar // in

	_cret = C.g_get_system_data_dirs()

	var _filenames []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_filenames = make([]string, i)
		for i := range src {
			_filenames[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _filenames
}

// GetTmpDir gets the directory to use for temporary files.
//
// On UNIX, this is taken from the TMPDIR environment variable. If the variable
// is not set, P_tmpdir is used, as defined by the system C library. Failing
// that, a hard-coded default of "/tmp" is returned.
//
// On Windows, the TEMP environment variable is used, with the root directory of
// the Windows installation (eg: "C:\") used as a default.
//
// The encoding of the returned string is system-defined. On Windows, it is
// always UTF-8. The return value is never NULL or the empty string.
func GetTmpDir() string {
	var _cret *C.gchar // in

	_cret = C.g_get_tmp_dir()

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// GetUserCacheDir returns a base directory in which to store non-essential,
// cached data specific to particular user.
//
// On UNIX platforms this is determined using the mechanisms described in the
// XDG Base Directory Specification
// (http://www.freedesktop.org/Standards/basedir-spec). In this case the
// directory retrieved will be XDG_CACHE_HOME.
//
// On Windows it follows XDG Base Directory Specification if XDG_CACHE_HOME is
// defined. If XDG_CACHE_HOME is undefined, the directory that serves as a
// common repository for temporary Internet files is used instead. A typical
// path is C:\Documents and Settings\username\Local Settings\Temporary Internet
// Files. See the documentation for CSIDL_INTERNET_CACHE
// (https://msdn.microsoft.com/en-us/library/windows/desktop/bb76249428v=vs.8529.aspx#csidl_internet_cache).
//
// The return value is cached and modifying it at runtime is not supported, as
// it’s not thread-safe to modify environment variables at runtime.
func GetUserCacheDir() string {
	var _cret *C.gchar // in

	_cret = C.g_get_user_cache_dir()

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// GetUserConfigDir returns a base directory in which to store user-specific
// application configuration information such as user preferences and settings.
//
// On UNIX platforms this is determined using the mechanisms described in the
// XDG Base Directory Specification
// (http://www.freedesktop.org/Standards/basedir-spec). In this case the
// directory retrieved will be XDG_CONFIG_HOME.
//
// On Windows it follows XDG Base Directory Specification if XDG_CONFIG_HOME is
// defined. If XDG_CONFIG_HOME is undefined, the folder to use for local (as
// opposed to roaming) application data is used instead. See the documentation
// for CSIDL_LOCAL_APPDATA
// (https://msdn.microsoft.com/en-us/library/windows/desktop/bb76249428v=vs.8529.aspx#csidl_local_appdata).
// Note that in this case on Windows it will be the same as what
// g_get_user_data_dir() returns.
//
// The return value is cached and modifying it at runtime is not supported, as
// it’s not thread-safe to modify environment variables at runtime.
func GetUserConfigDir() string {
	var _cret *C.gchar // in

	_cret = C.g_get_user_config_dir()

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// GetUserDataDir returns a base directory in which to access application data
// such as icons that is customized for a particular user.
//
// On UNIX platforms this is determined using the mechanisms described in the
// XDG Base Directory Specification
// (http://www.freedesktop.org/Standards/basedir-spec). In this case the
// directory retrieved will be XDG_DATA_HOME.
//
// On Windows it follows XDG Base Directory Specification if XDG_DATA_HOME is
// defined. If XDG_DATA_HOME is undefined, the folder to use for local (as
// opposed to roaming) application data is used instead. See the documentation
// for CSIDL_LOCAL_APPDATA
// (https://msdn.microsoft.com/en-us/library/windows/desktop/bb76249428v=vs.8529.aspx#csidl_local_appdata).
// Note that in this case on Windows it will be the same as what
// g_get_user_config_dir() returns.
//
// The return value is cached and modifying it at runtime is not supported, as
// it’s not thread-safe to modify environment variables at runtime.
func GetUserDataDir() string {
	var _cret *C.gchar // in

	_cret = C.g_get_user_data_dir()

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// GetUserName gets the user name of the current user. The encoding of the
// returned string is system-defined. On UNIX, it might be the preferred file
// name encoding, or something else, and there is no guarantee that it is even
// consistent on a machine. On Windows, it is always UTF-8.
func GetUserName() string {
	var _cret *C.gchar // in

	_cret = C.g_get_user_name()

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// GetUserRuntimeDir returns a directory that is unique to the current user on
// the local system.
//
// This is determined using the mechanisms described in the XDG Base Directory
// Specification (http://www.freedesktop.org/Standards/basedir-spec). This is
// the directory specified in the XDG_RUNTIME_DIR environment variable. In the
// case that this variable is not set, we return the value of
// g_get_user_cache_dir(), after verifying that it exists.
//
// The return value is cached and modifying it at runtime is not supported, as
// it’s not thread-safe to modify environment variables at runtime.
func GetUserRuntimeDir() string {
	var _cret *C.gchar // in

	_cret = C.g_get_user_runtime_dir()

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// GetUserSpecialDir returns the full path of a special directory using its
// logical id.
//
// On UNIX this is done using the XDG special user directories. For
// compatibility with existing practise, G_USER_DIRECTORY_DESKTOP falls back to
// $HOME/Desktop when XDG special user directories have not been set up.
//
// Depending on the platform, the user might be able to change the path of the
// special directory without requiring the session to restart; GLib will not
// reflect any change once the special directories are loaded.
//
// The function takes the following parameters:
//
//    - directory: logical id of special directory.
//
func GetUserSpecialDir(directory UserDirectory) string {
	var _arg1 C.GUserDirectory // out
	var _cret *C.gchar         // in

	_arg1 = C.GUserDirectory(directory)

	_cret = C.g_get_user_special_dir(_arg1)
	runtime.KeepAlive(directory)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// ParseDebugString parses a string containing debugging options into a guint
// containing bit flags. This is used within GDK and GTK+ to parse the debug
// options passed on the command line or through environment variables.
//
// If string is equal to "all", all flags are set. Any flags specified along
// with "all" in string are inverted; thus, "all,foo,bar" or "foo,bar,all" sets
// all flags except those corresponding to "foo" and "bar".
//
// If string is equal to "help", all the available keys in keys are printed out
// to standard error.
//
// The function takes the following parameters:
//
//    - str: list of debug options separated by colons, spaces, or commas, or
//    NULL.
//    - keys: pointer to an array of Key which associate strings with bit
//    flags.
//
func ParseDebugString(str string, keys []DebugKey) uint {
	var _arg1 *C.gchar     // out
	var _arg2 *C.GDebugKey // out
	var _arg3 C.guint
	var _cret C.guint // in

	if str != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg3 = (C.guint)(len(keys))
	_arg2 = (*C.GDebugKey)(C.calloc(C.size_t(len(keys)), C.size_t(C.sizeof_GDebugKey)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GDebugKey)(_arg2), len(keys))
		for i := range keys {
			out[i] = *(*C.GDebugKey)(gextras.StructNative(unsafe.Pointer((&keys[i]))))
		}
	}

	_cret = C.g_parse_debug_string(_arg1, _arg2, _arg3)
	runtime.KeepAlive(str)
	runtime.KeepAlive(keys)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
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
//
// The function takes the following parameters:
//
//    - applicationName: localized name of the application.
//
func SetApplicationName(applicationName string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(applicationName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_set_application_name(_arg1)
	runtime.KeepAlive(applicationName)
}

// SetPrgname sets the name of the program. This name should not be localized,
// in contrast to g_set_application_name().
//
// If you are using #GApplication the program name is set in
// g_application_run(). In case of GDK or GTK+ it is set in gdk_init(), which is
// called by gtk_init() and the Application::startup handler. The program name
// is found by taking the last component of argv[0].
//
// Note that for thread-safety reasons this function can only be called once.
//
// The function takes the following parameters:
//
//    - prgname: name of the program.
//
func SetPrgname(prgname string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(prgname)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_set_prgname(_arg1)
	runtime.KeepAlive(prgname)
}

// DebugKey associates a string with a bit flag. Used in g_parse_debug_string().
//
// An instance of this type is always passed by reference.
type DebugKey struct {
	*debugKey
}

// debugKey is the struct that's finalized.
type debugKey struct {
	native *C.GDebugKey
}

// Key: string.
func (d *DebugKey) Key() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(d.native.key)))
	return v
}

// Value: flag.
func (d *DebugKey) Value() uint {
	var v uint // out
	v = uint(d.native.value)
	return v
}
