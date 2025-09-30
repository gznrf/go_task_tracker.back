package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/gznrf/go_task_tracker.back.git/app"
	"github.com/gznrf/go_task_tracker.back.git/pkg/handler"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initDatabase() (*sqlx.DB, error) {

	db, err := repo.NewPostgresDB(repo.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func initHandler(db *sqlx.DB) *handler.Handler {
	r := repo.NewRepository(db)
	s := service.NewService(r)
	h := handler.NewHandler(s)

	return h
}

func main() {

	db, err := initDatabase()
	if err != nil {
		panic(errors.New("failed to connect to database"))
	}
	h := initHandler(db)

	srv := new(app.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), h.InitRoutes()); err != nil {
			err := fmt.Errorf("Ошибка запуска сервера %s", err.Error())
			panic(err)
		}
	}()

}
