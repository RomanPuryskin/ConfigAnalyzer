package rules

import "github.com/configAnalyzer/internal/entities"

type Rule interface {
	Check(map[string]any) []*entities.Issue
}

func BasicRules() []Rule {
	return []Rule{
		&debugModeRule{},
		&openBindRule{},
		&plaintextPasswordRule{},
		&tlsDisabledRule{},
		&weakAlgorithmRule{},
		&widePermissionsRule{},
	}
}

func recursiveTraversal(cfg map[string]any, prefix string, fn func(path string, val any)) {
	for k, v := range cfg {
		fullPath := k
		if prefix != "" {
			fullPath = prefix + "." + k
		}
		if nested, ok := v.(map[string]any); ok {
			recursiveTraversal(nested, fullPath, fn)
		} else {
			fn(fullPath, v)
		}
	}
}
