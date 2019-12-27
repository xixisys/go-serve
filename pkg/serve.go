package pkg

import (
	"context"
	"log"
	"net/http"
)

type StaticServe struct {
	srv  *http.Server
	Addr string
	Dir  string
}

func NewStaticServe(addr, dir string) *StaticServe {
	return &StaticServe{
		Addr: addr,
		Dir:  dir,
	}
}

func (s *StaticServe) Start() {
	s.srv = &http.Server{
		Addr:    s.Addr,
		Handler: http.FileServer(http.Dir(s.Dir)),
	}

	go func() {
		if err := s.srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()
}

func (s *StaticServe) End() {
	s.srv.Shutdown(context.TODO())
}
