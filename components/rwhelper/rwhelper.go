package rwhelper

import (
	"encoding/json"
	"net/http"
)

//SuccessRespond function
func SuccessRespond(fields map[string]interface{}, writer http.ResponseWriter) {
	fields["status"] = "success"
	message, err := json.Marshal(fields)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("An error occured internally"))
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(message)

}

//ErrorResponse function
func ErrorResponse(statusCode int, error string, writer http.ResponseWriter) {
	//Create a new map and fill it
	fields := make(map[string]interface{})
	fields["status"] = "error"
	fields["message"] = error
	message, err := json.Marshal(fields)

	if err != nil {
		//An error occurred processing the json
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("An error occured internally"))
	}
	//Send header, status code and output to writer
	writer.WriteHeader(statusCode)
	writer.Write(message)
}
