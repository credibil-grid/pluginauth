[![Build Status](https://github.com/credibil/pluginauth/workflows/Release/badge.svg?branch=main)](https://github.com/credibil/pluginauth/actions)

# Traefik Ory Auth Plugin

This [Traefik](https://traefik.io) middleware plugin validates requests are made by authenticated clients (have a valid Ory session).

The plugin calls the Ory `<host>/session/whoami` API and, if successful, will inject the user, tenant and permissions into the request headers.

## Usage

To use this plugin it must be defined in the Traefik static configuration and referenced in the `http.middlewares` section of a route's dynamic configuration.

```yaml
# Static configuration

experimental:
  plugins:
    oryauth:
      moduleName: github.com/credibil/pluginauth
      version: v0.0.15
```

```yaml
# Dynamic configuration

http:
  routers:
    app-1-default:
      rule: "PathPrefix(`/app1`)"
      middlewares:
        - check-auth
      service: app-1
  
  middlewares:
    check-auth:
      plugin:
        oryauth:
          host: "https://auth.staging.amlify.com"
          headers:
            User: "TiccTech-User"
            Tenant: "TiccTech-Tenant"
            Permissions: "Permissions"
  
  services:
    ...
```

## Development

A Traefik plugin is a simple http server implementing Traefik's plugin protocol by exporting:

- A type `type Config struct { ... }`
- A function `func CreateConfig() *Config`
- A function `func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error)`

### Manifest

Traefik uses the `.traefik.yml` manifest to provide Plugins Catalog with information about the plugin and to check the integrity of the plugin and catch errors on startup.

### Vendoring

Traefik does not support Go modules](https://blog.golang.org/using-go-modules) so all dependencies need to be [vendored](https://golang.org/ref/mod#vendoring).

### Logging

Currently, Traefik only supports logging using `os.Stdout.WriteString("...")` or `os.Stderr.WriteString("...")`.

### Repository Topic

In addition to containing a manifest, the Plugins Catalog requires a repository to have a topic set named `traefik-plugin`.

### Versioning

The Plugins Catalog requires the plugin be versioned with a git tag.

### Issues

If something goes wrong with the plugin, the Plugins Catalog will create a GitHub issue in the plugin repository and stops trying to add the repo. Closing the issue will allow the Plugins Catalog to try again.
