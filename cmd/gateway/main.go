package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
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

	// Create a new mux instance
	r := http.NewServeMux()

	// Create a rate limiter to limit by IP address
	// Limited to 10 requests per second
	lmt := tollbooth.NewLimiter(10, &limiter.ExpirableOptions{
		DefaultExpirationTTL: time.Hour,
	})

	lmt.SetIPLookups([]string{"X-Forwarded-For", "X-Real-IP", "RemoteAddr"})
	lmt.SetMethods([]string{"POST"})
	lmt.SetHeaderEntryExpirationTTL(time.Hour)
	lmt.SetMessage("You are being rate limited!")

	// Iterate over all of the services and create handlers
	for _, service := range config.C.Services {
		log.Printf("Creating handler for %s microservice.\n", service.Name)

		targetUrl, err := url.Parse(service.URL)
		if err != nil {
			log.Fatalln("Error creating handler:", err.Error())
		}

		r.Handle(service.Path, tollbooth.LimitHandler(lmt, http.StripPrefix(service.Path, httputil.NewSingleHostReverseProxy(targetUrl))))
	}

	// Configure the HTTP server
	log.Printf("Starting gateway on: %s:%d", config.C.HTTP.Bind, config.C.HTTP.Port)
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.C.HTTP.Bind, config.C.HTTP.Port),
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln("Error starting HTTP server:", err.Error())
	}
}
