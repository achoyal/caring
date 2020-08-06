package auth

import (
	"fmt"
	"log"
	"net/http"

	"gitlab.com/achoyal/caringtest/components/rwhelper"
)

const uname = "caring"

//AuthCheck function
func CheckAuth(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if ok := recover(); ok != nil {
			// Write logs
			log.Println(ok)
		}
	}()
	response := make(map[string]interface{})
	//uname := os.Getenv("username")
	username := r.Header.Get("username")
	if uname == username {
		fmt.Println("inside success")
		rwhelper.SuccessRespond(response, w)
		//response["status"] = 200
		return
	}
	rwhelper.ErrorResponse(http.StatusUnauthorized, "Invalid user request", w)
}
