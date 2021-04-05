package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// processLine parses single line
func processLine(input string) (string, error) {
	valueIdx := strings.Index(input, "value: ")
	sourceIdx := strings.Index(input, "#source ")

	if valueIdx == -1 || sourceIdx == -1 || valueIdx >= sourceIdx {
		// not the string we are interested in, skipping
		return input, nil
	}

	filePath := input[sourceIdx+len("#source "):]
	file, err := os.Open(filePath)
	if err != nil {
		return input, fmt.Errorf("can not open file %s: %w", filePath, err)
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return input, fmt.Errorf("can not read file: %w", err)
	}

	var jsonContent interface{} // anything
	if err := json.Unmarshal(fileBytes, &jsonContent); err != nil {
		return input, fmt.Errorf("can not parse json from file %s: %w", filePath, err)
	}

	jsonFile, err := json.Marshal(jsonContent)
	if err != nil {
		return input, fmt.Errorf("can not marshal json for value %#v: %w", jsonContent, err)
	}

	output := input[:valueIdx] + "value: " + string(jsonFile) + " #source " + filePath

	return output, nil
}

// processLines processes all lines from reader to writer
func processLines(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		input := scanner.Text()
		processed, err := processLine(input)
		if err != nil {
			log.Println(err.Error())
			fmt.Fprintln(out, input)
		} else {
			fmt.Fprintln(out, processed)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func main() {
	processLines(os.Stdin, os.Stdout)
}
