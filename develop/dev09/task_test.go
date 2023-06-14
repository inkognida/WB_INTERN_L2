package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestWget(t *testing.T) {
	fullCmd := []string{
		"wget google.com",
	}
	originalCmd := []string{
		"go run task.go https://www.google.com/",
	}

	for i := range fullCmd {
		t.Run("wget", func(t *testing.T) {
			cmd := exec.Command("bash", "-c", fullCmd[i])
			origCmd := exec.Command("bash", "-c", originalCmd[i])

			err := cmd.Run()
			if err != nil {
				t.Error(err)
			}

			err = origCmd.Run()
			if err != nil {
				t.Error(err)
			}

			expected, _ := ioutil.ReadFile("index.html")
			original, _ := ioutil.ReadFile("www.google.com.html")
			if !bytes.Equal(expected, original) {
				t.Log(fullCmd[i], originalCmd[i])
			}
		})
	}
}
