package yaml_parser

import (
	"gopkg.in/yaml.v3"
)

type Parser[T any] struct{}

func (Parser[T]) Parse(content string) (T, error) {
	var cfg T
	err := yaml.Unmarshal([]byte(content), &cfg)

	return cfg, err
}
