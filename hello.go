// Package helloworld provides a set of Cloud Functions samples.
package helloworld

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("/", helloGet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// helloGet is an HTTP Cloud Function.
func helloGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
