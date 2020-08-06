package route

import (
	"github.com/gorilla/mux"
	"gitlab.com/achoyal/caringtest/auth/controllers/auth"
)

// GetRoutes registers a new route with a matcher for the URL path.
// See Route.Path() and Route.HandlerFunc().
func GetRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	a := r.PathPrefix("/auth").Subrouter()
	a.HandleFunc("/", auth.CheckAuth).Methods("GET")
	return r
}
