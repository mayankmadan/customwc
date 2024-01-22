# CustomWC
A Simple wc clone for codingchallenges.fyi (https://codingchallenges.fyi/challenges/challenge-wc)

This is a custom implementation of the classic word count utility in Go. It counts bytes, lines, and words in an input file.

### Features
- Counts the number of bytes, lines and words in a file
- Supports stdin as input if no file is specified
- Custom Processor interface allows adding new metrics easily
- Clean separation of file reading and metric counting logic

### Usage
`customwc [flags] [filename]`



### Flags
- -c : Count bytes
- -l : Count lines
- -w : Count words
If no flags are specified, all metrics are counted.

### Examples
Count words only in file:

`customwc -w file.txt`



Count all metrics from stdin:

`cat file.txt | customwc`



Processors
The Processor interface defines the metric counting API:

```
type Processor interface {
  process(line []byte, isPrefix bool) 
  getValue() int
}
```


The following processors are implemented:

`ByteCountProcessor` - Counts total bytes
`LineCountProcessor` - Counts total lines
`WordCountProcessor` - Counts total words
Adding a new processor just requires implementing the Processor interface.
