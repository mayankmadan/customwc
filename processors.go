package main

import "bytes"

type Processor interface {
	process(line []byte, isPrefix bool)
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

func (processor *ByteCountProcessor) process(line []byte, isPrefix bool) {
	processor.value = processor.value + len(line)
}

type LineCountProcessor struct {
	ValueGetter
}

func (processor *LineCountProcessor) process(line []byte, isPrefix bool) {
	if !isPrefix {
		processor.value += 1
	}
}

type WordCountProcessor struct {
	ValueGetter
}

func (processor *WordCountProcessor) process(line []byte, isPrefix bool) {
	processor.value += len(bytes.Fields(line))
}
