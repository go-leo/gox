package xlsx

import (
	"bytes"
	"errors"
	"github.com/xuri/excelize/v2"
)

const DefaultSheetName = "Sheet1"

// ToXLSX converts the given fields and rows to a xlsx file.
func ToXLSX(fields []string, rows [][]any) ([]byte, error) {
	file := excelize.NewFile()
	index, err := file.NewSheet(DefaultSheetName)
	if err != nil {
		return nil, errors.Join(err, file.Close())
	}
	for col, value := range fields {
		cell, err := excelize.CoordinatesToCellName(col+1, 1)
		if err != nil {
			return nil, errors.Join(err, file.Close())
		}
		if err := file.SetCellValue(DefaultSheetName, cell, value); err != nil {
			return nil, errors.Join(err, file.Close())
		}
	}
	for row, line := range rows {
		for col, value := range line {
			cell, err := excelize.CoordinatesToCellName(col+1, row+2)
			if err != nil {
				return nil, errors.Join(err, file.Close())
			}
			if err := file.SetCellValue(DefaultSheetName, cell, value); err != nil {
				return nil, errors.Join(err, file.Close())

			}
		}
	}
	file.SetActiveSheet(index)
	buf, err := file.WriteToBuffer()
	if err != nil {
		return nil, errors.Join(err, file.Close())
	}
	return buf.Bytes(), file.Close()
}

// FromXLSX converts the given xlsx file to fields and rows.
func FromXLSX(data []byte) ([]string, [][]any, error) {
	file, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return nil, nil, err
	}
	sheet, err := file.GetRows(file.GetSheetName(0))
	if err != nil {
		return nil, nil, errors.Join(err, file.Close())
	}
	if len(sheet) <= 0 {
		return nil, nil, nil
	}
	fields := sheet[0]
	sheetContent := sheet[1:]
	rows := make([][]any, 0, len(sheetContent)-1)
	for i := 0; i < len(sheetContent); i++ {
		row := make([]any, 0, len(sheetContent[i]))
		for _, cell := range sheetContent[i] {
			row = append(row, cell)
		}
		rows = append(rows, row)
	}
	return fields, rows, file.Close()
}
