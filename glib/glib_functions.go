package glib

/*
#include <stdio.h>
#include <glib.h>
#include <glib-object.h>
#include <glib-unix.h>
#include <glib/gstdio.h>
#cgo pkg-config: glib-2.0 gobject-2.0
*/
import "C"

import (
	"errors"
	"reflect"
	"unsafe"
)

func init() {
	var _ unsafe.Pointer
	var _ reflect.SliceHeader
	_ = errors.New("")
}

/*
A wrapper for the POSIX access() function. This function is used to
test a pathname for one or several of read, write or execute
permissions, or just existence.

On Windows, the file protection mechanism is not at all POSIX-like,
and the underlying function in the C library only checks the
FAT-style READONLY attribute, and does not look at the ACL of a
file at all. This function is this in practise almost useless on
Windows. Software that needs to handle file permissions on Windows
more exactly should use the Win32 API.

See your C library manual for more details about access().
*/
func Access(filename string, mode int) (return__ int) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	var __cgo__return__ C.int
	__cgo__return__ = C.g_access(__cgo__filename, C.int(mode))
	C.free(unsafe.Pointer(__cgo__filename))
	return__ = int(__cgo__return__)
	return
}

/*
Determines the numeric value of a character as a decimal digit.
Differs from g_unichar_digit_value() because it takes a char, so
there's no worry about sign extension if characters are signed.
*/
func AsciiDigitValue(c byte) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_ascii_digit_value(C.gchar(c))
	return__ = int(__cgo__return__)
	return
}

/*
Converts a #gdouble to a string, using the '.' as
decimal point.

This function generates enough precision that converting
the string back using g_ascii_strtod() gives the same machine-number
(on machines with IEEE compatible 64bit doubles). It is
guaranteed that the size of the resulting string will never
be larger than @G_ASCII_DTOSTR_BUF_SIZE bytes.
*/
func AsciiDtostr(buffer string, buf_len int, d float64) (return__ string) {
	__cgo__buffer := (*C.gchar)(unsafe.Pointer(C.CString(buffer)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_ascii_dtostr(__cgo__buffer, C.gint(buf_len), C.gdouble(d))
	C.free(unsafe.Pointer(__cgo__buffer))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts a #gdouble to a string, using the '.' as
decimal point. To format the number you pass in
a printf()-style format string. Allowed conversion
specifiers are 'e', 'E', 'f', 'F', 'g' and 'G'.

If you just want to want to serialize the value into a
string, use g_ascii_dtostr().
*/
func AsciiFormatd(buffer string, buf_len int, format string, d float64) (return__ string) {
	__cgo__buffer := (*C.gchar)(unsafe.Pointer(C.CString(buffer)))
	__cgo__format := (*C.gchar)(unsafe.Pointer(C.CString(format)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_ascii_formatd(__cgo__buffer, C.gint(buf_len), __cgo__format, C.gdouble(d))
	C.free(unsafe.Pointer(__cgo__buffer))
	C.free(unsafe.Pointer(__cgo__format))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Compare two strings, ignoring the case of ASCII characters.

Unlike the BSD strcasecmp() function, this only recognizes standard
ASCII letters and ignores the locale, treating all non-ASCII
bytes as if they are not letters.

This function should be used only on strings that are known to be
in encodings where the bytes corresponding to ASCII letters always
represent themselves. This includes UTF-8 and the ISO-8859-*
charsets, but not for instance double-byte encodings like the
Windows Codepage 932, where the trailing bytes of double-byte
characters include all ASCII letters. If you compare two CP932
strings using this function, you will get false matches.

Both @s1 and @s2 must be non-%NULL.
*/
func AsciiStrcasecmp(s1 string, s2 string) (return__ int) {
	__cgo__s1 := (*C.gchar)(unsafe.Pointer(C.CString(s1)))
	__cgo__s2 := (*C.gchar)(unsafe.Pointer(C.CString(s2)))
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_ascii_strcasecmp(__cgo__s1, __cgo__s2)
	C.free(unsafe.Pointer(__cgo__s1))
	C.free(unsafe.Pointer(__cgo__s2))
	return__ = int(__cgo__return__)
	return
}

/*
Converts all upper case ASCII letters to lower case ASCII letters.
*/
func AsciiStrdown(str string, len_ int64) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_ascii_strdown(__cgo__str, C.gssize(len_))
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Compare @s1 and @s2, ignoring the case of ASCII characters and any
characters after the first @n in each string.

Unlike the BSD strcasecmp() function, this only recognizes standard
ASCII letters and ignores the locale, treating all non-ASCII
characters as if they are not letters.

The same warning as in g_ascii_strcasecmp() applies: Use this
function only on strings known to be in encodings where bytes
corresponding to ASCII letters always represent themselves.
*/
func AsciiStrncasecmp(s1 string, s2 string, n int64) (return__ int) {
	__cgo__s1 := (*C.gchar)(unsafe.Pointer(C.CString(s1)))
	__cgo__s2 := (*C.gchar)(unsafe.Pointer(C.CString(s2)))
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_ascii_strncasecmp(__cgo__s1, __cgo__s2, C.gsize(n))
	C.free(unsafe.Pointer(__cgo__s1))
	C.free(unsafe.Pointer(__cgo__s2))
	return__ = int(__cgo__return__)
	return
}

/*
Converts a string to a #gdouble value.

This function behaves like the standard strtod() function
does in the C locale. It does this without actually changing
the current locale, since that would not be thread-safe.
A limitation of the implementation is that this function
will still accept localized versions of infinities and NANs.

This function is typically used when reading configuration
files or other non-user input that should be locale independent.
To handle input from the user you should normally use the
locale-sensitive system strtod() function.

To convert from a #gdouble to a string in a locale-insensitive
way, use g_ascii_dtostr().

If the correct value would cause overflow, plus or minus %HUGE_VAL
is returned (according to the sign of the value), and %ERANGE is
stored in %errno. If the correct value would cause underflow,
zero is returned and %ERANGE is stored in %errno.

This function resets %errno before calling strtod() so that
you can reliably detect overflow and underflow.
*/
func AsciiStrtod(nptr string, endptr **C.gchar) (return__ float64) {
	__cgo__nptr := (*C.gchar)(unsafe.Pointer(C.CString(nptr)))
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.g_ascii_strtod(__cgo__nptr, endptr)
	C.free(unsafe.Pointer(__cgo__nptr))
	return__ = float64(__cgo__return__)
	return
}

/*
Converts a string to a #gint64 value.
This function behaves like the standard strtoll() function
does in the C locale. It does this without actually
changing the current locale, since that would not be
thread-safe.

This function is typically used when reading configuration
files or other non-user input that should be locale independent.
To handle input from the user you should normally use the
locale-sensitive system strtoll() function.

If the correct value would cause overflow, %G_MAXINT64 or %G_MININT64
is returned, and `ERANGE` is stored in `errno`.
If the base is outside the valid range, zero is returned, and
`EINVAL` is stored in `errno`. If the
string conversion fails, zero is returned, and @endptr returns @nptr
(if @endptr is non-%NULL).
*/
func AsciiStrtoll(nptr string, endptr **C.gchar, base uint) (return__ int64) {
	__cgo__nptr := (*C.gchar)(unsafe.Pointer(C.CString(nptr)))
	var __cgo__return__ C.gint64
	__cgo__return__ = C.g_ascii_strtoll(__cgo__nptr, endptr, C.guint(base))
	C.free(unsafe.Pointer(__cgo__nptr))
	return__ = int64(__cgo__return__)
	return
}

/*
Converts a string to a #guint64 value.
This function behaves like the standard strtoull() function
does in the C locale. It does this without actually
changing the current locale, since that would not be
thread-safe.

This function is typically used when reading configuration
files or other non-user input that should be locale independent.
To handle input from the user you should normally use the
locale-sensitive system strtoull() function.

If the correct value would cause overflow, %G_MAXUINT64
is returned, and `ERANGE` is stored in `errno`.
If the base is outside the valid range, zero is returned, and
`EINVAL` is stored in `errno`.
If the string conversion fails, zero is returned, and @endptr returns
@nptr (if @endptr is non-%NULL).
*/
func AsciiStrtoull(nptr string, endptr **C.gchar, base uint) (return__ uint64) {
	__cgo__nptr := (*C.gchar)(unsafe.Pointer(C.CString(nptr)))
	var __cgo__return__ C.guint64
	__cgo__return__ = C.g_ascii_strtoull(__cgo__nptr, endptr, C.guint(base))
	C.free(unsafe.Pointer(__cgo__nptr))
	return__ = uint64(__cgo__return__)
	return
}

/*
Converts all lower case ASCII letters to upper case ASCII letters.
*/
func AsciiStrup(str string, len_ int64) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_ascii_strup(__cgo__str, C.gssize(len_))
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Convert a character to ASCII lower case.

Unlike the standard C library tolower() function, this only
recognizes standard ASCII letters and ignores the locale, returning
all non-ASCII characters unchanged, even if they are lower case
letters in a particular character set. Also unlike the standard
library function, this takes and returns a char, not an int, so
don't call it on %EOF but no need to worry about casting to #guchar
before passing a possibly non-ASCII character in.
*/
func AsciiTolower(c byte) (return__ byte) {
	var __cgo__return__ C.gchar
	__cgo__return__ = C.g_ascii_tolower(C.gchar(c))
	return__ = byte(__cgo__return__)
	return
}

/*
Convert a character to ASCII upper case.

Unlike the standard C library toupper() function, this only
recognizes standard ASCII letters and ignores the locale, returning
all non-ASCII characters unchanged, even if they are upper case
letters in a particular character set. Also unlike the standard
library function, this takes and returns a char, not an int, so
don't call it on %EOF but no need to worry about casting to #guchar
before passing a possibly non-ASCII character in.
*/
func AsciiToupper(c byte) (return__ byte) {
	var __cgo__return__ C.gchar
	__cgo__return__ = C.g_ascii_toupper(C.gchar(c))
	return__ = byte(__cgo__return__)
	return
}

/*
Determines the numeric value of a character as a hexidecimal
digit. Differs from g_unichar_xdigit_value() because it takes
a char, so there's no worry about sign extension if characters
are signed.
*/
func AsciiXdigitValue(c byte) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_ascii_xdigit_value(C.gchar(c))
	return__ = int(__cgo__return__)
	return
}

// g_assert_warning is not generated due to explicit ignore

func AssertionMessage(domain string, file string, line int, func_ string, message string) {
	__cgo__domain := C.CString(domain)
	__cgo__file := C.CString(file)
	__cgo__func_ := C.CString(func_)
	__cgo__message := C.CString(message)
	C.g_assertion_message(__cgo__domain, __cgo__file, C.int(line), __cgo__func_, __cgo__message)
	C.free(unsafe.Pointer(__cgo__domain))
	C.free(unsafe.Pointer(__cgo__file))
	C.free(unsafe.Pointer(__cgo__func_))
	C.free(unsafe.Pointer(__cgo__message))
	return
}

// g_assertion_message_cmpnum is not generated due to long double param

func AssertionMessageCmpstr(domain string, file string, line int, func_ string, expr string, arg1 string, cmp string, arg2 string) {
	__cgo__domain := C.CString(domain)
	__cgo__file := C.CString(file)
	__cgo__func_ := C.CString(func_)
	__cgo__expr := C.CString(expr)
	__cgo__arg1 := C.CString(arg1)
	__cgo__cmp := C.CString(cmp)
	__cgo__arg2 := C.CString(arg2)
	C.g_assertion_message_cmpstr(__cgo__domain, __cgo__file, C.int(line), __cgo__func_, __cgo__expr, __cgo__arg1, __cgo__cmp, __cgo__arg2)
	C.free(unsafe.Pointer(__cgo__domain))
	C.free(unsafe.Pointer(__cgo__file))
	C.free(unsafe.Pointer(__cgo__func_))
	C.free(unsafe.Pointer(__cgo__expr))
	C.free(unsafe.Pointer(__cgo__arg1))
	C.free(unsafe.Pointer(__cgo__cmp))
	C.free(unsafe.Pointer(__cgo__arg2))
	return
}

func AssertionMessageError(domain string, file string, line int, func_ string, expr string, error_ *Error, error_domain C.GQuark, error_code int) {
	__cgo__domain := C.CString(domain)
	__cgo__file := C.CString(file)
	__cgo__func_ := C.CString(func_)
	__cgo__expr := C.CString(expr)
	C.g_assertion_message_error(__cgo__domain, __cgo__file, C.int(line), __cgo__func_, __cgo__expr, (*C.GError)(unsafe.Pointer(error_)), error_domain, C.int(error_code))
	C.free(unsafe.Pointer(__cgo__domain))
	C.free(unsafe.Pointer(__cgo__file))
	C.free(unsafe.Pointer(__cgo__func_))
	C.free(unsafe.Pointer(__cgo__expr))
	return
}

func AssertionMessageExpr(domain string, file string, line int, func_ string, expr string) {
	__cgo__domain := C.CString(domain)
	__cgo__file := C.CString(file)
	__cgo__func_ := C.CString(func_)
	__cgo__expr := C.CString(expr)
	C.g_assertion_message_expr(__cgo__domain, __cgo__file, C.int(line), __cgo__func_, __cgo__expr)
	C.free(unsafe.Pointer(__cgo__domain))
	C.free(unsafe.Pointer(__cgo__file))
	C.free(unsafe.Pointer(__cgo__func_))
	C.free(unsafe.Pointer(__cgo__expr))
	return
}

// g_atexit is not generated due to deprecation attr

// g_atomic_int_add is not generated due to explicit ignore

// g_atomic_int_and is not generated due to explicit ignore

// g_atomic_int_compare_and_exchange is not generated due to explicit ignore

// g_atomic_int_dec_and_test is not generated due to explicit ignore

// g_atomic_int_exchange_and_add is not generated due to deprecation attr

// g_atomic_int_get is not generated due to explicit ignore

// g_atomic_int_inc is not generated due to explicit ignore

// g_atomic_int_or is not generated due to explicit ignore

// g_atomic_int_set is not generated due to explicit ignore

// g_atomic_int_xor is not generated due to explicit ignore

// g_atomic_pointer_add is not generated due to explicit ignore

// g_atomic_pointer_and is not generated due to explicit ignore

// g_atomic_pointer_compare_and_exchange is not generated due to explicit ignore

// g_atomic_pointer_get is not generated due to explicit ignore

// g_atomic_pointer_or is not generated due to explicit ignore

// g_atomic_pointer_set is not generated due to explicit ignore

// g_atomic_pointer_xor is not generated due to explicit ignore

/*
Decode a sequence of Base-64 encoded text into binary data.  Note
that the returned binary data is not necessarily zero-terminated,
so it should not be used as a character string.
*/
func Base64Decode(text string) (out_len int64, return__ []byte) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	var __cgo__out_len C.gsize
	var __cgo__return__ *C.guchar
	__cgo__return__ = C.g_base64_decode(__cgo__text, &__cgo__out_len)
	C.free(unsafe.Pointer(__cgo__text))
	out_len = int64(__cgo__out_len)
	defer func() { return__ = C.GoBytes(unsafe.Pointer(__cgo__return__), C.int(out_len)) }()
	return
}

// g_base64_decode_inplace is not generated due to inout param

// g_base64_decode_step is not generated due to inout param

/*
Encode a sequence of binary data into its Base-64 stringified
representation.
*/
func Base64Encode(data []byte, len_ int64) (return__ string) {
	__header__data := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_base64_encode((*C.guchar)(unsafe.Pointer(__header__data.Data)), C.gsize(len_))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

// g_base64_encode_close is not generated due to inout param

// g_base64_encode_step is not generated due to inout param

// g_basename is not generated due to deprecation attr

/*
Sets the indicated @lock_bit in @address.  If the bit is already
set, this call will block until g_bit_unlock() unsets the
corresponding bit.

Attempting to lock on two different bits within the same integer is
not supported and will very probably cause deadlocks.

The value of the bit that is set is (1u << @bit).  If @bit is not
between 0 and 31 then the result is undefined.

This function accesses @address atomically.  All other accesses to
@address must be atomic in order for this function to work
reliably.
*/
func BitLock(address *C.gint, lock_bit int) {
	C.g_bit_lock(address, C.gint(lock_bit))
	return
}

/*
Find the position of the first bit set in @mask, searching
from (but not including) @nth_bit upwards. Bits are numbered
from 0 (least significant) to sizeof(#gulong) * 8 - 1 (31 or 63,
usually). To start searching from the 0th bit, set @nth_bit to -1.
*/
func BitNthLsf(mask int64, nth_bit int) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_bit_nth_lsf(C.gulong(mask), C.gint(nth_bit))
	return__ = int(__cgo__return__)
	return
}

/*
Find the position of the first bit set in @mask, searching
from (but not including) @nth_bit downwards. Bits are numbered
from 0 (least significant) to sizeof(#gulong) * 8 - 1 (31 or 63,
usually). To start searching from the last bit, set @nth_bit to
-1 or GLIB_SIZEOF_LONG * 8.
*/
func BitNthMsf(mask int64, nth_bit int) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_bit_nth_msf(C.gulong(mask), C.gint(nth_bit))
	return__ = int(__cgo__return__)
	return
}

/*
Gets the number of bits used to hold @number,
e.g. if @number is 4, 3 bits are needed.
*/
func BitStorage(number int64) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_bit_storage(C.gulong(number))
	return__ = uint(__cgo__return__)
	return
}

/*
Sets the indicated @lock_bit in @address, returning %TRUE if
successful.  If the bit is already set, returns %FALSE immediately.

Attempting to lock on two different bits within the same integer is
not supported.

The value of the bit that is set is (1u << @bit).  If @bit is not
between 0 and 31 then the result is undefined.

This function accesses @address atomically.  All other accesses to
@address must be atomic in order for this function to work
reliably.
*/
func BitTrylock(address *C.gint, lock_bit int) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_bit_trylock(address, C.gint(lock_bit))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Clears the indicated @lock_bit in @address.  If another thread is
currently blocked in g_bit_lock() on this same bit then it will be
woken up.

This function accesses @address atomically.  All other accesses to
@address must be atomic in order for this function to work
reliably.
*/
func BitUnlock(address *C.gint, lock_bit int) {
	C.g_bit_unlock(address, C.gint(lock_bit))
	return
}

func BookmarkFileErrorQuark() (return__ C.GQuark) {
	return__ = C.g_bookmark_file_error_quark()
	return
}

// g_build_filename is not generated due to varargs

/*
Behaves exactly like g_build_filename(), but takes the path elements
as a string array, instead of varargs. This function is mainly
meant for language bindings.
*/
func BuildFilenamev(args []string) (return__ string) {
	__header__args := (*reflect.SliceHeader)(unsafe.Pointer(&args))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_build_filenamev((**C.gchar)(unsafe.Pointer(__header__args.Data)))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

// g_build_path is not generated due to varargs

/*
Behaves exactly like g_build_path(), but takes the path elements
as a string array, instead of varargs. This function is mainly
meant for language bindings.
*/
func BuildPathv(separator string, args []string) (return__ string) {
	__cgo__separator := (*C.gchar)(unsafe.Pointer(C.CString(separator)))
	__header__args := (*reflect.SliceHeader)(unsafe.Pointer(&args))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_build_pathv(__cgo__separator, (**C.gchar)(unsafe.Pointer(__header__args.Data)))
	C.free(unsafe.Pointer(__cgo__separator))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Frees the memory allocated by the #GByteArray. If @free_segment is
%TRUE it frees the actual byte data. If the reference count of
@array is greater than one, the #GByteArray wrapper is preserved but
the size of @array will be set to zero.
*/
func ByteArrayFree(array *ByteArray, free_segment bool) (return__ *C.guint8) {
	__cgo__free_segment := C.gboolean(0)
	if free_segment {
		__cgo__free_segment = C.gboolean(1)
	}
	return__ = C.g_byte_array_free((*C.GByteArray)(unsafe.Pointer(array)), __cgo__free_segment)
	return
}

/*
Transfers the data from the #GByteArray into a new immutable #GBytes.

The #GByteArray is freed unless the reference count of @array is greater
than one, the #GByteArray wrapper is preserved but the size of @array
will be set to zero.

This is identical to using g_bytes_new_take() and g_byte_array_free()
together.
*/
func ByteArrayFreeToBytes(array *ByteArray) (return__ *Bytes) {
	var __cgo__return__ *C.GBytes
	__cgo__return__ = C.g_byte_array_free_to_bytes((*C.GByteArray)(unsafe.Pointer(array)))
	return__ = (*Bytes)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Creates a new #GByteArray with a reference count of 1.
*/
func ByteArrayNew() (return__ *ByteArray) {
	var __cgo__return__ *C.GByteArray
	__cgo__return__ = C.g_byte_array_new()
	return__ = (*ByteArray)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Create byte array containing the data. The data will be owned by the array
and will be freed with g_free(), i.e. it could be allocated using g_strdup().
*/
func ByteArrayNewTake(data []byte, len_ int64) (return__ *ByteArray) {
	__header__data := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	var __cgo__return__ *C.GByteArray
	__cgo__return__ = C.g_byte_array_new_take((*C.guint8)(unsafe.Pointer(__header__data.Data)), C.gsize(len_))
	return__ = (*ByteArray)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Atomically decrements the reference count of @array by one. If the
reference count drops to 0, all memory allocated by the array is
released. This function is thread-safe and may be called from any
thread.
*/
func ByteArrayUnref(array *ByteArray) {
	C.g_byte_array_unref((*C.GByteArray)(unsafe.Pointer(array)))
	return
}

/*
A wrapper for the POSIX chdir() function. The function changes the
current directory of the process to @path.

See your C library manual for more details about chdir().
*/
func Chdir(path string) (return__ int) {
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	var __cgo__return__ C.int
	__cgo__return__ = C.g_chdir(__cgo__path)
	C.free(unsafe.Pointer(__cgo__path))
	return__ = int(__cgo__return__)
	return
}

/*
Checks that the GLib library in use is compatible with the
given version. Generally you would pass in the constants
#GLIB_MAJOR_VERSION, #GLIB_MINOR_VERSION, #GLIB_MICRO_VERSION
as the three arguments to this function; that produces
a check that the library in use is compatible with
the version of GLib the application or module was compiled
against.

Compatibility is defined by two things: first the version
of the running library is newer than the version
@required_major.required_minor.@required_micro. Second
the running library must be binary compatible with the
version @required_major.required_minor.@required_micro
(same major version.)
*/
func CheckVersion(required_major uint, required_minor uint, required_micro uint) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.glib_check_version(C.guint(required_major), C.guint(required_minor), C.guint(required_micro))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the length in bytes of digests of type @checksum_type
*/
func ChecksumTypeGetLength(checksum_type C.GChecksumType) (return__ int64) {
	var __cgo__return__ C.gssize
	__cgo__return__ = C.g_checksum_type_get_length(checksum_type)
	return__ = int64(__cgo__return__)
	return
}

/*
Sets a function to be called when the child indicated by @pid
exits, at a default priority, #G_PRIORITY_DEFAULT.

If you obtain @pid from g_spawn_async() or g_spawn_async_with_pipes()
you will need to pass #G_SPAWN_DO_NOT_REAP_CHILD as flag to
the spawn function for the child watching to work.

Note that on platforms where #GPid must be explicitly closed
(see g_spawn_close_pid()) @pid must not be closed while the
source is still active. Typically, you will want to call
g_spawn_close_pid() in the callback function for the source.

GLib supports only a single callback per process id.

This internally creates a main loop source using
g_child_watch_source_new() and attaches it to the main loop context
using g_source_attach(). You can do these steps manually if you
need greater control.
*/
func ChildWatchAdd(pid C.GPid, function C.GChildWatchFunc, data unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_child_watch_add(pid, function, (C.gpointer)(data))
	return__ = uint(__cgo__return__)
	return
}

/*
Sets a function to be called when the child indicated by @pid
exits, at the priority @priority.

If you obtain @pid from g_spawn_async() or g_spawn_async_with_pipes()
you will need to pass #G_SPAWN_DO_NOT_REAP_CHILD as flag to
the spawn function for the child watching to work.

In many programs, you will want to call g_spawn_check_exit_status()
in the callback to determine whether or not the child exited
successfully.

Also, note that on platforms where #GPid must be explicitly closed
(see g_spawn_close_pid()) @pid must not be closed while the source
is still active.  Typically, you should invoke g_spawn_close_pid()
in the callback function for the source.

GLib supports only a single callback per process id.

This internally creates a main loop source using
g_child_watch_source_new() and attaches it to the main loop context
using g_source_attach(). You can do these steps manually if you
need greater control.
*/
func ChildWatchAddFull(priority int, pid C.GPid, function C.GChildWatchFunc, data unsafe.Pointer, notify C.GDestroyNotify) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_child_watch_add_full(C.gint(priority), pid, function, (C.gpointer)(data), notify)
	return__ = uint(__cgo__return__)
	return
}

/*
Creates a new child_watch source.

The source will not initially be associated with any #GMainContext
and must be added to one with g_source_attach() before it will be
executed.

Note that child watch sources can only be used in conjunction with
`g_spawn...` when the %G_SPAWN_DO_NOT_REAP_CHILD flag is used.

Note that on platforms where #GPid must be explicitly closed
(see g_spawn_close_pid()) @pid must not be closed while the
source is still active. Typically, you will want to call
g_spawn_close_pid() in the callback function for the source.

Note further that using g_child_watch_source_new() is not
compatible with calling `waitpid` with a nonpositive first
argument in the application. Calling waitpid() for individual
pids will still work fine.

Similarly, on POSIX platforms, the @pid passed to this function must
be greater than 0 (i.e. this function must wait for a specific child,
and cannot wait for one of many children by using a nonpositive argument).
*/
func ChildWatchSourceNew(pid C.GPid) (return__ *Source) {
	var __cgo__return__ *C.GSource
	__cgo__return__ = C.g_child_watch_source_new(pid)
	return__ = (*Source)(unsafe.Pointer(__cgo__return__))
	return
}

/*
If @err is %NULL, does nothing. If @err is non-%NULL,
calls g_error_free() on *@err and sets *@err to %NULL.
*/
func ClearError() (__err__ error) {
	var __cgo_error__ *C.GError
	C.g_clear_error(&__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

// g_clear_pointer is not generated due to explicit ignore

/*
This wraps the close() call; in case of error, %errno will be
preserved, but the error will also be stored as a #GError in @error.

Besides using #GError, there is another major reason to prefer this
function over the call provided by the system; on Unix, it will
attempt to correctly handle %EINTR, which has platform-specific
semantics.
*/
func Close(fd int) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_close(C.gint(fd), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Computes the checksum for a binary @data. This is a
convenience wrapper for g_checksum_new(), g_checksum_get_string()
and g_checksum_free().

The hexadecimal string returned will be in lower case.
*/
func ComputeChecksumForBytes(checksum_type C.GChecksumType, data *Bytes) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_compute_checksum_for_bytes(checksum_type, (*C.GBytes)(unsafe.Pointer(data)))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Computes the checksum for a binary @data of @length. This is a
convenience wrapper for g_checksum_new(), g_checksum_get_string()
and g_checksum_free().

The hexadecimal string returned will be in lower case.
*/
func ComputeChecksumForData(checksum_type C.GChecksumType, data []byte, length int64) (return__ string) {
	__header__data := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_compute_checksum_for_data(checksum_type, (*C.guchar)(unsafe.Pointer(__header__data.Data)), C.gsize(length))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Computes the checksum of a string.

The hexadecimal string returned will be in lower case.
*/
func ComputeChecksumForString(checksum_type C.GChecksumType, str string, length int64) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_compute_checksum_for_string(checksum_type, __cgo__str, C.gssize(length))
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Computes the HMAC for a binary @data of @length. This is a
convenience wrapper for g_hmac_new(), g_hmac_get_string()
and g_hmac_unref().

The hexadecimal string returned will be in lower case.
*/
func ComputeHmacForData(digest_type C.GChecksumType, key []byte, key_len int64, data *C.guchar, length int64) (return__ string) {
	__header__key := (*reflect.SliceHeader)(unsafe.Pointer(&key))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_compute_hmac_for_data(digest_type, (*C.guchar)(unsafe.Pointer(__header__key.Data)), C.gsize(key_len), data, C.gsize(length))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Computes the HMAC for a string.

The hexadecimal string returned will be in lower case.
*/
func ComputeHmacForString(digest_type C.GChecksumType, key []byte, key_len int64, str string, length int64) (return__ string) {
	__header__key := (*reflect.SliceHeader)(unsafe.Pointer(&key))
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_compute_hmac_for_string(digest_type, (*C.guchar)(unsafe.Pointer(__header__key.Data)), C.gsize(key_len), __cgo__str, C.gssize(length))
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts a string from one character set to another.

Note that you should use g_iconv() for streaming conversions.
Despite the fact that @byes_read can return information about partial
characters, the g_convert_... functions are not generally suitable
for streaming. If the underlying converter maintains internal state,
then this won't be preserved across successive calls to g_convert(),
g_convert_with_iconv() or g_convert_with_fallback(). (An example of
this is the GNU C converter for CP1255 which does not emit a base
character until it knows that the next character is not a mark that
could combine with the base character.)

Using extensions such as "//TRANSLIT" may not work (or may not work
well) on many platforms.  Consider using g_str_to_ascii() instead.
*/
func Convert(str string, len_ int64, to_codeset string, from_codeset string) (bytes_read int64, bytes_written int64, return__ string, __err__ error) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	__cgo__to_codeset := (*C.gchar)(unsafe.Pointer(C.CString(to_codeset)))
	__cgo__from_codeset := (*C.gchar)(unsafe.Pointer(C.CString(from_codeset)))
	var __cgo__bytes_read C.gsize
	var __cgo__bytes_written C.gsize
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_convert(__cgo__str, C.gssize(len_), __cgo__to_codeset, __cgo__from_codeset, &__cgo__bytes_read, &__cgo__bytes_written, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__str))
	C.free(unsafe.Pointer(__cgo__to_codeset))
	C.free(unsafe.Pointer(__cgo__from_codeset))
	bytes_read = int64(__cgo__bytes_read)
	bytes_written = int64(__cgo__bytes_written)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

func ConvertErrorQuark() (return__ C.GQuark) {
	return__ = C.g_convert_error_quark()
	return
}

/*
Converts a string from one character set to another, possibly
including fallback sequences for characters not representable
in the output. Note that it is not guaranteed that the specification
for the fallback sequences in @fallback will be honored. Some
systems may do an approximate conversion from @from_codeset
to @to_codeset in their iconv() functions,
in which case GLib will simply return that approximate conversion.

Note that you should use g_iconv() for streaming conversions.
Despite the fact that @byes_read can return information about partial
characters, the g_convert_... functions are not generally suitable
for streaming. If the underlying converter maintains internal state,
then this won't be preserved across successive calls to g_convert(),
g_convert_with_iconv() or g_convert_with_fallback(). (An example of
this is the GNU C converter for CP1255 which does not emit a base
character until it knows that the next character is not a mark that
could combine with the base character.)
*/
func ConvertWithFallback(str string, len_ int64, to_codeset string, from_codeset string, fallback string, bytes_read *C.gsize, bytes_written *C.gsize) (return__ string, __err__ error) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	__cgo__to_codeset := (*C.gchar)(unsafe.Pointer(C.CString(to_codeset)))
	__cgo__from_codeset := (*C.gchar)(unsafe.Pointer(C.CString(from_codeset)))
	__cgo__fallback := (*C.gchar)(unsafe.Pointer(C.CString(fallback)))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_convert_with_fallback(__cgo__str, C.gssize(len_), __cgo__to_codeset, __cgo__from_codeset, __cgo__fallback, bytes_read, bytes_written, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__str))
	C.free(unsafe.Pointer(__cgo__to_codeset))
	C.free(unsafe.Pointer(__cgo__from_codeset))
	C.free(unsafe.Pointer(__cgo__fallback))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Converts a string from one character set to another.

Note that you should use g_iconv() for streaming conversions.
Despite the fact that @byes_read can return information about partial
characters, the g_convert_... functions are not generally suitable
for streaming. If the underlying converter maintains internal state,
then this won't be preserved across successive calls to g_convert(),
g_convert_with_iconv() or g_convert_with_fallback(). (An example of
this is the GNU C converter for CP1255 which does not emit a base
character until it knows that the next character is not a mark that
could combine with the base character.)
*/
func ConvertWithIconv(str string, len_ int64, converter C.GIConv, bytes_read *C.gsize, bytes_written *C.gsize) (return__ string, __err__ error) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_convert_with_iconv(__cgo__str, C.gssize(len_), converter, bytes_read, bytes_written, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Frees all the data elements of the datalist.
The data elements' destroy functions are called
if they have been set.
*/
func DatalistClear(datalist **C.GData) {
	C.g_datalist_clear(datalist)
	return
}

/*
Calls the given function for each data element of the datalist. The
function is called with each data element's #GQuark id and data,
together with the given @user_data parameter. Note that this
function is NOT thread-safe. So unless @datalist can be protected
from any modifications during invocation of this function, it should
not be called.
*/
func DatalistForeach(datalist **C.GData, func_ C.GDataForeachFunc, user_data unsafe.Pointer) {
	C.g_datalist_foreach(datalist, func_, (C.gpointer)(user_data))
	return
}

/*
Gets a data element, using its string identifier. This is slower than
g_datalist_id_get_data() because it compares strings.
*/
func DatalistGetData(datalist **C.GData, key string) (return__ unsafe.Pointer) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_datalist_get_data(datalist, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Gets flags values packed in together with the datalist.
See g_datalist_set_flags().
*/
func DatalistGetFlags(datalist **C.GData) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_datalist_get_flags(datalist)
	return__ = uint(__cgo__return__)
	return
}

/*
This is a variant of g_datalist_id_get_data() which
returns a 'duplicate' of the value. @dup_func defines the
meaning of 'duplicate' in this context, it could e.g.
take a reference on a ref-counted object.

If the @key_id is not set in the datalist then @dup_func
will be called with a %NULL argument.

Note that @dup_func is called while the datalist is locked, so it
is not allowed to read or modify the datalist.

This function can be useful to avoid races when multiple
threads are using the same datalist and the same key.
*/
func DatalistIdDupData(datalist **C.GData, key_id C.GQuark, dup_func C.GDuplicateFunc, user_data unsafe.Pointer) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_datalist_id_dup_data(datalist, key_id, dup_func, (C.gpointer)(user_data))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Retrieves the data element corresponding to @key_id.
*/
func DatalistIdGetData(datalist **C.GData, key_id C.GQuark) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_datalist_id_get_data(datalist, key_id)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Removes an element, without calling its destroy notification
function.
*/
func DatalistIdRemoveNoNotify(datalist **C.GData, key_id C.GQuark) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_datalist_id_remove_no_notify(datalist, key_id)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Compares the member that is associated with @key_id in
@datalist to @oldval, and if they are the same, replace
@oldval with @newval.

This is like a typical atomic compare-and-exchange
operation, for a member of @datalist.

If the previous value was replaced then ownership of the
old value (@oldval) is passed to the caller, including
the registred destroy notify for it (passed out in @old_destroy).
Its up to the caller to free this as he wishes, which may
or may not include using @old_destroy as sometimes replacement
should not destroy the object in the normal way.
*/
func DatalistIdReplaceData(datalist **C.GData, key_id C.GQuark, oldval unsafe.Pointer, newval unsafe.Pointer, destroy C.GDestroyNotify, old_destroy *C.GDestroyNotify) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_datalist_id_replace_data(datalist, key_id, (C.gpointer)(oldval), (C.gpointer)(newval), destroy, old_destroy)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the data corresponding to the given #GQuark id, and the
function to be called when the element is removed from the datalist.
Any previous data with the same key is removed, and its destroy
function is called.
*/
func DatalistIdSetDataFull(datalist **C.GData, key_id C.GQuark, data unsafe.Pointer, destroy_func C.GDestroyNotify) {
	C.g_datalist_id_set_data_full(datalist, key_id, (C.gpointer)(data), destroy_func)
	return
}

/*
Resets the datalist to %NULL. It does not free any memory or call
any destroy functions.
*/
func DatalistInit(datalist **C.GData) {
	C.g_datalist_init(datalist)
	return
}

/*
Turns on flag values for a data list. This function is used
to keep a small number of boolean flags in an object with
a data list without using any additional space. It is
not generally useful except in circumstances where space
is very tight. (It is used in the base #GObject type, for
example.)
*/
func DatalistSetFlags(datalist **C.GData, flags uint) {
	C.g_datalist_set_flags(datalist, C.guint(flags))
	return
}

/*
Turns off flag values for a data list. See g_datalist_unset_flags()
*/
func DatalistUnsetFlags(datalist **C.GData, flags uint) {
	C.g_datalist_unset_flags(datalist, C.guint(flags))
	return
}

/*
Destroys the dataset, freeing all memory allocated, and calling any
destroy functions set for data elements.
*/
func DatasetDestroy(dataset_location unsafe.Pointer) {
	C.g_dataset_destroy((C.gconstpointer)(dataset_location))
	return
}

/*
Calls the given function for each data element which is associated
with the given location. Note that this function is NOT thread-safe.
So unless @datalist can be protected from any modifications during
invocation of this function, it should not be called.
*/
func DatasetForeach(dataset_location unsafe.Pointer, func_ C.GDataForeachFunc, user_data unsafe.Pointer) {
	C.g_dataset_foreach((C.gconstpointer)(dataset_location), func_, (C.gpointer)(user_data))
	return
}

/*
Gets the data element corresponding to a #GQuark.
*/
func DatasetIdGetData(dataset_location unsafe.Pointer, key_id C.GQuark) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_dataset_id_get_data((C.gconstpointer)(dataset_location), key_id)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Removes an element, without calling its destroy notification
function.
*/
func DatasetIdRemoveNoNotify(dataset_location unsafe.Pointer, key_id C.GQuark) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_dataset_id_remove_no_notify((C.gconstpointer)(dataset_location), key_id)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Sets the data element associated with the given #GQuark id, and also
the function to call when the data element is destroyed. Any
previous data with the same key is removed, and its destroy function
is called.
*/
func DatasetIdSetDataFull(dataset_location unsafe.Pointer, key_id C.GQuark, data unsafe.Pointer, destroy_func C.GDestroyNotify) {
	C.g_dataset_id_set_data_full((C.gconstpointer)(dataset_location), key_id, (C.gpointer)(data), destroy_func)
	return
}

/*
Returns the number of days in a month, taking leap
years into account.
*/
func DateGetDaysInMonth(month C.GDateMonth, year C.GDateYear) (return__ uint8) {
	var __cgo__return__ C.guint8
	__cgo__return__ = C.g_date_get_days_in_month(month, year)
	return__ = uint8(__cgo__return__)
	return
}

/*
Returns the number of weeks in the year, where weeks
are taken to start on Monday. Will be 52 or 53. The
date must be valid. (Years always have 52 7-day periods,
plus 1 or 2 extra days depending on whether it's a leap
year. This function is basically telling you how many
Mondays are in the year, i.e. there are 53 Mondays if
one of the extra days happens to be a Monday.)
*/
func DateGetMondayWeeksInYear(year C.GDateYear) (return__ uint8) {
	var __cgo__return__ C.guint8
	__cgo__return__ = C.g_date_get_monday_weeks_in_year(year)
	return__ = uint8(__cgo__return__)
	return
}

/*
Returns the number of weeks in the year, where weeks
are taken to start on Sunday. Will be 52 or 53. The
date must be valid. (Years always have 52 7-day periods,
plus 1 or 2 extra days depending on whether it's a leap
year. This function is basically telling you how many
Sundays are in the year, i.e. there are 53 Sundays if
one of the extra days happens to be a Sunday.)
*/
func DateGetSundayWeeksInYear(year C.GDateYear) (return__ uint8) {
	var __cgo__return__ C.guint8
	__cgo__return__ = C.g_date_get_sunday_weeks_in_year(year)
	return__ = uint8(__cgo__return__)
	return
}

/*
Returns %TRUE if the year is a leap year.

For the purposes of this function, leap year is every year
divisible by 4 unless that year is divisible by 100. If it
is divisible by 100 it would be a leap year only if that year
is also divisible by 400.
*/
func DateIsLeapYear(year C.GDateYear) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_date_is_leap_year(year)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Generates a printed representation of the date, in a
[locale][setlocale]-specific way.
Works just like the platform's C library strftime() function,
but only accepts date-related formats; time-related formats
give undefined results. Date must be valid. Unlike strftime()
(which uses the locale encoding), works on a UTF-8 format
string and stores a UTF-8 result.

This function does not provide any conversion specifiers in
addition to those implemented by the platform's C library.
For example, don't expect that using g_date_strftime() would
make the \%F provided by the C99 strftime() work on Windows
where the C library only complies to C89.
*/
func DateStrftime(s string, slen int64, format string, date *Date) (return__ int64) {
	__cgo__s := (*C.gchar)(unsafe.Pointer(C.CString(s)))
	__cgo__format := (*C.gchar)(unsafe.Pointer(C.CString(format)))
	var __cgo__return__ C.gsize
	__cgo__return__ = C.g_date_strftime(__cgo__s, C.gsize(slen), __cgo__format, (*C.GDate)(unsafe.Pointer(date)))
	C.free(unsafe.Pointer(__cgo__s))
	C.free(unsafe.Pointer(__cgo__format))
	return__ = int64(__cgo__return__)
	return
}

/*
A comparison function for #GDateTimes that is suitable
as a #GCompareFunc. Both #GDateTimes must be non-%NULL.
*/
func DateTimeCompare(dt1 unsafe.Pointer, dt2 unsafe.Pointer) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_date_time_compare((C.gconstpointer)(dt1), (C.gconstpointer)(dt2))
	return__ = int(__cgo__return__)
	return
}

/*
Checks to see if @dt1 and @dt2 are equal.

Equal here means that they represent the same moment after converting
them to the same time zone.
*/
func DateTimeEqual(dt1 unsafe.Pointer, dt2 unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_date_time_equal((C.gconstpointer)(dt1), (C.gconstpointer)(dt2))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Hashes @datetime into a #guint, suitable for use within #GHashTable.
*/
func DateTimeHash(datetime unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_date_time_hash((C.gconstpointer)(datetime))
	return__ = uint(__cgo__return__)
	return
}

/*
Returns %TRUE if the day of the month is valid (a day is valid if it's
between 1 and 31 inclusive).
*/
func DateValidDay(day C.GDateDay) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_date_valid_day(day)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if the day-month-year triplet forms a valid, existing day
in the range of days #GDate understands (Year 1 or later, no more than
a few thousand years in the future).
*/
func DateValidDmy(day C.GDateDay, month C.GDateMonth, year C.GDateYear) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_date_valid_dmy(day, month, year)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if the Julian day is valid. Anything greater than zero
is basically a valid Julian, though there is a 32-bit limit.
*/
func DateValidJulian(julian_date uint32) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_date_valid_julian(C.guint32(julian_date))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if the month value is valid. The 12 #GDateMonth
enumeration values are the only valid months.
*/
func DateValidMonth(month C.GDateMonth) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_date_valid_month(month)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if the weekday is valid. The seven #GDateWeekday enumeration
values are the only valid weekdays.
*/
func DateValidWeekday(weekday C.GDateWeekday) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_date_valid_weekday(weekday)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if the year is valid. Any year greater than 0 is valid,
though there is a 16-bit limit to what #GDate will understand.
*/
func DateValidYear(year C.GDateYear) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_date_valid_year(year)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This is a variant of g_dgettext() that allows specifying a locale
category instead of always using `LC_MESSAGES`. See g_dgettext() for
more information about how this functions differs from calling
dcgettext() directly.
*/
func Dcgettext(domain string, msgid string, category int) (return__ string) {
	__cgo__domain := (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	__cgo__msgid := (*C.gchar)(unsafe.Pointer(C.CString(msgid)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dcgettext(__cgo__domain, __cgo__msgid, C.gint(category))
	C.free(unsafe.Pointer(__cgo__domain))
	C.free(unsafe.Pointer(__cgo__msgid))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
This function is a wrapper of dgettext() which does not translate
the message if the default domain as set with textdomain() has no
translations for the current locale.

The advantage of using this function over dgettext() proper is that
libraries using this function (like GTK+) will not use translations
if the application using the library does not have translations for
the current locale.  This results in a consistent English-only
interface instead of one having partial translations.  For this
feature to work, the call to textdomain() and setlocale() should
precede any g_dgettext() invocations.  For GTK+, it means calling
textdomain() before gtk_init or its variants.

This function disables translations if and only if upon its first
call all the following conditions hold:

- @domain is not %NULL

- textdomain() has been called to set a default text domain

- there is no translations available for the default text domain
  and the current locale

- current locale is not "C" or any English locales (those
  starting with "en_")

Note that this behavior may not be desired for example if an application
has its untranslated messages in a language other than English. In those
cases the application should call textdomain() after initializing GTK+.

Applications should normally not use this function directly,
but use the _() macro for translations.
*/
func Dgettext(domain string, msgid string) (return__ string) {
	__cgo__domain := (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	__cgo__msgid := (*C.gchar)(unsafe.Pointer(C.CString(msgid)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dgettext(__cgo__domain, __cgo__msgid)
	C.free(unsafe.Pointer(__cgo__domain))
	C.free(unsafe.Pointer(__cgo__msgid))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Creates a subdirectory in the preferred directory for temporary
files (as returned by g_get_tmp_dir()).

@tmpl should be a string in the GLib file name encoding containing
a sequence of six 'X' characters, as the parameter to g_mkstemp().
However, unlike these functions, the template should only be a
basename, no directory components are allowed. If template is
%NULL, a default template is used.

Note that in contrast to g_mkdtemp() (and mkdtemp()) @tmpl is not
modified, and might thus be a read-only literal string.
*/
func DirMakeTmp(tmpl string) (return__ string, __err__ error) {
	__cgo__tmpl := (*C.gchar)(unsafe.Pointer(C.CString(tmpl)))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dir_make_tmp(__cgo__tmpl, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__tmpl))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Compares two #gpointer arguments and returns %TRUE if they are equal.
It can be passed to g_hash_table_new() as the @key_equal_func
parameter, when using opaque pointers compared by pointer value as
keys in a #GHashTable.

This equality function is also appropriate for keys that are integers
stored in pointers, such as `GINT_TO_POINTER (n)`.
*/
func DirectEqual(v1 unsafe.Pointer, v2 unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_direct_equal((C.gconstpointer)(v1), (C.gconstpointer)(v2))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Converts a gpointer to a hash value.
It can be passed to g_hash_table_new() as the @hash_func parameter,
when using opaque pointers compared by pointer value as keys in a
#GHashTable.

This hash function is also appropriate for keys that are integers
stored in pointers, such as `GINT_TO_POINTER (n)`.
*/
func DirectHash(v unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_direct_hash((C.gconstpointer)(v))
	return__ = uint(__cgo__return__)
	return
}

/*
This function is a wrapper of dngettext() which does not translate
the message if the default domain as set with textdomain() has no
translations for the current locale.

See g_dgettext() for details of how this differs from dngettext()
proper.
*/
func Dngettext(domain string, msgid string, msgid_plural string, n int64) (return__ string) {
	__cgo__domain := (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	__cgo__msgid := (*C.gchar)(unsafe.Pointer(C.CString(msgid)))
	__cgo__msgid_plural := (*C.gchar)(unsafe.Pointer(C.CString(msgid_plural)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dngettext(__cgo__domain, __cgo__msgid, __cgo__msgid_plural, C.gulong(n))
	C.free(unsafe.Pointer(__cgo__domain))
	C.free(unsafe.Pointer(__cgo__msgid))
	C.free(unsafe.Pointer(__cgo__msgid_plural))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Compares the two #gdouble values being pointed to and returns
%TRUE if they are equal.
It can be passed to g_hash_table_new() as the @key_equal_func
parameter, when using non-%NULL pointers to doubles as keys in a
#GHashTable.
*/
func DoubleEqual(v1 unsafe.Pointer, v2 unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_double_equal((C.gconstpointer)(v1), (C.gconstpointer)(v2))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Converts a pointer to a #gdouble to a hash value.
It can be passed to g_hash_table_new() as the @hash_func parameter,
It can be passed to g_hash_table_new() as the @hash_func parameter,
when using non-%NULL pointers to doubles as keys in a #GHashTable.
*/
func DoubleHash(v unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_double_hash((C.gconstpointer)(v))
	return__ = uint(__cgo__return__)
	return
}

/*
This function is a variant of g_dgettext() which supports
a disambiguating message context. GNU gettext uses the
'\004' character to separate the message context and
message id in @msgctxtid.
If 0 is passed as @msgidoffset, this function will fall back to
trying to use the deprecated convention of using "|" as a separation
character.

This uses g_dgettext() internally. See that functions for differences
with dgettext() proper.

Applications should normally not use this function directly,
but use the C_() macro for translations with context.
*/
func Dpgettext(domain string, msgctxtid string, msgidoffset int64) (return__ string) {
	__cgo__domain := (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	__cgo__msgctxtid := (*C.gchar)(unsafe.Pointer(C.CString(msgctxtid)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dpgettext(__cgo__domain, __cgo__msgctxtid, C.gsize(msgidoffset))
	C.free(unsafe.Pointer(__cgo__domain))
	C.free(unsafe.Pointer(__cgo__msgctxtid))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
This function is a variant of g_dgettext() which supports
a disambiguating message context. GNU gettext uses the
'\004' character to separate the message context and
message id in @msgctxtid.

This uses g_dgettext() internally. See that functions for differences
with dgettext() proper.

This function differs from C_() in that it is not a macro and
thus you may use non-string-literals as context and msgid arguments.
*/
func Dpgettext2(domain string, context string, msgid string) (return__ string) {
	__cgo__domain := (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	__cgo__context := (*C.gchar)(unsafe.Pointer(C.CString(context)))
	__cgo__msgid := (*C.gchar)(unsafe.Pointer(C.CString(msgid)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_dpgettext2(__cgo__domain, __cgo__context, __cgo__msgid)
	C.free(unsafe.Pointer(__cgo__domain))
	C.free(unsafe.Pointer(__cgo__context))
	C.free(unsafe.Pointer(__cgo__msgid))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the value of the environment variable @variable in the
provided list @envp.
*/
func EnvironGetenv(envp []string, variable string) (return__ string) {
	__header__envp := (*reflect.SliceHeader)(unsafe.Pointer(&envp))
	__cgo__variable := (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_environ_getenv((**C.gchar)(unsafe.Pointer(__header__envp.Data)), __cgo__variable)
	C.free(unsafe.Pointer(__cgo__variable))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Sets the environment variable @variable in the provided list
@envp to @value.
*/
func EnvironSetenv(envp []string, variable string, value string, overwrite bool) (return__ []string) {
	__header__envp := (*reflect.SliceHeader)(unsafe.Pointer(&envp))
	__cgo__variable := (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	__cgo__value := (*C.gchar)(unsafe.Pointer(C.CString(value)))
	__cgo__overwrite := C.gboolean(0)
	if overwrite {
		__cgo__overwrite = C.gboolean(1)
	}
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.g_environ_setenv((**C.gchar)(unsafe.Pointer(__header__envp.Data)), __cgo__variable, __cgo__value, __cgo__overwrite)
	C.free(unsafe.Pointer(__cgo__variable))
	C.free(unsafe.Pointer(__cgo__value))
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Removes the environment variable @variable from the provided
environment @envp.
*/
func EnvironUnsetenv(envp []string, variable string) (return__ []string) {
	__header__envp := (*reflect.SliceHeader)(unsafe.Pointer(&envp))
	__cgo__variable := (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.g_environ_unsetenv((**C.gchar)(unsafe.Pointer(__header__envp.Data)), __cgo__variable)
	C.free(unsafe.Pointer(__cgo__variable))
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Gets a #GFileError constant based on the passed-in @err_no.
For example, if you pass in `EEXIST` this function returns
#G_FILE_ERROR_EXIST. Unlike `errno` values, you can portably
assume that all #GFileError values will exist.

Normally a #GFileError value goes into a #GError returned
from a function that manipulates files. So you would use
g_file_error_from_errno() when constructing a #GError.
*/
func FileErrorFromErrno(err_no int) (return__ C.GFileError) {
	return__ = C.g_file_error_from_errno(C.gint(err_no))
	return
}

func FileErrorQuark() (return__ C.GQuark) {
	return__ = C.g_file_error_quark()
	return
}

/*
Reads an entire file into allocated memory, with good error
checking.

If the call was successful, it returns %TRUE and sets @contents to the file
contents and @length to the length of the file contents in bytes. The string
stored in @contents will be nul-terminated, so for text files you can pass
%NULL for the @length argument. If the call was not successful, it returns
%FALSE and sets @error. The error domain is #G_FILE_ERROR. Possible error
codes are those in the #GFileError enumeration. In the error case,
@contents is set to %NULL and @length is set to zero.
*/
func FileGetContents(filename string) (contents []byte, length int64, return__ bool, __err__ error) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	var __cgo__contents *C.gchar
	var __cgo__length C.gsize
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_get_contents(__cgo__filename, &__cgo__contents, &__cgo__length, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__filename))
	defer func() { contents = C.GoBytes(unsafe.Pointer(__cgo__contents), C.int(length)) }()
	length = int64(__cgo__length)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Opens a file for writing in the preferred directory for temporary
files (as returned by g_get_tmp_dir()).

@tmpl should be a string in the GLib file name encoding containing
a sequence of six 'X' characters, as the parameter to g_mkstemp().
However, unlike these functions, the template should only be a
basename, no directory components are allowed. If template is
%NULL, a default template is used.

Note that in contrast to g_mkstemp() (and mkstemp()) @tmpl is not
modified, and might thus be a read-only literal string.

Upon success, and if @name_used is non-%NULL, the actual name used
is returned in @name_used. This string should be freed with g_free()
when not needed any longer. The returned name is in the GLib file
name encoding.
*/
func FileOpenTmp(tmpl string) (name_used string, return__ int, __err__ error) {
	__cgo__tmpl := (*C.gchar)(unsafe.Pointer(C.CString(tmpl)))
	var __cgo__name_used *C.gchar
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_file_open_tmp(__cgo__tmpl, &__cgo__name_used, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__tmpl))
	name_used = C.GoString((*C.char)(unsafe.Pointer(__cgo__name_used)))
	return__ = int(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Reads the contents of the symbolic link @filename like the POSIX
readlink() function.  The returned string is in the encoding used
for filenames. Use g_filename_to_utf8() to convert it to UTF-8.
*/
func FileReadLink(filename string) (return__ string, __err__ error) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_file_read_link(__cgo__filename, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__filename))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Writes all of @contents to a file named @filename, with good error checking.
If a file called @filename already exists it will be overwritten.

This write is atomic in the sense that it is first written to a temporary
file which is then renamed to the final name. Notes:

- On UNIX, if @filename already exists hard links to @filename will break.
  Also since the file is recreated, existing permissions, access control
  lists, metadata etc. may be lost. If @filename is a symbolic link,
  the link itself will be replaced, not the linked file.

- On Windows renaming a file will not remove an existing file with the
  new name, so on Windows there is a race condition between the existing
  file being removed and the temporary file being renamed.

- On Windows there is no way to remove a file that is open to some
  process, or mapped into memory. Thus, this function will fail if
  @filename already exists and is open.

If the call was successful, it returns %TRUE. If the call was not successful,
it returns %FALSE and sets @error. The error domain is #G_FILE_ERROR.
Possible error codes are those in the #GFileError enumeration.

Note that the name for the temporary file is constructed by appending up
to 7 characters to @filename.
*/
func FileSetContents(filename string, contents []byte, length int64) (return__ bool, __err__ error) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	__header__contents := (*reflect.SliceHeader)(unsafe.Pointer(&contents))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_set_contents(__cgo__filename, (*C.gchar)(unsafe.Pointer(__header__contents.Data)), C.gssize(length), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__filename))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Returns %TRUE if any of the tests in the bitfield @test are
%TRUE. For example, `(G_FILE_TEST_EXISTS | G_FILE_TEST_IS_DIR)`
will return %TRUE if the file exists; the check whether it's a
directory doesn't matter since the existence test is %TRUE. With
the current set of available tests, there's no point passing in
more than one test at a time.

Apart from %G_FILE_TEST_IS_SYMLINK all tests follow symbolic links,
so for a symbolic link to a regular file g_file_test() will return
%TRUE for both %G_FILE_TEST_IS_SYMLINK and %G_FILE_TEST_IS_REGULAR.

Note, that for a dangling symbolic link g_file_test() will return
%TRUE for %G_FILE_TEST_IS_SYMLINK and %FALSE for all other flags.

You should never use g_file_test() to test whether it is safe
to perform an operation, because there is always the possibility
of the condition changing before you actually perform the operation.
For example, you might think you could use %G_FILE_TEST_IS_SYMLINK
to know whether it is safe to write to a file without being
tricked into writing into a different location. It doesn't work!
|[<!-- language="C" -->
 // DON'T DO THIS
 if (!g_file_test (filename, G_FILE_TEST_IS_SYMLINK))
   {
     fd = g_open (filename, O_WRONLY);
     // write to fd
   }
]|

Another thing to note is that %G_FILE_TEST_EXISTS and
%G_FILE_TEST_IS_EXECUTABLE are implemented using the access()
system call. This usually doesn't matter, but if your program
is setuid or setgid it means that these tests will give you
the answer for the real user ID and group ID, rather than the
effective user ID and group ID.

On Windows, there are no symlinks, so testing for
%G_FILE_TEST_IS_SYMLINK will always return %FALSE. Testing for
%G_FILE_TEST_IS_EXECUTABLE will just check that the file exists and
its name indicates that it is executable, checking for well-known
extensions and those listed in the `PATHEXT` environment variable.
*/
func FileTest(filename string, test C.GFileTest) (return__ bool) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_file_test(__cgo__filename, test)
	C.free(unsafe.Pointer(__cgo__filename))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the display basename for the particular filename, guaranteed
to be valid UTF-8. The display name might not be identical to the filename,
for instance there might be problems converting it to UTF-8, and some files
can be translated in the display.

If GLib cannot make sense of the encoding of @filename, as a last resort it
replaces unknown characters with U+FFFD, the Unicode replacement character.
You can search the result for the UTF-8 encoding of this character (which is
"\357\277\275" in octal notation) to find out if @filename was in an invalid
encoding.

You must pass the whole absolute pathname to this functions so that
translation of well known locations can be done.

This function is preferred over g_filename_display_name() if you know the
whole path, as it allows translation.
*/
func FilenameDisplayBasename(filename string) (return__ string) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_filename_display_basename(__cgo__filename)
	C.free(unsafe.Pointer(__cgo__filename))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts a filename into a valid UTF-8 string. The conversion is
not necessarily reversible, so you should keep the original around
and use the return value of this function only for display purposes.
Unlike g_filename_to_utf8(), the result is guaranteed to be non-%NULL
even if the filename actually isn't in the GLib file name encoding.

If GLib cannot make sense of the encoding of @filename, as a last resort it
replaces unknown characters with U+FFFD, the Unicode replacement character.
You can search the result for the UTF-8 encoding of this character (which is
"\357\277\275" in octal notation) to find out if @filename was in an invalid
encoding.

If you know the whole pathname of the file you should use
g_filename_display_basename(), since that allows location-based
translation of filenames.
*/
func FilenameDisplayName(filename string) (return__ string) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_filename_display_name(__cgo__filename)
	C.free(unsafe.Pointer(__cgo__filename))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts an escaped ASCII-encoded URI to a local filename in the
encoding used for filenames.
*/
func FilenameFromUri(uri string) (hostname string, return__ string, __err__ error) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	var __cgo__hostname *C.gchar
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_filename_from_uri(__cgo__uri, &__cgo__hostname, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__uri))
	hostname = C.GoString((*C.char)(unsafe.Pointer(__cgo__hostname)))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Converts a string from UTF-8 to the encoding GLib uses for
filenames. Note that on Windows GLib uses UTF-8 for filenames;
on other platforms, this function indirectly depends on the
[current locale][setlocale].
*/
func FilenameFromUtf8(utf8string string, len_ int64) (bytes_read int64, bytes_written int64, return__ []byte, __err__ error) {
	__cgo__utf8string := (*C.gchar)(unsafe.Pointer(C.CString(utf8string)))
	var __cgo__bytes_read C.gsize
	var __cgo__bytes_written C.gsize
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_filename_from_utf8(__cgo__utf8string, C.gssize(len_), &__cgo__bytes_read, &__cgo__bytes_written, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__utf8string))
	bytes_read = int64(__cgo__bytes_read)
	bytes_written = int64(__cgo__bytes_written)
	defer func() { return__ = C.GoBytes(unsafe.Pointer(__cgo__return__), C.int(bytes_written)) }()
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Converts an absolute filename to an escaped ASCII-encoded URI, with the path
component following Section 3.3. of RFC 2396.
*/
func FilenameToUri(filename string, hostname string) (return__ string, __err__ error) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	__cgo__hostname := (*C.gchar)(unsafe.Pointer(C.CString(hostname)))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_filename_to_uri(__cgo__filename, __cgo__hostname, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__filename))
	C.free(unsafe.Pointer(__cgo__hostname))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Converts a string which is in the encoding used by GLib for
filenames into a UTF-8 string. Note that on Windows GLib uses UTF-8
for filenames; on other platforms, this function indirectly depends on
the [current locale][setlocale].
*/
func FilenameToUtf8(opsysstring string, len_ int64, bytes_read *C.gsize, bytes_written *C.gsize) (return__ string, __err__ error) {
	__cgo__opsysstring := (*C.gchar)(unsafe.Pointer(C.CString(opsysstring)))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_filename_to_utf8(__cgo__opsysstring, C.gssize(len_), bytes_read, bytes_written, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__opsysstring))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Locates the first executable named @program in the user's path, in the
same way that execvp() would locate it. Returns an allocated string
with the absolute path name, or %NULL if the program is not found in
the path. If @program is already an absolute path, returns a copy of
@program if @program exists and is executable, and %NULL otherwise.

On Windows, if @program does not have a file type suffix, tries
with the suffixes .exe, .cmd, .bat and .com, and the suffixes in
the `PATHEXT` environment variable.

On Windows, it looks for the file in the same way as CreateProcess()
would. This means first in the directory where the executing
program was loaded from, then in the current directory, then in the
Windows 32-bit system directory, then in the Windows directory, and
finally in the directories in the `PATH` environment variable. If
the program is found, the return value contains the full name
including the type suffix.
*/
func FindProgramInPath(program string) (return__ string) {
	__cgo__program := (*C.gchar)(unsafe.Pointer(C.CString(program)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_find_program_in_path(__cgo__program)
	C.free(unsafe.Pointer(__cgo__program))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Formats a size (for example the size of a file) into a human readable
string.  Sizes are rounded to the nearest size prefix (kB, MB, GB)
and are displayed rounded to the nearest tenth. E.g. the file size
3292528 bytes will be converted into the string "3.2 MB".

The prefix units base is 1000 (i.e. 1 kB is 1000 bytes).

This string should be freed with g_free() when not needed any longer.

See g_format_size_full() for more options about how the size might be
formatted.
*/
func FormatSize(size uint64) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_format_size(C.guint64(size))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

// g_format_size_for_display is not generated due to deprecation attr

/*
Formats a size.

This function is similar to g_format_size() but allows for flags
that modify the output. See #GFormatSizeFlags.
*/
func FormatSizeFull(size uint64, flags C.GFormatSizeFlags) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_format_size_full(C.guint64(size), flags)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

// g_fprintf is not generated due to varargs

/*
Frees the memory pointed to by @mem.
If @mem is %NULL it simply returns.
*/
func Free(mem unsafe.Pointer) {
	C.g_free((C.gpointer)(mem))
	return
}

/*
Gets a human-readable name for the application, as set by
g_set_application_name(). This name should be localized if
possible, and is intended for display to the user.  Contrast with
g_get_prgname(), which gets a non-localized name. If
g_set_application_name() has not been called, returns the result of
g_get_prgname() (which may be %NULL if g_set_prgname() has also not
been called).
*/
func GetApplicationName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_get_application_name()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Obtains the character set for the [current locale][setlocale]; you
might use this character set as an argument to g_convert(), to convert
from the current locale's encoding to some other encoding. (Frequently
g_locale_to_utf8() and g_locale_from_utf8() are nice shortcuts, though.)

On Windows the character set returned by this function is the
so-called system default ANSI code-page. That is the character set
used by the "narrow" versions of C library and Win32 functions that
handle file names. It might be different from the character set
used by the C library's current locale.

The return value is %TRUE if the locale's encoding is UTF-8, in that
case you can perhaps avoid calling g_convert().

The string returned in @charset is not allocated, and should not be
freed.
*/
func GetCharset(charset **C.char) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_get_charset(charset)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the character set for the current locale.
*/
func GetCodeset() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_get_codeset()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the current directory.

The returned string should be freed when no longer needed.
The encoding of the returned string is system defined.
On Windows, it is always UTF-8.

Since GLib 2.40, this function will return the value of the "PWD"
environment variable if it is set and it happens to be the same as
the current directory.  This can make a difference in the case that
the current directory is the target of a symbolic link.
*/
func GetCurrentDir() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_get_current_dir()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Equivalent to the UNIX gettimeofday() function, but portable.

You may find g_get_real_time() to be more convenient.
*/
func GetCurrentTime(result *TimeVal) {
	C.g_get_current_time((*C.GTimeVal)(unsafe.Pointer(result)))
	return
}

/*
Gets the list of environment variables for the current process.

The list is %NULL terminated and each item in the list is of the
form 'NAME=VALUE'.

This is equivalent to direct access to the 'environ' global variable,
except portable.

The return value is freshly allocated and it should be freed with
g_strfreev() when it is no longer needed.
*/
func GetEnviron() (return__ []string) {
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.g_get_environ()
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Determines the preferred character sets used for filenames.
The first character set from the @charsets is the filename encoding, the
subsequent character sets are used when trying to generate a displayable
representation of a filename, see g_filename_display_name().

On Unix, the character sets are determined by consulting the
environment variables `G_FILENAME_ENCODING` and `G_BROKEN_FILENAMES`.
On Windows, the character set used in the GLib API is always UTF-8
and said environment variables have no effect.

`G_FILENAME_ENCODING` may be set to a comma-separated list of
character set names. The special token "&commat;locale" is taken
to  mean the character set for the [current locale][setlocale].
If `G_FILENAME_ENCODING` is not set, but `G_BROKEN_FILENAMES` is,
the character set of the current locale is taken as the filename
encoding. If neither environment variable  is set, UTF-8 is taken
as the filename encoding, but the character set of the current locale
is also put in the list of encodings.

The returned @charsets belong to GLib and must not be freed.

Note that on Unix, regardless of the locale character set or
`G_FILENAME_ENCODING` value, the actual file names present
on a system might be in any random encoding or just gibberish.
*/
func GetFilenameCharsets(charsets ***C.gchar) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_get_filename_charsets(charsets)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the current user's home directory.

As with most UNIX tools, this function will return the value of the
`HOME` environment variable if it is set to an existing absolute path
name, falling back to the `passwd` file in the case that it is unset.

If the path given in `HOME` is non-absolute, does not exist, or is
not a directory, the result is undefined.

Before version 2.36 this function would ignore the `HOME` environment
variable, taking the value from the `passwd` database instead. This was
changed to increase the compatibility of GLib with other programs (and
the XDG basedir specification) and to increase testability of programs
based on GLib (by making it easier to run them from test frameworks).

If your program has a strong requirement for either the new or the
old behaviour (and if you don't wish to increase your GLib
dependency to ensure that the new behaviour is in effect) then you
should either directly check the `HOME` environment variable yourself
or unset it before calling any functions in GLib.
*/
func GetHomeDir() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_get_home_dir()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Return a name for the machine.

The returned name is not necessarily a fully-qualified domain name,
or even present in DNS or some other name service at all. It need
not even be unique on your local network or site, but usually it
is. Callers should not rely on the return value having any specific
properties like uniqueness for security purposes. Even if the name
of the machine is changed while an application is running, the
return value from this function does not change. The returned
string is owned by GLib and should not be modified or freed. If no
name can be determined, a default fixed string "localhost" is
returned.
*/
func GetHostName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_get_host_name()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Computes a list of applicable locale names, which can be used to
e.g. construct locale-dependent filenames or search paths. The returned
list is sorted from most desirable to least desirable and always contains
the default locale "C".

For example, if LANGUAGE=de:en_US, then the returned list is
"de", "en_US", "en", "C".

This function consults the environment variables `LANGUAGE`, `LC_ALL`,
`LC_MESSAGES` and `LANG` to find the list of locales specified by the
user.
*/
func GetLanguageNames() (return__ []string) {
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.g_get_language_names()
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Returns a list of derived variants of @locale, which can be used to
e.g. construct locale-dependent filenames or search paths. The returned
list is sorted from most desirable to least desirable.
This function handles territory, charset and extra locale modifiers.

For example, if @locale is "fr_BE", then the returned list
is "fr_BE", "fr".

If you need the list of variants for the current locale,
use g_get_language_names().
*/
func GetLocaleVariants(locale string) (return__ []string) {
	__cgo__locale := (*C.gchar)(unsafe.Pointer(C.CString(locale)))
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.g_get_locale_variants(__cgo__locale)
	C.free(unsafe.Pointer(__cgo__locale))
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Queries the system monotonic time.

The monotonic clock will always increase and doesn't suffer
discontinuities when the user (or NTP) changes the system time.  It
may or may not continue to tick during times where the machine is
suspended.

We try to use the clock that corresponds as closely as possible to
the passage of time as measured by system calls such as poll() but it
may not always be possible to do this.
*/
func GetMonotonicTime() (return__ int64) {
	var __cgo__return__ C.gint64
	__cgo__return__ = C.g_get_monotonic_time()
	return__ = int64(__cgo__return__)
	return
}

/*
Determine the approximate number of threads that the system will
schedule simultaneously for this process.  This is intended to be
used as a parameter to g_thread_pool_new() for CPU bound tasks and
similar cases.
*/
func GetNumProcessors() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_get_num_processors()
	return__ = uint(__cgo__return__)
	return
}

/*
Gets the name of the program. This name should not be localized,
in contrast to g_get_application_name().

If you are using GDK or GTK+ the program name is set in gdk_init(),
which is called by gtk_init(). The program name is found by taking
the last component of @argv[0].
*/
func GetPrgname() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_get_prgname()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the real name of the user. This usually comes from the user's
entry in the `passwd` file. The encoding of the returned string is
system-defined. (On Windows, it is, however, always UTF-8.) If the
real user name cannot be determined, the string "Unknown" is
returned.
*/
func GetRealName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_get_real_name()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Queries the system wall-clock time.

This call is functionally equivalent to g_get_current_time() except
that the return value is often more convenient than dealing with a
#GTimeVal.

You should only use this call if you are actually interested in the real
wall-clock time.  g_get_monotonic_time() is probably more useful for
measuring intervals.
*/
func GetRealTime() (return__ int64) {
	var __cgo__return__ C.gint64
	__cgo__return__ = C.g_get_real_time()
	return__ = int64(__cgo__return__)
	return
}

/*
Returns an ordered list of base directories in which to access
system-wide configuration information.

On UNIX platforms this is determined using the mechanisms described
in the
[XDG Base Directory Specification](http://www.freedesktop.org/Standards/basedir-spec).
In this case the list of directories retrieved will be `XDG_CONFIG_DIRS`.

On Windows is the directory that contains application data for all users.
A typical path is C:\Documents and Settings\All Users\Application Data.
This folder is used for application data that is not user specific.
For example, an application can store a spell-check dictionary, a database
of clip art, or a log file in the CSIDL_COMMON_APPDATA folder.
This information will not roam and is available to anyone using the computer.
*/
func GetSystemConfigDirs() (return__ []string) {
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.g_get_system_config_dirs()
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Returns an ordered list of base directories in which to access
system-wide application data.

On UNIX platforms this is determined using the mechanisms described
in the
[XDG Base Directory Specification](http://www.freedesktop.org/Standards/basedir-spec)
In this case the list of directories retrieved will be XDG_DATA_DIRS.

On Windows the first elements in the list are the Application Data
and Documents folders for All Users. (These can be determined only
on Windows 2000 or later and are not present in the list on other
Windows versions.) See documentation for CSIDL_COMMON_APPDATA and
CSIDL_COMMON_DOCUMENTS.

Then follows the "share" subfolder in the installation folder for
the package containing the DLL that calls this function, if it can
be determined.

Finally the list contains the "share" subfolder in the installation
folder for GLib, and in the installation folder for the package the
application's .exe file belongs to.

The installation folders above are determined by looking up the
folder where the module (DLL or EXE) in question is located. If the
folder's name is "bin", its parent is used, otherwise the folder
itself.

Note that on Windows the returned list can vary depending on where
this function is called.
*/
func GetSystemDataDirs() (return__ []string) {
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.g_get_system_data_dirs()
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Gets the directory to use for temporary files.

On UNIX, this is taken from the `TMPDIR` environment variable.
If the variable is not set, `P_tmpdir` is
used, as defined by the system C library. Failing that, a
hard-coded default of "/tmp" is returned.

On Windows, the `TEMP` environment variable is used, with the
root directory of the Windows installation (eg: "C:\") used
as a default.

The encoding of the returned string is system-defined. On Windows,
it is always UTF-8. The return value is never %NULL or the empty
string.
*/
func GetTmpDir() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_get_tmp_dir()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns a base directory in which to store non-essential, cached
data specific to particular user.

On UNIX platforms this is determined using the mechanisms described
in the
[XDG Base Directory Specification](http://www.freedesktop.org/Standards/basedir-spec).
In this case the directory retrieved will be XDG_CACHE_HOME.

On Windows is the directory that serves as a common repository for
temporary Internet files. A typical path is
C:\Documents and Settings\username\Local Settings\Temporary Internet Files.
See documentation for CSIDL_INTERNET_CACHE.
*/
func GetUserCacheDir() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_get_user_cache_dir()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns a base directory in which to store user-specific application
configuration information such as user preferences and settings.

On UNIX platforms this is determined using the mechanisms described
in the
[XDG Base Directory Specification](http://www.freedesktop.org/Standards/basedir-spec).
In this case the directory retrieved will be `XDG_CONFIG_HOME`.

On Windows this is the folder to use for local (as opposed to
roaming) application data. See documentation for
CSIDL_LOCAL_APPDATA. Note that on Windows it thus is the same as
what g_get_user_data_dir() returns.
*/
func GetUserConfigDir() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_get_user_config_dir()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns a base directory in which to access application data such
as icons that is customized for a particular user.

On UNIX platforms this is determined using the mechanisms described
in the
[XDG Base Directory Specification](http://www.freedesktop.org/Standards/basedir-spec).
In this case the directory retrieved will be `XDG_DATA_HOME`.

On Windows this is the folder to use for local (as opposed to
roaming) application data. See documentation for
CSIDL_LOCAL_APPDATA. Note that on Windows it thus is the same as
what g_get_user_config_dir() returns.
*/
func GetUserDataDir() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_get_user_data_dir()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the user name of the current user. The encoding of the returned
string is system-defined. On UNIX, it might be the preferred file name
encoding, or something else, and there is no guarantee that it is even
consistent on a machine. On Windows, it is always UTF-8.
*/
func GetUserName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_get_user_name()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns a directory that is unique to the current user on the local
system.

On UNIX platforms this is determined using the mechanisms described
in the
[XDG Base Directory Specification](http://www.freedesktop.org/Standards/basedir-spec).
This is the directory
specified in the `XDG_RUNTIME_DIR` environment variable.
In the case that this variable is not set, GLib will issue a warning
message to stderr and return the value of g_get_user_cache_dir().

On Windows this is the folder to use for local (as opposed to
roaming) application data. See documentation for
CSIDL_LOCAL_APPDATA.  Note that on Windows it thus is the same as
what g_get_user_config_dir() returns.
*/
func GetUserRuntimeDir() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_get_user_runtime_dir()
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the full path of a special directory using its logical id.

On UNIX this is done using the XDG special user directories.
For compatibility with existing practise, %G_USER_DIRECTORY_DESKTOP
falls back to `$HOME/Desktop` when XDG special user directories have
not been set up.

Depending on the platform, the user might be able to change the path
of the special directory without requiring the session to restart; GLib
will not reflect any change once the special directories are loaded.
*/
func GetUserSpecialDir(directory C.GUserDirectory) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_get_user_special_dir(directory)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the value of an environment variable.

The name and value are in the GLib file name encoding. On UNIX,
this means the actual bytes which might or might not be in some
consistent character set and encoding. On Windows, it is in UTF-8.
On Windows, in case the environment variable's value contains
references to other environment variables, they are expanded.
*/
func Getenv(variable string) (return__ string) {
	__cgo__variable := (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_getenv(__cgo__variable)
	C.free(unsafe.Pointer(__cgo__variable))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
This is a convenience function for using a #GHashTable as a set.  It
is equivalent to calling g_hash_table_replace() with @key as both the
key and the value.

When a hash table only ever contains keys that have themselves as the
corresponding value it is able to be stored more efficiently.  See
the discussion in the section description.
*/
func HashTableAdd(hash_table *HashTable, key unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_hash_table_add((*C.GHashTable)(unsafe.Pointer(hash_table)), (C.gpointer)(key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if @key is in @hash_table.
*/
func HashTableContains(hash_table *HashTable, key unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_hash_table_contains((*C.GHashTable)(unsafe.Pointer(hash_table)), (C.gconstpointer)(key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Destroys all keys and values in the #GHashTable and decrements its
reference count by 1. If keys and/or values are dynamically allocated,
you should either free them first or create the #GHashTable with destroy
notifiers using g_hash_table_new_full(). In the latter case the destroy
functions you supplied will be called on all keys and values during the
destruction phase.
*/
func HashTableDestroy(hash_table *HashTable) {
	C.g_hash_table_destroy((*C.GHashTable)(unsafe.Pointer(hash_table)))
	return
}

/*
Inserts a new key and value into a #GHashTable.

If the key already exists in the #GHashTable its current
value is replaced with the new value. If you supplied a
@value_destroy_func when creating the #GHashTable, the old
value is freed using that function. If you supplied a
@key_destroy_func when creating the #GHashTable, the passed
key is freed using that function.
*/
func HashTableInsert(hash_table *HashTable, key unsafe.Pointer, value unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_hash_table_insert((*C.GHashTable)(unsafe.Pointer(hash_table)), (C.gpointer)(key), (C.gpointer)(value))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Looks up a key in the #GHashTable, returning the original key and the
associated value and a #gboolean which is %TRUE if the key was found. This
is useful if you need to free the memory allocated for the original key,
for example before calling g_hash_table_remove().

You can actually pass %NULL for @lookup_key to test
whether the %NULL key exists, provided the hash and equal functions
of @hash_table are %NULL-safe.
*/
func HashTableLookupExtended(hash_table *HashTable, lookup_key unsafe.Pointer, orig_key *C.gpointer, value *C.gpointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_hash_table_lookup_extended((*C.GHashTable)(unsafe.Pointer(hash_table)), (C.gconstpointer)(lookup_key), orig_key, value)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Removes a key and its associated value from a #GHashTable.

If the #GHashTable was created using g_hash_table_new_full(), the
key and value are freed using the supplied destroy functions, otherwise
you have to make sure that any dynamically allocated values are freed
yourself.
*/
func HashTableRemove(hash_table *HashTable, key unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_hash_table_remove((*C.GHashTable)(unsafe.Pointer(hash_table)), (C.gconstpointer)(key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Removes all keys and their associated values from a #GHashTable.

If the #GHashTable was created using g_hash_table_new_full(),
the keys and values are freed using the supplied destroy functions,
otherwise you have to make sure that any dynamically allocated
values are freed yourself.
*/
func HashTableRemoveAll(hash_table *HashTable) {
	C.g_hash_table_remove_all((*C.GHashTable)(unsafe.Pointer(hash_table)))
	return
}

/*
Inserts a new key and value into a #GHashTable similar to
g_hash_table_insert(). The difference is that if the key
already exists in the #GHashTable, it gets replaced by the
new key. If you supplied a @value_destroy_func when creating
the #GHashTable, the old value is freed using that function.
If you supplied a @key_destroy_func when creating the
#GHashTable, the old key is freed using that function.
*/
func HashTableReplace(hash_table *HashTable, key unsafe.Pointer, value unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_hash_table_replace((*C.GHashTable)(unsafe.Pointer(hash_table)), (C.gpointer)(key), (C.gpointer)(value))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the number of elements contained in the #GHashTable.
*/
func HashTableSize(hash_table *HashTable) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_hash_table_size((*C.GHashTable)(unsafe.Pointer(hash_table)))
	return__ = uint(__cgo__return__)
	return
}

/*
Removes a key and its associated value from a #GHashTable without
calling the key and value destroy functions.
*/
func HashTableSteal(hash_table *HashTable, key unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_hash_table_steal((*C.GHashTable)(unsafe.Pointer(hash_table)), (C.gconstpointer)(key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Removes all keys and their associated values from a #GHashTable
without calling the key and value destroy functions.
*/
func HashTableStealAll(hash_table *HashTable) {
	C.g_hash_table_steal_all((*C.GHashTable)(unsafe.Pointer(hash_table)))
	return
}

/*
Atomically decrements the reference count of @hash_table by one.
If the reference count drops to 0, all keys and values will be
destroyed, and all memory allocated by the hash table is released.
This function is MT-safe and may be called from any thread.
*/
func HashTableUnref(hash_table *HashTable) {
	C.g_hash_table_unref((*C.GHashTable)(unsafe.Pointer(hash_table)))
	return
}

/*
Destroys a #GHook, given its ID.
*/
func HookDestroy(hook_list *HookList, hook_id int64) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_hook_destroy((*C.GHookList)(unsafe.Pointer(hook_list)), C.gulong(hook_id))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Removes one #GHook from a #GHookList, marking it
inactive and calling g_hook_unref() on it.
*/
func HookDestroyLink(hook_list *HookList, hook *Hook) {
	C.g_hook_destroy_link((*C.GHookList)(unsafe.Pointer(hook_list)), (*C.GHook)(unsafe.Pointer(hook)))
	return
}

/*
Calls the #GHookList @finalize_hook function if it exists,
and frees the memory allocated for the #GHook.
*/
func HookFree(hook_list *HookList, hook *Hook) {
	C.g_hook_free((*C.GHookList)(unsafe.Pointer(hook_list)), (*C.GHook)(unsafe.Pointer(hook)))
	return
}

/*
Inserts a #GHook into a #GHookList, before a given #GHook.
*/
func HookInsertBefore(hook_list *HookList, sibling *Hook, hook *Hook) {
	C.g_hook_insert_before((*C.GHookList)(unsafe.Pointer(hook_list)), (*C.GHook)(unsafe.Pointer(sibling)), (*C.GHook)(unsafe.Pointer(hook)))
	return
}

/*
Prepends a #GHook on the start of a #GHookList.
*/
func HookPrepend(hook_list *HookList, hook *Hook) {
	C.g_hook_prepend((*C.GHookList)(unsafe.Pointer(hook_list)), (*C.GHook)(unsafe.Pointer(hook)))
	return
}

/*
Decrements the reference count of a #GHook.
If the reference count falls to 0, the #GHook is removed
from the #GHookList and g_hook_free() is called to free it.
*/
func HookUnref(hook_list *HookList, hook *Hook) {
	C.g_hook_unref((*C.GHookList)(unsafe.Pointer(hook_list)), (*C.GHook)(unsafe.Pointer(hook)))
	return
}

/*
Tests if @hostname contains segments with an ASCII-compatible
encoding of an Internationalized Domain Name. If this returns
%TRUE, you should decode the hostname with g_hostname_to_unicode()
before displaying it to the user.

Note that a hostname might contain a mix of encoded and unencoded
segments, and so it is possible for g_hostname_is_non_ascii() and
g_hostname_is_ascii_encoded() to both return %TRUE for a name.
*/
func HostnameIsAsciiEncoded(hostname string) (return__ bool) {
	__cgo__hostname := (*C.gchar)(unsafe.Pointer(C.CString(hostname)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_hostname_is_ascii_encoded(__cgo__hostname)
	C.free(unsafe.Pointer(__cgo__hostname))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Tests if @hostname is the string form of an IPv4 or IPv6 address.
(Eg, "192.168.0.1".)
*/
func HostnameIsIpAddress(hostname string) (return__ bool) {
	__cgo__hostname := (*C.gchar)(unsafe.Pointer(C.CString(hostname)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_hostname_is_ip_address(__cgo__hostname)
	C.free(unsafe.Pointer(__cgo__hostname))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Tests if @hostname contains Unicode characters. If this returns
%TRUE, you need to encode the hostname with g_hostname_to_ascii()
before using it in non-IDN-aware contexts.

Note that a hostname might contain a mix of encoded and unencoded
segments, and so it is possible for g_hostname_is_non_ascii() and
g_hostname_is_ascii_encoded() to both return %TRUE for a name.
*/
func HostnameIsNonAscii(hostname string) (return__ bool) {
	__cgo__hostname := (*C.gchar)(unsafe.Pointer(C.CString(hostname)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_hostname_is_non_ascii(__cgo__hostname)
	C.free(unsafe.Pointer(__cgo__hostname))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Converts @hostname to its canonical ASCII form; an ASCII-only
string containing no uppercase letters and not ending with a
trailing dot.
*/
func HostnameToAscii(hostname string) (return__ string) {
	__cgo__hostname := (*C.gchar)(unsafe.Pointer(C.CString(hostname)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_hostname_to_ascii(__cgo__hostname)
	C.free(unsafe.Pointer(__cgo__hostname))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts @hostname to its canonical presentation form; a UTF-8
string in Unicode normalization form C, containing no uppercase
letters, no forbidden characters, and no ASCII-encoded segments,
and not ending with a trailing dot.

Of course if @hostname is not an internationalized hostname, then
the canonical presentation form will be entirely ASCII.
*/
func HostnameToUnicode(hostname string) (return__ string) {
	__cgo__hostname := (*C.gchar)(unsafe.Pointer(C.CString(hostname)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_hostname_to_unicode(__cgo__hostname)
	C.free(unsafe.Pointer(__cgo__hostname))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Same as the standard UNIX routine iconv(), but
may be implemented via libiconv on UNIX flavors that lack
a native implementation.

GLib provides g_convert() and g_locale_to_utf8() which are likely
more convenient than the raw iconv wrappers.
*/
func Iconv(converter C.GIConv, inbuf **C.gchar, inbytes_left *C.gsize, outbuf **C.gchar, outbytes_left *C.gsize) (return__ int64) {
	var __cgo__return__ C.gsize
	__cgo__return__ = C.g_iconv(converter, inbuf, inbytes_left, outbuf, outbytes_left)
	return__ = int64(__cgo__return__)
	return
}

/*
Adds a function to be called whenever there are no higher priority
events pending to the default main loop. The function is given the
default idle priority, #G_PRIORITY_DEFAULT_IDLE.  If the function
returns %FALSE it is automatically removed from the list of event
sources and will not be called again.

This internally creates a main loop source using g_idle_source_new()
and attaches it to the main loop context using g_source_attach().
You can do these steps manually if you need greater control.
*/
func IdleAdd(function C.GSourceFunc, data unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_idle_add(function, (C.gpointer)(data))
	return__ = uint(__cgo__return__)
	return
}

/*
Adds a function to be called whenever there are no higher priority
events pending.  If the function returns %FALSE it is automatically
removed from the list of event sources and will not be called again.

This internally creates a main loop source using g_idle_source_new()
and attaches it to the main loop context using g_source_attach().
You can do these steps manually if you need greater control.
*/
func IdleAddFull(priority int, function C.GSourceFunc, data unsafe.Pointer, notify C.GDestroyNotify) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_idle_add_full(C.gint(priority), function, (C.gpointer)(data), notify)
	return__ = uint(__cgo__return__)
	return
}

/*
Removes the idle function with the given data.
*/
func IdleRemoveByData(data unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_idle_remove_by_data((C.gpointer)(data))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Creates a new idle source.

The source will not initially be associated with any #GMainContext
and must be added to one with g_source_attach() before it will be
executed. Note that the default priority for idle sources is
%G_PRIORITY_DEFAULT_IDLE, as compared to other sources which
have a default priority of %G_PRIORITY_DEFAULT.
*/
func IdleSourceNew() (return__ *Source) {
	var __cgo__return__ *C.GSource
	__cgo__return__ = C.g_idle_source_new()
	return__ = (*Source)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Compares the two #gint64 values being pointed to and returns
%TRUE if they are equal.
It can be passed to g_hash_table_new() as the @key_equal_func
parameter, when using non-%NULL pointers to 64-bit integers as keys in a
#GHashTable.
*/
func Int64Equal(v1 unsafe.Pointer, v2 unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_int64_equal((C.gconstpointer)(v1), (C.gconstpointer)(v2))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Converts a pointer to a #gint64 to a hash value.

It can be passed to g_hash_table_new() as the @hash_func parameter,
when using non-%NULL pointers to 64-bit integer values as keys in a
#GHashTable.
*/
func Int64Hash(v unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_int64_hash((C.gconstpointer)(v))
	return__ = uint(__cgo__return__)
	return
}

/*
Compares the two #gint values being pointed to and returns
%TRUE if they are equal.
It can be passed to g_hash_table_new() as the @key_equal_func
parameter, when using non-%NULL pointers to integers as keys in a
#GHashTable.

Note that this function acts on pointers to #gint, not on #gint
directly: if your hash table's keys are of the form
`GINT_TO_POINTER (n)`, use g_direct_equal() instead.
*/
func IntEqual(v1 unsafe.Pointer, v2 unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_int_equal((C.gconstpointer)(v1), (C.gconstpointer)(v2))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Converts a pointer to a #gint to a hash value.
It can be passed to g_hash_table_new() as the @hash_func parameter,
when using non-%NULL pointers to integer values as keys in a #GHashTable.

Note that this function acts on pointers to #gint, not on #gint
directly: if your hash table's keys are of the form
`GINT_TO_POINTER (n)`, use g_direct_hash() instead.
*/
func IntHash(v unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_int_hash((C.gconstpointer)(v))
	return__ = uint(__cgo__return__)
	return
}

/*
Returns a canonical representation for @string. Interned strings
can be compared for equality by comparing the pointers, instead of
using strcmp(). g_intern_static_string() does not copy the string,
therefore @string must not be freed or modified.
*/
func InternStaticString(string_ string) (return__ string) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_intern_static_string(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns a canonical representation for @string. Interned strings
can be compared for equality by comparing the pointers, instead of
using strcmp().
*/
func InternString(string_ string) (return__ string) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_intern_string(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Adds the #GIOChannel into the default main loop context
with the default priority.
*/
func IoAddWatch(channel *IOChannel, condition C.GIOCondition, func_ C.GIOFunc, user_data unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_io_add_watch((*C.GIOChannel)(unsafe.Pointer(channel)), condition, func_, (C.gpointer)(user_data))
	return__ = uint(__cgo__return__)
	return
}

/*
Adds the #GIOChannel into the default main loop context
with the given priority.

This internally creates a main loop source using g_io_create_watch()
and attaches it to the main loop context with g_source_attach().
You can do these steps manually if you need greater control.
*/
func IoAddWatchFull(channel *IOChannel, priority int, condition C.GIOCondition, func_ C.GIOFunc, user_data unsafe.Pointer, notify C.GDestroyNotify) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_io_add_watch_full((*C.GIOChannel)(unsafe.Pointer(channel)), C.gint(priority), condition, func_, (C.gpointer)(user_data), notify)
	return__ = uint(__cgo__return__)
	return
}

/*
Converts an `errno` error number to a #GIOChannelError.
*/
func IoChannelErrorFromErrno(en int) (return__ C.GIOChannelError) {
	return__ = C.g_io_channel_error_from_errno(C.gint(en))
	return
}

func IoChannelErrorQuark() (return__ C.GQuark) {
	return__ = C.g_io_channel_error_quark()
	return
}

/*
Creates a #GSource that's dispatched when @condition is met for the
given @channel. For example, if condition is #G_IO_IN, the source will
be dispatched when there's data available for reading.

g_io_add_watch() is a simpler interface to this same functionality, for
the case where you want to add the source to the default main loop context
at the default priority.

On Windows, polling a #GSource created to watch a channel for a socket
puts the socket in non-blocking mode. This is a side-effect of the
implementation and unavoidable.
*/
func IoCreateWatch(channel *IOChannel, condition C.GIOCondition) (return__ *Source) {
	var __cgo__return__ *C.GSource
	__cgo__return__ = C.g_io_create_watch((*C.GIOChannel)(unsafe.Pointer(channel)), condition)
	return__ = (*Source)(unsafe.Pointer(__cgo__return__))
	return
}

func KeyFileErrorQuark() (return__ C.GQuark) {
	return__ = C.g_key_file_error_quark()
	return
}

/*
Gets the names of all variables set in the environment.

Programs that want to be portable to Windows should typically use
this function and g_getenv() instead of using the environ array
from the C library directly. On Windows, the strings in the environ
array are in system codepage encoding, while in most of the typical
use cases for environment variables in GLib-using programs you want
the UTF-8 encoding that this function and g_getenv() provide.
*/
func Listenv() (return__ []string) {
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.g_listenv()
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Converts a string from UTF-8 to the encoding used for strings by
the C runtime (usually the same as that used by the operating
system) in the [current locale][setlocale]. On Windows this means
the system codepage.
*/
func LocaleFromUtf8(utf8string string, len_ int64, bytes_read *C.gsize, bytes_written *C.gsize) (return__ string, __err__ error) {
	__cgo__utf8string := (*C.gchar)(unsafe.Pointer(C.CString(utf8string)))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_locale_from_utf8(__cgo__utf8string, C.gssize(len_), bytes_read, bytes_written, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__utf8string))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Converts a string which is in the encoding used for strings by
the C runtime (usually the same as that used by the operating
system) in the [current locale][setlocale] into a UTF-8 string.
*/
func LocaleToUtf8(opsysstring string, len_ int64, bytes_read *C.gsize, bytes_written *C.gsize) (return__ string, __err__ error) {
	__cgo__opsysstring := (*C.gchar)(unsafe.Pointer(C.CString(opsysstring)))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_locale_to_utf8(__cgo__opsysstring, C.gssize(len_), bytes_read, bytes_written, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__opsysstring))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

// g_log is not generated due to varargs

/*
The default log handler set up by GLib; g_log_set_default_handler()
allows to install an alternate default log handler.
This is used if no log handler has been set for the particular log
domain and log level combination. It outputs the message to stderr
or stdout and if the log level is fatal it calls abort(). It automatically
prints a new-line character after the message, so one does not need to be
manually included in @message.

The behavior of this log handler can be influenced by a number of
environment variables:

- `G_MESSAGES_PREFIXED`: A :-separated list of log levels for which
  messages should be prefixed by the program name and PID of the
  aplication.

- `G_MESSAGES_DEBUG`: A space-separated list of log domains for
  which debug and informational messages are printed. By default
  these messages are not printed.

stderr is used for levels %G_LOG_LEVEL_ERROR, %G_LOG_LEVEL_CRITICAL,
%G_LOG_LEVEL_WARNING and %G_LOG_LEVEL_MESSAGE. stdout is used for
the rest.
*/
func LogDefaultHandler(log_domain string, log_level C.GLogLevelFlags, message string, unused_data unsafe.Pointer) {
	__cgo__log_domain := (*C.gchar)(unsafe.Pointer(C.CString(log_domain)))
	__cgo__message := (*C.gchar)(unsafe.Pointer(C.CString(message)))
	C.g_log_default_handler(__cgo__log_domain, log_level, __cgo__message, (C.gpointer)(unused_data))
	C.free(unsafe.Pointer(__cgo__log_domain))
	C.free(unsafe.Pointer(__cgo__message))
	return
}

/*
Removes the log handler.
*/
func LogRemoveHandler(log_domain string, handler_id uint) {
	__cgo__log_domain := (*C.gchar)(unsafe.Pointer(C.CString(log_domain)))
	C.g_log_remove_handler(__cgo__log_domain, C.guint(handler_id))
	C.free(unsafe.Pointer(__cgo__log_domain))
	return
}

/*
Sets the message levels which are always fatal, in any log domain.
When a message with any of these levels is logged the program terminates.
You can only set the levels defined by GLib to be fatal.
%G_LOG_LEVEL_ERROR is always fatal.

You can also make some message levels fatal at runtime by setting
the `G_DEBUG` environment variable (see
[Running GLib Applications](glib-running.html)).
*/
func LogSetAlwaysFatal(fatal_mask C.GLogLevelFlags) (return__ C.GLogLevelFlags) {
	return__ = C.g_log_set_always_fatal(fatal_mask)
	return
}

/*
Installs a default log handler which is used if no
log handler has been set for the particular log domain
and log level combination. By default, GLib uses
g_log_default_handler() as default log handler.
*/
func LogSetDefaultHandler(log_func C.GLogFunc, user_data unsafe.Pointer) (return__ C.GLogFunc) {
	return__ = C.g_log_set_default_handler(log_func, (C.gpointer)(user_data))
	return
}

/*
Sets the log levels which are fatal in the given domain.
%G_LOG_LEVEL_ERROR is always fatal.
*/
func LogSetFatalMask(log_domain string, fatal_mask C.GLogLevelFlags) (return__ C.GLogLevelFlags) {
	__cgo__log_domain := (*C.gchar)(unsafe.Pointer(C.CString(log_domain)))
	return__ = C.g_log_set_fatal_mask(__cgo__log_domain, fatal_mask)
	C.free(unsafe.Pointer(__cgo__log_domain))
	return
}

/*
Sets the log handler for a domain and a set of log levels.
To handle fatal and recursive messages the @log_levels parameter
must be combined with the #G_LOG_FLAG_FATAL and #G_LOG_FLAG_RECURSION
bit flags.

Note that since the #G_LOG_LEVEL_ERROR log level is always fatal, if
you want to set a handler for this log level you must combine it with
#G_LOG_FLAG_FATAL.

Here is an example for adding a log handler for all warning messages
in the default domain:
|[<!-- language="C" -->
g_log_set_handler (NULL, G_LOG_LEVEL_WARNING | G_LOG_FLAG_FATAL
                   | G_LOG_FLAG_RECURSION, my_log_handler, NULL);
]|

This example adds a log handler for all critical messages from GTK+:
|[<!-- language="C" -->
g_log_set_handler ("Gtk", G_LOG_LEVEL_CRITICAL | G_LOG_FLAG_FATAL
                   | G_LOG_FLAG_RECURSION, my_log_handler, NULL);
]|

This example adds a log handler for all messages from GLib:
|[<!-- language="C" -->
g_log_set_handler ("GLib", G_LOG_LEVEL_MASK | G_LOG_FLAG_FATAL
                   | G_LOG_FLAG_RECURSION, my_log_handler, NULL);
]|
*/
func LogSetHandler(log_domain string, log_levels C.GLogLevelFlags, log_func C.GLogFunc, user_data unsafe.Pointer) (return__ uint) {
	__cgo__log_domain := (*C.gchar)(unsafe.Pointer(C.CString(log_domain)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_log_set_handler(__cgo__log_domain, log_levels, log_func, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__log_domain))
	return__ = uint(__cgo__return__)
	return
}

// g_logv is not generated due to varargs

/*
Returns the global default main context. This is the main context
used for main loop functions when a main loop is not explicitly
specified, and corresponds to the "main" main loop. See also
g_main_context_get_thread_default().
*/
func MainContextDefault() (return__ *MainContext) {
	var __cgo__return__ *C.GMainContext
	__cgo__return__ = C.g_main_context_default()
	return__ = (*MainContext)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Gets the thread-default #GMainContext for this thread. Asynchronous
operations that want to be able to be run in contexts other than
the default one should call this method or
g_main_context_ref_thread_default() to get a #GMainContext to add
their #GSources to. (Note that even in single-threaded
programs applications may sometimes want to temporarily push a
non-default context, so it is not safe to assume that this will
always return %NULL if you are running in the default thread.)

If you need to hold a reference on the context, use
g_main_context_ref_thread_default() instead.
*/
func MainContextGetThreadDefault() (return__ *MainContext) {
	var __cgo__return__ *C.GMainContext
	__cgo__return__ = C.g_main_context_get_thread_default()
	return__ = (*MainContext)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Gets the thread-default #GMainContext for this thread, as with
g_main_context_get_thread_default(), but also adds a reference to
it with g_main_context_ref(). In addition, unlike
g_main_context_get_thread_default(), if the thread-default context
is the global default context, this will return that #GMainContext
(with a ref added to it) rather than returning %NULL.
*/
func MainContextRefThreadDefault() (return__ *MainContext) {
	var __cgo__return__ *C.GMainContext
	__cgo__return__ = C.g_main_context_ref_thread_default()
	return__ = (*MainContext)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Returns the currently firing source for this thread.
*/
func MainCurrentSource() (return__ *Source) {
	var __cgo__return__ *C.GSource
	__cgo__return__ = C.g_main_current_source()
	return__ = (*Source)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Returns the depth of the stack of calls to
g_main_context_dispatch() on any #GMainContext in the current thread.
 That is, when called from the toplevel, it gives 0. When
called from within a callback from g_main_context_iteration()
(or g_main_loop_run(), etc.) it returns 1. When called from within
a callback to a recursive call to g_main_context_iteration(),
it returns 2. And so forth.

This function is useful in a situation like the following:
Imagine an extremely simple "garbage collected" system.

|[<!-- language="C" -->
static GList *free_list;

gpointer
allocate_memory (gsize size)
{
  gpointer result = g_malloc (size);
  free_list = g_list_prepend (free_list, result);
  return result;
}

void
free_allocated_memory (void)
{
  GList *l;
  for (l = free_list; l; l = l->next);
    g_free (l->data);
  g_list_free (free_list);
  free_list = NULL;
 }

[...]

while (TRUE);
 {
   g_main_context_iteration (NULL, TRUE);
   free_allocated_memory();
  }
]|

This works from an application, however, if you want to do the same
thing from a library, it gets more difficult, since you no longer
control the main loop. You might think you can simply use an idle
function to make the call to free_allocated_memory(), but that
doesn't work, since the idle function could be called from a
recursive callback. This can be fixed by using g_main_depth()

|[<!-- language="C" -->
gpointer
allocate_memory (gsize size)
{
  FreeListBlock *block = g_new (FreeListBlock, 1);
  block->mem = g_malloc (size);
  block->depth = g_main_depth ();
  free_list = g_list_prepend (free_list, block);
  return block->mem;
}

void
free_allocated_memory (void)
{
  GList *l;

  int depth = g_main_depth ();
  for (l = free_list; l; );
    {
      GList *next = l->next;
      FreeListBlock *block = l->data;
      if (block->depth > depth)
        {
          g_free (block->mem);
          g_free (block);
          free_list = g_list_delete_link (free_list, l);
        }

      l = next;
    }
  }
]|

There is a temptation to use g_main_depth() to solve
problems with reentrancy. For instance, while waiting for data
to be received from the network in response to a menu item,
the menu item might be selected again. It might seem that
one could make the menu item's callback return immediately
and do nothing if g_main_depth() returns a value greater than 1.
However, this should be avoided since the user then sees selecting
the menu item do nothing. Furthermore, you'll find yourself adding
these checks all over your code, since there are doubtless many,
many things that the user could do. Instead, you can use the
following techniques:

1. Use gtk_widget_set_sensitive() or modal dialogs to prevent
   the user from interacting with elements while the main
   loop is recursing.

2. Avoid main loop recursion in situations where you can't handle
   arbitrary  callbacks. Instead, structure your code so that you
   simply return to the main loop and then get called again when
   there is more work to do.
*/
func MainDepth() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_main_depth()
	return__ = int(__cgo__return__)
	return
}

/*
Allocates @n_bytes bytes of memory.
If @n_bytes is 0 it returns %NULL.
*/
func Malloc(n_bytes int64) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_malloc(C.gsize(n_bytes))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Allocates @n_bytes bytes of memory, initialized to 0's.
If @n_bytes is 0 it returns %NULL.
*/
func Malloc0(n_bytes int64) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_malloc0(C.gsize(n_bytes))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
This function is similar to g_malloc0(), allocating (@n_blocks * @n_block_bytes) bytes,
but care is taken to detect possible overflow during multiplication.
*/
func Malloc0N(n_blocks int64, n_block_bytes int64) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_malloc0_n(C.gsize(n_blocks), C.gsize(n_block_bytes))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
This function is similar to g_malloc(), allocating (@n_blocks * @n_block_bytes) bytes,
but care is taken to detect possible overflow during multiplication.
*/
func MallocN(n_blocks int64, n_block_bytes int64) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_malloc_n(C.gsize(n_blocks), C.gsize(n_block_bytes))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

// g_markup_collect_attributes is not generated due to varargs

func MarkupErrorQuark() (return__ C.GQuark) {
	return__ = C.g_markup_error_quark()
	return
}

/*
Escapes text so that the markup parser will parse it verbatim.
Less than, greater than, ampersand, etc. are replaced with the
corresponding entities. This function would typically be used
when writing out a file to be parsed with the markup parser.

Note that this function doesn't protect whitespace and line endings
from being processed according to the XML rules for normalization
of line endings and attribute values.

Note also that this function will produce character references in
the range of &#x1; ... &#x1f; for all control sequences
except for tabstop, newline and carriage return.  The character
references in this range are not valid XML 1.0, but they are
valid XML 1.1 and will be accepted by the GMarkup parser.
*/
func MarkupEscapeText(text string, length int64) (return__ string) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_markup_escape_text(__cgo__text, C.gssize(length))
	C.free(unsafe.Pointer(__cgo__text))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

// g_markup_printf_escaped is not generated due to varargs

// g_markup_vprintf_escaped is not generated due to varargs

/*
Checks whether the allocator used by g_malloc() is the system's
malloc implementation. If it returns %TRUE memory allocated with
malloc() can be used interchangeable with memory allocated using g_malloc().
This function is useful for avoiding an extra copy of allocated memory returned
by a non-GLib-based API.

A different allocator can be set using g_mem_set_vtable().
*/
func MemIsSystemMalloc() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_mem_is_system_malloc()
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Outputs a summary of memory usage.

It outputs the frequency of allocations of different sizes,
the total number of bytes which have been allocated,
the total number of bytes which have been freed,
and the difference between the previous two values, i.e. the number of bytes
still in use.

Note that this function will not output anything unless you have
previously installed the #glib_mem_profiler_table with g_mem_set_vtable().
*/
func MemProfile() {
	C.g_mem_profile()
	return
}

/*
Sets the #GMemVTable to use for memory allocation. You can use this
to provide custom memory allocation routines.

The @vtable only needs to provide malloc(), realloc(), and free()
functions; GLib can provide default implementations of the others.
The malloc() and realloc() implementations should return %NULL on
failure, GLib will handle error-checking for you. @vtable is copied,
so need not persist after this function has been called.

Note that this function must be called before using any other GLib
functions.
*/
func MemSetVtable(vtable *MemVTable) {
	C.g_mem_set_vtable((*C.GMemVTable)(unsafe.Pointer(vtable)))
	return
}

/*
Allocates @byte_size bytes of memory, and copies @byte_size bytes into it
from @mem. If @mem is %NULL it returns %NULL.
*/
func Memdup(mem unsafe.Pointer, byte_size uint) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_memdup((C.gconstpointer)(mem), C.guint(byte_size))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Create a directory if it doesn't already exist. Create intermediate
parent directories as needed, too.
*/
func MkdirWithParents(pathname string, mode int) (return__ int) {
	__cgo__pathname := (*C.gchar)(unsafe.Pointer(C.CString(pathname)))
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_mkdir_with_parents(__cgo__pathname, C.gint(mode))
	C.free(unsafe.Pointer(__cgo__pathname))
	return__ = int(__cgo__return__)
	return
}

/*
Creates a temporary directory. See the mkdtemp() documentation
on most UNIX-like systems.

The parameter is a string that should follow the rules for
mkdtemp() templates, i.e. contain the string "XXXXXX".
g_mkdtemp() is slightly more flexible than mkdtemp() in that the
sequence does not have to occur at the very end of the template
and you can pass a @mode and additional @flags. The X string will
be modified to form the name of a directory that didn't exist.
The string should be in the GLib file name encoding. Most importantly,
on Windows it should be in UTF-8.
*/
func Mkdtemp(tmpl string) (return__ string) {
	__cgo__tmpl := (*C.gchar)(unsafe.Pointer(C.CString(tmpl)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_mkdtemp(__cgo__tmpl)
	C.free(unsafe.Pointer(__cgo__tmpl))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Creates a temporary directory. See the mkdtemp() documentation
on most UNIX-like systems.

The parameter is a string that should follow the rules for
mkdtemp() templates, i.e. contain the string "XXXXXX".
g_mkdtemp() is slightly more flexible than mkdtemp() in that the
sequence does not have to occur at the very end of the template
and you can pass a @mode. The X string will be modified to form
the name of a directory that didn't exist. The string should be
in the GLib file name encoding. Most importantly, on Windows it
should be in UTF-8.
*/
func MkdtempFull(tmpl string, mode int) (return__ string) {
	__cgo__tmpl := (*C.gchar)(unsafe.Pointer(C.CString(tmpl)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_mkdtemp_full(__cgo__tmpl, C.gint(mode))
	C.free(unsafe.Pointer(__cgo__tmpl))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Opens a temporary file. See the mkstemp() documentation
on most UNIX-like systems.

The parameter is a string that should follow the rules for
mkstemp() templates, i.e. contain the string "XXXXXX".
g_mkstemp() is slightly more flexible than mkstemp() in that the
sequence does not have to occur at the very end of the template.
The X string will be modified to form the name of a file that
didn't exist. The string should be in the GLib file name encoding.
Most importantly, on Windows it should be in UTF-8.
*/
func Mkstemp(tmpl string) (return__ int) {
	__cgo__tmpl := (*C.gchar)(unsafe.Pointer(C.CString(tmpl)))
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_mkstemp(__cgo__tmpl)
	C.free(unsafe.Pointer(__cgo__tmpl))
	return__ = int(__cgo__return__)
	return
}

/*
Opens a temporary file. See the mkstemp() documentation
on most UNIX-like systems.

The parameter is a string that should follow the rules for
mkstemp() templates, i.e. contain the string "XXXXXX".
g_mkstemp_full() is slightly more flexible than mkstemp()
in that the sequence does not have to occur at the very end of the
template and you can pass a @mode and additional @flags. The X
string will be modified to form the name of a file that didn't exist.
The string should be in the GLib file name encoding. Most importantly,
on Windows it should be in UTF-8.
*/
func MkstempFull(tmpl string, flags int, mode int) (return__ int) {
	__cgo__tmpl := (*C.gchar)(unsafe.Pointer(C.CString(tmpl)))
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_mkstemp_full(__cgo__tmpl, C.gint(flags), C.gint(mode))
	C.free(unsafe.Pointer(__cgo__tmpl))
	return__ = int(__cgo__return__)
	return
}

/*
Set the pointer at the specified location to %NULL.
*/
func NullifyPointer(nullify_location *C.gpointer) {
	C.g_nullify_pointer(nullify_location)
	return
}

/*
Prompts the user with
`[E]xit, [H]alt, show [S]tack trace or [P]roceed`.
This function is intended to be used for debugging use only.
The following example shows how it can be used together with
the g_log() functions.

|[<!-- language="C" -->
#include <glib.h>

static void
log_handler (const gchar   *log_domain,
             GLogLevelFlags log_level,
             const gchar   *message,
             gpointer       user_data)
{
  g_log_default_handler (log_domain, log_level, message, user_data);

  g_on_error_query (MY_PROGRAM_NAME);
}

int
main (int argc, char *argv[])
{
  g_log_set_handler (MY_LOG_DOMAIN,
                     G_LOG_LEVEL_WARNING |
                     G_LOG_LEVEL_ERROR |
                     G_LOG_LEVEL_CRITICAL,
                     log_handler,
                     NULL);
  ...
]|

If "[E]xit" is selected, the application terminates with a call
to _exit(0).

If "[S]tack" trace is selected, g_on_error_stack_trace() is called.
This invokes gdb, which attaches to the current process and shows
a stack trace. The prompt is then shown again.

If "[P]roceed" is selected, the function returns.

This function may cause different actions on non-UNIX platforms.
*/
func OnErrorQuery(prg_name string) {
	__cgo__prg_name := (*C.gchar)(unsafe.Pointer(C.CString(prg_name)))
	C.g_on_error_query(__cgo__prg_name)
	C.free(unsafe.Pointer(__cgo__prg_name))
	return
}

/*
Invokes gdb, which attaches to the current process and shows a
stack trace. Called by g_on_error_query() when the "[S]tack trace"
option is selected. You can get the current process's program name
with g_get_prgname(), assuming that you have called gtk_init() or
gdk_init().

This function may cause different actions on non-UNIX platforms.
*/
func OnErrorStackTrace(prg_name string) {
	__cgo__prg_name := (*C.gchar)(unsafe.Pointer(C.CString(prg_name)))
	C.g_on_error_stack_trace(__cgo__prg_name)
	C.free(unsafe.Pointer(__cgo__prg_name))
	return
}

// g_once_init_enter is not generated due to explicit ignore

// g_once_init_leave is not generated due to explicit ignore

func OptionErrorQuark() (return__ C.GQuark) {
	return__ = C.g_option_error_quark()
	return
}

/*
Parses a string containing debugging options
into a %guint containing bit flags. This is used
within GDK and GTK+ to parse the debug options passed on the
command line or through environment variables.

If @string is equal to "all", all flags are set. Any flags
specified along with "all" in @string are inverted; thus,
"all,foo,bar" or "foo,bar,all" sets all flags except those
corresponding to "foo" and "bar".

If @string is equal to "help", all the available keys in @keys
are printed out to standard error.
*/
func ParseDebugString(string_ string, keys *DebugKey, nkeys uint) (return__ uint) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_parse_debug_string(__cgo__string_, (*C.GDebugKey)(unsafe.Pointer(keys)), C.guint(nkeys))
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = uint(__cgo__return__)
	return
}

/*
Gets the last component of the filename.

If @file_name ends with a directory separator it gets the component
before the last slash. If @file_name consists only of directory
separators (and on Windows, possibly a drive letter), a single
separator is returned. If @file_name is empty, it gets ".".
*/
func PathGetBasename(file_name string) (return__ string) {
	__cgo__file_name := (*C.gchar)(unsafe.Pointer(C.CString(file_name)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_path_get_basename(__cgo__file_name)
	C.free(unsafe.Pointer(__cgo__file_name))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the directory components of a file name.

If the file name has no directory components "." is returned.
The returned string should be freed when no longer needed.
*/
func PathGetDirname(file_name string) (return__ string) {
	__cgo__file_name := (*C.gchar)(unsafe.Pointer(C.CString(file_name)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_path_get_dirname(__cgo__file_name)
	C.free(unsafe.Pointer(__cgo__file_name))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns %TRUE if the given @file_name is an absolute file name.
Note that this is a somewhat vague concept on Windows.

On POSIX systems, an absolute file name is well-defined. It always
starts from the single root directory. For example "/usr/local".

On Windows, the concepts of current drive and drive-specific
current directory introduce vagueness. This function interprets as
an absolute file name one that either begins with a directory
separator such as "\Users\tml" or begins with the root on a drive,
for example "C:\Windows". The first case also includes UNC paths
such as "\\myserver\docs\foo". In all cases, either slashes or
backslashes are accepted.

Note that a file name relative to the current drive root does not
truly specify a file uniquely over time and across processes, as
the current drive is a per-process value and can be changed.

File names relative the current directory on some specific drive,
such as "D:foo/bar", are not interpreted as absolute by this
function, but they obviously are not relative to the normal current
directory as returned by getcwd() or g_get_current_dir()
either. Such paths should be avoided, or need to be handled using
Windows-specific code.
*/
func PathIsAbsolute(file_name string) (return__ bool) {
	__cgo__file_name := (*C.gchar)(unsafe.Pointer(C.CString(file_name)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_path_is_absolute(__cgo__file_name)
	C.free(unsafe.Pointer(__cgo__file_name))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns a pointer into @file_name after the root component,
i.e. after the "/" in UNIX or "C:\" under Windows. If @file_name
is not an absolute path it returns %NULL.
*/
func PathSkipRoot(file_name string) (return__ string) {
	__cgo__file_name := (*C.gchar)(unsafe.Pointer(C.CString(file_name)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_path_skip_root(__cgo__file_name)
	C.free(unsafe.Pointer(__cgo__file_name))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Matches a string against a compiled pattern. Passing the correct
length of the string given is mandatory. The reversed string can be
omitted by passing %NULL, this is more efficient if the reversed
version of the string to be matched is not at hand, as
g_pattern_match() will only construct it if the compiled pattern
requires reverse matches.

Note that, if the user code will (possibly) match a string against a
multitude of patterns containing wildcards, chances are high that
some patterns will require a reversed string. In this case, it's
more efficient to provide the reversed string to avoid multiple
constructions thereof in the various calls to g_pattern_match().

Note also that the reverse of a UTF-8 encoded string can in general
not be obtained by g_strreverse(). This works only if the string
does not contain any multibyte characters. GLib offers the
g_utf8_strreverse() function to reverse UTF-8 encoded strings.
*/
func PatternMatch(pspec *PatternSpec, string_length uint, string_ string, string_reversed string) (return__ bool) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	__cgo__string_reversed := (*C.gchar)(unsafe.Pointer(C.CString(string_reversed)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_pattern_match((*C.GPatternSpec)(unsafe.Pointer(pspec)), C.guint(string_length), __cgo__string_, __cgo__string_reversed)
	C.free(unsafe.Pointer(__cgo__string_))
	C.free(unsafe.Pointer(__cgo__string_reversed))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Matches a string against a pattern given as a string. If this
function is to be called in a loop, it's more efficient to compile
the pattern once with g_pattern_spec_new() and call
g_pattern_match_string() repeatedly.
*/
func PatternMatchSimple(pattern string, string_ string) (return__ bool) {
	__cgo__pattern := (*C.gchar)(unsafe.Pointer(C.CString(pattern)))
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_pattern_match_simple(__cgo__pattern, __cgo__string_)
	C.free(unsafe.Pointer(__cgo__pattern))
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Matches a string against a compiled pattern. If the string is to be
matched against more than one pattern, consider using
g_pattern_match() instead while supplying the reversed string.
*/
func PatternMatchString(pspec *PatternSpec, string_ string) (return__ bool) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_pattern_match_string((*C.GPatternSpec)(unsafe.Pointer(pspec)), __cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// g_pointer_bit_lock is not generated due to explicit ignore

// g_pointer_bit_trylock is not generated due to explicit ignore

// g_pointer_bit_unlock is not generated due to explicit ignore

/*
Polls @fds, as with the poll() system call, but portably. (On
systems that don't have poll(), it is emulated using select().)
This is used internally by #GMainContext, but it can be called
directly if you need to block until a file descriptor is ready, but
don't want to run the full main loop.

Each element of @fds is a #GPollFD describing a single file
descriptor to poll. The %fd field indicates the file descriptor,
and the %events field indicates the events to poll for. On return,
the %revents fields will be filled with the events that actually
occurred.

On POSIX systems, the file descriptors in @fds can be any sort of
file descriptor, but the situation is much more complicated on
Windows. If you need to use g_poll() in code that has to run on
Windows, the easiest solution is to construct all of your
#GPollFDs with g_io_channel_win32_make_pollfd().
*/
func Poll(fds *PollFD, nfds uint, timeout int) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_poll((*C.GPollFD)(unsafe.Pointer(fds)), C.guint(nfds), C.gint(timeout))
	return__ = int(__cgo__return__)
	return
}

// g_prefix_error is not generated due to varargs

// g_print is not generated due to varargs

// g_printerr is not generated due to varargs

// g_printf is not generated due to varargs

// g_printf_string_upper_bound is not generated due to varargs

/*
If @dest is %NULL, free @src; otherwise, moves @src into *@dest.
The error variable @dest points to must be %NULL.
*/
func PropagateError(dest **C.GError, src *Error) {
	C.g_propagate_error(dest, (*C.GError)(unsafe.Pointer(src)))
	return
}

// g_propagate_prefixed_error is not generated due to varargs

/*
This is just like the standard C qsort() function, but
the comparison routine accepts a user data argument.

This is guaranteed to be a stable sort since version 2.32.
*/
func QsortWithData(pbase unsafe.Pointer, total_elems int, size int64, compare_func C.GCompareDataFunc, user_data unsafe.Pointer) {
	C.g_qsort_with_data((C.gconstpointer)(pbase), C.gint(total_elems), C.gsize(size), compare_func, (C.gpointer)(user_data))
	return
}

/*
Gets the #GQuark identifying the given (static) string. If the
string does not currently have an associated #GQuark, a new #GQuark
is created, linked to the given string.

Note that this function is identical to g_quark_from_string() except
that if a new #GQuark is created the string itself is used rather
than a copy. This saves memory, but can only be used if the string
will continue to exist until the program terminates. It can be used
with statically allocated strings in the main program, but not with
statically allocated memory in dynamically loaded modules, if you
expect to ever unload the module again (e.g. do not use this
function in GTK+ theme engines).
*/
func QuarkFromStaticString(string_ string) (return__ C.GQuark) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	return__ = C.g_quark_from_static_string(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return
}

/*
Gets the #GQuark identifying the given string. If the string does
not currently have an associated #GQuark, a new #GQuark is created,
using a copy of the string.
*/
func QuarkFromString(string_ string) (return__ C.GQuark) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	return__ = C.g_quark_from_string(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return
}

/*
Gets the string associated with the given #GQuark.
*/
func QuarkToString(quark C.GQuark) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_quark_to_string(quark)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the #GQuark associated with the given string, or 0 if string is
%NULL or it has no associated #GQuark.

If you want the GQuark to be created if it doesn't already exist,
use g_quark_from_string() or g_quark_from_static_string().
*/
func QuarkTryString(string_ string) (return__ C.GQuark) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	return__ = C.g_quark_try_string(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return
}

/*
Returns a random #gdouble equally distributed over the range [0..1).
*/
func RandomDouble() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.g_random_double()
	return__ = float64(__cgo__return__)
	return
}

/*
Returns a random #gdouble equally distributed over the range
[@begin..@end).
*/
func RandomDoubleRange(begin float64, end float64) (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.g_random_double_range(C.gdouble(begin), C.gdouble(end))
	return__ = float64(__cgo__return__)
	return
}

/*
Return a random #guint32 equally distributed over the range
[0..2^32-1].
*/
func RandomInt() (return__ uint32) {
	var __cgo__return__ C.guint32
	__cgo__return__ = C.g_random_int()
	return__ = uint32(__cgo__return__)
	return
}

/*
Returns a random #gint32 equally distributed over the range
[@begin..@end-1].
*/
func RandomIntRange(begin int32, end int32) (return__ int32) {
	var __cgo__return__ C.gint32
	__cgo__return__ = C.g_random_int_range(C.gint32(begin), C.gint32(end))
	return__ = int32(__cgo__return__)
	return
}

/*
Sets the seed for the global random number generator, which is used
by the g_random_* functions, to @seed.
*/
func RandomSetSeed(seed uint32) {
	C.g_random_set_seed(C.guint32(seed))
	return
}

/*
Reallocates the memory pointed to by @mem, so that it now has space for
@n_bytes bytes of memory. It returns the new address of the memory, which may
have been moved. @mem may be %NULL, in which case it's considered to
have zero-length. @n_bytes may be 0, in which case %NULL will be returned
and @mem will be freed unless it is %NULL.
*/
func Realloc(mem unsafe.Pointer, n_bytes int64) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_realloc((C.gpointer)(mem), C.gsize(n_bytes))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
This function is similar to g_realloc(), allocating (@n_blocks * @n_block_bytes) bytes,
but care is taken to detect possible overflow during multiplication.
*/
func ReallocN(mem unsafe.Pointer, n_blocks int64, n_block_bytes int64) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_realloc_n((C.gpointer)(mem), C.gsize(n_blocks), C.gsize(n_block_bytes))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Checks whether @replacement is a valid replacement string
(see g_regex_replace()), i.e. that all escape sequences in
it are valid.

If @has_references is not %NULL then @replacement is checked
for pattern references. For instance, replacement text 'foo\n'
does not contain references and may be evaluated without information
about actual match, but '\0\1' (whole match followed by first
subpattern) requires valid #GMatchInfo object.
*/
func RegexCheckReplacement(replacement string) (has_references bool, return__ bool, __err__ error) {
	__cgo__replacement := (*C.gchar)(unsafe.Pointer(C.CString(replacement)))
	var __cgo__has_references C.gboolean
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_regex_check_replacement(__cgo__replacement, &__cgo__has_references, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__replacement))
	has_references = __cgo__has_references == C.gboolean(1)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

func RegexErrorQuark() (return__ C.GQuark) {
	return__ = C.g_regex_error_quark()
	return
}

/*
Escapes the nul characters in @string to "\x00".  It can be used
to compile a regex with embedded nul characters.

For completeness, @length can be -1 for a nul-terminated string.
In this case the output string will be of course equal to @string.
*/
func RegexEscapeNul(string_ string, length int) (return__ string) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_regex_escape_nul(__cgo__string_, C.gint(length))
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Escapes the special characters used for regular expressions
in @string, for instance "a.b*c" becomes "a\.b\*c". This
function is useful to dynamically generate regular expressions.

@string can contain nul characters that are replaced with "\0",
in this case remember to specify the correct length of @string
in @length.
*/
func RegexEscapeString(string_ []byte, length int) (return__ string) {
	__header__string_ := (*reflect.SliceHeader)(unsafe.Pointer(&string_))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_regex_escape_string((*C.gchar)(unsafe.Pointer(__header__string_.Data)), C.gint(length))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Scans for a match in @string for @pattern.

This function is equivalent to g_regex_match() but it does not
require to compile the pattern with g_regex_new(), avoiding some
lines of code when you need just to do a match without extracting
substrings, capture counts, and so on.

If this function is to be called on the same @pattern more than
once, it's more efficient to compile the pattern once with
g_regex_new() and then use g_regex_match().
*/
func RegexMatchSimple(pattern string, string_ string, compile_options C.GRegexCompileFlags, match_options C.GRegexMatchFlags) (return__ bool) {
	__cgo__pattern := (*C.gchar)(unsafe.Pointer(C.CString(pattern)))
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_regex_match_simple(__cgo__pattern, __cgo__string_, compile_options, match_options)
	C.free(unsafe.Pointer(__cgo__pattern))
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Breaks the string on the pattern, and returns an array of
the tokens. If the pattern contains capturing parentheses,
then the text for each of the substrings will also be returned.
If the pattern does not match anywhere in the string, then the
whole string is returned as the first token.

This function is equivalent to g_regex_split() but it does
not require to compile the pattern with g_regex_new(), avoiding
some lines of code when you need just to do a split without
extracting substrings, capture counts, and so on.

If this function is to be called on the same @pattern more than
once, it's more efficient to compile the pattern once with
g_regex_new() and then use g_regex_split().

As a special case, the result of splitting the empty string ""
is an empty vector, not a vector containing a single string.
The reason for this special case is that being able to represent
a empty vector is typically more useful than consistent handling
of empty elements. If you do need to represent empty elements,
you'll need to check for the empty string before calling this
function.

A pattern that can match empty strings splits @string into
separate characters wherever it matches the empty string between
characters. For example splitting "ab c" using as a separator
"\s*", you will get "a", "b" and "c".
*/
func RegexSplitSimple(pattern string, string_ string, compile_options C.GRegexCompileFlags, match_options C.GRegexMatchFlags) (return__ []string) {
	__cgo__pattern := (*C.gchar)(unsafe.Pointer(C.CString(pattern)))
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.g_regex_split_simple(__cgo__pattern, __cgo__string_, compile_options, match_options)
	C.free(unsafe.Pointer(__cgo__pattern))
	C.free(unsafe.Pointer(__cgo__string_))
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Resets the cache used for g_get_user_special_dir(), so
that the latest on-disk version is used. Call this only
if you just changed the data on disk yourself.

Due to threadsafety issues this may cause leaking of strings
that were previously returned from g_get_user_special_dir()
that can't be freed. We ensure to only leak the data for
the directories that actually changed value though.
*/
func ReloadUserSpecialDirsCache() {
	C.g_reload_user_special_dirs_cache()
	return
}

func ReturnIfFailWarning(log_domain string, pretty_function string, expression string) {
	__cgo__log_domain := C.CString(log_domain)
	__cgo__pretty_function := C.CString(pretty_function)
	__cgo__expression := C.CString(expression)
	C.g_return_if_fail_warning(__cgo__log_domain, __cgo__pretty_function, __cgo__expression)
	C.free(unsafe.Pointer(__cgo__log_domain))
	C.free(unsafe.Pointer(__cgo__pretty_function))
	C.free(unsafe.Pointer(__cgo__expression))
	return
}

/*
A wrapper for the POSIX rmdir() function. The rmdir() function
deletes a directory from the filesystem.

See your C library manual for more details about how rmdir() works
on your system.
*/
func Rmdir(filename string) (return__ int) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	var __cgo__return__ C.int
	__cgo__return__ = C.g_rmdir(__cgo__filename)
	C.free(unsafe.Pointer(__cgo__filename))
	return__ = int(__cgo__return__)
	return
}

/*
Moves the item pointed to by @src to the position indicated by @dest.
After calling this function @dest will point to the position immediately
after @src. It is allowed for @src and @dest to point into different
sequences.
*/
func SequenceMove(src *SequenceIter, dest *SequenceIter) {
	C.g_sequence_move((*C.GSequenceIter)(unsafe.Pointer(src)), (*C.GSequenceIter)(unsafe.Pointer(dest)))
	return
}

/*
Inserts the (@begin, @end) range at the destination pointed to by ptr.
The @begin and @end iters must point into the same sequence. It is
allowed for @dest to point to a different sequence than the one pointed
into by @begin and @end.

If @dest is NULL, the range indicated by @begin and @end is
removed from the sequence. If @dest iter points to a place within
the (@begin, @end) range, the range does not move.
*/
func SequenceMoveRange(dest *SequenceIter, begin *SequenceIter, end *SequenceIter) {
	C.g_sequence_move_range((*C.GSequenceIter)(unsafe.Pointer(dest)), (*C.GSequenceIter)(unsafe.Pointer(begin)), (*C.GSequenceIter)(unsafe.Pointer(end)))
	return
}

/*
Removes the item pointed to by @iter. It is an error to pass the
end iterator to this function.

If the sequence has a data destroy function associated with it, this
function is called on the data for the removed item.
*/
func SequenceRemove(iter *SequenceIter) {
	C.g_sequence_remove((*C.GSequenceIter)(unsafe.Pointer(iter)))
	return
}

/*
Removes all items in the (@begin, @end) range.

If the sequence has a data destroy function associated with it, this
function is called on the data for the removed items.
*/
func SequenceRemoveRange(begin *SequenceIter, end *SequenceIter) {
	C.g_sequence_remove_range((*C.GSequenceIter)(unsafe.Pointer(begin)), (*C.GSequenceIter)(unsafe.Pointer(end)))
	return
}

/*
Changes the data for the item pointed to by @iter to be @data. If
the sequence has a data destroy function associated with it, that
function is called on the existing data that @iter pointed to.
*/
func SequenceSet(iter *SequenceIter, data unsafe.Pointer) {
	C.g_sequence_set((*C.GSequenceIter)(unsafe.Pointer(iter)), (C.gpointer)(data))
	return
}

/*
Swaps the items pointed to by @a and @b. It is allowed for @a and @b
to point into difference sequences.
*/
func SequenceSwap(a *SequenceIter, b *SequenceIter) {
	C.g_sequence_swap((*C.GSequenceIter)(unsafe.Pointer(a)), (*C.GSequenceIter)(unsafe.Pointer(b)))
	return
}

/*
Sets a human-readable name for the application. This name should be
localized if possible, and is intended for display to the user.
Contrast with g_set_prgname(), which sets a non-localized name.
g_set_prgname() will be called automatically by gtk_init(),
but g_set_application_name() will not.

Note that for thread safety reasons, this function can only
be called once.

The application name will be used in contexts such as error messages,
or when displaying an application's name in the task list.
*/
func SetApplicationName(application_name string) {
	__cgo__application_name := (*C.gchar)(unsafe.Pointer(C.CString(application_name)))
	C.g_set_application_name(__cgo__application_name)
	C.free(unsafe.Pointer(__cgo__application_name))
	return
}

// g_set_error is not generated due to varargs

/*
Does nothing if @err is %NULL; if @err is non-%NULL, then *@err
must be %NULL. A new #GError is created and assigned to *@err.
Unlike g_set_error(), @message is not a printf()-style format string.
Use this function if @message contains text you don't have control over,
that could include printf() escape sequences.
*/
func SetErrorLiteral(err **C.GError, domain C.GQuark, code int, message string) {
	__cgo__message := (*C.gchar)(unsafe.Pointer(C.CString(message)))
	C.g_set_error_literal(err, domain, C.gint(code), __cgo__message)
	C.free(unsafe.Pointer(__cgo__message))
	return
}

/*
Sets the name of the program. This name should not be localized,
in contrast to g_set_application_name().

Note that for thread-safety reasons this function can only be called once.
*/
func SetPrgname(prgname string) {
	__cgo__prgname := (*C.gchar)(unsafe.Pointer(C.CString(prgname)))
	C.g_set_prgname(__cgo__prgname)
	C.free(unsafe.Pointer(__cgo__prgname))
	return
}

/*
Sets the print handler.

Any messages passed to g_print() will be output via
the new handler. The default handler simply outputs
the message to stdout. By providing your own handler
you can redirect the output, to a GTK+ widget or a
log file for example.
*/
func SetPrintHandler(func_ C.GPrintFunc) (return__ C.GPrintFunc) {
	return__ = C.g_set_print_handler(func_)
	return
}

/*
Sets the handler for printing error messages.

Any messages passed to g_printerr() will be output via
the new handler. The default handler simply outputs the
message to stderr. By providing your own handler you can
redirect the output, to a GTK+ widget or a log file for
example.
*/
func SetPrinterrHandler(func_ C.GPrintFunc) (return__ C.GPrintFunc) {
	return__ = C.g_set_printerr_handler(func_)
	return
}

/*
Sets an environment variable. Both the variable's name and value
should be in the GLib file name encoding. On UNIX, this means that
they can be arbitrary byte strings. On Windows, they should be in
UTF-8.

Note that on some systems, when variables are overwritten, the memory
used for the previous variables and its value isn't reclaimed.

You should be mindful fo the fact that environment variable handling
in UNIX is not thread-safe, and your program may crash if one thread
calls g_setenv() while another thread is calling getenv(). (And note
that many functions, such as gettext(), call getenv() internally.)
This function is only safe to use at the very start of your program,
before creating any other threads (or creating objects that create
worker threads of their own).

If you need to set up the environment for a child process, you can
use g_get_environ() to get an environment array, modify that with
g_environ_setenv() and g_environ_unsetenv(), and then pass that
array directly to execvpe(), g_spawn_async(), or the like.
*/
func Setenv(variable string, value string, overwrite bool) (return__ bool) {
	__cgo__variable := (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	__cgo__value := (*C.gchar)(unsafe.Pointer(C.CString(value)))
	__cgo__overwrite := C.gboolean(0)
	if overwrite {
		__cgo__overwrite = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_setenv(__cgo__variable, __cgo__value, __cgo__overwrite)
	C.free(unsafe.Pointer(__cgo__variable))
	C.free(unsafe.Pointer(__cgo__value))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

func ShellErrorQuark() (return__ C.GQuark) {
	return__ = C.g_shell_error_quark()
	return
}

/*
Parses a command line into an argument vector, in much the same way
the shell would, but without many of the expansions the shell would
perform (variable expansion, globs, operators, filename expansion,
etc. are not supported). The results are defined to be the same as
those you would get from a UNIX98 /bin/sh, as long as the input
contains none of the unsupported shell expansions. If the input
does contain such expansions, they are passed through
literally. Possible errors are those from the #G_SHELL_ERROR
domain. Free the returned vector with g_strfreev().
*/
func ShellParseArgv(command_line string) (argcp int, argvp []string, return__ bool, __err__ error) {
	__cgo__command_line := (*C.gchar)(unsafe.Pointer(C.CString(command_line)))
	var __cgo__argcp C.gint
	var __cgo__argvp **C.gchar
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_shell_parse_argv(__cgo__command_line, &__cgo__argcp, &__cgo__argvp, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__command_line))
	argcp = int(__cgo__argcp)
	var __slice__argvp []*C.gchar
	__header__argvp := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__argvp))
	__header__argvp.Len = 4294967296
	__header__argvp.Cap = 4294967296
	__header__argvp.Data = uintptr(unsafe.Pointer(__cgo__argvp))
	for _, p := range __slice__argvp {
		if p == nil {
			break
		}
		argvp = append(argvp, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Quotes a string so that the shell (/bin/sh) will interpret the
quoted string to mean @unquoted_string. If you pass a filename to
the shell, for example, you should first quote it with this
function.  The return value must be freed with g_free(). The
quoting style used is undefined (single or double quotes may be
used).
*/
func ShellQuote(unquoted_string string) (return__ string) {
	__cgo__unquoted_string := (*C.gchar)(unsafe.Pointer(C.CString(unquoted_string)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_shell_quote(__cgo__unquoted_string)
	C.free(unsafe.Pointer(__cgo__unquoted_string))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Unquotes a string as the shell (/bin/sh) would. Only handles
quotes; if a string contains file globs, arithmetic operators,
variables, backticks, redirections, or other special-to-the-shell
features, the result will be different from the result a real shell
would produce (the variables, backticks, etc. will be passed
through literally instead of being expanded). This function is
guaranteed to succeed if applied to the result of
g_shell_quote(). If it fails, it returns %NULL and sets the
error. The @quoted_string need not actually contain quoted or
escaped text; g_shell_unquote() simply goes through the string and
unquotes/unescapes anything that the shell would. Both single and
double quotes are handled, as are escapes including escaped
newlines. The return value must be freed with g_free(). Possible
errors are in the #G_SHELL_ERROR domain.

Shell quoting rules are a bit strange. Single quotes preserve the
literal string exactly. escape sequences are not allowed; not even
\' - if you want a ' in the quoted text, you have to do something
like 'foo'\''bar'.  Double quotes allow $, `, ", \, and newline to
be escaped with backslash. Otherwise double quotes preserve things
literally.
*/
func ShellUnquote(quoted_string string) (return__ string, __err__ error) {
	__cgo__quoted_string := (*C.gchar)(unsafe.Pointer(C.CString(quoted_string)))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_shell_unquote(__cgo__quoted_string, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__quoted_string))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Allocates a block of memory from the slice allocator.
The block adress handed out can be expected to be aligned
to at least 1 * sizeof (void*),
though in general slices are 2 * sizeof (void*) bytes aligned,
if a malloc() fallback implementation is used instead,
the alignment may be reduced in a libc dependent fashion.
Note that the underlying slice allocation mechanism can
be changed with the [`G_SLICE=always-malloc`][G_SLICE]
environment variable.
*/
func SliceAlloc(block_size int64) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_slice_alloc(C.gsize(block_size))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Allocates a block of memory via g_slice_alloc() and initializes
the returned memory to 0. Note that the underlying slice allocation
mechanism can be changed with the [`G_SLICE=always-malloc`][G_SLICE]
environment variable.
*/
func SliceAlloc0(block_size int64) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_slice_alloc0(C.gsize(block_size))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Allocates a block of memory from the slice allocator
and copies @block_size bytes into it from @mem_block.
*/
func SliceCopy(block_size int64, mem_block unsafe.Pointer) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_slice_copy(C.gsize(block_size), (C.gconstpointer)(mem_block))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Frees a block of memory.

The memory must have been allocated via g_slice_alloc() or
g_slice_alloc0() and the @block_size has to match the size
specified upon allocation. Note that the exact release behaviour
can be changed with the [`G_DEBUG=gc-friendly`][G_DEBUG] environment
variable, also see [`G_SLICE`][G_SLICE] for related debugging options.
*/
func SliceFree1(block_size int64, mem_block unsafe.Pointer) {
	C.g_slice_free1(C.gsize(block_size), (C.gpointer)(mem_block))
	return
}

/*
Frees a linked list of memory blocks of structure type @type.

The memory blocks must be equal-sized, allocated via
g_slice_alloc() or g_slice_alloc0() and linked together by a
@next pointer (similar to #GSList). The offset of the @next
field in each block is passed as third argument.
Note that the exact release behaviour can be changed with the
[`G_DEBUG=gc-friendly`][G_DEBUG] environment variable, also see
[`G_SLICE`][G_SLICE] for related debugging options.
*/
func SliceFreeChainWithOffset(block_size int64, mem_chain unsafe.Pointer, next_offset int64) {
	C.g_slice_free_chain_with_offset(C.gsize(block_size), (C.gpointer)(mem_chain), C.gsize(next_offset))
	return
}

// g_slice_get_config is not generated due to explicit ignore

// g_slice_get_config_state is not generated due to explicit ignore

// g_slice_set_config is not generated due to explicit ignore

// g_snprintf is not generated due to varargs

/*
Removes the source with the given id from the default main context.

The id of a #GSource is given by g_source_get_id(), or will be
returned by the functions g_source_attach(), g_idle_add(),
g_idle_add_full(), g_timeout_add(), g_timeout_add_full(),
g_child_watch_add(), g_child_watch_add_full(), g_io_add_watch(), and
g_io_add_watch_full().

See also g_source_destroy(). You must use g_source_destroy() for sources
added to a non-default main context.

It is a programmer error to attempt to remove a non-existent source.
*/
func SourceRemove(tag uint) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_source_remove(C.guint(tag))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Removes a source from the default main loop context given the
source functions and user data. If multiple sources exist with the
same source functions and user data, only one will be destroyed.
*/
func SourceRemoveByFuncsUserData(funcs *SourceFuncs, user_data unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_source_remove_by_funcs_user_data((*C.GSourceFuncs)(unsafe.Pointer(funcs)), (C.gpointer)(user_data))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Removes a source from the default main loop context given the user
data for the callback. If multiple sources exist with the same user
data, only one will be destroyed.
*/
func SourceRemoveByUserData(user_data unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_source_remove_by_user_data((C.gpointer)(user_data))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the name of a source using its ID.

This is a convenience utility to set source names from the return
value of g_idle_add(), g_timeout_add(), etc.
*/
func SourceSetNameById(tag uint, name string) {
	__cgo__name := C.CString(name)
	C.g_source_set_name_by_id(C.guint(tag), __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Gets the smallest prime number from a built-in array of primes which
is larger than @num. This is used within GLib to calculate the optimum
size of a #GHashTable.

The built-in array of primes ranges from 11 to 13845163 such that
each prime is approximately 1.5-2 times the previous prime.
*/
func SpacedPrimesClosest(num uint) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_spaced_primes_closest(C.guint(num))
	return__ = uint(__cgo__return__)
	return
}

/*
See g_spawn_async_with_pipes() for a full description; this function
simply calls the g_spawn_async_with_pipes() without any pipes.

You should call g_spawn_close_pid() on the returned child process
reference when you don't need it any more.

If you are writing a GTK+ application, and the program you are
spawning is a graphical application, too, then you may want to
use gdk_spawn_on_screen() instead to ensure that the spawned program
opens its windows on the right screen.

Note that the returned @child_pid on Windows is a handle to the child
process and not its identifier. Process handles and process identifiers
are different concepts on Windows.
*/
func SpawnAsync(working_directory string, argv []string, envp []string, flags C.GSpawnFlags, child_setup C.GSpawnChildSetupFunc, user_data unsafe.Pointer) (child_pid C.GPid, return__ bool, __err__ error) {
	__cgo__working_directory := (*C.gchar)(unsafe.Pointer(C.CString(working_directory)))
	__header__argv := (*reflect.SliceHeader)(unsafe.Pointer(&argv))
	__header__envp := (*reflect.SliceHeader)(unsafe.Pointer(&envp))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_spawn_async(__cgo__working_directory, (**C.gchar)(unsafe.Pointer(__header__argv.Data)), (**C.gchar)(unsafe.Pointer(__header__envp.Data)), flags, child_setup, (C.gpointer)(user_data), &child_pid, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__working_directory))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Executes a child program asynchronously (your program will not
block waiting for the child to exit). The child program is
specified by the only argument that must be provided, @argv.
@argv should be a %NULL-terminated array of strings, to be passed
as the argument vector for the child. The first string in @argv
is of course the name of the program to execute. By default, the
name of the program must be a full path. If @flags contains the
%G_SPAWN_SEARCH_PATH flag, the `PATH` environment variable is
used to search for the executable. If @flags contains the
%G_SPAWN_SEARCH_PATH_FROM_ENVP flag, the `PATH` variable from
@envp is used to search for the executable. If both the
%G_SPAWN_SEARCH_PATH and %G_SPAWN_SEARCH_PATH_FROM_ENVP flags
are set, the `PATH` variable from @envp takes precedence over
the environment variable.

If the program name is not a full path and %G_SPAWN_SEARCH_PATH flag is not
used, then the program will be run from the current directory (or
@working_directory, if specified); this might be unexpected or even
dangerous in some cases when the current directory is world-writable.

On Windows, note that all the string or string vector arguments to
this function and the other g_spawn*() functions are in UTF-8, the
GLib file name encoding. Unicode characters that are not part of
the system codepage passed in these arguments will be correctly
available in the spawned program only if it uses wide character API
to retrieve its command line. For C programs built with Microsoft's
tools it is enough to make the program have a wmain() instead of
main(). wmain() has a wide character argument vector as parameter.

At least currently, mingw doesn't support wmain(), so if you use
mingw to develop the spawned program, it will have to call the
undocumented function __wgetmainargs() to get the wide character
argument vector and environment. See gspawn-win32-helper.c in the
GLib sources or init.c in the mingw runtime sources for a prototype
for that function. Alternatively, you can retrieve the Win32 system
level wide character command line passed to the spawned program
using the GetCommandLineW() function.

On Windows the low-level child process creation API CreateProcess()
doesn't use argument vectors, but a command line. The C runtime
library's spawn*() family of functions (which g_spawn_async_with_pipes()
eventually calls) paste the argument vector elements together into
a command line, and the C runtime startup code does a corresponding
reconstruction of an argument vector from the command line, to be
passed to main(). Complications arise when you have argument vector
elements that contain spaces of double quotes. The spawn*() functions
don't do any quoting or escaping, but on the other hand the startup
code does do unquoting and unescaping in order to enable receiving
arguments with embedded spaces or double quotes. To work around this
asymmetry, g_spawn_async_with_pipes() will do quoting and escaping on
argument vector elements that need it before calling the C runtime
spawn() function.

The returned @child_pid on Windows is a handle to the child
process, not its identifier. Process handles and process
identifiers are different concepts on Windows.

@envp is a %NULL-terminated array of strings, where each string
has the form `KEY=VALUE`. This will become the child's environment.
If @envp is %NULL, the child inherits its parent's environment.

@flags should be the bitwise OR of any flags you want to affect the
function's behaviour. The %G_SPAWN_DO_NOT_REAP_CHILD means that the
child will not automatically be reaped; you must use a child watch to
be notified about the death of the child process. Eventually you must
call g_spawn_close_pid() on the @child_pid, in order to free
resources which may be associated with the child process. (On Unix,
using a child watch is equivalent to calling waitpid() or handling
the %SIGCHLD signal manually. On Windows, calling g_spawn_close_pid()
is equivalent to calling CloseHandle() on the process handle returned
in @child_pid). See g_child_watch_add().

%G_SPAWN_LEAVE_DESCRIPTORS_OPEN means that the parent's open file
descriptors will be inherited by the child; otherwise all descriptors
except stdin/stdout/stderr will be closed before calling exec() in
the child. %G_SPAWN_SEARCH_PATH means that @argv[0] need not be an
absolute path, it will be looked for in the `PATH` environment
variable. %G_SPAWN_SEARCH_PATH_FROM_ENVP means need not be an
absolute path, it will be looked for in the `PATH` variable from
@envp. If both %G_SPAWN_SEARCH_PATH and %G_SPAWN_SEARCH_PATH_FROM_ENVP
are used, the value from @envp takes precedence over the environment.
%G_SPAWN_STDOUT_TO_DEV_NULL means that the child's standard output
will be discarded, instead of going to the same location as the parent's
standard output. If you use this flag, @standard_output must be %NULL.
%G_SPAWN_STDERR_TO_DEV_NULL means that the child's standard error
will be discarded, instead of going to the same location as the parent's
standard error. If you use this flag, @standard_error must be %NULL.
%G_SPAWN_CHILD_INHERITS_STDIN means that the child will inherit the parent's
standard input (by default, the child's standard input is attached to
/dev/null). If you use this flag, @standard_input must be %NULL.
%G_SPAWN_FILE_AND_ARGV_ZERO means that the first element of @argv is
the file to execute, while the remaining elements are the actual
argument vector to pass to the file. Normally g_spawn_async_with_pipes()
uses @argv[0] as the file to execute, and passes all of @argv to the child.

@child_setup and @user_data are a function and user data. On POSIX
platforms, the function is called in the child after GLib has
performed all the setup it plans to perform (including creating
pipes, closing file descriptors, etc.) but before calling exec().
That is, @child_setup is called just before calling exec() in the
child. Obviously actions taken in this function will only affect
the child, not the parent.

On Windows, there is no separate fork() and exec() functionality.
Child processes are created and run with a single API call,
CreateProcess(). There is no sensible thing @child_setup
could be used for on Windows so it is ignored and not called.

If non-%NULL, @child_pid will on Unix be filled with the child's
process ID. You can use the process ID to send signals to the child,
or to use g_child_watch_add() (or waitpid()) if you specified the
%G_SPAWN_DO_NOT_REAP_CHILD flag. On Windows, @child_pid will be
filled with a handle to the child process only if you specified the
%G_SPAWN_DO_NOT_REAP_CHILD flag. You can then access the child
process using the Win32 API, for example wait for its termination
with the WaitFor*() functions, or examine its exit code with
GetExitCodeProcess(). You should close the handle with CloseHandle()
or g_spawn_close_pid() when you no longer need it.

If non-%NULL, the @standard_input, @standard_output, @standard_error
locations will be filled with file descriptors for writing to the child's
standard input or reading from its standard output or standard error.
The caller of g_spawn_async_with_pipes() must close these file descriptors
when they are no longer in use. If these parameters are %NULL, the
corresponding pipe won't be created.

If @standard_input is NULL, the child's standard input is attached to
/dev/null unless %G_SPAWN_CHILD_INHERITS_STDIN is set.

If @standard_error is NULL, the child's standard error goes to the same
location as the parent's standard error unless %G_SPAWN_STDERR_TO_DEV_NULL
is set.

If @standard_output is NULL, the child's standard output goes to the same
location as the parent's standard output unless %G_SPAWN_STDOUT_TO_DEV_NULL
is set.

@error can be %NULL to ignore errors, or non-%NULL to report errors.
If an error is set, the function returns %FALSE. Errors are reported
even if they occur in the child (for example if the executable in
@argv[0] is not found). Typically the `message` field of returned
errors should be displayed to users. Possible errors are those from
the #G_SPAWN_ERROR domain.

If an error occurs, @child_pid, @standard_input, @standard_output,
and @standard_error will not be filled with valid values.

If @child_pid is not %NULL and an error does not occur then the returned
process reference must be closed using g_spawn_close_pid().

If you are writing a GTK+ application, and the program you
are spawning is a graphical application, too, then you may
want to use gdk_spawn_on_screen_with_pipes() instead to ensure that
the spawned program opens its windows on the right screen.
*/
func SpawnAsyncWithPipes(working_directory string, argv []string, envp []string, flags C.GSpawnFlags, child_setup C.GSpawnChildSetupFunc, user_data unsafe.Pointer) (child_pid C.GPid, standard_input int, standard_output int, standard_error int, return__ bool, __err__ error) {
	__cgo__working_directory := (*C.gchar)(unsafe.Pointer(C.CString(working_directory)))
	__header__argv := (*reflect.SliceHeader)(unsafe.Pointer(&argv))
	__header__envp := (*reflect.SliceHeader)(unsafe.Pointer(&envp))
	var __cgo__standard_input C.gint
	var __cgo__standard_output C.gint
	var __cgo__standard_error C.gint
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_spawn_async_with_pipes(__cgo__working_directory, (**C.gchar)(unsafe.Pointer(__header__argv.Data)), (**C.gchar)(unsafe.Pointer(__header__envp.Data)), flags, child_setup, (C.gpointer)(user_data), &child_pid, &__cgo__standard_input, &__cgo__standard_output, &__cgo__standard_error, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__working_directory))
	standard_input = int(__cgo__standard_input)
	standard_output = int(__cgo__standard_output)
	standard_error = int(__cgo__standard_error)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Set @error if @exit_status indicates the child exited abnormally
(e.g. with a nonzero exit code, or via a fatal signal).

The g_spawn_sync() and g_child_watch_add() family of APIs return an
exit status for subprocesses encoded in a platform-specific way.
On Unix, this is guaranteed to be in the same format waitpid() returns,
and on Windows it is guaranteed to be the result of GetExitCodeProcess().

Prior to the introduction of this function in GLib 2.34, interpreting
@exit_status required use of platform-specific APIs, which is problematic
for software using GLib as a cross-platform layer.

Additionally, many programs simply want to determine whether or not
the child exited successfully, and either propagate a #GError or
print a message to standard error. In that common case, this function
can be used. Note that the error message in @error will contain
human-readable information about the exit status.

The @domain and @code of @error have special semantics in the case
where the process has an "exit code", as opposed to being killed by
a signal. On Unix, this happens if WIFEXITED() would be true of
@exit_status. On Windows, it is always the case.

The special semantics are that the actual exit code will be the
code set in @error, and the domain will be %G_SPAWN_EXIT_ERROR.
This allows you to differentiate between different exit codes.

If the process was terminated by some means other than an exit
status, the domain will be %G_SPAWN_ERROR, and the code will be
%G_SPAWN_ERROR_FAILED.

This function just offers convenience; you can of course also check
the available platform via a macro such as %G_OS_UNIX, and use
WIFEXITED() and WEXITSTATUS() on @exit_status directly. Do not attempt
to scan or parse the error message string; it may be translated and/or
change in future versions of GLib.
*/
func SpawnCheckExitStatus(exit_status int) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_spawn_check_exit_status(C.gint(exit_status), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
On some platforms, notably Windows, the #GPid type represents a resource
which must be closed to prevent resource leaking. g_spawn_close_pid()
is provided for this purpose. It should be used on all platforms, even
though it doesn't do anything under UNIX.
*/
func SpawnClosePid(pid C.GPid) {
	C.g_spawn_close_pid(pid)
	return
}

/*
A simple version of g_spawn_async() that parses a command line with
g_shell_parse_argv() and passes it to g_spawn_async(). Runs a
command line in the background. Unlike g_spawn_async(), the
%G_SPAWN_SEARCH_PATH flag is enabled, other flags are not. Note
that %G_SPAWN_SEARCH_PATH can have security implications, so
consider using g_spawn_async() directly if appropriate. Possible
errors are those from g_shell_parse_argv() and g_spawn_async().

The same concerns on Windows apply as for g_spawn_command_line_sync().
*/
func SpawnCommandLineAsync(command_line string) (return__ bool, __err__ error) {
	__cgo__command_line := (*C.gchar)(unsafe.Pointer(C.CString(command_line)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_spawn_command_line_async(__cgo__command_line, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__command_line))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
A simple version of g_spawn_sync() with little-used parameters
removed, taking a command line instead of an argument vector.  See
g_spawn_sync() for full details. @command_line will be parsed by
g_shell_parse_argv(). Unlike g_spawn_sync(), the %G_SPAWN_SEARCH_PATH flag
is enabled. Note that %G_SPAWN_SEARCH_PATH can have security
implications, so consider using g_spawn_sync() directly if
appropriate. Possible errors are those from g_spawn_sync() and those
from g_shell_parse_argv().

If @exit_status is non-%NULL, the platform-specific exit status of
the child is stored there; see the documentation of
g_spawn_check_exit_status() for how to use and interpret this.

On Windows, please note the implications of g_shell_parse_argv()
parsing @command_line. Parsing is done according to Unix shell rules, not
Windows command interpreter rules.
Space is a separator, and backslashes are
special. Thus you cannot simply pass a @command_line containing
canonical Windows paths, like "c:\\program files\\app\\app.exe", as
the backslashes will be eaten, and the space will act as a
separator. You need to enclose such paths with single quotes, like
"'c:\\program files\\app\\app.exe' 'e:\\folder\\argument.txt'".
*/
func SpawnCommandLineSync(command_line string) (standard_output *C.gchar, standard_error *C.gchar, exit_status int, return__ bool, __err__ error) {
	__cgo__command_line := (*C.gchar)(unsafe.Pointer(C.CString(command_line)))
	var __cgo__exit_status C.gint
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_spawn_command_line_sync(__cgo__command_line, &standard_output, &standard_error, &__cgo__exit_status, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__command_line))
	exit_status = int(__cgo__exit_status)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

func SpawnErrorQuark() (return__ C.GQuark) {
	return__ = C.g_spawn_error_quark()
	return
}

func SpawnExitErrorQuark() (return__ C.GQuark) {
	return__ = C.g_spawn_exit_error_quark()
	return
}

/*
Executes a child synchronously (waits for the child to exit before returning).
All output from the child is stored in @standard_output and @standard_error,
if those parameters are non-%NULL. Note that you must set the
%G_SPAWN_STDOUT_TO_DEV_NULL and %G_SPAWN_STDERR_TO_DEV_NULL flags when
passing %NULL for @standard_output and @standard_error.

If @exit_status is non-%NULL, the platform-specific exit status of
the child is stored there; see the documentation of
g_spawn_check_exit_status() for how to use and interpret this.
Note that it is invalid to pass %G_SPAWN_DO_NOT_REAP_CHILD in
@flags.

If an error occurs, no data is returned in @standard_output,
@standard_error, or @exit_status.

This function calls g_spawn_async_with_pipes() internally; see that
function for full details on the other parameters and details on
how these functions work on Windows.
*/
func SpawnSync(working_directory string, argv []string, envp []string, flags C.GSpawnFlags, child_setup C.GSpawnChildSetupFunc, user_data unsafe.Pointer) (standard_output *C.gchar, standard_error *C.gchar, exit_status int, return__ bool, __err__ error) {
	__cgo__working_directory := (*C.gchar)(unsafe.Pointer(C.CString(working_directory)))
	__header__argv := (*reflect.SliceHeader)(unsafe.Pointer(&argv))
	__header__envp := (*reflect.SliceHeader)(unsafe.Pointer(&envp))
	var __cgo__exit_status C.gint
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_spawn_sync(__cgo__working_directory, (**C.gchar)(unsafe.Pointer(__header__argv.Data)), (**C.gchar)(unsafe.Pointer(__header__envp.Data)), flags, child_setup, (C.gpointer)(user_data), &standard_output, &standard_error, &__cgo__exit_status, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__working_directory))
	exit_status = int(__cgo__exit_status)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

// g_sprintf is not generated due to varargs

/*
Copies a nul-terminated string into the dest buffer, include the
trailing nul, and return a pointer to the trailing nul byte.
This is useful for concatenating multiple strings together
without having to repeatedly scan for the end.
*/
func Stpcpy(dest string, src string) (return__ string) {
	__cgo__dest := (*C.gchar)(unsafe.Pointer(C.CString(dest)))
	__cgo__src := C.CString(src)
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_stpcpy(__cgo__dest, __cgo__src)
	C.free(unsafe.Pointer(__cgo__dest))
	C.free(unsafe.Pointer(__cgo__src))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Compares two strings for byte-by-byte equality and returns %TRUE
if they are equal. It can be passed to g_hash_table_new() as the
@key_equal_func parameter, when using non-%NULL strings as keys in a
#GHashTable.

Note that this function is primarily meant as a hash table comparison
function. For a general-purpose, %NULL-safe string comparison function,
see g_strcmp0().
*/
func StrEqual(v1 unsafe.Pointer, v2 unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_str_equal((C.gconstpointer)(v1), (C.gconstpointer)(v2))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Looks whether the string @str begins with @prefix.
*/
func StrHasPrefix(str string, prefix string) (return__ bool) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	__cgo__prefix := (*C.gchar)(unsafe.Pointer(C.CString(prefix)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_str_has_prefix(__cgo__str, __cgo__prefix)
	C.free(unsafe.Pointer(__cgo__str))
	C.free(unsafe.Pointer(__cgo__prefix))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Looks whether the string @str ends with @suffix.
*/
func StrHasSuffix(str string, suffix string) (return__ bool) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	__cgo__suffix := (*C.gchar)(unsafe.Pointer(C.CString(suffix)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_str_has_suffix(__cgo__str, __cgo__suffix)
	C.free(unsafe.Pointer(__cgo__str))
	C.free(unsafe.Pointer(__cgo__suffix))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Converts a string to a hash value.

This function implements the widely used "djb" hash apparently
posted by Daniel Bernstein to comp.lang.c some time ago.  The 32
bit unsigned hash value starts at 5381 and for each byte 'c' in
the string, is updated: `hash = hash * 33 + c`. This function
uses the signed value of each byte.

It can be passed to g_hash_table_new() as the @hash_func parameter,
when using non-%NULL strings as keys in a #GHashTable.
*/
func StrHash(v unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_str_hash((C.gconstpointer)(v))
	return__ = uint(__cgo__return__)
	return
}

/*
Determines if a string is pure ASCII. A string is pure ASCII if it
contains no bytes with the high bit set.
*/
func StrIsAscii(str string) (return__ bool) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_str_is_ascii(__cgo__str)
	C.free(unsafe.Pointer(__cgo__str))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if a search conducted for @search_term should match
@potential_hit.

This function calls g_str_tokenize_and_fold() on both
@search_term and @potential_hit.  ASCII alternates are never taken
for @search_term but will be taken for @potential_hit according to
the value of @accept_alternates.

A hit occurs when each folded token in @search_term is a prefix of a
folded token from @potential_hit.

Depending on how you're performing the search, it will typically be
faster to call g_str_tokenize_and_fold() on each string in
your corpus and build an index on the returned folded tokens, then
call g_str_tokenize_and_fold() on the search term and
perform lookups into that index.

As some examples, searching for "fred" would match the potential hit
"Smith, Fred" and also "Frédéric".  Searching for "Fréd" would match
"Frédéric" but not "Frederic" (due to the one-directional nature of
accent matching).  Searching "fo" would match "Foo" and "Bar Foo
Baz", but not "SFO" (because no word as "fo" as a prefix).
*/
func StrMatchString(search_term string, potential_hit string, accept_alternates bool) (return__ bool) {
	__cgo__search_term := (*C.gchar)(unsafe.Pointer(C.CString(search_term)))
	__cgo__potential_hit := (*C.gchar)(unsafe.Pointer(C.CString(potential_hit)))
	__cgo__accept_alternates := C.gboolean(0)
	if accept_alternates {
		__cgo__accept_alternates = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_str_match_string(__cgo__search_term, __cgo__potential_hit, __cgo__accept_alternates)
	C.free(unsafe.Pointer(__cgo__search_term))
	C.free(unsafe.Pointer(__cgo__potential_hit))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Transliterate @str to plain ASCII.

For best results, @str should be in composed normalised form.

This function performs a reasonably good set of character
replacements.  The particular set of replacements that is done may
change by version or even by runtime environment.

If the source language of @str is known, it can used to improve the
accuracy of the translation by passing it as @from_locale.  It should
be a valid POSIX locale string (of the form
"language[_territory][.codeset][@modifier]").

If @from_locale is %NULL then the current locale is used.

If you want to do translation for no specific locale, and you want it
to be done independently of the currently locale, specify "C" for
@from_locale.
*/
func StrToAscii(str string, from_locale string) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	__cgo__from_locale := (*C.gchar)(unsafe.Pointer(C.CString(from_locale)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_str_to_ascii(__cgo__str, __cgo__from_locale)
	C.free(unsafe.Pointer(__cgo__str))
	C.free(unsafe.Pointer(__cgo__from_locale))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Tokenises @string and performs folding on each token.

A token is a non-empty sequence of alphanumeric characters in the
source string, separated by non-alphanumeric characters.  An
"alphanumeric" character for this purpose is one that matches
g_unichar_isalnum() or g_unichar_ismark().

Each token is then (Unicode) normalised and case-folded.  If
@ascii_alternates is non-%NULL and some of the returned tokens
contain non-ASCII characters, ASCII alternatives will be generated.

The number of ASCII alternatives that are generated and the method
for doing so is unspecified, but @translit_locale (if specified) may
improve the transliteration if the language of the source string is
known.
*/
func StrTokenizeAndFold(string_ string, translit_locale string) (ascii_alternates []string, return__ []string) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	__cgo__translit_locale := (*C.gchar)(unsafe.Pointer(C.CString(translit_locale)))
	var __cgo__ascii_alternates **C.gchar
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.g_str_tokenize_and_fold(__cgo__string_, __cgo__translit_locale, &__cgo__ascii_alternates)
	C.free(unsafe.Pointer(__cgo__string_))
	C.free(unsafe.Pointer(__cgo__translit_locale))
	var __slice__ascii_alternates []*C.gchar
	__header__ascii_alternates := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__ascii_alternates))
	__header__ascii_alternates.Len = 4294967296
	__header__ascii_alternates.Cap = 4294967296
	__header__ascii_alternates.Data = uintptr(unsafe.Pointer(__cgo__ascii_alternates))
	for _, p := range __slice__ascii_alternates {
		if p == nil {
			break
		}
		ascii_alternates = append(ascii_alternates, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
For each character in @string, if the character is not in @valid_chars,
replaces the character with @substitutor. Modifies @string in place,
and return @string itself, not a copy. The return value is to allow
nesting such as
|[<!-- language="C" -->
  g_ascii_strup (g_strcanon (str, "abc", '?'))
]|
*/
func Strcanon(string_ string, valid_chars string, substitutor byte) (return__ string) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	__cgo__valid_chars := (*C.gchar)(unsafe.Pointer(C.CString(valid_chars)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strcanon(__cgo__string_, __cgo__valid_chars, C.gchar(substitutor))
	C.free(unsafe.Pointer(__cgo__string_))
	C.free(unsafe.Pointer(__cgo__valid_chars))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

// g_strcasecmp is not generated due to deprecation attr

/*
Removes trailing whitespace from a string.

This function doesn't allocate or reallocate any memory;
it modifies @string in place. Therefore, it cannot be used
on statically allocated strings.

The pointer to @string is returned to allow the nesting of functions.

Also see g_strchug() and g_strstrip().
*/
func Strchomp(string_ string) (return__ string) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strchomp(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Removes leading whitespace from a string, by moving the rest
of the characters forward.

This function doesn't allocate or reallocate any memory;
it modifies @string in place. Therefore, it cannot be used on
statically allocated strings.

The pointer to @string is returned to allow the nesting of functions.

Also see g_strchomp() and g_strstrip().
*/
func Strchug(string_ string) (return__ string) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strchug(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Compares @str1 and @str2 like strcmp(). Handles %NULL
gracefully by sorting it before non-%NULL strings.
Comparing two %NULL pointers returns 0.
*/
func Strcmp0(str1 string, str2 string) (return__ int) {
	__cgo__str1 := C.CString(str1)
	__cgo__str2 := C.CString(str2)
	var __cgo__return__ C.int
	__cgo__return__ = C.g_strcmp0(__cgo__str1, __cgo__str2)
	C.free(unsafe.Pointer(__cgo__str1))
	C.free(unsafe.Pointer(__cgo__str2))
	return__ = int(__cgo__return__)
	return
}

/*
Replaces all escaped characters with their one byte equivalent.

This function does the reverse conversion of g_strescape().
*/
func Strcompress(source string) (return__ string) {
	__cgo__source := (*C.gchar)(unsafe.Pointer(C.CString(source)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strcompress(__cgo__source)
	C.free(unsafe.Pointer(__cgo__source))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

// g_strconcat is not generated due to varargs

/*
Converts any delimiter characters in @string to @new_delimiter.
Any characters in @string which are found in @delimiters are
changed to the @new_delimiter character. Modifies @string in place,
and returns @string itself, not a copy. The return value is to
allow nesting such as
|[<!-- language="C" -->
  g_ascii_strup (g_strdelimit (str, "abc", '?'))
]|
*/
func Strdelimit(string_ string, delimiters string, new_delimiter byte) (return__ string) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	__cgo__delimiters := (*C.gchar)(unsafe.Pointer(C.CString(delimiters)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strdelimit(__cgo__string_, __cgo__delimiters, C.gchar(new_delimiter))
	C.free(unsafe.Pointer(__cgo__string_))
	C.free(unsafe.Pointer(__cgo__delimiters))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

// g_strdown is not generated due to deprecation attr

/*
Duplicates a string. If @str is %NULL it returns %NULL.
The returned string should be freed with g_free()
when no longer needed.
*/
func Strdup(str string) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strdup(__cgo__str)
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

// g_strdup_printf is not generated due to varargs

// g_strdup_vprintf is not generated due to varargs

/*
Copies %NULL-terminated array of strings. The copy is a deep copy;
the new array should be freed by first freeing each string, then
the array itself. g_strfreev() does this for you. If called
on a %NULL value, g_strdupv() simply returns %NULL.
*/
func Strdupv(str_array **C.gchar) (return__ []string) {
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.g_strdupv(str_array)
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Returns a string corresponding to the given error code, e.g.
"no such process". You should use this function in preference to
strerror(), because it returns a string in UTF-8 encoding, and since
not all platforms support the strerror() function.
*/
func Strerror(errnum int) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strerror(C.gint(errnum))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Escapes the special characters '\b', '\f', '\n', '\r', '\t', '\v', '\'
and '&quot;' in the string @source by inserting a '\' before
them. Additionally all characters in the range 0x01-0x1F (everything
below SPACE) and in the range 0x7F-0xFF (all non-ASCII chars) are
replaced with a '\' followed by their octal representation.
Characters supplied in @exceptions are not escaped.

g_strcompress() does the reverse conversion.
*/
func Strescape(source string, exceptions string) (return__ string) {
	__cgo__source := (*C.gchar)(unsafe.Pointer(C.CString(source)))
	__cgo__exceptions := (*C.gchar)(unsafe.Pointer(C.CString(exceptions)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strescape(__cgo__source, __cgo__exceptions)
	C.free(unsafe.Pointer(__cgo__source))
	C.free(unsafe.Pointer(__cgo__exceptions))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Frees a %NULL-terminated array of strings, and the array itself.
If called on a %NULL value, g_strfreev() simply returns.
*/
func Strfreev(str_array **C.gchar) {
	C.g_strfreev(str_array)
	return
}

/*
Creates a new #GString, initialized with the given string.
*/
func StringNew(init string) (return__ *String) {
	__cgo__init := (*C.gchar)(unsafe.Pointer(C.CString(init)))
	var __cgo__return__ *C.GString
	__cgo__return__ = C.g_string_new(__cgo__init)
	C.free(unsafe.Pointer(__cgo__init))
	return__ = (*String)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Creates a new #GString with @len bytes of the @init buffer.
Because a length is provided, @init need not be nul-terminated,
and can contain embedded nul bytes.

Since this function does not stop at nul bytes, it is the caller's
responsibility to ensure that @init has at least @len addressable
bytes.
*/
func StringNewLen(init string, len_ int64) (return__ *String) {
	__cgo__init := (*C.gchar)(unsafe.Pointer(C.CString(init)))
	var __cgo__return__ *C.GString
	__cgo__return__ = C.g_string_new_len(__cgo__init, C.gssize(len_))
	C.free(unsafe.Pointer(__cgo__init))
	return__ = (*String)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Creates a new #GString, with enough space for @dfl_size
bytes. This is useful if you are going to add a lot of
text to the string and don't want it to be reallocated
too often.
*/
func StringSizedNew(dfl_size int64) (return__ *String) {
	var __cgo__return__ *C.GString
	__cgo__return__ = C.g_string_sized_new(C.gsize(dfl_size))
	return__ = (*String)(unsafe.Pointer(__cgo__return__))
	return
}

/*
An auxiliary function for gettext() support (see Q_()).
*/
func StripContext(msgid string, msgval string) (return__ string) {
	__cgo__msgid := (*C.gchar)(unsafe.Pointer(C.CString(msgid)))
	__cgo__msgval := (*C.gchar)(unsafe.Pointer(C.CString(msgval)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strip_context(__cgo__msgid, __cgo__msgval)
	C.free(unsafe.Pointer(__cgo__msgid))
	C.free(unsafe.Pointer(__cgo__msgval))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

// g_strjoin is not generated due to varargs

/*
Joins a number of strings together to form one long string, with the
optional @separator inserted between each of them. The returned string
should be freed with g_free().
*/
func Strjoinv(separator string, str_array **C.gchar) (return__ string) {
	__cgo__separator := (*C.gchar)(unsafe.Pointer(C.CString(separator)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strjoinv(__cgo__separator, str_array)
	C.free(unsafe.Pointer(__cgo__separator))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Portability wrapper that calls strlcat() on systems which have it,
and emulates it otherwise. Appends nul-terminated @src string to @dest,
guaranteeing nul-termination for @dest. The total size of @dest won't
exceed @dest_size.

At most @dest_size - 1 characters will be copied. Unlike strncat(),
@dest_size is the full size of dest, not the space left over. This
function does not allocate memory. It always nul-terminates (unless
@dest_size == 0 or there were no nul characters in the @dest_size
characters of dest to start with).

Caveat: this is supposedly a more secure alternative to strcat() or
strncat(), but for real security g_strconcat() is harder to mess up.
*/
func Strlcat(dest string, src string, dest_size int64) (return__ int64) {
	__cgo__dest := (*C.gchar)(unsafe.Pointer(C.CString(dest)))
	__cgo__src := (*C.gchar)(unsafe.Pointer(C.CString(src)))
	var __cgo__return__ C.gsize
	__cgo__return__ = C.g_strlcat(__cgo__dest, __cgo__src, C.gsize(dest_size))
	C.free(unsafe.Pointer(__cgo__dest))
	C.free(unsafe.Pointer(__cgo__src))
	return__ = int64(__cgo__return__)
	return
}

/*
Portability wrapper that calls strlcpy() on systems which have it,
and emulates strlcpy() otherwise. Copies @src to @dest; @dest is
guaranteed to be nul-terminated; @src must be nul-terminated;
@dest_size is the buffer size, not the number of bytes to copy.

At most @dest_size - 1 characters will be copied. Always nul-terminates
(unless @dest_size is 0). This function does not allocate memory. Unlike
strncpy(), this function doesn't pad @dest (so it's often faster). It
returns the size of the attempted result, strlen (src), so if
@retval >= @dest_size, truncation occurred.

Caveat: strlcpy() is supposedly more secure than strcpy() or strncpy(),
but if you really want to avoid screwups, g_strdup() is an even better
idea.
*/
func Strlcpy(dest string, src string, dest_size int64) (return__ int64) {
	__cgo__dest := (*C.gchar)(unsafe.Pointer(C.CString(dest)))
	__cgo__src := (*C.gchar)(unsafe.Pointer(C.CString(src)))
	var __cgo__return__ C.gsize
	__cgo__return__ = C.g_strlcpy(__cgo__dest, __cgo__src, C.gsize(dest_size))
	C.free(unsafe.Pointer(__cgo__dest))
	C.free(unsafe.Pointer(__cgo__src))
	return__ = int64(__cgo__return__)
	return
}

// g_strncasecmp is not generated due to deprecation attr

/*
Duplicates the first @n bytes of a string, returning a newly-allocated
buffer @n + 1 bytes long which will always be nul-terminated. If @str
is less than @n bytes long the buffer is padded with nuls. If @str is
%NULL it returns %NULL. The returned value should be freed when no longer
needed.

To copy a number of characters from a UTF-8 encoded string,
use g_utf8_strncpy() instead.
*/
func Strndup(str string, n int64) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strndup(__cgo__str, C.gsize(n))
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Creates a new string @length bytes long filled with @fill_char.
The returned string should be freed when no longer needed.
*/
func Strnfill(length int64, fill_char byte) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strnfill(C.gsize(length), C.gchar(fill_char))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Reverses all of the bytes in a string. For example,
`g_strreverse ("abcdef")` will result in "fedcba".

Note that g_strreverse() doesn't work on UTF-8 strings
containing multibyte characters. For that purpose, use
g_utf8_strreverse().
*/
func Strreverse(string_ string) (return__ string) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strreverse(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Searches the string @haystack for the last occurrence
of the string @needle.
*/
func Strrstr(haystack string, needle string) (return__ string) {
	__cgo__haystack := (*C.gchar)(unsafe.Pointer(C.CString(haystack)))
	__cgo__needle := (*C.gchar)(unsafe.Pointer(C.CString(needle)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strrstr(__cgo__haystack, __cgo__needle)
	C.free(unsafe.Pointer(__cgo__haystack))
	C.free(unsafe.Pointer(__cgo__needle))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Searches the string @haystack for the last occurrence
of the string @needle, limiting the length of the search
to @haystack_len.
*/
func StrrstrLen(haystack string, haystack_len int64, needle string) (return__ string) {
	__cgo__haystack := (*C.gchar)(unsafe.Pointer(C.CString(haystack)))
	__cgo__needle := (*C.gchar)(unsafe.Pointer(C.CString(needle)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strrstr_len(__cgo__haystack, C.gssize(haystack_len), __cgo__needle)
	C.free(unsafe.Pointer(__cgo__haystack))
	C.free(unsafe.Pointer(__cgo__needle))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns a string describing the given signal, e.g. "Segmentation fault".
You should use this function in preference to strsignal(), because it
returns a string in UTF-8 encoding, and since not all platforms support
the strsignal() function.
*/
func Strsignal(signum int) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strsignal(C.gint(signum))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Splits a string into a maximum of @max_tokens pieces, using the given
@delimiter. If @max_tokens is reached, the remainder of @string is
appended to the last token.

As a special case, the result of splitting the empty string "" is an empty
vector, not a vector containing a single string. The reason for this
special case is that being able to represent a empty vector is typically
more useful than consistent handling of empty elements. If you do need
to represent empty elements, you'll need to check for the empty string
before calling g_strsplit().
*/
func Strsplit(string_ string, delimiter string, max_tokens int) (return__ []string) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	__cgo__delimiter := (*C.gchar)(unsafe.Pointer(C.CString(delimiter)))
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.g_strsplit(__cgo__string_, __cgo__delimiter, C.gint(max_tokens))
	C.free(unsafe.Pointer(__cgo__string_))
	C.free(unsafe.Pointer(__cgo__delimiter))
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Splits @string into a number of tokens not containing any of the characters
in @delimiter. A token is the (possibly empty) longest string that does not
contain any of the characters in @delimiters. If @max_tokens is reached, the
remainder is appended to the last token.

For example the result of g_strsplit_set ("abc:def/ghi", ":/", -1) is a
%NULL-terminated vector containing the three strings "abc", "def",
and "ghi".

The result if g_strsplit_set (":def/ghi:", ":/", -1) is a %NULL-terminated
vector containing the four strings "", "def", "ghi", and "".

As a special case, the result of splitting the empty string "" is an empty
vector, not a vector containing a single string. The reason for this
special case is that being able to represent a empty vector is typically
more useful than consistent handling of empty elements. If you do need
to represent empty elements, you'll need to check for the empty string
before calling g_strsplit_set().

Note that this function works on bytes not characters, so it can't be used
to delimit UTF-8 strings for anything but ASCII characters.
*/
func StrsplitSet(string_ string, delimiters string, max_tokens int) (return__ []string) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	__cgo__delimiters := (*C.gchar)(unsafe.Pointer(C.CString(delimiters)))
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.g_strsplit_set(__cgo__string_, __cgo__delimiters, C.gint(max_tokens))
	C.free(unsafe.Pointer(__cgo__string_))
	C.free(unsafe.Pointer(__cgo__delimiters))
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Searches the string @haystack for the first occurrence
of the string @needle, limiting the length of the search
to @haystack_len.
*/
func StrstrLen(haystack string, haystack_len int64, needle string) (return__ string) {
	__cgo__haystack := (*C.gchar)(unsafe.Pointer(C.CString(haystack)))
	__cgo__needle := (*C.gchar)(unsafe.Pointer(C.CString(needle)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_strstr_len(__cgo__haystack, C.gssize(haystack_len), __cgo__needle)
	C.free(unsafe.Pointer(__cgo__haystack))
	C.free(unsafe.Pointer(__cgo__needle))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts a string to a #gdouble value.
It calls the standard strtod() function to handle the conversion, but
if the string is not completely converted it attempts the conversion
again with g_ascii_strtod(), and returns the best match.

This function should seldom be used. The normal situation when reading
numbers not for human consumption is to use g_ascii_strtod(). Only when
you know that you must expect both locale formatted and C formatted numbers
should you use this. Make sure that you don't pass strings such as comma
separated lists of values, since the commas may be interpreted as a decimal
point in some locales, causing unexpected results.
*/
func Strtod(nptr string, endptr **C.gchar) (return__ float64) {
	__cgo__nptr := (*C.gchar)(unsafe.Pointer(C.CString(nptr)))
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.g_strtod(__cgo__nptr, endptr)
	C.free(unsafe.Pointer(__cgo__nptr))
	return__ = float64(__cgo__return__)
	return
}

// g_strup is not generated due to deprecation attr

func StrvGetType() (return__ C.GType) {
	return__ = C.g_strv_get_type()
	return
}

/*
Returns the length of the given %NULL-terminated
string array @str_array.
*/
func StrvLength(str_array **C.gchar) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_strv_length(str_array)
	return__ = uint(__cgo__return__)
	return
}

/*
Create a new test case, similar to g_test_create_case(). However
the test is assumed to use no fixture, and test suites are automatically
created on the fly and added to the root fixture, based on the
slash-separated portions of @testpath. The @test_data argument
will be passed as first argument to @test_func.

If @testpath includes the component "subprocess" anywhere in it,
the test will be skipped by default, and only run if explicitly
required via the `-p` command-line option or g_test_trap_subprocess().
*/
func TestAddDataFunc(testpath string, test_data unsafe.Pointer, test_func C.GTestDataFunc) {
	__cgo__testpath := C.CString(testpath)
	C.g_test_add_data_func(__cgo__testpath, (C.gconstpointer)(test_data), test_func)
	C.free(unsafe.Pointer(__cgo__testpath))
	return
}

/*
Create a new test case, as with g_test_add_data_func(), but freeing
@test_data after the test run is complete.
*/
func TestAddDataFuncFull(testpath string, test_data unsafe.Pointer, test_func C.GTestDataFunc, data_free_func C.GDestroyNotify) {
	__cgo__testpath := C.CString(testpath)
	C.g_test_add_data_func_full(__cgo__testpath, (C.gpointer)(test_data), test_func, data_free_func)
	C.free(unsafe.Pointer(__cgo__testpath))
	return
}

/*
Create a new test case, similar to g_test_create_case(). However
the test is assumed to use no fixture, and test suites are automatically
created on the fly and added to the root fixture, based on the
slash-separated portions of @testpath.

If @testpath includes the component "subprocess" anywhere in it,
the test will be skipped by default, and only run if explicitly
required via the `-p` command-line option or g_test_trap_subprocess().
*/
func TestAddFunc(testpath string, test_func C.GTestFunc) {
	__cgo__testpath := C.CString(testpath)
	C.g_test_add_func(__cgo__testpath, test_func)
	C.free(unsafe.Pointer(__cgo__testpath))
	return
}

func TestAddVtable(testpath string, data_size int64, test_data unsafe.Pointer, data_setup C.GTestFixtureFunc, data_test C.GTestFixtureFunc, data_teardown C.GTestFixtureFunc) {
	__cgo__testpath := C.CString(testpath)
	C.g_test_add_vtable(__cgo__testpath, C.gsize(data_size), (C.gconstpointer)(test_data), data_setup, data_test, data_teardown)
	C.free(unsafe.Pointer(__cgo__testpath))
	return
}

func TestAssertExpectedMessagesInternal(domain string, file string, line int, func_ string) {
	__cgo__domain := C.CString(domain)
	__cgo__file := C.CString(file)
	__cgo__func_ := C.CString(func_)
	C.g_test_assert_expected_messages_internal(__cgo__domain, __cgo__file, C.int(line), __cgo__func_)
	C.free(unsafe.Pointer(__cgo__domain))
	C.free(unsafe.Pointer(__cgo__file))
	C.free(unsafe.Pointer(__cgo__func_))
	return
}

/*
This function adds a message to test reports that
associates a bug URI with a test case.
Bug URIs are constructed from a base URI set with g_test_bug_base()
and @bug_uri_snippet.
*/
func TestBug(bug_uri_snippet string) {
	__cgo__bug_uri_snippet := C.CString(bug_uri_snippet)
	C.g_test_bug(__cgo__bug_uri_snippet)
	C.free(unsafe.Pointer(__cgo__bug_uri_snippet))
	return
}

/*
Specify the base URI for bug reports.

The base URI is used to construct bug report messages for
g_test_message() when g_test_bug() is called.
Calling this function outside of a test case sets the
default base URI for all test cases. Calling it from within
a test case changes the base URI for the scope of the test
case only.
Bug URIs are constructed by appending a bug specific URI
portion to @uri_pattern, or by replacing the special string
'\%s' within @uri_pattern if that is present.
*/
func TestBugBase(uri_pattern string) {
	__cgo__uri_pattern := C.CString(uri_pattern)
	C.g_test_bug_base(__cgo__uri_pattern)
	C.free(unsafe.Pointer(__cgo__uri_pattern))
	return
}

// g_test_build_filename is not generated due to varargs

/*
Create a new #GTestCase, named @test_name, this API is fairly
low level, calling g_test_add() or g_test_add_func() is preferable.
When this test is executed, a fixture structure of size @data_size
will be allocated and filled with 0s. Then @data_setup is called
to initialize the fixture. After fixture setup, the actual test
function @data_test is called. Once the test run completed, the
fixture structure is torn down  by calling @data_teardown and
after that the memory is released.

Splitting up a test run into fixture setup, test function and
fixture teardown is most usful if the same fixture is used for
multiple tests. In this cases, g_test_create_case() will be
called with the same fixture, but varying @test_name and
@data_test arguments.
*/
func TestCreateCase(test_name string, data_size int64, test_data unsafe.Pointer, data_setup C.GTestFixtureFunc, data_test C.GTestFixtureFunc, data_teardown C.GTestFixtureFunc) (return__ *TestCase) {
	__cgo__test_name := C.CString(test_name)
	var __cgo__return__ *C.GTestCase
	__cgo__return__ = C.g_test_create_case(__cgo__test_name, C.gsize(data_size), (C.gconstpointer)(test_data), data_setup, data_test, data_teardown)
	C.free(unsafe.Pointer(__cgo__test_name))
	return__ = (*TestCase)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Create a new test suite with the name @suite_name.
*/
func TestCreateSuite(suite_name string) (return__ *TestSuite) {
	__cgo__suite_name := C.CString(suite_name)
	var __cgo__return__ *C.GTestSuite
	__cgo__return__ = C.g_test_create_suite(__cgo__suite_name)
	C.free(unsafe.Pointer(__cgo__suite_name))
	return__ = (*TestSuite)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Indicates that a message with the given @log_domain and @log_level,
with text matching @pattern, is expected to be logged. When this
message is logged, it will not be printed, and the test case will
not abort.

Use g_test_assert_expected_messages() to assert that all
previously-expected messages have been seen and suppressed.

You can call this multiple times in a row, if multiple messages are
expected as a result of a single call. (The messages must appear in
the same order as the calls to g_test_expect_message().)

For example:

|[<!-- language="C" -->
  // g_main_context_push_thread_default() should fail if the
  // context is already owned by another thread.
  g_test_expect_message (G_LOG_DOMAIN,
                         G_LOG_LEVEL_CRITICAL,
                         "assertion*acquired_context*failed");
  g_main_context_push_thread_default (bad_context);
  g_test_assert_expected_messages ();
]|

Note that you cannot use this to test g_error() messages, since
g_error() intentionally never returns even if the program doesn't
abort; use g_test_trap_subprocess() in this case.

If messages at %G_LOG_LEVEL_DEBUG are emitted, but not explicitly
expected via g_test_expect_message() then they will be ignored.
*/
func TestExpectMessage(log_domain string, log_level C.GLogLevelFlags, pattern string) {
	__cgo__log_domain := (*C.gchar)(unsafe.Pointer(C.CString(log_domain)))
	__cgo__pattern := (*C.gchar)(unsafe.Pointer(C.CString(pattern)))
	C.g_test_expect_message(__cgo__log_domain, log_level, __cgo__pattern)
	C.free(unsafe.Pointer(__cgo__log_domain))
	C.free(unsafe.Pointer(__cgo__pattern))
	return
}

/*
Indicates that a test failed. This function can be called
multiple times from the same test. You can use this function
if your test failed in a recoverable way.

Do not use this function if the failure of a test could cause
other tests to malfunction.

Calling this function will not stop the test from running, you
need to return from the test function yourself. So you can
produce additional diagnostic messages or even continue running
the test.

If not called from inside a test, this function does nothing.
*/
func TestFail() {
	C.g_test_fail()
	return
}

/*
Returns whether a test has already failed. This will
be the case when g_test_fail(), g_test_incomplete()
or g_test_skip() have been called, but also if an
assertion has failed.

This can be useful to return early from a test if
continuing after a failed assertion might be harmful.

The return value of this function is only meaningful
if it is called from inside a test function.
*/
func TestFailed() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_test_failed()
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the pathname of the directory containing test files of the type
specified by @file_type.

This is approximately the same as calling g_test_build_filename("."),
but you don't need to free the return value.
*/
func TestGetDir(file_type C.GTestFileType) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_test_get_dir(file_type)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

// g_test_get_filename is not generated due to varargs

/*
Get the toplevel test suite for the test path API.
*/
func TestGetRoot() (return__ *TestSuite) {
	var __cgo__return__ *C.GTestSuite
	__cgo__return__ = C.g_test_get_root()
	return__ = (*TestSuite)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Indicates that a test failed because of some incomplete
functionality. This function can be called multiple times
from the same test.

Calling this function will not stop the test from running, you
need to return from the test function yourself. So you can
produce additional diagnostic messages or even continue running
the test.

If not called from inside a test, this function does nothing.
*/
func TestIncomplete(msg string) {
	__cgo__msg := (*C.gchar)(unsafe.Pointer(C.CString(msg)))
	C.g_test_incomplete(__cgo__msg)
	C.free(unsafe.Pointer(__cgo__msg))
	return
}

// g_test_init is not generated due to varargs

/*
Installs a non-error fatal log handler which can be
used to decide whether log messages which are counted
as fatal abort the program.

The use case here is that you are running a test case
that depends on particular libraries or circumstances
and cannot prevent certain known critical or warning
messages. So you install a handler that compares the
domain and message to precisely not abort in such a case.

Note that the handler is reset at the beginning of
any test case, so you have to set it inside each test
function which needs the special behavior.

This handler has no effect on g_error messages.
*/
func TestLogSetFatalHandler(log_func C.GTestLogFatalFunc, user_data unsafe.Pointer) {
	C.g_test_log_set_fatal_handler(log_func, (C.gpointer)(user_data))
	return
}

func TestLogTypeName(log_type C.GTestLogType) (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_test_log_type_name(log_type)
	return__ = C.GoString(__cgo__return__)
	return
}

// g_test_maximized_result is not generated due to varargs

// g_test_message is not generated due to varargs

// g_test_minimized_result is not generated due to varargs

/*
This function enqueus a callback @destroy_func to be executed
during the next test case teardown phase. This is most useful
to auto destruct allocted test resources at the end of a test run.
Resources are released in reverse queue order, that means enqueueing
callback A before callback B will cause B() to be called before
A() during teardown.
*/
func TestQueueDestroy(destroy_func C.GDestroyNotify, destroy_data unsafe.Pointer) {
	C.g_test_queue_destroy(destroy_func, (C.gpointer)(destroy_data))
	return
}

/*
Enqueue a pointer to be released with g_free() during the next
teardown phase. This is equivalent to calling g_test_queue_destroy()
with a destroy callback of g_free().
*/
func TestQueueFree(gfree_pointer unsafe.Pointer) {
	C.g_test_queue_free((C.gpointer)(gfree_pointer))
	return
}

/*
Get a reproducible random floating point number,
see g_test_rand_int() for details on test case random numbers.
*/
func TestRandDouble() (return__ float64) {
	var __cgo__return__ C.double
	__cgo__return__ = C.g_test_rand_double()
	return__ = float64(__cgo__return__)
	return
}

/*
Get a reproducible random floating pointer number out of a specified range,
see g_test_rand_int() for details on test case random numbers.
*/
func TestRandDoubleRange(range_start float64, range_end float64) (return__ float64) {
	var __cgo__return__ C.double
	__cgo__return__ = C.g_test_rand_double_range(C.double(range_start), C.double(range_end))
	return__ = float64(__cgo__return__)
	return
}

/*
Get a reproducible random integer number.

The random numbers generated by the g_test_rand_*() family of functions
change with every new test program start, unless the --seed option is
given when starting test programs.

For individual test cases however, the random number generator is
reseeded, to avoid dependencies between tests and to make --seed
effective for all test cases.
*/
func TestRandInt() (return__ int32) {
	var __cgo__return__ C.gint32
	__cgo__return__ = C.g_test_rand_int()
	return__ = int32(__cgo__return__)
	return
}

/*
Get a reproducible random integer number out of a specified range,
see g_test_rand_int() for details on test case random numbers.
*/
func TestRandIntRange(begin int32, end int32) (return__ int32) {
	var __cgo__return__ C.gint32
	__cgo__return__ = C.g_test_rand_int_range(C.gint32(begin), C.gint32(end))
	return__ = int32(__cgo__return__)
	return
}

/*
Runs all tests under the toplevel suite which can be retrieved
with g_test_get_root(). Similar to g_test_run_suite(), the test
cases to be run are filtered according to test path arguments
(`-p testpath`) as parsed by g_test_init(). g_test_run_suite()
or g_test_run() may only be called once in a program.

In general, the tests and sub-suites within each suite are run in
the order in which they are defined. However, note that prior to
GLib 2.36, there was a bug in the `g_test_add_*`
functions which caused them to create multiple suites with the same
name, meaning that if you created tests "/foo/simple",
"/bar/simple", and "/foo/using-bar" in that order, they would get
run in that order (since g_test_run() would run the first "/foo"
suite, then the "/bar" suite, then the second "/foo" suite). As of
2.36, this bug is fixed, and adding the tests in that order would
result in a running order of "/foo/simple", "/foo/using-bar",
"/bar/simple". If this new ordering is sub-optimal (because it puts
more-complicated tests before simpler ones, making it harder to
figure out exactly what has failed), you can fix it by changing the
test paths to group tests by suite in a way that will result in the
desired running order. Eg, "/simple/foo", "/simple/bar",
"/complex/foo-using-bar".

However, you should never make the actual result of a test depend
on the order that tests are run in. If you need to ensure that some
particular code runs before or after a given test case, use
g_test_add(), which lets you specify setup and teardown functions.

If all tests are skipped, this function will return 0 if
producing TAP output, or 77 (treated as "skip test" by Automake) otherwise.
*/
func TestRun() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.g_test_run()
	return__ = int(__cgo__return__)
	return
}

/*
Execute the tests within @suite and all nested #GTestSuites.
The test suites to be executed are filtered according to
test path arguments (`-p testpath`) as parsed by g_test_init().
See the g_test_run() documentation for more information on the
order that tests are run in.

g_test_run_suite() or g_test_run() may only be called once
in a program.
*/
func TestRunSuite(suite *TestSuite) (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.g_test_run_suite((*C.GTestSuite)(unsafe.Pointer(suite)))
	return__ = int(__cgo__return__)
	return
}

/*
Changes the behaviour of g_assert_cmpstr(), g_assert_cmpint(),
g_assert_cmpuint(), g_assert_cmphex(), g_assert_cmpfloat(),
g_assert_true(), g_assert_false(), g_assert_null(), g_assert_no_error(),
g_assert_error(), g_test_assert_expected_messages() and the various
g_test_trap_assert_*() macros to not abort to program, but instead
call g_test_fail() and continue. (This also changes the behavior of
g_test_fail() so that it will not cause the test program to abort
after completing the failed test.)

Note that the g_assert_not_reached() and g_assert() are not
affected by this.

This function can only be called after g_test_init().
*/
func TestSetNonfatalAssertions() {
	C.g_test_set_nonfatal_assertions()
	return
}

/*
Indicates that a test was skipped.

Calling this function will not stop the test from running, you
need to return from the test function yourself. So you can
produce additional diagnostic messages or even continue running
the test.

If not called from inside a test, this function does nothing.
*/
func TestSkip(msg string) {
	__cgo__msg := (*C.gchar)(unsafe.Pointer(C.CString(msg)))
	C.g_test_skip(__cgo__msg)
	C.free(unsafe.Pointer(__cgo__msg))
	return
}

/*
Returns %TRUE (after g_test_init() has been called) if the test
program is running under g_test_trap_subprocess().
*/
func TestSubprocess() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_test_subprocess()
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the time since the last start of the timer with g_test_timer_start().
*/
func TestTimerElapsed() (return__ float64) {
	var __cgo__return__ C.double
	__cgo__return__ = C.g_test_timer_elapsed()
	return__ = float64(__cgo__return__)
	return
}

/*
Report the last result of g_test_timer_elapsed().
*/
func TestTimerLast() (return__ float64) {
	var __cgo__return__ C.double
	__cgo__return__ = C.g_test_timer_last()
	return__ = float64(__cgo__return__)
	return
}

/*
Start a timing test. Call g_test_timer_elapsed() when the task is supposed
to be done. Call this function again to restart the timer.
*/
func TestTimerStart() {
	C.g_test_timer_start()
	return
}

func TestTrapAssertions(domain string, file string, line int, func_ string, assertion_flags uint64, pattern string) {
	__cgo__domain := C.CString(domain)
	__cgo__file := C.CString(file)
	__cgo__func_ := C.CString(func_)
	__cgo__pattern := C.CString(pattern)
	C.g_test_trap_assertions(__cgo__domain, __cgo__file, C.int(line), __cgo__func_, C.guint64(assertion_flags), __cgo__pattern)
	C.free(unsafe.Pointer(__cgo__domain))
	C.free(unsafe.Pointer(__cgo__file))
	C.free(unsafe.Pointer(__cgo__func_))
	C.free(unsafe.Pointer(__cgo__pattern))
	return
}

// g_test_trap_fork is not generated due to deprecation attr

/*
Check the result of the last g_test_trap_subprocess() call.
*/
func TestTrapHasPassed() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_test_trap_has_passed()
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Check the result of the last g_test_trap_subprocess() call.
*/
func TestTrapReachedTimeout() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_test_trap_reached_timeout()
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Respawns the test program to run only @test_path in a subprocess.
This can be used for a test case that might not return, or that
might abort.

If @test_path is %NULL then the same test is re-run in a subprocess.
You can use g_test_subprocess() to determine whether the test is in
a subprocess or not.

@test_path can also be the name of the parent test, followed by
"`/subprocess/`" and then a name for the specific subtest (or just
ending with "`/subprocess`" if the test only has one child test);
tests with names of this form will automatically be skipped in the
parent process.

If @usec_timeout is non-0, the test subprocess is aborted and
considered failing if its run time exceeds it.

The subprocess behavior can be configured with the
#GTestSubprocessFlags flags.

You can use methods such as g_test_trap_assert_passed(),
g_test_trap_assert_failed(), and g_test_trap_assert_stderr() to
check the results of the subprocess. (But note that
g_test_trap_assert_stdout() and g_test_trap_assert_stderr()
cannot be used if @test_flags specifies that the child should
inherit the parent stdout/stderr.)

If your `main ()` needs to behave differently in
the subprocess, you can call g_test_subprocess() (after calling
g_test_init()) to see whether you are in a subprocess.

The following example tests that calling
`my_object_new(1000000)` will abort with an error
message.

|[<!-- language="C" -->
  static void
  test_create_large_object_subprocess (void)
  {
    if (g_test_subprocess ())
      {
        my_object_new (1000000);
        return;
      }

    // Reruns this same test in a subprocess
    g_test_trap_subprocess (NULL, 0, 0);
    g_test_trap_assert_failed ();
    g_test_trap_assert_stderr ("*ERROR*too large*");
  }

  int
  main (int argc, char **argv)
  {
    g_test_init (&argc, &argv, NULL);

    g_test_add_func ("/myobject/create_large_object",
                     test_create_large_object);
    return g_test_run ();
  }
]|
*/
func TestTrapSubprocess(test_path string, usec_timeout uint64, test_flags C.GTestSubprocessFlags) {
	__cgo__test_path := C.CString(test_path)
	C.g_test_trap_subprocess(__cgo__test_path, C.guint64(usec_timeout), test_flags)
	C.free(unsafe.Pointer(__cgo__test_path))
	return
}

func ThreadErrorQuark() (return__ C.GQuark) {
	return__ = C.g_thread_error_quark()
	return
}

/*
Terminates the current thread.

If another thread is waiting for us using g_thread_join() then the
waiting thread will be woken up and get @retval as the return value
of g_thread_join().

Calling g_thread_exit() with a parameter @retval is equivalent to
returning @retval from the function @func, as given to g_thread_new().

You must only call g_thread_exit() from a thread that you created
yourself with g_thread_new() or related APIs. You must not call
this function from a thread created with another threading library
or or from within a #GThreadPool.
*/
func ThreadExit(retval unsafe.Pointer) {
	C.g_thread_exit((C.gpointer)(retval))
	return
}

/*
This function will return the maximum @interval that a
thread will wait in the thread pool for new tasks before
being stopped.

If this function returns 0, threads waiting in the thread
pool for new work are not stopped.
*/
func ThreadPoolGetMaxIdleTime() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_thread_pool_get_max_idle_time()
	return__ = uint(__cgo__return__)
	return
}

/*
Returns the maximal allowed number of unused threads.
*/
func ThreadPoolGetMaxUnusedThreads() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_thread_pool_get_max_unused_threads()
	return__ = int(__cgo__return__)
	return
}

/*
Returns the number of currently unused threads.
*/
func ThreadPoolGetNumUnusedThreads() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_thread_pool_get_num_unused_threads()
	return__ = uint(__cgo__return__)
	return
}

/*
This function will set the maximum @interval that a thread
waiting in the pool for new tasks can be idle for before
being stopped. This function is similar to calling
g_thread_pool_stop_unused_threads() on a regular timeout,
except this is done on a per thread basis.

By setting @interval to 0, idle threads will not be stopped.

The default value is 15000 (15 seconds).
*/
func ThreadPoolSetMaxIdleTime(interval uint) {
	C.g_thread_pool_set_max_idle_time(C.guint(interval))
	return
}

/*
Sets the maximal number of unused threads to @max_threads.
If @max_threads is -1, no limit is imposed on the number
of unused threads.

The default value is 2.
*/
func ThreadPoolSetMaxUnusedThreads(max_threads int) {
	C.g_thread_pool_set_max_unused_threads(C.gint(max_threads))
	return
}

/*
Stops all currently unused threads. This does not change the
maximal number of unused threads. This function can be used to
regularly stop all unused threads e.g. from g_timeout_add().
*/
func ThreadPoolStopUnusedThreads() {
	C.g_thread_pool_stop_unused_threads()
	return
}

/*
This functions returns the #GThread corresponding to the
current thread. Note that this function does not increase
the reference count of the returned struct.

This function will return a #GThread even for threads that
were not created by GLib (i.e. those created by other threading
APIs). This may be useful for thread identification purposes
(i.e. comparisons) but you must not use GLib functions (such
as g_thread_join()) on these threads.
*/
func ThreadSelf() (return__ *Thread) {
	var __cgo__return__ *C.GThread
	__cgo__return__ = C.g_thread_self()
	return__ = (*Thread)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Causes the calling thread to voluntarily relinquish the CPU, so
that other threads can run.

This function is often used as a method to make busy wait less evil.
*/
func ThreadYield() {
	C.g_thread_yield()
	return
}

/*
Converts a string containing an ISO 8601 encoded date and time
to a #GTimeVal and puts it into @time_.

@iso_date must include year, month, day, hours, minutes, and
seconds. It can optionally include fractions of a second and a time
zone indicator. (In the absence of any time zone indication, the
timestamp is assumed to be in local time.)
*/
func TimeValFromIso8601(iso_date string) (time_ C.GTimeVal, return__ bool) {
	__cgo__iso_date := (*C.gchar)(unsafe.Pointer(C.CString(iso_date)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_time_val_from_iso8601(__cgo__iso_date, &time_)
	C.free(unsafe.Pointer(__cgo__iso_date))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets a function to be called at regular intervals, with the default
priority, #G_PRIORITY_DEFAULT.  The function is called repeatedly
until it returns %FALSE, at which point the timeout is automatically
destroyed and the function will not be called again.  The first call
to the function will be at the end of the first @interval.

Note that timeout functions may be delayed, due to the processing of other
event sources. Thus they should not be relied on for precise timing.
After each call to the timeout function, the time of the next
timeout is recalculated based on the current time and the given interval
(it does not try to 'catch up' time lost in delays).

If you want to have a timer in the "seconds" range and do not care
about the exact time of the first call of the timer, use the
g_timeout_add_seconds() function; this function allows for more
optimizations and more efficient system power usage.

This internally creates a main loop source using g_timeout_source_new()
and attaches it to the main loop context using g_source_attach(). You can
do these steps manually if you need greater control.

The interval given is in terms of monotonic time, not wall clock
time.  See g_get_monotonic_time().
*/
func TimeoutAdd(interval uint, function C.GSourceFunc, data unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_timeout_add(C.guint(interval), function, (C.gpointer)(data))
	return__ = uint(__cgo__return__)
	return
}

/*
Sets a function to be called at regular intervals, with the given
priority.  The function is called repeatedly until it returns
%FALSE, at which point the timeout is automatically destroyed and
the function will not be called again.  The @notify function is
called when the timeout is destroyed.  The first call to the
function will be at the end of the first @interval.

Note that timeout functions may be delayed, due to the processing of other
event sources. Thus they should not be relied on for precise timing.
After each call to the timeout function, the time of the next
timeout is recalculated based on the current time and the given interval
(it does not try to 'catch up' time lost in delays).

This internally creates a main loop source using g_timeout_source_new()
and attaches it to the main loop context using g_source_attach(). You can
do these steps manually if you need greater control.

The interval given in terms of monotonic time, not wall clock time.
See g_get_monotonic_time().
*/
func TimeoutAddFull(priority int, interval uint, function C.GSourceFunc, data unsafe.Pointer, notify C.GDestroyNotify) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_timeout_add_full(C.gint(priority), C.guint(interval), function, (C.gpointer)(data), notify)
	return__ = uint(__cgo__return__)
	return
}

/*
Sets a function to be called at regular intervals with the default
priority, #G_PRIORITY_DEFAULT. The function is called repeatedly until
it returns %FALSE, at which point the timeout is automatically destroyed
and the function will not be called again.

This internally creates a main loop source using
g_timeout_source_new_seconds() and attaches it to the main loop context
using g_source_attach(). You can do these steps manually if you need
greater control. Also see g_timeout_add_seconds_full().

Note that the first call of the timer may not be precise for timeouts
of one second. If you need finer precision and have such a timeout,
you may want to use g_timeout_add() instead.

The interval given is in terms of monotonic time, not wall clock
time.  See g_get_monotonic_time().
*/
func TimeoutAddSeconds(interval uint, function C.GSourceFunc, data unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_timeout_add_seconds(C.guint(interval), function, (C.gpointer)(data))
	return__ = uint(__cgo__return__)
	return
}

/*
Sets a function to be called at regular intervals, with @priority.
The function is called repeatedly until it returns %FALSE, at which
point the timeout is automatically destroyed and the function will
not be called again.

Unlike g_timeout_add(), this function operates at whole second granularity.
The initial starting point of the timer is determined by the implementation
and the implementation is expected to group multiple timers together so that
they fire all at the same time.
To allow this grouping, the @interval to the first timer is rounded
and can deviate up to one second from the specified interval.
Subsequent timer iterations will generally run at the specified interval.

Note that timeout functions may be delayed, due to the processing of other
event sources. Thus they should not be relied on for precise timing.
After each call to the timeout function, the time of the next
timeout is recalculated based on the current time and the given @interval

If you want timing more precise than whole seconds, use g_timeout_add()
instead.

The grouping of timers to fire at the same time results in a more power
and CPU efficient behavior so if your timer is in multiples of seconds
and you don't require the first timer exactly one second from now, the
use of g_timeout_add_seconds() is preferred over g_timeout_add().

This internally creates a main loop source using
g_timeout_source_new_seconds() and attaches it to the main loop context
using g_source_attach(). You can do these steps manually if you need
greater control.

The interval given is in terms of monotonic time, not wall clock
time.  See g_get_monotonic_time().
*/
func TimeoutAddSecondsFull(priority int, interval uint, function C.GSourceFunc, data unsafe.Pointer, notify C.GDestroyNotify) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_timeout_add_seconds_full(C.gint(priority), C.guint(interval), function, (C.gpointer)(data), notify)
	return__ = uint(__cgo__return__)
	return
}

/*
Creates a new timeout source.

The source will not initially be associated with any #GMainContext
and must be added to one with g_source_attach() before it will be
executed.

The interval given is in terms of monotonic time, not wall clock
time.  See g_get_monotonic_time().
*/
func TimeoutSourceNew(interval uint) (return__ *Source) {
	var __cgo__return__ *C.GSource
	__cgo__return__ = C.g_timeout_source_new(C.guint(interval))
	return__ = (*Source)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Creates a new timeout source.

The source will not initially be associated with any #GMainContext
and must be added to one with g_source_attach() before it will be
executed.

The scheduling granularity/accuracy of this timeout source will be
in seconds.

The interval given in terms of monotonic time, not wall clock time.
See g_get_monotonic_time().
*/
func TimeoutSourceNewSeconds(interval uint) (return__ *Source) {
	var __cgo__return__ *C.GSource
	__cgo__return__ = C.g_timeout_source_new_seconds(C.guint(interval))
	return__ = (*Source)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Returns the height of a #GTrashStack.

Note that execution of this function is of O(N) complexity
where N denotes the number of items on the stack.
*/
func TrashStackHeight(stack_p **C.GTrashStack) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_trash_stack_height(stack_p)
	return__ = uint(__cgo__return__)
	return
}

/*
Pushes a piece of memory onto a #GTrashStack.
*/
func TrashStackPush(stack_p **C.GTrashStack, data_p unsafe.Pointer) {
	C.g_trash_stack_push(stack_p, (C.gpointer)(data_p))
	return
}

/*
Attempts to allocate @n_bytes, and returns %NULL on failure.
Contrast with g_malloc(), which aborts the program on failure.
*/
func TryMalloc(n_bytes int64) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_try_malloc(C.gsize(n_bytes))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Attempts to allocate @n_bytes, initialized to 0's, and returns %NULL on
failure. Contrast with g_malloc0(), which aborts the program on failure.
*/
func TryMalloc0(n_bytes int64) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_try_malloc0(C.gsize(n_bytes))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
This function is similar to g_try_malloc0(), allocating (@n_blocks * @n_block_bytes) bytes,
but care is taken to detect possible overflow during multiplication.
*/
func TryMalloc0N(n_blocks int64, n_block_bytes int64) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_try_malloc0_n(C.gsize(n_blocks), C.gsize(n_block_bytes))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
This function is similar to g_try_malloc(), allocating (@n_blocks * @n_block_bytes) bytes,
but care is taken to detect possible overflow during multiplication.
*/
func TryMallocN(n_blocks int64, n_block_bytes int64) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_try_malloc_n(C.gsize(n_blocks), C.gsize(n_block_bytes))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Attempts to realloc @mem to a new size, @n_bytes, and returns %NULL
on failure. Contrast with g_realloc(), which aborts the program
on failure. If @mem is %NULL, behaves the same as g_try_malloc().
*/
func TryRealloc(mem unsafe.Pointer, n_bytes int64) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_try_realloc((C.gpointer)(mem), C.gsize(n_bytes))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
This function is similar to g_try_realloc(), allocating (@n_blocks * @n_block_bytes) bytes,
but care is taken to detect possible overflow during multiplication.
*/
func TryReallocN(mem unsafe.Pointer, n_blocks int64, n_block_bytes int64) (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.g_try_realloc_n((C.gpointer)(mem), C.gsize(n_blocks), C.gsize(n_block_bytes))
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Convert a string from UCS-4 to UTF-16. A 0 character will be
added to the result after the converted text.
*/
func Ucs4ToUtf16(str *C.gunichar, len_ int64, items_read *C.glong, items_written *C.glong) (return__ *C.gunichar2, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_ucs4_to_utf16(str, C.glong(len_), items_read, items_written, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Convert a string from a 32-bit fixed width representation as UCS-4.
to UTF-8. The result will be terminated with a 0 byte.
*/
func Ucs4ToUtf8(str *C.gunichar, len_ int64, items_read *C.glong, items_written *C.glong) (return__ string, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_ucs4_to_utf8(str, C.glong(len_), items_read, items_written, &__cgo_error__)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Determines the break type of @c. @c should be a Unicode character
(to derive a character from UTF-8 encoded text, use
g_utf8_get_char()). The break type is used to find word and line
breaks ("text boundaries"), Pango implements the Unicode boundary
resolution algorithms and normally you would use a function such
as pango_break() instead of caring about break types yourself.
*/
func UnicharBreakType(c rune) (return__ C.GUnicodeBreakType) {
	return__ = C.g_unichar_break_type(C.gunichar(c))
	return
}

/*
Determines the canonical combining class of a Unicode character.
*/
func UnicharCombiningClass(uc rune) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_unichar_combining_class(C.gunichar(uc))
	return__ = int(__cgo__return__)
	return
}

/*
Performs a single composition step of the
Unicode canonical composition algorithm.

This function includes algorithmic Hangul Jamo composition,
but it is not exactly the inverse of g_unichar_decompose().
No composition can have either of @a or @b equal to zero.
To be precise, this function composes if and only if
there exists a Primary Composite P which is canonically
equivalent to the sequence <@a,@b>.  See the Unicode
Standard for the definition of Primary Composite.

If @a and @b do not compose a new character, @ch is set to zero.

See
[UAX#15](http://unicode.org/reports/tr15/)
for details.
*/
func UnicharCompose(a rune, b rune, ch *C.gunichar) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_compose(C.gunichar(a), C.gunichar(b), ch)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Performs a single decomposition step of the
Unicode canonical decomposition algorithm.

This function does not include compatibility
decompositions. It does, however, include algorithmic
Hangul Jamo decomposition, as well as 'singleton'
decompositions which replace a character by a single
other character. In the case of singletons *@b will
be set to zero.

If @ch is not decomposable, *@a is set to @ch and *@b
is set to zero.

Note that the way Unicode decomposition pairs are
defined, it is guaranteed that @b would not decompose
further, but @a may itself decompose.  To get the full
canonical decomposition for @ch, one would need to
recursively call this function on @a.  Or use
g_unichar_fully_decompose().

See
[UAX#15](http://unicode.org/reports/tr15/)
for details.
*/
func UnicharDecompose(ch rune, a *C.gunichar, b *C.gunichar) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_decompose(C.gunichar(ch), a, b)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines the numeric value of a character as a decimal
digit.
*/
func UnicharDigitValue(c rune) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_unichar_digit_value(C.gunichar(c))
	return__ = int(__cgo__return__)
	return
}

/*
Computes the canonical or compatibility decomposition of a
Unicode character.  For compatibility decomposition,
pass %TRUE for @compat; for canonical decomposition
pass %FALSE for @compat.

The decomposed sequence is placed in @result.  Only up to
@result_len characters are written into @result.  The length
of the full decomposition (irrespective of @result_len) is
returned by the function.  For canonical decomposition,
currently all decompositions are of length at most 4, but
this may change in the future (very unlikely though).
At any rate, Unicode does guarantee that a buffer of length
18 is always enough for both compatibility and canonical
decompositions, so that is the size recommended. This is provided
as %G_UNICHAR_MAX_DECOMPOSITION_LENGTH.

See
[UAX#15](http://unicode.org/reports/tr15/)
for details.
*/
func UnicharFullyDecompose(ch rune, compat bool, result *C.gunichar, result_len int64) (return__ int64) {
	__cgo__compat := C.gboolean(0)
	if compat {
		__cgo__compat = C.gboolean(1)
	}
	var __cgo__return__ C.gsize
	__cgo__return__ = C.g_unichar_fully_decompose(C.gunichar(ch), __cgo__compat, result, C.gsize(result_len))
	return__ = int64(__cgo__return__)
	return
}

/*
In Unicode, some characters are "mirrored". This means that their
images are mirrored horizontally in text that is laid out from right
to left. For instance, "(" would become its mirror image, ")", in
right-to-left text.

If @ch has the Unicode mirrored property and there is another unicode
character that typically has a glyph that is the mirror image of @ch's
glyph and @mirrored_ch is set, it puts that character in the address
pointed to by @mirrored_ch.  Otherwise the original character is put.
*/
func UnicharGetMirrorChar(ch rune, mirrored_ch *C.gunichar) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_get_mirror_char(C.gunichar(ch), mirrored_ch)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Looks up the #GUnicodeScript for a particular character (as defined
by Unicode Standard Annex \#24). No check is made for @ch being a
valid Unicode character; if you pass in invalid character, the
result is undefined.

This function is equivalent to pango_script_for_unichar() and the
two are interchangeable.
*/
func UnicharGetScript(ch rune) (return__ C.GUnicodeScript) {
	return__ = C.g_unichar_get_script(C.gunichar(ch))
	return
}

/*
Determines whether a character is alphanumeric.
Given some UTF-8 text, obtain a character value
with g_utf8_get_char().
*/
func UnicharIsalnum(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_isalnum(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether a character is alphabetic (i.e. a letter).
Given some UTF-8 text, obtain a character value with
g_utf8_get_char().
*/
func UnicharIsalpha(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_isalpha(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether a character is a control character.
Given some UTF-8 text, obtain a character value with
g_utf8_get_char().
*/
func UnicharIscntrl(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_iscntrl(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines if a given character is assigned in the Unicode
standard.
*/
func UnicharIsdefined(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_isdefined(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether a character is numeric (i.e. a digit).  This
covers ASCII 0-9 and also digits in other languages/scripts.  Given
some UTF-8 text, obtain a character value with g_utf8_get_char().
*/
func UnicharIsdigit(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_isdigit(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether a character is printable and not a space
(returns %FALSE for control characters, format characters, and
spaces). g_unichar_isprint() is similar, but returns %TRUE for
spaces. Given some UTF-8 text, obtain a character value with
g_utf8_get_char().
*/
func UnicharIsgraph(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_isgraph(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether a character is a lowercase letter.
Given some UTF-8 text, obtain a character value with
g_utf8_get_char().
*/
func UnicharIslower(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_islower(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether a character is a mark (non-spacing mark,
combining mark, or enclosing mark in Unicode speak).
Given some UTF-8 text, obtain a character value
with g_utf8_get_char().

Note: in most cases where isalpha characters are allowed,
ismark characters should be allowed to as they are essential
for writing most European languages as well as many non-Latin
scripts.
*/
func UnicharIsmark(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_ismark(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether a character is printable.
Unlike g_unichar_isgraph(), returns %TRUE for spaces.
Given some UTF-8 text, obtain a character value with
g_utf8_get_char().
*/
func UnicharIsprint(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_isprint(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether a character is punctuation or a symbol.
Given some UTF-8 text, obtain a character value with
g_utf8_get_char().
*/
func UnicharIspunct(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_ispunct(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether a character is a space, tab, or line separator
(newline, carriage return, etc.).  Given some UTF-8 text, obtain a
character value with g_utf8_get_char().

(Note: don't use this to do word breaking; you have to use
Pango or equivalent to get word breaking right, the algorithm
is fairly complex.)
*/
func UnicharIsspace(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_isspace(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines if a character is titlecase. Some characters in
Unicode which are composites, such as the DZ digraph
have three case variants instead of just two. The titlecase
form is used at the beginning of a word where only the
first letter is capitalized. The titlecase form of the DZ
digraph is U+01F2 LATIN CAPITAL LETTTER D WITH SMALL LETTER Z.
*/
func UnicharIstitle(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_istitle(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines if a character is uppercase.
*/
func UnicharIsupper(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_isupper(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines if a character is typically rendered in a double-width
cell.
*/
func UnicharIswide(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_iswide(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines if a character is typically rendered in a double-width
cell under legacy East Asian locales.  If a character is wide according to
g_unichar_iswide(), then it is also reported wide with this function, but
the converse is not necessarily true. See the
[Unicode Standard Annex #11](http://www.unicode.org/reports/tr11/)
for details.

If a character passes the g_unichar_iswide() test then it will also pass
this test, but not the other way around.  Note that some characters may
pas both this test and g_unichar_iszerowidth().
*/
func UnicharIswideCjk(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_iswide_cjk(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines if a character is a hexidecimal digit.
*/
func UnicharIsxdigit(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_isxdigit(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines if a given character typically takes zero width when rendered.
The return value is %TRUE for all non-spacing and enclosing marks
(e.g., combining accents), format characters, zero-width
space, but not U+00AD SOFT HYPHEN.

A typical use of this function is with one of g_unichar_iswide() or
g_unichar_iswide_cjk() to determine the number of cells a string occupies
when displayed on a grid display (terminals).  However, note that not all
terminals support zero-width rendering of zero-width marks.
*/
func UnicharIszerowidth(c rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_iszerowidth(C.gunichar(c))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Converts a single character to UTF-8.
*/
func UnicharToUtf8(c rune, outbuf string) (return__ int) {
	__cgo__outbuf := (*C.gchar)(unsafe.Pointer(C.CString(outbuf)))
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_unichar_to_utf8(C.gunichar(c), __cgo__outbuf)
	C.free(unsafe.Pointer(__cgo__outbuf))
	return__ = int(__cgo__return__)
	return
}

/*
Converts a character to lower case.
*/
func UnicharTolower(c rune) (return__ rune) {
	var __cgo__return__ C.gunichar
	__cgo__return__ = C.g_unichar_tolower(C.gunichar(c))
	return__ = rune(__cgo__return__)
	return
}

/*
Converts a character to the titlecase.
*/
func UnicharTotitle(c rune) (return__ rune) {
	var __cgo__return__ C.gunichar
	__cgo__return__ = C.g_unichar_totitle(C.gunichar(c))
	return__ = rune(__cgo__return__)
	return
}

/*
Converts a character to uppercase.
*/
func UnicharToupper(c rune) (return__ rune) {
	var __cgo__return__ C.gunichar
	__cgo__return__ = C.g_unichar_toupper(C.gunichar(c))
	return__ = rune(__cgo__return__)
	return
}

/*
Classifies a Unicode character by type.
*/
func UnicharType(c rune) (return__ C.GUnicodeType) {
	return__ = C.g_unichar_type(C.gunichar(c))
	return
}

/*
Checks whether @ch is a valid Unicode character. Some possible
integer values of @ch will not be valid. 0 is considered a valid
character, though it's normally a string terminator.
*/
func UnicharValidate(ch rune) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unichar_validate(C.gunichar(ch))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines the numeric value of a character as a hexidecimal
digit.
*/
func UnicharXdigitValue(c rune) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_unichar_xdigit_value(C.gunichar(c))
	return__ = int(__cgo__return__)
	return
}

// g_unicode_canonical_decomposition is not generated due to deprecation attr

/*
Computes the canonical ordering of a string in-place.
This rearranges decomposed characters in the string
according to their combining classes.  See the Unicode
manual for more information.
*/
func UnicodeCanonicalOrdering(string_ *C.gunichar, len_ int64) {
	C.g_unicode_canonical_ordering(string_, C.gsize(len_))
	return
}

/*
Looks up the Unicode script for @iso15924.  ISO 15924 assigns four-letter
codes to scripts.  For example, the code for Arabic is 'Arab'.
This function accepts four letter codes encoded as a @guint32 in a
big-endian fashion.  That is, the code expected for Arabic is
0x41726162 (0x41 is ASCII code for 'A', 0x72 is ASCII code for 'r', etc).

See
[Codes for the representation of names of scripts](http://unicode.org/iso15924/codelists.html)
for details.
*/
func UnicodeScriptFromIso15924(iso15924 uint32) (return__ C.GUnicodeScript) {
	return__ = C.g_unicode_script_from_iso15924(C.guint32(iso15924))
	return
}

/*
Looks up the ISO 15924 code for @script.  ISO 15924 assigns four-letter
codes to scripts.  For example, the code for Arabic is 'Arab'.  The
four letter codes are encoded as a @guint32 by this function in a
big-endian fashion.  That is, the code returned for Arabic is
0x41726162 (0x41 is ASCII code for 'A', 0x72 is ASCII code for 'r', etc).

See
[Codes for the representation of names of scripts](http://unicode.org/iso15924/codelists.html)
for details.
*/
func UnicodeScriptToIso15924(script C.GUnicodeScript) (return__ uint32) {
	var __cgo__return__ C.guint32
	__cgo__return__ = C.g_unicode_script_to_iso15924(script)
	return__ = uint32(__cgo__return__)
	return
}

func UnixErrorQuark() (return__ C.GQuark) {
	return__ = C.g_unix_error_quark()
	return
}

/*
Sets a function to be called when the IO condition, as specified by
@condition becomes true for @fd.

@function will be called when the specified IO condition becomes
%TRUE.  The function is expected to clear whatever event caused the
IO condition to become true and return %TRUE in order to be notified
when it happens again.  If @function returns %FALSE then the watch
will be cancelled.

The return value of this function can be passed to g_source_remove()
to cancel the watch at any time that it exists.

The source will never close the fd -- you must do it yourself.
*/
func UnixFdAdd(fd int, condition C.GIOCondition, function C.GUnixFDSourceFunc, user_data unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_unix_fd_add(C.gint(fd), condition, function, (C.gpointer)(user_data))
	return__ = uint(__cgo__return__)
	return
}

/*
Sets a function to be called when the IO condition, as specified by
@condition becomes true for @fd.

This is the same as g_unix_fd_add(), except that it allows you to
specify a non-default priority and a provide a #GDestroyNotify for
@user_data.
*/
func UnixFdAddFull(priority int, fd int, condition C.GIOCondition, function C.GUnixFDSourceFunc, user_data unsafe.Pointer, notify C.GDestroyNotify) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_unix_fd_add_full(C.gint(priority), C.gint(fd), condition, function, (C.gpointer)(user_data), notify)
	return__ = uint(__cgo__return__)
	return
}

/*
Creates a #GSource to watch for a particular IO condition on a file
descriptor.

The source will never close the fd -- you must do it yourself.
*/
func UnixFdSourceNew(fd int, condition C.GIOCondition) (return__ *Source) {
	var __cgo__return__ *C.GSource
	__cgo__return__ = C.g_unix_fd_source_new(C.gint(fd), condition)
	return__ = (*Source)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Similar to the UNIX pipe() call, but on modern systems like Linux
uses the pipe2() system call, which atomically creates a pipe with
the configured flags. The only supported flag currently is
%FD_CLOEXEC. If for example you want to configure %O_NONBLOCK, that
must still be done separately with fcntl().

This function does not take %O_CLOEXEC, it takes %FD_CLOEXEC as if
for fcntl(); these are different on Linux/glibc.
*/
func UnixOpenPipe(fds *C.gint, flags int) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unix_open_pipe(fds, C.gint(flags), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Control the non-blocking state of the given file descriptor,
according to @nonblock. On most systems this uses %O_NONBLOCK, but
on some older ones may use %O_NDELAY.
*/
func UnixSetFdNonblocking(fd int, nonblock bool) (return__ bool, __err__ error) {
	__cgo__nonblock := C.gboolean(0)
	if nonblock {
		__cgo__nonblock = C.gboolean(1)
	}
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_unix_set_fd_nonblocking(C.gint(fd), __cgo__nonblock, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
A convenience function for g_unix_signal_source_new(), which
attaches to the default #GMainContext.  You can remove the watch
using g_source_remove().
*/
func UnixSignalAdd(signum int, handler C.GSourceFunc, user_data unsafe.Pointer) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_unix_signal_add(C.gint(signum), handler, (C.gpointer)(user_data))
	return__ = uint(__cgo__return__)
	return
}

/*
A convenience function for g_unix_signal_source_new(), which
attaches to the default #GMainContext.  You can remove the watch
using g_source_remove().
*/
func UnixSignalAddFull(priority int, signum int, handler C.GSourceFunc, user_data unsafe.Pointer, notify C.GDestroyNotify) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.g_unix_signal_add_full(C.gint(priority), C.gint(signum), handler, (C.gpointer)(user_data), notify)
	return__ = uint(__cgo__return__)
	return
}

/*
Create a #GSource that will be dispatched upon delivery of the UNIX
signal @signum.  In GLib versions before 2.36, only `SIGHUP`, `SIGINT`,
`SIGTERM` can be monitored.  In GLib 2.36, `SIGUSR1` and `SIGUSR2`
were added.

Note that unlike the UNIX default, all sources which have created a
watch will be dispatched, regardless of which underlying thread
invoked g_unix_signal_source_new().

For example, an effective use of this function is to handle `SIGTERM`
cleanly; flushing any outstanding files, and then calling
g_main_loop_quit ().  It is not safe to do any of this a regular
UNIX signal handler; your handler may be invoked while malloc() or
another library function is running, causing reentrancy if you
attempt to use it from the handler.  None of the GLib/GObject API
is safe against this kind of reentrancy.

The interaction of this source when combined with native UNIX
functions like sigprocmask() is not defined.

The source will not initially be associated with any #GMainContext
and must be added to one with g_source_attach() before it will be
executed.
*/
func UnixSignalSourceNew(signum int) (return__ *Source) {
	var __cgo__return__ *C.GSource
	__cgo__return__ = C.g_unix_signal_source_new(C.gint(signum))
	return__ = (*Source)(unsafe.Pointer(__cgo__return__))
	return
}

/*
A wrapper for the POSIX unlink() function. The unlink() function
deletes a name from the filesystem. If this was the last link to the
file and no processes have it opened, the diskspace occupied by the
file is freed.

See your C library manual for more details about unlink(). Note
that on Windows, it is in general not possible to delete files that
are open to some process, or mapped into memory.
*/
func Unlink(filename string) (return__ int) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	var __cgo__return__ C.int
	__cgo__return__ = C.g_unlink(__cgo__filename)
	C.free(unsafe.Pointer(__cgo__filename))
	return__ = int(__cgo__return__)
	return
}

/*
Removes an environment variable from the environment.

Note that on some systems, when variables are overwritten, the
memory used for the previous variables and its value isn't reclaimed.

You should be mindful of the fact that environment variable handling
in UNIX is not thread-safe, and your program may crash if one thread
calls g_unsetenv() while another thread is calling getenv(). (And note
that many functions, such as gettext(), call getenv() internally.) This
function is only safe to use at the very start of your program, before
creating any other threads (or creating objects that create worker
threads of their own).

If you need to set up the environment for a child process, you can
use g_get_environ() to get an environment array, modify that with
g_environ_setenv() and g_environ_unsetenv(), and then pass that
array directly to execvpe(), g_spawn_async(), or the like.
*/
func Unsetenv(variable string) {
	__cgo__variable := (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	C.g_unsetenv(__cgo__variable)
	C.free(unsafe.Pointer(__cgo__variable))
	return
}

/*
Escapes a string for use in a URI.

Normally all characters that are not "unreserved" (i.e. ASCII alphanumerical
characters plus dash, dot, underscore and tilde) are escaped.
But if you specify characters in @reserved_chars_allowed they are not
escaped. This is useful for the "reserved" characters in the URI
specification, since those are allowed unescaped in some portions of
a URI.
*/
func UriEscapeString(unescaped string, reserved_chars_allowed string, allow_utf8 bool) (return__ string) {
	__cgo__unescaped := C.CString(unescaped)
	__cgo__reserved_chars_allowed := C.CString(reserved_chars_allowed)
	__cgo__allow_utf8 := C.gboolean(0)
	if allow_utf8 {
		__cgo__allow_utf8 = C.gboolean(1)
	}
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_uri_escape_string(__cgo__unescaped, __cgo__reserved_chars_allowed, __cgo__allow_utf8)
	C.free(unsafe.Pointer(__cgo__unescaped))
	C.free(unsafe.Pointer(__cgo__reserved_chars_allowed))
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Splits an URI list conforming to the text/uri-list
mime type defined in RFC 2483 into individual URIs,
discarding any comments. The URIs are not validated.
*/
func UriListExtractUris(uri_list string) (return__ []string) {
	__cgo__uri_list := (*C.gchar)(unsafe.Pointer(C.CString(uri_list)))
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.g_uri_list_extract_uris(__cgo__uri_list)
	C.free(unsafe.Pointer(__cgo__uri_list))
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Gets the scheme portion of a URI string. RFC 3986 decodes the scheme as:
|[
URI = scheme ":" hier-part [ "?" query ] [ "#" fragment ]
]|
Common schemes include "file", "http", "svn+ssh", etc.
*/
func UriParseScheme(uri string) (return__ string) {
	__cgo__uri := C.CString(uri)
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_uri_parse_scheme(__cgo__uri)
	C.free(unsafe.Pointer(__cgo__uri))
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Unescapes a segment of an escaped string.

If any of the characters in @illegal_characters or the character zero appears
as an escaped character in @escaped_string then that is an error and %NULL
will be returned. This is useful it you want to avoid for instance having a
slash being expanded in an escaped path element, which might confuse pathname
handling.
*/
func UriUnescapeSegment(escaped_string string, escaped_string_end string, illegal_characters string) (return__ string) {
	__cgo__escaped_string := C.CString(escaped_string)
	__cgo__escaped_string_end := C.CString(escaped_string_end)
	__cgo__illegal_characters := C.CString(illegal_characters)
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_uri_unescape_segment(__cgo__escaped_string, __cgo__escaped_string_end, __cgo__illegal_characters)
	C.free(unsafe.Pointer(__cgo__escaped_string))
	C.free(unsafe.Pointer(__cgo__escaped_string_end))
	C.free(unsafe.Pointer(__cgo__illegal_characters))
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Unescapes a whole escaped string.

If any of the characters in @illegal_characters or the character zero appears
as an escaped character in @escaped_string then that is an error and %NULL
will be returned. This is useful it you want to avoid for instance having a
slash being expanded in an escaped path element, which might confuse pathname
handling.
*/
func UriUnescapeString(escaped_string string, illegal_characters string) (return__ string) {
	__cgo__escaped_string := C.CString(escaped_string)
	__cgo__illegal_characters := C.CString(illegal_characters)
	var __cgo__return__ *C.char
	__cgo__return__ = C.g_uri_unescape_string(__cgo__escaped_string, __cgo__illegal_characters)
	C.free(unsafe.Pointer(__cgo__escaped_string))
	C.free(unsafe.Pointer(__cgo__illegal_characters))
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Pauses the current thread for the given number of microseconds.

There are 1 million microseconds per second (represented by the
#G_USEC_PER_SEC macro). g_usleep() may have limited precision,
depending on hardware and operating system; don't rely on the exact
length of the sleep.
*/
func Usleep(microseconds int64) {
	C.g_usleep(C.gulong(microseconds))
	return
}

/*
Convert a string from UTF-16 to UCS-4. The result will be
nul-terminated.
*/
func Utf16ToUcs4(str *C.gunichar2, len_ int64, items_read *C.glong, items_written *C.glong) (return__ *C.gunichar, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.g_utf16_to_ucs4(str, C.glong(len_), items_read, items_written, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Convert a string from UTF-16 to UTF-8. The result will be
terminated with a 0 byte.

Note that the input is expected to be already in native endianness,
an initial byte-order-mark character is not handled specially.
g_convert() can be used to convert a byte buffer of UTF-16 data of
ambiguous endianess.

Further note that this function does not validate the result
string; it may e.g. include embedded NUL characters. The only
validation done by this function is to ensure that the input can
be correctly interpreted as UTF-16, i.e. it doesn't contain
things unpaired surrogates.
*/
func Utf16ToUtf8(str *C.gunichar2, len_ int64, items_read *C.glong, items_written *C.glong) (return__ string, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_utf16_to_utf8(str, C.glong(len_), items_read, items_written, &__cgo_error__)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Converts a string into a form that is independent of case. The
result will not correspond to any particular case, but can be
compared for equality or ordered with the results of calling
g_utf8_casefold() on other strings.

Note that calling g_utf8_casefold() followed by g_utf8_collate() is
only an approximation to the correct linguistic case insensitive
ordering, though it is a fairly good one. Getting this exactly
right would require a more sophisticated collation function that
takes case sensitivity into account. GLib does not currently
provide such a function.
*/
func Utf8Casefold(str string, len_ int64) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_utf8_casefold(__cgo__str, C.gssize(len_))
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Compares two strings for ordering using the linguistically
correct rules for the [current locale][setlocale].
When sorting a large number of strings, it will be significantly
faster to obtain collation keys with g_utf8_collate_key() and
compare the keys with strcmp() when sorting instead of sorting
the original strings.
*/
func Utf8Collate(str1 string, str2 string) (return__ int) {
	__cgo__str1 := (*C.gchar)(unsafe.Pointer(C.CString(str1)))
	__cgo__str2 := (*C.gchar)(unsafe.Pointer(C.CString(str2)))
	var __cgo__return__ C.gint
	__cgo__return__ = C.g_utf8_collate(__cgo__str1, __cgo__str2)
	C.free(unsafe.Pointer(__cgo__str1))
	C.free(unsafe.Pointer(__cgo__str2))
	return__ = int(__cgo__return__)
	return
}

/*
Converts a string into a collation key that can be compared
with other collation keys produced by the same function using
strcmp().

The results of comparing the collation keys of two strings
with strcmp() will always be the same as comparing the two
original keys with g_utf8_collate().

Note that this function depends on the [current locale][setlocale].
*/
func Utf8CollateKey(str string, len_ int64) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_utf8_collate_key(__cgo__str, C.gssize(len_))
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts a string into a collation key that can be compared
with other collation keys produced by the same function using strcmp().

In order to sort filenames correctly, this function treats the dot '.'
as a special case. Most dictionary orderings seem to consider it
insignificant, thus producing the ordering "event.c" "eventgenerator.c"
"event.h" instead of "event.c" "event.h" "eventgenerator.c". Also, we
would like to treat numbers intelligently so that "file1" "file10" "file5"
is sorted as "file1" "file5" "file10".

Note that this function depends on the [current locale][setlocale].
*/
func Utf8CollateKeyForFilename(str string, len_ int64) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_utf8_collate_key_for_filename(__cgo__str, C.gssize(len_))
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Finds the start of the next UTF-8 character in the string after @p.

@p does not have to be at the beginning of a UTF-8 character. No check
is made to see if the character found is actually valid other than
it starts with an appropriate byte.
*/
func Utf8FindNextChar(p string, end string) (return__ string) {
	__cgo__p := (*C.gchar)(unsafe.Pointer(C.CString(p)))
	__cgo__end := (*C.gchar)(unsafe.Pointer(C.CString(end)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_utf8_find_next_char(__cgo__p, __cgo__end)
	C.free(unsafe.Pointer(__cgo__p))
	C.free(unsafe.Pointer(__cgo__end))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Given a position @p with a UTF-8 encoded string @str, find the start
of the previous UTF-8 character starting before @p. Returns %NULL if no
UTF-8 characters are present in @str before @p.

@p does not have to be at the beginning of a UTF-8 character. No check
is made to see if the character found is actually valid other than
it starts with an appropriate byte.
*/
func Utf8FindPrevChar(str string, p string) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	__cgo__p := (*C.gchar)(unsafe.Pointer(C.CString(p)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_utf8_find_prev_char(__cgo__str, __cgo__p)
	C.free(unsafe.Pointer(__cgo__str))
	C.free(unsafe.Pointer(__cgo__p))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts a sequence of bytes encoded as UTF-8 to a Unicode character.

If @p does not point to a valid UTF-8 encoded character, results
are undefined. If you are not sure that the bytes are complete
valid Unicode characters, you should use g_utf8_get_char_validated()
instead.
*/
func Utf8GetChar(p string) (return__ rune) {
	__cgo__p := (*C.gchar)(unsafe.Pointer(C.CString(p)))
	var __cgo__return__ C.gunichar
	__cgo__return__ = C.g_utf8_get_char(__cgo__p)
	C.free(unsafe.Pointer(__cgo__p))
	return__ = rune(__cgo__return__)
	return
}

/*
Convert a sequence of bytes encoded as UTF-8 to a Unicode character.
This function checks for incomplete characters, for invalid characters
such as characters that are out of the range of Unicode, and for
overlong encodings of valid characters.
*/
func Utf8GetCharValidated(p string, max_len int64) (return__ rune) {
	__cgo__p := (*C.gchar)(unsafe.Pointer(C.CString(p)))
	var __cgo__return__ C.gunichar
	__cgo__return__ = C.g_utf8_get_char_validated(__cgo__p, C.gssize(max_len))
	C.free(unsafe.Pointer(__cgo__p))
	return__ = rune(__cgo__return__)
	return
}

/*
Converts a string into canonical form, standardizing
such issues as whether a character with an accent
is represented as a base character and combining
accent or as a single precomposed character. The
string has to be valid UTF-8, otherwise %NULL is
returned. You should generally call g_utf8_normalize()
before comparing two Unicode strings.

The normalization mode %G_NORMALIZE_DEFAULT only
standardizes differences that do not affect the
text content, such as the above-mentioned accent
representation. %G_NORMALIZE_ALL also standardizes
the "compatibility" characters in Unicode, such
as SUPERSCRIPT THREE to the standard forms
(in this case DIGIT THREE). Formatting information
may be lost but for most text operations such
characters should be considered the same.

%G_NORMALIZE_DEFAULT_COMPOSE and %G_NORMALIZE_ALL_COMPOSE
are like %G_NORMALIZE_DEFAULT and %G_NORMALIZE_ALL,
but returned a result with composed forms rather
than a maximally decomposed form. This is often
useful if you intend to convert the string to
a legacy encoding or pass it to a system with
less capable Unicode handling.
*/
func Utf8Normalize(str string, len_ int64, mode C.GNormalizeMode) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_utf8_normalize(__cgo__str, C.gssize(len_), mode)
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts from an integer character offset to a pointer to a position
within the string.

Since 2.10, this function allows to pass a negative @offset to
step backwards. It is usually worth stepping backwards from the end
instead of forwards if @offset is in the last fourth of the string,
since moving forward is about 3 times faster than moving backward.

Note that this function doesn't abort when reaching the end of @str.
Therefore you should be sure that @offset is within string boundaries
before calling that function. Call g_utf8_strlen() when unsure.
This limitation exists as this function is called frequently during
text rendering and therefore has to be as fast as possible.
*/
func Utf8OffsetToPointer(str string, offset int64) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_utf8_offset_to_pointer(__cgo__str, C.glong(offset))
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts from a pointer to position within a string to a integer
character offset.

Since 2.10, this function allows @pos to be before @str, and returns
a negative offset in this case.
*/
func Utf8PointerToOffset(str string, pos string) (return__ int64) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	__cgo__pos := (*C.gchar)(unsafe.Pointer(C.CString(pos)))
	var __cgo__return__ C.glong
	__cgo__return__ = C.g_utf8_pointer_to_offset(__cgo__str, __cgo__pos)
	C.free(unsafe.Pointer(__cgo__str))
	C.free(unsafe.Pointer(__cgo__pos))
	return__ = int64(__cgo__return__)
	return
}

/*
Finds the previous UTF-8 character in the string before @p.

@p does not have to be at the beginning of a UTF-8 character. No check
is made to see if the character found is actually valid other than
it starts with an appropriate byte. If @p might be the first
character of the string, you must use g_utf8_find_prev_char() instead.
*/
func Utf8PrevChar(p string) (return__ string) {
	__cgo__p := (*C.gchar)(unsafe.Pointer(C.CString(p)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_utf8_prev_char(__cgo__p)
	C.free(unsafe.Pointer(__cgo__p))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Finds the leftmost occurrence of the given Unicode character
in a UTF-8 encoded string, while limiting the search to @len bytes.
If @len is -1, allow unbounded search.
*/
func Utf8Strchr(p string, len_ int64, c rune) (return__ string) {
	__cgo__p := (*C.gchar)(unsafe.Pointer(C.CString(p)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_utf8_strchr(__cgo__p, C.gssize(len_), C.gunichar(c))
	C.free(unsafe.Pointer(__cgo__p))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts all Unicode characters in the string that have a case
to lowercase. The exact manner that this is done depends
on the current locale, and may result in the number of
characters in the string changing.
*/
func Utf8Strdown(str string, len_ int64) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_utf8_strdown(__cgo__str, C.gssize(len_))
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Computes the length of the string in characters, not including
the terminating nul character. If the @max'th byte falls in the
middle of a character, the last (partial) character is not counted.
*/
func Utf8Strlen(p string, max int64) (return__ int64) {
	__cgo__p := (*C.gchar)(unsafe.Pointer(C.CString(p)))
	var __cgo__return__ C.glong
	__cgo__return__ = C.g_utf8_strlen(__cgo__p, C.gssize(max))
	C.free(unsafe.Pointer(__cgo__p))
	return__ = int64(__cgo__return__)
	return
}

/*
Like the standard C strncpy() function, but copies a given number
of characters instead of a given number of bytes. The @src string
must be valid UTF-8 encoded text. (Use g_utf8_validate() on all
text before trying to use UTF-8 utility functions with it.)
*/
func Utf8Strncpy(dest string, src string, n int64) (return__ string) {
	__cgo__dest := (*C.gchar)(unsafe.Pointer(C.CString(dest)))
	__cgo__src := (*C.gchar)(unsafe.Pointer(C.CString(src)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_utf8_strncpy(__cgo__dest, __cgo__src, C.gsize(n))
	C.free(unsafe.Pointer(__cgo__dest))
	C.free(unsafe.Pointer(__cgo__src))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Find the rightmost occurrence of the given Unicode character
in a UTF-8 encoded string, while limiting the search to @len bytes.
If @len is -1, allow unbounded search.
*/
func Utf8Strrchr(p string, len_ int64, c rune) (return__ string) {
	__cgo__p := (*C.gchar)(unsafe.Pointer(C.CString(p)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_utf8_strrchr(__cgo__p, C.gssize(len_), C.gunichar(c))
	C.free(unsafe.Pointer(__cgo__p))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Reverses a UTF-8 string. @str must be valid UTF-8 encoded text.
(Use g_utf8_validate() on all text before trying to use UTF-8
utility functions with it.)

This function is intended for programmatic uses of reversed strings.
It pays no attention to decomposed characters, combining marks, byte
order marks, directional indicators (LRM, LRO, etc) and similar
characters which might need special handling when reversing a string
for display purposes.

Note that unlike g_strreverse(), this function returns
newly-allocated memory, which should be freed with g_free() when
no longer needed.
*/
func Utf8Strreverse(str string, len_ int64) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_utf8_strreverse(__cgo__str, C.gssize(len_))
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Converts all Unicode characters in the string that have a case
to uppercase. The exact manner that this is done depends
on the current locale, and may result in the number of
characters in the string increasing. (For instance, the
German ess-zet will be changed to SS.)
*/
func Utf8Strup(str string, len_ int64) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_utf8_strup(__cgo__str, C.gssize(len_))
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Copies a substring out of a UTF-8 encoded string.
The substring will contain @end_pos - @start_pos characters.
*/
func Utf8Substring(str string, start_pos int64, end_pos int64) (return__ string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_utf8_substring(__cgo__str, C.glong(start_pos), C.glong(end_pos))
	C.free(unsafe.Pointer(__cgo__str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Convert a string from UTF-8 to a 32-bit fixed width
representation as UCS-4. A trailing 0 character will be added to the
string after the converted text.
*/
func Utf8ToUcs4(str string, len_ int64, items_read *C.glong, items_written *C.glong) (return__ *C.gunichar, __err__ error) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo_error__ *C.GError
	return__ = C.g_utf8_to_ucs4(__cgo__str, C.glong(len_), items_read, items_written, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__str))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Convert a string from UTF-8 to a 32-bit fixed width
representation as UCS-4, assuming valid UTF-8 input.
This function is roughly twice as fast as g_utf8_to_ucs4()
but does no error checking on the input. A trailing 0 character
will be added to the string after the converted text.
*/
func Utf8ToUcs4Fast(str string, len_ int64, items_written *C.glong) (return__ *C.gunichar) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	return__ = C.g_utf8_to_ucs4_fast(__cgo__str, C.glong(len_), items_written)
	C.free(unsafe.Pointer(__cgo__str))
	return
}

/*
Convert a string from UTF-8 to UTF-16. A 0 character will be
added to the result after the converted text.
*/
func Utf8ToUtf16(str string, len_ int64, items_read *C.glong, items_written *C.glong) (return__ *C.gunichar2, __err__ error) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	var __cgo_error__ *C.GError
	return__ = C.g_utf8_to_utf16(__cgo__str, C.glong(len_), items_read, items_written, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__str))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Validates UTF-8 encoded text. @str is the text to validate;
if @str is nul-terminated, then @max_len can be -1, otherwise
@max_len should be the number of bytes to validate.
If @end is non-%NULL, then the end of the valid range
will be stored there (i.e. the start of the first invalid
character if some bytes were invalid, or the end of the text
being validated otherwise).

Note that g_utf8_validate() returns %FALSE if @max_len is
positive and any of the @max_len bytes are nul.

Returns %TRUE if all of @str was valid. Many GLib and GTK+
routines require valid UTF-8 as input; so data read from a file
or the network should be checked with g_utf8_validate() before
doing anything else with it.
*/
func Utf8Validate(str []byte, max_len int64) (end string, return__ bool) {
	__header__str := (*reflect.SliceHeader)(unsafe.Pointer(&str))
	var __cgo__end *C.gchar
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_utf8_validate((*C.gchar)(unsafe.Pointer(__header__str.Data)), C.gssize(max_len), &__cgo__end)
	end = C.GoString((*C.char)(unsafe.Pointer(__cgo__end)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// g_variant_get_gtype is not generated due to explicit ignore

/*
Determines if a given string is a valid D-Bus object path.  You
should ensure that a string is a valid D-Bus object path before
passing it to g_variant_new_object_path().

A valid object path starts with '/' followed by zero or more
sequences of characters separated by '/' characters.  Each sequence
must contain only the characters "[A-Z][a-z][0-9]_".  No sequence
(including the one following the final '/' character) may be empty.
*/
func VariantIsObjectPath(string_ string) (return__ bool) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_variant_is_object_path(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines if a given string is a valid D-Bus type signature.  You
should ensure that a string is a valid D-Bus type signature before
passing it to g_variant_new_signature().

D-Bus type signatures consist of zero or more definite #GVariantType
strings in sequence.
*/
func VariantIsSignature(string_ string) (return__ bool) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_variant_is_signature(__cgo__string_)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Parses a #GVariant from a text representation.

A single #GVariant is parsed from the content of @text.

The format is described [here][gvariant-text].

The memory at @limit will never be accessed and the parser behaves as
if the character at @limit is the nul terminator.  This has the
effect of bounding @text.

If @endptr is non-%NULL then @text is permitted to contain data
following the value that this function parses and @endptr will be
updated to point to the first character past the end of the text
parsed by this function.  If @endptr is %NULL and there is extra data
then an error is returned.

If @type is non-%NULL then the value will be parsed to have that
type.  This may result in additional parse errors (in the case that
the parsed value doesn't fit the type) but may also result in fewer
errors (in the case that the type would have been ambiguous, such as
with empty arrays).

In the event that the parsing is successful, the resulting #GVariant
is returned.

In case of any error, %NULL will be returned.  If @error is non-%NULL
then it will be set to reflect the error that occurred.

Officially, the language understood by the parser is "any string
produced by g_variant_print()".
*/
func VariantParse(type_ *VariantType, text string, limit string, endptr **C.gchar) (return__ *Variant, __err__ error) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	__cgo__limit := (*C.gchar)(unsafe.Pointer(C.CString(limit)))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GVariant
	__cgo__return__ = C.g_variant_parse((*C.GVariantType)(unsafe.Pointer(type_)), __cgo__text, __cgo__limit, endptr, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__text))
	C.free(unsafe.Pointer(__cgo__limit))
	return__ = (*Variant)(unsafe.Pointer(__cgo__return__))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Pretty-prints a message showing the context of a #GVariant parse
error within the string for which parsing was attempted.

The resulting string is suitable for output to the console or other
monospace media where newlines are treated in the usual way.

The message will typically look something like one of the following:

|[
unterminated string constant:
  (1, 2, 3, 'abc
            ^^^^
]|

or

|[
unable to find a common type:
  [1, 2, 3, 'str']
   ^        ^^^^^
]|

The format of the message may change in a future version.

@error must have come from a failed attempt to g_variant_parse() and
@source_str must be exactly the same string that caused the error.
If @source_str was not nul-terminated when you passed it to
g_variant_parse() then you must add nul termination before using this
function.
*/
func VariantParseErrorPrintContext(error_ *Error, source_str string) (return__ string) {
	__cgo__source_str := (*C.gchar)(unsafe.Pointer(C.CString(source_str)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.g_variant_parse_error_print_context((*C.GError)(unsafe.Pointer(error_)), __cgo__source_str)
	C.free(unsafe.Pointer(__cgo__source_str))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

func VariantParseErrorQuark() (return__ C.GQuark) {
	return__ = C.g_variant_parse_error_quark()
	return
}

// g_variant_parser_get_error_quark is not generated due to deprecation attr

func VariantTypeChecked(arg0 string) (return__ *VariantType) {
	__cgo__arg0 := (*C.gchar)(unsafe.Pointer(C.CString(arg0)))
	var __cgo__return__ *C.GVariantType
	__cgo__return__ = C.g_variant_type_checked_(__cgo__arg0)
	C.free(unsafe.Pointer(__cgo__arg0))
	return__ = (*VariantType)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Checks if @type_string is a valid GVariant type string.  This call is
equivalent to calling g_variant_type_string_scan() and confirming
that the following character is a nul terminator.
*/
func VariantTypeStringIsValid(type_string string) (return__ bool) {
	__cgo__type_string := (*C.gchar)(unsafe.Pointer(C.CString(type_string)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_variant_type_string_is_valid(__cgo__type_string)
	C.free(unsafe.Pointer(__cgo__type_string))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Scan for a single complete and valid GVariant type string in @string.
The memory pointed to by @limit (or bytes beyond it) is never
accessed.

If a valid type string is found, @endptr is updated to point to the
first character past the end of the string that was found and %TRUE
is returned.

If there is no valid type string starting at @string, or if the type
string does not end before @limit then %FALSE is returned.

For the simple case of checking if a string is a valid type string,
see g_variant_type_string_is_valid().
*/
func VariantTypeStringScan(string_ string, limit string) (endptr string, return__ bool) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	__cgo__limit := (*C.gchar)(unsafe.Pointer(C.CString(limit)))
	var __cgo__endptr *C.gchar
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.g_variant_type_string_scan(__cgo__string_, __cgo__limit, &__cgo__endptr)
	C.free(unsafe.Pointer(__cgo__string_))
	C.free(unsafe.Pointer(__cgo__limit))
	endptr = C.GoString((*C.char)(unsafe.Pointer(__cgo__endptr)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// g_vasprintf is not generated due to varargs

// g_vfprintf is not generated due to varargs

// g_vprintf is not generated due to varargs

// g_vsnprintf is not generated due to varargs

// g_vsprintf is not generated due to varargs

func WarnMessage(domain string, file string, line int, func_ string, warnexpr string) {
	__cgo__domain := C.CString(domain)
	__cgo__file := C.CString(file)
	__cgo__func_ := C.CString(func_)
	__cgo__warnexpr := C.CString(warnexpr)
	C.g_warn_message(__cgo__domain, __cgo__file, C.int(line), __cgo__func_, __cgo__warnexpr)
	C.free(unsafe.Pointer(__cgo__domain))
	C.free(unsafe.Pointer(__cgo__file))
	C.free(unsafe.Pointer(__cgo__func_))
	C.free(unsafe.Pointer(__cgo__warnexpr))
	return
}
