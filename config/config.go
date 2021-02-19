package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Token		string		`json:"token"`
	Path		string		`json:"path"`
	Fork		bool		`json:"fork"`
	Private		bool		`json:"private"`
	Public		bool		`json:"public"`
	Archived	bool		`json:"archived"`
	Disabled	bool		`json:"disabled"`
	Excludes	[]string	`json:"excludes"`
	excludesMap	map[string]bool
}

func LoadConfig() Config {

	config := Config{
		Path:		"/var/git_backup",
		Fork:		true,
		Private:	true,
		Public:		true,
		Archived:	true,
		Disabled:	true,
		Excludes:	[]string{},
	}

	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Error open config file")
	}

	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Error read config file")
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatal("Error parse config file")
	}

	if len(config.Token) == 0 {
		log.Fatal("Please specify GitHub personal token")
	}

	config.excludesMap = make(map[string]bool, len(config.Excludes))
	for _, s := range config.Excludes {
		config.excludesMap[s] = true
	}

	return config
}

func (c *Config) IsExclude(repoName string) bool {
	return c.excludesMap[repoName]
}