package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

/*
	When defining a new variable you can define the variable in 2 ways

	var foo string = "example one"
	or
	bar := "example two"

	in the first example you need to define the type that you are assigning to the variable
	in the second example the type of the variable will be infered by your LSP
*/

/*
	Adding error handling to the GO function

	Adding an error in the return of your function is common practice in GO
	This allows for the implicit error handling in code

	Any GO function can have more than one return
*/

func Hello(name string) (string, error) {
	if (name == "") {
		return "", errors.New("No name given")
	}

	message := fmt.Sprintf(randomFormat(), name);
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v welcome!",
		"Great to see you %v",
		"Hail, %v! Well met!",
	};

	return formats[rand.Intn(len(formats))]
}

/*
	A map in GO is the equivalent of a Dictionary in JS

	you define the type of the indexing value in the [] and then that is followed by the type you expect to be stored

	eg. map[string]string or map[string]int


	From the GO LSP doc
	- The make built-in function allocates and initializes an object of type 
		slice, map, or chan (only). Like new, the first argument is a type, not a
		value. Unlike new, make's return type is the same as the type of its
		argument, not a pointer to it.

	When allocating a map you can specify a size to allocate the memory to eg
	`make(map[string]int, 3)`
	will create a map of size 3

*/

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name);
		if err != nil {
			return nil, err;
		}
		messages[name] = message;
	}

	return messages, nil
}
