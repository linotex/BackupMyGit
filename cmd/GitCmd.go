package cmd

import (
	"log"
	"os/exec"
)

type GitCmd struct {

}

func (g *GitCmd) Clone(url string, destination string) error {

	err := g.runCommand(destination, "clone", "--mirror", url, ".git")
	if err != nil {
		log.Println("Cannot clone repo ", url)
		return err
	}

	err = g.runCommand(destination, "config", "--unset", "core.bare")
	if err != nil {
		return err
	}

	err = g.runCommand(destination, "reset", "--hard")
	if err != nil {
		return err
	}

	return nil
}

func (g *GitCmd) Update(destination string) error {

	err := g.runCommand(destination, "pull")
	if err != nil {
		log.Println("Cannot update repo")
		return err
	}

	return nil
}

func (g *GitCmd) runCommand(dir string, args... string) error {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	return cmd.Run()
}
