package main

import (
	"github.com/matheushermes/IAResumeScanner/internal/server"

	_ "github.com/matheushermes/IAResumeScanner/init"
)

func main() {
	server := server.NewServer()
	server.RunServer()
}