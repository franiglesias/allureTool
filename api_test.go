package main

import (
	"allureTool/config"
	"allureTool/source/api"
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestApiGetProject(t *testing.T) {
	projects := config.ReadFrom("./../data/projects.csv")

	for _, project := range projects {
		downloadProject(project)
	}
}

func downloadProject(project string) {
	c := config.GetConfig()
	env, err := c.LoadEnv("./.env")
	if err != nil {
		return
	}

	var pClient = api.Client{
		BaseUrl: api.PathString(env.BaseUrl),
		Server:  api.PathString(env.Server),
		Credentials: api.Credentials{
			Username: env.Username,
			Password: env.Password,
		},
	}

	cookie, _ := pClient.Login()
	bytes, err := pClient.Download("backend", cookie)
	if err != nil {
		return
	}

	tempZip := ("/tmp/") + project + ".zip"
	writeZipFile(bytes, tempZip)
	_ = unzipSource(project, tempZip, "./../data/allure", "/tmp")
	_ = os.Remove(tempZip)
}

func unzipFile(f *zip.File, destination string) error {
	filePath := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}
	return nil
}

func unzipSource(project, source, destination, temporal string) error {
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	temporal, err = filepath.Abs(temporal)
	if err != nil {
		return err
	}

	// 3. Iterate over zip files inside the archive and unzip each of them
	prefix := "allure-report"
	extension := ".csv"
	behaviorsFile := prefix + "/data/behaviors" + extension

	for _, f := range reader.File {
		if f.Name != behaviorsFile {
			continue
		}

		err := unzipFile(f, temporal)
		if err != nil {
			return err
		}
		err = os.Rename(temporal+string(os.PathSeparator)+behaviorsFile, destination+string(os.PathSeparator)+project+extension)
		if err != nil {
			return err
		}
		err = os.RemoveAll(temporal + string(os.PathSeparator) + prefix)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeZipFile(bytes []byte, name string) {
	f, err := os.Create(name)
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}
	defer f.Close()

	_, err = f.Write(bytes)
	err = f.Sync()
	if err != nil {
		return
	}

	reader, err := zip.OpenReader(name)
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}

	defer reader.Close()

	if len(bytes) == 0 {
		_ = fmt.Errorf("bad things happend")
	}
}
