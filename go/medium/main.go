package main

import (
	"encoding/json"
	"fmt"
	"log"

	yini "github.com/YINI-lang/yini-parser-go"
)

const configFile = "settings.yini"

func main() {
	// Empty options select the parser's default lenient mode.
	result, err := yini.ParseFile(configFile, yini.Options{})
	if err != nil {
		log.Fatalf("parse %s: %v", configFile, err)
	}

	// Values converts sections and inline objects to maps, and arrays to slices.
	config := result.Document.Values()
	settings, ok := config["Settings"].(map[string]any)
	if !ok {
		log.Fatal("Settings section is missing")
	}
	network, ok := settings["Network"].(map[string]any)
	if !ok {
		log.Fatal("Network section is missing")
	}
	capabilities, ok := settings["Capabilities"].(map[string]any)
	if !ok {
		log.Fatal("Capabilities section is missing")
	}
	pools, ok := capabilities["pools"].(map[string]any)
	if !ok {
		log.Fatal("Capabilities.pools object is missing")
	}
	origins, ok := network["allowedOrigins"].([]any)
	if !ok {
		log.Fatal("Network.allowedOrigins array is missing")
	}
	// Backticks in YINI let section names contain spaces.
	database, ok := settings["DB Config"].(map[string]any)
	if !ok {
		log.Fatal("DB Config section is missing")
	}
	security, ok := database["Security"].(map[string]any)
	if !ok {
		log.Fatal("DB Config.Security section is missing")
	}

	fmt.Println("Application Settings")
	fmt.Println("--------------------")
	fmt.Printf("Service ID: %v\n", settings["serviceId"])
	fmt.Printf("Release: %v\n", settings["release"])
	fmt.Printf("Debug mode: %v\n", settings["debugMode"])

	fmt.Println("\nNetwork")
	fmt.Println("-------")
	fmt.Printf("Bind address: %v\n", network["bindAddress"])
	fmt.Printf("Bind port: %v\n", network["bindPort"])
	fmt.Println("Allowed origins:")
	for _, origin := range origins {
		fmt.Printf("- %v\n", origin)
	}

	fmt.Println("\nCapabilities")
	fmt.Println("------------")
	fmt.Printf("Search enabled: %v\n", capabilities["enableSearch"])
	fmt.Printf("Pool units: %v to %v\n", pools["min_units"], pools["max_units"])

	fmt.Println("\nDatabase")
	fmt.Println("--------")
	fmt.Printf("Host: %v\n", database["host"])
	fmt.Printf("SSL: %v\n", database["ssl"])
	fmt.Printf("Username: %v\n", security["username"])

	formatted, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatalf("format parsed settings: %v", err)
	}
	fmt.Printf("\nParsed settings as JSON:\n%s\n", formatted)

	// Strict mode is explicit and requires the /END marker in settings.yini.
	if _, err := yini.ParseFile(configFile, yini.Options{Mode: yini.ModeStrict}); err != nil {
		log.Fatalf("parse %s in strict mode: %v", configFile, err)
	}
	fmt.Println("\nStrict mode: loaded successfully")
}
