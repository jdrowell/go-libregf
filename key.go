package libregf

/*
#cgo LDFLAGS: -lregf
#include <libregf.h>
*/
import "C"

import (
    "fmt"
    "unsafe"
)

// Key is an opaque struct to group related method calls
type Key   C.libregf_key_t

// Free frees memory allocated in C for the hidden Key struct.
// It wraps libregf_key_free().
func (key *Key) Free() error { 
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_key_free((**C.libregf_key_t)(unsafe.Pointer(&key)), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return fmt.Errorf("%s", pe.String())
    } else {
        return nil
    }
}

// NameLen returns the length (in bytes) of the Key's name.
// It wraps libregf_key_get_utf8_name_size().
// You usually call it to know how much space to allocate before calling
// libregf_key_get_utf8_name().
// You don't need to call this function if you call Name(), which calls NameLen().
func (key *Key) NameLen() (int, error) { 
    var namelen C.size_t
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_key_get_utf8_name_size((*C.libregf_key_t)(key), (*C.size_t)(&namelen), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return -1, fmt.Errorf("%s", pe.String())
    } else {
        return int(namelen), nil
    }
}

// Name returns the Key's name.
// It wraps libregf_key_get_utf8_name().
// It returns a regular Go string, so you don't have to worry about the
// underlying C calls and memory allocations.
func (key *Key) Name() (string, error) { 
    namelen, err := key.NameLen()
    if err != nil { return "", err }

    buffer := make([]byte, namelen+1)
    cstr := C.CString(string(buffer[:namelen]))
    defer C.free(unsafe.Pointer(cstr))
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_key_get_utf8_name((*C.libregf_key_t)(key), (*C.uint8_t)(unsafe.Pointer(cstr)), C.ulong(namelen), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return "", fmt.Errorf("%s", pe.String())
    } else {
        return C.GoString(cstr), nil
    }
}

// ClassNameLen returns the length (in bytes) of the Key's class name.
// It wraps libregf_key_get_utf8_class_name_size().
// You usually call it to know how much space to allocate before calling
// libregf_key_get_utf8_class_name().
// You don't need to call this function if you call ClassName(), which calls ClassNameLen().
func (key *Key) ClassNameLen() (int, error) { 
    var namelen C.size_t
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_key_get_utf8_class_name_size((*C.libregf_key_t)(key), (*C.size_t)(&namelen), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return -1, fmt.Errorf("%s", pe.String())
    } else {
        return int(namelen), nil
    }
}

// ClassName returns the Key's class name.
// It wraps libregf_key_get_utf8_class_name().
// It returns a regular Go string, so you don't have to worry about the
// underlying C calls and memory allocations.
func (key *Key) ClassName() (string, error) { 
    namelen, err := key.ClassNameLen()
    if err != nil { return "", err }

    buffer := make([]byte, namelen+1)
    cstr := C.CString(string(buffer[:namelen]))
    defer C.free(unsafe.Pointer(cstr))
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_key_get_utf8_class_name((*C.libregf_key_t)(key), (*C.uint8_t)(unsafe.Pointer(cstr)), C.ulong(namelen), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return "", fmt.Errorf("%s", pe.String())
    } else {
        return C.GoString(cstr), nil
    }
}

// ValuesLen returns the number of Values present inside a Key.
// It wraps libregf_key_get_number_of_values().
func (key *Key) ValuesLen() (int, error) { 
    var num C.int
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_key_get_number_of_values((*C.libregf_key_t)(key), &num, (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return -1, fmt.Errorf("%s", pe.String())
    } else {
        return int(num), nil
    }
}

// ValueAt returns the Value at a given position inside a Key.
// It wraps libregf_key_get_value().
func (key *Key) ValueAt(index int) (*Value, error) { 
    var cerr Error
    ppe := unsafe.Pointer(&cerr)
    var value Value
    ppvalue := unsafe.Pointer(&value)

    res := int(C.libregf_key_get_value((*C.libregf_key_t)(key), C.int(index), (**C.libregf_value_t)(ppvalue), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()
    pvalue := *(**Key)(ppvalue)

    if res != 1 {
        return nil, fmt.Errorf("%s", pe.String())
    } else {
        return (*Value)(pvalue), nil
    }
}

// Value returns the Value present inside a Key by its name.
// It wraps libregf_key_get_value_by_utf8_name().
func (key *Key) Value(path string) (*Value, error) { 
    var value Value
    ppvalue := unsafe.Pointer(&value)
    var cerr Error
    ppe := unsafe.Pointer(&cerr)
    bpath := append([]byte(path), 0)

    res := int(C.libregf_key_get_value_by_utf8_name((*C.libregf_key_t)(key), (*C.uint8_t)(unsafe.Pointer(&bpath[0])), C.ulong(len(bpath)-1), (**C.libregf_value_t)(ppvalue), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()
    pvalue := *(**Value)(ppvalue)

    if res != 1 {
        return nil, fmt.Errorf("%s", pe.String())
    } else {
        return (*Value)(pvalue), nil
    }
}

// SubkeysLen returns the count of sub-Keys present inside a Key.
// It wraps libregf_key_get_number_of_sub_keys().
func (key *Key) SubkeysLen() (int, error) { 
    var num C.int
    var cerr Error
    ppe := unsafe.Pointer(&cerr)

    res := int(C.libregf_key_get_number_of_sub_keys((*C.libregf_key_t)(key), &num, (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return -1, fmt.Errorf("%s", pe.String())
    } else {
        return int(num), nil
    }
}

// SubkeyAt returns the Key at a given position inside another Key.
// It wraps libregf_key_get_sub_key().
func (key *Key) SubkeyAt(index int) (*Key, error) { 
    var cerr Error
    ppe := unsafe.Pointer(&cerr)
    var subkey Key
    ppsubkey := unsafe.Pointer(&subkey)

    res := int(C.libregf_key_get_sub_key((*C.libregf_key_t)(key), C.int(index), (**C.libregf_key_t)(ppsubkey), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()
    psubkey := *(**Key)(ppsubkey)

    if res != 1 {
        return nil, fmt.Errorf("%s", pe.String())
    } else {
        return (*Key)(psubkey), nil
    }
}

// SubkeyByName returns the Key present inside another Key by its name.
// It wraps libregf_key_get_sub_key_by_utf8_name().
func (key *Key) SubkeyByName(name string) (*Key, error) { 
    var cerr Error
    ppe := unsafe.Pointer(&cerr)
    var subkey Key
    ppsubkey := unsafe.Pointer(&subkey)
    bname := []byte(name)
    //bname = append(bname, 0)
    fmt.Printf("bname: %v\n", bname)

    res := int(C.libregf_key_get_sub_key_by_utf8_name((*C.libregf_key_t)(key), (*C.uint8_t)(unsafe.Pointer(&bname[0])), C.ulong(len(bname)), (**C.libregf_key_t)(ppsubkey), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()
    psubkey := *(**Key)(ppsubkey)

    if res != 1 {
        return nil, fmt.Errorf("%s", pe.String())
    } else {
        return (*Key)(psubkey), nil
    }
}

