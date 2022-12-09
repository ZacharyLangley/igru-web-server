package gardens_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/cockroachdb/copyist"
)

const (
	DriverName     = "postgres"
	DockerArgs     = "-p 26888:26257 cockroachdb/cockroach:v20.2.4 start-single-node --insecure"
	DataSourceName = "postgresql://root@localhost:26888?sslmode=disable"
)

const resetScript = `
DROP TABLE IF EXISTS gardens;
CREATE TABLE gardens;
`

func resetDB() {
	db, err := sql.Open(DriverName, DataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if _, err := db.Exec(resetScript); err != nil {
		panic(err)
	}
}

func init() {
	copyist.Register("postgres")
	copyist.SetSessionInit(resetDB)
}

func TestMain(m *testing.M) {
	log.Println("Setting Up...")
	db, _ := sql.Open(DriverName, DataSourceName)
	log.Println("Opening Database...")
	defer db.Close()
	log.Println("Running Tests...")
	exitVal := m.Run()
	log.Println("Testing Completed...")
	os.Exit(exitVal)
}
