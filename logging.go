package main

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

// logRequests prints one line per request to the console and turns panics into 500s.
//
//	POST /analyze 200 2400B 412µs htmx=true
func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		defer func() {
			if p := recover(); p != nil {
				if p == http.ErrAbortHandler {
					panic(p)
				}
				log.Printf("panic: %v\n%s", p, debug.Stack())
				if !rec.wroteHeader {
					http.Error(rec, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}
			log.Printf("%s %s %d %dB %s htmx=%t",
				r.Method, r.URL.Path, rec.status, rec.bytes,
				time.Since(start).Round(time.Microsecond), r.Header.Get("HX-Request") == "true")
		}()

		next.ServeHTTP(rec, r)
	})
}

// statusRecorder captures the status code and response size for the request log.
type statusRecorder struct {
	http.ResponseWriter
	status      int
	bytes       int64
	wroteHeader bool
}

func (s *statusRecorder) WriteHeader(code int) {
	if !s.wroteHeader {
		s.status = code
		s.wroteHeader = true
	}
	s.ResponseWriter.WriteHeader(code)
}

func (s *statusRecorder) Write(b []byte) (int, error) {
	s.wroteHeader = true
	n, err := s.ResponseWriter.Write(b)
	s.bytes += int64(n)
	return n, err
}

// Unwrap lets http.ResponseController reach the underlying writer.
func (s *statusRecorder) Unwrap() http.ResponseWriter { return s.ResponseWriter }
