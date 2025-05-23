package threadsafe_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uncle-gua/threadsafe"
)

func TestNewArray(t *testing.T) {
	arr := threadsafe.NewArray[int](5)
	assert.Equal(t, 5, arr.Length())
}

func TestArrayGetSet(t *testing.T) {
	arr := threadsafe.NewArray[int](3)
	ok := arr.Set(2, 42)
	assert.True(t, ok)
	value, ok := arr.Get(2)
	assert.True(t, ok)
	assert.Equal(t, 42, value)
}

func TestArrayGetInvalidIndex(t *testing.T) {
	arr := threadsafe.NewArray[int](1)
	value, ok := arr.Get(10)
	assert.False(t, ok)
	assert.Equal(t, 0, value)
}

func TestArraySetInvalidIndex(t *testing.T) {
	arr := threadsafe.NewArray[int](1)
	ok := arr.Set(10, 42)
	assert.False(t, ok)
}

func TestArrayAppend(t *testing.T) {
	arr := threadsafe.NewArray[int](0)
	arr.Append(42)
	assert.Equal(t, 1, arr.Length())
	value, ok := arr.Get(0)
	assert.True(t, ok)
	assert.Equal(t, 42, value)
}

func TestArrayRemove(t *testing.T) {
	arr := threadsafe.NewArray[int](3)
	arr.Set(0, 1)
	arr.Set(1, 2)
	arr.Set(2, 3)
	ok := arr.Remove(1)
	assert.True(t, ok)
	assert.Equal(t, 2, arr.Length())
	value, ok := arr.Get(1)
	assert.True(t, ok)
	assert.Equal(t, 3, value)
}

func TestArrayRemoveInvalidIndex(t *testing.T) {
	arr := threadsafe.NewArray[int](1)
	ok := arr.Remove(10)
	assert.False(t, ok)
}

func TestArrayContains(t *testing.T) {
	arr := threadsafe.NewArray[int](3)
	arr.Set(0, 1)
	arr.Set(1, 2)
	arr.Set(2, 3)
	assert.True(t, arr.Contains(2))
	assert.False(t, arr.Contains(42))
}

func TestArrayClear(t *testing.T) {
	arr := threadsafe.NewArray[int](3)
	arr.Set(0, 1)
	arr.Set(1, 2)
	arr.Set(2, 3)
	arr.Clear()
	assert.Equal(t, 0, arr.Length())
}

func TestArrayInsert(t *testing.T) {
	arr := threadsafe.NewArray[int](3)
	arr.Set(0, 1)
	arr.Set(1, 2)
	arr.Set(2, 3)
	ok := arr.Insert(1, 42)
	assert.True(t, ok)
	assert.Equal(t, 4, arr.Length())
	value, ok := arr.Get(1)
	assert.True(t, ok)
	assert.Equal(t, 42, value)
}

func TestArrayInsertInvalidIndex(t *testing.T) {
	arr := threadsafe.NewArray[int](1)
	ok := arr.Insert(10, 42)
	assert.False(t, ok)
}

func TestArrayCopy(t *testing.T) {
	arr := threadsafe.NewArray[int](3)
	arr.Set(0, 1)
	arr.Set(1, 2)
	arr.Set(2, 3)
	copyArr := arr.Copy()
	assert.Equal(t, arr.Length(), copyArr.Length())
	for i := 0; i < arr.Length(); i++ {
		origValue, _ := arr.Get(i)
		copyValue, _ := copyArr.Get(i)
		assert.Equal(t, origValue, copyValue)
	}
}
func TestArrayValues(t *testing.T) {
	arr := threadsafe.NewArray[int](10)

	// Set values in the array
	for i := 0; i < arr.Length(); i++ {
		arr.Set(i, i*10)
	}

	// Get all values from the array
	values := arr.Values()

	// Check the length of the values slice
	assert.Equal(t, arr.Length(), len(values))

	// Check each value
	for i, v := range values {
		expectedValue, _ := arr.Get(i)
		assert.Equal(t, expectedValue, v)
	}
}
