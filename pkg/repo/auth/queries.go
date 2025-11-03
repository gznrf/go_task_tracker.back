package r_auth

import (
	"fmt"

	"github.com/gznrf/go_task_tracker.back.git/app"
)

var (
	registerQuery = fmt.Sprintf(`INSERT INTO %s (name, email, password_hash) VALUES ($1, $2, $3) RETURNING id`, app.UserTable)
	loginQuery    = fmt.Sprintf(`SELECT * FROM %s WHERE email = $1 AND password_hash = $2`, app.UserTable)
)
