package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

	func main() {
		db, err := sql.Open("mysql",
			"user:password@tcp(127.0.0.1:3306)/hello")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
	}

	// i := 0
	//  for i := 0; i <= 100; i++ {
	// MakeRequest(i)
	//  }
	// fmt.Scanln()
}

// func MakeRequest(i int) {

// 	response, errorResponse := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i))
// 	if errorResponse != nil {
// 		log.Fatalln(errorResponse)
// 	}

// 	body, errorRead := ioutil.ReadAll(response.Body)
// 	if errorRead != nil {
// 		log.Fatalln(errorRead)
// 	}

// 	fmt.Println(string(body))
// }
