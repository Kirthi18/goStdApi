package main

import (

	"net/http"

	"students/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// var DB *gorm.DB

// func Connect() {
// 	connection, err := gorm.Open(mysql.Open("root:QSGGSQ139@tcp(127.0.0.1:3306)/StudentDB"), &gorm.Config{})

// 	if err != nil {
// 		panic("could not connect to the database")
// 	}

// 	DB = connection

// 	connection.AutoMigrate(&entities.Users{})
// }

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/students", controllers.GetStudents).Methods("GET")
	r.HandleFunc("/students/{studentId}", controllers.GetStudent).Methods("GET")
	r.HandleFunc("/students", controllers.CreateStudents).Methods("POST")
	r.HandleFunc("/students/{studentId}", controllers.UpdateStudents).Methods("PUT")
	r.HandleFunc("/students/{studentId}", controllers.DeleteStudents).Methods("DELETE")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/user", controllers.LoggedUser).Methods("GET")
	r.HandleFunc("/logout", controllers.LogOut).Methods("POST")
	r.HandleFunc("/Register", controllers.RegisterUser).Methods("POST")
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://18.206.139.96:9070"})
	cred := handlers.AllowCredentials()

	http.ListenAndServe(":9070", handlers.CORS(headers, methods, origins, cred)(r))
	
}


