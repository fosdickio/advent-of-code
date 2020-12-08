/*
Package utils/input implements a simple library for handling reading input from external sources.
*/
package utils

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// ReadIntInputFileIntoSlice reads input from an external file into a slice.  The input files contains a newline
// separated list of integer data.
func ReadIntInputFileIntoSlice(filePath string) (inputData []int, err error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	split := strings.Split(strings.TrimSpace(string(content)), "\n")

	data := make([]int, 0, len(split))
	for _, c := range split {
		i, err := strconv.Atoi(c)
		if err != nil {
			return nil, err
		}
		data = append(data, i)
	}

	return data, nil
}

// ReadInputFileIntoGrid reads input from an external file into a string slice.  The input files contains a newline
// separated list of string data.
func ReadInputFileIntoGrid(filePath string) (inputGrid []string, err error) {
	input, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return strings.Fields(string(input)), nil
}
