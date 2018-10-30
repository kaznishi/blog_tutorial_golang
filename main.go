package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/kaznishi/blog_tutorial_golang/controller"
	"github.com/kaznishi/blog_tutorial_golang/controller/middleware"
	"github.com/kaznishi/blog_tutorial_golang/model/repository"
	"github.com/kaznishi/blog_tutorial_golang/service"
	"github.com/spf13/viper"
	"gopkg.in/boj/redistore.v1"
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
	sessionStore := initSessionStore()

	dbConn := initDB()
	repositoryManager := repository.NewRepositoryManager(dbConn)
	articleRepository := repositoryManager.NewArticleRepository()
	userRepository := repositoryManager.NewUserRepository()

	articleService := service.NewArticleService(articleRepository)
	userService := service.NewUserService(userRepository)
	sessionService := service.NewSessionService(sessionStore,userRepository)

	articleController := controller.NewArticleController(articleService)
	adminController := controller.NewAdminController(articleService, userService)
	loginController := controller.NewLoginController(sessionService)

	smw := middleware.NewSessionMiddleware(sessionService)

	m := mux.NewRouter()
	m.HandleFunc("/", smw.SessionStart(articleController.Index)).Methods("GET")
	m.HandleFunc("/view/{id:[0-9]+}", smw.SessionStart(articleController.View)).Methods("GET")
	m.HandleFunc("/login", smw.SessionStart(loginController.Login))
	m.HandleFunc("/logout", smw.SessionStart(loginController.Logout))
	m.HandleFunc("/admin/", smw.SessionStart(smw.ForOnlyLoginUser(adminController.Index))).Methods("GET")
	m.HandleFunc("/admin/article/new", smw.SessionStart(smw.ForOnlyLoginUser(adminController.NewArticle)))
	m.HandleFunc("/admin/article/edit/{id:[0-9]+}", smw.SessionStart(smw.ForOnlyLoginUser(adminController.EditArticle)))
	m.HandleFunc("/admin/user/list", smw.SessionStart(smw.ForOnlyLoginUser(adminController.ListUser)))

	m.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(viper.GetString("server.address"), m)
}

func initSessionStore() sessions.Store {
	store, err := redistore.NewRediStore(10, "tcp", "redis:6379", "", []byte("a6b0e040989e6131daccca9290cb64a0444b52dfc3bf22b8b77f938542f79757"))
	if err != nil {
		panic(err)
	}
	return store
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
