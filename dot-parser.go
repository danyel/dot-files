package dotparser

import (
	"log"
	"os"
)

type DotFileParser[T any] interface {
	Parse(content string) (T, error)
}

func ReadConfig[T any](parser DotFileParser[T], configFile string) (T, error) {
	var zero T
	data, err := os.ReadFile(configFile)

	if err != nil {
		log.Fatal(err)
		return zero, err
	}

	result, err := parser.Parse(string(data))

	if err != nil {
		return zero, err
	}
	log.Printf("data is parsed as %v", result)
	return result, nil
}
