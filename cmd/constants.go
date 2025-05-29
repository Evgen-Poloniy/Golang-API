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
