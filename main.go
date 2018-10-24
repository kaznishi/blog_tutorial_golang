package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"net/url"
	"os"
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
	repository := NewRepository(dbConn)
	controller := NewController(repository)

	m := mux.NewRouter()
	m.HandleFunc("/", controller.Index).Methods("GET")

//	router := config.Router()
	http.ListenAndServe(viper.GetString("server.address"), m)
}

type Repository struct {
	DB *sql.DB
}

func NewRepository(dbConn *sql.DB) Repository {
	return Repository{
		DB: dbConn,
	}
}

type Controller struct {
	Repository Repository
}

func NewController(repository Repository) Controller {
	return Controller{
		Repository: repository,
	}
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "一覧ページ")
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