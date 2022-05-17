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
		BaseDir:  "",
		filters:  "",
		BaseUrl:  "https://example.com",
		Server:   "/server/path",
		Password: "secret",
		Username: "username",
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
		BaseDir:  "./data",
		filters:  "filters.csv",
		BaseUrl:  "",
		Server:   "",
		Password: "",
		Username: "",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("conf got %+v, want %+v", got, want)
	}

	clean("./config.yml")
}

func TestLoadAllConfiguration(t *testing.T) {
	_ = setFakeEnvFileForTesting(".test.env", t)
	_ = setFakeConfigFileForTesting("config.yml", t)

	got := GetConfig()

	want := Config{
		output:   "output.csv",
		reports:  "allure",
		BaseDir:  "./data",
		filters:  "filters.csv",
		BaseUrl:  "https://example.com",
		Server:   "/server/path",
		Password: "secret",
		Username: "username",
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
		BaseDir:  "./data",
		filters:  "filters.csv",
		BaseUrl:  "",
		Server:   "",
		Password: "",
		Username: "",
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
	got, _ := c.LoadFlags("program", args)

	want := Config{
		output:   "custom.csv",
		reports:  "custom",
		BaseDir:  "./custom",
		filters:  "custom-filters.csv",
		BaseUrl:  "",
		Server:   "",
		Password: "",
		Username: "",
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
		t.Logf(err.Error())
		t.Fail()
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
		t.Log(err.Error())
		t.Fail()
	}
	return err
}
