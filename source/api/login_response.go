package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Data struct {
	AccessToken  string   `json:"access_token"`
	ExpiresIn    float32  `json:"expires_in"`
	RefreshToken string   `json:"refresh_token"`
	Roles        []string `json:"roles"`
}

type MetaData struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Data     Data     `json:"data"`
	MetaData MetaData `json:"meta_data"`
}

func (r LoginResponse) accessToken() string {
	return r.Data.AccessToken
}

func (r LoginResponse) maxAge() int {
	return int(r.Data.ExpiresIn)
}

func GetLoginResponse(response *http.Response) (LoginResponse, error) {
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return LoginResponse{}, err
	}

	var lr LoginResponse
	err = json.Unmarshal(bodyBytes, &lr)

	if err != nil {
		return LoginResponse{}, err
	}

	return lr, nil
}
