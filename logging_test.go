package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestLogRequests(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	t.Cleanup(func() { log.SetOutput(os.Stderr) })

	mux := http.NewServeMux()
	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("hello")) })
	mux.HandleFunc("/boom", func(http.ResponseWriter, *http.Request) { panic("kaboom") })
	h := logRequests(mux)

	tests := []struct {
		path       string
		htmx       bool
		wantStatus int
		wantLog    []string
	}{
		{"/ok", true, http.StatusOK, []string{"GET /ok 200 5B", "htmx=true"}},
		{"/missing", false, http.StatusNotFound, []string{"GET /missing 404", "htmx=false"}},
		{"/boom", false, http.StatusInternalServerError, []string{"panic: kaboom", "GET /boom 500"}},
	}
	for _, tc := range tests {
		buf.Reset()
		req := httptest.NewRequest(http.MethodGet, tc.path, nil)
		if tc.htmx {
			req.Header.Set("HX-Request", "true")
		}
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)

		if rec.Code != tc.wantStatus {
			t.Errorf("%s: status = %d, want %d", tc.path, rec.Code, tc.wantStatus)
		}
		for _, s := range tc.wantLog {
			if !strings.Contains(buf.String(), s) {
				t.Errorf("%s: log missing %q\n%s", tc.path, s, buf.String())
			}
		}
	}
}
