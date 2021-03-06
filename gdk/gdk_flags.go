package gdk

/*
#include <gdk/gdk.h>
#include <gdk/gdkprivate.h>
#include <glib.h>
#include <stdlib.h>
*/
import "C"

var (
	// DragAction
	ACTION_DEFAULT = C.GdkDragAction(C.GDK_ACTION_DEFAULT)
	ACTION_COPY    = C.GdkDragAction(C.GDK_ACTION_COPY)
	ACTION_MOVE    = C.GdkDragAction(C.GDK_ACTION_MOVE)
	ACTION_LINK    = C.GdkDragAction(C.GDK_ACTION_LINK)
	ACTION_PRIVATE = C.GdkDragAction(C.GDK_ACTION_PRIVATE)
	ACTION_ASK     = C.GdkDragAction(C.GDK_ACTION_ASK)

	// EventMask
	EXPOSURE_MASK            = C.GdkEventMask(C.GDK_EXPOSURE_MASK)
	POINTER_MOTION_MASK      = C.GdkEventMask(C.GDK_POINTER_MOTION_MASK)
	POINTER_MOTION_HINT_MASK = C.GdkEventMask(C.GDK_POINTER_MOTION_HINT_MASK)
	BUTTON_MOTION_MASK       = C.GdkEventMask(C.GDK_BUTTON_MOTION_MASK)
	BUTTON1_MOTION_MASK      = C.GdkEventMask(C.GDK_BUTTON1_MOTION_MASK)
	BUTTON2_MOTION_MASK      = C.GdkEventMask(C.GDK_BUTTON2_MOTION_MASK)
	BUTTON3_MOTION_MASK      = C.GdkEventMask(C.GDK_BUTTON3_MOTION_MASK)
	BUTTON_PRESS_MASK        = C.GdkEventMask(C.GDK_BUTTON_PRESS_MASK)
	BUTTON_RELEASE_MASK      = C.GdkEventMask(C.GDK_BUTTON_RELEASE_MASK)
	KEY_PRESS_MASK           = C.GdkEventMask(C.GDK_KEY_PRESS_MASK)
	KEY_RELEASE_MASK         = C.GdkEventMask(C.GDK_KEY_RELEASE_MASK)
	ENTER_NOTIFY_MASK        = C.GdkEventMask(C.GDK_ENTER_NOTIFY_MASK)
	LEAVE_NOTIFY_MASK        = C.GdkEventMask(C.GDK_LEAVE_NOTIFY_MASK)
	FOCUS_CHANGE_MASK        = C.GdkEventMask(C.GDK_FOCUS_CHANGE_MASK)
	STRUCTURE_MASK           = C.GdkEventMask(C.GDK_STRUCTURE_MASK)
	PROPERTY_CHANGE_MASK     = C.GdkEventMask(C.GDK_PROPERTY_CHANGE_MASK)
	VISIBILITY_NOTIFY_MASK   = C.GdkEventMask(C.GDK_VISIBILITY_NOTIFY_MASK)
	PROXIMITY_IN_MASK        = C.GdkEventMask(C.GDK_PROXIMITY_IN_MASK)
	PROXIMITY_OUT_MASK       = C.GdkEventMask(C.GDK_PROXIMITY_OUT_MASK)
	SUBSTRUCTURE_MASK        = C.GdkEventMask(C.GDK_SUBSTRUCTURE_MASK)
	SCROLL_MASK              = C.GdkEventMask(C.GDK_SCROLL_MASK)
	TOUCH_MASK               = C.GdkEventMask(C.GDK_TOUCH_MASK)
	SMOOTH_SCROLL_MASK       = C.GdkEventMask(C.GDK_SMOOTH_SCROLL_MASK)
	ALL_EVENTS_MASK          = C.GdkEventMask(C.GDK_ALL_EVENTS_MASK)

	// FrameClockPhase
	FRAME_CLOCK_PHASE_NONE          = C.GdkFrameClockPhase(C.GDK_FRAME_CLOCK_PHASE_NONE)
	FRAME_CLOCK_PHASE_FLUSH_EVENTS  = C.GdkFrameClockPhase(C.GDK_FRAME_CLOCK_PHASE_FLUSH_EVENTS)
	FRAME_CLOCK_PHASE_BEFORE_PAINT  = C.GdkFrameClockPhase(C.GDK_FRAME_CLOCK_PHASE_BEFORE_PAINT)
	FRAME_CLOCK_PHASE_UPDATE        = C.GdkFrameClockPhase(C.GDK_FRAME_CLOCK_PHASE_UPDATE)
	FRAME_CLOCK_PHASE_LAYOUT        = C.GdkFrameClockPhase(C.GDK_FRAME_CLOCK_PHASE_LAYOUT)
	FRAME_CLOCK_PHASE_PAINT         = C.GdkFrameClockPhase(C.GDK_FRAME_CLOCK_PHASE_PAINT)
	FRAME_CLOCK_PHASE_RESUME_EVENTS = C.GdkFrameClockPhase(C.GDK_FRAME_CLOCK_PHASE_RESUME_EVENTS)
	FRAME_CLOCK_PHASE_AFTER_PAINT   = C.GdkFrameClockPhase(C.GDK_FRAME_CLOCK_PHASE_AFTER_PAINT)

	// ModifierType
	SHIFT_MASK                = C.GdkModifierType(C.GDK_SHIFT_MASK)
	LOCK_MASK                 = C.GdkModifierType(C.GDK_LOCK_MASK)
	CONTROL_MASK              = C.GdkModifierType(C.GDK_CONTROL_MASK)
	MOD1_MASK                 = C.GdkModifierType(C.GDK_MOD1_MASK)
	MOD2_MASK                 = C.GdkModifierType(C.GDK_MOD2_MASK)
	MOD3_MASK                 = C.GdkModifierType(C.GDK_MOD3_MASK)
	MOD4_MASK                 = C.GdkModifierType(C.GDK_MOD4_MASK)
	MOD5_MASK                 = C.GdkModifierType(C.GDK_MOD5_MASK)
	BUTTON1_MASK              = C.GdkModifierType(C.GDK_BUTTON1_MASK)
	BUTTON2_MASK              = C.GdkModifierType(C.GDK_BUTTON2_MASK)
	BUTTON3_MASK              = C.GdkModifierType(C.GDK_BUTTON3_MASK)
	BUTTON4_MASK              = C.GdkModifierType(C.GDK_BUTTON4_MASK)
	BUTTON5_MASK              = C.GdkModifierType(C.GDK_BUTTON5_MASK)
	MODIFIER_RESERVED_13_MASK = C.GdkModifierType(C.GDK_MODIFIER_RESERVED_13_MASK)
	MODIFIER_RESERVED_14_MASK = C.GdkModifierType(C.GDK_MODIFIER_RESERVED_14_MASK)
	MODIFIER_RESERVED_15_MASK = C.GdkModifierType(C.GDK_MODIFIER_RESERVED_15_MASK)
	MODIFIER_RESERVED_16_MASK = C.GdkModifierType(C.GDK_MODIFIER_RESERVED_16_MASK)
	MODIFIER_RESERVED_17_MASK = C.GdkModifierType(C.GDK_MODIFIER_RESERVED_17_MASK)
	MODIFIER_RESERVED_18_MASK = C.GdkModifierType(C.GDK_MODIFIER_RESERVED_18_MASK)
	MODIFIER_RESERVED_19_MASK = C.GdkModifierType(C.GDK_MODIFIER_RESERVED_19_MASK)
	MODIFIER_RESERVED_20_MASK = C.GdkModifierType(C.GDK_MODIFIER_RESERVED_20_MASK)
	MODIFIER_RESERVED_21_MASK = C.GdkModifierType(C.GDK_MODIFIER_RESERVED_21_MASK)
	MODIFIER_RESERVED_22_MASK = C.GdkModifierType(C.GDK_MODIFIER_RESERVED_22_MASK)
	MODIFIER_RESERVED_23_MASK = C.GdkModifierType(C.GDK_MODIFIER_RESERVED_23_MASK)
	MODIFIER_RESERVED_24_MASK = C.GdkModifierType(C.GDK_MODIFIER_RESERVED_24_MASK)
	MODIFIER_RESERVED_25_MASK = C.GdkModifierType(C.GDK_MODIFIER_RESERVED_25_MASK)
	SUPER_MASK                = C.GdkModifierType(C.GDK_SUPER_MASK)
	HYPER_MASK                = C.GdkModifierType(C.GDK_HYPER_MASK)
	META_MASK                 = C.GdkModifierType(C.GDK_META_MASK)
	MODIFIER_RESERVED_29_MASK = C.GdkModifierType(C.GDK_MODIFIER_RESERVED_29_MASK)
	RELEASE_MASK              = C.GdkModifierType(C.GDK_RELEASE_MASK)
	MODIFIER_MASK             = C.GdkModifierType(C.GDK_MODIFIER_MASK)

	// WMDecoration
	DECOR_ALL      = C.GdkWMDecoration(C.GDK_DECOR_ALL)
	DECOR_BORDER   = C.GdkWMDecoration(C.GDK_DECOR_BORDER)
	DECOR_RESIZEH  = C.GdkWMDecoration(C.GDK_DECOR_RESIZEH)
	DECOR_TITLE    = C.GdkWMDecoration(C.GDK_DECOR_TITLE)
	DECOR_MENU     = C.GdkWMDecoration(C.GDK_DECOR_MENU)
	DECOR_MINIMIZE = C.GdkWMDecoration(C.GDK_DECOR_MINIMIZE)
	DECOR_MAXIMIZE = C.GdkWMDecoration(C.GDK_DECOR_MAXIMIZE)

	// WMFunction
	FUNC_ALL      = C.GdkWMFunction(C.GDK_FUNC_ALL)
	FUNC_RESIZE   = C.GdkWMFunction(C.GDK_FUNC_RESIZE)
	FUNC_MOVE     = C.GdkWMFunction(C.GDK_FUNC_MOVE)
	FUNC_MINIMIZE = C.GdkWMFunction(C.GDK_FUNC_MINIMIZE)
	FUNC_MAXIMIZE = C.GdkWMFunction(C.GDK_FUNC_MAXIMIZE)
	FUNC_CLOSE    = C.GdkWMFunction(C.GDK_FUNC_CLOSE)

	// WindowAttributesType
	WA_TITLE     = C.GdkWindowAttributesType(C.GDK_WA_TITLE)
	WA_X         = C.GdkWindowAttributesType(C.GDK_WA_X)
	WA_Y         = C.GdkWindowAttributesType(C.GDK_WA_Y)
	WA_CURSOR    = C.GdkWindowAttributesType(C.GDK_WA_CURSOR)
	WA_VISUAL    = C.GdkWindowAttributesType(C.GDK_WA_VISUAL)
	WA_WMCLASS   = C.GdkWindowAttributesType(C.GDK_WA_WMCLASS)
	WA_NOREDIR   = C.GdkWindowAttributesType(C.GDK_WA_NOREDIR)
	WA_TYPE_HINT = C.GdkWindowAttributesType(C.GDK_WA_TYPE_HINT)

	// WindowHints
	HINT_POS         = C.GdkWindowHints(C.GDK_HINT_POS)
	HINT_MIN_SIZE    = C.GdkWindowHints(C.GDK_HINT_MIN_SIZE)
	HINT_MAX_SIZE    = C.GdkWindowHints(C.GDK_HINT_MAX_SIZE)
	HINT_BASE_SIZE   = C.GdkWindowHints(C.GDK_HINT_BASE_SIZE)
	HINT_ASPECT      = C.GdkWindowHints(C.GDK_HINT_ASPECT)
	HINT_RESIZE_INC  = C.GdkWindowHints(C.GDK_HINT_RESIZE_INC)
	HINT_WIN_GRAVITY = C.GdkWindowHints(C.GDK_HINT_WIN_GRAVITY)
	HINT_USER_POS    = C.GdkWindowHints(C.GDK_HINT_USER_POS)
	HINT_USER_SIZE   = C.GdkWindowHints(C.GDK_HINT_USER_SIZE)

	// WindowState
	WINDOW_STATE_WITHDRAWN  = C.GdkWindowState(C.GDK_WINDOW_STATE_WITHDRAWN)
	WINDOW_STATE_ICONIFIED  = C.GdkWindowState(C.GDK_WINDOW_STATE_ICONIFIED)
	WINDOW_STATE_MAXIMIZED  = C.GdkWindowState(C.GDK_WINDOW_STATE_MAXIMIZED)
	WINDOW_STATE_STICKY     = C.GdkWindowState(C.GDK_WINDOW_STATE_STICKY)
	WINDOW_STATE_FULLSCREEN = C.GdkWindowState(C.GDK_WINDOW_STATE_FULLSCREEN)
	WINDOW_STATE_ABOVE      = C.GdkWindowState(C.GDK_WINDOW_STATE_ABOVE)
	WINDOW_STATE_BELOW      = C.GdkWindowState(C.GDK_WINDOW_STATE_BELOW)
	WINDOW_STATE_FOCUSED    = C.GdkWindowState(C.GDK_WINDOW_STATE_FOCUSED)
	WINDOW_STATE_TILED      = C.GdkWindowState(C.GDK_WINDOW_STATE_TILED)
)
