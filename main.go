package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/andygrunwald/go-jira"
	"github.com/atotto/clipboard"
	"github.com/gosimple/slug"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os/user"
)

var args struct {
	BranchPrefix string `arg:"-p,--prefix" help:"Branch type prefix default is /bugfix" default:"bugfix/"`
	IssueKey     string `arg:"-k,--key required" help:"Jira Issue Key"`
}

type conf struct {
	BaseUrl  string `yaml:"base_url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func main() {
	arg.MustParse(&args)

	c := readConf("brancify")

	tp := jira.BasicAuthTransport{
		Username: c.Username,
		Password: c.Password,
	}

	j, _ := jira.NewClient(tp.Client(), c.BaseUrl)
	issue, _, err := j.Issue.Get(args.IssueKey, nil)
	if err != nil {
		log.Fatalf("Jira: %v", err)
	}

	summarySlug := slug.Make(issue.Fields.Summary)
	branchName := fmt.Sprintf("%s%s-%s", args.BranchPrefix, issue.Key, summarySlug)
	fmt.Println(branchName)
	copyToClipboard(branchName)
}

func copyToClipboard(summarySlug string) {
	_ = clipboard.WriteAll(summarySlug)
	fmt.Println("Copied to clipboard!")
}

func readConf(filename string) *conf {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	yamlFile, err := ioutil.ReadFile(usr.HomeDir + "/." + filename)
	if err != nil {
		log.Fatalf("Config File: %v", err)
	}

	c := &conf{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
