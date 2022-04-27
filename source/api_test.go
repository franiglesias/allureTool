package source

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type AllureApi struct {
	baseUrl     string
	server      string
	endpoint    string
	credentials Credentials
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Data struct {
		AccessToken  string   `json:"access_token"`
		ExpiresIn    float32  `json:"expires_in"`
		RefreshToken string   `json:"refresh_token"`
		Roles        []string `json:"roles"`
	} `json:"data"`
	MetaData struct {
		Message string `json:"message"`
	} `json:"meta_data"`
}

func (r AllureApi) Uri() string {
	return r.baseUrl + r.server + "/" + r.endpoint
}

func (r AllureApi) Login() (*http.Cookie, error) {
	loginEndpoint := AllureApi{
		baseUrl:  r.baseUrl,
		server:   r.server,
		endpoint: "login",
	}

	jsonReq, err := json.Marshal(r.credentials)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(loginEndpoint.Uri(), "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var loginResponse LoginResponse

	err = json.Unmarshal(bodyBytes, &loginResponse)

	if err != nil {
		return nil, err
	}

	return &http.Cookie{
		Name:   "access_token_cookie",
		Value:  loginResponse.Data.AccessToken,
		MaxAge: int(loginResponse.Data.ExpiresIn),
	}, nil
}

func (r AllureApi) Get() (string, error) {
	cookie, err := r.Login()
	if err != nil {
		return "", err
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return "", err
	}

	client := http.Client{
		Jar: jar,
	}

	request, err := http.NewRequest("GET", r.Uri(), nil)
	if err != nil {
		return "", err
	}

	if cookie != nil {
		request.AddCookie(cookie)
	}

	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

type Project struct {
	name   string
	export AllureApi
}

func NewProject(name string) *Project {
	return &Project{
		name: name,
		export: AllureApi{
			baseUrl:  "http://allure.lni.core.dev.navify.com",
			server:   "/allure-api/allure-docker-service",
			endpoint: "report/export?project_id=" + name,
			credentials: Credentials{
				Password: "suh2QALJ",
				Username: "viewer",
			},
		},
	}
}

func (p Project) Download() {
	response, _ := p.export.Get()

	temporal := "/tmp/"
	writeZipFile(response, temporal+p.name+".zip")
	_ = unzipSource(p.name, temporal+p.name+".zip", "./../data/allure", "/tmp")
	_ = os.Remove(temporal + p.name + ".zip")
}

func TestApiGetProject(t *testing.T) {
	projects := []string{
		"backend",
		"data-services",
		"default",
		"infrastructure",
		"lni",
		"lni-anonymization",
		"system",
		"system-test",
	}

	for _, project := range projects {
		downloadProject(project)
	}
}

func downloadProject(project string) {
	exportProject := AllureApi{
		baseUrl:  "http://allure.lni.core.dev.navify.com",
		server:   "/allure-api/allure-docker-service",
		endpoint: "report/export?project_id=" + project,
		credentials: Credentials{
			Password: "suh2QALJ",
			Username: "viewer",
		},
	}

	response, _ := exportProject.Get()

	temporal := "/tmp/"
	writeZipFile(response, temporal+project+".zip")
	_ = unzipSource(project, temporal+project+".zip", "./../data/allure", "/tmp")
	_ = os.Remove(temporal + project + ".zip")
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

func writeZipFile(response string, name string) {
	f, err := os.Create(name)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer f.Close()

	_, err = f.WriteString(response)
	f.Sync()

	reader, err := zip.OpenReader(name)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	defer reader.Close()

	if response == "" {
		fmt.Errorf("Bad things happend")
	}
}
