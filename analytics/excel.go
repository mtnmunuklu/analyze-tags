package analytics

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type ExcelParams struct {
	SheetName string
	Data      map[string][]string
	Output    string
}

func (e *ExcelParams) ToExcel() error {
	file := excelize.NewFile()
	index, err := file.NewSheet(e.SheetName)
	if err != nil {
		return err
	}

	file.SetCellValue(e.SheetName, "A1", "Rule")
	file.SetCellValue(e.SheetName, "B1", "Tag")

	row := 2
	for rule, tags := range e.Data {
		for _, tag := range tags {
			file.SetCellValue(e.SheetName, fmt.Sprintf("A%d", row), rule)
			file.SetCellValue(e.SheetName, fmt.Sprintf("B%d", row), tag)
			row++
		}
	}

	file.SetActiveSheet(index)

	err = file.SaveAs(e.Output)
	if err != nil {
		return err
	}

	return nil
}
