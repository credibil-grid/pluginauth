package pluginauth

import (
	// "bytes"
	"context"
	"net/http"
	"strings"

	"github.com/ory/client-go"
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
	headers map[string]string
	ory     *client.APIClient
	next    http.Handler
	name    string
}

// New created a new Auth plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {

	conf := client.NewConfiguration()
	conf.Servers = client.ServerConfigurations{{URL: config.Address}}

	return &Auth{
		headers: config.ResponseHeaders,
		ory:     client.NewAPIClient(conf),
		next:    next,
		name:    name,
	}, nil
}

func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// get session token/cookie
	token := ""
	if h := r.Header.Get("Authorization"); strings.HasPrefix(strings.ToLower(h), "bearer ") {
		token = h[7:]
	}
	cookies := r.Header.Values("Cookie")

	// call Ory API
	session, _, err := a.ory.FrontendApi.ToSession(context.Background()).
		XSessionToken(token).
		Cookie(strings.Join(cookies, "; ")).
		Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set response headers
	r.Header.Set(a.headers["User"], session.Identity.Id)
	r.Header.Set(a.headers["Tenant"], session.Identity.MetadataPublic["tenantId"].(string))
	r.Header.Set(a.headers["Permissions"], session.Identity.MetadataPublic["permissions"].(string))

	a.next.ServeHTTP(w, r)
}
