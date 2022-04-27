package config

import (
	"flag"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	output   string
	reports  string
	baseDir  string
	filters  string
	baseUrl  string
	server   string
	password string
	username string
}

func GetConfig() Config {
	output := flag.String("output", "output.csv", "File to generate results report")
	source := flag.String("source", "allure", "Folder where report files are stored")
	filters := flag.String("filters", "filters.csv", "List of labels we want to find")
	baseDir := flag.String("base", "./data/", "Base folder for working")
	flag.Parse()

	return Config{
		output:  *output,
		reports: *source,
		baseDir: strings.TrimSuffix(*baseDir, string(os.PathSeparator)),
		filters: *filters,
	}
}

func (c Config) PathToReports() string {
	return filepath.Join(c.baseDir, c.reports) + string(os.PathSeparator)
}

func (c Config) OutputFile() string {
	return filepath.Join(c.baseDir, c.output)
}

func (c Config) FiltersFile() string {
	return filepath.Join(c.baseDir, c.filters)
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
		baseDir:  c.baseDir,
		filters:  c.filters,
		baseUrl:  viper.Get("ALLURE_BASE_URL").(string),
		server:   viper.Get("ALLURE_SERVER_PATH").(string),
		password: viper.Get("ALLURE_PASSWORD").(string),
		username: viper.Get("ALLURE_USERNAME").(string),
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
		baseDir:  viper.Get("base").(string),
		filters:  viper.Get("filters").(string),
		baseUrl:  c.baseUrl,
		server:   c.server,
		password: c.password,
		username: c.username,
	}

	return n, nil
}
