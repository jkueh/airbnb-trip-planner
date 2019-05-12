package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"gopkg.in/yaml.v2"
)

var config Config
var verbose bool
var debug bool

var gracefulStop = make(chan os.Signal, 1)
var gracefulStopRequested = false

func init() {

	// Set up the SIGTERM and SIGINT notifiers
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		log.Printf("Received signal %+v - Waiting for job to finish...\n", sig)
		gracefulStopRequested = true
	}()

	verboseString := strings.ToLower(os.Getenv("VERBOSE"))
	debugString := strings.ToLower(os.Getenv("DEBUG"))

	if debugString == "true" {
		verbose = true
		debug = true
	} else if verboseString == "true" {
		verbose = true
	}

	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("API Key not specified.")
	}
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = "my_trip.yml"
		log.Println("No CONFIG_FILE specified. Defaulting to:", configFile)
	} else {
		if debug {
			log.Println("Using config file: " + configFile)
		}
	}
	fileContents, err := ioutil.ReadFile(configFile)
	errExit(err)
	errExit(yaml.Unmarshal(fileContents, &config))
}
