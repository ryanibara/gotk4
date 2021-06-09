// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_uri_get_type()), F: marshalURI},
	})
}

// URIError: error codes returned by #GUri methods.
type URIError int

const (
	// URIErrorFailed: generic error if no more specific error is available. See
	// the error message for details.
	URIErrorFailed URIError = 0
	// URIErrorBadScheme: the scheme of a URI could not be parsed.
	URIErrorBadScheme URIError = 1
	// URIErrorBadUser: the user/userinfo of a URI could not be parsed.
	URIErrorBadUser URIError = 2
	// URIErrorBadPassword: the password of a URI could not be parsed.
	URIErrorBadPassword URIError = 3
	// URIErrorBadAuthParams: the authentication parameters of a URI could not
	// be parsed.
	URIErrorBadAuthParams URIError = 4
	// URIErrorBadHost: the host of a URI could not be parsed.
	URIErrorBadHost URIError = 5
	// URIErrorBadPort: the port of a URI could not be parsed.
	URIErrorBadPort URIError = 6
	// URIErrorBadPath: the path of a URI could not be parsed.
	URIErrorBadPath URIError = 7
	// URIErrorBadQuery: the query of a URI could not be parsed.
	URIErrorBadQuery URIError = 8
	// URIErrorBadFragment: the fragment of a URI could not be parsed.
	URIErrorBadFragment URIError = 9
)

// URIFlags flags that describe a URI.
//
// When parsing a URI, if you need to choose different flags based on the type
// of URI, you can use g_uri_peek_scheme() on the URI string to check the scheme
// first, and use that to decide what flags to parse it with.
type URIFlags int

const (
	// URIFlagsNone: no flags set.
	URIFlagsNone URIFlags = 0
	// URIFlagsParseRelaxed: parse the URI more relaxedly than the RFC 3986
	// (https://tools.ietf.org/html/rfc3986) grammar specifies, fixing up or
	// ignoring common mistakes in URIs coming from external sources. This is
	// also needed for some obscure URI schemes where `;` separates the host
	// from the path. Don’t use this flag unless you need to.
	URIFlagsParseRelaxed URIFlags = 1
	// URIFlagsHasPassword: the userinfo field may contain a password, which
	// will be separated from the username by `:`.
	URIFlagsHasPassword URIFlags = 2
	// URIFlagsHasAuthParams: the userinfo may contain additional
	// authentication-related parameters, which will be separated from the
	// username and/or password by `;`.
	URIFlagsHasAuthParams URIFlags = 4
	// URIFlagsEncoded: when parsing a URI, this indicates that `%`-encoded
	// characters in the userinfo, path, query, and fragment fields should not
	// be decoded. (And likewise the host field if G_URI_FLAGS_NON_DNS is also
	// set.) When building a URI, it indicates that you have already `%`-encoded
	// the components, and so #GUri should not do any encoding itself.
	URIFlagsEncoded URIFlags = 8
	// URIFlagsNonDns: the host component should not be assumed to be a DNS
	// hostname or IP address (for example, for `smb` URIs with NetBIOS
	// hostnames).
	URIFlagsNonDns URIFlags = 16
	// URIFlagsEncodedQuery: same as G_URI_FLAGS_ENCODED, for the query field
	// only.
	URIFlagsEncodedQuery URIFlags = 32
	// URIFlagsEncodedPath: same as G_URI_FLAGS_ENCODED, for the path only.
	URIFlagsEncodedPath URIFlags = 64
	// URIFlagsEncodedFragment: same as G_URI_FLAGS_ENCODED, for the fragment
	// only.
	URIFlagsEncodedFragment URIFlags = 128
	// URIFlagsSchemeNormalize: a scheme-based normalization will be applied.
	// For example, when parsing an HTTP URI changing omitted path to `/` and
	// omitted port to `80`; and when building a URI, changing empty path to `/`
	// and default port `80`). This only supports a subset of known schemes.
	// (Since: 2.68)
	URIFlagsSchemeNormalize URIFlags = 256
)

// URIHideFlags flags describing what parts of the URI to hide in
// g_uri_to_string_partial(). Note that G_URI_HIDE_PASSWORD and
// G_URI_HIDE_AUTH_PARAMS will only work if the #GUri was parsed with the
// corresponding flags.
type URIHideFlags int

const (
	// URIHideFlagsNone: no flags set.
	URIHideFlagsNone URIHideFlags = 0
	// URIHideFlagsUserinfo: hide the userinfo.
	URIHideFlagsUserinfo URIHideFlags = 1
	// URIHideFlagsPassword: hide the password.
	URIHideFlagsPassword URIHideFlags = 2
	// URIHideFlagsAuthParams: hide the auth_params.
	URIHideFlagsAuthParams URIHideFlags = 4
	// URIHideFlagsQuery: hide the query.
	URIHideFlagsQuery URIHideFlags = 8
	// URIHideFlagsFragment: hide the fragment.
	URIHideFlagsFragment URIHideFlags = 16
)

// URIParamsFlags flags modifying the way parameters are handled by
// g_uri_parse_params() and ParamsIter.
type URIParamsFlags int

const (
	// URIParamsFlagsNone: no flags set.
	URIParamsFlagsNone URIParamsFlags = 0
	// URIParamsFlagsCaseInsensitive: parameter names are case insensitive.
	URIParamsFlagsCaseInsensitive URIParamsFlags = 1
	// URIParamsFlagsWwwForm: replace `+` with space character. Only useful for
	// URLs on the web, using the `https` or `http` schemas.
	URIParamsFlagsWwwForm URIParamsFlags = 2
	// URIParamsFlagsParseRelaxed: see G_URI_FLAGS_PARSE_RELAXED.
	URIParamsFlagsParseRelaxed URIParamsFlags = 4
)

// URI: the #GUri type and related functions can be used to parse URIs into
// their components, and build valid URIs from individual components.
//
// Note that #GUri scope is to help manipulate URIs in various applications,
// following RFC 3986 (https://tools.ietf.org/html/rfc3986). In particular, it
// doesn't intend to cover web browser needs, and doesn't implement the WHATWG
// URL (https://url.spec.whatwg.org/) standard. No APIs are provided to help
// prevent homograph attacks
// (https://en.wikipedia.org/wiki/IDN_homograph_attack), so #GUri is not
// suitable for formatting URIs for display to the user for making
// security-sensitive decisions.
//
//
// Relative and absolute URIs
//
// As defined in RFC 3986 (https://tools.ietf.org/html/rfc3986#section-4), the
// hierarchical nature of URIs means that they can either be ‘relative
// references’ (sometimes referred to as ‘relative URIs’) or ‘URIs’ (for
// clarity, ‘URIs’ are referred to in this documentation as ‘absolute URIs’ —
// although in constrast to RFC 3986
// (https://tools.ietf.org/html/rfc3986#section-4.3), fragment identifiers are
// always allowed).
//
// Relative references have one or more components of the URI missing. In
// particular, they have no scheme. Any other component, such as hostname,
// query, etc. may be missing, apart from a path, which has to be specified (but
// may be empty). The path may be relative, starting with `./` rather than `/`.
//
// For example, a valid relative reference is `./path?query`, `/?query#fragment`
// or `//example.com`.
//
// Absolute URIs have a scheme specified. Any other components of the URI which
// are missing are specified as explicitly unset in the URI, rather than being
// resolved relative to a base URI using g_uri_parse_relative().
//
// For example, a valid absolute URI is `file:///home/bob` or
// `https://search.com?query=string`.
//
// A #GUri instance is always an absolute URI. A string may be an absolute URI
// or a relative reference; see the documentation for individual functions as to
// what forms they accept.
//
//
// Parsing URIs
//
// The most minimalist APIs for parsing URIs are g_uri_split() and
// g_uri_split_with_user(). These split a URI into its component parts, and
// return the parts; the difference between the two is that g_uri_split() treats
// the ‘userinfo’ component of the URI as a single element, while
// g_uri_split_with_user() can (depending on the Flags you pass) treat it as
// containing a username, password, and authentication parameters.
// Alternatively, g_uri_split_network() can be used when you are only interested
// in the components that are needed to initiate a network connection to the
// service (scheme, host, and port).
//
// g_uri_parse() is similar to g_uri_split(), but instead of returning
// individual strings, it returns a #GUri structure (and it requires that the
// URI be an absolute URI).
//
// g_uri_resolve_relative() and g_uri_parse_relative() allow you to resolve a
// relative URI relative to a base URI. g_uri_resolve_relative() takes two
// strings and returns a string, and g_uri_parse_relative() takes a #GUri and a
// string and returns a #GUri.
//
// All of the parsing functions take a Flags argument describing exactly how to
// parse the URI; see the documentation for that type for more details on the
// specific flags that you can pass. If you need to choose different flags based
// on the type of URI, you can use g_uri_peek_scheme() on the URI string to
// check the scheme first, and use that to decide what flags to parse it with.
//
// For example, you might want to use G_URI_PARAMS_WWW_FORM when parsing the
// params for a web URI, so compare the result of g_uri_peek_scheme() against
// `http` and `https`.
//
//
// Building URIs
//
// g_uri_join() and g_uri_join_with_user() can be used to construct valid URI
// strings from a set of component strings. They are the inverse of
// g_uri_split() and g_uri_split_with_user().
//
// Similarly, g_uri_build() and g_uri_build_with_user() can be used to construct
// a #GUri from a set of component strings.
//
// As with the parsing functions, the building functions take a Flags argument.
// In particular, it is important to keep in mind whether the URI components you
// are using are already `%`-encoded. If so, you must pass the
// G_URI_FLAGS_ENCODED flag.
//
// `file://` URIs
//
// Note that Windows and Unix both define special rules for parsing `file://`
// URIs (involving non-UTF-8 character sets on Unix, and the interpretation of
// path separators on Windows). #GUri does not implement these rules. Use
// g_filename_from_uri() and g_filename_to_uri() if you want to properly convert
// between `file://` URIs and local filenames.
//
//
// URI Equality
//
// Note that there is no `g_uri_equal ()` function, because comparing URIs
// usefully requires scheme-specific knowledge that #GUri does not have. #GUri
// can help with normalization if you use the various encoded Flags as well as
// G_URI_FLAGS_SCHEME_NORMALIZE however it is not comprehensive. For example,
// `data:,foo` and `data:;base64,Zm9v` resolve to the same thing according to
// the `data:` URI specification which GLib does not handle.
type URI struct {
	native C.GUri
}

// WrapURI wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapURI(ptr unsafe.Pointer) *URI {
	if ptr == nil {
		return nil
	}

	return (*URI)(ptr)
}

func marshalURI(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapURI(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (u *URI) Native() unsafe.Pointer {
	return unsafe.Pointer(&u.native)
}

// AuthParams gets @uri's authentication parameters, which may contain
// `%`-encoding, depending on the flags with which @uri was created. (If @uri
// was not created with G_URI_FLAGS_HAS_AUTH_PARAMS then this will be nil.)
//
// Depending on the URI scheme, g_uri_parse_params() may be useful for further
// parsing this information.
func (u *URI) AuthParams() string {
	var _arg0 *C.GUri

	_arg0 = (*C.GUri)(unsafe.Pointer(u.Native()))

	var _cret *C.gchar

	cret = C.g_uri_get_auth_params(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Flags gets @uri's flags set upon construction.
func (u *URI) Flags() URIFlags {
	var _arg0 *C.GUri

	_arg0 = (*C.GUri)(unsafe.Pointer(u.Native()))

	var _cret C.GUriFlags

	cret = C.g_uri_get_flags(_arg0)

	var _uriFlags URIFlags

	_uriFlags = URIFlags(_cret)

	return _uriFlags
}

// Fragment gets @uri's fragment, which may contain `%`-encoding, depending on
// the flags with which @uri was created.
func (u *URI) Fragment() string {
	var _arg0 *C.GUri

	_arg0 = (*C.GUri)(unsafe.Pointer(u.Native()))

	var _cret *C.gchar

	cret = C.g_uri_get_fragment(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Host gets @uri's host. This will never have `%`-encoded characters, unless it
// is non-UTF-8 (which can only be the case if @uri was created with
// G_URI_FLAGS_NON_DNS).
//
// If @uri contained an IPv6 address literal, this value will be just that
// address, without the brackets around it that are necessary in the string form
// of the URI. Note that in this case there may also be a scope ID attached to
// the address. Eg, `fe80::1234%“em1` (or `fe80::1234%“25em1` if the string is
// still encoded).
func (u *URI) Host() string {
	var _arg0 *C.GUri

	_arg0 = (*C.GUri)(unsafe.Pointer(u.Native()))

	var _cret *C.gchar

	cret = C.g_uri_get_host(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Password gets @uri's password, which may contain `%`-encoding, depending on
// the flags with which @uri was created. (If @uri was not created with
// G_URI_FLAGS_HAS_PASSWORD then this will be nil.)
func (u *URI) Password() string {
	var _arg0 *C.GUri

	_arg0 = (*C.GUri)(unsafe.Pointer(u.Native()))

	var _cret *C.gchar

	cret = C.g_uri_get_password(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Path gets @uri's path, which may contain `%`-encoding, depending on the flags
// with which @uri was created.
func (u *URI) Path() string {
	var _arg0 *C.GUri

	_arg0 = (*C.GUri)(unsafe.Pointer(u.Native()))

	var _cret *C.gchar

	cret = C.g_uri_get_path(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Port gets @uri's port.
func (u *URI) Port() int {
	var _arg0 *C.GUri

	_arg0 = (*C.GUri)(unsafe.Pointer(u.Native()))

	var _cret C.gint

	cret = C.g_uri_get_port(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Query gets @uri's query, which may contain `%`-encoding, depending on the
// flags with which @uri was created.
//
// For queries consisting of a series of `name=value` parameters, ParamsIter or
// g_uri_parse_params() may be useful.
func (u *URI) Query() string {
	var _arg0 *C.GUri

	_arg0 = (*C.GUri)(unsafe.Pointer(u.Native()))

	var _cret *C.gchar

	cret = C.g_uri_get_query(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Scheme gets @uri's scheme. Note that this will always be all-lowercase,
// regardless of the string or strings that @uri was created from.
func (u *URI) Scheme() string {
	var _arg0 *C.GUri

	_arg0 = (*C.GUri)(unsafe.Pointer(u.Native()))

	var _cret *C.gchar

	cret = C.g_uri_get_scheme(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// User gets the ‘username’ component of @uri's userinfo, which may contain
// `%`-encoding, depending on the flags with which @uri was created. If @uri was
// not created with G_URI_FLAGS_HAS_PASSWORD or G_URI_FLAGS_HAS_AUTH_PARAMS,
// this is the same as g_uri_get_userinfo().
func (u *URI) User() string {
	var _arg0 *C.GUri

	_arg0 = (*C.GUri)(unsafe.Pointer(u.Native()))

	var _cret *C.gchar

	cret = C.g_uri_get_user(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Userinfo gets @uri's userinfo, which may contain `%`-encoding, depending on
// the flags with which @uri was created.
func (u *URI) Userinfo() string {
	var _arg0 *C.GUri

	_arg0 = (*C.GUri)(unsafe.Pointer(u.Native()))

	var _cret *C.gchar

	cret = C.g_uri_get_userinfo(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// ParseRelative parses @uri_ref according to @flags and, if it is a [relative
// URI][relative-absolute-uris], resolves it relative to @base_uri. If the
// result is not a valid absolute URI, it will be discarded, and an error
// returned.
func (b *URI) ParseRelative(uriRef string, flags URIFlags) (*URI, error) {
	var _arg0 *C.GUri
	var _arg1 *C.gchar
	var _arg2 C.GUriFlags

	_arg0 = (*C.GUri)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.gchar)(C.CString(uriRef))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.GUriFlags)(flags)

	var _cret *C.GUri
	var _cerr *C.GError

	cret = C.g_uri_parse_relative(_arg0, _arg1, _arg2, _cerr)

	var _uri *URI
	var _goerr error

	_uri = WrapURI(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_uri, func(v *URI) {
		C.free(unsafe.Pointer(v.Native()))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _uri, _goerr
}

// Ref increments the reference count of @uri by one.
func (u *URI) Ref() *URI {
	var _arg0 *C.GUri

	_arg0 = (*C.GUri)(unsafe.Pointer(u.Native()))

	var _cret *C.GUri

	cret = C.g_uri_ref(_arg0)

	var _ret *URI

	_ret = WrapURI(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_ret, func(v *URI) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _ret
}

// String returns a string representing @uri.
//
// This is not guaranteed to return a string which is identical to the string
// that @uri was parsed from. However, if the source URI was syntactically
// correct (according to RFC 3986), and it was parsed with G_URI_FLAGS_ENCODED,
// then g_uri_to_string() is guaranteed to return a string which is at least
// semantically equivalent to the source URI (according to RFC 3986).
//
// If @uri might contain sensitive details, such as authentication parameters,
// or private data in its query string, and the returned string is going to be
// logged, then consider using g_uri_to_string_partial() to redact parts.
func (u *URI) String() string {
	var _arg0 *C.GUri

	_arg0 = (*C.GUri)(unsafe.Pointer(u.Native()))

	var _cret *C.char

	cret = C.g_uri_to_string(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ToStringPartial returns a string representing @uri, subject to the options in
// @flags. See g_uri_to_string() and HideFlags for more details.
func (u *URI) ToStringPartial(flags URIHideFlags) string {
	var _arg0 *C.GUri
	var _arg1 C.GUriHideFlags

	_arg0 = (*C.GUri)(unsafe.Pointer(u.Native()))
	_arg1 = (C.GUriHideFlags)(flags)

	var _cret *C.char

	cret = C.g_uri_to_string_partial(_arg0, _arg1)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Unref: atomically decrements the reference count of @uri by one.
//
// When the reference count reaches zero, the resources allocated by @uri are
// freed
func (u *URI) Unref() {
	var _arg0 *C.GUri

	_arg0 = (*C.GUri)(unsafe.Pointer(u.Native()))

	C.g_uri_unref(_arg0)
}

// URIParamsIter: many URI schemes include one or more attribute/value pairs as
// part of the URI value. For example
// `scheme://server/path?query=string&is=there` has two attributes –
// `query=string` and `is=there` – in its query part.
//
// A ParamsIter structure represents an iterator that can be used to iterate
// over the attribute/value pairs of a URI query string. ParamsIter structures
// are typically allocated on the stack and then initialized with
// g_uri_params_iter_init(). See the documentation for g_uri_params_iter_init()
// for a usage example.
type URIParamsIter struct {
	native C.GUriParamsIter
}

// WrapURIParamsIter wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapURIParamsIter(ptr unsafe.Pointer) *URIParamsIter {
	if ptr == nil {
		return nil
	}

	return (*URIParamsIter)(ptr)
}

func marshalURIParamsIter(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapURIParamsIter(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (u *URIParamsIter) Native() unsafe.Pointer {
	return unsafe.Pointer(&u.native)
}

// Init initializes an attribute/value pair iterator.
//
// The iterator keeps pointers to the @params and @separators arguments, those
// variables must thus outlive the iterator and not be modified during the
// iteration.
//
// If G_URI_PARAMS_WWW_FORM is passed in @flags, `+` characters in the param
// string will be replaced with spaces in the output. For example, `foo=bar+baz`
// will give attribute `foo` with value `bar baz`. This is commonly used on the
// web (the `https` and `http` schemes only), but is deprecated in favour of the
// equivalent of encoding spaces as `20`.
//
// Unlike with g_uri_parse_params(), G_URI_PARAMS_CASE_INSENSITIVE has no effect
// if passed to @flags for g_uri_params_iter_init(). The caller is responsible
// for doing their own case-insensitive comparisons.
//
//    GUriParamsIter iter;
//    GError *error = NULL;
//    gchar *unowned_attr, *unowned_value;
//
//    g_uri_params_iter_init (&iter, "foo=bar&baz=bar&Foo=frob&baz=bar2", -1, "&", G_URI_PARAMS_NONE);
//    while (g_uri_params_iter_next (&iter, &unowned_attr, &unowned_value, &error))
//      {
//        g_autofree gchar *attr = g_steal_pointer (&unowned_attr);
//        g_autofree gchar *value = g_steal_pointer (&unowned_value);
//        // do something with attr and value; this code will be called 4 times
//        // for the params string in this example: once with attr=foo and value=bar,
//        // then with baz/bar, then Foo/frob, then baz/bar2.
//      }
//    if (error)
//      // handle parsing error
func (i *URIParamsIter) Init(params string, length int, separators string, flags URIParamsFlags) {
	var _arg0 *C.GUriParamsIter
	var _arg1 *C.gchar
	var _arg2 C.gssize
	var _arg3 *C.gchar
	var _arg4 C.GUriParamsFlags

	_arg0 = (*C.GUriParamsIter)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(params))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(length)
	_arg3 = (*C.gchar)(C.CString(separators))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (C.GUriParamsFlags)(flags)

	C.g_uri_params_iter_init(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// Next advances @iter and retrieves the next attribute/value. false is returned
// if an error has occurred (in which case @error is set), or if the end of the
// iteration is reached (in which case @attribute and @value are set to nil and
// the iterator becomes invalid). If true is returned, g_uri_params_iter_next()
// may be called again to receive another attribute/value pair.
//
// Note that the same @attribute may be returned multiple times, since URIs
// allow repeated attributes.
func (i *URIParamsIter) Next() (attribute string, value string, goerr error) {
	var _arg0 *C.GUriParamsIter

	_arg0 = (*C.GUriParamsIter)(unsafe.Pointer(i.Native()))

	var _arg1 *C.gchar
	var _arg2 *C.gchar
	var _cerr *C.GError

	C.g_uri_params_iter_next(_arg0, &_arg1, &_arg2, _cerr)

	var _attribute string
	var _value string
	var _goerr error

	_attribute = C.GoString(_arg1)
	defer C.free(unsafe.Pointer(_arg1))
	_value = C.GoString(_arg2)
	defer C.free(unsafe.Pointer(_arg2))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _attribute, _value, _goerr
}