// Example GCP server that responds with metadata
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/compute/metadata"
)

// type Metadata struct {
// 	HostName    string
// 	AZ          string
// 	MachineType string
// }

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	c := metadata.NewClient(&http.Client{})
	p, err := c.ProjectID()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
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

func getMetadata(path string) []byte {
	client := &http.Client{}
	url := "http://metadata.google.internal/computeMetadata/v1/instance" + path
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Metadata-flavor", "Google")
	resp, err := client.Do(req)
	// check for response error
	if err != nil {
		log.Fatal(err)
	}
	// read response body
	data, _ := ioutil.ReadAll(resp.Body)

	// close response body
	resp.Body.Close()

	return data
}
