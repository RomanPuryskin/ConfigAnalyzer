package rules

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/configAnalyzer/internal/entities"
	"github.com/stretchr/testify/require"
)

func TestDebugModeRule(t *testing.T) {
	tests := []struct {
		name         string
		cfg          map[string]any
		expected     int
		problemLevel entities.ProblemLevel
	}{
		{
			name: "logLevel_debug",
			cfg: map[string]any{
				"log_level": "debug",
			},
			expected:     1,
			problemLevel: entities.ProblemLevelLow,
		},
		{
			name: "logLevel_not_Debug",
			cfg: map[string]any{
				"log_level": "info",
			},
			expected: 0,
		},
		{
			name: "logLevel_debug_and_nested_and_2",
			cfg: map[string]any{
				"logging": map[string]any{
					"level": "debug",
				},
				"log_level": "debug",
			},
			expected:     2,
			problemLevel: entities.ProblemLevelLow,
		},
	}

	rule := &debugModeRule{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			issues := rule.Check(tt.cfg)

			require.Len(t, issues, tt.expected)

			if tt.expected > 0 {
				require.Equal(t, tt.problemLevel, issues[0].ProblemLevel)
			}
		})
	}
}

func TestOpenBindRule(t *testing.T) {
	tests := []struct {
		name         string
		cfg          map[string]any
		expected     int
		problemLevel entities.ProblemLevel
	}{
		{
			name: "open_bind",
			cfg: map[string]any{
				"host": "0.0.0.0",
			},
			expected:     1,
			problemLevel: entities.ProblemLevelMedium,
		},
		{
			name: "localhost",
			cfg: map[string]any{
				"host": "127.0.0.1",
			},
			expected:     0,
			problemLevel: entities.ProblemLevelMedium,
		},
		{
			name: "open_nested_bind_and_2",
			cfg: map[string]any{
				"server": map[string]any{
					"listen": "0.0.0.0:8080",
				},
				"host": "0.0.0.0",
			},
			expected:     2,
			problemLevel: entities.ProblemLevelMedium,
		},
	}

	rule := &openBindRule{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			issues := rule.Check(tt.cfg)

			require.Len(t, issues, tt.expected)

			if tt.expected > 0 {
				require.Equal(t, tt.problemLevel, issues[0].ProblemLevel)
			}
		})
	}
}

func TestPlaintextPasswordRule(t *testing.T) {
	tests := []struct {
		name         string
		cfg          map[string]any
		expected     int
		problemLevel entities.ProblemLevel
	}{
		{
			name: "plaintext password",
			cfg: map[string]any{
				"password": "secret123",
			},
			expected:     1,
			problemLevel: entities.ProblemLevelHigh,
		},
		{
			name: "env variable",
			cfg: map[string]any{
				"password": "${DB_PASSWORD}",
			},
			expected:     0,
			problemLevel: entities.ProblemLevelHigh,
		},
		{
			name: "empty password",
			cfg: map[string]any{
				"password": "",
			},
			expected:     0,
			problemLevel: entities.ProblemLevelHigh,
		},
		{
			name: "nested_api_key_and_2",
			cfg: map[string]any{
				"auth": map[string]any{
					"api_key": "abcd",
				},
				"password": "secret123",
			},
			expected:     2,
			problemLevel: entities.ProblemLevelHigh,
		},
	}

	rule := &plaintextPasswordRule{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			issues := rule.Check(tt.cfg)

			require.Len(t, issues, tt.expected)

			if tt.expected > 0 {
				require.Equal(t, tt.problemLevel, issues[0].ProblemLevel)
			}
		})
	}
}

func TestTLSDisabledRule(t *testing.T) {
	tests := []struct {
		name         string
		cfg          map[string]any
		expected     int
		problemLevel entities.ProblemLevel
	}{
		{
			name: "tls_disabled",
			cfg: map[string]any{
				"tls_enabled": false,
			},
			expected:     1,
			problemLevel: entities.ProblemLevelHigh,
		},
		{
			name: "tls_enabled",
			cfg: map[string]any{
				"tls_enabled": true,
			},
			expected:     0,
			problemLevel: entities.ProblemLevelHigh,
		},
		{
			name: "http_url",
			cfg: map[string]any{
				"url": "http://example.com",
			},
			expected:     1,
			problemLevel: entities.ProblemLevelHigh,
		},
		{
			name: "https_url",
			cfg: map[string]any{
				"url": "https://example.com",
			},
			expected:     0,
			problemLevel: entities.ProblemLevelHigh,
		},
		{
			name: "2_issues",
			cfg: map[string]any{
				"tls_enabled": false,
				"url":         "http://example.com",
			},
			expected:     2,
			problemLevel: entities.ProblemLevelHigh,
		},
	}

	rule := &tlsDisabledRule{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			issues := rule.Check(tt.cfg)

			require.Len(t, issues, tt.expected)

			if tt.expected > 0 {
				require.Equal(t, tt.problemLevel, issues[0].ProblemLevel)
			}
		})
	}
}

func TestWeakAlgorithmRule(t *testing.T) {
	tests := []struct {
		name         string
		cfg          map[string]any
		expected     int
		problemLevel entities.ProblemLevel
	}{
		{
			name: "md5",
			cfg: map[string]any{
				"algorithm": "md5",
			},
			expected:     1,
			problemLevel: entities.ProblemLevelHigh,
		},
		{
			name: "sha1",
			cfg: map[string]any{
				"hash": "sha1",
			},
			expected:     1,
			problemLevel: entities.ProblemLevelHigh,
		},
		{
			name: "nested_and_2",
			cfg: map[string]any{
				"crypto": map[string]any{
					"signing": "sha-1",
				},
				"hash": "sha1",
			},
			expected:     2,
			problemLevel: entities.ProblemLevelHigh,
		},
	}

	rule := &weakAlgorithmRule{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			issues := rule.Check(tt.cfg)

			require.Len(t, issues, tt.expected)

			if tt.expected > 0 {
				require.Equal(t, tt.problemLevel, issues[0].ProblemLevel)
			}
		})
	}
}

func TestWidePermissionsRule(t *testing.T) {
	tests := []struct {
		name         string
		cfg          map[string]any
		expected     int
		problemLevel entities.ProblemLevel
	}{
		{
			name: "admin role",
			cfg: map[string]any{
				"role": "admin",
			},
			expected:     1,
			problemLevel: entities.ProblemLevelHigh,
		},
		{
			name: "777 permissions",
			cfg: map[string]any{
				"permission": "0777",
			},
			expected:     1,
			problemLevel: entities.ProblemLevelHigh,
		},
		{
			name: "nested_and_2",
			cfg: map[string]any{
				"crypto": map[string]any{
					"permission": "0777",
				},
				"role": "admin",
			},
			expected:     2,
			problemLevel: entities.ProblemLevelHigh,
		},
	}

	rule := &widePermissionsRule{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			issues := rule.Check(tt.cfg)

			require.Len(t, issues, tt.expected)

			if tt.expected > 0 {
				require.Equal(t, tt.problemLevel, issues[0].ProblemLevel)
			}
		})
	}
}

func TestFilePermissionRule(t *testing.T) {
	tests := []struct {
		name     string
		perm     os.FileMode
		expected int
	}{
		{
			name:     "secure permissions",
			perm:     0o600,
			expected: 2,
		},
		{
			name:     "world readable",
			perm:     0o644,
			expected: 2,
		},
		{
			name:     "world writable",
			perm:     0o666,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dir := t.TempDir()

			file := filepath.Join(dir, "config.yaml")

			err := os.WriteFile(file, []byte("test"), tt.perm)
			require.NoError(t, err)

			err = os.Chmod(file, tt.perm)
			require.NoError(t, err)

			rule := &FilePermissionRule{
				Path: file,
			}

			issues := rule.Check(nil)

			require.Len(t, issues, tt.expected)
		})
	}
}
