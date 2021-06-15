// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_stack_transition_type_get_type()), F: marshalStackTransitionType},
		{T: externglib.Type(C.gtk_stack_get_type()), F: marshalStack},
		{T: externglib.Type(C.gtk_stack_page_get_type()), F: marshalStackPage},
	})
}

// StackTransitionType: possible transitions between pages in a `GtkStack`
// widget.
//
// New values may be added to this enumeration over time.
type StackTransitionType int

const (
	// StackTransitionTypeNone: no transition
	StackTransitionTypeNone StackTransitionType = 0
	// StackTransitionTypeCrossfade: a cross-fade
	StackTransitionTypeCrossfade StackTransitionType = 1
	// StackTransitionTypeSlideRight: slide from left to right
	StackTransitionTypeSlideRight StackTransitionType = 2
	// StackTransitionTypeSlideLeft: slide from right to left
	StackTransitionTypeSlideLeft StackTransitionType = 3
	// StackTransitionTypeSlideUp: slide from bottom up
	StackTransitionTypeSlideUp StackTransitionType = 4
	// StackTransitionTypeSlideDown: slide from top down
	StackTransitionTypeSlideDown StackTransitionType = 5
	// StackTransitionTypeSlideLeftRight: slide from left or right according to
	// the children order
	StackTransitionTypeSlideLeftRight StackTransitionType = 6
	// StackTransitionTypeSlideUpDown: slide from top down or bottom up
	// according to the order
	StackTransitionTypeSlideUpDown StackTransitionType = 7
	// StackTransitionTypeOverUp: cover the old page by sliding up
	StackTransitionTypeOverUp StackTransitionType = 8
	// StackTransitionTypeOverDown: cover the old page by sliding down
	StackTransitionTypeOverDown StackTransitionType = 9
	// StackTransitionTypeOverLeft: cover the old page by sliding to the left
	StackTransitionTypeOverLeft StackTransitionType = 10
	// StackTransitionTypeOverRight: cover the old page by sliding to the right
	StackTransitionTypeOverRight StackTransitionType = 11
	// StackTransitionTypeUnderUp: uncover the new page by sliding up
	StackTransitionTypeUnderUp StackTransitionType = 12
	// StackTransitionTypeUnderDown: uncover the new page by sliding down
	StackTransitionTypeUnderDown StackTransitionType = 13
	// StackTransitionTypeUnderLeft: uncover the new page by sliding to the left
	StackTransitionTypeUnderLeft StackTransitionType = 14
	// StackTransitionTypeUnderRight: uncover the new page by sliding to the
	// right
	StackTransitionTypeUnderRight StackTransitionType = 15
	// StackTransitionTypeOverUpDown: cover the old page sliding up or uncover
	// the new page sliding down, according to order
	StackTransitionTypeOverUpDown StackTransitionType = 16
	// StackTransitionTypeOverDownUp: cover the old page sliding down or uncover
	// the new page sliding up, according to order
	StackTransitionTypeOverDownUp StackTransitionType = 17
	// StackTransitionTypeOverLeftRight: cover the old page sliding left or
	// uncover the new page sliding right, according to order
	StackTransitionTypeOverLeftRight StackTransitionType = 18
	// StackTransitionTypeOverRightLeft: cover the old page sliding right or
	// uncover the new page sliding left, according to order
	StackTransitionTypeOverRightLeft StackTransitionType = 19
	// StackTransitionTypeRotateLeft: pretend the pages are sides of a cube and
	// rotate that cube to the left
	StackTransitionTypeRotateLeft StackTransitionType = 20
	// StackTransitionTypeRotateRight: pretend the pages are sides of a cube and
	// rotate that cube to the right
	StackTransitionTypeRotateRight StackTransitionType = 21
	// StackTransitionTypeRotateLeftRight: pretend the pages are sides of a cube
	// and rotate that cube to the left or right according to the children order
	StackTransitionTypeRotateLeftRight StackTransitionType = 22
)

func marshalStackTransitionType(p uintptr) (interface{}, error) {
	return StackTransitionType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Stack: `GtkStack` is a container which only shows one of its children at a
// time.
//
// In contrast to `GtkNotebook`, `GtkStack` does not provide a means for users
// to change the visible child. Instead, a separate widget such as
// [class@Gtk.StackSwitcher] or [class@Gtk.StackSidebar] can be used with
// `GtkStack` to provide this functionality.
//
// Transitions between pages can be animated as slides or fades. This can be
// controlled with [method@Gtk.Stack.set_transition_type]. These animations
// respect the [property@Gtk.Settings:gtk-enable-animations] setting.
//
// `GtkStack` maintains a [class@Gtk.StackPage] object for each added child,
// which holds additional per-child properties. You obtain the `GtkStackPage`
// for a child with [method@Gtk.Stack.get_page] and you can obtain a
// `GtkSelectionModel` containing all the pages with
// [method@Gtk.Stack.get_pages].
//
//
// GtkStack as GtkBuildable
//
// To set child-specific properties in a .ui file, create `GtkStackPage` objects
// explicitly, and set the child widget as a property on it:
//
// “`xml <object class="GtkStack" id="stack"> <child> <object
// class="GtkStackPage"> <property name="name">page1</property> <property
// name="title">In the beginning…</property> <property name="child"> <object
// class="GtkLabel"> <property name="label">It was dark</property> </object>
// </property> </object> </child> “`
//
//
// CSS nodes
//
// `GtkStack` has a single CSS node named stack.
//
//
// Accessibility
//
// `GtkStack` uses the GTK_ACCESSIBLE_ROLE_TAB_PANEL for the stack pages, which
// are the accessible parent objects of the child widgets.
type Stack interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget

	// AddChild adds a child to @stack.
	AddChild(child Widget) StackPage
	// AddNamed adds a child to @stack.
	//
	// The child is identified by the @name.
	AddNamed(child Widget, name string) StackPage
	// AddTitled adds a child to @stack.
	//
	// The child is identified by the @name. The @title will be used by
	// `GtkStackSwitcher` to represent @child in a tab bar, so it should be
	// short.
	AddTitled(child Widget, name string, title string) StackPage
	// ChildByName finds the child with the name given as the argument.
	//
	// Returns nil if there is no child with this name.
	ChildByName(name string) Widget
	// Hhomogeneous gets whether @stack is horizontally homogeneous.
	Hhomogeneous() bool
	// InterpolateSize returns whether the Stack is set up to interpolate
	// between the sizes of children on page switch.
	InterpolateSize() bool
	// Page returns the `GtkStackPage` object for @child.
	Page(child Widget) StackPage
	// Pages returns a `GListModel` that contains the pages of the stack.
	//
	// This can be used to keep an up-to-date view. The model also implements
	// [iface@Gtk.SelectionModel] and can be used to track and modify the
	// visible page.
	Pages() SelectionModel
	// TransitionDuration returns the amount of time (in milliseconds) that
	// transitions between pages in @stack will take.
	TransitionDuration() uint
	// TransitionRunning returns whether the @stack is currently in a transition
	// from one page to another.
	TransitionRunning() bool
	// TransitionType gets the type of animation that will be used for
	// transitions between pages in @stack.
	TransitionType() StackTransitionType
	// Vhomogeneous gets whether @stack is vertically homogeneous.
	Vhomogeneous() bool
	// VisibleChild gets the currently visible child of @stack.
	//
	// Returns nil if there are no visible children.
	VisibleChild() Widget
	// VisibleChildName returns the name of the currently visible child of
	// @stack.
	//
	// Returns nil if there is no visible child.
	VisibleChildName() string
	// Remove removes a child widget from @stack.
	Remove(child Widget)
	// SetHhomogeneous sets the `GtkStack` to be horizontally homogeneous or
	// not.
	//
	// If it is homogeneous, the `GtkStack` will request the same width for all
	// its children. If it isn't, the stack may change width when a different
	// child becomes visible.
	SetHhomogeneous(hhomogeneous bool)
	// SetInterpolateSize sets whether or not @stack will interpolate its size
	// when changing the visible child.
	//
	// If the [property@Gtk.Stack:interpolate-size] property is set to true,
	// @stack will interpolate its size between the current one and the one
	// it'll take after changing the visible child, according to the set
	// transition duration.
	SetInterpolateSize(interpolateSize bool)
	// SetTransitionDuration sets the duration that transitions between pages in
	// @stack will take.
	SetTransitionDuration(duration uint)
	// SetTransitionType sets the type of animation that will be used for
	// transitions between pages in @stack.
	//
	// Available types include various kinds of fades and slides.
	//
	// The transition type can be changed without problems at runtime, so it is
	// possible to change the animation based on the page that is about to
	// become current.
	SetTransitionType(transition StackTransitionType)
	// SetVhomogeneous sets the Stack to be vertically homogeneous or not.
	//
	// If it is homogeneous, the `GtkStack` will request the same height for all
	// its children. If it isn't, the stack may change height when a different
	// child becomes visible.
	SetVhomogeneous(vhomogeneous bool)
	// SetVisibleChild makes @child the visible child of @stack.
	//
	// If @child is different from the currently visible child, the transition
	// between the two will be animated with the current transition type of
	// @stack.
	//
	// Note that the @child widget has to be visible itself (see
	// [method@Gtk.Widget.show]) in order to become the visible child of @stack.
	SetVisibleChild(child Widget)
	// SetVisibleChildFull makes the child with the given name visible.
	//
	// Note that the child widget has to be visible itself (see
	// [method@Gtk.Widget.show]) in order to become the visible child of @stack.
	SetVisibleChildFull(name string, transition StackTransitionType)
	// SetVisibleChildName makes the child with the given name visible.
	//
	// If @child is different from the currently visible child, the transition
	// between the two will be animated with the current transition type of
	// @stack.
	//
	// Note that the child widget has to be visible itself (see
	// [method@Gtk.Widget.show]) in order to become the visible child of @stack.
	SetVisibleChildName(name string)
}

// stack implements the Stack class.
type stack struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ Stack = (*stack)(nil)

// WrapStack wraps a GObject to the right type. It is
// primarily used internally.
func WrapStack(obj *externglib.Object) Stack {
	return stack{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalStack(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStack(obj), nil
}

// NewStack constructs a class Stack.
func NewStack() Stack {
	var _cret C.GtkStack // in

	_cret = C.gtk_stack_new()

	var _stack Stack // out

	_stack = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Stack)

	return _stack
}

// AddChild adds a child to @stack.
func (s stack) AddChild(child Widget) StackPage {
	var _arg0 *C.GtkStack     // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.GtkStackPage // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_stack_add_child(_arg0, _arg1)

	var _stackPage StackPage // out

	_stackPage = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(StackPage)

	return _stackPage
}

// AddNamed adds a child to @stack.
//
// The child is identified by the @name.
func (s stack) AddNamed(child Widget, name string) StackPage {
	var _arg0 *C.GtkStack     // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.char         // out
	var _cret *C.GtkStackPage // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_stack_add_named(_arg0, _arg1, _arg2)

	var _stackPage StackPage // out

	_stackPage = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(StackPage)

	return _stackPage
}

// AddTitled adds a child to @stack.
//
// The child is identified by the @name. The @title will be used by
// `GtkStackSwitcher` to represent @child in a tab bar, so it should be
// short.
func (s stack) AddTitled(child Widget, name string, title string) StackPage {
	var _arg0 *C.GtkStack     // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.char         // out
	var _arg3 *C.char         // out
	var _cret *C.GtkStackPage // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.gtk_stack_add_titled(_arg0, _arg1, _arg2, _arg3)

	var _stackPage StackPage // out

	_stackPage = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(StackPage)

	return _stackPage
}

// ChildByName finds the child with the name given as the argument.
//
// Returns nil if there is no child with this name.
func (s stack) ChildByName(name string) Widget {
	var _arg0 *C.GtkStack  // out
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_stack_get_child_by_name(_arg0, _arg1)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// Hhomogeneous gets whether @stack is horizontally homogeneous.
func (s stack) Hhomogeneous() bool {
	var _arg0 *C.GtkStack // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_get_hhomogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InterpolateSize returns whether the Stack is set up to interpolate
// between the sizes of children on page switch.
func (s stack) InterpolateSize() bool {
	var _arg0 *C.GtkStack // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_get_interpolate_size(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Page returns the `GtkStackPage` object for @child.
func (s stack) Page(child Widget) StackPage {
	var _arg0 *C.GtkStack     // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.GtkStackPage // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_stack_get_page(_arg0, _arg1)

	var _stackPage StackPage // out

	_stackPage = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(StackPage)

	return _stackPage
}

// Pages returns a `GListModel` that contains the pages of the stack.
//
// This can be used to keep an up-to-date view. The model also implements
// [iface@Gtk.SelectionModel] and can be used to track and modify the
// visible page.
func (s stack) Pages() SelectionModel {
	var _arg0 *C.GtkStack          // out
	var _cret *C.GtkSelectionModel // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_get_pages(_arg0)

	var _selectionModel SelectionModel // out

	_selectionModel = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(SelectionModel)

	return _selectionModel
}

// TransitionDuration returns the amount of time (in milliseconds) that
// transitions between pages in @stack will take.
func (s stack) TransitionDuration() uint {
	var _arg0 *C.GtkStack // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_get_transition_duration(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// TransitionRunning returns whether the @stack is currently in a transition
// from one page to another.
func (s stack) TransitionRunning() bool {
	var _arg0 *C.GtkStack // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_get_transition_running(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TransitionType gets the type of animation that will be used for
// transitions between pages in @stack.
func (s stack) TransitionType() StackTransitionType {
	var _arg0 *C.GtkStack              // out
	var _cret C.GtkStackTransitionType // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_get_transition_type(_arg0)

	var _stackTransitionType StackTransitionType // out

	_stackTransitionType = StackTransitionType(_cret)

	return _stackTransitionType
}

// Vhomogeneous gets whether @stack is vertically homogeneous.
func (s stack) Vhomogeneous() bool {
	var _arg0 *C.GtkStack // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_get_vhomogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VisibleChild gets the currently visible child of @stack.
//
// Returns nil if there are no visible children.
func (s stack) VisibleChild() Widget {
	var _arg0 *C.GtkStack  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_get_visible_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// VisibleChildName returns the name of the currently visible child of
// @stack.
//
// Returns nil if there is no visible child.
func (s stack) VisibleChildName() string {
	var _arg0 *C.GtkStack // out
	var _cret *C.char     // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_get_visible_child_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Remove removes a child widget from @stack.
func (s stack) Remove(child Widget) {
	var _arg0 *C.GtkStack  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_stack_remove(_arg0, _arg1)
}

// SetHhomogeneous sets the `GtkStack` to be horizontally homogeneous or
// not.
//
// If it is homogeneous, the `GtkStack` will request the same width for all
// its children. If it isn't, the stack may change width when a different
// child becomes visible.
func (s stack) SetHhomogeneous(hhomogeneous bool) {
	var _arg0 *C.GtkStack // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	if hhomogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_stack_set_hhomogeneous(_arg0, _arg1)
}

// SetInterpolateSize sets whether or not @stack will interpolate its size
// when changing the visible child.
//
// If the [property@Gtk.Stack:interpolate-size] property is set to true,
// @stack will interpolate its size between the current one and the one
// it'll take after changing the visible child, according to the set
// transition duration.
func (s stack) SetInterpolateSize(interpolateSize bool) {
	var _arg0 *C.GtkStack // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	if interpolateSize {
		_arg1 = C.TRUE
	}

	C.gtk_stack_set_interpolate_size(_arg0, _arg1)
}

// SetTransitionDuration sets the duration that transitions between pages in
// @stack will take.
func (s stack) SetTransitionDuration(duration uint) {
	var _arg0 *C.GtkStack // out
	var _arg1 C.guint     // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	_arg1 = (C.guint)(duration)

	C.gtk_stack_set_transition_duration(_arg0, _arg1)
}

// SetTransitionType sets the type of animation that will be used for
// transitions between pages in @stack.
//
// Available types include various kinds of fades and slides.
//
// The transition type can be changed without problems at runtime, so it is
// possible to change the animation based on the page that is about to
// become current.
func (s stack) SetTransitionType(transition StackTransitionType) {
	var _arg0 *C.GtkStack              // out
	var _arg1 C.GtkStackTransitionType // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkStackTransitionType)(transition)

	C.gtk_stack_set_transition_type(_arg0, _arg1)
}

// SetVhomogeneous sets the Stack to be vertically homogeneous or not.
//
// If it is homogeneous, the `GtkStack` will request the same height for all
// its children. If it isn't, the stack may change height when a different
// child becomes visible.
func (s stack) SetVhomogeneous(vhomogeneous bool) {
	var _arg0 *C.GtkStack // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	if vhomogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_stack_set_vhomogeneous(_arg0, _arg1)
}

// SetVisibleChild makes @child the visible child of @stack.
//
// If @child is different from the currently visible child, the transition
// between the two will be animated with the current transition type of
// @stack.
//
// Note that the @child widget has to be visible itself (see
// [method@Gtk.Widget.show]) in order to become the visible child of @stack.
func (s stack) SetVisibleChild(child Widget) {
	var _arg0 *C.GtkStack  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_stack_set_visible_child(_arg0, _arg1)
}

// SetVisibleChildFull makes the child with the given name visible.
//
// Note that the child widget has to be visible itself (see
// [method@Gtk.Widget.show]) in order to become the visible child of @stack.
func (s stack) SetVisibleChildFull(name string, transition StackTransitionType) {
	var _arg0 *C.GtkStack              // out
	var _arg1 *C.char                  // out
	var _arg2 C.GtkStackTransitionType // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.GtkStackTransitionType)(transition)

	C.gtk_stack_set_visible_child_full(_arg0, _arg1, _arg2)
}

// SetVisibleChildName makes the child with the given name visible.
//
// If @child is different from the currently visible child, the transition
// between the two will be animated with the current transition type of
// @stack.
//
// Note that the child widget has to be visible itself (see
// [method@Gtk.Widget.show]) in order to become the visible child of @stack.
func (s stack) SetVisibleChildName(name string) {
	var _arg0 *C.GtkStack // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_stack_set_visible_child_name(_arg0, _arg1)
}

// StackPage: `GtkStackPage` is an auxiliary class used by `GtkStack`.
type StackPage interface {
	gextras.Objector
	Accessible

	// Child returns the stack child to which @self belongs.
	Child() Widget
	// IconName returns the icon name of the page.
	IconName() string
	// Name returns the name of the page.
	Name() string
	// NeedsAttention returns whether the page is marked as “needs attention”.
	NeedsAttention() bool
	// Title gets the page title.
	Title() string
	// UseUnderline gets whether underlines in the page title indicate
	// mnemonics.
	UseUnderline() bool
	// Visible returns whether @page is visible in its `GtkStack`.
	//
	// This is independent from the [property@Gtk.Widget:visible] property of
	// its widget.
	Visible() bool
	// SetIconName sets the icon name of the page.
	SetIconName(setting string)
	// SetName sets the name of the page.
	SetName(setting string)
	// SetNeedsAttention sets whether the page is marked as “needs attention”.
	SetNeedsAttention(setting bool)
	// SetTitle sets the page title.
	SetTitle(setting string)
	// SetUseUnderline sets whether underlines in the page title indicate
	// mnemonics.
	SetUseUnderline(setting bool)
	// SetVisible sets whether @page is visible in its `GtkStack`.
	SetVisible(visible bool)
}

// stackPage implements the StackPage class.
type stackPage struct {
	gextras.Objector
	Accessible
}

var _ StackPage = (*stackPage)(nil)

// WrapStackPage wraps a GObject to the right type. It is
// primarily used internally.
func WrapStackPage(obj *externglib.Object) StackPage {
	return stackPage{
		Objector:   obj,
		Accessible: WrapAccessible(obj),
	}
}

func marshalStackPage(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStackPage(obj), nil
}

// Child returns the stack child to which @self belongs.
func (s stackPage) Child() Widget {
	var _arg0 *C.GtkStackPage // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_page_get_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// IconName returns the icon name of the page.
func (s stackPage) IconName() string {
	var _arg0 *C.GtkStackPage // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_page_get_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Name returns the name of the page.
func (s stackPage) Name() string {
	var _arg0 *C.GtkStackPage // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_page_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// NeedsAttention returns whether the page is marked as “needs attention”.
func (s stackPage) NeedsAttention() bool {
	var _arg0 *C.GtkStackPage // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_page_get_needs_attention(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Title gets the page title.
func (s stackPage) Title() string {
	var _arg0 *C.GtkStackPage // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_page_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// UseUnderline gets whether underlines in the page title indicate
// mnemonics.
func (s stackPage) UseUnderline() bool {
	var _arg0 *C.GtkStackPage // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_page_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Visible returns whether @page is visible in its `GtkStack`.
//
// This is independent from the [property@Gtk.Widget:visible] property of
// its widget.
func (s stackPage) Visible() bool {
	var _arg0 *C.GtkStackPage // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_stack_page_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetIconName sets the icon name of the page.
func (s stackPage) SetIconName(setting string) {
	var _arg0 *C.GtkStackPage // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(setting))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_stack_page_set_icon_name(_arg0, _arg1)
}

// SetName sets the name of the page.
func (s stackPage) SetName(setting string) {
	var _arg0 *C.GtkStackPage // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(setting))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_stack_page_set_name(_arg0, _arg1)
}

// SetNeedsAttention sets whether the page is marked as “needs attention”.
func (s stackPage) SetNeedsAttention(setting bool) {
	var _arg0 *C.GtkStackPage // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_stack_page_set_needs_attention(_arg0, _arg1)
}

// SetTitle sets the page title.
func (s stackPage) SetTitle(setting string) {
	var _arg0 *C.GtkStackPage // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(setting))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_stack_page_set_title(_arg0, _arg1)
}

// SetUseUnderline sets whether underlines in the page title indicate
// mnemonics.
func (s stackPage) SetUseUnderline(setting bool) {
	var _arg0 *C.GtkStackPage // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_stack_page_set_use_underline(_arg0, _arg1)
}

// SetVisible sets whether @page is visible in its `GtkStack`.
func (s stackPage) SetVisible(visible bool) {
	var _arg0 *C.GtkStackPage // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(s.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_stack_page_set_visible(_arg0, _arg1)
}
