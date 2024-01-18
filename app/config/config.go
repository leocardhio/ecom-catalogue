package config

import (
	"errors"
	"flag"
	"path"
	"path/filepath"
	"runtime"

	"github.com/leocardhio/ecom-catalogue/util"
	"github.com/spf13/viper"
)

const (
	InvalidEnvErr = "invalid env"
	MissingEnvErr = "missing env"
)

var (
	env        string
	_, b, _, _ = runtime.Caller(0)
	envPath    = path.Join(filepath.Dir(b), "env")
)

type Config struct {
	SQLHost     string `mapstructure:"SQL_HOST"`
	SQLPort     string `mapstructure:"SQL_PORT"`
	SQLDatabase string `mapstructure:"SQL_DATABASE"`
	SQLDriver   string `mapstructure:"SQL_DRIVER"`

	SQLUser     string `mapstructure:"SQL_USER"`
	SQLPassword string `mapstructure:"SQL_PASSWORD"`

	ESHost string `mapstructure:"ES_HOST"`
	ESPort string `mapstructure:"ES_PORT"`
}

func init() {
	args := util.DeclareFlag()
	flag.Parse()

	env = *args.Env
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() *Config {
	var err error

	switch env {
	case "dev":
		err = c.loadDev()
	case "test":
		err = c.loadTest()
	case "prod":
		err = c.loadProd()
	case "":
		err = errors.New(MissingEnvErr)
	default:
		err = errors.New(InvalidEnvErr)
	}

	if err != nil {
		panic(err)
	}

	return c
}

func (c *Config) loadDev() error {
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AddConfigPath(envPath)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}

func (c *Config) loadTest() error {
	// TODO: Implement me
	return errors.New("not implemented")
}

func (c *Config) loadProd() error {
	// TODO: Implement me
	return errors.New("not implemented")
}
