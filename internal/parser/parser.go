package parser

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrUnsupportedFileExtension = errors.New("unsupported file extension")
	ErrUnsupportedDataFormat    = errors.New("unsupported data format")
)

type Parser interface {
	Parse(r io.Reader) (map[string]any, error)
}

// парсер для файлов
type FileParser struct {
	supportedParsers map[string]Parser
}

func NewFileParser() *FileParser {
	return &FileParser{
		supportedParsers: map[string]Parser{

			// все поддерживаемые и реализованные парсеры
			".json": NewJsonParser(),
			".yaml": NewYamlParser(),
			".yml":  NewYamlParser(),
		},
	}
}

func (fp *FileParser) Run(path string) (map[string]any, error) {
	// пробуем открыть файл
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	// создаем нужный парсер
	ext := strings.ToLower(filepath.Ext(path))
	parser, ok := fp.supportedParsers[ext]
	if !ok {

		// соберем инфо и всех поддерживаемых форматах
		formats := []string{}
		for key := range fp.supportedParsers {
			formats = append(formats, key)
		}
		return nil, fmt.Errorf("%w supported: %s", ErrUnsupportedFileExtension, strings.Join(formats, " "))
	}

	return parser.Parse(file)
}

// парсер для текста ( из stdin )
type DataParser struct {
	supportedParsers map[string]Parser
}

func NewDataParser() *DataParser {
	return &DataParser{
		supportedParsers: map[string]Parser{

			// все поддерживаемые и реализованные парсеры
			"json": NewJsonParser(),
			"yaml": NewYamlParser(),
		},
	}
}

func (dp *DataParser) Run(data []byte) (map[string]any, error) {

	// попробуем распарсить во все поддерживаемые форматы
	for _, p := range dp.supportedParsers {

		if m, err := p.Parse(strings.NewReader(string(data))); err == nil {
			return m, nil
		}
	}

	// не удалось распарсить ни в один поддерживаемый формат

	// соберем инфо о всех поддерживаемых форматах
	formats := []string{}
	for key := range dp.supportedParsers {
		formats = append(formats, key)
	}

	return nil, fmt.Errorf("%w supported: %s", ErrUnsupportedDataFormat, strings.Join(formats, " "))
}
