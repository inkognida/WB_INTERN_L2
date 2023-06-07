package main

import "fmt"

/*

	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

Применимость паттерна:
алгоритм создания сложного объекта не должен зависеть от того,
из каких частей состоит объект и как они стыкуются между собой;
процесс конструирования должен обеспечивать различные представления конструируемого объекта.

Преимущества паттерна:
- позволяет изменять внутреннее представление продукта;
изолирует код, реализующий конструирование и представление;
- дает более тонкий контроль над процессом конструирования.

Недостатки использования паттерна "Строитель":
- увеличение количества классов в системе.
- затраты на создание дополнительных классов и методов.

Пример на объекте Product:
*/


// Product конечный продукт
type Product struct {
	Design string
	Analyse string
	Ad string
	Soft string
}

// Builder интерфейс для "строительства"
type Builder interface {
	makeDesign()
	makeAnalyse()
	makeAd()
	makeSoft()
	getProduct() Product
}

// AppBuilder билдер для создания приложения
type AppBuilder struct {
	app Product
}

// makeDesign создаем дизайн
func (a *AppBuilder) makeDesign() {
	a.app.Design = "new design"
}

// makeDesign создаем аналитику
func (a *AppBuilder) makeAnalyse() {
	a.app.Analyse = "new analyse"
}

// makeDesign создаем рекламу
func (a *AppBuilder) makeAd() {
	a.app.Ad = "new ad"
}

// makeDesign создаем софт
func (a *AppBuilder) makeSoft() {
	a.app.Soft = "new soft"
}

// getProduct получаем продукт
func (a *AppBuilder) getProduct() Product {
	return a.app
}

// NewAppBuilder создает новый строитель приложения
func NewAppBuilder() *AppBuilder {
	return &AppBuilder{}
}

// AppDirector "директор" строительства приложения
type AppDirector struct {
	builder Builder
}

// Develop разрабатывает приложение
func (d *AppDirector) Develop() Product {
	d.builder.makeAnalyse()
	d.builder.makeDesign()
	d.builder.makeAd()
	d.builder.makeSoft()

	return d.builder.getProduct()
}

// NewAppDirector создает нового директора приложения
func NewAppDirector(builder Builder) *AppDirector {
	return &AppDirector{
		builder: builder,
	}
}

func main() {
	// создаем builder
	builder := NewAppBuilder()

	// создаем директора для builder
	director := NewAppDirector(builder)

	// создаем app (develop)
	app := director.Develop()
	fmt.Println(app)
}
