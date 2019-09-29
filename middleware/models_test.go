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
	rows := sqlmock.NewRows([]string{"totalcount",
		"userid", "email", "firstname", "lastname", "nickname",
		"tagid", "tagname", "attribute", "number", "timestamp",
		"orgid", "orgname", "catid", "catname", "parent"}).
		AddRow(2, 1, "bernd@prager.ws", "Bernd", "Prager", nil, 2, "language", "German", nil, nil, 1, "Home", 6, "Male", 1).
		AddRow(2, 1, "bernd@prager.ws", "Bernd", "Prager", nil, 3, "creative", nil, nil, nil, 1, "Home", 6, "Male", 1).
		AddRow(2, 2, "tene@prager.ws", "Tené", "Closson-Prager", nil, 2, "language", "English", nil, nil, 1, "Home", 5, "Female", 1).
		AddRow(2, 2, "tene@prager.ws", "Tené", "Closson-Prager", nil, 3, "creative", nil, nil, nil, 1, "Home", 5, "Female", 1)

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
