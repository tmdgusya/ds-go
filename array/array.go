package array

type Array[T any] struct {
	Size    int
	room    []T
	pointer int
}

// init array with given size that a has
func (a *Array[T]) init() {
	a.room = make([]T, a.Size)
}

// add element into the array
// if there is no room for storing array, then array is going to grow automatically.
// the growth size is considered to be twice current size
func (a *Array[T]) Add(ele T) {
	// if there is no room for storing element, then grow array to twice size
	if len(a.room) == a.Size {
		a.room = append(a.room, make([]T, a.Size)...)
		a.Size *= 2
	}

	a.room[a.pointer] = ele
	// grow pointer after adding element
	a.pointer++
}

// Get element from array by index
func (a *Array[T]) Get(index int) T {
	return a.room[index]
}
