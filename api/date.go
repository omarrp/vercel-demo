package demo

import (
	"fmt"
	"net/http"
	"time"
)

// Date basic serveless function
func Date(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("Mon Jan 2 15:04:05 MST 2006")
	fmt.Fprintf(w, currentTime)
}
