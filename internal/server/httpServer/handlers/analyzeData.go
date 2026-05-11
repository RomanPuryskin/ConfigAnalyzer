package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/configAnalyzer/internal/analyzer"
	"github.com/configAnalyzer/internal/parser"
	"github.com/configAnalyzer/internal/server/httpServer"
)

type AnalyzeDataHanlder struct {
	parser   *parser.DataParser
	analyzer *analyzer.Analyzer
}

func NewAnalyzerHandler(p *parser.DataParser, a *analyzer.Analyzer) *AnalyzeDataHanlder {
	return &AnalyzeDataHanlder{
		parser:   p,
		analyzer: a,
	}
}

func (a *AnalyzeDataHanlder) Analyze() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req httpServer.AnalyzeFileRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
			return
		}
		if req.Content == "" {
			http.Error(w, "content is required", http.StatusBadRequest)
			return
		}

		// парсим
		cfg, err := a.parser.Run([]byte(req.Content))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// запускаем анализатор
		issues := a.analyzer.Run(cfg)

		resp := httpServer.AnalyzeFileResponse{Issues: issues, Count: len(issues)}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
