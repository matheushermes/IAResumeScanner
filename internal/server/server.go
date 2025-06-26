package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/matheushermes/IAResumeScanner/configs"
	"github.com/matheushermes/IAResumeScanner/internal/server/routes"
)

type Server struct {
	port 	string
	server 	*gin.Engine
}

func NewServer() Server {
	return Server{
		port: configs.API_PORT,
		server: gin.Default(),
	}
}

func (s *Server) RunServer() {
	corsConfig := cors.Config{
			AllowAllOrigins: true,
			AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:   []string{"Content-Length"},
	}
	s.server.Use(cors.New(corsConfig))

	router := routes.ConfigRouter(s.server)

	//Configurando timeout;
	srv := &http.Server{
		Addr:         ":" + s.port,
		Handler:      router,
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 10 * time.Minute,
		IdleTimeout:  10 * time.Minute,
	}

	fmt.Printf("Server run on port %s", s.port)
	log.Fatal(srv.ListenAndServe())
}