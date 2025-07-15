package memberships

import (
	"database/sql"
	"log"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	rows, err := db.Query("select email from users")
	if err != nil {
		log.Println("error query ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var email string

		if err = rows.Scan(&email); err != nil {
			log.Println("error scan ", err)
		}

		log.Println(email)
	}

	return &repository{db: db}
}
