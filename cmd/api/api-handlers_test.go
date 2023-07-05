package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
	"webapp/pkg/data"

	"github.com/go-chi/chi/v5"
)

func Test_app_authenticate(t *testing.T) {
	var theTests = []struct {
		name               string
		requestBody        string
		expectedStatusCode int
	}{
		{"valid user", `{"email": "admin@example.com", "password":"secret"}`, http.StatusOK},
		{"not json", `I'm not JSON`, http.StatusUnauthorized},
		{"empty json", `{}`, http.StatusUnauthorized},
		{"empty email", `{"email": ""}`, http.StatusUnauthorized},
		{"empty password", `{"email": "admin@example.com"}`, http.StatusUnauthorized},
		{"invalid user", `{"email": "admin@somerother.com", "password":"secret"}`, http.StatusUnauthorized},
	}

	for _, e := range theTests {
		var reader io.Reader
		reader = strings.NewReader(e.requestBody)
		req, _ := http.NewRequest("POST", "/auth", reader)
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(app.autheticate)

		handler.ServeHTTP(rr, req)

		if e.expectedStatusCode != rr.Code {
			t.Errorf("%s: returned wrong status code: expected %d but got %d", e.name, e.expectedStatusCode, rr.Code)
		}
	}
}

func Test_app_refresh(t *testing.T) {
	var tests = []struct {
		name               string
		token              string
		expectedStatusCode int
		resetRefreshTime   bool
	}{
		{name: "valid token", token: "", expectedStatusCode: http.StatusOK, resetRefreshTime: true},
		{name: "valid but not yet ready to expire", token: "", expectedStatusCode: http.StatusTooEarly, resetRefreshTime: false},
		{name: "expired token", token: expiredToken, expectedStatusCode: http.StatusBadRequest, resetRefreshTime: false},
	}

	testUser := data.User{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
		Email:     "admin@example.com",
	}

	oldRefreshTime := refreshTokenExpiry

	for _, e := range tests {
		var tkn string
		if e.token == "" {
			if e.resetRefreshTime {
				refreshTokenExpiry = time.Second * 1
			}

			tokens, _ := app.generateTokenPair(&testUser)
			tkn = tokens.RefreshToken
		} else {
			tkn = e.token
		}

		postData := url.Values{
			"refresh_token": {tkn},
		}

		req, _ := http.NewRequest("POST", "/refresh-token", strings.NewReader(postData.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(app.refresh)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatusCode {
			t.Errorf("%s: expected status of %d, but got %d", e.name, e.expectedStatusCode, rr.Code)
		}
		refreshTokenExpiry = oldRefreshTime
	}

}

func Test_app_userHandlers(t *testing.T) {
	var tests = []struct {
		name               string
		method             string
		json               string
		paramID            string
		handler            http.HandlerFunc
		expectedStatusCode int
	}{
		{"all users", "GET", "", "", app.allUsers, http.StatusOK},
		{"delete user", "DELETE", "", "1", app.deleteUser, http.StatusNoContent},
		{"getUser valid", "GET", "", "1", app.getUser, http.StatusOK},
		{"getUser invalid", "GET", "", "100", app.getUser, http.StatusBadRequest},
	}

	for _, e := range tests {
		var req *http.Request
		if e.json == "" {
			req, _ = http.NewRequest(e.method, "/", nil)
		} else {
			req, _ = http.NewRequest(e.method, "/", strings.NewReader(e.json))
		}

		if e.paramID != "" {
			chiCtx := chi.NewRouteContext()
			chiCtx.URLParams.Add("userID", e.paramID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(e.handler)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatusCode {
			t.Errorf("%s: wrong status returned; expected %d but got %d", e.name, e.expectedStatusCode, rr.Code)
		}
	}
}
