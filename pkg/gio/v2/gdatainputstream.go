// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
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
		{T: externglib.Type(C.g_data_input_stream_get_type()), F: marshalDataInputStream},
	})
}

// DataInputStream: data input stream implements Stream and includes functions
// for reading structured data directly from a binary input stream.
type DataInputStream interface {
	BufferedInputStream
	Seekable

	// ByteOrder gets the byte order for the data input stream.
	ByteOrder() DataStreamByteOrder
	// NewlineType gets the current newline type for the @stream.
	NewlineType() DataStreamNewlineType
	// ReadByte reads an unsigned 8-bit/1-byte value from @stream.
	ReadByte(cancellable Cancellable) (guint8 byte, err error)
	// ReadInt16 reads a 16-bit/2-byte value from @stream.
	//
	// In order to get the correct byte order for this read operation, see
	// g_data_input_stream_get_byte_order() and
	// g_data_input_stream_set_byte_order().
	ReadInt16(cancellable Cancellable) (gint16 int16, err error)
	// ReadInt32 reads a signed 32-bit/4-byte value from @stream.
	//
	// In order to get the correct byte order for this read operation, see
	// g_data_input_stream_get_byte_order() and
	// g_data_input_stream_set_byte_order().
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	ReadInt32(cancellable Cancellable) (gint32 int32, err error)
	// ReadInt64 reads a 64-bit/8-byte value from @stream.
	//
	// In order to get the correct byte order for this read operation, see
	// g_data_input_stream_get_byte_order() and
	// g_data_input_stream_set_byte_order().
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	ReadInt64(cancellable Cancellable) (gint64 int64, err error)
	// ReadLine reads a line from the data input stream. Note that no encoding
	// checks or conversion is performed; the input is not guaranteed to be
	// UTF-8, and may in fact have embedded NUL characters.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	ReadLine(cancellable Cancellable) (length uint, guint8s []byte, err error)
	// ReadLineAsync: the asynchronous version of
	// g_data_input_stream_read_line(). It is an error to have two outstanding
	// calls to this function.
	//
	// When the operation is finished, @callback will be called. You can then
	// call g_data_input_stream_read_line_finish() to get the result of the
	// operation.
	ReadLineAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// ReadLineFinish: finish an asynchronous call started by
	// g_data_input_stream_read_line_async(). Note the warning about string
	// encoding in g_data_input_stream_read_line() applies here as well.
	ReadLineFinish(result AsyncResult) (length uint, guint8s []byte, err error)
	// ReadLineFinishUTF8: finish an asynchronous call started by
	// g_data_input_stream_read_line_async().
	ReadLineFinishUTF8(result AsyncResult) (length uint, utf8 string, err error)
	// ReadLineUTF8 reads a UTF-8 encoded line from the data input stream.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	ReadLineUTF8(cancellable Cancellable) (length uint, utf8 string, err error)
	// ReadUint16 reads an unsigned 16-bit/2-byte value from @stream.
	//
	// In order to get the correct byte order for this read operation, see
	// g_data_input_stream_get_byte_order() and
	// g_data_input_stream_set_byte_order().
	ReadUint16(cancellable Cancellable) (guint16 uint16, err error)
	// ReadUint32 reads an unsigned 32-bit/4-byte value from @stream.
	//
	// In order to get the correct byte order for this read operation, see
	// g_data_input_stream_get_byte_order() and
	// g_data_input_stream_set_byte_order().
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	ReadUint32(cancellable Cancellable) (guint32 uint32, err error)
	// ReadUint64 reads an unsigned 64-bit/8-byte value from @stream.
	//
	// In order to get the correct byte order for this read operation, see
	// g_data_input_stream_get_byte_order().
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	ReadUint64(cancellable Cancellable) (guint64 uint64, err error)
	// ReadUntil reads a string from the data input stream, up to the first
	// occurrence of any of the stop characters.
	//
	// Note that, in contrast to g_data_input_stream_read_until_async(), this
	// function consumes the stop character that it finds.
	//
	// Don't use this function in new code. Its functionality is inconsistent
	// with g_data_input_stream_read_until_async(). Both functions will be
	// marked as deprecated in a future release. Use
	// g_data_input_stream_read_upto() instead, but note that that function does
	// not consume the stop character.
	ReadUntil(stopChars string, cancellable Cancellable) (length uint, utf8 string, err error)
	// ReadUntilAsync: the asynchronous version of
	// g_data_input_stream_read_until(). It is an error to have two outstanding
	// calls to this function.
	//
	// Note that, in contrast to g_data_input_stream_read_until(), this function
	// does not consume the stop character that it finds. You must read it for
	// yourself.
	//
	// When the operation is finished, @callback will be called. You can then
	// call g_data_input_stream_read_until_finish() to get the result of the
	// operation.
	//
	// Don't use this function in new code. Its functionality is inconsistent
	// with g_data_input_stream_read_until(). Both functions will be marked as
	// deprecated in a future release. Use g_data_input_stream_read_upto_async()
	// instead.
	ReadUntilAsync(stopChars string, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// ReadUntilFinish: finish an asynchronous call started by
	// g_data_input_stream_read_until_async().
	ReadUntilFinish(result AsyncResult) (length uint, utf8 string, err error)
	// ReadUpto reads a string from the data input stream, up to the first
	// occurrence of any of the stop characters.
	//
	// In contrast to g_data_input_stream_read_until(), this function does not
	// consume the stop character. You have to use
	// g_data_input_stream_read_byte() to get it before calling
	// g_data_input_stream_read_upto() again.
	//
	// Note that @stop_chars may contain '\0' if @stop_chars_len is specified.
	//
	// The returned string will always be nul-terminated on success.
	ReadUpto(stopChars string, stopCharsLen int, cancellable Cancellable) (length uint, utf8 string, err error)
	// ReadUptoAsync: the asynchronous version of
	// g_data_input_stream_read_upto(). It is an error to have two outstanding
	// calls to this function.
	//
	// In contrast to g_data_input_stream_read_until(), this function does not
	// consume the stop character. You have to use
	// g_data_input_stream_read_byte() to get it before calling
	// g_data_input_stream_read_upto() again.
	//
	// Note that @stop_chars may contain '\0' if @stop_chars_len is specified.
	//
	// When the operation is finished, @callback will be called. You can then
	// call g_data_input_stream_read_upto_finish() to get the result of the
	// operation.
	ReadUptoAsync(stopChars string, stopCharsLen int, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// ReadUptoFinish: finish an asynchronous call started by
	// g_data_input_stream_read_upto_async().
	//
	// Note that this function does not consume the stop character. You have to
	// use g_data_input_stream_read_byte() to get it before calling
	// g_data_input_stream_read_upto_async() again.
	//
	// The returned string will always be nul-terminated on success.
	ReadUptoFinish(result AsyncResult) (length uint, utf8 string, err error)
	// SetByteOrder: this function sets the byte order for the given @stream.
	// All subsequent reads from the @stream will be read in the given @order.
	SetByteOrder(order DataStreamByteOrder)
	// SetNewlineType sets the newline type for the @stream.
	//
	// Note that using G_DATA_STREAM_NEWLINE_TYPE_ANY is slightly unsafe. If a
	// read chunk ends in "CR" we must read an additional byte to know if this
	// is "CR" or "CR LF", and this might block if there is no more data
	// available.
	SetNewlineType(typ DataStreamNewlineType)
}

// dataInputStream implements the DataInputStream interface.
type dataInputStream struct {
	BufferedInputStream
	Seekable
}

var _ DataInputStream = (*dataInputStream)(nil)

// WrapDataInputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapDataInputStream(obj *externglib.Object) DataInputStream {
	return DataInputStream{
		BufferedInputStream: WrapBufferedInputStream(obj),
		Seekable:            WrapSeekable(obj),
	}
}

func marshalDataInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDataInputStream(obj), nil
}

// NewDataInputStream constructs a class DataInputStream.
func NewDataInputStream(baseStream InputStream) DataInputStream {
	var arg1 *C.GInputStream

	arg1 = (*C.GInputStream)(baseStream.Native())

	ret := C.g_data_input_stream_new(arg1)

	var ret0 DataInputStream

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(DataInputStream)

	return ret0
}

// ByteOrder gets the byte order for the data input stream.
func (s dataInputStream) ByteOrder() DataStreamByteOrder {
	var arg0 *C.GDataInputStream

	arg0 = (*C.GDataInputStream)(s.Native())

	ret := C.g_data_input_stream_get_byte_order(arg0)

	var ret0 DataStreamByteOrder

	ret0 = DataStreamByteOrder(ret)

	return ret0
}

// NewlineType gets the current newline type for the @stream.
func (s dataInputStream) NewlineType() DataStreamNewlineType {
	var arg0 *C.GDataInputStream

	arg0 = (*C.GDataInputStream)(s.Native())

	ret := C.g_data_input_stream_get_newline_type(arg0)

	var ret0 DataStreamNewlineType

	ret0 = DataStreamNewlineType(ret)

	return ret0
}

// ReadByte reads an unsigned 8-bit/1-byte value from @stream.
func (s dataInputStream) ReadByte(cancellable Cancellable) (guint8 byte, err error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_data_input_stream_read_byte(arg0, arg1, &gError)

	var ret0 byte
	var goError error

	ret0 = byte(ret)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// ReadInt16 reads a 16-bit/2-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order() and
// g_data_input_stream_set_byte_order().
func (s dataInputStream) ReadInt16(cancellable Cancellable) (gint16 int16, err error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_data_input_stream_read_int16(arg0, arg1, &gError)

	var ret0 int16
	var goError error

	ret0 = int16(ret)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// ReadInt32 reads a signed 32-bit/4-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order() and
// g_data_input_stream_set_byte_order().
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
func (s dataInputStream) ReadInt32(cancellable Cancellable) (gint32 int32, err error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_data_input_stream_read_int32(arg0, arg1, &gError)

	var ret0 int32
	var goError error

	ret0 = int32(ret)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// ReadInt64 reads a 64-bit/8-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order() and
// g_data_input_stream_set_byte_order().
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
func (s dataInputStream) ReadInt64(cancellable Cancellable) (gint64 int64, err error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_data_input_stream_read_int64(arg0, arg1, &gError)

	var ret0 int64
	var goError error

	ret0 = int64(ret)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// ReadLine reads a line from the data input stream. Note that no encoding
// checks or conversion is performed; the input is not guaranteed to be
// UTF-8, and may in fact have embedded NUL characters.
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
func (s dataInputStream) ReadLine(cancellable Cancellable) (length uint, guint8s []byte, err error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.gsize // out
	var arg2 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GDataInputStream)(s.Native())
	arg2 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_data_input_stream_read_line(arg0, &arg1, arg2, &gError)

	var ret0 uint
	var ret1 []byte
	var goError error

	ret0 = uint(arg1)

	{
		var length uint
		for p := unsafe.Pointer(ret); *p != 0; p = unsafe.Pointer(uintptr(p) + 1) {
			length++
		}

		ret1 = make([]byte, length)
		for i := 0; i < length; i++ {
			src := (C.guint8)(unsafe.Pointer(uintptr(unsafe.Pointer(ret)) + i))
			ret1[i] = byte(src)
		}
	}

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, ret1, goError
}

// ReadLineAsync: the asynchronous version of
// g_data_input_stream_read_line(). It is an error to have two outstanding
// calls to this function.
//
// When the operation is finished, @callback will be called. You can then
// call g_data_input_stream_read_line_finish() to get the result of the
// operation.
func (s dataInputStream) ReadLineAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GDataInputStream
	var arg1 C.gint
	var arg2 *C.GCancellable
	var arg3 C.GAsyncReadyCallback
	var arg4 C.gpointer

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = C.gint(ioPriority)
	arg2 = (*C.GCancellable)(cancellable.Native())
	arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	arg4 = C.gpointer(box.Assign(callback))

	C.g_data_input_stream_read_line_async(arg0, arg1, arg2, arg3, arg4)
}

// ReadLineFinish: finish an asynchronous call started by
// g_data_input_stream_read_line_async(). Note the warning about string
// encoding in g_data_input_stream_read_line() applies here as well.
func (s dataInputStream) ReadLineFinish(result AsyncResult) (length uint, guint8s []byte, err error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GAsyncResult
	var arg2 *C.gsize // out
	var gError *C.GError

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (*C.GAsyncResult)(result.Native())

	ret := C.g_data_input_stream_read_line_finish(arg0, arg1, &arg2, &gError)

	var ret0 uint
	var ret1 []byte
	var goError error

	ret0 = uint(arg2)

	{
		var length uint
		for p := unsafe.Pointer(ret); *p != 0; p = unsafe.Pointer(uintptr(p) + 1) {
			length++
		}

		ret1 = make([]byte, length)
		for i := 0; i < length; i++ {
			src := (C.guint8)(unsafe.Pointer(uintptr(unsafe.Pointer(ret)) + i))
			ret1[i] = byte(src)
		}
	}

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, ret1, goError
}

// ReadLineFinishUTF8: finish an asynchronous call started by
// g_data_input_stream_read_line_async().
func (s dataInputStream) ReadLineFinishUTF8(result AsyncResult) (length uint, utf8 string, err error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GAsyncResult
	var arg2 *C.gsize // out
	var gError *C.GError

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (*C.GAsyncResult)(result.Native())

	ret := C.g_data_input_stream_read_line_finish_utf8(arg0, arg1, &arg2, &gError)

	var ret0 uint
	var ret1 string
	var goError error

	ret0 = uint(arg2)

	ret1 = C.GoString(ret)
	C.free(unsafe.Pointer(ret))

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, ret1, goError
}

// ReadLineUTF8 reads a UTF-8 encoded line from the data input stream.
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
func (s dataInputStream) ReadLineUTF8(cancellable Cancellable) (length uint, utf8 string, err error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.gsize // out
	var arg2 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GDataInputStream)(s.Native())
	arg2 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_data_input_stream_read_line_utf8(arg0, &arg1, arg2, &gError)

	var ret0 uint
	var ret1 string
	var goError error

	ret0 = uint(arg1)

	ret1 = C.GoString(ret)
	C.free(unsafe.Pointer(ret))

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, ret1, goError
}

// ReadUint16 reads an unsigned 16-bit/2-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order() and
// g_data_input_stream_set_byte_order().
func (s dataInputStream) ReadUint16(cancellable Cancellable) (guint16 uint16, err error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_data_input_stream_read_uint16(arg0, arg1, &gError)

	var ret0 uint16
	var goError error

	ret0 = uint16(ret)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// ReadUint32 reads an unsigned 32-bit/4-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order() and
// g_data_input_stream_set_byte_order().
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
func (s dataInputStream) ReadUint32(cancellable Cancellable) (guint32 uint32, err error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_data_input_stream_read_uint32(arg0, arg1, &gError)

	var ret0 uint32
	var goError error

	ret0 = uint32(ret)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// ReadUint64 reads an unsigned 64-bit/8-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order().
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
func (s dataInputStream) ReadUint64(cancellable Cancellable) (guint64 uint64, err error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_data_input_stream_read_uint64(arg0, arg1, &gError)

	var ret0 uint64
	var goError error

	ret0 = uint64(ret)

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, goError
}

// ReadUntil reads a string from the data input stream, up to the first
// occurrence of any of the stop characters.
//
// Note that, in contrast to g_data_input_stream_read_until_async(), this
// function consumes the stop character that it finds.
//
// Don't use this function in new code. Its functionality is inconsistent
// with g_data_input_stream_read_until_async(). Both functions will be
// marked as deprecated in a future release. Use
// g_data_input_stream_read_upto() instead, but note that that function does
// not consume the stop character.
func (s dataInputStream) ReadUntil(stopChars string, cancellable Cancellable) (length uint, utf8 string, err error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.gchar
	var arg2 *C.gsize // out
	var arg3 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (*C.gchar)(C.CString(stopChars))
	defer C.free(unsafe.Pointer(arg1))
	arg3 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_data_input_stream_read_until(arg0, arg1, &arg2, arg3, &gError)

	var ret0 uint
	var ret1 string
	var goError error

	ret0 = uint(arg2)

	ret1 = C.GoString(ret)
	C.free(unsafe.Pointer(ret))

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, ret1, goError
}

// ReadUntilAsync: the asynchronous version of
// g_data_input_stream_read_until(). It is an error to have two outstanding
// calls to this function.
//
// Note that, in contrast to g_data_input_stream_read_until(), this function
// does not consume the stop character that it finds. You must read it for
// yourself.
//
// When the operation is finished, @callback will be called. You can then
// call g_data_input_stream_read_until_finish() to get the result of the
// operation.
//
// Don't use this function in new code. Its functionality is inconsistent
// with g_data_input_stream_read_until(). Both functions will be marked as
// deprecated in a future release. Use g_data_input_stream_read_upto_async()
// instead.
func (s dataInputStream) ReadUntilAsync(stopChars string, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GDataInputStream
	var arg1 *C.gchar
	var arg2 C.gint
	var arg3 *C.GCancellable
	var arg4 C.GAsyncReadyCallback
	var arg5 C.gpointer

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (*C.gchar)(C.CString(stopChars))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gint(ioPriority)
	arg3 = (*C.GCancellable)(cancellable.Native())
	arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	arg5 = C.gpointer(box.Assign(callback))

	C.g_data_input_stream_read_until_async(arg0, arg1, arg2, arg3, arg4, arg5)
}

// ReadUntilFinish: finish an asynchronous call started by
// g_data_input_stream_read_until_async().
func (s dataInputStream) ReadUntilFinish(result AsyncResult) (length uint, utf8 string, err error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GAsyncResult
	var arg2 *C.gsize // out
	var gError *C.GError

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (*C.GAsyncResult)(result.Native())

	ret := C.g_data_input_stream_read_until_finish(arg0, arg1, &arg2, &gError)

	var ret0 uint
	var ret1 string
	var goError error

	ret0 = uint(arg2)

	ret1 = C.GoString(ret)
	C.free(unsafe.Pointer(ret))

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, ret1, goError
}

// ReadUpto reads a string from the data input stream, up to the first
// occurrence of any of the stop characters.
//
// In contrast to g_data_input_stream_read_until(), this function does not
// consume the stop character. You have to use
// g_data_input_stream_read_byte() to get it before calling
// g_data_input_stream_read_upto() again.
//
// Note that @stop_chars may contain '\0' if @stop_chars_len is specified.
//
// The returned string will always be nul-terminated on success.
func (s dataInputStream) ReadUpto(stopChars string, stopCharsLen int, cancellable Cancellable) (length uint, utf8 string, err error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.gchar
	var arg2 C.gssize
	var arg3 *C.gsize // out
	var arg4 *C.GCancellable
	var gError *C.GError

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (*C.gchar)(C.CString(stopChars))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gssize(stopCharsLen)
	arg4 = (*C.GCancellable)(cancellable.Native())

	ret := C.g_data_input_stream_read_upto(arg0, arg1, arg2, &arg3, arg4, &gError)

	var ret0 uint
	var ret1 string
	var goError error

	ret0 = uint(arg3)

	ret1 = C.GoString(ret)
	C.free(unsafe.Pointer(ret))

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, ret1, goError
}

// ReadUptoAsync: the asynchronous version of
// g_data_input_stream_read_upto(). It is an error to have two outstanding
// calls to this function.
//
// In contrast to g_data_input_stream_read_until(), this function does not
// consume the stop character. You have to use
// g_data_input_stream_read_byte() to get it before calling
// g_data_input_stream_read_upto() again.
//
// Note that @stop_chars may contain '\0' if @stop_chars_len is specified.
//
// When the operation is finished, @callback will be called. You can then
// call g_data_input_stream_read_upto_finish() to get the result of the
// operation.
func (s dataInputStream) ReadUptoAsync(stopChars string, stopCharsLen int, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GDataInputStream
	var arg1 *C.gchar
	var arg2 C.gssize
	var arg3 C.gint
	var arg4 *C.GCancellable
	var arg5 C.GAsyncReadyCallback
	var arg6 C.gpointer

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (*C.gchar)(C.CString(stopChars))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gssize(stopCharsLen)
	arg3 = C.gint(ioPriority)
	arg4 = (*C.GCancellable)(cancellable.Native())
	arg5 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	arg6 = C.gpointer(box.Assign(callback))

	C.g_data_input_stream_read_upto_async(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// ReadUptoFinish: finish an asynchronous call started by
// g_data_input_stream_read_upto_async().
//
// Note that this function does not consume the stop character. You have to
// use g_data_input_stream_read_byte() to get it before calling
// g_data_input_stream_read_upto_async() again.
//
// The returned string will always be nul-terminated on success.
func (s dataInputStream) ReadUptoFinish(result AsyncResult) (length uint, utf8 string, err error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GAsyncResult
	var arg2 *C.gsize // out
	var gError *C.GError

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (*C.GAsyncResult)(result.Native())

	ret := C.g_data_input_stream_read_upto_finish(arg0, arg1, &arg2, &gError)

	var ret0 uint
	var ret1 string
	var goError error

	ret0 = uint(arg2)

	ret1 = C.GoString(ret)
	C.free(unsafe.Pointer(ret))

	if gError != nil {
		goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))
		C.g_error_free(gError)
	}

	return ret0, ret1, goError
}

// SetByteOrder: this function sets the byte order for the given @stream.
// All subsequent reads from the @stream will be read in the given @order.
func (s dataInputStream) SetByteOrder(order DataStreamByteOrder) {
	var arg0 *C.GDataInputStream
	var arg1 C.GDataStreamByteOrder

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (C.GDataStreamByteOrder)(order)

	C.g_data_input_stream_set_byte_order(arg0, arg1)
}

// SetNewlineType sets the newline type for the @stream.
//
// Note that using G_DATA_STREAM_NEWLINE_TYPE_ANY is slightly unsafe. If a
// read chunk ends in "CR" we must read an additional byte to know if this
// is "CR" or "CR LF", and this might block if there is no more data
// available.
func (s dataInputStream) SetNewlineType(typ DataStreamNewlineType) {
	var arg0 *C.GDataInputStream
	var arg1 C.GDataStreamNewlineType

	arg0 = (*C.GDataInputStream)(s.Native())
	arg1 = (C.GDataStreamNewlineType)(typ)

	C.g_data_input_stream_set_newline_type(arg0, arg1)
}

type DataInputStreamPrivate struct {
	native C.GDataInputStreamPrivate
}

// WrapDataInputStreamPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDataInputStreamPrivate(ptr unsafe.Pointer) *DataInputStreamPrivate {
	if ptr == nil {
		return nil
	}

	return (*DataInputStreamPrivate)(ptr)
}

func marshalDataInputStreamPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDataInputStreamPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DataInputStreamPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}
