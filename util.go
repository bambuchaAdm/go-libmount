package libmount;

import "C"
import "fmt"

type LibmountError struct {
	messages string
}

func BuildError(format string, arg ...interface{}) LibmountError {
	return LibmountError{fmt.Sprintf(format,arg)}
}

func convertFromBool(value bool) C.int {
	if value {
		return C.int(1)
	} 
	return C.int(0)
}

func convertToBool(value C.int) bool {
	if value == 0 {
		return false
	}
	return true
}
