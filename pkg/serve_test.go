package pkg

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewStaticServe(t *testing.T) {
	t.Skip()
}

func TestStaticServe_End(t *testing.T) {
	t.Skip()
}

func TestStaticServe_Start(t *testing.T) {
	as := assert.New(t)

	s := NewStaticServe(":8000", "../test/testdata")
	defer s.End()
	s.Start()

	resp, err := http.Get("http://localhost:8000/file1.html")
	if err != nil {
		t.Fatal(err)
	}

	as.Equal(200, resp.StatusCode)

	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)
	as.Equal("file1", buf.String())
}
