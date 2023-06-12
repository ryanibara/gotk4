// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeEditableLabel = coreglib.Type(C.gtk_editable_label_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEditableLabel, F: marshalEditableLabel},
	})
}

// EditableLabelOverrides contains methods that are overridable.
type EditableLabelOverrides struct {
}

func defaultEditableLabelOverrides(v *EditableLabel) EditableLabelOverrides {
	return EditableLabelOverrides{}
}

// EditableLabel: GtkEditableLabel is a label that allows users to edit the text
// by switching to an “edit mode”.
//
// !An example GtkEditableLabel (editable-label.png)
//
// GtkEditableLabel does not have API of its own, but it implements the
// gtk.Editable interface.
//
// The default bindings for activating the edit mode is to click or press the
// Enter key. The default bindings for leaving the edit mode are the Enter key
// (to save the results) or the Escape key (to cancel the editing).
//
// CSS nodes
//
//    editablelabel[.editing]
//    ╰── stack
//        ├── label
//        ╰── text
//
// GtkEditableLabel has a main node with the name editablelabel. When the entry
// is in editing mode, it gets the .editing style class.
//
// For all the subnodes added to the text node in various situations, see
// gtk.Text.
type EditableLabel struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Editable
}

var (
	_ Widgetter         = (*EditableLabel)(nil)
	_ coreglib.Objector = (*EditableLabel)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*EditableLabel, *EditableLabelClass, EditableLabelOverrides](
		GTypeEditableLabel,
		initEditableLabelClass,
		wrapEditableLabel,
		defaultEditableLabelOverrides,
	)
}

func initEditableLabelClass(gclass unsafe.Pointer, overrides EditableLabelOverrides, classInitFunc func(*EditableLabelClass)) {
	if classInitFunc != nil {
		class := (*EditableLabelClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapEditableLabel(obj *coreglib.Object) *EditableLabel {
	return &EditableLabel{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		Editable: Editable{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
		},
	}
}

func marshalEditableLabel(p uintptr) (interface{}, error) {
	return wrapEditableLabel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewEditableLabel creates a new GtkEditableLabel widget.
//
// The function takes the following parameters:
//
//   - str: text for the label.
//
// The function returns the following values:
//
//   - editableLabel: new GtkEditableLabel.
//
func NewEditableLabel(str string) *EditableLabel {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_editable_label_new(_arg1)
	runtime.KeepAlive(str)

	var _editableLabel *EditableLabel // out

	_editableLabel = wrapEditableLabel(coreglib.Take(unsafe.Pointer(_cret)))

	return _editableLabel
}

// Editing returns whether the label is currently in “editing mode”.
//
// The function returns the following values:
//
//   - ok: TRUE if self is currently in editing mode.
//
func (self *EditableLabel) Editing() bool {
	var _arg0 *C.GtkEditableLabel // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkEditableLabel)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_editable_label_get_editing(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StartEditing switches the label into “editing mode”.
func (self *EditableLabel) StartEditing() {
	var _arg0 *C.GtkEditableLabel // out

	_arg0 = (*C.GtkEditableLabel)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.gtk_editable_label_start_editing(_arg0)
	runtime.KeepAlive(self)
}

// StopEditing switches the label out of “editing mode”.
//
// If commit is TRUE, the resulting text is kept as the gtk.Editable:text
// property value, otherwise the resulting text is discarded and the label will
// keep its previous gtk.Editable:text property value.
//
// The function takes the following parameters:
//
//   - commit: whether to set the edited text on the label.
//
func (self *EditableLabel) StopEditing(commit bool) {
	var _arg0 *C.GtkEditableLabel // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkEditableLabel)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if commit {
		_arg1 = C.TRUE
	}

	C.gtk_editable_label_stop_editing(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(commit)
}

// EditableLabelClass: instance of this type is always passed by reference.
type EditableLabelClass struct {
	*editableLabelClass
}

// editableLabelClass is the struct that's finalized.
type editableLabelClass struct {
	native *C.GtkEditableLabelClass
}

func (e *EditableLabelClass) ParentClass() *WidgetClass {
	valptr := &e.native.parent_class
	var _v *WidgetClass // out
	_v = (*WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
