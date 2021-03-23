package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"

	"github.com/feirm/gateway/internal/config"
)

func main() {
	log.Println("Starting Feirm Microservice Gateway...")

	// Load in the configuration file
	log.Println("Loading configuration file...")
	if err := config.Load(); err != nil {
		log.Fatalf("Unable to load configuration file: %s", err.Error())
		return
	}

	// WaitGroup
	var wg sync.WaitGroup

	amountOfServices := len(config.C.Services)
	wg.Add(amountOfServices)

	// Iterate over all of the services and create handlers
	for _, service := range config.C.Services {
		log.Printf("Creating handler for %s microservice.\n", service.Name)

		targetUrl, err := url.Parse(service.URL)
		if err != nil {
			log.Fatalln("Error creating handler:", err.Error())
		}

		go http.Handle(service.Path, http.StripPrefix(service.Path, httputil.NewSingleHostReverseProxy(targetUrl)))

		defer wg.Done()
	}

	log.Printf("Starting gateway on: %s:%d", config.C.HTTP.Bind, config.C.HTTP.Port)
	if err := http.ListenAndServe(":"+fmt.Sprintf("%d", config.C.HTTP.Port), nil); err != nil {
		log.Fatalln("Error starting HTTP server:", err.Error())
	}
}
