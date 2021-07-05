// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_entry_icon_position_get_type()), F: marshalEntryIconPosition},
		{T: externglib.Type(C.gtk_entry_get_type()), F: marshalEntry},
	})
}

// EntryIconPosition specifies the side of the entry at which an icon is placed.
type EntryIconPosition int

const (
	// primary: at the beginning of the entry (depending on the text direction).
	EntryIconPositionPrimary EntryIconPosition = 0
	// secondary: at the end of the entry (depending on the text direction).
	EntryIconPositionSecondary EntryIconPosition = 1
)

func marshalEntryIconPosition(p uintptr) (interface{}, error) {
	return EntryIconPosition(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Entry: `GtkEntry` is a single line text entry widget.
//
// !An example GtkEntry (entry.png)
//
// A fairly large set of key bindings are supported by default. If the entered
// text is longer than the allocation of the widget, the widget will scroll so
// that the cursor position is visible.
//
// When using an entry for passwords and other sensitive information, it can be
// put into “password mode” using [method@Gtk.Entry.set_visibility]. In this
// mode, entered text is displayed using a “invisible” character. By default,
// GTK picks the best invisible character that is available in the current font,
// but it can be changed with [method@Gtk.Entry.set_invisible_char].
//
// `GtkEntry` has the ability to display progress or activity information behind
// the text. To make an entry display such information, use
// [method@Gtk.Entry.set_progress_fraction] or
// [method@Gtk.Entry.set_progress_pulse_step].
//
// Additionally, `GtkEntry` can show icons at either side of the entry. These
// icons can be activatable by clicking, can be set up as drag source and can
// have tooltips. To add an icon, use [method@Gtk.Entry.set_icon_from_gicon] or
// one of the various other functions that set an icon from an icon name or a
// paintable. To trigger an action when the user clicks an icon, connect to the
// [signal@Gtk.Entry::icon-press] signal. To allow DND operations from an icon,
// use [method@Gtk.Entry.set_icon_drag_source]. To set a tooltip on an icon, use
// [method@Gtk.Entry.set_icon_tooltip_text] or the corresponding function for
// markup.
//
// Note that functionality or information that is only available by clicking on
// an icon in an entry may not be accessible at all to users which are not able
// to use a mouse or other pointing device. It is therefore recommended that any
// such functionality should also be available by other means, e.g. via the
// context menu of the entry.
//
//
// CSS nodes
//
// “` entry[.flat][.warning][.error] ├── text[.readonly] ├── image.left ├──
// image.right ╰── [progress[.pulse]] “`
//
// `GtkEntry` has a main node with the name entry. Depending on the properties
// of the entry, the style classes .read-only and .flat may appear. The style
// classes .warning and .error may also be used with entries.
//
// When the entry shows icons, it adds subnodes with the name image and the
// style class .left or .right, depending on where the icon appears.
//
// When the entry shows progress, it adds a subnode with the name progress. The
// node has the style class .pulse when the shown progress is pulsing.
//
// For all the subnodes added to the text node in various situations, see
// [class@Gtk.Text].
//
//
// GtkEntry as GtkBuildable
//
// The `GtkEntry` implementation of the `GtkBuildable` interface supports a
// custom <attributes> element, which supports any number of <attribute>
// elements. The <attribute> element has attributes named “name“, “value“,
// “start“ and “end“ and allows you to specify Attribute values for this label.
//
// An example of a UI definition fragment specifying Pango attributes: “`xml
// <object class="GtkEnry"> <attributes> <attribute name="weight"
// value="PANGO_WEIGHT_BOLD"/> <attribute name="background" value="red"
// start="5" end="10"/> </attributes> </object> “`
//
// The start and end attributes specify the range of characters to which the
// Pango attribute applies. If start and end are not specified, the attribute is
// applied to the whole text. Note that specifying ranges does not make much
// sense with translatable attributes. Use markup embedded in the translatable
// content instead.
//
//
// Accessibility
//
// `GtkEntry` uses the GTK_ACCESSIBLE_ROLE_TEXT_BOX role.
type Entry interface {
	Widget

	// AsAccessible casts the class to the Accessible interface.
	AsAccessible() Accessible
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsCellEditable casts the class to the CellEditable interface.
	AsCellEditable() CellEditable
	// AsConstraintTarget casts the class to the ConstraintTarget interface.
	AsConstraintTarget() ConstraintTarget
	// AsEditable casts the class to the Editable interface.
	AsEditable() Editable

	// ActivatesDefault retrieves the value set by
	// gtk_entry_set_activates_default().
	ActivatesDefault() bool
	// Alignment gets the value set by gtk_entry_set_alignment().
	//
	// See also: [property@Gtk.Editable:xalign]
	Alignment() float32
	// Attributes gets the attribute list of the `GtkEntry`.
	//
	// See [method@Gtk.Entry.set_attributes].
	Attributes() pango.AttrList
	// Buffer: get the `GtkEntryBuffer` object which holds the text for this
	// widget.
	Buffer() EntryBuffer
	// Completion returns the auxiliary completion object currently in use by
	// @entry.
	Completion() EntryCompletion
	// CurrentIconDragSource returns the index of the icon which is the source
	// of the current DND operation, or -1.
	CurrentIconDragSource() int
	// ExtraMenu gets the menu model set with gtk_entry_set_extra_menu().
	ExtraMenu() gio.MenuModel
	// HasFrame gets the value set by gtk_entry_set_has_frame().
	HasFrame() bool
	// IconActivatable returns whether the icon is activatable.
	IconActivatable(iconPos EntryIconPosition) bool
	// IconArea gets the area where entry’s icon at @icon_pos is drawn.
	//
	// This function is useful when drawing something to the entry in a draw
	// callback.
	//
	// If the entry is not realized or has no icon at the given position,
	// @icon_area is filled with zeros. Otherwise, @icon_area will be filled
	// with the icon's allocation, relative to @entry's allocation.
	IconArea(iconPos EntryIconPosition) gdk.Rectangle
	// IconAtPos finds the icon at the given position and return its index.
	//
	// The position’s coordinates are relative to the @entry’s top left corner.
	// If @x, @y doesn’t lie inside an icon, -1 is returned. This function is
	// intended for use in a [signal@Gtk.Widget::query-tooltip] signal handler.
	IconAtPos(x int, y int) int
	// IconName retrieves the icon name used for the icon.
	//
	// nil is returned if there is no icon or if the icon was set by some other
	// method (e.g., by `GdkPaintable` or gicon).
	IconName(iconPos EntryIconPosition) string
	// IconSensitive returns whether the icon appears sensitive or insensitive.
	IconSensitive(iconPos EntryIconPosition) bool
	// IconStorageType gets the type of representation being used by the icon to
	// store image data.
	//
	// If the icon has no image data, the return value will be GTK_IMAGE_EMPTY.
	IconStorageType(iconPos EntryIconPosition) ImageType
	// IconTooltipMarkup gets the contents of the tooltip on the icon at the
	// specified position in @entry.
	IconTooltipMarkup(iconPos EntryIconPosition) string
	// IconTooltipText gets the contents of the tooltip on the icon at the
	// specified position in @entry.
	IconTooltipText(iconPos EntryIconPosition) string
	// InputHints gets the input hints of this `GtkEntry`.
	InputHints() InputHints
	// InputPurpose gets the input purpose of the `GtkEntry`.
	InputPurpose() InputPurpose
	// InvisibleChar retrieves the character displayed in place of the actual
	// text in “password mode”.
	InvisibleChar() uint32
	// MaxLength retrieves the maximum allowed length of the text in @entry.
	//
	// See [method@Gtk.Entry.set_max_length].
	MaxLength() int
	// OverwriteMode gets whether the `GtkEntry` is in overwrite mode.
	OverwriteMode() bool
	// PlaceholderText retrieves the text that will be displayed when @entry is
	// empty and unfocused
	PlaceholderText() string
	// ProgressFraction returns the current fraction of the task that’s been
	// completed.
	//
	// See [method@Gtk.Entry.set_progress_fraction].
	ProgressFraction() float64
	// ProgressPulseStep retrieves the pulse step set with
	// gtk_entry_set_progress_pulse_step().
	ProgressPulseStep() float64
	// Tabs gets the tabstops of the `GtkEntry.
	//
	// See [method@Gtk.Entry.set_tabs].
	Tabs() pango.TabArray
	// TextLength retrieves the current length of the text in @entry.
	//
	// This is equivalent to getting @entry's `GtkEntryBuffer` and calling
	// [method@Gtk.EntryBuffer.get_length] on it.
	TextLength() uint16
	// Visibility retrieves whether the text in @entry is visible.
	//
	// See [method@Gtk.Entry.set_visibility].
	Visibility() bool
	// GrabFocusWithoutSelectingEntry causes @entry to have keyboard focus.
	//
	// It behaves like [method@Gtk.Widget.grab_focus], except that it doesn't
	// select the contents of the entry. You only want to call this on some
	// special entries which the user usually doesn't want to replace all text
	// in, such as search-as-you-type entries.
	GrabFocusWithoutSelectingEntry() bool
	// ProgressPulseEntry indicates that some progress is made, but you don’t
	// know how much.
	//
	// Causes the entry’s progress indicator to enter “activity mode”, where a
	// block bounces back and forth. Each call to gtk_entry_progress_pulse()
	// causes the block to move by a little bit (the amount of movement per
	// pulse is determined by [method@Gtk.Entry.set_progress_pulse_step]).
	ProgressPulseEntry()
	// ResetImContextEntry: reset the input method context of the entry if
	// needed.
	//
	// This can be necessary in the case where modifying the buffer would
	// confuse on-going input method behavior.
	ResetImContextEntry()
	// SetActivatesDefaultEntry sets whether pressing Enter in the @entry will
	// activate the default widget for the window containing the entry.
	//
	// This usually means that the dialog containing the entry will be closed,
	// since the default widget is usually one of the dialog buttons.
	SetActivatesDefaultEntry(setting bool)
	// SetAlignmentEntry sets the alignment for the contents of the entry.
	//
	// This controls the horizontal positioning of the contents when the
	// displayed text is shorter than the width of the entry.
	//
	// See also: [property@Gtk.Editable:xalign]
	SetAlignmentEntry(xalign float32)
	// SetAttributesEntry sets a `PangoAttrList`.
	//
	// The attributes in the list are applied to the entry text.
	//
	// Since the attributes will be applies to text that changes as the user
	// types, it makes most sense to use attributes with unlimited extent.
	SetAttributesEntry(attrs pango.AttrList)
	// SetBufferEntry: set the `GtkEntryBuffer` object which holds the text for
	// this widget.
	SetBufferEntry(buffer EntryBuffer)
	// SetCompletionEntry sets @completion to be the auxiliary completion object
	// to use with @entry.
	//
	// All further configuration of the completion mechanism is done on
	// @completion using the `GtkEntryCompletion` API. Completion is disabled if
	// @completion is set to nil.
	SetCompletionEntry(completion EntryCompletion)
	// SetExtraMenuEntry sets a menu model to add when constructing the context
	// menu for @entry.
	SetExtraMenuEntry(model gio.MenuModel)
	// SetHasFrameEntry sets whether the entry has a beveled frame around it.
	SetHasFrameEntry(setting bool)
	// SetIconActivatableEntry sets whether the icon is activatable.
	SetIconActivatableEntry(iconPos EntryIconPosition, activatable bool)
	// SetIconDragSourceEntry sets up the icon at the given position as drag
	// source.
	//
	// This makes it so that GTK will start a drag operation when the user
	// clicks and drags the icon.
	SetIconDragSourceEntry(iconPos EntryIconPosition, provider gdk.ContentProvider, actions gdk.DragAction)
	// SetIconFromIconNameEntry sets the icon shown in the entry at the
	// specified position from the current icon theme.
	//
	// If the icon name isn’t known, a “broken image” icon will be displayed
	// instead.
	//
	// If @icon_name is nil, no icon will be shown in the specified position.
	SetIconFromIconNameEntry(iconPos EntryIconPosition, iconName string)
	// SetIconSensitiveEntry sets the sensitivity for the specified icon.
	SetIconSensitiveEntry(iconPos EntryIconPosition, sensitive bool)
	// SetIconTooltipMarkupEntry sets @tooltip as the contents of the tooltip
	// for the icon at the specified position.
	//
	// @tooltip is assumed to be marked up with Pango Markup.
	//
	// Use nil for @tooltip to remove an existing tooltip.
	//
	// See also [method@Gtk.Widget.set_tooltip_markup] and
	// [method@Gtk.Entry.set_icon_tooltip_text].
	SetIconTooltipMarkupEntry(iconPos EntryIconPosition, tooltip string)
	// SetIconTooltipTextEntry sets @tooltip as the contents of the tooltip for
	// the icon at the specified position.
	//
	// Use nil for @tooltip to remove an existing tooltip.
	//
	// See also [method@Gtk.Widget.set_tooltip_text] and
	// [method@Gtk.Entry.set_icon_tooltip_markup].
	//
	// If you unset the widget tooltip via [method@Gtk.Widget.set_tooltip_text]
	// or [method@Gtk.Widget.set_tooltip_markup], this sets
	// [property@Gtk.Widget:has-tooltip] to false, which suppresses icon
	// tooltips too. You can resolve this by then calling
	// [method@Gtk.Widget.set_has_tooltip] to set
	// [property@Gtk.Widget:has-tooltip] back to true, or setting at least one
	// non-empty tooltip on any icon achieves the same result.
	SetIconTooltipTextEntry(iconPos EntryIconPosition, tooltip string)
	// SetInputHintsEntry: set additional hints which allow input methods to
	// fine-tune their behavior.
	SetInputHintsEntry(hints InputHints)
	// SetInputPurposeEntry sets the input purpose which can be used by input
	// methods to adjust their behavior.
	SetInputPurposeEntry(purpose InputPurpose)
	// SetInvisibleCharEntry sets the character to use in place of the actual
	// text in “password mode”.
	//
	// See [method@Gtk.Entry.set_visibility] for how to enable “password mode”.
	//
	// By default, GTK picks the best invisible char available in the current
	// font. If you set the invisible char to 0, then the user will get no
	// feedback at all; there will be no text on the screen as they type.
	SetInvisibleCharEntry(ch uint32)
	// SetMaxLengthEntry sets the maximum allowed length of the contents of the
	// widget.
	//
	// If the current contents are longer than the given length, then they will
	// be truncated to fit.
	//
	// This is equivalent to getting @entry's `GtkEntryBuffer` and calling
	// [method@Gtk.EntryBuffer.set_max_length] on it.
	SetMaxLengthEntry(max int)
	// SetOverwriteModeEntry sets whether the text is overwritten when typing in
	// the `GtkEntry`.
	SetOverwriteModeEntry(overwrite bool)
	// SetPlaceholderTextEntry sets text to be displayed in @entry when it is
	// empty.
	//
	// This can be used to give a visual hint of the expected contents of the
	// `GtkEntry`.
	SetPlaceholderTextEntry(text string)
	// SetProgressFractionEntry causes the entry’s progress indicator to “fill
	// in” the given fraction of the bar.
	//
	// The fraction should be between 0.0 and 1.0, inclusive.
	SetProgressFractionEntry(fraction float64)
	// SetProgressPulseStepEntry sets the fraction of total entry width to move
	// the progress bouncing block for each pulse.
	//
	// Use [method@Gtk.Entry.progress_pulse] to pulse the progress.
	SetProgressPulseStepEntry(fraction float64)
	// SetTabsEntry sets a `PangoTabArray`.
	//
	// The tabstops in the array are applied to the entry text.
	SetTabsEntry(tabs pango.TabArray)
	// SetVisibilityEntry sets whether the contents of the entry are visible or
	// not.
	//
	// When visibility is set to false, characters are displayed as the
	// invisible char, and will also appear that way when the text in the entry
	// widget is copied elsewhere.
	//
	// By default, GTK picks the best invisible character available in the
	// current font, but it can be changed with
	// [method@Gtk.Entry.set_invisible_char].
	//
	// Note that you probably want to set [property@Gtk.Entry:input-purpose] to
	// GTK_INPUT_PURPOSE_PASSWORD or GTK_INPUT_PURPOSE_PIN to inform input
	// methods about the purpose of this entry, in addition to setting
	// visibility to false.
	SetVisibilityEntry(visible bool)
	// UnsetInvisibleCharEntry unsets the invisible char, so that the default
	// invisible char is used again. See [method@Gtk.Entry.set_invisible_char].
	UnsetInvisibleCharEntry()
}

// entry implements the Entry class.
type entry struct {
	Widget
}

// WrapEntry wraps a GObject to the right type. It is
// primarily used internally.
func WrapEntry(obj *externglib.Object) Entry {
	return entry{
		Widget: WrapWidget(obj),
	}
}

func marshalEntry(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEntry(obj), nil
}

// NewEntry creates a new entry.
func NewEntry() Entry {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_entry_new()

	var _entry Entry // out

	_entry = WrapEntry(externglib.Take(unsafe.Pointer(_cret)))

	return _entry
}

// NewEntryWithBuffer creates a new entry with the specified text buffer.
func NewEntryWithBuffer(buffer EntryBuffer) Entry {
	var _arg1 *C.GtkEntryBuffer // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))

	_cret = C.gtk_entry_new_with_buffer(_arg1)

	var _entry Entry // out

	_entry = WrapEntry(externglib.Take(unsafe.Pointer(_cret)))

	return _entry
}

func (e entry) ActivatesDefault() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_activates_default(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e entry) Alignment() float32 {
	var _arg0 *C.GtkEntry // out
	var _cret C.float     // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_alignment(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

func (e entry) Attributes() pango.AttrList {
	var _arg0 *C.GtkEntry      // out
	var _cret *C.PangoAttrList // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_attributes(_arg0)

	var _attrList pango.AttrList // out

	_attrList = (pango.AttrList)(unsafe.Pointer(_cret))
	C.pango_attr_list_ref(_cret)

	return _attrList
}

func (e entry) Buffer() EntryBuffer {
	var _arg0 *C.GtkEntry       // out
	var _cret *C.GtkEntryBuffer // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_buffer(_arg0)

	var _entryBuffer EntryBuffer // out

	_entryBuffer = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(EntryBuffer)

	return _entryBuffer
}

func (e entry) Completion() EntryCompletion {
	var _arg0 *C.GtkEntry           // out
	var _cret *C.GtkEntryCompletion // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_completion(_arg0)

	var _entryCompletion EntryCompletion // out

	_entryCompletion = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(EntryCompletion)

	return _entryCompletion
}

func (e entry) CurrentIconDragSource() int {
	var _arg0 *C.GtkEntry // out
	var _cret C.int       // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_current_icon_drag_source(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (e entry) ExtraMenu() gio.MenuModel {
	var _arg0 *C.GtkEntry   // out
	var _cret *C.GMenuModel // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_extra_menu(_arg0)

	var _menuModel gio.MenuModel // out

	_menuModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gio.MenuModel)

	return _menuModel
}

func (e entry) HasFrame() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_has_frame(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e entry) IconActivatable(iconPos EntryIconPosition) bool {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	_cret = C.gtk_entry_get_icon_activatable(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e entry) IconArea(iconPos EntryIconPosition) gdk.Rectangle {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 *C.GdkRectangle        // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	C.gtk_entry_get_icon_area(_arg0, _arg1, &_arg2)

	var _iconArea gdk.Rectangle // out

	_iconArea = (gdk.Rectangle)(unsafe.Pointer(_arg2))

	return _iconArea
}

func (e entry) IconAtPos(x int, y int) int {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.int       // out
	var _arg2 C.int       // out
	var _cret C.int       // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.int(x)
	_arg2 = C.int(y)

	_cret = C.gtk_entry_get_icon_at_pos(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (e entry) IconName(iconPos EntryIconPosition) string {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _cret *C.char                // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	_cret = C.gtk_entry_get_icon_name(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (e entry) IconSensitive(iconPos EntryIconPosition) bool {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	_cret = C.gtk_entry_get_icon_sensitive(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e entry) IconStorageType(iconPos EntryIconPosition) ImageType {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _cret C.GtkImageType         // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	_cret = C.gtk_entry_get_icon_storage_type(_arg0, _arg1)

	var _imageType ImageType // out

	_imageType = ImageType(_cret)

	return _imageType
}

func (e entry) IconTooltipMarkup(iconPos EntryIconPosition) string {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _cret *C.char                // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	_cret = C.gtk_entry_get_icon_tooltip_markup(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (e entry) IconTooltipText(iconPos EntryIconPosition) string {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _cret *C.char                // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	_cret = C.gtk_entry_get_icon_tooltip_text(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (e entry) InputHints() InputHints {
	var _arg0 *C.GtkEntry     // out
	var _cret C.GtkInputHints // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_input_hints(_arg0)

	var _inputHints InputHints // out

	_inputHints = InputHints(_cret)

	return _inputHints
}

func (e entry) InputPurpose() InputPurpose {
	var _arg0 *C.GtkEntry       // out
	var _cret C.GtkInputPurpose // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_input_purpose(_arg0)

	var _inputPurpose InputPurpose // out

	_inputPurpose = InputPurpose(_cret)

	return _inputPurpose
}

func (e entry) InvisibleChar() uint32 {
	var _arg0 *C.GtkEntry // out
	var _cret C.gunichar  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_invisible_char(_arg0)

	var _gunichar uint32 // out

	_gunichar = uint32(_cret)

	return _gunichar
}

func (e entry) MaxLength() int {
	var _arg0 *C.GtkEntry // out
	var _cret C.int       // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_max_length(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (e entry) OverwriteMode() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_overwrite_mode(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e entry) PlaceholderText() string {
	var _arg0 *C.GtkEntry // out
	var _cret *C.char     // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_placeholder_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (e entry) ProgressFraction() float64 {
	var _arg0 *C.GtkEntry // out
	var _cret C.double    // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_progress_fraction(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (e entry) ProgressPulseStep() float64 {
	var _arg0 *C.GtkEntry // out
	var _cret C.double    // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_progress_pulse_step(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (e entry) Tabs() pango.TabArray {
	var _arg0 *C.GtkEntry      // out
	var _cret *C.PangoTabArray // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_tabs(_arg0)

	var _tabArray pango.TabArray // out

	_tabArray = (pango.TabArray)(unsafe.Pointer(_cret))

	return _tabArray
}

func (e entry) TextLength() uint16 {
	var _arg0 *C.GtkEntry // out
	var _cret C.guint16   // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_text_length(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

func (e entry) Visibility() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_visibility(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e entry) GrabFocusWithoutSelectingEntry() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_grab_focus_without_selecting(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e entry) ProgressPulseEntry() {
	var _arg0 *C.GtkEntry // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	C.gtk_entry_progress_pulse(_arg0)
}

func (e entry) ResetImContextEntry() {
	var _arg0 *C.GtkEntry // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	C.gtk_entry_reset_im_context(_arg0)
}

func (e entry) SetActivatesDefaultEntry(setting bool) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_entry_set_activates_default(_arg0, _arg1)
}

func (e entry) SetAlignmentEntry(xalign float32) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.float     // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.float(xalign)

	C.gtk_entry_set_alignment(_arg0, _arg1)
}

func (e entry) SetAttributesEntry(attrs pango.AttrList) {
	var _arg0 *C.GtkEntry      // out
	var _arg1 *C.PangoAttrList // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.PangoAttrList)(unsafe.Pointer(attrs))

	C.gtk_entry_set_attributes(_arg0, _arg1)
}

func (e entry) SetBufferEntry(buffer EntryBuffer) {
	var _arg0 *C.GtkEntry       // out
	var _arg1 *C.GtkEntryBuffer // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))

	C.gtk_entry_set_buffer(_arg0, _arg1)
}

func (e entry) SetCompletionEntry(completion EntryCompletion) {
	var _arg0 *C.GtkEntry           // out
	var _arg1 *C.GtkEntryCompletion // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	C.gtk_entry_set_completion(_arg0, _arg1)
}

func (e entry) SetExtraMenuEntry(model gio.MenuModel) {
	var _arg0 *C.GtkEntry   // out
	var _arg1 *C.GMenuModel // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	C.gtk_entry_set_extra_menu(_arg0, _arg1)
}

func (e entry) SetHasFrameEntry(setting bool) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_entry_set_has_frame(_arg0, _arg1)
}

func (e entry) SetIconActivatableEntry(iconPos EntryIconPosition, activatable bool) {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 C.gboolean             // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)
	if activatable {
		_arg2 = C.TRUE
	}

	C.gtk_entry_set_icon_activatable(_arg0, _arg1, _arg2)
}

func (e entry) SetIconDragSourceEntry(iconPos EntryIconPosition, provider gdk.ContentProvider, actions gdk.DragAction) {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 *C.GdkContentProvider  // out
	var _arg3 C.GdkDragAction        // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)
	_arg2 = (*C.GdkContentProvider)(unsafe.Pointer(provider.Native()))
	_arg3 = C.GdkDragAction(actions)

	C.gtk_entry_set_icon_drag_source(_arg0, _arg1, _arg2, _arg3)
}

func (e entry) SetIconFromIconNameEntry(iconPos EntryIconPosition, iconName string) {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 *C.char                // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)
	_arg2 = (*C.char)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_entry_set_icon_from_icon_name(_arg0, _arg1, _arg2)
}

func (e entry) SetIconSensitiveEntry(iconPos EntryIconPosition, sensitive bool) {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 C.gboolean             // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)
	if sensitive {
		_arg2 = C.TRUE
	}

	C.gtk_entry_set_icon_sensitive(_arg0, _arg1, _arg2)
}

func (e entry) SetIconTooltipMarkupEntry(iconPos EntryIconPosition, tooltip string) {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 *C.char                // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)
	_arg2 = (*C.char)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_entry_set_icon_tooltip_markup(_arg0, _arg1, _arg2)
}

func (e entry) SetIconTooltipTextEntry(iconPos EntryIconPosition, tooltip string) {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 *C.char                // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)
	_arg2 = (*C.char)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_entry_set_icon_tooltip_text(_arg0, _arg1, _arg2)
}

func (e entry) SetInputHintsEntry(hints InputHints) {
	var _arg0 *C.GtkEntry     // out
	var _arg1 C.GtkInputHints // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkInputHints(hints)

	C.gtk_entry_set_input_hints(_arg0, _arg1)
}

func (e entry) SetInputPurposeEntry(purpose InputPurpose) {
	var _arg0 *C.GtkEntry       // out
	var _arg1 C.GtkInputPurpose // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkInputPurpose(purpose)

	C.gtk_entry_set_input_purpose(_arg0, _arg1)
}

func (e entry) SetInvisibleCharEntry(ch uint32) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gunichar  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.gunichar(ch)

	C.gtk_entry_set_invisible_char(_arg0, _arg1)
}

func (e entry) SetMaxLengthEntry(max int) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.int(max)

	C.gtk_entry_set_max_length(_arg0, _arg1)
}

func (e entry) SetOverwriteModeEntry(overwrite bool) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	if overwrite {
		_arg1 = C.TRUE
	}

	C.gtk_entry_set_overwrite_mode(_arg0, _arg1)
}

func (e entry) SetPlaceholderTextEntry(text string) {
	var _arg0 *C.GtkEntry // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_entry_set_placeholder_text(_arg0, _arg1)
}

func (e entry) SetProgressFractionEntry(fraction float64) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.double    // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.double(fraction)

	C.gtk_entry_set_progress_fraction(_arg0, _arg1)
}

func (e entry) SetProgressPulseStepEntry(fraction float64) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.double    // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.double(fraction)

	C.gtk_entry_set_progress_pulse_step(_arg0, _arg1)
}

func (e entry) SetTabsEntry(tabs pango.TabArray) {
	var _arg0 *C.GtkEntry      // out
	var _arg1 *C.PangoTabArray // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.PangoTabArray)(unsafe.Pointer(tabs))

	C.gtk_entry_set_tabs(_arg0, _arg1)
}

func (e entry) SetVisibilityEntry(visible bool) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_entry_set_visibility(_arg0, _arg1)
}

func (e entry) UnsetInvisibleCharEntry() {
	var _arg0 *C.GtkEntry // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	C.gtk_entry_unset_invisible_char(_arg0)
}

func (e entry) AsAccessible() Accessible {
	return WrapAccessible(gextras.InternObject(e))
}

func (e entry) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(e))
}

func (e entry) AsCellEditable() CellEditable {
	return WrapCellEditable(gextras.InternObject(e))
}

func (e entry) AsConstraintTarget() ConstraintTarget {
	return WrapConstraintTarget(gextras.InternObject(e))
}

func (e entry) AsEditable() Editable {
	return WrapEditable(gextras.InternObject(e))
}
