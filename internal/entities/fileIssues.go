package entities

import (
	"fmt"
	"strings"
)

type FileIssues struct {
	Path   string   `json:"path"`
	Issues []*Issue `json:"issues"`
	Err    error    `json:"error"`
}

func GetFileIssuesInfo(f *FileIssues) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[%s]: \n", f.Path))

	for _, i := range f.Issues {
		sb.WriteString(" " + GetIssueInfo(i) + "\n")
	}
	if f.Err != nil {
		sb.WriteString(fmt.Sprintf("error:  %s\n", f.Err.Error()))
	}
	return sb.String()
}
