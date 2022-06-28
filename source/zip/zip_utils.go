package zip

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
	target := NewTarget(project, destination)

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
		if target.IsNotNamed(f.Name) {
			continue
		}

		err := obtainFile(f, temporal, fs, target)
		if err != nil {
			return err
		}
	}

	return nil
}

func obtainFile(f *zip.File, temporal string, fs afero.Fs, target Target) error {
	err := unzipFile(f, temporal, fs)
	if err != nil {
		return err
	}
	err = target.MoveToDestinationInFs(fs)
	if err != nil {
		return err
	}
	err = target.RemoveFromFs(fs)
	if err != nil {
		return err
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

	err := fs.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
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
