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
	HostName    string
	AZ          string
	MachineType string
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	data := getMetadata()
	fmt.Println(*data)
	// fmt.Println(p)
	// Start HTTP server.
	// log.Printf("listening on port %s", port)
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatal(err)
	// }

	// metadata := Metadata{HostName: string(getMetadata("/hostname"))}
	// fmt.Println(metadata)
	// hostname := getMetadata("/hostname")
	// zone := strings.SplitAfter(string(getMetadata("/zone")), "/")
	// zoneText := zone[len(zone)-1]
	// machineType := strings.SplitAfter(string(getMetadata("/machine-type")), "/")
	// machineTypeText := machineType[len(machineType)-1]

}

func handler(w http.ResponseWriter, r *http.Request) {
	// hostname := getMetadata("/hostname")
	// zone := strings.SplitAfter(string(getMetadata("/zone")), "/")
	// zoneText := zone[len(zone)-1]
	// machineType := strings.SplitAfter(string(getMetadata("/machine-type")), "/")
	// machineTypeText := machineType[len(machineType)-1]

	// fmt.Fprintf(w, "hostname: %s\nzone: %s\nmachineType: %s\n", hostname, zoneText, machineTypeText)
}

func getMetadata() *InstanceMetadata {

	c := metadata.NewClient(&http.Client{})
	hostname, err := c.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	zone, err := c.Zone()
	if err != nil {
		log.Fatal(err)
	}
	machineType, err := c.Get("instance/machine-type")
	if err != nil {
		log.Fatal(err)
	}
	machineTypeText := strings.SplitAfter(string(machineType), "/")

	return &InstanceMetadata{
		HostName:    hostname,
		AZ:          zone,
		MachineType: machineTypeText[len(machineTypeText)-1],
	}

}
