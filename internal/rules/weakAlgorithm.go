package rules

import (
	"strings"

	"github.com/configAnalyzer/internal/entities"
	"github.com/configAnalyzer/pkg/utils"
)

type weakAlgorithmRule struct {
}

var algorithmVariations = []string{"algorithm", "algo", "cipher", "digest", "hash", "encryption", "signing"}
var weakAlgorithmVariations = []string{"md5", "sha1", "sha-1", "des", "3des", "rc4", "rc2", "blowfish"}

func (w *weakAlgorithmRule) Check(cfg map[string]any) []*entities.Issue {
	var issues []*entities.Issue

	recursiveTraversal(cfg, "", func(path string, val any) {
		key := strings.ToLower(path)
		if !utils.IsContains(key, algorithmVariations...) {
			return
		}
		s := utils.ToLowerString(val)
		for _, weak := range weakAlgorithmVariations {
			if strings.Contains(s, weak) {
				issues = append(issues, &entities.Issue{
					ProblemLevel:   entities.ProblemLevelHigh,
					Title:          "Небезопасный алгоритм: " + strings.ToUpper(weak),
					Description:    "Алгоритм «" + s + "» считается устаревшим и небезопасным.",
					Recommendation: "Замените на современный алгоритм: SHA-256/SHA-3, AES-GCM, ChaCha20-Poly1305, bcrypt/argon2 для паролей.",
					Path:           path,
				})
				break
			}
		}
	})
	return issues
}
