package main

import (
        "fmt"
        "io"
        "log"
        "net/http"
        "os"
        "strconv"
)

func main() {
        MakeRequest()
}

func MakeRequest() {

        directoryName := "storage/posts"
        erroMakeDirection := os.MkdirAll(directoryName, 077)
        if erroMakeDirection != nil {
                log.Fatalln(erroMakeDirection)
        }
        fmt.Println("Directory " + directoryName + " create")

        errorChangeDirection := os.Chdir(directoryName)
        if errorChangeDirection != nil {
                log.Fatalln(errorChangeDirection)
        }
        fmt.Println("Directory change to " + directoryName)

        for i := 1; i < 101; i++ {
                response, errorResponse := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i))
                if errorResponse != nil {
                        log.Fatalln(errorResponse)
                }

                file, errorCreateFile := os.Create(strconv.Itoa(i) + ".txt")
                if errorCreateFile != nil {
                        log.Fatalln(errorCreateFile)
                }

                fmt.Println("file " + strconv.Itoa(i) + ".txt is create")
                defer file.Close()

                body, errorRead := io.ReadAll(response.Body)
                if errorRead != nil {
                        log.Fatalln(errorRead)
                }
                file.WriteString(string(body))
        }
}
