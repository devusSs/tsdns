package main

import (
	"encoding/json"
	"fmt"
)

var (
	version string
	commit  string
	date    string
)

type build struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func getBuild() *build {
	checkBuild()

	return &build{
		Version: version,
		Commit:  commit,
		Date:    date,
	}
}

func (b *build) String() string {
	return fmt.Sprintf("%+v", *b)
}

func (b *build) json() string {
	m, err := json.Marshal(b)
	if err != nil {
		return fmt.Sprintf("error marshaling build info: %v", err)
	}
	return string(m)
}

func checkBuild() {
	if version == "" {
		version = "development"
	}

	if commit == "" {
		commit = "unknown"
	}

	if date == "" {
		date = "unknown"
	}
}
