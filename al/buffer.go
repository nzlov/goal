package al

// #include "al.h"
import "C"

/**
 * BUFFER
 * Buffer objects are storage space for sample data.
 * Buffers are referred to by Sources. One Buffer can be used
 * by multiple Sources.
 *
 * Properties include: -
 *
 * Frequency (Query only)    AL_FREQUENCY      ALint
 * Size (Query only)         AL_SIZE           ALint
 * Bits (Query only)         AL_BITS           ALint
 * Channels (Query only)     AL_CHANNELS       ALint
 */
type Buffer Object

func GenBuffer() Buffer {
	var b C.ALuint
	C.alGenBuffers(1, &b)
	return Buffer(b)
}

func GenBuffers(buffers []Buffer) {
	if len(buffers) > 0 {
		C.alGenBuffers(C.ALsizei(len(buffers)), (*C.ALuint)(&buffers[0]))
	}
}

func (buffer Buffer) Delete() {
	b := C.ALuint(buffer)
	C.alDeleteBuffers(1, &b)
}

func DeleteBuffers(buffers []Buffer) {
	if len(buffers) > 0 {
		C.alDeleteBuffers(C.ALsizei(len(buffers)), (*C.ALuint)(&buffers[0]))
	}
}

func (buffer Buffer) IsBuffer() bool{
	return goBool(C.alIsBuffer(C.ALuint(buffer)))
}

func (buffer Buffer) BufferData(format ALenum, data interface{}, size int, freq int){
	C.alBufferData(C.ALuint(buffer), C.ALenum(format), ptr(data), C.ALsizei(size), C.ALsizei(freq))
}

/*
 * Set Buffer parameters
 */

func (buffer Buffer) Bufferf(param ALenum, value float32) {
	C.alBufferf(C.ALuint(buffer), C.ALenum(param), C.ALfloat(value))
}

func (buffer Buffer) Buffer3f(
	param ALenum, value1, value2, value3 float32) {
	C.alBuffer3f(
		C.ALuint(buffer), C.ALenum(param),
		C.ALfloat(value1), C.ALfloat(value2), C.ALfloat(value3))
}

func (buffer Buffer) Bufferfv(param ALenum, values []float32) {
	C.alBufferfv(C.ALuint(buffer), C.ALenum(param), (*C.ALfloat)(&values[0]))
}

func (buffer Buffer) Bufferi(param ALenum, value int) {
	C.alBufferi(C.ALuint(buffer), C.ALenum(param), C.ALint(value))
}

func (buffer Buffer) Buffer3i(
	param ALenum, value1, value2, value3 int) {
	C.alBuffer3i(
		C.ALuint(buffer), C.ALenum(param),
		C.ALint(value1), C.ALint(value2), C.ALint(value3))
}

func (buffer Buffer) Bufferiv(param ALenum, values []int32) {
	C.alBufferiv(C.ALuint(buffer), C.ALenum(param), (*C.ALint)(&values[0]))
}


/*
 * Get Buffer parameters
 */

func (buffer Buffer) GetBufferf(param ALenum) (value float32) {
	C.alGetBufferf(
		C.ALuint(buffer), C.ALenum(param),
		(*C.ALfloat)(&value))
	return
}

func (buffer Buffer) GetBuffer3f(param ALenum) (value1, value2, value3 float32) {
	C.alGetBuffer3f(
		C.ALuint(buffer), C.ALenum(param),
		(*C.ALfloat)(&value1),
		(*C.ALfloat)(&value2),
		(*C.ALfloat)(&value3),
	)
	return
}

func (buffer Buffer) GetBufferfv(param ALenum, values []float32) {
	C.alGetBufferfv(C.ALuint(buffer), C.ALenum(param), (*C.ALfloat)(&values[0]))
}

func (buffer Buffer) GetBufferi(param ALenum) (value int32) {
	C.alGetBufferi(C.ALuint(buffer), C.ALenum(param), (*C.ALint)(&value))
	return
}

func (buffer Buffer) GetBuffer3i(param ALenum) (value1, value2, value3 int32) {
	C.alGetBuffer3i(
		C.ALuint(buffer), C.ALenum(param),
		(*C.ALint)(&value1),
		(*C.ALint)(&value2),
		(*C.ALint)(&value3),
	)
	return
}

func (buffer Buffer) GetBufferiv(param ALenum, values []int32) {
	C.alGetBufferiv(C.ALuint(buffer), C.ALenum(param), (*C.ALint)(&values[0]))
}

