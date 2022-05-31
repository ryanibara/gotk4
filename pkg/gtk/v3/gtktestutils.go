// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// TestCreateSimpleWindow: create a simple window with window title window_title
// and text contents dialog_text. The window will quit any running
// gtk_main()-loop when destroyed, and it will automatically be destroyed upon
// test function teardown.
//
// Deprecated: This testing infrastructure is phased out in favor of reftests.
//
// The function takes the following parameters:
//
//    - windowTitle: title of the window to be displayed.
//    - dialogText: text inside the window to be displayed.
//
// The function returns the following values:
//
//    - widget pointer to the newly created GtkWindow.
//
func TestCreateSimpleWindow(windowTitle, dialogText string) Widgetter {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(windowTitle)))
	defer C.free(unsafe.Pointer(_arg0))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(dialogText)))
	defer C.free(unsafe.Pointer(_arg1))
	*(*string)(unsafe.Pointer(&args[0])) = _arg0
	*(*string)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "test_create_simple_window").Invoke(args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(windowTitle)
	runtime.KeepAlive(dialogText)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// TestFindLabel: this function will search widget and all its descendants for a
// GtkLabel widget with a text string matching label_pattern. The label_pattern
// may contain asterisks “*” and question marks “?” as placeholders,
// g_pattern_match() is used for the matching. Note that locales other than "C“
// tend to alter (translate” label strings, so this function is genrally only
// useful in test programs with predetermined locales, see gtk_test_init() for
// more details.
//
// The function takes the following parameters:
//
//    - widget: valid label or container widget.
//    - labelPattern: shell-glob pattern to match a label string.
//
// The function returns the following values:
//
//    - ret: gtkLabel widget if any is found.
//
func TestFindLabel(widget Widgetter, labelPattern string) Widgetter {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(labelPattern)))
	defer C.free(unsafe.Pointer(_arg1))
	*(*Widgetter)(unsafe.Pointer(&args[0])) = _arg0
	*(*string)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "test_find_label").Invoke(args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(widget)
	runtime.KeepAlive(labelPattern)

	var _ret Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_ret = rv
	}

	return _ret
}

// TestRegisterAllTypes: force registration of all core Gtk+ and Gdk object
// types. This allowes to refer to any of those object types via
// g_type_from_name() after calling this function.
func TestRegisterAllTypes() {
	girepository.MustFind("Gtk", "test_register_all_types").Invoke(nil, nil)
}

// TestSliderGetValue: retrive the literal adjustment value for GtkRange based
// widgets and spin buttons. Note that the value returned by this function is
// anything between the lower and upper bounds of the adjustment belonging to
// widget, and is not a percentage as passed in to gtk_test_slider_set_perc().
//
// Deprecated: This testing infrastructure is phased out in favor of reftests.
//
// The function takes the following parameters:
//
//    - widget: valid widget pointer.
//
// The function returns the following values:
//
//    - gdouble: gtk_adjustment_get_value (adjustment) for an adjustment
//      belonging to widget.
//
func TestSliderGetValue(widget Widgetter) float64 {
	var args [1]girepository.Argument
	var _arg0 *C.void  // out
	var _cret C.double // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	*(*Widgetter)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "test_slider_get_value").Invoke(args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(widget)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// TestSliderSetPerc: this function will adjust the slider position of all
// GtkRange based widgets, such as scrollbars or scales, it’ll also adjust spin
// buttons. The adjustment value of these widgets is set to a value between the
// lower and upper limits, according to the percentage argument.
//
// Deprecated: This testing infrastructure is phased out in favor of reftests.
//
// The function takes the following parameters:
//
//    - widget: valid widget pointer.
//    - percentage: value between 0 and 100.
//
func TestSliderSetPerc(widget Widgetter, percentage float64) {
	var args [2]girepository.Argument
	var _arg0 *C.void  // out
	var _arg1 C.double // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg1 = C.double(percentage)
	*(*Widgetter)(unsafe.Pointer(&args[0])) = _arg0
	*(*float64)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "test_slider_set_perc").Invoke(args[:], nil)

	runtime.KeepAlive(widget)
	runtime.KeepAlive(percentage)
}

// TestSpinButtonClick: this function will generate a button click in the
// upwards or downwards spin button arrow areas, usually leading to an increase
// or decrease of spin button’s value.
//
// Deprecated: This testing infrastructure is phased out in favor of reftests.
//
// The function takes the following parameters:
//
//    - spinner: valid GtkSpinButton widget.
//    - button: number of the pointer button for the event, usually 1, 2 or 3.
//    - upwards: TRUE for upwards arrow click, FALSE for downwards arrow click.
//
// The function returns the following values:
//
//    - ok: whether all actions neccessary for the button click simulation were
//      carried out successfully.
//
func TestSpinButtonClick(spinner *SpinButton, button uint32, upwards bool) bool {
	var args [3]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.guint    // out
	var _arg2 C.gboolean // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinner).Native()))
	_arg1 = C.guint(button)
	if upwards {
		_arg2 = C.TRUE
	}
	*(**SpinButton)(unsafe.Pointer(&args[0])) = _arg0
	*(*uint32)(unsafe.Pointer(&args[1])) = _arg1
	*(*bool)(unsafe.Pointer(&args[2])) = _arg2

	_gret := girepository.MustFind("Gtk", "test_spin_button_click").Invoke(args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(spinner)
	runtime.KeepAlive(button)
	runtime.KeepAlive(upwards)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TestTextGet: retrive the text string of widget if it is a GtkLabel,
// GtkEditable (entry and text widgets) or GtkTextView.
//
// Deprecated: This testing infrastructure is phased out in favor of reftests.
//
// The function takes the following parameters:
//
//    - widget: valid widget pointer.
//
// The function returns the following values:
//
//    - utf8: new 0-terminated C string, needs to be released with g_free().
//
func TestTextGet(widget Widgetter) string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	*(*Widgetter)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "test_text_get").Invoke(args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(widget)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// TestTextSet: set the text string of widget to string if it is a GtkLabel,
// GtkEditable (entry and text widgets) or GtkTextView.
//
// Deprecated: This testing infrastructure is phased out in favor of reftests.
//
// The function takes the following parameters:
//
//    - widget: valid widget pointer.
//    - str: 0-terminated C string.
//
func TestTextSet(widget Widgetter, str string) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))
	*(*Widgetter)(unsafe.Pointer(&args[0])) = _arg0
	*(*string)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "test_text_set").Invoke(args[:], nil)

	runtime.KeepAlive(widget)
	runtime.KeepAlive(str)
}

// TestWidgetWaitForDraw enters the main loop and waits for widget to be
// “drawn”. In this context that means it waits for the frame clock of widget to
// have run a full styling, layout and drawing cycle.
//
// This function is intended to be used for syncing with actions that depend on
// widget relayouting or on interaction with the display server.
//
// The function takes the following parameters:
//
//    - widget to wait for.
//
func TestWidgetWaitForDraw(widget Widgetter) {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	*(*Widgetter)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "test_widget_wait_for_draw").Invoke(args[:], nil)

	runtime.KeepAlive(widget)
}
