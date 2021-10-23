// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_snapshot_get_type()), F: marshalSnapshotter},
	})
}

// Snapshot: base type for snapshot operations.
//
// The subclass of GdkSnapshot used by GTK is gtk.Snapshot.
type Snapshot struct {
	*externglib.Object
}

// Snapshotter describes types inherited from class Snapshot.
// To get the original type, the caller must assert this to an interface or
// another type.
type Snapshotter interface {
	externglib.Objector

	// BaseSnapshot returns the underlying base class.
	BaseSnapshot() *Snapshot
}

var _ Snapshotter = (*Snapshot)(nil)

func wrapSnapshot(obj *externglib.Object) *Snapshot {
	return &Snapshot{
		Object: obj,
	}
}

func marshalSnapshotter(p uintptr) (interface{}, error) {
	return wrapSnapshot(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// BaseSnapshot returns v.
func (v *Snapshot) BaseSnapshot() *Snapshot {
	return v
}
