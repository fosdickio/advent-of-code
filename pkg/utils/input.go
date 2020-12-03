package utils

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

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
