package demo

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Random returns random integer
func Random(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(100)
	fmt.Fprintf(w, strconv.Itoa(number))
}
