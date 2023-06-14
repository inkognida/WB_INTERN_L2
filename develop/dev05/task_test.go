package main

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestGrep(t *testing.T) {
	fullCmd := []string{
		"grep -A 2 aboba test.txt",
		"grep -B 2 aboba test.txt",
		"grep -C 2 aboba test.txt",
		"grep -c ab test.txt",
		"grep -i ab test.txt",
		"grep -n ab test.txt",
	}
	originalCmd := []string{
		"go build task.go && ./task -A 2 aboba test.txt",
		"go build task.go && ./task -B 2 aboba test.txt",
		"go build task.go && ./task -C 2 aboba test.txt",
		"go build task.go && ./task -c ab test.txt",
		"go build task.go && ./task -i ab test.txt",
		"go build task.go && ./task -n ab test.txt",
	}

	for i := range fullCmd {
		t.Run("grep", func(t *testing.T) {
			cmd := exec.Command("bash", "-c", fullCmd[i])
			origCmd := exec.Command("bash", "-c", originalCmd[i])

			var stdout bytes.Buffer
			cmd.Stdout = &stdout

			var origStdout bytes.Buffer
			origCmd.Stdout = &origStdout

			err := cmd.Run()
			if err != nil {
				t.Error(err)
			}

			err = origCmd.Run()
			if err != nil {
				t.Error(err)
			}

			expected := stdout.String()
			original := origStdout.String()
			if expected != original {
				t.Errorf("Unexpected compare %v \n %v", expected, original)
				t.Log(fullCmd[i], originalCmd[i])
			}
		})
	}
}
