package data

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
)

const databaseName = "./data/books.db"
const driverName = "sqlite3"

const structTag = "sql"

// Table names
const (
	GeneralLedgerTable  = "glAccounts"
	JournalEntriesTable = "journalEntries"
)

const (
	sqlCreateGeneralLedgerTable = `
		CREATE TABLE IF NOT EXISTS glAccounts (
			code INTEGER PRIMARY KEY,
			name TEXT UNIQUE
		);
	`

	sqlCreateJournalEntriesTable = `
		CREATE TABLE IF NOT EXISTS journalEntries (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			debitAccount INTEGER REFERENCES glAccounts (code),
			creditAccount INTEGER,
			amount REAL,
			date TEXT,
			memo TEXT
		);
	`

	sqlInsert = "INSERT INTO %s (%s) VALUES (%s);"
)

type database struct {
	*sql.DB
}

// newDatabase returns a new Database object with an open connection to the sqlite file.
// Returns an error if the databse cannot be opened.
func New() (database, error) {
	db, err := sql.Open(driverName, databaseName)
	if err != nil {
		return database{}, err
	}

	return database{db}, nil
}

// Insert takes a table name, and a struct and inserts them into the database. For the
// stuct to be insertable, its fields need to have tags. The tag name is sql so, if you
// had a database table for people and you wanted to insert a person you might have a
// struct that looks like this:
//
//	type Person struct {
//		Name     `sql="name"`
//		Birthday `sql="dob"`
//		Age
//	}
//
// This would insert a person into the database with Name as name and Birthday as dob. It
// would ignore Age since there is no tag there.
//
// It's the responsibility of the caller to make sure that what they are inserting fits
// into the table. If the field names or values don't fit into the database, then an error
// will be thrown.
func (d *database) Insert(table string, item interface{}) error {
	fmt.Print("Inserting item\n")
	// the vars below will be formatted into the insert string where it's missing table
	// name, field names and values (which should all be "?"s)
	var fieldsString string
	var valuesString string
	var insertSql string

	fields := make([]string, 0, 6)      // will hold all field names (by tag)
	values := make([]interface{}, 0, 6) // will hold fild values (same order as fields)

	itemV := reflect.ValueOf(item) // reflects about values
	itemT := itemV.Type()          // reflects about type

	if itemT.Kind() != reflect.Struct {
		return errors.New("item must be a stuct")
	}

	// iterate over struct values
	for i := 0; i < itemT.NumField(); i++ {

		// unexported fields will panic on Interface() so make sure we can use Interface()
		if itemV.Field(i).CanInterface() {
			field := itemT.Field(i)                // get field by index
			value := itemV.Field(i).Interface()    // get field value
			tag, ok := field.Tag.Lookup(structTag) // get field tag

			if ok {
				fields = append(fields, tag)
				values = append(values, value)
			}
		}
	}

	for i := 0; i < len(fields); i++ {
		if i != 0 {
			valuesString += ", "
			fieldsString += ", "
		}

		fieldsString += fields[i]
		valuesString += "?" // we will prepare values to guard against SQL injection
	}

	insertSql = fmt.Sprintf(sqlInsert, table, fieldsString, valuesString)

	fmt.Println(insertSql)

	stmt, err := d.Prepare(insertSql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(values...)

	if err != nil {
		return err
	}

	fmt.Print("Item inserted\n")

	return nil
}

// initialize creates the database and tables if they do not already exist.
func (d *database) Initialize() error {
	err := d.createTable(sqlCreateGeneralLedgerTable)
	if err != nil {
		return err
	}

	err = d.createTable(sqlCreateJournalEntriesTable)
	if err != nil {
		return err
	}

	return nil
}

func (d *database) createTable(sqlCreate string) error {
	fmt.Println(sqlCreate)
	stmt, err := d.Prepare(sqlCreate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
