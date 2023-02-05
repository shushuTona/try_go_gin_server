package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("expected code %v; got %v", "200", w.Code)
	}

	if w.Body. String() != "{\"message\":\"pong\"}" {
		t.Errorf("expected body %v; got %v", "pong", w.Body.String())
	}
}