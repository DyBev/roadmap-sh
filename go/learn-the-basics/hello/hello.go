package main

import (
	"fmt";
	"log";
	"greetings";
);

func main() {
	log.SetPrefix("Greetings Package: ")
	message, err := greetings.Hello("Michael");

	if (err != nil) {
		log.Fatal(err);
	}

	fmt.Println(message);

	names := []string{"Michael", "Gladys", "James", "Peter"}
	messages, err := greetings.Hellos(names);

	if (err != nil) {
		log.Fatal(err);
	}

	fmt.Println(messages);
}
