package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"github.com/mitchellh/go-ps"
)

//Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
//
//
//- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
//- pwd - показать путь до текущего каталога
//- echo <args> - вывод аргумента в STDOUT
//- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
//- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*
//
//Так же требуется поддерживать функционал fork/exec-команд
//
//
//Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).
//
//
//*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
//в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
//и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор,
//пока не будет введена команда выхода (например \quit).

func cd(args string) {
	err := os.Chdir(args)
	if err != nil {
		log.Println(err)
	}
}

func pwd(out io.Writer) {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	_, _ = fmt.Fprintln(out, dir)
}

func echo(args string,out io.Writer) {
	_, _ = fmt.Fprintln(out, strings.Replace(args, "echo", "", 1))
}

func kill(args string,out io.Writer) {
	pid, err := strconv.Atoi(args)
	if err != nil {
		_, _ = fmt.Fprintln(out,err)
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintln(out, err)
	}
	if err = proc.Kill(); err != nil {
		_, _ = fmt.Fprintln(out, err)
	}
}

func psCmd(out io.Writer) {
	sliceProc, _ := ps.Processes()
	for _, proc := range sliceProc {
		_, _ = fmt.Fprintf(out,"Process : %v id: %v\n", proc.Executable(), proc.Pid())
	}
}


func execute(input string, in io.Reader, out io.Writer) error {
	commands := strings.Split(strings.Trim(input, "\n"), "|")

	// выполняем команду
	for i, cmd := range commands {
		args := strings.Split(strings.TrimSpace(cmd), " ")
		// формируем вывод для последней команды
		if i == len(commands)-1 {
			out = os.Stdout
		}
		switch args[0] {
		case "cd":
			cd(args[1])
		case "pwd":
			pwd(out)
		case "echo":
			echo(args[1],out)
		case "kill":
			kill(args[1],out)
		case "ps":
			psCmd(out)
		case `\quit`:
			_, _ = fmt.Fprintln(out,"Exit")
			os.Exit(0)
		default:
			ex := exec.Command(args[0], args[1:]...)
			ex.Stdin = in
			ex.Stdout = out

			err := ex.Run()
			if err != nil {
				_, _ = fmt.Fprintln(out, err.Error())
			}
		}
	}

	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// считываем данные
	for scanner.Scan() {
		// формируем in out
		var in, out bytes.Buffer
		scan := bufio.NewScanner(os.Stdin)
		in.Write(scan.Bytes())

		// выполняем команды
		if err := execute(scanner.Text(), &in, &out); err != nil {
			log.Fatalln(err)
		}

		// меняем in, out для пайпов
		in = out
		out.Reset()
	}
}