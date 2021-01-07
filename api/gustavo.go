package demo

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Gustavo teste com gustavo
func Gustavo(w http.ResponseWriter, r *http.Request) {
	// defer resp.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, r.URL.Path)
}
