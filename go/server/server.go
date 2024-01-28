package server

import (
	"log"

	"github.com/labstack/echo/v4"
)

type StaticFile struct {
	Prefix   string
	Filepath string
}

type Server struct {
	App  *echo.Echo
	addr string
}

func NewServer(app *echo.Echo, addr string) *Server {
	return &Server{
		App:  app,
		addr: addr,
	}
}

func (s *Server) Start() {
	go log.Println("Application running on", s.addr)
	s.App.Logger.Fatal(s.App.Start(s.addr))
}

func (s *Server) ServeStaticFiles(files []StaticFile) {
	for _, file := range files {
		s.App.File(file.Prefix, file.Filepath)
	}
}

func (s *Server) MakeHTTPHandler(path string, method string, handler echo.HandlerFunc) {
	switch method {
	case "GET":
		s.App.GET(path, handler)
	case "POST":
		s.App.POST(path, handler)
	case "PUT":
		s.App.PUT(path, handler)
	case "PATCH":
		s.App.PATCH(path, handler)
	case "DELETE":
		s.App.DELETE(path, handler)
	}
}
