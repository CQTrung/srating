package main

import (
	"srating/cli"
	"srating/utils"
)

// @securityDefinitions.basic BearerAuth
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	if err := utils.LoadConfig("."); err != nil {
		utils.LogFatal(err, "Error loading config")
	}
	rootCmd := cli.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		utils.LogFatal(err, "Error executing root command")
	}
}
