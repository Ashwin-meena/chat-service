package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)
type Response struct{
	Success bool `json:"success"`
	Data    interface{} `json:"data"`
}

// Use to send success response to user...
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
     var resp Response 
	w.WriteHeader(statusCode)
    if(statusCode == http.StatusOK){
     resp.Success = true
	}else {
		resp.Success = false
	} 
	resp.Data=data;
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
// Use to send error response to user...
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}