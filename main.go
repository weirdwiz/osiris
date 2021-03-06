package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/weirdwiz/osiris/app"
	"github.com/weirdwiz/osiris/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/teach", controllers.TeacherNew).Methods("POST")
	router.HandleFunc("/api/skill/all", controllers.GetSkills).Methods("GET")
	router.HandleFunc("/api/skill/{id}/teachers", controllers.GetAllTeachers).Methods("GET")
	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
