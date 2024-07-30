package main

import (
	"context"
	"deliver/config"
	"deliver/internal/constants"
	"deliver/pkg/setup"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

const dir = "cmd/migration/schema"

var (
	flags          = flag.NewFlagSet("migrate", flag.ExitOnError)
	fieldMigration = logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryMigration}
)

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}

	db, err := setup.SetupPostgresConnection(config.GetConfig())
	if err != nil {
		logrus.WithFields(fieldMigration).Panic(err.Error())
	}
	defer db.Close()

	command := args[0]
	switch command {
	case "up", "down", "redo", "status":
		err = goose.RunContext(context.Background(), command, db.DB, dir, args...)
	default:
		err = goose.RunContext(context.Background(), "help", db.DB, dir, args...)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Println("Usage: make [OPTIONS] COMMAND")
	fmt.Println("Options:")
	fmt.Println("  -h, --help		Show this help message")
	fmt.Println("Commands:")
	fmt.Println("  up			Migrate the database to the most recent version available")
	fmt.Println("  down			Roll back the version by 1")
	fmt.Println("  redo			Roll back the most recently applied migration, then run it again")
	fmt.Println("  status		Print the status of all migrations")
}
