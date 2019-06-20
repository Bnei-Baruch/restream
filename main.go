package main

import (
	"flag"
	"sync"
)

const (
	EXEC_PATH = "EXEC_PATH"
	GET_CMD   = "GET_CMD"
	PUT_CMD   = "PUT_CMD"
)

var (
	// http port
	port    = flag.String("p", "8081", "-p=8081")
	db_path = "db.json"
	mu      sync.Mutex
)

func main() {
	flag.Parse()
	a := App{}
	a.Initialize()
	a.Run(":" + *port)
}
