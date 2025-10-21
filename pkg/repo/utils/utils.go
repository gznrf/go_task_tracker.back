package r_utils

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func CreateConnection(db *sqlx.DB, tableName string, fromId, toId int64) (int64, error) {
	var createdConnectionId int64

	query := fmt.Sprintf(`INSERT INTO %s (from_id, to_id) VALUES ($1, $2) RETURNING id`, tableName)

	row := db.QueryRow(query, fromId, toId)
	if err := row.Scan(&createdConnectionId); err != nil {
		return 0, err
	}

	return createdConnectionId, nil
}

// RemoveConnection idType - to_id/from_id
func RemoveConnection(db *sqlx.DB, tableName string, id int64, idType string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE %s = $1`, tableName, idType)
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
