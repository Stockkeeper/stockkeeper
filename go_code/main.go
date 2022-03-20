package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Stockkeeper/stockkeeper/config"
	"github.com/Stockkeeper/stockkeeper/server"
)

const (
	generateConfigFileCommandName        = "gen-config"
	generateConfigFileCommandDescription = "Generates a config file."
	startServerCommandName               = "start-server"
	startServerCommandDescription        = "Starts the API server."
	upgradeSchemaUpCommandName           = "db-up"
	upgradeSchemaUpCommandDescription    = "Upgrades the application database schema."
)

func main() {

	_ = flag.NewFlagSet(startServerCommandName, flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Printf(`
Usage: stockkeeper <command>

Commands:

%-16v %v
%-16v %v
%-16v %v

`,
			generateConfigFileCommandName,
			generateConfigFileCommandDescription,
			startServerCommandName,
			startServerCommandDescription,
			upgradeSchemaUpCommandName,
			upgradeSchemaUpCommandDescription,
		)
		os.Exit(1)
	}

	command := os.Args[1]
	flag.Parse()

	switch command {
	case startServerCommandName:
		startServer()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Usage: stockkeeper <command>")
		fmt.Println("Commands:")
		flag.PrintDefaults()
	}

}

func startServer() {
	c, err := config.ParseConfig("stockkeeper.yaml")
	if err != nil {
		panic(err)
	}

	listenAddr := fmt.Sprintf("%v:%v", c.ServerHost, c.ServerPort)
	fmt.Printf("Starting server on %v\n", listenAddr)
	if err := server.NewServer(listenAddr).ListenAndServe(); err != nil {
		panic(err)
	}
}
