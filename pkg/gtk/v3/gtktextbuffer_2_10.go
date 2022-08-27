// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// CopyTargetList: this function returns the list of targets this text buffer
// can provide for copying and as DND source. The targets in the list are added
// with info values from the TextBufferTargetInfo enum, using
// gtk_target_list_add_rich_text_targets() and
// gtk_target_list_add_text_targets().
//
// The function returns the following values:
//
//    - targetList: TargetList.
//
func (buffer *TextBuffer) CopyTargetList() *TargetList {
	var _arg0 *C.GtkTextBuffer // out
	var _cret *C.GtkTargetList // in

	_arg0 = (*C.GtkTextBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_cret = C.gtk_text_buffer_get_copy_target_list(_arg0)
	runtime.KeepAlive(buffer)

	var _targetList *TargetList // out

	_targetList = (*TargetList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gtk_target_list_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_targetList)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_target_list_unref((*C.GtkTargetList)(intern.C))
		},
	)

	return _targetList
}

// HasSelection indicates whether the buffer has some text currently selected.
//
// The function returns the following values:
//
//    - ok: TRUE if the there is text selected.
//
func (buffer *TextBuffer) HasSelection() bool {
	var _arg0 *C.GtkTextBuffer // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkTextBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_cret = C.gtk_text_buffer_get_has_selection(_arg0)
	runtime.KeepAlive(buffer)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PasteTargetList: this function returns the list of targets this text buffer
// supports for pasting and as DND destination. The targets in the list are
// added with info values from the TextBufferTargetInfo enum, using
// gtk_target_list_add_rich_text_targets() and
// gtk_target_list_add_text_targets().
//
// The function returns the following values:
//
//    - targetList: TargetList.
//
func (buffer *TextBuffer) PasteTargetList() *TargetList {
	var _arg0 *C.GtkTextBuffer // out
	var _cret *C.GtkTargetList // in

	_arg0 = (*C.GtkTextBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_cret = C.gtk_text_buffer_get_paste_target_list(_arg0)
	runtime.KeepAlive(buffer)

	var _targetList *TargetList // out

	_targetList = (*TargetList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gtk_target_list_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_targetList)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_target_list_unref((*C.GtkTargetList)(intern.C))
		},
	)

	return _targetList
}
