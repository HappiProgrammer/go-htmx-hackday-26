package main

import (
	"context"
	"embed"
	"errors"
	"html/template"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sort"
	"strings"
	"syscall"
	"time"
	"unicode"
)

//go:embed templates/*
var templatesFS embed.FS

// version is set at build time with -ldflags "-X main.version=...".
var version = "dev"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.SetOutput(os.Stdout)
	tmpl := template.Must(template.ParseFS(templatesFS, "templates/*.html"))

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           logRequests(newServer(tmpl)),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	// Cloud Run sends SIGTERM before it stops an instance; finish in-flight requests.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("server starting on :%s (version=%s, %s, revision=%q)", port, version, runtime.Version(), os.Getenv("K_REVISION"))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server failed: %v", err)
		}
	}()

	<-ctx.Done()
	log.Print("shutdown signal received, draining connections")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}
	log.Print("server stopped")
}

// newServer wires the routes. Kept separate from main so tests can use it.
func newServer(tmpl *template.Template) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		// The page embeds the empty-state fragment so first paint matches later swaps.
		platform, instance := detectPlatform()
		render(w, tmpl, "index.html", page{
			Results:  analyze(""),
			Platform: platform,
			Instance: instance,
			Version:  version,
		})
	})
	mux.HandleFunc("POST /analyze", func(w http.ResponseWriter, r *http.Request) {
		handleHTMXFragment(w, r, tmpl)
	})
	return mux
}

// handleHTMXFragment returns only the #results fragment, never a full page.
func handleHTMXFragment(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
	start := time.Now()
	a := analyze(r.FormValue("text"))
	a.Elapsed = time.Since(start).Round(100 * time.Nanosecond)
	render(w, tmpl, "results", a)
}

func render(w http.ResponseWriter, tmpl *template.Template, name string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		log.Printf("render %s: %v", name, err)
		http.Error(w, "template error", http.StatusInternalServerError)
	}
}

// detectPlatform reads the environment variables each host sets on the container.
func detectPlatform() (platform, instance string) {
	switch {
	case os.Getenv("K_SERVICE") != "":
		return "Cloud Run", os.Getenv("K_REVISION")
	case os.Getenv("RENDER") != "":
		return "Render", os.Getenv("RENDER_SERVICE_NAME")
	}
	return "", ""
}

type page struct {
	Results  Analysis
	Platform string // "Cloud Run", "Render", or empty when running locally
	Instance string // revision or service name on that platform
	Version  string
}

type WordCount struct {
	Word    string
	Count   int
	Percent int // bar width relative to the most frequent word
}

type Analysis struct {
	Empty          bool
	Words          int
	Characters     int
	Sentences      int
	ReadingMinutes int
	TopWords       []WordCount
	Elapsed        time.Duration
}

// analyze computes simple text statistics. Reading time assumes 200 wpm.
func analyze(text string) Analysis {
	text = strings.TrimSpace(text)
	if text == "" {
		return Analysis{Empty: true}
	}

	words := strings.FieldsFunc(strings.ToLower(text), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '\''
	})

	sentences := 0
	inSentence := false
	for _, r := range text {
		switch {
		case r == '.' || r == '!' || r == '?':
			if inSentence {
				sentences++
			}
			inSentence = false
		case !unicode.IsSpace(r):
			inSentence = true
		}
	}
	if inSentence {
		sentences++
	}

	counts := map[string]int{}
	for _, w := range words {
		counts[w]++
	}
	top := make([]WordCount, 0, len(counts))
	for w, c := range counts {
		top = append(top, WordCount{Word: w, Count: c})
	}
	sort.Slice(top, func(i, j int) bool {
		if top[i].Count != top[j].Count {
			return top[i].Count > top[j].Count
		}
		return top[i].Word < top[j].Word
	})
	if len(top) > 5 {
		top = top[:5]
	}
	for i := range top {
		top[i].Percent = top[i].Count * 100 / top[0].Count
	}

	return Analysis{
		Words:          len(words),
		Characters:     len([]rune(text)),
		Sentences:      sentences,
		ReadingMinutes: int(math.Max(1, math.Ceil(float64(len(words))/200))),
		TopWords:       top,
	}
}
