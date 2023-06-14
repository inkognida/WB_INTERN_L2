package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Реализовать утилиту аналог консольной команды cut (man cut).
//Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.
//
//Реализовать поддержку утилитой следующих ключей:
//-f - "fields" - выбрать поля (колонки)
//-d - "delimiter" - использовать другой разделитель
//-s - "separated" - только строки с разделителем

func cut(fields []string, delim string, sep bool) {
	scanner := bufio.NewScanner(os.Stdin)

	// итерируемся по вводу
	for scanner.Scan() {
		line := scanner.Text()
		// проверяем -s флаг
		if sep && !strings.Contains(line, delim) {
			continue
		}

		fieldsArr := strings.Split(line, delim)

		// формируем вывод
		var outputFields []string
		for _, f := range fields {
			index, err := strconv.Atoi(f)
			if err != nil {
				log.Fatalln(err)
			}

			// проверяем индекс
			if index >= 1 && index <= len(fieldsArr) {
				outputFields = append(outputFields, fieldsArr[index-1])
			}
		}

		fmt.Println(strings.Join(outputFields, delim))
	}
}

func main() {
	// парсим флаги
	f := flag.String("f", "", "fields")
	d := flag.String("d", "", "delimiter")
	s := flag.Bool("s", false, "separated")
	flag.Parse()

	// сплитим значения f
	fields := strings.Split(*f, ",")

	// утилита cut
	cut(fields, *d, *s)
}

