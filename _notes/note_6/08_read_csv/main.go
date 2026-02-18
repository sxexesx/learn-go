package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("start reading file...")

	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	tmpFile, err := os.CreateTemp(".", "temp_data.csv")
	if err != nil {
		fmt.Printf("error creating temp file: %v", err)
		return
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	writer := csv.NewWriter(tmpFile)
	defer writer.Flush()

	header, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading header:", err)
		return
	}
	fmt.Printf("%v", header)

	if err := writer.Write(header); err != nil {
		fmt.Printf("error writing header: %v", err)
	}
	writer.Flush()

	var counter int
	now := time.Now()

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading row:", err)
			return
		}

		salary, err := strconv.ParseInt(row[4], 10, 0)
		if err != nil {
			return
		}

		person := Person{
			Name:   "Mr." + row[0],
			Age:    row[1],
			Email:  row[2],
			City:   row[3],
			Salary: fmt.Sprintf("%d", salary+10000),
		}
		counter++

		if err := writer.Write([]string{
			person.Name,
			person.Age,
			person.Email,
			person.City,
			person.Salary,
		}); err != nil {
			fmt.Println(err)
			break
		}

		if counter%1000 == 0 {
			writer.Flush()
		}
	}
	writer.Flush()

	file.Close()
	tmpFile.Close()

	if err := os.Rename(tmpFile.Name(), "data.csv"); err != nil {
		fmt.Printf("Error replacing file: %v", err)
		return
	}

	fmt.Printf("Task done for %v\n", time.Since(now))
}

type Person struct {
	Name   string
	Age    string
	Email  string
	City   string
	Salary string
}
