package main;

import (
	"fmt";
)

func main() {
	var correct bool = true;

	if correct {
		fmt.Println("Correct is ture");
	}

	if !correct {
		fmt.Println("Correct is false");
	}

	tempMap := map[string]string{"hello": "hello", "world":"world"};
	if value, ok := tempMap["hello"]; ok {
		fmt.Printf("hello is in map with value: %s\n", value);
	}

	if _, ok := tempMap["goodbye"]; !ok {
		fmt.Println("goodbye is not in map");
	}

	array := []string{"hello", "world", "goodbye"}
	for index, value := range array {
		switch value {
		case "hello":
			fmt.Printf("index: %d || value: %s\n", index, value);
		case "world":
			fmt.Printf("index: %d || value: %s\n", index, value);
		default:
			fmt.Printf("Default switch statement || value: %s\n", value);
		}
	}

}
