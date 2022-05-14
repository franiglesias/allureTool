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
		{name: "Lack of trailing backslash", base: "https://allure.com", server: "server", want: "https://allure.com/server"},
		{name: "Server comes with backslash", base: "https://allure.com", server: "server/", want: "https://allure.com/server"},
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
				t.Errorf("Expected %s, got %s", test.want, got)
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
		t.Errorf(err.Error())
	}

	if got.Name != want.Name {
		t.Errorf("Expected %s, got %s", want.Name, got.Name)
	}
	if got.Value != want.Value {
		t.Errorf("Expected %s, got %s", want.Value, got.Value)
	}
	if got.MaxAge != want.MaxAge {
		t.Errorf("Expected %d, got %d", want.MaxAge, got.MaxAge)
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
		t.Errorf("It should have failed because invalid login, but was ok")
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
		t.Errorf("It should have found something")
	}

	if string(contents) != payload.String() {
		t.Errorf("Expected info, got %s", contents)
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
