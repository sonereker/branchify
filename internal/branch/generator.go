package branch

import (
	"fmt"
	"github.com/gosimple/slug"
)

type name struct {
	Prefix   string
	IssueKey string
	Summary  string
}

//NewName returns a new `Name` instance
func NewName(prefix string, issueKey string, summary string) *name {
	return &name{Prefix: prefix, IssueKey: issueKey, Summary: summary}
}

//Generate returns branch name generated from prefix, issueKey and summarySlug
func (s *name) Generate() string {
	summarySlug := slug.Make(s.Summary)
	branchName := fmt.Sprintf("%s%s-%s", s.Prefix, s.IssueKey, summarySlug)
	return branchName
}
