package main

import (
	"bkirk/utilities"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const delim = "\t"

// PreWaitlist make UPH pre waitlist file
func main() {
	start := time.Now()

	//constants
	inputFilePath := "Z:/Pre Waitlist Data.txt"
	outputFilePath := "Z:/HSB1010_Pre_WL.txt"

	//creat output file
	outFile, err := os.Create(outputFilePath)
	utilities.Check(err)
	defer outFile.Close()

	//open input file for reading
	inFile, err := os.Open(inputFilePath)
	utilities.Check(err)
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		record := buildRecord(line)
		record.WriteToFile(outFile)
	}

	check(scanner.Err())

	fmt.Println("Execution time: " + time.Since(start).String())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// func delim() string {
// 	return "\t"
// }

func buildRecord(line string) utilities.Record {
	record := utilities.Record{}
	data := strings.Split(line, delim)
	var mrn = data[3]
	var refDate = data[4]
	var evalDate = data[5]
	var episodeType = translateProgram(data[7])
	var startDate = getStartDate(refDate, evalDate)

	record.InsertItem(1, "*")
	record.InsertItem(3, mrn)
	record.InsertItem(20, "T")
	record.InsertItem(50, "1")
	record.InsertItem(70, startDate)
	record.InsertItem(30035, episodeType)

	return record
}

func translateProgram(program string) string {
	switch program {
	case "Kidney Recipient":
		return "2"
	case "Living Donor":
		return "1"
	default:
		return ""
	}
}

func getStartDate(refDate string, evalDate string) string {
	if refDate != "" {
		return refDate
	} else if evalDate != "" {
		return evalDate
	} else {
		return ""
	}
}
