package main

import "testing"

var config Config = ParseConfig(`
{
  "defaultBrowser": "Safari",
  "browsers": [
	{ "name": "Safari", "path": "/Applications/Safari.app" },
	{ "name": "Chrome", "path": "/Applications/Google Chrome.app" }
  ],
  "matchers": [
	{ "regexp": "https?.*?github.com", "browser": "Chrome" }
  ]
}`)

func TestGetDefaultBrowser(t *testing.T) {
	if config.GetDefaultBrowser().Path != "/Applications/Safari.app" {
		t.Errorf("default browser path was not `/Applications/Safari.app`, got: %s", config.GetDefaultBrowser().Path)
	}
}

func TestGetBrowserForUrl(t *testing.T) {
	githubBrowser := config.GetBrowserForUrl("https://github.com/")
	defaultBrowser := config.GetBrowserForUrl("https://reddit.com/")

	if githubBrowser.Path != "/Applications/Google Chrome.app" {
		t.Errorf("expected github browser to be `/Applications/Google Chrome.app`, got: %s", githubBrowser.Path)
	}

	if defaultBrowser.Path != "/Applications/Safari.app" {
		t.Errorf("expected safari browser to be `/Applications/Google Chrome.app`, got: %s", githubBrowser.Path)
	}
}
