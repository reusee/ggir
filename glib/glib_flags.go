package glib

/*
#include <glib.h>
#include <glib-object.h>
#include <glib-unix.h>
#include <glib/gstdio.h>
#include <stdlib.h>
*/
import "C"

var (
	// AsciiType
	ASCII_ALNUM  = C.GAsciiType(C.G_ASCII_ALNUM)
	ASCII_ALPHA  = C.GAsciiType(C.G_ASCII_ALPHA)
	ASCII_CNTRL  = C.GAsciiType(C.G_ASCII_CNTRL)
	ASCII_DIGIT  = C.GAsciiType(C.G_ASCII_DIGIT)
	ASCII_GRAPH  = C.GAsciiType(C.G_ASCII_GRAPH)
	ASCII_LOWER  = C.GAsciiType(C.G_ASCII_LOWER)
	ASCII_PRINT  = C.GAsciiType(C.G_ASCII_PRINT)
	ASCII_PUNCT  = C.GAsciiType(C.G_ASCII_PUNCT)
	ASCII_SPACE  = C.GAsciiType(C.G_ASCII_SPACE)
	ASCII_UPPER  = C.GAsciiType(C.G_ASCII_UPPER)
	ASCII_XDIGIT = C.GAsciiType(C.G_ASCII_XDIGIT)

	// FileTest
	FILE_TEST_IS_REGULAR    = C.GFileTest(C.G_FILE_TEST_IS_REGULAR)
	FILE_TEST_IS_SYMLINK    = C.GFileTest(C.G_FILE_TEST_IS_SYMLINK)
	FILE_TEST_IS_DIR        = C.GFileTest(C.G_FILE_TEST_IS_DIR)
	FILE_TEST_IS_EXECUTABLE = C.GFileTest(C.G_FILE_TEST_IS_EXECUTABLE)
	FILE_TEST_EXISTS        = C.GFileTest(C.G_FILE_TEST_EXISTS)

	// FormatSizeFlags
	FORMAT_SIZE_DEFAULT     = C.GFormatSizeFlags(C.G_FORMAT_SIZE_DEFAULT)
	FORMAT_SIZE_LONG_FORMAT = C.GFormatSizeFlags(C.G_FORMAT_SIZE_LONG_FORMAT)
	FORMAT_SIZE_IEC_UNITS   = C.GFormatSizeFlags(C.G_FORMAT_SIZE_IEC_UNITS)

	// HookFlagMask
	HOOK_FLAG_ACTIVE  = C.GHookFlagMask(C.G_HOOK_FLAG_ACTIVE)
	HOOK_FLAG_IN_CALL = C.GHookFlagMask(C.G_HOOK_FLAG_IN_CALL)
	HOOK_FLAG_MASK    = C.GHookFlagMask(C.G_HOOK_FLAG_MASK)

	// IOCondition
	IO_IN   = C.GIOCondition(C.G_IO_IN)
	IO_OUT  = C.GIOCondition(C.G_IO_OUT)
	IO_PRI  = C.GIOCondition(C.G_IO_PRI)
	IO_ERR  = C.GIOCondition(C.G_IO_ERR)
	IO_HUP  = C.GIOCondition(C.G_IO_HUP)
	IO_NVAL = C.GIOCondition(C.G_IO_NVAL)

	// IOFlags
	IO_FLAG_APPEND       = C.GIOFlags(C.G_IO_FLAG_APPEND)
	IO_FLAG_NONBLOCK     = C.GIOFlags(C.G_IO_FLAG_NONBLOCK)
	IO_FLAG_IS_READABLE  = C.GIOFlags(C.G_IO_FLAG_IS_READABLE)
	IO_FLAG_IS_WRITABLE  = C.GIOFlags(C.G_IO_FLAG_IS_WRITABLE)
	IO_FLAG_IS_WRITEABLE = C.GIOFlags(C.G_IO_FLAG_IS_WRITEABLE)
	IO_FLAG_IS_SEEKABLE  = C.GIOFlags(C.G_IO_FLAG_IS_SEEKABLE)
	IO_FLAG_MASK         = C.GIOFlags(C.G_IO_FLAG_MASK)
	IO_FLAG_GET_MASK     = C.GIOFlags(C.G_IO_FLAG_GET_MASK)
	IO_FLAG_SET_MASK     = C.GIOFlags(C.G_IO_FLAG_SET_MASK)

	// KeyFileFlags
	KEY_FILE_NONE              = C.GKeyFileFlags(C.G_KEY_FILE_NONE)
	KEY_FILE_KEEP_COMMENTS     = C.GKeyFileFlags(C.G_KEY_FILE_KEEP_COMMENTS)
	KEY_FILE_KEEP_TRANSLATIONS = C.GKeyFileFlags(C.G_KEY_FILE_KEEP_TRANSLATIONS)

	// LogLevelFlags
	LOG_FLAG_RECURSION = C.GLogLevelFlags(C.G_LOG_FLAG_RECURSION)
	LOG_FLAG_FATAL     = C.GLogLevelFlags(C.G_LOG_FLAG_FATAL)
	LOG_LEVEL_ERROR    = C.GLogLevelFlags(C.G_LOG_LEVEL_ERROR)
	LOG_LEVEL_CRITICAL = C.GLogLevelFlags(C.G_LOG_LEVEL_CRITICAL)
	LOG_LEVEL_WARNING  = C.GLogLevelFlags(C.G_LOG_LEVEL_WARNING)
	LOG_LEVEL_MESSAGE  = C.GLogLevelFlags(C.G_LOG_LEVEL_MESSAGE)
	LOG_LEVEL_INFO     = C.GLogLevelFlags(C.G_LOG_LEVEL_INFO)
	LOG_LEVEL_DEBUG    = C.GLogLevelFlags(C.G_LOG_LEVEL_DEBUG)
	LOG_LEVEL_MASK     = C.GLogLevelFlags(C.G_LOG_LEVEL_MASK)

	// MarkupCollectType
	MARKUP_COLLECT_INVALID  = C.GMarkupCollectType(C.G_MARKUP_COLLECT_INVALID)
	MARKUP_COLLECT_STRING   = C.GMarkupCollectType(C.G_MARKUP_COLLECT_STRING)
	MARKUP_COLLECT_STRDUP   = C.GMarkupCollectType(C.G_MARKUP_COLLECT_STRDUP)
	MARKUP_COLLECT_BOOLEAN  = C.GMarkupCollectType(C.G_MARKUP_COLLECT_BOOLEAN)
	MARKUP_COLLECT_TRISTATE = C.GMarkupCollectType(C.G_MARKUP_COLLECT_TRISTATE)
	MARKUP_COLLECT_OPTIONAL = C.GMarkupCollectType(C.G_MARKUP_COLLECT_OPTIONAL)

	// MarkupParseFlags
	MARKUP_DO_NOT_USE_THIS_UNSUPPORTED_FLAG = C.GMarkupParseFlags(C.G_MARKUP_DO_NOT_USE_THIS_UNSUPPORTED_FLAG)
	MARKUP_TREAT_CDATA_AS_TEXT              = C.GMarkupParseFlags(C.G_MARKUP_TREAT_CDATA_AS_TEXT)
	MARKUP_PREFIX_ERROR_POSITION            = C.GMarkupParseFlags(C.G_MARKUP_PREFIX_ERROR_POSITION)
	MARKUP_IGNORE_QUALIFIED                 = C.GMarkupParseFlags(C.G_MARKUP_IGNORE_QUALIFIED)

	// OptionFlags
	OPTION_FLAG_NONE         = C.GOptionFlags(C.G_OPTION_FLAG_NONE)
	OPTION_FLAG_HIDDEN       = C.GOptionFlags(C.G_OPTION_FLAG_HIDDEN)
	OPTION_FLAG_IN_MAIN      = C.GOptionFlags(C.G_OPTION_FLAG_IN_MAIN)
	OPTION_FLAG_REVERSE      = C.GOptionFlags(C.G_OPTION_FLAG_REVERSE)
	OPTION_FLAG_NO_ARG       = C.GOptionFlags(C.G_OPTION_FLAG_NO_ARG)
	OPTION_FLAG_FILENAME     = C.GOptionFlags(C.G_OPTION_FLAG_FILENAME)
	OPTION_FLAG_OPTIONAL_ARG = C.GOptionFlags(C.G_OPTION_FLAG_OPTIONAL_ARG)
	OPTION_FLAG_NOALIAS      = C.GOptionFlags(C.G_OPTION_FLAG_NOALIAS)

	// RegexCompileFlags
	REGEX_CASELESS          = C.GRegexCompileFlags(C.G_REGEX_CASELESS)
	REGEX_MULTILINE         = C.GRegexCompileFlags(C.G_REGEX_MULTILINE)
	REGEX_DOTALL            = C.GRegexCompileFlags(C.G_REGEX_DOTALL)
	REGEX_EXTENDED          = C.GRegexCompileFlags(C.G_REGEX_EXTENDED)
	REGEX_ANCHORED          = C.GRegexCompileFlags(C.G_REGEX_ANCHORED)
	REGEX_DOLLAR_ENDONLY    = C.GRegexCompileFlags(C.G_REGEX_DOLLAR_ENDONLY)
	REGEX_UNGREEDY          = C.GRegexCompileFlags(C.G_REGEX_UNGREEDY)
	REGEX_RAW               = C.GRegexCompileFlags(C.G_REGEX_RAW)
	REGEX_NO_AUTO_CAPTURE   = C.GRegexCompileFlags(C.G_REGEX_NO_AUTO_CAPTURE)
	REGEX_OPTIMIZE          = C.GRegexCompileFlags(C.G_REGEX_OPTIMIZE)
	REGEX_FIRSTLINE         = C.GRegexCompileFlags(C.G_REGEX_FIRSTLINE)
	REGEX_DUPNAMES          = C.GRegexCompileFlags(C.G_REGEX_DUPNAMES)
	REGEX_NEWLINE_CR        = C.GRegexCompileFlags(C.G_REGEX_NEWLINE_CR)
	REGEX_NEWLINE_LF        = C.GRegexCompileFlags(C.G_REGEX_NEWLINE_LF)
	REGEX_NEWLINE_CRLF      = C.GRegexCompileFlags(C.G_REGEX_NEWLINE_CRLF)
	REGEX_NEWLINE_ANYCRLF   = C.GRegexCompileFlags(C.G_REGEX_NEWLINE_ANYCRLF)
	REGEX_BSR_ANYCRLF       = C.GRegexCompileFlags(C.G_REGEX_BSR_ANYCRLF)
	REGEX_JAVASCRIPT_COMPAT = C.GRegexCompileFlags(C.G_REGEX_JAVASCRIPT_COMPAT)

	// RegexMatchFlags
	REGEX_MATCH_ANCHORED         = C.GRegexMatchFlags(C.G_REGEX_MATCH_ANCHORED)
	REGEX_MATCH_NOTBOL           = C.GRegexMatchFlags(C.G_REGEX_MATCH_NOTBOL)
	REGEX_MATCH_NOTEOL           = C.GRegexMatchFlags(C.G_REGEX_MATCH_NOTEOL)
	REGEX_MATCH_NOTEMPTY         = C.GRegexMatchFlags(C.G_REGEX_MATCH_NOTEMPTY)
	REGEX_MATCH_PARTIAL          = C.GRegexMatchFlags(C.G_REGEX_MATCH_PARTIAL)
	REGEX_MATCH_NEWLINE_CR       = C.GRegexMatchFlags(C.G_REGEX_MATCH_NEWLINE_CR)
	REGEX_MATCH_NEWLINE_LF       = C.GRegexMatchFlags(C.G_REGEX_MATCH_NEWLINE_LF)
	REGEX_MATCH_NEWLINE_CRLF     = C.GRegexMatchFlags(C.G_REGEX_MATCH_NEWLINE_CRLF)
	REGEX_MATCH_NEWLINE_ANY      = C.GRegexMatchFlags(C.G_REGEX_MATCH_NEWLINE_ANY)
	REGEX_MATCH_NEWLINE_ANYCRLF  = C.GRegexMatchFlags(C.G_REGEX_MATCH_NEWLINE_ANYCRLF)
	REGEX_MATCH_BSR_ANYCRLF      = C.GRegexMatchFlags(C.G_REGEX_MATCH_BSR_ANYCRLF)
	REGEX_MATCH_BSR_ANY          = C.GRegexMatchFlags(C.G_REGEX_MATCH_BSR_ANY)
	REGEX_MATCH_PARTIAL_SOFT     = C.GRegexMatchFlags(C.G_REGEX_MATCH_PARTIAL_SOFT)
	REGEX_MATCH_PARTIAL_HARD     = C.GRegexMatchFlags(C.G_REGEX_MATCH_PARTIAL_HARD)
	REGEX_MATCH_NOTEMPTY_ATSTART = C.GRegexMatchFlags(C.G_REGEX_MATCH_NOTEMPTY_ATSTART)

	// SpawnFlags
	SPAWN_DEFAULT                = C.GSpawnFlags(C.G_SPAWN_DEFAULT)
	SPAWN_LEAVE_DESCRIPTORS_OPEN = C.GSpawnFlags(C.G_SPAWN_LEAVE_DESCRIPTORS_OPEN)
	SPAWN_DO_NOT_REAP_CHILD      = C.GSpawnFlags(C.G_SPAWN_DO_NOT_REAP_CHILD)
	SPAWN_SEARCH_PATH            = C.GSpawnFlags(C.G_SPAWN_SEARCH_PATH)
	SPAWN_STDOUT_TO_DEV_NULL     = C.GSpawnFlags(C.G_SPAWN_STDOUT_TO_DEV_NULL)
	SPAWN_STDERR_TO_DEV_NULL     = C.GSpawnFlags(C.G_SPAWN_STDERR_TO_DEV_NULL)
	SPAWN_CHILD_INHERITS_STDIN   = C.GSpawnFlags(C.G_SPAWN_CHILD_INHERITS_STDIN)
	SPAWN_FILE_AND_ARGV_ZERO     = C.GSpawnFlags(C.G_SPAWN_FILE_AND_ARGV_ZERO)
	SPAWN_SEARCH_PATH_FROM_ENVP  = C.GSpawnFlags(C.G_SPAWN_SEARCH_PATH_FROM_ENVP)
	SPAWN_CLOEXEC_PIPES          = C.GSpawnFlags(C.G_SPAWN_CLOEXEC_PIPES)

	// TestSubprocessFlags
	TEST_SUBPROCESS_INHERIT_STDIN  = C.GTestSubprocessFlags(C.G_TEST_SUBPROCESS_INHERIT_STDIN)
	TEST_SUBPROCESS_INHERIT_STDOUT = C.GTestSubprocessFlags(C.G_TEST_SUBPROCESS_INHERIT_STDOUT)
	TEST_SUBPROCESS_INHERIT_STDERR = C.GTestSubprocessFlags(C.G_TEST_SUBPROCESS_INHERIT_STDERR)

	// TestTrapFlags
	TEST_TRAP_SILENCE_STDOUT = C.GTestTrapFlags(C.G_TEST_TRAP_SILENCE_STDOUT)
	TEST_TRAP_SILENCE_STDERR = C.GTestTrapFlags(C.G_TEST_TRAP_SILENCE_STDERR)
	TEST_TRAP_INHERIT_STDIN  = C.GTestTrapFlags(C.G_TEST_TRAP_INHERIT_STDIN)

	// TraverseFlags
	TRAVERSE_LEAVES     = C.GTraverseFlags(C.G_TRAVERSE_LEAVES)
	TRAVERSE_NON_LEAVES = C.GTraverseFlags(C.G_TRAVERSE_NON_LEAVES)
	TRAVERSE_ALL        = C.GTraverseFlags(C.G_TRAVERSE_ALL)
	TRAVERSE_MASK       = C.GTraverseFlags(C.G_TRAVERSE_MASK)
	TRAVERSE_LEAFS      = C.GTraverseFlags(C.G_TRAVERSE_LEAFS)
	TRAVERSE_NON_LEAFS  = C.GTraverseFlags(C.G_TRAVERSE_NON_LEAFS)
)
