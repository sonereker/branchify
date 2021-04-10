package branch

import (
	"fmt"
	"github.com/gosimple/slug"
)

type Name struct {
	Prefix   string
	IssueKey string
	Summary  string
}

func NewName(prefix string, issueKey string, summary string) *Name {
	return &Name{Prefix: prefix, IssueKey: issueKey, Summary: summary}
}

//Generate returns branch name generated from prefix, issueKey and summarySlug
func (s *Name) Generate() string {
	summarySlug := slug.Make(s.Summary)
	branchName := fmt.Sprintf("%s%s-%s", s.Prefix, s.IssueKey, summarySlug)
	return branchName
}
