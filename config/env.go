package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment int

const (
	Production Environment = iota
	Development
	Debug
)

// Type Vars
//
// Environment variables
type Vars struct {
	ENVIRONMENT Environment
	SECRET      string
	HOST        string
}

func (v Vars) Validate() {
	switch {
	case v.SECRET == "":
		log.Fatal("missing env secret")
	case v.HOST == "":
		log.Fatal("missing env host")
	}
}

// Env() returns Vars struct of environment variables
func Env() Vars {
	var v Vars
	// Load if not a test. This isn't required during testing.
	if flag.Lookup("test.v") == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
		v = Vars{
			SECRET: os.Getenv("SECRET"),
		}
		switch os.Getenv("ENVIRONMENT") {
		case "production":
			v.ENVIRONMENT = Production
		case "development":
			v.ENVIRONMENT = Development
		case "debug":
			v.ENVIRONMENT = Debug
		default:
			log.Fatal("invalid environment")
		}

		switch v.ENVIRONMENT {
		case Production:
			v.HOST = os.Getenv("PRODUCTION_HOST")
		case Development:
			v.HOST = os.Getenv("DEVELOPMENT_HOST")
		case Debug:
			v.HOST = os.Getenv("DEBUG_HOST")
		}
		v.Validate()
	}

	return v
}
