package main

import (
	"fmt"
	"log"
	"os"

	"github.com/isjhar/iet/pkg"

	"github.com/akamensky/argparse"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Create new parser object
	parser := argparse.NewParser("generate-script", "script generation")
	// Create string flag
	methodPointer := parser.String("m", "method", &argparse.Options{Help: "Method get, create, update, delete, or all, get is default."})
	structName := parser.String("n", "name", &argparse.Options{Help: "Model name. ex User"})
	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		return
	}
	// Finally print the collected string
	method := *methodPointer
	switch method {
	case "get":
		scriptGenerator := pkg.ScriptGenerator{}
		err := scriptGenerator.GenerateGet(*structName)
		if err != nil {
			log.Panic(err)
		}
	default:
		log.Print("no action")
	}
}
