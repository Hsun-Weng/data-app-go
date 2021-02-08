package main

import (
	"data-app-go/config"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var configPath = flag.String("config", "config.yml", "config file")

func main() {
	flag.Parse()
	appConfig := config.ReadConfig(*configPath)

	profile := "debug"
	switch profile {
	case "release":
		profile = "release"
	case "test":
		profile = "test"
	}
	gin.SetMode(profile)

	db, err := config.NewMongoDatabase(appConfig)
	if err != nil {
		log.Fatalf("Database Connection Failed: %s", err)
		panic(err)
	}
	router := SetupRouter(db)

	endPoint := fmt.Sprintf(":%s", appConfig.Server.Port)

	server := &http.Server{
		Addr:    endPoint,
		Handler: router,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
