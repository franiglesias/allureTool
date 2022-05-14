package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MakeDownloadServer(t *testing.T, stubStatus int, payload bytes.Buffer, method string, path string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		const cookieName = "access_token_cookie"

		if r.Method != method {
			t.Errorf("Expected %s, but got %s", method, r.Method)
		}

		if r.URL.Path != path {
			t.Errorf("Expecting request to %s, got %s", path, r.URL.Path)
		}

		_, err := r.Cookie(cookieName)
		if err != nil {
			t.Errorf("Expecting Cookie \"%s\", but not sent", cookieName)
		}

		writer.Header().Set("Content-Type", "application/zip")
		writer.Header().Set("Content-Disposition:", " attachment; filename=allure-docker-service-report.zip")
		writer.WriteHeader(stubStatus)

		_, err = writer.Write(payload.Bytes())
		if err != nil {
			return
		}
	}))
	return server
}

func MakeLoginServer(t *testing.T, stubStatus int, response LoginResponse, method string, path string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			t.Errorf("Expected %s, but got %s", method, r.Method)
		}
		if r.URL.Path != path {
			t.Errorf("Expecting request to %s, got %s", path, r.URL.Path)
		}

		writer.WriteHeader(stubStatus)

		body, _ := json.Marshal(response)
		_, err := writer.Write(body)
		if err != nil {
			return
		}
	}))

	return server
}
