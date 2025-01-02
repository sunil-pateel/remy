package main

import (
    "fmt"
	"log/slog"
	"github.com/sunil-pateel/personal-website/internal/server"
)

func main() {
    const PORT int = 5001
    server := server.NewServer(PORT)

    domain := fmt.Sprintf("http://localhost:%d/", PORT)
    slog.Info("Starting server", "domain", domain)

    server.ListenAndServe()
    
}
