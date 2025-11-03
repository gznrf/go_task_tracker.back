package r_column

import (
	"fmt"

	"github.com/gznrf/go_task_tracker.back.git/app"
)

var (
	createQuery       = fmt.Sprintf(`INSERT INTO %s (board_id, name) VALUES ($1, $2) RETURNING id`, app.ColumnTable)
	getQuery          = fmt.Sprintf(`SELECT * FROM %s`, app.ColumnTable)
	getByBoardIdQuery = fmt.Sprintf(`SELECT * FROM %s WHERE board_id = $1`, app.ColumnTable)
	getByIdQuery      = fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, app.ColumnTable)
	updateQuery       = fmt.Sprintf(`UPDATE %s SET board_id = $2, name = $1 WHERE id = $3`, app.ColumnTable)
	deleteQuery       = fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, app.ColumnTable)
)
