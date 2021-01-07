package demo

import (
	"fmt"
	"net/http"
)

// Hello basic serveless function
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
