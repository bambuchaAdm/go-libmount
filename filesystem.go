package libmount;

//#cgo pkg-config: mount
//#include <libmount/libmount.h>
//#include <stdlib.h>
import "C"

import "unsafe"

type Filesystem struct {
	handler * C.struct_libmnt_fs
}

func (fs * Filesystem) AppendAttributes(attributes string) error {
	attrstr := C.CString(attributes)
	defer C.free(unsafe.Pointer(attrstr))
	result := C.mnt_fs_append_attributes(fs.handler, attrstr)
	if result == 0 {
		return nil
	}
	return BuildError("Error on appending attribute to fs object. Exit code = %d", int(result))
}

func (fs * Filesystem) AppendOptions(options string) error {
	optstr := C.CString(options)
	defer C.free(unsafe.Pointer(optstr))
	result := C.mnt_fs_append_options(fs.handler, optstr)
	if result == 0 {
		return nil
	}
	return BuildError("Error on appending option to fs object. Exit code = %d", int(result))
} 

type MissingAttribute struct {
	name string
}

func (err MissingAttribute) Error() string {
	return "Missing attribute with name = " + err.name
}

func (fs * Filesystem) GetAttribute(name string) (string,error) {
	attr := C.CString(name)
	defer C.free(unsafe.Pointer(attr))
	var buffer * C.char
	var length C.size_t
	result := C.mnt_fs_get_attribute(fs.handler,attr,&buffer,&length)
	value := C.GoStringN(buffer,C.int(length))
	switch {
	case result == 0 :
		return value, nil
	case result == 1:
		return "", MissingAttribute{name}
	}
	return "", BuildError("Error on getting attrigute form fs object. Exit code = %d", int(result))
}

func (fs * Filesystem) GetAttributes() string {
	result := C.mnt_fs_get_attributes(fs.handler)
	return C.GoString(result)
}

func (fs * Filesystem) GetBindSource() string {
	result := C.mnt_fs_get_bindsrc(fs.handler)
	return C.GoString(result)
}

func (fs * Filesystem) GetDevice() Device {
	result := C.mnt_fs_get_devno(fs.handler)
	return Device{uint(result)}
}

func (fs * Filesystem) GetFrequency() int {
	result := C.mnt_fs_get_freq(fs.handler)
	return int(result)
}

func (fs * Filesystem) GetPosibleOptions() string {
	result := C.mnt_fs_get_fs_options(fs.handler)
	return C.GoString(result)
}

func (fs * Filesystem) GetType() string {
	result := C.mnt_fs_get_fstype(fs.handler)
	return C.GoString(result)
}

func (fs * Filesystem) GetId() (int,error) {
	result := C.mnt_fs_get_id(fs.handler)
	if result > 0 {
		return int(result), nil
	} 
	return 0, BuildError("Error on getting filesystem id. Exit code = %d", result)
}

type MissingOption struct {
	name string
}

func (err MissingOption) Error() string {
	return "Missing option with name = " + err.name
}

func (fs * Filesystem) GetOption(name string) (string,error) {
	opt := C.CString(name)
	defer C.free(unsafe.Pointer(opt))
	var buffer * C.char;
	var length C.size_t
	result := C.mnt_fs_get_option(fs.handler,opt,&buffer,&length)
	switch {
	case result == 0:
		return C.GoStringN(buffer,C.int(length)), nil
	case result == 0:
		return "", MissingOption{name}
	}
	return "",BuildError("Error on getting option form fs object. Exit code = %d", int(result))
}

func (fs * Filesystem) GetOptions() (string,error) {
	result := C.mnt_fs_get_options(fs.handler)
	if result != nil {
		return C.GoString(result), nil
	}
	return "", BuildError("Error on getting options form fs object.")
}

func (fs * Filesystem) GetParentId() (int,error) {
	result := C.mnt_fs_get_parent_id(fs.handler)
	if result > 0 {
		return int(result), nil
	} 
	return 0, BuildError("Error on getting parent id. Exit code = %d", int(result))
}

func (fs * Filesystem) GetPassNumber() int {
	result := C.mnt_fs_get_passno(fs.handler)
	return int(result)
}

type MissingRoot struct {
}

func (err MissingRoot) Error() string {
	return "Missing root of the mount"
}

func (fs * Filesystem) GetRoot() (string,error) {
	result := C.mnt_fs_get_root(fs.handler)
	if result != nil {
		return C.GoString(result), nil
	}
	return "", MissingRoot{}
}

type MissingSource struct {
}

func (err MissingSource) Error() string {
	return "Missing source. Error or didn't specified."
}

func (fs * Filesystem) GetSource() (string, error) {
	result := C.mnt_fs_get_source(fs.handler)
	return C.GoString(result), nil
}

func (fs * Filesystem) GetSourcePath() (string, error) {
	result := C.mnt_fs_get_srcpath(fs.handler)
	if result != nil {
		return C.GoString(result), nil
	}
	return "", BuildError("Error od path not defined")
}

func (fs * Filesystem) GetTag() (string,string, error) {
	var name * C.char
	var value * C.char
	if C.mnt_fs_get_tag(fs.handler,&name,&value) != 0 {
		return "", "", BuildError("Error on getting tag.")
	}
	return C.GoString(name), C.GoString(value), nil
		
}
