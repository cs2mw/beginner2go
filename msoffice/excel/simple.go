package main

import (
	"fmt"
	"log"

	"baliance.com/gooxml/spreadsheet"
)

func main() {
	workbook, _ := spreadsheet.Open("test.xlsx")

	sheet, err := workbook.GetSheet("Sheet1")
	if err != nil {
		log.Fatalf("error finding sheet: %s", err)
	} else {
		// rows
		for r := 0; r < 5; r++ {
			row := sheet.AddRow()
			// and cells
			for c := 0; c < 5; c++ {
				cell := row.AddCell()
				cell.SetString(fmt.Sprintf("row %d cell %d", r, c))
			}
		}

		if err := workbook.Validate(); err != nil {
			log.Fatalf("error validating sheet: %s", err)
		}

		workbook.SaveToFile("test.xlsx")
	}
}
