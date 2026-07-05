package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

const salesCSV = `name,city,total
Ana,Santos,42.50
Bruno,Recife,19.90
`

func ParseDelimited(data string, comma rune) ([][]string, error) {
	reader := csv.NewReader(strings.NewReader(data))
	reader.Comma = comma
	reader.FieldsPerRecord = -1

	var rows [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		rows = append(rows, record)
	}

	return rows, nil
}

func FormatRows(rows [][]string) string {
	var builder strings.Builder
	for _, row := range rows {
		fmt.Fprintf(&builder, "%s\n", strings.Join(row, " | "))
	}
	return builder.String()
}

func main() {
	rows, err := ParseDelimited(salesCSV, ',')
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse delimited data: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(FormatRows(rows))
}
