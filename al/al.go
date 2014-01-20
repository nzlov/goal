package al

// #cgo darwin LDFLAGS: -lalut -framework OpenAL 
// #cgo windows LDFLAGS: -lopenal32
// #cgo linux LDFLAGS: -lopenal
// #include "al.h"
import "C"
import (
	"unsafe"
	"reflect"
)


type ALenum C.ALenum

func alBool(v bool) C.ALboolean {
	if v {
		return 1
	}
	return 0
}

func goBool(v C.ALboolean) bool {
	return v != 0
}

func alString(s string) (r *C.ALchar) {
	if s == ""{
		r = nil
	} else{
		r = (*C.ALchar)(C.CString(s))
	}
	return
}

func goString(s *C.ALchar) string {
	return C.GoString((*C.char)(s))
}

func freeString(ptr *C.ALchar) {
	C.free(unsafe.Pointer(ptr))
}

func ptr(v interface{}) unsafe.Pointer {

	if v == nil {
		return unsafe.Pointer(nil)
	}

	rv := reflect.ValueOf(v)
	var et reflect.Value
	switch rv.Type().Kind() {
	case reflect.Uintptr:
		offset, _ := v.(uintptr)
		return unsafe.Pointer(offset)
	case reflect.Ptr:
		et = rv.Elem()
	case reflect.Slice:
		et = rv.Index(0)
	default:
		panic("type must be a pointer, a slice, uintptr or nil")
	}

	return unsafe.Pointer(et.UnsafeAddr())
}

//---------------------------------------------------------------------

/*
 * Renderer State management
 */
func Enable(capability ALenum) {
	C.alEnable(C.ALenum(capability))
}

func Disable(capability ALenum) {
	C.alDisable(C.ALenum(capability))
}

func IsEnabled(capability ALenum) bool {
	return goBool(C.alIsEnabled(C.ALenum(capability)))
}

/*
 * State retrieval
 */
func GetString(param ALenum) string {
	return goString(C.alGetString(C.ALenum(param)))
}

func GetBooleanv(param ALenum, data []bool) {
	C.alGetBooleanv(C.ALenum(param), (*C.ALboolean)(unsafe.Pointer(&data[0])))
}

func GetIntegerv(param ALenum, data []int32) {
	C.alGetIntegerv(C.ALenum(param), (*C.ALint)(&data[0]))
}

func GetFloatv(param ALenum, data []float32) {
	C.alGetFloatv(C.ALenum(param), (*C.ALfloat)(&data[0]))
}

func GetDoublev(param ALenum, data []float64) {
	C.alGetDoublev(C.ALenum(param), (*C.ALdouble)(&data[0]))
}

func GetBoolean(param ALenum) bool {
	return goBool(C.alGetBoolean(C.ALenum(param)))
}

func GetInteger(param ALenum) int {
	return int(C.alGetInteger(C.ALenum(param)))
}

func GetFloat(param ALenum) float32 {
	return float32(C.alGetFloat(C.ALenum(param)))
}

func GetDouble(param ALenum) float64 {
	return float64(C.alGetDouble(C.ALenum(param)))
}

/*
 * Error support.
 * Obtain the most recent error generated in the AL state machine.
 */
func GetError() ALenum {
	return ALenum(C.alGetError())
}

/* 
 * Extension support.
 * Query for the presence of an extension, and obtain any appropriate
 * function pointers and enum values.
 */
func IsExtensionPresent(extname string) bool {
	s := alString(extname)
	b := goBool(C.alIsExtensionPresent(s))
	C.free(unsafe.Pointer(s))
	return b
}

func GetProcAddress(fname string) uintptr {
	s := alString(fname)
	p := uintptr(C.alGetProcAddress(s))
	freeString(s)
	return p
}

func GetEnumValue(ename string) ALenum {
	s := alString(ename)
	e := ALenum(C.alGetEnumValue(s))
	freeString(s)
	return e
}

/*
 * Global Parameters
 */
func DopplerFactor(value float32) {
	C.alDopplerFactor(C.ALfloat(value))
}

func DopplerVelocity(value float32) {
	C.alDopplerVelocity(C.ALfloat(value))
}

func SpeedOfSound(value float32) {
	C.alSpeedOfSound(C.ALfloat(value))
}

func DistanceModel(distanceModel ALenum) {
	C.alDistanceModel(C.ALenum(distanceModel))
}
