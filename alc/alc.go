package alc

// #cgo darwin LDFLAGS: -lalut -framework OpenAL 
// #cgo windows LDFLAGS: -lopenal32
// #cgo linux LDFLAGS: -lopenal
// #include "alc.h"
import "C"

import (
	"unsafe"
	"reflect"
)

type ALCenum C.ALCenum

func alcBool(v bool) C.ALCboolean {
	if v {
		return 1
	}
	return 0
}

func goBool(v C.ALCboolean) bool {
	return v != 0
}

func alcString(s string) (r *C.ALCchar) {
	if s == ""{
		r = nil
	} else{
		r = (*C.ALCchar)(C.CString(s))	
	}
	return
}

func goString(s *C.ALCchar) string {
	return C.GoString((*C.char)(s))
}

func freeString(ptr *C.ALCchar) {
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

//---------------------------------------------------------------
type Context C.ALCcontext

func CreateContext(device *Device, attrlist []int32) (context *Context) {
	return (*Context)(C.alcCreateContext(
		(*C.ALCcontext)(context),
		(*C.ALCint)(&attrlist[0]),
	))
}

func (context *Context) MakeCurrent() bool{
	return goBool(C.alcMakeContextCurrent((*C.ALCcontext)(context)))
}

func (context *Context) Precess(){
	C.alcProcessContext((*C.ALCcontext)(context))
}

func (context *Context) Suspend(){
	C.alcSuspendContext((*C.ALCcontext)(context))
}

func (context *Context) Destroy(){
	C.alcDestroyContext((*C.ALCcontext)(context))
}

func GetCurrentContext() *Context{
	return (*Context)(C.alcGetCurrentContext())
}

func (context *Context) GetDevice() *Device{
	return (*Device)(C.alcGetContextsDevice((*C.ALCcontext)(context)))
}

type Device C.ALCdevice

func OpenDevice(name string) *Device {
	var s *C.ALCchar
	s = alcString(name)	
	defer freeString(s)
	return (*Device)(C.alcOpenDevice(s))
}

func (device *Device) CloseDevice() bool {
	return goBool(C.alcCloseDevice((*C.ALCdevice)(device)))
}

/*
 * Error support.
 * Obtain the most recent Context error
 */
func (device *Device) GetError() ALCenum {
	return ALCenum(C.alcGetError((*C.ALCdevice)(device)))
}

/* 
 * Extension support.
 * Query for the presence of an extension, and obtain any appropriate
 * function pointers and enum values.
 */
func (device *Device) IsExtensionPresent(extname string) bool {
	s := alcString(extname)
	b := goBool(C.alcIsExtensionPresent((*C.ALCdevice)(device), s))
	freeString(s)
	return b
}

func (device *Device) GetProcAddress(fname string) uintptr {
	s := alcString(fname)
	p := uintptr(C.alcGetProcAddress((*C.ALCdevice)(device), s))
	freeString(s)
	return p
}

func (device *Device) GetEnumValue(enumname string) ALCenum {
	s := alcString(enumname)
	e := ALCenum(C.alcGetEnumValue((*C.ALCdevice)(device), s))
	freeString(s)
	return e
}

/*
 * Query functions
 */
func (device *Device) GetString(param ALCenum) string {
	return goString(C.alcGetString((*C.ALCdevice)(device), C.ALCenum(param)))
}

func (device *Device) GetIntegerv(param ALCenum, data []int32) {
	C.alcGetIntegerv(
		(*C.ALCdevice)(device),
		C.ALCenum(param),
		(C.ALCsizei)(len(data)),
		(*C.ALCint)(&data[0]),
	)
}

/*
 * Capture functions
 */
func CaptureOpenDevice(devicename string, frequency int, format ALCenum, buffersize int) *Device {
	s := alcString(devicename)
	d := (*Device)(C.alcCaptureOpenDevice(
		s, C.ALCuint(frequency), C.ALCenum(format), C.ALCsizei(buffersize)))
	freeString(s)
	return d
}

func (device *Device) CaptureCloseDevice() bool {
	return goBool(C.alcCaptureCloseDevice((*C.ALCdevice)(device)))
}

func (device *Device) CaptureStart() {
	C.alcCaptureStart((*C.ALCdevice)(device))
}

func (device *Device) CaptureStop() {
	C.alcCaptureStop((*C.ALCdevice)(device))
}

func (device *Device) CaptureSamples(buffer interface{}, samples int) {
	C.alcCaptureSamples((*C.ALCdevice)(device), ptr(buffer), C.ALCsizei(samples))
}
