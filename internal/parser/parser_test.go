package parser

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseFile(t *testing.T) {
	tests := []struct {
		name string
		path string
		err  error
	}{
		{
			name: "unsupported_extension",
			path: "../../test_data/test_txt.txt",
			err:  ErrUnsupportedFileExtension,
		},
		{
			name: "success_yaml_parse",
			path: "../../test_data/test_yaml.yaml",
			err:  nil,
		},
		{
			name: "success_json_parse",
			path: "../../test_data/test_json.json",
			err:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fileParser := NewFileParser()
			_, err := fileParser.Run(tt.path)

			require.Equal(t, errors.Is(err, tt.err), true)
		})
	}
}

func TestParseStdin(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		err  error
	}{
		{
			name: "unsupported_format",
			data: []byte("unssuproted format"),
			err:  ErrUnsupportedDataFormat,
		},
		{
			name: "success_yaml_parse",
			data: []byte(`version: 2.2
storage:
    digest-algorithm: MD5
log:
    output: stdout
    level: debug`),
			err: nil,
		},
		{
			name: "success_json_parse",
			data: []byte(`
			{
				"storage": {
					"digest-algorithm": "MD5"
				},
				"log": {
					"output": "stdout",
					"level": "debug"
				}
			}
			`),
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dataParser := NewDataParser()
			_, err := dataParser.Run(tt.data)

			require.Equal(t, errors.Is(err, tt.err), true)
		})
	}
}
