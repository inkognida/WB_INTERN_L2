package main

import "fmt"

/*
Фасад — это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку.

Применимость паттерна "Фасад":
-когда требуется упростить сложную систему, выделив ее ключевой функционал через унифицированный интерфейс.
-когда нужно скрыть сложность и детали взаимодействия между классами и подсистемами.
-когда необходимо предоставить простой и удобный интерфейс для работы с подсистемой.

Преимущества использования паттерна "Фасад":
- упрощение использования сложной системы путем предоставления унифицированного интерфейса.
- сокрытие деталей и сложности взаимодействия между классами и подсистемами.
- повышение уровня абстракции и уменьшение связанности между компонентами системы.
- улучшение читаемости и поддерживаемости кода.

Недостатки использования паттерна "Фасад":
- введение дополнительного уровня абстракции может снизить производительность в некоторых случаях.
- если требуется большая гибкость и возможность изменения подсистемы, паттерн "Фасад" может стать слишком ограничивающим.
*/

// JsonReader ридер для json файлов
type JsonReader struct {

}

// Log печатает лог JsonReader
func (j *JsonReader) Log() {
	fmt.Println("Some json log")
}

// XmlReader ридер для xml файлов
type XmlReader struct {

}

// Log печатает лог XmlReader
func (x *XmlReader) Log() {
	fmt.Println("Some xml log")
}

// Reader общая
type Reader struct {
	readerJson *JsonReader
	readerXml  *XmlReader
}

func NewReader() *Reader {
	return &Reader{
		readerJson: &JsonReader{},
		readerXml:  &XmlReader{},
	}
}

func (r *Reader) Read() {
	r.readerJson.Log()
	r.readerXml.Log()
}


func main() {
	reader := NewReader()
	reader.Read()
}