package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
)

type conf struct {
	BaseUrl  string `yaml:"base_url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Client struct {
	Client *jira.Client
}

func New() *Client {
	c := readConf("brancify")

	tp := jira.BasicAuthTransport{
		Username: c.Username,
		Password: c.Password,
	}

	client, _ := jira.NewClient(tp.Client(), c.BaseUrl)
	return &Client{Client: client}
}

//GetSummary returns summary value of the issue
func (j *Client) GetSummary(issueKey string) string {
	issue, _, err := j.Client.Issue.Get(issueKey, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Jira Error: %s\n", err.Error())
		os.Exit(1)
	}

	return issue.Fields.Summary
}

func readConf(filename string) *conf {
	usr, _ := user.Current()
	yamlFile, err := ioutil.ReadFile(usr.HomeDir + "/." + filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Config file is missing: %s\n", err.Error())
		os.Exit(1)
	}

	c := &conf{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Config file error: %s\n", err.Error())
		os.Exit(1)
	}

	return c
}
