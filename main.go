package main

import (
	"fmt"
	"github.com/p4ndabk/ibge-integration/internal/api/route"
	"log"
	"net/http"
)

func main() {
	router := route.InitRouter()

	fmt.Println("Server is running on port 4001")
	log.Fatal(http.ListenAndServe(":4001", router))
}
