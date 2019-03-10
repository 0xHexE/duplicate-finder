package main

import (
	"github.com/alexflint/go-arg"
	"strings"
)

func ParseArgument() ArgumentResult {
	var returnValue ArgumentResult
	returnValue.MatchPercentage = 100
	returnValue.Output = "output.csv"
	argResult := arg.MustParse(&returnValue)
	for _, key := range returnValue.MatchKeys {
		if len(strings.Split(key, ":")) != 2 {
			argResult.Fail("Invalid format for keys. Example INPUT_FILE_KEY:MATCH_FILE_KEY[,MATCH_FILE_OTHER_KEY]")
		}
	}
	returnValue.MatchPercentage = returnValue.MatchPercentage / 100

	returnValue.SplitKeys = map[string][]string{}

	returnValue.SplitKeys["input"] = []string{}
	returnValue.SplitKeys["output"] = []string{}

	for _, key := range returnValue.MatchKeys {
		keys := strings.Split(key, ":")
		returnValue.SplitKeys["input"] = append(returnValue.SplitKeys["input"], keys[0])

		for _, otherKeys := range strings.Split(keys[1], ",") {
			returnValue.SplitKeys["output"] = append(returnValue.SplitKeys["output"], otherKeys)
		}
	}

	return returnValue
}

type ArgumentResult struct {
	InputFile       string   `arg:"required"`
	MatchFile       string   `arg:"required"`
	MatchKeys       []string `arg:"required"`
	Output          string
	MatchPercentage float32

	SplitKeys map[string][]string `arg:"-"`
}
