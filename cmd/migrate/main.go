package main

import (
	"fmt"
	"log"
	"os"

	"github.com/isjhar/iet/internal/config"
	"github.com/isjhar/iet/internal/data/repositories"

	"github.com/akamensky/argparse"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Create new parser object
	parser := argparse.NewParser("migration", "handle migration database")
	// Create string flag
	modePointer := parser.String("m", "mode", &argparse.Options{Help: "Mode migrate, seed, or refresh, migrate is default."})
	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		return
	}
	config.LoadConfig()
	repositories.Connect()
	// Finally print the collected string
	mode := *modePointer
	switch mode {
	case "refresh":
		err = repositories.Refresh()
		if err != nil {
			log.Panic(err)
		}
	case "seed":
		err = repositories.MigrateSeed()
		if err != nil {
			log.Panic(err)
		}
	default:
		err = repositories.MigrateDatabase()
		if err != nil {
			log.Panic(err)
		}
	}

	// err := repositories.MigrateDatabase()
	// if err != nil {
	// 	log.Panic(err)
	// }
	// if err != nil {
	// 	log.Panic(err)
	// }
}
