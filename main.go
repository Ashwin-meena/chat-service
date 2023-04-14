package main

import (
    "fmt"
    "log"
	"chat/router"
    //"github.com/joho/godotenv"
	"net/http"
)

func main() {
    //loadEnv()
    serveApplication()
}
 //** For load env file **// Commented because 
// func loadEnv() {
//     err := godotenv.Load(".env")
//     if err != nil {
//         log.Fatal("Error loading .env file")
//     }
// }

func serveApplication() {
    r := router.Router()
	r.Use()
    fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}