package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestProcessingLine(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		processed string
		errMsg    string
	}{
		{
			name:      "not matched",
			input:     "asdf aasdf asdf",
			processed: "asdf aasdf asdf",
		},
		{
			name:      "empty string",
			input:     "",
			processed: "",
		},
		{
			name:      "wrong string same keywords",
			input:     "#source value:",
			processed: "#source value:",
		},
		{
			name: "wrong string same keywords, indent is kepy",
			input: "		#source value:",
			processed: "		#source value:",
		},
		{
			name:      "file not found",
			input:     "value: asdf #source testdata/not-found-file.json",
			processed: "value: asdf #source testdata/not-found-file.json",
			errMsg:    "can not open file",
		},
		{
			name:      "file not found",
			input:     "value: asdf #source testdata/openapi.inplaced.yaml",
			processed: "value: asdf #source testdata/openapi.inplaced.yaml",
			errMsg:    "can not parse json from file",
		},
		{
			name: "success pastests and replaces, indent is correct",
			input: `	value: asdf #source testdata/user-basic.json`,
			processed: `	value: {"id":42,"name":"Nick Fury"} #source testdata/user-basic.json`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			processed, err := processLine(tc.input)

			if tc.processed != processed {
				t.Errorf("%s != %s", tc.processed, processed)
			}
			if tc.errMsg != "" && !strings.Contains(err.Error(), tc.errMsg) {
				t.Errorf("%s not in %s", err.Error(), tc.errMsg)
			}
		})
	}
}

func TestProcessingLines(t *testing.T) {
	tests := []struct {
		name       string
		inputFile  string
		outputFile string
	}{
		{
			name:       "basic",
			inputFile:  "testdata/openapi.only-comment.yaml",
			outputFile: "testdata/openapi.inplaced.yaml",
		},
		{
			name:       "with error in file",
			inputFile:  "testdata/openapi.with-error.yaml",
			outputFile: "testdata/openapi.with-error.yaml",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			original, err := os.Open(tc.inputFile)
			if err != nil {
				t.Error(err)

			}
			expected, err := os.Open(tc.outputFile)
			if err != nil {
				t.Error(err)
			}
			expectedBytes, err := ioutil.ReadAll(expected)
			if err != nil {
				t.Error(err)
			}

			buff := bytes.NewBuffer(nil)
			processLines(original, buff)

			if !bytes.Equal(expectedBytes, buff.Bytes()) {
				t.Fail()
			}
		})
	}
}
