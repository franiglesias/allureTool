package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

type Client struct {
	BaseUrl     PathString
	Server      PathString
	Credentials Credentials
}

func (api Client) Uri() PathString {
	return api.BaseUrl.WithSchema().WithTrailingSlash() + api.Server.WithoutSlashes()
}

func (api Client) Endpoint(endpoint string) string {
	return string(api.Uri()) + "/" + endpoint
}

func (api Client) Login() (*http.Cookie, error) {

	jsonReq, err := json.Marshal(api.Credentials)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}

	response, e := client.Post(
		api.Endpoint("login"),
		"application/json",
		bytes.NewBuffer(jsonReq),
	)

	if e != nil {
		fmt.Errorf("Something happended")
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("invalid credentials")
	}

	loginResponse, err := GetLoginResponse(response)

	if err != nil {
		return nil, err
	}

	return &http.Cookie{
		Name:   "access_token_cookie",
		Value:  loginResponse.accessToken(),
		MaxAge: loginResponse.maxAge(),
	}, nil
}

func (api Client) Download(projectId string, cookie *http.Cookie) ([]byte, error) {

	client, err := makeHttpClient()
	if err != nil {
		return []byte(""), err
	}

	request, err := http.NewRequest("GET", api.Endpoint("report/export?project_id="+projectId), nil)
	if err != nil {
		return []byte(""), err
	}
	if cookie != nil {
		request.AddCookie(cookie)
	}

	resp, err := client.Do(request)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return bodyBytes, nil
}

func makeHttpClient() (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Jar: jar,
	}
	return client, nil
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}