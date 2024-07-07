package handles_test

import (
	"graqler/stm/stm/assert"
	. "graqler/stm/stm/handles"
	"testing"
)

func TestMapPutGet(t *testing.T) {

	t.Run("Checking retrieval of values added to an Int32HandleMap", func(t *testing.T) {
		const handleCount = 1000000
		const expectedCapacity uint32 = 0x200000

		const cloneCount = 4000
		const expectedCloneCapacity uint32 = 16384

		const color = Blue

		h := make([]Int32Handle, handleCount+1)
		for i := 1; i <= handleCount; i += 1 {
			h[i] = MakeInt32Handle(int32(i))
		}

		m1 := MakeInt32HandleMap(32)
		var m2 Int32HandleMap
		for i := 1; i <= handleCount; i += 1 {
			m1.Put(h[i].HandleId(), h[i].Read(color)*2)

			if i == cloneCount {
				m2 = m1.Clone(2000)
			}
		}

		m1.Freeze()
		m2.Freeze()

		assert.Equals(t, uint32(handleCount), m1.Size())
		assert.Equals(t, expectedCapacity, m1.Capacity())

		assert.Equals(t, uint32(cloneCount), m2.Size())
		assert.Equals(t, expectedCloneCapacity, m2.Capacity())

		for i := 1; i <= handleCount; i += 1 {
			assert.True(t, m1.Has(h[i].HandleId()), "Missing m1 key")

			expected := int32(i * 2)
			actual, found := m1.Get(h[i].HandleId())
			assert.True(t, found, "Key not found in m1")
			assert.Equals(t, expected, actual)
		}

		for i := 1; i <= cloneCount; i += 1 {
			assert.True(t, m2.Has(h[i].HandleId()), "Missing m2 key")

			expected := int32(i * 2)
			actual, found := m2.Get(h[i].HandleId())
			assert.True(t, found, "Key not found in m2")
			assert.Equals(t, expected, actual)
		}

	})

}
