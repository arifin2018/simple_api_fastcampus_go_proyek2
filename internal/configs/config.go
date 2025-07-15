package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var config *Config

type option struct {
	configFolders []string
	configFile    string
	configType    string
}

func Init(opts ...Option) error {
	opt := &option{
		configFolders: getDefaultConfigFolder(),
		configFile:    getDefaultConfigFile(),
		configType:    getDefaultConfigType(),
	}

	for _, v := range opts {
		v(opt)
	}

	for _, v := range opt.configFolders {
		viper.AddConfigPath(v)
	}

	viper.SetConfigName(opt.configFile)
	viper.SetConfigType(opt.configType)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	fmt.Println(viper.AllSettings())
	fmt.Println(viper.AllKeys())
	fmt.Println(viper.GetViper().AllSettings())

	return viper.Unmarshal(&config)
}

type Option func(o *option)

func getDefaultConfigFolder() []string {
	return []string{"./internal/configs2"}
}

func getDefaultConfigFile() string {
	return "config2"
}

func getDefaultConfigType() string {
	return "yaml"
}

func WithConfigFolder(configFolder []string) Option {
	return func(o *option) {
		o.configFolders = configFolder
	}
}

func WithConfigFile(configFile string) Option {
	return func(o *option) {
		o.configFile = configFile
	}
}

func WithConfigType(configType string) Option {
	return func(o *option) {
		o.configType = configType
	}
}

func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
