package nvenc
/*
#include <nvEncodeAPI.h>
#cgo CFLAGS: -I/home/pihpah/repos/github/nvidia-codec-sdk/Interface
#cgo LDFLAGS: -lnvidia-encode
 */
import "C"
import (
//	"unsafe"
	"fmt"
)

type apiptr *C.NV_ENCODE_API_FUNCTION_LIST

type Encoder struct {
	api apiptr
}

func NewEncoder() *Encoder {
	api := &C.NV_ENCODE_API_FUNCTION_LIST{}
	status := C.NvEncodeAPICreateInstance(apiptr(api))
	if ! success(status) {
		panic(fmt.Sprintf("nvenc: %s", getStatusText(status)))
	}
	encoder := &Encoder{}
	encoder.api = api
	return encoder
}
