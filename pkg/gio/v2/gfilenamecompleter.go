// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_filename_completer_get_type()), F: marshalFilenameCompleter},
	})
}

// FilenameCompleter completes partial file and directory names given a partial
// string by looking in the file system for clues. Can return a list of possible
// completion strings for widget implementations.
type FilenameCompleter interface {
	gextras.Objector

	// CompletionSuffix obtains a completion for @initial_text from @completer.
	CompletionSuffix(initialText string) string
	// Completions gets an array of completion strings for a given initial text.
	Completions(initialText string) []string
	// SetDirsOnly: if @dirs_only is true, @completer will only complete
	// directory names, and not file names.
	SetDirsOnly(dirsOnly bool)
}

// filenameCompleter implements the FilenameCompleter interface.
type filenameCompleter struct {
	gextras.Objector
}

var _ FilenameCompleter = (*filenameCompleter)(nil)

// WrapFilenameCompleter wraps a GObject to the right type. It is
// primarily used internally.
func WrapFilenameCompleter(obj *externglib.Object) FilenameCompleter {
	return FilenameCompleter{
		Objector: obj,
	}
}

func marshalFilenameCompleter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFilenameCompleter(obj), nil
}

// NewFilenameCompleter constructs a class FilenameCompleter.
func NewFilenameCompleter() FilenameCompleter {
	ret := C.g_filename_completer_new()

	var ret0 FilenameCompleter

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(FilenameCompleter)

	return ret0
}

// CompletionSuffix obtains a completion for @initial_text from @completer.
func (c filenameCompleter) CompletionSuffix(initialText string) string {
	var arg0 *C.GFilenameCompleter
	var arg1 *C.char

	arg0 = (*C.GFilenameCompleter)(c.Native())
	arg1 = (*C.gchar)(C.CString(initialText))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.g_filename_completer_get_completion_suffix(arg0, arg1)

	var ret0 string

	ret0 = C.GoString(ret)
	C.free(unsafe.Pointer(ret))

	return ret0
}

// Completions gets an array of completion strings for a given initial text.
func (c filenameCompleter) Completions(initialText string) []string {
	var arg0 *C.GFilenameCompleter
	var arg1 *C.char

	arg0 = (*C.GFilenameCompleter)(c.Native())
	arg1 = (*C.gchar)(C.CString(initialText))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.g_filename_completer_get_completions(arg0, arg1)

	var ret0 []string

	{
		var length uint
		for p := unsafe.Pointer(ret); *p != 0; p = unsafe.Pointer(uintptr(p) + 1) {
			length++
		}

		ret0 = make([]string, length)
		for i := 0; i < length; i++ {
			src := (*C.gchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ret)) + i))
			ret0[i] = C.GoString(src)
			C.free(unsafe.Pointer(src))
		}
	}

	return ret0
}

// SetDirsOnly: if @dirs_only is true, @completer will only complete
// directory names, and not file names.
func (c filenameCompleter) SetDirsOnly(dirsOnly bool) {
	var arg0 *C.GFilenameCompleter
	var arg1 C.gboolean

	arg0 = (*C.GFilenameCompleter)(c.Native())
	if dirsOnly {
		arg1 = C.TRUE
	}

	C.g_filename_completer_set_dirs_only(arg0, arg1)
}
