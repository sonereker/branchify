package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/andygrunwald/go-jira"
	"github.com/gosimple/slug"
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
	"os/user"
)

var args struct {
	BranchPrefix string `arg:"-p,--prefix" help:"Branch type prefix default is /bugfix" default:"bugfix/"`
	IssueKey     string `arg:"-k,--key required" help:"Jira Issue Key"`
}

func main() {
	arg.MustParse(&args)

	baseUrl, username, password := getConfig()

	tp := jira.BasicAuthTransport{
		Username: username,
		Password: password,
	}

	j, _ := jira.NewClient(tp.Client(), baseUrl)
	issue, _, _ := j.Issue.Get(args.IssueKey, nil)

	summarySlug := slug.Make(issue.Fields.Summary)
	fmt.Printf("%s%s-%s\n", args.BranchPrefix, issue.Key, summarySlug)
}

func getConfig() (string, string, string) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	config, err := yaml.ReadFile(usr.HomeDir + "/.brancify")
	if err != nil {
		fmt.Println(err)
	}

	baseUrl, err := config.Get("base_url")
	if err != nil {
		fmt.Println(err)
	}

	username, err := config.Get("username")
	if err != nil {
		fmt.Println(err)
	}
	password, err := config.Get("password")
	if err != nil {
		fmt.Println(err)
	}
	return baseUrl, username, password
}
