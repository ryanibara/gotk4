// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_FilenameCompleterClass_got_completion_data(void*);
// extern void _gotk4_gio2_FilenameCompleter_ConnectGotCompletionData(gpointer, guintptr);
import "C"

// GTypeFilenameCompleter returns the GType for the type FilenameCompleter.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFilenameCompleter() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "FilenameCompleter").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalFilenameCompleter)
	return gtype
}

// FilenameCompleterOverrider contains methods that are overridable.
type FilenameCompleterOverrider interface {
	GotCompletionData()
}

// FilenameCompleter completes partial file and directory names given a partial
// string by looking in the file system for clues. Can return a list of possible
// completion strings for widget implementations.
type FilenameCompleter struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*FilenameCompleter)(nil)
)

func classInitFilenameCompleterer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gio", "FilenameCompleterClass")

	if _, ok := goval.(interface{ GotCompletionData() }); ok {
		o := pclass.StructFieldOffset("got_completion_data")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_FilenameCompleterClass_got_completion_data)
	}
}

//export _gotk4_gio2_FilenameCompleterClass_got_completion_data
func _gotk4_gio2_FilenameCompleterClass_got_completion_data(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ GotCompletionData() })

	iface.GotCompletionData()
}

func wrapFilenameCompleter(obj *coreglib.Object) *FilenameCompleter {
	return &FilenameCompleter{
		Object: obj,
	}
}

func marshalFilenameCompleter(p uintptr) (interface{}, error) {
	return wrapFilenameCompleter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_FilenameCompleter_ConnectGotCompletionData
func _gotk4_gio2_FilenameCompleter_ConnectGotCompletionData(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectGotCompletionData is emitted when the file name completion information
// comes available.
func (completer *FilenameCompleter) ConnectGotCompletionData(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(completer, "got-completion-data", false, unsafe.Pointer(C._gotk4_gio2_FilenameCompleter_ConnectGotCompletionData), f)
}

// NewFilenameCompleter creates a new filename completer.
//
// The function returns the following values:
//
//    - filenameCompleter: Completer.
//
func NewFilenameCompleter() *FilenameCompleter {
	_info := girepository.MustFind("Gio", "FilenameCompleter")
	_gret := _info.InvokeClassMethod("new_FilenameCompleter", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _filenameCompleter *FilenameCompleter // out

	_filenameCompleter = wrapFilenameCompleter(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _filenameCompleter
}

// CompletionSuffix obtains a completion for initial_text from completer.
//
// The function takes the following parameters:
//
//    - initialText: text to be completed.
//
// The function returns the following values:
//
//    - utf8 (optional): completed string, or NULL if no completion exists. This
//      string is not owned by GIO, so remember to g_free() it when finished.
//
func (completer *FilenameCompleter) CompletionSuffix(initialText string) string {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completer).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(initialText)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gio", "FilenameCompleter")
	_gret := _info.InvokeClassMethod("get_completion_suffix", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(completer)
	runtime.KeepAlive(initialText)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Completions gets an array of completion strings for a given initial text.
//
// The function takes the following parameters:
//
//    - initialText: text to be completed.
//
// The function returns the following values:
//
//    - utf8s: array of strings with possible completions for initial_text. This
//      array must be freed by g_strfreev() when finished.
//
func (completer *FilenameCompleter) Completions(initialText string) []string {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completer).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(initialText)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gio", "FilenameCompleter")
	_gret := _info.InvokeClassMethod("get_completions", _args[:], nil)
	_cret := *(***C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(completer)
	runtime.KeepAlive(initialText)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.void
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// SetDirsOnly: if dirs_only is TRUE, completer will only complete directory
// names, and not file names.
//
// The function takes the following parameters:
//
//    - dirsOnly: #gboolean.
//
func (completer *FilenameCompleter) SetDirsOnly(dirsOnly bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completer).Native()))
	if dirsOnly {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gio", "FilenameCompleter")
	_info.InvokeClassMethod("set_dirs_only", _args[:], nil)

	runtime.KeepAlive(completer)
	runtime.KeepAlive(dirsOnly)
}
