package rules

import (
	"strings"

	"github.com/configAnalyzer/internal/entities"
	"github.com/configAnalyzer/pkg/utils"
)

var passwordVariations = []string{"password", "passwd", "secret", "pwd", "pass", "token", "api_key", "apikey"}

type plaintextPasswordRule struct {
}

func (p *plaintextPasswordRule) Check(cfg map[string]any) []*entities.Issue {
	var issues []*entities.Issue

	recursiveTraversal(cfg, "", func(path string, val any) {
		key := strings.ToLower(path)
		if !utils.IsContains(key, passwordVariations...) {
			return
		}
		s := utils.ToLowerString(val)
		if s == "" || s == "null" || strings.HasPrefix(s, "${") || strings.HasPrefix(s, "env:") {
			return
		}
		issues = append(issues, &entities.Issue{
			ProblemLevel:   entities.ProblemLevelHigh,
			Title:          "Пароль/секрет в открытом виде",
			Description:    "Ключ «" + path + "» содержит значение, похожее на секрет, прямо в конфиг-файле.",
			Recommendation: "Используйте переменные окружения, секрет-менеджер (Vault, AWS Secrets Manager и т.п.).",
			Path:           path,
		})

	})
	return issues
}
