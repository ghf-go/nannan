package main

import (
	"github.com/ghf-go/nannan/app"
	"os"
)

func main() {
	os.Setenv("app.web",":9081")
	app.Run()
}
