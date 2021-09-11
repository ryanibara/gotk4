// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
// void _gotk4_glib2_SpawnChildSetupFunc(gpointer);
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
	// SpawnErrorAcces: execv() returned EACCES
	SpawnErrorAcces SpawnError = 3
	// SpawnErrorPerm: execv() returned EPERM
	SpawnErrorPerm SpawnError = 4
	// SpawnErrorTooBig: execv() returned E2BIG
	SpawnErrorTooBig SpawnError = 5
	// SpawnError2Big: deprecated alias for G_SPAWN_ERROR_TOO_BIG (deprecated
	// since GLib 2.32)
	SpawnError2Big SpawnError = 5
	// SpawnErrorNoexec: execv() returned ENOEXEC
	SpawnErrorNoexec SpawnError = 6
	// SpawnErrorNametoolong: execv() returned ENAMETOOLONG
	SpawnErrorNametoolong SpawnError = 7
	// SpawnErrorNoent: execv() returned ENOENT
	SpawnErrorNoent SpawnError = 8
	// SpawnErrorNOMEM: execv() returned ENOMEM
	SpawnErrorNOMEM SpawnError = 9
	// SpawnErrorNotdir: execv() returned ENOTDIR
	SpawnErrorNotdir SpawnError = 10
	// SpawnErrorLoop: execv() returned ELOOP
	SpawnErrorLoop SpawnError = 11
	// SpawnErrorTxtbusy: execv() returned ETXTBUSY
	SpawnErrorTxtbusy SpawnError = 12
	// SpawnErrorIO: execv() returned EIO
	SpawnErrorIO SpawnError = 13
	// SpawnErrorNfile: execv() returned ENFILE
	SpawnErrorNfile SpawnError = 14
	// SpawnErrorMfile: execv() returned EMFILE
	SpawnErrorMfile SpawnError = 15
	// SpawnErrorInval: execv() returned EINVAL
	SpawnErrorInval SpawnError = 16
	// SpawnErrorIsdir: execv() returned EISDIR
	SpawnErrorIsdir SpawnError = 17
	// SpawnErrorLibbad: execv() returned ELIBBAD
	SpawnErrorLibbad SpawnError = 18
	// SpawnErrorFailed: some other fatal failure, error->message should
	// explain.
	SpawnErrorFailed SpawnError = 19
)

// String returns the name in string for SpawnError.
func (s SpawnError) String() string {
	switch s {
	case SpawnErrorFork:
		return "Fork"
	case SpawnErrorRead:
		return "Read"
	case SpawnErrorChdir:
		return "Chdir"
	case SpawnErrorAcces:
		return "Acces"
	case SpawnErrorPerm:
		return "Perm"
	case SpawnErrorTooBig:
		return "TooBig"
	case SpawnErrorNoexec:
		return "Noexec"
	case SpawnErrorNametoolong:
		return "Nametoolong"
	case SpawnErrorNoent:
		return "Noent"
	case SpawnErrorNOMEM:
		return "NOMEM"
	case SpawnErrorNotdir:
		return "Notdir"
	case SpawnErrorLoop:
		return "Loop"
	case SpawnErrorTxtbusy:
		return "Txtbusy"
	case SpawnErrorIO:
		return "IO"
	case SpawnErrorNfile:
		return "Nfile"
	case SpawnErrorMfile:
		return "Mfile"
	case SpawnErrorInval:
		return "Inval"
	case SpawnErrorIsdir:
		return "Isdir"
	case SpawnErrorLibbad:
		return "Libbad"
	case SpawnErrorFailed:
		return "Failed"
	default:
		return fmt.Sprintf("SpawnError(%d)", s)
	}
}

// SpawnFlags flags passed to g_spawn_sync(), g_spawn_async() and
// g_spawn_async_with_pipes().
type SpawnFlags int

const (
	// SpawnDefault: no flags, default behaviour
	SpawnDefault SpawnFlags = 0b0
	// SpawnLeaveDescriptorsOpen parent's open file descriptors will be
	// inherited by the child; otherwise all descriptors except stdin, stdout
	// and stderr will be closed before calling exec() in the child.
	SpawnLeaveDescriptorsOpen SpawnFlags = 0b1
	// SpawnDoNotReapChild: child will not be automatically reaped; you must use
	// g_child_watch_add() yourself (or call waitpid() or handle SIGCHLD
	// yourself), or the child will become a zombie.
	SpawnDoNotReapChild SpawnFlags = 0b10
	// SpawnSearchPath: argv[0] need not be an absolute path, it will be looked
	// for in the user's PATH.
	SpawnSearchPath SpawnFlags = 0b100
	// SpawnStdoutToDevNull child's standard output will be discarded, instead
	// of going to the same location as the parent's standard output.
	SpawnStdoutToDevNull SpawnFlags = 0b1000
	// SpawnStderrToDevNull child's standard error will be discarded.
	SpawnStderrToDevNull SpawnFlags = 0b10000
	// SpawnChildInheritsStdin: child will inherit the parent's standard input
	// (by default, the child's standard input is attached to /dev/null).
	SpawnChildInheritsStdin SpawnFlags = 0b100000
	// SpawnFileAndArgvZero: first element of argv is the file to execute, while
	// the remaining elements are the actual argument vector to pass to the
	// file. Normally g_spawn_async_with_pipes() uses argv[0] as the file to
	// execute, and passes all of argv to the child.
	SpawnFileAndArgvZero SpawnFlags = 0b1000000
	// SpawnSearchPathFromEnvp: if argv[0] is not an absolute path, it will be
	// looked for in the PATH from the passed child environment. Since: 2.34
	SpawnSearchPathFromEnvp SpawnFlags = 0b10000000
	// SpawnCloexecPipes: create all pipes with the O_CLOEXEC flag set. Since:
	// 2.40
	SpawnCloexecPipes SpawnFlags = 0b100000000
)

// String returns the names in string for SpawnFlags.
func (s SpawnFlags) String() string {
	if s == 0 {
		return "SpawnFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(203)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case SpawnDefault:
			builder.WriteString("Default|")
		case SpawnLeaveDescriptorsOpen:
			builder.WriteString("LeaveDescriptorsOpen|")
		case SpawnDoNotReapChild:
			builder.WriteString("DoNotReapChild|")
		case SpawnSearchPath:
			builder.WriteString("SearchPath|")
		case SpawnStdoutToDevNull:
			builder.WriteString("StdoutToDevNull|")
		case SpawnStderrToDevNull:
			builder.WriteString("StderrToDevNull|")
		case SpawnChildInheritsStdin:
			builder.WriteString("ChildInheritsStdin|")
		case SpawnFileAndArgvZero:
			builder.WriteString("FileAndArgvZero|")
		case SpawnSearchPathFromEnvp:
			builder.WriteString("SearchPathFromEnvp|")
		case SpawnCloexecPipes:
			builder.WriteString("CloexecPipes|")
		default:
			builder.WriteString(fmt.Sprintf("SpawnFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s SpawnFlags) Has(other SpawnFlags) bool {
	return (s & other) == other
}

// SpawnChildSetupFunc specifies the type of the setup function passed to
// g_spawn_async(), g_spawn_sync() and g_spawn_async_with_pipes(), which can, in
// very limited ways, be used to affect the child's execution.
//
// On POSIX platforms, the function is called in the child after GLib has
// performed all the setup it plans to perform, but before calling exec().
// Actions taken in this function will only affect the child, not the parent.
//
// On Windows, the function is called in the parent. Its usefulness on Windows
// is thus questionable. In many cases executing the child setup function in the
// parent can have ill effects, and you should be very careful when porting
// software to Windows that uses child setup functions.
//
// However, even on POSIX, you are extremely limited in what you can safely do
// from a ChildSetupFunc, because any mutexes that were held by other threads in
// the parent process at the time of the fork() will still be locked in the
// child process, and they will never be unlocked (since the threads that held
// them don't exist in the child). POSIX allows only async-signal-safe functions
// (see signal(7)) to be called in the child between fork() and exec(), which
// drastically limits the usefulness of child setup functions.
//
// In particular, it is not safe to call any function which may call malloc(),
// which includes POSIX functions such as setenv(). If you need to set up the
// child environment differently from the parent, you should use
// g_get_environ(), g_environ_setenv(), and g_environ_unsetenv(), and then pass
// the complete environment list to the g_spawn... function.
type SpawnChildSetupFunc func()

//export _gotk4_glib2_SpawnChildSetupFunc
func _gotk4_glib2_SpawnChildSetupFunc(arg0 C.gpointer) {
	v := gbox.Get(uintptr(arg0))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(SpawnChildSetupFunc)
	fn()
}

// SpawnAsync: see g_spawn_async_with_pipes() for a full description; this
// function simply calls the g_spawn_async_with_pipes() without any pipes.
//
// You should call g_spawn_close_pid() on the returned child process reference
// when you don't need it any more.
//
// If you are writing a GTK+ application, and the program you are spawning is a
// graphical application too, then to ensure that the spawned program opens its
// windows on the right screen, you may want to use AppLaunchContext,
// LaunchContext, or set the DISPLAY environment variable.
//
// Note that the returned child_pid on Windows is a handle to the child process
// and not its identifier. Process handles and process identifiers are different
// concepts on Windows.
func SpawnAsync(workingDirectory string, argv []string, envp []string, flags SpawnFlags, childSetup SpawnChildSetupFunc) (Pid, error) {
	var _arg1 *C.gchar               // out
	var _arg2 **C.gchar              // out
	var _arg3 **C.gchar              // out
	var _arg4 C.GSpawnFlags          // out
	var _arg5 C.GSpawnChildSetupFunc // out
	var _arg6 C.gpointer
	var _arg7 C.GPid    // in
	var _cerr *C.GError // in

	if workingDirectory != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(workingDirectory)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	{
		_arg2 = (**C.gchar)(C.malloc(C.ulong(len(argv)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(argv)+1)
			var zero *C.gchar
			out[len(argv)] = zero
			for i := range argv {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(argv[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	{
		_arg3 = (**C.gchar)(C.malloc(C.ulong(len(envp)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg3))
		{
			out := unsafe.Slice(_arg3, len(envp)+1)
			var zero *C.gchar
			out[len(envp)] = zero
			for i := range envp {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(envp[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	_arg4 = C.GSpawnFlags(flags)
	if childSetup != nil {
		_arg5 = (*[0]byte)(C._gotk4_glib2_SpawnChildSetupFunc)
		_arg6 = C.gpointer(gbox.AssignOnce(childSetup))
	}

	C.g_spawn_async(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, &_arg7, &_cerr)
	runtime.KeepAlive(workingDirectory)
	runtime.KeepAlive(argv)
	runtime.KeepAlive(envp)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(childSetup)

	var _childPid Pid // out
	var _goerr error  // out

	_childPid = int(_arg7)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _childPid, _goerr
}

// SpawnAsyncWithFds: identical to g_spawn_async_with_pipes_and_fds() but with
// n_fds set to zero, so no FD assignments are used.
func SpawnAsyncWithFds(workingDirectory string, argv []string, envp []string, flags SpawnFlags, childSetup SpawnChildSetupFunc, stdinFd int, stdoutFd int, stderrFd int) (Pid, error) {
	var _arg1 *C.gchar               // out
	var _arg2 **C.gchar              // out
	var _arg3 **C.gchar              // out
	var _arg4 C.GSpawnFlags          // out
	var _arg5 C.GSpawnChildSetupFunc // out
	var _arg6 C.gpointer
	var _arg7 C.GPid    // in
	var _arg8 C.gint    // out
	var _arg9 C.gint    // out
	var _arg10 C.gint   // out
	var _cerr *C.GError // in

	if workingDirectory != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(workingDirectory)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	{
		_arg2 = (**C.gchar)(C.malloc(C.ulong(len(argv)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(argv)+1)
			var zero *C.gchar
			out[len(argv)] = zero
			for i := range argv {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(argv[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	{
		_arg3 = (**C.gchar)(C.malloc(C.ulong(len(envp)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg3))
		{
			out := unsafe.Slice(_arg3, len(envp)+1)
			var zero *C.gchar
			out[len(envp)] = zero
			for i := range envp {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(envp[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	_arg4 = C.GSpawnFlags(flags)
	if childSetup != nil {
		_arg5 = (*[0]byte)(C._gotk4_glib2_SpawnChildSetupFunc)
		_arg6 = C.gpointer(gbox.AssignOnce(childSetup))
	}
	_arg8 = C.gint(stdinFd)
	_arg9 = C.gint(stdoutFd)
	_arg10 = C.gint(stderrFd)

	C.g_spawn_async_with_fds(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, &_arg7, _arg8, _arg9, _arg10, &_cerr)
	runtime.KeepAlive(workingDirectory)
	runtime.KeepAlive(argv)
	runtime.KeepAlive(envp)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(childSetup)
	runtime.KeepAlive(stdinFd)
	runtime.KeepAlive(stdoutFd)
	runtime.KeepAlive(stderrFd)

	var _childPid Pid // out
	var _goerr error  // out

	_childPid = int(_arg7)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _childPid, _goerr
}

// SpawnAsyncWithPipes: identical to g_spawn_async_with_pipes_and_fds() but with
// n_fds set to zero, so no FD assignments are used.
func SpawnAsyncWithPipes(workingDirectory string, argv []string, envp []string, flags SpawnFlags, childSetup SpawnChildSetupFunc) (childPid Pid, standardInput int, standardOutput int, standardError int, goerr error) {
	var _arg1 *C.gchar               // out
	var _arg2 **C.gchar              // out
	var _arg3 **C.gchar              // out
	var _arg4 C.GSpawnFlags          // out
	var _arg5 C.GSpawnChildSetupFunc // out
	var _arg6 C.gpointer
	var _arg7 C.GPid    // in
	var _arg8 C.gint    // in
	var _arg9 C.gint    // in
	var _arg10 C.gint   // in
	var _cerr *C.GError // in

	if workingDirectory != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(workingDirectory)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	{
		_arg2 = (**C.gchar)(C.malloc(C.ulong(len(argv)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(argv)+1)
			var zero *C.gchar
			out[len(argv)] = zero
			for i := range argv {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(argv[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	{
		_arg3 = (**C.gchar)(C.malloc(C.ulong(len(envp)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg3))
		{
			out := unsafe.Slice(_arg3, len(envp)+1)
			var zero *C.gchar
			out[len(envp)] = zero
			for i := range envp {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(envp[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	_arg4 = C.GSpawnFlags(flags)
	if childSetup != nil {
		_arg5 = (*[0]byte)(C._gotk4_glib2_SpawnChildSetupFunc)
		_arg6 = C.gpointer(gbox.AssignOnce(childSetup))
	}

	C.g_spawn_async_with_pipes(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, &_arg7, &_arg8, &_arg9, &_arg10, &_cerr)
	runtime.KeepAlive(workingDirectory)
	runtime.KeepAlive(argv)
	runtime.KeepAlive(envp)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(childSetup)

	var _childPid Pid       // out
	var _standardInput int  // out
	var _standardOutput int // out
	var _standardError int  // out
	var _goerr error        // out

	_childPid = int(_arg7)
	_standardInput = int(_arg8)
	_standardOutput = int(_arg9)
	_standardError = int(_arg10)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _childPid, _standardInput, _standardOutput, _standardError, _goerr
}

// SpawnAsyncWithPipesAndFds executes a child program asynchronously (your
// program will not block waiting for the child to exit). The child program is
// specified by the only argument that must be provided, argv. argv should be a
// NULL-terminated array of strings, to be passed as the argument vector for the
// child. The first string in argv is of course the name of the program to
// execute. By default, the name of the program must be a full path. If flags
// contains the G_SPAWN_SEARCH_PATH flag, the PATH environment variable is used
// to search for the executable. If flags contains the
// G_SPAWN_SEARCH_PATH_FROM_ENVP flag, the PATH variable from envp is used to
// search for the executable. If both the G_SPAWN_SEARCH_PATH and
// G_SPAWN_SEARCH_PATH_FROM_ENVP flags are set, the PATH variable from envp
// takes precedence over the environment variable.
//
// If the program name is not a full path and G_SPAWN_SEARCH_PATH flag is not
// used, then the program will be run from the current directory (or
// working_directory, if specified); this might be unexpected or even dangerous
// in some cases when the current directory is world-writable.
//
// On Windows, note that all the string or string vector arguments to this
// function and the other g_spawn*() functions are in UTF-8, the GLib file name
// encoding. Unicode characters that are not part of the system codepage passed
// in these arguments will be correctly available in the spawned program only if
// it uses wide character API to retrieve its command line. For C programs built
// with Microsoft's tools it is enough to make the program have a wmain()
// instead of main(). wmain() has a wide character argument vector as parameter.
//
// At least currently, mingw doesn't support wmain(), so if you use mingw to
// develop the spawned program, it should call g_win32_get_command_line() to get
// arguments in UTF-8.
//
// On Windows the low-level child process creation API CreateProcess() doesn't
// use argument vectors, but a command line. The C runtime library's spawn*()
// family of functions (which g_spawn_async_with_pipes() eventually calls) paste
// the argument vector elements together into a command line, and the C runtime
// startup code does a corresponding reconstruction of an argument vector from
// the command line, to be passed to main(). Complications arise when you have
// argument vector elements that contain spaces or double quotes. The spawn*()
// functions don't do any quoting or escaping, but on the other hand the startup
// code does do unquoting and unescaping in order to enable receiving arguments
// with embedded spaces or double quotes. To work around this asymmetry,
// g_spawn_async_with_pipes() will do quoting and escaping on argument vector
// elements that need it before calling the C runtime spawn() function.
//
// The returned child_pid on Windows is a handle to the child process, not its
// identifier. Process handles and process identifiers are different concepts on
// Windows.
//
// envp is a NULL-terminated array of strings, where each string has the form
// KEY=VALUE. This will become the child's environment. If envp is NULL, the
// child inherits its parent's environment.
//
// flags should be the bitwise OR of any flags you want to affect the function's
// behaviour. The G_SPAWN_DO_NOT_REAP_CHILD means that the child will not
// automatically be reaped; you must use a child watch (g_child_watch_add()) to
// be notified about the death of the child process, otherwise it will stay
// around as a zombie process until this process exits. Eventually you must call
// g_spawn_close_pid() on the child_pid, in order to free resources which may be
// associated with the child process. (On Unix, using a child watch is
// equivalent to calling waitpid() or handling the SIGCHLD signal manually. On
// Windows, calling g_spawn_close_pid() is equivalent to calling CloseHandle()
// on the process handle returned in child_pid). See g_child_watch_add().
//
// Open UNIX file descriptors marked as FD_CLOEXEC will be automatically closed
// in the child process. G_SPAWN_LEAVE_DESCRIPTORS_OPEN means that other open
// file descriptors will be inherited by the child; otherwise all descriptors
// except stdin/stdout/stderr will be closed before calling exec() in the child.
// G_SPAWN_SEARCH_PATH means that argv[0] need not be an absolute path, it will
// be looked for in the PATH environment variable. G_SPAWN_SEARCH_PATH_FROM_ENVP
// means need not be an absolute path, it will be looked for in the PATH
// variable from envp. If both G_SPAWN_SEARCH_PATH and
// G_SPAWN_SEARCH_PATH_FROM_ENVP are used, the value from envp takes precedence
// over the environment.
//
// G_SPAWN_STDOUT_TO_DEV_NULL means that the child's standard output will be
// discarded, instead of going to the same location as the parent's standard
// output. If you use this flag, stdout_pipe_out must be NULL.
//
// G_SPAWN_STDERR_TO_DEV_NULL means that the child's standard error will be
// discarded, instead of going to the same location as the parent's standard
// error. If you use this flag, stderr_pipe_out must be NULL.
//
// G_SPAWN_CHILD_INHERITS_STDIN means that the child will inherit the parent's
// standard input (by default, the child's standard input is attached to
// /dev/null). If you use this flag, stdin_pipe_out must be NULL.
//
// It is valid to pass the same FD in multiple parameters (e.g. you can pass a
// single FD for both stdout_fd and stderr_fd, and include it in source_fds
// too).
//
// source_fds and target_fds allow zero or more FDs from this process to be
// remapped to different FDs in the spawned process. If n_fds is greater than
// zero, source_fds and target_fds must both be non-NULL and the same length.
// Each FD in source_fds is remapped to the FD number at the same index in
// target_fds. The source and target FD may be equal to simply propagate an FD
// to the spawned process. FD remappings are processed after standard FDs, so
// any target FDs which equal stdin_fd, stdout_fd or stderr_fd will overwrite
// them in the spawned process.
//
// G_SPAWN_FILE_AND_ARGV_ZERO means that the first element of argv is the file
// to execute, while the remaining elements are the actual argument vector to
// pass to the file. Normally g_spawn_async_with_pipes() uses argv[0] as the
// file to execute, and passes all of argv to the child.
//
// child_setup and user_data are a function and user data. On POSIX platforms,
// the function is called in the child after GLib has performed all the setup it
// plans to perform (including creating pipes, closing file descriptors, etc.)
// but before calling exec(). That is, child_setup is called just before calling
// exec() in the child. Obviously actions taken in this function will only
// affect the child, not the parent.
//
// On Windows, there is no separate fork() and exec() functionality. Child
// processes are created and run with a single API call, CreateProcess(). There
// is no sensible thing child_setup could be used for on Windows so it is
// ignored and not called.
//
// If non-NULL, child_pid will on Unix be filled with the child's process ID.
// You can use the process ID to send signals to the child, or to use
// g_child_watch_add() (or waitpid()) if you specified the
// G_SPAWN_DO_NOT_REAP_CHILD flag. On Windows, child_pid will be filled with a
// handle to the child process only if you specified the
// G_SPAWN_DO_NOT_REAP_CHILD flag. You can then access the child process using
// the Win32 API, for example wait for its termination with the WaitFor*()
// functions, or examine its exit code with GetExitCodeProcess(). You should
// close the handle with CloseHandle() or g_spawn_close_pid() when you no longer
// need it.
//
// If non-NULL, the stdin_pipe_out, stdout_pipe_out, stderr_pipe_out locations
// will be filled with file descriptors for writing to the child's standard
// input or reading from its standard output or standard error. The caller of
// g_spawn_async_with_pipes() must close these file descriptors when they are no
// longer in use. If these parameters are NULL, the corresponding pipe won't be
// created.
//
// If stdin_pipe_out is NULL, the child's standard input is attached to
// /dev/null unless G_SPAWN_CHILD_INHERITS_STDIN is set.
//
// If stderr_pipe_out is NULL, the child's standard error goes to the same
// location as the parent's standard error unless G_SPAWN_STDERR_TO_DEV_NULL is
// set.
//
// If stdout_pipe_out is NULL, the child's standard output goes to the same
// location as the parent's standard output unless G_SPAWN_STDOUT_TO_DEV_NULL is
// set.
//
// error can be NULL to ignore errors, or non-NULL to report errors. If an error
// is set, the function returns FALSE. Errors are reported even if they occur in
// the child (for example if the executable in argv[0] is not found). Typically
// the message field of returned errors should be displayed to users. Possible
// errors are those from the SPAWN_ERROR domain.
//
// If an error occurs, child_pid, stdin_pipe_out, stdout_pipe_out, and
// stderr_pipe_out will not be filled with valid values.
//
// If child_pid is not NULL and an error does not occur then the returned
// process reference must be closed using g_spawn_close_pid().
//
// On modern UNIX platforms, GLib can use an efficient process launching
// codepath driven internally by posix_spawn(). This has the advantage of
// avoiding the fork-time performance costs of cloning the parent process
// address space, and avoiding associated memory overcommit checks that are not
// relevant in the context of immediately executing a distinct process. This
// optimized codepath will be used provided that the following conditions are
// met:
//
// 1. G_SPAWN_DO_NOT_REAP_CHILD is set 2. G_SPAWN_LEAVE_DESCRIPTORS_OPEN is set
// 3. G_SPAWN_SEARCH_PATH_FROM_ENVP is not set 4. working_directory is NULL 5.
// child_setup is NULL 6. The program is of a recognised binary format, or has a
// shebang. Otherwise, GLib will have to execute the program through the shell,
// which is not done using the optimized codepath.
//
// If you are writing a GTK+ application, and the program you are spawning is a
// graphical application too, then to ensure that the spawned program opens its
// windows on the right screen, you may want to use AppLaunchContext,
// LaunchContext, or set the DISPLAY environment variable.
func SpawnAsyncWithPipesAndFds(workingDirectory string, argv []string, envp []string, flags SpawnFlags, childSetup SpawnChildSetupFunc, stdinFd int, stdoutFd int, stderrFd int, sourceFds []int, targetFds []int) (childPidOut Pid, stdinPipeOut int, stdoutPipeOut int, stderrPipeOut int, goerr error) {
	var _arg1 *C.gchar               // out
	var _arg2 **C.gchar              // out
	var _arg3 **C.gchar              // out
	var _arg4 C.GSpawnFlags          // out
	var _arg5 C.GSpawnChildSetupFunc // out
	var _arg6 C.gpointer
	var _arg7 C.gint   // out
	var _arg8 C.gint   // out
	var _arg9 C.gint   // out
	var _arg10 *C.gint // out
	var _arg12 C.gsize
	var _arg11 *C.gint  // out
	var _arg13 C.GPid   // in
	var _arg14 C.gint   // in
	var _arg15 C.gint   // in
	var _arg16 C.gint   // in
	var _cerr *C.GError // in

	if workingDirectory != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(workingDirectory)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	{
		_arg2 = (**C.gchar)(C.malloc(C.ulong(len(argv)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(argv)+1)
			var zero *C.gchar
			out[len(argv)] = zero
			for i := range argv {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(argv[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	{
		_arg3 = (**C.gchar)(C.malloc(C.ulong(len(envp)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg3))
		{
			out := unsafe.Slice(_arg3, len(envp)+1)
			var zero *C.gchar
			out[len(envp)] = zero
			for i := range envp {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(envp[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	_arg4 = C.GSpawnFlags(flags)
	if childSetup != nil {
		_arg5 = (*[0]byte)(C._gotk4_glib2_SpawnChildSetupFunc)
		_arg6 = C.gpointer(gbox.AssignOnce(childSetup))
	}
	_arg7 = C.gint(stdinFd)
	_arg8 = C.gint(stdoutFd)
	_arg9 = C.gint(stderrFd)
	_arg12 = (C.gsize)(len(sourceFds))
	_arg10 = (*C.gint)(C.malloc(C.ulong(len(sourceFds)) * C.ulong(C.sizeof_gint)))
	defer C.free(unsafe.Pointer(_arg10))
	{
		out := unsafe.Slice((*C.gint)(_arg10), len(sourceFds))
		for i := range sourceFds {
			out[i] = C.gint(sourceFds[i])
		}
	}
	_arg12 = (C.gsize)(len(targetFds))
	_arg11 = (*C.gint)(C.malloc(C.ulong(len(targetFds)) * C.ulong(C.sizeof_gint)))
	defer C.free(unsafe.Pointer(_arg11))
	{
		out := unsafe.Slice((*C.gint)(_arg11), len(targetFds))
		for i := range targetFds {
			out[i] = C.gint(targetFds[i])
		}
	}

	C.g_spawn_async_with_pipes_and_fds(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10, _arg11, _arg12, &_arg13, &_arg14, &_arg15, &_arg16, &_cerr)
	runtime.KeepAlive(workingDirectory)
	runtime.KeepAlive(argv)
	runtime.KeepAlive(envp)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(childSetup)
	runtime.KeepAlive(stdinFd)
	runtime.KeepAlive(stdoutFd)
	runtime.KeepAlive(stderrFd)
	runtime.KeepAlive(sourceFds)
	runtime.KeepAlive(targetFds)

	var _childPidOut Pid   // out
	var _stdinPipeOut int  // out
	var _stdoutPipeOut int // out
	var _stderrPipeOut int // out
	var _goerr error       // out

	_childPidOut = int(_arg13)
	_stdinPipeOut = int(_arg14)
	_stdoutPipeOut = int(_arg15)
	_stderrPipeOut = int(_arg16)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _childPidOut, _stdinPipeOut, _stdoutPipeOut, _stderrPipeOut, _goerr
}

// SpawnCheckExitStatus: set error if exit_status indicates the child exited
// abnormally (e.g. with a nonzero exit code, or via a fatal signal).
//
// The g_spawn_sync() and g_child_watch_add() family of APIs return an exit
// status for subprocesses encoded in a platform-specific way. On Unix, this is
// guaranteed to be in the same format waitpid() returns, and on Windows it is
// guaranteed to be the result of GetExitCodeProcess().
//
// Prior to the introduction of this function in GLib 2.34, interpreting
// exit_status required use of platform-specific APIs, which is problematic for
// software using GLib as a cross-platform layer.
//
// Additionally, many programs simply want to determine whether or not the child
// exited successfully, and either propagate a #GError or print a message to
// standard error. In that common case, this function can be used. Note that the
// error message in error will contain human-readable information about the exit
// status.
//
// The domain and code of error have special semantics in the case where the
// process has an "exit code", as opposed to being killed by a signal. On Unix,
// this happens if WIFEXITED() would be true of exit_status. On Windows, it is
// always the case.
//
// The special semantics are that the actual exit code will be the code set in
// error, and the domain will be G_SPAWN_EXIT_ERROR. This allows you to
// differentiate between different exit codes.
//
// If the process was terminated by some means other than an exit status, the
// domain will be G_SPAWN_ERROR, and the code will be G_SPAWN_ERROR_FAILED.
//
// This function just offers convenience; you can of course also check the
// available platform via a macro such as G_OS_UNIX, and use WIFEXITED() and
// WEXITSTATUS() on exit_status directly. Do not attempt to scan or parse the
// error message string; it may be translated and/or change in future versions
// of GLib.
func SpawnCheckExitStatus(exitStatus int) error {
	var _arg1 C.gint    // out
	var _cerr *C.GError // in

	_arg1 = C.gint(exitStatus)

	C.g_spawn_check_exit_status(_arg1, &_cerr)
	runtime.KeepAlive(exitStatus)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SpawnClosePid: on some platforms, notably Windows, the #GPid type represents
// a resource which must be closed to prevent resource leaking.
// g_spawn_close_pid() is provided for this purpose. It should be used on all
// platforms, even though it doesn't do anything under UNIX.
func SpawnClosePid(pid Pid) {
	var _arg1 C.GPid // out

	_arg1 = C.int(pid)

	C.g_spawn_close_pid(_arg1)
	runtime.KeepAlive(pid)
}

// SpawnCommandLineAsync: simple version of g_spawn_async() that parses a
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

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(commandLine)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_spawn_command_line_async(_arg1, &_cerr)
	runtime.KeepAlive(commandLine)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SpawnCommandLineSync: simple version of g_spawn_sync() with little-used
// parameters removed, taking a command line instead of an argument vector. See
// g_spawn_sync() for full details. command_line will be parsed by
// g_shell_parse_argv(). Unlike g_spawn_sync(), the G_SPAWN_SEARCH_PATH flag is
// enabled. Note that G_SPAWN_SEARCH_PATH can have security implications, so
// consider using g_spawn_sync() directly if appropriate. Possible errors are
// those from g_spawn_sync() and those from g_shell_parse_argv().
//
// If exit_status is non-NULL, the platform-specific exit status of the child is
// stored there; see the documentation of g_spawn_check_exit_status() for how to
// use and interpret this.
//
// On Windows, please note the implications of g_shell_parse_argv() parsing
// command_line. Parsing is done according to Unix shell rules, not Windows
// command interpreter rules. Space is a separator, and backslashes are special.
// Thus you cannot simply pass a command_line containing canonical Windows
// paths, like "c:\\program files\\app\\app.exe", as the backslashes will be
// eaten, and the space will act as a separator. You need to enclose such paths
// with single quotes, like "'c:\\program files\\app\\app.exe'
// 'e:\\folder\\argument.txt'".
func SpawnCommandLineSync(commandLine string) (standardOutput []byte, standardError []byte, exitStatus int, goerr error) {
	var _arg1 *C.gchar  // out
	var _arg2 *C.gchar  // in
	var _arg3 *C.gchar  // in
	var _arg4 C.gint    // in
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(commandLine)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_spawn_command_line_sync(_arg1, &_arg2, &_arg3, &_arg4, &_cerr)
	runtime.KeepAlive(commandLine)

	var _standardOutput []byte // out
	var _standardError []byte  // out
	var _exitStatus int        // out
	var _goerr error           // out

	if _arg2 != nil {
		defer C.free(unsafe.Pointer(_arg2))
		{
			var i int
			var z C.gchar
			for p := _arg2; *p != z; p = &unsafe.Slice(p, i+1)[i] {
				i++
			}

			src := unsafe.Slice(_arg2, i)
			_standardOutput = make([]byte, i)
			for i := range src {
				_standardOutput[i] = byte(src[i])
			}
		}
	}
	if _arg3 != nil {
		defer C.free(unsafe.Pointer(_arg3))
		{
			var i int
			var z C.gchar
			for p := _arg3; *p != z; p = &unsafe.Slice(p, i+1)[i] {
				i++
			}

			src := unsafe.Slice(_arg3, i)
			_standardError = make([]byte, i)
			for i := range src {
				_standardError[i] = byte(src[i])
			}
		}
	}
	_exitStatus = int(_arg4)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _standardOutput, _standardError, _exitStatus, _goerr
}

// SpawnSync executes a child synchronously (waits for the child to exit before
// returning). All output from the child is stored in standard_output and
// standard_error, if those parameters are non-NULL. Note that you must set the
// G_SPAWN_STDOUT_TO_DEV_NULL and G_SPAWN_STDERR_TO_DEV_NULL flags when passing
// NULL for standard_output and standard_error.
//
// If exit_status is non-NULL, the platform-specific exit status of the child is
// stored there; see the documentation of g_spawn_check_exit_status() for how to
// use and interpret this. Note that it is invalid to pass
// G_SPAWN_DO_NOT_REAP_CHILD in flags, and on POSIX platforms, the same
// restrictions as for g_child_watch_source_new() apply.
//
// If an error occurs, no data is returned in standard_output, standard_error,
// or exit_status.
//
// This function calls g_spawn_async_with_pipes() internally; see that function
// for full details on the other parameters and details on how these functions
// work on Windows.
func SpawnSync(workingDirectory string, argv []string, envp []string, flags SpawnFlags, childSetup SpawnChildSetupFunc) (standardOutput []byte, standardError []byte, exitStatus int, goerr error) {
	var _arg1 *C.gchar               // out
	var _arg2 **C.gchar              // out
	var _arg3 **C.gchar              // out
	var _arg4 C.GSpawnFlags          // out
	var _arg5 C.GSpawnChildSetupFunc // out
	var _arg6 C.gpointer
	var _arg7 *C.gchar  // in
	var _arg8 *C.gchar  // in
	var _arg9 C.gint    // in
	var _cerr *C.GError // in

	if workingDirectory != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(workingDirectory)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	{
		_arg2 = (**C.gchar)(C.malloc(C.ulong(len(argv)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(argv)+1)
			var zero *C.gchar
			out[len(argv)] = zero
			for i := range argv {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(argv[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	{
		_arg3 = (**C.gchar)(C.malloc(C.ulong(len(envp)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg3))
		{
			out := unsafe.Slice(_arg3, len(envp)+1)
			var zero *C.gchar
			out[len(envp)] = zero
			for i := range envp {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(envp[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	_arg4 = C.GSpawnFlags(flags)
	if childSetup != nil {
		_arg5 = (*[0]byte)(C._gotk4_glib2_SpawnChildSetupFunc)
		_arg6 = C.gpointer(gbox.AssignOnce(childSetup))
	}

	C.g_spawn_sync(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, &_arg7, &_arg8, &_arg9, &_cerr)
	runtime.KeepAlive(workingDirectory)
	runtime.KeepAlive(argv)
	runtime.KeepAlive(envp)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(childSetup)

	var _standardOutput []byte // out
	var _standardError []byte  // out
	var _exitStatus int        // out
	var _goerr error           // out

	if _arg7 != nil {
		defer C.free(unsafe.Pointer(_arg7))
		{
			var i int
			var z C.gchar
			for p := _arg7; *p != z; p = &unsafe.Slice(p, i+1)[i] {
				i++
			}

			src := unsafe.Slice(_arg7, i)
			_standardOutput = make([]byte, i)
			for i := range src {
				_standardOutput[i] = byte(src[i])
			}
		}
	}
	if _arg8 != nil {
		defer C.free(unsafe.Pointer(_arg8))
		{
			var i int
			var z C.gchar
			for p := _arg8; *p != z; p = &unsafe.Slice(p, i+1)[i] {
				i++
			}

			src := unsafe.Slice(_arg8, i)
			_standardError = make([]byte, i)
			for i := range src {
				_standardError[i] = byte(src[i])
			}
		}
	}
	_exitStatus = int(_arg9)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _standardOutput, _standardError, _exitStatus, _goerr
}
