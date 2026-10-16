package main

import (
	"fmt"
	"log"

	yini "github.com/YINI-lang/yini-parser-go"
)

func main() {
	// Empty options select the parser's default lenient mode.
	result, err := yini.ParseFile("config.yini", yini.Options{})
	if err != nil {
		log.Fatal(err)
	}

	// Values converts the parsed document to ordinary Go maps and slices.
	config := result.Document.Values()
	app, ok := config["App"].(map[string]any)
	if !ok {
		log.Fatal("App section is missing")
	}
	server, ok := app["Server"].(map[string]any)
	if !ok {
		log.Fatal("Server section is missing")
	}

	fmt.Println("Application")
	fmt.Println("-----------")
	fmt.Printf("Name: %v\n", app["name"])
	fmt.Printf("Version: %v\n", app["version"])
	fmt.Printf("Debug: %v\n", app["debug"])
	fmt.Printf("Features: %v\n", app["features"])

	fmt.Println()
	fmt.Println("Server")
	fmt.Println("------")
	fmt.Printf("Host: %v\n", server["host"])
	fmt.Printf("Port: %v\n", server["port"])
	fmt.Printf("Use TLS: %v\n", server["useTLS"])
}
