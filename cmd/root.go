package cmd

import (
	"os"

	"github.com/jcbritobr/james/dengine"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "James CLI",
		Version: "0.0.1",
		Long:    "James is a CLI for create gnome desktop launchers",
		Run: func(cmd *cobra.Command, args []string) {
			enBuilder := dengine.NewDesktopData()
			enBuilder.AddField(dengine.NameEntry, opName)
			enBuilder.AddField(dengine.CommentEntry, opComment)
			enBuilder.AddField(dengine.ExecEntry, opExec)
			enBuilder.AddField(dengine.IconEntry, opIcon)
			enBuilder.AddField(dengine.TerminalEntry, opTerminal)
			enBuilder.AddField(dengine.TypeEntry, opType)
			data := enBuilder.BuildFileData()
			file, err := os.Create(opFilename)

			if err != nil {
				panic(err)
			}

			file.WriteString(data)
		},
	}

	opName     string
	opExec     string
	opComment  string
	opType     string
	opTerminal string
	opIcon     string
	opFilename string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&opName, "name", "n", "default", "The --name -n sets Name entry on desktop file")
	rootCmd.PersistentFlags().StringVarP(&opExec, "exec", "e", "default", "The --exec -e sets Exec entry on desktop file")
	rootCmd.PersistentFlags().StringVarP(&opComment, "comment", "c", "default", "The --comment -c sets Comment entry on desktop file")
	rootCmd.PersistentFlags().StringVarP(&opType, "type", "t", "Application", "The --type -t sets Type entry on desktop file")
	rootCmd.PersistentFlags().StringVarP(&opTerminal, "terminal", "r", "false", "The --terminal -r sets Terminal entry on desktop file")
	rootCmd.PersistentFlags().StringVarP(&opIcon, "icon", "i", "default", "The --icon -i sets Icon entry on desktop file")
	rootCmd.PersistentFlags().StringVarP(&opFilename, "filename", "f", "default", "The --filename -f creates the <filename>.desktop")
}

func Execute() error {
	return rootCmd.Execute()
}
