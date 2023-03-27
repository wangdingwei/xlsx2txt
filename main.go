package main

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s filename", os.Args[0])
		fmt.Fprintln(os.Stderr)
		os.Exit(1)
	}

	file := os.Args[1];
	fmt.Println(file)

    f, err := excelize.OpenFile(file)
    if err != nil {
        fmt.Fprintln(os.Stderr, "open failed:", err)
		os.Exit(1)
    }
    defer func() {
        // Close the spreadsheet.
        if err := f.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
        }
    }()


	for _, sheet := range f.GetSheetList() {
		println(sheet)
	}
	os.Exit(0)

    // Get value from cell by given worksheet name and cell reference.
    cell, err := f.GetCellValue("Sheet1", "B2")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(cell)
    // Get all the rows in the Sheet1.
    rows, err := f.GetRows("Sheet1")
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
    }
}
