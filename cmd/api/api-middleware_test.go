package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"webapp/pkg/data"
)

func Test_app_enableCORS(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	var tests = []struct {
		name         string
		method       string
		expectHeader bool
	}{
		{"preflight", "OPTIONS", true},
		{"get", "GET", false},
	}

	for _, e := range tests {
		handlerToTest := app.enableCORS(nextHandler)

		req := httptest.NewRequest(e.method, "http://testing", nil)
		rr := httptest.NewRecorder()

		handlerToTest.ServeHTTP(rr, req)

		if e.expectHeader && rr.Header().Get("Access-Control-Allow-Credentials") == "" {
			t.Errorf("%s: expected header, but did not find it", e.name)
		}

		if !e.expectHeader && rr.Header().Get("Access-Control-Allow-Credentials") != "" {
			t.Errorf("%s: expected no header, but got one", e.name)
		}
	}
}

func Test_app_authRequire(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	testUser := data.User{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
		Email:     "admin@example.com",
	}

	tokens, _ := app.generateTokenPair(&testUser)

	var tests = []struct {
		name               string
		token              string
		expectedAuthorized bool
		setHeader          bool
	}{
		{name: "valid token", token: fmt.Sprintf("Bearer %s", tokens.Token), expectedAuthorized: true, setHeader: true},
		{name: "no token", token: "", expectedAuthorized: false, setHeader: false},
		{name: "invalid token", token: fmt.Sprintf("Bearer %s", expiredToken), expectedAuthorized: false, setHeader: true},
	}

	for _, e := range tests {
		req, _ := http.NewRequest("GET", "/", nil)
		if e.setHeader {
			req.Header.Set("Authorization", e.token)
		}
		rr := httptest.NewRecorder()

		handlerToTest := app.authRequire(nextHandler)
		handlerToTest.ServeHTTP(rr, req)

		if e.expectedAuthorized && rr.Code == http.StatusUnauthorized {
			t.Errorf("%s: got code 401, and should not have", e.name)
		}

		if !e.expectedAuthorized && rr.Code != http.StatusUnauthorized {
			t.Errorf("%s: dit not get code 401, and should have", e.name)
		}
	}
}
