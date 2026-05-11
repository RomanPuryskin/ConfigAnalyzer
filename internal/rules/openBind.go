package rules

import (
	"strings"

	"github.com/configAnalyzer/internal/entities"
	"github.com/configAnalyzer/pkg/utils"
)

var bindVariations = []string{"host", "bind", "address", "listen", "addr"}

type openBindRule struct {
}

func (o *openBindRule) Check(cfg map[string]any) []*entities.Issue {
	var issues []*entities.Issue

	recursiveTraversal(cfg, "", func(path string, val any) {
		key := strings.ToLower(path)
		if !utils.IsContains(key, bindVariations...) {
			return
		}
		if strings.Contains(utils.ToLowerString(val), "0.0.0.0") {
			issues = append(issues, &entities.Issue{
				ProblemLevel:   entities.ProblemLevelMedium,
				Title:          "Слушает на всех интерфейсах (0.0.0.0)",
				Description:    "Сервис доступен на всех сетевых интерфейсах без ограничений.",
				Recommendation: "Укажите конкретный адрес (например, 127.0.0.1) или защитите доступ брандмауэром.",
				Path:           path,
			})
		}
	})
	return issues
}
