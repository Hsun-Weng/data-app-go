package main

import (
	"data-app-go/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	gin.SetMode("debug")

	appConfig := config.ReadConfig()
	db, err := config.NewMongoDatabase(appConfig)
	if err != nil {
		log.Fatalf("Database Connection Failed: %s", err)
		panic(err)
	}
	router := SetupRouter(db)

	endPoint := fmt.Sprintf(":%s", appConfig.Server.Port)

	server := &http.Server{
		Addr:  endPoint,
		Handler: router,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}