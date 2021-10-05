// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_file_chooser_widget_get_type()), F: marshalFileChooserWidgetter},
	})
}

// FileChooserWidget is a widget for choosing files. It exposes the FileChooser
// interface, and you should use the methods of this interface to interact with
// the widget.
//
//
// CSS nodes
//
// GtkFileChooserWidget has a single CSS node with name filechooser.
type FileChooserWidget struct {
	Box

	FileChooser
	*externglib.Object
}

func wrapFileChooserWidget(obj *externglib.Object) *FileChooserWidget {
	return &FileChooserWidget{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					Object: obj,
				},
			},
			Orientable: Orientable{
				Object: obj,
			},
			Object: obj,
		},
		FileChooser: FileChooser{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalFileChooserWidgetter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFileChooserWidget(obj), nil
}

// NewFileChooserWidget creates a new FileChooserWidget. This is a file chooser
// widget that can be embedded in custom windows, and it is the same widget that
// is used by FileChooserDialog.
func NewFileChooserWidget(action FileChooserAction) *FileChooserWidget {
	var _arg1 C.GtkFileChooserAction // out
	var _cret *C.GtkWidget           // in

	_arg1 = C.GtkFileChooserAction(action)

	_cret = C.gtk_file_chooser_widget_new(_arg1)
	runtime.KeepAlive(action)

	var _fileChooserWidget *FileChooserWidget // out

	_fileChooserWidget = wrapFileChooserWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _fileChooserWidget
}

func (*FileChooserWidget) privateFileChooserWidget() {}

// ConnectDesktopFolder signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to make the file chooser show the user's Desktop folder in the
// file list.
//
// The default binding for this signal is Alt + D.
func (f *FileChooserWidget) ConnectDesktopFolder(f func()) glib.SignalHandle {
	return f.Connect("desktop-folder", f)
}

// ConnectDownFolder signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to make the file chooser go to a child of the current folder in
// the file hierarchy. The subfolder that will be used is displayed in the path
// bar widget of the file chooser. For example, if the path bar is showing
// "/foo/bar/baz", with bar currently displayed, then this will cause the file
// chooser to switch to the "baz" subfolder.
//
// The default binding for this signal is Alt + Down.
func (f *FileChooserWidget) ConnectDownFolder(f func()) glib.SignalHandle {
	return f.Connect("down-folder", f)
}

// ConnectHomeFolder signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to make the file chooser show the user's home folder in the file
// list.
//
// The default binding for this signal is Alt + Home.
func (f *FileChooserWidget) ConnectHomeFolder(f func()) glib.SignalHandle {
	return f.Connect("home-folder", f)
}

// ConnectLocationPopup signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to make the file chooser show a "Location" prompt which the user
// can use to manually type the name of the file he wishes to select.
//
// The default bindings for this signal are Control + L with a path string of ""
// (the empty string). It is also bound to / with a path string of "/" (a
// slash): this lets you type / and immediately type a path name. On Unix
// systems, this is bound to ~ (tilde) with a path string of "~" itself for
// access to home directories.
func (f *FileChooserWidget) ConnectLocationPopup(f func(path string)) glib.SignalHandle {
	return f.Connect("location-popup", f)
}

// ConnectLocationPopupOnPaste signal is a [keybinding signal][GtkBindingSignal]
// which gets emitted when the user asks for it.
//
// This is used to make the file chooser show a "Location" prompt when the user
// pastes into a FileChooserWidget.
//
// The default binding for this signal is Control + V.
func (f *FileChooserWidget) ConnectLocationPopupOnPaste(f func()) glib.SignalHandle {
	return f.Connect("location-popup-on-paste", f)
}

// ConnectLocationTogglePopup signal is a [keybinding signal][GtkBindingSignal]
// which gets emitted when the user asks for it.
//
// This is used to toggle the visibility of a "Location" prompt which the user
// can use to manually type the name of the file he wishes to select.
//
// The default binding for this signal is Control + L.
func (f *FileChooserWidget) ConnectLocationTogglePopup(f func()) glib.SignalHandle {
	return f.Connect("location-toggle-popup", f)
}

// ConnectPlacesShortcut signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to move the focus to the places sidebar.
//
// The default binding for this signal is Alt + P.
func (f *FileChooserWidget) ConnectPlacesShortcut(f func()) glib.SignalHandle {
	return f.Connect("places-shortcut", f)
}

// ConnectQuickBookmark signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to make the file chooser switch to the bookmark specified in the
// bookmark_index parameter. For example, if you have three bookmarks, you can
// pass 0, 1, 2 to this signal to switch to each of them, respectively.
//
// The default binding for this signal is Alt + 1, Alt + 2, etc. until Alt + 0.
// Note that in the default binding, that Alt + 1 is actually defined to switch
// to the bookmark at index 0, and so on successively; Alt + 0 is defined to
// switch to the bookmark at index 10.
func (f *FileChooserWidget) ConnectQuickBookmark(f func(bookmarkIndex int)) glib.SignalHandle {
	return f.Connect("quick-bookmark", f)
}

// ConnectRecentShortcut signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to make the file chooser show the Recent location.
//
// The default binding for this signal is Alt + R.
func (f *FileChooserWidget) ConnectRecentShortcut(f func()) glib.SignalHandle {
	return f.Connect("recent-shortcut", f)
}

// ConnectSearchShortcut signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to make the file chooser show the search entry.
//
// The default binding for this signal is Alt + S.
func (f *FileChooserWidget) ConnectSearchShortcut(f func()) glib.SignalHandle {
	return f.Connect("search-shortcut", f)
}

// ConnectShowHidden signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to make the file chooser display hidden files.
//
// The default binding for this signal is Control + H.
func (f *FileChooserWidget) ConnectShowHidden(f func()) glib.SignalHandle {
	return f.Connect("show-hidden", f)
}

// ConnectUpFolder signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted when the user asks for it.
//
// This is used to make the file chooser go to the parent of the current folder
// in the file hierarchy.
//
// The default binding for this signal is Alt + Up.
func (f *FileChooserWidget) ConnectUpFolder(f func()) glib.SignalHandle {
	return f.Connect("up-folder", f)
}
