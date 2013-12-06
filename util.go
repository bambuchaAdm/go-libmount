package libmount;

import "C"
import "fmt"

type LibmountError struct {
	messages string
}

func BuildError(format string, arg ...interface{}) LibmountError {
	return LibmountError{fmt.Sprintf(format,arg)}
}

func convertBool(value bool) C.int {
	if value {
		return C.int(1)
	} 
	return C.int(0)
}
