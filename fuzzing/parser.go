package main

import (
	"encoding/csv"
	"strings"
)

func ParseCSVLine(line string) ([]string, error) {
	r := csv.NewReader(strings.NewReader(line))
	r.FieldsPerRecord = -1
	return r.Read()
}
