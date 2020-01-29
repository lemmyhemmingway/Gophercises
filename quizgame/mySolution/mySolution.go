package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func mapCsvData(data [][]string) map[string]string {
	quizMap := map[string]string{}
	for _, items := range data {
		result := items[1]
		question := items[0]
		quizMap[result] = string(question)
	}
	return quizMap
}

func main() {
	records := readCsvFile("problems.csv")
	quizMap := mapCsvData(records)
	correct := 0
	wrong := 0
	answer := ""
	for key, value := range quizMap {
		fmt.Print(value, "= ")
		fmt.Scan(&answer)
		if key == answer {
			correct++
		} else {
			wrong++
		}
	}
	fmt.Println("correct: ", correct, "wrong:", wrong)
}
