package main

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func testServer(t *testing.T) http.Handler {
	t.Helper()
	tmpl := template.Must(template.ParseFS(templatesFS, "templates/*.html"))
	return newServer(tmpl)
}

func TestRoutes(t *testing.T) {
	srv := testServer(t)

	tests := []struct {
		name        string
		method      string
		path        string
		form        url.Values
		wantStatus  int
		wantContain []string
		wantAbsent  []string
	}{
		{
			name:        "healthz",
			method:      http.MethodGet,
			path:        "/healthz",
			wantStatus:  http.StatusOK,
			wantContain: []string{"ok"},
		},
		{
			name:        "index page",
			method:      http.MethodGet,
			path:        "/",
			wantStatus:  http.StatusOK,
			wantContain: []string{`hx-post="/analyze"`, `hx-target="#results"`, "htmx.org", "data-version>dev<"},
		},
		{
			name:       "unknown path",
			method:     http.MethodGet,
			path:       "/nope",
			wantStatus: http.StatusNotFound,
		},
		{
			name:        "analyze returns fragment",
			method:      http.MethodPost,
			path:        "/analyze",
			form:        url.Values{"text": {"Go is fast. Go is simple!"}},
			wantStatus:  http.StatusOK,
			wantContain: []string{`data-stat="words">6<`, `data-stat="sentences">2<`, `<code class="font-medium">go</code>`, `data-count="2"`, "Rendered by Go in"},
			wantAbsent:  []string{"<html", "<head"},
		},
		{
			name:        "analyze empty input",
			method:      http.MethodPost,
			path:        "/analyze",
			form:        url.Values{"text": {"   "}},
			wantStatus:  http.StatusOK,
			wantContain: []string{"Nothing to analyze yet"},
		},
		{
			name:       "analyze escapes html",
			method:     http.MethodPost,
			path:       "/analyze",
			form:       url.Values{"text": {"<script>alert(1)</script>"}},
			wantStatus: http.StatusOK,
			wantAbsent: []string{"<script>"},
		},
		{
			name:       "analyze rejects GET",
			method:     http.MethodGet,
			path:       "/analyze",
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var req *http.Request
			if tc.form != nil {
				req = httptest.NewRequest(tc.method, tc.path, strings.NewReader(tc.form.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			} else {
				req = httptest.NewRequest(tc.method, tc.path, nil)
			}
			rec := httptest.NewRecorder()
			srv.ServeHTTP(rec, req)

			if rec.Code != tc.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tc.wantStatus)
			}
			body := rec.Body.String()
			for _, s := range tc.wantContain {
				if !strings.Contains(body, s) {
					t.Errorf("body missing %q\n%s", s, body)
				}
			}
			for _, s := range tc.wantAbsent {
				if strings.Contains(body, s) {
					t.Errorf("body should not contain %q\n%s", s, body)
				}
			}
		})
	}
}

func TestAnalyze(t *testing.T) {
	a := analyze("One two three. Four five? Six")
	if a.Words != 6 || a.Sentences != 3 || a.ReadingMinutes != 1 {
		t.Errorf("got %+v", a)
	}
	if a.TopWords[0].Percent != 100 {
		t.Errorf("top word percent = %d, want 100", a.TopWords[0].Percent)
	}
	if long := analyze(strings.Repeat("word ", 401)); long.ReadingMinutes != 3 {
		t.Errorf("reading minutes = %d, want 3", long.ReadingMinutes)
	}
}

func TestDetectPlatform(t *testing.T) {
	tests := []struct {
		name         string
		env          map[string]string
		wantPlatform string
		wantInstance string
	}{
		{"local", nil, "", ""},
		{"cloud run", map[string]string{"K_SERVICE": "go-htmx-hackday", "K_REVISION": "go-htmx-hackday-00007-abc"}, "Cloud Run", "go-htmx-hackday-00007-abc"},
		{"render", map[string]string{"RENDER": "true", "RENDER_SERVICE_NAME": "go-htmx"}, "Render", "go-htmx"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, k := range []string{"K_SERVICE", "K_REVISION", "RENDER", "RENDER_SERVICE_NAME"} {
				t.Setenv(k, tc.env[k])
			}
			platform, instance := detectPlatform()
			if platform != tc.wantPlatform || instance != tc.wantInstance {
				t.Errorf("got (%q, %q), want (%q, %q)", platform, instance, tc.wantPlatform, tc.wantInstance)
			}

			rec := httptest.NewRecorder()
			testServer(t).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))
			want := "Running locally"
			if tc.wantPlatform != "" {
				want = "Live on " + tc.wantPlatform
			}
			if !strings.Contains(rec.Body.String(), want) {
				t.Errorf("page missing %q", want)
			}
		})
	}
}
