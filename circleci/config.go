package circleci

import (
	"github.com/jszwedko/go-circleci"
)

type Config struct {
	Token        string
	Organization string
}

type Organization struct {
	name   string
	client *circleci.Client
}

func (c *Config) Client() (interface{}, error) {
	var org Organization

	org.name = c.Organization
	org.client = &circleci.Client{Token: c.Token}

	return &org, nil
}
