package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"

	rawformatter "github.com/dukeofdisaster/avlogparser"
)

func main() {
	// (1) All or some of the following
	r, err := os.Open(os.Args[1])
	rawformatter.CheckError("Can't write to file", err)
	defer r.Close()

	// create a new scanner object for our CLI supplied filename r
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	// this should be a function
	var myLogs []string
	for scanner.Scan() {
		myLogs = append(myLogs, scanner.Text())
	}

	fmt.Println(rawformatter.ValueExtract(myLogs[0], "payload"))
	for i := 0; i < len(rawformatter.Names); i++ {
		fmt.Println(rawformatter.Names[i] + " : " + rawformatter.ValueExtract(myLogs[0], rawformatter.Names[i]))
	}

	// Create a csv file now
	file, err := os.Create("generated.csv")
	rawformatter.CheckError("Can't write to file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	// write data to the writer; if there's an error it will be returned.
	// thus err is always nil/some error
	err = writer.Write(rawformatter.Names)
	for i := range myLogs {
		writer.Write(rawformatter.ExtractAll(myLogs[i]))
	}

}
