package main

import (
	"log"
	"net/http"

	"github.com/pranayjoshi/go-videochat-app/utils"
)

func main() {
	utils.AllRooms.Init()

	http.HandleFunc("/create", utils.CreateRoomRequestHandler)
	http.HandleFunc("/join", utils.JoinRoomRequestHandler)

	log.Println("Starting Server on Port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
