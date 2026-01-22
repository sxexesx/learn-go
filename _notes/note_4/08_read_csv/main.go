package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("start reading file...")

	bytes, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer bytes.Close()

	reader := csv.NewReader(bytes)
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading bytes:", err)
		return
	}

	persons := make([]Person, 0, len(data))
	for _, row := range data {
		persons = append(persons,
			Person{
				Name:   row[0],
				Age:    row[1],
				Email:  row[2],
				City:   row[3],
				Salary: row[4],
			})
	}
	lastPerson := persons[len(persons)-1]

	fmt.Printf("Last person: %v", lastPerson)
}

type Person struct {
	Name   string
	Age    string
	Email  string
	City   string
	Salary string
}
