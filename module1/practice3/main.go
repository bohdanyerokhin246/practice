package main

import (
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "strconv"
)

func main() {
        for i := 0; i <= 100; i++ {
                go MakeRequest(i)
        }
        fmt.Scanln()
}

func MakeRequest(i int) {

        response, errorResponse := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i))
        if errorResponse != nil {
                log.Fatalln(errorResponse)
        }

        body, errorRead := ioutil.ReadAll(response.Body)
        if errorRead != nil {
                log.Fatalln(errorRead)
        }

        fmt.Println(string(body))
}
