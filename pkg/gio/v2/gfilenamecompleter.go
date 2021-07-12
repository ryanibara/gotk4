// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_filename_completer_get_type()), F: marshalFilenameCompleterer},
	})
}

// FilenameCompleterOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FilenameCompleterOverrider interface {
	GotCompletionData()
}

// FilenameCompleterer describes FilenameCompleter's methods.
type FilenameCompleterer interface {
	// CompletionSuffix obtains a completion for @initial_text from @completer.
	CompletionSuffix(initialText string) string
	// Completions gets an array of completion strings for a given initial text.
	Completions(initialText string) []string
	// SetDirsOnly: if @dirs_only is true, @completer will only complete
	// directory names, and not file names.
	SetDirsOnly(dirsOnly bool)
}

// FilenameCompleter completes partial file and directory names given a partial
// string by looking in the file system for clues. Can return a list of possible
// completion strings for widget implementations.
type FilenameCompleter struct {
	*externglib.Object
}

var (
	_ FilenameCompleterer = (*FilenameCompleter)(nil)
	_ gextras.Nativer     = (*FilenameCompleter)(nil)
)

func wrapFilenameCompleter(obj *externglib.Object) *FilenameCompleter {
	return &FilenameCompleter{
		Object: obj,
	}
}

func marshalFilenameCompleterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFilenameCompleter(obj), nil
}

// NewFilenameCompleter creates a new filename completer.
func NewFilenameCompleter() *FilenameCompleter {
	var _cret *C.GFilenameCompleter // in

	_cret = C.g_filename_completer_new()

	var _filenameCompleter *FilenameCompleter // out

	_filenameCompleter = wrapFilenameCompleter(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _filenameCompleter
}

// CompletionSuffix obtains a completion for @initial_text from @completer.
func (completer *FilenameCompleter) CompletionSuffix(initialText string) string {
	var _arg0 *C.GFilenameCompleter // out
	var _arg1 *C.char               // out
	var _cret *C.char               // in

	_arg0 = (*C.GFilenameCompleter)(unsafe.Pointer(completer.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(initialText)))

	_cret = C.g_filename_completer_get_completion_suffix(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Completions gets an array of completion strings for a given initial text.
func (completer *FilenameCompleter) Completions(initialText string) []string {
	var _arg0 *C.GFilenameCompleter // out
	var _arg1 *C.char               // out
	var _cret **C.char

	_arg0 = (*C.GFilenameCompleter)(unsafe.Pointer(completer.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(initialText)))

	_cret = C.g_filename_completer_get_completions(_arg0, _arg1)

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

// SetDirsOnly: if @dirs_only is true, @completer will only complete directory
// names, and not file names.
func (completer *FilenameCompleter) SetDirsOnly(dirsOnly bool) {
	var _arg0 *C.GFilenameCompleter // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GFilenameCompleter)(unsafe.Pointer(completer.Native()))
	if dirsOnly {
		_arg1 = C.TRUE
	}

	C.g_filename_completer_set_dirs_only(_arg0, _arg1)
}
