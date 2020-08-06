package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"gitlab.com/achoyal/caringtest/auth/route"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	routes := route.GetRoutes()
	log.Println("Applicatio has been started on : 8001")
	http.Handle("/", routes)
	http.ListenAndServe(":8001", nil)
}
