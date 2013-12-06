package libmount;

import "C"

func convertBool(value bool) C.int {
	if value {
		return C.int(1)
	} 
	return C.int(0)
}
