package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gznrf/go_task_tracker.back.git/app"
	"github.com/gznrf/go_task_tracker.back.git/pkg/handler"
	"github.com/gznrf/go_task_tracker.back.git/pkg/repo"
	"github.com/gznrf/go_task_tracker.back.git/pkg/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initDatabase() (*sqlx.DB, error) {

	c := app.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	db, err := app.NewPostgresDB(c)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func initHandler(db *sqlx.DB) *handler.Handler {

	//auth
	r := repo.NewRepository(db)
	s := service.NewService(r)
	h := handler.NewHandler(s)

	return h
}

//dmitry.sv95@yandex.ru

func main() {
	err := initConfig()
	if err != nil {
		panic(err)
	}
	db, err := initDatabase()
	if err != nil {
		panic(err)
	}
	h := initHandler(db)
	srv := new(app.Server)
	go func() {
		if err := srv.Run(fmt.Sprintf(":"+viper.GetString("port")), *h.InitRoutes()); err != nil {
			err := fmt.Errorf("Server run error %s", err.Error())
			panic(err)
		}
	}()

	fmt.Printf("app started, API listenning on port :%s", viper.GetString("port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Println("app shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Printf("error on server shuting down: %s", err)
	}
	if err := db.Close(); err != nil {
		fmt.Printf("error on db connection closing: %s", err.Error())
	}

}
