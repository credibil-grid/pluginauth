[![Build Status](https://github.com/credibil/pluginauth/workflows/Release/badge.svg?branch=main)](https://github.com/credibil/pluginauth/actions)

# Traefik Ory auth plugin

A [Traefik](https://traefik.io) middleware plugin is just a [Go package](https://golang.org/ref/spec#Packages) that provides an `http.Handler` to perform specific processing of requests and responses.

Rather than being pre-compiled and linked, however, plugins are executed on the fly by [Yaegi](https://github.com/traefik/yaegi), an embedded Go interpreter.

## Development

Traefik uses `.traefik.yml` to check the integrity of the plugin and catch errors on startup. If an error occurs during loading, the plugin is disabled.

Once loaded, the plugin behaves like any other compiled middleware.

All dependencies need to be [vendored](https://golang.org/ref/mod#vendoring), with the vendored packages should be included in the repository. ([Go modules](https://blog.golang.org/using-go-modules) are not supported.)

## Installation

### Static Configuration

For a plugin to be active for a given Traefik instance, it must be declared in Traefik's static configuration.

```yaml
# Static configuration

experimental:
  plugins:
    oryauth:
      moduleName: github.com/credibil/pluginauth
      version: v0.0.15
```

### Dynamic Configuration

In the `http.middlewares` section:

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

## Defining a Plugin

A plugin package must define the following exported Go objects:

- A type `type Config struct { ... }`. The struct fields are arbitrary.
- A function `func CreateConfig() *Config`.
- A function `func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error)`.

## Logging

Currently, the only way to send logs to Traefik is to use `os.Stdout.WriteString("...")` or `os.Stderr.WriteString("...")`.

## Plugins Catalog

Every 30 minutes, the Plugins Catalog online service polls Github to find plugins and add them to its catalog.

### Prerequisites

To be recognized by Plugins Catalog, your repository must meet the following criteria:

- The `traefik-plugin` topic must be set.
- The `.traefik.yml` manifest must exist, and be filled with valid contents.

If your repository fails to meet either of these prerequisites, Plugins Catalog will not see it.

### Manifest

A manifest is also mandatory, and is named `.traefik.yml` and stored at the root of the project.

This YAML file provides Plugins Catalog with information about your plugin, such as a description, a full name, and so on.

Properties include:

- `displayName` (required): The name of your plugin as displayed in the Plugins Catalog web UI.
- `type` (required): For now, `middleware` is the only type available.
- `import` (required): The import path of your plugin.
- `summary` (required): A brief description of what your plugin is doing.
- `testData` (required): Configuration data for your plugin. This is mandatory, and Plugins Catalog will try to execute the plugin with the data you provide as part of its startup validity tests.
- `iconPath` (optional): A local path in the repository to the icon of the project.
- `bannerPath` (optional): A local path in the repository to the image that will be used when you will share your plugin page in social medias.

### Tags and Dependencies

Plugins Catalog needs the plugin to be versioned with a git tag.

If something goes wrong with the integration of your plugin, Plugins Catalog will create an issue inside your Github repository and will stop trying to add your repo until you close the issue.
