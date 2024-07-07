package handles

import (
	"math"
)

//-------------------------------------------------------------------------------------------------

// Heavily modified from the original code found here:
// https://github.com/kamstrup/intmap/blob/main/map64.go
// BSD-2-Clause License

//-------------------------------------------------------------------------------------------------

// int32Entry is a key/value pair in the map.
type int32Entry struct {
	key   HandleId
	value int32
}

//-------------------------------------------------------------------------------------------------

// Int32HandleMap is a hash map where the keys are STM object handles and the values are int32.
type Int32HandleMap struct {
	data            []int32Entry
	dataLenMinus1   uint32
	frozen          bool
	size            uint32
	growthThreshold uint32
}

//-------------------------------------------------------------------------------------------------

// MakeInt32HandleMap creates a new map with keys being STM handle IDs.
// The map can store up to the given capacity before reallocation and rehashing occurs.
func MakeInt32HandleMap(capacity uint32) Int32HandleMap {
	dataLen := arraySize(capacity)
	return Int32HandleMap{
		data:            make([]int32Entry, dataLen),
		dataLenMinus1:   dataLen - 1,
		frozen:          false,
		size:            0,
		growthThreshold: uint32(math.Floor(float64(dataLen) * handleMapFillFactor)),
	}
}

//-------------------------------------------------------------------------------------------------

// Capacity returns the number of elements in the underlying array.
func (m *Int32HandleMap) Capacity() uint32 {
	return m.dataLenMinus1 + 1
}

//-------------------------------------------------------------------------------------------------

// Clone makes a copy of this map with room for addedCapacity more entries in the (unfrozen) copy.
func (m *Int32HandleMap) Clone(addedCapacity uint32) Int32HandleMap {
	// Find the new needed capacity.
	capacity := m.size + addedCapacity

	// Allocate the clone.
	result := MakeInt32HandleMap(capacity)

	// If not growing, can just straight copy the data.
	if result.dataLenMinus1 == m.dataLenMinus1 {
		copy(result.data, m.data)
		result.size = m.size
		return result
	}

	// Otherwise, copy over all the values individually to spread them out in the new space.
	forEachEntry(m.data, func(k HandleId, v int32) {
		result.Put(k, v)
	})

	return result
}

//-------------------------------------------------------------------------------------------------

// ForEach iterates through key-value pairs in the map, calling function f for each entry.
// The iteration order is not defined.
func (m *Int32HandleMap) ForEach(f func(HandleId, int32)) {
	forEachEntry(m.data, f)
}

//-------------------------------------------------------------------------------------------------

// Freeze marks this map as no longer writable (future attempts will panic).
func (m *Int32HandleMap) Freeze() {
	m.frozen = true
}

//-------------------------------------------------------------------------------------------------

// Get returns the value if the key is found.
func (m *Int32HandleMap) Get(key HandleId) (int32, bool) {
	for idx := m.startIndex(key); true; idx = m.nextIndex(idx) {
		entry := m.data[idx]
		if entry.key == 0 {
			return 0, false
		}
		if entry.key == key {
			return entry.value, true
		}
	}

	panic("Unreachable")
}

//-------------------------------------------------------------------------------------------------

// Has checks if the given key exists in the map.
func (m *Int32HandleMap) Has(key HandleId) bool {
	for idx := m.startIndex(key); true; idx = m.nextIndex(idx) {
		entryKey := m.data[idx].key
		if entryKey == 0 {
			return false
		}
		if entryKey == key {
			return true
		}
	}

	panic("Unreachable")
}

//-------------------------------------------------------------------------------------------------

// Put adds or updates key with value val.
func (m *Int32HandleMap) Put(key HandleId, val int32) {
	if m.frozen {
		panic("Map is frozen")
	}

	for idx := m.startIndex(key); true; idx = m.nextIndex(idx) {
		entry := &m.data[idx]

		if entry.key == 0 { // end of chain
			entry.key = key
			entry.value = val
			m.size += 1
			if m.size >= m.growthThreshold {
				m.grow()
			}
			return
		}
		if entry.key == key { // overwrite existing value
			entry.value = val
			return
		}
	}
}

//-------------------------------------------------------------------------------------------------

// PutIfNotExists adds the key-value entry only if the key does not already exist
// in the map, and returns the current value associated with the key and a boolean
// indicating whether the value was newly added or not.
func (m *Int32HandleMap) PutIfNotExists(key HandleId, val int32) (int32, bool) {
	if m.frozen {
		panic("Map is frozen")
	}

	for idx := m.startIndex(key); true; idx = m.nextIndex(idx) {
		entry := &m.data[idx]

		if entry.key == 0 {
			entry.key = key
			entry.value = val
			m.size += 1
			if m.size >= m.growthThreshold {
				m.grow()
			}
			return val, true
		}
		if entry.key == key {
			return entry.value, false
		}
	}

	panic("Unreachable")
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of elements in the map.
func (m *Int32HandleMap) Size() uint32 {
	return m.size
}

//=================================================================================================

// forEachEntry iterates over a given array of entries.
func forEachEntry(entries []int32Entry, f func(key HandleId, val int32)) {
	for _, p := range entries {
		if p.key != 0 {
			f(p.key, p.value)
		}
	}
}

//-------------------------------------------------------------------------------------------------

// grow doubles the capacity of the map.
func (m *Int32HandleMap) grow() {
	oldData := m.data
	dataLen := uint32(2 * len(m.data))

	// Resize
	m.data = make([]int32Entry, dataLen)
	m.dataLenMinus1 = dataLen - 1
	m.size = 0
	m.growthThreshold = uint32(math.Floor(float64(dataLen) * handleMapFillFactor))

	// Re-add all the entries.
	forEachEntry(oldData, func(k HandleId, v int32) {
		m.Put(k, v)
	})
}

//-------------------------------------------------------------------------------------------------

// nextIndex returns the next index in the map's data following the given one (modulo the map capacity).
func (m *Int32HandleMap) nextIndex(priorIdx uint32) uint32 {
	return (priorIdx + 1) & m.dataLenMinus1
}

//-------------------------------------------------------------------------------------------------

// startIndex returns the array index for given key (modulo the map capacity).
func (m *Int32HandleMap) startIndex(key HandleId) uint32 {
	return uint32(key & HandleId(m.dataLenMinus1))
}

//-------------------------------------------------------------------------------------------------
