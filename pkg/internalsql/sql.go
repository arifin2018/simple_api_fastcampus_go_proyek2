package internalsql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(drivername string, dataDBsourcename string) (*sql.DB, error) {
	db, err := sql.Open(drivername, dataDBsourcename)
	if err != nil {
		log.Fatalf("error connection to database %+v", err)
		return nil, err
	}

	return db, nil
}
