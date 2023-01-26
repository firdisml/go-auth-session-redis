package main

import (
	"goth/model"
	"goth/routes"
)

func main() {
	model.Setup()
	routes.Setup()
}
