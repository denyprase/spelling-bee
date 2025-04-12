// db wrapper
package models

import "database/sql"

type DB struct {
	Conn *sql.DB
}
