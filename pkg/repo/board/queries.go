package r_board

import (
	"fmt"

	"github.com/gznrf/go_task_tracker.back.git/app"
)

var (
	createQuery         = fmt.Sprintf(`INSERT INTO %s (project_id, name) VALUES ($1, $2) RETURNING id`, app.BoardTable)
	getQuery            = fmt.Sprintf(`SELECT * FROM %s`, app.BoardTable)
	getByIdQuery        = fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, app.BoardTable)
	getByProjectIdQuery = fmt.Sprintf(`SELECT * FROM %s WHERE project_id = $1`, app.BoardTable)
	updateQuery         = fmt.Sprintf(`UPDATE %s SET project_id = $1, name = $2 WHERE id = $3 RETURNING id`, app.BoardTable)
	deleteQuery         = fmt.Sprintf(`DELETE FROM %s WHERE id = $1 RETURNING id`, app.BoardTable)
)
