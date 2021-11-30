package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "log"
)

func main() {
    resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}

    body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println(len(body))
}

