// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// TCPConnectionClass: instance of this type is always passed by reference.
type TCPConnectionClass struct {
	*tcpConnectionClass
}

// tcpConnectionClass is the struct that's finalized.
type tcpConnectionClass struct {
	native *C.GTcpConnectionClass
}

func (t *TCPConnectionClass) ParentClass() *SocketConnectionClass {
	valptr := &t.native.parent_class
	var _v *SocketConnectionClass // out
	_v = (*SocketConnectionClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
