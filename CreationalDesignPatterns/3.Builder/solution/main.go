package main

import (
	"design-patterns/CreationalDesignPatterns/3.Builder/solution/internal"
)

func main() {
	director := internal.NewDirector()
	builder := internal.NewBuilder()

	service := director.BuildService(builder)

	service.DoBusiness()
}
