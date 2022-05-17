package config

import (
	"flag"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	output   string
	reports  string
	BaseDir  string
	filters  string
	BaseUrl  string
	Server   string
	Password string
	Username string
}

func GetConfig() Config {
	c := Config{}
	c, _ = c.LoadConf("config.yml")
	c, _ = c.LoadEnv(".test.env")

	return c
}

func (c Config) PathToReports() string {
	return filepath.Join(c.BaseDir, c.reports) + string(os.PathSeparator)
}

func (c Config) OutputFile() string {
	return filepath.Join(c.BaseDir, c.output)
}

func (c Config) FiltersFile() string {
	return filepath.Join(c.BaseDir, c.filters)
}

func (c Config) LoadEnv(envFile string) (Config, error) {
	viper.SetConfigName(envFile)
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	n := Config{
		output:   c.output,
		reports:  c.reports,
		BaseDir:  c.BaseDir,
		filters:  c.filters,
		BaseUrl:  viper.Get("ALLURE_BASE_URL").(string),
		Server:   viper.Get("ALLURE_SERVER_PATH").(string),
		Password: viper.Get("ALLURE_PASSWORD").(string),
		Username: viper.Get("ALLURE_USERNAME").(string),
	}

	return n, nil
}

func (c Config) LoadConf(confFile string) (Config, error) {
	viper.SetConfigName(confFile)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	n := Config{
		output:   viper.Get("output").(string),
		reports:  viper.Get("source").(string),
		BaseDir:  viper.Get("base").(string),
		filters:  viper.Get("filters").(string),
		BaseUrl:  c.BaseUrl,
		Server:   c.Server,
		Password: c.Password,
		Username: c.Username,
	}

	return n, nil
}

func (c Config) LoadFlags(program string, args []string) (Config, error) {
	f := flag.NewFlagSet(program, flag.ContinueOnError)

	n := c
	f.StringVar(&n.output, "output", "output.csv", "File to generate results report")
	f.StringVar(&n.reports, "source", "allure", "Folder where report files are stored")
	f.StringVar(&n.filters, "filters", "filters.csv", "List of labels we want to find")
	f.StringVar(&n.BaseDir, "base", "./data/", "Base folder for working")

	err := f.Parse(args)

	if err != nil {
		return n, err
	}

	return n, nil
}
