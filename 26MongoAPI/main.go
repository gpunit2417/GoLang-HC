package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gpunit2417/MongoAPI/router"
)

func main() {
	fmt.Println("Mongodb in golang")
	r := router.Router();
	fmt.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server started at port 4000")
}

