package main

/*
=== Задача на распаковку ===
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.

*функция UnpackString не проходит go vet из-за пакета main
*/

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidString = errors.New("invalid string")
)

// Symbol структура символа
type Symbol struct {
	symbol bool
	digit  bool
	escape bool
	r rune
}

// SymbolType определяет тип c
func SymbolType(c rune, previous Symbol) Symbol {
	current := Symbol{}
	current.r = c

	if previous.escape {
		if unicode.IsDigit(c) || c == '\\' {
			current.symbol = true
		}
		return current
	}

	if unicode.IsDigit(c) {
		current.digit = true
	} else if c == '\\' {
		current.escape = true
	} else {
		current.symbol = true
	}

	return current
}

func UnpackString(s string) (string, error) {
	// проверка пустой строки
	if len(s) == 0 {
		return "", nil
	}

	// res результирующая строка
	res := strings.Builder{}

	// previous предыдущий элемент
	previous := Symbol{}

	// проверяем 0 элемент на валидность
	// current текущий элемент
	current := SymbolType([]rune(s)[0], previous)
	if current.digit {
		return "", ErrInvalidString
	} else if !current.escape {
		res.WriteRune([]rune(s)[0])
	}

	// удаляем 0 элемент, преобразуя s
	s = s[1:]
	for _, r := range s {
		previous = current
		current = SymbolType(r, previous)

		// записываем символ
		if current.symbol {
			res.WriteRune(current.r)
		} else if current.digit { // обработка числа
			repeat, err := strconv.Atoi(string(r))
			if err != nil {
				log.Println(err)
			}
			// запись повторений
			res.WriteString(strings.Repeat(string(previous.r), repeat - 1))
		} else if !current.escape {
			return "", ErrInvalidString
		}
	}

	// проверка на \
	if current.escape {
		return "", ErrInvalidString
	}

	// возвращаем значение строки
	return res.String(), nil
}

func main() {
	s, err := UnpackString("tes2t")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(s)
}