package project

import (
	"allureTool/config"
	"allureTool/source/api"
	"allureTool/source/zip"
)

type Project struct {
	Name   string
	Config config.Config
}

func (p Project) TmpZip() string {
	return "/tmp/" + p.Name + ".zip"
}

func (p Project) GetData() error {
	err := p.downloadZip(api.MakeClient(p.Config))
	if err != nil {
		return err
	}

	err = p.extractData()
	if err != nil {
		return err
	}

	return p.cleanAfter()
}

func (p Project) downloadZip(api api.Client) error {
	cookie, err := api.Login()
	if err != nil {
		return err
	}
	bytes, err := api.Download(p.Name, cookie)
	if err != nil {
		return err
	}
	return config.NewDataFile(p.TmpZip(), p.Config.Fs).WriteBytes(bytes)
}

func (p Project) extractData() error {
	return zip.UnzipSource(p.Name, p.TmpZip(), p.Config.PathToReports(), p.Config.Fs)
}

func (p Project) cleanAfter() error {
	return p.Config.Fs.Remove(p.TmpZip())
}
