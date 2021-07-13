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
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_drag_source_get_type()), F: marshalDragSourcer},
	})
}

// DragSourcer describes DragSource's methods.
type DragSourcer interface {
	// DragCancel cancels a currently ongoing drag operation.
	DragCancel()
	// Actions gets the actions that are currently set on the GtkDragSource.
	Actions() gdk.DragAction
	// Content gets the current content provider of a GtkDragSource.
	Content() *gdk.ContentProvider
	// Drag returns the underlying GdkDrag object for an ongoing drag.
	Drag() *gdk.Drag
	// SetActions sets the actions on the GtkDragSource.
	SetActions(actions gdk.DragAction)
	// SetContent sets a content provider on a GtkDragSource.
	SetContent(content *gdk.ContentProvider)
	// SetIcon sets a paintable to use as icon during DND operations.
	SetIcon(paintable gdk.Paintabler, hotX int, hotY int)
}

// DragSource: GtkDragSource is an event controller to initiate Drag-And-Drop
// operations.
//
// GtkDragSource can be set up with the necessary ingredients for a DND
// operation ahead of time. This includes the source for the data that is being
// transferred, in the form of a gdk.ContentProvider, the desired action, and
// the icon to use during the drag operation. After setting it up, the drag
// source must be added to a widget as an event controller, using
// gtk.Widget.AddController().
//
//    static void
//    my_widget_init (MyWidget *self)
//    {
//      GtkDragSource *drag_source = gtk_drag_source_new ();
//
//      g_signal_connect (drag_source, "prepare", G_CALLBACK (on_drag_prepare), self);
//      g_signal_connect (drag_source, "drag-begin", G_CALLBACK (on_drag_begin), self);
//
//      gtk_widget_add_controller (GTK_WIDGET (self), GTK_EVENT_CONTROLLER (drag_source));
//    }
//
//
// Setting up the content provider and icon ahead of time only makes sense when
// the data does not change. More commonly, you will want to set them up just in
// time. To do so, GtkDragSource has gtk.DragSource::prepare and
// gtk.DragSource::drag-begin signals.
//
// The ::prepare signal is emitted before a drag is started, and can be used to
// set the content provider and actions that the drag should be started with.
//
//    static GdkContentProvider *
//    on_drag_prepare (GtkDragSource *source,
//                     double         x,
//                     double         y,
//                     MyWidget      *self)
//    {
//      // This widget supports two types of content: GFile objects
//      // and GdkPixbuf objects; GTK will handle the serialization
//      // of these types automatically
//      GFile *file = my_widget_get_file (self);
//      GdkPixbuf *pixbuf = my_widget_get_pixbuf (self);
//
//      return gdk_content_provider_new_union ((GdkContentProvider *[2]) {
//          gdk_content_provider_new_typed (G_TYPE_FILE, file),
//          gdk_content_provider_new_typed (GDK_TYPE_PIXBUF, pixbuf),
//        }, 2);
//    }
//
//
// The ::drag-begin signal is emitted after the GdkDrag object has been created,
// and can be used to set up the drag icon.
//
//    static void
//    on_drag_begin (GtkDragSource *source,
//                   GtkDrag       *drag,
//                   MyWidget      *self)
//    {
//      // Set the widget as the drag icon
//      GdkPaintable *paintable = gtk_widget_paintable_new (GTK_WIDGET (self));
//      gtk_drag_source_set_icon (source, paintable, 0, 0);
//      g_object_unref (paintable);
//    }
//
//
// During the DND operation, GtkDragSource emits signals that can be used to
// obtain updates about the status of the operation, but it is not normally
// necessary to connect to any signals, except for one case: when the supported
// actions include GDK_ACTION_MOVE, you need to listen for the
// gtk.DragSource::drag-end signal and delete the data after it has been
// transferred.
type DragSource struct {
	GestureSingle
}

var (
	_ DragSourcer     = (*DragSource)(nil)
	_ gextras.Nativer = (*DragSource)(nil)
)

func wrapDragSource(obj *externglib.Object) *DragSource {
	return &DragSource{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalDragSourcer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDragSource(obj), nil
}

// NewDragSource creates a new GtkDragSource object.
func NewDragSource() *DragSource {
	var _cret *C.GtkDragSource // in

	_cret = C.gtk_drag_source_new()

	var _dragSource *DragSource // out

	_dragSource = wrapDragSource(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dragSource
}

// DragCancel cancels a currently ongoing drag operation.
func (source *DragSource) DragCancel() {
	var _arg0 *C.GtkDragSource // out

	_arg0 = (*C.GtkDragSource)(unsafe.Pointer(source.Native()))

	C.gtk_drag_source_drag_cancel(_arg0)
}

// Actions gets the actions that are currently set on the GtkDragSource.
func (source *DragSource) Actions() gdk.DragAction {
	var _arg0 *C.GtkDragSource // out
	var _cret C.GdkDragAction  // in

	_arg0 = (*C.GtkDragSource)(unsafe.Pointer(source.Native()))

	_cret = C.gtk_drag_source_get_actions(_arg0)

	var _dragAction gdk.DragAction // out

	_dragAction = gdk.DragAction(_cret)

	return _dragAction
}

// Content gets the current content provider of a GtkDragSource.
func (source *DragSource) Content() *gdk.ContentProvider {
	var _arg0 *C.GtkDragSource      // out
	var _cret *C.GdkContentProvider // in

	_arg0 = (*C.GtkDragSource)(unsafe.Pointer(source.Native()))

	_cret = C.gtk_drag_source_get_content(_arg0)

	var _contentProvider *gdk.ContentProvider // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_contentProvider = &gdk.ContentProvider{
			Object: obj,
		}
	}

	return _contentProvider
}

// Drag returns the underlying GdkDrag object for an ongoing drag.
func (source *DragSource) Drag() *gdk.Drag {
	var _arg0 *C.GtkDragSource // out
	var _cret *C.GdkDrag       // in

	_arg0 = (*C.GtkDragSource)(unsafe.Pointer(source.Native()))

	_cret = C.gtk_drag_source_get_drag(_arg0)

	var _drag *gdk.Drag // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_drag = &gdk.Drag{
			Object: obj,
		}
	}

	return _drag
}

// SetActions sets the actions on the GtkDragSource.
//
// During a DND operation, the actions are offered to potential drop targets. If
// actions include GDK_ACTION_MOVE, you need to listen to the
// gtk.DragSource::drag-end signal and handle delete_data being TRUE.
//
// This function can be called before a drag is started, or in a handler for the
// gtk.DragSource::prepare signal.
func (source *DragSource) SetActions(actions gdk.DragAction) {
	var _arg0 *C.GtkDragSource // out
	var _arg1 C.GdkDragAction  // out

	_arg0 = (*C.GtkDragSource)(unsafe.Pointer(source.Native()))
	_arg1 = C.GdkDragAction(actions)

	C.gtk_drag_source_set_actions(_arg0, _arg1)
}

// SetContent sets a content provider on a GtkDragSource.
//
// When the data is requested in the cause of a DND operation, it will be
// obtained from the content provider.
//
// This function can be called before a drag is started, or in a handler for the
// gtk.DragSource::prepare signal.
//
// You may consider setting the content provider back to NULL in a
// gtk.DragSource::drag-end signal handler.
func (source *DragSource) SetContent(content *gdk.ContentProvider) {
	var _arg0 *C.GtkDragSource      // out
	var _arg1 *C.GdkContentProvider // out

	_arg0 = (*C.GtkDragSource)(unsafe.Pointer(source.Native()))
	_arg1 = (*C.GdkContentProvider)(unsafe.Pointer(content.Native()))

	C.gtk_drag_source_set_content(_arg0, _arg1)
}

// SetIcon sets a paintable to use as icon during DND operations.
//
// The hotspot coordinates determine the point on the icon that gets aligned
// with the hotspot of the cursor.
//
// If paintable is NULL, a default icon is used.
//
// This function can be called before a drag is started, or in a
// gtk.DragSource::prepare or gtk.DragSource::drag-begin signal handler.
func (source *DragSource) SetIcon(paintable gdk.Paintabler, hotX int, hotY int) {
	var _arg0 *C.GtkDragSource // out
	var _arg1 *C.GdkPaintable  // out
	var _arg2 C.int            // out
	var _arg3 C.int            // out

	_arg0 = (*C.GtkDragSource)(unsafe.Pointer(source.Native()))
	_arg1 = (*C.GdkPaintable)(unsafe.Pointer((paintable).(gextras.Nativer).Native()))
	_arg2 = C.int(hotX)
	_arg3 = C.int(hotY)

	C.gtk_drag_source_set_icon(_arg0, _arg1, _arg2, _arg3)
}
