// Copyright 2017 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	color   = "yellow" 
	version = "v1.0.0"
)

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", versionHandler)
	s := http.Server{Addr: ":8080"}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	log.Printf("Shutdown signal received, exiting...")

	s.Shutdown(context.Background())
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Color", color)
	message := fmt.Sprintf("Hello World %s\n", version)
	fmt.Fprintf(w, message)
}
