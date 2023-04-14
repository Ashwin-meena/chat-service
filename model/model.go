package model 

// struct of message..
type Message struct{
	From     string    `json:"from" validate:"required"`
	To       string    `json:"to" validate:"required"`
	Message  string    `json:"message" validate:"required"`
}
