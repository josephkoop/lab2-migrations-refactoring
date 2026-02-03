package main

import(
	"net/http"
	"github.com/Joseph-Koop/lab2-josephkoop/internal/routes"
)

func main(){
	mux := http.NewServeMux()

	routes.SetupRoutes(mux)
}