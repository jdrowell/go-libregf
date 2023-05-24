package libregf

/*
#cgo LDFLAGS: -lregf
#include <libregf.h>
*/
import "C"

import (
    "fmt"
    "strings"
    "unsafe"
)

// Value is an opaque struct to group related method calls
type Value       C.libregf_value_t
// MultiString is an opaque struct to group related method calls
type MultiString C.libregf_multi_string_t

// Free frees memory allocated in C for the hidden Value struct.
// It wraps libregf_value_free().
func (value *Value) Free() error { 
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_value_free((**C.libregf_value_t)(unsafe.Pointer(&value)), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return fmt.Errorf("%s", pe.String())
    } else {
        return nil
    }
}

// NameLen returns the length (in bytes) of the Value's name.
// It wraps libregf_value_get_utf8_name_size().
// You usually call it to know how much space to allocate before calling
// libregf_value_get_utf8_name().
// You don't need to call this function if you call Name(), which calls NameLen().
func (value *Value) NameLen() (int, error) { 
    var namelen C.size_t
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_value_get_utf8_name_size((*C.libregf_value_t)(value), (*C.size_t)(&namelen), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return -1, fmt.Errorf("%s", pe.String())
    } else {
        return int(namelen), nil
    }
}

// Name returns the Value's name.
// It wraps libregf_value_get_utf8_name().
// It returns a regular Go string, so you don't have to worry about the
// underlying C calls and memory allocations.
func (value *Value) Name() (string, error) { 
    namelen, err := value.NameLen()
    if err != nil { return "", err }

    buffer := make([]byte, namelen+1)
    cstr := C.CString(string(buffer[:namelen]))
    defer C.free(unsafe.Pointer(cstr))
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_value_get_utf8_name((*C.libregf_value_t)(value), (*C.uint8_t)(unsafe.Pointer(cstr)), C.ulong(namelen), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return "", fmt.Errorf("%s", pe.String())
    } else {
        return C.GoString(cstr), nil
    }
}

// Type returns the value's type.
// It wraps libregf_value_get_value_type().
func (value *Value) Type() (int, error) { 
    var _type C.uint32_t
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_value_get_value_type((*C.libregf_value_t)(value), &_type, (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return -1, fmt.Errorf("%s", pe.String())
    } else {
        return int(_type), nil
    }
}

// TStringLen returns the length (in bytes) of a value of type LIBREGF_VALUE_TYPE_STRING
// It wraps libregf_value_get_value_utf8_string_size().
// You don't need to call this function if you call TString(), which calls TStringLen().
func (value *Value) TStringLen() (int, error) { 
    var tlen C.size_t
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_value_get_value_utf8_string_size((*C.libregf_value_t)(value), (*C.size_t)(&tlen), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return -1, fmt.Errorf("%s", pe.String())
    } else {
        return int(tlen), nil
    }
}

// TString returns a value of type LIBREGF_VALUE_TYPE_STRING as a Go string
// It wraps libregf_value_get_value_utf8_string().
func (value *Value) TString() (string, error) { 
    tlen, err := value.TStringLen()
    if err != nil { return "", err }

    buffer := make([]byte, tlen+1)
    cstr := C.CString(string(buffer[:tlen]))
    defer C.free(unsafe.Pointer(cstr))
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_value_get_value_utf8_string((*C.libregf_value_t)(value), (*C.uint8_t)(unsafe.Pointer(cstr)), C.ulong(tlen), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return "", fmt.Errorf("%s", pe.String())
    } else {
        return C.GoString(cstr), nil
    }
}

// TBinaryLen returns the length (in bytes) of a value of type LIBREGF_VALUE_TYPE_BINARY_DATA
// It wraps libregf_value_get_value_binary_data_size().
// You don't need to call this function if you call TBinary(), which calls TBinaryLen().
func (value *Value) TBinaryLen() (int, error) { 
    var tlen C.size_t
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_value_get_value_binary_data_size((*C.libregf_value_t)(value), (*C.size_t)(&tlen), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return -1, fmt.Errorf("%s", pe.String())
    } else {
        return int(tlen), nil
    }
}

// TBinary returns a value of type LIBREGF_VALUE_TYPE_BINARY_DATANG as a Go []byte
// It wraps libregf_value_get_value_binary_data().
func (value *Value) TBinary() ([]byte, error) { 
    tlen, err := value.TBinaryLen()
    if err != nil { return []byte{}, err }

    buffer := make([]byte, tlen)
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_value_get_value_binary_data((*C.libregf_value_t)(value), (*C.uint8_t)(unsafe.Pointer(&buffer[0])), C.ulong(tlen), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return []byte{}, fmt.Errorf("%s", pe.String())
    } else {
        return buffer, nil
    }
}

// Tint32 returns a value of type LIBREGF_VALUE_TYPE_INTEGER_32BIT_LITTLE_ENDIAN as a Go int
// It wraps libregf_value_get_value_32bit().
func (value *Value) Tint32() (int, error) { 
    var cint C.uint32_t
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_value_get_value_32bit((*C.libregf_value_t)(value), &cint, (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return -1, fmt.Errorf("%s", pe.String())
    } else {
        return int(cint), nil
    }
}

// Tint64 returns a value of type LIBREGF_VALUE_TYPE_INTEGER_64BIT_LITTLE_ENDIAN as a Go int
// It wraps libregf_value_get_value_64bit().
func (value *Value) Tint64() (int, error) { 
    var cint C.uint64_t
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_value_get_value_64bit((*C.libregf_value_t)(value), &cint, (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return -1, fmt.Errorf("%s", pe.String())
    } else {
        return int(cint), nil
    }
}

// TMultiString returns a value of type LIBREGF_VALUE_TYPE_MULTI_VALUE_STRING as a
// pointer to a MultiString struct. You can only access the inner strings through
// the (*MultiString) methods.
// It wraps libregf_value_get_value_multi_string().
func (value *Value) TMultiString() (*MultiString, error) { 
    var ms MultiString
    ppms := unsafe.Pointer(&ms)
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_value_get_value_multi_string((*C.libregf_value_t)(value), (**C.libregf_multi_string_t)(ppms), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()
    pms := *(**MultiString)(ppms)

    if res != 1 {
        return nil, fmt.Errorf("%s", pe.String())
    } else {
        return pms, nil
    }
}

// String returns any possible value as a string. These results may be
// truncated depending on the type and size of the underlying value.
func (value *Value) String() (string, error) {
    _type, err := value.Type()
    if err != nil { return "", err }

    switch _type {
    case C.LIBREGF_VALUE_TYPE_STRING:
        tstr, err := value.TString()
        if err != nil { return "", err }
        return tstr, nil
    case C.LIBREGF_VALUE_TYPE_EXPANDABLE_STRING:
        tstr, err := value.TString()
        if err != nil { return "", err }
        return tstr, nil
    case C.LIBREGF_VALUE_TYPE_MULTI_VALUE_STRING:
        ms, err := value.TMultiString()
        if err != nil { return "", err }
        strs, err := ms.Strings()
        if err != nil { return "", err }
        l := len(strs)
        extra := ""
        if l > 4 {
            l = 4
            extra = "…"
        }
        return strings.Join(strs[:l], ", ") + extra, nil
    case C.LIBREGF_VALUE_TYPE_BINARY_DATA:
        tbin, err := value.TBinary()
        if err != nil { return "", err }
        l := len(tbin)
        extra := ""
        if l > 40 {
            l = 40
            extra = "…"
        }
        return fmt.Sprintf("%x", tbin[:l]) + extra, nil
    case C.LIBREGF_VALUE_TYPE_INTEGER_32BIT_LITTLE_ENDIAN:
        i, err := value.Tint32()
        if err != nil { return "", err }
        return fmt.Sprintf("%d", i), nil
    case C.LIBREGF_VALUE_TYPE_INTEGER_64BIT_LITTLE_ENDIAN:
        i, err := value.Tint64()
        if err != nil { return "", err }
        return fmt.Sprintf("%d", i), nil
    default:
        return fmt.Sprintf("[Please implement value type %d]", _type), nil
    }
}

// Free frees the memory allocated by C to an opaque *MultiString.
// It wraps libregf_multi_string_free().
// Most of the time you will just defer a call to Free() right after calling
// a function that allocates a MultiString.
func (ms *MultiString) Free() error { 
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_multi_string_free((**C.libregf_multi_string_t)(unsafe.Pointer(&ms)), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return fmt.Errorf("%s", pe.String())
    } else {
        return nil
    }
}

// StringsLen returns the number of strings contained in a MultiString.
// It wraps libregf_multi_string_get_number_of_strings().
func (ms *MultiString) StringsLen() (int, error) { 
    var slen C.int
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_multi_string_get_number_of_strings((*C.libregf_multi_string_t)(ms), &slen, (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return -1, fmt.Errorf("%s", pe.String())
    } else {
        return int(slen), nil
    }
}

// StringLenAt returns the length (in bytes) of a particular string
// in a MultiString.
// It wraps libregf_multi_string_get_utf8_string_size().
// You don't need to call this function if you call StringAt(), which calls StringLenAt().
func (ms *MultiString) StringLenAt(index int) (int, error) { 
    var slen C.size_t
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_multi_string_get_utf8_string_size((*C.libregf_multi_string_t)(ms), C.int(index), (*C.size_t)(&slen), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return -1, fmt.Errorf("%s", pe.String())
    } else {
        return int(slen), nil
    }
}

// StringAt returns the string at a position inside a MultiString.
// It wraps libregf_multi_string_get_utf8_string().
func (ms *MultiString) StringAt(index int) (string, error) { 
    slen, err := ms.StringLenAt(index)
    if err != nil { return "", err }

    var cerr Error
    ppe := unsafe.Pointer(&cerr)
    buffer := make([]byte, slen+1)
    cstr := C.CString(string(buffer[:slen]))
    defer C.free(unsafe.Pointer(cstr))

    res := int(C.libregf_multi_string_get_utf8_string((*C.libregf_multi_string_t)(ms), C.int(index), (*C.uint8_t)(unsafe.Pointer(cstr)), C.ulong(slen), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return "", fmt.Errorf("%s", pe.String())
    } else {
        return C.GoString(cstr), nil
    }
}

// Strings returns all of the strings inside a MultiString as a []string.
func (ms *MultiString) Strings() ([]string, error) {
    slen, err := ms.StringsLen()
    if err != nil { return []string{}, err }

    strs := make([]string, slen)
    for i := 0; i < slen; i++ {
        s, err := ms.StringAt(i)
        if err != nil { return []string{}, nil }

        strs[i] = s
    }

    return strs, nil
} 
