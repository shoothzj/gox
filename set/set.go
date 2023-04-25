package set

// Set represents a generic set
type Set[T comparable] map[T]struct{}

func (set Set[T]) Add(element T) {
	set[element] = struct{}{}
}

func (set Set[T]) Remove(element T) {
	delete(set, element)
}

func (set Set[T]) Contains(element T) bool {
	_, exists := set[element]
	return exists
}

func (set Set[T]) Len() int {
	return len(set)
}

func (set Set[T]) Elements() []T {
	elements := make([]T, 0, len(set))
	for element := range set {
		elements = append(elements, element)
	}
	return elements
}
