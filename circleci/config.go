package circleci

import (
	"github.com/TomTucka/go-circleci"
)

type Config struct {
	Token        string
	Organization string
	vcs          string
}

type Organization struct {
	name   string
	VCS    string
	client *circleci.Client
}

func (c *Config) Client() (interface{}, error) {
	var org Organization

	org.name = c.Organization
	org.client = &circleci.Client{Token: c.Token}
	org.VCS = c.vcs

	return &org, nil
}
