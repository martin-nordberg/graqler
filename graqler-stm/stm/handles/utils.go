package handles

import (
	"math"
)

//-------------------------------------------------------------------------------------------------

type HandleId = uint64

//-------------------------------------------------------------------------------------------------

type BlueGreen = int64

const (
	Blue  BlueGreen = 0
	Green BlueGreen = 1
)

func OppositeColor(color BlueGreen) BlueGreen {
	if color == Blue {
		return Green
	}
	return Blue
}

//-------------------------------------------------------------------------------------------------

// See https://probablydance.com/2018/06/16/fibonacci-hashing-the-optimization-that-the-world-forgot-or-a-better-alternative-to-integer-modulo/
const two64DivPhi = 11400714819323198485

//-------------------------------------------------------------------------------------------------

const handleMapFillFactor = 0.7

//-------------------------------------------------------------------------------------------------

func arraySize(capacity uint32) uint32 {
	return nextPowerOf2(uint32(math.Ceil(float64(capacity) / handleMapFillFactor)))
}

//-------------------------------------------------------------------------------------------------

func nextPowerOf2(x uint32) uint32 {
	if x == math.MaxUint32 {
		return x
	}

	if x < 2 {
		return 2
	}

	// See: https://stackoverflow.com/questions/1322510/given-an-integer-how-do-i-find-the-next-largest-power-of-two-using-bit-twiddlin/1322548#1322548
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16

	return x + 1
}

//-------------------------------------------------------------------------------------------------
