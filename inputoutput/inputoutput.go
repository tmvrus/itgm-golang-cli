package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"
	"github.com/schollz/progressbar/v2"
)

const (
	even int = iota
	odd
	prime
	total

	outputFormatRawText      = "Raw text"
	outputFormatTable        = "Table"
	outputFormatColoredTable = "Colored Table"

	defaultFlagValue = ""
)

func main() {
	result := make([]int, 4, 4)

	sourceFileName := ""
	flag.StringVar(&sourceFileName, "source", "", "data source file name")
	flag.Parse()
	if sourceFileName == defaultFlagValue {
		fmt.Printf("source file not provided\n")
		flag.Usage()
		os.Exit(1)
	}

	sourceFile, err := os.Open(sourceFileName)
	if err != nil {
		fmt.Printf("failed to Open source file with name %q, error: %s", sourceFileName, err.Error())
		os.Exit(1)
	}

	s, err := sourceFile.Stat()
	if err != nil {
		fmt.Printf("failed to get Stat for source file, error: %s\n", err.Error())
		os.Exit(1)
	}

	fileSize := int(s.Size())
	if fileSize <= 0 {
		fmt.Printf("source file %q is empty, nothing to process\n", sourceFileName)
		os.Exit(1)
	}

	prompt := promptui.Select{
		Label: "Select output format",
		Items: []string{outputFormatRawText, outputFormatTable, outputFormatColoredTable},
	}

	_, format, err := prompt.Run()
	if err != nil {
		fmt.Printf("failed to read user input, error: %s\n", err.Error())
		os.Exit(1)
	}

	bar := progressbar.NewOptions(
		fileSize,
		progressbar.OptionSetDescription("Processing input data"),
		progressbar.OptionSetBytes(fileSize),
	)

	r := bufio.NewReader(sourceFile)
	for {
		line, _, err := r.ReadLine()

		if err == io.EOF {
			time.Sleep(time.Second)
			break
		}

		if err != nil {
			fmt.Printf("failed to ReadLine from stdIn, error: %s", err.Error())
			os.Exit(1)
		}

		num, err := strconv.ParseInt(string(line), 10, 32)
		if err != nil {
			fmt.Printf("failed to ParseInt input data %q, error: %s", line, err.Error())
			os.Exit(1)
		}

		result[total]++
		if num%2 == 0 {
			result[even]++
		} else if isPrime(int(num)) {
			result[odd]++
			result[prime]++
		} else {
			result[odd]++
		}

		if err := bar.Add(len(line) + 1); err != nil {
			fmt.Printf("failed to Add line size to progress bar, error: %s", err.Error())
			os.Exit(1)
		}

		time.Sleep(time.Millisecond)
	}

	if err := bar.Clear(); err != nil {
		fmt.Printf("failed to Clear progress bar, error: %s", err.Error())
		os.Exit(1)
	}

	if format == outputFormatRawText {
		fmt.Printf("even: %d\n", result[even])
		fmt.Printf("odd: %d\n", result[odd])
		fmt.Printf("prime: %d\n", result[prime])
		fmt.Printf("total: %d\n", result[total])
	} else {
		renderTable(result, format == outputFormatColoredTable)
	}

}

func isPrime(a int) bool {
	if a <= 1 {
		return false
	}

	for i := 2; i < a; i++ {
		if a%i == 0 {
			return false
		}
	}

	return true
}

func renderTable(result []int, colored bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Type", "Count"})
	table.AppendBulk([][]string{
		{"Even", fmt.Sprintf("%d", result[even])},
		{"Odd", fmt.Sprintf("%d", result[odd])},
		{"Prime", fmt.Sprintf("%d", result[prime])},
		{"Total", fmt.Sprintf("%d", result[total])},
	})

	if colored {
		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold},
		)
		table.SetColumnColor(
			tablewriter.Colors{tablewriter.FgGreenColor},
			tablewriter.Colors{tablewriter.FgHiYellowColor},
		)
	}

	table.Render()
}
