// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_stack_transition_type_get_type()), F: marshalStackTransitionType},
		{T: externglib.Type(C.gtk_stack_get_type()), F: marshalStacker},
		{T: externglib.Type(C.gtk_stack_page_get_type()), F: marshalStackPager},
	})
}

// StackTransitionType: possible transitions between pages in a `GtkStack`
// widget.
//
// New values may be added to this enumeration over time.
type StackTransitionType int

const (
	// None: no transition
	StackTransitionTypeNone StackTransitionType = iota
	// Crossfade: cross-fade
	StackTransitionTypeCrossfade
	// SlideRight: slide from left to right
	StackTransitionTypeSlideRight
	// SlideLeft: slide from right to left
	StackTransitionTypeSlideLeft
	// SlideUp: slide from bottom up
	StackTransitionTypeSlideUp
	// SlideDown: slide from top down
	StackTransitionTypeSlideDown
	// SlideLeftRight: slide from left or right according to the children order
	StackTransitionTypeSlideLeftRight
	// SlideUpDown: slide from top down or bottom up according to the order
	StackTransitionTypeSlideUpDown
	// OverUp: cover the old page by sliding up
	StackTransitionTypeOverUp
	// OverDown: cover the old page by sliding down
	StackTransitionTypeOverDown
	// OverLeft: cover the old page by sliding to the left
	StackTransitionTypeOverLeft
	// OverRight: cover the old page by sliding to the right
	StackTransitionTypeOverRight
	// UnderUp: uncover the new page by sliding up
	StackTransitionTypeUnderUp
	// UnderDown: uncover the new page by sliding down
	StackTransitionTypeUnderDown
	// UnderLeft: uncover the new page by sliding to the left
	StackTransitionTypeUnderLeft
	// UnderRight: uncover the new page by sliding to the right
	StackTransitionTypeUnderRight
	// OverUpDown: cover the old page sliding up or uncover the new page sliding
	// down, according to order
	StackTransitionTypeOverUpDown
	// OverDownUp: cover the old page sliding down or uncover the new page
	// sliding up, according to order
	StackTransitionTypeOverDownUp
	// OverLeftRight: cover the old page sliding left or uncover the new page
	// sliding right, according to order
	StackTransitionTypeOverLeftRight
	// OverRightLeft: cover the old page sliding right or uncover the new page
	// sliding left, according to order
	StackTransitionTypeOverRightLeft
	// RotateLeft: pretend the pages are sides of a cube and rotate that cube to
	// the left
	StackTransitionTypeRotateLeft
	// RotateRight: pretend the pages are sides of a cube and rotate that cube
	// to the right
	StackTransitionTypeRotateRight
	// RotateLeftRight: pretend the pages are sides of a cube and rotate that
	// cube to the left or right according to the children order
	StackTransitionTypeRotateLeftRight
)

func marshalStackTransitionType(p uintptr) (interface{}, error) {
	return StackTransitionType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Stacker describes Stack's methods.
type Stacker interface {
	// AddChild adds a child to @stack.
	AddChild(child Widgeter) *StackPage
	// AddNamed adds a child to @stack.
	AddNamed(child Widgeter, name string) *StackPage
	// AddTitled adds a child to @stack.
	AddTitled(child Widgeter, name string, title string) *StackPage
	// ChildByName finds the child with the name given as the argument.
	ChildByName(name string) *Widget
	// Hhomogeneous gets whether @stack is horizontally homogeneous.
	Hhomogeneous() bool
	// InterpolateSize returns whether the Stack is set up to interpolate
	// between the sizes of children on page switch.
	InterpolateSize() bool
	// Page returns the `GtkStackPage` object for @child.
	Page(child Widgeter) *StackPage
	// Pages returns a `GListModel` that contains the pages of the stack.
	Pages() *SelectionModel
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
	VisibleChild() *Widget
	// VisibleChildName returns the name of the currently visible child of
	// @stack.
	VisibleChildName() string
	// Remove removes a child widget from @stack.
	Remove(child Widgeter)
	// SetHhomogeneous sets the `GtkStack` to be horizontally homogeneous or
	// not.
	SetHhomogeneous(hhomogeneous bool)
	// SetInterpolateSize sets whether or not @stack will interpolate its size
	// when changing the visible child.
	SetInterpolateSize(interpolateSize bool)
	// SetTransitionDuration sets the duration that transitions between pages in
	// @stack will take.
	SetTransitionDuration(duration uint)
	// SetTransitionType sets the type of animation that will be used for
	// transitions between pages in @stack.
	SetTransitionType(transition StackTransitionType)
	// SetVhomogeneous sets the Stack to be vertically homogeneous or not.
	SetVhomogeneous(vhomogeneous bool)
	// SetVisibleChild makes @child the visible child of @stack.
	SetVisibleChild(child Widgeter)
	// SetVisibleChildFull makes the child with the given name visible.
	SetVisibleChildFull(name string, transition StackTransitionType)
	// SetVisibleChildName makes the child with the given name visible.
	SetVisibleChildName(name string)
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
type Stack struct {
	Widget
}

var (
	_ Stacker         = (*Stack)(nil)
	_ gextras.Nativer = (*Stack)(nil)
)

func wrapStack(obj *externglib.Object) *Stack {
	return &Stack{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
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
	}
}

func marshalStacker(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStack(obj), nil
}

// NewStack creates a new `GtkStack`.
func NewStack() *Stack {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_stack_new()

	var _stack *Stack // out

	_stack = wrapStack(externglib.Take(unsafe.Pointer(_cret)))

	return _stack
}

// AddChild adds a child to @stack.
func (stack *Stack) AddChild(child Widgeter) *StackPage {
	var _arg0 *C.GtkStack     // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.GtkStackPage // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	_cret = C.gtk_stack_add_child(_arg0, _arg1)

	var _stackPage *StackPage // out

	_stackPage = wrapStackPage(externglib.Take(unsafe.Pointer(_cret)))

	return _stackPage
}

// AddNamed adds a child to @stack.
//
// The child is identified by the @name.
func (stack *Stack) AddNamed(child Widgeter, name string) *StackPage {
	var _arg0 *C.GtkStack     // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.char         // out
	var _cret *C.GtkStackPage // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(name)))

	_cret = C.gtk_stack_add_named(_arg0, _arg1, _arg2)

	var _stackPage *StackPage // out

	_stackPage = wrapStackPage(externglib.Take(unsafe.Pointer(_cret)))

	return _stackPage
}

// AddTitled adds a child to @stack.
//
// The child is identified by the @name. The @title will be used by
// `GtkStackSwitcher` to represent @child in a tab bar, so it should be short.
func (stack *Stack) AddTitled(child Widgeter, name string, title string) *StackPage {
	var _arg0 *C.GtkStack     // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.char         // out
	var _arg3 *C.char         // out
	var _cret *C.GtkStackPage // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(name)))
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(title)))

	_cret = C.gtk_stack_add_titled(_arg0, _arg1, _arg2, _arg3)

	var _stackPage *StackPage // out

	_stackPage = wrapStackPage(externglib.Take(unsafe.Pointer(_cret)))

	return _stackPage
}

// ChildByName finds the child with the name given as the argument.
//
// Returns nil if there is no child with this name.
func (stack *Stack) ChildByName(name string) *Widget {
	var _arg0 *C.GtkStack  // out
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))

	_cret = C.gtk_stack_get_child_by_name(_arg0, _arg1)

	var _widget *Widget // out

	_widget = wrapWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _widget
}

// Hhomogeneous gets whether @stack is horizontally homogeneous.
func (stack *Stack) Hhomogeneous() bool {
	var _arg0 *C.GtkStack // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_hhomogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InterpolateSize returns whether the Stack is set up to interpolate between
// the sizes of children on page switch.
func (stack *Stack) InterpolateSize() bool {
	var _arg0 *C.GtkStack // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_interpolate_size(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Page returns the `GtkStackPage` object for @child.
func (stack *Stack) Page(child Widgeter) *StackPage {
	var _arg0 *C.GtkStack     // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.GtkStackPage // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	_cret = C.gtk_stack_get_page(_arg0, _arg1)

	var _stackPage *StackPage // out

	_stackPage = wrapStackPage(externglib.Take(unsafe.Pointer(_cret)))

	return _stackPage
}

// Pages returns a `GListModel` that contains the pages of the stack.
//
// This can be used to keep an up-to-date view. The model also implements
// [iface@Gtk.SelectionModel] and can be used to track and modify the visible
// page.
func (stack *Stack) Pages() *SelectionModel {
	var _arg0 *C.GtkStack          // out
	var _cret *C.GtkSelectionModel // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_pages(_arg0)

	var _selectionModel *SelectionModel // out

	_selectionModel = wrapSelectionModel(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _selectionModel
}

// TransitionDuration returns the amount of time (in milliseconds) that
// transitions between pages in @stack will take.
func (stack *Stack) TransitionDuration() uint {
	var _arg0 *C.GtkStack // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_transition_duration(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// TransitionRunning returns whether the @stack is currently in a transition
// from one page to another.
func (stack *Stack) TransitionRunning() bool {
	var _arg0 *C.GtkStack // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_transition_running(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TransitionType gets the type of animation that will be used for transitions
// between pages in @stack.
func (stack *Stack) TransitionType() StackTransitionType {
	var _arg0 *C.GtkStack              // out
	var _cret C.GtkStackTransitionType // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_transition_type(_arg0)

	var _stackTransitionType StackTransitionType // out

	_stackTransitionType = StackTransitionType(_cret)

	return _stackTransitionType
}

// Vhomogeneous gets whether @stack is vertically homogeneous.
func (stack *Stack) Vhomogeneous() bool {
	var _arg0 *C.GtkStack // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

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
func (stack *Stack) VisibleChild() *Widget {
	var _arg0 *C.GtkStack  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_visible_child(_arg0)

	var _widget *Widget // out

	_widget = wrapWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _widget
}

// VisibleChildName returns the name of the currently visible child of @stack.
//
// Returns nil if there is no visible child.
func (stack *Stack) VisibleChildName() string {
	var _arg0 *C.GtkStack // out
	var _cret *C.char     // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_visible_child_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Remove removes a child widget from @stack.
func (stack *Stack) Remove(child Widgeter) {
	var _arg0 *C.GtkStack  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_stack_remove(_arg0, _arg1)
}

// SetHhomogeneous sets the `GtkStack` to be horizontally homogeneous or not.
//
// If it is homogeneous, the `GtkStack` will request the same width for all its
// children. If it isn't, the stack may change width when a different child
// becomes visible.
func (stack *Stack) SetHhomogeneous(hhomogeneous bool) {
	var _arg0 *C.GtkStack // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	if hhomogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_stack_set_hhomogeneous(_arg0, _arg1)
}

// SetInterpolateSize sets whether or not @stack will interpolate its size when
// changing the visible child.
//
// If the [property@Gtk.Stack:interpolate-size] property is set to true, @stack
// will interpolate its size between the current one and the one it'll take
// after changing the visible child, according to the set transition duration.
func (stack *Stack) SetInterpolateSize(interpolateSize bool) {
	var _arg0 *C.GtkStack // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	if interpolateSize {
		_arg1 = C.TRUE
	}

	C.gtk_stack_set_interpolate_size(_arg0, _arg1)
}

// SetTransitionDuration sets the duration that transitions between pages in
// @stack will take.
func (stack *Stack) SetTransitionDuration(duration uint) {
	var _arg0 *C.GtkStack // out
	var _arg1 C.guint     // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = C.guint(duration)

	C.gtk_stack_set_transition_duration(_arg0, _arg1)
}

// SetTransitionType sets the type of animation that will be used for
// transitions between pages in @stack.
//
// Available types include various kinds of fades and slides.
//
// The transition type can be changed without problems at runtime, so it is
// possible to change the animation based on the page that is about to become
// current.
func (stack *Stack) SetTransitionType(transition StackTransitionType) {
	var _arg0 *C.GtkStack              // out
	var _arg1 C.GtkStackTransitionType // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = C.GtkStackTransitionType(transition)

	C.gtk_stack_set_transition_type(_arg0, _arg1)
}

// SetVhomogeneous sets the Stack to be vertically homogeneous or not.
//
// If it is homogeneous, the `GtkStack` will request the same height for all its
// children. If it isn't, the stack may change height when a different child
// becomes visible.
func (stack *Stack) SetVhomogeneous(vhomogeneous bool) {
	var _arg0 *C.GtkStack // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	if vhomogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_stack_set_vhomogeneous(_arg0, _arg1)
}

// SetVisibleChild makes @child the visible child of @stack.
//
// If @child is different from the currently visible child, the transition
// between the two will be animated with the current transition type of @stack.
//
// Note that the @child widget has to be visible itself (see
// [method@Gtk.Widget.show]) in order to become the visible child of @stack.
func (stack *Stack) SetVisibleChild(child Widgeter) {
	var _arg0 *C.GtkStack  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_stack_set_visible_child(_arg0, _arg1)
}

// SetVisibleChildFull makes the child with the given name visible.
//
// Note that the child widget has to be visible itself (see
// [method@Gtk.Widget.show]) in order to become the visible child of @stack.
func (stack *Stack) SetVisibleChildFull(name string, transition StackTransitionType) {
	var _arg0 *C.GtkStack              // out
	var _arg1 *C.char                  // out
	var _arg2 C.GtkStackTransitionType // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	_arg2 = C.GtkStackTransitionType(transition)

	C.gtk_stack_set_visible_child_full(_arg0, _arg1, _arg2)
}

// SetVisibleChildName makes the child with the given name visible.
//
// If @child is different from the currently visible child, the transition
// between the two will be animated with the current transition type of @stack.
//
// Note that the child widget has to be visible itself (see
// [method@Gtk.Widget.show]) in order to become the visible child of @stack.
func (stack *Stack) SetVisibleChildName(name string) {
	var _arg0 *C.GtkStack // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))

	C.gtk_stack_set_visible_child_name(_arg0, _arg1)
}

// StackPager describes StackPage's methods.
type StackPager interface {
	// Child returns the stack child to which @self belongs.
	Child() *Widget
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

// StackPage: `GtkStackPage` is an auxiliary class used by `GtkStack`.
type StackPage struct {
	*externglib.Object

	Accessible
}

var (
	_ StackPager      = (*StackPage)(nil)
	_ gextras.Nativer = (*StackPage)(nil)
)

func wrapStackPage(obj *externglib.Object) *StackPage {
	return &StackPage{
		Object: obj,
		Accessible: Accessible{
			Object: obj,
		},
	}
}

func marshalStackPager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStackPage(obj), nil
}

// Child returns the stack child to which @self belongs.
func (self *StackPage) Child() *Widget {
	var _arg0 *C.GtkStackPage // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_stack_page_get_child(_arg0)

	var _widget *Widget // out

	_widget = wrapWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _widget
}

// IconName returns the icon name of the page.
func (self *StackPage) IconName() string {
	var _arg0 *C.GtkStackPage // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_stack_page_get_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Name returns the name of the page.
func (self *StackPage) Name() string {
	var _arg0 *C.GtkStackPage // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_stack_page_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// NeedsAttention returns whether the page is marked as “needs attention”.
func (self *StackPage) NeedsAttention() bool {
	var _arg0 *C.GtkStackPage // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_stack_page_get_needs_attention(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Title gets the page title.
func (self *StackPage) Title() string {
	var _arg0 *C.GtkStackPage // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_stack_page_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UseUnderline gets whether underlines in the page title indicate mnemonics.
func (self *StackPage) UseUnderline() bool {
	var _arg0 *C.GtkStackPage // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_stack_page_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Visible returns whether @page is visible in its `GtkStack`.
//
// This is independent from the [property@Gtk.Widget:visible] property of its
// widget.
func (self *StackPage) Visible() bool {
	var _arg0 *C.GtkStackPage // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_stack_page_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetIconName sets the icon name of the page.
func (self *StackPage) SetIconName(setting string) {
	var _arg0 *C.GtkStackPage // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(setting)))

	C.gtk_stack_page_set_icon_name(_arg0, _arg1)
}

// SetName sets the name of the page.
func (self *StackPage) SetName(setting string) {
	var _arg0 *C.GtkStackPage // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(setting)))

	C.gtk_stack_page_set_name(_arg0, _arg1)
}

// SetNeedsAttention sets whether the page is marked as “needs attention”.
func (self *StackPage) SetNeedsAttention(setting bool) {
	var _arg0 *C.GtkStackPage // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_stack_page_set_needs_attention(_arg0, _arg1)
}

// SetTitle sets the page title.
func (self *StackPage) SetTitle(setting string) {
	var _arg0 *C.GtkStackPage // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(setting)))

	C.gtk_stack_page_set_title(_arg0, _arg1)
}

// SetUseUnderline sets whether underlines in the page title indicate mnemonics.
func (self *StackPage) SetUseUnderline(setting bool) {
	var _arg0 *C.GtkStackPage // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_stack_page_set_use_underline(_arg0, _arg1)
}

// SetVisible sets whether @page is visible in its `GtkStack`.
func (self *StackPage) SetVisible(visible bool) {
	var _arg0 *C.GtkStackPage // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkStackPage)(unsafe.Pointer(self.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_stack_page_set_visible(_arg0, _arg1)
}
