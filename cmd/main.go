package main

import (
	"github.com/matheushermes/IAResumeScanner/configs"
	"github.com/matheushermes/IAResumeScanner/internal/server"
)

func main() {
	configs.LoadingEnvironmentVariables()
	server := server.NewServer()
	server.RunServer()
}