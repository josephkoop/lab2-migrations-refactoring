package routes

import (
	"log"
	"net/http"
	"github.com/Joseph-Koop/lab2-josephkoop/internal/handlers"
	"github.com/Joseph-Koop/lab2-josephkoop/internal/middleware"
)

func SetupRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/contact", handlers.Contact)
	mux.HandleFunc("/hobby", handlers.Hobby)

	middlewareFunctions := middleware.LoggingMiddleware(middleware.CountingMiddleware(mux))

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", middlewareFunctions)
	log.Fatal(err)

}
