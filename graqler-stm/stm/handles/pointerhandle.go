package handles

import (
	"sync/atomic"
)

//-------------------------------------------------------------------------------------------------

type PointerHandle[T any] struct {
	handleId HandleId
	value    [2]*T
}

//-------------------------------------------------------------------------------------------------

func MakePointerHandle[T any](value *T) PointerHandle[T] {
	return PointerHandle[T]{
		value:    [2]*T{value, value},
		handleId: getNextPointerHandleId(),
	}
}

//-------------------------------------------------------------------------------------------------

func (h *PointerHandle[T]) HandleId() HandleId {
	return h.handleId
}

//-------------------------------------------------------------------------------------------------

func (h *PointerHandle[T]) Read(blueOrGreen BlueGreen) *T {
	return h.value[blueOrGreen]
}

//-------------------------------------------------------------------------------------------------

func (h *PointerHandle[T]) Write(blueOrGreen BlueGreen, value *T) {
	h.value[blueOrGreen] = value
}

//=================================================================================================

var nextPointerHandleId atomic.Uint64

//-------------------------------------------------------------------------------------------------

func init() {
	nextPointerHandleId.Store(1)
}

//-------------------------------------------------------------------------------------------------

func getNextPointerHandleId() HandleId {
	for {
		var h = nextPointerHandleId.Load()
		var result = h * two64DivPhi
		if nextPointerHandleId.CompareAndSwap(h, result) {
			return result
		}
	}
}

//-------------------------------------------------------------------------------------------------
