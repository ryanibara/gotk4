// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tool_palette_get_type()), F: marshalToolPalette},
	})
}

// ToolPalette: a ToolPalette allows you to add ToolItems to a palette-like
// container with different categories and drag and drop support.
//
// A ToolPalette is created with a call to gtk_tool_palette_new().
//
// ToolItems cannot be added directly to a ToolPalette - instead they are added
// to a ToolItemGroup which can than be added to a ToolPalette. To add a
// ToolItemGroup to a ToolPalette, use gtk_container_add().
//
//    static void
//    passive_canvas_drag_data_received (GtkWidget        *widget,
//                                       GdkDragContext   *context,
//                                       gint              x,
//                                       gint              y,
//                                       GtkSelectionData *selection,
//                                       guint             info,
//                                       guint             time,
//                                       gpointer          data)
//    {
//      GtkWidget *palette;
//      GtkWidget *item;
//
//      // Get the dragged item
//      palette = gtk_widget_get_ancestor (gtk_drag_get_source_widget (context),
//                                         GTK_TYPE_TOOL_PALETTE);
//      if (palette != NULL)
//        item = gtk_tool_palette_get_drag_item (GTK_TOOL_PALETTE (palette),
//                                               selection);
//
//      // Do something with item
//    }
//
//    GtkWidget *target, palette;
//
//    palette = gtk_tool_palette_new ();
//    target = gtk_drawing_area_new ();
//
//    g_signal_connect (G_OBJECT (target), "drag-data-received",
//                      G_CALLBACK (passive_canvas_drag_data_received), NULL);
//    gtk_tool_palette_add_drag_dest (GTK_TOOL_PALETTE (palette), target,
//                                    GTK_DEST_DEFAULT_ALL,
//                                    GTK_TOOL_PALETTE_DRAG_ITEMS,
//                                    GDK_ACTION_COPY);
//
//
// CSS nodes
//
// GtkToolPalette has a single CSS node named toolpalette.
type ToolPalette interface {
	Container
	Buildable
	Orientable
	Scrollable

	// AddDragDest sets @palette as drag source (see
	// gtk_tool_palette_set_drag_source()) and sets @widget as a drag
	// destination for drags from @palette. See gtk_drag_dest_set().
	AddDragDest(widget Widget, flags DestDefaults, targets ToolPaletteDragTargets, actions gdk.DragAction)
	// Exclusive gets whether @group is exclusive or not. See
	// gtk_tool_palette_set_exclusive().
	Exclusive(group ToolItemGroup) bool
	// Expand gets whether group should be given extra space. See
	// gtk_tool_palette_set_expand().
	Expand(group ToolItemGroup) bool
	// GroupPosition gets the position of @group in @palette as index. See
	// gtk_tool_palette_set_group_position().
	GroupPosition(group ToolItemGroup) int
	// IconSize gets the size of icons in the tool palette. See
	// gtk_tool_palette_set_icon_size().
	IconSize() int
	// SetDragSource sets the tool palette as a drag source. Enables all groups
	// and items in the tool palette as drag sources on button 1 and button 3
	// press with copy and move actions. See gtk_drag_source_set().
	SetDragSource(targets ToolPaletteDragTargets)
	// SetExclusive sets whether the group should be exclusive or not. If an
	// exclusive group is expanded all other groups are collapsed.
	SetExclusive(group ToolItemGroup, exclusive bool)
	// SetExpand sets whether the group should be given extra space.
	SetExpand(group ToolItemGroup, expand bool)
	// SetGroupPosition sets the position of the group as an index of the tool
	// palette. If position is 0 the group will become the first child, if
	// position is -1 it will become the last child.
	SetGroupPosition(group ToolItemGroup, position int)
	// SetIconSize sets the size of icons in the tool palette.
	SetIconSize(iconSize int)
	// SetStyle sets the style (text, icons or both) of items in the tool
	// palette.
	SetStyle(style ToolbarStyle)
	// UnsetIconSize unsets the tool palette icon size set with
	// gtk_tool_palette_set_icon_size(), so that user preferences will be used
	// to determine the icon size.
	UnsetIconSize()
	// UnsetStyle unsets a toolbar style set with gtk_tool_palette_set_style(),
	// so that user preferences will be used to determine the toolbar style.
	UnsetStyle()
}

// toolPalette implements the ToolPalette interface.
type toolPalette struct {
	Container
	Buildable
	Orientable
	Scrollable
}

var _ ToolPalette = (*toolPalette)(nil)

// WrapToolPalette wraps a GObject to the right type. It is
// primarily used internally.
func WrapToolPalette(obj *externglib.Object) ToolPalette {
	return ToolPalette{
		Container:  WrapContainer(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
		Scrollable: WrapScrollable(obj),
	}
}

func marshalToolPalette(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToolPalette(obj), nil
}

// AddDragDest sets @palette as drag source (see
// gtk_tool_palette_set_drag_source()) and sets @widget as a drag
// destination for drags from @palette. See gtk_drag_dest_set().
func (p toolPalette) AddDragDest(widget Widget, flags DestDefaults, targets ToolPaletteDragTargets, actions gdk.DragAction) {
	var _arg0 *C.GtkToolPalette
	var _arg1 *C.GtkWidget
	var _arg2 C.GtkDestDefaults
	var _arg3 C.GtkToolPaletteDragTargets
	var _arg4 C.GdkDragAction

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = (C.GtkDestDefaults)(flags)
	_arg3 = (C.GtkToolPaletteDragTargets)(targets)
	_arg4 = (C.GdkDragAction)(actions)

	C.gtk_tool_palette_add_drag_dest(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// Exclusive gets whether @group is exclusive or not. See
// gtk_tool_palette_set_exclusive().
func (p toolPalette) Exclusive(group ToolItemGroup) bool {
	var _arg0 *C.GtkToolPalette
	var _arg1 *C.GtkToolItemGroup

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	var _cret C.gboolean

	_cret = C.gtk_tool_palette_get_exclusive(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Expand gets whether group should be given extra space. See
// gtk_tool_palette_set_expand().
func (p toolPalette) Expand(group ToolItemGroup) bool {
	var _arg0 *C.GtkToolPalette
	var _arg1 *C.GtkToolItemGroup

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	var _cret C.gboolean

	_cret = C.gtk_tool_palette_get_expand(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// GroupPosition gets the position of @group in @palette as index. See
// gtk_tool_palette_set_group_position().
func (p toolPalette) GroupPosition(group ToolItemGroup) int {
	var _arg0 *C.GtkToolPalette
	var _arg1 *C.GtkToolItemGroup

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	var _cret C.gint

	_cret = C.gtk_tool_palette_get_group_position(_arg0, _arg1)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// IconSize gets the size of icons in the tool palette. See
// gtk_tool_palette_set_icon_size().
func (p toolPalette) IconSize() int {
	var _arg0 *C.GtkToolPalette

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))

	var _cret C.GtkIconSize

	_cret = C.gtk_tool_palette_get_icon_size(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// SetDragSource sets the tool palette as a drag source. Enables all groups
// and items in the tool palette as drag sources on button 1 and button 3
// press with copy and move actions. See gtk_drag_source_set().
func (p toolPalette) SetDragSource(targets ToolPaletteDragTargets) {
	var _arg0 *C.GtkToolPalette
	var _arg1 C.GtkToolPaletteDragTargets

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	_arg1 = (C.GtkToolPaletteDragTargets)(targets)

	C.gtk_tool_palette_set_drag_source(_arg0, _arg1)
}

// SetExclusive sets whether the group should be exclusive or not. If an
// exclusive group is expanded all other groups are collapsed.
func (p toolPalette) SetExclusive(group ToolItemGroup, exclusive bool) {
	var _arg0 *C.GtkToolPalette
	var _arg1 *C.GtkToolItemGroup
	var _arg2 C.gboolean

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	if exclusive {
		_arg2 = C.gboolean(1)
	}

	C.gtk_tool_palette_set_exclusive(_arg0, _arg1, _arg2)
}

// SetExpand sets whether the group should be given extra space.
func (p toolPalette) SetExpand(group ToolItemGroup, expand bool) {
	var _arg0 *C.GtkToolPalette
	var _arg1 *C.GtkToolItemGroup
	var _arg2 C.gboolean

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	if expand {
		_arg2 = C.gboolean(1)
	}

	C.gtk_tool_palette_set_expand(_arg0, _arg1, _arg2)
}

// SetGroupPosition sets the position of the group as an index of the tool
// palette. If position is 0 the group will become the first child, if
// position is -1 it will become the last child.
func (p toolPalette) SetGroupPosition(group ToolItemGroup, position int) {
	var _arg0 *C.GtkToolPalette
	var _arg1 *C.GtkToolItemGroup
	var _arg2 C.gint

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg2 = C.gint(position)

	C.gtk_tool_palette_set_group_position(_arg0, _arg1, _arg2)
}

// SetIconSize sets the size of icons in the tool palette.
func (p toolPalette) SetIconSize(iconSize int) {
	var _arg0 *C.GtkToolPalette
	var _arg1 C.GtkIconSize

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	_arg1 = C.GtkIconSize(iconSize)

	C.gtk_tool_palette_set_icon_size(_arg0, _arg1)
}

// SetStyle sets the style (text, icons or both) of items in the tool
// palette.
func (p toolPalette) SetStyle(style ToolbarStyle) {
	var _arg0 *C.GtkToolPalette
	var _arg1 C.GtkToolbarStyle

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	_arg1 = (C.GtkToolbarStyle)(style)

	C.gtk_tool_palette_set_style(_arg0, _arg1)
}

// UnsetIconSize unsets the tool palette icon size set with
// gtk_tool_palette_set_icon_size(), so that user preferences will be used
// to determine the icon size.
func (p toolPalette) UnsetIconSize() {
	var _arg0 *C.GtkToolPalette

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))

	C.gtk_tool_palette_unset_icon_size(_arg0)
}

// UnsetStyle unsets a toolbar style set with gtk_tool_palette_set_style(),
// so that user preferences will be used to determine the toolbar style.
func (p toolPalette) UnsetStyle() {
	var _arg0 *C.GtkToolPalette

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))

	C.gtk_tool_palette_unset_style(_arg0)
}
