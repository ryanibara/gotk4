// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk4_EntryClass_activate(void*);
// extern void _gotk4_gtk4_Entry_ConnectActivate(gpointer, guintptr);
import "C"

// GTypeEntryIconPosition returns the GType for the type EntryIconPosition.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeEntryIconPosition() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "EntryIconPosition").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalEntryIconPosition)
	return gtype
}

// GTypeEntry returns the GType for the type Entry.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeEntry() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Entry").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalEntry)
	return gtype
}

// EntryIconPosition specifies the side of the entry at which an icon is placed.
type EntryIconPosition C.gint

const (
	// EntryIconPrimary: at the beginning of the entry (depending on the text
	// direction).
	EntryIconPrimary EntryIconPosition = iota
	// EntryIconSecondary: at the end of the entry (depending on the text
	// direction).
	EntryIconSecondary
)

func marshalEntryIconPosition(p uintptr) (interface{}, error) {
	return EntryIconPosition(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for EntryIconPosition.
func (e EntryIconPosition) String() string {
	switch e {
	case EntryIconPrimary:
		return "Primary"
	case EntryIconSecondary:
		return "Secondary"
	default:
		return fmt.Sprintf("EntryIconPosition(%d)", e)
	}
}

// EntryOverrider contains methods that are overridable.
type EntryOverrider interface {
	Activate()
}

// Entry: GtkEntry is a single line text entry widget.
//
// !An example GtkEntry (entry.png)
//
// A fairly large set of key bindings are supported by default. If the entered
// text is longer than the allocation of the widget, the widget will scroll so
// that the cursor position is visible.
//
// When using an entry for passwords and other sensitive information, it can be
// put into “password mode” using gtk.Entry.SetVisibility(). In this mode,
// entered text is displayed using a “invisible” character. By default, GTK
// picks the best invisible character that is available in the current font, but
// it can be changed with gtk.Entry.SetInvisibleChar().
//
// GtkEntry has the ability to display progress or activity information behind
// the text. To make an entry display such information, use
// gtk.Entry.SetProgressFraction() or gtk.Entry.SetProgressPulseStep().
//
// Additionally, GtkEntry can show icons at either side of the entry. These
// icons can be activatable by clicking, can be set up as drag source and can
// have tooltips. To add an icon, use gtk.Entry.SetIconFromGIcon() or one of the
// various other functions that set an icon from an icon name or a paintable. To
// trigger an action when the user clicks an icon, connect to the
// gtk.Entry::icon-press signal. To allow DND operations from an icon, use
// gtk.Entry.SetIconDragSource(). To set a tooltip on an icon, use
// gtk.Entry.SetIconTooltipText() or the corresponding function for markup.
//
// Note that functionality or information that is only available by clicking on
// an icon in an entry may not be accessible at all to users which are not able
// to use a mouse or other pointing device. It is therefore recommended that any
// such functionality should also be available by other means, e.g. via the
// context menu of the entry.
//
// CSS nodes
//
//    entry[.flat][.warning][.error]
//    ├── text[.readonly]
//    ├── image.left
//    ├── image.right
//    ╰── [progress[.pulse]]
//
//
// GtkEntry has a main node with the name entry. Depending on the properties of
// the entry, the style classes .read-only and .flat may appear. The style
// classes .warning and .error may also be used with entries.
//
// When the entry shows icons, it adds subnodes with the name image and the
// style class .left or .right, depending on where the icon appears.
//
// When the entry shows progress, it adds a subnode with the name progress. The
// node has the style class .pulse when the shown progress is pulsing.
//
// For all the subnodes added to the text node in various situations, see
// gtk.Text.
//
//
// GtkEntry as GtkBuildable
//
// The GtkEntry implementation of the GtkBuildable interface supports a custom
// <attributes> element, which supports any number of <attribute> elements. The
// <attribute> element has attributes named “name“, “value“, “start“ and “end“
// and allows you to specify Attribute values for this label.
//
// An example of a UI definition fragment specifying Pango attributes:
//
//    <object class="GtkEnry">
//      <attributes>
//        <attribute name="weight" value="PANGO_WEIGHT_BOLD"/>
//        <attribute name="background" value="red" start="5" end="10"/>
//      </attributes>
//    </object>
//
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
// GtkEntry uses the GTK_ACCESSIBLE_ROLE_TEXT_BOX role.
type Entry struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	coreglib.InitiallyUnowned
	Accessible
	Buildable
	CellEditable
	ConstraintTarget
	Editable
}

var (
	_ Widgetter         = (*Entry)(nil)
	_ coreglib.Objector = (*Entry)(nil)
)

func classInitEntrier(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "EntryClass")

	if _, ok := goval.(interface{ Activate() }); ok {
		o := pclass.StructFieldOffset("activate")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_EntryClass_activate)
	}
}

//export _gotk4_gtk4_EntryClass_activate
func _gotk4_gtk4_EntryClass_activate(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Activate() })

	iface.Activate()
}

func wrapEntry(obj *coreglib.Object) *Entry {
	return &Entry{
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
		InitiallyUnowned: coreglib.InitiallyUnowned{
			Object: obj,
		},
		Accessible: Accessible{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		CellEditable: CellEditable{
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
		ConstraintTarget: ConstraintTarget{
			Object: obj,
		},
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

func marshalEntry(p uintptr) (interface{}, error) {
	return wrapEntry(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_Entry_ConnectActivate
func _gotk4_gtk4_Entry_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectActivate is emitted when the entry is activated.
//
// The keybindings for this signal are all forms of the Enter key.
func (entry *Entry) ConnectActivate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(entry, "activate", false, unsafe.Pointer(C._gotk4_gtk4_Entry_ConnectActivate), f)
}

// NewEntry creates a new entry.
//
// The function returns the following values:
//
//    - entry: new GtkEntry.
//
func NewEntry() *Entry {
	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("new_Entry", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _entry *Entry // out

	_entry = wrapEntry(coreglib.Take(unsafe.Pointer(_cret)))

	return _entry
}

// NewEntryWithBuffer creates a new entry with the specified text buffer.
//
// The function takes the following parameters:
//
//    - buffer to use for the new GtkEntry.
//
// The function returns the following values:
//
//    - entry: new GtkEntry.
//
func NewEntryWithBuffer(buffer *EntryBuffer) *Entry {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("new_Entry_with_buffer", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(buffer)

	var _entry *Entry // out

	_entry = wrapEntry(coreglib.Take(unsafe.Pointer(_cret)))

	return _entry
}

// ActivatesDefault retrieves the value set by
// gtk_entry_set_activates_default().
//
// The function returns the following values:
//
//    - ok: TRUE if the entry will activate the default widget.
//
func (entry *Entry) ActivatesDefault() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_activates_default", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Alignment gets the value set by gtk_entry_set_alignment().
//
// See also: gtk.Editable:xalign.
//
// The function returns the following values:
//
//    - gfloat: alignment.
//
func (entry *Entry) Alignment() float32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_alignment", _args[:], nil)
	_cret = *(*C.float)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _gfloat float32 // out

	_gfloat = float32(*(*C.float)(unsafe.Pointer(&_cret)))

	return _gfloat
}

// Attributes gets the attribute list of the GtkEntry.
//
// See gtk.Entry.SetAttributes().
//
// The function returns the following values:
//
//    - attrList (optional): attribute list, or NULL if none was set.
//
func (entry *Entry) Attributes() *pango.AttrList {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_attributes", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _attrList *pango.AttrList // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_attrList = (*pango.AttrList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.pango_attr_list_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_attrList)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _attrList
}

// Buffer: get the GtkEntryBuffer object which holds the text for this widget.
//
// The function returns the following values:
//
//    - entryBuffer: GtkEntryBuffer object.
//
func (entry *Entry) Buffer() *EntryBuffer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_buffer", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _entryBuffer *EntryBuffer // out

	_entryBuffer = wrapEntryBuffer(coreglib.Take(unsafe.Pointer(_cret)))

	return _entryBuffer
}

// Completion returns the auxiliary completion object currently in use by entry.
//
// The function returns the following values:
//
//    - entryCompletion (optional): auxiliary completion object currently in use
//      by entry.
//
func (entry *Entry) Completion() *EntryCompletion {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_completion", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _entryCompletion *EntryCompletion // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_entryCompletion = wrapEntryCompletion(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _entryCompletion
}

// CurrentIconDragSource returns the index of the icon which is the source of
// the current DND operation, or -1.
//
// The function returns the following values:
//
//    - gint: index of the icon which is the source of the current DND operation,
//      or -1.
//
func (entry *Entry) CurrentIconDragSource() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_current_icon_drag_source", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// ExtraMenu gets the menu model set with gtk_entry_set_extra_menu().
//
// The function returns the following values:
//
//    - menuModel (optional): menu model.
//
func (entry *Entry) ExtraMenu() gio.MenuModeller {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_extra_menu", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _menuModel gio.MenuModeller // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gio.MenuModeller)
				return ok
			})
			rv, ok := casted.(gio.MenuModeller)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.MenuModeller")
			}
			_menuModel = rv
		}
	}

	return _menuModel
}

// HasFrame gets the value set by gtk_entry_set_has_frame().
//
// The function returns the following values:
//
//    - ok: whether the entry has a beveled frame.
//
func (entry *Entry) HasFrame() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_has_frame", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IconAtPos finds the icon at the given position and return its index.
//
// The position’s coordinates are relative to the entry’s top left corner. If x,
// y doesn’t lie inside an icon, -1 is returned. This function is intended for
// use in a gtk.Widget::query-tooltip signal handler.
//
// The function takes the following parameters:
//
//    - x coordinate of the position to find, relative to entry.
//    - y coordinate of the position to find, relative to entry.
//
// The function returns the following values:
//
//    - gint: index of the icon at the given position, or -1.
//
func (entry *Entry) IconAtPos(x, y int32) int32 {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(x)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(y)

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_icon_at_pos", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// InvisibleChar retrieves the character displayed in place of the actual text
// in “password mode”.
//
// The function returns the following values:
//
//    - gunichar: current invisible char, or 0, if the entry does not show
//      invisible text at all.
//
func (entry *Entry) InvisibleChar() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_invisible_char", _args[:], nil)
	_cret = *(*C.gunichar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _gunichar uint32 // out

	_gunichar = uint32(*(*C.gunichar)(unsafe.Pointer(&_cret)))

	return _gunichar
}

// MaxLength retrieves the maximum allowed length of the text in entry.
//
// See gtk.Entry.SetMaxLength().
//
// The function returns the following values:
//
//    - gint: maximum allowed number of characters in GtkEntry, or 0 if there is
//      no maximum.
//
func (entry *Entry) MaxLength() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_max_length", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// OverwriteMode gets whether the GtkEntry is in overwrite mode.
//
// The function returns the following values:
//
//    - ok: whether the text is overwritten when typing.
//
func (entry *Entry) OverwriteMode() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_overwrite_mode", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// PlaceholderText retrieves the text that will be displayed when entry is empty
// and unfocused.
//
// The function returns the following values:
//
//    - utf8 (optional): pointer to the placeholder text as a string. This string
//      points to internally allocated storage in the widget and must not be
//      freed, modified or stored. If no placeholder text has been set, NULL will
//      be returned.
//
func (entry *Entry) PlaceholderText() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_placeholder_text", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ProgressFraction returns the current fraction of the task that’s been
// completed.
//
// See gtk.Entry.SetProgressFraction().
//
// The function returns the following values:
//
//    - gdouble: fraction from 0.0 to 1.0.
//
func (entry *Entry) ProgressFraction() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_progress_fraction", _args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// ProgressPulseStep retrieves the pulse step set with
// gtk_entry_set_progress_pulse_step().
//
// The function returns the following values:
//
//    - gdouble: fraction from 0.0 to 1.0.
//
func (entry *Entry) ProgressPulseStep() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_progress_pulse_step", _args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// Tabs gets the tabstops of the `GtkEntry.
//
// See gtk.Entry.SetTabs().
//
// The function returns the following values:
//
//    - tabArray (optional): tabstops, or NULL if none was set.
//
func (entry *Entry) Tabs() *pango.TabArray {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_tabs", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _tabArray *pango.TabArray // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_tabArray = (*pango.TabArray)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _tabArray
}

// TextLength retrieves the current length of the text in entry.
//
// This is equivalent to getting entry's GtkEntryBuffer and calling
// gtk.EntryBuffer.GetLength() on it.
//
// The function returns the following values:
//
//    - guint16: current number of characters in GtkEntry, or 0 if there are
//      none.
//
func (entry *Entry) TextLength() uint16 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_text_length", _args[:], nil)
	_cret = *(*C.guint16)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _guint16 uint16 // out

	_guint16 = uint16(*(*C.guint16)(unsafe.Pointer(&_cret)))

	return _guint16
}

// Visibility retrieves whether the text in entry is visible.
//
// See gtk.Entry.SetVisibility().
//
// The function returns the following values:
//
//    - ok: TRUE if the text is currently visible.
//
func (entry *Entry) Visibility() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("get_visibility", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// GrabFocusWithoutSelecting causes entry to have keyboard focus.
//
// It behaves like gtk.Widget.GrabFocus(), except that it doesn't select the
// contents of the entry. You only want to call this on some special entries
// which the user usually doesn't want to replace all text in, such as
// search-as-you-type entries.
//
// The function returns the following values:
//
//    - ok: TRUE if focus is now inside self.
//
func (entry *Entry) GrabFocusWithoutSelecting() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	_gret := girepository.MustFind("Gtk", "Entry").InvokeMethod("grab_focus_without_selecting", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ProgressPulse indicates that some progress is made, but you don’t know how
// much.
//
// Causes the entry’s progress indicator to enter “activity mode”, where a block
// bounces back and forth. Each call to gtk_entry_progress_pulse() causes the
// block to move by a little bit (the amount of movement per pulse is determined
// by gtk.Entry.SetProgressPulseStep()).
func (entry *Entry) ProgressPulse() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	girepository.MustFind("Gtk", "Entry").InvokeMethod("progress_pulse", _args[:], nil)

	runtime.KeepAlive(entry)
}

// ResetIMContext: reset the input method context of the entry if needed.
//
// This can be necessary in the case where modifying the buffer would confuse
// on-going input method behavior.
func (entry *Entry) ResetIMContext() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	girepository.MustFind("Gtk", "Entry").InvokeMethod("reset_im_context", _args[:], nil)

	runtime.KeepAlive(entry)
}

// SetActivatesDefault sets whether pressing Enter in the entry will activate
// the default widget for the window containing the entry.
//
// This usually means that the dialog containing the entry will be closed, since
// the default widget is usually one of the dialog buttons.
//
// The function takes the following parameters:
//
//    - setting: TRUE to activate window’s default widget on Enter keypress.
//
func (entry *Entry) SetActivatesDefault(setting bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	if setting {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Entry").InvokeMethod("set_activates_default", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(setting)
}

// SetAlignment sets the alignment for the contents of the entry.
//
// This controls the horizontal positioning of the contents when the displayed
// text is shorter than the width of the entry.
//
// See also: gtk.Editable:xalign.
//
// The function takes the following parameters:
//
//    - xalign: horizontal alignment, from 0 (left) to 1 (right). Reversed for
//      RTL layouts.
//
func (entry *Entry) SetAlignment(xalign float32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	*(*C.float)(unsafe.Pointer(&_args[1])) = C.float(xalign)

	girepository.MustFind("Gtk", "Entry").InvokeMethod("set_alignment", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(xalign)
}

// SetAttributes sets a PangoAttrList.
//
// The attributes in the list are applied to the entry text.
//
// Since the attributes will be applies to text that changes as the user types,
// it makes most sense to use attributes with unlimited extent.
//
// The function takes the following parameters:
//
//    - attrs: PangoAttrList.
//
func (entry *Entry) SetAttributes(attrs *pango.AttrList) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(attrs)))

	girepository.MustFind("Gtk", "Entry").InvokeMethod("set_attributes", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(attrs)
}

// SetBuffer: set the GtkEntryBuffer object which holds the text for this
// widget.
//
// The function takes the following parameters:
//
//    - buffer: GtkEntryBuffer.
//
func (entry *Entry) SetBuffer(buffer *EntryBuffer) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	girepository.MustFind("Gtk", "Entry").InvokeMethod("set_buffer", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(buffer)
}

// SetCompletion sets completion to be the auxiliary completion object to use
// with entry.
//
// All further configuration of the completion mechanism is done on completion
// using the GtkEntryCompletion API. Completion is disabled if completion is set
// to NULL.
//
// The function takes the following parameters:
//
//    - completion (optional): GtkEntryCompletion or NULL.
//
func (entry *Entry) SetCompletion(completion *EntryCompletion) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	if completion != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	}

	girepository.MustFind("Gtk", "Entry").InvokeMethod("set_completion", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(completion)
}

// SetExtraMenu sets a menu model to add when constructing the context menu for
// entry.
//
// The function takes the following parameters:
//
//    - model (optional): GMenuModel.
//
func (entry *Entry) SetExtraMenu(model gio.MenuModeller) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	if model != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	girepository.MustFind("Gtk", "Entry").InvokeMethod("set_extra_menu", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(model)
}

// SetHasFrame sets whether the entry has a beveled frame around it.
//
// The function takes the following parameters:
//
//    - setting: new value.
//
func (entry *Entry) SetHasFrame(setting bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	if setting {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Entry").InvokeMethod("set_has_frame", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(setting)
}

// SetInvisibleChar sets the character to use in place of the actual text in
// “password mode”.
//
// See gtk.Entry.SetVisibility() for how to enable “password mode”.
//
// By default, GTK picks the best invisible char available in the current font.
// If you set the invisible char to 0, then the user will get no feedback at
// all; there will be no text on the screen as they type.
//
// The function takes the following parameters:
//
//    - ch: unicode character.
//
func (entry *Entry) SetInvisibleChar(ch uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	*(*C.gunichar)(unsafe.Pointer(&_args[1])) = C.gunichar(ch)

	girepository.MustFind("Gtk", "Entry").InvokeMethod("set_invisible_char", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(ch)
}

// SetMaxLength sets the maximum allowed length of the contents of the widget.
//
// If the current contents are longer than the given length, then they will be
// truncated to fit.
//
// This is equivalent to getting entry's GtkEntryBuffer and calling
// gtk.EntryBuffer.SetMaxLength() on it.
//
// The function takes the following parameters:
//
//    - max: maximum length of the entry, or 0 for no maximum. (other than the
//      maximum length of entries.) The value passed in will be clamped to the
//      range 0-65536.
//
func (entry *Entry) SetMaxLength(max int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(max)

	girepository.MustFind("Gtk", "Entry").InvokeMethod("set_max_length", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(max)
}

// SetOverwriteMode sets whether the text is overwritten when typing in the
// GtkEntry.
//
// The function takes the following parameters:
//
//    - overwrite: new value.
//
func (entry *Entry) SetOverwriteMode(overwrite bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	if overwrite {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Entry").InvokeMethod("set_overwrite_mode", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(overwrite)
}

// SetPlaceholderText sets text to be displayed in entry when it is empty.
//
// This can be used to give a visual hint of the expected contents of the
// GtkEntry.
//
// The function takes the following parameters:
//
//    - text (optional): string to be displayed when entry is empty and
//      unfocused, or NULL.
//
func (entry *Entry) SetPlaceholderText(text string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	if text != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(text)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "Entry").InvokeMethod("set_placeholder_text", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(text)
}

// SetProgressFraction causes the entry’s progress indicator to “fill in” the
// given fraction of the bar.
//
// The fraction should be between 0.0 and 1.0, inclusive.
//
// The function takes the following parameters:
//
//    - fraction of the task that’s been completed.
//
func (entry *Entry) SetProgressFraction(fraction float64) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(fraction)

	girepository.MustFind("Gtk", "Entry").InvokeMethod("set_progress_fraction", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(fraction)
}

// SetProgressPulseStep sets the fraction of total entry width to move the
// progress bouncing block for each pulse.
//
// Use gtk.Entry.ProgressPulse() to pulse the progress.
//
// The function takes the following parameters:
//
//    - fraction between 0.0 and 1.0.
//
func (entry *Entry) SetProgressPulseStep(fraction float64) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(fraction)

	girepository.MustFind("Gtk", "Entry").InvokeMethod("set_progress_pulse_step", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(fraction)
}

// SetTabs sets a PangoTabArray.
//
// The tabstops in the array are applied to the entry text.
//
// The function takes the following parameters:
//
//    - tabs (optional): PangoTabArray.
//
func (entry *Entry) SetTabs(tabs *pango.TabArray) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	if tabs != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(tabs)))
	}

	girepository.MustFind("Gtk", "Entry").InvokeMethod("set_tabs", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(tabs)
}

// SetVisibility sets whether the contents of the entry are visible or not.
//
// When visibility is set to FALSE, characters are displayed as the invisible
// char, and will also appear that way when the text in the entry widget is
// copied elsewhere.
//
// By default, GTK picks the best invisible character available in the current
// font, but it can be changed with gtk.Entry.SetInvisibleChar().
//
// Note that you probably want to set gtk.Entry:input-purpose to
// GTK_INPUT_PURPOSE_PASSWORD or GTK_INPUT_PURPOSE_PIN to inform input methods
// about the purpose of this entry, in addition to setting visibility to FALSE.
//
// The function takes the following parameters:
//
//    - visible: TRUE if the contents of the entry are displayed as plaintext.
//
func (entry *Entry) SetVisibility(visible bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	if visible {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Entry").InvokeMethod("set_visibility", _args[:], nil)

	runtime.KeepAlive(entry)
	runtime.KeepAlive(visible)
}

// UnsetInvisibleChar unsets the invisible char, so that the default invisible
// char is used again. See gtk.Entry.SetInvisibleChar().
func (entry *Entry) UnsetInvisibleChar() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	girepository.MustFind("Gtk", "Entry").InvokeMethod("unset_invisible_char", _args[:], nil)

	runtime.KeepAlive(entry)
}
