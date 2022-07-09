// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_ListView_ConnectActivate(gpointer, guint, guintptr);
import "C"

// GTypeListView returns the GType for the type ListView.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeListView() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ListView").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalListView)
	return gtype
}

// ListView: GtkListView presents a large dynamic list of items.
//
// GtkListView uses its factory to generate one row widget for each visible item
// and shows them in a linear display, either vertically or horizontally.
//
// The gtk.ListView:show-separators property offers a simple way to display
// separators between the rows.
//
// GtkListView allows the user to select items according to the selection
// characteristics of the model. For models that allow multiple selected items,
// it is possible to turn on _rubberband selection_, using
// gtk.ListView:enable-rubberband.
//
// If you need multiple columns with headers, see gtk.ColumnView.
//
// To learn more about the list widget framework, see the overview
// (section-list-widget.html).
//
// An example of using GtkListView:
//
//    static void
//    setup_listitem_cb (GtkListItemFactory *factory,
//                       GtkListItem        *list_item)
//    {
//      GtkWidget *image;
//
//      image = gtk_image_new ();
//      gtk_image_set_icon_size (GTK_IMAGE (image), GTK_ICON_SIZE_LARGE);
//      gtk_list_item_set_child (list_item, image);
//    }
//
//    static void
//    bind_listitem_cb (GtkListItemFactory *factory,
//                      GtkListItem        *list_item)
//    {
//      GtkWidget *image;
//      GAppInfo *app_info;
//
//      image = gtk_list_item_get_child (list_item);
//      app_info = gtk_list_item_get_item (list_item);
//      gtk_image_set_from_gicon (GTK_IMAGE (image), g_app_info_get_icon (app_info));
//    }
//
//    static void
//    activate_cb (GtkListView  *list,
//                 guint         position,
//                 gpointer      unused)
//    {
//      GAppInfo *app_info;
//
//      app_info = g_list_model_get_item (G_LIST_MODEL (gtk_list_view_get_model (list)), position);
//      g_app_info_launch (app_info, NULL, NULL, NULL);
//      g_object_unref (app_info);
//    }
//
//    ...
//
//      model = create_application_list ();
//
//      factory = gtk_signal_list_item_factory_new ();
//      g_signal_connect (factory, "setup", G_CALLBACK (setup_listitem_cb), NULL);
//      g_signal_connect (factory, "bind", G_CALLBACK (bind_listitem_cb), NULL);
//
//      list = gtk_list_view_new (GTK_SELECTION_MODEL (gtk_single_selection_new (model)), factory);
//
//      g_signal_connect (list, "activate", G_CALLBACK (activate_cb), NULL);
//
//      gtk_scrolled_window_set_child (GTK_SCROLLED_WINDOW (sw), list);
//
//
// CSS nodes
//
//    listview[.separators][.rich-list][.navigation-sidebar][.data-table]
//    ├── row
//    │
//    ├── row
//    │
//    ┊
//    ╰── [rubberband]
//
//
// GtkListView uses a single CSS node named listview. It may carry the
// .separators style class, when GtkListView:show-separators property is set.
// Each child widget uses a single CSS node named row. For rubberband selection,
// a node with name rubberband is used.
//
// The main listview node may also carry style classes to select the style of
// list presentation (ListContainers.html#list-styles): .rich-list,
// .navigation-sidebar or .data-table.
//
//
// Accessibility
//
// GtkListView uses the GTK_ACCESSIBLE_ROLE_LIST role, and the list items use
// the GTK_ACCESSIBLE_ROLE_LIST_ITEM role.
type ListView struct {
	_ [0]func() // equal guard
	ListBase
}

var (
	_ ListBaser = (*ListView)(nil)
)

func wrapListView(obj *coreglib.Object) *ListView {
	return &ListView{
		ListBase: ListBase{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
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
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
			Scrollable: Scrollable{
				Object: obj,
			},
		},
	}
}

func marshalListView(p uintptr) (interface{}, error) {
	return wrapListView(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_ListView_ConnectActivate
func _gotk4_gtk4_ListView_ConnectActivate(arg0 C.gpointer, arg1 C.guint, arg2 C.guintptr) {
	var f func(position uint32)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(position uint32))
	}

	var _position uint32 // out

	_position = uint32(arg1)

	f(_position)
}

// ConnectActivate is emitted when a row has been activated by the user, usually
// via activating the GtkListView|list.activate-item action.
//
// This allows for a convenient way to handle activation in a listview. See
// gtk.ListItem.SetActivatable() for details on how to use this signal.
func (self *ListView) ConnectActivate(f func(position uint32)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "activate", false, unsafe.Pointer(C._gotk4_gtk4_ListView_ConnectActivate), f)
}

// NewListView creates a new GtkListView that uses the given factory for mapping
// items to widgets.
//
// The function takes ownership of the arguments, so you can write code like
//
//    list_view = gtk_list_view_new (create_model (),
//      gtk_builder_list_item_factory_new_from_resource ("/resource.ui"));.
//
// The function takes the following parameters:
//
//    - model (optional) to use, or NULL.
//    - factory (optional) to populate items with, or NULL.
//
// The function returns the following values:
//
//    - listView: new GtkListView using the given model and factory.
//
func NewListView(model SelectionModeller, factory *ListItemFactory) *ListView {
	var _args [2]girepository.Argument

	if model != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(model).Native()))
	}
	if factory != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(factory).Native()))
	}

	_gret := girepository.MustFind("Gtk", "ListView").InvokeMethod("new_ListView", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(model)
	runtime.KeepAlive(factory)

	var _listView *ListView // out

	_listView = wrapListView(coreglib.Take(unsafe.Pointer(_cret)))

	return _listView
}

// EnableRubberband returns whether rows can be selected by dragging with the
// mouse.
//
// The function returns the following values:
//
//    - ok: TRUE if rubberband selection is enabled.
//
func (self *ListView) EnableRubberband() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "ListView").InvokeMethod("get_enable_rubberband", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Factory gets the factory that's currently used to populate list items.
//
// The function returns the following values:
//
//    - listItemFactory (optional): factory in use.
//
func (self *ListView) Factory() *ListItemFactory {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "ListView").InvokeMethod("get_factory", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _listItemFactory *ListItemFactory // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_listItemFactory = wrapListItemFactory(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _listItemFactory
}

// Model gets the model that's currently used to read the items displayed.
//
// The function returns the following values:
//
//    - selectionModel (optional): model in use.
//
func (self *ListView) Model() *SelectionModel {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "ListView").InvokeMethod("get_model", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _selectionModel *SelectionModel // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_selectionModel = wrapSelectionModel(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _selectionModel
}

// ShowSeparators returns whether the list box should show separators between
// rows.
//
// The function returns the following values:
//
//    - ok: TRUE if the list box shows separators.
//
func (self *ListView) ShowSeparators() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "ListView").InvokeMethod("get_show_separators", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SingleClickActivate returns whether rows will be activated on single click
// and selected on hover.
//
// The function returns the following values:
//
//    - ok: TRUE if rows are activated on single click.
//
func (self *ListView) SingleClickActivate() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "ListView").InvokeMethod("get_single_click_activate", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetEnableRubberband sets whether selections can be changed by dragging with
// the mouse.
//
// The function takes the following parameters:
//
//    - enableRubberband: TRUE to enable rubberband selection.
//
func (self *ListView) SetEnableRubberband(enableRubberband bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if enableRubberband {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "ListView").InvokeMethod("set_enable_rubberband", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(enableRubberband)
}

// SetFactory sets the GtkListItemFactory to use for populating list items.
//
// The function takes the following parameters:
//
//    - factory (optional) to use or NULL for none.
//
func (self *ListView) SetFactory(factory *ListItemFactory) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if factory != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	}

	girepository.MustFind("Gtk", "ListView").InvokeMethod("set_factory", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(factory)
}

// SetModel sets the model to use.
//
// This must be a gtk.SelectionModel to use.
//
// The function takes the following parameters:
//
//    - model (optional) to use or NULL for none.
//
func (self *ListView) SetModel(model SelectionModeller) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if model != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	girepository.MustFind("Gtk", "ListView").InvokeMethod("set_model", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}

// SetShowSeparators sets whether the list box should show separators between
// rows.
//
// The function takes the following parameters:
//
//    - showSeparators: TRUE to show separators.
//
func (self *ListView) SetShowSeparators(showSeparators bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if showSeparators {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "ListView").InvokeMethod("set_show_separators", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(showSeparators)
}

// SetSingleClickActivate sets whether rows should be activated on single click
// and selected on hover.
//
// The function takes the following parameters:
//
//    - singleClickActivate: TRUE to activate items on single click.
//
func (self *ListView) SetSingleClickActivate(singleClickActivate bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if singleClickActivate {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "ListView").InvokeMethod("set_single_click_activate", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(singleClickActivate)
}
