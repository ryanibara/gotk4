// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_about_dialog_get_type()), F: marshalAboutDialog},
	})
}

// AboutDialog: the GtkAboutDialog offers a simple way to display information
// about a program like its logo, name, copyright, website and license. It is
// also possible to give credits to the authors, documenters, translators and
// artists who have worked on the program. An about dialog is typically opened
// when the user selects the `About` option from the `Help` menu. All parts of
// the dialog are optional.
//
// About dialogs often contain links and email addresses. GtkAboutDialog
// displays these as clickable links. By default, it calls
// gtk_show_uri_on_window() when a user clicks one. The behaviour can be
// overridden with the AboutDialog::activate-link signal.
//
// To specify a person with an email address, use a string like "Edgar Allan Poe
// <edgar\@poe.com>". To specify a website with a title, use a string like "GTK+
// team http://www.gtk.org".
//
// To make constructing a GtkAboutDialog as convenient as possible, you can use
// the function gtk_show_about_dialog() which constructs and shows a dialog and
// keeps it around so that it can be shown again.
//
// Note that GTK+ sets a default title of `_("About s")` on the dialog window
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
type AboutDialog interface {
	Dialog
	Buildable

	// AddCreditSection creates a new section in the Credits page.
	AddCreditSection(sectionName string, people []string)
	// Artists returns the string which are displayed in the artists tab of the
	// secondary credits dialog.
	Artists() []string
	// Authors returns the string which are displayed in the authors tab of the
	// secondary credits dialog.
	Authors() []string
	// Comments returns the comments string.
	Comments() string
	// Copyright returns the copyright string.
	Copyright() string
	// Documenters returns the string which are displayed in the documenters tab
	// of the secondary credits dialog.
	Documenters() []string
	// License returns the license information.
	License() string
	// LicenseType retrieves the license set using
	// gtk_about_dialog_set_license_type()
	LicenseType() License
	// Logo returns the pixbuf displayed as logo in the about dialog.
	Logo() gdkpixbuf.Pixbuf
	// LogoIconName returns the icon name displayed as logo in the about dialog.
	LogoIconName() string
	// ProgramName returns the program name displayed in the about dialog.
	ProgramName() string
	// TranslatorCredits returns the translator credits string which is
	// displayed in the translators tab of the secondary credits dialog.
	TranslatorCredits() string
	// Version returns the version string.
	Version() string
	// Website returns the website URL.
	Website() string
	// WebsiteLabel returns the label used for the website link.
	WebsiteLabel() string
	// WrapLicense returns whether the license text in @about is automatically
	// wrapped.
	WrapLicense() bool
	// SetArtists sets the strings which are displayed in the artists tab of the
	// secondary credits dialog.
	SetArtists(artists []string)
	// SetAuthors sets the strings which are displayed in the authors tab of the
	// secondary credits dialog.
	SetAuthors(authors []string)
	// SetComments sets the comments string to display in the about dialog. This
	// should be a short string of one or two lines.
	SetComments(comments string)
	// SetCopyright sets the copyright string to display in the about dialog.
	// This should be a short string of one or two lines.
	SetCopyright(copyright string)
	// SetDocumenters sets the strings which are displayed in the documenters
	// tab of the secondary credits dialog.
	SetDocumenters(documenters []string)
	// SetLicense sets the license information to be displayed in the secondary
	// license dialog. If @license is nil, the license button is hidden.
	SetLicense(license string)
	// SetLicenseType sets the license of the application showing the @about
	// dialog from a list of known licenses.
	//
	// This function overrides the license set using
	// gtk_about_dialog_set_license().
	SetLicenseType(licenseType License)
	// SetLogo sets the pixbuf to be displayed as logo in the about dialog. If
	// it is nil, the default window icon set with gtk_window_set_default_icon()
	// will be used.
	SetLogo(logo gdkpixbuf.Pixbuf)
	// SetLogoIconName sets the pixbuf to be displayed as logo in the about
	// dialog. If it is nil, the default window icon set with
	// gtk_window_set_default_icon() will be used.
	SetLogoIconName(iconName string)
	// SetProgramName sets the name to display in the about dialog. If this is
	// not set, it defaults to g_get_application_name().
	SetProgramName(name string)
	// SetTranslatorCredits sets the translator credits string which is
	// displayed in the translators tab of the secondary credits dialog.
	//
	// The intended use for this string is to display the translator of the
	// language which is currently used in the user interface. Using gettext(),
	// a simple way to achieve that is to mark the string for translation:
	//
	//    GtkWidget *about = gtk_about_dialog_new ();
	//    gtk_about_dialog_set_translator_credits (GTK_ABOUT_DIALOG (about),
	//                                             _("translator-credits"));
	//
	// It is a good idea to use the customary msgid “translator-credits” for
	// this purpose, since translators will already know the purpose of that
	// msgid, and since AboutDialog will detect if “translator-credits” is
	// untranslated and hide the tab.
	SetTranslatorCredits(translatorCredits string)
	// SetVersion sets the version string to display in the about dialog.
	SetVersion(version string)
	// SetWebsite sets the URL to use for the website link.
	SetWebsite(website string)
	// SetWebsiteLabel sets the label to be used for the website link.
	SetWebsiteLabel(websiteLabel string)
	// SetWrapLicense sets whether the license text in @about is automatically
	// wrapped.
	SetWrapLicense(wrapLicense bool)
}

// aboutDialog implements the AboutDialog interface.
type aboutDialog struct {
	Dialog
	Buildable
}

var _ AboutDialog = (*aboutDialog)(nil)

// WrapAboutDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapAboutDialog(obj *externglib.Object) AboutDialog {
	return AboutDialog{
		Dialog:    WrapDialog(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalAboutDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAboutDialog(obj), nil
}

// NewAboutDialog constructs a class AboutDialog.
func NewAboutDialog() AboutDialog {
	var cret C.GtkAboutDialog
	var ret1 AboutDialog

	cret = C.gtk_about_dialog_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(AboutDialog)

	return ret1
}

// AddCreditSection creates a new section in the Credits page.
func (a aboutDialog) AddCreditSection(sectionName string, people []string) {
	var arg0 *C.GtkAboutDialog
	var arg1 *C.gchar
	var arg2 **C.gchar

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	arg1 = (*C.gchar)(C.CString(sectionName))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.malloc(len(people) * (unsafe.Sizeof(int(0)) + 1))
	defer C.free(unsafe.Pointer(arg2))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg2), int(len(people)))

		for i := range people {
			out[i] = (*C.gchar)(C.CString(people[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_about_dialog_add_credit_section(arg0, sectionName, people)
}

// Artists returns the string which are displayed in the artists tab of the
// secondary credits dialog.
func (a aboutDialog) Artists() []string {
	var arg0 *C.GtkAboutDialog

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	var cret **C.gchar
	var ret1 []string

	cret = C.gtk_about_dialog_get_artists(arg0)

	{
		var length int
		for p := cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		ret1 = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
			ret1[i] = C.GoString(src)
		}
	}

	return ret1
}

// Authors returns the string which are displayed in the authors tab of the
// secondary credits dialog.
func (a aboutDialog) Authors() []string {
	var arg0 *C.GtkAboutDialog

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	var cret **C.gchar
	var ret1 []string

	cret = C.gtk_about_dialog_get_authors(arg0)

	{
		var length int
		for p := cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		ret1 = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
			ret1[i] = C.GoString(src)
		}
	}

	return ret1
}

// Comments returns the comments string.
func (a aboutDialog) Comments() string {
	var arg0 *C.GtkAboutDialog

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_about_dialog_get_comments(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Copyright returns the copyright string.
func (a aboutDialog) Copyright() string {
	var arg0 *C.GtkAboutDialog

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_about_dialog_get_copyright(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Documenters returns the string which are displayed in the documenters tab
// of the secondary credits dialog.
func (a aboutDialog) Documenters() []string {
	var arg0 *C.GtkAboutDialog

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	var cret **C.gchar
	var ret1 []string

	cret = C.gtk_about_dialog_get_documenters(arg0)

	{
		var length int
		for p := cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		ret1 = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
			ret1[i] = C.GoString(src)
		}
	}

	return ret1
}

// License returns the license information.
func (a aboutDialog) License() string {
	var arg0 *C.GtkAboutDialog

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_about_dialog_get_license(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// LicenseType retrieves the license set using
// gtk_about_dialog_set_license_type()
func (a aboutDialog) LicenseType() License {
	var arg0 *C.GtkAboutDialog

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	var cret C.GtkLicense
	var ret1 License

	cret = C.gtk_about_dialog_get_license_type(arg0)

	ret1 = License(cret)

	return ret1
}

// Logo returns the pixbuf displayed as logo in the about dialog.
func (a aboutDialog) Logo() gdkpixbuf.Pixbuf {
	var arg0 *C.GtkAboutDialog

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	var cret *C.GdkPixbuf
	var ret1 gdkpixbuf.Pixbuf

	cret = C.gtk_about_dialog_get_logo(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdkpixbuf.Pixbuf)

	return ret1
}

// LogoIconName returns the icon name displayed as logo in the about dialog.
func (a aboutDialog) LogoIconName() string {
	var arg0 *C.GtkAboutDialog

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_about_dialog_get_logo_icon_name(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// ProgramName returns the program name displayed in the about dialog.
func (a aboutDialog) ProgramName() string {
	var arg0 *C.GtkAboutDialog

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_about_dialog_get_program_name(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// TranslatorCredits returns the translator credits string which is
// displayed in the translators tab of the secondary credits dialog.
func (a aboutDialog) TranslatorCredits() string {
	var arg0 *C.GtkAboutDialog

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_about_dialog_get_translator_credits(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Version returns the version string.
func (a aboutDialog) Version() string {
	var arg0 *C.GtkAboutDialog

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_about_dialog_get_version(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Website returns the website URL.
func (a aboutDialog) Website() string {
	var arg0 *C.GtkAboutDialog

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_about_dialog_get_website(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// WebsiteLabel returns the label used for the website link.
func (a aboutDialog) WebsiteLabel() string {
	var arg0 *C.GtkAboutDialog

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_about_dialog_get_website_label(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// WrapLicense returns whether the license text in @about is automatically
// wrapped.
func (a aboutDialog) WrapLicense() bool {
	var arg0 *C.GtkAboutDialog

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_about_dialog_get_wrap_license(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// SetArtists sets the strings which are displayed in the artists tab of the
// secondary credits dialog.
func (a aboutDialog) SetArtists(artists []string) {
	var arg0 *C.GtkAboutDialog
	var arg1 **C.gchar

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	arg1 = C.malloc(len(artists) * (unsafe.Sizeof(int(0)) + 1))
	defer C.free(unsafe.Pointer(arg1))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg1), int(len(artists)))

		for i := range artists {
			out[i] = (*C.gchar)(C.CString(artists[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_about_dialog_set_artists(arg0, artists)
}

// SetAuthors sets the strings which are displayed in the authors tab of the
// secondary credits dialog.
func (a aboutDialog) SetAuthors(authors []string) {
	var arg0 *C.GtkAboutDialog
	var arg1 **C.gchar

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	arg1 = C.malloc(len(authors) * (unsafe.Sizeof(int(0)) + 1))
	defer C.free(unsafe.Pointer(arg1))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg1), int(len(authors)))

		for i := range authors {
			out[i] = (*C.gchar)(C.CString(authors[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_about_dialog_set_authors(arg0, authors)
}

// SetComments sets the comments string to display in the about dialog. This
// should be a short string of one or two lines.
func (a aboutDialog) SetComments(comments string) {
	var arg0 *C.GtkAboutDialog
	var arg1 *C.gchar

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	arg1 = (*C.gchar)(C.CString(comments))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_about_dialog_set_comments(arg0, comments)
}

// SetCopyright sets the copyright string to display in the about dialog.
// This should be a short string of one or two lines.
func (a aboutDialog) SetCopyright(copyright string) {
	var arg0 *C.GtkAboutDialog
	var arg1 *C.gchar

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	arg1 = (*C.gchar)(C.CString(copyright))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_about_dialog_set_copyright(arg0, copyright)
}

// SetDocumenters sets the strings which are displayed in the documenters
// tab of the secondary credits dialog.
func (a aboutDialog) SetDocumenters(documenters []string) {
	var arg0 *C.GtkAboutDialog
	var arg1 **C.gchar

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	arg1 = C.malloc(len(documenters) * (unsafe.Sizeof(int(0)) + 1))
	defer C.free(unsafe.Pointer(arg1))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg1), int(len(documenters)))

		for i := range documenters {
			out[i] = (*C.gchar)(C.CString(documenters[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_about_dialog_set_documenters(arg0, documenters)
}

// SetLicense sets the license information to be displayed in the secondary
// license dialog. If @license is nil, the license button is hidden.
func (a aboutDialog) SetLicense(license string) {
	var arg0 *C.GtkAboutDialog
	var arg1 *C.gchar

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	arg1 = (*C.gchar)(C.CString(license))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_about_dialog_set_license(arg0, license)
}

// SetLicenseType sets the license of the application showing the @about
// dialog from a list of known licenses.
//
// This function overrides the license set using
// gtk_about_dialog_set_license().
func (a aboutDialog) SetLicenseType(licenseType License) {
	var arg0 *C.GtkAboutDialog
	var arg1 C.GtkLicense

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	arg1 = (C.GtkLicense)(licenseType)

	C.gtk_about_dialog_set_license_type(arg0, licenseType)
}

// SetLogo sets the pixbuf to be displayed as logo in the about dialog. If
// it is nil, the default window icon set with gtk_window_set_default_icon()
// will be used.
func (a aboutDialog) SetLogo(logo gdkpixbuf.Pixbuf) {
	var arg0 *C.GtkAboutDialog
	var arg1 *C.GdkPixbuf

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GdkPixbuf)(unsafe.Pointer(logo.Native()))

	C.gtk_about_dialog_set_logo(arg0, logo)
}

// SetLogoIconName sets the pixbuf to be displayed as logo in the about
// dialog. If it is nil, the default window icon set with
// gtk_window_set_default_icon() will be used.
func (a aboutDialog) SetLogoIconName(iconName string) {
	var arg0 *C.GtkAboutDialog
	var arg1 *C.gchar

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_about_dialog_set_logo_icon_name(arg0, iconName)
}

// SetProgramName sets the name to display in the about dialog. If this is
// not set, it defaults to g_get_application_name().
func (a aboutDialog) SetProgramName(name string) {
	var arg0 *C.GtkAboutDialog
	var arg1 *C.gchar

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_about_dialog_set_program_name(arg0, name)
}

// SetTranslatorCredits sets the translator credits string which is
// displayed in the translators tab of the secondary credits dialog.
//
// The intended use for this string is to display the translator of the
// language which is currently used in the user interface. Using gettext(),
// a simple way to achieve that is to mark the string for translation:
//
//    GtkWidget *about = gtk_about_dialog_new ();
//    gtk_about_dialog_set_translator_credits (GTK_ABOUT_DIALOG (about),
//                                             _("translator-credits"));
//
// It is a good idea to use the customary msgid “translator-credits” for
// this purpose, since translators will already know the purpose of that
// msgid, and since AboutDialog will detect if “translator-credits” is
// untranslated and hide the tab.
func (a aboutDialog) SetTranslatorCredits(translatorCredits string) {
	var arg0 *C.GtkAboutDialog
	var arg1 *C.gchar

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	arg1 = (*C.gchar)(C.CString(translatorCredits))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_about_dialog_set_translator_credits(arg0, translatorCredits)
}

// SetVersion sets the version string to display in the about dialog.
func (a aboutDialog) SetVersion(version string) {
	var arg0 *C.GtkAboutDialog
	var arg1 *C.gchar

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	arg1 = (*C.gchar)(C.CString(version))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_about_dialog_set_version(arg0, version)
}

// SetWebsite sets the URL to use for the website link.
func (a aboutDialog) SetWebsite(website string) {
	var arg0 *C.GtkAboutDialog
	var arg1 *C.gchar

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	arg1 = (*C.gchar)(C.CString(website))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_about_dialog_set_website(arg0, website)
}

// SetWebsiteLabel sets the label to be used for the website link.
func (a aboutDialog) SetWebsiteLabel(websiteLabel string) {
	var arg0 *C.GtkAboutDialog
	var arg1 *C.gchar

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	arg1 = (*C.gchar)(C.CString(websiteLabel))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_about_dialog_set_website_label(arg0, websiteLabel)
}

// SetWrapLicense sets whether the license text in @about is automatically
// wrapped.
func (a aboutDialog) SetWrapLicense(wrapLicense bool) {
	var arg0 *C.GtkAboutDialog
	var arg1 C.gboolean

	arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	if wrapLicense {
		arg1 = C.gboolean(1)
	}

	C.gtk_about_dialog_set_wrap_license(arg0, wrapLicense)
}

type AboutDialogPrivate struct {
	native C.GtkAboutDialogPrivate
}

// WrapAboutDialogPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAboutDialogPrivate(ptr unsafe.Pointer) *AboutDialogPrivate {
	if ptr == nil {
		return nil
	}

	return (*AboutDialogPrivate)(ptr)
}

func marshalAboutDialogPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAboutDialogPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AboutDialogPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}
