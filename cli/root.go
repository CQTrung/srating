package cli

import (
	"srating/bootstrap"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	root := &cobra.Command{}
	app := bootstrap.NewApplication()
	root.AddCommand(server(app), migrate(app.DB))
	return root
}
