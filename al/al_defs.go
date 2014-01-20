package al

// #include "al.h"
import "C"

//Constants
const (
	/* "no distance model" or "no buffer" */
	NONE = C.AL_NONE

	/* Boolean False. */
	FALSE = C.AL_FALSE

	/** Boolean True. */
	TRUE = C.AL_TRUE

	/** Indicate Source has relative coordinates. */
	SOURCE_RELATIVE = C.AL_SOURCE_RELATIVE

	/**
	 * Directional source, inner cone angle, in degrees.
	 * Range:    [0-360] 
	 * Default:  360
	 */
	CONE_INNER_ANGLE = C.AL_CONE_INNER_ANGLE

	/**
	 * Directional source, outer cone angle, in degrees.
	 * Range:    [0-360] 
	 * Default:  360
	 */
	CONE_OUTER_ANGLE = C.AL_CONE_OUTER_ANGLE

	/**
	 * Specify the pitch to be applied, either at source,
	 *  or on mixer results, at listener.
	 * Range:   [0.5-2.0]
	 * Default: 1.0
	 */
	PITCH = C.AL_PITCH

	/** 
	 * Specify the current location in three dimensional space.
	 * OpenAL, like OpenGL, uses a right handed coordinate system,
	 *  where in a frontal default view X (thumb) points right, 
	 *  Y points up (index finger), and Z points towards the
	 *  viewer/camera (middle finger). 
	 * To switch from a left handed coordinate system, flip the
	 *  sign on the Z coordinate.
	 * Listener position is always in the world coordinate system.
	 */
	POSITION = C.AL_POSITION

	/** Specify the current direction. */
	DIRECTION = C.AL_DIRECTION

	/** Specify the current velocity in three dimensional space. */
	VELOCITY = C.AL_VELOCITY

	/**
	 * Indicate whether source is looping.
	 * Type: ALboolean?
	 * Range:   [AL_TRUE, AL_FALSE]
	 * Default: FALSE.
	 */
	LOOPING = C.AL_LOOPING

	/**
	 * Indicate the buffer to provide sound samples. 
	 * Type: ALuint.
	 * Range: any valid Buffer id.
	 */
	BUFFER = C.AL_BUFFER

	/**
	 * Indicate the gain (volume amplification) applied. 
	 * Type:   ALfloat.
	 * Range:  ]0.0-  ]
	 * A value of 1.0 means un-attenuated/unchanged.
	 * Each division by 2 equals an attenuation of -6dB.
	 * Each multiplicaton with 2 equals an amplification of +6dB.
	 * A value of 0.0 is meaningless with respect to a logarithmic
	 *  scale; it is interpreted as zero volume - the channel
	 *  is effectively disabled.
	 */
	GAIN = C.AL_GAIN

	/*
	 * Indicate minimum source attenuation
	 * Type: ALfloat
	 * Range:  [0.0 - 1.0]
	 *
	 * Logarthmic
	 */
	MIN_GAIN = C.AL_MIN_GAIN

	/**
	 * Indicate maximum source attenuation
	 * Type: ALfloat
	 * Range:  [0.0 - 1.0]
	 *
	 * Logarthmic
	 */
	MAX_GAIN = C.AL_MAX_GAIN

	/**
	 * Indicate listener orientation.
	 *
	 * at/up 
	 */
	ORIENTATION = C.AL_ORIENTATION

	/**
	 * Source state information.
	 */
	SOURCE_STATE = C.AL_SOURCE_STATE
	INITIAL      = C.AL_INITIAL
	PLAYING      = C.AL_PLAYING
	PAUSED       = C.AL_PAUSED
	STOPPED      = C.AL_STOPPED

	/**
	 * Buffer Queue params
	 */
	BUFFERS_QUEUED    = C.AL_BUFFERS_QUEUED
	BUFFERS_PROCESSED = C.AL_BUFFERS_PROCESSED

	/**
	 * Source buffer position information
	 */
	SEC_OFFSET    = C.AL_SEC_OFFSET
	SAMPLE_OFFSET = C.AL_SAMPLE_OFFSET
	BYTE_OFFSET   = C.AL_BYTE_OFFSET

	/*
	 * Source type (Static, Streaming or undetermined)
	 * Source is Static if a Buffer has been attached using AL_BUFFER
	 * Source is Streaming if one or more Buffers have been attached using alSourceQueueBuffers
	 * Source is undetermined when it has the NULL buffer attached
	 */
	SOURCE_TYPE  = C.AL_SOURCE_TYPE
	STATIC       = C.AL_STATIC
	STREAMING    = C.AL_STREAMING
	UNDETERMINED = C.AL_UNDETERMINED

	/** Sound samples: format specifier. */
	FORMAT_MONO8    = C.AL_FORMAT_MONO8
	FORMAT_MONO16   = C.AL_FORMAT_MONO16
	FORMAT_STEREO8  = C.AL_FORMAT_STEREO8
	FORMAT_STEREO16 = C.AL_FORMAT_STEREO16

	/**
	 * source specific reference distance
	 * Type: ALfloat
	 * Range:  0.0 - +inf
	 *
	 * At 0.0, no distance attenuation occurs.  Default is
	 * 1.0.
	 */
	REFERENCE_DISTANCE = C.AL_REFERENCE_DISTANCE

	/**
	 * source specific rolloff factor
	 * Type: ALfloat
	 * Range:  0.0 - +inf
	 *
	 */
	ROLLOFF_FACTOR = C.AL_ROLLOFF_FACTOR

	/**
	 * Directional source, outer cone gain.
	 *
	 * Default:  0.0
	 * Range:    [0.0 - 1.0]
	 * Logarithmic
	 */
	CONE_OUTER_GAIN = C.AL_CONE_OUTER_GAIN

	/**
	 * Indicate distance above which sources are not
	 * attenuated using the inverse clamped distance model.
	 *
	 * Default: +inf
	 * Type: ALfloat
	 * Range:  0.0 - +inf
	 */
	MAX_DISTANCE = C.AL_MAX_DISTANCE

	/** 
	 * Sound samples: frequency, in units of Hertz [Hz].
	 * This is the number of samples per second. Half of the
	 *  sample frequency marks the maximum significant
	 *  frequency component.
	 */
	FREQUENCY = C.AL_FREQUENCY
	BITS      = C.AL_BITS
	CHANNELS  = C.AL_CHANNELS
	SIZE      = C.AL_SIZE

	/**
	 * Buffer state.
	 *
	 * Not supported for public use (yet).
	 */
	UNUSED    = C.AL_UNUSED
	PENDING   = C.AL_PENDING
	PROCESSED = C.AL_PROCESSED

	/** Errors: No Error. */
	NO_ERROR = C.AL_NO_ERROR

	/** 
	 * Invalid Name paramater passed to AL call.
	 */
	INVALID_NAME = C.AL_INVALID_NAME

	/** 
	 * Invalid parameter passed to AL call.
	 */
	INVALID_ENUM = C.AL_INVALID_ENUM

	/** 
	 * Invalid enum parameter value.
	 */
	INVALID_VALUE = C.AL_INVALID_VALUE

	/** 
	 * Illegal call.
	 */
	INVALID_OPERATION = C.AL_INVALID_OPERATION

	/**
	 * No mojo.
	 */
	OUT_OF_MEMORY = C.AL_OUT_OF_MEMORY

	/** Context strings: Vendor Name. */
	VENDOR     = C.AL_VENDOR
	VERSION    = C.AL_VERSION
	RENDERER   = C.AL_RENDERER
	EXTENSIONS = C.AL_EXTENSIONS

	/** Global tweakage. */

	/**
	 * Doppler scale.  Default 1.0
	 */
	DOPPLER_FACTOR = C.AL_DOPPLER_FACTOR

	/**
	 * Tweaks speed of propagation.
	 */
	DOPPLER_VELOCITY = C.AL_DOPPLER_VELOCITY

	/**
	 * Speed of Sound in units per second
	 */
	SPEED_OF_SOUND = C.AL_SPEED_OF_SOUND

	/**
	 * Distance models
	 *
	 * used in conjunction with DistanceModel
	 *
	 * implicit: NONE, which disances distance attenuation.
	 */
	DISTANCE_MODEL            = C.AL_DISTANCE_MODEL
	INVERSE_DISTANCE          = C.AL_INVERSE_DISTANCE
	INVERSE_DISTANCE_CLAMPED  = C.AL_INVERSE_DISTANCE_CLAMPED
	LINEAR_DISTANCE           = C.AL_LINEAR_DISTANCE
	LINEAR_DISTANCE_CLAMPED   = C.AL_LINEAR_DISTANCE_CLAMPED
	EXPONENT_DISTANCE         = C.AL_EXPONENT_DISTANCE
	EXPONENT_DISTANCE_CLAMPED = C.AL_EXPONENT_DISTANCE_CLAMPED
)
