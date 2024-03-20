package array

import "testing"

func addElements(times int, array *Array[int]) {
	for i := 0; i < times; i++ {
		array.Add(i)
	}
}

func Test_array_should_be_same_size_as_defined(t *testing.T) {
	array := Array[int]{Size: 10}
	array.init()

	if array.Size != 10 {
		t.Errorf("Array size should be 10, but got %d", array.Size)
	}
}

func Test_array_should_be_double_in_size_when_adding_element(t *testing.T) {
	array := Array[int]{Size: 10}
	array.init()

	addElements(11, &array)

	if array.Size != 20 {
		t.Errorf("Array size should be 20, but got %d", array.Size)
	}
}

func Test_get_element_from_array(t *testing.T) {
	array := Array[int]{Size: 10}
	array.init()

	array.Add(1)

	if array.Get(0) != 1 && array.Size != 10 {
		t.Errorf("Array should have 1, but got %d", array.Get(0))
	}
}
