package main

import (
	"fmt"
	"strings"
)

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

Цепочка обязанностей — это поведенческий паттерн, позволяющий передавать запрос по цепочке потенциальных обработчиков,
пока один из них не обработает запрос.



Применимость паттерна "Цепочка вызовов":
Когда имеется несколько объектов, которые могут обработать запрос, и неизвестно заранее, какой из них это будет.
Когда нужно передать запрос через последовательность обработчиков до тех пор, пока один из них не обработает его.
Когда необходимо динамически указывать набор объектов, способных обработать запрос.

Преимущества использования паттерна "Цепочка вызовов":
Уменьшение связанности: Паттерн позволяет избежать прямой зависимости между отправителем запроса и получателем, что облегчает модификацию и расширение системы.
Гибкость и расширяемость: Можно легко добавлять и изменять обработчики в цепочке без внесения изменений в клиентский код.
Возможность отмены обработки: Клиентский код может прервать обработку запроса, если ни один из обработчиков не может его обработать.

Недостатки использования паттерна "Цепочка вызовов":
Запрос может быть необработанным: Если ни один из обработчиков не может обработать запрос, он может остаться необработанным без явной обработки.

Реальные примеры использования паттерна "Цепочка вызовов" на практике:
Обработка событий: Цепочка обработчиков может использоваться для обработки событий, где каждый обработчик проверяет,
может ли он обработать событие, и передает его следующему обработчику в цепочке.
*/

// Request запрос на исполнение
type Request struct {
	Body string
}

// Handler обработчик запросов
type Handler interface {
	Handle(request Request)
	NextHandler(handler Handler)
}

// UserHandler обработчик запросов для пользователя
type UserHandler struct {
	nextHandler Handler
}

// NextHandler устанавливаем обработчик для дальнейших запросов
func (l *UserHandler) NextHandler(handler Handler) {
	l.nextHandler = handler
}

// Handle обрабатываем запрос
func (l *UserHandler) Handle(request Request) {
	if request.Body != "" && strings.EqualFold(request.Body, "user") {
		fmt.Println("Login done")
	} else if l.nextHandler != nil {
		l.nextHandler.Handle(request)
	} else {
		fmt.Println("No handler for this")
	}
}

// RootHandler обработчик запросов для рута
type RootHandler struct {
	nextHandler Handler
}

func (a *RootHandler) NextHandler(handler Handler) {
	a.nextHandler = handler
}

func (a *RootHandler) Handle(request Request) {
	if request.Body != "" && strings.EqualFold(request.Body, "root") {
		fmt.Println("Root done")
	} else if a.nextHandler != nil {
		a.nextHandler.Handle(request)
	} else {
		fmt.Println("No handler for this")
	}
}

func main() {
	// обработчик запросов
	handler := &UserHandler{nextHandler: &RootHandler{nextHandler: nil}}

	// обрабатываем запрос для пользователя
	handler.Handle(Request{Body: "user"})

	// обрабатываем запрос для рута
	handler.Handle(Request{Body: "root"})
}