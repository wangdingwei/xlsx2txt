package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

var writer *bufio.Writer

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s infile outfile", os.Args[0])
		fmt.Fprintln(os.Stderr)
		os.Exit(1)
	}

	infileName := os.Args[1]
	outfileName := os.Args[2]
	// fmt.Println(infileName, outfileName)

	outfile, err := os.Create(outfileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer outfile.Close()

	writer = bufio.NewWriter(outfile)

	xlsxFile, err := excelize.OpenFile(infileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "open failed:", err)
		os.Exit(1)
	}
	defer func() {
		if err := xlsxFile.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	for _, sheet := range xlsxFile.GetSheetList() {

		fmt.Fprintln(writer)
		fmt.Fprintf(writer, "====== [sheet: %s] ======", sheet)
		fmt.Fprintln(writer)
		printSheet(xlsxFile, sheet)
		fmt.Fprintln(writer)
	}
	writer.Flush()
}

func printSheet(xlsxFile *excelize.File, sheet string) {
	rows, err := xlsxFile.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		len := len(row)
		for idx, cell := range row {
			fmt.Fprint(writer, cell)
			if idx < len-1 { // last
				fmt.Fprint(writer, "|")
			}
		}
		fmt.Fprintln(writer)
	}
}
