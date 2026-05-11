package parser

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

type yamlParser struct {
}

func NewYamlParser() *yamlParser {
	return &yamlParser{}
}

func (y *yamlParser) Parse(r io.Reader) (map[string]any, error) {
	var out map[string]any
	if err := yaml.NewDecoder(r).Decode(&out); err != nil {
		return nil, fmt.Errorf("yaml parse: %w", err)
	}
	return out, nil
}
