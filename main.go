package main

import (
	"fmt"
	"log"
	"database/sql"
	"net/http"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/kaznishi/blog_tutorial_golang/controller"
	"github.com/kaznishi/blog_tutorial_golang/model/repository"
	"github.com/kaznishi/blog_tutorial_golang/service"
)

func init() {
	viper.SetConfigFile(`config.yaml`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbConn := initDB()

	repositoryManager := repository.NewRepositoryManager(dbConn)
	articleRepository := repositoryManager.NewArticleRepository()
	articleService := service.NewArticleService(articleRepository)
	articleController := controller.NewArticleController(articleService)

	m := mux.NewRouter()
	m.HandleFunc("/", articleController.Index).Methods("GET")
	m.HandleFunc("/view/{id:[0-9]+}", articleController.View).Methods("GET")

	http.ListenAndServe(viper.GetString("server.address"), m)
}


func initDB() *sql.DB {

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Tokyo")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return dbConn
}