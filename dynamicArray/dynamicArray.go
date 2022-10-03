package main

import (
	"fmt"
)

type DynamicArray struct {
	Array []int
	Len int // length user think array is
	Capacity int // Actual array size
}

func (arr *DynamicArray) Append(elt int) {
	// append to array if array has room for an element arr.len < arr.capacity
	if arr.Len < arr.Capacity {
		arr.Array[arr.Len] = elt
		arr.Len += 1
	// create array with double capacity
	} else {
		tempArray := arr.Array
		// create array with double capacity
		arr.Capacity = arr.Capacity*2
		arr.Array = make([]int, arr.Capacity)
		// loop to copy all elements in the new array
		for i := 0; i < len(tempArray); i++ {
			arr.Array[i] = tempArray[i]
		}
		// insert element
		arr.Array[arr.Len] = elt
		arr.Len += 1
	}
}

func (arr *DynamicArray) RemoveAt(rmIndex int) {
	if rmIndex >= arr.Len && rmIndex < 0 {
		fmt.Println("Error rmIndex out of range")
		//os.exit(3)
	}
	newArr := make([]int, arr.Capacity)
	i, j := 0, 0
	for j < len(arr.Array) {
		if j == rmIndex {
			i--
		} else {
			newArr[i] = arr.Array[j]
		}
		i++
		j++
	}
	arr.Array = newArr
	arr.Len--
}

func (arr DynamicArray) Print() {
	for i := 0; i < arr.Len; i++ {
		fmt.Println(arr.Array[i])
	}
}

func main() {
	fmt.Println("Hello")
	dynaArray := DynamicArray{
		Array: []int{0, 0},
		Len: 0,
		Capacity: 2,
	}
	dynaArray.Append(22)
	dynaArray.Append(11)
	dynaArray.Append(45)
	dynaArray.Append(116)
	dynaArray.RemoveAt(3)
	dynaArray.Print()
	fmt.Println(dynaArray.Len)
}