// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// NewDBusMessage creates a new empty BusMessage.
//
// The function returns the following values:
//
//    - dBusMessage Free with g_object_unref().
//
func NewDBusMessage() *DBusMessage {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("new_DBusMessage", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _dBusMessage *DBusMessage // out

	_dBusMessage = wrapDBusMessage(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusMessage
}

// NewDBusMessageMethodCall creates a new BusMessage for a method call.
//
// The function takes the following parameters:
//
//    - name (optional): valid D-Bus name or NULL.
//    - path: valid object path.
//    - interface_ (optional): valid D-Bus interface name or NULL.
//    - method: valid method name.
//
// The function returns the following values:
//
//    - dBusMessage Free with g_object_unref().
//
func NewDBusMessageMethodCall(name, path, interface_, method string) *DBusMessage {
	var _args [4]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out
	var _arg3 *C.void // out
	var _cret *C.void // in

	if name != "" {
		_arg0 = (*C.void)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg0))
	}
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))
	if interface_ != "" {
		_arg2 = (*C.void)(unsafe.Pointer(C.CString(interface_)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	_arg3 = (*C.void)(unsafe.Pointer(C.CString(method)))
	defer C.free(unsafe.Pointer(_arg3))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(**C.void)(unsafe.Pointer(&_args[2])) = _arg2
	*(**C.void)(unsafe.Pointer(&_args[3])) = _arg3

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("new_DBusMessage_method_call", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(name)
	runtime.KeepAlive(path)
	runtime.KeepAlive(interface_)
	runtime.KeepAlive(method)

	var _dBusMessage *DBusMessage // out

	_dBusMessage = wrapDBusMessage(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusMessage
}

// NewDBusMessageSignal creates a new BusMessage for a signal emission.
//
// The function takes the following parameters:
//
//    - path: valid object path.
//    - interface_: valid D-Bus interface name.
//    - signal: valid signal name.
//
// The function returns the following values:
//
//    - dBusMessage Free with g_object_unref().
//
func NewDBusMessageSignal(path, interface_, signal string) *DBusMessage {
	var _args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg0))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(interface_)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.void)(unsafe.Pointer(C.CString(signal)))
	defer C.free(unsafe.Pointer(_arg2))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(**C.void)(unsafe.Pointer(&_args[2])) = _arg2

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("new_DBusMessage_signal", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(path)
	runtime.KeepAlive(interface_)
	runtime.KeepAlive(signal)

	var _dBusMessage *DBusMessage // out

	_dBusMessage = wrapDBusMessage(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusMessage
}

// Copy copies message. The copy is a deep copy and the returned BusMessage is
// completely identical except that it is guaranteed to not be locked.
//
// This operation can fail if e.g. message contains file descriptors and the
// per-process or system-wide open files limit is reached.
//
// The function returns the following values:
//
//    - dBusMessage: new BusMessage or NULL if error is set. Free with
//      g_object_unref().
//
func (message *DBusMessage) Copy() (*DBusMessage, error) {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("copy", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(message)

	var _dBusMessage *DBusMessage // out
	var _goerr error              // out

	_dBusMessage = wrapDBusMessage(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dBusMessage, _goerr
}

// Arg0: convenience to get the first item in the body of message.
//
// The function returns the following values:
//
//    - utf8 (optional): string item or NULL if the first item in the body of
//      message is not a string.
//
func (message *DBusMessage) Arg0() string {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("get_arg0", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(message)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Body gets the body of a message.
//
// The function returns the following values:
//
//    - variant (optional) or NULL if the body is empty. Do not free, it is owned
//      by message.
//
func (message *DBusMessage) Body() *glib.Variant {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("get_body", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(message)

	var _variant *glib.Variant // out

	if _cret != nil {
		_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.g_variant_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_variant)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}

	return _variant
}

// Destination: convenience getter for the
// G_DBUS_MESSAGE_HEADER_FIELD_DESTINATION header field.
//
// The function returns the following values:
//
//    - utf8 (optional): value.
//
func (message *DBusMessage) Destination() string {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("get_destination", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(message)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ErrorName: convenience getter for the G_DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME
// header field.
//
// The function returns the following values:
//
//    - utf8 (optional): value.
//
func (message *DBusMessage) ErrorName() string {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("get_error_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(message)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// HeaderFields gets an array of all header fields on message that are set.
//
// The function returns the following values:
//
//    - guint8s: array of header fields terminated by
//      G_DBUS_MESSAGE_HEADER_FIELD_INVALID. Each element is a #guchar. Free with
//      g_free().
//
func (message *DBusMessage) HeaderFields() []byte {
	var _args [1]girepository.Argument
	var _arg0 *C.void   // out
	var _cret *C.guchar // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("get_header_fields", _args[:], nil)
	_cret = *(**C.guchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(message)

	var _guint8s []byte // out

	{
		var i int
		var z C.guchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_guint8s = make([]byte, i)
		for i := range src {
			_guint8s[i] = byte(src[i])
		}
	}

	return _guint8s
}

// Interface: convenience getter for the G_DBUS_MESSAGE_HEADER_FIELD_INTERFACE
// header field.
//
// The function returns the following values:
//
//    - utf8 (optional): value.
//
func (message *DBusMessage) Interface() string {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("get_interface", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(message)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Locked checks whether message is locked. To monitor changes to this value,
// conncet to the #GObject::notify signal to listen for changes on the
// BusMessage:locked property.
//
// The function returns the following values:
//
//    - ok: TRUE if message is locked, FALSE otherwise.
//
func (message *DBusMessage) Locked() bool {
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("get_locked", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(message)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Member: convenience getter for the G_DBUS_MESSAGE_HEADER_FIELD_MEMBER header
// field.
//
// The function returns the following values:
//
//    - utf8 (optional): value.
//
func (message *DBusMessage) Member() string {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("get_member", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(message)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Path: convenience getter for the G_DBUS_MESSAGE_HEADER_FIELD_PATH header
// field.
//
// The function returns the following values:
//
//    - utf8 (optional): value.
//
func (message *DBusMessage) Path() string {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("get_path", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(message)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ReplySerial: convenience getter for the
// G_DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL header field.
//
// The function returns the following values:
//
//    - guint32: value.
//
func (message *DBusMessage) ReplySerial() uint32 {
	var _args [1]girepository.Argument
	var _arg0 *C.void   // out
	var _cret C.guint32 // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("get_reply_serial", _args[:], nil)
	_cret = *(*C.guint32)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(message)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// Sender: convenience getter for the G_DBUS_MESSAGE_HEADER_FIELD_SENDER header
// field.
//
// The function returns the following values:
//
//    - utf8 (optional): value.
//
func (message *DBusMessage) Sender() string {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("get_sender", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(message)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Serial gets the serial for message.
//
// The function returns the following values:
//
//    - guint32: #guint32.
//
func (message *DBusMessage) Serial() uint32 {
	var _args [1]girepository.Argument
	var _arg0 *C.void   // out
	var _cret C.guint32 // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("get_serial", _args[:], nil)
	_cret = *(*C.guint32)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(message)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// Signature: convenience getter for the G_DBUS_MESSAGE_HEADER_FIELD_SIGNATURE
// header field.
//
// The function returns the following values:
//
//    - utf8: value.
//
func (message *DBusMessage) Signature() string {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("get_signature", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(message)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Lock: if message is locked, does nothing. Otherwise locks the message.
func (message *DBusMessage) Lock() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gio", "DBusMessage").InvokeMethod("lock", _args[:], nil)

	runtime.KeepAlive(message)
}

// NewMethodErrorLiteral creates a new BusMessage that is an error reply to
// method_call_message.
//
// The function takes the following parameters:
//
//    - errorName: valid D-Bus error name.
//    - errorMessage d-Bus error message.
//
// The function returns the following values:
//
//    - dBusMessage Free with g_object_unref().
//
func (methodCallMessage *DBusMessage) NewMethodErrorLiteral(errorName, errorMessage string) *DBusMessage {
	var _args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(methodCallMessage).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(errorName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.void)(unsafe.Pointer(C.CString(errorMessage)))
	defer C.free(unsafe.Pointer(_arg2))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(**C.void)(unsafe.Pointer(&_args[2])) = _arg2

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("new_method_error_literal", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(methodCallMessage)
	runtime.KeepAlive(errorName)
	runtime.KeepAlive(errorMessage)

	var _dBusMessage *DBusMessage // out

	_dBusMessage = wrapDBusMessage(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusMessage
}

// NewMethodReply creates a new BusMessage that is a reply to
// method_call_message.
//
// The function returns the following values:
//
//    - dBusMessage Free with g_object_unref().
//
func (methodCallMessage *DBusMessage) NewMethodReply() *DBusMessage {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(methodCallMessage).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("new_method_reply", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(methodCallMessage)

	var _dBusMessage *DBusMessage // out

	_dBusMessage = wrapDBusMessage(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusMessage
}

// Print produces a human-readable multi-line description of message.
//
// The contents of the description has no ABI guarantees, the contents and
// formatting is subject to change at any time. Typical output looks something
// like this:
//
//    Flags:   none
//    Version: 0
//    Serial:  4
//    Headers:
//      path -> objectpath '/org/gtk/GDBus/TestObject'
//      interface -> 'org.gtk.GDBus.TestInterface'
//      member -> 'GimmeStdout'
//      destination -> ':1.146'
//    Body: ()
//    UNIX File Descriptors:
//      (none)
//
// or
//
//    Flags:   no-reply-expected
//    Version: 0
//    Serial:  477
//    Headers:
//      reply-serial -> uint32 4
//      destination -> ':1.159'
//      sender -> ':1.146'
//      num-unix-fds -> uint32 1
//    Body: ()
//    UNIX File Descriptors:
//      fd 12: dev=0:10,mode=020620,ino=5,uid=500,gid=5,rdev=136:2,size=0,atime=1273085037,mtime=1273085851,ctime=1272982635.
//
// The function takes the following parameters:
//
//    - indent: indentation level.
//
// The function returns the following values:
//
//    - utf8: string that should be freed with g_free().
//
func (message *DBusMessage) Print(indent uint32) string {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	_arg1 = C.guint(indent)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.guint)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gio", "DBusMessage").InvokeMethod("print", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(message)
	runtime.KeepAlive(indent)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SetBody sets the body message. As a side-effect the
// G_DBUS_MESSAGE_HEADER_FIELD_SIGNATURE header field is set to the type string
// of body (or cleared if body is NULL).
//
// If body is floating, message assumes ownership of body.
//
// The function takes the following parameters:
//
//    - body: either NULL or a #GVariant that is a tuple.
//
func (message *DBusMessage) SetBody(body *glib.Variant) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(body)))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "DBusMessage").InvokeMethod("set_body", _args[:], nil)

	runtime.KeepAlive(message)
	runtime.KeepAlive(body)
}

// SetDestination: convenience setter for the
// G_DBUS_MESSAGE_HEADER_FIELD_DESTINATION header field.
//
// The function takes the following parameters:
//
//    - value (optional) to set.
//
func (message *DBusMessage) SetDestination(value string) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	if value != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "DBusMessage").InvokeMethod("set_destination", _args[:], nil)

	runtime.KeepAlive(message)
	runtime.KeepAlive(value)
}

// SetErrorName: convenience setter for the
// G_DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME header field.
//
// The function takes the following parameters:
//
//    - value to set.
//
func (message *DBusMessage) SetErrorName(value string) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	if message != nil {
		_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	}
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "DBusMessage").InvokeMethod("set_error_name", _args[:], nil)

	runtime.KeepAlive(message)
	runtime.KeepAlive(value)
}

// SetInterface: convenience setter for the
// G_DBUS_MESSAGE_HEADER_FIELD_INTERFACE header field.
//
// The function takes the following parameters:
//
//    - value (optional) to set.
//
func (message *DBusMessage) SetInterface(value string) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	if value != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "DBusMessage").InvokeMethod("set_interface", _args[:], nil)

	runtime.KeepAlive(message)
	runtime.KeepAlive(value)
}

// SetMember: convenience setter for the G_DBUS_MESSAGE_HEADER_FIELD_MEMBER
// header field.
//
// The function takes the following parameters:
//
//    - value (optional) to set.
//
func (message *DBusMessage) SetMember(value string) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	if value != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "DBusMessage").InvokeMethod("set_member", _args[:], nil)

	runtime.KeepAlive(message)
	runtime.KeepAlive(value)
}

// SetPath: convenience setter for the G_DBUS_MESSAGE_HEADER_FIELD_PATH header
// field.
//
// The function takes the following parameters:
//
//    - value (optional) to set.
//
func (message *DBusMessage) SetPath(value string) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	if value != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "DBusMessage").InvokeMethod("set_path", _args[:], nil)

	runtime.KeepAlive(message)
	runtime.KeepAlive(value)
}

// SetReplySerial: convenience setter for the
// G_DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL header field.
//
// The function takes the following parameters:
//
//    - value to set.
//
func (message *DBusMessage) SetReplySerial(value uint32) {
	var _args [2]girepository.Argument
	var _arg0 *C.void   // out
	var _arg1 C.guint32 // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	_arg1 = C.guint32(value)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.guint32)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "DBusMessage").InvokeMethod("set_reply_serial", _args[:], nil)

	runtime.KeepAlive(message)
	runtime.KeepAlive(value)
}

// SetSender: convenience setter for the G_DBUS_MESSAGE_HEADER_FIELD_SENDER
// header field.
//
// The function takes the following parameters:
//
//    - value (optional) to set.
//
func (message *DBusMessage) SetSender(value string) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	if value != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "DBusMessage").InvokeMethod("set_sender", _args[:], nil)

	runtime.KeepAlive(message)
	runtime.KeepAlive(value)
}

// SetSerial sets the serial for message.
//
// The function takes the following parameters:
//
//    - serial: #guint32.
//
func (message *DBusMessage) SetSerial(serial uint32) {
	var _args [2]girepository.Argument
	var _arg0 *C.void   // out
	var _arg1 C.guint32 // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	_arg1 = C.guint32(serial)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.guint32)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "DBusMessage").InvokeMethod("set_serial", _args[:], nil)

	runtime.KeepAlive(message)
	runtime.KeepAlive(serial)
}

// SetSignature: convenience setter for the
// G_DBUS_MESSAGE_HEADER_FIELD_SIGNATURE header field.
//
// The function takes the following parameters:
//
//    - value (optional) to set.
//
func (message *DBusMessage) SetSignature(value string) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	if value != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "DBusMessage").InvokeMethod("set_signature", _args[:], nil)

	runtime.KeepAlive(message)
	runtime.KeepAlive(value)
}

// ToGError: if message is not of type G_DBUS_MESSAGE_TYPE_ERROR does nothing
// and returns FALSE.
//
// Otherwise this method encodes the error in message as a #GError using
// g_dbus_error_set_dbus_error() using the information in the
// G_DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME header field of message as well as the
// first string item in message's body.
func (message *DBusMessage) ToGError() error {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gio", "DBusMessage").InvokeMethod("to_gerror", _args[:], nil)

	runtime.KeepAlive(message)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// DBusMessageBytesNeeded: utility function to calculate how many bytes are
// needed to completely deserialize the D-Bus message stored at blob.
//
// The function takes the following parameters:
//
//    - blob representing a binary D-Bus message.
//
// The function returns the following values:
//
//    - gssize: number of bytes needed or -1 if error is set (e.g. if blob
//      contains invalid data or not enough data is available to determine the
//      size).
//
func DBusMessageBytesNeeded(blob []byte) (int, error) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gsize
	var _cret C.gssize // in
	var _cerr *C.void  // in

	_arg1 = (C.gsize)(len(blob))
	if len(blob) > 0 {
		_arg0 = (*C.void)(unsafe.Pointer(&blob[0]))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "bytes_needed").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(blob)

	var _gssize int  // out
	var _goerr error // out
	_out1 = *(*C.gssize)(unsafe.Pointer(&_outs[1]))

	_gssize = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gssize, _goerr
}
