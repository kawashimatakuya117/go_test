// Package helloworld provides a set of Cloud Functions samples.
package helloworld

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("/", helloGet)
}

// helloGet is an HTTP Cloud Function.
func helloGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func ClientHello(url string) map[string]interface{} {
	var client = &http.Client{}
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Error decoding response: %v", err)
		return nil
	}
	log.Print(result)
	return result
}

func ServerClientHello(w http.ResponseWriter, r *http.Request, url string) {
	res := ClientHello(url)
	fmt.Fprint(w, res["hoge"])
}
