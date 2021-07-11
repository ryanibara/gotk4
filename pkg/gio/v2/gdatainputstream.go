// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
//
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_data_input_stream_get_type()), F: marshalDataInputStreamer},
	})
}

// DataInputStreamer describes DataInputStream's methods.
type DataInputStreamer interface {
	// ByteOrder gets the byte order for the data input stream.
	ByteOrder() DataStreamByteOrder
	// NewlineType gets the current newline type for the @stream.
	NewlineType() DataStreamNewlineType
	// ReadByte reads an unsigned 8-bit/1-byte value from @stream.
	ReadByte(cancellable Cancellabler) (byte, error)
	// ReadInt16 reads a 16-bit/2-byte value from @stream.
	ReadInt16(cancellable Cancellabler) (int16, error)
	// ReadInt32 reads a signed 32-bit/4-byte value from @stream.
	ReadInt32(cancellable Cancellabler) (int32, error)
	// ReadInt64 reads a 64-bit/8-byte value from @stream.
	ReadInt64(cancellable Cancellabler) (int64, error)
	// ReadLine reads a line from the data input stream.
	ReadLine(cancellable Cancellabler) (uint, []byte, error)
	// ReadLineAsync asynchronous version of g_data_input_stream_read_line().
	ReadLineAsync(ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback)
	// ReadLineFinish: finish an asynchronous call started by
	// g_data_input_stream_read_line_async().
	ReadLineFinish(result AsyncResulter) (uint, []byte, error)
	// ReadLineFinishUTF8: finish an asynchronous call started by
	// g_data_input_stream_read_line_async().
	ReadLineFinishUTF8(result AsyncResulter) (uint, string, error)
	// ReadLineUTF8 reads a UTF-8 encoded line from the data input stream.
	ReadLineUTF8(cancellable Cancellabler) (uint, string, error)
	// ReadUint16 reads an unsigned 16-bit/2-byte value from @stream.
	ReadUint16(cancellable Cancellabler) (uint16, error)
	// ReadUint32 reads an unsigned 32-bit/4-byte value from @stream.
	ReadUint32(cancellable Cancellabler) (uint32, error)
	// ReadUint64 reads an unsigned 64-bit/8-byte value from @stream.
	ReadUint64(cancellable Cancellabler) (uint64, error)
	// ReadUntil reads a string from the data input stream, up to the first
	// occurrence of any of the stop characters.
	ReadUntil(stopChars string, cancellable Cancellabler) (uint, string, error)
	// ReadUntilAsync asynchronous version of g_data_input_stream_read_until().
	ReadUntilAsync(stopChars string, ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback)
	// ReadUntilFinish: finish an asynchronous call started by
	// g_data_input_stream_read_until_async().
	ReadUntilFinish(result AsyncResulter) (uint, string, error)
	// ReadUpto reads a string from the data input stream, up to the first
	// occurrence of any of the stop characters.
	ReadUpto(stopChars string, stopCharsLen int, cancellable Cancellabler) (uint, string, error)
	// ReadUptoAsync asynchronous version of g_data_input_stream_read_upto().
	ReadUptoAsync(stopChars string, stopCharsLen int, ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback)
	// ReadUptoFinish: finish an asynchronous call started by
	// g_data_input_stream_read_upto_async().
	ReadUptoFinish(result AsyncResulter) (uint, string, error)
}

// DataInputStream: data input stream implements Stream and includes functions
// for reading structured data directly from a binary input stream.
type DataInputStream struct {
	BufferedInputStream
}

var (
	_ DataInputStreamer = (*DataInputStream)(nil)
	_ gextras.Nativer   = (*DataInputStream)(nil)
)

func wrapDataInputStream(obj *externglib.Object) DataInputStreamer {
	return &DataInputStream{
		BufferedInputStream: BufferedInputStream{
			FilterInputStream: FilterInputStream{
				InputStream: InputStream{
					Object: obj,
				},
			},
			Seekable: Seekable{
				Object: obj,
			},
		},
	}
}

func marshalDataInputStreamer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDataInputStream(obj), nil
}

// NewDataInputStream creates a new data input stream for the @base_stream.
func NewDataInputStream(baseStream InputStreamer) *DataInputStream {
	var _arg1 *C.GInputStream     // out
	var _cret *C.GDataInputStream // in

	_arg1 = (*C.GInputStream)(unsafe.Pointer((baseStream).(gextras.Nativer).Native()))

	_cret = C.g_data_input_stream_new(_arg1)

	var _dataInputStream *DataInputStream // out

	_dataInputStream = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*DataInputStream)

	return _dataInputStream
}

// ByteOrder gets the byte order for the data input stream.
func (stream *DataInputStream) ByteOrder() DataStreamByteOrder {
	var _arg0 *C.GDataInputStream    // out
	var _cret C.GDataStreamByteOrder // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_data_input_stream_get_byte_order(_arg0)

	var _dataStreamByteOrder DataStreamByteOrder // out

	_dataStreamByteOrder = DataStreamByteOrder(_cret)

	return _dataStreamByteOrder
}

// NewlineType gets the current newline type for the @stream.
func (stream *DataInputStream) NewlineType() DataStreamNewlineType {
	var _arg0 *C.GDataInputStream      // out
	var _cret C.GDataStreamNewlineType // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_data_input_stream_get_newline_type(_arg0)

	var _dataStreamNewlineType DataStreamNewlineType // out

	_dataStreamNewlineType = DataStreamNewlineType(_cret)

	return _dataStreamNewlineType
}

// ReadByte reads an unsigned 8-bit/1-byte value from @stream.
func (stream *DataInputStream) ReadByte(cancellable Cancellabler) (byte, error) {
	var _arg0 *C.GDataInputStream // out
	var _arg1 *C.GCancellable     // out
	var _cret C.guchar            // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_data_input_stream_read_byte(_arg0, _arg1, &_cerr)

	var _guint8 byte // out
	var _goerr error // out

	_guint8 = byte(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _guint8, _goerr
}

// ReadInt16 reads a 16-bit/2-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order() and
// g_data_input_stream_set_byte_order().
func (stream *DataInputStream) ReadInt16(cancellable Cancellabler) (int16, error) {
	var _arg0 *C.GDataInputStream // out
	var _arg1 *C.GCancellable     // out
	var _cret C.gint16            // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_data_input_stream_read_int16(_arg0, _arg1, &_cerr)

	var _gint16 int16 // out
	var _goerr error  // out

	_gint16 = int16(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gint16, _goerr
}

// ReadInt32 reads a signed 32-bit/4-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order() and
// g_data_input_stream_set_byte_order().
//
// If @cancellable is not nil, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned.
func (stream *DataInputStream) ReadInt32(cancellable Cancellabler) (int32, error) {
	var _arg0 *C.GDataInputStream // out
	var _arg1 *C.GCancellable     // out
	var _cret C.gint32            // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_data_input_stream_read_int32(_arg0, _arg1, &_cerr)

	var _gint32 int32 // out
	var _goerr error  // out

	_gint32 = int32(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gint32, _goerr
}

// ReadInt64 reads a 64-bit/8-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order() and
// g_data_input_stream_set_byte_order().
//
// If @cancellable is not nil, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned.
func (stream *DataInputStream) ReadInt64(cancellable Cancellabler) (int64, error) {
	var _arg0 *C.GDataInputStream // out
	var _arg1 *C.GCancellable     // out
	var _cret C.gint64            // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_data_input_stream_read_int64(_arg0, _arg1, &_cerr)

	var _gint64 int64 // out
	var _goerr error  // out

	_gint64 = int64(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gint64, _goerr
}

// ReadLine reads a line from the data input stream. Note that no encoding
// checks or conversion is performed; the input is not guaranteed to be UTF-8,
// and may in fact have embedded NUL characters.
//
// If @cancellable is not nil, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned.
func (stream *DataInputStream) ReadLine(cancellable Cancellabler) (uint, []byte, error) {
	var _arg0 *C.GDataInputStream // out
	var _arg1 C.gsize             // in
	var _arg2 *C.GCancellable     // out
	var _cret *C.char
	var _cerr *C.GError // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_data_input_stream_read_line(_arg0, &_arg1, _arg2, &_cerr)

	var _length uint // out
	var _guint8s []byte
	var _goerr error // out

	_length = uint(_arg1)
	{
		var i int
		var z C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_guint8s = make([]byte, i)
		for i := range src {
			_guint8s[i] = byte(src[i])
		}
	}
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _length, _guint8s, _goerr
}

// ReadLineAsync asynchronous version of g_data_input_stream_read_line(). It is
// an error to have two outstanding calls to this function.
//
// When the operation is finished, @callback will be called. You can then call
// g_data_input_stream_read_line_finish() to get the result of the operation.
func (stream *DataInputStream) ReadLineAsync(ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback) {
	var _arg0 *C.GDataInputStream   // out
	var _arg1 C.gint                // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = C.gint(ioPriority)
	_arg2 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(gbox.Assign(callback))

	C.g_data_input_stream_read_line_async(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// ReadLineFinish: finish an asynchronous call started by
// g_data_input_stream_read_line_async(). Note the warning about string encoding
// in g_data_input_stream_read_line() applies here as well.
func (stream *DataInputStream) ReadLineFinish(result AsyncResulter) (uint, []byte, error) {
	var _arg0 *C.GDataInputStream // out
	var _arg1 *C.GAsyncResult     // out
	var _arg2 C.gsize             // in
	var _cret *C.char
	var _cerr *C.GError // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.g_data_input_stream_read_line_finish(_arg0, _arg1, &_arg2, &_cerr)

	var _length uint // out
	var _guint8s []byte
	var _goerr error // out

	_length = uint(_arg2)
	{
		var i int
		var z C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_guint8s = make([]byte, i)
		for i := range src {
			_guint8s[i] = byte(src[i])
		}
	}
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _length, _guint8s, _goerr
}

// ReadLineFinishUTF8: finish an asynchronous call started by
// g_data_input_stream_read_line_async().
func (stream *DataInputStream) ReadLineFinishUTF8(result AsyncResulter) (uint, string, error) {
	var _arg0 *C.GDataInputStream // out
	var _arg1 *C.GAsyncResult     // out
	var _arg2 C.gsize             // in
	var _cret *C.char             // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.g_data_input_stream_read_line_finish_utf8(_arg0, _arg1, &_arg2, &_cerr)

	var _length uint // out
	var _utf8 string // out
	var _goerr error // out

	_length = uint(_arg2)
	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _length, _utf8, _goerr
}

// ReadLineUTF8 reads a UTF-8 encoded line from the data input stream.
//
// If @cancellable is not nil, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned.
func (stream *DataInputStream) ReadLineUTF8(cancellable Cancellabler) (uint, string, error) {
	var _arg0 *C.GDataInputStream // out
	var _arg1 C.gsize             // in
	var _arg2 *C.GCancellable     // out
	var _cret *C.char             // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_data_input_stream_read_line_utf8(_arg0, &_arg1, _arg2, &_cerr)

	var _length uint // out
	var _utf8 string // out
	var _goerr error // out

	_length = uint(_arg1)
	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _length, _utf8, _goerr
}

// ReadUint16 reads an unsigned 16-bit/2-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order() and
// g_data_input_stream_set_byte_order().
func (stream *DataInputStream) ReadUint16(cancellable Cancellabler) (uint16, error) {
	var _arg0 *C.GDataInputStream // out
	var _arg1 *C.GCancellable     // out
	var _cret C.guint16           // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_data_input_stream_read_uint16(_arg0, _arg1, &_cerr)

	var _guint16 uint16 // out
	var _goerr error    // out

	_guint16 = uint16(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _guint16, _goerr
}

// ReadUint32 reads an unsigned 32-bit/4-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order() and
// g_data_input_stream_set_byte_order().
//
// If @cancellable is not nil, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned.
func (stream *DataInputStream) ReadUint32(cancellable Cancellabler) (uint32, error) {
	var _arg0 *C.GDataInputStream // out
	var _arg1 *C.GCancellable     // out
	var _cret C.guint32           // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_data_input_stream_read_uint32(_arg0, _arg1, &_cerr)

	var _guint32 uint32 // out
	var _goerr error    // out

	_guint32 = uint32(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _guint32, _goerr
}

// ReadUint64 reads an unsigned 64-bit/8-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order().
//
// If @cancellable is not nil, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned.
func (stream *DataInputStream) ReadUint64(cancellable Cancellabler) (uint64, error) {
	var _arg0 *C.GDataInputStream // out
	var _arg1 *C.GCancellable     // out
	var _cret C.guint64           // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_data_input_stream_read_uint64(_arg0, _arg1, &_cerr)

	var _guint64 uint64 // out
	var _goerr error    // out

	_guint64 = uint64(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _guint64, _goerr
}

// ReadUntil reads a string from the data input stream, up to the first
// occurrence of any of the stop characters.
//
// Note that, in contrast to g_data_input_stream_read_until_async(), this
// function consumes the stop character that it finds.
//
// Don't use this function in new code. Its functionality is inconsistent with
// g_data_input_stream_read_until_async(). Both functions will be marked as
// deprecated in a future release. Use g_data_input_stream_read_upto() instead,
// but note that that function does not consume the stop character.
//
// Deprecated: Use g_data_input_stream_read_upto() instead, which has more
// consistent behaviour regarding the stop character.
func (stream *DataInputStream) ReadUntil(stopChars string, cancellable Cancellabler) (uint, string, error) {
	var _arg0 *C.GDataInputStream // out
	var _arg1 *C.gchar            // out
	var _arg2 C.gsize             // in
	var _arg3 *C.GCancellable     // out
	var _cret *C.char             // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stopChars)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg3 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_data_input_stream_read_until(_arg0, _arg1, &_arg2, _arg3, &_cerr)

	var _length uint // out
	var _utf8 string // out
	var _goerr error // out

	_length = uint(_arg2)
	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _length, _utf8, _goerr
}

// ReadUntilAsync asynchronous version of g_data_input_stream_read_until(). It
// is an error to have two outstanding calls to this function.
//
// Note that, in contrast to g_data_input_stream_read_until(), this function
// does not consume the stop character that it finds. You must read it for
// yourself.
//
// When the operation is finished, @callback will be called. You can then call
// g_data_input_stream_read_until_finish() to get the result of the operation.
//
// Don't use this function in new code. Its functionality is inconsistent with
// g_data_input_stream_read_until(). Both functions will be marked as deprecated
// in a future release. Use g_data_input_stream_read_upto_async() instead.
//
// Deprecated: Use g_data_input_stream_read_upto_async() instead, which has more
// consistent behaviour regarding the stop character.
func (stream *DataInputStream) ReadUntilAsync(stopChars string, ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback) {
	var _arg0 *C.GDataInputStream   // out
	var _arg1 *C.gchar              // out
	var _arg2 C.gint                // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stopChars)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(ioPriority)
	_arg3 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(gbox.Assign(callback))

	C.g_data_input_stream_read_until_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// ReadUntilFinish: finish an asynchronous call started by
// g_data_input_stream_read_until_async().
//
// Deprecated: Use g_data_input_stream_read_upto_finish() instead, which has
// more consistent behaviour regarding the stop character.
func (stream *DataInputStream) ReadUntilFinish(result AsyncResulter) (uint, string, error) {
	var _arg0 *C.GDataInputStream // out
	var _arg1 *C.GAsyncResult     // out
	var _arg2 C.gsize             // in
	var _cret *C.char             // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.g_data_input_stream_read_until_finish(_arg0, _arg1, &_arg2, &_cerr)

	var _length uint // out
	var _utf8 string // out
	var _goerr error // out

	_length = uint(_arg2)
	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _length, _utf8, _goerr
}

// ReadUpto reads a string from the data input stream, up to the first
// occurrence of any of the stop characters.
//
// In contrast to g_data_input_stream_read_until(), this function does not
// consume the stop character. You have to use g_data_input_stream_read_byte()
// to get it before calling g_data_input_stream_read_upto() again.
//
// Note that @stop_chars may contain '\0' if @stop_chars_len is specified.
//
// The returned string will always be nul-terminated on success.
func (stream *DataInputStream) ReadUpto(stopChars string, stopCharsLen int, cancellable Cancellabler) (uint, string, error) {
	var _arg0 *C.GDataInputStream // out
	var _arg1 *C.gchar            // out
	var _arg2 C.gssize            // out
	var _arg3 C.gsize             // in
	var _arg4 *C.GCancellable     // out
	var _cret *C.char             // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stopChars)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(stopCharsLen)
	_arg4 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_data_input_stream_read_upto(_arg0, _arg1, _arg2, &_arg3, _arg4, &_cerr)

	var _length uint // out
	var _utf8 string // out
	var _goerr error // out

	_length = uint(_arg3)
	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _length, _utf8, _goerr
}

// ReadUptoAsync asynchronous version of g_data_input_stream_read_upto(). It is
// an error to have two outstanding calls to this function.
//
// In contrast to g_data_input_stream_read_until(), this function does not
// consume the stop character. You have to use g_data_input_stream_read_byte()
// to get it before calling g_data_input_stream_read_upto() again.
//
// Note that @stop_chars may contain '\0' if @stop_chars_len is specified.
//
// When the operation is finished, @callback will be called. You can then call
// g_data_input_stream_read_upto_finish() to get the result of the operation.
func (stream *DataInputStream) ReadUptoAsync(stopChars string, stopCharsLen int, ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback) {
	var _arg0 *C.GDataInputStream   // out
	var _arg1 *C.gchar              // out
	var _arg2 C.gssize              // out
	var _arg3 C.gint                // out
	var _arg4 *C.GCancellable       // out
	var _arg5 C.GAsyncReadyCallback // out
	var _arg6 C.gpointer

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stopChars)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(stopCharsLen)
	_arg3 = C.gint(ioPriority)
	_arg4 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))
	_arg5 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg6 = C.gpointer(gbox.Assign(callback))

	C.g_data_input_stream_read_upto_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// ReadUptoFinish: finish an asynchronous call started by
// g_data_input_stream_read_upto_async().
//
// Note that this function does not consume the stop character. You have to use
// g_data_input_stream_read_byte() to get it before calling
// g_data_input_stream_read_upto_async() again.
//
// The returned string will always be nul-terminated on success.
func (stream *DataInputStream) ReadUptoFinish(result AsyncResulter) (uint, string, error) {
	var _arg0 *C.GDataInputStream // out
	var _arg1 *C.GAsyncResult     // out
	var _arg2 C.gsize             // in
	var _cret *C.char             // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GDataInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.g_data_input_stream_read_upto_finish(_arg0, _arg1, &_arg2, &_cerr)

	var _length uint // out
	var _utf8 string // out
	var _goerr error // out

	_length = uint(_arg2)
	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _length, _utf8, _goerr
}
