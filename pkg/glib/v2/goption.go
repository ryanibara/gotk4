// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <glib.h>
import "C"

// GType values.
var (
	GTypeOptionGroup = coreglib.Type(C.g_option_group_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeOptionGroup, F: marshalOptionGroup},
	})
}

// OptionArg enum values determine which type of extra argument the options
// expect to find. If an option expects an extra argument, it can be specified
// in several ways; with a short option: -x arg, with a long option: --name arg
// or combined in a single argument: --name=arg.
type OptionArg C.gint

const (
	// OptionArgNone: no extra argument. This is useful for simple flags.
	OptionArgNone OptionArg = iota
	// OptionArgString: option takes a UTF-8 string argument.
	OptionArgString
	// OptionArgInt: option takes an integer argument.
	OptionArgInt
	// OptionArgCallback: option provides a callback (of type ArgFunc) to parse
	// the extra argument.
	OptionArgCallback
	// OptionArgFilename: option takes a filename as argument, which will be in
	// the GLib filename encoding rather than UTF-8.
	OptionArgFilename
	// OptionArgStringArray: option takes a string argument, multiple uses of
	// the option are collected into an array of strings.
	OptionArgStringArray
	// OptionArgFilenameArray: option takes a filename as argument, multiple
	// uses of the option are collected into an array of strings.
	OptionArgFilenameArray
	// OptionArgDouble: option takes a double argument. The argument can be
	// formatted either for the user's locale or for the "C" locale. Since 2.12.
	OptionArgDouble
	// OptionArgInt64: option takes a 64-bit integer. Like G_OPTION_ARG_INT but
	// for larger numbers. The number can be in decimal base, or in hexadecimal
	// (when prefixed with 0x, for example, 0xffffffff). Since 2.12.
	OptionArgInt64
)

// String returns the name in string for OptionArg.
func (o OptionArg) String() string {
	switch o {
	case OptionArgNone:
		return "None"
	case OptionArgString:
		return "String"
	case OptionArgInt:
		return "Int"
	case OptionArgCallback:
		return "Callback"
	case OptionArgFilename:
		return "Filename"
	case OptionArgStringArray:
		return "StringArray"
	case OptionArgFilenameArray:
		return "FilenameArray"
	case OptionArgDouble:
		return "Double"
	case OptionArgInt64:
		return "Int64"
	default:
		return fmt.Sprintf("OptionArg(%d)", o)
	}
}

// OptionError: error codes returned by option parsing.
type OptionError C.gint

const (
	// OptionErrorUnknownOption: option was not known to the parser. This error
	// will only be reported, if the parser hasn't been instructed to ignore
	// unknown options, see g_option_context_set_ignore_unknown_options().
	OptionErrorUnknownOption OptionError = iota
	// OptionErrorBadValue: value couldn't be parsed.
	OptionErrorBadValue
	// OptionErrorFailed callback failed.
	OptionErrorFailed
)

// String returns the name in string for OptionError.
func (o OptionError) String() string {
	switch o {
	case OptionErrorUnknownOption:
		return "UnknownOption"
	case OptionErrorBadValue:
		return "BadValue"
	case OptionErrorFailed:
		return "Failed"
	default:
		return fmt.Sprintf("OptionError(%d)", o)
	}
}

// OptionFlags flags which modify individual options.
type OptionFlags C.guint

const (
	// OptionFlagNone: no flags. Since: 2.42.
	OptionFlagNone OptionFlags = 0b0
	// OptionFlagHidden: option doesn't appear in --help output.
	OptionFlagHidden OptionFlags = 0b1
	// OptionFlagInMain: option appears in the main section of the --help
	// output, even if it is defined in a group.
	OptionFlagInMain OptionFlags = 0b10
	// OptionFlagReverse: for options of the G_OPTION_ARG_NONE kind, this flag
	// indicates that the sense of the option is reversed.
	OptionFlagReverse OptionFlags = 0b100
	// OptionFlagNoArg: for options of the G_OPTION_ARG_CALLBACK kind, this flag
	// indicates that the callback does not take any argument (like a
	// G_OPTION_ARG_NONE option). Since 2.8.
	OptionFlagNoArg OptionFlags = 0b1000
	// OptionFlagFilename: for options of the G_OPTION_ARG_CALLBACK kind, this
	// flag indicates that the argument should be passed to the callback in the
	// GLib filename encoding rather than UTF-8. Since 2.8.
	OptionFlagFilename OptionFlags = 0b10000
	// OptionFlagOptionalArg: for options of the G_OPTION_ARG_CALLBACK kind,
	// this flag indicates that the argument supply is optional. If no argument
	// is given then data of GOptionParseFunc will be set to NULL. Since 2.8.
	OptionFlagOptionalArg OptionFlags = 0b100000
	// OptionFlagNoalias: this flag turns off the automatic conflict resolution
	// which prefixes long option names with groupname- if there is a conflict.
	// This option should only be used in situations where aliasing is necessary
	// to model some legacy commandline interface. It is not safe to use this
	// option, unless all option groups are under your direct control. Since
	// 2.8.
	OptionFlagNoalias OptionFlags = 0b1000000
)

// String returns the names in string for OptionFlags.
func (o OptionFlags) String() string {
	if o == 0 {
		return "OptionFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(141)

	for o != 0 {
		next := o & (o - 1)
		bit := o - next

		switch bit {
		case OptionFlagNone:
			builder.WriteString("None|")
		case OptionFlagHidden:
			builder.WriteString("Hidden|")
		case OptionFlagInMain:
			builder.WriteString("InMain|")
		case OptionFlagReverse:
			builder.WriteString("Reverse|")
		case OptionFlagNoArg:
			builder.WriteString("NoArg|")
		case OptionFlagFilename:
			builder.WriteString("Filename|")
		case OptionFlagOptionalArg:
			builder.WriteString("OptionalArg|")
		case OptionFlagNoalias:
			builder.WriteString("Noalias|")
		default:
			builder.WriteString(fmt.Sprintf("OptionFlags(0b%b)|", bit))
		}

		o = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if o contains other.
func (o OptionFlags) Has(other OptionFlags) bool {
	return (o & other) == other
}

// OptionEntry struct defines a single option. To have an effect, they must be
// added to a Group with g_option_context_add_main_entries() or
// g_option_group_add_entries().
//
// An instance of this type is always passed by reference.
type OptionEntry struct {
	*optionEntry
}

// optionEntry is the struct that's finalized.
type optionEntry struct {
	native *C.GOptionEntry
}

// LongName: long name of an option can be used to specify it in a commandline
// as --long_name. Every option must have a long name. To resolve conflicts if
// multiple option groups contain the same long name, it is also possible to
// specify the option as --groupname-long_name.
func (o *OptionEntry) LongName() string {
	valptr := &o.native.long_name
	var _v string // out
	_v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return _v
}

// ShortName: if an option has a short name, it can be specified -short_name in
// a commandline. short_name must be a printable ASCII character different from
// '-', or zero if the option has no short name.
func (o *OptionEntry) ShortName() byte {
	valptr := &o.native.short_name
	var _v byte // out
	_v = byte(*valptr)
	return _v
}

// Flags from Flags.
func (o *OptionEntry) Flags() int {
	valptr := &o.native.flags
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Arg: type of the option, as a Arg.
func (o *OptionEntry) Arg() OptionArg {
	valptr := &o.native.arg
	var _v OptionArg // out
	_v = OptionArg(*valptr)
	return _v
}

// ArgData: if the arg type is G_OPTION_ARG_CALLBACK, then arg_data must point
// to a ArgFunc callback function, which will be called to handle the extra
// argument. Otherwise, arg_data is a pointer to a location to store the value,
// the required type of the location depends on the arg type: -
// G_OPTION_ARG_NONE: gboolean - G_OPTION_ARG_STRING: gchar* - G_OPTION_ARG_INT:
// gint - G_OPTION_ARG_FILENAME: gchar* - G_OPTION_ARG_STRING_ARRAY: gchar** -
// G_OPTION_ARG_FILENAME_ARRAY: gchar** - G_OPTION_ARG_DOUBLE: gdouble If arg
// type is G_OPTION_ARG_STRING or G_OPTION_ARG_FILENAME, the location will
// contain a newly allocated string if the option was given. That string needs
// to be freed by the callee using g_free(). Likewise if arg type is
// G_OPTION_ARG_STRING_ARRAY or G_OPTION_ARG_FILENAME_ARRAY, the data should be
// freed using g_strfreev().
func (o *OptionEntry) ArgData() unsafe.Pointer {
	valptr := &o.native.arg_data
	var _v unsafe.Pointer // out
	_v = (unsafe.Pointer)(unsafe.Pointer(*valptr))
	return _v
}

// Description: description for the option in --help output. The description is
// translated using the translate_func of the group, see
// g_option_group_set_translation_domain().
func (o *OptionEntry) Description() string {
	valptr := &o.native.description
	var _v string // out
	_v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return _v
}

// ArgDescription: placeholder to use for the extra argument parsed by the
// option in --help output. The arg_description is translated using the
// translate_func of the group, see g_option_group_set_translation_domain().
func (o *OptionEntry) ArgDescription() string {
	valptr := &o.native.arg_description
	var _v string // out
	_v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return _v
}

// ShortName: if an option has a short name, it can be specified -short_name in
// a commandline. short_name must be a printable ASCII character different from
// '-', or zero if the option has no short name.
func (o *OptionEntry) SetShortName(shortName byte) {
	valptr := &o.native.short_name
	*valptr = C.gchar(shortName)
}

// Flags from Flags.
func (o *OptionEntry) SetFlags(flags int) {
	valptr := &o.native.flags
	*valptr = C.gint(flags)
}

// OptionGroup: GOptionGroup struct defines the options in a single group. The
// struct has only private fields and should not be directly accessed.
//
// All options in a group share the same translation function. Libraries which
// need to parse commandline options are expected to provide a function for
// getting a GOptionGroup holding their options, which the application can then
// add to its Context.
//
// An instance of this type is always passed by reference.
type OptionGroup struct {
	*optionGroup
}

// optionGroup is the struct that's finalized.
type optionGroup struct {
	native *C.GOptionGroup
}

func marshalOptionGroup(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &OptionGroup{&optionGroup{(*C.GOptionGroup)(b)}}, nil
}

// AddEntries adds the options specified in entries to group.
//
// The function takes the following parameters:
//
//    - entries: NULL-terminated array of Entrys.
//
func (group *OptionGroup) AddEntries(entries []OptionEntry) {
	var _arg0 *C.GOptionGroup // out
	var _arg1 *C.GOptionEntry // out

	_arg0 = (*C.GOptionGroup)(gextras.StructNative(unsafe.Pointer(group)))
	{
		_arg1 = (*C.GOptionEntry)(C.calloc(C.size_t((len(entries) + 1)), C.size_t(C.sizeof_GOptionEntry)))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(entries)+1)
			var zero C.GOptionEntry
			out[len(entries)] = zero
			for i := range entries {
				out[i] = *(*C.GOptionEntry)(gextras.StructNative(unsafe.Pointer((&entries[i]))))
			}
		}
	}

	C.g_option_group_add_entries(_arg0, _arg1)
	runtime.KeepAlive(group)
	runtime.KeepAlive(entries)
}

// SetTranslationDomain: convenience function to use gettext() for translating
// user-visible strings.
//
// The function takes the following parameters:
//
//    - domain to use.
//
func (group *OptionGroup) SetTranslationDomain(domain string) {
	var _arg0 *C.GOptionGroup // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GOptionGroup)(gextras.StructNative(unsafe.Pointer(group)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_option_group_set_translation_domain(_arg0, _arg1)
	runtime.KeepAlive(group)
	runtime.KeepAlive(domain)
}
