package main

import (
	"flag"
	"sync"
)

var (
	// http port
	port = flag.String("p", "8081", "-p=8081")
	db_path = "db.json"
	mu sync.Mutex
)

func main() {
	flag.Parse()
	a := App{}
	a.Initialize()
	a.Run(":" + *port)
}
