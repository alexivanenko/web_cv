package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"path"
	"runtime"
	"strings"

	"github.com/go-ini/ini"
)

const VERSION = "0.0.1"

var config *ini.File
var stdLogger *log.Logger
var rootDir string

func getValue(key string) *ini.Key {
	return config.Section("").Key(key)
}

func String(key string) string {
	return getValue(key).String()
}

func Is(key string) bool {
	if val, err := getValue(key).Bool(); err != nil {
		return false
	} else {
		return val
	}
}

func Log(msg string) {
	stdLogger.Printf("[LOG] %v | %s\n", time.Now().Format("2016/01/06 - 16:07:08"), msg)
}

func GetVersion() string {
	return VERSION
}

func GetRootDir() string {
	if rootDir == "" {
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			panic("No caller information")
		}

		rootDir = path.Dir(filename)
		rootDir = strings.Replace(rootDir, "/config", "", 1)
	}

	return rootDir
}

func init() {
	stdLogger = log.New(os.Stdout, "", 0)
	Log(fmt.Sprintf("Initializing application ver %s", GetVersion()))

	var err error
	configPath := fmt.Sprintf("%s/config.ini", GetRootDir())

	Log(fmt.Sprintf("Loading base config from %s", configPath))

	if config, err = ini.Load(configPath); err != nil {
		panic(err)
	}
}
