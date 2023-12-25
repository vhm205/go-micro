package main

import (
	"context"

	"github.com/vhm205/go-micro/application"
)

func main() {
	app := application.Init()
	app.Run(context.TODO())
}
