package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	rows := readOrders("test_task_data.csv")
	addingStations(rows)
}

func readOrders(name string) [][]string {

	f, err := os.Open(name)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}

	defer f.Close()

	r := csv.NewReader(f)

	r.Comma = ';'

	rows, err := r.ReadAll()

	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}
	return rows
}

func addingStations(rows [][]string) {

	var DataMatrix [150][4]int
	for i := range rows {

		Train_Number, err := strconv.Atoi(rows[i][0])
		if err != nil {
			fmt.Println(err)
		}

		DataMatrix[i][0] = Train_Number

		Poisoning_Station_Number, err := strconv.Atoi(rows[i][1])
		if err != nil {
			fmt.Println(err)
		}

		DataMatrix[i][1] = Poisoning_Station_Number

		Arrival_Station_Number, err := strconv.Atoi(rows[i][2])
		if err != nil {
			fmt.Println(err)
		}

		DataMatrix[i][2] = Arrival_Station_Number

		// Price, err := strconv.Atoi(rows[i][3])
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// DataMatrix[i][3] = Price

	}

	fmt.Println(DataMatrix)
}

// func writeOrders(name string, rows [][]string) {

// 	f, err := os.Create(name)
// 	if err != nil {
// 		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
// 	}

// 	defer func() {
// 		e := f.Close()
// 		if e != nil {
// 			log.Fatalf("Cannot close '%s': %s\n", name, e.Error())
// 		}
// 	}()

// 	w := csv.NewWriter(f)
// 	err = w.WriteAll(rows)
// }
