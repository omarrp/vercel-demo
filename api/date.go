package demo

import (
	"fmt"
	"net/http"
	"time"
)

// Date basic serveless function
func Date(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)
	fmt.Fprintf(w, currentTime)
}
