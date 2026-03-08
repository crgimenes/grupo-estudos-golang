package main

import (
	"encoding/csv"
	"io"
	"strings"
)

func ParseCSV(data string, comma rune) ([][]string, error) {
	r := csv.NewReader(strings.NewReader(data))
	r.Comma = comma
	r.FieldsPerRecord = -1
	rows := [][]string{}
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		rows = append(rows, rec)
	}
	return rows, nil
}
