package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Configuration struct {
	Server     Host     `yaml:server`
	Repository Repo     `yaml:repository`
	Shared     Shared   `yaml:shared`
	Tasks      []string `yaml:tasks`
	Cluster    Cluster  `yaml:cluster`
}

type Host struct {
	Address string `yaml:"address"`
	User    string `yaml:"user"`
	Port    int    `yaml:"port"`
	Dir     string `yaml:"dir"`
	Project string `yaml:"project"`
}

type Repo struct {
	Url    string `yaml:"url"`
	Branch string `yaml:"branch"`
	Tag    string `yaml:"tag"`
}

type Shared struct {
	Folders []string `yaml:"folders"`
	Files   []string `yaml:"files"`
}

type Cluster struct {
	Hosts []string `yaml:"hosts"`
	Rsync Rsync    `yaml:"rsync"`
	Cmds  []string `yaml:"cmds"`
}

type Rsync struct {
	Excludes []string `yaml:"excludes"`
}

func (c *Configuration) ReadFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	errY := yaml.Unmarshal(file, &c)
	if errY != nil {
		return errY
	}

	return nil
}
