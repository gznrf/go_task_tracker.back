package r_user

import (
	"fmt"

	"github.com/gznrf/go_task_tracker.back.git/app"
)

var (
	getQuery     = fmt.Sprintf(`SELECT * FROM %s`, app.UserTable)
	getByIdQuery = fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, app.UserTable)
	updateQuery  = fmt.Sprintf(`UPDATE %s SET name = $1, email = $2, password_hash = $3 WHERE id = $4`, app.UserTable)
	deleteQuery  = fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, app.UserTable)
)
