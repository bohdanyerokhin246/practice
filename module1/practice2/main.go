package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	MakeRequest()
}

func MakeRequest() {
	response, errorResponse := http.Get("https://jsonplaceholder.typicode.com/posts/")
	if errorResponse != nil {
		log.Fatalln(errorResponse)
	}

	body, errorRead := ioutil.ReadAll(response.Body)
	if errorRead != nil {
		log.Fatalln(errorRead)
	}

	fmt.Println(string(body))
}
