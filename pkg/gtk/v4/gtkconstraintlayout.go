// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_constraint_layout_get_type()), F: marshalConstraintLayouter},
		{T: externglib.Type(C.gtk_constraint_layout_child_get_type()), F: marshalConstraintLayoutChilder},
	})
}

// ConstraintLayout: layout manager using constraints to describe relations
// between widgets.
//
// GtkConstraintLayout is a layout manager that uses relations between widget
// attributes, expressed via gtk.Constraint instances, to measure and allocate
// widgets.
//
//
// How do constraints work
//
// Constraints are objects defining the relationship between attributes of a
// widget; you can read the description of the gtk.Constraint class to have a
// more in depth definition.
//
// By taking multiple constraints and applying them to the children of a widget
// using GtkConstraintLayout, it's possible to describe complex layout policies;
// each constraint applied to a child or to the parent widgets contributes to
// the full description of the layout, in terms of parameters for resolving the
// value of each attribute.
//
// It is important to note that a layout is defined by the totality of
// constraints; removing a child, or a constraint, from an existing layout
// without changing the remaining constraints may result in an unstable or
// unsolvable layout.
//
// Constraints have an implicit "reading order"; you should start describing
// each edge of each child, as well as their relationship with the parent
// container, from the top left (or top right, in RTL languages), horizontally
// first, and then vertically.
//
// A constraint-based layout with too few constraints can become "unstable",
// that is: have more than one solution. The behavior of an unstable layout is
// undefined.
//
// A constraint-based layout with conflicting constraints may be unsolvable, and
// lead to an unstable layout. You can use the gtk.Constraint:strength property
// of gtk.Constraint to "nudge" the layout towards a solution.
//
//
// GtkConstraintLayout as GtkBuildable
//
// GtkConstraintLayout implements the gtk.Buildable interface and has a custom
// "constraints" element which allows describing constraints in a gtk.Builder UI
// file.
//
// An example of a UI definition fragment specifying a constraint:
//
//      <object class="GtkConstraintLayout">
//        <constraints>
//          <constraint target="button" target-attribute="start"
//                      relation="eq"
//                      source="super" source-attribute="start"
//                      constant="12"
//                      strength="required" />
//          <constraint target="button" target-attribute="width"
//                      relation="ge"
//                      constant="250"
//                      strength="strong" />
//        </constraints>
//      </object>
//
//
// The definition above will add two constraints to the GtkConstraintLayout:
//
//    - a required constraint between the leading edge of "button" and
//      the leading edge of the widget using the constraint layout, plus
//      12 pixels
//    - a strong, constant constraint making the width of "button" greater
//      than, or equal to 250 pixels
//
// The "target" and "target-attribute" attributes are required.
//
// The "source" and "source-attribute" attributes of the "constraint" element
// are optional; if they are not specified, the constraint is assumed to be a
// constant.
//
// The "relation" attribute is optional; if not specified, the constraint is
// assumed to be an equality.
//
// The "strength" attribute is optional; if not specified, the constraint is
// assumed to be required.
//
// The "source" and "target" attributes can be set to "super" to indicate that
// the constraint target is the widget using the GtkConstraintLayout.
//
// There can be "constant" and "multiplier" attributes.
//
// Additionally, the "constraints" element can also contain a description of the
// ConstraintGuides used by the layout:
//
//      <constraints>
//        <guide min-width="100" max-width="500" name="hspace"/>
//        <guide min-height="64" nat-height="128" name="vspace" strength="strong"/>
//      </constraints>
//
//
// The "guide" element has the following optional attributes:
//
//    - "min-width", "nat-width", and "max-width", describe the minimum,
//      natural, and maximum width of the guide, respectively
//    - "min-height", "nat-height", and "max-height", describe the minimum,
//      natural, and maximum height of the guide, respectively
//    - "strength" describes the strength of the constraint on the natural
//      size of the guide; if not specified, the constraint is assumed to
//      have a medium strength
//    - "name" describes a name for the guide, useful when debugging
//
//
// Using the Visual Format Language
//
// Complex constraints can be described using a compact syntax called VFL, or
// *Visual Format Language*.
//
// The Visual Format Language describes all the constraints on a row or column,
// typically starting from the leading edge towards the trailing one. Each
// element of the layout is composed by "views", which identify a
// gtk.ConstraintTarget.
//
// For instance:
//
//      [button]-[textField]
//
//
// Describes a constraint that binds the trailing edge of "button" to the
// leading edge of "textField", leaving a default space between the two.
//
// Using VFL is also possible to specify predicates that describe constraints on
// attributes like width and height:
//
//      // Width must be greater than, or equal to 50
//      [button(>=50)]
//
//      // Width of button1 must be equal to width of button2
//      [button1(==button2)]
//
//
// The default orientation for a VFL description is horizontal, unless otherwise
// specified:
//
//      // horizontal orientation, default attribute: width
//      H:[button(>=150)]
//
//      // vertical orientation, default attribute: height
//      V:[button1(==button2)]
//
//
// It's also possible to specify multiple predicates, as well as their strength:
//
//      // minimum width of button must be 150
//      // natural width of button can be 250
//      [button(>=150required, ==250medium)]
//
//
// Finally, it's also possible to use simple arithmetic operators:
//
//      // width of button1 must be equal to width of button2
//      // divided by 2 plus 12
//      [button1(button2 / 2 + 12)]
type ConstraintLayout struct {
	LayoutManager

	Buildable
	*externglib.Object
}

func wrapConstraintLayout(obj *externglib.Object) *ConstraintLayout {
	return &ConstraintLayout{
		LayoutManager: LayoutManager{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalConstraintLayouter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapConstraintLayout(obj), nil
}

// NewConstraintLayout creates a new GtkConstraintLayout layout manager.
func NewConstraintLayout() *ConstraintLayout {
	var _cret *C.GtkLayoutManager // in

	_cret = C.gtk_constraint_layout_new()

	var _constraintLayout *ConstraintLayout // out

	_constraintLayout = wrapConstraintLayout(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _constraintLayout
}

// AddConstraint adds a constraint to the layout manager.
//
// The gtk.Constraint:source and gtk.Constraint:target properties of constraint
// can be:
//
//    - set to NULL to indicate that the constraint refers to the
//      widget using layout
//    - set to the gtk.Widget using layout
//    - set to a child of the gtk.Widget using layout
//    - set to a gtk.ConstraintGuide that is part of layout
//
// The layout acquires the ownership of constraint after calling this function.
func (layout *ConstraintLayout) AddConstraint(constraint *Constraint) {
	var _arg0 *C.GtkConstraintLayout // out
	var _arg1 *C.GtkConstraint       // out

	_arg0 = (*C.GtkConstraintLayout)(unsafe.Pointer(layout.Native()))
	_arg1 = (*C.GtkConstraint)(unsafe.Pointer(constraint.Native()))
	constraint.Ref()

	C.gtk_constraint_layout_add_constraint(_arg0, _arg1)
}

// AddConstraintsFromDescription creates a list of constraints from a VFL
// description.
//
// The Visual Format Language, VFL, is based on Apple's AutoLayout VFL
// (https://developer.apple.com/library/content/documentation/UserExperience/Conceptual/AutolayoutPG/VisualFormatLanguage.html).
//
// The views dictionary is used to match gtk.ConstraintTarget instances to the
// symbolic view name inside the VFL.
//
// The VFL grammar is:
//
//           <visualFormatString> = (<orientation>)?
//                                  (<superview><connection>)?
//                                  <view>(<connection><view>)*
//                                  (<connection><superview>)?
//                  <orientation> = 'H' | 'V'
//                    <superview> = '|'
//                   <connection> = '' | '-' <predicateList> '-' | '-'
//                <predicateList> = <simplePredicate> | <predicateListWithParens>
//              <simplePredicate> = <metricName> | <positiveNumber>
//      <predicateListWithParens> = '(' <predicate> (',' <predicate>)* ')'
//                    <predicate> = (<relation>)? <objectOfPredicate> (<operatorList>)? ('@' <priority>)?
//                     <relation> = '==' | '<=' | '>='
//            <objectOfPredicate> = <constant> | <viewName> | ('.' <attributeName>)?
//                     <priority> = <positiveNumber> | 'required' | 'strong' | 'medium' | 'weak'
//                     <constant> = <number>
//                 <operatorList> = (<multiplyOperator>)? (<addOperator>)?
//             <multiplyOperator> = [ '*' | '/' ] <positiveNumber>
//                  <addOperator> = [ '+' | '-' ] <positiveNumber>
//                     <viewName> = A-Za-z_ ([A-Za-z0-9_]*) // A C identifier
//                   <metricName> = A-Za-z_ ([A-Za-z0-9_]*) // A C identifier
//                <attributeName> = 'top' | 'bottom' | 'left' | 'right' | 'width' | 'height' |
//                                  'start' | 'end' | 'centerX' | 'centerY' | 'baseline'
//               <positiveNumber> // A positive real number parseable by g_ascii_strtod()
//                       <number> // A real number parseable by g_ascii_strtod()
//
//
// **Note**: The VFL grammar used by GTK is slightly different than the one
// defined by Apple, as it can use symbolic values for the constraint's strength
// instead of numeric values; additionally, GTK allows adding simple arithmetic
// operations inside predicates.
//
// Examples of VFL descriptions are:
//
//      // Default spacing
//      [button]-[textField]
//
//      // Width constraint
//      [button(>=50)]
//
//      // Connection to super view
//      |-50-[purpleBox]-50-|
//
//      // Vertical layout
//      V:[topField]-10-[bottomField]
//
//      // Flush views
//      [maroonView][blueView]
//
//      // Priority
//      [button(100strong)]
//
//      // Equal widths
//      [button1(==button2)]
//
//      // Multiple predicates
//      [flexibleButton(>=70,<=100)]
//
//      // A complete line of layout
//      |-[find]-[findNext]-[findField(>=20)]-|
//
//      // Operators
//      [button1(button2 / 3 + 50)]
//
//      // Named attributes
//      [button1(==button2.height)]
func (layout *ConstraintLayout) AddConstraintsFromDescription(lines []string, hspacing int, vspacing int, views map[string]ConstraintTargetter) ([]Constraint, error) {
	var _arg0 *C.GtkConstraintLayout // out
	var _arg1 **C.char               // out
	var _arg2 C.gsize
	var _arg3 C.int         // out
	var _arg4 C.int         // out
	var _arg5 *C.GHashTable // out
	var _cret *C.GList      // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkConstraintLayout)(unsafe.Pointer(layout.Native()))
	_arg2 = (C.gsize)(len(lines))
	_arg1 = (**C.char)(C.malloc(C.ulong(len(lines)) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((**C.char)(_arg1), len(lines))
		for i := range lines {
			out[i] = (*C.char)(unsafe.Pointer(C.CString(lines[i])))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}
	_arg3 = C.int(hspacing)
	_arg4 = C.int(vspacing)
	_arg5 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range views {
		var kdst *C.gchar               // out
		var vdst *C.GtkConstraintTarget // out
		kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
		defer C.free(unsafe.Pointer(kdst))
		vdst = (*C.GtkConstraintTarget)(unsafe.Pointer(vsrc.Native()))
		C.g_hash_table_insert(_arg5, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg5)

	_cret = C.gtk_constraint_layout_add_constraints_from_descriptionv(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, &_cerr)

	var _list []Constraint // out
	var _goerr error       // out

	_list = make([]Constraint, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkConstraint)(v)
		var dst Constraint // out
		dst = *wrapConstraint(externglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// AddGuide adds a guide to layout.
//
// A guide can be used as the source or target of constraints, like a widget,
// but it is not visible.
//
// The layout acquires the ownership of guide after calling this function.
func (layout *ConstraintLayout) AddGuide(guide *ConstraintGuide) {
	var _arg0 *C.GtkConstraintLayout // out
	var _arg1 *C.GtkConstraintGuide  // out

	_arg0 = (*C.GtkConstraintLayout)(unsafe.Pointer(layout.Native()))
	_arg1 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))
	guide.Ref()

	C.gtk_constraint_layout_add_guide(_arg0, _arg1)
}

// ObserveConstraints returns a GListModel to track the constraints that are
// part of the layout.
//
// Calling this function will enable extra internal bookkeeping to track
// constraints and emit signals on the returned listmodel. It may slow down
// operations a lot.
//
// Applications should try hard to avoid calling this function because of the
// slowdowns.
func (layout *ConstraintLayout) ObserveConstraints() gio.ListModeller {
	var _arg0 *C.GtkConstraintLayout // out
	var _cret *C.GListModel          // in

	_arg0 = (*C.GtkConstraintLayout)(unsafe.Pointer(layout.Native()))

	_cret = C.gtk_constraint_layout_observe_constraints(_arg0)

	var _listModel gio.ListModeller // out

	_listModel = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(gio.ListModeller)

	return _listModel
}

// ObserveGuides returns a GListModel to track the guides that are part of the
// layout.
//
// Calling this function will enable extra internal bookkeeping to track guides
// and emit signals on the returned listmodel. It may slow down operations a
// lot.
//
// Applications should try hard to avoid calling this function because of the
// slowdowns.
func (layout *ConstraintLayout) ObserveGuides() gio.ListModeller {
	var _arg0 *C.GtkConstraintLayout // out
	var _cret *C.GListModel          // in

	_arg0 = (*C.GtkConstraintLayout)(unsafe.Pointer(layout.Native()))

	_cret = C.gtk_constraint_layout_observe_guides(_arg0)

	var _listModel gio.ListModeller // out

	_listModel = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(gio.ListModeller)

	return _listModel
}

// RemoveAllConstraints removes all constraints from the layout manager.
func (layout *ConstraintLayout) RemoveAllConstraints() {
	var _arg0 *C.GtkConstraintLayout // out

	_arg0 = (*C.GtkConstraintLayout)(unsafe.Pointer(layout.Native()))

	C.gtk_constraint_layout_remove_all_constraints(_arg0)
}

// RemoveConstraint removes constraint from the layout manager, so that it no
// longer influences the layout.
func (layout *ConstraintLayout) RemoveConstraint(constraint *Constraint) {
	var _arg0 *C.GtkConstraintLayout // out
	var _arg1 *C.GtkConstraint       // out

	_arg0 = (*C.GtkConstraintLayout)(unsafe.Pointer(layout.Native()))
	_arg1 = (*C.GtkConstraint)(unsafe.Pointer(constraint.Native()))

	C.gtk_constraint_layout_remove_constraint(_arg0, _arg1)
}

// RemoveGuide removes guide from the layout manager, so that it no longer
// influences the layout.
func (layout *ConstraintLayout) RemoveGuide(guide *ConstraintGuide) {
	var _arg0 *C.GtkConstraintLayout // out
	var _arg1 *C.GtkConstraintGuide  // out

	_arg0 = (*C.GtkConstraintLayout)(unsafe.Pointer(layout.Native()))
	_arg1 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))

	C.gtk_constraint_layout_remove_guide(_arg0, _arg1)
}

// ConstraintLayoutChild: GtkLayoutChild subclass for children in a
// GtkConstraintLayout.
type ConstraintLayoutChild struct {
	LayoutChild
}

func wrapConstraintLayoutChild(obj *externglib.Object) *ConstraintLayoutChild {
	return &ConstraintLayoutChild{
		LayoutChild: LayoutChild{
			Object: obj,
		},
	}
}

func marshalConstraintLayoutChilder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapConstraintLayoutChild(obj), nil
}

func (*ConstraintLayoutChild) privateConstraintLayoutChild() {}
