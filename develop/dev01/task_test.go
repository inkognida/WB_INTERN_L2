package main

import (
	"bytes"
	"log"
	"os"
	"testing"
)

// TestMain структура для тестов
func TestMain(m *testing.M) {
	log.SetOutput(&bytes.Buffer{})

	code := m.Run()

	os.Exit(code)
}


