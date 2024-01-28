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
	app  *echo.Echo
	addr string
}

func NewServer(app *echo.Echo, addr string) *Server {
	return &Server{
		app:  app,
		addr: addr,
	}
}

func (s *Server) Start() {
	go log.Println("Application running on", s.addr)
	s.app.Logger.Fatal(s.app.Start(s.addr))
}

func (s *Server) ServeStaticFiles(files []StaticFile) {
	for _, file := range files {
		s.app.File(file.Prefix, file.Filepath)
	}
}

func (s *Server) MakeHTTPHandler(path string, method string, handler echo.HandlerFunc) {
	switch method {
	case "GET":
		s.app.GET(path, handler)
	case "POST":
		s.app.POST(path, handler)
	case "PUT":
		s.app.PUT(path, handler)
	case "PATCH":
		s.app.PATCH(path, handler)
	case "DELETE":
		s.app.DELETE(path, handler)
	}
}
