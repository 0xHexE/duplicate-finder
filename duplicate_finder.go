package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/httpsOmkar/string-similarity"
	"os"
	"strings"
)

func main() {
	args := ParseArgument()

	inputFile := ReadFile(args.InputFile)
	matchFile := ReadFile(args.MatchFile)

	inputTitles := ParseTitles(inputFile[0], args.SplitKeys["input"])
	outputTitle := ParseTitles(matchFile[0], args.SplitKeys["output"])

	var result []string

	for _, inputData := range inputFile[1:] {
		var matchedData []string

		for _, outputData := range matchFile[1:] {
			var matchedKeys = map[string]bool{}

			for keyIndex, index := range args.SplitKeys["input"] {
				if MatchRecord(inputData[inputTitles[index]], outputData[outputTitle[args.SplitKeys["output"][keyIndex]]], args.MatchPercentage) {
					matchedKeys[index] = true
				} else if matchedKeys[index] != true {
					matchedKeys[index] = false
				}
			}

			isMatched := true

			for _, v := range matchedKeys {
				if v == false {
					isMatched = false
					break
				}
			}

			if isMatched {
				matchedData = append(matchedData, strings.Join(outputData, ","))
			}
		}

		if len(matchedData) != 0 {
			result = append(result, strings.Join(inputData, ","), strings.Join(matchedData, "\n"))
		}
	}

	WriteFileFromString(strings.Join(result, "\n\n"), args.Output)
}

func WriteFileFromString(input, path string) {
	handler, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	_, err = handler.WriteString(input)

	if err != nil {
		panic(err)
	}
}

func MatchRecord(string1, string2 string, matchPercentage float32) bool {
	if matchPercentage == 1.0 {
		return string1 == string2
	}

	return string_similarity.CompareString(string1, string2) >= float64(matchPercentage)
}

func ParseTitles(inputHead []string, keys []string) map[string]int {
	var returnValue = make(map[string]int, 0)
	for _, key := range keys {
		for index, headTitle := range inputHead {
			if key == headTitle {
				returnValue[headTitle] = index
				break
			}
		}
		if _, ok := returnValue[key]; !ok {
			panic(fmt.Sprintf("%s Key not found", key))
		}
	}
	return returnValue
}

func ReadFile(filePath string) [][]string {
	fileHandler, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fileHandler.Close()
	fileReader := csv.NewReader(bufio.NewReader(fileHandler))
	fileData, err := fileReader.ReadAll()
	if err != nil {
		panic(err)
	}
	return fileData
}
