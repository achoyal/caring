package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/achoyal/caringtest/components/helper"
	"gitlab.com/achoyal/caringtest/components/rwhelper"
)

// GetRoutes registers a new route with a matcher for the URL path.
// See Route.Path() and Route.HandlerFunc().
func GetRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", notFoundHandler)
	a := r.PathPrefix("/user").Subrouter()
	a.HandleFunc("/profile", userProfile).Methods("GET")
	a.Use(authMiddleware)
	b := r.PathPrefix("/microservice").Subrouter()
	b.HandleFunc("/name", microserviceName).Methods("GET")
	return r
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is error handler for path123"+r.URL.Path[1:])
}

func userProfile(w http.ResponseWriter, r *http.Request) {
	URI := "http://localhost:8002/user/profile"
	serviceParams := helper.ServiceParams{
		URI:    URI,
		Method: r.Method,
	}
	response, err := helper.CallService(serviceParams)
	if err != nil {
		log.Fatalln(err)
	}
	res := make(map[string]interface{})
	if response.StatusCode == 200 {
		/**get response body */
		data, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(data))
		json.Unmarshal(data, &res)
		rwhelper.SuccessRespond(res, w)
		return
	}
	rwhelper.ErrorResponse(http.StatusUnauthorized, "Invalid user request", w)
	return
}

func microserviceName(w http.ResponseWriter, r *http.Request) {
	URI := "http://localhost:8002/microservice/name"
	serviceParams := helper.ServiceParams{
		URI:    URI,
		Method: r.Method,
	}
	response, err := helper.CallService(serviceParams)
	if err != nil {
		log.Fatalln(err)
	}
	res := make(map[string]interface{})
	if response.StatusCode == 200 {
		/**get response body */
		data, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(data))
		json.Unmarshal(data, &res)
		/*End of get response */
		rwhelper.SuccessRespond(res, w)
		return
	}
	rwhelper.ErrorResponse(http.StatusUnauthorized, "Invalid user request", w)
	return
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.Header.Get("username"))
		URI := "http://localhost:8001/auth/"
		serviceParams := helper.ServiceParams{
			URI:    URI,
			Method: r.Method,
			Token:  r.Header.Get("username"),
		}
		response, err := helper.CallService(serviceParams)
		if err != nil {
			log.Fatalln(err)
		}
		if response.StatusCode == 200 {
			next.ServeHTTP(w, r)
			return
		}
		rwhelper.ErrorResponse(http.StatusUnauthorized, "Invalid user request", w)
		return
	})
}
