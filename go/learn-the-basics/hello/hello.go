package main

import (
	"fmt";
	"log";
	"learn-the-basics/greetings";
);

func main() {
	log.SetPrefix("Greetings Package: ")
	message, err := greetings.Hello("Dylan");

	if (err != nil) {
		log.Fatal(err);
	}

	fmt.Println(message);
}
