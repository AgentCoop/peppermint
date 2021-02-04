package nvenc
/*
#include <nvEncodeAPI.h>
#cgo CFLAGS: -I/home/pihpah/repos/github/nvidia-codec-sdk/Interface
*/
import "C"
import "C"

type apiCallStatus C.NVENCSTATUS

var apiCallStatusMap = map[C.NVENCSTATUS]string{
	/**
	 * This indicates that API call returned with no errors.
	 */
	C.NV_ENC_SUCCESS: "success",
	/**
	 * This indicates that no encode capable devices were detected.
	 */
	C.NV_ENC_ERR_NO_ENCODE_DEVICE: "no encode device",
	/**
	 * This indicates that devices pass by the client is not supported.
	 */
	C.NV_ENC_ERR_UNSUPPORTED_DEVICE: "unsupported device",
	/**
	 * This indicates that the encoder device supplied by the client is not
	 * valid.
	 */
	C.NV_ENC_ERR_INVALID_ENCODERDEVICE: "invalid encoder device",
	/**
	 * This indicates that device passed to the API call is invalid.
	 */
	C.NV_ENC_ERR_INVALID_DEVICE: "invalid device",
	/**
	 * This indicates that device passed to the API call is no longer available and
	 * needs to be reinitialized. The clients need to destroy the current encoder
	 * session by freeing the allocated input output buffers and destroying the device
	 * and create a new encoding session.
	 */
	C.NV_ENC_ERR_DEVICE_NOT_EXIST: "device not exist",
	/**
	 * This indicates that one or more of the pointers passed to the API call
	 * is invalid.
	 */
	C.NV_ENC_ERR_INVALID_PTR: "invalid API call pointer",
	/**
	 * This indicates that completion event passed in ::NvEncEncodePicture() call
	 * is invalid.
	 */
	C.NV_ENC_ERR_INVALID_EVENT: "invalid event",
	/**
	 * This indicates that one or more of the parameter passed to the API call
	 * is invalid.
	 */
	C.NV_ENC_ERR_INVALID_PARAM: "invalid param",
	/**
	 * This indicates that an API call was made in wrong sequence/order.
	 */
	C.NV_ENC_ERR_INVALID_CALL: "invalid API call",
	/**
	 * This indicates that the API call failed because it was unable to allocate
	 * enough memory to perform the requested operation.
	 */
	C.NV_ENC_ERR_OUT_OF_MEMORY: "out of memory",
	/**
	 * This indicates that the encoder has not been initialized with
	 * ::NvEncInitializeEncoder() or that initialization has failed.
	 * The client cannot allocate input or output buffers or do any encoding
	 * related operation before successfully initializing the encoder.
	 */
	C.NV_ENC_ERR_ENCODER_NOT_INITIALIZED: "encoder not initialized",
	/**
	 * This indicates that an unsupported parameter was passed by the client.
	 */
	C.NV_ENC_ERR_UNSUPPORTED_PARAM: "unsupported param",
	/**
	 * This indicates that the ::NvEncLockBitstream() failed to lock the output
	 * buffer. This happens when the client makes a non blocking lock call to
	 * access the output bitstream by passing NV_ENC_LOCK_BITSTREAM::doNotWait flag.
	 * This is not a fatal error and client should retry the same operation after
	 * few milliseconds.
	 */
	C.NV_ENC_ERR_LOCK_BUSY: "failed to lock the output buffer",
	/**
	 * This indicates that the size of the user buffer passed by the client is
	 * insufficient for the requested operation.
	 */
	C.NV_ENC_ERR_NOT_ENOUGH_BUFFER: "not enough buffer",
	/**
	 * This indicates that an invalid struct version was used by the client.
	 */
	C.NV_ENC_ERR_INVALID_VERSION: "client invalid version",
	/**
	 * This indicates that ::NvEncMapInputResource() API failed to map the client
	 * provided input resource.
	 */
	C.NV_ENC_ERR_MAP_FAILED: "::NvEncMapInputResource() failed",
	/**
	 * This indicates encode driver requires more input buffers to produce an output
	 * bitstream. If this error is returned from ::NvEncEncodePicture() API, this
	 * is not a fatal error. If the client is encoding with B frames then,
	 * ::NvEncEncodePicture() API might be buffering the input frame for re-ordering.
	 *
	 * A client operating in synchronous mode cannot call ::NvEncLockBitstream()
	 * API on the output bitstream buffer if ::NvEncEncodePicture() returned the
	 * ::NV_ENC_ERR_NEED_MORE_INPUT error code.
	 * The client must continue providing input frames until encode driver returns
	 * ::NV_ENC_SUCCESS. After receiving ::NV_ENC_SUCCESS status the client can call
	 * ::NvEncLockBitstream() API on the output buffers in the same order in which
	 * it has called ::NvEncEncodePicture().
	 */
	C.NV_ENC_ERR_NEED_MORE_INPUT: "need more input",
	/**
	 * This indicates that the HW encoder is busy encoding and is unable to encode
	 * the input. The client should call ::NvEncEncodePicture() again after few
	 * milliseconds.
	 */
	C.NV_ENC_ERR_ENCODER_BUSY: "encoder busy",

	/**
	 * This indicates that the completion event passed in ::NvEncEncodePicture()
	 * API has not been registered with encoder driver using ::NvEncRegisterAsyncEvent().
	 */
	C.NV_ENC_ERR_EVENT_NOT_REGISTERD: "event not registered",
	/**
	 * This indicates that an unknown internal error has occurred.
	 */
	C.NV_ENC_ERR_GENERIC: "unknown internal error",
	/**
	 * This indicates that the client is attempting to use a feature
	 * that is not available for the license type for the current system.
	 */
	C.NV_ENC_ERR_INCOMPATIBLE_CLIENT_KEY: "incompatible client key",

	/**
	 * This indicates that the client is attempting to use a feature
	 * that is not implemented for the current version.
	 */
	C.NV_ENC_ERR_UNIMPLEMENTED: "unimplemented",
	/**
	 * This indicates that the ::NvEncRegisterResource API failed to register the resource.
	 */
	C.NV_ENC_ERR_RESOURCE_REGISTER_FAILED: ":NvEncRegisterResource() failed",
	/**
	 * This indicates that the client is attempting to unregister a resource
	 * that has not been successfully registered.
	 */
	C.NV_ENC_ERR_RESOURCE_NOT_REGISTERED: "resource not registered",
	/**
	 * This indicates that the client is attempting to unmap a resource
	 * that has not been successfully mapped.
	 */
	C.NV_ENC_ERR_RESOURCE_NOT_MAPPED: "resource not mapped",
}

func success(status C.NVENCSTATUS) bool {
	return status == C.NV_ENC_SUCCESS
}

func getStatusText(status C.NVENCSTATUS) string {
	return apiCallStatusMap[status]
}
