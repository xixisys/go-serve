package main

import (
	"go-serve/pkg"
	"os"
	"os/signal"
)

func main() {
	s := pkg.NewStaticServe(":8000", "test/testdata")

	s.Start()

	c := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		s.End()
		done <- true
	}()

	<-done
}
