package books

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nolwn/go-books/data"
)

// Books contains the general ledger and open journal entries. It should be created
// with New().
type Books struct {
	generalLedgerAccounts
	journalEntries []journalEntry
}

type journalEntry struct {
	AccountDebited  int     `sql:"debitAccount"`
	AccountCredited int     `sql:"creditAccount"`
	Amount          float64 `sql:"amount"`
	Date            string  `sql:"date"`
	Memo            string  `sql:"memo"`
}

// New returns a new Books object.
func New() Books {
	return Books{generalLedgerAccounts{make(map[int]account)}, []journalEntry{}}
}

// AddEntry adds a new journal entry. An error will be returned if the dateStr is invalid
// or if the credit or debit accounts don't exist.
func (b *Books) AddEntry(
	dateString string,
	debitAccount int,
	creditAccount int,
	amount float64,
	memo string,
) (err error) {
	db, err := data.New()
	badDateMsg := "maformed date"

	// passing an error pointer to reduce the number of error checks
	// b.checkAccount(debitAccount, &err)
	// b.checkAccount(creditAccount, &err)

	if err != nil {
		return
	}
	month, day, year, err := getMonthDayYear(dateString)

	if err != nil {
		return errors.New(badDateMsg)
	}

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	entry := journalEntry{
		debitAccount,
		creditAccount,
		amount,
		date.Local().Format(time.RFC3339),
		memo,
	}

	err = db.Insert(data.JournalEntriesTable, entry)
	if err != nil {
		return err
	}

	b.journalEntries = append(
		b.journalEntries,
		entry,
	)

	return nil
}

// checkAccount takes an error pointer and will set it if the account doesn't exist.
// TODO: this function needs to check against the database before it can be turned back on
func (b *Books) checkAccount(account int, err *error) {
	if _, ok := b.accounts[account]; !ok {
		*err = fmt.Errorf(
			"the account code %d does not appear in the General Ledger",
			account,
		)
	}
}

// getMonthDayYear takes a date in the format mm/dd/yy and returns a month, day and year
// (as integers). If the date string is not valid, an error will be returned.
func getMonthDayYear(dateString string) (month, day, year int, err error) {
	monDayYear := strings.Split(dateString, "/") // expected date format: mm/dd/yy

	if len(monDayYear) != 3 {
		err = errors.New("date format isn't correct")
		return
	}

	asInts := [3]int{}

	for i, part := range monDayYear {
		var asInt int
		asInt, err = strconv.Atoi(part)

		if err != nil {
			return
		}

		asInts[i] = asInt
	}

	month = asInts[0]
	day = asInts[1]
	year = asInts[2]

	if month > 12 || month < 1 {
		err = errors.New("month out of range")
	}

	return
}
