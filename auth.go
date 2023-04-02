package pluginauth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// Config holds the plugin configuration.
type Config struct {
	Host    string            `json:"host,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Headers: make(map[string]string),
	}
}

// Auth a Auth plugin.
type Auth struct {
	host    string
	headers map[string]string
	next    http.Handler
	name    string
}

// New created a new Auth plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Auth{
		host:    config.Host,
		headers: config.Headers,
		next:    next,
		name:    name,
	}, nil
}

func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// call Ory whoami API
	url := fmt.Sprintf("%s/sessions/whoami", a.host)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// token or cookie?
	if auth := r.Header.Get("Authorization"); strings.HasPrefix(strings.ToLower(auth), "bearer ") {
		req.Header.Set("X-Session-Token", auth[7:])
	} else {
		req.Header.Set("Cookie", r.Header.Get("Cookie"))
	}
	// req.Header.Set("Cache-Control", "max-age=60")

	start := time.Now()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// debug latency
	end := time.Since(start).Milliseconds()
	os.Stderr.WriteString(fmt.Sprintf("latency: %dms ", end))

	// If the response is 200, set a user ID header. If the response is not 200,
	// the user is not authenticated. This is OK since some of our services are
	// public.
	if res.StatusCode == http.StatusOK {

		var session struct {
			Identity struct {
				Id     string `json:"id"`
				Active bool   `json:"active"`
			} `json:"identity"`
		}

		defer res.Body.Close()
		if err := json.NewDecoder(res.Body).Decode(&session); err != nil {
			os.Stderr.WriteString(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// set response headers
		r.Header.Set(a.headers["credibil-user"], session.Identity.Id)
	}

	a.next.ServeHTTP(w, r)
}
