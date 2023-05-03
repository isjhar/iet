package main

import (
	"isjhar/template/echo-golang/data/repositories"
	"log"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	err := repositories.Migrate()
	if err != nil {
		log.Panic(err)
	}
	if err != nil {
		log.Panic(err)
	}
}
