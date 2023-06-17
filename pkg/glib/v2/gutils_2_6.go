// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// GetSystemConfigDirs returns an ordered list of base directories in which to
// access system-wide configuration information.
//
// On UNIX platforms this is determined using the mechanisms
// described in the XDG Base Directory Specification
// (http://www.freedesktop.org/Standards/basedir-spec). In this case the list of
// directories retrieved will be XDG_CONFIG_DIRS.
//
// On Windows it follows XDG Base Directory Specification if XDG_CONFIG_DIRS
// is defined. If XDG_CONFIG_DIRS is undefined, the directory that contains
// application data for all users is used instead. A typical path is
// C:\Documents and Settings\All Users\Application Data. This folder is used for
// application data that is not user specific. For example, an application can
// store a spell-check dictionary, a database of clip art, or a log file in the
// CSIDL_COMMON_APPDATA folder. This information will not roam and is available
// to anyone using the computer.
//
// The return value is cached and modifying it at runtime is not supported,
// as it’s not thread-safe to modify environment variables at runtime.
//
// The function returns the following values:
//
//   - filenames: a NULL-terminated array of strings owned by GLib that must not
//     be modified or freed.
//
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
// On UNIX platforms this is determined using the mechanisms
// described in the XDG Base Directory Specification
// (http://www.freedesktop.org/Standards/basedir-spec) In this case the list of
// directories retrieved will be XDG_DATA_DIRS.
//
// On Windows it follows XDG Base Directory Specification if XDG_DATA_DIRS is
// defined. If XDG_DATA_DIRS is undefined, the first elements in the list are
// the Application Data and Documents folders for All Users. (These can be
// determined only on Windows 2000 or later and are not present in the list
// on other Windows versions.) See documentation for CSIDL_COMMON_APPDATA and
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
// The return value is cached and modifying it at runtime is not supported,
// as it’s not thread-safe to modify environment variables at runtime.
//
// The function returns the following values:
//
//   - filenames: a NULL-terminated array of strings owned by GLib that must not
//     be modified or freed.
//
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

// GetUserCacheDir returns a base directory in which to store non-essential,
// cached data specific to particular user.
//
// On UNIX platforms this is determined using the mechanisms
// described in the XDG Base Directory Specification
// (http://www.freedesktop.org/Standards/basedir-spec). In this case the
// directory retrieved will be XDG_CACHE_HOME.
//
// On Windows it follows XDG Base Directory Specification if XDG_CACHE_HOME
// is defined. If XDG_CACHE_HOME is undefined, the directory that serves
// as a common repository for temporary Internet files is used instead.
// A typical path is C:\Documents and Settings\username\Local Settings\Temporary
// Internet Files. See the documentation for CSIDL_INTERNET_CACHE
// (https://msdn.microsoft.com/en-us/library/windows/desktop/bb76249428v=vs.8529.aspx#csidl_internet_cache).
//
// The return value is cached and modifying it at runtime is not supported,
// as it’s not thread-safe to modify environment variables at runtime.
//
// The function returns the following values:
//
//   - filename: string owned by GLib that must not be modified or freed.
//
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
// On UNIX platforms this is determined using the mechanisms
// described in the XDG Base Directory Specification
// (http://www.freedesktop.org/Standards/basedir-spec). In this case the
// directory retrieved will be XDG_CONFIG_HOME.
//
// On Windows it follows XDG Base Directory Specification if
// XDG_CONFIG_HOME is defined. If XDG_CONFIG_HOME is undefined,
// the folder to use for local (as opposed to roaming) application
// data is used instead. See the documentation for CSIDL_LOCAL_APPDATA
// (https://msdn.microsoft.com/en-us/library/windows/desktop/bb76249428v=vs.8529.aspx#csidl_local_appdata).
// Note that in this case on Windows it will be the same as what
// g_get_user_data_dir() returns.
//
// The return value is cached and modifying it at runtime is not supported,
// as it’s not thread-safe to modify environment variables at runtime.
//
// The function returns the following values:
//
//   - filename: string owned by GLib that must not be modified or freed.
//
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
// On UNIX platforms this is determined using the mechanisms
// described in the XDG Base Directory Specification
// (http://www.freedesktop.org/Standards/basedir-spec). In this case the
// directory retrieved will be XDG_DATA_HOME.
//
// On Windows it follows XDG Base Directory Specification if
// XDG_DATA_HOME is defined. If XDG_DATA_HOME is undefined,
// the folder to use for local (as opposed to roaming) application
// data is used instead. See the documentation for CSIDL_LOCAL_APPDATA
// (https://msdn.microsoft.com/en-us/library/windows/desktop/bb76249428v=vs.8529.aspx#csidl_local_appdata).
// Note that in this case on Windows it will be the same as what
// g_get_user_config_dir() returns.
//
// The return value is cached and modifying it at runtime is not supported,
// as it’s not thread-safe to modify environment variables at runtime.
//
// The function returns the following values:
//
//   - filename: string owned by GLib that must not be modified or freed.
//
func GetUserDataDir() string {
	var _cret *C.gchar // in

	_cret = C.g_get_user_data_dir()

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}
