package main

// ReverseArray reverses the input array.
func ReverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// InsertElement inserts an element into an array at a specified position.
func InsertElement(arr []int, element, index int) []int {
	if index > len(arr) {
		index = len(arr)
	}
	arr = append(arr[:index], append([]int{element}, arr[index:]...)...)
	return arr
}

// SearchElement searches for the first occurrence of an element in an array.
func SearchElement(arr []int, element int) int {
	for i, v := range arr {
		if v == element {
			return i
		}
	}
	return -1
}

// RemoveElement removes the first occurrence of an element from an array.
func RemoveElement(arr []int, element int) []int {
	index := SearchElement(arr, element)
	if index == -1 {
		return arr
	}
	return append(arr[:index], arr[index+1:]...)
}
