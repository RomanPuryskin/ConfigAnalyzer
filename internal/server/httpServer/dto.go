package httpServer

import "github.com/configAnalyzer/internal/entities"

type AnalyzeFileRequest struct {
	Content string `json:"content"`
}

type AnalyzeFileResponse struct {
	Issues []*entities.Issue `json:"issues"`
	Count  int               `json:"count"`
}
