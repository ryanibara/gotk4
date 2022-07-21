// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern GtkIconSize _gotk4_gtk3_ToolShellIface_get_icon_size(GtkToolShell*);
// extern GtkOrientation _gotk4_gtk3_ToolShellIface_get_orientation(GtkToolShell*);
// extern GtkOrientation _gotk4_gtk3_ToolShellIface_get_text_orientation(GtkToolShell*);
// extern GtkReliefStyle _gotk4_gtk3_ToolShellIface_get_relief_style(GtkToolShell*);
// extern GtkSizeGroup* _gotk4_gtk3_ToolShellIface_get_text_size_group(GtkToolShell*);
// extern GtkToolbarStyle _gotk4_gtk3_ToolShellIface_get_style(GtkToolShell*);
// extern PangoEllipsizeMode _gotk4_gtk3_ToolShellIface_get_ellipsize_mode(GtkToolShell*);
// extern gfloat _gotk4_gtk3_ToolShellIface_get_text_alignment(GtkToolShell*);
// extern void _gotk4_gtk3_ToolShellIface_rebuild_menu(GtkToolShell*);
import "C"

// GTypeToolShell returns the GType for the type ToolShell.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeToolShell() coreglib.Type {
	gtype := coreglib.Type(C.gtk_tool_shell_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalToolShell)
	return gtype
}

// ToolShellOverrider contains methods that are overridable.
type ToolShellOverrider interface {
	// EllipsizeMode retrieves the current ellipsize mode for the tool shell.
	// Tool items must not call this function directly, but rely on
	// gtk_tool_item_get_ellipsize_mode() instead.
	//
	// The function returns the following values:
	//
	//    - ellipsizeMode: current ellipsize mode of shell.
	//
	EllipsizeMode() pango.EllipsizeMode
	// The function returns the following values:
	//
	IconSize() IconSize
	// Orientation retrieves the current orientation for the tool shell. Tool
	// items must not call this function directly, but rely on
	// gtk_tool_item_get_orientation() instead.
	//
	// The function returns the following values:
	//
	//    - orientation: current orientation of shell.
	//
	Orientation() Orientation
	// ReliefStyle returns the relief style of buttons on shell. Tool items must
	// not call this function directly, but rely on
	// gtk_tool_item_get_relief_style() instead.
	//
	// The function returns the following values:
	//
	//    - reliefStyle: relief style of buttons on shell.
	//
	ReliefStyle() ReliefStyle
	// Style retrieves whether the tool shell has text, icons, or both. Tool
	// items must not call this function directly, but rely on
	// gtk_tool_item_get_toolbar_style() instead.
	//
	// The function returns the following values:
	//
	//    - toolbarStyle: current style of shell.
	//
	Style() ToolbarStyle
	// TextAlignment retrieves the current text alignment for the tool shell.
	// Tool items must not call this function directly, but rely on
	// gtk_tool_item_get_text_alignment() instead.
	//
	// The function returns the following values:
	//
	//    - gfloat: current text alignment of shell.
	//
	TextAlignment() float32
	// TextOrientation retrieves the current text orientation for the tool
	// shell. Tool items must not call this function directly, but rely on
	// gtk_tool_item_get_text_orientation() instead.
	//
	// The function returns the following values:
	//
	//    - orientation: current text orientation of shell.
	//
	TextOrientation() Orientation
	// TextSizeGroup retrieves the current text size group for the tool shell.
	// Tool items must not call this function directly, but rely on
	// gtk_tool_item_get_text_size_group() instead.
	//
	// The function returns the following values:
	//
	//    - sizeGroup: current text size group of shell.
	//
	TextSizeGroup() *SizeGroup
	// RebuildMenu: calling this function signals the tool shell that the
	// overflow menu item for tool items have changed. If there is an overflow
	// menu and if it is visible when this function it called, the menu will be
	// rebuilt.
	//
	// Tool items must not call this function directly, but rely on
	// gtk_tool_item_rebuild_menu() instead.
	RebuildMenu()
}

// ToolShell interface allows container widgets to provide additional
// information when embedding ToolItem widgets.
//
// ToolShell wraps an interface. This means the user can get the
// underlying type by calling Cast().
type ToolShell struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*ToolShell)(nil)
)

// ToolSheller describes ToolShell's interface methods.
type ToolSheller interface {
	coreglib.Objector

	// EllipsizeMode retrieves the current ellipsize mode for the tool shell.
	EllipsizeMode() pango.EllipsizeMode
	// IconSize retrieves the icon size for the tool shell.
	IconSize() int32
	// Orientation retrieves the current orientation for the tool shell.
	Orientation() Orientation
	// ReliefStyle returns the relief style of buttons on shell.
	ReliefStyle() ReliefStyle
	// Style retrieves whether the tool shell has text, icons, or both.
	Style() ToolbarStyle
	// TextAlignment retrieves the current text alignment for the tool shell.
	TextAlignment() float32
	// TextOrientation retrieves the current text orientation for the tool
	// shell.
	TextOrientation() Orientation
	// TextSizeGroup retrieves the current text size group for the tool shell.
	TextSizeGroup() *SizeGroup
	// RebuildMenu: calling this function signals the tool shell that the
	// overflow menu item for tool items have changed.
	RebuildMenu()
}

var _ ToolSheller = (*ToolShell)(nil)

func ifaceInitToolSheller(gifacePtr, data C.gpointer) {
	iface := (*C.GtkToolShellIface)(unsafe.Pointer(gifacePtr))
	iface.get_ellipsize_mode = (*[0]byte)(C._gotk4_gtk3_ToolShellIface_get_ellipsize_mode)
	iface.get_icon_size = (*[0]byte)(C._gotk4_gtk3_ToolShellIface_get_icon_size)
	iface.get_orientation = (*[0]byte)(C._gotk4_gtk3_ToolShellIface_get_orientation)
	iface.get_relief_style = (*[0]byte)(C._gotk4_gtk3_ToolShellIface_get_relief_style)
	iface.get_style = (*[0]byte)(C._gotk4_gtk3_ToolShellIface_get_style)
	iface.get_text_alignment = (*[0]byte)(C._gotk4_gtk3_ToolShellIface_get_text_alignment)
	iface.get_text_orientation = (*[0]byte)(C._gotk4_gtk3_ToolShellIface_get_text_orientation)
	iface.get_text_size_group = (*[0]byte)(C._gotk4_gtk3_ToolShellIface_get_text_size_group)
	iface.rebuild_menu = (*[0]byte)(C._gotk4_gtk3_ToolShellIface_rebuild_menu)
}

//export _gotk4_gtk3_ToolShellIface_get_ellipsize_mode
func _gotk4_gtk3_ToolShellIface_get_ellipsize_mode(arg0 *C.GtkToolShell) (cret C.PangoEllipsizeMode) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ToolShellOverrider)

	ellipsizeMode := iface.EllipsizeMode()

	cret = C.PangoEllipsizeMode(ellipsizeMode)

	return cret
}

//export _gotk4_gtk3_ToolShellIface_get_icon_size
func _gotk4_gtk3_ToolShellIface_get_icon_size(arg0 *C.GtkToolShell) (cret C.GtkIconSize) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ToolShellOverrider)

	iconSize := iface.IconSize()

	cret = C.GtkIconSize(iconSize)

	return cret
}

//export _gotk4_gtk3_ToolShellIface_get_orientation
func _gotk4_gtk3_ToolShellIface_get_orientation(arg0 *C.GtkToolShell) (cret C.GtkOrientation) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ToolShellOverrider)

	orientation := iface.Orientation()

	cret = C.GtkOrientation(orientation)

	return cret
}

//export _gotk4_gtk3_ToolShellIface_get_relief_style
func _gotk4_gtk3_ToolShellIface_get_relief_style(arg0 *C.GtkToolShell) (cret C.GtkReliefStyle) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ToolShellOverrider)

	reliefStyle := iface.ReliefStyle()

	cret = C.GtkReliefStyle(reliefStyle)

	return cret
}

//export _gotk4_gtk3_ToolShellIface_get_style
func _gotk4_gtk3_ToolShellIface_get_style(arg0 *C.GtkToolShell) (cret C.GtkToolbarStyle) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ToolShellOverrider)

	toolbarStyle := iface.Style()

	cret = C.GtkToolbarStyle(toolbarStyle)

	return cret
}

//export _gotk4_gtk3_ToolShellIface_get_text_alignment
func _gotk4_gtk3_ToolShellIface_get_text_alignment(arg0 *C.GtkToolShell) (cret C.gfloat) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ToolShellOverrider)

	gfloat := iface.TextAlignment()

	cret = C.gfloat(gfloat)

	return cret
}

//export _gotk4_gtk3_ToolShellIface_get_text_orientation
func _gotk4_gtk3_ToolShellIface_get_text_orientation(arg0 *C.GtkToolShell) (cret C.GtkOrientation) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ToolShellOverrider)

	orientation := iface.TextOrientation()

	cret = C.GtkOrientation(orientation)

	return cret
}

//export _gotk4_gtk3_ToolShellIface_get_text_size_group
func _gotk4_gtk3_ToolShellIface_get_text_size_group(arg0 *C.GtkToolShell) (cret *C.GtkSizeGroup) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ToolShellOverrider)

	sizeGroup := iface.TextSizeGroup()

	cret = (*C.GtkSizeGroup)(unsafe.Pointer(coreglib.InternObject(sizeGroup).Native()))

	return cret
}

//export _gotk4_gtk3_ToolShellIface_rebuild_menu
func _gotk4_gtk3_ToolShellIface_rebuild_menu(arg0 *C.GtkToolShell) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ToolShellOverrider)

	iface.RebuildMenu()
}

func wrapToolShell(obj *coreglib.Object) *ToolShell {
	return &ToolShell{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
		},
	}
}

func marshalToolShell(p uintptr) (interface{}, error) {
	return wrapToolShell(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// EllipsizeMode retrieves the current ellipsize mode for the tool shell. Tool
// items must not call this function directly, but rely on
// gtk_tool_item_get_ellipsize_mode() instead.
//
// The function returns the following values:
//
//    - ellipsizeMode: current ellipsize mode of shell.
//
func (shell *ToolShell) EllipsizeMode() pango.EllipsizeMode {
	var _arg0 *C.GtkToolShell      // out
	var _cret C.PangoEllipsizeMode // in

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(coreglib.InternObject(shell).Native()))

	_cret = C.gtk_tool_shell_get_ellipsize_mode(_arg0)
	runtime.KeepAlive(shell)

	var _ellipsizeMode pango.EllipsizeMode // out

	_ellipsizeMode = pango.EllipsizeMode(_cret)

	return _ellipsizeMode
}

// IconSize retrieves the icon size for the tool shell. Tool items must not call
// this function directly, but rely on gtk_tool_item_get_icon_size() instead.
//
// The function returns the following values:
//
//    - gint: current size (IconSize) for icons of shell.
//
func (shell *ToolShell) IconSize() int32 {
	var _arg0 *C.GtkToolShell // out
	var _cret C.GtkIconSize   // in

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(coreglib.InternObject(shell).Native()))

	_cret = C.gtk_tool_shell_get_icon_size(_arg0)
	runtime.KeepAlive(shell)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// Orientation retrieves the current orientation for the tool shell. Tool items
// must not call this function directly, but rely on
// gtk_tool_item_get_orientation() instead.
//
// The function returns the following values:
//
//    - orientation: current orientation of shell.
//
func (shell *ToolShell) Orientation() Orientation {
	var _arg0 *C.GtkToolShell  // out
	var _cret C.GtkOrientation // in

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(coreglib.InternObject(shell).Native()))

	_cret = C.gtk_tool_shell_get_orientation(_arg0)
	runtime.KeepAlive(shell)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

// ReliefStyle returns the relief style of buttons on shell. Tool items must not
// call this function directly, but rely on gtk_tool_item_get_relief_style()
// instead.
//
// The function returns the following values:
//
//    - reliefStyle: relief style of buttons on shell.
//
func (shell *ToolShell) ReliefStyle() ReliefStyle {
	var _arg0 *C.GtkToolShell  // out
	var _cret C.GtkReliefStyle // in

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(coreglib.InternObject(shell).Native()))

	_cret = C.gtk_tool_shell_get_relief_style(_arg0)
	runtime.KeepAlive(shell)

	var _reliefStyle ReliefStyle // out

	_reliefStyle = ReliefStyle(_cret)

	return _reliefStyle
}

// Style retrieves whether the tool shell has text, icons, or both. Tool items
// must not call this function directly, but rely on
// gtk_tool_item_get_toolbar_style() instead.
//
// The function returns the following values:
//
//    - toolbarStyle: current style of shell.
//
func (shell *ToolShell) Style() ToolbarStyle {
	var _arg0 *C.GtkToolShell   // out
	var _cret C.GtkToolbarStyle // in

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(coreglib.InternObject(shell).Native()))

	_cret = C.gtk_tool_shell_get_style(_arg0)
	runtime.KeepAlive(shell)

	var _toolbarStyle ToolbarStyle // out

	_toolbarStyle = ToolbarStyle(_cret)

	return _toolbarStyle
}

// TextAlignment retrieves the current text alignment for the tool shell. Tool
// items must not call this function directly, but rely on
// gtk_tool_item_get_text_alignment() instead.
//
// The function returns the following values:
//
//    - gfloat: current text alignment of shell.
//
func (shell *ToolShell) TextAlignment() float32 {
	var _arg0 *C.GtkToolShell // out
	var _cret C.gfloat        // in

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(coreglib.InternObject(shell).Native()))

	_cret = C.gtk_tool_shell_get_text_alignment(_arg0)
	runtime.KeepAlive(shell)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// TextOrientation retrieves the current text orientation for the tool shell.
// Tool items must not call this function directly, but rely on
// gtk_tool_item_get_text_orientation() instead.
//
// The function returns the following values:
//
//    - orientation: current text orientation of shell.
//
func (shell *ToolShell) TextOrientation() Orientation {
	var _arg0 *C.GtkToolShell  // out
	var _cret C.GtkOrientation // in

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(coreglib.InternObject(shell).Native()))

	_cret = C.gtk_tool_shell_get_text_orientation(_arg0)
	runtime.KeepAlive(shell)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

// TextSizeGroup retrieves the current text size group for the tool shell. Tool
// items must not call this function directly, but rely on
// gtk_tool_item_get_text_size_group() instead.
//
// The function returns the following values:
//
//    - sizeGroup: current text size group of shell.
//
func (shell *ToolShell) TextSizeGroup() *SizeGroup {
	var _arg0 *C.GtkToolShell // out
	var _cret *C.GtkSizeGroup // in

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(coreglib.InternObject(shell).Native()))

	_cret = C.gtk_tool_shell_get_text_size_group(_arg0)
	runtime.KeepAlive(shell)

	var _sizeGroup *SizeGroup // out

	_sizeGroup = wrapSizeGroup(coreglib.Take(unsafe.Pointer(_cret)))

	return _sizeGroup
}

// RebuildMenu: calling this function signals the tool shell that the overflow
// menu item for tool items have changed. If there is an overflow menu and if it
// is visible when this function it called, the menu will be rebuilt.
//
// Tool items must not call this function directly, but rely on
// gtk_tool_item_rebuild_menu() instead.
func (shell *ToolShell) RebuildMenu() {
	var _arg0 *C.GtkToolShell // out

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(coreglib.InternObject(shell).Native()))

	C.gtk_tool_shell_rebuild_menu(_arg0)
	runtime.KeepAlive(shell)
}
