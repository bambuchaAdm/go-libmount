package context;

//#cgo pkg-config: mount
//#include <libmount/libmount.h>
import "C"

type Context struct {
	 handler * C.struct_libmnt_context
}

func createContext() (* Context) {
	return &Context{C.mnt_new_context()}
}

func (context * Context) free() {
	C.mnt_free_context(context.handler)
}
