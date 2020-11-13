package circleci

import (
	"github.com/mattermost/go-circleci"
)

type Config struct {
	Token        string
	Organization string
	vcs          string
}

type Organization struct {
	name   string
	VCS    circleci.VcsType
	client *circleci.Client
}


func (c *Config) Client() (interface{}, error) {
	var org Organization

	org.name = c.Organization
	org.client = &circleci.Client{Token: c.Token, Version: circleci.APIVersion2}
	org.VCS = circleci.VcsType(c.vcs)

	return &org, nil
}
