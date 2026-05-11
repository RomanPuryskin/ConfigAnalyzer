package rules

import (
	"fmt"
	"os"

	"github.com/configAnalyzer/internal/entities"
)

type FilePermissionRule struct {
	Path string
}

func (r *FilePermissionRule) Check(_ map[string]any) []*entities.Issue {
	info, err := os.Stat(r.Path)
	if err != nil {
		return nil
	}
	mode := info.Mode().Perm()
	var issues []*entities.Issue

	if mode&0o004 != 0 {
		issues = append(issues, &entities.Issue{
			ProblemLevel:   entities.ProblemLevelMedium,
			Title:          "Конфиг-файл доступен для чтения всем",
			Description:    fmt.Sprintf("Права файла %s: %s — другие пользователи могут прочитать файл.", r.Path, mode),
			Recommendation: "Установите права 600 или 640: chmod 600 " + r.Path,
			Path:           r.Path,
		})
	}
	if mode&0o002 != 0 {
		issues = append(issues, &entities.Issue{
			ProblemLevel:   entities.ProblemLevelHigh,
			Title:          "Конфиг-файл доступен для записи всем",
			Description:    fmt.Sprintf("Права файла %s: %s — посторонние могут изменить конфигурацию.", r.Path, mode),
			Recommendation: "Установите права 600: chmod 600 " + r.Path,
			Path:           r.Path,
		})
	}
	return issues
}
