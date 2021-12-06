package main

import (
	"BackupMyGit/api"
	"BackupMyGit/cmd"
	"BackupMyGit/config"
	"flag"
	"fmt"
	"log"
	"os"
)

const Version = "0.1"

func main() {

	version := flag.Bool("v", false, "Version of tool")
	flag.Parse()

	if *version {
		fmt.Println(Version)
		return
	}

	cfg := config.LoadConfig()

	repoList := getRepoList(cfg)

	_, _ = createPath(cfg.Path, false)

	gitCmd := cmd.GitCmd{}

	for _, repo := range repoList {
		fmt.Println()
		log.Println("Updating repo", repo.FullName)
		path := cfg.Path + "/" + repo.FullName

		log.Println("Repo path", path)
		isExists, err := createPath(path, true)
		if err != nil {
			continue
		}

		if !isExists {
			log.Println("Cloning repo...")
			err = gitCmd.Clone(repo.SshURL, path)

			if err != nil {
				log.Println("Cannot clone repo", repo.SshURL)
				log.Println(err)

				continue
			}
		} else {
			log.Println("Updating repo...")
			err = gitCmd.Update(path)

			if err != nil {
				log.Println("Error update repo")
				log.Println(err)

				continue
			}
		}

		log.Println("Success")
	}
}

func getRepoList(cfg config.Config) []api.Repo {
	client := api.NewClient(cfg.Token)
	repoList := client.GetRepoList()
	filterRepos := []api.Repo{}

	log.Println("Found", len(repoList), "repos")

	for _, repo := range repoList {
		if (!cfg.Fork && repo.Fork) ||
			(!cfg.Private && repo.Private) ||
			(!cfg.Public && !repo.Private) ||
			(!cfg.Archived && repo.Archived) ||
			(!cfg.Disabled && repo.Disabled) ||
			cfg.IsExclude(repo.FullName) {
			continue
		}

		filterRepos = append(filterRepos, repo)
	}

	log.Println("Processing", len(filterRepos), "repos")

	return filterRepos
}

func createPath(path string, skipError bool) (bool, error) {
	_, err := os.Stat(path)
	isExists := true
	if os.IsNotExist(err) {
		isExists = false
		log.Println("Path", path, "is not exists. Creating...")
		err = os.MkdirAll(path, os.ModePerm)

		if err != nil && !skipError {
			log.Fatal("Cannot create path", path)
			return isExists, err
		}

		return isExists, err
	}

	return isExists, nil
}
