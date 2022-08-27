// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// NewDBusMessage creates a new empty BusMessage.
//
// The function returns the following values:
//
//    - dBusMessage Free with g_object_unref().
//
func NewDBusMessage() *DBusMessage {
	var _cret *C.GDBusMessage // in

	_cret = C.g_dbus_message_new()

	var _dBusMessage *DBusMessage // out

	_dBusMessage = wrapDBusMessage(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusMessage
}

// NewDBusMessageFromBlob creates a new BusMessage from the data stored at blob.
// The byte order that the message was in can be retrieved using
// g_dbus_message_get_byte_order().
//
// If the blob cannot be parsed, contains invalid fields, or contains invalid
// headers, G_IO_ERROR_INVALID_ARGUMENT will be returned.
//
// The function takes the following parameters:
//
//    - blob representing a binary D-Bus message.
//    - capabilities describing what protocol features are supported.
//
// The function returns the following values:
//
//    - dBusMessage: new BusMessage or NULL if error is set. Free with
//      g_object_unref().
//
func NewDBusMessageFromBlob(blob []byte, capabilities DBusCapabilityFlags) (*DBusMessage, error) {
	var _arg1 *C.guchar // out
	var _arg2 C.gsize
	var _arg3 C.GDBusCapabilityFlags // out
	var _cret *C.GDBusMessage        // in
	var _cerr *C.GError              // in

	_arg2 = (C.gsize)(len(blob))
	if len(blob) > 0 {
		_arg1 = (*C.guchar)(unsafe.Pointer(&blob[0]))
	}
	_arg3 = C.GDBusCapabilityFlags(capabilities)

	_cret = C.g_dbus_message_new_from_blob(_arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(blob)
	runtime.KeepAlive(capabilities)

	var _dBusMessage *DBusMessage // out
	var _goerr error              // out

	_dBusMessage = wrapDBusMessage(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dBusMessage, _goerr
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
	var _arg1 *C.gchar        // out
	var _arg2 *C.gchar        // out
	var _arg3 *C.gchar        // out
	var _arg4 *C.gchar        // out
	var _cret *C.GDBusMessage // in

	if name != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg2))
	if interface_ != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(interface_)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(method)))
	defer C.free(unsafe.Pointer(_arg4))

	_cret = C.g_dbus_message_new_method_call(_arg1, _arg2, _arg3, _arg4)
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
	var _arg1 *C.gchar        // out
	var _arg2 *C.gchar        // out
	var _arg3 *C.gchar        // out
	var _cret *C.GDBusMessage // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(interface_)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(signal)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.g_dbus_message_new_signal(_arg1, _arg2, _arg3)
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
	var _arg0 *C.GDBusMessage // out
	var _cret *C.GDBusMessage // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	_cret = C.g_dbus_message_copy(_arg0, &_cerr)
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
	var _arg0 *C.GDBusMessage // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	_cret = C.g_dbus_message_get_arg0(_arg0)
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
	var _arg0 *C.GDBusMessage // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	_cret = C.g_dbus_message_get_body(_arg0)
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
	var _arg0 *C.GDBusMessage // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	_cret = C.g_dbus_message_get_destination(_arg0)
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
	var _arg0 *C.GDBusMessage // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	_cret = C.g_dbus_message_get_error_name(_arg0)
	runtime.KeepAlive(message)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Flags gets the flags for message.
//
// The function returns the following values:
//
//    - dBusMessageFlags flags that are set (typically values from the
//      BusMessageFlags enumeration bitwise ORed together).
//
func (message *DBusMessage) Flags() DBusMessageFlags {
	var _arg0 *C.GDBusMessage     // out
	var _cret C.GDBusMessageFlags // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	_cret = C.g_dbus_message_get_flags(_arg0)
	runtime.KeepAlive(message)

	var _dBusMessageFlags DBusMessageFlags // out

	_dBusMessageFlags = DBusMessageFlags(_cret)

	return _dBusMessageFlags
}

// Header gets a header field on message.
//
// The caller is responsible for checking the type of the returned #GVariant
// matches what is expected.
//
// The function takes the following parameters:
//
//    - headerField: 8-bit unsigned integer (typically a value from the
//      BusMessageHeaderField enumeration).
//
// The function returns the following values:
//
//    - variant (optional) with the value if the header was found, NULL
//      otherwise. Do not free, it is owned by message.
//
func (message *DBusMessage) Header(headerField DBusMessageHeaderField) *glib.Variant {
	var _arg0 *C.GDBusMessage           // out
	var _arg1 C.GDBusMessageHeaderField // out
	var _cret *C.GVariant               // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	_arg1 = C.GDBusMessageHeaderField(headerField)

	_cret = C.g_dbus_message_get_header(_arg0, _arg1)
	runtime.KeepAlive(message)
	runtime.KeepAlive(headerField)

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

// HeaderFields gets an array of all header fields on message that are set.
//
// The function returns the following values:
//
//    - guint8s: array of header fields terminated by
//      G_DBUS_MESSAGE_HEADER_FIELD_INVALID. Each element is a #guchar. Free with
//      g_free().
//
func (message *DBusMessage) HeaderFields() []byte {
	var _arg0 *C.GDBusMessage // out
	var _cret *C.guchar       // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	_cret = C.g_dbus_message_get_header_fields(_arg0)
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
	var _arg0 *C.GDBusMessage // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	_cret = C.g_dbus_message_get_interface(_arg0)
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
	var _arg0 *C.GDBusMessage // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	_cret = C.g_dbus_message_get_locked(_arg0)
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
	var _arg0 *C.GDBusMessage // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	_cret = C.g_dbus_message_get_member(_arg0)
	runtime.KeepAlive(message)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// MessageType gets the type of message.
//
// The function returns the following values:
//
//    - dBusMessageType: 8-bit unsigned integer (typically a value from the
//      BusMessageType enumeration).
//
func (message *DBusMessage) MessageType() DBusMessageType {
	var _arg0 *C.GDBusMessage    // out
	var _cret C.GDBusMessageType // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	_cret = C.g_dbus_message_get_message_type(_arg0)
	runtime.KeepAlive(message)

	var _dBusMessageType DBusMessageType // out

	_dBusMessageType = DBusMessageType(_cret)

	return _dBusMessageType
}

// Path: convenience getter for the G_DBUS_MESSAGE_HEADER_FIELD_PATH header
// field.
//
// The function returns the following values:
//
//    - utf8 (optional): value.
//
func (message *DBusMessage) Path() string {
	var _arg0 *C.GDBusMessage // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	_cret = C.g_dbus_message_get_path(_arg0)
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
	var _arg0 *C.GDBusMessage // out
	var _cret C.guint32       // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	_cret = C.g_dbus_message_get_reply_serial(_arg0)
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
	var _arg0 *C.GDBusMessage // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	_cret = C.g_dbus_message_get_sender(_arg0)
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
	var _arg0 *C.GDBusMessage // out
	var _cret C.guint32       // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	_cret = C.g_dbus_message_get_serial(_arg0)
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
	var _arg0 *C.GDBusMessage // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	_cret = C.g_dbus_message_get_signature(_arg0)
	runtime.KeepAlive(message)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Lock: if message is locked, does nothing. Otherwise locks the message.
func (message *DBusMessage) Lock() {
	var _arg0 *C.GDBusMessage // out

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	C.g_dbus_message_lock(_arg0)
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
	var _arg0 *C.GDBusMessage // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.gchar        // out
	var _cret *C.GDBusMessage // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(methodCallMessage).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(errorName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(errorMessage)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_dbus_message_new_method_error_literal(_arg0, _arg1, _arg2)
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
	var _arg0 *C.GDBusMessage // out
	var _cret *C.GDBusMessage // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(methodCallMessage).Native()))

	_cret = C.g_dbus_message_new_method_reply(_arg0)
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
func (message *DBusMessage) Print(indent uint) string {
	var _arg0 *C.GDBusMessage // out
	var _arg1 C.guint         // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	_arg1 = C.guint(indent)

	_cret = C.g_dbus_message_print(_arg0, _arg1)
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
	var _arg0 *C.GDBusMessage // out
	var _arg1 *C.GVariant     // out

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(body)))

	C.g_dbus_message_set_body(_arg0, _arg1)
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
	var _arg0 *C.GDBusMessage // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	if value != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.g_dbus_message_set_destination(_arg0, _arg1)
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
	var _arg0 *C.GDBusMessage // out
	var _arg1 *C.gchar        // out

	if message != nil {
		_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_dbus_message_set_error_name(_arg0, _arg1)
	runtime.KeepAlive(message)
	runtime.KeepAlive(value)
}

// SetFlags sets the flags to set on message.
//
// The function takes the following parameters:
//
//    - flags flags for message that are set (typically values from the
//      BusMessageFlags enumeration bitwise ORed together).
//
func (message *DBusMessage) SetFlags(flags DBusMessageFlags) {
	var _arg0 *C.GDBusMessage     // out
	var _arg1 C.GDBusMessageFlags // out

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	_arg1 = C.GDBusMessageFlags(flags)

	C.g_dbus_message_set_flags(_arg0, _arg1)
	runtime.KeepAlive(message)
	runtime.KeepAlive(flags)
}

// SetHeader sets a header field on message.
//
// If value is floating, message assumes ownership of value.
//
// The function takes the following parameters:
//
//    - headerField: 8-bit unsigned integer (typically a value from the
//      BusMessageHeaderField enumeration).
//    - value (optional) to set the header field or NULL to clear the header
//      field.
//
func (message *DBusMessage) SetHeader(headerField DBusMessageHeaderField, value *glib.Variant) {
	var _arg0 *C.GDBusMessage           // out
	var _arg1 C.GDBusMessageHeaderField // out
	var _arg2 *C.GVariant               // out

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	_arg1 = C.GDBusMessageHeaderField(headerField)
	if value != nil {
		_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))
	}

	C.g_dbus_message_set_header(_arg0, _arg1, _arg2)
	runtime.KeepAlive(message)
	runtime.KeepAlive(headerField)
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
	var _arg0 *C.GDBusMessage // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	if value != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.g_dbus_message_set_interface(_arg0, _arg1)
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
	var _arg0 *C.GDBusMessage // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	if value != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.g_dbus_message_set_member(_arg0, _arg1)
	runtime.KeepAlive(message)
	runtime.KeepAlive(value)
}

// SetMessageType sets message to be of type.
//
// The function takes the following parameters:
//
//    - typ: 8-bit unsigned integer (typically a value from the BusMessageType
//      enumeration).
//
func (message *DBusMessage) SetMessageType(typ DBusMessageType) {
	var _arg0 *C.GDBusMessage    // out
	var _arg1 C.GDBusMessageType // out

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	_arg1 = C.GDBusMessageType(typ)

	C.g_dbus_message_set_message_type(_arg0, _arg1)
	runtime.KeepAlive(message)
	runtime.KeepAlive(typ)
}

// SetPath: convenience setter for the G_DBUS_MESSAGE_HEADER_FIELD_PATH header
// field.
//
// The function takes the following parameters:
//
//    - value (optional) to set.
//
func (message *DBusMessage) SetPath(value string) {
	var _arg0 *C.GDBusMessage // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	if value != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.g_dbus_message_set_path(_arg0, _arg1)
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
	var _arg0 *C.GDBusMessage // out
	var _arg1 C.guint32       // out

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	_arg1 = C.guint32(value)

	C.g_dbus_message_set_reply_serial(_arg0, _arg1)
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
	var _arg0 *C.GDBusMessage // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	if value != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.g_dbus_message_set_sender(_arg0, _arg1)
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
	var _arg0 *C.GDBusMessage // out
	var _arg1 C.guint32       // out

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	_arg1 = C.guint32(serial)

	C.g_dbus_message_set_serial(_arg0, _arg1)
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
	var _arg0 *C.GDBusMessage // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	if value != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.g_dbus_message_set_signature(_arg0, _arg1)
	runtime.KeepAlive(message)
	runtime.KeepAlive(value)
}

// ToBlob serializes message to a blob. The byte order returned by
// g_dbus_message_get_byte_order() will be used.
//
// The function takes the following parameters:
//
//    - capabilities describing what protocol features are supported.
//
// The function returns the following values:
//
//    - guint8s: pointer to a valid binary D-Bus message of out_size bytes
//      generated by message or NULL if error is set. Free with g_free().
//
func (message *DBusMessage) ToBlob(capabilities DBusCapabilityFlags) ([]byte, error) {
	var _arg0 *C.GDBusMessage        // out
	var _arg2 C.GDBusCapabilityFlags // out
	var _cret *C.guchar              // in
	var _arg1 C.gsize                // in
	var _cerr *C.GError              // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))
	_arg2 = C.GDBusCapabilityFlags(capabilities)

	_cret = C.g_dbus_message_to_blob(_arg0, &_arg1, _arg2, &_cerr)
	runtime.KeepAlive(message)
	runtime.KeepAlive(capabilities)

	var _guint8s []byte // out
	var _goerr error    // out

	defer C.free(unsafe.Pointer(_cret))
	_guint8s = make([]byte, _arg1)
	copy(_guint8s, unsafe.Slice((*byte)(unsafe.Pointer(_cret)), _arg1))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _guint8s, _goerr
}

// ToGError: if message is not of type G_DBUS_MESSAGE_TYPE_ERROR does nothing
// and returns FALSE.
//
// Otherwise this method encodes the error in message as a #GError using
// g_dbus_error_set_dbus_error() using the information in the
// G_DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME header field of message as well as the
// first string item in message's body.
func (message *DBusMessage) ToGError() error {
	var _arg0 *C.GDBusMessage // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GDBusMessage)(unsafe.Pointer(coreglib.InternObject(message).Native()))

	C.g_dbus_message_to_gerror(_arg0, &_cerr)
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
	var _arg1 *C.guchar // out
	var _arg2 C.gsize
	var _cret C.gssize  // in
	var _cerr *C.GError // in

	_arg2 = (C.gsize)(len(blob))
	if len(blob) > 0 {
		_arg1 = (*C.guchar)(unsafe.Pointer(&blob[0]))
	}

	_cret = C.g_dbus_message_bytes_needed(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(blob)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gssize, _goerr
}
