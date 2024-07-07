package main

import (
	"fmt"
)

func main() {
	// String manipulations
	str := "hello world"
	fmt.Println("Original string:", str)

	reversedStr := ReverseString(str)
	fmt.Println("Reversed string:", reversedStr)

	insertedStr := InsertString(str, "beautiful ", 6)
	fmt.Println("String after insertion:", insertedStr)

	removedStr := RemoveSubstring(str, "world")
	fmt.Println("String after removal:", removedStr)

	searchIndex := SearchString(str, "world")
	fmt.Println("Index of 'world':", searchIndex)

	// Array manipulations
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", arr)

	reversedArr := ReverseArray(arr)
	fmt.Println("Reversed array:", reversedArr)

	insertedArr := InsertElement(arr, 10, 2)
	fmt.Println("Array after insertion:", insertedArr)

	removedArr := RemoveElement(arr, 3)
	fmt.Println("Array after removal:", removedArr)

	searchElemIndex := SearchElement(arr, 4)
	fmt.Println("Index of element 4:", searchElemIndex)
}
