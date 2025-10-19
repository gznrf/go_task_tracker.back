package r_utils

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func CreateConnection(db *sqlx.DB, tableName string, fromId, toId int64) error {
	query := fmt.Sprintf("INSERT INTO %s (fromId, toId) VALUES ($1, $2)", tableName)

	row := db.QueryRow(query, fromId, toId)
	if err := row.Scan(); err != nil {
		return err
	}

	return nil
}

// RemoveConnection idType - to_id/from_id
func RemoveConnection(db *sqlx.DB, tableName string, id int64, idType string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE %s = $1`, idType, tableName)
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
