package main 

import (
	"fmt"
	"flag"
	"errors"
	"log"
	"os"
	"bufio"
	"io"
	_"strings"
	"strconv"
)

func main() {
	fieldFlag := flag.String("f", "1", "Get the n-th field from a line")
	delimiterFlag := flag.String("d", string('\t'), "Split input by this delimited")
	flag.Parse()

	var reader *bufio.Reader
	var f *os.File

	if len(flag.Args()) == 0 ||  flag.Args()[0] == "-" {
		reader = bufio.NewReader(os.Stdin)
	} else {
		f, err := os.Open(flag.Args()[0])

		if err != nil {
			log.Fatal(errors.New("Failed to open input file"))
		}

		reader = bufio.NewReader(f)
	}

	defer f.Close()

	fmt.Println(*fieldFlag)
	fmt.Println(*delimiterFlag)

	for {
		b, err := reader.ReadBytes('\n')

		fmt.Println(err)

		if err == io.EOF {
			break
		}

		
		fmt.Println(b)

		// var splittedFieldFlag []string

		// if strings.Contains(*fieldFlag, ",") {
		// 	*fieldFlag = strings.ReplaceAll(*fieldFlag, " ", "")
		// 	splittedFieldFlag = strings.Split(*fieldFlag, ",")
		// } else if strings.Contains(*fieldFlag, " ") {
		// 	splittedFieldFlag = strings.Fields(*fieldFlag)
		// } else {
		// 	splittedFieldFlag = []string{*fieldFlag}
		// }

		// splitted := strings.Split(string(b), *delimiterFlag)

		// if getMaxIdx(splittedFieldFlag) > len(splitted) {
		// 	continue
		// }

		// intSlice := stringToIntSlice(splittedFieldFlag)

		// for idx, fFlag := range intSlice {
		// 	if idx == len(intSlice) - 1 {
		// 		fmt.Printf("%s", splitted[fFlag - 1])
		// 	} else {
		// 		fmt.Printf("%s%s", splitted[fFlag - 1], *delimiterFlag)
		// 	}
		// }
		// fmt.Printf("\n")
	}
}

func getMaxIdx(splitted []string) int {
	var max int = -999999
	for _, sp := range splitted {
		i, err := strconv.Atoi(sp)
		if err != nil {
			log.Fatal(err)
		}

		if i > max {
			max = i
		}
	}

	return max
}

func stringToIntSlice(splitted []string) []int {
	var ret []int
	for _, sp := range splitted {
		i, err := strconv.Atoi(sp)
		if err != nil {
			log.Fatal(err)
		}

		ret = append(ret, i)
	}

	return ret
}