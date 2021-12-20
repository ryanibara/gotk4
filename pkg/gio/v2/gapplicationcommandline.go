// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_application_command_line_get_type()), F: marshalApplicationCommandLiner},
	})
}

// ApplicationCommandLineOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
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
	*externglib.Object
}

var (
	_ externglib.Objector = (*ApplicationCommandLine)(nil)
)

func wrapApplicationCommandLine(obj *externglib.Object) *ApplicationCommandLine {
	return &ApplicationCommandLine{
		Object: obj,
	}
}

func marshalApplicationCommandLiner(p uintptr) (interface{}, error) {
	return wrapApplicationCommandLine(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
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
func (cmdline *ApplicationCommandLine) CreateFileForArg(arg string) Filer {
	var _arg0 *C.GApplicationCommandLine // out
	var _arg1 *C.gchar                   // out
	var _cret *C.GFile                   // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(cmdline.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(arg)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_application_command_line_create_file_for_arg(_arg0, _arg1)
	runtime.KeepAlive(cmdline)
	runtime.KeepAlive(arg)

	var _file Filer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Filer is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.Cast()
		rv, ok := casted.(Filer)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.Filer")
		}
		_file = rv
	}

	return _file
}

// Arguments gets the list of arguments that was passed on the command line.
//
// The strings in the array may contain non-UTF-8 data on UNIX (such as
// filenames or arguments given in the system locale) but are always in UTF-8 on
// Windows.
//
// If you wish to use the return value with Context, you must use
// g_option_context_parse_strv().
//
// The return value is NULL-terminated and should be freed using g_strfreev().
//
// The function returns the following values:
//
//    - filenames: the string array containing the arguments (the argv).
//
func (cmdline *ApplicationCommandLine) Arguments() []string {
	var _arg0 *C.GApplicationCommandLine // out
	var _cret **C.gchar                  // in
	var _arg1 C.int                      // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(cmdline.Native()))

	_cret = C.g_application_command_line_get_arguments(_arg0, &_arg1)
	runtime.KeepAlive(cmdline)

	var _filenames []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		src := unsafe.Slice(_cret, _arg1)
		_filenames = make([]string, _arg1)
		for i := 0; i < int(_arg1); i++ {
			_filenames[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _filenames
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
	var _arg0 *C.GApplicationCommandLine // out
	var _cret *C.gchar                   // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(cmdline.Native()))

	_cret = C.g_application_command_line_get_cwd(_arg0)
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
	var _arg0 *C.GApplicationCommandLine // out
	var _cret **C.gchar                  // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(cmdline.Native()))

	_cret = C.g_application_command_line_get_environ(_arg0)
	runtime.KeepAlive(cmdline)

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

// ExitStatus gets the exit status of cmdline. See
// g_application_command_line_set_exit_status() for more information.
//
// The function returns the following values:
//
//    - gint: exit status.
//
func (cmdline *ApplicationCommandLine) ExitStatus() int {
	var _arg0 *C.GApplicationCommandLine // out
	var _cret C.int                      // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(cmdline.Native()))

	_cret = C.g_application_command_line_get_exit_status(_arg0)
	runtime.KeepAlive(cmdline)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IsRemote determines if cmdline represents a remote invocation.
//
// The function returns the following values:
//
//    - ok: TRUE if the invocation was remote.
//
func (cmdline *ApplicationCommandLine) IsRemote() bool {
	var _arg0 *C.GApplicationCommandLine // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(cmdline.Native()))

	_cret = C.g_application_command_line_get_is_remote(_arg0)
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
	var _arg0 *C.GApplicationCommandLine // out
	var _cret *C.GVariantDict            // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(cmdline.Native()))

	_cret = C.g_application_command_line_get_options_dict(_arg0)
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
	var _arg0 *C.GApplicationCommandLine // out
	var _cret *C.GVariant                // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(cmdline.Native()))

	_cret = C.g_application_command_line_get_platform_data(_arg0)
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
	var _arg0 *C.GApplicationCommandLine // out
	var _cret *C.GInputStream            // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(cmdline.Native()))

	_cret = C.g_application_command_line_get_stdin(_arg0)
	runtime.KeepAlive(cmdline)

	var _inputStream InputStreamer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			casted := object.Cast()
			rv, ok := casted.(InputStreamer)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.InputStreamer")
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
	var _arg0 *C.GApplicationCommandLine // out
	var _arg1 *C.gchar                   // out
	var _cret *C.gchar                   // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(cmdline.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_application_command_line_getenv(_arg0, _arg1)
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
func (cmdline *ApplicationCommandLine) SetExitStatus(exitStatus int) {
	var _arg0 *C.GApplicationCommandLine // out
	var _arg1 C.int                      // out

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(cmdline.Native()))
	_arg1 = C.int(exitStatus)

	C.g_application_command_line_set_exit_status(_arg0, _arg1)
	runtime.KeepAlive(cmdline)
	runtime.KeepAlive(exitStatus)
}
