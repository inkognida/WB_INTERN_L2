package main

import "fmt"

/*
		Реализовать паттерн «состояние».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

Применимость паттерна "Состояние":
Когда поведение объекта зависит от его состояния, и объект должен изменять свое поведение во время выполнения программы.
Когда у вас есть множество условных операторов, которые определяют поведение объекта в разных состояниях, 
и вы хотите избежать большого количества вложенных условных операторов.

Преимущества использования паттерна "Состояние":
Изолирует различное поведение объекта в отдельные классы-состояния, что делает код более читаемым и поддерживаемым.
Упрощает добавление новых состояний и изменение поведения объекта без изменения существующего кода.
Улучшает расширяемость системы, так как новые состояния могут быть легко добавлены без модификации других состояний или контекста.

Недостатки использования паттерна "Состояние":
Может привести к увеличению количества классов в системе из-за необходимости создания классов для каждого состояния и контекста.
Сложность управления состояниями может возрасти, особенно если состояния должны взаимодействовать друг с другом или переходы между 
состояниями зависят от внешних условий.

Реальные примеры использования паттерна "Состояние" на практике:
Реализация платежных систем: Когда платежная система находится в разных состояниях, таких как "ожидание оплаты", "обработка платежа" и "платеж завершен", 
каждое состояние может быть представлено отдельным классом-состоянием с соответствующим поведением.
*/

// DriveMode интерфейс с типом движениям
type DriveMode interface {
	Drive()
}

type Car struct {
	drivingMode DriveMode
}

// Drive движение с определенным типом
func (c *Car) Drive() {
	c.drivingMode.Drive()
}

// SetDrivingMode установка типа движения
func (c *Car) SetDrivingMode(mode DriveMode) {
	c.drivingMode = mode
}

// NewCar создание новой машины
func NewCar() *Car {
	return &Car{nil}
}

// SlowMode медленный тип
type SlowMode struct {

}

// Drive медленное движение
func (s *SlowMode) Drive() {
	fmt.Println("Driving slow")
}

// SlowMode быстрый тип
type FastMode struct {

}

func (f *FastMode) Drive() {
	fmt.Println("Driving fast")
}

func main()  {
	// создаем машину
	car := NewCar()

	// создаем типы движения
	slowM := &SlowMode{}
	fastM := &FastMode{}

	// устанавливаем тип движения
	car.SetDrivingMode(slowM)
	// выполняем движение
	car.Drive()

	car.SetDrivingMode(fastM)
	car.Drive()

}