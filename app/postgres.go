package app

import (
	"fmt"
	"os/exec"

	"github.com/jmoiron/sqlx"
)

const (
	UserTable    = "user_info"
	ProjectTable = "project"
	//ProjectsUsersTable = "projects_users"
	BoardTable  = "board"
	ColumnTable = "column_info"
	TaskTable   = "task"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	databaseConnectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	db, err := sqlx.Open("postgres", databaseConnectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = createTables(databaseConnectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(connString string) error {
	cmd := exec.Command(
		"migrate",
		"-path", "./migration",
		"-database", connString,
		"-verbose",
		"up",
	)

	res, err := cmd.Output()
	if err != nil {
		return err
	}

	fmt.Println(string(res))
	return nil
}
