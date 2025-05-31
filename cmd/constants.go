package main

import (
	"time"
)

const (
	envPath string = "./config/.env"
)

const (
	maxHeaderBytes = 1 << 20
	writeTimeout   = 10 * time.Second
	readTimeout    = 10 * time.Second
)

const (
	DEBUG            string = "debug"
	DEBUG_WITHOUT_DB string = "debug_without_string"
	PRODUCTION       string = "production"
)
