package postgres_backend

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/straumur/straumur"
	"testing"
	"time"
)

func TestDB(t *testing.T) {

	const connection = "dbname=teststream host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		t.Error("Error:", err)
		return
	}

	//truncate table
	clear := func() {
		_, err := db.Exec(`TRUNCATE TABLE event RESTART IDENTITY;`)
		if err != nil {
			t.Fatal(err)
		}
	}

	_, err = db.Exec(`drop table if exists migration_info, event;`)
	if err != nil {
		t.Fatal("Error:", err)
	}

	// With migrations applied
	_, err = NewPostgresDataSource(connection)
	if err != nil {
		t.Fatal("PostgresDataSource has error:", err)
	}

	// With no migrations applied
	p, err := NewPostgresDataSource(connection)
	if err != nil {
		t.Fatal("PostgresDataSource has error:", err)
	}

	straumur.RunDataBackendSuite(t, p, clear)
}

func TestDuplicateInsertion(t *testing.T) {

	const connection = "dbname=teststream host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		t.Error("Error:", err)
		return
	}

	//truncate table
	clear := func() {
		_, err := db.Exec(`TRUNCATE TABLE event RESTART IDENTITY;`)
		if err != nil {
			t.Fatal(err)
		}
	}

	clear()

	now := time.Now()

	e := straumur.NewEvent(
		"foo.bar",
		nil,
		nil,
		"My event",
		3,
		"mysystem",
		[]string{"ns/foo", "ns/moo"},
		nil,
		[]string{"someone"},
		nil)
	e.Created = now

	e2 := straumur.NewEvent(
		"foo.bar",
		nil,
		nil,
		"My event",
		3,
		"mysystem",
		[]string{"ns/foo", "ns/moo"},
		nil,
		[]string{"someone"},
		nil)
	e2.Created = now

	e3 := straumur.NewEvent(
		"foo.bar",
		nil,
		nil,
		"My event",
		3,
		"mysystem",
		[]string{"ns/foo", "ns/moo"},
		nil,
		[]string{"someone"},
		nil)

	p, err := NewPostgresDataSource(connection)
	if err != nil {
		t.Fatal("PostgresDataSource has error:", err)
	}

	p.Save(e)
	p.Save(e2)
	p.Save(e3)
}
