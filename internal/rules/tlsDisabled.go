package rules

import (
	"strings"

	"github.com/configAnalyzer/internal/entities"
	"github.com/configAnalyzer/pkg/utils"
)

var tlsVariations = []string{"tls", "ssl"}
var URLVariations = []string{"url", "endpoint", "dsn", "connection"}

type tlsDisabledRule struct {
}

func (t *tlsDisabledRule) Check(cfg map[string]any) []*entities.Issue {
	var issues []*entities.Issue

	recursiveTraversal(cfg, "", func(path string, val any) {

		key := strings.ToLower(path)
		s := utils.ToLowerString(val)

		// tls.enabled = false  /  tls.disabled = true
		if utils.IsContains(key, tlsVariations...) {
			if utils.IsContains(key, "enabled", "enable") && s == "false" {
				issues = append(issues, &entities.Issue{
					ProblemLevel:   entities.ProblemLevelHigh,
					Title:          "TLS/SSL отключён",
					Description:    "Передача данных происходит без шифрования.",
					Recommendation: "Включите TLS и используйте действительные сертификаты.",
					Path:           path,
				})
			}
			if utils.IsContains(key, "disabled", "disable") && s == "true" {
				issues = append(issues, &entities.Issue{
					ProblemLevel:   entities.ProblemLevelHigh,
					Title:          "TLS/SSL отключён",
					Description:    "Передача данных происходит без шифрования.",
					Recommendation: "Включите TLS и используйте действительные сертификаты.",
					Path:           path,
				})
			}
			if utils.IsContains(key, "verify", "insecure_skip") {
				if (utils.IsContains(key, "verify") && s == "false") ||
					(utils.IsContains(key, "insecure_skip", "skip_verify") && s == "true") {
					issues = append(issues, &entities.Issue{
						ProblemLevel:   entities.ProblemLevelHigh,
						Title:          "Проверка TLS-сертификата отключена",
						Description:    "Приложение не проверяет TLS-сертификат удалённой стороны.",
						Recommendation: "Включите проверку сертификата; при необходимости добавьте CA в доверенные.",
						Path:           path,
					})
				}
			}
		}

		// http:// schema in url keys
		if utils.IsContains(key, URLVariations...) {
			if strings.HasPrefix(s, "http://") {
				issues = append(issues, &entities.Issue{
					ProblemLevel:   entities.ProblemLevelHigh,
					Title:          "Небезопасная схема URL (http://)",
					Description:    "Соединение без шифрования: «" + s + "».",
					Recommendation: "Используйте https:// для всех внешних соединений.",
					Path:           path,
				})
			}
		}
	})
	return issues
}
