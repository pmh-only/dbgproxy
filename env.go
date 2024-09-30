package main

import (
	"fmt"
	"log"
	"os"
)

var LISTEN_HOST = tryEnv("LISTEN_HOST", "0.0.0.0")
var LISTEN_PORT = tryEnv("LISTEN_PORT", "3000")

var TARGET_HOST = tryEnv("TARGET_HOST", "127.0.0.1")
var TARGET_PORT = getEnv("TARGET_PORT")

var LISTEN_ADDRESS = fmt.Sprint(LISTEN_HOST, ":", LISTEN_PORT)
var TARGET_ADDRESS = fmt.Sprint(TARGET_HOST, ":", TARGET_PORT)

// ---

func tryEnv(envName string, defaults string) string {
	envValue, ok := os.LookupEnv(envName)
	if !ok {
		return defaults
	}

	return envValue
}

func getEnv(envName string) string {
	envValue, ok := os.LookupEnv(envName)
	if !ok {
		log.Fatalf(
			"Important environment variable: \"%s\" not provided.\n",
			envName)
	}

	return envValue
}
