package pluginauth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {

	ctx := context.Background()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// Ory session/whoami API
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rsp := `{
			"identity": {
				"id": "9f425a8d-7efc-4768-8f23-7647a74fdf13",
				"traits": {
					"email": "foo@ory.sh"
				},
				"metadata_public": {
					"tenantId": "4248d7ff-ef45-43f6-8f07-10cff998aadf",
					"permissions": "read:users,write:users"
				}
			}
		}`
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(rsp))
		w.WriteHeader(http.StatusOK)
	}))

	cfg := CreateConfig()
	cfg.Host = srv.URL
	cfg.Headers = map[string]string{
		"User":        "TiccTech-User",
		"Tenant":      "TiccTech-Tenant",
		"Permissions": "Permissions",
	}

	handler, err := New(ctx, next, cfg, "auth-plugin")
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rec, req)

	assertHeader(t, req, "TiccTech-User", "9f425a8d-7efc-4768-8f23-7647a74fdf13")
	assertHeader(t, req, "TiccTech-Tenant", "4248d7ff-ef45-43f6-8f07-10cff998aadf")
	assertHeader(t, req, "Permissions", "read:users,write:users")
}

func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()
	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}
