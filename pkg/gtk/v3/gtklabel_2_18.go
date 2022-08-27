// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// CurrentURI returns the URI for the currently active link in the label. The
// active link is the one under the mouse pointer or, in a selectable label, the
// link in which the text cursor is currently positioned.
//
// This function is intended for use in a Label::activate-link handler or for
// use in a Widget::query-tooltip handler.
//
// The function returns the following values:
//
//    - utf8: currently active URI. The string is owned by GTK+ and must not be
//      freed or modified.
//
func (label *Label) CurrentURI() string {
	var _arg0 *C.GtkLabel // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(coreglib.InternObject(label).Native()))

	_cret = C.gtk_label_get_current_uri(_arg0)
	runtime.KeepAlive(label)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// TrackVisitedLinks returns whether the label is currently keeping track of
// clicked links.
//
// The function returns the following values:
//
//    - ok: TRUE if clicked links are remembered.
//
func (label *Label) TrackVisitedLinks() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(coreglib.InternObject(label).Native()))

	_cret = C.gtk_label_get_track_visited_links(_arg0)
	runtime.KeepAlive(label)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetTrackVisitedLinks sets whether the label should keep track of clicked
// links (and use a different color for them).
//
// The function takes the following parameters:
//
//    - trackLinks: TRUE to track visited links.
//
func (label *Label) SetTrackVisitedLinks(trackLinks bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(coreglib.InternObject(label).Native()))
	if trackLinks {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_track_visited_links(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(trackLinks)
}
