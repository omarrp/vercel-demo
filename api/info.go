package demo

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
)

// Info returns info about environment
func Info(w http.ResponseWriter, r *http.Request) {
	nc := runtime.NumCPU()
	ng := runtime.NumGoroutine()
	// ncc := runtime.NumCgoCall()
	gmp := runtime.GOMAXPROCS()
	os := runtime.GOOS
	arch := runtime.GOARCH
	info := map[string]string{
		"NumCPU":       strconv.Itoa(nc),
		"NumGoroutine": strconv.Itoa(ng),
		// "NumCgoCall":   strconv.Itoa(ncc),
		"GOMAXPROCS": strconv.Itoa(gmp),
		"OS":         os,
		"ARCH":       arch,
	}
	for k, v := range info {
		fmt.Fprintf(w, k, v)
	}
}
