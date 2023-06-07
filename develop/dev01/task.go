package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)
/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	// получаем точное время
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Println("Got error:", err)
		os.Exit(1)
	}

	// выводим результат
	fmt.Println("Current time", time.Now())
	fmt.Println("Exact time", exactTime)
}