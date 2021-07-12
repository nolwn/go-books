package main

import (
	"fmt"
	"strconv"

	"github.com/nolwn/go-books/books"
)

// Convert an arg from a string to an int. Passing an error pointer means multiple
// arguments can be converted and the error only needs to be checked once when they are
// all done.
func argToInt(arg string, argError *error) int {
	argInt, err := strconv.Atoi(arg)

	if err != nil {
		*argError = err
	}

	return argInt
}

// Convert an arg from a string to a float64. Passing an error pointer means multiple
// arguments can be converted and the error only needs to be checked once when they are
// all done.
func argToFloat(arg string, argError *error) float64 {
	argFloat, err := strconv.ParseFloat(arg, 64)

	if err != nil {
		*argError = err
	}

	return argFloat
}

func runCommand(command string, args []string, bks books.Books) {
	switch command {
	// user wants to add a new GL account
	case "add-account":
		// after filename and command (0 and 1) we need a code and a name for the account
		if len(args) != 4 {
			fmt.Printf("provide an account code and a name\n")
			break
		}

		code, err := strconv.Atoi(args[2])
		name := args[3]

		if err != nil {
			fmt.Printf("invalid GL code: %s\n", args[2])
			break
		}

		err = bks.AddAccount(code, name)

		if err != nil {
			fmt.Printf("could not add account: %s\n", err)
		}

		fmt.Printf("GL %v", bks)

	// user wants to add a new journal entry
	case "add-journal-entry":
		// after filename and command (0 and 1) we need a date, debit account, credit
		// account, amount and a description of the transaction.
		if len(args) != 7 {
			fmt.Print("provide an a date (mm/dd/yyy), debit account, credit account, " +
				"amount, and memo\n")
			return
		}

		var argErr error

		dateStr := args[2]
		debitAccount := argToInt(args[3], &argErr)
		creditAccount := argToInt(args[4], &argErr)
		amount := argToFloat(args[5], &argErr)
		memo := args[6]

		if argErr != nil {
			fmt.Print(
				"debitAccount, and creditAccount must be integers and amount must be " +
					"a decimal number\n",
			)
		}

		err := bks.AddEntry(dateStr, debitAccount, creditAccount, amount, memo)
		if err != nil {
			fmt.Printf("could not create journal entry: %s", err)
		}
	}
}
