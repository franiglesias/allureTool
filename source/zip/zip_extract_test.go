package zip

import (
	"archive/zip"
	"github.com/spf13/afero"
	"testing"
)

func TestZipExtract(t *testing.T) {
	fs := afero.NewMemMapFs()

	// create a zip archive.zip with a file original.file inside
	err := ZipAFile("/archive.zip", "/original.file", []byte("Some initial content"), fs)

	if err != nil {
		t.Logf("Error creating zip file" + err.Error())
		t.Fail()
	}

	err = ZipExtractFrom("/archive.zip", "/original.file", "/myData.file", fs)

	if err != nil {
		t.Logf("Error reading zip file" + err.Error())
		t.Fail()
	}

	stat, err := fs.Stat("myData.file")
	if err != nil {
		t.Logf("Desired file not found" + err.Error())
		t.Fail()
	}

	if stat.Name() != "myData.file" {
		t.Logf("Expected %s, got %s", "myData.file", stat.Name())
	}
}

func ZipAFile(zipArchive string, fileName string, contents []byte, fs afero.Fs) error {
	zipFile, err := fs.Create(zipArchive)
	if err != nil {
		return err
	}
	zipWriter := zip.NewWriter(zipFile)
	defer func(zipWriter *zip.Writer) {
		_ = zipWriter.Close()
	}(zipWriter)

	fileInZip, err := zipWriter.Create(fileName)
	if err != nil {
		return err
	}
	_, err = fileInZip.Write(contents)
	if err != nil {
		return err
	}
	err = zipWriter.Close()
	if err != nil {
		return err
	}
	return nil
}

type VirtualZip struct {
	archive string
	fs      afero.Fs
}

func (vz *VirtualZip) ReadAt(p []byte, offset int64) (int, error) {
	reader, err := vz.fs.Open(vz.archive)
	if err != nil {
		return 0, err
	}
	defer func(reader afero.File) {
		_ = reader.Close()
	}(reader)

	bytesRead, err := reader.Read(p)
	if err != nil {
		return 0, err
	}

	return bytesRead, nil
}

func (vz VirtualZip) Size() (int64, error) {
	data, err := afero.ReadFile(vz.fs, vz.archive)
	if err != nil {
		return 0, err
	}

	return int64(len(data)), nil
}

func ZipExtractFrom(zipArchive, fileToExtract, destination string, fs afero.Fs) error {

	readerAt := VirtualZip{
		archive: zipArchive,
		fs:      fs,
	}

	size, err := readerAt.Size()

	if err != nil {
		return err
	}

	zipReader, err := zip.NewReader(&readerAt, size)

	if err != nil {
		return err
	}

	for _, file := range zipReader.File {
		if file.Name != fileToExtract {
			continue
		}
	}

	return nil
}
