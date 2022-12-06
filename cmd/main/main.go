package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/space-devops/mountebank-sidecar/pkg/handlers"
	"github.com/space-devops/mountebank-sidecar/pkg/middleware"
	"github.com/space-devops/mountebank-sidecar/pkg/utils"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello world from Golang")

	r := mux.NewRouter()

	r.Use(middleware.CorrelationMiddleware)

	r.HandleFunc("/", handlers.WelcomeHandler).GetHandler()

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", utils.ServerPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
