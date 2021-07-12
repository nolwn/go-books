package main

import (
	"fmt"
	"os"

	"github.com/nolwn/go-books/books"
	"github.com/nolwn/go-books/data"
)

func main() {
	args := os.Args
	bks := books.New()
	db, err := data.New()

	if err != nil {
		fmt.Printf("couldn't initialize database: %s", err)
	}

	db.Initialize() // create tables if they don't exist

	if len(args) == 1 { // a command should be passed to the program
		fmt.Print("No command found.")
	}

	command := args[1]
	runCommand(command, args, bks)
}
