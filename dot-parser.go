package dotparser

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type DotFileConfiguration struct {
	Project        string
	ConfigFileName string
	GetHomeDirFunc func() (string, error)
}

func ConstructPath(dotFileConfiguration DotFileConfiguration) string {
	GetHomeDirFunc := dotFileConfiguration.GetHomeDirFunc
	if GetHomeDirFunc == nil {
		GetHomeDirFunc = os.UserHomeDir
	}
	dir, err := GetHomeDirFunc()
	if err != nil {
		panic(err)
	}

	return filepath.Join(dir, ".config", dotFileConfiguration.Project, dotFileConfiguration.ConfigFileName)
}

type YamlParser[T any] struct{}

func (YamlParser[T]) Parse(content string) (T, error) {
	var cfg T
	err := yaml.Unmarshal([]byte(content), &cfg)

	return cfg, err
}

type DotFileParser[T any] interface {
	Parse(content string) (T, error)
}

//goland:noinspection GoUnusedExportedFunction
func ReadConfig[T any](parser DotFileParser[T], configFile DotFileConfiguration) (T, error) {
	var zero T
	data, err := os.ReadFile(ConstructPath(configFile))

	if err != nil {
		log.Fatal(err)
		return zero, err
	}

	result, err := parser.Parse(string(data))

	if err != nil {
		return zero, err
	}
	return result, nil
}
