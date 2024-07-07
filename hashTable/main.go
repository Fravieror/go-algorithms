package main

import "fmt"

func main() {
	// Create a map
	myMap := make(map[string]string)

	// Insert key-value pairs
	myMap["name"] = "John"
	myMap["age"] = "30"
	myMap["location"] = "New York"

	// Lookup
	fmt.Println(myMap["name"])    // Output: John
	fmt.Println(myMap["age"])     // Output: 30
	fmt.Println(myMap["address"]) // Output: (empty string, key does not exist)

	// Delete
	delete(myMap, "age")
	_, exists := myMap["age"]
	fmt.Println(exists) // Output: false
}
