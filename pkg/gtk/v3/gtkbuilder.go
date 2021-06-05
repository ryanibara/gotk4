// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gobject/v2"
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
		{T: externglib.Type(C.gtk_builder_get_type()), F: marshalBuilder},
	})
}

// Builder: a GtkBuilder is an auxiliary object that reads textual descriptions
// of a user interface and instantiates the described objects. To create a
// GtkBuilder from a user interface description, call
// gtk_builder_new_from_file(), gtk_builder_new_from_resource() or
// gtk_builder_new_from_string().
//
// In the (unusual) case that you want to add user interface descriptions from
// multiple sources to the same GtkBuilder you can call gtk_builder_new() to get
// an empty builder and populate it by (multiple) calls to
// gtk_builder_add_from_file(), gtk_builder_add_from_resource() or
// gtk_builder_add_from_string().
//
// A GtkBuilder holds a reference to all objects that it has constructed and
// drops these references when it is finalized. This finalization can cause the
// destruction of non-widget objects or widgets which are not contained in a
// toplevel window. For toplevel windows constructed by a builder, it is the
// responsibility of the user to call gtk_widget_destroy() to get rid of them
// and all the widgets they contain.
//
// The functions gtk_builder_get_object() and gtk_builder_get_objects() can be
// used to access the widgets in the interface by the names assigned to them
// inside the UI description. Toplevel windows returned by these functions will
// stay around until the user explicitly destroys them with
// gtk_widget_destroy(). Other widgets will either be part of a larger hierarchy
// constructed by the builder (in which case you should not have to worry about
// their lifecycle), or without a parent, in which case they have to be added to
// some container to make use of them. Non-widget objects need to be reffed with
// g_object_ref() to keep them beyond the lifespan of the builder.
//
// The function gtk_builder_connect_signals() and variants thereof can be used
// to connect handlers to the named signals in the description.
//
//
// GtkBuilder UI Definitions
//
// GtkBuilder parses textual descriptions of user interfaces which are specified
// in an XML format which can be roughly described by the RELAX NG schema below.
// We refer to these descriptions as “GtkBuilder UI definitions” or just “UI
// definitions” if the context is clear. Do not confuse GtkBuilder UI
// Definitions with [GtkUIManager UI Definitions][XML-UI], which are more
// limited in scope. It is common to use `.ui` as the filename extension for
// files containing GtkBuilder UI definitions.
//
// RELAX NG Compact Syntax
// (https://gitlab.gnome.org/GNOME/gtk/-/blob/gtk-3-24/gtk/gtkbuilder.rnc)
//
// The toplevel element is <interface>. It optionally takes a “domain”
// attribute, which will make the builder look for translated strings using
// dgettext() in the domain specified. This can also be done by calling
// gtk_builder_set_translation_domain() on the builder. Objects are described by
// <object> elements, which can contain <property> elements to set properties,
// <signal> elements which connect signals to handlers, and <child> elements,
// which describe child objects (most often widgets inside a container, but also
// e.g. actions in an action group, or columns in a tree model). A <child>
// element contains an <object> element which describes the child object. The
// target toolkit version(s) are described by <requires> elements, the “lib”
// attribute specifies the widget library in question (currently the only
// supported value is “gtk+”) and the “version” attribute specifies the target
// version in the form “<major>.<minor>”. The builder will error out if the
// version requirements are not met.
//
// Typically, the specific kind of object represented by an <object> element is
// specified by the “class” attribute. If the type has not been loaded yet, GTK+
// tries to find the get_type() function from the class name by applying
// heuristics. This works in most cases, but if necessary, it is possible to
// specify the name of the get_type() function explictly with the "type-func"
// attribute. As a special case, GtkBuilder allows to use an object that has
// been constructed by a UIManager in another part of the UI definition by
// specifying the id of the UIManager in the “constructor” attribute and the
// name of the object in the “id” attribute.
//
// Objects may be given a name with the “id” attribute, which allows the
// application to retrieve them from the builder with gtk_builder_get_object().
// An id is also necessary to use the object as property value in other parts of
// the UI definition. GTK+ reserves ids starting and ending with ___ (3
// underscores) for its own purposes.
//
// Setting properties of objects is pretty straightforward with the <property>
// element: the “name” attribute specifies the name of the property, and the
// content of the element specifies the value. If the “translatable” attribute
// is set to a true value, GTK+ uses gettext() (or dgettext() if the builder has
// a translation domain set) to find a translation for the value. This happens
// before the value is parsed, so it can be used for properties of any type, but
// it is probably most useful for string properties. It is also possible to
// specify a context to disambiguate short strings, and comments which may help
// the translators.
//
// GtkBuilder can parse textual representations for the most common property
// types: characters, strings, integers, floating-point numbers, booleans
// (strings like “TRUE”, “t”, “yes”, “y”, “1” are interpreted as true, strings
// like “FALSE”, “f”, “no”, “n”, “0” are interpreted as false), enumerations
// (can be specified by their name, nick or integer value), flags (can be
// specified by their name, nick, integer value, optionally combined with “|”,
// e.g. “GTK_VISIBLE|GTK_REALIZED”) and colors (in a format understood by
// gdk_rgba_parse()).
//
// GVariants can be specified in the format understood by g_variant_parse(), and
// pixbufs can be specified as a filename of an image file to load.
//
// Objects can be referred to by their name and by default refer to objects
// declared in the local xml fragment and objects exposed via
// gtk_builder_expose_object(). In general, GtkBuilder allows forward references
// to objects — declared in the local xml; an object doesn’t have to be
// constructed before it can be referred to. The exception to this rule is that
// an object has to be constructed before it can be used as the value of a
// construct-only property.
//
// It is also possible to bind a property value to another object's property
// value using the attributes "bind-source" to specify the source object of the
// binding, "bind-property" to specify the source property and optionally
// "bind-flags" to specify the binding flags. Internally builder implements this
// using GBinding objects. For more information see g_object_bind_property()
//
// Signal handlers are set up with the <signal> element. The “name” attribute
// specifies the name of the signal, and the “handler” attribute specifies the
// function to connect to the signal. By default, GTK+ tries to find the handler
// using g_module_symbol(), but this can be changed by passing a custom
// BuilderConnectFunc to gtk_builder_connect_signals_full(). The remaining
// attributes, “after”, “swapped” and “object”, have the same meaning as the
// corresponding parameters of the g_signal_connect_object() or
// g_signal_connect_data() functions. A “last_modification_time” attribute is
// also allowed, but it does not have a meaning to the builder.
//
// Sometimes it is necessary to refer to widgets which have implicitly been
// constructed by GTK+ as part of a composite widget, to set properties on them
// or to add further children (e.g. the @vbox of a Dialog). This can be achieved
// by setting the “internal-child” property of the <child> element to a true
// value. Note that GtkBuilder still requires an <object> element for the
// internal child, even if it has already been constructed.
//
// A number of widgets have different places where a child can be added (e.g.
// tabs vs. page content in notebooks). This can be reflected in a UI definition
// by specifying the “type” attribute on a <child> The possible values for the
// “type” attribute are described in the sections describing the widget-specific
// portions of UI definitions.
//
// A GtkBuilder UI Definition
//
//    <interface>
//      <object class="GtkDialog" id="dialog1">
//        <child internal-child="vbox">
//          <object class="GtkBox" id="vbox1">
//            <property name="border-width">10</property>
//            <child internal-child="action_area">
//              <object class="GtkButtonBox" id="hbuttonbox1">
//                <property name="border-width">20</property>
//                <child>
//                  <object class="GtkButton" id="ok_button">
//                    <property name="label">gtk-ok</property>
//                    <property name="use-stock">TRUE</property>
//                    <signal name="clicked" handler="ok_button_clicked"/>
//                  </object>
//                </child>
//              </object>
//            </child>
//          </object>
//        </child>
//      </object>
//    </interface>
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
// Additionally, since 3.10 a special <template> tag has been added to the
// format allowing one to define a widget class’s components. See the [GtkWidget
// documentation][composite-templates] for details.
type Builder interface {
	gextras.Objector

	// AddFromFile parses a file containing a [GtkBuilder UI
	// definition][BUILDER-UI] and merges it with the current contents of
	// @builder.
	//
	// Most users will probably want to use gtk_builder_new_from_file().
	//
	// If an error occurs, 0 will be returned and @error will be assigned a
	// #GError from the K_BUILDER_ERROR, MARKUP_ERROR or FILE_ERROR domain.
	//
	// It’s not really reasonable to attempt to handle failures of this call.
	// You should not use this function with untrusted files (ie: files that are
	// not part of your application). Broken Builder files can easily crash your
	// program, and it’s possible that memory was leaked leading up to the
	// reported failure. The only reasonable thing to do when an error is
	// detected is to call g_error().
	AddFromFile(filename string) (guint uint, err error)
	// AddFromResource parses a resource file containing a [GtkBuilder UI
	// definition][BUILDER-UI] and merges it with the current contents of
	// @builder.
	//
	// Most users will probably want to use gtk_builder_new_from_resource().
	//
	// If an error occurs, 0 will be returned and @error will be assigned a
	// #GError from the K_BUILDER_ERROR, MARKUP_ERROR or RESOURCE_ERROR domain.
	//
	// It’s not really reasonable to attempt to handle failures of this call.
	// The only reasonable thing to do when an error is detected is to call
	// g_error().
	AddFromResource(resourcePath string) (guint uint, err error)
	// AddFromString parses a string containing a [GtkBuilder UI
	// definition][BUILDER-UI] and merges it with the current contents of
	// @builder.
	//
	// Most users will probably want to use gtk_builder_new_from_string().
	//
	// Upon errors 0 will be returned and @error will be assigned a #GError from
	// the K_BUILDER_ERROR, MARKUP_ERROR or VARIANT_PARSE_ERROR domain.
	//
	// It’s not really reasonable to attempt to handle failures of this call.
	// The only reasonable thing to do when an error is detected is to call
	// g_error().
	AddFromString(buffer string, length uint) (guint uint, err error)
	// AddObjectsFromFile parses a file containing a [GtkBuilder UI
	// definition][BUILDER-UI] building only the requested objects and merges
	// them with the current contents of @builder.
	//
	// Upon errors 0 will be returned and @error will be assigned a #GError from
	// the K_BUILDER_ERROR, MARKUP_ERROR or FILE_ERROR domain.
	//
	// If you are adding an object that depends on an object that is not its
	// child (for instance a TreeView that depends on its TreeModel), you have
	// to explicitly list all of them in @object_ids.
	AddObjectsFromFile(filename string, objectIds []string) (guint uint, err error)
	// AddObjectsFromResource parses a resource file containing a [GtkBuilder UI
	// definition][BUILDER-UI] building only the requested objects and merges
	// them with the current contents of @builder.
	//
	// Upon errors 0 will be returned and @error will be assigned a #GError from
	// the K_BUILDER_ERROR, MARKUP_ERROR or RESOURCE_ERROR domain.
	//
	// If you are adding an object that depends on an object that is not its
	// child (for instance a TreeView that depends on its TreeModel), you have
	// to explicitly list all of them in @object_ids.
	AddObjectsFromResource(resourcePath string, objectIds []string) (guint uint, err error)
	// AddObjectsFromString parses a string containing a [GtkBuilder UI
	// definition][BUILDER-UI] building only the requested objects and merges
	// them with the current contents of @builder.
	//
	// Upon errors 0 will be returned and @error will be assigned a #GError from
	// the K_BUILDER_ERROR or MARKUP_ERROR domain.
	//
	// If you are adding an object that depends on an object that is not its
	// child (for instance a TreeView that depends on its TreeModel), you have
	// to explicitly list all of them in @object_ids.
	AddObjectsFromString(buffer string, length uint, objectIds []string) (guint uint, err error)
	// ConnectSignals: this method is a simpler variation of
	// gtk_builder_connect_signals_full(). It uses symbols explicitly added to
	// @builder with prior calls to gtk_builder_add_callback_symbol(). In the
	// case that symbols are not explicitly added; it uses #GModule’s
	// introspective features (by opening the module nil) to look at the
	// application’s symbol table. From here it tries to match the signal
	// handler names given in the interface description with symbols in the
	// application and connects the signals. Note that this function can only be
	// called once, subsequent calls will do nothing.
	//
	// Note that unless gtk_builder_add_callback_symbol() is called for all
	// signal callbacks which are referenced by the loaded XML, this function
	// will require that #GModule be supported on the platform.
	//
	// If you rely on #GModule support to lookup callbacks in the symbol table,
	// the following details should be noted:
	//
	// When compiling applications for Windows, you must declare signal
	// callbacks with MODULE_EXPORT, or they will not be put in the symbol
	// table. On Linux and Unices, this is not necessary; applications should
	// instead be compiled with the -Wl,--export-dynamic CFLAGS, and linked
	// against gmodule-export-2.0.
	ConnectSignals(userData interface{})
	// ConnectSignalsFull: this function can be thought of the interpreted
	// language binding version of gtk_builder_connect_signals(), except that it
	// does not require GModule to function correctly.
	ConnectSignalsFull(fn BuilderConnectFunc)
	// ExposeObject: add @object to the @builder object pool so it can be
	// referenced just like any other object built by builder.
	ExposeObject(name string, object gextras.Objector)
	// ExtendWithTemplate: main private entry point for building composite
	// container components from template XML.
	//
	// This is exported purely to let gtk-builder-tool validate templates,
	// applications have no need to call this function.
	ExtendWithTemplate(widget Widget, templateType externglib.Type, buffer string, length uint) (guint uint, err error)
	// Application gets the Application associated with the builder.
	//
	// The Application is used for creating action proxies as requested from XML
	// that the builder is loading.
	//
	// By default, the builder uses the default application: the one from
	// g_application_get_default(). If you want to use another application for
	// constructing proxies, use gtk_builder_set_application().
	Application() Application
	// Object gets the object named @name. Note that this function does not
	// increment the reference count of the returned object.
	Object(name string) gextras.Objector
	// Objects gets all objects that have been constructed by @builder. Note
	// that this function does not increment the reference counts of the
	// returned objects.
	Objects() *glib.SList
	// TranslationDomain gets the translation domain of @builder.
	TranslationDomain() string
	// TypeFromName looks up a type by name, using the virtual function that
	// Builder has for that purpose. This is mainly used when implementing the
	// Buildable interface on a type.
	TypeFromName(typeName string) externglib.Type
	// SetApplication sets the application associated with @builder.
	//
	// You only need this function if there is more than one #GApplication in
	// your process. @application cannot be nil.
	SetApplication(application Application)
	// SetTranslationDomain sets the translation domain of @builder. See
	// Builder:translation-domain.
	SetTranslationDomain(domain string)
	// ValueFromString: this function demarshals a value from a string. This
	// function calls g_value_init() on the @value argument, so it need not be
	// initialised beforehand.
	//
	// This function can handle char, uchar, boolean, int, uint, long, ulong,
	// enum, flags, float, double, string, Color, RGBA and Adjustment type
	// values. Support for Widget type values is still to come.
	//
	// Upon errors false will be returned and @error will be assigned a #GError
	// from the K_BUILDER_ERROR domain.
	ValueFromString(pspec gobject.ParamSpec, string string) (value externglib.Value, err error)
	// ValueFromStringType: like gtk_builder_value_from_string(), this function
	// demarshals a value from a string, but takes a #GType instead of Spec.
	// This function calls g_value_init() on the @value argument, so it need not
	// be initialised beforehand.
	//
	// Upon errors false will be returned and @error will be assigned a #GError
	// from the K_BUILDER_ERROR domain.
	ValueFromStringType(typ externglib.Type, string string) (value externglib.Value, err error)
}

// builder implements the Builder interface.
type builder struct {
	gextras.Objector
}

var _ Builder = (*builder)(nil)

// WrapBuilder wraps a GObject to the right type. It is
// primarily used internally.
func WrapBuilder(obj *externglib.Object) Builder {
	return Builder{
		Objector: obj,
	}
}

func marshalBuilder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBuilder(obj), nil
}

// NewBuilder constructs a class Builder.
func NewBuilder() Builder {
	var cret C.GtkBuilder
	var ret1 Builder

	cret = C.gtk_builder_new()

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Builder)

	return ret1
}

// NewBuilderFromFile constructs a class Builder.
func NewBuilderFromFile(filename string) Builder {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkBuilder
	var ret1 Builder

	cret = C.gtk_builder_new_from_file(filename)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Builder)

	return ret1
}

// NewBuilderFromResource constructs a class Builder.
func NewBuilderFromResource(resourcePath string) Builder {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkBuilder
	var ret1 Builder

	cret = C.gtk_builder_new_from_resource(resourcePath)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Builder)

	return ret1
}

// NewBuilderFromString constructs a class Builder.
func NewBuilderFromString(string string, length int) Builder {
	var arg1 *C.gchar
	var arg2 C.gssize

	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gssize(length)

	var cret C.GtkBuilder
	var ret1 Builder

	cret = C.gtk_builder_new_from_string(string, length)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Builder)

	return ret1
}

// AddFromFile parses a file containing a [GtkBuilder UI
// definition][BUILDER-UI] and merges it with the current contents of
// @builder.
//
// Most users will probably want to use gtk_builder_new_from_file().
//
// If an error occurs, 0 will be returned and @error will be assigned a
// #GError from the K_BUILDER_ERROR, MARKUP_ERROR or FILE_ERROR domain.
//
// It’s not really reasonable to attempt to handle failures of this call.
// You should not use this function with untrusted files (ie: files that are
// not part of your application). Broken Builder files can easily crash your
// program, and it’s possible that memory was leaked leading up to the
// reported failure. The only reasonable thing to do when an error is
// detected is to call g_error().
func (b builder) AddFromFile(filename string) (guint uint, err error) {
	var arg0 *C.GtkBuilder
	var arg1 *C.gchar

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	var errout *C.GError
	var goerr error
	var cret C.guint
	var ret2 uint

	cret = C.gtk_builder_add_from_file(arg0, filename, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = C.guint(cret)

	return goerr, ret2
}

// AddFromResource parses a resource file containing a [GtkBuilder UI
// definition][BUILDER-UI] and merges it with the current contents of
// @builder.
//
// Most users will probably want to use gtk_builder_new_from_resource().
//
// If an error occurs, 0 will be returned and @error will be assigned a
// #GError from the K_BUILDER_ERROR, MARKUP_ERROR or RESOURCE_ERROR domain.
//
// It’s not really reasonable to attempt to handle failures of this call.
// The only reasonable thing to do when an error is detected is to call
// g_error().
func (b builder) AddFromResource(resourcePath string) (guint uint, err error) {
	var arg0 *C.GtkBuilder
	var arg1 *C.gchar

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(arg1))

	var errout *C.GError
	var goerr error
	var cret C.guint
	var ret2 uint

	cret = C.gtk_builder_add_from_resource(arg0, resourcePath, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = C.guint(cret)

	return goerr, ret2
}

// AddFromString parses a string containing a [GtkBuilder UI
// definition][BUILDER-UI] and merges it with the current contents of
// @builder.
//
// Most users will probably want to use gtk_builder_new_from_string().
//
// Upon errors 0 will be returned and @error will be assigned a #GError from
// the K_BUILDER_ERROR, MARKUP_ERROR or VARIANT_PARSE_ERROR domain.
//
// It’s not really reasonable to attempt to handle failures of this call.
// The only reasonable thing to do when an error is detected is to call
// g_error().
func (b builder) AddFromString(buffer string, length uint) (guint uint, err error) {
	var arg0 *C.GtkBuilder
	var arg1 *C.gchar
	var arg2 C.gsize

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(buffer))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gsize(length)

	var errout *C.GError
	var goerr error
	var cret C.guint
	var ret2 uint

	cret = C.gtk_builder_add_from_string(arg0, buffer, length, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = C.guint(cret)

	return goerr, ret2
}

// AddObjectsFromFile parses a file containing a [GtkBuilder UI
// definition][BUILDER-UI] building only the requested objects and merges
// them with the current contents of @builder.
//
// Upon errors 0 will be returned and @error will be assigned a #GError from
// the K_BUILDER_ERROR, MARKUP_ERROR or FILE_ERROR domain.
//
// If you are adding an object that depends on an object that is not its
// child (for instance a TreeView that depends on its TreeModel), you have
// to explicitly list all of them in @object_ids.
func (b builder) AddObjectsFromFile(filename string, objectIds []string) (guint uint, err error) {
	var arg0 *C.GtkBuilder
	var arg1 *C.gchar
	var arg2 **C.gchar

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.malloc(len(objectIds) * (unsafe.Sizeof(int(0)) + 1))
	defer C.free(unsafe.Pointer(arg2))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg2), int(len(objectIds)))

		for i := range objectIds {
			out[i] = (*C.gchar)(C.CString(objectIds[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	var errout *C.GError
	var goerr error
	var cret C.guint
	var ret2 uint

	cret = C.gtk_builder_add_objects_from_file(arg0, filename, objectIds, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = C.guint(cret)

	return goerr, ret2
}

// AddObjectsFromResource parses a resource file containing a [GtkBuilder UI
// definition][BUILDER-UI] building only the requested objects and merges
// them with the current contents of @builder.
//
// Upon errors 0 will be returned and @error will be assigned a #GError from
// the K_BUILDER_ERROR, MARKUP_ERROR or RESOURCE_ERROR domain.
//
// If you are adding an object that depends on an object that is not its
// child (for instance a TreeView that depends on its TreeModel), you have
// to explicitly list all of them in @object_ids.
func (b builder) AddObjectsFromResource(resourcePath string, objectIds []string) (guint uint, err error) {
	var arg0 *C.GtkBuilder
	var arg1 *C.gchar
	var arg2 **C.gchar

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.malloc(len(objectIds) * (unsafe.Sizeof(int(0)) + 1))
	defer C.free(unsafe.Pointer(arg2))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg2), int(len(objectIds)))

		for i := range objectIds {
			out[i] = (*C.gchar)(C.CString(objectIds[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	var errout *C.GError
	var goerr error
	var cret C.guint
	var ret2 uint

	cret = C.gtk_builder_add_objects_from_resource(arg0, resourcePath, objectIds, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = C.guint(cret)

	return goerr, ret2
}

// AddObjectsFromString parses a string containing a [GtkBuilder UI
// definition][BUILDER-UI] building only the requested objects and merges
// them with the current contents of @builder.
//
// Upon errors 0 will be returned and @error will be assigned a #GError from
// the K_BUILDER_ERROR or MARKUP_ERROR domain.
//
// If you are adding an object that depends on an object that is not its
// child (for instance a TreeView that depends on its TreeModel), you have
// to explicitly list all of them in @object_ids.
func (b builder) AddObjectsFromString(buffer string, length uint, objectIds []string) (guint uint, err error) {
	var arg0 *C.GtkBuilder
	var arg1 *C.gchar
	var arg2 C.gsize
	var arg3 **C.gchar

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(buffer))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gsize(length)
	arg3 = C.malloc(len(objectIds) * (unsafe.Sizeof(int(0)) + 1))
	defer C.free(unsafe.Pointer(arg3))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg3), int(len(objectIds)))

		for i := range objectIds {
			out[i] = (*C.gchar)(C.CString(objectIds[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	var errout *C.GError
	var goerr error
	var cret C.guint
	var ret2 uint

	cret = C.gtk_builder_add_objects_from_string(arg0, buffer, length, objectIds, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = C.guint(cret)

	return goerr, ret2
}

// ConnectSignals: this method is a simpler variation of
// gtk_builder_connect_signals_full(). It uses symbols explicitly added to
// @builder with prior calls to gtk_builder_add_callback_symbol(). In the
// case that symbols are not explicitly added; it uses #GModule’s
// introspective features (by opening the module nil) to look at the
// application’s symbol table. From here it tries to match the signal
// handler names given in the interface description with symbols in the
// application and connects the signals. Note that this function can only be
// called once, subsequent calls will do nothing.
//
// Note that unless gtk_builder_add_callback_symbol() is called for all
// signal callbacks which are referenced by the loaded XML, this function
// will require that #GModule be supported on the platform.
//
// If you rely on #GModule support to lookup callbacks in the symbol table,
// the following details should be noted:
//
// When compiling applications for Windows, you must declare signal
// callbacks with MODULE_EXPORT, or they will not be put in the symbol
// table. On Linux and Unices, this is not necessary; applications should
// instead be compiled with the -Wl,--export-dynamic CFLAGS, and linked
// against gmodule-export-2.0.
func (b builder) ConnectSignals(userData interface{}) {
	var arg0 *C.GtkBuilder
	var arg1 C.gpointer

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = C.gpointer(userData)

	C.gtk_builder_connect_signals(arg0, userData)
}

// ConnectSignalsFull: this function can be thought of the interpreted
// language binding version of gtk_builder_connect_signals(), except that it
// does not require GModule to function correctly.
func (b builder) ConnectSignalsFull(fn BuilderConnectFunc) {
	var arg0 *C.GtkBuilder

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))

	C.gtk_builder_connect_signals_full(arg0, fn, userData)
}

// ExposeObject: add @object to the @builder object pool so it can be
// referenced just like any other object built by builder.
func (b builder) ExposeObject(name string, object gextras.Objector) {
	var arg0 *C.GtkBuilder
	var arg1 *C.gchar
	var arg2 *C.GObject

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GObject)(unsafe.Pointer(object.Native()))

	C.gtk_builder_expose_object(arg0, name, object)
}

// ExtendWithTemplate: main private entry point for building composite
// container components from template XML.
//
// This is exported purely to let gtk-builder-tool validate templates,
// applications have no need to call this function.
func (b builder) ExtendWithTemplate(widget Widget, templateType externglib.Type, buffer string, length uint) (guint uint, err error) {
	var arg0 *C.GtkBuilder
	var arg1 *C.GtkWidget
	var arg2 C.GType
	var arg3 *C.gchar
	var arg4 C.gsize

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 := C.GType(templateType)
	arg3 = (*C.gchar)(C.CString(buffer))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.gsize(length)

	var errout *C.GError
	var goerr error
	var cret C.guint
	var ret2 uint

	cret = C.gtk_builder_extend_with_template(arg0, widget, templateType, buffer, length, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = C.guint(cret)

	return goerr, ret2
}

// Application gets the Application associated with the builder.
//
// The Application is used for creating action proxies as requested from XML
// that the builder is loading.
//
// By default, the builder uses the default application: the one from
// g_application_get_default(). If you want to use another application for
// constructing proxies, use gtk_builder_set_application().
func (b builder) Application() Application {
	var arg0 *C.GtkBuilder

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))

	var cret *C.GtkApplication
	var ret1 Application

	cret = C.gtk_builder_get_application(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Application)

	return ret1
}

// Object gets the object named @name. Note that this function does not
// increment the reference count of the returned object.
func (b builder) Object(name string) gextras.Objector {
	var arg0 *C.GtkBuilder
	var arg1 *C.gchar

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GObject
	var ret1 gextras.Objector

	cret = C.gtk_builder_get_object(arg0, name)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gextras.Objector)

	return ret1
}

// Objects gets all objects that have been constructed by @builder. Note
// that this function does not increment the reference counts of the
// returned objects.
func (b builder) Objects() *glib.SList {
	var arg0 *C.GtkBuilder

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))

	var cret *C.GSList
	var ret1 *glib.SList

	cret = C.gtk_builder_get_objects(arg0)

	ret1 = glib.WrapSList(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *glib.SList) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// TranslationDomain gets the translation domain of @builder.
func (b builder) TranslationDomain() string {
	var arg0 *C.GtkBuilder

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_builder_get_translation_domain(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// TypeFromName looks up a type by name, using the virtual function that
// Builder has for that purpose. This is mainly used when implementing the
// Buildable interface on a type.
func (b builder) TypeFromName(typeName string) externglib.Type {
	var arg0 *C.GtkBuilder
	var arg1 *C.char

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(typeName))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GType
	var ret1 externglib.Type

	cret = C.gtk_builder_get_type_from_name(arg0, typeName)

	ret1 = externglib.Type(cret)

	return ret1
}

// SetApplication sets the application associated with @builder.
//
// You only need this function if there is more than one #GApplication in
// your process. @application cannot be nil.
func (b builder) SetApplication(application Application) {
	var arg0 *C.GtkBuilder
	var arg1 *C.GtkApplication

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))

	C.gtk_builder_set_application(arg0, application)
}

// SetTranslationDomain sets the translation domain of @builder. See
// Builder:translation-domain.
func (b builder) SetTranslationDomain(domain string) {
	var arg0 *C.GtkBuilder
	var arg1 *C.gchar

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_builder_set_translation_domain(arg0, domain)
}

// ValueFromString: this function demarshals a value from a string. This
// function calls g_value_init() on the @value argument, so it need not be
// initialised beforehand.
//
// This function can handle char, uchar, boolean, int, uint, long, ulong,
// enum, flags, float, double, string, Color, RGBA and Adjustment type
// values. Support for Widget type values is still to come.
//
// Upon errors false will be returned and @error will be assigned a #GError
// from the K_BUILDER_ERROR domain.
func (b builder) ValueFromString(pspec gobject.ParamSpec, string string) (value externglib.Value, err error) {
	var arg0 *C.GtkBuilder
	var arg1 *C.GParamSpec
	var arg2 *C.gchar

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))
	arg2 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg2))

	var arg3 C.GValue
	var ret3 *externglib.Value
	var errout *C.GError
	var goerr error

	C.gtk_builder_value_from_string(arg0, pspec, string, &arg3, &errout)

	*ret3 = externglib.ValueFromNative(unsafe.Pointer(arg3))
	goerr = gerror.Take(unsafe.Pointer(errout))

	return ret3, goerr
}

// ValueFromStringType: like gtk_builder_value_from_string(), this function
// demarshals a value from a string, but takes a #GType instead of Spec.
// This function calls g_value_init() on the @value argument, so it need not
// be initialised beforehand.
//
// Upon errors false will be returned and @error will be assigned a #GError
// from the K_BUILDER_ERROR domain.
func (b builder) ValueFromStringType(typ externglib.Type, string string) (value externglib.Value, err error) {
	var arg0 *C.GtkBuilder
	var arg1 C.GType
	var arg2 *C.gchar

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 := C.GType(typ)
	arg2 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg2))

	var arg3 C.GValue
	var ret3 *externglib.Value
	var errout *C.GError
	var goerr error

	C.gtk_builder_value_from_string_type(arg0, typ, string, &arg3, &errout)

	*ret3 = externglib.ValueFromNative(unsafe.Pointer(arg3))
	goerr = gerror.Take(unsafe.Pointer(errout))

	return ret3, goerr
}

type BuilderPrivate struct {
	native C.GtkBuilderPrivate
}

// WrapBuilderPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBuilderPrivate(ptr unsafe.Pointer) *BuilderPrivate {
	if ptr == nil {
		return nil
	}

	return (*BuilderPrivate)(ptr)
}

func marshalBuilderPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapBuilderPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (b *BuilderPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}
