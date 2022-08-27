// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// ActionGroup retrieves the Group that was registered using prefix. The
// resulting Group may have been registered to widget or any Widget in its
// ancestry.
//
// If no action group was found matching prefix, then NULL is returned.
//
// The function takes the following parameters:
//
//    - prefix: “prefix” of the action group.
//
// The function returns the following values:
//
//    - actionGroup (optional) or NULL.
//
func (widget *Widget) ActionGroup(prefix string) *gio.ActionGroup {
	var _arg0 *C.GtkWidget    // out
	var _arg1 *C.gchar        // out
	var _cret *C.GActionGroup // in

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(prefix)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_widget_get_action_group(_arg0, _arg1)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(prefix)

	var _actionGroup *gio.ActionGroup // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_actionGroup = &gio.ActionGroup{
				Object: obj,
			}
		}
	}

	return _actionGroup
}

// ListActionPrefixes retrieves a NULL-terminated array of strings containing
// the prefixes of Group's available to widget.
//
// The function returns the following values:
//
//    - utf8s: NULL-terminated array of strings.
//
func (widget *Widget) ListActionPrefixes() []string {
	var _arg0 *C.GtkWidget // out
	var _cret **C.gchar    // in

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_cret = C.gtk_widget_list_action_prefixes(_arg0)
	runtime.KeepAlive(widget)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}
