package user

import (
	"log"
	"net/http"

	"gitlab.com/achoyal/caringtest/components/rwhelper"
	"gitlab.com/achoyal/caringtest/user/models"
)

//GetProfile function
func GetProfile(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("AUth Function Called")
	defer func() {
		if ok := recover(); ok != nil {
			rwhelper.ErrorResponse(http.StatusBadGateway, "Invalid user request", w)
			log.Println(ok)
		}
	}()
	var userData models.User
	userData.Username = "Amit"
	userData.Email = "choyalamit1432@gmail.con"
	userData.Age = 32
	userData.Phone = 9717274997
	response := make(map[string]interface{})
	response["data"] = userData
	rwhelper.SuccessRespond(response, w)
	return
}

//MsName function
func MsName(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("AUth Function Called")
	defer func() {
		if ok := recover(); ok != nil {
			rwhelper.ErrorResponse(http.StatusBadGateway, "Invalid user request", w)
			log.Println(ok)
		}
	}()
	var MsName models.MsName
	MsName.Name = "user-microservices"
	response := make(map[string]interface{})
	response["data"] = MsName
	rwhelper.SuccessRespond(response, w)
	return
}
