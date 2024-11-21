package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // PostgreSQL 드라이버 임포트
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:1234/simple_bank?sslmode=disable"
)

var (
	testQueries *Queries
	testStore   Store
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(conn)
	testStore = NewStore(conn)

	code := m.Run()
	// 종료 시 데이터베이스 연결 닫기
	conn.Close()

	// 테스트 코드 종료
	os.Exit(code)
}
