package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

//Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).
//
//
//Реализовать поддержку утилитой следующих ключей:
//-A - "after" печатать +N строк после совпадения
//-B - "before" печатать +N строк до совпадения
//-C - "context" (A+B) печатать ±N строк вокруг совпадения
//-c - "count" (количество строк)
//-i - "ignore-case" (игнорировать регистр)
//-v - "invert" (вместо совпадения, исключать)
//-F - "fixed", точное совпадение со строкой, не паттерн
//-n - "line num", напечатать номер строки


// Flags используемые флаги
type Flags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

// grep поиск строк по паттерну
func grep(flags Flags, pattern string, lines []string) {
	// обработка C флага
	count := len(lines)
	if flags.context > 0 {
		flags.before, flags.after = flags.context, flags.context
	}

	for i, str := range lines {
		// обработка i флага
		if flags.ignoreCase {
			str = strings.ToLower(str)
		}
		match := strings.Contains(str, pattern)

		// обработка v флага
		if flags.invert {
			match = !match
		}

		// обработка F флага
		if flags.fixed && !strings.EqualFold(str, pattern) {
			match = false
		}

		// печатаем строки в случае match
		if match {
			l := math.Max(0, float64(i-flags.before))
			r := math.Min(float64(len(lines)-1), float64(i+flags.after))
			for j := l; j <= r; j++ {
				if count >= 1 {
					if flags.lineNum {
						if !flags.count {
							fmt.Printf("%v:%v\n", j+1, lines[int(j)])
						}
						count--
					} else {
						if !flags.count {
							fmt.Printf("%v\n", lines[int(j)])
						}
						count--
					}
				} else {
					log.Fatalln(count)
				}
			}
		}
	}

	// c флаг
	if flags.count {
		fmt.Printf("%v\n", len(lines)-count)
	}
}

// readLines считываем данные
func readLines(file io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(file)
	var flags []string

	for scanner.Scan() {
		text := scanner.Text()
		flags = append(flags, text)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return flags, nil
}

// getPattern возвращает паттерн
func getPattern(flags Flags) string {
	pattern := flag.Arg(0)

	if flags.ignoreCase {
		pattern = strings.ToLower(pattern)
	}

	return pattern
}

func getFlags() Flags {
	A := flag.Int("A", 0, `output n flags after match`)
	B := flag.Int("B", 0, `output n flags before match`)
	C := flag.Int("C", 0, `output n flags around match`)
	c := flag.Bool("c", false, `flags count`)
	i := flag.Bool("i", false, `ignore case`)
	v := flag.Bool("v", false, `avoid`)
	F := flag.Bool("F", false, `exact match`)
	n := flag.Bool("n", false, `output line number`)
	flag.Parse()

	// заполнение структуру флагов
	flags := Flags{
		after:      *A,
		before:     *B,
		context:    *C,
		count:      *c,
		ignoreCase: *i,
		invert:     *v,
		fixed:      *F,
		lineNum:    *n,
	}

	return flags
}

func main() {
	// парсим флаги
	flags := getFlags()

	// открываем файл
	file, err := os.Open(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// считываем данные из файла
	lines, err := readLines(file)
	if err != nil {
		log.Fatal(err)
	}

	// ищем строки по паттерну с флагами
	grep(flags, getPattern(flags), lines)
}