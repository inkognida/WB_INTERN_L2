package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

Стратегия — это поведенческий паттерн, выносит набор алгоритмов в собственные классы и делает их взаимозаменимыми

Применимость паттерна "Стратегия":
Когда у вас есть несколько алгоритмов, которые выполняют схожие действия, и вы хотите, чтобы клиенты могли выбирать нужный алгоритм во время выполнения.
Когда у вас есть класс с одним методом, который имеет несколько вариантов реализации, и вы хотите изолировать каждый вариант в отдельный класс.

Преимущества использования паттерна "Стратегия":
Упрощает добавление новых стратегий и изменение поведения программы без изменения существующего кода.
Позволяет избежать множественных условных операторов, так как каждая стратегия инкапсулируется в отдельном классе.
Улучшает читаемость и понятность кода, так как каждая стратегия представляет отдельный алгоритм.

Недостатки использования паттерна "Стратегия":
Увеличивает количество классов в программе из-за необходимости создания отдельного класса для каждой стратегии.
Усложняет коммуникацию между стратегиями и контекстом, если стратегии должны взаимодействовать друг с другом.

Реальные примеры использования паттерна "Стратегия" на практике:
Сортировка данных: Различные алгоритмы сортировки, такие как сортировка пузырьком, сортировка выбором и сортировка слиянием,
могут быть реализованы с помощью стратегий. Клиентский код может выбирать нужный алгоритм сортировки во время выполнения.
*/


// SearchAlgorithm интерфейс алгоритма для поиска
type SearchAlgorithm interface {
	Search(d *Cache)
}

// KeySearch поиск по ключю в цикле
type KeySearch struct {

}

// Search выполнение поиска
func (k *KeySearch) Search(c *Cache) {
	fmt.Println("key looking by using loop", c.cache)
}

// KeyExistence поиск по наличию ключа
type KeyExistence struct {

}

// Search выполнение поиска
func (ke *KeyExistence) Search(c *Cache) {
	fmt.Println("key existence check", c.cache)
}

// Cache структура данных (кэш)
type Cache struct {
	cache map[int]struct{}
	searchAlgo SearchAlgorithm
}

// NewCache создание Cache
func NewCache() *Cache {
	return &Cache{
		cache:      make(map[int]struct{}),
		searchAlgo: nil,
	}
}

// SetSearchAlgorithm установка алгоритма поиска
func (c *Cache) SetSearchAlgorithm(algo SearchAlgorithm) {
	c.searchAlgo = algo
}

// Add добавление элемента
func (c *Cache) Add(k int) {
	c.cache[k] = struct{}{}
}

// ExecuteSearch выполнение поиска
func (c *Cache) ExecuteSearch() {
	c.searchAlgo.Search(c)
}

func main() {
	// создаем кэш
	cache := NewCache()
	// создаем алгоритмы поиска
	keySearch := &KeySearch{}
	keyExistence := &KeyExistence{}

	// добавляем элементы
	cache.Add(1)
	cache.Add(2)

	// устанавливаем алгоритм поиска для cache
	cache.SetSearchAlgorithm(keySearch)
	// выполняем поиск
	cache.ExecuteSearch()

	cache.SetSearchAlgorithm(keyExistence)
	cache.ExecuteSearch()
}