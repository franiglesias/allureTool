package zip_repository

import (
	"archive/zip"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"io"
	"log"
	"os"
	"path/filepath"
)

type ZipArchive struct {
	Fs   afero.Fs
	Path string
	Tmp  string
}

type ZipArchiveError struct {
	rootCause error
	file      string
}

func (e ZipArchiveError) Error() string {
	return fmt.Sprintf("cannot use ZipArchive %s, %s", e.file, e.rootCause.Error())
}

func (z ZipArchive) ExtractFile(desiredFile string) error {
	reader, err := z.getZipReader()
	if err != nil {
		return errors.Wrap(err, "cannot extract file")
	}

	file, err := z.locateFileInArchive(desiredFile, reader)
	if err != nil {
		return errors.Wrap(err, "cannot extract file")
	}

	err = z.extractFileToTempDir(file)
	if err != nil {
		return errors.Wrap(err, "cannot extract file")
	}

	return nil
}

func (z ZipArchive) getZipReader() (*zip.Reader, error) {
	fileInfo, err := z.Fs.Stat(z.Path)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get zip.Reader")
	}

	file, err := z.Fs.OpenFile(z.Path, os.O_RDONLY, os.FileMode(0644))
	if err != nil {
		return nil, errors.Wrap(err, "cannot get zip.Reader")
	}

	reader, err := zip.NewReader(file, fileInfo.Size())
	if err != nil {
		return nil, errors.Wrap(err, "cannot get zip.Reader")
	}

	return reader, nil
}

func (z ZipArchive) extractFileToTempDir(f *zip.File) error {
	filePath := filepath.Join(z.Tmp, f.Name)

	err := z.Fs.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "cannot extract file from zip")
	}

	destinationFile, err := z.Fs.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	defer closeFile(destinationFile)

	if err != nil {
		return errors.Wrap(err, "cannot extract file from zip")
	}

	fileInArchive, err := f.Open()
	defer fileInArchive.Close()
	if err != nil {
		return errors.Wrap(err, "cannot extract file from zip")
	}

	_, err = io.Copy(destinationFile, fileInArchive)
	if err != nil {
		return errors.Wrap(err, "cannot extract file from zip")
	}

	return nil
}

func closeFile(destinationFile afero.File) {
	err := destinationFile.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (z ZipArchive) locateFileInArchive(desiredFile string, reader *zip.Reader) (*zip.File, error) {
	for _, file := range reader.File {
		if file.Name == desiredFile {
			return file, nil
		}
	}

	return nil, fmt.Errorf("file %s not found in archive", desiredFile)
}

func MakeZipArchive(fs afero.Fs, file string) ZipArchive {
	return ZipArchive{
		Fs:   fs,
		Path: file,
	}
}
