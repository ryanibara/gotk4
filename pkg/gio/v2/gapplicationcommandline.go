// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern GInputStream* _gotk4_gio2_ApplicationCommandLineClass_get_stdin(GApplicationCommandLine*);
// extern void _gotk4_gio2_ApplicationCommandLineClass_print_literal(GApplicationCommandLine*, gchar*);
// extern void _gotk4_gio2_ApplicationCommandLineClass_printerr_literal(GApplicationCommandLine*, gchar*);
import "C"

// glib.Type values for gapplicationcommandline.go.
var GTypeApplicationCommandLine = coreglib.Type(C.g_application_command_line_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeApplicationCommandLine, F: marshalApplicationCommandLine},
	})
}

// ApplicationCommandLineOverrider contains methods that are overridable.
type ApplicationCommandLineOverrider interface {
	// Stdin gets the stdin of the invoking process.
	//
	// The Stream can be used to read data passed to the standard input of the
	// invoking process. This doesn't work on all platforms. Presently, it is
	// only available on UNIX when using a D-Bus daemon capable of passing file
	// descriptors. If stdin is not available then NULL will be returned. In the
	// future, support may be expanded to other platforms.
	//
	// You must only call this function once per commandline invocation.
	//
	// The function returns the following values:
	//
	//    - inputStream (optional) for stdin.
	//
	Stdin() InputStreamer
	// The function takes the following parameters:
	//
	PrintLiteral(message string)
	// The function takes the following parameters:
	//
	PrinterrLiteral(message string)
}

// ApplicationCommandLine represents a command-line invocation of an
// application. It is created by #GApplication and emitted in the
// #GApplication::command-line signal and virtual function.
//
// The class contains the list of arguments that the program was invoked with.
// It is also possible to query if the commandline invocation was local (ie: the
// current process is running in direct response to the invocation) or remote
// (ie: some other process forwarded the commandline to this process).
//
// The GApplicationCommandLine object can provide the argc and argv parameters
// for use with the Context command-line parsing API, with the
// g_application_command_line_get_arguments() function. See
// [gapplication-example-cmdline3.c][gapplication-example-cmdline3] for an
// example.
//
// The exit status of the originally-invoked process may be set and messages can
// be printed to stdout or stderr of that process. The lifecycle of the
// originally-invoked process is tied to the lifecycle of this object (ie: the
// process exits when the last reference is dropped).
//
// The main use for CommandLine (and the #GApplication::command-line signal) is
// 'Emacs server' like use cases: You can set the EDITOR environment variable to
// have e.g. git use your favourite editor to edit commit messages, and if you
// already have an instance of the editor running, the editing will happen in
// the running instance, instead of opening a new one. An important aspect of
// this use case is that the process that gets started by git does not return
// until the editing is done.
//
// Normally, the commandline is completely handled in the
// #GApplication::command-line handler. The launching instance exits once the
// signal handler in the primary instance has returned, and the return value of
// the signal handler becomes the exit status of the launching instance.
//
//    static gboolean
//    my_cmdline_handler (gpointer data)
//    {
//      GApplicationCommandLine *cmdline = data;
//
//      // do the heavy lifting in an idle
//
//      g_application_command_line_set_exit_status (cmdline, 0);
//      g_object_unref (cmdline); // this releases the application
//
//      return G_SOURCE_REMOVE;
//    }
//
//    static int
//    command_line (GApplication            *application,
//                  GApplicationCommandLine *cmdline)
//    {
//      // keep the application running until we are done with this commandline
//      g_application_hold (application);
//
//      g_object_set_data_full (G_OBJECT (cmdline),
//                              "application", application,
//                              (GDestroyNotify)g_application_release);
//
//      g_object_ref (cmdline);
//      g_idle_add (my_cmdline_handler, cmdline);
//
//      return 0;
//    }
//
// In this example the commandline is not completely handled before the
// #GApplication::command-line handler returns. Instead, we keep a reference to
// the CommandLine object and handle it later (in this example, in an idle).
// Note that it is necessary to hold the application until you are done with the
// commandline.
//
// The complete example can be found here: gapplication-example-cmdline3.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gapplication-example-cmdline3.c).
type ApplicationCommandLine struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ApplicationCommandLine)(nil)
)

func classInitApplicationCommandLiner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GApplicationCommandLineClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GApplicationCommandLineClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Stdin() InputStreamer }); ok {
		pclass.get_stdin = (*[0]byte)(C._gotk4_gio2_ApplicationCommandLineClass_get_stdin)
	}

	if _, ok := goval.(interface{ PrintLiteral(message string) }); ok {
		pclass.print_literal = (*[0]byte)(C._gotk4_gio2_ApplicationCommandLineClass_print_literal)
	}

	if _, ok := goval.(interface{ PrinterrLiteral(message string) }); ok {
		pclass.printerr_literal = (*[0]byte)(C._gotk4_gio2_ApplicationCommandLineClass_printerr_literal)
	}
}

//export _gotk4_gio2_ApplicationCommandLineClass_get_stdin
func _gotk4_gio2_ApplicationCommandLineClass_get_stdin(arg0 *C.GApplicationCommandLine) (cret *C.GInputStream) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Stdin() InputStreamer })

	inputStream := iface.Stdin()

	if inputStream != nil {
		cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(inputStream).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(inputStream).Native()))
	}

	return cret
}

//export _gotk4_gio2_ApplicationCommandLineClass_print_literal
func _gotk4_gio2_ApplicationCommandLineClass_print_literal(arg0 *C.GApplicationCommandLine, arg1 *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ PrintLiteral(message string) })

	var _message string // out

	_message = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	iface.PrintLiteral(_message)
}

//export _gotk4_gio2_ApplicationCommandLineClass_printerr_literal
func _gotk4_gio2_ApplicationCommandLineClass_printerr_literal(arg0 *C.GApplicationCommandLine, arg1 *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ PrinterrLiteral(message string) })

	var _message string // out

	_message = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	iface.PrinterrLiteral(_message)
}

func wrapApplicationCommandLine(obj *coreglib.Object) *ApplicationCommandLine {
	return &ApplicationCommandLine{
		Object: obj,
	}
}

func marshalApplicationCommandLine(p uintptr) (interface{}, error) {
	return wrapApplicationCommandLine(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CreateFileForArg creates a #GFile corresponding to a filename that was given
// as part of the invocation of cmdline.
//
// This differs from g_file_new_for_commandline_arg() in that it resolves
// relative pathnames using the current working directory of the invoking
// process rather than the local process.
//
// The function takes the following parameters:
//
//    - arg: argument from cmdline.
//
// The function returns the following values:
//
//    - file: new #GFile.
//
func (cmdline *ApplicationCommandLine) CreateFileForArg(arg string) *File {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cmdline).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(arg)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**ApplicationCommandLine)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gio", "ApplicationCommandLine").InvokeMethod("create_file_for_arg", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cmdline)
	runtime.KeepAlive(arg)

	var _file *File // out

	_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _file
}

// Cwd gets the working directory of the command line invocation. The string may
// contain non-utf8 data.
//
// It is possible that the remote application did not send a working directory,
// so this may be NULL.
//
// The return value should not be modified or freed and is valid for as long as
// cmdline exists.
//
// The function returns the following values:
//
//    - filename (optional): current directory, or NULL.
//
func (cmdline *ApplicationCommandLine) Cwd() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cmdline).Native()))
	*(**ApplicationCommandLine)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "ApplicationCommandLine").InvokeMethod("get_cwd", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cmdline)

	var _filename string // out

	if _cret != nil {
		_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _filename
}

// Environ gets the contents of the 'environ' variable of the command line
// invocation, as would be returned by g_get_environ(), ie as a NULL-terminated
// list of strings in the form 'NAME=VALUE'. The strings may contain non-utf8
// data.
//
// The remote application usually does not send an environment. Use
// G_APPLICATION_SEND_ENVIRONMENT to affect that. Even with this flag set it is
// possible that the environment is still not available (due to invocation
// messages from other applications).
//
// The return value should not be modified or freed and is valid for as long as
// cmdline exists.
//
// See g_application_command_line_getenv() if you are only interested in the
// value of a single environment variable.
//
// The function returns the following values:
//
//    - filenames: the environment strings, or NULL if they were not sent.
//
func (cmdline *ApplicationCommandLine) Environ() []string {
	var args [1]girepository.Argument
	var _arg0 *C.void   // out
	var _cret **C.gchar // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cmdline).Native()))
	*(**ApplicationCommandLine)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "ApplicationCommandLine").InvokeMethod("get_environ", args[:], nil)
	_cret = *(***C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cmdline)

	var _filenames []string // out

	{
		var i int
		var z *C.void
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

// ExitStatus gets the exit status of cmdline. See
// g_application_command_line_set_exit_status() for more information.
//
// The function returns the following values:
//
//    - gint: exit status.
//
func (cmdline *ApplicationCommandLine) ExitStatus() int32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.int   // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cmdline).Native()))
	*(**ApplicationCommandLine)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "ApplicationCommandLine").InvokeMethod("get_exit_status", args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cmdline)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// IsRemote determines if cmdline represents a remote invocation.
//
// The function returns the following values:
//
//    - ok: TRUE if the invocation was remote.
//
func (cmdline *ApplicationCommandLine) IsRemote() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cmdline).Native()))
	*(**ApplicationCommandLine)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "ApplicationCommandLine").InvokeMethod("get_is_remote", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cmdline)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// OptionsDict gets the options there were passed to
// g_application_command_line().
//
// If you did not override local_command_line() then these are the same options
// that were parsed according to the Entrys added to the application with
// g_application_add_main_option_entries() and possibly modified from your
// GApplication::handle-local-options handler.
//
// If no options were sent then an empty dictionary is returned so that you
// don't need to check for NULL.
//
// The function returns the following values:
//
//    - variantDict with the options.
//
func (cmdline *ApplicationCommandLine) OptionsDict() *glib.VariantDict {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cmdline).Native()))
	*(**ApplicationCommandLine)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "ApplicationCommandLine").InvokeMethod("get_options_dict", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cmdline)

	var _variantDict *glib.VariantDict // out

	_variantDict = (*glib.VariantDict)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_variant_dict_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variantDict)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_dict_unref((*C.GVariantDict)(intern.C))
		},
	)

	return _variantDict
}

// PlatformData gets the platform data associated with the invocation of
// cmdline.
//
// This is a #GVariant dictionary containing information about the context in
// which the invocation occurred. It typically contains information like the
// current working directory and the startup notification ID.
//
// For local invocation, it will be NULL.
//
// The function returns the following values:
//
//    - variant (optional): platform data, or NULL.
//
func (cmdline *ApplicationCommandLine) PlatformData() *glib.Variant {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cmdline).Native()))
	*(**ApplicationCommandLine)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "ApplicationCommandLine").InvokeMethod("get_platform_data", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cmdline)

	var _variant *glib.Variant // out

	if _cret != nil {
		_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_variant)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}

	return _variant
}

// Stdin gets the stdin of the invoking process.
//
// The Stream can be used to read data passed to the standard input of the
// invoking process. This doesn't work on all platforms. Presently, it is only
// available on UNIX when using a D-Bus daemon capable of passing file
// descriptors. If stdin is not available then NULL will be returned. In the
// future, support may be expanded to other platforms.
//
// You must only call this function once per commandline invocation.
//
// The function returns the following values:
//
//    - inputStream (optional) for stdin.
//
func (cmdline *ApplicationCommandLine) Stdin() InputStreamer {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cmdline).Native()))
	*(**ApplicationCommandLine)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "ApplicationCommandLine").InvokeMethod("get_stdin", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cmdline)

	var _inputStream InputStreamer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(InputStreamer)
				return ok
			})
			rv, ok := casted.(InputStreamer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.InputStreamer")
			}
			_inputStream = rv
		}
	}

	return _inputStream
}

// Env gets the value of a particular environment variable of the command line
// invocation, as would be returned by g_getenv(). The strings may contain
// non-utf8 data.
//
// The remote application usually does not send an environment. Use
// G_APPLICATION_SEND_ENVIRONMENT to affect that. Even with this flag set it is
// possible that the environment is still not available (due to invocation
// messages from other applications).
//
// The return value should not be modified or freed and is valid for as long as
// cmdline exists.
//
// The function takes the following parameters:
//
//    - name: environment variable to get.
//
// The function returns the following values:
//
//    - utf8 (optional): value of the variable, or NULL if unset or unsent.
//
func (cmdline *ApplicationCommandLine) env(name string) string {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cmdline).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**ApplicationCommandLine)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gio", "ApplicationCommandLine").InvokeMethod("getenv", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cmdline)
	runtime.KeepAlive(name)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SetExitStatus sets the exit status that will be used when the invoking
// process exits.
//
// The return value of the #GApplication::command-line signal is passed to this
// function when the handler returns. This is the usual way of setting the exit
// status.
//
// In the event that you want the remote invocation to continue running and want
// to decide on the exit status in the future, you can use this call. For the
// case of a remote invocation, the remote process will typically exit when the
// last reference is dropped on cmdline. The exit status of the remote process
// will be equal to the last value that was set with this function.
//
// In the case that the commandline invocation is local, the situation is
// slightly more complicated. If the commandline invocation results in the
// mainloop running (ie: because the use-count of the application increased to a
// non-zero value) then the application is considered to have been 'successful'
// in a certain sense, and the exit status is always zero. If the application
// use count is zero, though, the exit status of the local CommandLine is used.
//
// The function takes the following parameters:
//
//    - exitStatus: exit status.
//
func (cmdline *ApplicationCommandLine) SetExitStatus(exitStatus int32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.int   // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cmdline).Native()))
	_arg1 = C.int(exitStatus)
	*(**ApplicationCommandLine)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gio", "ApplicationCommandLine").InvokeMethod("set_exit_status", args[:], nil)

	runtime.KeepAlive(cmdline)
	runtime.KeepAlive(exitStatus)
}
