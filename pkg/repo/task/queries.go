package r_task

import (
	"fmt"

	"github.com/gznrf/go_task_tracker.back.git/app"
)

var (
	createQuery        = fmt.Sprintf(`INSERT INTO %s (column_id, creator_id, executor_id, number, name, description) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`, app.TaskTable)
	getQuery           = fmt.Sprintf(`SELECT * FROM %s`, app.TaskTable)
	getByColumnIdQuery = fmt.Sprintf(`SELECT * FROM %s WHERE column_id = $1`, app.TaskTable)
	getByIdQuery       = fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, app.TaskTable)
	updateQuery        = fmt.Sprintf(`UPDATE %s SET column_id = $1, creator_id = $2, executor_id = $3, number = $4, name = $5, description = $6 WHERE id = $7 RETURNING id`, app.TaskTable)
	deleteQuery        = fmt.Sprintf(`DELETE FROM %s WHERE id = $1 RETURNING id`, app.TaskTable)
)
