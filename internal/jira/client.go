package jira

import (
	"github.com/andygrunwald/go-jira"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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
		log.Fatalf("Jira: %v", err)
	}

	return issue.Fields.Summary
}

func readConf(filename string) *conf {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	yamlFile, err := ioutil.ReadFile(usr.HomeDir + "/." + filename)
	if err != nil {
		log.Fatalf("Config file is missing: %v", err)
	}

	c := &conf{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Config file error: %v", err)
	}

	return c
}
