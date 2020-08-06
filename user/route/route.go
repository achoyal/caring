package route

import (
	"fmt"

	"github.com/gorilla/mux"
	"gitlab.com/achoyal/caringtest/user/controllers/user"
)

// GetRoutes registers a new route with a matcher for the URL path.
// See Route.Path() and Route.HandlerFunc().
func GetRoutes() *mux.Router {
	fmt.Println("inside user route")
	r := mux.NewRouter().StrictSlash(true)
	a := r.PathPrefix("/").Subrouter()
	a.HandleFunc("/user/profile", user.GetProfile).Methods("GET")
	a.HandleFunc("/microservice/name", user.MsName).Methods("GET")
	//a := r.PathPrefix("/auth").Subrouter()
	//a.HandleFunc("/", auth.CheckAuth).Methods("GET")
	return r
}
