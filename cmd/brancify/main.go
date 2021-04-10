package main

import (
	"github.com/alexflint/go-arg"
	"github.com/atotto/clipboard"
	"github.com/sonereker/branchify/internal/branch"
	"github.com/sonereker/branchify/internal/jira"
	"log"
)

var args struct {
	BranchPrefix string `arg:"-p,--prefix" help:"Branch type prefix default is /bugfix" default:"bugfix/"`
	IssueKey     string `arg:"-k,--key required" help:"Jira Issue Key"`
}

func main() {
	arg.MustParse(&args)

	j := jira.New()
	summary := j.GetSummary(args.IssueKey)

	n := branch.NewName(args.BranchPrefix, args.IssueKey, summary)
	branchName := n.Generate()
	log.Println(branchName)

	// Copy generated branch name to clipboard
	err := clipboard.WriteAll(branchName)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Copied to clipboard!")
}
