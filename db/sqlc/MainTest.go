package db

import (
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:1234/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(t *testing.M) {

}
