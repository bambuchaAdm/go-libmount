package context;

//#cgo pkg-config: mount
//#include <libmount/libmount.h>
//#include <stdlib.h>
import "C"

import "unsafe"

type LibmountError struct {
	messages string
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
	cstring := C.CString(options)
	defer C.free(unsafe.Pointer(cstring))
	result := C.mnt_context_append_options(context.handler,cstring)
	if result == 0 {
		return nil;
	} else {
		return LibmountError{"Error on appending options to context"}
	}
}

