package controller

import (
	"encoding/json"
	"net/http"
	"fmt"
	"chat/response"
	"errors"
	"github.com/go-playground/validator"
	"chat/model"

)


// Make map for storing unread messages so tha user can retrive messages...
var messages = make(map[string][]model.Message);

// SendMessage send message to user
func SendMessage(w http.ResponseWriter, r *http.Request) {
	var messageBody model.Message
	    // Decode the json request...
    err := json.NewDecoder(r.Body).Decode(&messageBody)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to decode the request body.  %v", err)
		response.ERROR(w, http.StatusBadRequest, errors.New(errMsg));
		return
    }
	   // Validate the body...
	validate := validator.New()
     if err := validate.Struct(messageBody); err != nil {
		response.ERROR(w, http.StatusBadRequest, err);
		return
     }
	messages[messageBody.To]= append(messages[messageBody.To],messageBody);
	responseMessage := fmt.Sprintf("Message sent from %s to %s",messageBody.From,messageBody.To);
	//Send response to user...
    response.JSON(w, http.StatusOK,responseMessage)
}


//GetUnreadMessages get all unread message of user...
func GetUnreadMessages(w http.ResponseWriter, r *http.Request){
	   user := r.URL.Query().Get("user");
	   responses := messages[user];
	   // Delete the message of user form message because user recived the message...
	   delete(messages,user);
	   response.JSON(w, http.StatusOK, responses)
}


