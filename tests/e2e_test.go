// +build integration

package tests

import (
	"encoding/json"
	"github.com/rahulvramesh/ocha-url/internal/app/ocha"
	"github.com/rahulvramesh/ocha-url/internal/pkg/bigcache"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRecord(t *testing.T) {
	router := ocha.InitRoutes()
	bigcache.InitializeDataSource()

	w := httptest.NewRecorder()
	payload := strings.NewReader(`{"shortcode": "abcdef","url": "https://google.com"}`)

	req, _ := http.NewRequest("POST", "/shorten", payload)
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, `{"shortcode":"abcdef"}`, w.Body.String())
}

func TestCreateDuplicateRecord(t *testing.T) {
	router := ocha.InitRoutes()
	bigcache.InitializeDataSource()

	w := httptest.NewRecorder()

	// Payload 1
	payload := strings.NewReader(`{"shortcode": "abcdef","url": "https://google.com"}`)

	req, _ := http.NewRequest("POST", "/shorten", payload)
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, `{"shortcode":"abcdef"}`, w.Body.String())

	// Payload 2
	w = httptest.NewRecorder()
	payload = strings.NewReader(`{"shortcode": "abcdef","url": "https://google.com"}`)

	req2, _ := http.NewRequest("POST", "/shorten", payload)
	req2.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req2)

	assert.Equal(t, 409, w.Code)
	assert.Equal(t, `{"code":409,"message":"the the desired shortcode is already in use [shortcodes are case-sensitive]"}`, w.Body.String())
}

func TestGetByShortCode(t *testing.T) {
	router := ocha.InitRoutes()
	bigcache.InitializeDataSource()

	w := httptest.NewRecorder()
	payload := strings.NewReader(`{"shortcode": "abcdef","url": "https://google.com"}`)

	req, _ := http.NewRequest("POST", "/shorten", payload)
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/abcdef", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 302, w.Code)
}

func TestGetByShortNotFound(t *testing.T) {
	router := ocha.InitRoutes()
	bigcache.InitializeDataSource()

	w := httptest.NewRecorder()
	payload := strings.NewReader(`{"shortcode": "abcdef","url": "https://google.com"}`)

	req, _ := http.NewRequest("POST", "/shorten", payload)
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/bcdefg", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, `{"code":404,"message":"the shortcode cannot be found in the system"}`, w.Body.String())
}

func TestEmptyKey(t *testing.T) {
	router := ocha.InitRoutes()
	bigcache.InitializeDataSource()

	w := httptest.NewRecorder()
	payload := strings.NewReader(`{"shortcode": "","url": "https://google.com"}`)

	req, _ := http.NewRequest("POST", "/shorten", payload)
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	shortcode := make(map[string]string)
	_ = json.Unmarshal(w.Body.Bytes(), &shortcode)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/"+shortcode["shortcode"], nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 302, w.Code)
}

//TestRegexValidation - Check regex
func TestRegexValidation(t *testing.T) {
	router := ocha.InitRoutes()
	bigcache.InitializeDataSource()

	w := httptest.NewRecorder()
	payload := strings.NewReader(`{"shortcode": "dse^*)","url": "https://google.com"}`)

	req, _ := http.NewRequest("POST", "/shorten", payload)
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 422, w.Code)
	assert.Equal(t, `{"code":422,"message":"invalid characters found in shortcode"}`, w.Body.String())
}
