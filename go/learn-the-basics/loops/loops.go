package main

import (
	"fmt"
)

func main() {

	fmt.Println("For loop");
	for i := 0; i < 10; i++ {
		fmt.Println(i);
	}

	fmt.Println("Range loop");
	for i := range 10 {
		fmt.Println(i);
	}

	fmt.Println("Array range loop");
	array := []string{"one", "two", "tree", "fowa", "five", "sex", "seven", "eight", "nine", "ten"};
	for index, value := range array {
		fmt.Printf("index: %d || value: %s\n", index, value);
	}

	fmt.Println("Map range loop");
	tempMap := map[string]int{
		"test": 1,
		"tost": 2,
		"twist": 3,
		"antler": 4,
	}
	for key, value := range tempMap {
		fmt.Printf("key: %s || value: %d\n", key, value);
	}
}
