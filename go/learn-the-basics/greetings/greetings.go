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
