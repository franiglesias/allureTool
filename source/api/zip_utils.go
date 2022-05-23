package api

import (
	"archive/zip"
	"fmt"
	"github.com/spf13/afero"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func UnzipSource(project, source, destination string, fs afero.Fs) error {
	prefix := "allure-report"
	extension := ".csv"
	destinationFile := destination + project + extension
	behaviorsFile := prefix + "/data/behaviors" + extension

	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	temporal, err := filepath.Abs("/tmp")
	if err != nil {
		return err
	}

	for _, f := range reader.File {
		if f.Name != behaviorsFile {
			continue
		}

		err := unzipFile(f, temporal, fs)
		if err != nil {
			return err
		}
		err = fs.Rename(temporal+string(os.PathSeparator)+behaviorsFile, destinationFile)
		if err != nil {
			return err
		}
		err = fs.RemoveAll(temporal + string(os.PathSeparator) + prefix)
		if err != nil {
			return err
		}
	}

	return nil
}

func unzipFile(f *zip.File, destination string, fs afero.Fs) error {
	filePath := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	if f.FileInfo().IsDir() {
		if err := fs.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err := fs.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	destinationFile, err := fs.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
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
