package models

//User Data...
type User struct {
	Username string      `json:"username"`
	Dob      string      `json:"dob"`
	Age      int         `json:"age"`
	Email    string      `json:"email"`
	Phone    interface{} `json:"phone-number"`
}

//Msname
type MsName struct {
	Name string
}
