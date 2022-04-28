package config

import (
	"os"
	"reflect"
	"testing"
)
import "github.com/spf13/viper"

func TestLoadSecretConfigFromDotEnv(t *testing.T) {
	_ = setFakeEnvFileForTesting(".test.env", t)

	c := Config{}

	got, _ := c.LoadEnv(".test.env")

	want := Config{
		output:   "",
		reports:  "",
		baseDir:  "",
		filters:  "",
		baseUrl:  "https://example.com",
		server:   "/server/path",
		password: "secret",
		username: "username",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("conf got %+v, want %+v", got, want)
	}

	clean("./.test.env")
}

func TestLoadConfigFromConfigFile(t *testing.T) {
	_ = setFakeConfigFileForTesting("config.yml", t)

	c := Config{}

	got, _ := c.LoadConf("config.yml")

	want := Config{
		output:   "output.csv",
		reports:  "allure",
		baseDir:  "./data",
		filters:  "filters.csv",
		baseUrl:  "",
		server:   "",
		password: "",
		username: "",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("conf got %+v, want %+v", got, want)
	}

	clean("./config.yml")
}

func TestLoadAllConfiguration(t *testing.T) {
	_ = setFakeEnvFileForTesting(".test.env", t)
	_ = setFakeConfigFileForTesting("config.yml", t)

	c := Config{}

	c, _ = c.LoadConf("config.yml")
	got, _ := c.LoadEnv(".test.env")

	want := Config{
		output:   "output.csv",
		reports:  "allure",
		baseDir:  "./data",
		filters:  "filters.csv",
		baseUrl:  "https://example.com",
		server:   "/server/path",
		password: "secret",
		username: "username",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("conf got %+v, want %+v", got, want)
	}

	clean("./config.yml")
	clean("./.test.env")
}

func TestOverrideConfigurationWithCLIOptions(t *testing.T) {
	c := Config{
		output:   "output.csv",
		reports:  "allure",
		baseDir:  "./data",
		filters:  "filters.csv",
		baseUrl:  "",
		server:   "",
		password: "",
		username: "",
	}

	args := []string{
		"-output",
		"custom.csv",
		"-source",
		"custom",
		"-base",
		"./custom",
		"-filters",
		"custom-filters.csv",
	}
	got, _ := c.LoadFlags("prog", args)

	want := Config{
		output:   "custom.csv",
		reports:  "custom",
		baseDir:  "./custom",
		filters:  "custom-filters.csv",
		baseUrl:  "",
		server:   "",
		password: "",
		username: "",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("conf got \n%+v\n, want \n%+v", got, want)
	}
}

func clean(name string) {
	err := os.Remove(name)
	if err != nil {
		return
	}
}

func setFakeEnvFileForTesting(file string, t *testing.T) error {
	viper.Reset()
	viper.SetConfigType("env")

	viper.Set("ALLURE_BASE_URL", "https://example.com")
	viper.Set("ALLURE_SERVER_PATH", "/server/path")
	viper.Set("ALLURE_PASSWORD", "secret")
	viper.Set("ALLURE_USERNAME", "username")

	viper.AddConfigPath(".")
	err := viper.WriteConfigAs(file)
	if err != nil {
		t.Errorf(err.Error())
	}
	return err
}

func setFakeConfigFileForTesting(file string, t *testing.T) error {
	viper.Reset()
	viper.SetConfigType("yaml")

	viper.Set("output", "output.csv")
	viper.Set("source", "allure")
	viper.Set("filters", "filters.csv")
	viper.Set("base", "./data")

	viper.AddConfigPath(".")
	err := viper.WriteConfigAs(file)
	if err != nil {
		t.Errorf(err.Error())
	}
	return err
}
