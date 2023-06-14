package main

import (
	"fmt"
	"sort"
	"strings"
)

//Написать функцию поиска всех множеств анаграмм по словарю.
//
//
//Например:
//"пятак", "пятка" и "тяпка" - принадлежат одному множеству,
//"листок", "слиток" и "столик" - другому.
//
//
//Требования:
//Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
//Выходные данные: ссылка на мапу множеств анаграмм
//Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
//слово из множества.
//Массив должен быть отсортирован по возрастанию.
//Множества из одного элемента не должны попасть в результат.
//Все слова должны быть приведены к нижнему регистру.
//В результате каждое слово должно встречаться только один раз.

func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func findAnagrams(arr []string) map[string][]string {
	anagrams := make(map[string]map[string]struct{})

	for _, word := range arr {
		lowerWord := strings.ToLower(word)
		sortWord := sortString(lowerWord)
		if _, ok := anagrams[sortWord]; !ok {
			anagrams[sortWord] = make(map[string]struct{})
			anagrams[sortWord][lowerWord] = struct{}{}
		} else {
			anagrams[sortWord][lowerWord] = struct{}{}
		}
	}

	res := make(map[string][]string)
	for key := range anagrams {
		// множества из одного элемента
		if len(anagrams[key]) == 1 {
			continue
		}

		sorted := make([]string, 0, len(anagrams[key]))
		for k := range anagrams[key] {
			sorted = append(sorted, k)
		}
		sort.Strings(sorted)

		for _, word := range arr {
			lowerWord := strings.ToLower(word)
			if sortString(lowerWord) == key {
				res[lowerWord] = sorted
				break
			}
		}
	}

	return res
}

func main() {
	arr := []string{"пятак", "пятка" , "тяпка", "листок", "слиток" , "столик", "один", "дино", "два"}
	fmt.Println(findAnagrams(arr))
}