// Code generated by girgen. DO NOT EDIT.

package gio

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// SocketListenerClass class structure for Listener.
//
// An instance of this type is always passed by reference.
type SocketListenerClass struct {
	*socketListenerClass
}

// socketListenerClass is the struct that's finalized.
type socketListenerClass struct {
	native *C.GSocketListenerClass
}
