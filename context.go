package libmount;

//#cgo pkg-config: mount
//#include <libmount/libmount.h>
//#include <stdlib.h>
import "C"

import (
	"unsafe"
	"fmt"
)

type LibmountError struct {
	messages string
}

func BuildError(format string, arg ...interface{}) LibmountError {
	return LibmountError{fmt.Sprintf(format,arg)}
}

func (input LibmountError) Error() string {
	return input.messages
}

type Context struct {
	 handler * C.struct_libmnt_context
}

func createContext() (* Context) {
	return &Context{C.mnt_new_context()}
}

func (context * Context) free() {
	C.mnt_free_context(context.handler)
}

func (context * Context) AppendOptions(options string) (err error) {
	arg := C.CString(options)
	defer C.free(unsafe.Pointer(arg))
	result := int(C.mnt_context_append_options(context.handler,arg))
	if result == 0 {
		return nil;
	} else {
		return BuildError("Error on appending options to context. Exit code = %d",result)
	}
}

func (context * Context) DisableCanonicalize(option bool) (err error) {
	arg := convertBool(option)
	result := int(C.mnt_context_disable_canonicalize(context.handler,arg))
	if result == 0 {
		return nil;
	} else {
		return BuildError("Error on disables canonicalzied form. Exit code = %d",result)
	}
}

func (context * Context) DisableHelpers(option bool) (err error) {
	arg := convertBool(option)
	result := int(C.mnt_context_disable_helpers(context.handler,arg))
	if result == 0 {
		return nil;
	} else {
		return BuildError("Error on disables helpers. Exit code = %d",result)
	}
}

func (context * Context) DisableMtab(option bool) (err error) {
	arg := convertBool(option)
	result := int(C.mnt_context_disable_mtab(context.handler,arg))
	if result == 0 {
		return nil;
	} else {
		return BuildError("Error on disables mtab. Exit code = %d",result)
	}
}

func (context * Context) EnableFake(option bool) (err error) {
	arg := convertBool(option)
	result := int(C.mnt_context_enable_fake(context.handler,arg))
	if result == 0 {
		return nil;
	} else {
		return BuildError("Error on enabling fake. Exit code = %d",result)
	}
}

