package config

import (
	"github.com/spf13/afero"
	"os"
	"reflect"
	"testing"
)

func TestLoadSecretConfigFromDotEnv(t *testing.T) {
	c := Config{
		Fs:  afero.NewMemMapFs(),
		Env: ".test.env",
	}

	_ = setFakeEnvFileForTesting(".test.env", c.Fs)

	got, _ := c.LoadEnv()

	want := Config{
		Fs:       c.Fs,
		Env:      c.Env,
		Conf:     c.Conf,
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
		t.Fail()
	}
}

func TestLoadConfigFromConfigFile(t *testing.T) {
	c := Config{
		Fs:   afero.NewMemMapFs(),
		Conf: "config.yml",
	}

	_ = setFakeConfigFileForTesting("config.yml", c.Fs)

	got, _ := c.LoadConf()

	want := Config{
		Fs:       c.Fs,
		Env:      c.Env,
		Conf:     c.Conf,
		output:   "output.csv",
		reports:  "allure",
		BaseDir:  "./data",
		filters:  "filters-old.csv",
		projects: "projects.csv",
		BaseUrl:  "",
		Server:   "",
		Password: "",
		Username: "",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("conf got %+v, want %+v", got, want)
		t.Fail()
	}
}

func TestLoadAllConfiguration(t *testing.T) {
	c := Config{
		Fs:   afero.NewMemMapFs(),
		Env:  ".test.env",
		Conf: "config.yml",
	}

	_ = setFakeEnvFileForTesting(".test.env", c.Fs)
	_ = setFakeConfigFileForTesting("config.yml", c.Fs)

	got := c.Get()

	want := Config{
		Fs:       c.Fs,
		Env:      c.Env,
		Conf:     c.Conf,
		output:   "output.csv",
		reports:  "allure",
		BaseDir:  "./data",
		filters:  "filters-old.csv",
		projects: "projects.csv",
		BaseUrl:  "https://example.com",
		Server:   "/server/path",
		Password: "secret",
		Username: "username",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("conf got %+v, want %+v", got, want)
		t.Fail()
	}
}

func TestOverrideConfigurationWithCLIOptions(t *testing.T) {
	c := Config{
		output:   "output.csv",
		reports:  "allure",
		BaseDir:  "./data",
		filters:  "filters-old.csv",
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
		"custom-filters-old.csv",
	}
	got, _ := c.LoadFlags("program", args)

	want := Config{
		output:   "custom.csv",
		reports:  "custom",
		BaseDir:  "./custom",
		filters:  "custom-filters-old.csv",
		BaseUrl:  "",
		Server:   "",
		Password: "",
		Username: "",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("conf got \n%+v\n, want \n%+v", got, want)
	}
}

func setFakeEnvFileForTesting(file string, fs afero.Fs) error {
	data := []byte(
		"ALLURE_BASE_URL=https://example.com\n" +
			"ALLURE_PASSWORD=secret\n" +
			"ALLURE_SERVER_PATH=/server/path\n" +
			"ALLURE_USERNAME=username\n",
	)

	return afero.WriteFile(fs, file, data, os.ModeAppend)
}

func setFakeConfigFileForTesting(file string, fs afero.Fs) error {
	data := []byte(
		"base: ./data\n" +
			"filters: filters-old.csv\n" +
			"output: output.csv\n" +
			"source: allure\n" +
			"projects: projects.csv\n",
	)

	return afero.WriteFile(fs, file, data, os.ModeAppend)
}
