# Chat service

Chat service(this repo) is the backend app for sending message to one person to another and also one can read the undelivered message.

**Runtime and Language**: chat service currently uses go 1.20.2.

## Run Project

For run chat service backend app run below command

```
go run main.go

```

## Run Test case

For run chat service test case run below command

```
go test

```

## Api documentation:

1. Sending message api:

   Description: Send message from one user to another user.
   Endpoint: `/api/send`
   Method: `POST`
   Request Body: JSON object containing `from`,`to` and `message`. All field are required field.
   Request Body:
   {
   "from":"User name",
   "to": "User name whome message want to send",
   "message":"Message for user"
   }
   Response:
   {
   "success": true,
   "data": "Message sent from user1 to user2"
   }

2. Get unread message api:

   Description: Get all unread message of user and once use read message api delete the message of user.
   Endpoint: `/api/unread`
   Method: `GET`
   Query Parameter: `user` string type - Name Of user whose message are to be retrived.
   Request `api/unread?user=user2`
   Response:
   {
   "success": true,
   "data": [
   {
   "from": "user1",
   "to": "user2",
   "message": "Hello,user2"
   }
   ]
   }
