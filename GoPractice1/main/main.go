package main

import ( 
	"fmt"
	"log"

	"github.com/server"
)

func main() { 
	
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := server.Hellos(names)

	if err != nil { 
		log.Fatal(err)
	}

	fmt.Println(messages)
}