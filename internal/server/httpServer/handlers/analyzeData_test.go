package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/configAnalyzer/internal/analyzer"
	"github.com/configAnalyzer/internal/parser"
	"github.com/configAnalyzer/internal/server/httpServer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnalyze_JSON_Success(t *testing.T) {

	dataParser := parser.NewDataParser()
	analyzer := analyzer.NewAnalyzer()

	h := NewAnalyzerHandler(dataParser, analyzer)

	reqBody := httpServer.AnalyzeFileRequest{
		Content: `{
			"storage": {
				"digest-algorithm": "MD5"
			},
			"log": {
				"output": "stdout",
				"level": "debug"
			}
			}`,
	}

	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/analyze", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	h.Analyze().ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var resp httpServer.AnalyzeFileResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	require.NoError(t, err)

	assert.Equal(t, 2, resp.Count)
}

func TestAnalyze_YAML_Success(t *testing.T) {

	dataParser := parser.NewDataParser()
	analyzer := analyzer.NewAnalyzer()

	h := NewAnalyzerHandler(dataParser, analyzer)

	reqBody := httpServer.AnalyzeFileRequest{
		Content: `version: 2.2
storage:
  digest-algorithm: MD5
log:
  output: stdout
  level: debug
  `,
	}

	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/analyze", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	h.Analyze().ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var resp httpServer.AnalyzeFileResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	require.NoError(t, err)

	assert.Equal(t, 2, resp.Count)
}

func TestAnalyze_UnssuportedFormat(t *testing.T) {

	dataParser := parser.NewDataParser()
	analyzer := analyzer.NewAnalyzer()

	h := NewAnalyzerHandler(dataParser, analyzer)

	reqBody := httpServer.AnalyzeFileRequest{
		Content: "some text",
	}

	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/analyze", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	h.Analyze().ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), parser.ErrUnsupportedDataFormat.Error())
}
