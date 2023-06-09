package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты,
позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их, а также поддерживать отмену операций.


Применимость паттерна "Команда":
Когда необходимо параметризовать объекты с операциями и выполнять их асинхронно.
Когда нужна возможность отмены операций и управления историей выполненных команд.
Когда необходимо реализовать очередь операций или планировщик задач.

Преимущества использования паттерна "Команда":
Расширяемость: Новые команды могут быть легко добавлены без изменения существующих классов.
Отделение вызывающего объекта от получателя операции: Клиент не зависит от конкретного получателя и операции, которую он выполняет.
Управление историей и отмена операций: Паттерн "Команда" позволяет легко реализовать отмену операций и хранение истории выполненных команд.

Недостатки использования паттерна "Команда":
Увеличение количества классов: Внедрение паттерна "Команда" может привести к созданию большого числа классов, особенно при использовании отдельных команд для каждого типа операции.
Дополнительные накладные расходы: Использование объектов команд может привести к дополнительным накладным расходам памяти и производительности.
*/

// Command команда для выполнения
type Command interface {
	Execute()
}

// TurnOn команда для включения
type TurnOn struct {
	receiver *Receiver
}

func (on *TurnOn) Execute() {
	on.receiver.TurnOn()
}

// TurnOff команда для выключения
type TurnOff struct {
	receiver *Receiver
}

func (off *TurnOff) Execute() {
	off.receiver.TurnOff()
}

// Receiver получатель команд для выполнения
type Receiver struct {
}

func (r *Receiver) TurnOn() {
	fmt.Println("Turn on")
}

func (r *Receiver) TurnOff() {
	fmt.Println("Turn off")
}

// Invoker представляет инициатора команд.
type Invoker struct {
	command Command
}

// SetCommand устанавливает команду для выполнения
func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

// ExecuteCommand выполняет команду
func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	// получатель
	receiver := &Receiver{}

	// команды для включения/выключения
	on := &TurnOn{receiver: receiver}
	off := &TurnOff{receiver: receiver}

	// инициатор выполнения команд
	invoker := &Invoker{}
	// устанавливаем команда для выполнения
	invoker.SetCommand(on)
	// выполняем команду
	invoker.ExecuteCommand()


	invoker.SetCommand(off)
	invoker.ExecuteCommand()
}
