package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"students.com/api/config"
	"students.com/api/controllers"
)

func main() {
	fmt.Println(config.Config)
	r := mux.NewRouter()

	r.HandleFunc("/students", controllers.GetStudents).Methods("GET")
	r.HandleFunc("/students/{studentId}", controllers.GetStudent).Methods("GET")
	r.HandleFunc("/students", controllers.CreateStudents).Methods("POST")
	r.HandleFunc("/students/{studentId}", controllers.UpdateStudents).Methods("PUT")
	r.HandleFunc("/students/{studentId}", controllers.DeleteStudents).Methods("DELETE")
	r.HandleFunc("/login", controllers.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/user", controllers.LoggedUser).Methods("GET")
	r.HandleFunc("/logout", controllers.LogOut).Methods("POST")
	r.HandleFunc("/file/{studentId}", controllers.DownloadFile).Methods("GET")
	// r.HandleFunc("/refresh", controllers.Refresh).Methods("POST")
	// headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	// methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	// origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://"})
	cred := handlers.AllowCredentials()

	http.ListenAndServe(":9070", handlers.CORS(headers, methods, origins, cred)(r))
}

// http://localhost:4200/
