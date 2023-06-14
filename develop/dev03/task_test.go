package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestReadLines(t *testing.T) {
	input := "line1\nline2\nline3\n"
	expected := []Line{
		{Text: "line1", Key: "line1"},
		{Text: "line2", Key: "line2"},
		{Text: "line3", Key: "line3"},
	}
	reader := strings.NewReader(input)

	lines, err := readLines(reader)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(lines) != len(expected) {
		t.Fatalf("Unexpected number of lines. Expected: %d, Got: %d", len(expected), len(lines))
	}

	for i, line := range lines {
		if line.Text != expected[i].Text || line.Key != expected[i].Key {
			t.Errorf("Unexpected line. Expected: %+v, Got: %+v", expected[i], line)
		}
	}
}

func TestSortLinesKey(t *testing.T) {
	fullCmd := "sort -k 2 test.txt"
	cmd := exec.Command("bash", "-c", fullCmd)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		t.Error(err)
	}

	expected := stdout.String()

	file, err := os.Open("test.txt")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	lines, err := readLines(file)
	sorted := sortLines(lines, 2, false, false, false)

	var res string
	for _, v := range sorted {
		res += v.Text + "\n"
	}

	if len(res) != len(expected) {
		t.Fatalf("Unexpected number of lines. Expected: %d, Got: %d", len(expected), len(sorted))
	}

	if res != expected {
		t.Fatalf("Unexpected compare %v \n %v", res, expected)
	}
}

func TestReverseLines(t *testing.T) {
	fullCmd := "sort -r test.txt"
	cmd := exec.Command("bash", "-c", fullCmd)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		t.Error(err)
	}

	expected := stdout.String()

	file, err := os.Open("test.txt")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	lines, err := readLines(file)
	sorted := sortLines(lines, 1, false, true, false)

	var res string
	for _, v := range sorted {
		res += v.Text + "\n"
	}

	if len(res) != len(expected) {
		t.Fatalf("Unexpected number of lines. Expected: %d, Got: %d", len(expected), len(sorted))
	}

	if res != expected {
		t.Fatalf("Unexpected compare %v \n %v", res, expected)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	fullCmd := "sort -u test.txt"
	cmd := exec.Command("bash", "-c", fullCmd)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		t.Error(err)
	}

	expected := stdout.String()

	file, err := os.Open("test.txt")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	lines, err := readLines(file)
	sorted := sortLines(lines, 1, false, false, true)

	var res string
	for _, v := range sorted {
		res += v.Text + "\n"
	}

	if len(res) != len(expected) {
		t.Fatalf("Unexpected number of lines. Expected: %d, Got: %d", len(expected), len(sorted))
	}

	if res != expected {
		t.Fatalf("Unexpected compare %v \n %v", res, expected)
	}
}

func TestNumerical(t *testing.T) {
	fullCmd := "sort -n test.txt"
	cmd := exec.Command("bash", "-c", fullCmd)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		t.Error(err)
	}

	expected := stdout.String()

	file, err := os.Open("test.txt")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	lines, err := readLines(file)
	sorted := sortLines(lines, 1, true, false, false)

	var res string
	for _, v := range sorted {
		res += v.Text + "\n"
	}

	if len(res) != len(expected) {
		t.Fatalf("Unexpected number of lines. Expected: %d, Got: %d", len(expected), len(sorted))
	}

	if res != expected {
		t.Fatalf("Unexpected compare %v \n %v", res, expected)
	}
}
