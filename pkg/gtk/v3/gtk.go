// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_movement_step_get_type()), F: marshalMovementStep},
		{T: externglib.Type(C.gtk_notebook_tab_get_type()), F: marshalNotebookTab},
		{T: externglib.Type(C.gtk_resize_mode_get_type()), F: marshalResizeMode},
		{T: externglib.Type(C.gtk_scroll_step_get_type()), F: marshalScrollStep},
		{T: externglib.Type(C.gtk_debug_flag_get_type()), F: marshalDebugFlag},
		{T: externglib.Type(C.gtk_entry_icon_accessible_get_type()), F: marshalEntryIconAccessibler},
	})
}

// The function returns the following values:
//
func BuilderErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_builder_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func CSSProviderErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_css_provider_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func IconThemeErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_icon_theme_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

type MovementStep C.gint

const (
	// MovementLogicalPositions: move forward or back by graphemes.
	MovementLogicalPositions MovementStep = iota
	// MovementVisualPositions: move left or right by graphemes.
	MovementVisualPositions
	// MovementWords: move forward or back by words.
	MovementWords
	// MovementDisplayLines: move up or down lines (wrapped lines).
	MovementDisplayLines
	// MovementDisplayLineEnds: move to either end of a line.
	MovementDisplayLineEnds
	// MovementParagraphs: move up or down paragraphs (newline-ended lines).
	MovementParagraphs
	// MovementParagraphEnds: move to either end of a paragraph.
	MovementParagraphEnds
	// MovementPages: move by pages.
	MovementPages
	// MovementBufferEnds: move to ends of the buffer.
	MovementBufferEnds
	// MovementHorizontalPages: move horizontally by pages.
	MovementHorizontalPages
)

func marshalMovementStep(p uintptr) (interface{}, error) {
	return MovementStep(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for MovementStep.
func (m MovementStep) String() string {
	switch m {
	case MovementLogicalPositions:
		return "LogicalPositions"
	case MovementVisualPositions:
		return "VisualPositions"
	case MovementWords:
		return "Words"
	case MovementDisplayLines:
		return "DisplayLines"
	case MovementDisplayLineEnds:
		return "DisplayLineEnds"
	case MovementParagraphs:
		return "Paragraphs"
	case MovementParagraphEnds:
		return "ParagraphEnds"
	case MovementPages:
		return "Pages"
	case MovementBufferEnds:
		return "BufferEnds"
	case MovementHorizontalPages:
		return "HorizontalPages"
	default:
		return fmt.Sprintf("MovementStep(%d)", m)
	}
}

type NotebookTab C.gint

const (
	NotebookTabFirst NotebookTab = iota
	NotebookTabLast
)

func marshalNotebookTab(p uintptr) (interface{}, error) {
	return NotebookTab(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for NotebookTab.
func (n NotebookTab) String() string {
	switch n {
	case NotebookTabFirst:
		return "First"
	case NotebookTabLast:
		return "Last"
	default:
		return fmt.Sprintf("NotebookTab(%d)", n)
	}
}

// The function returns the following values:
//
func RecentChooserErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_recent_chooser_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func RecentManagerErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_recent_manager_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

type ResizeMode C.gint

const (
	// ResizeParent pass resize request to the parent.
	ResizeParent ResizeMode = iota
	// ResizeQueue: queue resizes on this widget.
	ResizeQueue
	// ResizeImmediate: resize immediately. Deprecated.
	ResizeImmediate
)

func marshalResizeMode(p uintptr) (interface{}, error) {
	return ResizeMode(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ResizeMode.
func (r ResizeMode) String() string {
	switch r {
	case ResizeParent:
		return "Parent"
	case ResizeQueue:
		return "Queue"
	case ResizeImmediate:
		return "Immediate"
	default:
		return fmt.Sprintf("ResizeMode(%d)", r)
	}
}

type ScrollStep C.gint

const (
	// ScrollSteps: scroll in steps.
	ScrollSteps ScrollStep = iota
	// ScrollPages: scroll by pages.
	ScrollPages
	// ScrollEnds: scroll to ends.
	ScrollEnds
	// ScrollHorizontalSteps: scroll in horizontal steps.
	ScrollHorizontalSteps
	// ScrollHorizontalPages: scroll by horizontal pages.
	ScrollHorizontalPages
	// ScrollHorizontalEnds: scroll to the horizontal ends.
	ScrollHorizontalEnds
)

func marshalScrollStep(p uintptr) (interface{}, error) {
	return ScrollStep(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ScrollStep.
func (s ScrollStep) String() string {
	switch s {
	case ScrollSteps:
		return "Steps"
	case ScrollPages:
		return "Pages"
	case ScrollEnds:
		return "Ends"
	case ScrollHorizontalSteps:
		return "HorizontalSteps"
	case ScrollHorizontalPages:
		return "HorizontalPages"
	case ScrollHorizontalEnds:
		return "HorizontalEnds"
	default:
		return fmt.Sprintf("ScrollStep(%d)", s)
	}
}

type DebugFlag C.guint

const (
	DebugMisc         DebugFlag = 0b1
	DebugPlugsocket   DebugFlag = 0b10
	DebugText         DebugFlag = 0b100
	DebugTree         DebugFlag = 0b1000
	DebugUpdates      DebugFlag = 0b10000
	DebugKeybindings  DebugFlag = 0b100000
	DebugMultihead    DebugFlag = 0b1000000
	DebugModules      DebugFlag = 0b10000000
	DebugGeometry     DebugFlag = 0b100000000
	DebugIcontheme    DebugFlag = 0b1000000000
	DebugPrinting     DebugFlag = 0b10000000000
	DebugBuilder      DebugFlag = 0b100000000000
	DebugSizeRequest  DebugFlag = 0b1000000000000
	DebugNoCSSCache   DebugFlag = 0b10000000000000
	DebugBaselines    DebugFlag = 0b100000000000000
	DebugPixelCache   DebugFlag = 0b1000000000000000
	DebugNoPixelCache DebugFlag = 0b10000000000000000
	DebugInteractive  DebugFlag = 0b100000000000000000
	DebugTouchscreen  DebugFlag = 0b1000000000000000000
	DebugActions      DebugFlag = 0b10000000000000000000
	DebugResize       DebugFlag = 0b100000000000000000000
	DebugLayout       DebugFlag = 0b1000000000000000000000
)

func marshalDebugFlag(p uintptr) (interface{}, error) {
	return DebugFlag(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for DebugFlag.
func (d DebugFlag) String() string {
	if d == 0 {
		return "DebugFlag(0)"
	}

	var builder strings.Builder
	builder.Grow(256)

	for d != 0 {
		next := d & (d - 1)
		bit := d - next

		switch bit {
		case DebugMisc:
			builder.WriteString("Misc|")
		case DebugPlugsocket:
			builder.WriteString("Plugsocket|")
		case DebugText:
			builder.WriteString("Text|")
		case DebugTree:
			builder.WriteString("Tree|")
		case DebugUpdates:
			builder.WriteString("Updates|")
		case DebugKeybindings:
			builder.WriteString("Keybindings|")
		case DebugMultihead:
			builder.WriteString("Multihead|")
		case DebugModules:
			builder.WriteString("Modules|")
		case DebugGeometry:
			builder.WriteString("Geometry|")
		case DebugIcontheme:
			builder.WriteString("Icontheme|")
		case DebugPrinting:
			builder.WriteString("Printing|")
		case DebugBuilder:
			builder.WriteString("Builder|")
		case DebugSizeRequest:
			builder.WriteString("SizeRequest|")
		case DebugNoCSSCache:
			builder.WriteString("NoCSSCache|")
		case DebugBaselines:
			builder.WriteString("Baselines|")
		case DebugPixelCache:
			builder.WriteString("PixelCache|")
		case DebugNoPixelCache:
			builder.WriteString("NoPixelCache|")
		case DebugInteractive:
			builder.WriteString("Interactive|")
		case DebugTouchscreen:
			builder.WriteString("Touchscreen|")
		case DebugActions:
			builder.WriteString("Actions|")
		case DebugResize:
			builder.WriteString("Resize|")
		case DebugLayout:
			builder.WriteString("Layout|")
		default:
			builder.WriteString(fmt.Sprintf("DebugFlag(0b%b)|", bit))
		}

		d = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if d contains other.
func (d DebugFlag) Has(other DebugFlag) bool {
	return (d & other) == other
}

type EntryIconAccessible struct {
	_ [0]func() // equal guard
	atk.ObjectClass

	*externglib.Object
	atk.Action
	atk.Component
}

var (
	_ externglib.Objector = (*EntryIconAccessible)(nil)
)

func wrapEntryIconAccessible(obj *externglib.Object) *EntryIconAccessible {
	return &EntryIconAccessible{
		ObjectClass: atk.ObjectClass{
			Object: obj,
		},
		Object: obj,
		Action: atk.Action{
			Object: obj,
		},
		Component: atk.Component{
			Object: obj,
		},
	}
}

func marshalEntryIconAccessibler(p uintptr) (interface{}, error) {
	return wrapEntryIconAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Init binds to the gtk_init() function. Argument parsing is not
// supported.
func Init() {
	C.gtk_init(nil, nil)
}
