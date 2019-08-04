package nps

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestShouldReturnAllTags(t *testing.T) {
	// Creates sqlmock database connection and a mock to manage expectations.
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Closes the database and prevents new queries from starting.
	defer db.Close()

	// Switch to mock db
	DB = sqlx.NewDb(db, "sqlMock")

	// Here we are creating rows in our mocked database.
	rows := sqlmock.NewRows([]string{"id", "name", "attribute", "number", "timestamp"}).
		AddRow(1, "male", nil, nil, nil).
		AddRow(2, "female", nil, nil, nil).
		AddRow(3, "complete", "nonsense", 42, "2014-02-12 16:04:15.879588-08")
		// This is most important part in our test. Here, literally, we are altering SQL query from
		// MenuByNameAndLanguage function and replacing result with our expected result.
	mock.ExpectQuery("^SELECT (.+) FROM tags.*").
		WillReturnRows(rows)

	r := new(Resolver)
	allUsers, err := r.Query().AllTags(context.TODO())
	if len(allUsers) != 3 {
		t.Errorf("Didn't get expected results, got: %d rows, want: %d.", len(allUsers), 3)
	}
}

func TestShouldReturnAllUsers(t *testing.T) {
	// Creates sqlmock database connection and a mock to manage expectations.
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Closes the database and prevents new queries from starting.
	defer db.Close()

	// Switch to mock db
	DB = sqlx.NewDb(db, "sqlMock")

	// Here we are creating rows in our mocked database.
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "nick_name"}).
		AddRow(1, "First1", "Last1", "Nick1").
		AddRow(2, "First2", "Last2", "Nick2")

		// This is most important part in our test. Here, literally, we are altering SQL query from
		// MenuByNameAndLanguage function and replacing result with our expected result.
	mock.ExpectQuery("^SELECT (.+) FROM users.*").
		WillReturnRows(rows)

	r := new(Resolver)
	allUsers, err := r.Query().AllUsers(context.TODO())
	if len(allUsers) != 2 {
		t.Errorf("Didn't get expected results, got: %d rows, want: %d.", len(allUsers), 2)
	}
}
