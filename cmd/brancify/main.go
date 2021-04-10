package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/atotto/clipboard"
	"github.com/sonereker/branchify/internal/branch"
	"github.com/sonereker/branchify/internal/jira"
	"os"
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
	fmt.Println(branchName)

	// Copy generated branch name to clipboard
	err := clipboard.WriteAll(branchName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't copy to clipboard: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Copied to clipboard!")
}
