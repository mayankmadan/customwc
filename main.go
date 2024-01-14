package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func processFile(file io.Reader, processors []Processor) error {
	fileReader := bufio.NewReader(file)

	for {
		line, isPrefix, err := fileReader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error reading file")
		}
		for _, processor := range processors {
			processor.process(line, isPrefix)
		}
	}

	return nil
}

func main() {
	countBytesFlag := flag.Bool("c", false, "Count number of bytes in the file")
	countLinesFlag := flag.Bool("l", false, "Count number of lines in the file")
	countWordFlag := flag.Bool("w", false, "Count number of words in the file")

	flag.Parse()
	input := flag.Args()

	var file *os.File
	defer file.Close()

	var fileName string = ""
	var err error

	if len(input) == 0 {
		file = os.Stdin
	} else if len(input) == 1 {
		fileName = input[0]
		file, err = os.Open(input[0])
	} else {
		err = fmt.Errorf("invalid arguments")
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

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

	err = processFile(file, processors)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := ""
	for _, processor := range processors {
		output += fmt.Sprintf("%d ", processor.getValue())
	}

	fmt.Printf("%s %s\n", output, fileName)
}
