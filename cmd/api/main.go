package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nosilex/crebito/cmd/api/route"
	"github.com/nosilex/crebito/internal/infrastructure/configuration"
	"github.com/nosilex/crebito/pkg/helper"
	"github.com/nosilex/crebito/pkg/http"
)

func main() {
	log.Println("⚡️Starting application...")

	//init port
	port, err := configuration.NewPort()
	if err != nil {
		log.Fatalf("error on init port: %s", err)
	}
	//init service
	service := configuration.NewService(port)
	//init handler
	handler := configuration.NewHandler(service)

	//run server
	http.Run(route.Setup(handler), fmt.Sprintf("0.0.0.0:%s", helper.Coalesce(os.Getenv("APP_PORT"), "8000")))
}
