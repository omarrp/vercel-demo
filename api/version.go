package demo

import (
	"fmt"
	"net/http"
	"runtime"
)

// Version returns go version
func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, runtime.Version())
}
