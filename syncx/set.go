package syncx

import "sync"

// Set represents a thread-safe generic set
type Set[T comparable] struct {
	m sync.Map
}

// Add adds an element to the set
func (s *Set[T]) Add(value T) {
	s.m.Store(value, struct{}{})
}

// Remove removes an element from the set
func (s *Set[T]) Remove(value T) {
	s.m.Delete(value)
}

// Get returns the element from the set
func (s *Set[T]) Get(value T) (T, bool) {
	_, ok := s.m.Load(value)
	return value, ok
}

// Contains checks if the set contains the element
func (s *Set[T]) Contains(value T) bool {
	_, ok := s.m.Load(value)
	return ok
}

// Range iterates over the elements in the set and applies the given function
func (s *Set[T]) Range(f func(value T) bool) {
	s.m.Range(func(key, _ interface{}) bool {
		return f(key.(T))
	})
}
