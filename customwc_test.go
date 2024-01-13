package main

import (
	"os"
	"testing"
)

func TestByteCountProcessor(t *testing.T) {
	fileName := "testfile.txt"
	processors := []Processor{&ByteCountProcessor{}}
	file, _ := os.Open(fileName)

	processFile(file, processors)

	expected := 3001
	actual := processors[0].getValue()

	file.Close()

	if expected != actual {
		t.Errorf("Expected %d, Got %d", expected, actual)
	}
}

func TestLineCountProcessor(t *testing.T) {
	fileName := "testfile.txt"
	processors := []Processor{&LineCountProcessor{}}
	file, _ := os.Open(fileName)

	processFile(file, processors)

	expected := 9
	actual := processors[0].getValue()

	file.Close()

	if expected != actual {
		t.Errorf("Expected %d, Got %d", expected, actual)
	}
}

func TestWordCountProcessor(t *testing.T) {
	fileName := "testfile.txt"
	processors := []Processor{&WordCountProcessor{}}
	file, _ := os.Open(fileName)

	processFile(file, processors)

	expected := 438
	actual := processors[0].getValue()

	file.Close()

	if expected != actual {
		t.Errorf("Expected %d, Got %d", expected, actual)
	}
}
