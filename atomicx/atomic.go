package atomx

import (
	"sync/atomic"
	"unsafe"
)

// GenericAtomic is a generic atomic utility for any type.
type GenericAtomic[T any] struct {
	value unsafe.Pointer
}

// NewAtomic creates a new GenericAtomic with the given initial value.
func NewAtomic[T any](initial T) *GenericAtomic[T] {
	ptr := unsafe.Pointer(&initial)
	return &GenericAtomic[T]{value: ptr}
}

// Load atomically loads the value.
func (a *GenericAtomic[T]) Load() T {
	return *(*T)(atomic.LoadPointer(&a.value))
}

// Store atomically stores the given value.
func (a *GenericAtomic[T]) Store(newValue T) {
	ptr := unsafe.Pointer(&newValue)
	atomic.StorePointer(&a.value, ptr)
}

// CompareAndSwap atomically swaps the old value with the new value if the old value matches the current value.
func (a *GenericAtomic[T]) CompareAndSwap(oldValue, newValue T) bool {
	oldPtr := unsafe.Pointer(&oldValue)
	newPtr := unsafe.Pointer(&newValue)
	return atomic.CompareAndSwapPointer(&a.value, oldPtr, newPtr)
}
