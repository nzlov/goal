package alc

// #include "alc.h"
import "C"

const (
	/* Boolean False. */
	FALSE = C.ALC_FALSE

	/* Boolean True. */
	TRUE = C.ALC_TRUE

	/**
	 * followed by <int> Hz
	 */
	FREQUENCY = C.ALC_FREQUENCY

	/**
	 * followed by <int> Hz
	 */
	REFRESH = C.ALC_REFRESH

	/**
	 * followed by AL_TRUE, AL_FALSE
	 */
	SYNC = C.ALC_SYNC

	/**
	 * followed by <int> Num of requested Mono (3D) Sources
	 */
	MONO_SOURCES = C.ALC_MONO_SOURCES

	/**
	 * followed by <int> Num of requested Stereo Sources
	 */
	STEREO_SOURCES = C.ALC_STEREO_SOURCES

	/**
	 * errors
	 */

	/**
	 * No error
	 */
	NO_ERROR = C.ALC_NO_ERROR

	/**
	 * No device
	 */
	INVALID_DEVICE = C.ALC_INVALID_DEVICE

	/**
	 * invalid context ID
	 */
	INVALID_CONTEXT = C.ALC_INVALID_CONTEXT

	/**
	 * bad enum
	 */
	INVALID_ENUM = C.ALC_INVALID_ENUM

	/**
	 * bad value
	 */
	INVALID_VALUE = C.ALC_INVALID_VALUE

	/**
	 * Out of memory.
	 */
	OUT_OF_MEMORY = C.ALC_OUT_OF_MEMORY

	/**
	 * The Specifier string for default device
	 */
	DEFAULT_DEVICE_SPECIFIER = C.ALC_DEFAULT_DEVICE_SPECIFIER
	DEVICE_SPECIFIER         = C.ALC_DEVICE_SPECIFIER
	EXTENSIONS               = C.ALC_EXTENSIONS

	MAJOR_VERSION = C.ALC_MAJOR_VERSION
	MINOR_VERSION = C.ALC_MINOR_VERSION

	ATTRIBUTES_SIZE = C.ALC_ATTRIBUTES_SIZE
	ALL_ATTRIBUTES  = C.ALC_ALL_ATTRIBUTES

	/**
	 * ALC_ENUMERATE_ALL_EXT enums
	 */
	DEFAULT_ALL_DEVICES_SPECIFIER = C.ALC_DEFAULT_ALL_DEVICES_SPECIFIER
	ALL_DEVICES_SPECIFIER         = C.ALC_ALL_DEVICES_SPECIFIER

	/**
	 * Capture extension
	 */
	CAPTURE_DEVICE_SPECIFIER         = C.ALC_CAPTURE_DEVICE_SPECIFIER
	CAPTURE_DEFAULT_DEVICE_SPECIFIER = C.ALC_CAPTURE_DEFAULT_DEVICE_SPECIFIER
	CAPTURE_SAMPLES                  = C.ALC_CAPTURE_SAMPLES
)
