# b-r-o-w-s-e

A default browser for macOS that uses a config file + regular expressions to
determine which browser to open URL's in.

## Installing

Running `script/install` via the CLI will build and install b-r-o-w-s-e.
Alternatively you can follow the steps of `script/install` manually.

## Configuration

After installing b-r-o-w-s-e you need to set it as the default browser in System
Preferences -> General.

You also want to add some configuration. Without it, b-r-o-w-s-e won't work.
Here's a sample config file to build on:

```json
{
  "defaultBrowser": "Safari",
  "browsers": [
    { "name": "Safari", "path": "/Applications/Safari.app" },
    { "name": "Chrome", "path": "/Applications/Google Chrome.app" },
    { "name": "Firefox", "path": "/Applications/Firefox.app" }
  ],
  "matchers": [
    { "regexp": "https?.*?github.com", "browser": "Chrome" }
  ]
}
```

### Defining Browsers

Browsers go under the `browsers` array in the configuration file and require a
`name` and `path` property. The `name` property is how rules know
what browser to target and `path` is where the application lives.

eg: `{ "name": "Firefox", "path": "/Applications/Firefox.app" }`

The configuration also supports a `defaultBrowser` property which is the default
browser b-r-o-w-s-e opens when no matchers match the url to open.

### Defining Matchers

Matchers are the bread and butter of `b-r-o-w-s-e`. In the `matchers` array of
the config, you define objects with a `regexp` that is run against the opened
URL and if it matches it opens the browser defined by its `browser` property.
