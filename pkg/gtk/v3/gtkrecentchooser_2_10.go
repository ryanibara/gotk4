// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeRecentChooserError = coreglib.Type(C.gtk_recent_chooser_error_get_type())
	GTypeRecentSortType     = coreglib.Type(C.gtk_recent_sort_type_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRecentChooserError, F: marshalRecentChooserError},
		coreglib.TypeMarshaler{T: GTypeRecentSortType, F: marshalRecentSortType},
	})
}

// RecentChooserError: these identify the various errors that can occur while
// calling RecentChooser functions.
type RecentChooserError C.gint

const (
	// RecentChooserErrorNotFound indicates that a file does not exist.
	RecentChooserErrorNotFound RecentChooserError = iota
	// RecentChooserErrorInvalidURI indicates a malformed URI.
	RecentChooserErrorInvalidURI
)

func marshalRecentChooserError(p uintptr) (interface{}, error) {
	return RecentChooserError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for RecentChooserError.
func (r RecentChooserError) String() string {
	switch r {
	case RecentChooserErrorNotFound:
		return "NotFound"
	case RecentChooserErrorInvalidURI:
		return "InvalidURI"
	default:
		return fmt.Sprintf("RecentChooserError(%d)", r)
	}
}

// RecentSortType: used to specify the sorting method to be applyed to the
// recently used resource list.
type RecentSortType C.gint

const (
	// RecentSortNone: do not sort the returned list of recently used resources.
	RecentSortNone RecentSortType = iota
	// RecentSortMru: sort the returned list with the most recently used items
	// first.
	RecentSortMru
	// RecentSortLru: sort the returned list with the least recently used items
	// first.
	RecentSortLru
	// RecentSortCustom: sort the returned list using a custom sorting function
	// passed using gtk_recent_chooser_set_sort_func().
	RecentSortCustom
)

func marshalRecentSortType(p uintptr) (interface{}, error) {
	return RecentSortType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for RecentSortType.
func (r RecentSortType) String() string {
	switch r {
	case RecentSortNone:
		return "None"
	case RecentSortMru:
		return "Mru"
	case RecentSortLru:
		return "Lru"
	case RecentSortCustom:
		return "Custom"
	default:
		return fmt.Sprintf("RecentSortType(%d)", r)
	}
}
