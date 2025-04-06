package threadsafe_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uncle-gua/threadsafe"
)

func TestNewStack(t *testing.T) {
	stack := threadsafe.NewStack[int]()
	assert.Equal(t, 0, stack.Len())
}

func TestStackPush(t *testing.T) {
	queue := threadsafe.NewStack[int]()
	queue.Push(42)
	queue.Push(43)
	queue.Push(44)

	value, ok := queue.Pop()
	assert.True(t, ok)
	assert.Equal(t, 44, value)
	assert.Equal(t, 2, queue.Len())

	value, ok = queue.Pop()
	assert.True(t, ok)
	assert.Equal(t, 43, value)
	assert.Equal(t, 1, queue.Len())

	value, ok = queue.Pop()
	assert.True(t, ok)
	assert.Equal(t, 42, value)
	assert.Equal(t, 0, queue.Len())

	value, ok = queue.Pop()
	assert.False(t, ok)
	assert.Zero(t, value)
	assert.Equal(t, 0, queue.Len())
}

func TestStackPop(t *testing.T) {
	stack := threadsafe.NewStack[int]()
	stack.Push(42)
	value, ok := stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, 42, value)
	assert.Equal(t, 0, stack.Len())
}

func TestStackPopEmpty(t *testing.T) {
	stack := threadsafe.NewStack[int]()
	value, ok := stack.Pop()
	assert.False(t, ok)
	assert.Zero(t, value)
}

func TestStackPeek(t *testing.T) {
	stack := threadsafe.NewStack[int]()
	stack.Push(42)
	value, ok := stack.Peek()
	assert.True(t, ok)
	assert.Equal(t, 42, value)
}

func TestStackPeekEmpty(t *testing.T) {
	stack := threadsafe.NewStack[int]()
	value, ok := stack.Peek()
	assert.False(t, ok)
	assert.Zero(t, value)
}

func TestStackLen(t *testing.T) {
	stack := threadsafe.NewStack[int]()
	stack.Push(42)
	stack.Push(43)
	assert.Equal(t, 2, stack.Len())
}

func TestStackIsEmpty(t *testing.T) {
	stack := threadsafe.NewStack[int]()
	assert.True(t, stack.IsEmpty())
	stack.Push(42)
	assert.False(t, stack.IsEmpty())
	stack.Pop()
	assert.True(t, stack.IsEmpty())
}

func TestStackClear(t *testing.T) {
	stack := threadsafe.NewStack[int]()
	stack.Push(42)
	stack.Push(43)
	stack.Clear()
	assert.Equal(t, 0, stack.Len())
	assert.True(t, stack.IsEmpty())
}

func TestStackValues(t *testing.T) {
	stack := threadsafe.NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	values := stack.Values()
	assert.Equal(t, 3, len(values))
	assert.Equal(t, 3, values[0])
	assert.Equal(t, 2, values[1])
	assert.Equal(t, 1, values[2])
}
