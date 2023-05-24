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

// Error is an opaque struct to group related method calls
type Error C.libregf_error_t

// Version returns the version number of the libregf C library
// It wraps libregf_get_version().
func Version() string {
    cVersion := C.libregf_get_version()
    return C.GoString(cVersion)
}

// String returns the string representation of an error.
// It wraps libregf_error_sprint().
func (err *Error) String() string {
    size := C.size_t(200)
    cstr := (*C.char)(C.malloc(size))
    defer C.free(unsafe.Pointer(cstr))
    res := C.libregf_error_sprint((*C.libregf_error_t)(err), cstr, size)

    if res == 0 {
        return "!!! Can't describe error !!!"
    }

    str := C.GoString(cstr)
    return fmt.Sprintf("libregf error (len:%d): %s", len(str), str)
}

// Free frees memory allocated in C for the hidden Error struct.
// It wraps libregf_error_free().
func (err *Error) Free() {
    C.libregf_error_free((**C.libregf_error_t)(unsafe.Pointer(&err)))
}
