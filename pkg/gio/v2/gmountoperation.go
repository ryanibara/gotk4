// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_mount_operation_get_type()), F: marshalMountOperationer},
	})
}

// MountOperationOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type MountOperationOverrider interface {
	Aborted()
	AskPassword(message, defaultUser, defaultDomain string, flags AskPasswordFlags)
	// AskQuestion: virtual implementation of Operation::ask-question.
	AskQuestion(message string, choices []string)
	// Reply emits the Operation::reply signal.
	Reply(result MountOperationResult)
	// ShowProcesses: virtual implementation of Operation::show-processes.
	ShowProcesses(message string, processes []glib.Pid, choices []string)
	ShowUnmountProgress(message string, timeLeft, bytesLeft int64)
}

// MountOperation provides a mechanism for interacting with the user. It can be
// used for authenticating mountable operations, such as loop mounting files,
// hard drive partitions or server locations. It can also be used to ask the
// user questions or show a list of applications preventing unmount or eject
// operations from completing.
//
// Note that Operation is used for more than just #GMount objects – for example
// it is also used in g_drive_start() and g_drive_stop().
//
// Users should instantiate a subclass of this that implements all the various
// callbacks to show the required dialogs, such as MountOperation. If no user
// interaction is desired (for example when automounting filesystems at login
// time), usually NULL can be passed, see each method taking a Operation for
// details.
//
// The term ‘TCRYPT’ is used to mean ‘compatible with TrueCrypt and VeraCrypt’.
// TrueCrypt (https://en.wikipedia.org/wiki/TrueCrypt) is a discontinued system
// for encrypting file containers, partitions or whole disks, typically used
// with Windows. VeraCrypt (https://www.veracrypt.fr/) is a maintained fork of
// TrueCrypt with various improvements and auditing fixes.
type MountOperation struct {
	*externglib.Object
}

func wrapMountOperation(obj *externglib.Object) *MountOperation {
	return &MountOperation{
		Object: obj,
	}
}

func marshalMountOperationer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMountOperation(obj), nil
}

// NewMountOperation creates a new mount operation.
func NewMountOperation() *MountOperation {
	var _cret *C.GMountOperation // in

	_cret = C.g_mount_operation_new()

	var _mountOperation *MountOperation // out

	_mountOperation = wrapMountOperation(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mountOperation
}

// Anonymous: check to see whether the mount operation is being used for an
// anonymous user.
func (op *MountOperation) Anonymous() bool {
	var _arg0 *C.GMountOperation // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))

	_cret = C.g_mount_operation_get_anonymous(_arg0)
	runtime.KeepAlive(op)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Choice gets a choice from the mount operation.
func (op *MountOperation) Choice() int {
	var _arg0 *C.GMountOperation // out
	var _cret C.int              // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))

	_cret = C.g_mount_operation_get_choice(_arg0)
	runtime.KeepAlive(op)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Domain gets the domain of the mount operation.
func (op *MountOperation) Domain() string {
	var _arg0 *C.GMountOperation // out
	var _cret *C.char            // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))

	_cret = C.g_mount_operation_get_domain(_arg0)
	runtime.KeepAlive(op)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// IsTcryptHiddenVolume: check to see whether the mount operation is being used
// for a TCRYPT hidden volume.
func (op *MountOperation) IsTcryptHiddenVolume() bool {
	var _arg0 *C.GMountOperation // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))

	_cret = C.g_mount_operation_get_is_tcrypt_hidden_volume(_arg0)
	runtime.KeepAlive(op)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsTcryptSystemVolume: check to see whether the mount operation is being used
// for a TCRYPT system volume.
func (op *MountOperation) IsTcryptSystemVolume() bool {
	var _arg0 *C.GMountOperation // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))

	_cret = C.g_mount_operation_get_is_tcrypt_system_volume(_arg0)
	runtime.KeepAlive(op)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Password gets a password from the mount operation.
func (op *MountOperation) Password() string {
	var _arg0 *C.GMountOperation // out
	var _cret *C.char            // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))

	_cret = C.g_mount_operation_get_password(_arg0)
	runtime.KeepAlive(op)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// PasswordSave gets the state of saving passwords for the mount operation.
func (op *MountOperation) PasswordSave() PasswordSave {
	var _arg0 *C.GMountOperation // out
	var _cret C.GPasswordSave    // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))

	_cret = C.g_mount_operation_get_password_save(_arg0)
	runtime.KeepAlive(op)

	var _passwordSave PasswordSave // out

	_passwordSave = PasswordSave(_cret)

	return _passwordSave
}

// Pim gets a PIM from the mount operation.
func (op *MountOperation) Pim() uint {
	var _arg0 *C.GMountOperation // out
	var _cret C.guint            // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))

	_cret = C.g_mount_operation_get_pim(_arg0)
	runtime.KeepAlive(op)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Username: get the user name from the mount operation.
func (op *MountOperation) Username() string {
	var _arg0 *C.GMountOperation // out
	var _cret *C.char            // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))

	_cret = C.g_mount_operation_get_username(_arg0)
	runtime.KeepAlive(op)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Reply emits the Operation::reply signal.
//
// The function takes the following parameters:
//
//    - result: OperationResult.
//
func (op *MountOperation) Reply(result MountOperationResult) {
	var _arg0 *C.GMountOperation      // out
	var _arg1 C.GMountOperationResult // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))
	_arg1 = C.GMountOperationResult(result)

	C.g_mount_operation_reply(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(result)
}

// SetAnonymous sets the mount operation to use an anonymous user if anonymous
// is TRUE.
//
// The function takes the following parameters:
//
//    - anonymous: boolean value.
//
func (op *MountOperation) SetAnonymous(anonymous bool) {
	var _arg0 *C.GMountOperation // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))
	if anonymous {
		_arg1 = C.TRUE
	}

	C.g_mount_operation_set_anonymous(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(anonymous)
}

// SetChoice sets a default choice for the mount operation.
//
// The function takes the following parameters:
//
//    - choice: integer.
//
func (op *MountOperation) SetChoice(choice int) {
	var _arg0 *C.GMountOperation // out
	var _arg1 C.int              // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))
	_arg1 = C.int(choice)

	C.g_mount_operation_set_choice(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(choice)
}

// SetDomain sets the mount operation's domain.
//
// The function takes the following parameters:
//
//    - domain to set.
//
func (op *MountOperation) SetDomain(domain string) {
	var _arg0 *C.GMountOperation // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))
	if domain != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(domain)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.g_mount_operation_set_domain(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(domain)
}

// SetIsTcryptHiddenVolume sets the mount operation to use a hidden volume if
// hidden_volume is TRUE.
//
// The function takes the following parameters:
//
//    - hiddenVolume: boolean value.
//
func (op *MountOperation) SetIsTcryptHiddenVolume(hiddenVolume bool) {
	var _arg0 *C.GMountOperation // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))
	if hiddenVolume {
		_arg1 = C.TRUE
	}

	C.g_mount_operation_set_is_tcrypt_hidden_volume(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(hiddenVolume)
}

// SetIsTcryptSystemVolume sets the mount operation to use a system volume if
// system_volume is TRUE.
//
// The function takes the following parameters:
//
//    - systemVolume: boolean value.
//
func (op *MountOperation) SetIsTcryptSystemVolume(systemVolume bool) {
	var _arg0 *C.GMountOperation // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))
	if systemVolume {
		_arg1 = C.TRUE
	}

	C.g_mount_operation_set_is_tcrypt_system_volume(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(systemVolume)
}

// SetPassword sets the mount operation's password to password.
//
// The function takes the following parameters:
//
//    - password to set.
//
func (op *MountOperation) SetPassword(password string) {
	var _arg0 *C.GMountOperation // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))
	if password != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(password)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.g_mount_operation_set_password(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(password)
}

// SetPasswordSave sets the state of saving passwords for the mount operation.
//
// The function takes the following parameters:
//
//    - save: set of Save flags.
//
func (op *MountOperation) SetPasswordSave(save PasswordSave) {
	var _arg0 *C.GMountOperation // out
	var _arg1 C.GPasswordSave    // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))
	_arg1 = C.GPasswordSave(save)

	C.g_mount_operation_set_password_save(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(save)
}

// SetPim sets the mount operation's PIM to pim.
//
// The function takes the following parameters:
//
//    - pim: unsigned integer.
//
func (op *MountOperation) SetPim(pim uint) {
	var _arg0 *C.GMountOperation // out
	var _arg1 C.guint            // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))
	_arg1 = C.guint(pim)

	C.g_mount_operation_set_pim(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(pim)
}

// SetUsername sets the user name within op to username.
//
// The function takes the following parameters:
//
//    - username: input username.
//
func (op *MountOperation) SetUsername(username string) {
	var _arg0 *C.GMountOperation // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(op.Native()))
	if username != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(username)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.g_mount_operation_set_username(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(username)
}

// ConnectAborted: emitted by the backend when e.g. a device becomes unavailable
// while a mount operation is in progress.
//
// Implementations of GMountOperation should handle this signal by dismissing
// open password dialogs.
func (op *MountOperation) ConnectAborted(f func()) externglib.SignalHandle {
	return op.Connect("aborted", f)
}

// ConnectAskPassword: emitted when a mount operation asks the user for a
// password.
//
// If the message contains a line break, the first line should be presented as a
// heading. For example, it may be used as the primary text in a MessageDialog.
func (op *MountOperation) ConnectAskPassword(f func(message, defaultUser, defaultDomain string, flags AskPasswordFlags)) externglib.SignalHandle {
	return op.Connect("ask-password", f)
}

// ConnectAskQuestion: emitted when asking the user a question and gives a list
// of choices for the user to choose from.
//
// If the message contains a line break, the first line should be presented as a
// heading. For example, it may be used as the primary text in a MessageDialog.
func (op *MountOperation) ConnectAskQuestion(f func(message string, choices []string)) externglib.SignalHandle {
	return op.Connect("ask-question", f)
}

// ConnectReply: emitted when the user has replied to the mount operation.
func (op *MountOperation) ConnectReply(f func(result MountOperationResult)) externglib.SignalHandle {
	return op.Connect("reply", f)
}

// ConnectShowProcesses: emitted when one or more processes are blocking an
// operation e.g. unmounting/ejecting a #GMount or stopping a #GDrive.
//
// Note that this signal may be emitted several times to update the list of
// blocking processes as processes close files. The application should only
// respond with g_mount_operation_reply() to the latest signal (setting
// Operation:choice to the choice the user made).
//
// If the message contains a line break, the first line should be presented as a
// heading. For example, it may be used as the primary text in a MessageDialog.
func (op *MountOperation) ConnectShowProcesses(f func(message string, processes []glib.Pid, choices []string)) externglib.SignalHandle {
	return op.Connect("show-processes", f)
}

// ConnectShowUnmountProgress: emitted when an unmount operation has been busy
// for more than some time (typically 1.5 seconds).
//
// When unmounting or ejecting a volume, the kernel might need to flush pending
// data in its buffers to the volume stable storage, and this operation can take
// a considerable amount of time. This signal may be emitted several times as
// long as the unmount operation is outstanding, and then one last time when the
// operation is completed, with bytes_left set to zero.
//
// Implementations of GMountOperation should handle this signal by showing an UI
// notification, and then dismiss it, or show another notification of
// completion, when bytes_left reaches zero.
//
// If the message contains a line break, the first line should be presented as a
// heading. For example, it may be used as the primary text in a MessageDialog.
func (op *MountOperation) ConnectShowUnmountProgress(f func(message string, timeLeft, bytesLeft int64)) externglib.SignalHandle {
	return op.Connect("show-unmount-progress", f)
}
