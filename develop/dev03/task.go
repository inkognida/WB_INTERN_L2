package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	keyColumn       int
	sortByNumeric   bool
	sortReverse     bool
	ignoreDuplicates bool
	sortByMonth     bool
	ignoreTrailing  bool
	checkSorted     bool
	sortByNumericSuffix bool
)

func init() {
	flag.IntVar(&keyColumn, "k", 0, "column index for sorting")
	flag.BoolVar(&sortByNumeric, "n", false, "sort by numeric value")
	flag.BoolVar(&sortReverse, "r", false, "sort in reverse order")
	flag.BoolVar(&ignoreDuplicates, "u", false, "ignore duplicate lines")
	flag.BoolVar(&sortByMonth, "M", false, "sort by month name")
	flag.BoolVar(&ignoreTrailing, "b", false, "ignore trailing whitespace")
	flag.BoolVar(&checkSorted, "c", false, "check if data is sorted")
	flag.BoolVar(&sortByNumericSuffix, "h", false, "sort by numeric value with suffixes")
	flag.Parse()
}

func main() {
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: sort-tool <filename>")
		os.Exit(1)
	}

	filename := args[0]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	lines := strings.Split(string(data), "\n")
	lines = removeEmptyLines(lines)

	sortLines(lines)

	output := strings.Join(lines, "\n")
	fmt.Println(output)
}

func removeEmptyLines(lines []string) []string {
	var cleanedLines []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			cleanedLines = append(cleanedLines, line)
		}
	}
	return cleanedLines
}

func sortLines(lines []string) {
	sort.SliceStable(lines, func(i, j int) bool {
		a := lines[i]
		b := lines[j]

		// Key column
		if keyColumn > 0 {
			aColumns := strings.Fields(a)
			bColumns := strings.Fields(b)

			if len(aColumns) >= keyColumn && len(bColumns) >= keyColumn {
				a = aColumns[keyColumn-1]
				b = bColumns[keyColumn-1]
			}
		}

		// Sort by numeric value
		if sortByNumeric {
			aNum, aErr := strconv.ParseFloat(a, 64)
			bNum, bErr := strconv.ParseFloat(b, 64)
			if aErr == nil && bErr == nil {
				return aNum < bNum
			}
		}

		// Sort by numeric value with suffixes
		if sortByNumericSuffix {
			aNum := parseNumericValueWithSuffix(a)
			bNum := parseNumericValueWithSuffix(b)
			return aNum < bNum
		}

		// Sort by month name
		if sortByMonth {
			aTime, aErr := time.Parse("January", a)
			bTime, bErr := time.Parse("January", b)
			if aErr == nil && bErr == nil {
				return aTime.Before(bTime)
			}
		}

		// Ignore trailing whitespace
		if ignoreTrailing {
			a = strings.TrimRight(a, " ")
			b = strings.TrimRight(b, " ")
		}

		// Default string comparison
		return a < b
	})

	// Reverse order
	if sortReverse {
		reverse(lines)
	}

	// Ignore duplicate lines
	if ignoreDuplicates {
		lines = removeDuplicates(lines)
	}

	// Check if data is sorted
	if checkSorted {
		if !isSorted(lines) {
			log.Println("Data is not sorted")
		} else {
			log.Println("Data is sorted")
		}
	}
}

func parseNumericValueWithSuffix(value string) float64 {
	if len(value) == 0 {
		return 0
	}

	lastChar := value[len(value)-1]
	multiplier := 1.0

	switch lastChar {
	case 'K':
		multiplier = 1e3
	case 'M':
		multiplier = 1e6
	case 'B':
		multiplier = 1e9
	}

	if multiplier != 1.0 {
		value = value[:len(value)-1]
	}

	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	return num * multiplier
}

func reverse(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

func removeDuplicates(lines []string) []string {
	var uniqueLines []string
	seen := make(map[string]bool)

	for _, line := range lines {
		if !seen[line] {
			uniqueLines = append(uniqueLines, line)
			seen[line] = true
		}
	}

	return uniqueLines
}

func isSorted(lines []string) bool {
	for i := 1; i < len(lines); i++ {
		if lines[i-1] > lines[i] {
			return false
		}
	}
	return true
}
