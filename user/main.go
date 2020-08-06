package main

import (
	"log"
	"net/http"

	"gitlab.com/achoyal/caringtest/user/route"
)

//main
func main() {
	routes := route.GetRoutes()
	log.Println("Applicatio has been started on : 8002")
	http.Handle("/", routes)
	http.ListenAndServe(":8002", nil)
}
