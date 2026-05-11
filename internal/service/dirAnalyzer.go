package service

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/configAnalyzer/internal/parser"

	"github.com/configAnalyzer/internal/analyzer"
	"github.com/configAnalyzer/internal/entities"
)

type DirAnalyzer struct {
	FileAnalyer analyzer.Analyzer
}

func NewDirAnalyzer(fa *analyzer.Analyzer) *DirAnalyzer {
	return &DirAnalyzer{
		FileAnalyer: *fa,
	}
}

func (da *DirAnalyzer) Run(dirPath string) ([]*entities.FileIssues, error) {

	res := []*entities.FileIssues{}

	_, err := os.ReadDir(dirPath)
	if err != nil {
		return res, fmt.Errorf("read dir: %w", err)
	}

	da.recursiveTraversal(dirPath, &res)

	return res, nil
}

func (da *DirAnalyzer) recursiveTraversal(dirParh string, res *[]*entities.FileIssues) {
	filepath.WalkDir(dirParh, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			*res = append(*res, &entities.FileIssues{
				Path:   path,
				Issues: []*entities.Issue{},
				Err:    err,
			})
			return err
		}

		if !d.IsDir() {

			fileParser := parser.NewFileParser()
			cfg, err := fileParser.Run(path)
			if err != nil {
				*res = append(*res, &entities.FileIssues{
					Path:   path,
					Issues: []*entities.Issue{},
					Err:    err,
				})

				return err
			}
			issues := da.FileAnalyer.Run(cfg)
			*res = append(*res, &entities.FileIssues{
				Path:   path,
				Issues: issues,
				Err:    err,
			})
		}
		return nil
	})
}
