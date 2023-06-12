// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GetOptionGroup returns a Group for the commandline arguments recognized by
// GTK+ and GDK.
//
// You should add this group to your Context with g_option_context_add_group(),
// if you are using g_option_context_parse() to parse your commandline
// arguments.
//
// The function takes the following parameters:
//
//   - openDefaultDisplay: whether to open the default display when parsing the
//     commandline arguments.
//
// The function returns the following values:
//
//   - optionGroup for the commandline arguments recognized by GTK+.
//
func GetOptionGroup(openDefaultDisplay bool) *glib.OptionGroup {
	var _arg1 C.gboolean      // out
	var _cret *C.GOptionGroup // in

	if openDefaultDisplay {
		_arg1 = C.TRUE
	}

	_cret = C.gtk_get_option_group(_arg1)
	runtime.KeepAlive(openDefaultDisplay)

	var _optionGroup *glib.OptionGroup // out

	_optionGroup = (*glib.OptionGroup)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_optionGroup)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_option_group_unref((*C.GOptionGroup)(intern.C))
		},
	)

	return _optionGroup
}
