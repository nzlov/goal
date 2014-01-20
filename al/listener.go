package al

// #include "al.h"
import "C"

/*
 * LISTENER
 * Listener represents the location and orientation of the
 * 'user' in 3D-space.
 *
 * Properties include: -
 *
 * Gain         AL_GAIN         ALfloat
 * Position     AL_POSITION     ALfloat[3]
 * Velocity     AL_VELOCITY     ALfloat[3]
 * Orientation  AL_ORIENTATION  ALfloat[6] (Forward then Up vectors)
*/
type Listener Object

/*
 * Set Listener parameters
 */
func Listenerf(param ALenum, value float32) {
	C.alListenerf(C.ALenum(param), C.ALfloat(value))
}

func Listener3f(param ALenum, value1, value2, value3 float32) {
	C.alListener3f(C.ALenum(param), C.ALfloat(value1), C.ALfloat(value2), C.ALfloat(value3))
}

func Listenerfv(param ALenum, values []float32) {
	C.alListenerfv(C.ALenum(param), (*C.ALfloat)(&values[0]))
}

func Listeneri(param ALenum, value int) {
	C.alListeneri(C.ALenum(param), C.ALint(value))
}

func Listener3i(param ALenum, value1, value2, value3 int) {
	C.alListener3i(C.ALenum(param), C.ALint(value1), C.ALint(value2), C.ALint(value3))
}

func Listeneriv(param ALenum, values []int32) {
	C.alListeneriv(C.ALenum(param), (*C.ALint)(&values[0]))
}

/*
 * Get Listener parameters
 */
func GetListenerf(param ALenum) (value float32) {
	var cvalue C.ALfloat
	C.alGetListenerf(C.ALenum(param), &cvalue)
	value = float32(cvalue)
	return
}

func GetListener3f(param ALenum) (value1, value2, value3 float32) {
	C.alGetListener3f(
		C.ALenum(param),
		(*C.ALfloat)(&value1),
		(*C.ALfloat)(&value2),
		(*C.ALfloat)(&value3),
	)
	return
}

func GetListenerfv(param ALenum, values []float32) {
	C.alGetListenerfv(C.ALenum(param), (*C.ALfloat)(&values[0]))
}

func GetListeneri(param ALenum) (value int32) {
	C.alGetListeneri(C.ALenum(param), (*C.ALint)(&value))
	return
}

func GetListener3i(param ALenum) (value1, value2, value3 int32) {
	C.alGetListener3i(
		C.ALenum(param),
		(*C.ALint)(&value1),
		(*C.ALint)(&value2),
		(*C.ALint)(&value3),
	)
	return
}

func GetListeneriv(param ALenum, values []int32) {
	C.alGetListeneriv(C.ALenum(param), (*C.ALint)(&values[0]))
}

