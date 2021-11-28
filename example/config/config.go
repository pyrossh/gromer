package config

import "os"

var DATABASE_URL string

func init() {
	DATABASE_URL = os.Getenv("DATABASE_URL")
}
