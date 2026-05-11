package analyzer

import (
	"github.com/configAnalyzer/internal/entities"
	"github.com/configAnalyzer/internal/rules"
)

type Analyzer struct {
	Rules []rules.Rule
}

func NewAnalyzer() *Analyzer {
	rules := rules.BasicRules()
	return &Analyzer{
		Rules: rules,
	}
}

func (a *Analyzer) AddRule(r rules.Rule) {
	a.Rules = append(a.Rules, r)
}

func (a *Analyzer) Run(cfg map[string]any) []*entities.Issue {
	var issues []*entities.Issue

	// применяем каждое правило
	for _, rule := range a.Rules {
		issues = append(issues, rule.Check(cfg)...)
	}
	return issues
}
