// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// PRIORITY_RESIZE: use this priority for functionality related to size
// allocation.
//
// It is used internally by GTK+ to compute the sizes of widgets. This priority
// is higher than GDK_PRIORITY_REDRAW to avoid resizing a widget which was just
// redrawn.
const PRIORITY_RESIZE = 110

// KeySnoopFunc: key snooper functions are called before normal event delivery.
// They can be used to implement custom key event handling.
type KeySnoopFunc func(grabWidget Widgetter, event *gdk.EventKey) (gint int)

// CheckVersion checks that the GTK+ library in use is compatible with the given
// version. Generally you would pass in the constants K_MAJOR_VERSION,
// K_MINOR_VERSION, K_MICRO_VERSION as the three arguments to this function;
// that produces a check that the library in use is compatible with the version
// of GTK+ the application or module was compiled against.
//
// Compatibility is defined by two things: first the version of the running
// library is newer than the version
// required_major.required_minor.required_micro. Second the running library must
// be binary compatible with the version
// required_major.required_minor.required_micro (same major version.)
//
// This function is primarily for GTK+ modules; the module can call this
// function to check that it wasn’t loaded into an incompatible version of GTK+.
// However, such a check isn’t completely reliable, since the module may be
// linked against an old version of GTK+ and calling the old version of
// gtk_check_version(), but still get loaded into an application using a newer
// version of GTK+.
//
// The function takes the following parameters:
//
//    - requiredMajor: required major version.
//    - requiredMinor: required minor version.
//    - requiredMicro: required micro version.
//
// The function returns the following values:
//
//    - utf8 (optional): NULL if the GTK+ library is compatible with the given
//      version, or a string describing the version mismatch. The returned string
//      is owned by GTK+ and should not be modified or freed.
//
func CheckVersion(requiredMajor, requiredMinor, requiredMicro uint) string {
	var _arg1 C.guint  // out
	var _arg2 C.guint  // out
	var _arg3 C.guint  // out
	var _cret *C.gchar // in

	_arg1 = C.guint(requiredMajor)
	_arg2 = C.guint(requiredMinor)
	_arg3 = C.guint(requiredMicro)

	_cret = C.gtk_check_version(_arg1, _arg2, _arg3)
	runtime.KeepAlive(requiredMajor)
	runtime.KeepAlive(requiredMinor)
	runtime.KeepAlive(requiredMicro)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// DisableSetlocale prevents gtk_init(), gtk_init_check(), gtk_init_with_args()
// and gtk_parse_args() from automatically calling setlocale (LC_ALL, ""). You
// would want to use this function if you wanted to set the locale for your
// program to something other than the user’s locale, or if you wanted to set
// different values for different locale categories.
//
// Most programs should not need to call this function.
func DisableSetlocale() {
	C.gtk_disable_setlocale()
}

// EventsPending checks if any events are pending.
//
// This can be used to update the UI and invoke timeouts etc. while doing some
// time intensive computation.
//
// Updating the UI during a long computation
//
//     // computation going on...
//
//     while (gtk_events_pending ())
//       gtk_main_iteration ();
//
//     // ...computation continued.
//
// The function returns the following values:
//
//    - ok: TRUE if any events are pending, FALSE otherwise.
//
func EventsPending() bool {
	var _cret C.gboolean // in

	_cret = C.gtk_events_pending()

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// False: analogical to gtk_true(), this function does nothing but always
// returns FALSE.
//
// The function returns the following values:
//
//    - ok: FALSE.
//
func False() bool {
	var _cret C.gboolean // in

	_cret = C.gtk_false()

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GetCurrentEvent obtains a copy of the event currently being processed by
// GTK+.
//
// For example, if you are handling a Button::clicked signal, the current event
// will be the EventButton that triggered the ::clicked signal.
//
// The function returns the following values:
//
//    - event (optional): copy of the current event, or NULL if there is no
//      current event. The returned event must be freed with gdk_event_free().
//
func GetCurrentEvent() *gdk.Event {
	var _cret *C.GdkEvent // in

	_cret = C.gtk_get_current_event()

	var _event *gdk.Event // out

	if _cret != nil {
		{
			v := (*gdk.Event)(gextras.NewStructNative(unsafe.Pointer(_cret)))
			_event = v
		}
	}

	return _event
}

// GetCurrentEventDevice: if there is a current event and it has a device,
// return that device, otherwise return NULL.
//
// The function returns the following values:
//
//    - device (optional) or NULL.
//
func GetCurrentEventDevice() gdk.Devicer {
	var _cret *C.GdkDevice // in

	_cret = C.gtk_get_current_event_device()

	var _device gdk.Devicer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gdk.Devicer)
				return ok
			})
			rv, ok := casted.(gdk.Devicer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
			}
			_device = rv
		}
	}

	return _device
}

// GetCurrentEventState: if there is a current event and it has a state field,
// place that state field in state and return TRUE, otherwise return FALSE.
//
// The function returns the following values:
//
//    - state: location to store the state of the current event.
//    - ok: TRUE if there was a current event and it had a state field.
//
func GetCurrentEventState() (gdk.ModifierType, bool) {
	var _arg1 C.GdkModifierType // in
	var _cret C.gboolean        // in

	_cret = C.gtk_get_current_event_state(&_arg1)

	var _state gdk.ModifierType // out
	var _ok bool                // out

	_state = gdk.ModifierType(_arg1)
	if _cret != 0 {
		_ok = true
	}

	return _state, _ok
}

// GetCurrentEventTime: if there is a current event and it has a timestamp,
// return that timestamp, otherwise return GDK_CURRENT_TIME.
//
// The function returns the following values:
//
//    - guint32: timestamp from the current event, or GDK_CURRENT_TIME.
//
func GetCurrentEventTime() uint32 {
	var _cret C.guint32 // in

	_cret = C.gtk_get_current_event_time()

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// GetDefaultLanguage returns the Language for the default language currently in
// effect. (Note that this can change over the life of an application.) The
// default language is derived from the current locale. It determines, for
// example, whether GTK+ uses the right-to-left or left-to-right text direction.
//
// This function is equivalent to pango_language_get_default(). See that
// function for details.
//
// The function returns the following values:
//
//    - language: default language as a Language, must not be freed.
//
func GetDefaultLanguage() *pango.Language {
	var _cret *C.PangoLanguage // in

	_cret = C.gtk_get_default_language()

	var _language *pango.Language // out

	_language = (*pango.Language)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _language
}

// GetEventWidget: if event is NULL or the event was not associated with any
// widget, returns NULL, otherwise returns the widget that received the event
// originally.
//
// The function takes the following parameters:
//
//    - event: Event.
//
// The function returns the following values:
//
//    - widget (optional) that originally received event, or NULL.
//
func GetEventWidget(event *gdk.Event) Widgetter {
	var _arg1 *C.GdkEvent  // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkEvent)(gextras.StructNative(unsafe.Pointer(event)))

	_cret = C.gtk_get_event_widget(_arg1)
	runtime.KeepAlive(event)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// GrabGetCurrent queries the current grab of the default window group.
//
// The function returns the following values:
//
//    - widget (optional) which currently has the grab or NULL if no grab is
//      active.
//
func GrabGetCurrent() Widgetter {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_grab_get_current()

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// KeySnooperRemove removes the key snooper function with the given id.
//
// Deprecated: Key snooping should not be done. Events should be handled by
// widgets.
//
// The function takes the following parameters:
//
//    - snooperHandlerId identifies the key snooper to remove.
//
func KeySnooperRemove(snooperHandlerId uint) {
	var _arg1 C.guint // out

	_arg1 = C.guint(snooperHandlerId)

	C.gtk_key_snooper_remove(_arg1)
	runtime.KeepAlive(snooperHandlerId)
}

// Main runs the main loop until gtk_main_quit() is called.
//
// You can nest calls to gtk_main(). In that case gtk_main_quit() will make the
// innermost invocation of the main loop return.
func Main() {
	C.gtk_main()
}

// MainDoEvent processes a single GDK event.
//
// This is public only to allow filtering of events between GDK and GTK+. You
// will not usually need to call this function directly.
//
// While you should not call this function directly, you might want to know how
// exactly events are handled. So here is what this function does with the
// event:
//
// 1. Compress enter/leave notify events. If the event passed build an
// enter/leave pair together with the next event (peeked from GDK), both events
// are thrown away. This is to avoid a backlog of (de-)highlighting widgets
// crossed by the pointer.
//
// 2. Find the widget which got the event. If the widget can’t be determined the
// event is thrown away unless it belongs to a INCR transaction.
//
// 3. Then the event is pushed onto a stack so you can query the currently
// handled event with gtk_get_current_event().
//
// 4. The event is sent to a widget. If a grab is active all events for widgets
// that are not in the contained in the grab widget are sent to the latter with
// a few exceptions: - Deletion and destruction events are still sent to the
// event widget for obvious reasons. - Events which directly relate to the
// visual representation of the event widget. - Leave events are delivered to
// the event widget if there was an enter event delivered to it before without
// the paired leave event. - Drag events are not redirected because it is
// unclear what the semantics of that would be. Another point of interest might
// be that all key events are first passed through the key snooper functions if
// there are any. Read the description of gtk_key_snooper_install() if you need
// this feature.
//
// 5. After finishing the delivery the event is popped from the event stack.
//
// The function takes the following parameters:
//
//    - event to process (normally passed by GDK).
//
func MainDoEvent(event *gdk.Event) {
	var _arg1 *C.GdkEvent // out

	_arg1 = (*C.GdkEvent)(gextras.StructNative(unsafe.Pointer(event)))

	C.gtk_main_do_event(_arg1)
	runtime.KeepAlive(event)
}

// MainIteration runs a single iteration of the mainloop.
//
// If no events are waiting to be processed GTK+ will block until the next event
// is noticed. If you don’t want to block look at gtk_main_iteration_do() or
// check if any events are pending with gtk_events_pending() first.
//
// The function returns the following values:
//
//    - ok: TRUE if gtk_main_quit() has been called for the innermost mainloop.
//
func MainIteration() bool {
	var _cret C.gboolean // in

	_cret = C.gtk_main_iteration()

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MainIterationDo runs a single iteration of the mainloop. If no events are
// available either return or block depending on the value of blocking.
//
// The function takes the following parameters:
//
//    - blocking: TRUE if you want GTK+ to block if no events are pending.
//
// The function returns the following values:
//
//    - ok: TRUE if gtk_main_quit() has been called for the innermost mainloop.
//
func MainIterationDo(blocking bool) bool {
	var _arg1 C.gboolean // out
	var _cret C.gboolean // in

	if blocking {
		_arg1 = C.TRUE
	}

	_cret = C.gtk_main_iteration_do(_arg1)
	runtime.KeepAlive(blocking)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MainLevel asks for the current nesting level of the main loop.
//
// The function returns the following values:
//
//    - guint: nesting level of the current invocation of the main loop.
//
func MainLevel() uint {
	var _cret C.guint // in

	_cret = C.gtk_main_level()

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// MainQuit makes the innermost invocation of the main loop return when it
// regains control.
func MainQuit() {
	C.gtk_main_quit()
}

// PropagateEvent sends an event to a widget, propagating the event to parent
// widgets if the event remains unhandled.
//
// Events received by GTK+ from GDK normally begin in gtk_main_do_event().
// Depending on the type of event, existence of modal dialogs, grabs, etc., the
// event may be propagated; if so, this function is used.
//
// gtk_propagate_event() calls gtk_widget_event() on each widget it decides to
// send the event to. So gtk_widget_event() is the lowest-level function; it
// simply emits the Widget::event and possibly an event-specific signal on a
// widget. gtk_propagate_event() is a bit higher-level, and gtk_main_do_event()
// is the highest level.
//
// All that said, you most likely don’t want to use any of these functions;
// synthesizing events is rarely needed. There are almost certainly better ways
// to achieve your goals. For example, use gdk_window_invalidate_rect() or
// gtk_widget_queue_draw() instead of making up expose events.
//
// The function takes the following parameters:
//
//    - widget: Widget.
//    - event: event.
//
func PropagateEvent(widget Widgetter, event *gdk.Event) {
	var _arg1 *C.GtkWidget // out
	var _arg2 *C.GdkEvent  // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg2 = (*C.GdkEvent)(gextras.StructNative(unsafe.Pointer(event)))

	C.gtk_propagate_event(_arg1, _arg2)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(event)
}

// True: all this function does it to return TRUE.
//
// This can be useful for example if you want to inhibit the deletion of a
// window. Of course you should not do this as the user expects a reaction from
// clicking the close icon of the window...
//
// A persistent window
//
//    #include <gtk/gtk.h>
//
//    int
//    main (int argc, char **argv)
//    {
//      GtkWidget *win, *but;
//      const char *text = "Close yourself. I mean it!";
//
//      gtk_init (&argc, &argv);
//
//      win = gtk_window_new (GTK_WINDOW_TOPLEVEL);
//      g_signal_connect (win,
//                        "delete-event",
//                        G_CALLBACK (gtk_true),
//                        NULL);
//      g_signal_connect (win, "destroy",
//                        G_CALLBACK (gtk_main_quit),
//                        NULL);
//
//      but = gtk_button_new_with_label (text);
//      g_signal_connect_swapped (but, "clicked",
//                                G_CALLBACK (gtk_object_destroy),
//                                win);
//      gtk_container_add (GTK_CONTAINER (win), but);
//
//      gtk_widget_show_all (win);
//
//      gtk_main ();
//
//      return 0;
//    }.
//
// The function returns the following values:
//
//    - ok: TRUE.
//
func True() bool {
	var _cret C.gboolean // in

	_cret = C.gtk_true()

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
