package parser

import (
	"encoding/json"
	"fmt"
	"io"
)

type jsonParser struct {
}

func NewJsonParser() *jsonParser {
	return &jsonParser{}
}

func (j *jsonParser) Parse(r io.Reader) (map[string]any, error) {
	var out map[string]any
	if err := json.NewDecoder(r).Decode(&out); err != nil {
		return nil, fmt.Errorf("json parse: %w", err)
	}
	return out, nil
}
