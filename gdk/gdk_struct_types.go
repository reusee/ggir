package gdk

/*
#include <gdk/gdk.h>
#include <gdk/gdkprivate.h>
#include <stdlib.h>
*/
import "C"

type Atom C.GdkAtom
type Color C.GdkColor
type EventAny C.GdkEventAny
type EventButton C.GdkEventButton
type EventConfigure C.GdkEventConfigure
type EventCrossing C.GdkEventCrossing
type EventDND C.GdkEventDND
type EventExpose C.GdkEventExpose
type EventFocus C.GdkEventFocus
type EventGrabBroken C.GdkEventGrabBroken
type EventKey C.GdkEventKey
type EventMotion C.GdkEventMotion
type EventOwnerChange C.GdkEventOwnerChange
type EventProperty C.GdkEventProperty
type EventProximity C.GdkEventProximity
type EventScroll C.GdkEventScroll
type EventSelection C.GdkEventSelection
type EventSequence C.GdkEventSequence
type EventSetting C.GdkEventSetting
type EventTouch C.GdkEventTouch
type EventVisibility C.GdkEventVisibility
type EventWindowState C.GdkEventWindowState
type FrameTimings C.GdkFrameTimings
type Geometry C.GdkGeometry
type KeymapKey C.GdkKeymapKey
type Point C.GdkPoint
type RGBA C.GdkRGBA
type TimeCoord C.GdkTimeCoord
type WindowAttr C.GdkWindowAttr
type WindowRedirect C.GdkWindowRedirect
