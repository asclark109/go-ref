package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestTAS(t *testing.T) {

	var tests = []struct {
		lock            string
		incrementAmount string
		threads         string
		expected        string
	}{
		{"tas", "100", "2", "200"},
		{"tas", "100", "4", "400"},
		{"tas", "100", "6", "600"},
		{"tas", "100", "8", "800"},
		{"tas", "1000", "12", "12000"},
		{"tas", "1000", "16", "16000"},
		{"tas", "1000", "24", "24000"},
		{"tas", "1000", "2", "2000"},
		{"tas", "10000", "4", "40000"},
		{"tas", "10000", "6", "60000"},
		{"tas", "10000", "8", "80000"},
		{"tas", "10000", "16", "160000"},
		{"tas", "1000000", "8", "8000000"},
		{"tas", "1000000", "16", "16000000"},
		{"tas", "1000000", "24", "24000000"},
		{"tas", "1000000", "32", "32000000"},
	}
	for num, test := range tests {
		testname := fmt.Sprintf("T=%v", num)
		t.Run(testname, func(t *testing.T) {
			var err error
			cmd := exec.Command("go", "run", "hw3/experiment", test.lock, test.incrementAmount, test.threads)
			out, err := cmd.Output()
			sOut := strings.TrimSpace(string(out))

			if err != nil || test.expected != sOut {
				t.Errorf("\nRan:%s\nExpected:%s\nGot:%s", cmd, test.expected, sOut)
			}
		})
	}
}
func TestSync(t *testing.T) {

	var tests = []struct {
		lock            string
		incrementAmount string
		threads         string
		expected        string
	}{
		{"sync", "100", "2", "200"},
		{"sync", "100", "4", "400"},
		{"sync", "100", "6", "600"},
		{"sync", "100", "8", "800"},
		{"sync", "1000", "12", "12000"},
		{"sync", "1000", "16", "16000"},
		{"sync", "1000", "24", "24000"},
		{"sync", "1000", "2", "2000"},
		{"sync", "10000", "4", "40000"},
		{"sync", "10000", "6", "60000"},
		{"sync", "10000", "8", "80000"},
		{"sync", "10000", "16", "160000"},
		{"sync", "1000000", "8", "8000000"},
		{"sync", "1000000", "16", "16000000"},
		{"sync", "1000000", "24", "24000000"},
		{"sync", "1000000", "32", "32000000"},
	}
	for num, test := range tests {
		testname := fmt.Sprintf("T=%v", num)
		t.Run(testname, func(t *testing.T) {
			var err error
			cmd := exec.Command("go", "run", "hw3/experiment", test.lock, test.incrementAmount, test.threads)
			out, err := cmd.Output()
			sOut := strings.TrimSpace(string(out))

			if err != nil || test.expected != sOut {
				t.Errorf("\nRan:%s\nExpected:%s\nGot:%s", cmd, test.expected, sOut)
			}
		})
	}
}
func TestTTAS(t *testing.T) {

	var tests = []struct {
		lock            string
		incrementAmount string
		threads         string
		expected        string
	}{
		{"ttas", "100", "2", "200"},
		{"ttas", "100", "4", "400"},
		{"ttas", "100", "6", "600"},
		{"ttas", "100", "8", "800"},
		{"ttas", "1000", "12", "12000"},
		{"ttas", "1000", "16", "16000"},
		{"ttas", "1000", "24", "24000"},
		{"ttas", "1000", "2", "2000"},
		{"ttas", "10000", "4", "40000"},
		{"ttas", "10000", "6", "60000"},
		{"ttas", "10000", "8", "80000"},
		{"ttas", "10000", "16", "160000"},
		{"ttas", "1000000", "8", "8000000"},
		{"ttas", "1000000", "16", "16000000"},
		{"ttas", "1000000", "24", "24000000"},
		{"ttas", "1000000", "32", "32000000"},
	}
	for num, test := range tests {
		testname := fmt.Sprintf("T=%v", num)
		t.Run(testname, func(t *testing.T) {
			var err error
			cmd := exec.Command("go", "run", "hw3/experiment", test.lock, test.incrementAmount, test.threads)
			out, err := cmd.Output()
			sOut := strings.TrimSpace(string(out))

			if err != nil || test.expected != sOut {
				t.Errorf("\nRan:%s\nExpected:%s\nGot:%s", cmd, test.expected, sOut)
			}
		})
	}
}
func TestEB(t *testing.T) {

	var tests = []struct {
		lock            string
		minDelay        string
		maxDelay        string
		incrementAmount string
		threads         string
		expected        string
	}{
		{"eb", "32", "1024", "100", "2", "200"},
		{"eb", "32", "1024", "100", "4", "400"},
		{"eb", "32", "1024", "100", "6", "600"},
		{"eb", "32", "1024", "100", "8", "800"},
		{"eb", "32", "1024", "1000", "12", "12000"},
		{"eb", "32", "1024", "1000", "16", "16000"},
		{"eb", "32", "1024", "1000", "24", "24000"},
		{"eb", "32", "1024", "1000", "2", "2000"},
		{"eb", "32", "1024", "10000", "4", "40000"},
		{"eb", "32", "1024", "10000", "6", "60000"},
		{"eb", "32", "1024", "10000", "8", "80000"},
		{"eb", "32", "1024", "10000", "16", "160000"},
		{"eb", "32", "1024", "1000000", "8", "8000000"},
		{"eb", "32", "1024", "1000000", "16", "16000000"},
		{"eb", "32", "1024", "1000000", "24", "24000000"},
		{"eb", "32", "1024", "1000000", "32", "32000000"},
	}
	for num, test := range tests {
		testname := fmt.Sprintf("T=%v", num)
		t.Run(testname, func(t *testing.T) {
			var err error
			cmd := exec.Command("go", "run", "hw3/experiment", test.lock, test.minDelay, test.maxDelay, test.incrementAmount, test.threads)
			out, err := cmd.Output()
			sOut := strings.TrimSpace(string(out))

			if err != nil || test.expected != sOut {
				t.Errorf("\nRan:%s\nExpected:%s\nGot:%s", cmd, test.expected, sOut)
			}
		})
	}
}
func TestF(t *testing.T) {

	var tests = []struct {
		lock            string
		minDelay        string
		maxDelay        string
		incrementAmount string
		threads         string
		expected        string
	}{
		{"ebf", "32", "1024", "100", "2", "200"},
		{"ebf", "32", "1024", "100", "4", "400"},
		{"ebf", "32", "1024", "100", "6", "600"},
		{"ebf", "32", "1024", "100", "8", "800"},
		{"ebf", "32", "1024", "1000", "12", "12000"},
		{"ebf", "32", "1024", "1000", "16", "16000"},
		{"ebf", "32", "1024", "1000", "24", "24000"},
		{"ebf", "32", "1024", "1000", "2", "2000"},
		{"ebf", "32", "1024", "10000", "4", "40000"},
		{"ebf", "32", "1024", "10000", "6", "60000"},
		{"ebf", "32", "1024", "10000", "8", "80000"},
		{"ebf", "32", "1024", "10000", "16", "160000"},
		{"ebf", "32", "1024", "1000000", "8", "8000000"},
		{"ebf", "32", "1024", "1000000", "16", "16000000"},
		{"ebf", "32", "1024", "1000000", "24", "24000000"},
		{"ebf", "32", "1024", "1000000", "32", "32000000"},
	}
	for num, test := range tests {
		testname := fmt.Sprintf("T=%v", num)
		t.Run(testname, func(t *testing.T) {
			var err error
			cmd := exec.Command("go", "run", "hw3/experiment", test.lock, test.minDelay, test.maxDelay, test.incrementAmount, test.threads)
			out, err := cmd.Output()
			sOut := strings.TrimSpace(string(out))

			if err != nil || test.expected != sOut {
				t.Errorf("\nRan:%s\nExpected:%s\nGot:%s", cmd, test.expected, sOut)
			}
		})
	}
}
