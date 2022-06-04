// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern GList* _gotk4_gtk3_RecentChooserIface_get_items(void*);
// extern GSList* _gotk4_gtk3_RecentChooserIface_list_filters(void*);
// extern gboolean _gotk4_gtk3_RecentChooserIface_select_uri(void*, void*, GError**);
// extern gboolean _gotk4_gtk3_RecentChooserIface_set_current_uri(void*, void*, GError**);
// extern gchar* _gotk4_gtk3_RecentChooserIface_get_current_uri(void*);
// extern gint _gotk4_gtk3_RecentSortFunc(void*, void*, gpointer);
// extern void _gotk4_gtk3_RecentChooserIface_add_filter(void*, void*);
// extern void _gotk4_gtk3_RecentChooserIface_item_activated(void*);
// extern void _gotk4_gtk3_RecentChooserIface_remove_filter(void*, void*);
// extern void _gotk4_gtk3_RecentChooserIface_select_all(void*);
// extern void _gotk4_gtk3_RecentChooserIface_selection_changed(void*);
// extern void _gotk4_gtk3_RecentChooserIface_unselect_all(void*);
// extern void _gotk4_gtk3_RecentChooserIface_unselect_uri(void*, void*);
// extern void _gotk4_gtk3_RecentChooser_ConnectItemActivated(gpointer, guintptr);
// extern void _gotk4_gtk3_RecentChooser_ConnectSelectionChanged(gpointer, guintptr);
// extern void callbackDelete(gpointer);
import "C"

// glib.Type values for gtkrecentchooser.go.
var (
	GTypeRecentChooserError = coreglib.Type(C.gtk_recent_chooser_error_get_type())
	GTypeRecentSortType     = coreglib.Type(C.gtk_recent_sort_type_get_type())
	GTypeRecentChooser      = coreglib.Type(C.gtk_recent_chooser_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeRecentChooserError, F: marshalRecentChooserError},
		{T: GTypeRecentSortType, F: marshalRecentSortType},
		{T: GTypeRecentChooser, F: marshalRecentChooser},
	})
}

// RecentChooserError: these identify the various errors that can occur while
// calling RecentChooser functions.
type RecentChooserError C.gint

const (
	// RecentChooserErrorNotFound indicates that a file does not exist.
	RecentChooserErrorNotFound RecentChooserError = iota
	// RecentChooserErrorInvalidURI indicates a malformed URI.
	RecentChooserErrorInvalidURI
)

func marshalRecentChooserError(p uintptr) (interface{}, error) {
	return RecentChooserError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for RecentChooserError.
func (r RecentChooserError) String() string {
	switch r {
	case RecentChooserErrorNotFound:
		return "NotFound"
	case RecentChooserErrorInvalidURI:
		return "InvalidURI"
	default:
		return fmt.Sprintf("RecentChooserError(%d)", r)
	}
}

// RecentSortType: used to specify the sorting method to be applyed to the
// recently used resource list.
type RecentSortType C.gint

const (
	// RecentSortNone: do not sort the returned list of recently used resources.
	RecentSortNone RecentSortType = iota
	// RecentSortMru: sort the returned list with the most recently used items
	// first.
	RecentSortMru
	// RecentSortLru: sort the returned list with the least recently used items
	// first.
	RecentSortLru
	// RecentSortCustom: sort the returned list using a custom sorting function
	// passed using gtk_recent_chooser_set_sort_func().
	RecentSortCustom
)

func marshalRecentSortType(p uintptr) (interface{}, error) {
	return RecentSortType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for RecentSortType.
func (r RecentSortType) String() string {
	switch r {
	case RecentSortNone:
		return "None"
	case RecentSortMru:
		return "Mru"
	case RecentSortLru:
		return "Lru"
	case RecentSortCustom:
		return "Custom"
	default:
		return fmt.Sprintf("RecentSortType(%d)", r)
	}
}

type RecentSortFunc func(a, b *RecentInfo) (gint int32)

//export _gotk4_gtk3_RecentSortFunc
func _gotk4_gtk3_RecentSortFunc(arg1 *C.void, arg2 *C.void, arg3 C.gpointer) (cret C.gint) {
	var fn RecentSortFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(RecentSortFunc)
	}

	var _a *RecentInfo // out
	var _b *RecentInfo // out

	_a = (*RecentInfo)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	C.gtk_recent_info_ref(arg1)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_a)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_recent_info_unref((*C.GtkRecentInfo)(intern.C))
		},
	)
	_b = (*RecentInfo)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	C.gtk_recent_info_ref(arg2)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_b)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_recent_info_unref((*C.GtkRecentInfo)(intern.C))
		},
	)

	gint := fn(_a, _b)

	cret = C.gint(gint)

	return cret
}

// RecentChooser is an interface that can be implemented by widgets displaying
// the list of recently used files. In GTK+, the main objects that implement
// this interface are RecentChooserWidget, RecentChooserDialog and
// RecentChooserMenu.
//
// Recently used files are supported since GTK+ 2.10.
//
// RecentChooser wraps an interface. This means the user can get the
// underlying type by calling Cast().
type RecentChooser struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*RecentChooser)(nil)
)

// RecentChooserer describes RecentChooser's interface methods.
type RecentChooserer interface {
	coreglib.Objector

	// AddFilter adds filter to the list of RecentFilter objects held by
	// chooser.
	AddFilter(filter *RecentFilter)
	// CurrentItem gets the RecentInfo currently selected by chooser.
	CurrentItem() *RecentInfo
	// CurrentURI gets the URI currently selected by chooser.
	CurrentURI() string
	// Filter gets the RecentFilter object currently used by chooser to affect
	// the display of the recently used resources.
	Filter() *RecentFilter
	// Items gets the list of recently used resources in form of RecentInfo
	// objects.
	Items() []*RecentInfo
	// Limit gets the number of items returned by gtk_recent_chooser_get_items()
	// and gtk_recent_chooser_get_uris().
	Limit() int32
	// LocalOnly gets whether only local resources should be shown in the
	// recently used resources selector.
	LocalOnly() bool
	// SelectMultiple gets whether chooser can select multiple items.
	SelectMultiple() bool
	// ShowIcons retrieves whether chooser should show an icon near the
	// resource.
	ShowIcons() bool
	// ShowNotFound retrieves whether chooser should show the recently used
	// resources that were not found.
	ShowNotFound() bool
	// ShowPrivate returns whether chooser should display recently used
	// resources registered as private.
	ShowPrivate() bool
	// ShowTips gets whether chooser should display tooltips containing the full
	// path of a recently user resource.
	ShowTips() bool
	// URIs gets the URI of the recently used resources.
	URIs() []string
	// ListFilters gets the RecentFilter objects held by chooser.
	ListFilters() []*RecentFilter
	// RemoveFilter removes filter from the list of RecentFilter objects held by
	// chooser.
	RemoveFilter(filter *RecentFilter)
	// SelectAll selects all the items inside chooser, if the chooser supports
	// multiple selection.
	SelectAll()
	// SelectURI selects uri inside chooser.
	SelectURI(uri string) error
	// SetCurrentURI sets uri as the current URI for chooser.
	SetCurrentURI(uri string) error
	// SetFilter sets filter as the current RecentFilter object used by chooser
	// to affect the displayed recently used resources.
	SetFilter(filter *RecentFilter)
	// SetLimit sets the number of items that should be returned by
	// gtk_recent_chooser_get_items() and gtk_recent_chooser_get_uris().
	SetLimit(limit int32)
	// SetLocalOnly sets whether only local resources, that is resources using
	// the file:// URI scheme, should be shown in the recently used resources
	// selector.
	SetLocalOnly(localOnly bool)
	// SetSelectMultiple sets whether chooser can select multiple items.
	SetSelectMultiple(selectMultiple bool)
	// SetShowIcons sets whether chooser should show an icon near the resource
	// when displaying it.
	SetShowIcons(showIcons bool)
	// SetShowNotFound sets whether chooser should display the recently used
	// resources that it didn’t find.
	SetShowNotFound(showNotFound bool)
	// SetShowPrivate: whether to show recently used resources marked registered
	// as private.
	SetShowPrivate(showPrivate bool)
	// SetShowTips sets whether to show a tooltips containing the full path of
	// each recently used resource in a RecentChooser widget.
	SetShowTips(showTips bool)
	// SetSortFunc sets the comparison function used when sorting to be
	// sort_func.
	SetSortFunc(sortFunc RecentSortFunc)
	// UnselectAll unselects all the items inside chooser.
	UnselectAll()
	// UnselectURI unselects uri inside chooser.
	UnselectURI(uri string)

	// Item-activated: this signal is emitted when the user "activates" a recent
	// item in the recent chooser.
	ConnectItemActivated(func()) coreglib.SignalHandle
	// Selection-changed: this signal is emitted when there is a change in the
	// set of selected recently used resources.
	ConnectSelectionChanged(func()) coreglib.SignalHandle
}

var _ RecentChooserer = (*RecentChooser)(nil)

func wrapRecentChooser(obj *coreglib.Object) *RecentChooser {
	return &RecentChooser{
		Object: obj,
	}
}

func marshalRecentChooser(p uintptr) (interface{}, error) {
	return wrapRecentChooser(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_RecentChooser_ConnectItemActivated
func _gotk4_gtk3_RecentChooser_ConnectItemActivated(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectItemActivated: this signal is emitted when the user "activates" a
// recent item in the recent chooser. This can happen by double-clicking on an
// item in the recently used resources list, or by pressing Enter.
func (chooser *RecentChooser) ConnectItemActivated(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(chooser, "item-activated", false, unsafe.Pointer(C._gotk4_gtk3_RecentChooser_ConnectItemActivated), f)
}

//export _gotk4_gtk3_RecentChooser_ConnectSelectionChanged
func _gotk4_gtk3_RecentChooser_ConnectSelectionChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectSelectionChanged: this signal is emitted when there is a change in the
// set of selected recently used resources. This can happen when a user modifies
// the selection with the mouse or the keyboard, or when explicitly calling
// functions to change the selection.
func (chooser *RecentChooser) ConnectSelectionChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(chooser, "selection-changed", false, unsafe.Pointer(C._gotk4_gtk3_RecentChooser_ConnectSelectionChanged), f)
}

// AddFilter adds filter to the list of RecentFilter objects held by chooser.
//
// If no previous filter objects were defined, this function will call
// gtk_recent_chooser_set_filter().
//
// The function takes the following parameters:
//
//    - filter: RecentFilter.
//
func (chooser *RecentChooser) AddFilter(filter *RecentFilter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))

	runtime.KeepAlive(chooser)
	runtime.KeepAlive(filter)
}

// CurrentItem gets the RecentInfo currently selected by chooser.
//
// The function returns the following values:
//
//    - recentInfo: RecentInfo. Use gtk_recent_info_unref() when when you have
//      finished using it.
//
func (chooser *RecentChooser) CurrentItem() *RecentInfo {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(chooser)

	var _recentInfo *RecentInfo // out

	_recentInfo = (*RecentInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_recentInfo)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_recent_info_unref((*C.GtkRecentInfo)(intern.C))
		},
	)

	return _recentInfo
}

// CurrentURI gets the URI currently selected by chooser.
//
// The function returns the following values:
//
//    - utf8: newly allocated string holding a URI.
//
func (chooser *RecentChooser) CurrentURI() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(chooser)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Filter gets the RecentFilter object currently used by chooser to affect the
// display of the recently used resources.
//
// The function returns the following values:
//
//    - recentFilter: RecentFilter object.
//
func (chooser *RecentChooser) Filter() *RecentFilter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(chooser)

	var _recentFilter *RecentFilter // out

	_recentFilter = wrapRecentFilter(coreglib.Take(unsafe.Pointer(_cret)))

	return _recentFilter
}

// Items gets the list of recently used resources in form of RecentInfo objects.
//
// The return value of this function is affected by the “sort-type” and “limit”
// properties of chooser.
//
// The function returns the following values:
//
//    - list: newly allocated list of RecentInfo objects. You should use
//      gtk_recent_info_unref() on every item of the list, and then free the list
//      itself using g_list_free().
//
func (chooser *RecentChooser) Items() []*RecentInfo {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(chooser)

	var _list []*RecentInfo // out

	_list = make([]*RecentInfo, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst *RecentInfo // out
		dst = (*RecentInfo)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gtk_recent_info_unref((*C.GtkRecentInfo)(intern.C))
			},
		)
		_list = append(_list, dst)
	})

	return _list
}

// Limit gets the number of items returned by gtk_recent_chooser_get_items() and
// gtk_recent_chooser_get_uris().
//
// The function returns the following values:
//
//    - gint: positive integer, or -1 meaning that all items are returned.
//
func (chooser *RecentChooser) Limit() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(chooser)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// LocalOnly gets whether only local resources should be shown in the recently
// used resources selector. See gtk_recent_chooser_set_local_only().
//
// The function returns the following values:
//
//    - ok: TRUE if only local resources should be shown.
//
func (chooser *RecentChooser) LocalOnly() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(chooser)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SelectMultiple gets whether chooser can select multiple items.
//
// The function returns the following values:
//
//    - ok: TRUE if chooser can select more than one item.
//
func (chooser *RecentChooser) SelectMultiple() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(chooser)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ShowIcons retrieves whether chooser should show an icon near the resource.
//
// The function returns the following values:
//
//    - ok: TRUE if the icons should be displayed, FALSE otherwise.
//
func (chooser *RecentChooser) ShowIcons() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(chooser)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ShowNotFound retrieves whether chooser should show the recently used
// resources that were not found.
//
// The function returns the following values:
//
//    - ok: TRUE if the resources not found should be displayed, and FALSE
//      otheriwse.
//
func (chooser *RecentChooser) ShowNotFound() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(chooser)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ShowPrivate returns whether chooser should display recently used resources
// registered as private.
//
// The function returns the following values:
//
//    - ok: TRUE if the recent chooser should show private items, FALSE
//      otherwise.
//
func (chooser *RecentChooser) ShowPrivate() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(chooser)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ShowTips gets whether chooser should display tooltips containing the full
// path of a recently user resource.
//
// The function returns the following values:
//
//    - ok: TRUE if the recent chooser should show tooltips, FALSE otherwise.
//
func (chooser *RecentChooser) ShowTips() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(chooser)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// URIs gets the URI of the recently used resources.
//
// The return value of this function is affected by the “sort-type” and “limit”
// properties of chooser.
//
// Since the returned array is NULL terminated, length may be NULL.
//
// The function returns the following values:
//
//    - utf8s: A newly allocated, NULL-terminated array of strings. Use
//      g_strfreev() to free it.
//
func (chooser *RecentChooser) URIs() []string {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = *(***C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(chooser)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		src := unsafe.Slice((**C.void)(_cret), _outs[0])
		_utf8s = make([]string, _outs[0])
		for i := 0; i < int(_outs[0]); i++ {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// ListFilters gets the RecentFilter objects held by chooser.
//
// The function returns the following values:
//
//    - sList: singly linked list of RecentFilter objects. You should just free
//      the returned list using g_slist_free().
//
func (chooser *RecentChooser) ListFilters() []*RecentFilter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(chooser)

	var _sList []*RecentFilter // out

	_sList = make([]*RecentFilter, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst *RecentFilter // out
		dst = wrapRecentFilter(coreglib.Take(unsafe.Pointer(src)))
		_sList = append(_sList, dst)
	})

	return _sList
}

// RemoveFilter removes filter from the list of RecentFilter objects held by
// chooser.
//
// The function takes the following parameters:
//
//    - filter: RecentFilter.
//
func (chooser *RecentChooser) RemoveFilter(filter *RecentFilter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))

	runtime.KeepAlive(chooser)
	runtime.KeepAlive(filter)
}

// SelectAll selects all the items inside chooser, if the chooser supports
// multiple selection.
func (chooser *RecentChooser) SelectAll() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	runtime.KeepAlive(chooser)
}

// SelectURI selects uri inside chooser.
//
// The function takes the following parameters:
//
//    - uri: URI.
//
func (chooser *RecentChooser) SelectURI(uri string) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_args[1]))

	runtime.KeepAlive(chooser)
	runtime.KeepAlive(uri)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetCurrentURI sets uri as the current URI for chooser.
//
// The function takes the following parameters:
//
//    - uri: URI.
//
func (chooser *RecentChooser) SetCurrentURI(uri string) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_args[1]))

	runtime.KeepAlive(chooser)
	runtime.KeepAlive(uri)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetFilter sets filter as the current RecentFilter object used by chooser to
// affect the displayed recently used resources.
//
// The function takes the following parameters:
//
//    - filter (optional): RecentFilter.
//
func (chooser *RecentChooser) SetFilter(filter *RecentFilter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	if filter != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	}

	runtime.KeepAlive(chooser)
	runtime.KeepAlive(filter)
}

// SetLimit sets the number of items that should be returned by
// gtk_recent_chooser_get_items() and gtk_recent_chooser_get_uris().
//
// The function takes the following parameters:
//
//    - limit: positive integer, or -1 for all items.
//
func (chooser *RecentChooser) SetLimit(limit int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(limit)

	runtime.KeepAlive(chooser)
	runtime.KeepAlive(limit)
}

// SetLocalOnly sets whether only local resources, that is resources using the
// file:// URI scheme, should be shown in the recently used resources selector.
// If local_only is TRUE (the default) then the shown resources are guaranteed
// to be accessible through the operating system native file system.
//
// The function takes the following parameters:
//
//    - localOnly: TRUE if only local files can be shown.
//
func (chooser *RecentChooser) SetLocalOnly(localOnly bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	if localOnly {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	runtime.KeepAlive(chooser)
	runtime.KeepAlive(localOnly)
}

// SetSelectMultiple sets whether chooser can select multiple items.
//
// The function takes the following parameters:
//
//    - selectMultiple: TRUE if chooser can select more than one item.
//
func (chooser *RecentChooser) SetSelectMultiple(selectMultiple bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	if selectMultiple {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	runtime.KeepAlive(chooser)
	runtime.KeepAlive(selectMultiple)
}

// SetShowIcons sets whether chooser should show an icon near the resource when
// displaying it.
//
// The function takes the following parameters:
//
//    - showIcons: whether to show an icon near the resource.
//
func (chooser *RecentChooser) SetShowIcons(showIcons bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	if showIcons {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	runtime.KeepAlive(chooser)
	runtime.KeepAlive(showIcons)
}

// SetShowNotFound sets whether chooser should display the recently used
// resources that it didn’t find. This only applies to local resources.
//
// The function takes the following parameters:
//
//    - showNotFound: whether to show the local items we didn’t find.
//
func (chooser *RecentChooser) SetShowNotFound(showNotFound bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	if showNotFound {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	runtime.KeepAlive(chooser)
	runtime.KeepAlive(showNotFound)
}

// SetShowPrivate: whether to show recently used resources marked registered as
// private.
//
// The function takes the following parameters:
//
//    - showPrivate: TRUE to show private items, FALSE otherwise.
//
func (chooser *RecentChooser) SetShowPrivate(showPrivate bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	if showPrivate {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	runtime.KeepAlive(chooser)
	runtime.KeepAlive(showPrivate)
}

// SetShowTips sets whether to show a tooltips containing the full path of each
// recently used resource in a RecentChooser widget.
//
// The function takes the following parameters:
//
//    - showTips: TRUE if tooltips should be shown.
//
func (chooser *RecentChooser) SetShowTips(showTips bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	if showTips {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	runtime.KeepAlive(chooser)
	runtime.KeepAlive(showTips)
}

// SetSortFunc sets the comparison function used when sorting to be sort_func.
// If the chooser has the sort type set to K_RECENT_SORT_CUSTOM then the chooser
// will sort using this function.
//
// To the comparison function will be passed two RecentInfo structs and
// sort_data; sort_func should return a positive integer if the first item comes
// before the second, zero if the two items are equal and a negative integer if
// the first item comes after the second.
//
// The function takes the following parameters:
//
//    - sortFunc: comparison function.
//
func (chooser *RecentChooser) SetSortFunc(sortFunc RecentSortFunc) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = (*[0]byte)(C._gotk4_gtk3_RecentSortFunc)
	_args[2] = C.gpointer(gbox.Assign(sortFunc))
	_args[3] = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	runtime.KeepAlive(chooser)
	runtime.KeepAlive(sortFunc)
}

// UnselectAll unselects all the items inside chooser.
func (chooser *RecentChooser) UnselectAll() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	runtime.KeepAlive(chooser)
}

// UnselectURI unselects uri inside chooser.
//
// The function takes the following parameters:
//
//    - uri: URI.
//
func (chooser *RecentChooser) UnselectURI(uri string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_args[1]))

	runtime.KeepAlive(chooser)
	runtime.KeepAlive(uri)
}
