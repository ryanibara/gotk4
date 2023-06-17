// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <glib.h>
import "C"

// GType values.
var (
	GTypeChecksum = coreglib.Type(C.g_checksum_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeChecksum, F: marshalChecksum},
	})
}

// ChecksumType: hashing algorithm to be used by #GChecksum when performing the
// digest of some data.
//
// Note that the Type enumeration may be extended at a later date to include new
// hashing algorithm types.
type ChecksumType C.gint

const (
	// ChecksumMD5: use the MD5 hashing algorithm.
	ChecksumMD5 ChecksumType = iota
	// ChecksumSHA1: use the SHA-1 hashing algorithm.
	ChecksumSHA1
	// ChecksumSHA256: use the SHA-256 hashing algorithm.
	ChecksumSHA256
	// ChecksumSHA512: use the SHA-512 hashing algorithm (Since: 2.36).
	ChecksumSHA512
	// ChecksumSHA384: use the SHA-384 hashing algorithm (Since: 2.51).
	ChecksumSHA384
)

// String returns the name in string for ChecksumType.
func (c ChecksumType) String() string {
	switch c {
	case ChecksumMD5:
		return "MD5"
	case ChecksumSHA1:
		return "SHA1"
	case ChecksumSHA256:
		return "SHA256"
	case ChecksumSHA512:
		return "SHA512"
	case ChecksumSHA384:
		return "SHA384"
	default:
		return fmt.Sprintf("ChecksumType(%d)", c)
	}
}

// ComputeChecksumForData computes the checksum for a binary data of length.
// This is a convenience wrapper for g_checksum_new(), g_checksum_get_string()
// and g_checksum_free().
//
// The hexadecimal string returned will be in lower case.
//
// The function takes the following parameters:
//
//   - checksumType: Type.
//   - data: binary blob to compute the digest of.
//
// The function returns the following values:
//
//   - utf8 (optional): digest of the binary data as a string in hexadecimal,
//     or NULL if g_checksum_new() fails for checksum_type. The returned string
//     should be freed with g_free() when done using it.
//
func ComputeChecksumForData(checksumType ChecksumType, data []byte) string {
	var _arg1 C.GChecksumType // out
	var _arg2 *C.guchar       // out
	var _arg3 C.gsize
	var _cret *C.gchar // in

	_arg1 = C.GChecksumType(checksumType)
	_arg3 = (C.gsize)(len(data))
	if len(data) > 0 {
		_arg2 = (*C.guchar)(unsafe.Pointer(&data[0]))
	}

	_cret = C.g_compute_checksum_for_data(_arg1, _arg2, _arg3)
	runtime.KeepAlive(checksumType)
	runtime.KeepAlive(data)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// ComputeChecksumForString computes the checksum of a string.
//
// The hexadecimal string returned will be in lower case.
//
// The function takes the following parameters:
//
//   - checksumType: Type.
//   - str: string to compute the checksum of.
//   - length of the string, or -1 if the string is null-terminated.
//
// The function returns the following values:
//
//   - utf8 (optional): checksum as a hexadecimal string, or NULL if
//     g_checksum_new() fails for checksum_type. The returned string should be
//     freed with g_free() when done using it.
//
func ComputeChecksumForString(checksumType ChecksumType, str string, length int) string {
	var _arg1 C.GChecksumType // out
	var _arg2 *C.gchar        // out
	var _arg3 C.gssize        // out
	var _cret *C.gchar        // in

	_arg1 = C.GChecksumType(checksumType)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gssize(length)

	_cret = C.g_compute_checksum_for_string(_arg1, _arg2, _arg3)
	runtime.KeepAlive(checksumType)
	runtime.KeepAlive(str)
	runtime.KeepAlive(length)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Checksum: opaque structure representing a checksumming operation.
// To create a new GChecksum, use g_checksum_new(). To free a GChecksum,
// use g_checksum_free().
//
// An instance of this type is always passed by reference.
type Checksum struct {
	*checksum
}

// checksum is the struct that's finalized.
type checksum struct {
	native *C.GChecksum
}

func marshalChecksum(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Checksum{&checksum{(*C.GChecksum)(b)}}, nil
}

// NewChecksum constructs a struct Checksum.
func NewChecksum(checksumType ChecksumType) *Checksum {
	var _arg1 C.GChecksumType // out
	var _cret *C.GChecksum    // in

	_arg1 = C.GChecksumType(checksumType)

	_cret = C.g_checksum_new(_arg1)
	runtime.KeepAlive(checksumType)

	var _checksum *Checksum // out

	if _cret != nil {
		_checksum = (*Checksum)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_checksum)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_checksum_free((*C.GChecksum)(intern.C))
			},
		)
	}

	return _checksum
}

// Copy copies a #GChecksum. If checksum has been closed, by calling
// g_checksum_get_string() or g_checksum_get_digest(), the copied checksum will
// be closed as well.
//
// The function returns the following values:
//
//   - ret: copy of the passed #GChecksum. Use g_checksum_free() when finished
//     using it.
//
func (checksum *Checksum) Copy() *Checksum {
	var _arg0 *C.GChecksum // out
	var _cret *C.GChecksum // in

	_arg0 = (*C.GChecksum)(gextras.StructNative(unsafe.Pointer(checksum)))

	_cret = C.g_checksum_copy(_arg0)
	runtime.KeepAlive(checksum)

	var _ret *Checksum // out

	_ret = (*Checksum)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_ret)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_checksum_free((*C.GChecksum)(intern.C))
		},
	)

	return _ret
}

// String gets the digest as a hexadecimal string.
//
// Once this function has been called the #GChecksum can no longer be updated
// with g_checksum_update().
//
// The hexadecimal characters will be lower case.
//
// The function returns the following values:
//
//   - utf8: hexadecimal representation of the checksum. The returned string is
//     owned by the checksum and should not be modified or freed.
//
func (checksum *Checksum) String() string {
	var _arg0 *C.GChecksum // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GChecksum)(gextras.StructNative(unsafe.Pointer(checksum)))

	_cret = C.g_checksum_get_string(_arg0)
	runtime.KeepAlive(checksum)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Reset resets the state of the checksum back to its initial state.
func (checksum *Checksum) Reset() {
	var _arg0 *C.GChecksum // out

	_arg0 = (*C.GChecksum)(gextras.StructNative(unsafe.Pointer(checksum)))

	C.g_checksum_reset(_arg0)
	runtime.KeepAlive(checksum)
}

// Update feeds data into an existing #GChecksum. The checksum must still be
// open, that is g_checksum_get_string() or g_checksum_get_digest() must not
// have been called on checksum.
//
// The function takes the following parameters:
//
//   - data: buffer used to compute the checksum.
//
func (checksum *Checksum) Update(data []byte) {
	var _arg0 *C.GChecksum // out
	var _arg1 *C.guchar    // out
	var _arg2 C.gssize

	_arg0 = (*C.GChecksum)(gextras.StructNative(unsafe.Pointer(checksum)))
	_arg2 = (C.gssize)(len(data))
	if len(data) > 0 {
		_arg1 = (*C.guchar)(unsafe.Pointer(&data[0]))
	}

	C.g_checksum_update(_arg0, _arg1, _arg2)
	runtime.KeepAlive(checksum)
	runtime.KeepAlive(data)
}

// ChecksumTypeGetLength gets the length in bytes of digests of type
// checksum_type.
//
// The function takes the following parameters:
//
//   - checksumType: Type.
//
// The function returns the following values:
//
//   - gssize: checksum length, or -1 if checksum_type is not supported.
//
func ChecksumTypeGetLength(checksumType ChecksumType) int {
	var _arg1 C.GChecksumType // out
	var _cret C.gssize        // in

	_arg1 = C.GChecksumType(checksumType)

	_cret = C.g_checksum_type_get_length(_arg1)
	runtime.KeepAlive(checksumType)

	var _gssize int // out

	_gssize = int(_cret)

	return _gssize
}
