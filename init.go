//
// Anthony Morehouse 2022
//
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
}

func main() {
	response, err := http.Get("https://api.eia.gov/v2/electricity/state-electricity-profiles/emissions-by-state-by-fuel/data/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
}
