// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// SpawnError: error codes returned by spawning processes.
type SpawnError int

const (
	// SpawnErrorFork: fork failed due to lack of memory.
	SpawnErrorFork SpawnError = 0
	// SpawnErrorRead: read or select on pipes failed.
	SpawnErrorRead SpawnError = 1
	// SpawnErrorChdir: changing to working directory failed.
	SpawnErrorChdir SpawnError = 2
	// SpawnErrorAcces: execv() returned `EACCES`
	SpawnErrorAcces SpawnError = 3
	// SpawnErrorPerm: execv() returned `EPERM`
	SpawnErrorPerm SpawnError = 4
	// SpawnErrorTooBig: execv() returned `E2BIG`
	SpawnErrorTooBig SpawnError = 5
	// SpawnError2Big: deprecated alias for G_SPAWN_ERROR_TOO_BIG (deprecated
	// since GLib 2.32)
	SpawnError2Big SpawnError = 5
	// SpawnErrorNoexec: execv() returned `ENOEXEC`
	SpawnErrorNoexec SpawnError = 6
	// SpawnErrorNametoolong: execv() returned `ENAMETOOLONG`
	SpawnErrorNametoolong SpawnError = 7
	// SpawnErrorNoent: execv() returned `ENOENT`
	SpawnErrorNoent SpawnError = 8
	// SpawnErrorNOMEM: execv() returned `ENOMEM`
	SpawnErrorNOMEM SpawnError = 9
	// SpawnErrorNotdir: execv() returned `ENOTDIR`
	SpawnErrorNotdir SpawnError = 10
	// SpawnErrorLoop: execv() returned `ELOOP`
	SpawnErrorLoop SpawnError = 11
	// SpawnErrorTxtbusy: execv() returned `ETXTBUSY`
	SpawnErrorTxtbusy SpawnError = 12
	// SpawnErrorIO: execv() returned `EIO`
	SpawnErrorIO SpawnError = 13
	// SpawnErrorNfile: execv() returned `ENFILE`
	SpawnErrorNfile SpawnError = 14
	// SpawnErrorMfile: execv() returned `EMFILE`
	SpawnErrorMfile SpawnError = 15
	// SpawnErrorInval: execv() returned `EINVAL`
	SpawnErrorInval SpawnError = 16
	// SpawnErrorIsdir: execv() returned `EISDIR`
	SpawnErrorIsdir SpawnError = 17
	// SpawnErrorLibbad: execv() returned `ELIBBAD`
	SpawnErrorLibbad SpawnError = 18
	// SpawnErrorFailed: some other fatal failure, `error->message` should
	// explain.
	SpawnErrorFailed SpawnError = 19
)

// SpawnFlags flags passed to g_spawn_sync(), g_spawn_async() and
// g_spawn_async_with_pipes().
type SpawnFlags int

const (
	// SpawnFlagsDefault: no flags, default behaviour
	SpawnFlagsDefault SpawnFlags = 0
	// SpawnFlagsLeaveDescriptorsOpen: the parent's open file descriptors will
	// be inherited by the child; otherwise all descriptors except stdin, stdout
	// and stderr will be closed before calling exec() in the child.
	SpawnFlagsLeaveDescriptorsOpen SpawnFlags = 1
	// SpawnFlagsDoNotReapChild: the child will not be automatically reaped; you
	// must use g_child_watch_add() yourself (or call waitpid() or handle
	// `SIGCHLD` yourself), or the child will become a zombie.
	SpawnFlagsDoNotReapChild SpawnFlags = 2
	// SpawnFlagsSearchPath: `argv[0]` need not be an absolute path, it will be
	// looked for in the user's `PATH`.
	SpawnFlagsSearchPath SpawnFlags = 4
	// SpawnFlagsStdoutToDevNull: the child's standard output will be discarded,
	// instead of going to the same location as the parent's standard output.
	SpawnFlagsStdoutToDevNull SpawnFlags = 8
	// SpawnFlagsStderrToDevNull: the child's standard error will be discarded.
	SpawnFlagsStderrToDevNull SpawnFlags = 16
	// SpawnFlagsChildInheritsStdin: the child will inherit the parent's
	// standard input (by default, the child's standard input is attached to
	// `/dev/null`).
	SpawnFlagsChildInheritsStdin SpawnFlags = 32
	// SpawnFlagsFileAndArgvZero: the first element of `argv` is the file to
	// execute, while the remaining elements are the actual argument vector to
	// pass to the file. Normally g_spawn_async_with_pipes() uses `argv[0]` as
	// the file to execute, and passes all of `argv` to the child.
	SpawnFlagsFileAndArgvZero SpawnFlags = 64
	// SpawnFlagsSearchPathFromEnvp: if `argv[0]` is not an absolute path, it
	// will be looked for in the `PATH` from the passed child environment.
	// Since: 2.34
	SpawnFlagsSearchPathFromEnvp SpawnFlags = 128
	// SpawnFlagsCloexecPipes: create all pipes with the `O_CLOEXEC` flag set.
	// Since: 2.40
	SpawnFlagsCloexecPipes SpawnFlags = 256
)

// SpawnCheckExitStatus: set @error if @exit_status indicates the child exited
// abnormally (e.g. with a nonzero exit code, or via a fatal signal).
//
// The g_spawn_sync() and g_child_watch_add() family of APIs return an exit
// status for subprocesses encoded in a platform-specific way. On Unix, this is
// guaranteed to be in the same format waitpid() returns, and on Windows it is
// guaranteed to be the result of GetExitCodeProcess().
//
// Prior to the introduction of this function in GLib 2.34, interpreting
// @exit_status required use of platform-specific APIs, which is problematic for
// software using GLib as a cross-platform layer.
//
// Additionally, many programs simply want to determine whether or not the child
// exited successfully, and either propagate a #GError or print a message to
// standard error. In that common case, this function can be used. Note that the
// error message in @error will contain human-readable information about the
// exit status.
//
// The @domain and @code of @error have special semantics in the case where the
// process has an "exit code", as opposed to being killed by a signal. On Unix,
// this happens if WIFEXITED() would be true of @exit_status. On Windows, it is
// always the case.
//
// The special semantics are that the actual exit code will be the code set in
// @error, and the domain will be G_SPAWN_EXIT_ERROR. This allows you to
// differentiate between different exit codes.
//
// If the process was terminated by some means other than an exit status, the
// domain will be G_SPAWN_ERROR, and the code will be G_SPAWN_ERROR_FAILED.
//
// This function just offers convenience; you can of course also check the
// available platform via a macro such as G_OS_UNIX, and use WIFEXITED() and
// WEXITSTATUS() on @exit_status directly. Do not attempt to scan or parse the
// error message string; it may be translated and/or change in future versions
// of GLib.
func SpawnCheckExitStatus(exitStatus int) error {
	var _arg1 C.gint    // out
	var _cerr *C.GError // in

	_arg1 = (C.gint)(exitStatus)

	C.g_spawn_check_exit_status(_arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SpawnCommandLineAsync: a simple version of g_spawn_async() that parses a
// command line with g_shell_parse_argv() and passes it to g_spawn_async(). Runs
// a command line in the background. Unlike g_spawn_async(), the
// G_SPAWN_SEARCH_PATH flag is enabled, other flags are not. Note that
// G_SPAWN_SEARCH_PATH can have security implications, so consider using
// g_spawn_async() directly if appropriate. Possible errors are those from
// g_shell_parse_argv() and g_spawn_async().
//
// The same concerns on Windows apply as for g_spawn_command_line_sync().
func SpawnCommandLineAsync(commandLine string) error {
	var _arg1 *C.gchar  // out
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(C.CString(commandLine))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_spawn_command_line_async(_arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SpawnCommandLineSync: a simple version of g_spawn_sync() with little-used
// parameters removed, taking a command line instead of an argument vector. See
// g_spawn_sync() for full details. @command_line will be parsed by
// g_shell_parse_argv(). Unlike g_spawn_sync(), the G_SPAWN_SEARCH_PATH flag is
// enabled. Note that G_SPAWN_SEARCH_PATH can have security implications, so
// consider using g_spawn_sync() directly if appropriate. Possible errors are
// those from g_spawn_sync() and those from g_shell_parse_argv().
//
// If @exit_status is non-nil, the platform-specific exit status of the child is
// stored there; see the documentation of g_spawn_check_exit_status() for how to
// use and interpret this.
//
// On Windows, please note the implications of g_shell_parse_argv() parsing
// @command_line. Parsing is done according to Unix shell rules, not Windows
// command interpreter rules. Space is a separator, and backslashes are special.
// Thus you cannot simply pass a @command_line containing canonical Windows
// paths, like "c:\\program files\\app\\app.exe", as the backslashes will be
// eaten, and the space will act as a separator. You need to enclose such paths
// with single quotes, like "'c:\\program files\\app\\app.exe'
// 'e:\\folder\\argument.txt'".
func SpawnCommandLineSync(commandLine string) (standardOutput []byte, standardError []byte, exitStatus int, goerr error) {
	var _arg1 *C.gchar // out
	var _arg2 *C.gchar
	var _arg3 *C.gchar
	var _arg4 C.gint    // in
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(C.CString(commandLine))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_spawn_command_line_sync(_arg1, &_arg2, &_arg3, &_arg4, &_cerr)

	var _standardOutput []byte
	var _standardError []byte
	var _exitStatus int // out
	var _goerr error    // out

	{
		var i int
		for p := _arg2; *p != nil; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_arg2, i)
		_standardOutput = make([]byte, i)
		for i := range src {
			_standardOutput[i] = (byte)(src[i])
		}
	}
	{
		var i int
		for p := _arg3; *p != nil; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_arg3, i)
		_standardError = make([]byte, i)
		for i := range src {
			_standardError[i] = (byte)(src[i])
		}
	}
	_exitStatus = (int)(_arg4)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _standardOutput, _standardError, _exitStatus, _goerr
}
