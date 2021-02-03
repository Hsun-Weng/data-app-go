package main

import (
	config2 "data-app-go/config"
	"data-app-go/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	gin.SetMode("debug")
	router := router.SetupRouter()
	config := config2.ReadConfig()
	log.Println(config)

	endPoint := fmt.Sprintf(":%d", 8080)
	server := &http.Server{
		Addr:  endPoint,
		Handler: router,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
