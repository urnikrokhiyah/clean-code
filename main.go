package main

import (
	"cleancode/config"
	"cleancode/routes"
)

func main() {
	config.InitDb()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
