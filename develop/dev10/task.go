package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Реализовать простейший telnet-клиент.
//
//Примеры вызовов:
//go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123
//
//
//Требования:
//Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
//После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
//Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
//При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера,
// программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout

func telnet(host string, port string, timeout *time.Duration) {
	// Установление соединения
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), *timeout)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// Обработка сигналов для завершения программы
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		conn.Close()
	}()

	// Копирование данных между stdin и сокетом
	go func() {
		_, err = io.Copy(conn, os.Stdin)
		if err != nil {
			fmt.Println("Error copying from stdin to the server:", err)
		}
	}()

	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		fmt.Println("Error copying from the server to stdout:", err)
	}
}

func main() {
	// Получение аргументов командной строки
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	// Проверка наличия аргументов хоста и порта
	if len(flag.Args()) != 2 {
		log.Fatalln("Usage: go-telnet [--timeout=<duration>] host port")
	}
	host := flag.Arg(0)
	port := flag.Arg(1)

	// соединение
	telnet(host, port, timeout)
}