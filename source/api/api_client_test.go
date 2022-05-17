package api

import (
	"bytes"
	"compress/gzip"
	"log"
	"net/http"
	"testing"
)

func TestApiClientUriShould(t *testing.T) {
	tests := []struct {
		name   string
		base   PathString
		server PathString
		want   PathString
	}{
		{name: "All fine", base: "https://allure.com/", server: "server", want: "https://allure.com/server"},
		{name: "Lack of trailing slash", base: "https://allure.com", server: "server", want: "https://allure.com/server"},
		{name: "Server with root slash", base: "https://allure.com", server: "/server", want: "https://allure.com/server"},
		{name: "Server comes with slash", base: "https://allure.com", server: "server/", want: "https://allure.com/server"},
		{name: "Schema defaults to https", base: "allure.com", server: "server", want: "https://allure.com/server"},
		{name: "Not secure http is allowed", base: "http://allure.com", server: "server", want: "http://allure.com/server"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			apiClient := Client{
				BaseUrl: test.base,
				Server:  test.server,
			}

			got := apiClient.Uri()

			if got != test.want {
				t.Logf("Expected %s, got %s", test.want, got)
				t.Fail()
			}
		})
	}
}

func TestApiClientShouldLogin(t *testing.T) {

	response := LoginResponse{
		Data: Data{
			AccessToken:  "some token",
			ExpiresIn:    100.0,
			RefreshToken: "refresh token",
			Roles:        []string{"role1"},
		},
		MetaData: MetaData{
			Message: "Successfully logged",
		},
	}

	server := MakeLoginServer(t, http.StatusOK, response, http.MethodPost, "/server/login")
	defer server.Close()
	apiClient := Client{
		BaseUrl: PathString(server.URL),
		Server:  "server",
	}

	got, err := apiClient.Login()

	want := http.Cookie{
		Name:   "access_token_cookie",
		Value:  "some token",
		MaxAge: 100,
	}

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	if got.Name != want.Name {
		t.Logf("Expected %s, got %s", want.Name, got.Name)
		t.Fail()
	}
	if got.Value != want.Value {
		t.Logf("Expected %s, got %s", want.Value, got.Value)
		t.Fail()
	}
	if got.MaxAge != want.MaxAge {
		t.Logf("Expected %d, got %d", want.MaxAge, got.MaxAge)
		t.Fail()
	}
}

func TestApiClientShouldHandleFailedLogin(t *testing.T) {

	response := LoginResponse{
		MetaData: MetaData{
			Message: "Invalid username/password",
		},
	}
	server := MakeLoginServer(t, http.StatusUnauthorized, response, http.MethodPost, "/server/login")
	defer server.Close()
	apiClient := Client{
		BaseUrl: PathString(server.URL),
		Server:  "server",
	}

	_, err := apiClient.Login()

	if err == nil {
		t.Logf("It should have failed because invalid login, but was ok")
		t.Fail()
	}
}

func TestApiShouldHandleProjectFile(t *testing.T) {
	payload := payload()

	server := MakeDownloadServer(t, http.StatusOK, payload, http.MethodGet, "/server/report/export")

	defer server.Close()

	apiClient := Client{
		BaseUrl: PathString(server.URL),
		Server:  "server",
	}

	cookie := &http.Cookie{Name: "access_token_cookie", Value: "some token"}

	contents, err := apiClient.Download("project", cookie)

	if err != nil {
		t.Logf("It should have found something")
		t.Fail()
	}

	if string(contents) != payload.String() {
		t.Logf("Expected info, got %s", contents)
		t.Fail()
	}
}

func payload() bytes.Buffer {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte("Information for a zip file")); err != nil {
		log.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		log.Fatal(err)
	}
	return b
}
