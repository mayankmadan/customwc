package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
)

type Processor interface {
	process(line []byte)
	getValue() int
}

type ValueGetter struct {
	value int
}

func (v ValueGetter) getValue() int {
	return v.value
}

type ByteCountProcessor struct {
	ValueGetter
}

func (processor *ByteCountProcessor) process(line []byte) {
	processor.value = processor.value + len(line)
}

type LineCountProcessor struct {
	ValueGetter
}

func (processor *LineCountProcessor) process(line []byte) {
	processor.value += 1
}

type WordCountProcessor struct {
	ValueGetter
}

func (processor *WordCountProcessor) process(line []byte) {
	processor.value += len(bytes.Fields(line))
}

func processFile(fileName string, processors []Processor) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Fatal Error: %e\n", err)
	}

	fileReader := bufio.NewReader(file)

	for {
		line, _, err := fileReader.ReadLine()
		if err != nil {
			break
		}
		for _, processor := range processors {
			processor.process(line)
		}
	}
}

func main() {
	fileName := os.Args[len(os.Args)-1]
	countBytesFlag := flag.Bool("c", false, "Count number of bytes in the file")
	countLinesFlag := flag.Bool("l", false, "Count number of lines in the file")
	countWordFlag := flag.Bool("w", false, "Count number of words in the file")

	flag.Parse()

	var processors []Processor = make([]Processor, 0, 3)

	if *countBytesFlag {
		processors = append(processors, &ByteCountProcessor{})
	}

	if *countLinesFlag {
		processors = append(processors, &LineCountProcessor{})
	}

	if *countWordFlag {
		processors = append(processors, &WordCountProcessor{})
	}

	if !*countBytesFlag && !*countLinesFlag && !*countWordFlag {
		processors = append(processors, &ByteCountProcessor{}, &WordCountProcessor{}, &LineCountProcessor{})
	}

	processFile(fileName, processors)

	output := ""
	for _, processor := range processors {
		output += fmt.Sprintf("%d ", processor.getValue())
	}

	fmt.Printf("%s %s\n", output, fileName)

}
