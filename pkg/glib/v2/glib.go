// Code generated by girgen. DO NOT EDIT.

package glib

import (
	_ "runtime/cgo"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib.h>
import "C"

// The function returns the following values:
//
func ConvertErrorQuark() Quark {
	var _cret C.GQuark // in

	_cret = C.g_convert_error_quark()

	var _quark Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func FileErrorQuark() Quark {
	var _cret C.GQuark // in

	_cret = C.g_file_error_quark()

	var _quark Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func MarkupErrorQuark() Quark {
	var _cret C.GQuark // in

	_cret = C.g_markup_error_quark()

	var _quark Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func NumberParserErrorQuark() Quark {
	var _cret C.GQuark // in

	_cret = C.g_number_parser_error_quark()

	var _quark Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func OptionErrorQuark() Quark {
	var _cret C.GQuark // in

	_cret = C.g_option_error_quark()

	var _quark Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func ShellErrorQuark() Quark {
	var _cret C.GQuark // in

	_cret = C.g_shell_error_quark()

	var _quark Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func SpawnErrorQuark() Quark {
	var _cret C.GQuark // in

	_cret = C.g_spawn_error_quark()

	var _quark Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func SpawnExitErrorQuark() Quark {
	var _cret C.GQuark // in

	_cret = C.g_spawn_exit_error_quark()

	var _quark Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func IOChannelErrorQuark() Quark {
	var _cret C.GQuark // in

	_cret = C.g_io_channel_error_quark()

	var _quark Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func KeyFileErrorQuark() Quark {
	var _cret C.GQuark // in

	_cret = C.g_key_file_error_quark()

	var _quark Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func RegexErrorQuark() Quark {
	var _cret C.GQuark // in

	_cret = C.g_regex_error_quark()

	var _quark Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func URIErrorQuark() Quark {
	var _cret C.GQuark // in

	_cret = C.g_uri_error_quark()

	var _quark Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func VariantParseErrorQuark() Quark {
	var _cret C.GQuark // in

	_cret = C.g_variant_parse_error_quark()

	var _quark Quark // out

	_quark = uint32(_cret)

	return _quark
}
