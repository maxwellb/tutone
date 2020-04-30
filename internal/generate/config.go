package generate

import (
	"errors"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/newrelic/tutone/internal/schema"
)

// Config is the package specific configuration file
type Config struct {
	Package       string                    `yaml:"package,omitempty"`
	Types         []schema.TypeInfo         `yaml:"types,omitempty"`
	Mutations     []schema.MutationInfo     `yaml:"mutations,omitempty"`
	Subscriptions []schema.SubscriptionInfo `yaml:"subscriptions,omitempty"`
	Queries       []schema.QueryInfo        `yaml:"queries,omitempty"`
}

func LoadConfig(file string) (*Config, error) {
	if file == "" {
		return nil, errors.New("config file name required")
	}
	log.WithFields(log.Fields{
		"file": file,
	}).Debug("loading package definition")

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	log.Tracef("definition: %+v", config)

	return &config, nil
}