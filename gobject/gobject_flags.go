package gobject

/*
#include <glib-object.h>
#include <gobject/gvaluecollector.h>
#include <stdlib.h>
*/
import "C"

var (
	// BindingFlags
	BINDING_DEFAULT        = C.GBindingFlags(C.G_BINDING_DEFAULT)
	BINDING_BIDIRECTIONAL  = C.GBindingFlags(C.G_BINDING_BIDIRECTIONAL)
	BINDING_SYNC_CREATE    = C.GBindingFlags(C.G_BINDING_SYNC_CREATE)
	BINDING_INVERT_BOOLEAN = C.GBindingFlags(C.G_BINDING_INVERT_BOOLEAN)

	// ConnectFlags
	CONNECT_AFTER   = C.GConnectFlags(C.G_CONNECT_AFTER)
	CONNECT_SWAPPED = C.GConnectFlags(C.G_CONNECT_SWAPPED)

	// ParamFlags
	PARAM_READABLE       = C.GParamFlags(C.G_PARAM_READABLE)
	PARAM_WRITABLE       = C.GParamFlags(C.G_PARAM_WRITABLE)
	PARAM_CONSTRUCT      = C.GParamFlags(C.G_PARAM_CONSTRUCT)
	PARAM_CONSTRUCT_ONLY = C.GParamFlags(C.G_PARAM_CONSTRUCT_ONLY)
	PARAM_LAX_VALIDATION = C.GParamFlags(C.G_PARAM_LAX_VALIDATION)
	PARAM_STATIC_NAME    = C.GParamFlags(C.G_PARAM_STATIC_NAME)
	PARAM_PRIVATE        = C.GParamFlags(C.G_PARAM_PRIVATE)
	PARAM_STATIC_NICK    = C.GParamFlags(C.G_PARAM_STATIC_NICK)
	PARAM_STATIC_BLURB   = C.GParamFlags(C.G_PARAM_STATIC_BLURB)
	PARAM_DEPRECATED     = C.GParamFlags(C.G_PARAM_DEPRECATED)

	// SignalFlags
	SIGNAL_RUN_FIRST    = C.GSignalFlags(C.G_SIGNAL_RUN_FIRST)
	SIGNAL_RUN_LAST     = C.GSignalFlags(C.G_SIGNAL_RUN_LAST)
	SIGNAL_RUN_CLEANUP  = C.GSignalFlags(C.G_SIGNAL_RUN_CLEANUP)
	SIGNAL_NO_RECURSE   = C.GSignalFlags(C.G_SIGNAL_NO_RECURSE)
	SIGNAL_DETAILED     = C.GSignalFlags(C.G_SIGNAL_DETAILED)
	SIGNAL_ACTION       = C.GSignalFlags(C.G_SIGNAL_ACTION)
	SIGNAL_NO_HOOKS     = C.GSignalFlags(C.G_SIGNAL_NO_HOOKS)
	SIGNAL_MUST_COLLECT = C.GSignalFlags(C.G_SIGNAL_MUST_COLLECT)
	SIGNAL_DEPRECATED   = C.GSignalFlags(C.G_SIGNAL_DEPRECATED)

	// SignalMatchType
	SIGNAL_MATCH_ID        = C.GSignalMatchType(C.G_SIGNAL_MATCH_ID)
	SIGNAL_MATCH_DETAIL    = C.GSignalMatchType(C.G_SIGNAL_MATCH_DETAIL)
	SIGNAL_MATCH_CLOSURE   = C.GSignalMatchType(C.G_SIGNAL_MATCH_CLOSURE)
	SIGNAL_MATCH_FUNC      = C.GSignalMatchType(C.G_SIGNAL_MATCH_FUNC)
	SIGNAL_MATCH_DATA      = C.GSignalMatchType(C.G_SIGNAL_MATCH_DATA)
	SIGNAL_MATCH_UNBLOCKED = C.GSignalMatchType(C.G_SIGNAL_MATCH_UNBLOCKED)

	// TypeDebugFlags
	TYPE_DEBUG_NONE    = C.GTypeDebugFlags(C.G_TYPE_DEBUG_NONE)
	TYPE_DEBUG_OBJECTS = C.GTypeDebugFlags(C.G_TYPE_DEBUG_OBJECTS)
	TYPE_DEBUG_SIGNALS = C.GTypeDebugFlags(C.G_TYPE_DEBUG_SIGNALS)
	TYPE_DEBUG_MASK    = C.GTypeDebugFlags(C.G_TYPE_DEBUG_MASK)

	// TypeFlags
	TYPE_FLAG_ABSTRACT       = C.GTypeFlags(C.G_TYPE_FLAG_ABSTRACT)
	TYPE_FLAG_VALUE_ABSTRACT = C.GTypeFlags(C.G_TYPE_FLAG_VALUE_ABSTRACT)

	// TypeFundamentalFlags
	TYPE_FLAG_CLASSED        = C.GTypeFundamentalFlags(C.G_TYPE_FLAG_CLASSED)
	TYPE_FLAG_INSTANTIATABLE = C.GTypeFundamentalFlags(C.G_TYPE_FLAG_INSTANTIATABLE)
	TYPE_FLAG_DERIVABLE      = C.GTypeFundamentalFlags(C.G_TYPE_FLAG_DERIVABLE)
	TYPE_FLAG_DEEP_DERIVABLE = C.GTypeFundamentalFlags(C.G_TYPE_FLAG_DEEP_DERIVABLE)
)