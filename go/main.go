// Example GCP server that responds with meta data of the server
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

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

func handler(w http.ResponseWriter, r *http.Request) {
	hostname := getMetadata()
	fmt.Fprintf(w, "%s\n", hostname)
}

func getMetadata() []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://metadata.google.internal/computeMetadata/v1/instance/hostname", nil)

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
