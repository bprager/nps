package nps

import (
	"context"
	"testing"
	"time"

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
	rows := sqlmock.NewRows([]string{"id", "name", "attribute", "number", "timestamp", "count"}).
		AddRow(1, "male", nil, nil, nil, 3).
		AddRow(2, "female", nil, nil, nil, 3).
		AddRow(3, "complete", "nonsense", 42, "2014-02-12 16:04:15.879588-08", 3)
		// This is most important part in our test. Here, literally, we are altering SQL query from
		// MenuByNameAndLanguage function and replacing result with our expected result.
	mock.ExpectQuery("^SELECT (.+) FROM tags.*").
		WillReturnRows(rows)

	r := new(Resolver)
	allTags, err := r.Query().AllTags(context.TODO(), 100, 0)
	if len(allTags.Tags) != 3 {
		t.Errorf("Didn't get expected results, got: %d rows, want: %d.", len(allTags.Tags), 3)
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
	rows := sqlmock.NewRows([]string{"id", "email", "first_name", "last_name", "nick_name", "orgs", "tags", "categories", "count"}).
		AddRow(1, "first@email", "First1", "Last1", "Nick1", 1, 2, 3, 3).
		AddRow(2, "second@email", nil, nil, nil, nil, nil, nil, 3).
		AddRow(2, nil, "First3", "Last3", "Nick3", nil, nil, nil, 3)

		// This is most important part in our test. Here, literally, we are altering SQL query from
		// MenuByNameAndLanguage function and replacing result with our expected result.
	mock.ExpectQuery("^SELECT (.+) FROM users.*").
		WillReturnRows(rows)

	r := new(Resolver)
	allUsers, err := r.Query().AllUsers(context.TODO(), 100, 0)
	if len(allUsers.Users) != 3 {
		t.Errorf("Didn't get expected results, got: %d rows, want: %d.", len(allUsers.Users), 3)
	}
}

func TestAddTag(t *testing.T) {
	// Creates sqlmock database connection and a mock to manage expectations.
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Closes the database and prevents new queries from starting.
	defer db.Close()

	// Switch to mock db
	DB = sqlx.NewDb(db, "sqlMock")

	mock.ExpectExec("INSERT INTO tags").
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := new(Resolver)

	attribute := "attribute"
	number := 123
	timestamp := time.Now().String()

	result, err := r.Mutation().AddTag(context.TODO(), "name", &attribute, &number, &timestamp)
	if !result {
		t.Errorf("Didn't get expected results, got: %t , want: %t.", result, true)
	}

}
