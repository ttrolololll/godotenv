// A simple package to read .env file and get the key-values
// for usage in application.

package dotenv

import (
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

type Dotenv struct {
	dotenvPath string
	varMap map[string]string
}

var instance *Dotenv
var once sync.Once

// Implement singleton pattern
// Good read: http://marcio.io/2015/07/singleton-pattern-in-go/
func GetInstance() *Dotenv {
	once.Do(func() {
		instance = &Dotenv{}
		instance.varMap = make(map[string]string)
		instance.Load()
	})

	return instance
}

// Gets the current .env file path
func (d Dotenv) GetPath() string {
	return d.dotenvPath
}

// Returns a value from .env
func (d Dotenv) Env(key string, backupVal ...string) string {
	value, ok := d.varMap[key]

	if ok {
		return value
	}

	backup := ""

	if len(backupVal) > 0 {
		backup = backupVal[0]
	}

	return backup
}

// Read .env file and load vars into map
func (d *Dotenv) Load(path ...string) bool {
	d.dotenvPath = ".env"

	if len(path) > 0 {
		d.dotenvPath = path[0]
	}

	content, err := ioutil.ReadFile(d.dotenvPath)
	if err != nil {
		log.Fatal(err)
		return false
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		keyVal := strings.Split(line, "=")

		if len(keyVal) == 2 {
			d.varMap[keyVal[0]] = keyVal[1]
		}
	}

	return true
}

// [TODO] To check for key-value formatting in .env file
// [TODO] Redesign it to be more resilient to errors
