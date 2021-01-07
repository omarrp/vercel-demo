package demo

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Dolar returns dolar exchange rate
func Dolar(w http.ResponseWriter, r *http.Request) {
	url := "https://api.ratesapi.io/api/latest?base=USD&symbols=BRL"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body))
}
