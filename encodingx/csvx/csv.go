package csvx

import (
	"bytes"
	"encoding/csv"
	"strings"
)

// ToCSV converts the given fields and rows to a csv file.
func ToCSV(fields []string, rows [][]string) []byte {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	for i, field := range fields {
		fields[i] = strings.Replace(field, "\n", "\\n", -1)
	}
	if len(fields) > 0 {
		_ = writer.Write(fields) // nolint
	}
	for _, row := range rows {
		for i, field := range row {
			field = strings.Replace(field, "\n", "\\n", -1)
			field = strings.Replace(field, "\r", "\\r", -1)
			row[i] = field
		}
		_ = writer.Write(row) // nolint
	}
	writer.Flush()
	return buf.Bytes()
}

// FromCSV converts the given csv file to fields and rows.
func FromCSV(data []byte) ([]string, [][]string, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	var fields []string
	var rows [][]string
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	if len(records) > 0 {
		fields = records[0]
	}
	if len(records) > 1 {
		rows = records[1:]
	}
	return fields, rows, nil
}
