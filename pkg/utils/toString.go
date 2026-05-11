package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ToLowerString(v any) string {
	if v == nil {
		return ""
	}
	switch t := v.(type) {
	case string:
		return strings.ToLower(t)
	case bool:
		if t {
			return "true"
		}
		return "false"
	case int:
		return fmt.Sprintf("%d", t)
	case int32:
		return fmt.Sprintf("%d", t)
	case int64:
		return fmt.Sprintf("%d", t)
	case float32:
		return fmt.Sprintf("%g", t)
	case float64:
		return fmt.Sprintf("%g", t)
	case json.Number:
		return strings.ToLower(t.String())
	default:
		return strings.ToLower(fmt.Sprintf("%v", t))
	}
}
