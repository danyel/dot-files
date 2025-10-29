package dotparser

import (
	"log"
	"os"

	"github.com/danyel/dot-files/configuration"
)

type DotFileParser[T any] interface {
	Parse(content string) (T, error)
}

//goland:noinspection GoUnusedExportedFunction
func ReadConfig[T any](parser DotFileParser[T], configFile configuration.DotFileConfiguration) (T, error) {
	var zero T
	data, err := os.ReadFile(configuration.ConstructPath(configFile))

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
