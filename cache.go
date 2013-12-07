package libmount

//#cgo pkg-config: mount
//#include <libmount/libmount.h>
//#include <stdlib.h>
import "C"

import (
	"unsafe"
)

type Cache struct {
	handler * C.struct_libmnt_cache
}

func createCache() (* Cache) {
	return &Cache{ C.mnt_new_cache() }
}

func (cache * Cache) free() {
	C.mnt_free_cache(cache.handler)
}

func (cache * Cache) DeviceHasTag(device string, tag string, value string) bool {
	dev := C.CString(device)
	defer C.free(unsafe.Pointer(dev))
	ctag := C.CString(tag)
	defer C.free(unsafe.Pointer(ctag))
	val := C.CString(value)
	defer C.free(unsafe.Pointer(val))
	result := C.mnt_cache_device_has_tag(cache.handler,dev,ctag,val)
	if result == 1 {
		return true
	} else {
		return false
	}
}

func (cache * Cache) FindTagValue(device string, token string) (string,error) {
	dev := C.CString(device)
	defer C.free(unsafe.Pointer(dev))
	ctoken := C.CString(token)
	defer C.free(unsafe.Pointer(ctoken))
	result := C.mnt_cache_find_tag_value(cache.handler,dev,ctoken)
	defer C.free(unsafe.Pointer(result))
	if result != nil {
		return C.GoString(result), nil;
	} else {
		return "", BuildError("Error on finding tag in cache")
	}
}


func (cache * Cache) ReadTags(device string) bool {
	dev := C.CString(device)
	defer C.free(unsafe.Pointer(dev))
	result := C.mnt_cache_read_tags(cache.handler,dev)
	if result == 1 {
		return true
	} else {
		return false
	}
}

func (cache * Cache) GetFstype(device string) (string,bool,error) {
	var amb * C.int 
	dev := C.CString(device)
	defer C.free(unsafe.Pointer(dev))
	result := C.mnt_get_fstype(dev,amb,cache.handler)
	ambivalent := convertToBool(*amb)
	if result != nil {
		return C.GoString(result), ambivalent, nil
	} else {
		return "", ambivalent, BuildError("Error on geting fstype.")
	}
}

func (cache * Cache) PrettyPath(device string) (string,error) {
	dev := C.CString(device)
	defer C.free(unsafe.Pointer(dev))
	result := C.mnt_pretty_path(dev,cache.handler)
	defer C.free(unsafe.Pointer(result))
	if result != nil {
		return C.GoString(result), nil
	} else {
		return "", BuildError("Error on getting preety path")
	}
}

func (cache * Cache) ResolvePath(device string) (string, error) {
	dev := C.CString(device)
	defer C.free(unsafe.Pointer(dev))
	result := C.mnt_resolve_path(dev,cache.handler)
	if result != nil {
		return C.GoString(result), nil
	} else {
		return "", BuildError("Error on resolving path")
	}
}

func (cache * Cache) ResolveSpec(device string) (string, error) {
	dev := C.CString(device)
	defer C.free(unsafe.Pointer(dev))
	result := C.mnt_resolve_spec(dev,cache.handler)
	if result != nil {
		return C.GoString(result), nil
	} else {
		return "", BuildError("Error on resolving path")
	}
}

func (cache * Cache) ResolveTag(token string,value string) (string,error) {
	tok := C.CString(token)
	defer C.free(unsafe.Pointer(tok))
	val := C.CString(value)
	defer C.free(unsafe.Pointer(val))
	result := C.mnt_resolve_tag(tok,val,cache.handler)
	if result != nil {
		return C.GoString(result), nil
	} else {
		return "", BuildError("Error on resolving tags")
	}
}
