// Copyright 2020-2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

// A configuration type and helpers.

package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// A Config provides a way to load config from file and retrieve
// value from content.
type Config struct {
	content interface{}
}

// CONF is the global config instance.
var CONF = &Config{}

// Get retrieves a value from CONF.
func Get(path string) interface{} {
	return CONF.Get(path)
}

// GetString retrieves a string value from CONF.
func GetString(path string) string {
	v, ok := CONF.Get(path).(string)
	if !ok {
		return ""
	}
	return v
}

func GetBool(path string) bool {
	v, ok := CONF.Get(path).(bool)
	if !ok {
		return false
	}
	return v
}

// Load loads config content from supported file formats.
func Load(file string) {
	// TODO(william): implements load configuration from YAML file.
	CONF.LoadJSON(file)
}

// LoadJSON loads config content from json file.
func (cfg *Config) LoadJSON(file string) {
	// check config file
	file, _ = filepath.Abs(file)
	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.Printf("the configuration %s is not exist.\n", file)
		log.Fatalln(err)
	}

	// read config file
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("load the configuration %s failed.\n", file)
		log.Fatalln(err)
	}

	// parse config file
	err = json.Unmarshal(content, &cfg.content)
	if err != nil {
		log.Printf("parse the configuration %s failed.\n", file)
		log.Fatalln(err)
	}
}

// Get retrieves a value from the configuration. The key is reated
// a as period-sprarated path, with each path segment used as a map key.
func (cfg *Config) Get(path string) interface{} {
	c := cfg.content
	for _, p := range strings.Split(path, ".") {
		// handle for silce
		if strings.HasPrefix(p, "[") && strings.HasSuffix(p, "]") {
			// parse array index
			is := strings.TrimFunc(p, func(r rune) bool {
				return r == '[' || r == ']'
			})
			i, err := strconv.Atoi(is)
			if err != nil {
				return nil
			}

			// get value from array
			switch v := c.(type) {
			case []interface{}:
				for j, k := range v {
					if i == j {
						c = k
					}
				}
			}
			continue
		}

		// handle for map
		m, ok := c.(map[string]interface{})
		if !ok {
			return nil
		}
		c = m[p]
	}
	return c
}
