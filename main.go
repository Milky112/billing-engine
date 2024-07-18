package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	handler "amartha.com/billing/handler"
	logger "amartha.com/billing/log"
	model "amartha.com/billing/model"
	repository "amartha.com/billing/repository"
	usecase "amartha.com/billing/usecase"
	"gopkg.in/yaml.v3"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	var db *sql.DB

	routes(router, db)
	defer db.Close()

	logger.GenerateLog()
	logger.CommonLog.Println("Server is Running")
	server := http.Server{Addr: ":8080", Handler: router}
	err := server.ListenAndServe()
	if err != nil {
		logger.CommonLog.Print("Error Start Server ", err)
	}

	logger.CommonLog.Print("Shuting down on progress...")

}

func routes(router *httprouter.Router, sqlDB *sql.DB) {
	var config model.Config

	getConfig(&config)
	dbConfig := config.Database

	//Convert DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.DatabaseUser, dbConfig.DatabasePassword, dbConfig.DatabaseConn, dbConfig.DatabasePort, dbConfig.DatabaseUser)

	repo := repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dsn,
	})

	uc := usecase.NewUsecase(usecase.NewUsecaseOptions{
		Repository: repo,
	})

	h := handler.NewServer(handler.NewServerOptions{
		Usecase: uc,
	})

	sqlDB = repo.Db

	router.GET("/GetOutstanding", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		h.GetSchedule(w, r, p)
	})
	router.POST("/AcquireLoan", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		h.AcquireLoan(w, r, p)
	})
	router.POST("/MakePayment", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		h.AcquireLoan(w, r, p)
	})
	router.POST("/UserInfo", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		h.AcquireLoan(w, r, p)
	})
}

func getConfig(config *model.Config) {
	data, err := os.ReadFile("config/billing-config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
		return
	}
}
