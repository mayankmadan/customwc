package main

import "testing"

func TestByteCountProcessor(t *testing.T) {
	fileName := "testfile.txt"
	processors := []Processor{&ByteCountProcessor{}}

	processFile(fileName, processors)

	expected := 3001
	actual := processors[0].getValue()

	if expected != actual {
		t.Errorf("Expected %d, Got %d", expected, actual)
	}
}

func TestLineCountProcessor(t *testing.T) {
	fileName := "testfile.txt"
	processors := []Processor{&LineCountProcessor{}}

	processFile(fileName, processors)

	expected := 9
	actual := processors[0].getValue()

	if expected != actual {
		t.Errorf("Expected %d, Got %d", expected, actual)
	}
}

func TestWordCountProcessor(t *testing.T) {
	fileName := "testfile.txt"
	processors := []Processor{&WordCountProcessor{}}

	processFile(fileName, processors)

	expected := 438
	actual := processors[0].getValue()

	if expected != actual {
		t.Errorf("Expected %d, Got %d", expected, actual)
	}
}
