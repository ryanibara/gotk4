// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkpagesetup.go.
var GTypePageSetup = coreglib.Type(C.gtk_page_setup_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypePageSetup, F: marshalPageSetup},
	})
}

// PageSetup: GtkPageSetup object stores the page size, orientation and margins.
//
// The idea is that you can get one of these from the page setup dialog and then
// pass it to the GtkPrintOperation when printing. The benefit of splitting this
// out of the GtkPrintSettings is that these affect the actual layout of the
// page, and thus need to be set long before user prints.
//
//
// Margins
//
// The margins specified in this object are the “print margins”, i.e. the parts
// of the page that the printer cannot print on. These are different from the
// layout margins that a word processor uses; they are typically used to
// determine the minimal size for the layout margins.
//
// To obtain a GtkPageSetup use gtk.PageSetup.New to get the defaults, or use
// gtk.PrintRunPageSetupDialog() to show the page setup dialog and receive the
// resulting page setup.
//
// A page setup dialog
//
//    static GtkPrintSettings *settings = NULL;
//    static GtkPageSetup *page_setup = NULL;
//
//    static void
//    do_page_setup (void)
//    {
//      GtkPageSetup *new_page_setup;
//
//      if (settings == NULL)
//        settings = gtk_print_settings_new ();
//
//      new_page_setup = gtk_print_run_page_setup_dialog (GTK_WINDOW (main_window),
//                                                        page_setup, settings);
//
//      if (page_setup)
//        g_object_unref (page_setup);
//
//      page_setup = new_page_setup;
//    }.
type PageSetup struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*PageSetup)(nil)
)

func wrapPageSetup(obj *coreglib.Object) *PageSetup {
	return &PageSetup{
		Object: obj,
	}
}

func marshalPageSetup(p uintptr) (interface{}, error) {
	return wrapPageSetup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPageSetup creates a new GtkPageSetup.
//
// The function returns the following values:
//
//    - pageSetup: new GtkPageSetup.
//
func NewPageSetup() *PageSetup {
	_gret := girepository.MustFind("Gtk", "PageSetup").InvokeMethod("new_PageSetup", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _pageSetup *PageSetup // out

	_pageSetup = wrapPageSetup(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _pageSetup
}

// NewPageSetupFromFile reads the page setup from the file file_name.
//
// Returns a new GtkPageSetup object with the restored page setup, or NULL if an
// error occurred. See gtk.PageSetup.ToFile().
//
// The function takes the following parameters:
//
//    - fileName: filename to read the page setup from.
//
// The function returns the following values:
//
//    - pageSetup: restored GtkPageSetup.
//
func NewPageSetupFromFile(fileName string) (*PageSetup, error) {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(_args[0]))

	_gret := girepository.MustFind("Gtk", "PageSetup").InvokeMethod("new_PageSetup_from_file", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fileName)

	var _pageSetup *PageSetup // out
	var _goerr error          // out

	_pageSetup = wrapPageSetup(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _pageSetup, _goerr
}

// NewPageSetupFromGVariant: desrialize a page setup from an a{sv} variant.
//
// The variant must be in the format produced by gtk.PageSetup.ToGVariant().
//
// The function takes the following parameters:
//
//    - variant: a{sv} GVariant.
//
// The function returns the following values:
//
//    - pageSetup: new GtkPageSetup object.
//
func NewPageSetupFromGVariant(variant *glib.Variant) *PageSetup {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(variant)))

	_gret := girepository.MustFind("Gtk", "PageSetup").InvokeMethod("new_PageSetup_from_gvariant", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(variant)

	var _pageSetup *PageSetup // out

	_pageSetup = wrapPageSetup(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _pageSetup
}

// NewPageSetupFromKeyFile reads the page setup from the group group_name in the
// key file key_file.
//
// Returns a new GtkPageSetup object with the restored page setup, or NULL if an
// error occurred.
//
// The function takes the following parameters:
//
//    - keyFile: GKeyFile to retrieve the page_setup from.
//    - groupName (optional): name of the group in the key_file to read, or NULL
//      to use the default name “Page Setup”.
//
// The function returns the following values:
//
//    - pageSetup: restored GtkPageSetup.
//
func NewPageSetupFromKeyFile(keyFile *glib.KeyFile, groupName string) (*PageSetup, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(keyFile)))
	if groupName != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(groupName)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	_gret := girepository.MustFind("Gtk", "PageSetup").InvokeMethod("new_PageSetup_from_key_file", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(keyFile)
	runtime.KeepAlive(groupName)

	var _pageSetup *PageSetup // out
	var _goerr error          // out

	_pageSetup = wrapPageSetup(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _pageSetup, _goerr
}

// Copy copies a GtkPageSetup.
//
// The function returns the following values:
//
//    - pageSetup: copy of other.
//
func (other *PageSetup) Copy() *PageSetup {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(other).Native()))

	_gret := girepository.MustFind("Gtk", "PageSetup").InvokeMethod("copy", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(other)

	var _pageSetup *PageSetup // out

	_pageSetup = wrapPageSetup(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _pageSetup
}

// PaperSize gets the paper size of the GtkPageSetup.
//
// The function returns the following values:
//
//    - paperSize: paper size.
//
func (setup *PageSetup) PaperSize() *PaperSize {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(setup).Native()))

	_gret := girepository.MustFind("Gtk", "PageSetup").InvokeMethod("get_paper_size", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(setup)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _paperSize
}

// LoadFile reads the page setup from the file file_name.
//
// See gtk.PageSetup.ToFile().
//
// The function takes the following parameters:
//
//    - fileName: filename to read the page setup from.
//
func (setup *PageSetup) LoadFile(fileName string) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "PageSetup").InvokeMethod("load_file", _args[:], nil)

	runtime.KeepAlive(setup)
	runtime.KeepAlive(fileName)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// LoadKeyFile reads the page setup from the group group_name in the key file
// key_file.
//
// The function takes the following parameters:
//
//    - keyFile: GKeyFile to retrieve the page_setup from.
//    - groupName (optional): name of the group in the key_file to read, or NULL
//      to use the default name “Page Setup”.
//
func (setup *PageSetup) LoadKeyFile(keyFile *glib.KeyFile, groupName string) error {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(keyFile)))
	if groupName != "" {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(groupName)))
		defer C.free(unsafe.Pointer(_args[2]))
	}

	girepository.MustFind("Gtk", "PageSetup").InvokeMethod("load_key_file", _args[:], nil)

	runtime.KeepAlive(setup)
	runtime.KeepAlive(keyFile)
	runtime.KeepAlive(groupName)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetPaperSize sets the paper size of the GtkPageSetup without changing the
// margins.
//
// See gtk.PageSetup.SetPaperSizeAndDefaultMargins().
//
// The function takes the following parameters:
//
//    - size: PaperSize.
//
func (setup *PageSetup) SetPaperSize(size *PaperSize) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(size)))

	girepository.MustFind("Gtk", "PageSetup").InvokeMethod("set_paper_size", _args[:], nil)

	runtime.KeepAlive(setup)
	runtime.KeepAlive(size)
}

// SetPaperSizeAndDefaultMargins sets the paper size of the GtkPageSetup and
// modifies the margins according to the new paper size.
//
// The function takes the following parameters:
//
//    - size: PaperSize.
//
func (setup *PageSetup) SetPaperSizeAndDefaultMargins(size *PaperSize) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(size)))

	girepository.MustFind("Gtk", "PageSetup").InvokeMethod("set_paper_size_and_default_margins", _args[:], nil)

	runtime.KeepAlive(setup)
	runtime.KeepAlive(size)
}

// ToFile: this function saves the information from setup to file_name.
//
// The function takes the following parameters:
//
//    - fileName: file to save to.
//
func (setup *PageSetup) ToFile(fileName string) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "PageSetup").InvokeMethod("to_file", _args[:], nil)

	runtime.KeepAlive(setup)
	runtime.KeepAlive(fileName)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// ToGVariant: serialize page setup to an a{sv} variant.
//
// The function returns the following values:
//
//    - variant: new, floating, #GVariant.
//
func (setup *PageSetup) ToGVariant() *glib.Variant {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(setup).Native()))

	_gret := girepository.MustFind("Gtk", "PageSetup").InvokeMethod("to_gvariant", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(setup)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	return _variant
}

// ToKeyFile: this function adds the page setup from setup to key_file.
//
// The function takes the following parameters:
//
//    - keyFile: GKeyFile to save the page setup to.
//    - groupName (optional): group to add the settings to in key_file, or NULL
//      to use the default name “Page Setup”.
//
func (setup *PageSetup) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(keyFile)))
	if groupName != "" {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(groupName)))
		defer C.free(unsafe.Pointer(_args[2]))
	}

	girepository.MustFind("Gtk", "PageSetup").InvokeMethod("to_key_file", _args[:], nil)

	runtime.KeepAlive(setup)
	runtime.KeepAlive(keyFile)
	runtime.KeepAlive(groupName)
}
