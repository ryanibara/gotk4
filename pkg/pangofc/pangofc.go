// Code generated by girgen. DO NOT EDIT.

package pangofc

import (
	"github.com/diamondburned/gotk4/internal/callback"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pangofc-fontmap.h>
// extern void callbackDelete(gpointer);
import "C"

//export callbackDelete
func callbackDelete(ptr C.gpointer) { callback.Delete(ptr) }
