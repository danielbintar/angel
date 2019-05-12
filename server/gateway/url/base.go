package url

import (
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	once sync.Once
)

type UrlConfig struct {
	Prefix string   `yaml:"prefix"`
	Routes []string `yaml:"routes"`
}

func Instance() map[string]string {
	urls := make(map[string]string)
	once.Do(func() {
		files, err := ioutil.ReadDir("./url")
		if err != nil { panic(err) }

		for _, file := range files {
			var config *UrlConfig
			if file.Name() == "base.go" { continue }

			yamlFile, err := ioutil.ReadFile("./url/" + file.Name())
			if err != nil { panic(err) }
			err = yaml.Unmarshal(yamlFile, &config)
			if err != nil { panic(err) }
			if len(config.Routes) == 0 { continue }

			key := file.Name()[:len(file.Name()) - 5]
			target := os.Getenv("URL_" + strings.ToUpper(key))

			for _, route := range config.Routes {
				urls[config.Prefix + route] = target
				if route == "/" {
					urls[config.Prefix] = target
				} else {
					urls[config.Prefix + route + "/"] = target
				}
			}
		}
	})

	return urls
}
