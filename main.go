package main

import (
	"code_structure/cmd"
	"code_structure/internals/adapters/menu"
	"code_structure/internals/application"
	"code_structure/internals/domain/selection"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"code_structure/internals/config"
)

func main() {
	app := cli.NewApp()
	app.Usage = "SOME STUFF"

	// Init envConfig
	envConfig, err := config.LoadEnv()
	if err != nil {
		logrus.WithError(err).Fatal("The envConfig could not be loaded.")
	}

	selectionCore := selection.NewCore()

	mealsAdapter := menuadapter.NewMenuAdapter(&http.Client{})

	appAPI := application.NewAPI(mealsAdapter, selectionCore)

	// Init command handlers
	baseCommand := cmd.NewBaseCommand(context.Background(), envConfig)

	httpCommand := cmd.NewHTTPCommand(baseCommand, appAPI)
	consumer := cmd.NewConsumerCommand(baseCommand)
	migrateCommand := cmd.NewMigrateCommand(baseCommand)

	app.Commands = []cli.Command{
		{
			Name:   "http",
			Usage:  "Start REST API service",
			Action: httpCommand.Run,
		},
		{
			Name:      "migrate",
			Aliases:   []string{"m"},
			Usage:     "Run database migrations to the specific version",
			Action:    migrateCommand.Run,
			ArgsUsage: "",
			Subcommands: []cli.Command{
				{
					Name:      "up",
					Aliases:   []string{"u"},
					Usage:     "Up the database migrations",
					Action:    migrateCommand.Up,
					ArgsUsage: "",
				},
				{
					Name:      "down",
					Aliases:   []string{"d"},
					Usage:     "Down the database migrations",
					Action:    migrateCommand.Down,
					ArgsUsage: "",
				},
			},
		},

		{
			Name:   "consumer",
			Usage:  "Run consumer",
			Action: consumer.Run,
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
