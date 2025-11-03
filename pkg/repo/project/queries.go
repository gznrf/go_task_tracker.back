package r_projects

import (
	"fmt"

	"github.com/gznrf/go_task_tracker.back.git/app"
)

var (
	createQuery      = fmt.Sprintf(`INSERT INTO %s (owner_id, name, description) VALUES ($1, $2, $3) RETURNING id`, app.ProjectTable)
	getQuery         = fmt.Sprintf(`SELECT * FROM %s`, app.ProjectTable)
	getByUserIdQuery = fmt.Sprintf(`SELECT * FROM %s WHERE owner_id = $1`, app.ProjectTable)
	getByIdQuery     = fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, app.ProjectTable)
	updateQuery      = fmt.Sprintf(`UPDATE %s SET owner_id = $1, name = $2, description = $3 WHERE id = $4 RETURNING id`, app.ProjectTable)
	deleteQuery      = fmt.Sprintf(`DELETE FROM %s WHERE id = $1 RETURNING id`, app.ProjectTable)
)
