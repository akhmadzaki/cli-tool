package main

import (
	"fmt"
	"flag"
	"os"
	"log"
	"bufio"
	"io"
	"strings"
)

func main() {
	var reader *bufio.Reader
	var f *os.File
	var filePath string

	byteFlag := flag.Bool("c", false, "Count number of bytes in a file")
	lineFlag := flag.Bool("l", false, "Count number of lines in a file")
	wordFlag := flag.Bool("w", false, "Count number of words in a file")
	charFlag := flag.Bool("m", false, "Count number of characters in a file")
	flag.Parse()
	
	if len(flag.Args()) > 1 {
		log.Fatal("Too many input file")
	}

	

	if len(flag.Args()) == 0 {
		reader = bufio.NewReader(os.Stdin)
	} else {
		filePath := flag.Args()[0]
		f, err := os.Open(filePath)

		if err != nil {
			log.Fatal(err)
		}

		reader = bufio.NewReader(f)
	}
	
	defer f.Close()

	if (*byteFlag) {
		byteCount, err := countByte(reader)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(byteCount, filePath)
	} else if (*lineFlag) {
		lineCount, err := countLine(reader)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(lineCount, filePath)
	} else if (*wordFlag) {
		wordCount, err := countWord(reader)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(wordCount, filePath)
	} else if (*charFlag) {
		charCount, err := countChar(reader)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(charCount, filePath)
	} else {
		lineCount, wordCount, byteCount, err := defaultCount(reader)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(lineCount, wordCount, byteCount, filePath)
	}
}

func defaultCount(reader *bufio.Reader) (int, int, int, error) {
	cntLine, cntWord, cntByte  := 0, 0, 0

	for {
		b, err := reader.ReadBytes('\n')

		if err == io.EOF {
			break
		}

		cntLine += 1
		cntWord += len(strings.Fields(string(b)))
		cntByte += len(b)
	}

	return cntLine, cntWord, cntByte, nil
}

func countByte(reader *bufio.Reader) (int, error) {
	cnt := 0

	for {
		b, err := reader.ReadBytes('\n')

		if err == io.EOF {
			break
		}

		cnt += len(b)
	}

	return cnt, nil
}

func countLine(reader *bufio.Reader) (int, error) {
	cnt := 0

	for {
		_, err := reader.ReadBytes('\n')

		if err == io.EOF {
			break
		}

		cnt += 1
	}

	return cnt, nil
}

func countWord(reader *bufio.Reader) (int, error) {
	cnt := 0

	for {
		b, err := reader.ReadBytes('\n')

		if err == io.EOF {
			break
		}

		cnt += len(strings.Fields(string(b)))
	}

	return cnt, nil
}

func countChar(reader *bufio.Reader) (int, error) {
	cnt := 0

	for {
		b, err := reader.ReadBytes('\n')

		if err == io.EOF {
			break
		}

		cnt += strings.Count(string(b), "") - 1
	}

	return cnt, nil
}