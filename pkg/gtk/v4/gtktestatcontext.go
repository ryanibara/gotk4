// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

func TestAccessibleAssertionMessageRole(domain string, file string, line int, fn string, expr string, accessible Accessibler, expectedRole AccessibleRole, actualRole AccessibleRole) {
	var _arg1 *C.char             // out
	var _arg2 *C.char             // out
	var _arg3 C.int               // out
	var _arg4 *C.char             // out
	var _arg5 *C.char             // out
	var _arg6 *C.GtkAccessible    // out
	var _arg7 C.GtkAccessibleRole // out
	var _arg8 C.GtkAccessibleRole // out

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(domain)))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(file)))
	_arg3 = C.int(line)
	_arg4 = (*C.char)(unsafe.Pointer(C.CString(fn)))
	_arg5 = (*C.char)(unsafe.Pointer(C.CString(expr)))
	_arg6 = (*C.GtkAccessible)(unsafe.Pointer((accessible).(gextras.Nativer).Native()))
	_arg7 = C.GtkAccessibleRole(expectedRole)
	_arg8 = C.GtkAccessibleRole(actualRole)

	C.gtk_test_accessible_assertion_message_role(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
}

// TestAccessibleHasProperty checks whether the Accessible has property set.
func TestAccessibleHasProperty(accessible Accessibler, property AccessibleProperty) bool {
	var _arg1 *C.GtkAccessible        // out
	var _arg2 C.GtkAccessibleProperty // out
	var _cret C.gboolean              // in

	_arg1 = (*C.GtkAccessible)(unsafe.Pointer((accessible).(gextras.Nativer).Native()))
	_arg2 = C.GtkAccessibleProperty(property)

	_cret = C.gtk_test_accessible_has_property(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TestAccessibleHasRelation checks whether the Accessible has relation set.
func TestAccessibleHasRelation(accessible Accessibler, relation AccessibleRelation) bool {
	var _arg1 *C.GtkAccessible        // out
	var _arg2 C.GtkAccessibleRelation // out
	var _cret C.gboolean              // in

	_arg1 = (*C.GtkAccessible)(unsafe.Pointer((accessible).(gextras.Nativer).Native()))
	_arg2 = C.GtkAccessibleRelation(relation)

	_cret = C.gtk_test_accessible_has_relation(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TestAccessibleHasRole checks whether the Accessible:accessible-role of the
// accessible is role.
func TestAccessibleHasRole(accessible Accessibler, role AccessibleRole) bool {
	var _arg1 *C.GtkAccessible    // out
	var _arg2 C.GtkAccessibleRole // out
	var _cret C.gboolean          // in

	_arg1 = (*C.GtkAccessible)(unsafe.Pointer((accessible).(gextras.Nativer).Native()))
	_arg2 = C.GtkAccessibleRole(role)

	_cret = C.gtk_test_accessible_has_role(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TestAccessibleHasState checks whether the Accessible has state set.
func TestAccessibleHasState(accessible Accessibler, state AccessibleState) bool {
	var _arg1 *C.GtkAccessible     // out
	var _arg2 C.GtkAccessibleState // out
	var _cret C.gboolean           // in

	_arg1 = (*C.GtkAccessible)(unsafe.Pointer((accessible).(gextras.Nativer).Native()))
	_arg2 = C.GtkAccessibleState(state)

	_cret = C.gtk_test_accessible_has_state(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
