package config

import (
	"testing"
	"os"
)


// Config for testing purpose
type Config struct {
	DBHost      string `json:"RDS_HOSTNAME"`
	DBPort      string `json:"RDS_PORT"`
	DBName      string `json:"RDS_DB_NAME"`
	DBUser      string `json:"RDS_USERNAME"`
	DBPass      string `json:"RDS_PASSWORD"`
	DSN         string
}


func TestParse(t *testing.T) {
	os.Setenv("RDS_HOSTNAME", "localhost")
	os.Setenv("RDS_PORT", "8000")

	c := Config{}
	if err := Parse(&c); err != nil {
		t.Errorf(err.Error())
		return
	}

	if c.DBHost != "localhost" {
		t.Errorf("Not expected as %q", c.DBHost)
		return
	}

	if c.DBPort != "8000" {
		t.Errorf("Not expected as %q", c.DBPort)
		return
	}
}