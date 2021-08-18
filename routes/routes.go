package routes

import (
	"go_crud/controllers"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func InitRoutes() {
	router := mux.NewRouter()
	// log to stdout
	// router.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(controllers.Index)))
	router.HandleFunc("/", controllers.List).Methods(http.MethodGet)
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/create", http.HandlerFunc(controllers.Create)).Methods(http.MethodPost)
	subRouter.HandleFunc("/{id}/edit", http.HandlerFunc(controllers.Edit)).Methods(http.MethodPut)
	subRouter.HandleFunc("/{id}/delete", http.HandlerFunc(controllers.Destroy)).Methods(http.MethodDelete)

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodOptions})

	server := &http.Server{
		Addr:    ":8000",
		Handler: handlers.CORS(originsOk, headersOk, methodsOk)(router)}

	log.Fatal(server.ListenAndServe())
}
