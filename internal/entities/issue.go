package entities

import "fmt"

type ProblemLevel string

const (
	ProblemLevelLow    ProblemLevel = "LOW"
	ProblemLevelMedium ProblemLevel = "MEDIUM"
	ProblemLevelHigh   ProblemLevel = "HIGH"
)

type Issue struct {
	ProblemLevel   ProblemLevel `json:"problem_level"`
	Title          string       `json:"title"`
	Description    string       `json:"description"`
	Recommendation string       `json:"recommendation"`
	Path           string       `json:"path"`
}

func GetIssueInfo(i *Issue) string {
	return fmt.Sprintf("[%s] %s\n  Место: %s\n  Описание: %s\n  Рекомендация: %s",
		i.ProblemLevel, i.Title, i.Path, i.Description, i.Recommendation)
}
