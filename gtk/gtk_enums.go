package gtk

/*
#include <gtk/gtk.h>
#include <stdlib.h>
*/
import "C"

var (
	// Align
	ALIGN_FILL     = C.GTK_ALIGN_FILL
	ALIGN_START    = C.GTK_ALIGN_START
	ALIGN_END      = C.GTK_ALIGN_END
	ALIGN_CENTER   = C.GTK_ALIGN_CENTER
	ALIGN_BASELINE = C.GTK_ALIGN_BASELINE

	// ArrowPlacement
	ARROWS_BOTH  = C.GTK_ARROWS_BOTH
	ARROWS_START = C.GTK_ARROWS_START
	ARROWS_END   = C.GTK_ARROWS_END

	// ArrowType
	ARROW_UP    = C.GTK_ARROW_UP
	ARROW_DOWN  = C.GTK_ARROW_DOWN
	ARROW_LEFT  = C.GTK_ARROW_LEFT
	ARROW_RIGHT = C.GTK_ARROW_RIGHT
	ARROW_NONE  = C.GTK_ARROW_NONE

	// AssistantPageType
	ASSISTANT_PAGE_CONTENT  = C.GTK_ASSISTANT_PAGE_CONTENT
	ASSISTANT_PAGE_INTRO    = C.GTK_ASSISTANT_PAGE_INTRO
	ASSISTANT_PAGE_CONFIRM  = C.GTK_ASSISTANT_PAGE_CONFIRM
	ASSISTANT_PAGE_SUMMARY  = C.GTK_ASSISTANT_PAGE_SUMMARY
	ASSISTANT_PAGE_PROGRESS = C.GTK_ASSISTANT_PAGE_PROGRESS
	ASSISTANT_PAGE_CUSTOM   = C.GTK_ASSISTANT_PAGE_CUSTOM

	// BaselinePosition
	BASELINE_POSITION_TOP    = C.GTK_BASELINE_POSITION_TOP
	BASELINE_POSITION_CENTER = C.GTK_BASELINE_POSITION_CENTER
	BASELINE_POSITION_BOTTOM = C.GTK_BASELINE_POSITION_BOTTOM

	// BorderStyle
	BORDER_STYLE_NONE   = C.GTK_BORDER_STYLE_NONE
	BORDER_STYLE_SOLID  = C.GTK_BORDER_STYLE_SOLID
	BORDER_STYLE_INSET  = C.GTK_BORDER_STYLE_INSET
	BORDER_STYLE_OUTSET = C.GTK_BORDER_STYLE_OUTSET
	BORDER_STYLE_HIDDEN = C.GTK_BORDER_STYLE_HIDDEN
	BORDER_STYLE_DOTTED = C.GTK_BORDER_STYLE_DOTTED
	BORDER_STYLE_DASHED = C.GTK_BORDER_STYLE_DASHED
	BORDER_STYLE_DOUBLE = C.GTK_BORDER_STYLE_DOUBLE
	BORDER_STYLE_GROOVE = C.GTK_BORDER_STYLE_GROOVE
	BORDER_STYLE_RIDGE  = C.GTK_BORDER_STYLE_RIDGE

	// BuilderError
	BUILDER_ERROR_INVALID_TYPE_FUNCTION  = C.GTK_BUILDER_ERROR_INVALID_TYPE_FUNCTION
	BUILDER_ERROR_UNHANDLED_TAG          = C.GTK_BUILDER_ERROR_UNHANDLED_TAG
	BUILDER_ERROR_MISSING_ATTRIBUTE      = C.GTK_BUILDER_ERROR_MISSING_ATTRIBUTE
	BUILDER_ERROR_INVALID_ATTRIBUTE      = C.GTK_BUILDER_ERROR_INVALID_ATTRIBUTE
	BUILDER_ERROR_INVALID_TAG            = C.GTK_BUILDER_ERROR_INVALID_TAG
	BUILDER_ERROR_MISSING_PROPERTY_VALUE = C.GTK_BUILDER_ERROR_MISSING_PROPERTY_VALUE
	BUILDER_ERROR_INVALID_VALUE          = C.GTK_BUILDER_ERROR_INVALID_VALUE
	BUILDER_ERROR_VERSION_MISMATCH       = C.GTK_BUILDER_ERROR_VERSION_MISMATCH
	BUILDER_ERROR_DUPLICATE_ID           = C.GTK_BUILDER_ERROR_DUPLICATE_ID
	BUILDER_ERROR_OBJECT_TYPE_REFUSED    = C.GTK_BUILDER_ERROR_OBJECT_TYPE_REFUSED
	BUILDER_ERROR_TEMPLATE_MISMATCH      = C.GTK_BUILDER_ERROR_TEMPLATE_MISMATCH

	// ButtonBoxStyle
	BUTTONBOX_SPREAD = C.GTK_BUTTONBOX_SPREAD
	BUTTONBOX_EDGE   = C.GTK_BUTTONBOX_EDGE
	BUTTONBOX_START  = C.GTK_BUTTONBOX_START
	BUTTONBOX_END    = C.GTK_BUTTONBOX_END
	BUTTONBOX_CENTER = C.GTK_BUTTONBOX_CENTER
	BUTTONBOX_EXPAND = C.GTK_BUTTONBOX_EXPAND

	// ButtonsType
	BUTTONS_NONE      = C.GTK_BUTTONS_NONE
	BUTTONS_OK        = C.GTK_BUTTONS_OK
	BUTTONS_CLOSE     = C.GTK_BUTTONS_CLOSE
	BUTTONS_CANCEL    = C.GTK_BUTTONS_CANCEL
	BUTTONS_YES_NO    = C.GTK_BUTTONS_YES_NO
	BUTTONS_OK_CANCEL = C.GTK_BUTTONS_OK_CANCEL

	// CellRendererAccelMode
	CELL_RENDERER_ACCEL_MODE_GTK   = C.GTK_CELL_RENDERER_ACCEL_MODE_GTK
	CELL_RENDERER_ACCEL_MODE_OTHER = C.GTK_CELL_RENDERER_ACCEL_MODE_OTHER

	// CellRendererMode
	CELL_RENDERER_MODE_INERT       = C.GTK_CELL_RENDERER_MODE_INERT
	CELL_RENDERER_MODE_ACTIVATABLE = C.GTK_CELL_RENDERER_MODE_ACTIVATABLE
	CELL_RENDERER_MODE_EDITABLE    = C.GTK_CELL_RENDERER_MODE_EDITABLE

	// CornerType
	CORNER_TOP_LEFT     = C.GTK_CORNER_TOP_LEFT
	CORNER_BOTTOM_LEFT  = C.GTK_CORNER_BOTTOM_LEFT
	CORNER_TOP_RIGHT    = C.GTK_CORNER_TOP_RIGHT
	CORNER_BOTTOM_RIGHT = C.GTK_CORNER_BOTTOM_RIGHT

	// CssProviderError
	CSS_PROVIDER_ERROR_FAILED        = C.GTK_CSS_PROVIDER_ERROR_FAILED
	CSS_PROVIDER_ERROR_SYNTAX        = C.GTK_CSS_PROVIDER_ERROR_SYNTAX
	CSS_PROVIDER_ERROR_IMPORT        = C.GTK_CSS_PROVIDER_ERROR_IMPORT
	CSS_PROVIDER_ERROR_NAME          = C.GTK_CSS_PROVIDER_ERROR_NAME
	CSS_PROVIDER_ERROR_DEPRECATED    = C.GTK_CSS_PROVIDER_ERROR_DEPRECATED
	CSS_PROVIDER_ERROR_UNKNOWN_VALUE = C.GTK_CSS_PROVIDER_ERROR_UNKNOWN_VALUE

	// CssSectionType
	CSS_SECTION_DOCUMENT         = C.GTK_CSS_SECTION_DOCUMENT
	CSS_SECTION_IMPORT           = C.GTK_CSS_SECTION_IMPORT
	CSS_SECTION_COLOR_DEFINITION = C.GTK_CSS_SECTION_COLOR_DEFINITION
	CSS_SECTION_BINDING_SET      = C.GTK_CSS_SECTION_BINDING_SET
	CSS_SECTION_RULESET          = C.GTK_CSS_SECTION_RULESET
	CSS_SECTION_SELECTOR         = C.GTK_CSS_SECTION_SELECTOR
	CSS_SECTION_DECLARATION      = C.GTK_CSS_SECTION_DECLARATION
	CSS_SECTION_VALUE            = C.GTK_CSS_SECTION_VALUE
	CSS_SECTION_KEYFRAMES        = C.GTK_CSS_SECTION_KEYFRAMES

	// DeleteType
	DELETE_CHARS             = C.GTK_DELETE_CHARS
	DELETE_WORD_ENDS         = C.GTK_DELETE_WORD_ENDS
	DELETE_WORDS             = C.GTK_DELETE_WORDS
	DELETE_DISPLAY_LINES     = C.GTK_DELETE_DISPLAY_LINES
	DELETE_DISPLAY_LINE_ENDS = C.GTK_DELETE_DISPLAY_LINE_ENDS
	DELETE_PARAGRAPH_ENDS    = C.GTK_DELETE_PARAGRAPH_ENDS
	DELETE_PARAGRAPHS        = C.GTK_DELETE_PARAGRAPHS
	DELETE_WHITESPACE        = C.GTK_DELETE_WHITESPACE

	// DirectionType
	DIR_TAB_FORWARD  = C.GTK_DIR_TAB_FORWARD
	DIR_TAB_BACKWARD = C.GTK_DIR_TAB_BACKWARD
	DIR_UP           = C.GTK_DIR_UP
	DIR_DOWN         = C.GTK_DIR_DOWN
	DIR_LEFT         = C.GTK_DIR_LEFT
	DIR_RIGHT        = C.GTK_DIR_RIGHT

	// DragResult
	DRAG_RESULT_SUCCESS         = C.GTK_DRAG_RESULT_SUCCESS
	DRAG_RESULT_NO_TARGET       = C.GTK_DRAG_RESULT_NO_TARGET
	DRAG_RESULT_USER_CANCELLED  = C.GTK_DRAG_RESULT_USER_CANCELLED
	DRAG_RESULT_TIMEOUT_EXPIRED = C.GTK_DRAG_RESULT_TIMEOUT_EXPIRED
	DRAG_RESULT_GRAB_BROKEN     = C.GTK_DRAG_RESULT_GRAB_BROKEN
	DRAG_RESULT_ERROR           = C.GTK_DRAG_RESULT_ERROR

	// EntryIconPosition
	ENTRY_ICON_PRIMARY   = C.GTK_ENTRY_ICON_PRIMARY
	ENTRY_ICON_SECONDARY = C.GTK_ENTRY_ICON_SECONDARY

	// ExpanderStyle
	EXPANDER_COLLAPSED      = C.GTK_EXPANDER_COLLAPSED
	EXPANDER_SEMI_COLLAPSED = C.GTK_EXPANDER_SEMI_COLLAPSED
	EXPANDER_SEMI_EXPANDED  = C.GTK_EXPANDER_SEMI_EXPANDED
	EXPANDER_EXPANDED       = C.GTK_EXPANDER_EXPANDED

	// FileChooserAction
	FILE_CHOOSER_ACTION_OPEN          = C.GTK_FILE_CHOOSER_ACTION_OPEN
	FILE_CHOOSER_ACTION_SAVE          = C.GTK_FILE_CHOOSER_ACTION_SAVE
	FILE_CHOOSER_ACTION_SELECT_FOLDER = C.GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER
	FILE_CHOOSER_ACTION_CREATE_FOLDER = C.GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER

	// FileChooserConfirmation
	FILE_CHOOSER_CONFIRMATION_CONFIRM         = C.GTK_FILE_CHOOSER_CONFIRMATION_CONFIRM
	FILE_CHOOSER_CONFIRMATION_ACCEPT_FILENAME = C.GTK_FILE_CHOOSER_CONFIRMATION_ACCEPT_FILENAME
	FILE_CHOOSER_CONFIRMATION_SELECT_AGAIN    = C.GTK_FILE_CHOOSER_CONFIRMATION_SELECT_AGAIN

	// FileChooserError
	FILE_CHOOSER_ERROR_NONEXISTENT         = C.GTK_FILE_CHOOSER_ERROR_NONEXISTENT
	FILE_CHOOSER_ERROR_BAD_FILENAME        = C.GTK_FILE_CHOOSER_ERROR_BAD_FILENAME
	FILE_CHOOSER_ERROR_ALREADY_EXISTS      = C.GTK_FILE_CHOOSER_ERROR_ALREADY_EXISTS
	FILE_CHOOSER_ERROR_INCOMPLETE_HOSTNAME = C.GTK_FILE_CHOOSER_ERROR_INCOMPLETE_HOSTNAME

	// IMPreeditStyle
	IM_PREEDIT_NOTHING  = C.GTK_IM_PREEDIT_NOTHING
	IM_PREEDIT_CALLBACK = C.GTK_IM_PREEDIT_CALLBACK
	IM_PREEDIT_NONE     = C.GTK_IM_PREEDIT_NONE

	// IMStatusStyle
	IM_STATUS_NOTHING  = C.GTK_IM_STATUS_NOTHING
	IM_STATUS_CALLBACK = C.GTK_IM_STATUS_CALLBACK
	IM_STATUS_NONE     = C.GTK_IM_STATUS_NONE

	// IconSize
	ICON_SIZE_INVALID       = C.GTK_ICON_SIZE_INVALID
	ICON_SIZE_MENU          = C.GTK_ICON_SIZE_MENU
	ICON_SIZE_SMALL_TOOLBAR = C.GTK_ICON_SIZE_SMALL_TOOLBAR
	ICON_SIZE_LARGE_TOOLBAR = C.GTK_ICON_SIZE_LARGE_TOOLBAR
	ICON_SIZE_BUTTON        = C.GTK_ICON_SIZE_BUTTON
	ICON_SIZE_DND           = C.GTK_ICON_SIZE_DND
	ICON_SIZE_DIALOG        = C.GTK_ICON_SIZE_DIALOG

	// IconThemeError
	ICON_THEME_NOT_FOUND = C.GTK_ICON_THEME_NOT_FOUND
	ICON_THEME_FAILED    = C.GTK_ICON_THEME_FAILED

	// IconViewDropPosition
	ICON_VIEW_NO_DROP    = C.GTK_ICON_VIEW_NO_DROP
	ICON_VIEW_DROP_INTO  = C.GTK_ICON_VIEW_DROP_INTO
	ICON_VIEW_DROP_LEFT  = C.GTK_ICON_VIEW_DROP_LEFT
	ICON_VIEW_DROP_RIGHT = C.GTK_ICON_VIEW_DROP_RIGHT
	ICON_VIEW_DROP_ABOVE = C.GTK_ICON_VIEW_DROP_ABOVE
	ICON_VIEW_DROP_BELOW = C.GTK_ICON_VIEW_DROP_BELOW

	// ImageType
	IMAGE_EMPTY     = C.GTK_IMAGE_EMPTY
	IMAGE_PIXBUF    = C.GTK_IMAGE_PIXBUF
	IMAGE_STOCK     = C.GTK_IMAGE_STOCK
	IMAGE_ICON_SET  = C.GTK_IMAGE_ICON_SET
	IMAGE_ANIMATION = C.GTK_IMAGE_ANIMATION
	IMAGE_ICON_NAME = C.GTK_IMAGE_ICON_NAME
	IMAGE_GICON     = C.GTK_IMAGE_GICON
	IMAGE_SURFACE   = C.GTK_IMAGE_SURFACE

	// InputPurpose
	INPUT_PURPOSE_FREE_FORM = C.GTK_INPUT_PURPOSE_FREE_FORM
	INPUT_PURPOSE_ALPHA     = C.GTK_INPUT_PURPOSE_ALPHA
	INPUT_PURPOSE_DIGITS    = C.GTK_INPUT_PURPOSE_DIGITS
	INPUT_PURPOSE_NUMBER    = C.GTK_INPUT_PURPOSE_NUMBER
	INPUT_PURPOSE_PHONE     = C.GTK_INPUT_PURPOSE_PHONE
	INPUT_PURPOSE_URL       = C.GTK_INPUT_PURPOSE_URL
	INPUT_PURPOSE_EMAIL     = C.GTK_INPUT_PURPOSE_EMAIL
	INPUT_PURPOSE_NAME      = C.GTK_INPUT_PURPOSE_NAME
	INPUT_PURPOSE_PASSWORD  = C.GTK_INPUT_PURPOSE_PASSWORD
	INPUT_PURPOSE_PIN       = C.GTK_INPUT_PURPOSE_PIN

	// Justification
	JUSTIFY_LEFT   = C.GTK_JUSTIFY_LEFT
	JUSTIFY_RIGHT  = C.GTK_JUSTIFY_RIGHT
	JUSTIFY_CENTER = C.GTK_JUSTIFY_CENTER
	JUSTIFY_FILL   = C.GTK_JUSTIFY_FILL

	// LevelBarMode
	LEVEL_BAR_MODE_CONTINUOUS = C.GTK_LEVEL_BAR_MODE_CONTINUOUS
	LEVEL_BAR_MODE_DISCRETE   = C.GTK_LEVEL_BAR_MODE_DISCRETE

	// License
	LICENSE_UNKNOWN       = C.GTK_LICENSE_UNKNOWN
	LICENSE_CUSTOM        = C.GTK_LICENSE_CUSTOM
	LICENSE_GPL_2_0       = C.GTK_LICENSE_GPL_2_0
	LICENSE_GPL_3_0       = C.GTK_LICENSE_GPL_3_0
	LICENSE_LGPL_2_1      = C.GTK_LICENSE_LGPL_2_1
	LICENSE_LGPL_3_0      = C.GTK_LICENSE_LGPL_3_0
	LICENSE_BSD           = C.GTK_LICENSE_BSD
	LICENSE_MIT_X11       = C.GTK_LICENSE_MIT_X11
	LICENSE_ARTISTIC      = C.GTK_LICENSE_ARTISTIC
	LICENSE_GPL_2_0_ONLY  = C.GTK_LICENSE_GPL_2_0_ONLY
	LICENSE_GPL_3_0_ONLY  = C.GTK_LICENSE_GPL_3_0_ONLY
	LICENSE_LGPL_2_1_ONLY = C.GTK_LICENSE_LGPL_2_1_ONLY
	LICENSE_LGPL_3_0_ONLY = C.GTK_LICENSE_LGPL_3_0_ONLY

	// MenuDirectionType
	MENU_DIR_PARENT = C.GTK_MENU_DIR_PARENT
	MENU_DIR_CHILD  = C.GTK_MENU_DIR_CHILD
	MENU_DIR_NEXT   = C.GTK_MENU_DIR_NEXT
	MENU_DIR_PREV   = C.GTK_MENU_DIR_PREV

	// MessageType
	MESSAGE_INFO     = C.GTK_MESSAGE_INFO
	MESSAGE_WARNING  = C.GTK_MESSAGE_WARNING
	MESSAGE_QUESTION = C.GTK_MESSAGE_QUESTION
	MESSAGE_ERROR    = C.GTK_MESSAGE_ERROR
	MESSAGE_OTHER    = C.GTK_MESSAGE_OTHER

	// MovementStep
	MOVEMENT_LOGICAL_POSITIONS = C.GTK_MOVEMENT_LOGICAL_POSITIONS
	MOVEMENT_VISUAL_POSITIONS  = C.GTK_MOVEMENT_VISUAL_POSITIONS
	MOVEMENT_WORDS             = C.GTK_MOVEMENT_WORDS
	MOVEMENT_DISPLAY_LINES     = C.GTK_MOVEMENT_DISPLAY_LINES
	MOVEMENT_DISPLAY_LINE_ENDS = C.GTK_MOVEMENT_DISPLAY_LINE_ENDS
	MOVEMENT_PARAGRAPHS        = C.GTK_MOVEMENT_PARAGRAPHS
	MOVEMENT_PARAGRAPH_ENDS    = C.GTK_MOVEMENT_PARAGRAPH_ENDS
	MOVEMENT_PAGES             = C.GTK_MOVEMENT_PAGES
	MOVEMENT_BUFFER_ENDS       = C.GTK_MOVEMENT_BUFFER_ENDS
	MOVEMENT_HORIZONTAL_PAGES  = C.GTK_MOVEMENT_HORIZONTAL_PAGES

	// NotebookTab
	NOTEBOOK_TAB_FIRST = C.GTK_NOTEBOOK_TAB_FIRST
	NOTEBOOK_TAB_LAST  = C.GTK_NOTEBOOK_TAB_LAST

	// NumberUpLayout
	NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_TOP_TO_BOTTOM = C.GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_TOP_TO_BOTTOM
	NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_BOTTOM_TO_TOP = C.GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_BOTTOM_TO_TOP
	NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_TOP_TO_BOTTOM = C.GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_TOP_TO_BOTTOM
	NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_BOTTOM_TO_TOP = C.GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_BOTTOM_TO_TOP
	NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_LEFT_TO_RIGHT = C.GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_LEFT_TO_RIGHT
	NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_RIGHT_TO_LEFT = C.GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_RIGHT_TO_LEFT
	NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_LEFT_TO_RIGHT = C.GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_LEFT_TO_RIGHT
	NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_RIGHT_TO_LEFT = C.GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_RIGHT_TO_LEFT

	// Orientation
	ORIENTATION_HORIZONTAL = C.GTK_ORIENTATION_HORIZONTAL
	ORIENTATION_VERTICAL   = C.GTK_ORIENTATION_VERTICAL

	// PackDirection
	PACK_DIRECTION_LTR = C.GTK_PACK_DIRECTION_LTR
	PACK_DIRECTION_RTL = C.GTK_PACK_DIRECTION_RTL
	PACK_DIRECTION_TTB = C.GTK_PACK_DIRECTION_TTB
	PACK_DIRECTION_BTT = C.GTK_PACK_DIRECTION_BTT

	// PackType
	PACK_START = C.GTK_PACK_START
	PACK_END   = C.GTK_PACK_END

	// PageOrientation
	PAGE_ORIENTATION_PORTRAIT          = C.GTK_PAGE_ORIENTATION_PORTRAIT
	PAGE_ORIENTATION_LANDSCAPE         = C.GTK_PAGE_ORIENTATION_LANDSCAPE
	PAGE_ORIENTATION_REVERSE_PORTRAIT  = C.GTK_PAGE_ORIENTATION_REVERSE_PORTRAIT
	PAGE_ORIENTATION_REVERSE_LANDSCAPE = C.GTK_PAGE_ORIENTATION_REVERSE_LANDSCAPE

	// PageSet
	PAGE_SET_ALL  = C.GTK_PAGE_SET_ALL
	PAGE_SET_EVEN = C.GTK_PAGE_SET_EVEN
	PAGE_SET_ODD  = C.GTK_PAGE_SET_ODD

	// PathPriorityType
	PATH_PRIO_LOWEST      = C.GTK_PATH_PRIO_LOWEST
	PATH_PRIO_GTK         = C.GTK_PATH_PRIO_GTK
	PATH_PRIO_APPLICATION = C.GTK_PATH_PRIO_APPLICATION
	PATH_PRIO_THEME       = C.GTK_PATH_PRIO_THEME
	PATH_PRIO_RC          = C.GTK_PATH_PRIO_RC
	PATH_PRIO_HIGHEST     = C.GTK_PATH_PRIO_HIGHEST

	// PathType
	PATH_WIDGET       = C.GTK_PATH_WIDGET
	PATH_WIDGET_CLASS = C.GTK_PATH_WIDGET_CLASS
	PATH_CLASS        = C.GTK_PATH_CLASS

	// PolicyType
	POLICY_ALWAYS    = C.GTK_POLICY_ALWAYS
	POLICY_AUTOMATIC = C.GTK_POLICY_AUTOMATIC
	POLICY_NEVER     = C.GTK_POLICY_NEVER

	// PositionType
	POS_LEFT   = C.GTK_POS_LEFT
	POS_RIGHT  = C.GTK_POS_RIGHT
	POS_TOP    = C.GTK_POS_TOP
	POS_BOTTOM = C.GTK_POS_BOTTOM

	// PrintDuplex
	PRINT_DUPLEX_SIMPLEX    = C.GTK_PRINT_DUPLEX_SIMPLEX
	PRINT_DUPLEX_HORIZONTAL = C.GTK_PRINT_DUPLEX_HORIZONTAL
	PRINT_DUPLEX_VERTICAL   = C.GTK_PRINT_DUPLEX_VERTICAL

	// PrintError
	PRINT_ERROR_GENERAL        = C.GTK_PRINT_ERROR_GENERAL
	PRINT_ERROR_INTERNAL_ERROR = C.GTK_PRINT_ERROR_INTERNAL_ERROR
	PRINT_ERROR_NOMEM          = C.GTK_PRINT_ERROR_NOMEM
	PRINT_ERROR_INVALID_FILE   = C.GTK_PRINT_ERROR_INVALID_FILE

	// PrintOperationAction
	PRINT_OPERATION_ACTION_PRINT_DIALOG = C.GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG
	PRINT_OPERATION_ACTION_PRINT        = C.GTK_PRINT_OPERATION_ACTION_PRINT
	PRINT_OPERATION_ACTION_PREVIEW      = C.GTK_PRINT_OPERATION_ACTION_PREVIEW
	PRINT_OPERATION_ACTION_EXPORT       = C.GTK_PRINT_OPERATION_ACTION_EXPORT

	// PrintOperationResult
	PRINT_OPERATION_RESULT_ERROR       = C.GTK_PRINT_OPERATION_RESULT_ERROR
	PRINT_OPERATION_RESULT_APPLY       = C.GTK_PRINT_OPERATION_RESULT_APPLY
	PRINT_OPERATION_RESULT_CANCEL      = C.GTK_PRINT_OPERATION_RESULT_CANCEL
	PRINT_OPERATION_RESULT_IN_PROGRESS = C.GTK_PRINT_OPERATION_RESULT_IN_PROGRESS

	// PrintPages
	PRINT_PAGES_ALL       = C.GTK_PRINT_PAGES_ALL
	PRINT_PAGES_CURRENT   = C.GTK_PRINT_PAGES_CURRENT
	PRINT_PAGES_RANGES    = C.GTK_PRINT_PAGES_RANGES
	PRINT_PAGES_SELECTION = C.GTK_PRINT_PAGES_SELECTION

	// PrintQuality
	PRINT_QUALITY_LOW    = C.GTK_PRINT_QUALITY_LOW
	PRINT_QUALITY_NORMAL = C.GTK_PRINT_QUALITY_NORMAL
	PRINT_QUALITY_HIGH   = C.GTK_PRINT_QUALITY_HIGH
	PRINT_QUALITY_DRAFT  = C.GTK_PRINT_QUALITY_DRAFT

	// PrintStatus
	PRINT_STATUS_INITIAL          = C.GTK_PRINT_STATUS_INITIAL
	PRINT_STATUS_PREPARING        = C.GTK_PRINT_STATUS_PREPARING
	PRINT_STATUS_GENERATING_DATA  = C.GTK_PRINT_STATUS_GENERATING_DATA
	PRINT_STATUS_SENDING_DATA     = C.GTK_PRINT_STATUS_SENDING_DATA
	PRINT_STATUS_PENDING          = C.GTK_PRINT_STATUS_PENDING
	PRINT_STATUS_PENDING_ISSUE    = C.GTK_PRINT_STATUS_PENDING_ISSUE
	PRINT_STATUS_PRINTING         = C.GTK_PRINT_STATUS_PRINTING
	PRINT_STATUS_FINISHED         = C.GTK_PRINT_STATUS_FINISHED
	PRINT_STATUS_FINISHED_ABORTED = C.GTK_PRINT_STATUS_FINISHED_ABORTED

	// RcTokenType
	RC_TOKEN_INVALID        = C.GTK_RC_TOKEN_INVALID
	RC_TOKEN_INCLUDE        = C.GTK_RC_TOKEN_INCLUDE
	RC_TOKEN_NORMAL         = C.GTK_RC_TOKEN_NORMAL
	RC_TOKEN_ACTIVE         = C.GTK_RC_TOKEN_ACTIVE
	RC_TOKEN_PRELIGHT       = C.GTK_RC_TOKEN_PRELIGHT
	RC_TOKEN_SELECTED       = C.GTK_RC_TOKEN_SELECTED
	RC_TOKEN_INSENSITIVE    = C.GTK_RC_TOKEN_INSENSITIVE
	RC_TOKEN_FG             = C.GTK_RC_TOKEN_FG
	RC_TOKEN_BG             = C.GTK_RC_TOKEN_BG
	RC_TOKEN_TEXT           = C.GTK_RC_TOKEN_TEXT
	RC_TOKEN_BASE           = C.GTK_RC_TOKEN_BASE
	RC_TOKEN_XTHICKNESS     = C.GTK_RC_TOKEN_XTHICKNESS
	RC_TOKEN_YTHICKNESS     = C.GTK_RC_TOKEN_YTHICKNESS
	RC_TOKEN_FONT           = C.GTK_RC_TOKEN_FONT
	RC_TOKEN_FONTSET        = C.GTK_RC_TOKEN_FONTSET
	RC_TOKEN_FONT_NAME      = C.GTK_RC_TOKEN_FONT_NAME
	RC_TOKEN_BG_PIXMAP      = C.GTK_RC_TOKEN_BG_PIXMAP
	RC_TOKEN_PIXMAP_PATH    = C.GTK_RC_TOKEN_PIXMAP_PATH
	RC_TOKEN_STYLE          = C.GTK_RC_TOKEN_STYLE
	RC_TOKEN_BINDING        = C.GTK_RC_TOKEN_BINDING
	RC_TOKEN_BIND           = C.GTK_RC_TOKEN_BIND
	RC_TOKEN_WIDGET         = C.GTK_RC_TOKEN_WIDGET
	RC_TOKEN_WIDGET_CLASS   = C.GTK_RC_TOKEN_WIDGET_CLASS
	RC_TOKEN_CLASS          = C.GTK_RC_TOKEN_CLASS
	RC_TOKEN_LOWEST         = C.GTK_RC_TOKEN_LOWEST
	RC_TOKEN_GTK            = C.GTK_RC_TOKEN_GTK
	RC_TOKEN_APPLICATION    = C.GTK_RC_TOKEN_APPLICATION
	RC_TOKEN_THEME          = C.GTK_RC_TOKEN_THEME
	RC_TOKEN_RC             = C.GTK_RC_TOKEN_RC
	RC_TOKEN_HIGHEST        = C.GTK_RC_TOKEN_HIGHEST
	RC_TOKEN_ENGINE         = C.GTK_RC_TOKEN_ENGINE
	RC_TOKEN_MODULE_PATH    = C.GTK_RC_TOKEN_MODULE_PATH
	RC_TOKEN_IM_MODULE_PATH = C.GTK_RC_TOKEN_IM_MODULE_PATH
	RC_TOKEN_IM_MODULE_FILE = C.GTK_RC_TOKEN_IM_MODULE_FILE
	RC_TOKEN_STOCK          = C.GTK_RC_TOKEN_STOCK
	RC_TOKEN_LTR            = C.GTK_RC_TOKEN_LTR
	RC_TOKEN_RTL            = C.GTK_RC_TOKEN_RTL
	RC_TOKEN_COLOR          = C.GTK_RC_TOKEN_COLOR
	RC_TOKEN_UNBIND         = C.GTK_RC_TOKEN_UNBIND
	RC_TOKEN_LAST           = C.GTK_RC_TOKEN_LAST

	// RecentChooserError
	RECENT_CHOOSER_ERROR_NOT_FOUND   = C.GTK_RECENT_CHOOSER_ERROR_NOT_FOUND
	RECENT_CHOOSER_ERROR_INVALID_URI = C.GTK_RECENT_CHOOSER_ERROR_INVALID_URI

	// RecentManagerError
	RECENT_MANAGER_ERROR_NOT_FOUND        = C.GTK_RECENT_MANAGER_ERROR_NOT_FOUND
	RECENT_MANAGER_ERROR_INVALID_URI      = C.GTK_RECENT_MANAGER_ERROR_INVALID_URI
	RECENT_MANAGER_ERROR_INVALID_ENCODING = C.GTK_RECENT_MANAGER_ERROR_INVALID_ENCODING
	RECENT_MANAGER_ERROR_NOT_REGISTERED   = C.GTK_RECENT_MANAGER_ERROR_NOT_REGISTERED
	RECENT_MANAGER_ERROR_READ             = C.GTK_RECENT_MANAGER_ERROR_READ
	RECENT_MANAGER_ERROR_WRITE            = C.GTK_RECENT_MANAGER_ERROR_WRITE
	RECENT_MANAGER_ERROR_UNKNOWN          = C.GTK_RECENT_MANAGER_ERROR_UNKNOWN

	// RecentSortType
	RECENT_SORT_NONE   = C.GTK_RECENT_SORT_NONE
	RECENT_SORT_MRU    = C.GTK_RECENT_SORT_MRU
	RECENT_SORT_LRU    = C.GTK_RECENT_SORT_LRU
	RECENT_SORT_CUSTOM = C.GTK_RECENT_SORT_CUSTOM

	// ReliefStyle
	RELIEF_NORMAL = C.GTK_RELIEF_NORMAL
	RELIEF_HALF   = C.GTK_RELIEF_HALF
	RELIEF_NONE   = C.GTK_RELIEF_NONE

	// ResizeMode
	RESIZE_PARENT    = C.GTK_RESIZE_PARENT
	RESIZE_QUEUE     = C.GTK_RESIZE_QUEUE
	RESIZE_IMMEDIATE = C.GTK_RESIZE_IMMEDIATE

	// ResponseType
	RESPONSE_NONE         = C.GTK_RESPONSE_NONE
	RESPONSE_REJECT       = C.GTK_RESPONSE_REJECT
	RESPONSE_ACCEPT       = C.GTK_RESPONSE_ACCEPT
	RESPONSE_DELETE_EVENT = C.GTK_RESPONSE_DELETE_EVENT
	RESPONSE_OK           = C.GTK_RESPONSE_OK
	RESPONSE_CANCEL       = C.GTK_RESPONSE_CANCEL
	RESPONSE_CLOSE        = C.GTK_RESPONSE_CLOSE
	RESPONSE_YES          = C.GTK_RESPONSE_YES
	RESPONSE_NO           = C.GTK_RESPONSE_NO
	RESPONSE_APPLY        = C.GTK_RESPONSE_APPLY
	RESPONSE_HELP         = C.GTK_RESPONSE_HELP

	// RevealerTransitionType
	REVEALER_TRANSITION_TYPE_NONE        = C.GTK_REVEALER_TRANSITION_TYPE_NONE
	REVEALER_TRANSITION_TYPE_CROSSFADE   = C.GTK_REVEALER_TRANSITION_TYPE_CROSSFADE
	REVEALER_TRANSITION_TYPE_SLIDE_RIGHT = C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_RIGHT
	REVEALER_TRANSITION_TYPE_SLIDE_LEFT  = C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_LEFT
	REVEALER_TRANSITION_TYPE_SLIDE_UP    = C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_UP
	REVEALER_TRANSITION_TYPE_SLIDE_DOWN  = C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_DOWN

	// ScrollStep
	SCROLL_STEPS            = C.GTK_SCROLL_STEPS
	SCROLL_PAGES            = C.GTK_SCROLL_PAGES
	SCROLL_ENDS             = C.GTK_SCROLL_ENDS
	SCROLL_HORIZONTAL_STEPS = C.GTK_SCROLL_HORIZONTAL_STEPS
	SCROLL_HORIZONTAL_PAGES = C.GTK_SCROLL_HORIZONTAL_PAGES
	SCROLL_HORIZONTAL_ENDS  = C.GTK_SCROLL_HORIZONTAL_ENDS

	// ScrollType
	SCROLL_NONE          = C.GTK_SCROLL_NONE
	SCROLL_JUMP          = C.GTK_SCROLL_JUMP
	SCROLL_STEP_BACKWARD = C.GTK_SCROLL_STEP_BACKWARD
	SCROLL_STEP_FORWARD  = C.GTK_SCROLL_STEP_FORWARD
	SCROLL_PAGE_BACKWARD = C.GTK_SCROLL_PAGE_BACKWARD
	SCROLL_PAGE_FORWARD  = C.GTK_SCROLL_PAGE_FORWARD
	SCROLL_STEP_UP       = C.GTK_SCROLL_STEP_UP
	SCROLL_STEP_DOWN     = C.GTK_SCROLL_STEP_DOWN
	SCROLL_PAGE_UP       = C.GTK_SCROLL_PAGE_UP
	SCROLL_PAGE_DOWN     = C.GTK_SCROLL_PAGE_DOWN
	SCROLL_STEP_LEFT     = C.GTK_SCROLL_STEP_LEFT
	SCROLL_STEP_RIGHT    = C.GTK_SCROLL_STEP_RIGHT
	SCROLL_PAGE_LEFT     = C.GTK_SCROLL_PAGE_LEFT
	SCROLL_PAGE_RIGHT    = C.GTK_SCROLL_PAGE_RIGHT
	SCROLL_START         = C.GTK_SCROLL_START
	SCROLL_END           = C.GTK_SCROLL_END

	// ScrollablePolicy
	SCROLL_MINIMUM = C.GTK_SCROLL_MINIMUM
	SCROLL_NATURAL = C.GTK_SCROLL_NATURAL

	// SelectionMode
	SELECTION_NONE     = C.GTK_SELECTION_NONE
	SELECTION_SINGLE   = C.GTK_SELECTION_SINGLE
	SELECTION_BROWSE   = C.GTK_SELECTION_BROWSE
	SELECTION_MULTIPLE = C.GTK_SELECTION_MULTIPLE

	// SensitivityType
	SENSITIVITY_AUTO = C.GTK_SENSITIVITY_AUTO
	SENSITIVITY_ON   = C.GTK_SENSITIVITY_ON
	SENSITIVITY_OFF  = C.GTK_SENSITIVITY_OFF

	// ShadowType
	SHADOW_NONE       = C.GTK_SHADOW_NONE
	SHADOW_IN         = C.GTK_SHADOW_IN
	SHADOW_OUT        = C.GTK_SHADOW_OUT
	SHADOW_ETCHED_IN  = C.GTK_SHADOW_ETCHED_IN
	SHADOW_ETCHED_OUT = C.GTK_SHADOW_ETCHED_OUT

	// SizeGroupMode
	SIZE_GROUP_NONE       = C.GTK_SIZE_GROUP_NONE
	SIZE_GROUP_HORIZONTAL = C.GTK_SIZE_GROUP_HORIZONTAL
	SIZE_GROUP_VERTICAL   = C.GTK_SIZE_GROUP_VERTICAL
	SIZE_GROUP_BOTH       = C.GTK_SIZE_GROUP_BOTH

	// SizeRequestMode
	SIZE_REQUEST_HEIGHT_FOR_WIDTH = C.GTK_SIZE_REQUEST_HEIGHT_FOR_WIDTH
	SIZE_REQUEST_WIDTH_FOR_HEIGHT = C.GTK_SIZE_REQUEST_WIDTH_FOR_HEIGHT
	SIZE_REQUEST_CONSTANT_SIZE    = C.GTK_SIZE_REQUEST_CONSTANT_SIZE

	// SortType
	SORT_ASCENDING  = C.GTK_SORT_ASCENDING
	SORT_DESCENDING = C.GTK_SORT_DESCENDING

	// SpinButtonUpdatePolicy
	UPDATE_ALWAYS   = C.GTK_UPDATE_ALWAYS
	UPDATE_IF_VALID = C.GTK_UPDATE_IF_VALID

	// SpinType
	SPIN_STEP_FORWARD  = C.GTK_SPIN_STEP_FORWARD
	SPIN_STEP_BACKWARD = C.GTK_SPIN_STEP_BACKWARD
	SPIN_PAGE_FORWARD  = C.GTK_SPIN_PAGE_FORWARD
	SPIN_PAGE_BACKWARD = C.GTK_SPIN_PAGE_BACKWARD
	SPIN_HOME          = C.GTK_SPIN_HOME
	SPIN_END           = C.GTK_SPIN_END
	SPIN_USER_DEFINED  = C.GTK_SPIN_USER_DEFINED

	// StackTransitionType
	STACK_TRANSITION_TYPE_NONE             = C.GTK_STACK_TRANSITION_TYPE_NONE
	STACK_TRANSITION_TYPE_CROSSFADE        = C.GTK_STACK_TRANSITION_TYPE_CROSSFADE
	STACK_TRANSITION_TYPE_SLIDE_RIGHT      = C.GTK_STACK_TRANSITION_TYPE_SLIDE_RIGHT
	STACK_TRANSITION_TYPE_SLIDE_LEFT       = C.GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT
	STACK_TRANSITION_TYPE_SLIDE_UP         = C.GTK_STACK_TRANSITION_TYPE_SLIDE_UP
	STACK_TRANSITION_TYPE_SLIDE_DOWN       = C.GTK_STACK_TRANSITION_TYPE_SLIDE_DOWN
	STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT = C.GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT
	STACK_TRANSITION_TYPE_SLIDE_UP_DOWN    = C.GTK_STACK_TRANSITION_TYPE_SLIDE_UP_DOWN
	STACK_TRANSITION_TYPE_OVER_UP          = C.GTK_STACK_TRANSITION_TYPE_OVER_UP
	STACK_TRANSITION_TYPE_OVER_DOWN        = C.GTK_STACK_TRANSITION_TYPE_OVER_DOWN
	STACK_TRANSITION_TYPE_OVER_LEFT        = C.GTK_STACK_TRANSITION_TYPE_OVER_LEFT
	STACK_TRANSITION_TYPE_OVER_RIGHT       = C.GTK_STACK_TRANSITION_TYPE_OVER_RIGHT
	STACK_TRANSITION_TYPE_UNDER_UP         = C.GTK_STACK_TRANSITION_TYPE_UNDER_UP
	STACK_TRANSITION_TYPE_UNDER_DOWN       = C.GTK_STACK_TRANSITION_TYPE_UNDER_DOWN
	STACK_TRANSITION_TYPE_UNDER_LEFT       = C.GTK_STACK_TRANSITION_TYPE_UNDER_LEFT
	STACK_TRANSITION_TYPE_UNDER_RIGHT      = C.GTK_STACK_TRANSITION_TYPE_UNDER_RIGHT
	STACK_TRANSITION_TYPE_OVER_UP_DOWN     = C.GTK_STACK_TRANSITION_TYPE_OVER_UP_DOWN

	// StateType
	STATE_NORMAL       = C.GTK_STATE_NORMAL
	STATE_ACTIVE       = C.GTK_STATE_ACTIVE
	STATE_PRELIGHT     = C.GTK_STATE_PRELIGHT
	STATE_SELECTED     = C.GTK_STATE_SELECTED
	STATE_INSENSITIVE  = C.GTK_STATE_INSENSITIVE
	STATE_INCONSISTENT = C.GTK_STATE_INCONSISTENT
	STATE_FOCUSED      = C.GTK_STATE_FOCUSED

	// TextBufferTargetInfo
	TEXT_BUFFER_TARGET_INFO_BUFFER_CONTENTS = C.GTK_TEXT_BUFFER_TARGET_INFO_BUFFER_CONTENTS
	TEXT_BUFFER_TARGET_INFO_RICH_TEXT       = C.GTK_TEXT_BUFFER_TARGET_INFO_RICH_TEXT
	TEXT_BUFFER_TARGET_INFO_TEXT            = C.GTK_TEXT_BUFFER_TARGET_INFO_TEXT

	// TextDirection
	TEXT_DIR_NONE = C.GTK_TEXT_DIR_NONE
	TEXT_DIR_LTR  = C.GTK_TEXT_DIR_LTR
	TEXT_DIR_RTL  = C.GTK_TEXT_DIR_RTL

	// TextWindowType
	TEXT_WINDOW_PRIVATE = C.GTK_TEXT_WINDOW_PRIVATE
	TEXT_WINDOW_WIDGET  = C.GTK_TEXT_WINDOW_WIDGET
	TEXT_WINDOW_TEXT    = C.GTK_TEXT_WINDOW_TEXT
	TEXT_WINDOW_LEFT    = C.GTK_TEXT_WINDOW_LEFT
	TEXT_WINDOW_RIGHT   = C.GTK_TEXT_WINDOW_RIGHT
	TEXT_WINDOW_TOP     = C.GTK_TEXT_WINDOW_TOP
	TEXT_WINDOW_BOTTOM  = C.GTK_TEXT_WINDOW_BOTTOM

	// ToolbarSpaceStyle
	TOOLBAR_SPACE_EMPTY = C.GTK_TOOLBAR_SPACE_EMPTY
	TOOLBAR_SPACE_LINE  = C.GTK_TOOLBAR_SPACE_LINE

	// ToolbarStyle
	TOOLBAR_ICONS      = C.GTK_TOOLBAR_ICONS
	TOOLBAR_TEXT       = C.GTK_TOOLBAR_TEXT
	TOOLBAR_BOTH       = C.GTK_TOOLBAR_BOTH
	TOOLBAR_BOTH_HORIZ = C.GTK_TOOLBAR_BOTH_HORIZ

	// TreeViewColumnSizing
	TREE_VIEW_COLUMN_GROW_ONLY = C.GTK_TREE_VIEW_COLUMN_GROW_ONLY
	TREE_VIEW_COLUMN_AUTOSIZE  = C.GTK_TREE_VIEW_COLUMN_AUTOSIZE
	TREE_VIEW_COLUMN_FIXED     = C.GTK_TREE_VIEW_COLUMN_FIXED

	// TreeViewDropPosition
	TREE_VIEW_DROP_BEFORE         = C.GTK_TREE_VIEW_DROP_BEFORE
	TREE_VIEW_DROP_AFTER          = C.GTK_TREE_VIEW_DROP_AFTER
	TREE_VIEW_DROP_INTO_OR_BEFORE = C.GTK_TREE_VIEW_DROP_INTO_OR_BEFORE
	TREE_VIEW_DROP_INTO_OR_AFTER  = C.GTK_TREE_VIEW_DROP_INTO_OR_AFTER

	// TreeViewGridLines
	TREE_VIEW_GRID_LINES_NONE       = C.GTK_TREE_VIEW_GRID_LINES_NONE
	TREE_VIEW_GRID_LINES_HORIZONTAL = C.GTK_TREE_VIEW_GRID_LINES_HORIZONTAL
	TREE_VIEW_GRID_LINES_VERTICAL   = C.GTK_TREE_VIEW_GRID_LINES_VERTICAL
	TREE_VIEW_GRID_LINES_BOTH       = C.GTK_TREE_VIEW_GRID_LINES_BOTH

	// Unit
	UNIT_NONE   = C.GTK_UNIT_NONE
	UNIT_POINTS = C.GTK_UNIT_POINTS
	UNIT_INCH   = C.GTK_UNIT_INCH
	UNIT_MM     = C.GTK_UNIT_MM

	// WidgetHelpType
	WIDGET_HELP_TOOLTIP    = C.GTK_WIDGET_HELP_TOOLTIP
	WIDGET_HELP_WHATS_THIS = C.GTK_WIDGET_HELP_WHATS_THIS

	// WindowPosition
	WIN_POS_NONE             = C.GTK_WIN_POS_NONE
	WIN_POS_CENTER           = C.GTK_WIN_POS_CENTER
	WIN_POS_MOUSE            = C.GTK_WIN_POS_MOUSE
	WIN_POS_CENTER_ALWAYS    = C.GTK_WIN_POS_CENTER_ALWAYS
	WIN_POS_CENTER_ON_PARENT = C.GTK_WIN_POS_CENTER_ON_PARENT

	// WindowType
	WINDOW_TOPLEVEL = C.GTK_WINDOW_TOPLEVEL
	WINDOW_POPUP    = C.GTK_WINDOW_POPUP

	// WrapMode
	WRAP_NONE      = C.GTK_WRAP_NONE
	WRAP_CHAR      = C.GTK_WRAP_CHAR
	WRAP_WORD      = C.GTK_WRAP_WORD
	WRAP_WORD_CHAR = C.GTK_WRAP_WORD_CHAR
)