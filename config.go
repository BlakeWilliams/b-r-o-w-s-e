package main

import (
	"encoding/json"
	"regexp"
)

type Config struct {
	Browsers           []Browser `json:"browsers"`
	DefaultBrowserName string    `json:"defaultBrowser"`
	Matchers           []Matcher `json:"matchers"`
}

type Browser struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Matcher struct {
	Regexp      string `json:"regexp"`
	BrowserName string `json:"browser"`
}

func ParseConfig(configJson string) Config {
	var config Config

	err := json.Unmarshal([]byte(configJson), &config)
	if err != nil {
		ShowError(
			"Couldn't parse config file",
			err.Error(),
		)
		panic(err)
	}

	return config
}

func (config Config) GetBrowser(name string) Browser {
	for _, browser := range config.Browsers {
		if browser.Name == name {
			return browser
		}
	}

	ShowError(
		"Couldn't find browser",
		"Couldn't find browser with name: "+name,
	)
	panic("couldn't find browser with name " + name)
}

func (config Config) GetDefaultBrowser() Browser {
	return config.GetBrowser(config.DefaultBrowserName)
}

func (config Config) GetBrowserForUrl(url string) Browser {
	for _, matcher := range config.Matchers {
		regex, err := regexp.Compile(matcher.Regexp)

		if err != nil {
			ShowError(
				"Couldn't compile regex",
				matcher.Regexp,
			)
			panic(err)
		}

		if regex.MatchString(url) {
			return config.GetBrowser(matcher.BrowserName)
		}
	}

	return config.GetDefaultBrowser()
}
