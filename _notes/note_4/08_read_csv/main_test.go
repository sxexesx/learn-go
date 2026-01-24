package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

const targetSizeGB = 10
const approxRowSizeBytes = 60

func TestMain(t *testing.T) {
	rowNum := (targetSizeGB * 1024 * 1024 * 1024) / approxRowSizeBytes

	file, err := os.Create("data_10Gb.csv")
	if err != nil {
		t.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	csvWriter.Write([]string{"Name", "Age", "Email", "City", "Salary"})

	now := time.Now()

	for i := 0; i < rowNum; i++ {
		row := []string{
			gofakeit.Name(),
			fmt.Sprintf("%d", gofakeit.IntRange(18, 65)),
			gofakeit.Email(),
			gofakeit.City(),
			fmt.Sprintf("%.2f", gofakeit.Float64Range(30000, 100000)),
		}

		csvWriter.Write(row)

		if i%10000 == 0 {
			// переодический flush
			csvWriter.Flush()
			info, _ := file.Stat()

			fileSize := info.Size()
			fmt.Printf("Записано: %d строк. Размер файла: %v\n", i, fileSize)
			if fileSize > int64(targetSizeGB*1024*1024*1024) {
				break
			}
		}
	}
	fmt.Printf("Сделано за %v", time.Since(now))
}
