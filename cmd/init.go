package cmd

import (
	"github.com/sottey/redo.vc/redovc"
	"github.com/spf13/cobra"
)

func init() {
	var (
		initCmdDesc     = "Initializes a new todo list in the current directory"
		initCmdLongDesc = `Initializes a new todo list in the current directory.

This will create a .todos.json in the directory you're in.  You can then start adding todos to it.`
	)

	var initCmd = &cobra.Command{
		Use:   "init",
		Long:  initCmdLongDesc,
		Short: initCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			redovc.NewApp().InitializeRepo()
		},
	}

	rootCmd.AddCommand(initCmd)
}
