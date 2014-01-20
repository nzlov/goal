package al


// #include "al.h"
import "C"

type Object C.ALuint

func (object Object) IsSource() bool{
	return goBool(C.alIsSource(C.ALuint(object)))
}

func (object Object) IsBuffer() bool{
	return goBool(C.alIsBuffer(C.ALuint(object)))
}
