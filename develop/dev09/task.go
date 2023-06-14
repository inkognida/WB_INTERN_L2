package main

import (
	"github.com/opesun/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

//Реализовать утилиту wget с возможностью скачивать сайты целиком.

// download скачиваем содержимое по url
func download(url string) {
	// resp получаем ответ
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// создаем файл
	link := strings.Split(url, "/")
	fileName := link[2] + ".html"

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// копируем resp.Body в файл
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
}

func files(url string) {
	x, _ := goquery.ParseUrl(url)
	// ищем все файлы
	for _, url := range x.Find("").Attrs("href") {
		var str []string
		switch {
		case strings.Contains(url, ".png"):
			str = strings.Split(url, "/")
			if err := downloadResources(str[len(str)-1], url); err != nil {
				log.Fatal(err)
			}
		case strings.Contains(url, ".jpg"):
			str = strings.Split(url, "/")
			if err := downloadResources(str[len(str)-1], url); err != nil {
				log.Fatal(err)
			}
		case strings.Contains(url, ".css"):
			str = strings.Split(url, "/")
			if err := downloadResources(str[len(str)-1], url); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func downloadResources(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Wrong input: ./task <url>")
	}

	// скачиваем страницу
	download(os.Args[1])
	files(os.Args[1])
}