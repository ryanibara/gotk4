// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_builder_error_get_type()), F: marshalBuilderError},
		{T: externglib.Type(C.gtk_builder_get_type()), F: marshalBuilder},
	})
}

// BuilderError: error codes that identify various errors that can occur while
// using Builder.
type BuilderError int

const (
	// InvalidTypeFunction: a type-func attribute didn’t name a function that
	// returns a #GType.
	BuilderErrorInvalidTypeFunction BuilderError = 0
	// UnhandledTag: the input contained a tag that Builder can’t handle.
	BuilderErrorUnhandledTag BuilderError = 1
	// MissingAttribute: an attribute that is required by Builder was missing.
	BuilderErrorMissingAttribute BuilderError = 2
	// InvalidAttribute: Builder found an attribute that it doesn’t understand.
	BuilderErrorInvalidAttribute BuilderError = 3
	// InvalidTag: Builder found a tag that it doesn’t understand.
	BuilderErrorInvalidTag BuilderError = 4
	// MissingPropertyValue: a required property value was missing.
	BuilderErrorMissingPropertyValue BuilderError = 5
	// InvalidValue: Builder couldn’t parse some attribute value.
	BuilderErrorInvalidValue BuilderError = 6
	// VersionMismatch: the input file requires a newer version of GTK.
	BuilderErrorVersionMismatch BuilderError = 7
	// DuplicateID: an object id occurred twice.
	BuilderErrorDuplicateID BuilderError = 8
	// ObjectTypeRefused: a specified object type is of the same type or derived
	// from the type of the composite class being extended with builder XML.
	BuilderErrorObjectTypeRefused BuilderError = 9
	// TemplateMismatch: the wrong type was specified in a composite class’s
	// template XML
	BuilderErrorTemplateMismatch BuilderError = 10
	// InvalidProperty: the specified property is unknown for the object class.
	BuilderErrorInvalidProperty BuilderError = 11
	// InvalidSignal: the specified signal is unknown for the object class.
	BuilderErrorInvalidSignal BuilderError = 12
	// InvalidID: an object id is unknown.
	BuilderErrorInvalidID BuilderError = 13
	// InvalidFunction: a function could not be found. This often happens when
	// symbols are set to be kept private. Compiling code with -rdynamic or
	// using the `gmodule-export-2.0` pkgconfig module can fix this problem.
	BuilderErrorInvalidFunction BuilderError = 14
)

func marshalBuilderError(p uintptr) (interface{}, error) {
	return BuilderError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Builder: a `GtkBuilder` reads XML descriptions of a user interface and
// instantiates the described objects.
//
// To create a `GtkBuilder` from a user interface description, call
// [ctor@Gtk.Builder.new_from_file], [ctor@Gtk.Builder.new_from_resource] or
// [ctor@Gtk.Builder.new_from_string].
//
// In the (unusual) case that you want to add user interface descriptions from
// multiple sources to the same `GtkBuilder` you can call [ctor@Gtk.Builder.new]
// to get an empty builder and populate it by (multiple) calls to
// [method@Gtk.Builder.add_from_file], [method@Gtk.Builder.add_from_resource] or
// [method@Gtk.Builder.add_from_string].
//
// A `GtkBuilder` holds a reference to all objects that it has constructed and
// drops these references when it is finalized. This finalization can cause the
// destruction of non-widget objects or widgets which are not contained in a
// toplevel window. For toplevel windows constructed by a builder, it is the
// responsibility of the user to call [method@Gtk.Window.destroy] to get rid of
// them and all the widgets they contain.
//
// The functions [method@Gtk.Builder.get_object] and
// [method@Gtk.Builder.get_objects] can be used to access the widgets in the
// interface by the names assigned to them inside the UI description. Toplevel
// windows returned by these functions will stay around until the user
// explicitly destroys them with [method@Gtk.Window.destroy]. Other widgets will
// either be part of a larger hierarchy constructed by the builder (in which
// case you should not have to worry about their lifecycle), or without a
// parent, in which case they have to be added to some container to make use of
// them. Non-widget objects need to be reffed with g_object_ref() to keep them
// beyond the lifespan of the builder.
//
//
// GtkBuilder UI Definitions
//
// `GtkBuilder` parses textual descriptions of user interfaces which are
// specified in XML format. We refer to these descriptions as “GtkBuilder UI
// definitions” or just “UI definitions” if the context is clear.
//
// The toplevel element is `<interface>`. It optionally takes a “domain”
// attribute, which will make the builder look for translated strings using
// `dgettext()` in the domain specified. This can also be done by calling
// [method@Gtk.Builder.set_translation_domain] on the builder.
//
// Objects are described by `<object>` elements, which can contain `<property>`
// elements to set properties, `<signal>` elements which connect signals to
// handlers, and `<child>` elements, which describe child objects (most often
// widgets inside a container, but also e.g. actions in an action group, or
// columns in a tree model). A `<child>` element contains an `<object>` element
// which describes the child object.
//
// The target toolkit version(s) are described by `<requires>` elements, the
// “lib” attribute specifies the widget library in question (currently the only
// supported value is “gtk”) and the “version” attribute specifies the target
// version in the form “`<major>`.`<minor>`”. `GtkBuilder` will error out if the
// version requirements are not met.
//
// Typically, the specific kind of object represented by an `<object>` element
// is specified by the “class” attribute. If the type has not been loaded yet,
// GTK tries to find the `get_type()` function from the class name by applying
// heuristics. This works in most cases, but if necessary, it is possible to
// specify the name of the `get_type()` function explicitly with the "type-func"
// attribute.
//
// Objects may be given a name with the “id” attribute, which allows the
// application to retrieve them from the builder with
// [method@Gtk.Builder.get_object]. An id is also necessary to use the object as
// property value in other parts of the UI definition. GTK reserves ids starting
// and ending with `___` (three consecutive underscores) for its own purposes.
//
// Setting properties of objects is pretty straightforward with the `<property>`
// element: the “name” attribute specifies the name of the property, and the
// content of the element specifies the value. If the “translatable” attribute
// is set to a true value, GTK uses `gettext()` (or `dgettext()` if the builder
// has a translation domain set) to find a translation for the value. This
// happens before the value is parsed, so it can be used for properties of any
// type, but it is probably most useful for string properties. It is also
// possible to specify a context to disambiguate short strings, and comments
// which may help the translators.
//
// `GtkBuilder` can parse textual representations for the most common property
// types: characters, strings, integers, floating-point numbers, booleans
// (strings like “TRUE”, “t”, “yes”, “y”, “1” are interpreted as true, strings
// like “FALSE”, “f”, “no”, “n”, “0” are interpreted as false), enumerations
// (can be specified by their name, nick or integer value), flags (can be
// specified by their name, nick, integer value, optionally combined with “|”,
// e.g. “GTK_INPUT_HINT_EMOJI|GTK_INPUT_HINT_LOWERCASE”) and colors (in a format
// understood by [method@Gdk.RGBA.parse]).
//
// `GVariant`s can be specified in the format understood by g_variant_parse(),
// and pixbufs can be specified as a filename of an image file to load.
//
// Objects can be referred to by their name and by default refer to objects
// declared in the local XML fragment and objects exposed via
// [method@Gtk.Builder.expose_object]. In general, `GtkBuilder` allows forward
// references to objects — declared in the local XML; an object doesn’t have to
// be constructed before it can be referred to. The exception to this rule is
// that an object has to be constructed before it can be used as the value of a
// construct-only property.
//
// It is also possible to bind a property value to another object's property
// value using the attributes "bind-source" to specify the source object of the
// binding, and optionally, "bind-property" and "bind-flags" to specify the
// source property and source binding flags respectively. Internally,
// `GtkBuilder` implements this using `GBinding` objects. For more information
// see g_object_bind_property().
//
// Sometimes it is necessary to refer to widgets which have implicitly been
// constructed by GTK as part of a composite widget, to set properties on them
// or to add further children (e.g. the content area of a `GtkDialog`). This can
// be achieved by setting the “internal-child” property of the `<child>` element
// to a true value. Note that Builder still requires an `<object>` element for
// the internal child, even if it has already been constructed.
//
// A number of widgets have different places where a child can be added (e.g.
// tabs vs. page content in notebooks). This can be reflected in a UI definition
// by specifying the “type” attribute on a `<child>` The possible values for the
// “type” attribute are described in the sections describing the widget-specific
// portions of UI definitions.
//
//
// Signal handlers and function pointers
//
// Signal handlers are set up with the `<signal>` element. The “name” attribute
// specifies the name of the signal, and the “handler” attribute specifies the
// function to connect to the signal. The remaining attributes, “after”,
// “swapped” and “object”, have the same meaning as the corresponding parameters
// of the g_signal_connect_object() or g_signal_connect_data() functions. A
// “last_modification_time” attribute is also allowed, but it does not have a
// meaning to the builder.
//
// If you rely on `GModule` support to lookup callbacks in the symbol table, the
// following details should be noted:
//
// When compiling applications for Windows, you must declare signal callbacks
// with G_MODULE_EXPORT, or they will not be put in the symbol table. On Linux
// and Unix, this is not necessary; applications should instead be compiled with
// the -Wl,--export-dynamic `CFLAGS`, and linked against `gmodule-export-2.0`.
//
//
// A GtkBuilder UI Definition
//
// “`xml <interface> <object class="GtkDialog" id="dialog1"> <child
// internal-child="vbox"> <object class="GtkBox" id="vbox1"> <child
// internal-child="action_area"> <object class="GtkBox" id="hbuttonbox1">
// <child> <object class="GtkButton" id="ok_button"> <property
// name="label">gtk-ok</property> <signal name="clicked"
// handler="ok_button_clicked"/> </object> </child> </object> </child> </object>
// </child> </object> </interface> “`
//
// Beyond this general structure, several object classes define their own XML
// DTD fragments for filling in the ANY placeholders in the DTD above. Note that
// a custom element in a <child> element gets parsed by the custom tag handler
// of the parent object, while a custom element in an <object> element gets
// parsed by the custom tag handler of the object.
//
// These XML fragments are explained in the documentation of the respective
// objects.
//
// A `<template>` tag can be used to define a widget class’s components. See the
// GtkWidget documentation
// (class.Widget.html#building-composite-widgets-from-template-xml) for details.
type Builder interface {
	gextras.Objector

	AddFromFileBuilder(filename string) error

	AddFromResourceBuilder(resourcePath string) error

	AddFromStringBuilder(buffer string, length int) error

	AddObjectsFromFileBuilder(filename string, objectIds []string) error

	AddObjectsFromResourceBuilder(resourcePath string, objectIds []string) error

	AddObjectsFromStringBuilder(buffer string, length int, objectIds []string) error

	ExposeObjectBuilder(name string, object gextras.Objector)

	ExtendWithTemplateBuilder(object gextras.Objector, templateType externglib.Type, buffer string, length int) error

	CurrentObject() gextras.Objector

	Object(name string) gextras.Objector

	Scope() BuilderScope

	TranslationDomain() string

	TypeFromName(typeName string) externglib.Type

	SetCurrentObjectBuilder(currentObject gextras.Objector)

	SetScopeBuilder(scope BuilderScope)

	SetTranslationDomainBuilder(domain string)

	ValueFromStringTypeBuilder(typ externglib.Type, _string string) (externglib.Value, error)
}

// builder implements the Builder class.
type builder struct {
	gextras.Objector
}

// WrapBuilder wraps a GObject to the right type. It is
// primarily used internally.
func WrapBuilder(obj *externglib.Object) Builder {
	return builder{
		Objector: obj,
	}
}

func marshalBuilder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBuilder(obj), nil
}

func NewBuilder() Builder {
	var _cret *C.GtkBuilder // in

	_cret = C.gtk_builder_new()

	var _builder Builder // out

	_builder = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Builder)

	return _builder
}

func NewBuilderFromFile(filename string) Builder {
	var _arg1 *C.char       // out
	var _cret *C.GtkBuilder // in

	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_builder_new_from_file(_arg1)

	var _builder Builder // out

	_builder = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Builder)

	return _builder
}

func NewBuilderFromResource(resourcePath string) Builder {
	var _arg1 *C.char       // out
	var _cret *C.GtkBuilder // in

	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_builder_new_from_resource(_arg1)

	var _builder Builder // out

	_builder = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Builder)

	return _builder
}

func NewBuilderFromString(_string string, length int) Builder {
	var _arg1 *C.char       // out
	var _arg2 C.gssize      // out
	var _cret *C.GtkBuilder // in

	_arg1 = (*C.char)(C.CString(_string))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(length)

	_cret = C.gtk_builder_new_from_string(_arg1, _arg2)

	var _builder Builder // out

	_builder = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Builder)

	return _builder
}

func (b builder) AddFromFileBuilder(filename string) error {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_builder_add_from_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (b builder) AddFromResourceBuilder(resourcePath string) error {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_builder_add_from_resource(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (b builder) AddFromStringBuilder(buffer string, length int) error {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _arg2 C.gssize      // out
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.char)(C.CString(buffer))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(length)

	C.gtk_builder_add_from_string(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (b builder) AddObjectsFromFileBuilder(filename string, objectIds []string) error {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _arg2 **C.char
	var _cerr *C.GError // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (**C.char)(C.malloc(C.ulong(len(objectIds)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(objectIds))
		for i := range objectIds {
			out[i] = (*C.char)(C.CString(objectIds[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_builder_add_objects_from_file(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (b builder) AddObjectsFromResourceBuilder(resourcePath string, objectIds []string) error {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _arg2 **C.char
	var _cerr *C.GError // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (**C.char)(C.malloc(C.ulong(len(objectIds)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(objectIds))
		for i := range objectIds {
			out[i] = (*C.char)(C.CString(objectIds[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_builder_add_objects_from_resource(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (b builder) AddObjectsFromStringBuilder(buffer string, length int, objectIds []string) error {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _arg2 C.gssize      // out
	var _arg3 **C.char
	var _cerr *C.GError // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.char)(C.CString(buffer))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(length)
	_arg3 = (**C.char)(C.malloc(C.ulong(len(objectIds)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice(_arg3, len(objectIds))
		for i := range objectIds {
			out[i] = (*C.char)(C.CString(objectIds[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_builder_add_objects_from_string(_arg0, _arg1, _arg2, _arg3, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (b builder) ExposeObjectBuilder(name string, object gextras.Objector) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _arg2 *C.GObject    // out

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GObject)(unsafe.Pointer(object.Native()))

	C.gtk_builder_expose_object(_arg0, _arg1, _arg2)
}

func (b builder) ExtendWithTemplateBuilder(object gextras.Objector, templateType externglib.Type, buffer string, length int) error {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.GObject    // out
	var _arg2 C.GType       // out
	var _arg3 *C.char       // out
	var _arg4 C.gssize      // out
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))
	_arg2 = (C.GType)(templateType)
	_arg3 = (*C.char)(C.CString(buffer))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.gssize(length)

	C.gtk_builder_extend_with_template(_arg0, _arg1, _arg2, _arg3, _arg4, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (b builder) CurrentObject() gextras.Objector {
	var _arg0 *C.GtkBuilder // out
	var _cret *C.GObject    // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_builder_get_current_object(_arg0)

	var _object gextras.Objector // out

	_object = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gextras.Objector)

	return _object
}

func (b builder) Object(name string) gextras.Objector {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _cret *C.GObject    // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_builder_get_object(_arg0, _arg1)

	var _object gextras.Objector // out

	_object = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gextras.Objector)

	return _object
}

func (b builder) Scope() BuilderScope {
	var _arg0 *C.GtkBuilder      // out
	var _cret *C.GtkBuilderScope // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_builder_get_scope(_arg0)

	var _builderScope BuilderScope // out

	_builderScope = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(BuilderScope)

	return _builderScope
}

func (b builder) TranslationDomain() string {
	var _arg0 *C.GtkBuilder // out
	var _cret *C.char       // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_builder_get_translation_domain(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (b builder) TypeFromName(typeName string) externglib.Type {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _cret C.GType       // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.char)(C.CString(typeName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_builder_get_type_from_name(_arg0, _arg1)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

func (b builder) SetCurrentObjectBuilder(currentObject gextras.Objector) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.GObject    // out

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer(currentObject.Native()))

	C.gtk_builder_set_current_object(_arg0, _arg1)
}

func (b builder) SetScopeBuilder(scope BuilderScope) {
	var _arg0 *C.GtkBuilder      // out
	var _arg1 *C.GtkBuilderScope // out

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkBuilderScope)(unsafe.Pointer(scope.Native()))

	C.gtk_builder_set_scope(_arg0, _arg1)
}

func (b builder) SetTranslationDomainBuilder(domain string) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.char)(C.CString(domain))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_builder_set_translation_domain(_arg0, _arg1)
}

func (b builder) ValueFromStringTypeBuilder(typ externglib.Type, _string string) (externglib.Value, error) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 C.GType       // out
	var _arg2 *C.char       // out
	var _arg3 C.GValue      // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	_arg1 = (C.GType)(typ)
	_arg2 = (*C.char)(C.CString(_string))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_builder_value_from_string_type(_arg0, _arg1, _arg2, &_arg3, &_cerr)

	var _value externglib.Value // out
	var _goerr error            // out

	{
		var refTmpIn *C.GValue
		var refTmpOut *externglib.Value

		in0 := &_arg3
		refTmpIn = in0

		refTmpOut = externglib.ValueFromNative(unsafe.Pointer(refTmpIn))

		_value = *refTmpOut
	}
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _value, _goerr
}
