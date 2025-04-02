package main

import (
	"log"

	"github.com/isjhar/iet/data/repositories"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	err := repositories.MigrateSeed()
	if err != nil {
		log.Panic(err)
	}
	if err != nil {
		log.Panic(err)
	}
}
