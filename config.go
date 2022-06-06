package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

// ConsulConfigurator  configuration settings
func ConsulConfigurator(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		// Setting for consul address
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		// setting for consul. DEFAULT: /micro/config/consul
		consul.WithPrefix(prefix),
		// prefix remover.
		consul.StripPrefix(true),
	)
	config, err := config.NewConfig()
	if err != nil {
		return config, err
	}
	// configuration loader
	err = config.Load(consulSource)
	return config, err
}
