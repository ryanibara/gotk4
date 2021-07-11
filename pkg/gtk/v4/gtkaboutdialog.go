// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
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
		{T: externglib.Type(C.gtk_license_get_type()), F: marshalLicense},
		{T: externglib.Type(C.gtk_about_dialog_get_type()), F: marshalAboutDialogger},
	})
}

// License: type of license for an application.
//
// This enumeration can be expanded at later date.
type License int

const (
	// Unknown: no license specified
	LicenseUnknown License = iota
	// Custom: license text is going to be specified by the developer
	LicenseCustom
	// GPL20: GNU General Public License, version 2.0 or later
	LicenseGPL20
	// GPL30: GNU General Public License, version 3.0 or later
	LicenseGPL30
	// LGPL21: GNU Lesser General Public License, version 2.1 or later
	LicenseLGPL21
	// LGPL30: GNU Lesser General Public License, version 3.0 or later
	LicenseLGPL30
	// BSD: BSD standard license
	LicenseBSD
	// MITX11: MIT/X11 standard license
	LicenseMITX11
	// Artistic: artistic License, version 2.0
	LicenseArtistic
	// GPL20Only: GNU General Public License, version 2.0 only
	LicenseGPL20Only
	// GPL30Only: GNU General Public License, version 3.0 only
	LicenseGPL30Only
	// LGPL21Only: GNU Lesser General Public License, version 2.1 only
	LicenseLGPL21Only
	// LGPL30Only: GNU Lesser General Public License, version 3.0 only
	LicenseLGPL30Only
	// AGPL30: GNU Affero General Public License, version 3.0 or later
	LicenseAGPL30
	// AGPL30Only: GNU Affero General Public License, version 3.0 only
	LicenseAGPL30Only
	// BSD3: 3-clause BSD licence
	LicenseBSD3
	// Apache20: apache License, version 2.0
	LicenseApache20
	// MPL20: mozilla Public License, version 2.0
	LicenseMPL20
)

func marshalLicense(p uintptr) (interface{}, error) {
	return License(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// AboutDialogger describes AboutDialog's methods.
type AboutDialogger interface {
	// AddCreditSection creates a new section in the "Credits" page.
	AddCreditSection(sectionName string, people []string)
	// Artists returns the string which are displayed in the "Artists" tab of
	// the secondary credits dialog.
	Artists() []string

	Authors() []string
	// Comments returns the comments string.
	Comments() string

	Copyright() string
	// Documenters returns the string which are displayed in the "Documenters"
	// tab of the secondary credits dialog.
	Documenters() []string
	// License returns the license information.
	License() string
	// LicenseType retrieves the license type.
	LicenseType() License
	// Logo returns the paintable displayed as logo in the about dialog.
	Logo() *gdk.Paintable
	// LogoIconName returns the icon name displayed as logo in the about dialog.
	LogoIconName() string
	// ProgramName returns the program name displayed in the about dialog.
	ProgramName() string
	// SystemInformation returns the system information that is shown in the
	// about dialog.
	SystemInformation() string
	// TranslatorCredits returns the translator credits string which is
	// displayed in the translators tab of the secondary credits dialog.
	TranslatorCredits() string
	// Version returns the version string.
	Version() string
	// Website returns the website URL.
	Website() string
	// WebsiteLabel returns the label used for the website link.
	WebsiteLabel() string
	// WrapLicense returns whether the license text in the about dialog is
	// automatically wrapped.
	WrapLicense() bool
	// SetArtists sets the strings which are displayed in the "Artists" tab of
	// the secondary credits dialog.
	SetArtists(artists []string)
	// SetAuthors sets the strings which are displayed in the "Authors" tab of
	// the secondary credits dialog.
	SetAuthors(authors []string)
	// SetComments sets the comments string to display in the about dialog.
	SetComments(comments string)
	// SetCopyright sets the copyright string to display in the about dialog.
	SetCopyright(copyright string)
	// SetDocumenters sets the strings which are displayed in the "Documenters"
	// tab of the credits dialog.
	SetDocumenters(documenters []string)
	// SetLicense sets the license information to be displayed in the secondary
	// license dialog.
	SetLicense(license string)
	// SetLogo sets the logo in the about dialog.
	SetLogo(logo gdk.Paintabler)
	// SetLogoIconName sets the icon name to be displayed as logo in the about
	// dialog.
	SetLogoIconName(iconName string)
	// SetProgramName sets the name to display in the about dialog.
	SetProgramName(name string)
	// SetSystemInformation sets the system information to be displayed in the
	// about dialog.
	SetSystemInformation(systemInformation string)
	// SetTranslatorCredits sets the translator credits string which is
	// displayed in the translators tab of the secondary credits dialog.
	SetTranslatorCredits(translatorCredits string)
	// SetVersion sets the version string to display in the about dialog.
	SetVersion(version string)
	// SetWebsite sets the URL to use for the website link.
	SetWebsite(website string)
	// SetWebsiteLabel sets the label to be used for the website link.
	SetWebsiteLabel(websiteLabel string)
	// SetWrapLicense sets whether the license text in the about dialog should
	// be automatically wrapped.
	SetWrapLicense(wrapLicense bool)
}

// AboutDialog: `GtkAboutDialog` offers a simple way to display information
// about a program.
//
// The shown information includes the programs' logo, name, copyright, website
// and license. It is also possible to give credits to the authors, documenters,
// translators and artists who have worked on the program.
//
// An about dialog is typically opened when the user selects the `About` option
// from the `Help` menu. All parts of the dialog are optional.
//
// !An example GtkAboutDialog (aboutdialog.png)
//
// About dialogs often contain links and email addresses. `GtkAboutDialog`
// displays these as clickable links. By default, it calls [func@Gtk.show_uri]
// when a user clicks one. The behaviour can be overridden with the
// [signal@Gtk.AboutDialog::activate-link] signal.
//
// To specify a person with an email address, use a string like `Edgar Allan Poe
// <edgar@poe.com>`. To specify a website with a title, use a string like `GTK
// team https://www.gtk.org`.
//
// To make constructing a `GtkAboutDialog` as convenient as possible, you can
// use the function [func@Gtk.show_about_dialog] which constructs and shows a
// dialog and keeps it around so that it can be shown again.
//
// Note that GTK sets a default title of `_("About s")` on the dialog window
// (where `s` is replaced by the name of the application, but in order to ensure
// proper translation of the title, applications should set the title property
// explicitly when constructing a `GtkAboutDialog`, as shown in the following
// example:
//
// “`c GFile *logo_file = g_file_new_for_path ("./logo.png"); GdkTexture
// *example_logo = gdk_texture_new_from_file (logo_file, NULL); g_object_unref
// (logo_file);
//
// gtk_show_about_dialog (NULL, "program-name", "ExampleCode", "logo",
// example_logo, "title", _("About ExampleCode"), NULL); “`
//
//
// CSS nodes
//
// `GtkAboutDialog` has a single CSS node with the name `window` and style class
// `.aboutdialog`.
type AboutDialog struct {
	Window
}

var (
	_ AboutDialogger  = (*AboutDialog)(nil)
	_ gextras.Nativer = (*AboutDialog)(nil)
)

func wrapAboutDialog(obj *externglib.Object) AboutDialogger {
	return &AboutDialog{
		Window: Window{
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
			Root: Root{
				Native: Native{
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
				},
			},
			ShortcutManager: ShortcutManager{
				Object: obj,
			},
		},
	}
}

func marshalAboutDialogger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAboutDialog(obj), nil
}

// NewAboutDialog creates a new `GtkAboutDialog`.
func NewAboutDialog() *AboutDialog {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_about_dialog_new()

	var _aboutDialog *AboutDialog // out

	_aboutDialog = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*AboutDialog)

	return _aboutDialog
}

// AddCreditSection creates a new section in the "Credits" page.
func (about *AboutDialog) AddCreditSection(sectionName string, people []string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.char           // out
	var _arg2 **C.char

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (*C.char)(C.CString(sectionName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (**C.char)(C.malloc(C.ulong(len(people)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(people))
		for i := range people {
			out[i] = (*C.char)(C.CString(people[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_about_dialog_add_credit_section(_arg0, _arg1, _arg2)
}

// Artists returns the string which are displayed in the "Artists" tab of the
// secondary credits dialog.
func (about *AboutDialog) Artists() []string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret **C.char

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_artists(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
		}
	}

	return _utf8s
}

// Authors returns the string which are displayed in the authors tab of the
// secondary credits dialog.
func (about *AboutDialog) Authors() []string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret **C.char

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_authors(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
		}
	}

	return _utf8s
}

// Comments returns the comments string.
func (about *AboutDialog) Comments() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_comments(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Copyright returns the copyright string.
func (about *AboutDialog) Copyright() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_copyright(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Documenters returns the string which are displayed in the "Documenters" tab
// of the secondary credits dialog.
func (about *AboutDialog) Documenters() []string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret **C.char

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_documenters(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
		}
	}

	return _utf8s
}

// License returns the license information.
func (about *AboutDialog) License() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_license(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// LicenseType retrieves the license type.
func (about *AboutDialog) LicenseType() License {
	var _arg0 *C.GtkAboutDialog // out
	var _cret C.GtkLicense      // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_license_type(_arg0)

	var _license License // out

	_license = (License)(_cret)

	return _license
}

// Logo returns the paintable displayed as logo in the about dialog.
func (about *AboutDialog) Logo() *gdk.Paintable {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.GdkPaintable   // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_logo(_arg0)

	var _paintable *gdk.Paintable // out

	_paintable = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gdk.Paintable)

	return _paintable
}

// LogoIconName returns the icon name displayed as logo in the about dialog.
func (about *AboutDialog) LogoIconName() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_logo_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// ProgramName returns the program name displayed in the about dialog.
func (about *AboutDialog) ProgramName() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_program_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// SystemInformation returns the system information that is shown in the about
// dialog.
func (about *AboutDialog) SystemInformation() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_system_information(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// TranslatorCredits returns the translator credits string which is displayed in
// the translators tab of the secondary credits dialog.
func (about *AboutDialog) TranslatorCredits() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_translator_credits(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Version returns the version string.
func (about *AboutDialog) Version() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_version(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Website returns the website URL.
func (about *AboutDialog) Website() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_website(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// WebsiteLabel returns the label used for the website link.
func (about *AboutDialog) WebsiteLabel() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_website_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// WrapLicense returns whether the license text in the about dialog is
// automatically wrapped.
func (about *AboutDialog) WrapLicense() bool {
	var _arg0 *C.GtkAboutDialog // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_wrap_license(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetArtists sets the strings which are displayed in the "Artists" tab of the
// secondary credits dialog.
func (about *AboutDialog) SetArtists(artists []string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 **C.char

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (**C.char)(C.malloc(C.ulong(len(artists)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(artists))
		for i := range artists {
			out[i] = (*C.char)(C.CString(artists[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_about_dialog_set_artists(_arg0, _arg1)
}

// SetAuthors sets the strings which are displayed in the "Authors" tab of the
// secondary credits dialog.
func (about *AboutDialog) SetAuthors(authors []string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 **C.char

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (**C.char)(C.malloc(C.ulong(len(authors)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(authors))
		for i := range authors {
			out[i] = (*C.char)(C.CString(authors[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_about_dialog_set_authors(_arg0, _arg1)
}

// SetComments sets the comments string to display in the about dialog.
//
// This should be a short string of one or two lines.
func (about *AboutDialog) SetComments(comments string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (*C.char)(C.CString(comments))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_comments(_arg0, _arg1)
}

// SetCopyright sets the copyright string to display in the about dialog.
//
// This should be a short string of one or two lines.
func (about *AboutDialog) SetCopyright(copyright string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (*C.char)(C.CString(copyright))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_copyright(_arg0, _arg1)
}

// SetDocumenters sets the strings which are displayed in the "Documenters" tab
// of the credits dialog.
func (about *AboutDialog) SetDocumenters(documenters []string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 **C.char

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (**C.char)(C.malloc(C.ulong(len(documenters)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(documenters))
		for i := range documenters {
			out[i] = (*C.char)(C.CString(documenters[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_about_dialog_set_documenters(_arg0, _arg1)
}

// SetLicense sets the license information to be displayed in the secondary
// license dialog.
//
// If `license` is `NULL`, the license button is hidden.
func (about *AboutDialog) SetLicense(license string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (*C.char)(C.CString(license))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_license(_arg0, _arg1)
}

// SetLogo sets the logo in the about dialog.
func (about *AboutDialog) SetLogo(logo gdk.Paintabler) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.GdkPaintable   // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (*C.GdkPaintable)(unsafe.Pointer((logo).(gextras.Nativer).Native()))

	C.gtk_about_dialog_set_logo(_arg0, _arg1)
}

// SetLogoIconName sets the icon name to be displayed as logo in the about
// dialog.
func (about *AboutDialog) SetLogoIconName(iconName string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (*C.char)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_logo_icon_name(_arg0, _arg1)
}

// SetProgramName sets the name to display in the about dialog.
//
// If `name` is not set, it defaults to `g_get_application_name()`.
func (about *AboutDialog) SetProgramName(name string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_program_name(_arg0, _arg1)
}

// SetSystemInformation sets the system information to be displayed in the about
// dialog.
//
// If `system_information` is `NULL`, the system information tab is hidden.
//
// See [property@Gtk.AboutDialog:system-information].
func (about *AboutDialog) SetSystemInformation(systemInformation string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (*C.char)(C.CString(systemInformation))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_system_information(_arg0, _arg1)
}

// SetTranslatorCredits sets the translator credits string which is displayed in
// the translators tab of the secondary credits dialog.
//
// The intended use for this string is to display the translator of the language
// which is currently used in the user interface. Using `gettext()`, a simple
// way to achieve that is to mark the string for translation:
//
// “`c GtkWidget *about = gtk_about_dialog_new ();
// gtk_about_dialog_set_translator_credits (GTK_ABOUT_DIALOG (about),
// _("translator-credits")); “`
//
// It is a good idea to use the customary `msgid` “translator-credits” for this
// purpose, since translators will already know the purpose of that `msgid`, and
// since `GtkAboutDialog` will detect if “translator-credits” is untranslated
// and hide the tab.
func (about *AboutDialog) SetTranslatorCredits(translatorCredits string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (*C.char)(C.CString(translatorCredits))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_translator_credits(_arg0, _arg1)
}

// SetVersion sets the version string to display in the about dialog.
func (about *AboutDialog) SetVersion(version string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (*C.char)(C.CString(version))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_version(_arg0, _arg1)
}

// SetWebsite sets the URL to use for the website link.
func (about *AboutDialog) SetWebsite(website string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (*C.char)(C.CString(website))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_website(_arg0, _arg1)
}

// SetWebsiteLabel sets the label to be used for the website link.
func (about *AboutDialog) SetWebsiteLabel(websiteLabel string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (*C.char)(C.CString(websiteLabel))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_website_label(_arg0, _arg1)
}

// SetWrapLicense sets whether the license text in the about dialog should be
// automatically wrapped.
func (about *AboutDialog) SetWrapLicense(wrapLicense bool) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	if wrapLicense {
		_arg1 = C.TRUE
	}

	C.gtk_about_dialog_set_wrap_license(_arg0, _arg1)
}
