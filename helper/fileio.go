package helper

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(file string) string {
	reader, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
