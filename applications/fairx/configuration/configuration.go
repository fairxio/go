package configuration

import (
	"github.com/fairxio/go/log"
	"github.com/spf13/viper"
)

const (
	CFG_LISTEN_ADDRESS = "httpListenAddress"
	CFG_LISTEN_PORT    = "httpListenPort"
	CFG_JWT_KEY        = "jwtKey"
)

type NodeConfiguration struct {
	ListenAddress string
	ListenPort    int
	JWTKey        string
}

var nodeConfiguration *NodeConfiguration

func Create() *NodeConfiguration {

	if nodeConfiguration == nil {

		log.Info("Reading Configuration")
		nc := NodeConfiguration{}
		nc.SetDefaults()
		nc.ReadConfig()

		nodeConfiguration = &nc

	}

	return nodeConfiguration

}

func (nc *NodeConfiguration) SetDefaults() {

	viper.SetDefault(CFG_LISTEN_ADDRESS, "0.0.0.0")
	viper.SetDefault(CFG_LISTEN_PORT, int(8080))
	viper.SetDefault(CFG_JWT_KEY, "onetwothreefour")

}

func (nc *NodeConfiguration) ReadConfig() error {

	viper.SetConfigName("fairx")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/fairx")
	viper.AddConfigPath("/opt/fairx")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Info("Unable to read in configuration:  %v.\n\nUsing defaults!", err)
	}

	return nil
}

func (nc *NodeConfiguration) GetListenAddress() string {
	return viper.GetString(CFG_LISTEN_ADDRESS)
}

func (nc *NodeConfiguration) GetListenPort() int {
	return viper.GetInt(CFG_LISTEN_PORT)
}

func (nc *NodeConfiguration) GetJWTKey() string {
	return viper.GetString(CFG_JWT_KEY)
}
