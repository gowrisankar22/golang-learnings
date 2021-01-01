package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"user-auth/controller"
	"user-auth/middleware"
	"user-auth/models"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error in loading in env file")
	}
	dsn := "host=localhost user=" + os.Getenv("Name") + " password=" + os.Getenv("Password") + " dbname=" + os.Getenv("Database") + " port=" + os.Getenv("Port") + " sslmode=disable "
	us, err := models.NewUserService(dsn)

	if err != nil {
		panic(err)
	}

	us.DestructiveReset()

	userC := controller.NewUserController(us)
	requireUser := middleware.RequireUser{UserService: us}

	r := mux.NewRouter()

	r.HandleFunc("/signup", userC.Create).Methods("POST")
	r.HandleFunc("/login", userC.Login)
	r.HandleFunc("/home", userC.Home)
	r.HandleFunc("/hello", requireUser.Applyfn(userC.Hello))

	http.ListenAndServe(":8080", r)

	//http.HandleFunc("/", controller.Home)
	//http.ListenAndServe(":8080", nil)

}
