package server

import (
    "fmt"
	"net/http"
	"time"
)

func NewServer(portNumber int) *http.Server {
    server := &http.Server {
        Addr: fmt.Sprintf(":%d", portNumber),
        ReadTimeout: time.Minute,
        ReadHeaderTimeout: time.Minute,
        WriteTimeout: time.Minute,
    }

    server.Handler = MakeRoutesHandler()

    return server
}
