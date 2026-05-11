package rules

import (
	"strings"

	"github.com/configAnalyzer/internal/entities"
	"github.com/configAnalyzer/pkg/utils"
)

var permissionVariations = []string{"perm", "permission", "role", "access", "scope"}
var modeVariations = []string{"*", "all", "admin", "root", "superuser", "777", "0777"}

type widePermissionsRule struct {
}

func (w *widePermissionsRule) Check(cfg map[string]any) []*entities.Issue {
	var issues []*entities.Issue

	recursiveTraversal(cfg, "", func(path string, val any) {
		key := strings.ToLower(path)
		s := utils.ToLowerString(val)
		if !utils.IsContains(key, permissionVariations...) {
			return
		}

		for _, m := range modeVariations {
			if s == m || strings.Contains(s, m) {
				issues = append(issues, &entities.Issue{
					ProblemLevel:   entities.ProblemLevelHigh,
					Title:          "Слишком широкие права доступа",
					Description:    "Значение «" + s + "» предоставляет неограниченный доступ.",
					Recommendation: "Применяйте принцип минимальных привилегий: укажите только необходимые права.",
					Path:           path,
				})
				break
			}
		}
	})
	return issues
}
