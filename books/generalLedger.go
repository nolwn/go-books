package books

import (
	"errors"

	"github.com/nolwn/go-books/data"
)

type account struct {
	Code int    `sql:"code"`
	Name string `sql:"name"`
}

type generalLedgerAccounts struct {
	accounts map[int]account
}

// AddAccount adds a new account to the general ledger. Both the code and the name must
// be unique.
func (g *generalLedgerAccounts) AddAccount(code int, name string) error {
	db, err := data.New()
	if err != nil {
		return err
	}

	_, exists := g.accounts[code]

	if exists {
		return errors.New("account with that code already exists")
	}

	acc := account{code, name}

	err = db.Insert(data.GeneralLedgerTable, acc)
	if err != nil {
		return err
	}

	g.accounts[code] = acc

	return nil
}
