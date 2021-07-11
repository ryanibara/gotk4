// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib.h>
import "C"

// ShellError: error codes returned by shell functions.
type ShellError int

const (
	// BadQuoting: mismatched or otherwise mangled quoting.
	ShellErrorBadQuoting ShellError = iota
	// EmptyString: string to be parsed was empty.
	ShellErrorEmptyString
	// Failed: some other error.
	ShellErrorFailed
)

// ShellParseArgv parses a command line into an argument vector, in much the
// same way the shell would, but without many of the expansions the shell would
// perform (variable expansion, globs, operators, filename expansion, etc. are
// not supported). The results are defined to be the same as those you would get
// from a UNIX98 /bin/sh, as long as the input contains none of the unsupported
// shell expansions. If the input does contain such expansions, they are passed
// through literally. Possible errors are those from the SHELL_ERROR domain.
// Free the returned vector with g_strfreev().
func ShellParseArgv(commandLine string) ([]string, error) {
	var _arg1 *C.gchar // out
	var _arg3 **C.gchar
	var _arg2 C.gint    // in
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(commandLine)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_shell_parse_argv(_arg1, &_arg2, &_arg3, &_cerr)

	var _argvp []string
	var _goerr error // out

	defer C.free(unsafe.Pointer(_arg3))
	{
		src := unsafe.Slice(_arg3, _arg2)
		_argvp = make([]string, _arg2)
		for i := 0; i < int(_arg2); i++ {
			_argvp[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _argvp, _goerr
}

// ShellQuote quotes a string so that the shell (/bin/sh) will interpret the
// quoted string to mean @unquoted_string. If you pass a filename to the shell,
// for example, you should first quote it with this function. The return value
// must be freed with g_free(). The quoting style used is undefined (single or
// double quotes may be used).
func ShellQuote(unquotedString string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(unquotedString)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_shell_quote(_arg1)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// ShellUnquote unquotes a string as the shell (/bin/sh) would. Only handles
// quotes; if a string contains file globs, arithmetic operators, variables,
// backticks, redirections, or other special-to-the-shell features, the result
// will be different from the result a real shell would produce (the variables,
// backticks, etc. will be passed through literally instead of being expanded).
// This function is guaranteed to succeed if applied to the result of
// g_shell_quote(). If it fails, it returns nil and sets the error. The
// @quoted_string need not actually contain quoted or escaped text;
// g_shell_unquote() simply goes through the string and unquotes/unescapes
// anything that the shell would. Both single and double quotes are handled, as
// are escapes including escaped newlines. The return value must be freed with
// g_free(). Possible errors are in the SHELL_ERROR domain.
//
// Shell quoting rules are a bit strange. Single quotes preserve the literal
// string exactly. escape sequences are not allowed; not even \' - if you want a
// ' in the quoted text, you have to do something like 'foo'\”bar'. Double
// quotes allow $, `, ", \, and newline to be escaped with backslash. Otherwise
// double quotes preserve things literally.
func ShellUnquote(quotedString string) (string, error) {
	var _arg1 *C.gchar  // out
	var _cret *C.gchar  // in
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(quotedString)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_shell_unquote(_arg1, &_cerr)

	var _filename string // out
	var _goerr error     // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _filename, _goerr
}
