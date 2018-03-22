package plugin

import (
	"github.com/arquillian/ike-prow-plugins/pkg/log"
	"github.com/arquillian/ike-prow-plugins/pkg/plugin/config"
	"github.com/arquillian/ike-prow-plugins/pkg/scm"
	"github.com/arquillian/ike-prow-plugins/pkg/github"
)

// TestKeeperConfiguration defines inclusion and exclusion patterns set of files will be matched against
// It's unmarshaled from test-keeper.yml configuration file
type TestKeeperConfiguration struct {
	config.PluginConfiguration `yaml:",inline"`
	Inclusion string           `yaml:"test_pattern,omitempty"`
	Exclusion string           `yaml:"skip_validation_for,omitempty"`
	Combine   bool             `yaml:"combine_defaults,omitempty"`
}

// LoadTestKeeperConfig loads a TestKeeperConfiguration for the given change
func LoadTestKeeperConfig(log log.Logger, change scm.RepositoryChange) TestKeeperConfiguration {

	configuration := TestKeeperConfiguration{Combine: true}
	loadableConfig := &github.LoadableConfig{PluginName: ProwPluginName, Change: change, BaseConfig: &configuration.PluginConfiguration}

	err := config.Load(&configuration, loadableConfig)

	if err != nil {
		log.Warnf("Config file was not loaded. Cause: %s", err)
		return configuration
	}

	return configuration
}
