// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/cgo"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
// GLogWriterOutput _gotk4_glib2_LogWriterFunc(GLogLevelFlags, GLogField*, gsize, gpointer);
// extern void callbackDelete(gpointer);
// void _gotk4_glib2_LogFunc(gchar*, GLogLevelFlags, gchar*, gpointer);
import "C"

// LOG_DOMAIN defines the log domain. See Log Domains (#log-domains).
//
// Libraries should define this so that any messages which they log can be
// differentiated from messages from other libraries and application code. But
// be careful not to define it in any public header files.
//
// Log domains must be unique, and it is recommended that they are the
// application or library name, optionally followed by a hyphen and a sub-domain
// name. For example, bloatpad or bloatpad-io.
//
// If undefined, it defaults to the default NULL (or "") log domain; this is not
// advisable, as it cannot be filtered against using the G_MESSAGES_DEBUG
// environment variable.
//
// For example, GTK+ uses this in its Makefile.am:
//
//    AM_CPPFLAGS = -DG_LOG_DOMAIN=\"Gtk\"
//
// Applications can choose to leave it as the default NULL (or "") domain.
// However, defining the domain offers the same advantages as above.
const LOG_DOMAIN = 0

// LOG_FATAL_MASK: GLib log levels that are considered fatal by default.
//
// This is not used if structured logging is enabled; see [Using Structured
// Logging][using-structured-logging].
const LOG_FATAL_MASK = 5

// LOG_LEVEL_USER_SHIFT: log levels below 1<<G_LOG_LEVEL_USER_SHIFT are used by
// GLib. Higher bits can be used for user-defined log levels.
const LOG_LEVEL_USER_SHIFT = 8

// LogWriterOutput: return values from WriterFuncs to indicate whether the given
// log entry was successfully handled by the writer, or whether there was an
// error in handling it (and hence a fallback writer should be used).
//
// If a WriterFunc ignores a log entry, it should return G_LOG_WRITER_HANDLED.
type LogWriterOutput int

const (
	// LogWriterHandled: log writer has handled the log entry.
	LogWriterHandled LogWriterOutput = 1
	// LogWriterUnhandled: log writer could not handle the log entry.
	LogWriterUnhandled LogWriterOutput = 0
)

// String returns the name in string for LogWriterOutput.
func (l LogWriterOutput) String() string {
	switch l {
	case LogWriterHandled:
		return "Handled"
	case LogWriterUnhandled:
		return "Unhandled"
	default:
		return fmt.Sprintf("LogWriterOutput(%d)", l)
	}
}

// LogLevelFlags flags specifying the level of log messages.
//
// It is possible to change how GLib treats messages of the various levels using
// g_log_set_handler() and g_log_set_fatal_mask().
type LogLevelFlags int

const (
	// LogFlagRecursion: internal flag
	LogFlagRecursion LogLevelFlags = 0b1
	// LogFlagFatal: internal flag
	LogFlagFatal LogLevelFlags = 0b10
	// LogLevelError: log level for errors, see g_error(). This level is also
	// used for messages produced by g_assert().
	LogLevelError LogLevelFlags = 0b100
	// LogLevelCritical: log level for critical warning messages, see
	// g_critical(). This level is also used for messages produced by
	// g_return_if_fail() and g_return_val_if_fail().
	LogLevelCritical LogLevelFlags = 0b1000
	// LogLevelWarning: log level for warnings, see g_warning()
	LogLevelWarning LogLevelFlags = 0b10000
	// LogLevelMessage: log level for messages, see g_message()
	LogLevelMessage LogLevelFlags = 0b100000
	// LogLevelInfo: log level for informational messages, see g_info()
	LogLevelInfo LogLevelFlags = 0b1000000
	// LogLevelDebug: log level for debug messages, see g_debug()
	LogLevelDebug LogLevelFlags = 0b10000000
	// LogLevelMask: mask including all log levels
	LogLevelMask LogLevelFlags = -4
)

// String returns the names in string for LogLevelFlags.
func (l LogLevelFlags) String() string {
	if l == 0 {
		return "LogLevelFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(132)

	for l != 0 {
		next := l & (l - 1)
		bit := l - next

		switch bit {
		case LogFlagRecursion:
			builder.WriteString("FlagRecursion|")
		case LogFlagFatal:
			builder.WriteString("FlagFatal|")
		case LogLevelError:
			builder.WriteString("LevelError|")
		case LogLevelCritical:
			builder.WriteString("LevelCritical|")
		case LogLevelWarning:
			builder.WriteString("LevelWarning|")
		case LogLevelMessage:
			builder.WriteString("LevelMessage|")
		case LogLevelInfo:
			builder.WriteString("LevelInfo|")
		case LogLevelDebug:
			builder.WriteString("LevelDebug|")
		case LogLevelMask:
			builder.WriteString("LevelMask|")
		default:
			builder.WriteString(fmt.Sprintf("LogLevelFlags(0b%b)|", bit))
		}

		l = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if l contains other.
func (l LogLevelFlags) Has(other LogLevelFlags) bool {
	return (l & other) == other
}

// LogFunc specifies the prototype of log handler functions.
//
// The default log handler, g_log_default_handler(), automatically appends a
// new-line character to message when printing it. It is advised that any custom
// log handler functions behave similarly, so that logging calls in user code do
// not need modifying to add a new-line character to the message if the log
// handler is changed.
//
// This is not used if structured logging is enabled; see [Using Structured
// Logging][using-structured-logging].
type LogFunc func(logDomain string, logLevel LogLevelFlags, message string)

//export _gotk4_glib2_LogFunc
func _gotk4_glib2_LogFunc(arg0 *C.gchar, arg1 C.GLogLevelFlags, arg2 *C.gchar, arg3 C.gpointer) {
	v := gbox.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var logDomain string       // out
	var logLevel LogLevelFlags // out
	var message string         // out

	logDomain = C.GoString((*C.gchar)(unsafe.Pointer(arg0)))
	logLevel = LogLevelFlags(arg1)
	message = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	fn := v.(LogFunc)
	fn(logDomain, logLevel, message)
}

// LogWriterFunc: writer function for log entries. A log entry is a collection
// of one or more Fields, using the standard [field names from journal
// specification](https://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html).
// See g_log_structured() for more information.
//
// Writer functions must ignore fields which they do not recognise, unless they
// can write arbitrary binary output, as field values may be arbitrary binary.
//
// log_level is guaranteed to be included in fields as the PRIORITY field, but
// is provided separately for convenience of deciding whether or where to output
// the log entry.
//
// Writer functions should return G_LOG_WRITER_HANDLED if they handled the log
// message successfully or if they deliberately ignored it. If there was an
// error handling the message (for example, if the writer function is meant to
// send messages to a remote logging server and there is a network error), it
// should return G_LOG_WRITER_UNHANDLED. This allows writer functions to be
// chained and fall back to simpler handlers in case of failure.
type LogWriterFunc func(logLevel LogLevelFlags, fields []LogField) (logWriterOutput LogWriterOutput)

//export _gotk4_glib2_LogWriterFunc
func _gotk4_glib2_LogWriterFunc(arg0 C.GLogLevelFlags, arg1 *C.GLogField, arg2 C.gsize, arg3 C.gpointer) (cret C.GLogWriterOutput) {
	v := gbox.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var logLevel LogLevelFlags // out
	var fields []LogField      // out

	logLevel = LogLevelFlags(arg0)
	{
		src := unsafe.Slice(arg1, arg2)
		fields = make([]LogField, arg2)
		for i := 0; i < int(arg2); i++ {
			fields[i] = *(*LogField)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
		}
	}

	fn := v.(LogWriterFunc)
	logWriterOutput := fn(logLevel, fields)

	cret = C.GLogWriterOutput(logWriterOutput)

	return cret
}

func AssertWarning(logDomain string, file string, line int, prettyFunction string, expression string) {
	var _arg1 *C.char // out
	var _arg2 *C.char // out
	var _arg3 C.int   // out
	var _arg4 *C.char // out
	var _arg5 *C.char // out

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(logDomain)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(file)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.int(line)
	_arg4 = (*C.char)(unsafe.Pointer(C.CString(prettyFunction)))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = (*C.char)(unsafe.Pointer(C.CString(expression)))
	defer C.free(unsafe.Pointer(_arg5))

	C.g_assert_warning(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(logDomain)
	runtime.KeepAlive(file)
	runtime.KeepAlive(line)
	runtime.KeepAlive(prettyFunction)
	runtime.KeepAlive(expression)
}

// LogDefaultHandler: default log handler set up by GLib;
// g_log_set_default_handler() allows to install an alternate default log
// handler. This is used if no log handler has been set for the particular log
// domain and log level combination. It outputs the message to stderr or stdout
// and if the log level is fatal it calls G_BREAKPOINT(). It automatically
// prints a new-line character after the message, so one does not need to be
// manually included in message.
//
// The behavior of this log handler can be influenced by a number of environment
// variables:
//
// - G_MESSAGES_PREFIXED: A :-separated list of log levels for which messages
// should be prefixed by the program name and PID of the application.
//
// - G_MESSAGES_DEBUG: A space-separated list of log domains for which debug and
// informational messages are printed. By default these messages are not
// printed.
//
// stderr is used for levels G_LOG_LEVEL_ERROR, G_LOG_LEVEL_CRITICAL,
// G_LOG_LEVEL_WARNING and G_LOG_LEVEL_MESSAGE. stdout is used for the rest,
// unless stderr was requested by g_log_writer_default_set_use_stderr().
//
// This has no effect if structured logging is enabled; see [Using Structured
// Logging][using-structured-logging].
func LogDefaultHandler(logDomain string, logLevel LogLevelFlags, message string, unusedData cgo.Handle) {
	var _arg1 *C.gchar         // out
	var _arg2 C.GLogLevelFlags // out
	var _arg3 *C.gchar         // out
	var _arg4 C.gpointer       // out

	if logDomain != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(logDomain)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = C.GLogLevelFlags(logLevel)
	if message != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(message)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	_arg4 = (C.gpointer)(unsafe.Pointer(unusedData))

	C.g_log_default_handler(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(logDomain)
	runtime.KeepAlive(logLevel)
	runtime.KeepAlive(message)
	runtime.KeepAlive(unusedData)
}

// LogRemoveHandler removes the log handler.
//
// This has no effect if structured logging is enabled; see [Using Structured
// Logging][using-structured-logging].
func LogRemoveHandler(logDomain string, handlerId uint) {
	var _arg1 *C.gchar // out
	var _arg2 C.guint  // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(logDomain)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint(handlerId)

	C.g_log_remove_handler(_arg1, _arg2)
	runtime.KeepAlive(logDomain)
	runtime.KeepAlive(handlerId)
}

// LogSetAlwaysFatal sets the message levels which are always fatal, in any log
// domain. When a message with any of these levels is logged the program
// terminates. You can only set the levels defined by GLib to be fatal.
// G_LOG_LEVEL_ERROR is always fatal.
//
// You can also make some message levels fatal at runtime by setting the G_DEBUG
// environment variable (see Running GLib Applications (glib-running.html)).
//
// Libraries should not call this function, as it affects all messages logged by
// a process, including those from other libraries.
//
// Structured log messages (using g_log_structured() and
// g_log_structured_array()) are fatal only if the default log writer is used;
// otherwise it is up to the writer function to determine which log messages are
// fatal. See [Using Structured Logging][using-structured-logging].
func LogSetAlwaysFatal(fatalMask LogLevelFlags) LogLevelFlags {
	var _arg1 C.GLogLevelFlags // out
	var _cret C.GLogLevelFlags // in

	_arg1 = C.GLogLevelFlags(fatalMask)

	_cret = C.g_log_set_always_fatal(_arg1)
	runtime.KeepAlive(fatalMask)

	var _logLevelFlags LogLevelFlags // out

	_logLevelFlags = LogLevelFlags(_cret)

	return _logLevelFlags
}

// LogSetFatalMask sets the log levels which are fatal in the given domain.
// G_LOG_LEVEL_ERROR is always fatal.
//
// This has no effect on structured log messages (using g_log_structured() or
// g_log_structured_array()). To change the fatal behaviour for specific log
// messages, programs must install a custom log writer function using
// g_log_set_writer_func(). See [Using Structured
// Logging][using-structured-logging].
//
// This function is mostly intended to be used with G_LOG_LEVEL_CRITICAL. You
// should typically not set G_LOG_LEVEL_WARNING, G_LOG_LEVEL_MESSAGE,
// G_LOG_LEVEL_INFO or G_LOG_LEVEL_DEBUG as fatal except inside of test
// programs.
func LogSetFatalMask(logDomain string, fatalMask LogLevelFlags) LogLevelFlags {
	var _arg1 *C.gchar         // out
	var _arg2 C.GLogLevelFlags // out
	var _cret C.GLogLevelFlags // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(logDomain)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GLogLevelFlags(fatalMask)

	_cret = C.g_log_set_fatal_mask(_arg1, _arg2)
	runtime.KeepAlive(logDomain)
	runtime.KeepAlive(fatalMask)

	var _logLevelFlags LogLevelFlags // out

	_logLevelFlags = LogLevelFlags(_cret)

	return _logLevelFlags
}

// LogStructuredArray: log a message with structured data. The message will be
// passed through to the log writer set by the application using
// g_log_set_writer_func(). If the message is fatal (i.e. its log level is
// G_LOG_LEVEL_ERROR), the program will be aborted at the end of this function.
//
// See g_log_structured() for more documentation.
//
// This assumes that log_level is already present in fields (typically as the
// PRIORITY field).
func LogStructuredArray(logLevel LogLevelFlags, fields []LogField) {
	var _arg1 C.GLogLevelFlags // out
	var _arg2 *C.GLogField     // out
	var _arg3 C.gsize

	_arg1 = C.GLogLevelFlags(logLevel)
	_arg3 = (C.gsize)(len(fields))
	_arg2 = (*C.GLogField)(C.malloc(C.ulong(len(fields)) * C.ulong(C.sizeof_GLogField)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GLogField)(_arg2), len(fields))
		for i := range fields {
			out[i] = *(*C.GLogField)(gextras.StructNative(unsafe.Pointer((&fields[i]))))
		}
	}

	C.g_log_structured_array(_arg1, _arg2, _arg3)
	runtime.KeepAlive(logLevel)
	runtime.KeepAlive(fields)
}

// LogVariant: log a message with structured data, accepting the data within a
// #GVariant. This version is especially useful for use in other languages, via
// introspection.
//
// The only mandatory item in the fields dictionary is the "MESSAGE" which must
// contain the text shown to the user.
//
// The values in the fields dictionary are likely to be of type String
// (VARIANT_TYPE_STRING). Array of bytes (VARIANT_TYPE_BYTESTRING) is also
// supported. In this case the message is handled as binary and will be
// forwarded to the log writer as such. The size of the array should not be
// higher than G_MAXSSIZE. Otherwise it will be truncated to this size. For
// other types g_variant_print() will be used to convert the value into a
// string.
//
// For more details on its usage and about the parameters, see
// g_log_structured().
func LogVariant(logDomain string, logLevel LogLevelFlags, fields *Variant) {
	var _arg1 *C.gchar         // out
	var _arg2 C.GLogLevelFlags // out
	var _arg3 *C.GVariant      // out

	if logDomain != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(logDomain)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = C.GLogLevelFlags(logLevel)
	_arg3 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(fields)))

	C.g_log_variant(_arg1, _arg2, _arg3)
	runtime.KeepAlive(logDomain)
	runtime.KeepAlive(logLevel)
	runtime.KeepAlive(fields)
}

// LogWriterDefault: format a structured log message and output it to the
// default log destination for the platform. On Linux, this is typically the
// systemd journal, falling back to stdout or stderr if running from the
// terminal or if output is being redirected to a file.
//
// Support for other platform-specific logging mechanisms may be added in
// future. Distributors of GLib may modify this function to impose their own
// (documented) platform-specific log writing policies.
//
// This is suitable for use as a WriterFunc, and is the default writer used if
// no other is set using g_log_set_writer_func().
//
// As with g_log_default_handler(), this function drops debug and informational
// messages unless their log domain (or all) is listed in the space-separated
// G_MESSAGES_DEBUG environment variable.
//
// g_log_writer_default() uses the mask set by g_log_set_always_fatal() to
// determine which messages are fatal. When using a custom writer func instead
// it is up to the writer function to determine which log messages are fatal.
func LogWriterDefault(logLevel LogLevelFlags, fields []LogField, userData cgo.Handle) LogWriterOutput {
	var _arg1 C.GLogLevelFlags // out
	var _arg2 *C.GLogField     // out
	var _arg3 C.gsize
	var _arg4 C.gpointer         // out
	var _cret C.GLogWriterOutput // in

	_arg1 = C.GLogLevelFlags(logLevel)
	_arg3 = (C.gsize)(len(fields))
	_arg2 = (*C.GLogField)(C.malloc(C.ulong(len(fields)) * C.ulong(C.sizeof_GLogField)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GLogField)(_arg2), len(fields))
		for i := range fields {
			out[i] = *(*C.GLogField)(gextras.StructNative(unsafe.Pointer((&fields[i]))))
		}
	}
	_arg4 = (C.gpointer)(unsafe.Pointer(userData))

	_cret = C.g_log_writer_default(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(logLevel)
	runtime.KeepAlive(fields)
	runtime.KeepAlive(userData)

	var _logWriterOutput LogWriterOutput // out

	_logWriterOutput = LogWriterOutput(_cret)

	return _logWriterOutput
}

// LogWriterDefaultSetUseStderr: configure whether the built-in log functions
// (g_log_default_handler() for the old-style API, and both
// g_log_writer_default() and g_log_writer_standard_streams() for the structured
// API) will output all log messages to stderr.
//
// By default, log messages of levels G_LOG_LEVEL_INFO and G_LOG_LEVEL_DEBUG are
// sent to stdout, and other log messages are sent to stderr. This is
// problematic for applications that intend to reserve stdout for structured
// output such as JSON or XML.
//
// This function sets global state. It is not thread-aware, and should be called
// at the very start of a program, before creating any other threads or creating
// objects that could create worker threads of their own.
func LogWriterDefaultSetUseStderr(useStderr bool) {
	var _arg1 C.gboolean // out

	if useStderr {
		_arg1 = C.TRUE
	}

	C.g_log_writer_default_set_use_stderr(_arg1)
	runtime.KeepAlive(useStderr)
}

// LogWriterDefaultWouldDrop: check whether g_log_writer_default() and
// g_log_default_handler() would ignore a message with the given domain and
// level.
//
// As with g_log_default_handler(), this function drops debug and informational
// messages unless their log domain (or all) is listed in the space-separated
// G_MESSAGES_DEBUG environment variable.
//
// This can be used when implementing log writers with the same filtering
// behaviour as the default, but a different destination or output format:
//
//      if (!g_log_writer_default_would_drop (G_LOG_LEVEL_DEBUG, G_LOG_DOMAIN))
//        {
//          gchar *result = expensive_computation (my_object);
//
//          g_debug ("my_object result: s", result);
//          g_free (result);
//        }
func LogWriterDefaultWouldDrop(logLevel LogLevelFlags, logDomain string) bool {
	var _arg1 C.GLogLevelFlags // out
	var _arg2 *C.char          // out
	var _cret C.gboolean       // in

	_arg1 = C.GLogLevelFlags(logLevel)
	if logDomain != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(logDomain)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.g_log_writer_default_would_drop(_arg1, _arg2)
	runtime.KeepAlive(logLevel)
	runtime.KeepAlive(logDomain)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LogWriterFormatFields: format a structured log message as a string suitable
// for outputting to the terminal (or elsewhere). This will include the values
// of all fields it knows how to interpret, which includes MESSAGE and
// GLIB_DOMAIN (see the documentation for g_log_structured()). It does not
// include values from unknown fields.
//
// The returned string does **not** have a trailing new-line character. It is
// encoded in the character set of the current locale, which is not necessarily
// UTF-8.
func LogWriterFormatFields(logLevel LogLevelFlags, fields []LogField, useColor bool) string {
	var _arg1 C.GLogLevelFlags // out
	var _arg2 *C.GLogField     // out
	var _arg3 C.gsize
	var _arg4 C.gboolean // out
	var _cret *C.gchar   // in

	_arg1 = C.GLogLevelFlags(logLevel)
	_arg3 = (C.gsize)(len(fields))
	_arg2 = (*C.GLogField)(C.malloc(C.ulong(len(fields)) * C.ulong(C.sizeof_GLogField)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GLogField)(_arg2), len(fields))
		for i := range fields {
			out[i] = *(*C.GLogField)(gextras.StructNative(unsafe.Pointer((&fields[i]))))
		}
	}
	if useColor {
		_arg4 = C.TRUE
	}

	_cret = C.g_log_writer_format_fields(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(logLevel)
	runtime.KeepAlive(fields)
	runtime.KeepAlive(useColor)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// LogWriterIsJournald: check whether the given output_fd file descriptor is a
// connection to the systemd journal, or something else (like a log file or
// stdout or stderr).
//
// Invalid file descriptors are accepted and return FALSE, which allows for the
// following construct without needing any additional error handling:
//
//    is_journald = g_log_writer_is_journald (fileno (stderr));
func LogWriterIsJournald(outputFd int) bool {
	var _arg1 C.gint     // out
	var _cret C.gboolean // in

	_arg1 = C.gint(outputFd)

	_cret = C.g_log_writer_is_journald(_arg1)
	runtime.KeepAlive(outputFd)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LogWriterJournald: format a structured log message and send it to the systemd
// journal as a set of key–value pairs. All fields are sent to the journal, but
// if a field has length zero (indicating program-specific data) then only its
// key will be sent.
//
// This is suitable for use as a WriterFunc.
//
// If GLib has been compiled without systemd support, this function is still
// defined, but will always return G_LOG_WRITER_UNHANDLED.
func LogWriterJournald(logLevel LogLevelFlags, fields []LogField, userData cgo.Handle) LogWriterOutput {
	var _arg1 C.GLogLevelFlags // out
	var _arg2 *C.GLogField     // out
	var _arg3 C.gsize
	var _arg4 C.gpointer         // out
	var _cret C.GLogWriterOutput // in

	_arg1 = C.GLogLevelFlags(logLevel)
	_arg3 = (C.gsize)(len(fields))
	_arg2 = (*C.GLogField)(C.malloc(C.ulong(len(fields)) * C.ulong(C.sizeof_GLogField)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GLogField)(_arg2), len(fields))
		for i := range fields {
			out[i] = *(*C.GLogField)(gextras.StructNative(unsafe.Pointer((&fields[i]))))
		}
	}
	_arg4 = (C.gpointer)(unsafe.Pointer(userData))

	_cret = C.g_log_writer_journald(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(logLevel)
	runtime.KeepAlive(fields)
	runtime.KeepAlive(userData)

	var _logWriterOutput LogWriterOutput // out

	_logWriterOutput = LogWriterOutput(_cret)

	return _logWriterOutput
}

// LogWriterStandardStreams: format a structured log message and print it to
// either stdout or stderr, depending on its log level. G_LOG_LEVEL_INFO and
// G_LOG_LEVEL_DEBUG messages are sent to stdout, or to stderr if requested by
// g_log_writer_default_set_use_stderr(); all other log levels are sent to
// stderr. Only fields which are understood by this function are included in the
// formatted string which is printed.
//
// If the output stream supports ANSI color escape sequences, they will be used
// in the output.
//
// A trailing new-line character is added to the log message when it is printed.
//
// This is suitable for use as a WriterFunc.
func LogWriterStandardStreams(logLevel LogLevelFlags, fields []LogField, userData cgo.Handle) LogWriterOutput {
	var _arg1 C.GLogLevelFlags // out
	var _arg2 *C.GLogField     // out
	var _arg3 C.gsize
	var _arg4 C.gpointer         // out
	var _cret C.GLogWriterOutput // in

	_arg1 = C.GLogLevelFlags(logLevel)
	_arg3 = (C.gsize)(len(fields))
	_arg2 = (*C.GLogField)(C.malloc(C.ulong(len(fields)) * C.ulong(C.sizeof_GLogField)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GLogField)(_arg2), len(fields))
		for i := range fields {
			out[i] = *(*C.GLogField)(gextras.StructNative(unsafe.Pointer((&fields[i]))))
		}
	}
	_arg4 = (C.gpointer)(unsafe.Pointer(userData))

	_cret = C.g_log_writer_standard_streams(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(logLevel)
	runtime.KeepAlive(fields)
	runtime.KeepAlive(userData)

	var _logWriterOutput LogWriterOutput // out

	_logWriterOutput = LogWriterOutput(_cret)

	return _logWriterOutput
}

// LogWriterSupportsColor: check whether the given output_fd file descriptor
// supports ANSI color escape sequences. If so, they can safely be used when
// formatting log messages.
func LogWriterSupportsColor(outputFd int) bool {
	var _arg1 C.gint     // out
	var _cret C.gboolean // in

	_arg1 = C.gint(outputFd)

	_cret = C.g_log_writer_supports_color(_arg1)
	runtime.KeepAlive(outputFd)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LogField: structure representing a single field in a structured log entry.
// See g_log_structured() for details.
//
// Log fields may contain arbitrary values, including binary with embedded nul
// bytes. If the field contains a string, the string must be UTF-8 encoded and
// have a trailing nul byte. Otherwise, length must be set to a non-negative
// value.
//
// An instance of this type is always passed by reference.
type LogField struct {
	*logField
}

// logField is the struct that's finalized.
type logField struct {
	native *C.GLogField
}

// Key: field name (UTF-8 string)
func (l *LogField) Key() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(l.native.key)))
	return v
}

// LogSetHandler sets the handler used for GLib logging and returns the
// new handler ID. It is a wrapper around g_log_set_handler and
// g_log_set_handler_full.
//
// To detach a log handler, use LogRemoveHandler.
func LogSetHandler(domain string, level LogLevelFlags, f LogFunc) uint {
	var log_domain *C.gchar
	if domain != "" {
		log_domain = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
		defer C.free(unsafe.Pointer(log_domain))
	}

	data := gbox.Assign(f)

	h := C.g_log_set_handler_full(
		log_domain,
		C.GLogLevelFlags(level),
		C.GLogFunc((*[0]byte)(C._gotk4_glib2_LogFunc)),
		C.gpointer(data),
		C.GDestroyNotify((*[0]byte)(C.callbackDelete)),
	)

	return uint(h)
}

// Value returns the field's value.
func (l *LogField) Value() string {
	if l.native.length == -1 {
		return C.GoString((*C.gchar)(unsafe.Pointer(l.native.value)))
	}
	return C.GoStringN((*C.gchar)(unsafe.Pointer(l.native.value)), C.int(l.native.length))
}

// LogSetWriter sets the log writer to the given callback, which should
// take in a list of pair of key-value strings and return true if the
// log has been successfully written. It is a wrapper around
// g_log_set_writer_func.
//
// Note that this function must ONLY be called ONCE. The GLib
// documentation states that it is an error to call it more than once.
func LogSetWriter(f LogWriterFunc) {
	data := gbox.Assign(f)
	C.g_log_set_writer_func(
		C.GLogWriterFunc((*[0]byte)(C._gotk4_glib2_LogWriterFunc)),
		C.gpointer(data),
		C.GDestroyNotify((*[0]byte)(C.callbackDelete)),
	)
}

// LogUseDefaultLogger calls LogUseLogger with Go's default standard
// logger. It is a convenient function for log.Default().
func LogUseDefaultLogger() { LogUseLogger(log.Default()) }

// LogUseLogger calls LogSetWriter with the given Go's standard logger.
// Note that either this or LogSetWriter must only be called once.
// The method will ignore all fields excet for "MESSAGE"; for more
// sophisticated, structured log writing, use LogSetWriter.
// The output format of the logs printed using this function is not
// guaranteed to not change. Users who rely on the format are better off
// using LogSetWriter.
func LogUseLogger(l *log.Logger) { LogSetWriter(LoggerHandler(l)) }

// LoggerHandler creates a new LogWriterFunc that LogUseLogger uses. For
// more information, see LogUseLogger's documentation.
func LoggerHandler(l *log.Logger) LogWriterFunc {
	// Treat Lshortfile and Llongfile the same, because we don't have
	// the full path in codeFile anyway.
	Lfile := l.Flags()&(log.Lshortfile|log.Llongfile) != 0

	// Support $G_MESSAGES_DEBUG.
	debugDomains := make(map[string]struct{})
	for _, debugDomain := range strings.Fields(os.Getenv("G_MESSAGES_DEBUG")) {
		debugDomains[debugDomain] = struct{}{}
	}

	// Special case: G_MESSAGES_DEBUG=all.
	_, debugAll := debugDomains["all"]

	return func(lvl LogLevelFlags, fields []LogField) LogWriterOutput {
		var message, codeFile, codeLine, codeFunc string
		domain := "GLib (no domain)"

		for _, field := range fields {
			if !Lfile {
				switch field.Key() {
				case "MESSAGE":
					message = field.Value()
				case "GLIB_DOMAIN":
					domain = field.Value()
				}
				// Skip setting code* if we don't have to.
				continue
			}

			switch field.Key() {
			case "MESSAGE":
				message = field.Value()
			case "CODE_FILE":
				codeFile = field.Value()
			case "CODE_LINE":
				codeLine = field.Value()
			case "CODE_FUNC":
				codeFunc = field.Value()
			case "GLIB_DOMAIN":
				domain = field.Value()
			}
		}

		if !debugAll && (lvl&LogLevelDebug != 0) && domain != "" {
			if _, ok := debugDomains[domain]; !ok {
				return LogWriterHandled
			}
		}

		f := l.Printf
		if lvl&LogFlagFatal != 0 {
			f = l.Fatalf
		}

		if !Lfile || (codeFile == "" && codeLine == "") {
			f("%s: %s: %s", lvl, domain, message)
			return LogWriterHandled
		}

		if codeFunc == "" {
			f("%s: %s: %s:%s: %s", lvl, domain, codeFile, codeLine, message)
			return LogWriterHandled
		}

		f("%s: %s: %s:%s (%s): %s", lvl, domain, codeFile, codeLine, codeFunc, message)
		return LogWriterHandled
	}
}
