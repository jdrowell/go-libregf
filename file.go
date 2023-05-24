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

// File is an opaque struct to group related method calls
type File C.libregf_file_t

// OpenFile opens a registry file by its path.
// It wraps libregf_file_initialize() and libregf_file_open().
func OpenFile(path string) (*File, error) {
    var file File
    ppfile := unsafe.Pointer(&file)
    var err Error
    ppe := unsafe.Pointer(&err)

    res := int(C.libregf_file_initialize((**C.libregf_file_t)(ppfile), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()
    
    if res != 1 {
        return nil, fmt.Errorf("%s", pe.String())
    }

    // pfile is just *ppfile
    //pfile := (*C.libregf_file_t)(*(**C.libregf_file_t)(ppfile))
    pfile := *(**File)(ppfile)

    res = int(C.libregf_file_open((*C.libregf_file_t)(pfile), C.CString(path), C.LIBREGF_ACCESS_FLAG_READ | C.LIBREGF_FILE_TYPE_REGISTRY, (**C.libregf_error_t)(ppe)))
    pe = *(**Error)(ppe)
    defer pe.Free()

    if res != 1 {
        return nil, fmt.Errorf("%s", pe.String())
    } else {
        return (*File)(pfile), nil
    }
}

// Close closes a registry file.
// It wraps libregf_file_close().
func (file *File) Close() {
    C.libregf_file_close((*C.libregf_file_t)(file), nil)
}

// RootKey returns the root Key of a registry file.
// It wraps libregf_file_get_root_key().
func (file *File) RootKey() (*Key, error) { 
    var key Key
    ppkey := unsafe.Pointer(&key)
    var err Error
    ppe := unsafe.Pointer(&err)

    res := int(C.libregf_file_get_root_key((*C.libregf_file_t)(file), (**C.libregf_key_t)(ppkey), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()
    // pkey is just *ppkey
    //pkey := (*C.libregf_key_t)(*(**C.libregf_key_t)(ppkey))
    pkey := *(**Key)(ppkey)

    if res != 1 {
        return nil, fmt.Errorf("%s", pe.String())
    } else {
        return (*Key)(pkey), nil
    }
}

// Key returns a Key by its path inside the registry.
// It wraps libregf_file_get_key_by_utf8_path().
func (file *File) Key(path string) (*Key, error) { 
    var key Key
    ppkey := unsafe.Pointer(&key)
    var cerr Error
    ppe := unsafe.Pointer(&cerr)
    bpath := append([]byte(path), 0)

    res := int(C.libregf_file_get_key_by_utf8_path((*C.libregf_file_t)(file), (*C.uint8_t)(unsafe.Pointer(&bpath[0])), C.ulong(len(bpath)-1), (**C.libregf_key_t)(ppkey), (**C.libregf_error_t)(ppe)))
    pe := *(**Error)(ppe)
    defer pe.Free()
    pkey := *(**Key)(ppkey)

    if res != 1 {
        return nil, fmt.Errorf("%s", pe.String())
    } else {
        return (*Key)(pkey), nil
    }
}

// Value returns the string representation of a "value of a Value" by its path inside the registry.
// It does so by considering that the last part of the path if the Value's name.
// This is a quick way to get a displayable value for a full registry path in one call.
func (file *File) Value(path string) (string, error) { 
    parts := strings.Split(path, "\\")
    l := len(parts)
    k := strings.Join(parts[:l-1], "\\")
    v := parts[l-1]
    key, err := file.Key(k)
    if err != nil { return "", err }
    value, err := key.Value(v) 
    if err != nil { return "", err }
    s, err := value.String()
    if err != nil { return "", err }

    return s, nil
}

