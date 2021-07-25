package main

import (
	"fmt"
	"serapis/internal/pkg/metrics"
	"serapis/internal/pkg/server"
)

func main() {
	fmt.Println("  ______                           __        \n /   __/ ________________  ______ |__| _____\n \\___ \\_/ __ \\_  __ \\__  \\ \\____ \\|  |/  __/\n /     \\  ___/|  | \\// __ \\|  |_> >  |\\___ \\\n/____  /\\___  >__|  (____  /   __/|__/___   >\n     \\/     \\/           \\/|__|          \\/ \n")
	metrics.Start()
	server.Start()
}
