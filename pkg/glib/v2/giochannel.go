// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_io_channel_get_type()), F: marshalIOChannel},
	})
}

const WIN32_MSG_HANDLE = 19981206

// IOChannelError: error codes returned by OChannel operations.
type IOChannelError C.gint

const (
	// IOChannelErrorFbig: file too large.
	IOChannelErrorFbig IOChannelError = iota
	// IOChannelErrorInval: invalid argument.
	IOChannelErrorInval
	// IOChannelErrorIO: IO error.
	IOChannelErrorIO
	// IOChannelErrorIsdir: file is a directory.
	IOChannelErrorIsdir
	// IOChannelErrorNospc: no space left on device.
	IOChannelErrorNospc
	// IOChannelErrorNxio: no such device or address.
	IOChannelErrorNxio
	// IOChannelErrorOverflow: value too large for defined datatype.
	IOChannelErrorOverflow
	// IOChannelErrorPipe: broken pipe.
	IOChannelErrorPipe
	// IOChannelErrorFailed: some other error.
	IOChannelErrorFailed
)

// String returns the name in string for IOChannelError.
func (i IOChannelError) String() string {
	switch i {
	case IOChannelErrorFbig:
		return "Fbig"
	case IOChannelErrorInval:
		return "Inval"
	case IOChannelErrorIO:
		return "IO"
	case IOChannelErrorIsdir:
		return "Isdir"
	case IOChannelErrorNospc:
		return "Nospc"
	case IOChannelErrorNxio:
		return "Nxio"
	case IOChannelErrorOverflow:
		return "Overflow"
	case IOChannelErrorPipe:
		return "Pipe"
	case IOChannelErrorFailed:
		return "Failed"
	default:
		return fmt.Sprintf("IOChannelError(%d)", i)
	}
}

// IOError is only used by the deprecated functions g_io_channel_read(),
// g_io_channel_write(), and g_io_channel_seek().
type IOError C.gint

const (
	// IOErrorNone: no error.
	IOErrorNone IOError = iota
	// IOErrorAgain: EAGAIN error occurred.
	IOErrorAgain
	// IOErrorInval: EINVAL error occurred.
	IOErrorInval
	// IOErrorUnknown: another error occurred.
	IOErrorUnknown
)

// String returns the name in string for IOError.
func (i IOError) String() string {
	switch i {
	case IOErrorNone:
		return "None"
	case IOErrorAgain:
		return "Again"
	case IOErrorInval:
		return "Inval"
	case IOErrorUnknown:
		return "Unknown"
	default:
		return fmt.Sprintf("IOError(%d)", i)
	}
}

// IOStatus statuses returned by most of the OFuncs functions.
type IOStatus C.gint

const (
	// IOStatusError: error occurred.
	IOStatusError IOStatus = iota
	// IOStatusNormal: success.
	IOStatusNormal
	// IOStatusEOF: end of file.
	IOStatusEOF
	// IOStatusAgain: resource temporarily unavailable.
	IOStatusAgain
)

// String returns the name in string for IOStatus.
func (i IOStatus) String() string {
	switch i {
	case IOStatusError:
		return "Error"
	case IOStatusNormal:
		return "Normal"
	case IOStatusEOF:
		return "EOF"
	case IOStatusAgain:
		return "Again"
	default:
		return fmt.Sprintf("IOStatus(%d)", i)
	}
}

// SeekType: enumeration specifying the base position for a
// g_io_channel_seek_position() operation.
type SeekType C.gint

const (
	// SeekCur: current position in the file.
	SeekCur SeekType = iota
	// SeekSet: start of the file.
	SeekSet
	// SeekEnd: end of the file.
	SeekEnd
)

// String returns the name in string for SeekType.
func (s SeekType) String() string {
	switch s {
	case SeekCur:
		return "Cur"
	case SeekSet:
		return "Set"
	case SeekEnd:
		return "End"
	default:
		return fmt.Sprintf("SeekType(%d)", s)
	}
}

// IOFlags specifies properties of a OChannel. Some of the flags can only be
// read with g_io_channel_get_flags(), but not changed with
// g_io_channel_set_flags().
type IOFlags C.guint

const (
	// IOFlagAppend turns on append mode, corresponds to O_APPEND (see the
	// documentation of the UNIX open() syscall).
	IOFlagAppend IOFlags = 0b1
	// IOFlagNonblock turns on nonblocking mode, corresponds to
	// O_NONBLOCK/O_NDELAY (see the documentation of the UNIX open() syscall).
	IOFlagNonblock IOFlags = 0b10
	// IOFlagIsReadable indicates that the io channel is readable. This flag
	// cannot be changed.
	IOFlagIsReadable IOFlags = 0b100
	// IOFlagIsWritable indicates that the io channel is writable. This flag
	// cannot be changed.
	IOFlagIsWritable IOFlags = 0b1000
	// IOFlagIsWriteable: misspelled version of G_IO_FLAG_IS_WRITABLE that
	// existed before the spelling was fixed in GLib 2.30. It is kept here for
	// compatibility reasons. Deprecated since 2.30.
	IOFlagIsWriteable IOFlags = 0b1000
	// IOFlagIsSeekable indicates that the io channel is seekable, i.e. that
	// g_io_channel_seek_position() can be used on it. This flag cannot be
	// changed.
	IOFlagIsSeekable IOFlags = 0b10000
	// IOFlagMask: mask that specifies all the valid flags.
	IOFlagMask IOFlags = 0b11111
	// IOFlagGetMask: mask of the flags that are returned from
	// g_io_channel_get_flags().
	IOFlagGetMask IOFlags = 0b11111
	// IOFlagSetMask: mask of the flags that the user can modify with
	// g_io_channel_set_flags().
	IOFlagSetMask IOFlags = 0b11
)

// String returns the names in string for IOFlags.
func (i IOFlags) String() string {
	if i == 0 {
		return "IOFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(135)

	for i != 0 {
		next := i & (i - 1)
		bit := i - next

		switch bit {
		case IOFlagAppend:
			builder.WriteString("Append|")
		case IOFlagNonblock:
			builder.WriteString("Nonblock|")
		case IOFlagIsReadable:
			builder.WriteString("IsReadable|")
		case IOFlagIsWritable:
			builder.WriteString("IsWritable|")
		case IOFlagIsSeekable:
			builder.WriteString("IsSeekable|")
		case IOFlagMask:
			builder.WriteString("Mask|")
		case IOFlagSetMask:
			builder.WriteString("SetMask|")
		default:
			builder.WriteString(fmt.Sprintf("IOFlags(0b%b)|", bit))
		}

		i = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if i contains other.
func (i IOFlags) Has(other IOFlags) bool {
	return (i & other) == other
}

// IOCreateWatch creates a #GSource that's dispatched when condition is met for
// the given channel. For example, if condition is IO_IN, the source will be
// dispatched when there's data available for reading.
//
// The callback function invoked by the #GSource should be added with
// g_source_set_callback(), but it has type OFunc (not Func).
//
// g_io_add_watch() is a simpler interface to this same functionality, for the
// case where you want to add the source to the default main loop context at the
// default priority.
//
// On Windows, polling a #GSource created to watch a channel for a socket puts
// the socket in non-blocking mode. This is a side-effect of the implementation
// and unavoidable.
//
// The function takes the following parameters:
//
//    - channel to watch.
//    - condition conditions to watch for.
//
// The function returns the following values:
//
//    - source: new #GSource.
//
func IOCreateWatch(channel *IOChannel, condition IOCondition) *Source {
	var _arg1 *C.GIOChannel  // out
	var _arg2 C.GIOCondition // out
	var _cret *C.GSource     // in

	_arg1 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))
	_arg2 = C.GIOCondition(condition)

	_cret = C.g_io_create_watch(_arg1, _arg2)
	runtime.KeepAlive(channel)
	runtime.KeepAlive(condition)

	var _source *Source // out

	_source = (*Source)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_source)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_source_unref((*C.GSource)(intern.C))
		},
	)

	return _source
}

// IOChannel: data structure representing an IO Channel. The fields should be
// considered private and should only be accessed with the following functions.
//
// An instance of this type is always passed by reference.
type IOChannel struct {
	*ioChannel
}

// ioChannel is the struct that's finalized.
type ioChannel struct {
	native *C.GIOChannel
}

func marshalIOChannel(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &IOChannel{&ioChannel{(*C.GIOChannel)(b)}}, nil
}

// NewIOChannelFile constructs a struct IOChannel.
func NewIOChannelFile(filename string, mode string) (*IOChannel, error) {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _cret *C.GIOChannel // in
	var _cerr *C.GError     // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(mode)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_io_channel_new_file(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(mode)

	var _ioChannel *IOChannel // out
	var _goerr error          // out

	_ioChannel = (*IOChannel)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_ioChannel)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_io_channel_unref((*C.GIOChannel)(intern.C))
		},
	)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _ioChannel, _goerr
}

// NewIOChannelUnix constructs a struct IOChannel.
func NewIOChannelUnix(fd int) *IOChannel {
	var _arg1 C.int         // out
	var _cret *C.GIOChannel // in

	_arg1 = C.int(fd)

	_cret = C.g_io_channel_unix_new(_arg1)
	runtime.KeepAlive(fd)

	var _ioChannel *IOChannel // out

	_ioChannel = (*IOChannel)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_ioChannel)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_io_channel_unref((*C.GIOChannel)(intern.C))
		},
	)

	return _ioChannel
}

// Close an IO channel. Any pending data to be written will be flushed, ignoring
// errors. The channel will not be freed until the last reference is dropped
// using g_io_channel_unref().
//
// Deprecated: Use g_io_channel_shutdown() instead.
func (channel *IOChannel) Close() {
	var _arg0 *C.GIOChannel // out

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))

	C.g_io_channel_close(_arg0)
	runtime.KeepAlive(channel)
}

// Flush flushes the write buffer for the GIOChannel.
//
// The function returns the following values:
//
//    - ioStatus status of the operation: One of IO_STATUS_NORMAL,
//      IO_STATUS_AGAIN, or IO_STATUS_ERROR.
//
func (channel *IOChannel) Flush() (IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))

	_cret = C.g_io_channel_flush(_arg0, &_cerr)
	runtime.KeepAlive(channel)

	var _ioStatus IOStatus // out
	var _goerr error       // out

	_ioStatus = IOStatus(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _ioStatus, _goerr
}

// BufferCondition: this function returns a OCondition depending on whether
// there is data to be read/space to write data in the internal buffers in the
// OChannel. Only the flags G_IO_IN and G_IO_OUT may be set.
//
// The function returns the following values:
//
//    - ioCondition: OCondition.
//
func (channel *IOChannel) BufferCondition() IOCondition {
	var _arg0 *C.GIOChannel  // out
	var _cret C.GIOCondition // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))

	_cret = C.g_io_channel_get_buffer_condition(_arg0)
	runtime.KeepAlive(channel)

	var _ioCondition IOCondition // out

	_ioCondition = IOCondition(_cret)

	return _ioCondition
}

// BufferSize gets the buffer size.
//
// The function returns the following values:
//
//    - gsize: size of the buffer.
//
func (channel *IOChannel) BufferSize() uint {
	var _arg0 *C.GIOChannel // out
	var _cret C.gsize       // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))

	_cret = C.g_io_channel_get_buffer_size(_arg0)
	runtime.KeepAlive(channel)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// Buffered returns whether channel is buffered.
//
// The function returns the following values:
//
//    - ok: TRUE if the channel is buffered.
//
func (channel *IOChannel) Buffered() bool {
	var _arg0 *C.GIOChannel // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))

	_cret = C.g_io_channel_get_buffered(_arg0)
	runtime.KeepAlive(channel)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Encoding gets the encoding for the input/output of the channel. The internal
// encoding is always UTF-8. The encoding NULL makes the channel safe for binary
// data.
//
// The function returns the following values:
//
//    - utf8: string containing the encoding, this string is owned by GLib and
//      must not be freed.
//
func (channel *IOChannel) Encoding() string {
	var _arg0 *C.GIOChannel // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))

	_cret = C.g_io_channel_get_encoding(_arg0)
	runtime.KeepAlive(channel)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Flags gets the current flags for a OChannel, including read-only flags such
// as G_IO_FLAG_IS_READABLE.
//
// The values of the flags G_IO_FLAG_IS_READABLE and G_IO_FLAG_IS_WRITABLE are
// cached for internal use by the channel when it is created. If they should
// change at some later point (e.g. partial shutdown of a socket with the UNIX
// shutdown() function), the user should immediately call
// g_io_channel_get_flags() to update the internal values of these flags.
//
// The function returns the following values:
//
//    - ioFlags flags which are set on the channel.
//
func (channel *IOChannel) Flags() IOFlags {
	var _arg0 *C.GIOChannel // out
	var _cret C.GIOFlags    // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))

	_cret = C.g_io_channel_get_flags(_arg0)
	runtime.KeepAlive(channel)

	var _ioFlags IOFlags // out

	_ioFlags = IOFlags(_cret)

	return _ioFlags
}

// LineTerm: this returns the string that OChannel uses to determine where in
// the file a line break occurs. A value of NULL indicates autodetection.
//
// The function takes the following parameters:
//
//    - length: location to return the length of the line terminator.
//
// The function returns the following values:
//
//    - utf8: line termination string. This value is owned by GLib and must not
//      be freed.
//
func (channel *IOChannel) LineTerm(length *int) string {
	var _arg0 *C.GIOChannel // out
	var _arg1 *C.gint       // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))
	_arg1 = (*C.gint)(unsafe.Pointer(length))

	_cret = C.g_io_channel_get_line_term(_arg0, _arg1)
	runtime.KeepAlive(channel)
	runtime.KeepAlive(length)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Init initializes a OChannel struct.
//
// This is called by each of the above functions when creating a OChannel, and
// so is not often needed by the application programmer (unless you are creating
// a new type of OChannel).
func (channel *IOChannel) Init() {
	var _arg0 *C.GIOChannel // out

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))

	C.g_io_channel_init(_arg0)
	runtime.KeepAlive(channel)
}

// ReadChars: replacement for g_io_channel_read() with the new API.
//
// The function takes the following parameters:
//
//    - buf: a buffer to read data into.
//
// The function returns the following values:
//
//    - bytesRead (optional): number of bytes read. This may be zero even on
//      success if count < 6 and the channel's encoding is non-NULL. This
//      indicates that the next UTF-8 character is too wide for the buffer.
//    - ioStatus status of the operation.
//
func (channel *IOChannel) ReadChars(buf []byte) (uint, IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 *C.gchar      // out
	var _arg2 C.gsize
	var _arg3 C.gsize     // in
	var _cret C.GIOStatus // in
	var _cerr *C.GError   // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))
	_arg2 = (C.gsize)(len(buf))
	_arg1 = (*C.gchar)(C.CBytes(buf))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_io_channel_read_chars(_arg0, _arg1, _arg2, &_arg3, &_cerr)
	runtime.KeepAlive(channel)
	runtime.KeepAlive(buf)

	var _bytesRead uint    // out
	var _ioStatus IOStatus // out
	var _goerr error       // out

	_bytesRead = uint(_arg3)
	_ioStatus = IOStatus(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _bytesRead, _ioStatus, _goerr
}

// ReadLine reads a line, including the terminating character(s), from a
// OChannel into a newly-allocated string. str_return will contain allocated
// memory if the return is G_IO_STATUS_NORMAL.
//
// The function returns the following values:
//
//    - strReturn: line read from the OChannel, including the line terminator.
//      This data should be freed with g_free() when no longer needed. This is a
//      nul-terminated string. If a length of zero is returned, this will be NULL
//      instead.
//    - length (optional): location to store length of the read data, or NULL.
//    - terminatorPos (optional): location to store position of line terminator,
//      or NULL.
//    - ioStatus status of the operation.
//
func (channel *IOChannel) ReadLine() (strReturn string, length uint, terminatorPos uint, ioStatus IOStatus, goerr error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 *C.gchar      // in
	var _arg2 C.gsize       // in
	var _arg3 C.gsize       // in
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))

	_cret = C.g_io_channel_read_line(_arg0, &_arg1, &_arg2, &_arg3, &_cerr)
	runtime.KeepAlive(channel)

	var _strReturn string   // out
	var _length uint        // out
	var _terminatorPos uint // out
	var _ioStatus IOStatus  // out
	var _goerr error        // out

	_strReturn = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	defer C.free(unsafe.Pointer(_arg1))
	_length = uint(_arg2)
	_terminatorPos = uint(_arg3)
	_ioStatus = IOStatus(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _strReturn, _length, _terminatorPos, _ioStatus, _goerr
}

// ReadToEnd reads all the remaining data from the file.
//
// The function returns the following values:
//
//    - strReturn: location to store a pointer to a string holding the remaining
//      data in the OChannel. This data should be freed with g_free() when no
//      longer needed. This data is terminated by an extra nul character, but
//      there may be other nuls in the intervening data.
//    - ioStatus: G_IO_STATUS_NORMAL on success. This function never returns
//      G_IO_STATUS_EOF.
//
func (channel *IOChannel) ReadToEnd() ([]byte, IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 *C.gchar      // in
	var _arg2 C.gsize       // in
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))

	_cret = C.g_io_channel_read_to_end(_arg0, &_arg1, &_arg2, &_cerr)
	runtime.KeepAlive(channel)

	var _strReturn []byte  // out
	var _ioStatus IOStatus // out
	var _goerr error       // out

	defer C.free(unsafe.Pointer(_arg1))
	_strReturn = make([]byte, _arg2)
	copy(_strReturn, unsafe.Slice((*byte)(unsafe.Pointer(_arg1)), _arg2))
	_ioStatus = IOStatus(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _strReturn, _ioStatus, _goerr
}

// ReadUnichar reads a Unicode character from channel. This function cannot be
// called on a channel with NULL encoding.
//
// The function returns the following values:
//
//    - thechar: location to return a character.
//    - ioStatus: OStatus.
//
func (channel *IOChannel) ReadUnichar() (uint32, IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.gunichar    // in
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))

	_cret = C.g_io_channel_read_unichar(_arg0, &_arg1, &_cerr)
	runtime.KeepAlive(channel)

	var _thechar uint32    // out
	var _ioStatus IOStatus // out
	var _goerr error       // out

	_thechar = uint32(_arg1)
	_ioStatus = IOStatus(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _thechar, _ioStatus, _goerr
}

// Seek sets the current position in the OChannel, similar to the standard
// library function fseek().
//
// Deprecated: Use g_io_channel_seek_position() instead.
//
// The function takes the following parameters:
//
//    - offset: offset, in bytes, which is added to the position specified by
//      type.
//    - typ: position in the file, which can be G_SEEK_CUR (the current
//      position), G_SEEK_SET (the start of the file), or G_SEEK_END (the end of
//      the file).
//
// The function returns the following values:
//
//    - ioError: G_IO_ERROR_NONE if the operation was successful.
//
func (channel *IOChannel) Seek(offset int64, typ SeekType) IOError {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.gint64      // out
	var _arg2 C.GSeekType   // out
	var _cret C.GIOError    // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))
	_arg1 = C.gint64(offset)
	_arg2 = C.GSeekType(typ)

	_cret = C.g_io_channel_seek(_arg0, _arg1, _arg2)
	runtime.KeepAlive(channel)
	runtime.KeepAlive(offset)
	runtime.KeepAlive(typ)

	var _ioError IOError // out

	_ioError = IOError(_cret)

	return _ioError
}

// SeekPosition: replacement for g_io_channel_seek() with the new API.
//
// The function takes the following parameters:
//
//    - offset in bytes from the position specified by type.
//    - typ The type G_SEEK_CUR is only allowed in those cases where a call to
//      g_io_channel_set_encoding () is allowed. See the documentation for
//      g_io_channel_set_encoding () for details.
//
// The function returns the following values:
//
//    - ioStatus status of the operation.
//
func (channel *IOChannel) SeekPosition(offset int64, typ SeekType) (IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.gint64      // out
	var _arg2 C.GSeekType   // out
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))
	_arg1 = C.gint64(offset)
	_arg2 = C.GSeekType(typ)

	_cret = C.g_io_channel_seek_position(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(channel)
	runtime.KeepAlive(offset)
	runtime.KeepAlive(typ)

	var _ioStatus IOStatus // out
	var _goerr error       // out

	_ioStatus = IOStatus(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _ioStatus, _goerr
}

// SetBufferSize sets the buffer size.
//
// The function takes the following parameters:
//
//    - size of the buffer, or 0 to let GLib pick a good size.
//
func (channel *IOChannel) SetBufferSize(size uint) {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.gsize       // out

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))
	_arg1 = C.gsize(size)

	C.g_io_channel_set_buffer_size(_arg0, _arg1)
	runtime.KeepAlive(channel)
	runtime.KeepAlive(size)
}

// SetBuffered: buffering state can only be set if the channel's encoding is
// NULL. For any other encoding, the channel must be buffered.
//
// A buffered channel can only be set unbuffered if the channel's internal
// buffers have been flushed. Newly created channels or channels which have
// returned G_IO_STATUS_EOF not require such a flush. For write-only channels, a
// call to g_io_channel_flush () is sufficient. For all other channels, the
// buffers may be flushed by a call to g_io_channel_seek_position (). This
// includes the possibility of seeking with seek type G_SEEK_CUR and an offset
// of zero. Note that this means that socket-based channels cannot be set
// unbuffered once they have had data read from them.
//
// On unbuffered channels, it is safe to mix read and write calls from the new
// and old APIs, if this is necessary for maintaining old code.
//
// The default state of the channel is buffered.
//
// The function takes the following parameters:
//
//    - buffered: whether to set the channel buffered or unbuffered.
//
func (channel *IOChannel) SetBuffered(buffered bool) {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))
	if buffered {
		_arg1 = C.TRUE
	}

	C.g_io_channel_set_buffered(_arg0, _arg1)
	runtime.KeepAlive(channel)
	runtime.KeepAlive(buffered)
}

// SetEncoding sets the encoding for the input/output of the channel. The
// internal encoding is always UTF-8. The default encoding for the external file
// is UTF-8.
//
// The encoding NULL is safe to use with binary data.
//
// The encoding can only be set if one of the following conditions is true:
//
// - The channel was just created, and has not been written to or read from yet.
//
// - The channel is write-only.
//
// - The channel is a file, and the file pointer was just repositioned by a call
// to g_io_channel_seek_position(). (This flushes all the internal buffers.)
//
// - The current encoding is NULL or UTF-8.
//
// - One of the (new API) read functions has just returned G_IO_STATUS_EOF (or,
// in the case of g_io_channel_read_to_end(), G_IO_STATUS_NORMAL).
//
// - One of the functions g_io_channel_read_chars() or
// g_io_channel_read_unichar() has returned G_IO_STATUS_AGAIN or
// G_IO_STATUS_ERROR. This may be useful in the case of
// G_CONVERT_ERROR_ILLEGAL_SEQUENCE. Returning one of these statuses from
// g_io_channel_read_line(), g_io_channel_read_line_string(), or
// g_io_channel_read_to_end() does not guarantee that the encoding can be
// changed.
//
// Channels which do not meet one of the above conditions cannot call
// g_io_channel_seek_position() with an offset of G_SEEK_CUR, and, if they are
// "seekable", cannot call g_io_channel_write_chars() after calling one of the
// API "read" functions.
//
// The function takes the following parameters:
//
//    - encoding (optional) type.
//
// The function returns the following values:
//
//    - ioStatus: G_IO_STATUS_NORMAL if the encoding was successfully set.
//
func (channel *IOChannel) SetEncoding(encoding string) (IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 *C.gchar      // out
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))
	if encoding != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(encoding)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.g_io_channel_set_encoding(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(channel)
	runtime.KeepAlive(encoding)

	var _ioStatus IOStatus // out
	var _goerr error       // out

	_ioStatus = IOStatus(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _ioStatus, _goerr
}

// SetFlags sets the (writeable) flags in channel to (flags &
// G_IO_FLAG_SET_MASK).
//
// The function takes the following parameters:
//
//    - flags to set on the IO channel.
//
// The function returns the following values:
//
//    - ioStatus status of the operation.
//
func (channel *IOChannel) SetFlags(flags IOFlags) (IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.GIOFlags    // out
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))
	_arg1 = C.GIOFlags(flags)

	_cret = C.g_io_channel_set_flags(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(channel)
	runtime.KeepAlive(flags)

	var _ioStatus IOStatus // out
	var _goerr error       // out

	_ioStatus = IOStatus(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _ioStatus, _goerr
}

// SetLineTerm: this sets the string that OChannel uses to determine where in
// the file a line break occurs.
//
// The function takes the following parameters:
//
//    - lineTerm (optional): line termination string. Use NULL for autodetect.
//      Autodetection breaks on "\n", "\r\n", "\r", "\0", and the Unicode
//      paragraph separator. Autodetection should not be used for anything other
//      than file-based channels.
//    - length of the termination string. If -1 is passed, the string is assumed
//      to be nul-terminated. This option allows termination strings with
//      embedded nuls.
//
func (channel *IOChannel) SetLineTerm(lineTerm string, length int) {
	var _arg0 *C.GIOChannel // out
	var _arg1 *C.gchar      // out
	var _arg2 C.gint        // out

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))
	if lineTerm != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(lineTerm)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = C.gint(length)

	C.g_io_channel_set_line_term(_arg0, _arg1, _arg2)
	runtime.KeepAlive(channel)
	runtime.KeepAlive(lineTerm)
	runtime.KeepAlive(length)
}

// Shutdown: close an IO channel. Any pending data to be written will be flushed
// if flush is TRUE. The channel will not be freed until the last reference is
// dropped using g_io_channel_unref().
//
// The function takes the following parameters:
//
//    - flush: if TRUE, flush pending.
//
// The function returns the following values:
//
//    - ioStatus status of the operation.
//
func (channel *IOChannel) Shutdown(flush bool) (IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.gboolean    // out
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))
	if flush {
		_arg1 = C.TRUE
	}

	_cret = C.g_io_channel_shutdown(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(channel)
	runtime.KeepAlive(flush)

	var _ioStatus IOStatus // out
	var _goerr error       // out

	_ioStatus = IOStatus(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _ioStatus, _goerr
}

// UnixGetFd returns the file descriptor of the OChannel.
//
// On Windows this function returns the file descriptor or socket of the
// OChannel.
//
// The function returns the following values:
//
//    - gint: file descriptor of the OChannel.
//
func (channel *IOChannel) UnixGetFd() int {
	var _arg0 *C.GIOChannel // out
	var _cret C.gint        // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))

	_cret = C.g_io_channel_unix_get_fd(_arg0)
	runtime.KeepAlive(channel)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Write writes data to a OChannel.
//
// Deprecated: Use g_io_channel_write_chars() instead.
//
// The function takes the following parameters:
//
//    - buf: buffer containing the data to write.
//    - count: number of bytes to write.
//    - bytesWritten: number of bytes actually written.
//
// The function returns the following values:
//
//    - ioError: G_IO_ERROR_NONE if the operation was successful.
//
func (channel *IOChannel) Write(buf string, count uint, bytesWritten *uint) IOError {
	var _arg0 *C.GIOChannel // out
	var _arg1 *C.gchar      // out
	var _arg2 C.gsize       // out
	var _arg3 *C.gsize      // out
	var _cret C.GIOError    // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(buf)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gsize(count)
	_arg3 = (*C.gsize)(unsafe.Pointer(bytesWritten))

	_cret = C.g_io_channel_write(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(channel)
	runtime.KeepAlive(buf)
	runtime.KeepAlive(count)
	runtime.KeepAlive(bytesWritten)

	var _ioError IOError // out

	_ioError = IOError(_cret)

	return _ioError
}

// WriteChars: replacement for g_io_channel_write() with the new API.
//
// On seekable channels with encodings other than NULL or UTF-8, generic mixing
// of reading and writing is not allowed. A call to g_io_channel_write_chars ()
// may only be made on a channel from which data has been read in the cases
// described in the documentation for g_io_channel_set_encoding ().
//
// The function takes the following parameters:
//
//    - buf: buffer to write data from.
//    - count: size of the buffer. If -1, the buffer is taken to be a
//      nul-terminated string.
//
// The function returns the following values:
//
//    - bytesWritten: number of bytes written. This can be nonzero even if the
//      return value is not G_IO_STATUS_NORMAL. If the return value is
//      G_IO_STATUS_NORMAL and the channel is blocking, this will always be equal
//      to count if count >= 0.
//    - ioStatus status of the operation.
//
func (channel *IOChannel) WriteChars(buf string, count int) (uint, IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 *C.gchar      // out
	var _arg2 C.gssize      // out
	var _arg3 C.gsize       // in
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))
	_arg1 = (*C.gchar)(C.calloc(C.size_t((len(buf) + 1)), C.size_t(C.sizeof_gchar)))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(_arg1)), len(buf)), buf)
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(count)

	_cret = C.g_io_channel_write_chars(_arg0, _arg1, _arg2, &_arg3, &_cerr)
	runtime.KeepAlive(channel)
	runtime.KeepAlive(buf)
	runtime.KeepAlive(count)

	var _bytesWritten uint // out
	var _ioStatus IOStatus // out
	var _goerr error       // out

	_bytesWritten = uint(_arg3)
	_ioStatus = IOStatus(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _bytesWritten, _ioStatus, _goerr
}

// WriteUnichar writes a Unicode character to channel. This function cannot be
// called on a channel with NULL encoding.
//
// The function takes the following parameters:
//
//    - thechar: character.
//
// The function returns the following values:
//
//    - ioStatus: OStatus.
//
func (channel *IOChannel) WriteUnichar(thechar uint32) (IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.gunichar    // out
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(channel)))
	_arg1 = C.gunichar(thechar)

	_cret = C.g_io_channel_write_unichar(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(channel)
	runtime.KeepAlive(thechar)

	var _ioStatus IOStatus // out
	var _goerr error       // out

	_ioStatus = IOStatus(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _ioStatus, _goerr
}

// IOChannelErrorFromErrno converts an errno error number to a OChannelError.
//
// The function takes the following parameters:
//
//    - en: errno error number, e.g. EINVAL.
//
// The function returns the following values:
//
//    - ioChannelError error number, e.g. G_IO_CHANNEL_ERROR_INVAL.
//
func IOChannelErrorFromErrno(en int) IOChannelError {
	var _arg1 C.gint            // out
	var _cret C.GIOChannelError // in

	_arg1 = C.gint(en)

	_cret = C.g_io_channel_error_from_errno(_arg1)
	runtime.KeepAlive(en)

	var _ioChannelError IOChannelError // out

	_ioChannelError = IOChannelError(_cret)

	return _ioChannelError
}

// IOFuncs: table of functions used to handle different types of OChannel in a
// generic way.
//
// An instance of this type is always passed by reference.
type IOFuncs struct {
	*ioFuncs
}

// ioFuncs is the struct that's finalized.
type ioFuncs struct {
	native *C.GIOFuncs
}
