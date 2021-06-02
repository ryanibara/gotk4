// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gobject/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_editable_get_type()), F: marshalEditable},
	})
}

// EditableDelegateGetProperty gets a property of the `GtkEditable` delegate for
// @object.
//
// This is helper function that should be called in the `get_property` function
// of your `GtkEditable` implementation, before handling your own properties.
func EditableDelegateGetProperty(object gextras.Objector, propID uint, value *externglib.Value, pspec gobject.ParamSpec) bool {
	var arg1 *C.GObject
	var arg2 C.guint
	var arg3 *C.GValue
	var arg4 *C.GParamSpec

	arg1 = (*C.GObject)(object.Native())
	arg2 = C.guint(propID)
	arg3 = (*C.GValue)(value.GValue)
	arg4 = (*C.GParamSpec)(pspec.Native())

	ret := C.gtk_editable_delegate_get_property(arg1, arg2, arg3, arg4)

	var ret0 bool

	ret0 = C.bool(ret) != C.false

	return ret0
}

// EditableDelegateSetProperty sets a property on the `GtkEditable` delegate for
// @object.
//
// This is a helper function that should be called in the `set_property`
// function of your `GtkEditable` implementation, before handling your own
// properties.
func EditableDelegateSetProperty(object gextras.Objector, propID uint, value *externglib.Value, pspec gobject.ParamSpec) bool {
	var arg1 *C.GObject
	var arg2 C.guint
	var arg3 *C.GValue
	var arg4 *C.GParamSpec

	arg1 = (*C.GObject)(object.Native())
	arg2 = C.guint(propID)
	arg3 = (*C.GValue)(value.GValue)
	arg4 = (*C.GParamSpec)(pspec.Native())

	ret := C.gtk_editable_delegate_set_property(arg1, arg2, arg3, arg4)

	var ret0 bool

	ret0 = C.bool(ret) != C.false

	return ret0
}

// EditableOverrider contains methods that are overridable. This
// interface is a subset of the interface Editable.
type EditableOverrider interface {
	Changed()
	// DeleteText deletes a sequence of characters.
	//
	// The characters that are deleted are those characters at positions from
	// @start_pos up to, but not including @end_pos. If @end_pos is negative,
	// then the characters deleted are those from @start_pos to the end of the
	// text.
	//
	// Note that the positions are specified in characters, not bytes.
	DeleteText(startPos int, endPos int)
	// DoDeleteText deletes a sequence of characters.
	//
	// The characters that are deleted are those characters at positions from
	// @start_pos up to, but not including @end_pos. If @end_pos is negative,
	// then the characters deleted are those from @start_pos to the end of the
	// text.
	//
	// Note that the positions are specified in characters, not bytes.
	DoDeleteText(startPos int, endPos int)
	// DoInsertText inserts @length bytes of @text into the contents of the
	// widget, at position @position.
	//
	// Note that the position is in characters, not in bytes. The function
	// updates @position to point after the newly inserted text.
	DoInsertText(text string, length int, position int)
	// Delegate gets the `GtkEditable` that @editable is delegating its
	// implementation to.
	//
	// Typically, the delegate is a [class@Gtk.Text] widget.
	Delegate() Editable
	// SelectionBounds retrieves the selection bound of the editable.
	//
	// @start_pos will be filled with the start of the selection and @end_pos
	// with end. If no text was selected both will be identical and false will
	// be returned.
	//
	// Note that positions are specified in characters, not bytes.
	SelectionBounds() (startPos int, endPos int, ok bool)
	// Text retrieves the contents of @editable.
	//
	// The returned string is owned by GTK and must not be modified or freed.
	Text() string
	// InsertText inserts @length bytes of @text into the contents of the
	// widget, at position @position.
	//
	// Note that the position is in characters, not in bytes. The function
	// updates @position to point after the newly inserted text.
	InsertText(text string, length int, position int)
	// SetSelectionBounds selects a region of text.
	//
	// The characters that are selected are those characters at positions from
	// @start_pos up to, but not including @end_pos. If @end_pos is negative,
	// then the characters selected are those characters from @start_pos to the
	// end of the text.
	//
	// Note that positions are specified in characters, not bytes.
	SetSelectionBounds(startPos int, endPos int)
}

// Editable: `GtkEditable` is an interface for text editing widgets.
//
// Typical examples of editable widgets are [class@Gtk.Entry] and
// [class@Gtk.SpinButton]. It contains functions for generically manipulating an
// editable widget, a large number of action signals used for key bindings, and
// several signals that an application can connect to modify the behavior of a
// widget.
//
// As an example of the latter usage, by connecting the following handler to
// [signal@Gtk.Editable::insert-text], an application can convert all entry into
// a widget into uppercase.
//
// Forcing entry to uppercase.
//
// “`c #include <ctype.h>
//
// void insert_text_handler (GtkEditable *editable, const char *text, int
// length, int *position, gpointer data) { char *result = g_utf8_strup (text,
// length);
//
//    g_signal_handlers_block_by_func (editable,
//                                 (gpointer) insert_text_handler, data);
//    gtk_editable_insert_text (editable, result, length, position);
//    g_signal_handlers_unblock_by_func (editable,
//                                       (gpointer) insert_text_handler, data);
//
//    g_signal_stop_emission_by_name (editable, "insert_text");
//
//    g_free (result);
//
// } “`
//
//
// Implementing GtkEditable
//
// The most likely scenario for implementing `GtkEditable` on your own widget is
// that you will embed a Text inside a complex widget, and want to delegate the
// editable functionality to that text widget. `GtkEditable` provides some
// utility functions to make this easy.
//
// In your class_init function, call [func@Gtk.Editable.install_properties],
// passing the first available property ID:
//
// “`c static void my_class_init (MyClass *class) { ...
// g_object_class_install_properties (object_class, NUM_PROPERTIES, props);
// gtk_editable_install_properties (object_clas, NUM_PROPERTIES); ... } “`
//
// In your interface_init function for the `GtkEditable` interface, provide an
// implementation for the get_delegate vfunc that returns your text widget:
//
// “`c GtkEditable * get_editable_delegate (GtkEditable *editable) { return
// GTK_EDITABLE (MY_WIDGET (editable)->text_widget); }
//
// static void my_editable_init (GtkEditableInterface *iface) {
// iface->get_delegate = get_editable_delegate; } “`
//
// You don't need to provide any other vfuncs. The default implementations work
// by forwarding to the delegate that the GtkEditableInterface.get_delegate()
// vfunc returns.
//
// In your instance_init function, create your text widget, and then call
// [method@Gtk.Editable.init_delegate]:
//
// “`c static void my_widget_init (MyWidget *self) { ... self->text_widget =
// gtk_text_new (); gtk_editable_init_delegate (GTK_EDITABLE (self)); ... } “`
//
// In your dispose function, call [method@Gtk.Editable.finish_delegate] before
// destroying your text widget:
//
// “`c static void my_widget_dispose (GObject *object) { ...
// gtk_editable_finish_delegate (GTK_EDITABLE (self)); g_clear_pointer
// (&self->text_widget, gtk_widget_unparent); ... } “`
//
// Finally, use [func@Gtk.Editable.delegate_set_property] in your `set_property`
// function (and similar for `get_property`), to set the editable properties:
//
// “`c ... if (gtk_editable_delegate_set_property (object, prop_id, value,
// pspec)) return;
//
//    switch (prop_id)
//    ...
//
// “`
//
// It is important to note that if you create a `GtkEditable` that uses a
// delegate, the low level [signal@Gtk.Editable::insert-text] and
// [signal@Gtk.Editable::delete-text] signals will be propagated from the
// "wrapper" editable to the delegate, but they will not be propagated from the
// delegate to the "wrapper" editable, as they would cause an infinite
// recursion. If you wish to connect to the [signal@Gtk.Editable::insert-text]
// and [signal@Gtk.Editable::delete-text] signals, you will need to connect to
// them on the delegate obtained via [method@Gtk.Editable.get_delegate].
type Editable interface {
	Widget
	EditableOverrider

	// DeleteSelection deletes the currently selected text of the editable.
	//
	// This call doesn’t do anything if there is no selected text.
	DeleteSelection()
	// FinishDelegate undoes the setup done by
	// [method@Gtk.Editable.init_delegate].
	//
	// This is a helper function that should be called from dispose, before
	// removing the delegate object.
	FinishDelegate()
	// Alignment gets the alignment of the editable.
	Alignment() float32
	// Chars retrieves a sequence of characters.
	//
	// The characters that are retrieved are those characters at positions from
	// @start_pos up to, but not including @end_pos. If @end_pos is negative,
	// then the characters retrieved are those characters from @start_pos to the
	// end of the text.
	//
	// Note that positions are specified in characters, not bytes.
	Chars(startPos int, endPos int) string
	// Editable retrieves whether @editable is editable.
	Editable() bool
	// EnableUndo gets if undo/redo actions are enabled for @editable
	EnableUndo() bool
	// MaxWidthChars retrieves the desired maximum width of @editable, in
	// characters.
	MaxWidthChars() int
	// Position retrieves the current position of the cursor relative to the
	// start of the content of the editable.
	//
	// Note that this position is in characters, not in bytes.
	Position() int
	// WidthChars gets the number of characters of space reserved for the
	// contents of the editable.
	WidthChars() int
	// InitDelegate sets up a delegate for `GtkEditable`.
	//
	// This is assuming that the get_delegate vfunc in the `GtkEditable`
	// interface has been set up for the @editable's type.
	//
	// This is a helper function that should be called in instance init, after
	// creating the delegate object.
	InitDelegate()
	// SelectRegion selects a region of text.
	//
	// The characters that are selected are those characters at positions from
	// @start_pos up to, but not including @end_pos. If @end_pos is negative,
	// then the characters selected are those characters from @start_pos to the
	// end of the text.
	//
	// Note that positions are specified in characters, not bytes.
	SelectRegion(startPos int, endPos int)
	// SetAlignment sets the alignment for the contents of the editable.
	//
	// This controls the horizontal positioning of the contents when the
	// displayed text is shorter than the width of the editable.
	SetAlignment(xalign float32)
	// SetEditable determines if the user can edit the text in the editable
	// widget.
	SetEditable(isEditable bool)
	// SetEnableUndo: if enabled, changes to @editable will be saved for
	// undo/redo actions.
	//
	// This results in an additional copy of text changes and are not stored in
	// secure memory. As such, undo is forcefully disabled when
	// [property@Gtk.Text:visibility] is set to false.
	SetEnableUndo(enableUndo bool)
	// SetMaxWidthChars sets the desired maximum width in characters of
	// @editable.
	SetMaxWidthChars(nChars int)
	// SetPosition sets the cursor position in the editable to the given value.
	//
	// The cursor is displayed before the character with the given (base 0)
	// index in the contents of the editable. The value must be less than or
	// equal to the number of characters in the editable. A value of -1
	// indicates that the position should be set after the last character of the
	// editable. Note that @position is in characters, not in bytes.
	SetPosition(position int)
	// SetText sets the text in the editable to the given value.
	//
	// This is replacing the current contents.
	SetText(text string)
	// SetWidthChars changes the size request of the editable to be about the
	// right size for @n_chars characters.
	//
	// Note that it changes the size request, the size can still be affected by
	// how you pack the widget into containers. If @n_chars is -1, the size
	// reverts to the default size.
	SetWidthChars(nChars int)
}

// editable implements the Editable interface.
type editable struct {
	Widget
}

var _ Editable = (*editable)(nil)

// WrapEditable wraps a GObject to a type that implements interface
// Editable. It is primarily used internally.
func WrapEditable(obj *externglib.Object) Editable {
	return Editable{
		Widget: WrapWidget(obj),
	}
}

func marshalEditable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEditable(obj), nil
}

// DeleteSelection deletes the currently selected text of the editable.
//
// This call doesn’t do anything if there is no selected text.
func (e editable) DeleteSelection() {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(e.Native())

	C.gtk_editable_delete_selection(arg0)
}

// DeleteText deletes a sequence of characters.
//
// The characters that are deleted are those characters at positions from
// @start_pos up to, but not including @end_pos. If @end_pos is negative,
// then the characters deleted are those from @start_pos to the end of the
// text.
//
// Note that the positions are specified in characters, not bytes.
func (e editable) DeleteText(startPos int, endPos int) {
	var arg0 *C.GtkEditable
	var arg1 C.int
	var arg2 C.int

	arg0 = (*C.GtkEditable)(e.Native())
	arg1 = C.int(startPos)
	arg2 = C.int(endPos)

	C.gtk_editable_delete_text(arg0, arg1, arg2)
}

// FinishDelegate undoes the setup done by
// [method@Gtk.Editable.init_delegate].
//
// This is a helper function that should be called from dispose, before
// removing the delegate object.
func (e editable) FinishDelegate() {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(e.Native())

	C.gtk_editable_finish_delegate(arg0)
}

// Alignment gets the alignment of the editable.
func (e editable) Alignment() float32 {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(e.Native())

	ret := C.gtk_editable_get_alignment(arg0)

	var ret0 float32

	ret0 = float32(ret)

	return ret0
}

// Chars retrieves a sequence of characters.
//
// The characters that are retrieved are those characters at positions from
// @start_pos up to, but not including @end_pos. If @end_pos is negative,
// then the characters retrieved are those characters from @start_pos to the
// end of the text.
//
// Note that positions are specified in characters, not bytes.
func (e editable) Chars(startPos int, endPos int) string {
	var arg0 *C.GtkEditable
	var arg1 C.int
	var arg2 C.int

	arg0 = (*C.GtkEditable)(e.Native())
	arg1 = C.int(startPos)
	arg2 = C.int(endPos)

	ret := C.gtk_editable_get_chars(arg0, arg1, arg2)

	var ret0 string

	ret0 = C.GoString(ret)
	C.free(unsafe.Pointer(ret))

	return ret0
}

// Delegate gets the `GtkEditable` that @editable is delegating its
// implementation to.
//
// Typically, the delegate is a [class@Gtk.Text] widget.
func (e editable) Delegate() Editable {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(e.Native())

	ret := C.gtk_editable_get_delegate(arg0)

	var ret0 Editable

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Editable)

	return ret0
}

// Editable retrieves whether @editable is editable.
func (e editable) Editable() bool {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(e.Native())

	ret := C.gtk_editable_get_editable(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != C.false

	return ret0
}

// EnableUndo gets if undo/redo actions are enabled for @editable
func (e editable) EnableUndo() bool {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(e.Native())

	ret := C.gtk_editable_get_enable_undo(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != C.false

	return ret0
}

// MaxWidthChars retrieves the desired maximum width of @editable, in
// characters.
func (e editable) MaxWidthChars() int {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(e.Native())

	ret := C.gtk_editable_get_max_width_chars(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// Position retrieves the current position of the cursor relative to the
// start of the content of the editable.
//
// Note that this position is in characters, not in bytes.
func (e editable) Position() int {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(e.Native())

	ret := C.gtk_editable_get_position(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// SelectionBounds retrieves the selection bound of the editable.
//
// @start_pos will be filled with the start of the selection and @end_pos
// with end. If no text was selected both will be identical and false will
// be returned.
//
// Note that positions are specified in characters, not bytes.
func (e editable) SelectionBounds() (startPos int, endPos int, ok bool) {
	var arg0 *C.GtkEditable
	var arg1 *C.int // out
	var arg2 *C.int // out

	arg0 = (*C.GtkEditable)(e.Native())

	ret := C.gtk_editable_get_selection_bounds(arg0, &arg1, &arg2)

	var ret0 int
	var ret1 int
	var ret2 bool

	ret0 = int(arg1)

	ret1 = int(arg2)

	ret2 = C.bool(ret) != C.false

	return ret0, ret1, ret2
}

// Text retrieves the contents of @editable.
//
// The returned string is owned by GTK and must not be modified or freed.
func (e editable) Text() string {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(e.Native())

	ret := C.gtk_editable_get_text(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// WidthChars gets the number of characters of space reserved for the
// contents of the editable.
func (e editable) WidthChars() int {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(e.Native())

	ret := C.gtk_editable_get_width_chars(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// InitDelegate sets up a delegate for `GtkEditable`.
//
// This is assuming that the get_delegate vfunc in the `GtkEditable`
// interface has been set up for the @editable's type.
//
// This is a helper function that should be called in instance init, after
// creating the delegate object.
func (e editable) InitDelegate() {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(e.Native())

	C.gtk_editable_init_delegate(arg0)
}

// InsertText inserts @length bytes of @text into the contents of the
// widget, at position @position.
//
// Note that the position is in characters, not in bytes. The function
// updates @position to point after the newly inserted text.
func (e editable) InsertText(text string, length int, position int) {
	var arg0 *C.GtkEditable
	var arg1 *C.char
	var arg2 C.int
	var arg3 *C.int

	arg0 = (*C.GtkEditable)(e.Native())
	arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.int(length)
	arg3 = (*C.int)(position)

	C.gtk_editable_insert_text(arg0, arg1, arg2, arg3)
}

// SelectRegion selects a region of text.
//
// The characters that are selected are those characters at positions from
// @start_pos up to, but not including @end_pos. If @end_pos is negative,
// then the characters selected are those characters from @start_pos to the
// end of the text.
//
// Note that positions are specified in characters, not bytes.
func (e editable) SelectRegion(startPos int, endPos int) {
	var arg0 *C.GtkEditable
	var arg1 C.int
	var arg2 C.int

	arg0 = (*C.GtkEditable)(e.Native())
	arg1 = C.int(startPos)
	arg2 = C.int(endPos)

	C.gtk_editable_select_region(arg0, arg1, arg2)
}

// SetAlignment sets the alignment for the contents of the editable.
//
// This controls the horizontal positioning of the contents when the
// displayed text is shorter than the width of the editable.
func (e editable) SetAlignment(xalign float32) {
	var arg0 *C.GtkEditable
	var arg1 C.float

	arg0 = (*C.GtkEditable)(e.Native())
	arg1 = C.float(xalign)

	C.gtk_editable_set_alignment(arg0, arg1)
}

// SetEditable determines if the user can edit the text in the editable
// widget.
func (e editable) SetEditable(isEditable bool) {
	var arg0 *C.GtkEditable
	var arg1 C.gboolean

	arg0 = (*C.GtkEditable)(e.Native())
	if isEditable {
		arg1 = C.TRUE
	}

	C.gtk_editable_set_editable(arg0, arg1)
}

// SetEnableUndo: if enabled, changes to @editable will be saved for
// undo/redo actions.
//
// This results in an additional copy of text changes and are not stored in
// secure memory. As such, undo is forcefully disabled when
// [property@Gtk.Text:visibility] is set to false.
func (e editable) SetEnableUndo(enableUndo bool) {
	var arg0 *C.GtkEditable
	var arg1 C.gboolean

	arg0 = (*C.GtkEditable)(e.Native())
	if enableUndo {
		arg1 = C.TRUE
	}

	C.gtk_editable_set_enable_undo(arg0, arg1)
}

// SetMaxWidthChars sets the desired maximum width in characters of
// @editable.
func (e editable) SetMaxWidthChars(nChars int) {
	var arg0 *C.GtkEditable
	var arg1 C.int

	arg0 = (*C.GtkEditable)(e.Native())
	arg1 = C.int(nChars)

	C.gtk_editable_set_max_width_chars(arg0, arg1)
}

// SetPosition sets the cursor position in the editable to the given value.
//
// The cursor is displayed before the character with the given (base 0)
// index in the contents of the editable. The value must be less than or
// equal to the number of characters in the editable. A value of -1
// indicates that the position should be set after the last character of the
// editable. Note that @position is in characters, not in bytes.
func (e editable) SetPosition(position int) {
	var arg0 *C.GtkEditable
	var arg1 C.int

	arg0 = (*C.GtkEditable)(e.Native())
	arg1 = C.int(position)

	C.gtk_editable_set_position(arg0, arg1)
}

// SetText sets the text in the editable to the given value.
//
// This is replacing the current contents.
func (e editable) SetText(text string) {
	var arg0 *C.GtkEditable
	var arg1 *C.char

	arg0 = (*C.GtkEditable)(e.Native())
	arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_editable_set_text(arg0, arg1)
}

// SetWidthChars changes the size request of the editable to be about the
// right size for @n_chars characters.
//
// Note that it changes the size request, the size can still be affected by
// how you pack the widget into containers. If @n_chars is -1, the size
// reverts to the default size.
func (e editable) SetWidthChars(nChars int) {
	var arg0 *C.GtkEditable
	var arg1 C.int

	arg0 = (*C.GtkEditable)(e.Native())
	arg1 = C.int(nChars)

	C.gtk_editable_set_width_chars(arg0, arg1)
}
