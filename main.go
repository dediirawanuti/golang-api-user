package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/golang-api-user/driver"
	"github.com/golang-api-user/middlewares"
	"github.com/golang-api-user/user/handler"
	"github.com/golang-api-user/user/repo"
	"github.com/golang-api-user/user/usecase"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	var dir string

	flag.StringVar(&dir, "dir", viper.GetString("file.path"), "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	db := driver.Config()

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	router := mux.NewRouter().StrictSlash(true)
	router.Use(middlewares.Logging)
	router.Use(middlewares.SetMiddlewareJSON)
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir(dir))))

	userRepo := repo.CreateRepo(db)
	userUsecase := usecase.CreateUsecase(userRepo)
	handler.CreateHandler(router, userUsecase)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to my restApi"))
	})

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"})

	log.Println("Server start at http://localhost:8080")
	logrus.Fatal(http.ListenAndServe(viper.GetString("server.port"), handlers.CORS(headersOk, originsOk, methodsOk)(router)))
}
