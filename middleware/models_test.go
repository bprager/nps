package nps

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestShouldReturnAllUsers(t *testing.T) {
	// Creates sqlmock database connection and a mock to manage expectations.
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Closes the database and prevents new queries from starting.
	defer db.Close()

	// Switch to mock db
	dbOld := DB
	DB = sqlx.NewDb(db, "sqlMock")
	defer func() { DB = dbOld }()

	// Here we are creating rows in our mocked database.
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "nick_name"}).
		AddRow(1, "First1", "Last1", "Nick1").
		AddRow(2, "First2", "Last2", "Nick2")

		// This is most important part in our test. Here, literally, we are altering SQL query from
		// MenuByNameAndLanguage function and replacing result with our expected result.
	mock.ExpectQuery("^SELECT (.+) FROM users*").
		WillReturnRows(rows)

	ctx := context.TODO()

	r := new(Resolver)
	allUsers, err := r.Query().AllUsers(ctx)

	t.Logf("%v", allUsers)
}
