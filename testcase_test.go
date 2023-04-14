package main

import (
	"bytes"
	"chat/controller"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"chat/model"
)

// TestSendMessage test function to test send message to user...
func TestSendMessage(t *testing.T){
	message := model.Message{
		From: "user1",
		To: "user2",
		Message: "Hello,user2",
	}
	payload,_ := json.Marshal(message);
	req,err := http.NewRequest("POST","/api/send",bytes.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler :=http.HandlerFunc(controller.SendMessage)
	handler.ServeHTTP(rr,req)
	if status := rr.Code; status != http.StatusOK{
		t.Errorf("Error:%v",status)
	}
	exceptedResponse := `{"success":true,"data":"Message sent from user1 to user2"}`
	var testResponse, apiResponse interface{}
      json.Unmarshal([]byte(rr.Body.String()),&testResponse)
      json.Unmarshal([]byte(exceptedResponse),&apiResponse)
	if  !reflect.DeepEqual(testResponse,apiResponse) {
		t.Errorf("Not return correct msg %v",rr.Body.String())
	}
}

// TestUnreadMessages test function to test unread message of user...
func TestUnreadMessages(t *testing.T) {
	req,err:=http.NewRequest("GET","/api/unread?user=user2",nil);
		if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler :=http.HandlerFunc(controller.GetUnreadMessages)
	handler.ServeHTTP(rr,req)

	if status := rr.Code; status != http.StatusOK{
		t.Errorf("Error:%v",status)
	}
	exceptedResponse := `{"success":true,"data":[{"from":"user1","to":"user2","message":"Hello,user2"}]}`;
	expectedNullResponse := `{"success":true,"data":null}`
	var testResponse, apiResponse,testNullResponse interface{}
    json.Unmarshal([]byte(rr.Body.String()),&testResponse)
    json.Unmarshal([]byte(exceptedResponse),&apiResponse)
	json.Unmarshal([]byte(expectedNullResponse),&testNullResponse)
	if  !reflect.DeepEqual(testResponse,apiResponse) {
		if !reflect.DeepEqual(testResponse,apiResponse){
		t.Errorf("Not return correct msg %v",rr.Body.String())
		}
	}

}