package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

// Move up to the project root so tests can find the /templates folder
func init() {
	os.Chdir("../../")
}

// Testing for 200 status code on home page.
func TestSuccessfullHome200(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()

	Home(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

// Testing for 200 status code on ascii-art page.
func TestSuccesfulAsciiArt200(t *testing.T) {
	formData := strings.NewReader("inputText=hello&bannerType=standard")

	req := httptest.NewRequest("POST", "/ascii-art", formData)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	AsciiArt(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", resp.StatusCode)
	}
}

// Testing for 404 status code.
func TestNotFound404(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", Home)
	mux.HandleFunc("POST /ascii-art", AsciiArt)

	req := httptest.NewRequest("GET", "/this-route-is-fake", nil)

	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected 404 Not Found, got %d", resp.StatusCode)
	}
}

// Testing for 400 status code.
func TestBadRequest400(t *testing.T) {
	data := url.Values{}

	data.Set("inputType", "😁😁😁")
	data.Set("bannerType", "standard")

	req := httptest.NewRequest("POST", "/ascii-art", strings.NewReader(data.Encode()))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	AsciiArt(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

// Testing for 500 status code.
func TestInternalServerError500(t *testing.T) {
	originalDir, _ := os.Getwd()

	os.Chdir(os.TempDir())

	defer os.Chdir(originalDir)

	req := httptest.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()
	Home(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", resp.StatusCode)
	}
}
