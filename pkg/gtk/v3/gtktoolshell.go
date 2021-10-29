// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tool_shell_get_type()), F: marshalToolSheller},
	})
}

// ToolShellOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ToolShellOverrider interface {
	// EllipsizeMode retrieves the current ellipsize mode for the tool shell.
	// Tool items must not call this function directly, but rely on
	// gtk_tool_item_get_ellipsize_mode() instead.
	EllipsizeMode() pango.EllipsizeMode
	IconSize() IconSize
	// Orientation retrieves the current orientation for the tool shell. Tool
	// items must not call this function directly, but rely on
	// gtk_tool_item_get_orientation() instead.
	Orientation() Orientation
	// ReliefStyle returns the relief style of buttons on shell. Tool items must
	// not call this function directly, but rely on
	// gtk_tool_item_get_relief_style() instead.
	ReliefStyle() ReliefStyle
	// Style retrieves whether the tool shell has text, icons, or both. Tool
	// items must not call this function directly, but rely on
	// gtk_tool_item_get_toolbar_style() instead.
	Style() ToolbarStyle
	// TextAlignment retrieves the current text alignment for the tool shell.
	// Tool items must not call this function directly, but rely on
	// gtk_tool_item_get_text_alignment() instead.
	TextAlignment() float32
	// TextOrientation retrieves the current text orientation for the tool
	// shell. Tool items must not call this function directly, but rely on
	// gtk_tool_item_get_text_orientation() instead.
	TextOrientation() Orientation
	// TextSizeGroup retrieves the current text size group for the tool shell.
	// Tool items must not call this function directly, but rely on
	// gtk_tool_item_get_text_size_group() instead.
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
type ToolShell struct {
	Widget
}

var (
	_ Widgetter = (*ToolShell)(nil)
)

// ToolSheller describes ToolShell's interface methods.
type ToolSheller interface {
	externglib.Objector

	// EllipsizeMode retrieves the current ellipsize mode for the tool shell.
	EllipsizeMode() pango.EllipsizeMode
	// IconSize retrieves the icon size for the tool shell.
	IconSize() int
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

func wrapToolShell(obj *externglib.Object) *ToolShell {
	return &ToolShell{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalToolSheller(p uintptr) (interface{}, error) {
	return wrapToolShell(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// EllipsizeMode retrieves the current ellipsize mode for the tool shell. Tool
// items must not call this function directly, but rely on
// gtk_tool_item_get_ellipsize_mode() instead.
func (shell *ToolShell) EllipsizeMode() pango.EllipsizeMode {
	var _arg0 *C.GtkToolShell      // out
	var _cret C.PangoEllipsizeMode // in

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(shell.Native()))

	_cret = C.gtk_tool_shell_get_ellipsize_mode(_arg0)
	runtime.KeepAlive(shell)

	var _ellipsizeMode pango.EllipsizeMode // out

	_ellipsizeMode = pango.EllipsizeMode(_cret)

	return _ellipsizeMode
}

// IconSize retrieves the icon size for the tool shell. Tool items must not call
// this function directly, but rely on gtk_tool_item_get_icon_size() instead.
func (shell *ToolShell) IconSize() int {
	var _arg0 *C.GtkToolShell // out
	var _cret C.GtkIconSize   // in

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(shell.Native()))

	_cret = C.gtk_tool_shell_get_icon_size(_arg0)
	runtime.KeepAlive(shell)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Orientation retrieves the current orientation for the tool shell. Tool items
// must not call this function directly, but rely on
// gtk_tool_item_get_orientation() instead.
func (shell *ToolShell) Orientation() Orientation {
	var _arg0 *C.GtkToolShell  // out
	var _cret C.GtkOrientation // in

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(shell.Native()))

	_cret = C.gtk_tool_shell_get_orientation(_arg0)
	runtime.KeepAlive(shell)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

// ReliefStyle returns the relief style of buttons on shell. Tool items must not
// call this function directly, but rely on gtk_tool_item_get_relief_style()
// instead.
func (shell *ToolShell) ReliefStyle() ReliefStyle {
	var _arg0 *C.GtkToolShell  // out
	var _cret C.GtkReliefStyle // in

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(shell.Native()))

	_cret = C.gtk_tool_shell_get_relief_style(_arg0)
	runtime.KeepAlive(shell)

	var _reliefStyle ReliefStyle // out

	_reliefStyle = ReliefStyle(_cret)

	return _reliefStyle
}

// Style retrieves whether the tool shell has text, icons, or both. Tool items
// must not call this function directly, but rely on
// gtk_tool_item_get_toolbar_style() instead.
func (shell *ToolShell) Style() ToolbarStyle {
	var _arg0 *C.GtkToolShell   // out
	var _cret C.GtkToolbarStyle // in

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(shell.Native()))

	_cret = C.gtk_tool_shell_get_style(_arg0)
	runtime.KeepAlive(shell)

	var _toolbarStyle ToolbarStyle // out

	_toolbarStyle = ToolbarStyle(_cret)

	return _toolbarStyle
}

// TextAlignment retrieves the current text alignment for the tool shell. Tool
// items must not call this function directly, but rely on
// gtk_tool_item_get_text_alignment() instead.
func (shell *ToolShell) TextAlignment() float32 {
	var _arg0 *C.GtkToolShell // out
	var _cret C.gfloat        // in

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(shell.Native()))

	_cret = C.gtk_tool_shell_get_text_alignment(_arg0)
	runtime.KeepAlive(shell)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// TextOrientation retrieves the current text orientation for the tool shell.
// Tool items must not call this function directly, but rely on
// gtk_tool_item_get_text_orientation() instead.
func (shell *ToolShell) TextOrientation() Orientation {
	var _arg0 *C.GtkToolShell  // out
	var _cret C.GtkOrientation // in

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(shell.Native()))

	_cret = C.gtk_tool_shell_get_text_orientation(_arg0)
	runtime.KeepAlive(shell)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

// TextSizeGroup retrieves the current text size group for the tool shell. Tool
// items must not call this function directly, but rely on
// gtk_tool_item_get_text_size_group() instead.
func (shell *ToolShell) TextSizeGroup() *SizeGroup {
	var _arg0 *C.GtkToolShell // out
	var _cret *C.GtkSizeGroup // in

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(shell.Native()))

	_cret = C.gtk_tool_shell_get_text_size_group(_arg0)
	runtime.KeepAlive(shell)

	var _sizeGroup *SizeGroup // out

	_sizeGroup = wrapSizeGroup(externglib.Take(unsafe.Pointer(_cret)))

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

	_arg0 = (*C.GtkToolShell)(unsafe.Pointer(shell.Native()))

	C.gtk_tool_shell_rebuild_menu(_arg0)
	runtime.KeepAlive(shell)
}
