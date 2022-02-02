// Example GCP server that responds with metadata
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/compute/metadata"
)

type InstanceMetadata struct {
	hostname    string
	zone        string
	machinetype string
}

func main() {
	// only pull traits once
	data := getMetadata()

	log.Print("starting server...")
	http.HandleFunc("/", data.handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func (data InstanceMetadata) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hostname: %s\nzone: %s\nmachineType: %s\n", data.hostname, data.zone, data.machinetype)
}

func getMetadata() *InstanceMetadata {
	// Use GCP SDK to pull metadata
	c := metadata.NewClient(&http.Client{})
	hostname, err := c.Hostname()
	if err != nil {
		hostname = "none"
	}
	zone, err := c.Zone()
	if err != nil {
		zone = "none"
	}
	machineType, err := c.Get("instance/machine-type")
	if err != nil {
		machineType = "none"
	}
	machineTypeText := strings.SplitAfter(string(machineType), "/")

	return &InstanceMetadata{
		hostname:    hostname,
		zone:        zone,
		machinetype: machineTypeText[len(machineTypeText)-1],
	}

}
