package handler

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHello(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rec := httptest.NewRecorder()

    Hello(rec, req)

    if rec.Code != http.StatusOK {
        t.Fatalf("expected status 200, got %d", rec.Code)
    }
}

func TestHealth(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/health", nil)
    rec := httptest.NewRecorder()

    Health(rec, req)

    if rec.Code != http.StatusOK {
        t.Fatalf("expected status 200, got %d", rec.Code)
    }
}
