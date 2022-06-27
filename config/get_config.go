package config

import (
	"flag"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	Env      string
	Conf     string
	Fs       afero.Fs
	output   string
	Reports  string
	Projects string
	BaseDir  string
	Filters  string
	BaseUrl  string
	Server   string
	Password string
	Username string
}

func (c Config) Get() Config {
	env, _ := c.LoadConf()
	cfg, _ := env.LoadEnv()

	return cfg
}

func (c Config) PathToReports() string {
	return filepath.Join(c.BaseDir, c.Reports) + string(os.PathSeparator)
}

func (c Config) OutputFile() string {
	return filepath.Join(c.BaseDir, c.output)
}

func (c Config) FiltersFile() string {
	return filepath.Join(c.BaseDir, c.Filters)
}

func (c Config) ProjectsFile() string {
	return filepath.Join(c.BaseDir, c.Projects)
}

func (c Config) LoadEnv() (Config, error) {
	viper.SetFs(c.Fs)
	viper.SetConfigFile(c.Env)
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	n := Config{
		Env:      c.Env,
		Conf:     c.Conf,
		Fs:       c.Fs,
		output:   c.output,
		Reports:  c.Reports,
		BaseDir:  c.BaseDir,
		Filters:  c.Filters,
		Projects: c.Projects,
		BaseUrl:  viper.Get("ALLURE_BASE_URL").(string),
		Server:   viper.Get("ALLURE_SERVER_PATH").(string),
		Password: viper.Get("ALLURE_PASSWORD").(string),
		Username: viper.Get("ALLURE_USERNAME").(string),
	}

	return n, nil
}

func (c Config) LoadConf() (Config, error) {
	viper.SetFs(c.Fs)
	viper.SetConfigFile(c.Conf)
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	n := Config{
		Env:      c.Env,
		Conf:     c.Conf,
		Fs:       c.Fs,
		output:   viper.Get("output").(string),
		Reports:  viper.Get("source").(string),
		BaseDir:  viper.Get("base").(string),
		Filters:  viper.Get("filters").(string),
		Projects: viper.Get("projects").(string),
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
	f.StringVar(&n.Reports, "source", "allure", "Folder where report files are stored")
	f.StringVar(&n.Filters, "filters", "filters-old.csv", "List of labels we want to find")
	f.StringVar(&n.BaseDir, "base", "./data/", "Base folder for working")
	f.StringVar(&n.BaseDir, "projects", "projects.csv", "List of projects to check")
	err := f.Parse(args)

	if err != nil {
		return n, err
	}

	return n, nil
}
