package al

// #include "al.h"
import "C"

/**
 * SOURCE
 * Sources represent individual sound objects in 3D-space.
 * Sources take the PCM data provided in the specified Buffer,
 * apply Source-specific modifications, and then
 * submit them to be mixed according to spatial arrangement etc.
 * 
 * Properties include: -
 *
 * Gain                              AL_GAIN                 ALfloat
 * Min Gain                          AL_MIN_GAIN             ALfloat
 * Max Gain                          AL_MAX_GAIN             ALfloat
 * Position                          AL_POSITION             ALfloat[3]
 * Velocity                          AL_VELOCITY             ALfloat[3]
 * Direction                         AL_DIRECTION            ALfloat[3]
 * Head Relative Mode                AL_SOURCE_RELATIVE      ALint (AL_TRUE or AL_FALSE)
 * Reference Distance                AL_REFERENCE_DISTANCE   ALfloat
 * Max Distance                      AL_MAX_DISTANCE         ALfloat
 * RollOff Factor                    AL_ROLLOFF_FACTOR       ALfloat
 * Inner Angle                       AL_CONE_INNER_ANGLE     ALint or ALfloat
 * Outer Angle                       AL_CONE_OUTER_ANGLE     ALint or ALfloat
 * Cone Outer Gain                   AL_CONE_OUTER_GAIN      ALint or ALfloat
 * Pitch                             AL_PITCH                ALfloat
 * Looping                           AL_LOOPING              ALint (AL_TRUE or AL_FALSE)
 * MS Offset                         AL_MSEC_OFFSET          ALint or ALfloat
 * Byte Offset                       AL_BYTE_OFFSET          ALint or ALfloat
 * Sample Offset                     AL_SAMPLE_OFFSET        ALint or ALfloat
 * Attached Buffer                   AL_BUFFER               ALint
 * State (Query only)                AL_SOURCE_STATE         ALint
 * Buffers Queued (Query only)       AL_BUFFERS_QUEUED       ALint
 * Buffers Processed (Query only)    AL_BUFFERS_PROCESSED    ALint
 */
type Source Object

func GenSource() Source {
	var b C.ALuint
	C.alGenSources(1, &b)
	return Source(b)
}

func GenSources(sources []Source) {
	if len(sources) > 0 {
		C.alGenSources(C.ALsizei(len(sources)), (*C.ALuint)(&sources[0]))
	}
}

func (source Source) Delete() {
	b := C.ALuint(source)
	C.alDeleteSources(1, &b)
}

func DeleteSources(sources []Source) {
	if len(sources) > 0 {
		C.alDeleteSources(C.ALsizei(len(sources)), (*C.ALuint)(&sources[0]))
	}
}

/*
 * Set Source parameters
 */
func (source Source) Sourcef(param ALenum, value float32) {
	C.alSourcef(C.ALuint(source), C.ALenum(param), C.ALfloat(value))
}

func (source Source) Source3f(
	param ALenum, value1, value2, value3 float32) {
	C.alSource3f(
		C.ALuint(source), C.ALenum(param),
		C.ALfloat(value1), C.ALfloat(value2), C.ALfloat(value3))
}

func (source Source) Sourcefv(param ALenum, values []float32) {
	C.alSourcefv(C.ALuint(source), C.ALenum(param), (*C.ALfloat)(&values[0]))
}

func (source Source) Sourcei(param ALenum, value int) {
	C.alSourcei(C.ALuint(source), C.ALenum(param), C.ALint(value))
}

func (source Source) Source3i(
	param ALenum, value1, value2, value3 int) {
	C.alSource3i(
		C.ALuint(source), C.ALenum(param),
		C.ALint(value1), C.ALint(value2), C.ALint(value3))
}

func (source Source) Sourceiv(param ALenum, values []int32) {
	C.alSourceiv(C.ALuint(source), C.ALenum(param), (*C.ALint)(&values[0]))
}

/*
 * Get Source parameters
 */
func (source Source) GetSourcef(param ALenum) (value float32) {
	C.alGetSourcef(
		C.ALuint(source), C.ALenum(param),
		(*C.ALfloat)(&value))
	return
}

func (source Source) GetSource3f(param ALenum) (value1, value2, value3 float32) {
	C.alGetSource3f(
		C.ALuint(source), C.ALenum(param),
		(*C.ALfloat)(&value1),
		(*C.ALfloat)(&value2),
		(*C.ALfloat)(&value3),
	)
	return
}

func (source Source) GetSourcefv(param ALenum, values []float32) {
	C.alGetSourcefv(C.ALuint(source), C.ALenum(param), (*C.ALfloat)(&values[0]))
}

func (source Source) GetSourcei(param ALenum) (value int32) {
	C.alGetSourcei(C.ALuint(source), C.ALenum(param), (*C.ALint)(&value))
	return
}

func (source Source) GetSource3i(param ALenum) (value1, value2, value3 int32) {
	C.alGetSource3i(
		C.ALuint(source), C.ALenum(param),
		(*C.ALint)(&value1),
		(*C.ALint)(&value2),
		(*C.ALint)(&value3),
	)
	return
}

func (source Source) GetSourceiv(param ALenum, values []int32) {
	C.alGetSourceiv(C.ALuint(source), C.ALenum(param), (*C.ALint)(&values[0]))
}

/*
 * Source vector based playback calls
 */
func (source Source) Playv(sources []Source) {
	C.alSourcePlayv(C.ALsizei(len(sources)), (*C.ALuint)(&sources[0]))
}

func (source Source) Stopv(sources []Source) {
	C.alSourceStopv(C.ALsizei(len(sources)), (*C.ALuint)(&sources[0]))
}

func (source Source) Rewindv(sources []Source) {
	C.alSourceRewindv(C.ALsizei(len(sources)), (*C.ALuint)(&sources[0]))
}

func (source Source) Pausev(sources []Source) {
	C.alSourcePausev(C.ALsizei(len(sources)), (*C.ALuint)(&sources[0]))
}

/*
 * Source based playback calls
 */
func (source Source) Play() {
	C.alSourcePlay(C.ALuint(source))
}

func (source Source) Stop() {
	C.alSourceStop(C.ALuint(source))
}

func (source Source) Rewind() {
	C.alSourceRewind(C.ALuint(source))
}

func (source Source) Pause() {
	C.alSourcePause(C.ALuint(source))
}

/*
 * Source Queuing 
 */
func (source Source) QueueBuffer(buffer Buffer){
	C.alSourceQueueBuffers(C.ALuint(source), 1, (*C.ALuint)(&buffer))
}

func (source Source) QueueBuffers(buffers []Buffer){
	C.alSourceQueueBuffers(
		C.ALuint(source), C.ALsizei(len(buffers)), (*C.ALuint)(&buffers[0]))
}

func (source Source) UnqueueBuffer() Buffer{
	var b C.ALuint
	C.alSourceUnqueueBuffers(C.ALuint(source), 1, &b)
	return Buffer(b)
}

func (source Source) UnqueueBuffers(buffers []Buffer){
	C.alSourceUnqueueBuffers(
		C.ALuint(source), C.ALsizei(len(buffers)), (*C.ALuint)(&buffers[0]))
}
