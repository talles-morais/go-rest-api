package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/talles-morais/go-rest-api/controllers"
	"github.com/talles-morais/go-rest-api/middleware"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.Use(middleware.HeaderSetupMiddleware)
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	router.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("Post")
	router.HandleFunc("/api/personalities/{id}", controllers.OnePersonality).Methods("Get")
	router.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	router.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
