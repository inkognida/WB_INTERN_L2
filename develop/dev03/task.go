package main

/*
sort
-k: указание колонки для сортировки
-n: сортировать по числовому значению
-r: сортировать в обратном порядке
-u: не выводить повторяющиеся строки
*/

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Line структура строки
type Line struct {
	Text string
	Key  string
	Column bool
}

type Lines []Line

// Len возвращает длину Lines
func (l Lines) Len() int {
	return len(l)
}

// Swap меняет значение элементов
func (l Lines) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Less сравниванием Lines по Key
func (l Lines) Less(i, j int) bool {
	if l[i].Column && l[j].Column {
		return l[i].Key < l[j].Key
	} else if l[i].Column && !l[j].Column {
		return false
	} else if !l[i].Column && l[j].Column {
		return false
	}
 	return l[i].Key < l[j].Key
}

func readLines(file io.Reader) ([]Line, error) {
	scanner := bufio.NewScanner(file)
	var lines []Line

	for scanner.Scan() {
		text := scanner.Text()
		key := text
		lines = append(lines, Line{Text: text, Key: key, Column: false})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func sortLines(lines []Line, column int, numerical, reverse, unique bool) []Line {
	// используем column для сортировки
	if column != 1 {
		for i := 0; i < len(lines); i++ {
			parts := strings.Fields(lines[i].Text)
			if column > 0 && column <= len(parts) {
				lines[i].Key = parts[column-1]
				lines[i].Column = true
			}
		}
	}
	// сортируем
	sort.Sort(Lines(lines))

	// применяем флаги сортировки
	if numerical {
		sort.SliceStable(lines, func(i, j int) bool {
			num1, err1 := strconv.ParseFloat(lines[i].Key, 64)
			num2, err2 := strconv.ParseFloat(lines[j].Key, 64)

			if err1 == nil && err2 == nil {
				return num1 < num2
			}

			return lines[i].Key < lines[j].Key
		})
	}

	// применяем reverse
	if reverse {
		reverseLines(lines)
	}

	// применяем unique
	if unique {
		lines = removeDuplicates(lines)
	}

	return lines
}

// reverseLines делает reverse lines
func reverseLines(lines []Line) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

// removeDuplicates удаляет дупликаты из lines
func removeDuplicates(lines []Line) []Line {
	uniqueLines := make([]Line, 0, len(lines))
	seen := make(map[string]bool)

	for _, line := range lines {
		if !seen[line.Key] {
			uniqueLines = append(uniqueLines, line)
			seen[line.Key] = true
		}
	}

	return uniqueLines
}

func main() {
	// парсим флаги
	k := flag.Int("k", 1, "sort key")
	n := flag.Bool("n", false, "sort fields value")
	r := flag.Bool("r", false, "sort reversed")
	u := flag.Bool("u", false, "save only unique keys")
	flag.Parse()

	// открываем файл
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// считываем данные из файла
	lines, err := readLines(file)
	if err != nil {
		log.Fatal(err)
	}

	// сортируем данные файла
	sortedLines := sortLines(lines, *k, *n, *r, *u)

	// выводим отсортированный файл
	for _, line := range sortedLines {
		fmt.Println(line.Text)
	}
}


