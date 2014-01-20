package alut

// #cgo  darwin  CFLAGS: -I/usr/local/include
// #cgo  darwin LDFLAGS: -L/usr/local/lib -lalut -framework OpenAL
// #cgo windows LDFLAGS: -lalut32 -lopenal32
// #cgo linux LDFLAGS: -lalut -lopenal
// #include <stdlib.h>
// #include <AL/alut.h>
import "C"

import (
	"github.com/nzlov/goal/al"
	"os"
	"reflect"
	"unsafe"
)

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
	if s == "" {
		r = nil
	} else {
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

//---------------------------------------------------------------

func Init() {
	argc := C.int(len(os.Args))
	argv := make([]*C.char, argc)
	for i, arg := range os.Args {
		argv[i] = C.CString(arg)
	}

	C.alutInit(&argc, &argv[0])

	for _, arg := range argv {
		C.free(unsafe.Pointer(arg))
	}
}

func InitWithoutContext() {
	argc := C.int(len(os.Args))
	argv := make([]*C.char, argc)
	for i, arg := range os.Args {
		argv[i] = C.CString(arg)
	}

	C.alutInitWithoutContext(&argc, &argv[0])

	for _, arg := range argv {
		C.free(unsafe.Pointer(arg))
	}
}

func Exit() {
	C.alutExit()
}

func GetError() al.ALenum {
	return al.ALenum(C.alutGetError())
}

func GetErrorString(e al.ALenum) string {
	s := unsafe.Pointer(C.alutGetErrorString(C.ALenum(e)))
	return C.GoString((*C.char)(s))
}

func CreateBufferFromFile(fileName string) al.Buffer {
	s := C.CString(fileName)
	b := al.Buffer(C.alutCreateBufferFromFile(s))
	C.free(unsafe.Pointer(s))
	return b
}

func CreateBufferFromFileImage(data interface{}, length int) {
	C.alutCreateBufferFromFileImage(ptr(data), C.ALsizei(length))
}

func CreateBufferHelloWorld() al.Buffer {
	return al.Buffer(C.alutCreateBufferHelloWorld())
}

func CreateBufferWavefrom(
	waveshape al.ALenum,
	frequency float32,
	phase float32,
	duration float32,
) al.Buffer {
	return al.Buffer(C.alutCreateBufferWaveform(
		C.ALenum(waveshape),
		C.ALfloat(frequency),
		C.ALfloat(phase),
		C.ALfloat(duration),
	))
}

func LoadMemoryFromFile(fileName string) (
	format al.ALenum, size int32, frequency float32, rdata unsafe.Pointer) {
	cfileName := C.CString(fileName)
	rdata = C.alutLoadMemoryFromFile(
		cfileName,
		(*C.ALenum)(&format),
		(*C.ALsizei)(&size),
		(*C.ALfloat)(&frequency),
	)
	C.free(unsafe.Pointer(cfileName))
	return
}

func LoadMemoryFromFileImage(data interface{}, length int) (
	format al.ALenum, size int32, frequency float32, rdata unsafe.Pointer) {
	rdata = C.alutLoadMemoryFromFileImage(
		ptr(data), C.ALsizei(length),
		(*C.ALenum)(&format),
		(*C.ALsizei)(&size),
		(*C.ALfloat)(&frequency),
	)
	return
}

func LoadMemoryHelloWorld() (
	format al.ALenum, size int32, frequency float32, data unsafe.Pointer) {
	data = C.alutLoadMemoryHelloWorld(
		(*C.ALenum)(&format),
		(*C.ALsizei)(&size),
		(*C.ALfloat)(&frequency),
	)
	return
}

func LoadMemoryWaveform(waveshape al.ALenum, frequency float32, phase float32, duration float32) (
	format al.ALenum, size int32, freq float32, data unsafe.Pointer) {
	data = C.alutLoadMemoryWaveform(
		C.ALenum(waveshape), C.ALfloat(frequency), C.ALfloat(phase), C.ALfloat(duration),
		(*C.ALenum)(&format),
		(*C.ALsizei)(&size),
		(*C.ALfloat)(&frequency),
	)
	return
}

func GetMimeTypes(loader al.ALenum) string {
	return C.GoString(C.alutGetMIMETypes(C.ALenum(loader)))
}

func GetMajorVersion() int {
	return int(C.alutGetMajorVersion())
}

func GetMinorVersion() int {
	return int(C.alutGetMinorVersion())
}

func Sleep(duration float32) {
	C.alutSleep(C.ALfloat(duration))
}
