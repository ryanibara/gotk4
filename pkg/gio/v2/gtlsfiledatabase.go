// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtlsfiledatabase.go.
var GTypeTLSFileDatabase = coreglib.Type(C.g_tls_file_database_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeTLSFileDatabase, F: marshalTLSFileDatabase},
	})
}

// TLSFileDatabaseOverrider contains methods that are overridable.
type TLSFileDatabaseOverrider interface {
}

// TLSFileDatabase is implemented by Database objects which load their
// certificate information from a file. It is an interface which TLS library
// specific subtypes implement.
//
// TLSFileDatabase wraps an interface. This means the user can get the
// underlying type by calling Cast().
type TLSFileDatabase struct {
	_ [0]func() // equal guard
	TLSDatabase
}

var (
	_ TLSDatabaser = (*TLSFileDatabase)(nil)
)

// TLSFileDatabaser describes TLSFileDatabase's interface methods.
type TLSFileDatabaser interface {
	coreglib.Objector

	baseTLSFileDatabase() *TLSFileDatabase
}

var _ TLSFileDatabaser = (*TLSFileDatabase)(nil)

func ifaceInitTLSFileDatabaser(gifacePtr, data C.gpointer) {
}

func wrapTLSFileDatabase(obj *coreglib.Object) *TLSFileDatabase {
	return &TLSFileDatabase{
		TLSDatabase: TLSDatabase{
			Object: obj,
		},
	}
}

func marshalTLSFileDatabase(p uintptr) (interface{}, error) {
	return wrapTLSFileDatabase(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *TLSFileDatabase) baseTLSFileDatabase() *TLSFileDatabase {
	return v
}

// BaseTLSFileDatabase returns the underlying base object.
func BaseTLSFileDatabase(obj TLSFileDatabaser) *TLSFileDatabase {
	return obj.baseTLSFileDatabase()
}

// NewTLSFileDatabase creates a new FileDatabase which uses anchor certificate
// authorities in anchors to verify certificate chains.
//
// The certificates in anchors must be PEM encoded.
//
// The function takes the following parameters:
//
//    - anchors: filename of anchor certificate authorities.
//
// The function returns the following values:
//
//    - tlsFileDatabase: new FileDatabase, or NULL on error.
//
func NewTLSFileDatabase(anchors string) (*TLSFileDatabase, error) {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(anchors)))
	defer C.free(unsafe.Pointer(_arg0))
	*(*string)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "new").Invoke(args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(anchors)

	var _tlsFileDatabase *TLSFileDatabase // out
	var _goerr error                      // out

	_tlsFileDatabase = wrapTLSFileDatabase(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsFileDatabase, _goerr
}
