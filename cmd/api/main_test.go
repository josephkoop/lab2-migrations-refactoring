//This is not my code this is straight AI-generated.


package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(home)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	expected := "Welcome! This is Joseph Koop's territory."
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("response body does not contain expected content")
	}
}

func TestAboutHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/about", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(about)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	expected := "About me"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("response body does not contain expected content")
	}
}

func TestContactHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/contact", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(contact)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	expected := "Email:"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("response body does not contain expected content")
	}
}

func TestHobbyHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hobby", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(hobby)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	expected := "Hobby: Chess"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("response body does not contain expected content")
	}
}
