package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/p4ndabk/ibge-integration/app/api/route"
)

func StartApplication() {
	err := InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	port := Env.ApiRestPort
	fmt.Println("Server is running on port ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), route.InitRouter()))
}