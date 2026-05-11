package rules

import (
	"strings"

	"github.com/configAnalyzer/internal/entities"
	"github.com/configAnalyzer/pkg/utils"
)

var logLevelVariations = []string{"level", "log_level", "loglevel", "logging"}

type debugModeRule struct {
}

func (m *debugModeRule) Check(cfg map[string]any) []*entities.Issue {
	var issues []*entities.Issue

	recursiveTraversal(cfg, "", func(path string, val any) {
		key := strings.ToLower(path)
		if !utils.IsContains(key, logLevelVariations...) {
			return
		}
		if utils.ToLowerString(val) == "debug" {
			issues = append(issues, &entities.Issue{
				ProblemLevel:   entities.ProblemLevelLow,
				Title:          "Логирование в debug-режиме",
				Description:    "Уровень логирования «debug» может раскрывать чувствительные данные в логах.",
				Recommendation: "Смените уровень на info или выше в продакшн-окружении.",
				Path:           path,
			})
		}
	})
	return issues
}
