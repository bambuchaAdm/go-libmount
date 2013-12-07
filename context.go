package libmount;

//#cgo pkg-config: mount
//#include <libmount/libmount.h>
//#include <stdlib.h>
import "C"

import (
	"unsafe"
)


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
	arg := convertFromBool(option)
	result := int(C.mnt_context_disable_canonicalize(context.handler,arg))
	if result == 0 {
		return nil;
	} else {
		return BuildError("Error on disables canonicalzied form. Exit code = %d",result)
	}
}

func (context * Context) DisableHelpers(option bool) (err error) {
	arg := convertFromBool(option)
	result := int(C.mnt_context_disable_helpers(context.handler,arg))
	if result == 0 {
		return nil;
	} else {
		return BuildError("Error on disables helpers. Exit code = %d",result)
	}
}

func (context * Context) DisableMtab(option bool) (err error) {
	arg := convertFromBool(option)
	result := int(C.mnt_context_disable_mtab(context.handler,arg))
	if result == 0 {
		return nil;
	} else {
		return BuildError("Error on disables mtab. Exit code = %d",result)
	}
}

func (context * Context) EnableFake(option bool) (err error) {
	arg := convertFromBool(option)
	result := int(C.mnt_context_enable_fake(context.handler,arg))
	if result == 0 {
		return nil;
	} else {
		return BuildError("Error on enabling fake. Exit code = %d",result)
	}
}

func (context * Context) EnableForce(option bool) (err error) {
	arg := convertFromBool(option)
	result := int(C.mnt_context_enable_force(context.handler,arg))
	if result == 0 {
		return nil;
	} else {
		return BuildError("Error on enabling force. Exit code = %d",result)
	}
}

func (context * Context) EnableFork(option bool) (err error) {
	arg := convertFromBool(option)
	result := int(C.mnt_context_enable_fork(context.handler,arg))
	if result == 0 {
		return nil;
	} else {
		return BuildError("Error on enabling forking. Exit code = %d",result)
	}
}

func (context * Context) EnableLazy(option bool) (err error) {
	arg := convertFromBool(option)
	result := int(C.mnt_context_enable_lazy(context.handler,arg))
	if result == 0 {
		return nil;
	} else {
		return BuildError("Error on enabling lazy mounting. Exit code = %d",result)
	}
}

func (context * Context) EnableLoopdel(option bool) (err error) {
	arg := convertFromBool(option)
	result := int(C.mnt_context_enable_loopdel(context.handler,arg))
	if result == 0 {
		return nil;
	} else {
		return BuildError("Error on enabling loop deletion. Exit code = %d",result)
	}
}

func (context * Context) EnableRdonlyUmount(option bool) (err error) {
	arg := convertFromBool(option)
	result := int(C.mnt_context_enable_rdonly_umount(context.handler,arg))
	if result == 0 {
		return nil;
	} else {
		return BuildError("Error on enabling read only remount. Exit code = %d",result)
	}
}

func (context * Context) EnableSloppy(option bool) (err error) {
	arg := convertFromBool(option)
	result := int(C.mnt_context_enable_sloppy(context.handler,arg))
	if result == 0 {
		return nil;
	} else {
		return BuildError("Error on enabling sloppy mounting. Exit code = %d",result)
	}
}

func (context * Context) EnableVerbose(option bool) (err error) {
	arg := convertFromBool(option)
	result := int(C.mnt_context_enable_verbose(context.handler,arg))
	if result == 0 {
		return nil;
	} else {
		return BuildError("Error on enabling sloppy mounting. Exit code = %d",result)
	}
}





