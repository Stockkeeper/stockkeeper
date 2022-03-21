package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Stockkeeper/stockkeeper/go/config"
	"github.com/Stockkeeper/stockkeeper/go/database"
	"github.com/Stockkeeper/stockkeeper/go/registry"
	"github.com/Stockkeeper/stockkeeper/go/server"
	"github.com/Stockkeeper/stockkeeper/go/storage"
)

const (
	startServerCommandName            = "start-server"
	startServerCommandDescription     = "Starts the API server."
	upgradeSchemaUpCommandName        = "db-up"
	upgradeSchemaUpCommandDescription = "Upgrades the application database schema."
)

func main() {

	_ = flag.NewFlagSet(startServerCommandName, flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Printf(`
Usage: stockkeeper <command>

Commands:

%-16v %v
%-16v %v

`,
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

	db := &database.FakeDatabase{}
	storage := &storage.FakeStorage{}

	registrySrv := registry.NewService(db, storage)

	s := server.NewServer(c, registrySrv)
	fmt.Printf("Starting server on %v\n", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
