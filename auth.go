package pluginauth

import (
	// "bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

// Config the plugin configuration.
type Config struct {
	Address         string            `json:"address,omitempty"`
	RequestHeaders  []string          `json:"requestHeaders,omitempty"`
	ResponseHeaders map[string]string `json:"responseHeaders,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		ResponseHeaders: make(map[string]string),
	}
}

// Auth a Auth plugin.
type Auth struct {
	next       http.Handler
	reqHeaders []string
	rspHeaders map[string]string
	name       string
	template   *template.Template
}

// New created a new Auth plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.RequestHeaders) == 0 {
		return nil, fmt.Errorf("headers cannot be empty")
	}

	return &Auth{
		reqHeaders: config.RequestHeaders,
		rspHeaders: config.ResponseHeaders,
		next:       next,
		name:       name,
		template:   template.New("demo").Delims("[[", "]]"),
	}, nil
}

func (a *Auth) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	os.Stderr.WriteString(fmt.Sprintf("%v", a.rspHeaders))

	// for key, value := range a.headers {
	// 	tmpl, err := a.template.Parse(value)
	// 	if err != nil {
	// 		http.Error(rw, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	writer := &bytes.Buffer{}

	// 	err = tmpl.Execute(writer, req)
	// 	if err != nil {
	// 		http.Error(rw, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	req.Header.Set(key, writer.String())
	// }

	a.next.ServeHTTP(rw, req)
}
