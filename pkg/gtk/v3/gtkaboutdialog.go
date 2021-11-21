// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
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
type License C.gint

const (
	// LicenseUnknown: no license specified.
	LicenseUnknown License = iota
	// LicenseCustom: license text is going to be specified by the developer.
	LicenseCustom
	// LicenseGPL20: GNU General Public License, version 2.0 or later.
	LicenseGPL20
	// LicenseGPL30: GNU General Public License, version 3.0 or later.
	LicenseGPL30
	// LicenseLGPL21: GNU Lesser General Public License, version 2.1 or later.
	LicenseLGPL21
	// LicenseLGPL30: GNU Lesser General Public License, version 3.0 or later.
	LicenseLGPL30
	// LicenseBSD: BSD standard license.
	LicenseBSD
	// LicenseMITX11: MIT/X11 standard license.
	LicenseMITX11
	// LicenseArtistic: artistic License, version 2.0.
	LicenseArtistic
	// LicenseGPL20_Only: GNU General Public License, version 2.0 only. Since
	// 3.12.
	LicenseGPL20_Only
	// LicenseGPL30_Only: GNU General Public License, version 3.0 only. Since
	// 3.12.
	LicenseGPL30_Only
	// LicenseLGPL21_Only: GNU Lesser General Public License, version 2.1 only.
	// Since 3.12.
	LicenseLGPL21_Only
	// LicenseLGPL30_Only: GNU Lesser General Public License, version 3.0 only.
	// Since 3.12.
	LicenseLGPL30_Only
	// LicenseAGPL30: GNU Affero General Public License, version 3.0 or later.
	// Since: 3.22.
	LicenseAGPL30
	// LicenseAGPL30_Only: GNU Affero General Public License, version 3.0 only.
	// Since: 3.22.27.
	LicenseAGPL30_Only
	// LicenseBSD3: 3-clause BSD licence. Since: 3.24.20.
	LicenseBSD3
	// LicenseApache20: apache License, version 2.0. Since: 3.24.20.
	LicenseApache20
	// LicenseMPL20: mozilla Public License, version 2.0. Since: 3.24.20.
	LicenseMPL20
)

func marshalLicense(p uintptr) (interface{}, error) {
	return License(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for License.
func (l License) String() string {
	switch l {
	case LicenseUnknown:
		return "Unknown"
	case LicenseCustom:
		return "Custom"
	case LicenseGPL20:
		return "GPL20"
	case LicenseGPL30:
		return "GPL30"
	case LicenseLGPL21:
		return "LGPL21"
	case LicenseLGPL30:
		return "LGPL30"
	case LicenseBSD:
		return "BSD"
	case LicenseMITX11:
		return "MITX11"
	case LicenseArtistic:
		return "Artistic"
	case LicenseGPL20_Only:
		return "GPL20_Only"
	case LicenseGPL30_Only:
		return "GPL30_Only"
	case LicenseLGPL21_Only:
		return "LGPL21_Only"
	case LicenseLGPL30_Only:
		return "LGPL30_Only"
	case LicenseAGPL30:
		return "AGPL30"
	case LicenseAGPL30_Only:
		return "AGPL30_Only"
	case LicenseBSD3:
		return "BSD3"
	case LicenseApache20:
		return "Apache20"
	case LicenseMPL20:
		return "MPL20"
	default:
		return fmt.Sprintf("License(%d)", l)
	}
}

// AboutDialogOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type AboutDialogOverrider interface {
	ActivateLink(uri string) bool
}

// AboutDialog offers a simple way to display information about a program like
// its logo, name, copyright, website and license. It is also possible to give
// credits to the authors, documenters, translators and artists who have worked
// on the program. An about dialog is typically opened when the user selects the
// About option from the Help menu. All parts of the dialog are optional.
//
// About dialogs often contain links and email addresses. GtkAboutDialog
// displays these as clickable links. By default, it calls
// gtk_show_uri_on_window() when a user clicks one. The behaviour can be
// overridden with the AboutDialog::activate-link signal.
//
// To specify a person with an email address, use a string like "Edgar Allan Poe
// <edgar\poe.com>". To specify a website with a title, use a string like "GTK+
// team http://www.gtk.org".
//
// To make constructing a GtkAboutDialog as convenient as possible, you can use
// the function gtk_show_about_dialog() which constructs and shows a dialog and
// keeps it around so that it can be shown again.
//
// Note that GTK+ sets a default title of _("About s") on the dialog window
// (where \s is replaced by the name of the application, but in order to ensure
// proper translation of the title, applications should set the title property
// explicitly when constructing a GtkAboutDialog, as shown in the following
// example:
//
//    GdkPixbuf *example_logo = gdk_pixbuf_new_from_file ("./logo.png", NULL);
//    gtk_show_about_dialog (NULL,
//                           "program-name", "ExampleCode",
//                           "logo", example_logo,
//                           "title", _("About ExampleCode"),
//                           NULL);
//
// It is also possible to show a AboutDialog like any other Dialog, e.g. using
// gtk_dialog_run(). In this case, you might need to know that the “Close”
// button returns the K_RESPONSE_CANCEL response id.
type AboutDialog struct {
	Dialog
}

var (
	_ Binner = (*AboutDialog)(nil)
)

func wrapAboutDialog(obj *externglib.Object) *AboutDialog {
	return &AboutDialog{
		Dialog: Dialog{
			Window: Window{
				Bin: Bin{
					Container: Container{
						Widget: Widget{
							InitiallyUnowned: externglib.InitiallyUnowned{
								Object: obj,
							},
							Object: obj,
							ImplementorIface: atk.ImplementorIface{
								Object: obj,
							},
							Buildable: Buildable{
								Object: obj,
							},
						},
					},
				},
			},
		},
	}
}

func marshalAboutDialogger(p uintptr) (interface{}, error) {
	return wrapAboutDialog(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAboutDialog creates a new AboutDialog.
func NewAboutDialog() *AboutDialog {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_about_dialog_new()

	var _aboutDialog *AboutDialog // out

	_aboutDialog = wrapAboutDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _aboutDialog
}

// AddCreditSection creates a new section in the Credits page.
//
// The function takes the following parameters:
//
//    - sectionName: name of the section.
//    - people who belong to that section.
//
func (about *AboutDialog) AddCreditSection(sectionName string, people []string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out
	var _arg2 **C.gchar         // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(sectionName)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		_arg2 = (**C.gchar)(C.malloc(C.size_t(uint((len(people) + 1)) * uint(unsafe.Sizeof(uint(0))))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(people)+1)
			var zero *C.gchar
			out[len(people)] = zero
			for i := range people {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(people[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.gtk_about_dialog_add_credit_section(_arg0, _arg1, _arg2)
	runtime.KeepAlive(about)
	runtime.KeepAlive(sectionName)
	runtime.KeepAlive(people)
}

// Artists returns the string which are displayed in the artists tab of the
// secondary credits dialog.
func (about *AboutDialog) Artists() []string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret **C.gchar         // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_artists(_arg0)
	runtime.KeepAlive(about)

	var _utf8s []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// Authors returns the string which are displayed in the authors tab of the
// secondary credits dialog.
func (about *AboutDialog) Authors() []string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret **C.gchar         // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_authors(_arg0)
	runtime.KeepAlive(about)

	var _utf8s []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// Comments returns the comments string.
func (about *AboutDialog) Comments() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_comments(_arg0)
	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Copyright returns the copyright string.
func (about *AboutDialog) Copyright() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_copyright(_arg0)
	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Documenters returns the string which are displayed in the documenters tab of
// the secondary credits dialog.
func (about *AboutDialog) Documenters() []string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret **C.gchar         // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_documenters(_arg0)
	runtime.KeepAlive(about)

	var _utf8s []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// License returns the license information.
func (about *AboutDialog) License() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_license(_arg0)
	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// LicenseType retrieves the license set using
// gtk_about_dialog_set_license_type().
func (about *AboutDialog) LicenseType() License {
	var _arg0 *C.GtkAboutDialog // out
	var _cret C.GtkLicense      // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_license_type(_arg0)
	runtime.KeepAlive(about)

	var _license License // out

	_license = License(_cret)

	return _license
}

// Logo returns the pixbuf displayed as logo in the about dialog.
func (about *AboutDialog) Logo() *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.GdkPixbuf      // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_logo(_arg0)
	runtime.KeepAlive(about)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	return _pixbuf
}

// LogoIconName returns the icon name displayed as logo in the about dialog.
func (about *AboutDialog) LogoIconName() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_logo_icon_name(_arg0)
	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ProgramName returns the program name displayed in the about dialog.
func (about *AboutDialog) ProgramName() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_program_name(_arg0)
	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// TranslatorCredits returns the translator credits string which is displayed in
// the translators tab of the secondary credits dialog.
func (about *AboutDialog) TranslatorCredits() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_translator_credits(_arg0)
	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Version returns the version string.
func (about *AboutDialog) Version() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_version(_arg0)
	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Website returns the website URL.
func (about *AboutDialog) Website() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_website(_arg0)
	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// WebsiteLabel returns the label used for the website link.
func (about *AboutDialog) WebsiteLabel() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_website_label(_arg0)
	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// WrapLicense returns whether the license text in about is automatically
// wrapped.
func (about *AboutDialog) WrapLicense() bool {
	var _arg0 *C.GtkAboutDialog // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))

	_cret = C.gtk_about_dialog_get_wrap_license(_arg0)
	runtime.KeepAlive(about)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetArtists sets the strings which are displayed in the artists tab of the
// secondary credits dialog.
//
// The function takes the following parameters:
//
//    - artists: NULL-terminated array of strings.
//
func (about *AboutDialog) SetArtists(artists []string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 **C.gchar         // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	{
		_arg1 = (**C.gchar)(C.malloc(C.size_t(uint((len(artists) + 1)) * uint(unsafe.Sizeof(uint(0))))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(artists)+1)
			var zero *C.gchar
			out[len(artists)] = zero
			for i := range artists {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(artists[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.gtk_about_dialog_set_artists(_arg0, _arg1)
	runtime.KeepAlive(about)
	runtime.KeepAlive(artists)
}

// SetAuthors sets the strings which are displayed in the authors tab of the
// secondary credits dialog.
//
// The function takes the following parameters:
//
//    - authors: NULL-terminated array of strings.
//
func (about *AboutDialog) SetAuthors(authors []string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 **C.gchar         // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	{
		_arg1 = (**C.gchar)(C.malloc(C.size_t(uint((len(authors) + 1)) * uint(unsafe.Sizeof(uint(0))))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(authors)+1)
			var zero *C.gchar
			out[len(authors)] = zero
			for i := range authors {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(authors[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.gtk_about_dialog_set_authors(_arg0, _arg1)
	runtime.KeepAlive(about)
	runtime.KeepAlive(authors)
}

// SetComments sets the comments string to display in the about dialog. This
// should be a short string of one or two lines.
//
// The function takes the following parameters:
//
//    - comments string.
//
func (about *AboutDialog) SetComments(comments string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	if comments != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(comments)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_about_dialog_set_comments(_arg0, _arg1)
	runtime.KeepAlive(about)
	runtime.KeepAlive(comments)
}

// SetCopyright sets the copyright string to display in the about dialog. This
// should be a short string of one or two lines.
//
// The function takes the following parameters:
//
//    - copyright string.
//
func (about *AboutDialog) SetCopyright(copyright string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	if copyright != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(copyright)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_about_dialog_set_copyright(_arg0, _arg1)
	runtime.KeepAlive(about)
	runtime.KeepAlive(copyright)
}

// SetDocumenters sets the strings which are displayed in the documenters tab of
// the secondary credits dialog.
//
// The function takes the following parameters:
//
//    - documenters: NULL-terminated array of strings.
//
func (about *AboutDialog) SetDocumenters(documenters []string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 **C.gchar         // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	{
		_arg1 = (**C.gchar)(C.malloc(C.size_t(uint((len(documenters) + 1)) * uint(unsafe.Sizeof(uint(0))))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(documenters)+1)
			var zero *C.gchar
			out[len(documenters)] = zero
			for i := range documenters {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(documenters[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.gtk_about_dialog_set_documenters(_arg0, _arg1)
	runtime.KeepAlive(about)
	runtime.KeepAlive(documenters)
}

// SetLicense sets the license information to be displayed in the secondary
// license dialog. If license is NULL, the license button is hidden.
//
// The function takes the following parameters:
//
//    - license information or NULL.
//
func (about *AboutDialog) SetLicense(license string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	if license != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(license)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_about_dialog_set_license(_arg0, _arg1)
	runtime.KeepAlive(about)
	runtime.KeepAlive(license)
}

// SetLicenseType sets the license of the application showing the about dialog
// from a list of known licenses.
//
// This function overrides the license set using gtk_about_dialog_set_license().
//
// The function takes the following parameters:
//
//    - licenseType: type of license.
//
func (about *AboutDialog) SetLicenseType(licenseType License) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 C.GtkLicense      // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = C.GtkLicense(licenseType)

	C.gtk_about_dialog_set_license_type(_arg0, _arg1)
	runtime.KeepAlive(about)
	runtime.KeepAlive(licenseType)
}

// SetLogo sets the pixbuf to be displayed as logo in the about dialog. If it is
// NULL, the default window icon set with gtk_window_set_default_icon() will be
// used.
//
// The function takes the following parameters:
//
//    - logo or NULL.
//
func (about *AboutDialog) SetLogo(logo *gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.GdkPixbuf      // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	if logo != nil {
		_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(logo.Native()))
	}

	C.gtk_about_dialog_set_logo(_arg0, _arg1)
	runtime.KeepAlive(about)
	runtime.KeepAlive(logo)
}

// SetLogoIconName sets the pixbuf to be displayed as logo in the about dialog.
// If it is NULL, the default window icon set with gtk_window_set_default_icon()
// will be used.
//
// The function takes the following parameters:
//
//    - iconName: icon name, or NULL.
//
func (about *AboutDialog) SetLogoIconName(iconName string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	if iconName != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_about_dialog_set_logo_icon_name(_arg0, _arg1)
	runtime.KeepAlive(about)
	runtime.KeepAlive(iconName)
}

// SetProgramName sets the name to display in the about dialog. If this is not
// set, it defaults to g_get_application_name().
//
// The function takes the following parameters:
//
//    - name: program name.
//
func (about *AboutDialog) SetProgramName(name string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_program_name(_arg0, _arg1)
	runtime.KeepAlive(about)
	runtime.KeepAlive(name)
}

// SetTranslatorCredits sets the translator credits string which is displayed in
// the translators tab of the secondary credits dialog.
//
// The intended use for this string is to display the translator of the language
// which is currently used in the user interface. Using gettext(), a simple way
// to achieve that is to mark the string for translation:
//
//    GtkWidget *about = gtk_about_dialog_new ();
//    gtk_about_dialog_set_translator_credits (GTK_ABOUT_DIALOG (about),
//                                             _("translator-credits"));
//
// It is a good idea to use the customary msgid “translator-credits” for this
// purpose, since translators will already know the purpose of that msgid, and
// since AboutDialog will detect if “translator-credits” is untranslated and
// hide the tab.
//
// The function takes the following parameters:
//
//    - translatorCredits: translator credits.
//
func (about *AboutDialog) SetTranslatorCredits(translatorCredits string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	if translatorCredits != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(translatorCredits)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_about_dialog_set_translator_credits(_arg0, _arg1)
	runtime.KeepAlive(about)
	runtime.KeepAlive(translatorCredits)
}

// SetVersion sets the version string to display in the about dialog.
//
// The function takes the following parameters:
//
//    - version string.
//
func (about *AboutDialog) SetVersion(version string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	if version != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(version)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_about_dialog_set_version(_arg0, _arg1)
	runtime.KeepAlive(about)
	runtime.KeepAlive(version)
}

// SetWebsite sets the URL to use for the website link.
//
// The function takes the following parameters:
//
//    - website: URL string starting with "http://".
//
func (about *AboutDialog) SetWebsite(website string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	if website != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(website)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_about_dialog_set_website(_arg0, _arg1)
	runtime.KeepAlive(about)
	runtime.KeepAlive(website)
}

// SetWebsiteLabel sets the label to be used for the website link.
//
// The function takes the following parameters:
//
//    - websiteLabel: label used for the website link.
//
func (about *AboutDialog) SetWebsiteLabel(websiteLabel string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(websiteLabel)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_website_label(_arg0, _arg1)
	runtime.KeepAlive(about)
	runtime.KeepAlive(websiteLabel)
}

// SetWrapLicense sets whether the license text in about is automatically
// wrapped.
//
// The function takes the following parameters:
//
//    - wrapLicense: whether to wrap the license.
//
func (about *AboutDialog) SetWrapLicense(wrapLicense bool) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(about.Native()))
	if wrapLicense {
		_arg1 = C.TRUE
	}

	C.gtk_about_dialog_set_wrap_license(_arg0, _arg1)
	runtime.KeepAlive(about)
	runtime.KeepAlive(wrapLicense)
}

// ConnectActivateLink: signal which gets emitted to activate a URI.
// Applications may connect to it to override the default behaviour, which is to
// call gtk_show_uri_on_window().
func (about *AboutDialog) ConnectActivateLink(f func(uri string) bool) externglib.SignalHandle {
	return about.Connect("activate-link", f)
}
