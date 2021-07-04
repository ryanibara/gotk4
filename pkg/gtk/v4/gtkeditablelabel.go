// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_editable_label_get_type()), F: marshalEditableLabel},
	})
}

// EditableLabel: a `GtkEditableLabel` is a label that allows users to edit the
// text by switching to an “edit mode”.
//
// !An example GtkEditableLabel (editable-label.png)
//
// `GtkEditableLabel` does not have API of its own, but it implements the
// [iface@Gtk.Editable] interface.
//
// The default bindings for activating the edit mode is to click or press the
// Enter key. The default bindings for leaving the edit mode are the Enter key
// (to save the results) or the Escape key (to cancel the editing).
//
//
// CSS nodes
//
// “` editablelabel[.editing] ╰── stack ├── label ╰── text “`
//
// `GtkEditableLabel` has a main node with the name editablelabel. When the
// entry is in editing mode, it gets the .editing style class.
//
// For all the subnodes added to the text node in various situations, see
// [class@Gtk.Text].
type EditableLabel interface {
	Editable

	Editing() bool

	StartEditingEditableLabel()

	StopEditingEditableLabel(commit bool)
}

// editableLabel implements the EditableLabel class.
type editableLabel struct {
	Widget
}

// WrapEditableLabel wraps a GObject to the right type. It is
// primarily used internally.
func WrapEditableLabel(obj *externglib.Object) EditableLabel {
	return editableLabel{
		Widget: WrapWidget(obj),
	}
}

func marshalEditableLabel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEditableLabel(obj), nil
}

func NewEditableLabel(str string) EditableLabel {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_editable_label_new(_arg1)

	var _editableLabel EditableLabel // out

	_editableLabel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(EditableLabel)

	return _editableLabel
}

func (s editableLabel) Editing() bool {
	var _arg0 *C.GtkEditableLabel // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkEditableLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_editable_label_get_editing(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s editableLabel) StartEditingEditableLabel() {
	var _arg0 *C.GtkEditableLabel // out

	_arg0 = (*C.GtkEditableLabel)(unsafe.Pointer(s.Native()))

	C.gtk_editable_label_start_editing(_arg0)
}

func (s editableLabel) StopEditingEditableLabel(commit bool) {
	var _arg0 *C.GtkEditableLabel // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkEditableLabel)(unsafe.Pointer(s.Native()))
	if commit {
		_arg1 = C.TRUE
	}

	C.gtk_editable_label_stop_editing(_arg0, _arg1)
}

func (e editableLabel) DeleteSelection() {
	WrapEditable(gextras.InternObject(e)).DeleteSelection()
}

func (e editableLabel) DeleteText(startPos int, endPos int) {
	WrapEditable(gextras.InternObject(e)).DeleteText(startPos, endPos)
}

func (e editableLabel) FinishDelegate() {
	WrapEditable(gextras.InternObject(e)).FinishDelegate()
}

func (e editableLabel) Alignment() float32 {
	return WrapEditable(gextras.InternObject(e)).Alignment()
}

func (e editableLabel) Chars(startPos int, endPos int) string {
	return WrapEditable(gextras.InternObject(e)).Chars(startPos, endPos)
}

func (e editableLabel) Delegate() Editable {
	return WrapEditable(gextras.InternObject(e)).Delegate()
}

func (e editableLabel) Editable() bool {
	return WrapEditable(gextras.InternObject(e)).Editable()
}

func (e editableLabel) EnableUndo() bool {
	return WrapEditable(gextras.InternObject(e)).EnableUndo()
}

func (e editableLabel) MaxWidthChars() int {
	return WrapEditable(gextras.InternObject(e)).MaxWidthChars()
}

func (e editableLabel) Position() int {
	return WrapEditable(gextras.InternObject(e)).Position()
}

func (e editableLabel) SelectionBounds() (startPos int, endPos int, ok bool) {
	return WrapEditable(gextras.InternObject(e)).SelectionBounds()
}

func (e editableLabel) Text() string {
	return WrapEditable(gextras.InternObject(e)).Text()
}

func (e editableLabel) WidthChars() int {
	return WrapEditable(gextras.InternObject(e)).WidthChars()
}

func (e editableLabel) InitDelegate() {
	WrapEditable(gextras.InternObject(e)).InitDelegate()
}

func (e editableLabel) SelectRegion(startPos int, endPos int) {
	WrapEditable(gextras.InternObject(e)).SelectRegion(startPos, endPos)
}

func (e editableLabel) SetAlignment(xalign float32) {
	WrapEditable(gextras.InternObject(e)).SetAlignment(xalign)
}

func (e editableLabel) SetEditable(isEditable bool) {
	WrapEditable(gextras.InternObject(e)).SetEditable(isEditable)
}

func (e editableLabel) SetEnableUndo(enableUndo bool) {
	WrapEditable(gextras.InternObject(e)).SetEnableUndo(enableUndo)
}

func (e editableLabel) SetMaxWidthChars(nChars int) {
	WrapEditable(gextras.InternObject(e)).SetMaxWidthChars(nChars)
}

func (e editableLabel) SetPosition(position int) {
	WrapEditable(gextras.InternObject(e)).SetPosition(position)
}

func (e editableLabel) SetText(text string) {
	WrapEditable(gextras.InternObject(e)).SetText(text)
}

func (e editableLabel) SetWidthChars(nChars int) {
	WrapEditable(gextras.InternObject(e)).SetWidthChars(nChars)
}

func (s editableLabel) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s editableLabel) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s editableLabel) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s editableLabel) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s editableLabel) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s editableLabel) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s editableLabel) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b editableLabel) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
